package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
)

var ErrNoExitAnalyzer = &analysis.Analyzer{
	Name: "noexit",
	Doc:  "os.Exit",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	return nil, nil
}

func main() {

	multichecker.Main(
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		ErrNoExitAnalyzer,
		//mychecks...,
	)
}
