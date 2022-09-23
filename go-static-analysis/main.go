package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	names := getInterfaceNames("input.go")
	fmt.Println(names)
}

func getInterfaceNames(filename string) []string {
	names := make([]string, 0)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	// search Interface and print Interface name
	ast.Inspect(f, func(n ast.Node) bool {
		// fmt.Printf("%#v\n", n)

		a, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if _, ok = a.Type.(*ast.InterfaceType); !ok {
			return true
		}

		fmt.Printf("%#v\n", a)
		fmt.Printf("%s\n", a.Name)

		names = append(names, a.Name.Name)
		return true
	})

	return names
}
