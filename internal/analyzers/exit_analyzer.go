package analyzers

import (
	"go/ast"

	log "github.com/sirupsen/logrus"
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

			log.Info(file.Name.Name)

			return true
		})
	}

	return nil, nil
}
