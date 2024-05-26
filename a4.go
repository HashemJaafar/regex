package main

// import (
// 	"fmt"
// 	"go/ast"
// 	"go/build"
// 	"go/parser"
// 	"go/token"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// // Function represents a function or method with its path and name.
// type Function struct {
// 	Path string
// 	Name string
// }

// // parseFunctions parses the Go files in the specified directory and returns a map of functions to their respective calls.
// func parseFunctions(dir string) (map[Function][]Function, error) {
// 	functionCalls := make(map[Function][]Function)

// 	// Walk through the directory and parse each Go file
// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
// 			// Parse the Go file
// 			fileCalls, err := parseGoFile(path)
// 			if err != nil {
// 				return err
// 			}
// 			// Merge the calls from this file into the main map
// 			for caller, callees := range fileCalls {
// 				functionCalls[caller] = append(functionCalls[caller], callees...)
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return functionCalls, nil
// }

// // parseGoFile parses a single Go file and extracts its function and method calls.
// func parseGoFile(filename string) (map[Function][]Function, error) {
// 	functionCalls := make(map[Function][]Function)

// 	// Parse the Go file
// 	fset := token.NewFileSet()
// 	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get package path
// 	importPath := filepath.Dir(filename)

// 	// Walk through the AST and extract function and method declarations and calls
// 	ast.Inspect(node, func(n ast.Node) bool {
// 		switch fn := n.(type) {
// 		case *ast.FuncDecl:
// 			caller := Function{
// 				Path: importPath,
// 				Name: fn.Name.Name,
// 			}
// 			if fn.Recv != nil {
// 				// This is a method
// 				caller.Name = fmt.Sprintf("(%s).%s", receiverType(fn.Recv.List[0].Type), fn.Name.Name)
// 			}
// 			ast.Inspect(fn.Body, func(n ast.Node) bool {
// 				switch call := n.(type) {
// 				case *ast.CallExpr:
// 					callee := extractFunctionCall(call, node.Imports, importPath)
// 					if callee != nil {
// 						functionCalls[caller] = append(functionCalls[caller], *callee)
// 					}
// 				}
// 				return true
// 			})
// 		}
// 		return true
// 	})

// 	return functionCalls, nil
// }

// // receiverType returns the type of the receiver as a string.
// func receiverType(expr ast.Expr) string {
// 	switch t := expr.(type) {
// 	case *ast.Ident:
// 		return t.Name
// 	case *ast.StarExpr:
// 		return receiverType(t.X)
// 	case *ast.SelectorExpr:
// 		return fmt.Sprintf("%s.%s", receiverType(t.X), t.Sel.Name)
// 	default:
// 		return "unknown"
// 	}
// }

// // extractFunctionCall extracts the function or method call details from an AST CallExpr node.
// func extractFunctionCall(call *ast.CallExpr, imports []*ast.ImportSpec, currentPkg string) *Function {
// 	switch fun := call.Fun.(type) {
// 	case *ast.Ident:
// 		return &Function{
// 			Path: "BuiltIn",
// 			Name: fun.Name,
// 		}
// 	case *ast.SelectorExpr:
// 		if ident, ok := fun.X.(*ast.Ident); ok {
// 			for _, imp := range imports {
// 				if imp.Name != nil && imp.Name.Name == ident.Name || filepath.Base(imp.Path.Value) == ident.Name {
// 					impPath := strings.Trim(imp.Path.Value, "\"")
// 					if isStandardLibrary(impPath) {
// 						return &Function{
// 							Path: impPath,
// 							Name: fun.Sel.Name,
// 						}
// 					} else {
// 						return &Function{
// 							Path: impPath,
// 							Name: fun.Sel.Name,
// 						}
// 					}
// 				}
// 			}
// 			// Handle unimported packages (assume it's from current package)
// 			return &Function{
// 				Path: currentPkg,
// 				Name: fun.Sel.Name,
// 			}
// 		}
// 	}
// 	return nil
// }

// // isStandardLibrary checks if the import path is part of the standard library.
// func isStandardLibrary(importPath string) bool {
// 	pkg, err := build.Import(importPath, "", build.FindOnly)
// 	if err != nil {
// 		return false
// 	}
// 	return pkg.Goroot
// }

// // generateDOT generates the DOT representation of the function call graph.
// func generateDOT(functionCalls map[Function][]Function) string {
// 	var dotBuilder strings.Builder
// 	dotBuilder.WriteString("digraph G {\n")

// 	for caller, callees := range functionCalls {
// 		callerName := fmt.Sprintf("%s.%s", caller.Path, caller.Name)
// 		callerName = convertBackSlashToForwardSlash(callerName)
// 		callerColor := determineColor(caller.Path)

// 		dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" [label=\"%s\" fontcolor=\"%s\"];\n", callerName, callerName, callerColor))
// 		for _, callee := range callees {
// 			calleeName := fmt.Sprintf("%s.%s", callee.Path, callee.Name)
// 			calleeName = convertBackSlashToForwardSlash(calleeName)
// 			calleeColor := determineColor(callee.Path)

// 			dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" [label=\"%s\" fontcolor=\"%s\"];\n", calleeName, calleeName, calleeColor))
// 			dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", callerName, calleeName))
// 		}
// 	}

// 	dotBuilder.WriteString("}\n")
// 	return dotBuilder.String()
// }

// // convertBackSlashToForwardSlash replaces all backslashes with forward slashes in a string.
// func convertBackSlashToForwardSlash(s string) string {
// 	return strings.ReplaceAll(s, "\\", "/")
// }

// // determineColor returns the color based on the function path.
// func determineColor(path string) string {
// 	if path == "BuiltIn" {
// 		return "blue"
// 	}
// 	if isStandardLibrary(path) {
// 		return "green"
// 	}
// 	if strings.Contains(path, "golang.org") || strings.Contains(path, "github.com") {
// 		return "red"
// 	}
// 	return "purple"
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatalf("Usage: %s <directory>\n", os.Args[0])
// 	}

// 	dir := os.Args[1]
// 	functionCalls, err := parseFunctions(dir)
// 	if err != nil {
// 		log.Fatalf("Failed to parse functions: %v\n", err)
// 	}

// 	dotContent := generateDOT(functionCalls)
// 	dotFileName := "function_call_graph.dot"
// 	err = os.WriteFile(dotFileName, []byte(dotContent), 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to write DOT file: %v\n", err)
// 	}

// 	fmt.Printf("Function call graph saved as %s\n", dotFileName)
// }
