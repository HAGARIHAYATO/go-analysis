package unusedFunc_test

import (
	"testing"

	"unusedFunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unusedFunc.Analyzer, "a")
}

