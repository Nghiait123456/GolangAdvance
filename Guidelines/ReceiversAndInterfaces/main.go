package main

import "fmt"

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func main() {
	sVals := map[int]S{1: {"A"}}

	// You can only call Read using a value
	fmt.Println(sVals[1].Read())

	//sVals[1].Write("23_cleary_example_smart_fetch")

	sPtrs := map[int]*S{1: {"A"}}
	// You can call both Read and Write using a pointer
	fmt.Println(sPtrs[1].Read())
	sPtrs[1].Write("23_cleary_example_smart_fetch")
	fmt.Println(sPtrs[1].Read())

	temp := &S{"A"}
	temp.Write("test2")
}

//explain:
/**
Methods with value receivers can be called on pointers as well as values. Methods with pointer receivers can only be called on pointers or addressable values

  if uncmt linr 24 sVals[1].Write("23_cleary_example_smart_fetch"), you run code have error:
  : cannot call pointer method on sVals[1]
  : cannot call pointer method on sVals[1]
  .Write() call from one pointer, it's reason error
  in line 22, sPtrs is pointer, it's use success for sPtrs[1].Read().  Read() is method with value receive. Read() run success with both value or pinter pass
  in line 29, sPtrs[1].Write("23_cleary_example_smart_fetch"), Write() is method with pointer receive. It's run success with pointer or address receive
  in line 31, view: temp := &S{"A"}, temp is receive value, it line 32 will run success in this case
*/
