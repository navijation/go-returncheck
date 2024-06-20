package analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NewAnalyzer returns a new analyzer.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "standaloneFuncCalls",
		Doc:      "Reports standalone function calls",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ExprStmt)(nil),     // to find expression statements
		(*ast.CallExpr)(nil),     // to find call expressions
		(*ast.SelectorExpr)(nil), // to find selector expressions
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch stmt := n.(type) {
		case *ast.ExprStmt:
			// Check if the expression is a call expression
			switch expr := stmt.X.(type) {
			case *ast.CallExpr:
				// Report the standalone function call
				pass.Reportf(expr.Pos(), "RETC1: return value for function call %s is ignored", formatNode(pass, expr))
			case *ast.SelectorExpr:
				// Check if the selector expression's X is a call expression
				if callExpr, ok := expr.X.(*ast.CallExpr); ok {
					// Report the standalone function call in the selector
					pass.Reportf(callExpr.Pos(), "RETC1: return value for function call %s is ignored", formatNode(pass, callExpr))
				}
			}
		}
	})

	return nil, nil
}

// formatNode formats the given AST node as a string.
func formatNode(pass *analysis.Pass, node ast.Node) string {
	return pass.Fset.Position(node.Pos()).String()
}
