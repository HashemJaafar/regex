package main

// import (
// 	"fmt"
// 	"go/ast"
// 	"log"
// 	"os"
// 	"strings"

// 	"golang.org/x/tools/go/packages"
// )

// // Function represents a function with its path and name.
// type Function struct {
// 	Path string
// 	Name string
// 	Type string // "workspace", "third_party", "standard", "builtin"
// }

// // Method represents a method with its receiver type, path, and name.
// type Method struct {
// 	Receiver string
// 	Path     string
// 	Name     string
// 	Type     string // "workspace", "third_party", "standard", "builtin"
// }

// // parseFunctions parses the Go files in the specified directory and returns a map of functions to their respective calls.
// func parseFunctions(dir string) (map[Function][]Function, map[Method][]Function, error) {
// 	functionCalls := make(map[Function][]Function)
// 	methodCalls := make(map[Method][]Function)

// 	// Load packages in the directory
// 	cfg := &packages.Config{Mode: packages.LoadSyntax, Dir: dir, Tests: false}
// 	pkgs, err := packages.Load(cfg, "./...")
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	for _, pkg := range pkgs {
// 		for _, file := range pkg.Syntax {
// 			importPath := pkg.PkgPath

// 			// Extract functions and methods from the AST
// 			ast.Inspect(file, func(n ast.Node) bool {
// 				switch fn := n.(type) {
// 				case *ast.FuncDecl:
// 					if fn.Recv == nil {
// 						// Regular function
// 						caller := Function{
// 							Path: importPath,
// 							Name: fn.Name.Name,
// 							Type: classifyFunction(pkg.PkgPath),
// 						}
// 						ast.Inspect(fn.Body, func(n ast.Node) bool {
// 							switch call := n.(type) {
// 							case *ast.CallExpr:
// 								callee := extractFunctionCall(call)
// 								if callee != nil {
// 									callee.Type = classifyFunction(callee.Path)
// 									functionCalls[caller] = append(functionCalls[caller], *callee)
// 								}
// 							}
// 							return true
// 						})
// 					} else {
// 						// Method
// 						for _, recv := range fn.Recv.List {
// 							if starExpr, ok := recv.Type.(*ast.StarExpr); ok {
// 								if ident, ok := starExpr.X.(*ast.Ident); ok {
// 									method := Method{
// 										Receiver: ident.Name,
// 										Path:     importPath,
// 										Name:     fn.Name.Name,
// 										Type:     classifyFunction(pkg.PkgPath),
// 									}
// 									ast.Inspect(fn.Body, func(n ast.Node) bool {
// 										switch call := n.(type) {
// 										case *ast.CallExpr:
// 											callee := extractFunctionCall(call)
// 											if callee != nil {
// 												callee.Type = classifyFunction(callee.Path)
// 												methodCalls[method] = append(methodCalls[method], *callee)
// 											}
// 										}
// 										return true
// 									})
// 								}
// 							}
// 						}
// 					}
// 				}
// 				return true
// 			})
// 		}
// 	}

// 	return functionCalls, methodCalls, nil
// }

// // classifyFunction classifies a function based on its import path.
// func classifyFunction(importPath string) string {
// 	if strings.HasPrefix(importPath, "golang.org/x/") || strings.HasPrefix(importPath, "github.com/") || strings.HasPrefix(importPath, "bitbucket.org/") {
// 		return "third_party"
// 	}
// 	if importPath == "builtin" {
// 		return "builtin"
// 	}
// 	if strings.Contains(importPath, ".") {
// 		return "third_party"
// 	}
// 	if importPath == "std" {
// 		return "standard"
// 	}
// 	return "workspace"
// }

// // extractFunctionCall extracts the function call details from an AST CallExpr node.
// func extractFunctionCall(call *ast.CallExpr) *Function {
// 	switch fun := call.Fun.(type) {
// 	case *ast.Ident:
// 		return &Function{
// 			Path: "builtin",
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
// func generateDOT(functionCalls map[Function][]Function, methodCalls map[Method][]Function) string {
// 	var dotBuilder strings.Builder
// 	dotBuilder.WriteString("digraph G {\n")
// 	dotBuilder.WriteString("\trankdir=LR;\n")

// 	// Workspace functions
// 	dotBuilder.WriteString("\tsubgraph cluster_work_space {\n")
// 	dotBuilder.WriteString("\t\tlabel=\"Work Space\";\n")
// 	dotBuilder.WriteString("\t\tbgcolor=lightblue;\n")
// 	generateSubgraph(&dotBuilder, functionCalls, "workspace")
// 	generateMethodSubgraph(&dotBuilder, methodCalls, "workspace")
// 	dotBuilder.WriteString("\t}\n")

// 	// Third-party functions
// 	dotBuilder.WriteString("\tsubgraph cluster_third_party {\n")
// 	dotBuilder.WriteString("\t\tlabel=\"Third Party Library\";\n")
// 	dotBuilder.WriteString("\t\tbgcolor=lightgreen;\n")
// 	generateSubgraph(&dotBuilder, functionCalls, "third_party")
// 	generateMethodSubgraph(&dotBuilder, methodCalls, "third_party")
// 	dotBuilder.WriteString("\t}\n")

// 	// Standard library functions
// 	dotBuilder.WriteString("\tsubgraph cluster_standard {\n")
// 	dotBuilder.WriteString("\t\tlabel=\"Standard Library\";\n")
// 	dotBuilder.WriteString("\t\tbgcolor=lightyellow;\n")
// 	generateSubgraph(&dotBuilder, functionCalls, "standard")
// 	generateMethodSubgraph(&dotBuilder, methodCalls, "standard")
// 	dotBuilder.WriteString("\t}\n")

// 	// Built-in functions
// 	dotBuilder.WriteString("\tsubgraph cluster_builtin {\n")
// 	dotBuilder.WriteString("\t\tlabel=\"Built-in\";\n")
// 	dotBuilder.WriteString("\t\tbgcolor=lightgrey;\n")
// 	generateSubgraph(&dotBuilder, functionCalls, "builtin")
// 	generateMethodSubgraph(&dotBuilder, methodCalls, "builtin")
// 	dotBuilder.WriteString("\t}\n")

// 	// Function calls
// 	for caller, callees := range functionCalls {
// 		callerName := fmt.Sprintf("%s.%s", caller.Path, caller.Name)
// 		callerName = convertBackSlashToForwardSlash(callerName)
// 		for _, callee := range callees {
// 			calleeName := fmt.Sprintf("%s.%s", callee.Path, callee.Name)
// 			calleeName = convertBackSlashToForwardSlash(calleeName)
// 			dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", callerName, calleeName))
// 		}
// 	}

// 	for caller, callees := range methodCalls {
// 		callerName := fmt.Sprintf("%s.%s", caller.Path, caller.Name)
// 		callerName = convertBackSlashToForwardSlash(callerName)
// 		for _, callee := range callees {
// 			calleeName := fmt.Sprintf("%s.%s", callee.Path, callee.Name)
// 			calleeName = convertBackSlashToForwardSlash(calleeName)
// 			dotBuilder.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", callerName, calleeName))
// 		}
// 	}

// 	dotBuilder.WriteString("}\n")
// 	return dotBuilder.String()
// }

// // generateSubgraph generates the subgraph for a specific function type.
// func generateSubgraph(dotBuilder *strings.Builder, functionCalls map[Function][]Function, functionType string) {
// 	for caller := range functionCalls {
// 		if caller.Type == functionType {
// 			functionName := fmt.Sprintf("%s.%s", caller.Path, caller.Name)
// 			functionName = convertBackSlashToForwardSlash(functionName)
// 			dotBuilder.WriteString(fmt.Sprintf("\t\t\"%s\";\n", functionName))
// 		}
// 	}
// }

// // generateMethodSubgraph generates the subgraph for methods of a specific type.
// func generateMethodSubgraph(dotBuilder *strings.Builder, methodCalls map[Method][]Function, methodType string) {
// 	for caller := range methodCalls {
// 		if caller.Type == methodType {
// 			methodName := fmt.Sprintf("%s.%s", caller.Path, caller.Name)
// 			methodName = convertBackSlashToForwardSlash(methodName)
// 			dotBuilder.WriteString(fmt.Sprintf("\t\t\"%s\";\n", methodName))
// 		}
// 	}
// }

// // convertBackSlashToForwardSlash converts backslashes to forward slashes in a string.
// func convertBackSlashToForwardSlash(s string) string {
// 	return strings.ReplaceAll(s, "\\", "/")
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatalf("Usage: %s <directory>\n", os.Args[0])
// 	}

// 	dir := os.Args[1]
// 	functionCalls, methodCalls, err := parseFunctions(dir)
// 	if err != nil {
// 		log.Fatalf("Failed to parse functions: %v\n", err)
// 	}

// 	dotContent := generateDOT(functionCalls, methodCalls)
// 	dotFileName := "function_call_graph.dot"
// 	err = os.WriteFile(dotFileName, []byte(dotContent), 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to write DOT file: %v\n", err)
// 	}

// 	fmt.Printf("Function call graph saved as %s\n", dotFileName)
// }
