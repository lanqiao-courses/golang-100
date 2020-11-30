package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 输入一个整数，返回一个字符串
	a := strconv.Itoa(1234)
	fmt.Printf("%T %v\n", a, a)
	// Output:
	// "1234"
}
