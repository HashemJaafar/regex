package main

// import (
// 	"fmt"
// 	"go/ast"
// 	"go/parser"
// 	"go/token"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"sort"
// 	"strings"
// )

// func main() {
// 	// Specify the directory to read Go files from
// 	dir := "C:/Users/hashem/Desktop/programing/my_big_project/dht"

// 	// Create a map to store occurrences of names
// 	nameOccurrences := make(map[string]int)

// 	// Create a map to store the types of names
// 	nameTypes := make(map[string]string)

// 	// Read all Go files in the directory
// 	files, err := ioutil.ReadDir(dir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		if strings.HasSuffix(file.Name(), ".go") {
// 			filePath := filepath.Join(dir, file.Name())
// 			parseFile(filePath, nameOccurrences, nameTypes)
// 		}
// 	}

// 	// Write the results to a text file
// 	writeResultsToFile(dir+"/words_count.txt", nameOccurrences, nameTypes)
// }

// func parseFile(filePath string, nameOccurrences map[string]int, nameTypes map[string]string) {
// 	fset := token.NewFileSet()
// 	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
// 	if err != nil {
// 		log.Printf("Error parsing file %s: %v\n", filePath, err)
// 		return
// 	}

// 	ast.Inspect(node, func(n ast.Node) bool {
// 		switch x := n.(type) {
// 		case *ast.FuncDecl:
// 			name := x.Name.Name
// 			nameOccurrences[name]++
// 			nameTypes[name] = "Func"
// 			if x.Recv != nil {
// 				for _, field := range x.Recv.List {
// 					for _, name := range field.Names {
// 						nameOccurrences[name.Name]++
// 						nameTypes[name.Name] = "Method"
// 					}
// 				}
// 			}
// 		case *ast.GenDecl:
// 			for _, spec := range x.Specs {
// 				switch spec := spec.(type) {
// 				case *ast.TypeSpec:
// 					name := spec.Name.Name
// 					nameOccurrences[name]++
// 					if _, ok := spec.Type.(*ast.InterfaceType); ok {
// 						nameTypes[name] = "Interface"
// 					} else {
// 						nameTypes[name] = "Type"
// 					}
// 				case *ast.ValueSpec:
// 					for _, name := range spec.Names {
// 						nameOccurrences[name.Name]++
// 						if x.Tok == token.CONST {
// 							nameTypes[name.Name] = "Const"
// 						} else {
// 							nameTypes[name.Name] = "Global Var"
// 						}
// 					}
// 				}
// 			}
// 		case *ast.AssignStmt:
// 			for _, lhs := range x.Lhs {
// 				if ident, ok := lhs.(*ast.Ident); ok {
// 					name := ident.Name
// 					nameOccurrences[name]++
// 					nameTypes[name] = "Local Var"
// 				}
// 			}
// 		}
// 		return true
// 	})
// }

// func writeResultsToFile(fileName string, nameOccurrences map[string]int, nameTypes map[string]string) {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		log.Fatalf("Could not create results file: %v\n", err)
// 	}
// 	defer file.Close()

// 	fmt.Fprintf(file, "%-10s %-15s %-20s\n", "Count", "Type", "Name")
// 	fmt.Fprintf(file, "%-10s %-15s %-20s\n", "-----", "----", "----")

// 	// Create a slice of keys from the map and sort it
// 	var keys []string
// 	for key := range nameOccurrences {
// 		keys = append(keys, key)
// 	}
// 	sort.Strings(keys)

// 	// Write sorted results to the file
// 	for _, name := range keys {
// 		count := nameOccurrences[name]
// 		nameType := nameTypes[name]
// 		_, err := fmt.Fprintf(file, "%-10d %-15s %-20s\n", count, nameType, name)
// 		if err != nil {
// 			log.Printf("Error writing to file: %v\n", err)
// 		}
// 	}
// }
