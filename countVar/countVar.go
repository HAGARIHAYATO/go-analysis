package countVar

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "countVar is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "countVar",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	count := 0
	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			for _, s := range n.Specs {
				switch s := s.(type) {
				case *ast.ValueSpec:
					count += len(s.Names)
				}
			}
		}
	})
	fmt.Println(count)
	return nil, nil
}

