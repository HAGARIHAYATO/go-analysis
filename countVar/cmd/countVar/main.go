package main

import (
	"countVar"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(countVar.Analyzer) }

