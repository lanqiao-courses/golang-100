package main

import (
	"fmt"
	"reflect"
)

func XProduct(params ...interface{}) [][]interface{} {
	a, b := reflect.ValueOf(params[0]), reflect.ValueOf(params[1])
	l := a.Len() * b.Len()
	r := make([][]interface{}, l)

	for i := 0; i < l; i++ {
		r[i] = []interface{}{
			a.Index(i % a.Len()).Interface(),
			b.Index((i / a.Len()) % b.Len()).Interface(),
		}
	}
	return r

}
func main() {
	r := XProduct([]int{1, 2}, []string{"a", "b"}) // [[1 a] [2 a] [1 b] [2 b]]
	fmt.Println(r)
}