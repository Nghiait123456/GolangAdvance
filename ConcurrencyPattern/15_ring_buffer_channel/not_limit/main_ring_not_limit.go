package main

import (
	"fmt"
	"time"
)

// A channel-based ring buffer removes the oldest item when the queue is full
// Ref:
// https://tanzu.vmware.com/content/blog/a-channel-based-ring-buffer-in-go

func NewRingBuffer(inCh, outCh chan int) *ringBuffer {
	return &ringBuffer{
		inCh:  inCh,
		outCh: outCh,
	}
}

// ringBuffer throttle buffer for implement async channel.
type ringBuffer struct {
	inCh  chan int
	outCh chan int
}

func (r *ringBuffer) Run() {
	for v := range r.inCh {
		select {
		case r.outCh <- v:
		}
	}
	close(r.outCh)
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int) // try to change outCh buffer to understand the result
	rb := NewRingBuffer(inCh, outCh)
	go rb.Run()

	go func() {
		for i := 0; i < 10; i++ {
			inCh <- i
		}
	}()

	go func() {
		for res := range outCh {
			fmt.Println(res)
		}
	}()

	time.Sleep(10 * time.Second)

}
