package locate

import (
	"context"
	"fmt"
	"go/ast"
	"go/types"
	"path/filepath"
)

type BenchInfo struct {
	Name    string
	Package string
}

func Benchmarks(ctx context.Context, rootDir string) ([]BenchInfo, error) {
	rootDir, err := filepath.Abs(rootDir)
	if err != nil {
		return nil, err
	}

	pkgs, err := packages(rootDir)
	if err != nil {
		return nil, err
	}

	var res []BenchInfo

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}

			ast.Inspect(file, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok || fn.Recv != nil {
					return true
				}

				// Must start with "Benchmark"
				if !fn.Name.IsExported() || len(fn.Name.Name) < 9 || fn.Name.Name[:9] != "Benchmark" {
					return true
				}

				// Must have one parameter: *testing.B
				if fn.Type.Params.NumFields() != 1 {
					return true
				}
				param := fn.Type.Params.List[0]
				typ := pkg.TypesInfo.TypeOf(param.Type)
				if ptr, ok := typ.(*types.Pointer); ok {
					if named, ok := ptr.Elem().(*types.Named); ok {
						if named.Obj().Pkg().Path() == "testing" && named.Obj().Name() == "B" {
							p, err := filepath.Rel(rootDir, pkg.Dir)
							if err != nil {
								fmt.Println(err)
							}
							res = append(res, BenchInfo{
								Name:    fn.Name.Name,
								Package: p,
							})
						}
					}
				}

				return true
			})
		}
	}

	return res, nil
}
