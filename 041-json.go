package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     string
	Lessons []string
}

func main() {
	s := Student{
		Name: "John",
		Age:  "17",
		Lessons: []string{
			"Mathematics",
			"Computer science",
			"Philosophy",
		},
	}

	jsonBytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nUgly print:\n%s\n", jsonBytes)

	jsonBytes, err = json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nPretty print:\n%s\n", jsonBytes)
}
