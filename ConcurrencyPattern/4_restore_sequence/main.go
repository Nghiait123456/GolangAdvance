package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

// the boring function return a channel to communicate with it.
func boring(msg string) <-chan Message { // <-chan Message means receives-only channel of Message.
	c := make(chan Message)
	waitForIt := make(chan bool) // share between all messages
	go func() {                  // we launch goroutine inside a function.
		for i := 0; ; i++ {
			fmt.Println("send data to channel")
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}

			// every time the goroutine send message.
			// This code waits until the value to be received.
			<-waitForIt
		}

	}()
	return c // return a channel to caller.
}

func main() {
	// merge 3 channels into 1 channel
	c := fanIn(boring("Joe_1"), boring("Joe_2"), boring("Joe_3"))

	for i := 0; i < 5; i++ {
		msg := <-c // wait to receive message
		fmt.Println(msg.str)
		msg.wait <- true // main goroutine allows the boring goroutine to send next value to message channel.
	}
	time.Sleep(10 * time.Second)
	fmt.Println("You're both boring. I'm leaving")
}

// main: goroutine                                          boring: goroutine
//    |                                                           |
//    |                                                           |
// wait for receiving msg from channel c                    c <- Message{} // send message
//   <-c                                                          |
//    |                                                           |
//    |                                                     <-waitForIt // wait for wake up signal
// send value to channel                                          |
// hey, boring. You can send next value to me                     |
//   wait <-true                                                  |
///                            REPEAT THE PROCESS

//imagine, we have pool save flag:

//|       |
//|       |
//|       |
//|       |
//| flat3 |
//| flat2 |
//|_flat1_|

/**
  explain:
     run code and get result:
		send data to channel
		send data to channel
		send data to channel
		Joe_3 0
		Joe_1 0
		Joe_2 0
		send data to channel
		Joe_2 1
		send data to channel
		Joe_1 1
		send data to channel
		send data to channel
		send data to channel

    we want get 5 times data form stream, after get full 5 times, we sleep for test routine is pending
    in line 60, for this time sleep, routine run after line <-waitForIt
    never have routine run continue, all routine pending until we want another data.
    we have control sequence data stream from flag line 29: waitForIt := make(chan bool)

*/
