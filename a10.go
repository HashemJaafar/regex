package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Object struct {
	Path         string
	Selector     string
	Name         string
	Type         string // "method" or "function" or struct or type or interface or var or const
	IsExportable bool
}

var calls = make(map[Object][]Object)

func parse(dir string) error {

	// Walk through the directory and parse each Go file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// Parse the Go file
			err := parseGoFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// parseGoFile parses a single Go file and extracts its function calls.
func parseGoFile(path string) error {

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	// Get package path
	importPath := filepath.Dir(path)

	// Walk through the AST and extract function declarations and calls
	ast.Inspect(node, func(n ast.Node) bool {
		switch n1 := n.(type) {
		case *ast.Comment:
		case *ast.CommentGroup:
		case *ast.Field:
		case *ast.FieldList:
		case *ast.BadExpr:
		case *ast.Ident:
		case *ast.Ellipsis:
		case *ast.BasicLit:
		case *ast.FuncLit:
		case *ast.CompositeLit:
		case *ast.ParenExpr:
		case *ast.SelectorExpr:
		case *ast.IndexExpr:
		case *ast.IndexListExpr:
		case *ast.SliceExpr:
		case *ast.TypeAssertExpr:
		case *ast.CallExpr:
		case *ast.StarExpr:
		case *ast.UnaryExpr:
		case *ast.BinaryExpr:
		case *ast.KeyValueExpr:
		case *ast.ArrayType:
		case *ast.StructType:
		case *ast.FuncType:
		case *ast.InterfaceType:
		case *ast.MapType:
		case *ast.ChanType:
		case *ast.BadStmt:
		case *ast.DeclStmt:
		case *ast.EmptyStmt:
		case *ast.LabeledStmt:
		case *ast.ExprStmt:
		case *ast.SendStmt:
		case *ast.IncDecStmt:
		case *ast.AssignStmt:
		case *ast.GoStmt:
		case *ast.DeferStmt:
		case *ast.ReturnStmt:
		case *ast.BranchStmt:
		case *ast.BlockStmt:
		case *ast.IfStmt:
		case *ast.CaseClause:
		case *ast.SwitchStmt:
		case *ast.TypeSwitchStmt:
		case *ast.CommClause:
		case *ast.SelectStmt:
		case *ast.ForStmt:
		case *ast.RangeStmt:
		case *ast.ImportSpec:
		case *ast.ValueSpec:
		case *ast.TypeSpec:
		case *ast.BadDecl:
		case *ast.GenDecl:
		case *ast.FuncDecl:
			if n1.Recv == nil {
				caller := Object{
					Path:         importPath,
					Selector:     "",
					Name:         n1.Name.Name,
					Type:         "function",
					IsExportable: n1.Name.IsExported(),
				}
				calls[caller] = nil
			} else {
				var receiverName string

				// Assuming only one receiver for simplicity
				if len(n1.Recv.List) > 0 {
					receiver := n1.Recv.List[0]
					if len(receiver.Names) > 0 {
						receiverName = receiver.Names[0].Name
					}
				}
				caller := Object{
					Path:         importPath,
					Selector:     receiverName,
					Name:         n1.Name.Name,
					Type:         "method",
					IsExportable: n1.Name.IsExported(),
				}
				calls[caller] = nil
			}
		case *ast.File:
		case *ast.Package:
		}
		return true
	})

	return nil
}

func generateDOT() string {
	return ""
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory>\n", os.Args[0])
	}

	dir := os.Args[1]
	err := parse(dir)
	if err != nil {
		log.Fatalf("Failed to parse functions: %v\n", err)
	}

	dotContent := generateDOT()
	dotFileName := "function_call_graph.dot"
	err = os.WriteFile(dotFileName, []byte(dotContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write DOT file: %v\n", err)
	}

	fmt.Printf("Function call graph saved as %s\n", dotFileName)
}
