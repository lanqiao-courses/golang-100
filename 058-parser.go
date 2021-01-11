package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	src := `
// Calculator package provides methods
// for basic int calculation
package calculator

// Import of fmt package
import "fmt"

// Add adds two integers
func Add(a, b int) int {
	// calculate the result
	result := a + b
	// return the result
	return result
}
`
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, c := range f.Comments {
		fmt.Println(c.Text())
	}
}
