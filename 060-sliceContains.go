package main

import (
	"fmt"
)

func SliceContains(s []int, x int) bool {
	if len(s) == 0 {
		return false
	}
	for k := range s {
		if s[k] == x {
			return true
		}
	}
	return false
}

func main() {
	tests := []struct {
		name string
		s    []int
		x    int
		want bool
	}{
		{
			name: "nil",
			s:    nil,
			x:    1,
		},
		{
			name: "empty",
			s:    []int{},
			x:    1,
		},
		{
			name: "value exists",
			s:    []int{2, 1},
			x:    1,
		},
		{
			name: "value does not exist",
			s:    []int{2, 1},
			x:    3,
		},
	}
	for _, tt := range tests {
		r := SliceContains(tt.s, tt.x)
		fmt.Println(tt.s, tt.x, r)
	}
}
