package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// RingBuffer Structure
type RingBuffer struct {
	Size      uint64        // Size of the Ringbuffer
	Container []interface{} // Array container of objects
	Reader    uint64        // Reader position
	Writer    uint64        // Writer Position
	muLockW   sync.Mutex
	muLockR   sync.Mutex
}

const (
	MAX_UINT_64         = 18446744073709551615
	TOLERANCE           = 100000000000000000
	MAX_LIMIT_OVER_FLOW = MAX_UINT_64 - TOLERANCE
)

// Create a new RingBuffer of initial size "size"
// Returns a pointer to the new RingBuffer
func NewRingBuffer(size uint64) *RingBuffer {
	rb := new(RingBuffer)
	rb.Size = size
	rb.Container = make([]interface{}, size)
	rb.Reader = 0
	rb.Writer = 0
	return rb
}

// Write object into the RingBuffer
func (r *RingBuffer) ReduceWhenOver(value uint64) uint64 {
	maxMultiples := value / r.Size
	return value - r.Size*maxMultiples
}

// Write object into the RingBuffer
func (r *RingBuffer) Write(v interface{}) {
	future := atomic.AddUint64(&r.Writer, 1)
	if future >= MAX_LIMIT_OVER_FLOW {
		r.muLockW.Lock()

		//replace new value
		reduce := r.ReduceWhenOver(future)
		replace := future - reduce
		atomic.StoreUint64(&r.Writer, replace)

		//write  value
		nowIndex := replace
		r.Container[(nowIndex-1)%r.Size] = v

		r.muLockW.Unlock()
		return
	}

	r.Container[(future-1)%r.Size] = v
}

// Read single object from the RingBuffer
func (r *RingBuffer) Read() interface{} {
	future := atomic.AddUint64(&r.Reader, 1)
	if future >= r.LatestWriteIndex() {
		r.muLockR.Lock()
		atomic.StoreUint64(&r.Reader, future-1)
		r.muLockR.Unlock()

		return nil
	}

	if future >= MAX_LIMIT_OVER_FLOW {
		r.muLockW.Lock()

		//replace new value
		reduce := r.ReduceWhenOver(future)
		replace := future - reduce
		atomic.StoreUint64(&r.Reader, replace)

		//write  value
		FutureNew := replace
		r.muLockW.Unlock()

		return r.Container[(FutureNew-1)%r.Size]
	}

	return r.Container[(future-1)%r.Size]
}

func (r *RingBuffer) LatestWriteIndex() uint64 {
	return atomic.LoadUint64(&r.Writer)
}

func (r *RingBuffer) LatestReadIndex() interface{} {
	return atomic.LoadUint64(&r.Reader)
}

func testWriteAndRead() {
	var wg sync.WaitGroup
	rb := NewRingBuffer(10)

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for {
			fmt.Println("start write 1")
			rb.Write(1)
			//time.Sleep(5 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for {
			fmt.Println("start read")
			data := rb.Read()
			if data != 1 && data != nil {
				fmt.Println("have errors, data = ", data)
				panic(" have errors in code")
			}
			//time.Sleep(5 * time.Second)
		}
	}()

	wg.Wait()
}

func testReadWaitWrite(totalCountTest uint64) {
	var wg sync.WaitGroup
	rb := NewRingBuffer(totalCountTest)
	var i uint64
	for i = 0; i < totalCountTest; i++ {
		fmt.Println("start write 1")
		rb.Write(1)
		//time.Sleep(5 * time.Second)
	}

	wg.Add(1)
	var count uint64
	count = 1
	go func() {
		defer func() {
			wg.Done()
		}()

		for {
			data := rb.Read()
			if data != 1 && data != nil {
				fmt.Println("have errors, data = ", data)
				panic(" have errors in code")
			}

			if data == 1 {
				fmt.Println("read done one data, count = ", count)
				atomic.AddUint64(&count, 1)
			}

			if data == nil {
				fmt.Println("empty ring, count = ", count)
			}

			if count == totalCountTest {
				fmt.Println("read success")
			}

			if count > totalCountTest {
				errM := fmt.Sprintf("read process error, count = ", count)
				fmt.Println(errM)
				panic(errM)
			}
		}
	}()

	wg.Wait()
}

func benchMarkRing(countW int) {
	var wg sync.WaitGroup
	rb := NewRingBuffer(500000)
	var totalCountWrite uint64
	var totalCountRead uint64

	var doneInitWorkerW, doneInitWorkerR bool

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 0; i < countW; i++ {
			fmt.Println("start one worker write")
			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()

				//fmt.Println("start loop write 1")
				for {
					rb.Write(1)
					atomic.AddUint64(&totalCountWrite, 1)
					//time.Sleep(5 * time.Second)
				}
			}()
		}

		doneInitWorkerW = true
		fmt.Println("done init write worker")
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 0; i < countW; i++ {
			fmt.Println("start one worker read")
			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()

				for {
					//fmt.Println("start loop read")
					data := rb.Read()
					atomic.AddUint64(&totalCountRead, 1)
					if data != 1 && data != nil {
						fmt.Println("have errors, data = ", data)
						panic(" have errors in code")
					}
					//time.Sleep(5 * time.Second)
				}
			}()
		}

		doneInitWorkerR = true
		fmt.Println("done init read worker")
	}()

	for {
		if doneInitWorkerR == true && doneInitWorkerW == true {
			fmt.Println("done init all worker read write")
			break
		}
	}

	fmt.Println("start init benchmark")
	// performance
	wg.Add(1)
	go func() {
		fmt.Println("--------------------------------------------------------")
		defer func() {
			wg.Done()
		}()

		var totalSecond, oldCountRead, oldCountWrite uint64
		totalSecond = 5
		tick := time.After(time.Duration(totalSecond) * time.Second)

		for {
			select {
			case <-tick:
				fmt.Println("start caculator benchmark")
				countR := atomic.LoadUint64(&totalCountRead)
				countW := atomic.LoadUint64(&totalCountWrite)

				fmt.Printf(" countR = %v, count W = %v, total ReadWrite ; %v,  countR/s: %v, countW/s: %v, totalReadWrite/s: %v \n",
					countR-oldCountRead, countW-oldCountWrite, countR+countW-oldCountWrite-oldCountRead, (countR-oldCountRead)/totalSecond, (countW-oldCountWrite)/totalSecond, ((countR+countW)-(oldCountRead+oldCountWrite))/totalSecond)
				tick = time.After(time.Duration(totalSecond) * time.Second)
				oldCountRead = countR
				oldCountWrite = countW
			}
		}
	}()

	wg.Wait()
}

func benchChannel() {
	var wg sync.WaitGroup
	channel := make(chan int)
	var totalCountWrite uint64
	var totalCountRead uint64

	var doneInitWorkerW, doneInitWorkerR bool

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 0; i < 1000; i++ {
			fmt.Println("start one worker write")
			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()

				//fmt.Println("start loop write 1")
				for {
					channel <- 1
					atomic.AddUint64(&totalCountWrite, 1)
					//time.Sleep(5 * time.Second)
				}
			}()
		}

		doneInitWorkerW = true
		fmt.Println("done init write worker")
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()

		for i := 0; i < 1000; i++ {
			fmt.Println("start one worker read")
			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()

				for {
					//fmt.Println("start loop read")
					data := <-channel
					atomic.AddUint64(&totalCountRead, 1)
					if data != 1 {
						fmt.Println("have errors, data = ", data)
						panic(" have errors in code")
					}
					//time.Sleep(5 * time.Second)
				}
			}()
		}

		doneInitWorkerR = true
		fmt.Println("done init read worker")
	}()

	for {
		if doneInitWorkerR == true && doneInitWorkerW == true {
			fmt.Println("done init all worker read write")
			break
		}
	}

	fmt.Println("start init benchmark")
	// performance
	wg.Add(1)
	go func() {
		fmt.Println("--------------------------------------------------------")
		defer func() {
			wg.Done()
		}()

		var totalSecond, oldCountRead, oldCountWrite uint64
		totalSecond = 5
		tick := time.After(time.Duration(totalSecond) * time.Second)

		for {
			select {
			case <-tick:
				fmt.Println("start calculator benchmark")
				countR := atomic.LoadUint64(&totalCountRead)
				countW := atomic.LoadUint64(&totalCountWrite)

				fmt.Printf(" countR = %v, count W = %v, total ReadWrite ; %v,  countR/s: %v, countW/s: %v, totalReadWrite/s: %v \n",
					countR-oldCountRead, countW-oldCountWrite, countR+countW-oldCountWrite-oldCountRead, (countR-oldCountRead)/totalSecond, (countW-oldCountWrite)/totalSecond, ((countR+countW)-(oldCountRead+oldCountWrite))/totalSecond)
				tick = time.After(time.Duration(totalSecond) * time.Second)
				oldCountRead = countR
				oldCountWrite = countW
			}
		}

	}()

	wg.Wait()
}

func main() {
	//testWriteAndRead()
	//testReadWaitWrite(50000)
	benchMarkRing(4000)
	//benchChannel()
}
