package __GroupSimilarDeclarations

//this is not good
//import "a"
//import "b"

//this is good
//import (
//	"a"
//	"b"
//)

// this is good
//const a = 1
//const b = 2
//
//var a = 1
//var b = 2
//
//type Area float64
//type Volume float64
//
//const (
//	a = 1
//	b = 2
//)
//
//var (
//	a = 1
//	b = 2
//)

type (
	Area   float64
	Volume float64
)

//type Operation int
//
//const (
//	Add Operation = iota + 1
//	Subtract
//	Multiply
//	EnvVar = "MY_ENV"
//)

//Only group related declarations. Do not group declarations that are unrelated.
type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
)

// this is bad
//const EnvVar = "MY_ENV"
//
//
//
//func f() string {
//	red := color.New(0xff0000)
//	green := color.New(0x00ff00)
//	blue := color.New(0x0000ff)
//
//	// ...
//}

//func (c *client) request() {
//	caller := c.name
//	format := "json"
//	timeout := 5*time.Second
//	var err error
//
//	// ...
//}

// this is good
//func f() string {
//	var (
//		red   = color.New(0xff0000)
//		green = color.New(0x00ff00)
//		blue  = color.New(0x0000ff)
//	)
//
//	// ...
//}

//func (c *client) request() {
//	var (
//		caller  = c.name
//		format  = "json"
//		timeout = 5*time.Second
//		err error
//	)
//
//	// ...
//}
