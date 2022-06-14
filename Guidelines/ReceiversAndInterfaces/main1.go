package main

import "fmt"

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func main1() {
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
	//i = s2Val

}

//explain:
/**
 in main.go edit function main to main1, in file main1.go edit function man1 to main.
 run code main1.go, code run success

 if uncmt line 34, i = s2Val, you run and get error:
       cannot use s2Val (type S2) as type F in assignment:
       S2 does not implement F (f method has pointer receiver)
why have this error:
  before we repeat some new declarations:
	  "Method sets
		The method set of a type determines the methods that can be called on an operand of that type. Every type has a (possibly empty) method set associated with it:

		The method set of a defined type T consists of all methods declared with receiver type T.
		The method set of a pointer to a defined type T (where T is neither a pointer nor an interface) is the set of all methods declared with receiver *T or T.
		The method set of an interface type is the intersection of the method sets of each type in the interface's type set (the resulting method set is usually just the set of declared methods in the interface).
		Further rules apply to structs (and pointer to structs) containing embedded fields, as described in the section on struct types. Any other type has an empty method set.

		In a method set, each method must have a unique non-blank method name."

		"This compile-time error arises when you try to assign or pass (or convert) a concrete type to an interface type; and the type itself does not implement the interface, only a pointer to the type."

		"An assignment to a variable of interface type is valid if the value being assigned implements the interface it is assigned to. It implements it if its method set is a superset of the interface.
		The method set of pointer types includes methods with both pointer and non-pointer receiver. The method set of non-pointer types only includes methods with non-pointer receiver."
   if you don't understand this declarations, don't worry. You focus special bug in this case:

	   in line 13: type S2 struct{} implement interface f()( in line 6), any value is stored in an interface value f() must have this method f()
	   => The method set of a type determines the methods that can be called on an operand of that type, detail in this case:
	   in line 15: func (s *S2) f() {} is method with pointer receive.  f() call from pointer s *S2, f() is method set of *S2, f() is not in that  S2, f() require pointer receive( type *S2) for call it.
	   => we attempt to assign a value of S2  to a variable of type F, we get the error.
	   cannot use s2Val (type S2) as type F in assignment,  S2 does not implement F (f method has pointer receiver),
	   => clearly is s2Val not pointer receive =>  error
*/
