package main

import (
	"fmt"
)

func IsOdd(n int) bool {
	return n%2 == 1
}
func main() {
	r := IsOdd(3) // true
	fmt.Println(r)
}
