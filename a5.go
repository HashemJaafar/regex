package main

// import (
// 	"fmt"
// 	"go/ast"
// 	"go/parser"
// 	"go/token"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// // Function represents a function with its path and name.
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

// // parseGoFile parses a single Go file and extracts its function calls.
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

// 	// Walk through the AST and extract function declarations and calls
// 	ast.Inspect(node, func(n ast.Node) bool {
// 		switch fn := n.(type) {
// 		case *ast.FuncDecl:
// 			caller := Function{
// 				Path: importPath,
// 				Name: fn.Name.Name,
// 			}
// 			ast.Inspect(fn.Body, func(n ast.Node) bool {
// 				switch call := n.(type) {
// 				case *ast.CallExpr:
// 					callee := extractFunctionCall(call)
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

// // extractFunctionCall extracts the function call details from an AST CallExpr node.
// func extractFunctionCall(call *ast.CallExpr) *Function {
// 	switch fun := call.Fun.(type) {
// 	case *ast.Ident:
// 		return &Function{
// 			Path: "BuiltIn",
// 			Name: fun.Name,
// 		}
// 	case *ast.SelectorExpr:
// 		if ident, ok := fun.X.(*ast.Ident); ok {
// 			return &Function{
// 				Path: ident.Name,
// 				Name: fun.Sel.Name,
// 			}
// 		}
// 	}
// 	return nil
// }

// // generateDOT generates the DOT representation of the function call graph.
// func generateDOT(functionCalls map[Function][]Function) string {
// 	var dotBuilder strings.Builder
// 	dotBuilder.WriteString("digraph G {\n")

// 	for caller, callees := range functionCalls {
// 		callerName := fmt.Sprintf("%s\\%s", caller.Path, caller.Name)
// 		dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" [label=\"%s\"];\n", callerName, callerName))
// 		for _, callee := range callees {
// 			calleeName := fmt.Sprintf("%s\\%s", callee.Path, callee.Name)
// 			dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", callerName, calleeName))
// 		}
// 	}

// 	dotBuilder.WriteString("}\n")
// 	return dotBuilder.String()
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
