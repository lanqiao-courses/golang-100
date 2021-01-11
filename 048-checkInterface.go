package main

import (
	"fmt"
)

type Checker interface {
	Check() bool
}

type P struct{}

func (p *P) Check() bool {
	return true
}

type V struct{}

func (v V) Check() bool {
	return true
}

type N struct{}

func main() {
	var pp interface{} = &P{}
	if _, ok := pp.(Checker); ok {
		fmt.Println("*P implements the checker interface")
	} else {
		fmt.Println("*P does not implement the checker interface")
	}

	var pv interface{} = P{}

	if _, ok := pv.(Checker); ok {
		fmt.Println("P implements the checker interface")
	} else {
		fmt.Println("P does not implement the checker interface")
	}

	var vp interface{} = &V{}

	if _, ok := vp.(Checker); ok {
		fmt.Println("*V implements the checker interface")
	} else {
		fmt.Println("*V does not implement the checker interface")
	}

	// V implements the Checker interface using value receiver
	// so V satisfies the interface
	var vv interface{} = V{}

	if _, ok := vv.(Checker); ok {
		fmt.Println("V implements the checker interface")
	} else {
		fmt.Println("V does not implement the checker interface")
	}

	// N does not implement the Checker interface at all
	// so *N does not satisfy the interface
	var np interface{} = &N{}

	if _, ok := np.(Checker); ok {
		fmt.Println("*N implements the checker interface")
	} else {
		fmt.Println("*N does not implement the checker interface")
	}

	// N does not implement the Checker interface at all
	// so N does not satisfy the interface
	var nv interface{} = N{}

	if _, ok := nv.(Checker); ok {
		fmt.Println("N implements the checker interface")
	} else {
		fmt.Println("N does not implement the checker interface")
	}
}
