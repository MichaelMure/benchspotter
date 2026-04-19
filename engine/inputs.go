package engine

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"strconv"
)

type InputInfo interface {
	Package() string
	Name() string
}

type BoolType struct {
	pkg  string
	name string
}

func (b BoolType) Package() string {
	return b.pkg
}

func (b BoolType) Name() string {
	return b.name
}

type IntType struct {
	pkg  string
	name string
	min  int
	max  int
}

func (i IntType) Package() string {
	return i.pkg
}

func (i IntType) Name() string {
	return i.name
}

func Inputs(rootDir string) (res []InputInfo, err error) {
	rootDir, err = filepath.Abs(rootDir)
	if err != nil {
		return nil, err
	}

	pkgs, err := packages(rootDir)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(n ast.Node) bool {
				call, ok := n.(*ast.CallExpr)
				if !ok {
					return true
				}
				sel, ok := call.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				selType := pkg.TypesInfo.Uses[sel.Sel]
				if selType == nil {
					return true
				}
				fn, ok := selType.(*types.Func)
				if !ok {
					return true
				}
				if fn.Pkg().Path() != "benchspotter/benchinput" {
					return true
				}
				sig := fn.Type().(*types.Signature)

				switch fn.Name() {
				case "Bool": // name string, default_ bool
					if sig.Params().Len() != 2 {
						return true
					}
					param1 := sig.Params().At(0)
					if param1.Type().String() != "string" {
						return true
					}
					param2 := sig.Params().At(1)
					if param2.Type().String() != "bool" {
						return true
					}
					var info BoolType
					info.pkg, err = filepath.Rel(rootDir, pkg.Dir)
					if err != nil {
						return false
					}
					info.name, err = extractName(call)
					if err != nil {
						return false
					}
					res = append(res, info)
				case "Int": // name string, default_, min, max int
					if sig.Params().Len() != 4 {
						return true
					}
					param1 := sig.Params().At(0)
					if param1.Type().String() != "string" {
						return true
					}
					param2 := sig.Params().At(1)
					if param2.Type().String() != "int" {
						return true
					}
					param3 := sig.Params().At(2)
					if param3.Type().String() != "int" {
						return true
					}
					param4 := sig.Params().At(3)
					if param4.Type().String() != "int" {
						return true
					}
					var info IntType
					info.pkg, err = filepath.Rel(rootDir, pkg.Dir)
					if err != nil {
						return false
					}
					info.name, err = extractName(call)
					if err != nil {
						return false
					}
					info.min, err = extractInt(call, 2)
					if err != nil {
						return false
					}
					info.max, err = extractInt(call, 3)
					if err != nil {
						return false
					}
					res = append(res, info)
				}
				return true
			})
		}
	}
	return res, nil
}

// func Inputs(ctx context.Context, sources billy.Filesystem) ([]InputInfo, error) {
// 	var res []InputInfo
//
// 	err := util.Walk(sources, ".", func(path string, info iofs.FileInfo, err error) error {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 		}
//
// 		if err != nil {
// 			return err
// 		}
//
// 		f, err := sources.Open(path)
// 		defer func() { _ = f.Close() }()
//
// 		if info.IsDir() || !strings.HasSuffix(path, "_test.go") {
// 			return nil
// 		}
//
// 		fileAst, err := parser.ParseFile(token.NewFileSet(), "", f, parser.SkipObjectResolution)
// 		if err != nil {
// 			return err
// 		}
//
// 		importMap := buildImportMap(fileAst)
//
// 		ast.Inspect(fileAst, func(n ast.Node) bool {
// 			call, ok := n.(*ast.CallExpr)
// 			if !ok {
// 				return true
// 			}
// 			sel, ok := call.Fun.(*ast.SelectorExpr)
// 			if !ok {
// 				return true
// 			}
//
// 			selType := pkg.TypesInfo.Uses[sel.Sel]
// 			if selType == nil {
// 				return true
// 			}
// 			fn, ok := selType.(*types.Func)
// 			if !ok {
// 				return true
// 			}
// 			if fn.Pkg().Path() != "benchspotter/benchinput" {
// 				return true
// 			}
// 			sig := fn.Type().(*types.Signature)
//
// 			switch fn.Name() {
// 			case "Bool": // name string, default_ bool
// 				if sig.Params().Len() != 2 {
// 					return true
// 				}
// 				param1 := sig.Params().At(0)
// 				if param1.Type().String() != "string" {
// 					return true
// 				}
// 				param2 := sig.Params().At(1)
// 				if param2.Type().String() != "bool" {
// 					return true
// 				}
// 				var info BoolType
// 				info.pkg, err = filepath.Rel(rootDir, pkg.Dir)
// 				if err != nil {
// 					return false
// 				}
// 				info.name, err = extractName(call)
// 				if err != nil {
// 					return false
// 				}
// 				res = append(res, info)
// 			case "Int": // name string, default_, min, max int
// 				if sig.Params().Len() != 4 {
// 					return true
// 				}
// 				param1 := sig.Params().At(0)
// 				if param1.Type().String() != "string" {
// 					return true
// 				}
// 				param2 := sig.Params().At(1)
// 				if param2.Type().String() != "int" {
// 					return true
// 				}
// 				param3 := sig.Params().At(2)
// 				if param3.Type().String() != "int" {
// 					return true
// 				}
// 				param4 := sig.Params().At(3)
// 				if param4.Type().String() != "int" {
// 					return true
// 				}
// 				var info IntType
// 				info.pkg, err = filepath.Rel(rootDir, pkg.Dir)
// 				if err != nil {
// 					return false
// 				}
// 				info.name, err = extractName(call)
// 				if err != nil {
// 					return false
// 				}
// 				info.min, err = extractInt(call, 2)
// 				if err != nil {
// 					return false
// 				}
// 				info.max, err = extractInt(call, 3)
// 				if err != nil {
// 				}
// 				res = append(res, info)
// 			}
// 			return true
// 		})
//
// 	})
//
// 	for _, pkg := range pkgs {
// 		for _, file := range pkg.Syntax {
//
// 		}
// 	}
// 	return res, nil
// }

func extractName(call *ast.CallExpr) (string, error) {
	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", fmt.Errorf("only string literals are supported for benchspotter inputs name")
	}
	name := lit.Value[1 : len(lit.Value)-1] // remove quotes
	return name, nil
}

func extractInt(call *ast.CallExpr, pos int) (int, error) {
	lit, ok := call.Args[pos].(*ast.BasicLit)
	if !ok || lit.Kind != token.INT {
		return 0, fmt.Errorf("only literals are supported for benchspotter inputs")
	}
	val, err := strconv.Atoi(lit.Value)
	if err != nil {
		return 0, err
	}
	return val, nil
}
