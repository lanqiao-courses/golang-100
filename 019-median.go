package main

import (
	"fmt"
	"math"
	"sort"
)

func Median(nums ...float64) float64 {
	m, n := int(math.Floor(float64(len(nums))/2.0)),
		nums[:]
	sort.Float64s(n)

	if len(nums)%2 == 0 {
		return (n[m-1] + n[m]) / 2.0
	}
	return n[m]
}
func main() {
	r1 := Median(5.0, 6.0, 50.0, 1.0, -5.0)
	r2 := Median(1.0, 2.0, 4.0, 5.0)
	fmt.Println("中位数:", r1, "中位数:", r2)
}
