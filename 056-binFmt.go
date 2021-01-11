package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 16; x++ {
		fmt.Printf("%b\n", x)
	}
	fmt.Println("fixed length of 8 digits:")
	for x := 0; x < 16; x++ {
		fmt.Printf("%08b\n", x)
	}
	// Output:
	// 0
	// 1
	// ...
}
