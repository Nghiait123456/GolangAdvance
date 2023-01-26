package __FunctionGroupingAndOrdering

// this is bad
//func (s *something) Cost() {
//	return calcCost(s.weights)
//}
//
//type something struct{ ... }
//
//func calcCost(n []int) int {...}
//
//func (s *something) Stop() {...}
//
//func newSomething() *something {
//	return &something{}
//}

// this is good
//type something struct{ ... }
//
//func newSomething() *something {
//	return &something{}
//}
//
//func (s *something) Cost() {
//	return calcCost(s.weights)
//}
//
//func (s *something) Stop() {...}
//
//func calcCost(n []int) int {...}

/**
Function Grouping and Ordering
Functions should be sorted in rough call order.
Functions in a file should be grouped by receiver.
Therefore, exported functions should appear first in a file, after struct, const, var definitions.

A newXYZ()/NewXYZ() may appear after the type is defined, but before the rest of the methods on the receiver.

Since functions are grouped by receiver, plain utility functions should appear towards the end of the file.
*/
