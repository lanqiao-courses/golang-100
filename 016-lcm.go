package main

import (
	"fmt"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func LCM(nums ...int) int {
	x := nums[0]
	for _, y := range nums[1:] {
		x = (x * y) / gcd(x, y)
	}
	return x
}
func main() {
	r1 := LCM(12, 7)                // 84
	r2 := LCM([]int{1, 3, 4, 5}...) // 60
	fmt.Println(r1, r2)
}
