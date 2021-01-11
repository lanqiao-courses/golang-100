package main

import (
	"fmt"
	"math"
)

func Clamp(n, a, b float64) float64 {
	return math.Max(math.Min(n, math.Max(a, b)), math.Min(a, b))
}
func main() {
	r1 := Clamp(2.0, 3.0, 5.0)   // 3.0
	r2 := Clamp(1.0, -1.0, -5.0) // -1.0
	fmt.Println(r1, r2)
}
