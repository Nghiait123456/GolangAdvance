package _10_UnnecessaryElse

import "fmt"

/**
f a variable is set in both branches of an if, it can be replaced with a single if.
*/

func bad() {
	var a int
	var b bool
	if b {
		a = 100
	} else {
		a = 10
	}

	fmt.Println(a, b)
}

func good() {
	var b bool
	a := 10
	if b {
		a = 100
	}

	fmt.Println(a, b)
}
