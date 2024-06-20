package analyzer

import (
	"fmt"
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
				pass.Reportf(expr.Pos(), "RETC1: return value for function call %s is ignored", functionName(expr))
			case *ast.SelectorExpr:
				// Check if the selector expression's X is a call expression
				if callExpr, ok := expr.X.(*ast.CallExpr); ok {
					// Report the standalone function call in the selector
					pass.Reportf(callExpr.Pos(), "RETC1: return value for function call %s is ignored", functionName(callExpr))
				}
			}
		}
	})

	return nil, nil
}

// functionName extracts the name of the function from a CallExpr.
func functionName(call *ast.CallExpr) string {
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		return fun.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", selectorName(fun.X), fun.Sel.Name)
	}
	return ""
}

// selectorName returns the name of a selector's base.
func selectorName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", selectorName(e.X), e.Sel.Name)
	}
	return ""
}
