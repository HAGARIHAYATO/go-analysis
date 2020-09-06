package findingNotExportedFunc

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "findingNotExportedFunc is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "findingNotExportedFunc",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if !n.Name.IsExported() {
				pass.Reportf(n.Pos(), "notExported")
			}
		}
	})

	return nil, nil
}

