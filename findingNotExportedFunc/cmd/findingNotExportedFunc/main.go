package main

import (
	"findingNotExportedFunc"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findingNotExportedFunc.Analyzer) }

