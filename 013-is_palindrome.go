package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	v := strings.ToLower(strings.Join(strings.Fields(s), ""))
	for i := range v {
		if v[len(v)-i-1] != v[i] {
			return false
		}
	}
	return true
}
func main() {
	r := IsPalindrome("taco cat")
	fmt.Println(r)
}
