package main

import (
	"fmt"

	errorsN "github.com/pkg/errors"
)

func warpError() error {
	e1 := errorsN.New("error")
	e2 := errorsN.Wrap(e1, "inner")
	e3 := errorsN.Wrap(e2, "middle")
	e4 := errorsN.Wrap(e3, "high")
	return errorsN.Wrap(e4, "outer")
}

func main() {
	type stackTracer interface {
		StackTrace() errorsN.StackTrace
	}

	//errDefault := errors.New("test")
	errCheck := warpError()
	fmt.Println("errCheck = ", errCheck)
	err, ok := errorsN.Cause(errCheck).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Println("start stack trade")
	fmt.Printf("%+v \n", st) // top two frames
	//fmt.Errorf("math: square root of negative number %g", f)
	// Example output:
	// github.com/pkg/errors_test.fn
	//	/home/dfc/src/github.com/pkg/errors/example_test.go:47
	// github.com/pkg/errors_test.Example_stackTrace
	//	/home/dfc/src/github.com/pkg/errors/example_test.go:127

	fmt.Println("end")
}
