package main

import (
	"fmt"
	"math"
)

func Rads(d float64) float64 {
	return d * math.Pi / 180.0
}
func main() {
	r := Rads(90.0) // ~1.5708
	fmt.Println(r)
}
