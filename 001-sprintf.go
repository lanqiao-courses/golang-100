package main

import (
	"fmt"
)

func main() {
	// sprintf函数返回转换后的字符类型
	c := fmt.Sprintf("%d", 1234)
	fmt.Printf("%T %v\n", c, c)
}
