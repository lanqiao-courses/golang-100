package main

import (
	"fmt"
)

func CelsiusToFahrenheit(d float64) float64 {
	return 1.8*d + 32.0
}
func main() {
	r := CelsiusToFahrenheit(-273.15) // -459.669
	fmt.Println(r)
}
