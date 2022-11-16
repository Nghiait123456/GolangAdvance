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
	go func() { b <- "b" }()
	go func() {
		a = nil // HL
		fmt.Println("start push to nil chan")
		a <- "ssss"
		fmt.Println("done push to nil chan")
	}()

	if rand.Intn(2) == 0 {
		a = nil // HL
		fmt.Println("nil a")
	} else {
		b = nil // HL
		fmt.Println("nil b")
	}
	go func() {
		select {
		case s := <-a:
			fmt.Println("have data in nil chan")
			fmt.Println("got", s)
		case s := <-b:
			fmt.Println("got", s)
		}
	}()

	time.Sleep(1000 * time.Second)
}
