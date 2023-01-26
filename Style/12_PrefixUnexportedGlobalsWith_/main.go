package PrefixUnexportedGlobalsWith_

/**
Prefix unexported top-level vars and consts with _ to make it clear when they are used that they are global symbols.

Rationale: Top-level variables and constants have a package scope. Using a generic name makes it easy to accidentally use the wrong value in a different file.
Prefix unexported top-level vars and consts with _ to make it clear when they are used that they are global symbols.

Rationale: Top-level variables and constants have a package scope. Using a generic name makes it easy to accidentally use the wrong value in a different file.
*/

// this is bad
//const (
//	defaultPort = 8080
//	defaultUser = "user"
//)
//
//// bar.go
//
//func Bar() {
//	defaultPort := 9090
//	...
//	fmt.Println("Default port", defaultPort)
//
//	// We will not see a compile error if the first line of
//	// Bar() is deleted.
//}

// this is good
const (
	_defaultPort = 8080
	_defaultUser = "user"
)
