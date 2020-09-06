package main

import (
	"findingNotExportedField"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findingNotExportedField.Analyzer) }

