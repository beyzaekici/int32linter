package linter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "int32linter",
	Doc:  "check for int32 usage",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if id, ok := n.(*ast.Ident); ok {
				if id.Name == "int32" {
					pass.Reportf(id.Pos(), "int32 type is used, consider using int or int64")
				}
			}
			return true
		})
	}
	return nil, nil
}
