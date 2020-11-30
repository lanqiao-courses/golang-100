package main

import (
	"fmt"
	"regexp"
)

func ContainsWhiteSpace(str string) bool {
	re := regexp.MustCompile(`\s`)
	return re.MatchString(str)
}
func main() {
	r1 := ContainsWhiteSpace("lorem")       // false
	r2 := ContainsWhiteSpace("lorem ipsum") // true
	fmt.Println(r1, r2)
}
