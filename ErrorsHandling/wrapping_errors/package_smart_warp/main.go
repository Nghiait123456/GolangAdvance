package main

import (
	"errors"
	"fmt"

	errorsN "github.com/pkg/errors"
)

func warpError() error {
	e0 := errors.New("it error0 ")
	e1 := errorsN.New("have error: " + e0.Error())
	e2 := errorsN.Wrap(e1, "inner")
	e3 := errorsN.Wrap(e2, "middle")
	e4 := errorsN.Wrap(e3, "high")
	return errorsN.Wrap(e4, "outer")
}

func main() {
	type stackTracer interface {
		StackTrace() errorsN.StackTrace
	}

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

	fmt.Println("test with error default")
	errDefault := errors.New("test")
	errD, ok := errorsN.Cause(errDefault).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	fmt.Println("errD: ", errD)

}
