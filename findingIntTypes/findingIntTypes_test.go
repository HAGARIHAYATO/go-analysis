package findingIntTypes_test

import (
	"testing"

	"findingIntTypes"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, findingIntTypes.Analyzer, "a")
}

