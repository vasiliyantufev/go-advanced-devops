package exitanalyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var ExitAnalyzer = &analysis.Analyzer{
	Name: "exit_analyzer",
	Doc:  "exit analyzer - os.Exit",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {

		ast.Inspect(file, func(node ast.Node) bool {

			switch x := node.(type) {
			case *ast.FuncDecl:
				if x.Name.Name != "main" {
					return false
				}
			case *ast.CallExpr:
				switch ce := x.Fun.(type) {
				case *ast.SelectorExpr:
					if ce.Sel.Name == "Exit" && fmt.Sprintf("%s", ce.X) == "os" {
						pass.Reportf(ce.Pos(), "undesirable use of os.Exit")
					}
				}
			}
			return true
		})
	}

	return nil, nil
}
