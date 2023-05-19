package main

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/analyzer/exitanalyzer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/analyzer/ignoreerrorsanalyzer"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
)

// make run_multichecker
func main() {

	multichecker.Main(
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		exitanalyzer.ExitAnalyzer,
		ignoreerrorsanalyzer.IgnoreErrorsAnalizer,
	)
}
