package main

import "fmt"

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func GCD(nums ...int) int {
	r := nums[0]
	for _, num := range nums[1:] {
		r = gcd(r, num)
	}
	return r
}
func main() {
	r1 := GCD(8, 36)               // 4
	r2 := GCD([]int{11, 8, 32}...) // 1
	fmt.Println(r1, r2)
}
