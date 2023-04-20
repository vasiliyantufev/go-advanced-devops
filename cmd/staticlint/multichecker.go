package main

import (
	"go/ast"

	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
)

// make run_multichecker
var ErrNoExitAnalyzer = &analysis.Analyzer{
	Name: "noexit",
	Doc:  "os.Exit",
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

func main() {

	multichecker.Main(
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		ErrNoExitAnalyzer,
	)
}
