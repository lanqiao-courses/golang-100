package main

import (
	"fmt"
	"strings"
)

func IsLower(s string) bool {
	return strings.ToLower(s) == s
}
func main() {
	r1 := IsLower("GO") // true
	r2 := IsLower("go") // false
	fmt.Println(r1, r2)
}
