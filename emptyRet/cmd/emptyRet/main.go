package main

import (
	"emptyRet"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(emptyRet.Analyzer) }

