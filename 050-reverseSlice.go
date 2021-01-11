package main

import (
	"fmt"
)

func ReverseSlice(s []int) []int {
	if len(s) < 2 {
		return s
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	d := []int{1, 2, 3, 4, 5}
	r := ReverseSlice(d)
	fmt.Println(r)
}
