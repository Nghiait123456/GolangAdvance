package main

import (
	"fmt"
	"sync"
)

// bad case
type SMapBad struct {
	sync.Mutex

	data map[string]string
}

func NewSMapBad() *SMapBad {
	return &SMapBad{
		data: make(map[string]string),
	}
}

func (m *SMapBad) Get(k string) string {
	m.Lock()
	defer m.Unlock()

	return m.data[k]
}

//good case
type SMap struct {
	mu sync.Mutex

	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}

func main() {
	fmt.Println("bad start mutex")
	muBad := new(sync.Mutex)
	muBad.Lock()
	fmt.Println(muBad, &muBad)

	fmt.Println("good start mutex")
	var muGood sync.Mutex
	muGood.Lock()
	fmt.Println(muGood, &muGood)

	fmt.Println("bad start mutex")
	var muBad1 *sync.RWMutex
	fmt.Println(muBad1, &muBad1)
	muBad1.Lock()

}

//we will explain some case in this code:
/**
1) The zero-value of sync.Mutex and sync.RWMutex is valid, so you almost never need a pointer to a mutex.

+) compile and run code, you get error: panic: runtime error: invalid memory address or nil pointer dereference
in line 60, var muBad1 *sync.RWMutex:  muBad1 is type pointer of sync.RWMutex. The mutexes are generally designed to work without any type of initialization.
if pointer to mutex, this pointer is nil pointer.
if you view line 62: 	muBad1.Lock(),  Lock function call from nill pointer ==> panic
cmt line 59 to line 62, run success

+) view to line 	muBad := new(sync.Mutex), you don't need it,
   The zero-value of sync.Mutex and sync.RWMutex is valid, you  never can pointer to muxtex.

2) If you use a struct by pointer, then the mutex should be a non-pointer field on it. Do not embed the mutex on the struct, even if the struct is not exported.

view in line 10: sync.Mutex, => it embedded muxte to struct, will use Lock, UnLock,  the Lock and Unlock methods are unintentionally part of the exported API of SMap.
this is not good

view in line 30: mu sync.Mutex, mu is not embedded in struct. The mutex and its methods are implementation details of SMap hidden from its callers.
this is good,  mu maybe custom, replace and independence with API of struct
*/
