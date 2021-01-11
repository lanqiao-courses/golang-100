package main

import "fmt"

func main() {
	// The map keys are a unique unordered collection.
	// So if you just ignore the values, a map can be viewed as a set
	// The most space optimized way to do this is to use the empty struct
	// as the map value type, since the empty struct (a struct with no fields)
	// occupies zero bytes of storage
	set := make(map[string]struct{})
	fmt.Println(set)
	// Add new elements in the set
	set["a"] = struct{}{}
	set["b"] = struct{}{}
	set["c"] = struct{}{}
	fmt.Println(set)

	// Check if an element exists in the set
	if _, exists := set["a"]; exists {
		fmt.Println("a exists in the set")
	}
	// Delete an element from the set
	delete(set, "b")
	fmt.Println(set)
	// Output:
	// map[]
	// map[a:{} b:{} c:{}]
	// a exists in the set
	// map[a:{} c:{}]
}
