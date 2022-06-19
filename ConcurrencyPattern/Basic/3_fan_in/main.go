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
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c // return a channel to caller.
}

// <-chan string only get the receive value
// fanIn spawns 2 goroutines to reads the value from 2 channels
// then it sends to value to result channel( `c` channel)
func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for { // infinite loop to read value from channel.
			v1 := <-c1 // read value from c2. This line will wait when receiving value.
			c <- v1
		}
	}()
	go func() {
		for {
			c <- <-c2 // read value from c2 and send it to c
		}
	}()
	return c
}

func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs { // spawn channel based on the number of input channel

		go func(cv <-chan string) { // cv is a channel value
			for {
				c <- <-cv
			}
		}(ci) // send each channel to

	}
	return c
}

func main() {
	// merge 2 channels into 1 channel
	// c := fanIn(boring("Joe"), boring("Ahn"))
	c := fanInSimple(boring("Joe"), boring("Ahn"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // now we can read from 1 channel
	}
	fmt.Println("You're both boring. I'm leaving")
}

/**
  Explain:
	   you run code and get result:
			Ahn 0
			Joe 0
			Ahn 1
			Joe 1
			Ahn 2
			You're both boring. I'm leaving
	   +) in func fanInSimple(cs ...<-chan string):
			  in line 42, we want get all channel need fan, we create one routine is point share global data
			  in line 43 to 51, for every one channel need fan, we create one routine, this routine loop forever for one job: move data form  channel to global channel
			  in line 52, we run global channel for endpoint
	   +) in this pattern, we don't care order stream data ( why problem don't care it). We care we need all data form many job, we always have one endpoint(global channel) for read data.
   Note:
   		this pattern is high performance?
           => maybe, most of case, it's good for speed
           if you want high performance, we share solution fit it in next part
           it's never use mutex(inverse channel), it maybe is ring buffer
*/
