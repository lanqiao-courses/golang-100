package main

import (
	"fmt"
	"time"
)

func today() {
	now := time.Now()
	fmt.Println("Today is", now.Weekday())
	if now.Weekday() == time.Tuesday {
		fmt.Println("The day is Tuesday")
	}
}
func main() {
	today()
}
