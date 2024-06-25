package analyzer

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NewAnalyzer returns a new analyzer.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "returncheck",
		Doc:      "Reports ignored return values",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ExprStmt)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		stmt, ok := n.(*ast.ExprStmt)
		if !ok {
			return
		}

		call, ok := stmt.X.(*ast.CallExpr)
		if !ok {
			return
		}

		// Skip if the function doesn't return any values
		if !returnsValues(pass, call) {
			return
		}

		name := functionName(call)
		if hasMultipleReturnValues(pass, call) {
			pass.Reportf(call.Pos(), "RETC1: return values for function call %s are ignored", name)
		} else {
			pass.Reportf(call.Pos(), "RETC1: return value for function call %s is ignored", name)
		}
	})

	return nil, nil
}

func returnsValues(pass *analysis.Pass, call *ast.CallExpr) bool {
	if t, ok := pass.TypesInfo.Types[call]; ok {
		return t.Type != types.Typ[types.Invalid] && t.Type != nil
	}
	return false
}

func hasMultipleReturnValues(pass *analysis.Pass, call *ast.CallExpr) bool {
	if t, ok := pass.TypesInfo.Types[call]; ok {
		if tuple, ok := t.Type.(*types.Tuple); ok {
			return tuple.Len() > 1
		}
	}
	return false
}

func functionName(call *ast.CallExpr) string {
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		return fun.Name + "()"
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s()", exprString(fun.X), fun.Sel.Name)
	}
	return ""
}

func exprString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprString(e.X), e.Sel.Name)
	case *ast.CallExpr:
		return fmt.Sprintf("%s()", functionName(e))
	}
	return ""
}