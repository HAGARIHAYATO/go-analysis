package main

import (
	"findingIntTypes"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findingIntTypes.Analyzer) }

