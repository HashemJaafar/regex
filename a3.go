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
// 	"text/template"
// )

// // Function represents a function in the code.
// type Function struct {
// 	PackagePath string
// 	Name        string
// 	Calls       []string
// }

// // parseFunctions parses the Go files in the specified directory and returns a map of function names to their respective Function structs.
// func parseFunctions(dir string) (map[string]*Function, error) {
// 	functions := make(map[string]*Function)

// 	// Walk through the directory and parse each Go file
// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
// 			// Parse the Go file
// 			fileFunctions, err := parseGoFile(path)
// 			if err != nil {
// 				return err
// 			}
// 			// Merge the functions from this file into the main map
// 			for _, fn := range fileFunctions {
// 				key := fmt.Sprintf("%s:%s", fn.PackagePath, fn.Name)
// 				if existingFn, ok := functions[key]; ok {
// 					// Merge calls if the function already exists
// 					existingFn.Calls = append(existingFn.Calls, fn.Calls...)
// 				} else {
// 					functions[key] = fn
// 				}
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return functions, nil
// }

// // parseGoFile parses a single Go file and extracts its functions and their calls.
// func parseGoFile(filename string) ([]*Function, error) {
// 	var functions []*Function

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
// 			function := &Function{
// 				PackagePath: importPath,
// 				Name:        fn.Name.Name,
// 			}
// 			ast.Inspect(fn.Body, func(n ast.Node) bool {
// 				switch call := n.(type) {
// 				case *ast.CallExpr:
// 					if ident, ok := call.Fun.(*ast.Ident); ok {
// 						calledFnName := ident.Name
// 						if !ast.IsExported(calledFnName) {
// 							// If the called function is not exported, prepend the package path
// 							calledFnName = fmt.Sprintf("%s:%s", importPath, calledFnName)
// 						}
// 						function.Calls = append(function.Calls, calledFnName)
// 					}
// 				}
// 				return true
// 			})
// 			functions = append(functions, function)
// 		}
// 		return true
// 	})

// 	return functions, nil
// }

// // generateDOT generates the DOT representation of the function call graph.
// func generateDOT(functions map[string]*Function) string {
// 	var dotTemplate = `
// digraph G {
// {{- range $key, $function := . }}
// 	"{{$key}}" [label="{{$key}}"];
// 	{{- range $function.Calls }}
// 		"{{$key}}" -> "{{.}}";
// 	{{- end }}
// {{- end }}
// }
// `

// 	tmpl, err := template.New("dot").Parse(dotTemplate)
// 	if err != nil {
// 		log.Fatal("Error parsing DOT template:", err)
// 	}

// 	var dotBuffer strings.Builder
// 	err = tmpl.Execute(&dotBuffer, functions)
// 	if err != nil {
// 		log.Fatal("Error executing DOT template:", err)
// 	}

// 	return dotBuffer.String()
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatalf("Usage: %s <directory>\n", os.Args[0])
// 	}

// 	dir := os.Args[1]
// 	functions, err := parseFunctions(dir)
// 	if err != nil {
// 		log.Fatalf("Failed to parse functions: %v\n", err)
// 	}

// 	dotContent := generateDOT(functions)
// 	dotFileName := "function_call_graph.dot"
// 	err = os.WriteFile(dotFileName, []byte(dotContent), 0644)
// 	if err != nil {
// 		log.Fatalf("Failed to write DOT file: %v\n", err)
// 	}

// 	fmt.Printf("Function call graph saved as %s\n", dotFileName)
// }
