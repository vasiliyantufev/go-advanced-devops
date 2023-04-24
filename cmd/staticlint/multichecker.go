package main

import (
	"github.com/vasiliyantufev/go-advanced-devops/internal/analyzers/exit_analyzer"
	"github.com/vasiliyantufev/go-advanced-devops/internal/analyzers/ignore_errors_analyzer"
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
		exit_analyzer.ExitAnalyzer,
		ignore_errors_analyzer.IgnoreErrorsAnalizer,
	)
}
