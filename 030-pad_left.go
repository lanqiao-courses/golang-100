package main

import (
	"fmt"
	"strconv"
)

func PadLeft(s string, l int) string {
	f := "%" + strconv.Itoa(l) + "v"
	return fmt.Sprintf(f, s)
}
func main() {
	r := PadLeft("go", 8) // "      go"
	fmt.Println(r)
}
