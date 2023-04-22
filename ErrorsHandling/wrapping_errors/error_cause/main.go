package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return errors.Wrap(err, "failed to do something")
	}
	return nil
}

func doSomethingElse() error {
	err := doSomethingMore()
	if err != nil {
		return errors.Wrap(err, "failed to do something else")
	}
	return nil
}

func doSomethingMore() error {
	return &MyError{"something went wrong"}
}

func main() {
	err := doSomething()
	if err != nil {
		switch err := errors.Cause(err).(type) {
		case *MyError:
			fmt.Println("Got a MyError:", err.Error())
		default:
			fmt.Println("Unknown error:", err.Error())
		}
	}
}
