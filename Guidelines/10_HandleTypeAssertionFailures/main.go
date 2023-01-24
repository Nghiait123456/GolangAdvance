package main

import (
	"fmt"
)

type I interface {
	walk()
}
type A struct{}

func (a A) walk() {}

type B struct{}

func (b B) walk() {}

func HandleTypeAssertions() {

	fmt.Println("handle type Type Assertion Failures")
	var i I
	t, ok := i.(A)
	if !ok {
		fmt.Println("handle the error gracefully")
	}

	fmt.Println("convert success, t= ", t)

	t1, ok1 := i.(B)
	if !ok1 {
		fmt.Println("handle the error gracefully")
	}

	fmt.Println("convert success, t= ", t1)
}

func main() {
	HandleTypeAssertions()
}
