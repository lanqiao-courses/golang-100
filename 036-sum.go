package main

import "fmt"

func Sum(nums ...float64) float64 {
	sum := float64(0)
	for _, num := range nums {
		sum += num
	}
	return sum
}
func main() {
	r1 := Sum(1.0, 4.0)                    // 5.0
	r2 := Sum([]float64{1.0, 2.0, 3.0}...) // 6.0
	fmt.Println(r1, r2)
}
