package main

import (
	"fmt"
	"math"
)

func MaxOf(nums ...float64) float64 {
	max := math.Inf(-1)
	for _, num := range nums {
		max = math.Max(num, max)
	}
	return max
}
func main() {
	r := MaxOf(3.0, 4.0, 2.0)
	fmt.Println(r)
}
