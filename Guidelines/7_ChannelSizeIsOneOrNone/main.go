package main

import (
	"fmt"
	"sync"
)

type TestChannel struct {
	cZero       chan int
	cOne        chan int
	cGreaterOne chan int
}

func main() {
	fmt.Println("channel zero, one, greater one and action with it")
	var wg sync.WaitGroup
	test := &TestChannel{
		cZero:       make(chan int),
		cOne:        make(chan int, 1),
		cGreaterOne: make(chan int, 64),
	}
	fmt.Println(test)

	// chan Zero block when empty chan
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("read from channel Zero with empty")
	//	read := <-24_cleary_example_smart_fetch.cZero
	//	fmt.Println("read= ", read)
	//	time.Sleep(20 * time.Second)
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	time.Sleep(10 * time.Second)
	//	fmt.Println("push first value to channel Zero")
	//	24_cleary_example_smart_fetch.cZero <- 1
	//	fmt.Println("done push first value to channel Zero")
	//	24_cleary_example_smart_fetch.cZero <- 1
	//	24_cleary_example_smart_fetch.cZero <- 1
	//	fmt.Println("done push  to channel Zero")
	//}()
	/////////////////////////////////////////////////////////////////////////////////////////////

	//chane One not block when empty Chan or receive and sender not ready
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	//time.Sleep(10 * time.Second)
	//	fmt.Println("read first from channel One with not empty")
	//	24_cleary_example_smart_fetch.cOne <- 1
	//	read := <-24_cleary_example_smart_fetch.cOne
	//	fmt.Println("read= ", read)
	//
	//	fmt.Println("read second from channel One with empty")
	//	read = <-24_cleary_example_smart_fetch.cOne
	//	fmt.Println("read= ", read)
	//	time.Sleep(20 * time.Second)
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("push first value to channel One")
	//	24_cleary_example_smart_fetch.cOne <- 1
	//	fmt.Println("done push first value to channel One")
	//	24_cleary_example_smart_fetch.cOne <- 1
	//	fmt.Println("done push seconds value to channel One")
	//	24_cleary_example_smart_fetch.cOne <- 1
	//	fmt.Println("done push  to channel One")
	//}()

	/////////////////////////////////////////////////////////////////////////////////////

	//chane greater One same channel one
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	//time.Sleep(10 * time.Second)
	//	fmt.Println("read first from channel greater than one  with not empty")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	read := <-24_cleary_example_smart_fetch.cGreaterOne
	//	fmt.Println("read= ", read)
	//
	//	fmt.Println("read second from channel  greater than one  with empty")
	//	read = <-24_cleary_example_smart_fetch.cGreaterOne
	//	fmt.Println("read= ", read)
	//	time.Sleep(20 * time.Second)
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("push first value to channel  greater than one ")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	fmt.Println("done push first value to channel greater than one ")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	fmt.Println("done push seconds value to channel  greater than one ")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	fmt.Println("done push three value to channel  greater than one ")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	fmt.Println("done push  to channel greater than one ")
	//}()

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//problem when use channel greater than 1
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 64; i++ {
	//		24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	}
	//
	//	fmt.Println("done push 64 times to channel")
	//	fmt.Println("start push 65 times to channel")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	fmt.Println("done push 65 times to channel")
	//	fmt.Println("start push 66 times to channel")
	//	24_cleary_example_smart_fetch.cGreaterOne <- 1
	//	fmt.Println("done push 66 times to channel")
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	value := <-24_cleary_example_smart_fetch.cGreaterOne
	//
	//	fmt.Println("get value = ", value)
	//	time.Sleep(100 * time.Second)
	//}()

	wg.Wait()
	fmt.Println("Main: Completed")
}

// have many problem when use channel, it's not simple same syntax "chan" :D, it's not simple explain it,
// but you can follow me step by step to understand it.
// we will explain this:
/**
  1) deference chanZero, chanOne, chanGreaterOne

  +) chanZero block Sends and Receive, available when send or receive ready
     un cmt line 23 ->to 44 : we have two routine Interactive with channel zero ( 24_cleary_example_smart_fetch.cZero )
     run and view result:
		read from channel Zero with empty
		push first value to channel Zero
		done push first value to channel Zero
		read=  1
     after sleep 3s, routine first read from channel, but it's blocked when chanel empty
     inline37: after 10 second, routine seconds write first value to channel, after write success, routine 1 will is not blocked and read success
     ==> channel zero blocked read until sends avaible
     in line  41,42, have two cmd write to channel : 24_cleary_example_smart_fetch.cZero <- 1, but it's blocked why not receive ready
     edit line 27 to : time.Sleep(10 * time.Second)  and edit line 37 to : 		time.Sleep(3 * time.Second), run again and get result:
		push first value to channel Zero
		read from channel Zero with empty
		read=  1
		done push first value to channel Zero
     ==> same before, send to channel blocked until receive ready
     ( please cmt all code 23 -> 44)
     ==========================================> SUMMARY ============================================================================================
                 zero channel blocked receive and write,  communication succeeds only when the sender and receiver are both ready.
     ================================================================================================================================================


    +) chanOne send succeeds without blocking if the channel is not full and receive succeeds without blocking if the buffer is not empty
       don't block if receive and write not ready in this currency time

     un cmt line 47 to 73, run and view result :
	read first from channel One with empty
		read=  1
		read second from channel One with empty
		push first value to channel One
		done push first value to channel One
		done push seconds value to channel One
		read from channel Zero with empty
		read=  1
		push first value to channel Zero
		done push first value to channel Zero
		read=  1
    same case one, we have 2 routine action to routine one 24_cleary_example_smart_fetch.cOne
    int line 55, fmt.Println("read= ", read), read not blocked when write not ready, it's only need channel not empty
    inline 71: fmt.Println("done push seconds value to channel One"), write to channel one not blocked when receive not ready, it's only need channel not full
    (please cmt all code 47 -> 73)
   ===================================> SUMMARY =============================================================================================================
        with channel one :  send succeeds without blocking if the channel is not full and receive succeeds without blocking if the buffer is not empty.
    ==============================================================================================================================================================


   +) channel greater than one same channel one
    un cmt line 78 to 106 and run, view result :
		read first from channel greater than one  with not empty
		read=  1
		read second from channel  greater than one  with empty
		push first value to channel  greater than one
		done push first value to channel greater than one
		done push seconds value to channel  greater than one
		done push three value to channel  greater than one
		done push  to channel greater than one
		read=  1
  (please cmt all code 78 -> 106)
  ======================================> SUMMARY ===============================================================================================================
                                                    channel greater than one same channel one
  ================================================================================================================================================================


  +) why channel greater than one is not good
     un cmt line 111 to 133 and run it, view result:
		 done push 64 times to channel
         start push 65 times to channel
	     done push 65 times to channel
	     start push 66 times to channel
	     get value =  1
     => we have plus times write to channel size 64 to full, write is blocked.
     q/a: why channel greater one is bad when use it.
         please remember the point of problem; if channel size >= 1, channel will block if empty or full
         if channel size = 1, clearly target when use it :  only one data write and push and don't need two:  write and read ready
         if channel size > 1, target when use it same with target when use unBuffer( zero channel),  but have one reason optimal memory.
                              but in concurrency, not simple define one size for many concurrency action to channel, it's not simple.
                              if you calculate wrong, routine is blocked a few time, it's very hard debug
                              i don't say never use channel greater than 1, but careful when use it
                              only use it when you control all problem when blocked (when full size)
      (please cmt all code 111 -> 133)
======================================> SUMMARY ===============================================================================================================
                      careful when use channel greater than one, only use when handle all problem it ( blocked why full)
================================================================================================================================================================

                                  ===> channel is simple syntax but not simple run good ;)
*/
