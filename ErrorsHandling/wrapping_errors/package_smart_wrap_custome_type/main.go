package main

import (
	"errors"
	"fmt"

	errorsN "github.com/pkg/errors"
)

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

type MyError1 struct {
	message string
}

func (e *MyError1) Error() string {
	return e.message
}

func warpErrorCustomType() error {
	e0 := &MyError{
		"my Error",
	}
	e1 := errorsN.Wrap(e0, "inner")
	e2 := errorsN.Wrap(e1, "middle")
	e3 := errorsN.Wrap(e2, "high")
	return errorsN.Wrap(e3, "outer")
}

func main() {
	type stackTracer interface {
		StackTrace() errorsN.StackTrace
	}

	errCheck := warpErrorCustomType()
	fmt.Println("errCheck = ", errCheck)
	err, ok := errorsN.Cause(errCheck).(stackTracer)
	if !ok {
		fmt.Println("oops, err does not implement stackTracer")
	}

	fmt.Println("-----------------------------------------------------test error cause -------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------- error Cause dont check after warp -------------------------------------------------------------------")
	switch errT := errorsN.Cause(errCheck).(type) {
	case *MyError:
		fmt.Println("in type MyError: ", errT)
		break
	case *MyError1:
		fmt.Println("in type MyError: ", errT)
		break
	default:
		fmt.Println("default: ", errT)
		// unknown error
	}

	fmt.Println("-----------------------------------------------------other way fast check warp error -------------------------------------------------------------------")
	var MyErr *MyError
	fmt.Println(errors.As(errCheck, &MyErr))

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
		fmt.Println("oops, err does not implement stackTracer")
	}

	fmt.Println("errD: ", errD)

}
