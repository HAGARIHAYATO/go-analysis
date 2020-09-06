package findingNotExportedFunc_test

import (
	"testing"

	"findingNotExportedFunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, findingNotExportedFunc.Analyzer, "a")
}

