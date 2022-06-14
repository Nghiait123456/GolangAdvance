package main

import "fmt"

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func main() {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	fmt.Printf("%v %v %v %v", s1Val, s1Ptr, s2Val, s2Ptr)

	var i F
	i = s1Val
	fmt.Printf("%v", i)
	i = s1Ptr
	fmt.Printf("%v", i)
	i = s2Ptr
	fmt.Printf("%v", i)

	// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
	i = s2Val

}

//explain:
/**
  in main.go edit function main to main1, in file main1.go edit function man1 to main.
  run code main1.go, code run success

  if uncmt line 34, i = s2Val
*/
