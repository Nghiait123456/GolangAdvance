package _4_LocalVariableDeclarations

/**
Short variable declarations (:=) should be used if a variable is being set to some value explicitly.
However, there are cases where the default value is clearer when the var keyword is used. Declaring Empty Slices, for example.
*/

// this is bad
//var s = "foo"
//func f(list []int) {
//	filtered := []int{}
//	for _, v := range list {
//		if v > 10 {
//			filtered = append(filtered, v)
//		}
//	}
//}

// this is good
//s := "foo"
//func f(list []int) {
//	var filtered []int
//	for _, v := range list {
//		if v > 10 {
//			filtered = append(filtered, v)
//		}
//	}
//}
