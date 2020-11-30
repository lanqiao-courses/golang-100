package main

import "fmt"

func FindIndexInt(arr []int, f func(int) bool) int {
	for i, v := range arr {
		if f(v) {
			return i
		}
	}
	return -1
}
func main() {
	r1 := FindIndexInt([]int{1, 1, 2}, func(x int) bool { return x%2 == 0 }) // 2
	fmt.Println(r1)
}
