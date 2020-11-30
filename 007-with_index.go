package main

import (
	"fmt"
	"reflect"
)

func WithIndex(params ...interface{}) map[int]interface{} {
	arr, m := reflect.ValueOf(params[0]),
		make(map[int]interface{})
	for i := 0; i < arr.Len(); i++ {
		m[i] = arr.Index(i).Interface()
	}
	return m
}
func main() {
	r := WithIndex([]int{4, 3, 2, 1}) // [0:4 1:3 2:2 3:1]
	fmt.Println(r)
}
