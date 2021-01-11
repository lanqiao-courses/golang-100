package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%[1]d %[1]d\n", 5)
	fmt.Printf("%[2]d %[2]d %[1]d %[1]d\n", 5, 6)
	// Output:
	// 5 5
	// 6 6 5 5
}
