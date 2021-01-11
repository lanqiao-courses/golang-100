package main

import (
	"log"
	"os"
)

func main() {
	// You can create as many loggers you want using the log.New method
	// Each logger can write logs in different targets, use different prefix
	// and different settings
	logger1 := log.New(os.Stdout, "logger1: ", log.LstdFlags)
	logger2 := log.New(os.Stdout, "logger2: ", log.LstdFlags|log.Lmicroseconds)
	logger1.Println("Message from logger1")
	logger2.Println("Message from logger2")
	// Output:
	// logger1: 20xx/xx/xx xx:xx:xx Message from logger1
	// logger2: 20xx/xx/xx xx:xx:xx 000000 Message from logger2
}
