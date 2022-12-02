package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	IntA int
	IntB int
	Msg  string
	Code int
}

func (e *DivisionError) Error() string {
	return e.Msg
}

type DefaultError struct {
	Msg  string
	Code int
}

func (e *DefaultError) Error() string {
	return e.Msg
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			Msg:  fmt.Sprintf("cannot divide '%d' by zero", a),
			IntA: a,
			IntB: b,
			Code: 0,
		}
	}
	return a / b, nil
}

func main() {
	a, b := 10, 0
	result, err := Divide(a, b)
	if err != nil {
		var divErr *DivisionError
		switch {
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n",
				divErr.IntA, divErr.IntB, divErr.Error())
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
	}

	fmt.Println("test default error")

	errDefault := &DefaultError{
		Code: 10,
		Msg:  "default",
	}

	if errDefault != nil {
		var defaultError *DefaultError
		var divErr *DivisionError
		switch {
		case errors.As(errDefault, &divErr):
			fmt.Printf("have div error \n")
		case errors.As(errDefault, &defaultError):
			fmt.Printf("have default error \n")

		default:
			fmt.Printf("unexpected  error \n")
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}
