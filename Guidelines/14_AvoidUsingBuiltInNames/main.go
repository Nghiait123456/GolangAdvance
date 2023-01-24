package main

//var error string

// `error` shadows the builtin
// this is example bad case
func handleErrorMessage(error string) {
	// `error` shadows the builtin
}

type Foo struct {
	// While these fields technically don't
	// constitute shadowing, grepping for
	// `error` or `string` strings is now
	// ambiguous.
	error  error
	string string
}

func (f Foo) Error() error {
	// `error` and `f.error` are
	// visually similar
	return f.error
}

func (f Foo) String() string {
	// `string` and `f.string` are
	// visually similar
	return f.string
}

//this is good case

var errorMessage string

// `error` refers to the builtin

// or

func handleErrorMessageGood(msg string) {
	// `error` refers to the builtin
}

type FooGood struct {
	// `error` and `string` strings are
	// now unambiguous.
	err error
	str string
}

func (f FooGood) Error() error {
	return f.err
}

func (f FooGood) String() string {
	return f.str
}

func main() {

}
