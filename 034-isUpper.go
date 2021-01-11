package main

import (
	"fmt"
	"strings"
)

func IsUpper(s string) bool {
	return strings.ToUpper(s) == s
}
func main() {
	r1 := IsUpper("GO") // true
	r2 := IsUpper("Go") // false
	fmt.Println(r1, r2)
}
