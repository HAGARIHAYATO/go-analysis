package pkgFunc

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "pkgFunc is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "pkgFunc",
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
		case *ast.FuncDecl:
			count += 1
			for _, l := range n.Body.List {
				switch l := l.(type) {
				case *ast.AssignStmt:
					for _, f := range l.Rhs {
						_, ok := f.(*ast.FuncLit)
						if ok {
							count += 1
						}
					}
				}
			}
		}
	})
	fmt.Println(count)
	return nil, nil
}

