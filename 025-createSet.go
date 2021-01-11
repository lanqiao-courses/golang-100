package main

import "fmt"

func main() {
	set := make(map[string]struct{})
	fmt.Println(set)

	set["a"] = struct{}{}
	set["b"] = struct{}{}
	set["c"] = struct{}{}
	fmt.Println(set)

	if _, exists := set["a"]; exists {
		fmt.Println("a exists in the set")
	}

	delete(set, "b")
	fmt.Println(set)
	// Output:
	// map[]
	// map[a:{} b:{} c:{}]
	// a exists in the set
	// map[a:{} c:{}]
}
