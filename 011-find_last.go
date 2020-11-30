package main

import "fmt"

func FindLastInt(arr []int, f func(int) bool) (int, error) {
	for i := len(arr) - 1; i >= 0; i-- {
		if f(arr[i]) {
			return arr[i], nil
		}
	}
	return 0, fmt.Errorf("No matches found")
}
func FindLastFloat64(arr []float64, f func(float64) bool) (float64, error) {
	for i := len(arr) - 1; i >= 0; i-- {
		if f(arr[i]) {
			return arr[i], nil
		}
	}
	return 0.0, fmt.Errorf("No matches found")
}
func FindLastBool(arr []bool, f func(bool) bool) (bool, error) {
	for i := len(arr) - 1; i >= 0; i-- {
		if f(arr[i]) {
			return arr[i], nil
		}
	}
	return false, fmt.Errorf("No matches found")
}
func FindLastString(arr []string, f func(string) bool) (string, error) {
	for i := len(arr) - 1; i >= 0; i-- {
		if f(arr[i]) {
			return arr[i], nil
		}
	}
	return "", fmt.Errorf("No matches found")
}
func main() {
	r1, ch1 := FindLastInt([]int{1, 1, 2}, func(x int) bool { return x%2 == 0 }) // 2 nil
	fmt.Println(r1, ch1)
}
