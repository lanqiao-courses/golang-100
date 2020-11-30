package main

import (
	"fmt"
)

func IntRange(f, t, s int) []int {
	arr := make([]int, (t-f+1)/s)
	for i := range arr {
		arr[i] = i*s + f
	}
	return arr
}
func main() {
	r1 := IntRange(0, 9, 2)
	fmt.Println(r1)
}
