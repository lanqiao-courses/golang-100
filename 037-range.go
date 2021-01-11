package main

import (
	"fmt"
	"math"
)

func IsInRange(n, a, b float64) bool {
	s, e := math.Min(a, b), math.Max(a, b)
	return n >= s && n < e
}
func main() {
	r1 := IsInRange(3, 2, 5) // true
	r2 := IsInRange(2, 3, 5) // false
	fmt.Println(r1, r2)
}
