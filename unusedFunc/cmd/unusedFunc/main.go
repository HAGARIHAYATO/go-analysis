package main

import (
	"unusedFunc"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unusedFunc.Analyzer) }

