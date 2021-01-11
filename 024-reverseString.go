package main

import (
	"fmt"
)

func ReverseString(s string) string {
	o := make([]rune, len(s))
	for i, c := range s {
		o[len(s)-i-1] = c
	}
	return string(o)
}
func main() {
	r := ReverseString("hello") // "olleh"
	fmt.Println(r)
}
