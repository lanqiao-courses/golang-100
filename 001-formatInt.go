package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 输入int64类型的1234，设置为10进制
	b := strconv.FormatInt(int64(1234), 10)
	fmt.Printf("%T %v\n", b, b)
}
