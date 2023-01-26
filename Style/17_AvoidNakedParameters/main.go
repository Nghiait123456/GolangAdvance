package AvoidNakedParameters

//Naked parameters in function calls can hurt readability. Add C-style comments (/* ... */)
//for parameter names when their meaning is not obvious

//this is bad
// func printInfo(name string, isLocal, done bool)

//printInfo("foo", true, true)

// this is good
// func printInfo(name string, isLocal, done bool)

//printInfo("foo", true /* isLocal */, true /* done */)

/**
Better yet, replace naked bool types with custom types for more readable and type-safe code. This allows more than just two states (true/false) for that parameter in the future.
*/

// this is good
//type Region int
//
//const (
//	UnknownRegion Region = iota
//	Local
//)
//
//type Status int
//
//const (
//	StatusReady Status = iota + 1
//	StatusDone
//	// Maybe we will have a StatusInProgress in the future.
//)
//
//func printInfo(name string, region Region, status Status)