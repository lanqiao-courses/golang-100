package main

import (
	"fmt"
	"time"
)

func f() {
	// first we calculate the exact time the function
	// started the execution
	start := time.Now()
	// Then we calculate the execution time duration
	// inside a deferred function, since it will be execute after f() returns
	defer func(start time.Time) {
		dur := time.Since(start)
		fmt.Printf("f() took %0.0fs to execute", dur.Seconds())
	}(start)
	// pause the execution for 2 seconds
	time.Sleep(2 * time.Second)
}

func main() {
	f()
}
