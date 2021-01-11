package main

import (
	"fmt"
	"math/rand"
)

func RandIntSliceInRange(min, max, n int) []int {
	arr := make([]int, n)

	for i := range arr {
		arr[i] = rand.Intn(max-min) + min
	}
	return arr
}
func main() {
	r1 := RandIntSliceInRange(12, 35, 10)
	fmt.Println(r1)
}
