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

var MySentinelEr = errorsN.New("MySentinelEr")

func warpErrorSentinelTypeDefault() error {
	e0 := MySentinelEr
	e1 := fmt.Errorf("inner %v", e0)
	return fmt.Errorf("middle %v", e1)
}

func main() {
	type stackTracer interface {
		StackTrace() errorsN.StackTrace
	}

	errCheck := warpErrorSentinelTypeDefault()
	fmt.Println("errCheck = ", errCheck)
	err, ok := errorsN.Cause(errCheck).(stackTracer)
	if !ok {
		fmt.Println("oops, err does not implement stackTracer")
	}

	fmt.Println("-----------------------------------------------------if you use fmt.Error, dont have way check type error, this is weakness -------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------test error cause -------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------- error Cause dont check sentinel error, it just check custom error -------------------------------------------------------------------")
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

	fmt.Println("-----------------------------------------------------other way fast check warp sentinel error -------------------------------------------------------------------")
	fmt.Println(errors.Is(errCheck, MySentinelEr))

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
