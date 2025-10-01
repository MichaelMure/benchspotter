package engine

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	iofs "io/fs"
	"iter"
	"path/filepath"
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

		f, err := sources.Open(path)
		defer func() { _ = f.Close() }()

		if info.IsDir() || !strings.HasSuffix(path, "_test.go") {
			return nil
		}

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

			param := fn.Type.Params.List[0]
			if starExpr, ok := param.Type.(*ast.StarExpr); ok {
				if selectorExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
					if pkgIdent, ok := selectorExpr.X.(*ast.Ident); ok {
						if actualPackage, ok := importMap[pkgIdent.Name]; ok {
							if actualPackage == "testing" && selectorExpr.Sel.Name == "B" {
								res = append(res, BenchInfo{
									Name:    fn.Name.Name,
									Package: filepath.Dir(path),
								})
							}
						}
					}
				}
			}
			return true
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

func RunBenches(ctx context.Context, storage billy.Filesystem, id string, benches []BenchInfo) func() (*benchfmt.Result, error) {
	next, _ := iter.Pull2(func(yield func(*benchfmt.Result, error) bool) {
		out, err := storage.Create(filepath.Join(id, benchFilename))
		if err != nil {
			yield(nil, err)
			return
		}
		defer func() { _ = out.Close() }()
		w := benchfmt.NewWriter(out)

		for _, infos := range benches {
			cmd := execabs.CommandContext(ctx, "go", "test", "-bench",
				"^\\Q"+infos.Name+"\\E$", "-benchmem", "-run", "^$", ".")
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

			var res *benchfmt.Result
			for r.Scan() {
				line := r.Result()
				switch line := line.(type) {
				case *benchfmt.Result:
					res = line
					// 	// weird dance to set a non-internal config, the API is not meant for that
					// 	res.SetConfig("git-commit", "qsbdjkqsdnbkqqsd")
					// 	idx, _ := res.ConfigIndex("git-commit")
					// 	res.Config[idx].File = true // not internal, meaning it will print in the writer
				}

				err = w.Write(res)
				if err != nil {
					yield(nil, err)
					return
				}
			}

			if !yield(res, nil) {
				return
			}
		}
	})
	return func() (*benchfmt.Result, error) {
		res, err, _ := next()
		return res, err
	}
}
