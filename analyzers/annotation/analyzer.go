package annotation

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const annotationPkgPath = "github.com/openmcp-project/openmcp-operator/api/constants"

var Analyzer = &analysis.Analyzer{
	Name: "operation_annotation",
	Doc:  "checks that switches on OperationAnnotation handle both Ignore and Reconcile cases",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			switchStmt, ok := n.(*ast.SwitchStmt)
			if !ok {
				return true
			}
			touchesPkg, hasIgnore, hasReconcile := inspectSwitchCases(pass, switchStmt)
			if !touchesPkg {
				return true
			}
			if !hasIgnore {
				pass.Reportf(switchStmt.Pos(), "switch on OperationAnnotation missing case for OperationAnnotationValueIgnore")
			}
			if !hasReconcile {
				pass.Reportf(switchStmt.Pos(), "switch on OperationAnnotation missing case for OperationAnnotationValueReconcile")
			}
			return true
		})
	}
	return nil, nil
}

func inspectSwitchCases(pass *analysis.Pass, sw *ast.SwitchStmt) (touchesPkg, hasIgnore, hasReconcile bool) {
	for _, stmt := range sw.Body.List {
		cc, ok := stmt.(*ast.CaseClause)
		if !ok || cc.List == nil {
			continue
		}
		for _, expr := range cc.List {
			name, pkg := resolveConst(pass, expr)
			if pkg != annotationPkgPath {
				continue
			}
			touchesPkg = true
			switch name {
			case "OperationAnnotationValueIgnore":
				hasIgnore = true
			case "OperationAnnotationValueReconcile":
				hasReconcile = true
			}
		}
	}
	return
}

// resolveConst resolves a selector expression (pkg.Name) to its constant name and package path
// using type information, so import aliases are handled transparently.
func resolveConst(pass *analysis.Pass, expr ast.Expr) (name, pkgPath string) {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return "", ""
	}
	obj, ok := pass.TypesInfo.Uses[sel.Sel].(*types.Const)
	if !ok {
		return "", ""
	}
	return obj.Name(), obj.Pkg().Path()
}
