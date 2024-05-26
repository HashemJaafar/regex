package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

func Test(t *testing.T) {
	src := `
package main

import(
	 "fmt"
	 f "github.com/example/package"
)

func main() {
	fmt.Println("Hello, world!")
	add(1, 2)
	f.gg()
}

type s struct{}
func (a s)add(a, b int) int {
	return a + b
}
`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok {
			// Extract the callee path
			calleePath := extractCalleePath(callExpr.Fun)
			fmt.Println("Callee function path:", calleePath)
		}
		return true
	})
}

func extractCalleePath(expr ast.Expr) string {
	switch fun := expr.(type) {
	case *ast.Ident:
		// Identifiers represent functions defined within the same package.
		// You can determine if it's a built-in function, standard library function,
		// or a function within the workspace based on its name and package path.
		return "WorkspacePath." + fun.Name // Example: "WorkspacePath.functionName"
	case *ast.SelectorExpr:
		// Selector expressions represent functions from imported packages.
		// You can extract the package path and function name from the selector.
		// Example: "github.com/example/package.functionName"
		path, err := ResolveStandardLibraryPackage(fun.Sel.Name)
		if err != nil {
			fmt.Println(err)
		}
		return path
	// Handle other cases (e.g., *ast.CallExpr, etc.) as needed
	default:
		return "" // Unknown or unsupported expression type
	}
}

func extractSelectorPath(sel *ast.SelectorExpr) string {
	// Extract package path and function name from the selector expression
	pkgPath := extractPackagePath(sel.X)
	return pkgPath + "." + sel.Sel.Name
}

func extractPackagePath(expr ast.Expr) string {
	// Extract the package path from the expression
	// For simplicity, assume we're dealing with an *ast.Ident representing the package name.
	// In a real implementation, you may need to handle other cases.
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name // Example: "github.com/example/package"
	}
	return "" // Package path not found or unsupported expression type
}

// ResolveImportPath resolves an import path to a package path.
func ResolveStandardLibraryPackage(importPath string) (string, error) {
	pkg, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return pkg.Dir, nil
}
