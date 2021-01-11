package main

import (
	"fmt"
)

func TruncateString(s string, l int) string {
	r := s
	if len(s) > l {
		if l > 3 {
			l -= 3
		}
		r = s[0:l] + "..."
	}
	return r
}
func main() {
	r := TruncateString("boomerang", 7) // "boom..."
	fmt.Println(r)
}
