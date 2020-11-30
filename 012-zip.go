package main

import (
	"fmt"
	"reflect"
)

func Zip(params ...interface{}) [][]interface{} {
	l := 0
	for i := range params {
		arr := reflect.ValueOf(params[i])
		if l < arr.Len() {
			l = arr.Len()
		}
	}
	r := make([][]interface{}, l)

	for i := 0; i < l; i++ {
		r[i] = make([]interface{}, 0)
		for j := range params {
			v := reflect.ValueOf(params[j])
			if v.Len() > i {
				r[i] = append(r[i], v.Index(i).Interface())
			}
		}
	}
	return r
}
func main() {
	s := []string{"a", "b"}
	i := []int{1, 2}
	b := []bool{true, false}
	r1 := Zip(s, i, b) // [[a 1 true] [b 2 false]]
	fmt.Println(r1)
}
