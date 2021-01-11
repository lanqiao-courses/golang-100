package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Input strings: ")
	inp := bufio.NewReader(os.Stdin)

	// res 为输出的Lisp字符串
	// s为输入的字符串
	var res string
	s, err := inp.ReadString('\n')
	if err != nil {
		log.Fatal("wrong input")
	}
	res = strings.Join(strings.Fields(s), "-")
	fmt.Println(res)
}
