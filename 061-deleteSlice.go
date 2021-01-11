package main

import (
	"errors"
	"fmt"
)

// DeleteSlice removes an element at a specific index of an slice.
// An error is return in case the index is out of bounds or if the slice is nil or empty.
func DeleteSlice(s []int, i int) ([]int, error) {
	if len(s) == 0 {
		return nil, errors.New("Cannot delete an element from a nil or empty slice")
	}
	if i < 0 || i > len(s)-1 {
		return nil, errors.New("Index out of bounds")
	}

	return append(s[:i], s[i+1:]...), nil
}

func main() {
	d, i := []int{1, 2, 3, 4, 5}, 2
	r, _ := DeleteSlice(d, i)
	fmt.Println("result:", r)
}
