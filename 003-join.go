package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Join(params ...interface{}) string {
	arr, sp := reflect.ValueOf(params[0]),
		reflect.ValueOf(params[1]).String()
	ars := make([]string, arr.Len())

	for i := 0; i < arr.Len(); i++ {
		ars[i] = fmt.Sprintf("%v", arr.Index(i))
	}

	return strings.Join(ars, sp)

}
func main() {
	w := Join([]int{1, 2, 3}, ".")
	fmt.Printf("%T %v", w, w)
}
