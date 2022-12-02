package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")
var DefaultError = errors.New("DefaultError")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, DefaultError):
			fmt.Println("default error")
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
