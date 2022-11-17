//go:build ignore && OMIT
// +build ignore,OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	a, b := make(chan string), make(chan string)
	go func() { a <- "a" }()
	go func() {
		for {
			b <- "b"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		fmt.Println("nil a")
		a = nil // HL
		fmt.Println("start push to nil chan")
		a <- "ssss"
		fmt.Println("done push to nil chan")
	}()

	go func() {
		a = nil // HL
		fmt.Println("get from nil channel")
		<-a
		fmt.Println("done get from nil channel")
	}()

	go func() {
		for {
			select {
			case s := <-a:
				fmt.Println("have data in nil chan")
				fmt.Println("got", s)
			case s := <-b:
				fmt.Println("got", s)
			}
		}

	}()

	time.Sleep(1000 * time.Second)
}
