package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe") // HL
	ann := boring("Ann") // HL
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings. // HL
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller. // HL
}

/**
  explain:
	run anonymous_embedded and get result:
		Joe: 0
		Ann: 0
		Joe: 1
		Ann: 1
		Joe: 2
		Ann: 2
		Joe: 3
		Ann: 3
		Joe: 4
		Ann: 4
		You're both boring; I'm leaving.

    in line 10,11, we create 2 routine run 2 job and get 2 channel to stream data
    in line 12 to 15; we get data from 2 this job:
        channel is zero, it's block receive and write until two receive and write together ready
             => order stream data is order use chanel, first is <-joe, after is <-ann, <-joe, <-ann, ....
*/
