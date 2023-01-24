package main

import "fmt"

type Fooer interface {
	DummyName()
}

type Foo struct {
	Name string
}

func (f Foo) DummyName() {
	fmt.Printf(f.Name)
}

func main() {
	var f1 = Foo{Name: "f1"}
	var f2 = &Foo{Name: "f2"}
	fmt.Println("You almost never need a pointer to an interface. You should be passing interfaces as values—the underlying data can still be a pointer. An interface is two fields:\n\nA pointer to some type-specific information. You can think of this as \"type.\"\nData pointer. If the data stored is a pointer, it’s stored directly. If the data stored is a value, then a pointer to the value is stored")
	fmt.Println("Ex: Fooer interface access both value or pointer")
	DoFoo(f1)
	DoFoo(f2)
}

func DoFoo(f Fooer) {
	fmt.Printf("[%T] %+v\n", f, f)
}

//explain: 1). You almost never need a pointer to an interface. You should be passing interfaces as values—the underlying data can still be a pointer.
/**
if run code, result:
[main.Foo] {Name:f1}
[*main.Foo] &{Name:f2}

view in func DoFoo(f Fooer) { ... }, f varible in DoFoo is just variblean interface, is not pointer to an interface.
However, when pass and storing f2, the interface holds a pointer to a Foo structure.
you don't need use *f Foo to handle it.


if change func DoFoo(f *Fooer) { }, error : *Fooer is pointer to interface, not interface

*/

//explain 2: 2). An interface is two fields:
//
//A pointer to some type-specific information. You can think of this as "type."
//Data pointer. If the data stored is a pointer, it’s stored directly. If the data stored is a value, then a pointer to the value is stored.
//If you want interface methods to modify the underlying data, you must use a pointer.

/**
  please un all cmt in code, if you set variable for : SetName(name string), you can pass pointer f *Foo
  will run code, error: Foo does not implement Fooer (SetName method has pointer receiver)
  you  cmt line 26: DoFoo(f1) and run again: f1f2f1Newf2New, variable change succes

  if change function func (f *Foo) SetName(name string to  func (f Foo) SetName(name string and run code:
  you will get result : f1f2f1f2
  ==> variable not change if don't passed one pointer
*/
