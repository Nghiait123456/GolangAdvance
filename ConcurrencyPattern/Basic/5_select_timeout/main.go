package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the boring function return a channel to communicate with it.
func boring(msg string) <-chan string { // <-chan string means receives-only channel of string.
	c := make(chan string)
	go func() { // we launch goroutine inside a function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}

	}()
	return c // return a channel to caller.
}

func main() {
	c := boring("Joe")

	// timeout for the whole conversation
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}

/**
  explain it:
    we run and get result:
		Joe 0
		Joe 1
		Joe 2
		Joe 3
		Joe 4
		Joe 5
		Joe 6
		Joe 7
		Joe 8
		You talk too much.
	we run 2 job, 1 job generate data and push to channel, 1 job get data and show
	in line 26: 	timeout := time.After(5 * time.Second) return one channel timer
    in line 27 to 34, for {
		select {
		case
      => each all signal from many channel, when case <-timeout: have data, ==> has timeout => will action and return job receive.
*/
