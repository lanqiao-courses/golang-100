package main

import (
	"fmt"
)

func IsPowerOf2(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
func main() {
	r1 := IsPowerOf2(0) // false
	r2 := IsPowerOf2(1) // true
	r3 := IsPowerOf2(8) // true
	fmt.Println(r1, r2, r3)
}
