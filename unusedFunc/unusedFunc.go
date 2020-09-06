package unusedFunc

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "unusedFunc is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "unusedFunc",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	m := map[string]bool{}
	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			// TODO
		}
	})
	fmt.Println(m)
	return nil, nil
}

