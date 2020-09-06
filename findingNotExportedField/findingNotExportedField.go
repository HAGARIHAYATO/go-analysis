package findingNotExportedField

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "findingNotExportedField is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "findingNotExportedField",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.StructType:
			for _, l := range n.Fields.List {
				if l.Tag == nil {
					continue
				}
				for _, name := range l.Names {
					if !name.IsExported() {
						pass.Reportf(l.Pos(), "is not exported with tag")
					}
				}
			}

		}
	})

	return nil, nil
}
