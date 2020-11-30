package main

import "fmt"

func ConcatInt(a, b []int) []int {
	return append(a, b...)
}
func ConcatFloat64(a, b []float64) []float64 {
	return append(a, b...)
}
func ConcatBool(a, b []bool) []bool {
	return append(a, b...)
}
func ConcatStrings(a, b []string) []string {
	return append(a, b...)
}
func main() {
	d1, d2 := []string{"a", "b", "c"}, []string{"d", "e"}
	r := ConcatStrings(d1, d2) // [a b c d e]
	fmt.Println("date1 ", d1, "\ndate2 ", d2, "\nresult: ", r)
}
