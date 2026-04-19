package engine

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	iofs "io/fs"
	"iter"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"golang.org/x/perf/benchfmt"
	"golang.org/x/sys/execabs"
)

const benchFilename = "results.bench"

type BenchInfo struct {
	Name    string
	Package string
}

func (i BenchInfo) Regex() string {
	var res strings.Builder
	for i, segment := range strings.Split(i.Name, "/") {
		if i > 0 {
			res.WriteString(`/`)
		}
		res.WriteString(`^\Q`)
		res.WriteString(segment)
		res.WriteString(`\E$`)
	}
	return res.String()
}

func LocateBenchmarks(ctx context.Context, sources billy.Filesystem) ([]BenchInfo, error) {
	// TODO: this doesn't support dot import of "testing"

	var res []BenchInfo

	err := util.Walk(sources, ".", func(path string, info iofs.FileInfo, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, "_test.go") {
			return nil
		}

		f, err := sources.Open(path)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()

		fileAst, err := parser.ParseFile(token.NewFileSet(), "", f, parser.SkipObjectResolution)
		if err != nil {
			return err
		}

		importMap := buildImportMap(fileAst)

		ast.Inspect(fileAst, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok || fn.Recv != nil {
				return true
			}

			// Must start with "Benchmark"
			if !fn.Name.IsExported() || !strings.HasPrefix(fn.Name.Name, "Benchmark") {
				return true
			}

			// Must have one parameter: *testing.B
			if fn.Type.Params.NumFields() != 1 {
				return true
			}
			param1 := fn.Type.Params.List[0]
			if !isTestingBParam(param1.Type, importMap) {
				return true
			}

			// Now, hunt for sub-benchmarks.
			// First, we need the name of the *testing.B parameter.
			// If unnamed (ie, BenchmarkXX(*testing.B), we can stop as there won't be any sub-benchmarks.
			if len(param1.Names) == 0 || fn.Body == nil {
				return false
			}
			testingVarName := param1.Names[0].Name

			// Record the top-level benchmark, but only if there is no .Run() sub call.
			// Recursively find all sub-benchmarks.
			if !collectSubBenchmarks(&res, filepath.Dir(path), fn.Name.Name, testingVarName, fn.Body, importMap) {
				res = append(res, BenchInfo{
					Name:    fn.Name.Name,
					Package: filepath.Dir(path),
				})
			}

			// Stop traversing the AST.
			return false
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// buildImportMap creates a mapping from local package names to their actual import paths
func buildImportMap(file *ast.File) map[string]string {
	importMap := make(map[string]string)

	for _, imp := range file.Imports {
		if imp.Path == nil {
			continue
		}

		// Remove quotes from import path
		path := imp.Path.Value[1 : len(imp.Path.Value)-1]

		var localName string
		if imp.Name != nil {
			if imp.Name.Name == "_" {
				continue // skip blank imports
			}
			localName = imp.Name.Name
		} else {
			localName = filepath.Base(path)
		}

		importMap[localName] = path
	}

	return importMap
}

// isTestingBParam checks whether the given parameter type is *testing.B
func isTestingBParam(paramType ast.Expr, importMap map[string]string) bool {
	starExpr, ok := paramType.(*ast.StarExpr)
	if !ok {
		return false
	}
	selectorExpr, ok := starExpr.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	pkgIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}
	actualPackage, ok := importMap[pkgIdent.Name]
	if !ok {
		return false
	}
	return actualPackage == "testing" && selectorExpr.Sel.Name == "B"
}

// collectSubBenchmarks recursively finds all sub-benchmarks in the given function body.
// It returns true if a sub-benchmark was found, false otherwise.
func collectSubBenchmarks(dst *[]BenchInfo, pkgDir, prefix, testingVarName string, node ast.Node, importMap map[string]string) (found bool) {
	ast.Inspect(node, func(n ast.Node) bool {
		// Search for b.Run(...)
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		xIdent, ok := sel.X.(*ast.Ident)
		if !ok || xIdent.Name != testingVarName {
			return true
		}
		if sel.Sel == nil || sel.Sel.Name != "Run" {
			return true
		}

		// Expect 2 parameters: b.Run("name", func(b *testing.B) { ... })
		if len(call.Args) < 2 {
			return true
		}

		// Expect the first parameter to be a string literal: the name of the sub-benchmark.
		nameLit, ok := call.Args[0].(*ast.BasicLit)
		if !ok || nameLit.Kind != token.STRING {
			return true // dynamic names can't be determined here
		}
		subName, err := strconv.Unquote(nameLit.Value)
		if err != nil || subName == "" {
			return true
		}

		// Expect the second parameter to be a function literal: the sub-benchmark body.
		fnLit, ok := call.Args[1].(*ast.FuncLit)
		if !ok || fnLit.Type == nil || fnLit.Type.Params == nil || fnLit.Type.Params.NumFields() != 1 {
			return true
		}
		subParam := fnLit.Type.Params.List[0]
		if !isTestingBParam(subParam.Type, importMap) {
			return true
		}

		// Record the sub-benchmark, but only if there is no .Run() sub call.
		found = true
		fullName := prefix + "/" + subName

		// Recurse into the sub-benchmark body to find nested b.Run calls.
		if len(subParam.Names) == 0 || fnLit.Body == nil {
			return false
		}
		subTestingVarName := subParam.Names[0].Name

		if !collectSubBenchmarks(dst, pkgDir, fullName, subTestingVarName, fnLit.Body, importMap) {
			*dst = append(*dst, BenchInfo{
				Name:    fullName,
				Package: pkgDir,
			})
		}

		// Stop traversing the AST.
		return false
	})
	return
}

func RunBenches(ctx context.Context, storage billy.Filesystem, id string, benches []BenchInfo, count int) func() (*benchfmt.Result, error) {
	next, _ := iter.Pull2(func(yield func(*benchfmt.Result, error) bool) {
		out, err := storage.Create(filepath.Join(sessionDir, id, benchFilename))
		if err != nil {
			yield(nil, err)
			return
		}
		defer func() { _ = out.Close() }()
		w := benchfmt.NewWriter(out)

		for _, infos := range benches {
			cmd := execabs.CommandContext(ctx, "go", "test",
				"-bench", infos.Regex(),
				"-benchmem", "-count", strconv.Itoa(count), "-run", "^$", ".")
			cmd.Dir = infos.Package

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				yield(nil, err)
				return
			}

			err = cmd.Start()
			if err != nil {
				yield(nil, err)
				return
			}

			r := benchfmt.NewReader(stdout, "")

			for r.Scan() {
				line := r.Result()

				err = w.Write(line)
				if err != nil {
					yield(nil, err)
					return
				}

				if line, ok := line.(*benchfmt.Result); ok {
					if !yield(line, nil) {
						return
					}
				}
			}
			if err := r.Err(); err != nil {
				yield(nil, err)
				return
			}
			if err := cmd.Wait(); err != nil {
				yield(nil, err)
				return
			}
		}
	})
	return func() (*benchfmt.Result, error) {
		res, err, _ := next()
		return res, err
	}
}
