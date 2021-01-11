package main

import (
	"fmt"
	"strings"
)

func Decapitalize(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}
func main() {
	r := Decapitalize("Boomerang") // "boomerang"
	fmt.Println(r)
}
