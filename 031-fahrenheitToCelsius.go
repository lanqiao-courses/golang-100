package main

import (
	"fmt"
)

func FahrenheitToCelsius(d float64) float64 {
	return (d - 32.0) * 5.0 / 9.0
}
func main() {
	r := FahrenheitToCelsius(451.0)
	fmt.Println(r)
}
