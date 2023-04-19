package main

import (
	"github.com/vasiliyantufev/go-advanced-devops/cmd/errcheckanalyzer"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/structtag"
)

func main() {
	multichecker.Main(
		errcheckanalyzer.ErrCheckAnalyzer, // или errcheckanalyzer.ErrCheckAnalyzer, если анализатор импортируется
		printf.Analyzer,
		shadow.Analyzer,
		shift.Analyzer, // добавляем анализатор в вызов multichecker
		structtag.Analyzer,
	)
}
