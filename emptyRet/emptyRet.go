package emptyRet

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "emptyRet is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "emptyRet",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			for _, b := range n.Body.List {
				switch b := b.(type) {
				case *ast.ReturnStmt:
					for _, val := range b.Results {
						val, _ := val.(*ast.BasicLit)
						if val.Value == "\"\"" {
							pass.Reportf(val.Pos(), "return empty string")
						}
					}
				case *ast.AssignStmt:
					for _, r := range b.Rhs {
						r, _ := r.(*ast.FuncLit)
						walk(r, pass)
					}
				}
			}
		}
	})
	return nil, nil
}

func walk(n *ast.FuncLit, pass *analysis.Pass) {
	for _, b := range n.Body.List {
		switch b := b.(type) {
		case *ast.ReturnStmt:
			for _, val := range b.Results {
				val, _ := val.(*ast.BasicLit)
				if val.Value == "\"\"" {
					pass.Reportf(val.Pos(), "return empty string")
				}
			}
		case *ast.AssignStmt:
			for _, r := range b.Rhs {
				r, _ := r.(*ast.FuncLit)
				walk(r, pass)
			}
		}
	}
}
