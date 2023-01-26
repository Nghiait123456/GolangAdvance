package NilIsAValidSlice

/**
nil is a valid slice of length 0. This means that,

You should not return a slice of length zero explicitly. Return nil instead.
*/

// this is bad
//if x == "" {
//return []int{}
//}

// this is good
//if x == "" {
//	return nil
//}

/**
To check if a slice is empty, always use len(s) == 0. Do not check for nil
*/

//func isEmptyBad(s []string) bool {
//	return s == nil
//}

//func isEmptyGood(s []string) bool {
//	return len(s) == 0
//}

/**
The zero value (a slice declared with var) is usable immediately without make().
*/

//this is bad
//nums := []int{}
// or, nums := make([]int)

//if add1 {
//nums = append(nums, 1)
//}
//
//if add2 {
//nums = append(nums, 2)
//}

//this is good
//var nums []int
//
//if add1 {
//nums = append(nums, 1)
//}
//
//if add2 {
//nums = append(nums, 2)
//}

/**
Remember that, while it is a valid slice, a nil slice is not equivalent to an allocated slice of length 0 - one is nil and the other is not - and the two may be treated differently in different situations (such as serialization).
*/
