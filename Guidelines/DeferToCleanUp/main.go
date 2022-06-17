package main

import (
	"fmt"
	"sync"
)

type Test struct {
	mu    sync.Mutex
	count int
}

func (t *Test) notGoodNotUseDefer() int {
	t.mu.Lock()
	if t.count > 10 {
		t.mu.Unlock()
		return t.count
	}

	t.count++
	newCount := t.count
	t.mu.Unlock()

	return newCount
}

func (t *Test) GoodUseDefer() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.count > 10 {
		return t.count
	}

	t.count++
	newCount := t.count
	return newCount
}
func main() {
	test := &Test{
		count: 0,
	}

	fmt.Println(test.notGoodNotUseDefer())
	fmt.Println(test.GoodUseDefer())
}

// 1) Use defer to clean up resources such as files and locks.
// 2) Defer has an extremely small overhead and should be avoided only if you can prove that your function execution time is in the order of nanoseconds.
//The readability win of using defers is worth the miniscule cost of using them. This is especially true for larger methods that have more than simple memory accesses, where the other computations are more significant than the defer.

/**
  we will explain 1:
  notGoodNotUseDefer and not good code, GoodUseDefer is goodCode
  why:
      in line 16, 21, everytime break function, it always call t.mu.Unlock(), call before break
      in function maybe have many point break, you can put t.mu.Unlock() to it, if lost it, it has risk and bug very hard debug (bug race condition)
      function is boring and repeat code : t.mu.Unlock()

      in line 29: defer t.mu.Unlock(), only one line, when before function close, function in defer will call, code very clear and don't miss case unLock

  best practice:
     summary:  goodWay user defer is always better than badWay not use defer.
     ==> no, but most of case, you should use defer
     but, notGoodNotUseDefer() run faster GoodUseDefer() why defer run slowly with not defer.
     if you can speed and sure you don't miss case, you should not use defer, when really sure you can SPEED and sure is not miss CASE.
     ==> be careful, why race condition is bug very hard debug
*/
