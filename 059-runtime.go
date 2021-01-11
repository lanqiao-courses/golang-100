package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("running program's operating system target: %s\n", runtime.GOOS)
	fmt.Printf("running program's architecture target: %s\n", runtime.GOARCH)
	// Output:
	// running program's operating system target: linux
	// running program's architecture target: amd64
}
