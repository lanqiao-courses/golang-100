package main

import (
	"errors"
	"fmt"
)

func SplitSliceInChunks(a []int, chuckSize int) ([][]int, error) {
	if chuckSize < 1 {
		return nil, errors.New("chuckSize must be greater that zero")
	}
	chunks := make([][]int, 0, (len(a)+chuckSize-1)/chuckSize)

	for chuckSize < len(a) {
		a, chunks = a[chuckSize:], append(chunks, a[0:chuckSize:chuckSize])
	}
	chunks = append(chunks, a)
	return chunks, nil
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunks, err := SplitSliceInChunks(a, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", chunks)
	// Output:
	// [[0 1 2 3] [4 5 6 7] [8 9]]
}
