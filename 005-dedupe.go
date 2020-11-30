package main

import (
	"fmt"
)

func DedupeInts(arr []int) []int {
	m, uniq := make(map[int]bool), make([]int, 0)
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v], uniq = true, append(uniq, v)
		}
	}
	return uniq
}
func main() {
	date := []int{1, 2, 1, 2, 3, 3, 4}
	res := DedupeInts(date)
	fmt.Printf("date: %v\nres: %v", date, res)
}
