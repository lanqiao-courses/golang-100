package main

import "fmt"

func FrequenciesInt(arr []int) map[int]int {
	m := make(map[int]int)
	for _, v := range arr {
		if f, ok := m[v]; ok {
			m[v] = f + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
func FrequenciesFloat64(arr []float64) map[float64]int {
	m := make(map[float64]int)
	for _, v := range arr {
		if f, ok := m[v]; ok {
			m[v] = f + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
func FrequenciesBool(arr []bool) map[bool]int {
	m := make(map[bool]int)
	for _, v := range arr {
		if f, ok := m[v]; ok {
			m[v] = f + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
func FrequenciesString(arr []string) map[string]int {
	m := make(map[string]int)
	for _, v := range arr {
		if f, ok := m[v]; ok {
			m[v] = f + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
func main() {
	i := []int{1, 1, 3, 1, 4}
	s := []string{"a", "d", "a", "d", "c"}
	r1 := FrequenciesInt(i)
	r2 := FrequenciesString(s)
	fmt.Println(i, "\t", r1, "\n", s, "\t", r2)
}
