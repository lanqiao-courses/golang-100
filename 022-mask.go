package main

import (
	"fmt"
	"strings"
)

func Mask(cc string, n int, m rune) string {
	return strings.Repeat(string(m), len(cc)-n) + cc[len(cc)-n:]
}
func main() {
	r := Mask("1234567890", 4, '*') // "******7890"
	fmt.Println(r)
}
