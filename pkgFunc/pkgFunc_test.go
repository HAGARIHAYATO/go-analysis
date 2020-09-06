package pkgFunc_test

import (
	"testing"

	"pkgFunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, pkgFunc.Analyzer, "a")
}

