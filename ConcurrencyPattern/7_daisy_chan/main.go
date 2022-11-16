package main

import (
	"fmt"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right // get the value from the right and add 1 to it
}

func main() {
	const n = 10000
	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	//routine show leftmost pending until leftmost available write from another
	go func() {
		fmt.Println(<-leftmost, leftmost)
	}()

	fmt.Printf("&left = %v, &right= %v \n, leftmost %v \n", left, right, leftmost)
	for i := 0; i < n; i++ {
		fmt.Println("----------------------------------------start one loop -------------------------------------------")
		right = make(chan int)
		fmt.Printf("left = %v, right = %v in loop \n", left, right)
		go f(left, right)
		left = right
		fmt.Println("----------------------------------------end one loop -------------------------------------------")
	}

	fmt.Printf("out of loop: left = %v, right = %v \n", left, right)

	fmt.Printf("start push data channel to righ,left = %v, right = %v \n", left, right)
	go func(c chan int) { c <- 1 }(right)

	time.Sleep(1000 * time.Second)

}
