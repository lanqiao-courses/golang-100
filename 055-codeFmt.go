package main

import (
	"fmt"
	"go/format"
	"log"
)

func main() {
	unformatted := `
package main
       import "fmt"

func  main(   )  {
    x :=    12
fmt.Printf(   "%d",   x  )
	}
`
	formatted, err := format.Source([]byte(unformatted))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(formatted))
	// Output:
	// package main
	//
	// import "fmt"
	//
	// func main() {
	//	x := 12
	//	fmt.Printf("%d", x)
	// }
}
