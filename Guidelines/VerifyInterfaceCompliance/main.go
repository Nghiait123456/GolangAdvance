package main

import "fmt"

type Somether interface {
	Method() bool
}

type MyType string

func (mt MyType) Method2() bool {
	return true
}

//func (mt MyType) Method() bool {
//	return true
//}

func main() {
	var _ Somether = (*MyType)(nil)
	fmt.Println("start main")
}

//explain:
/**
  +) if you run code, it's have error:

  cannot use (*MyType)(nil) (type *MyType) as type Somether in assignment:
        *MyType does not implement Somether (missing Method method)

  golang compile check and if have anything error in
  		 1) Exported types that are required to implement specific interfaces as part of their API contract
		 2) Exported or unexported types that are part of a collection of types implementing the same interface
		 3) Other cases where violating an interface would break users
 => compile will stop
 if don't have error in list 1,2,3, it would compile and do nothing


  +) if you cmt line 23: var _ Somether = (*MyType)(nil),  ad run code,  it's run success, why golang don't automatic check matching implement of all code

  +) you un cmt line 23, un cmt line 15,16,17, and run,  it's work and all rule interface matching.
*/
