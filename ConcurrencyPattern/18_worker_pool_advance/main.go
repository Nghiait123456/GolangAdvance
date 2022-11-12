package main

import (
	"fmt"
	"github.com/gammazero/workerpool"
)

func main() {
	fmt.Printf("-------------------start--------------------------------------------------------------------------")
	// if want sync, use only one worker
	// if concurrency, use > 1 worker
	wp := workerpool.New(1)

	for i := 0; i < 100; i++ {
		// always use new variable for one job, if you don't ware race conditions data
		a := i
		wp.Submit(func() {
			fmt.Println("Job : ", a, &a)
		})
	}

	fmt.Printf("-------------------stopWait--------------------------------------------------------------------------")
	wp.StopWait()
	fmt.Printf("-------------------end--------------------------------------------------------------------------")
}
