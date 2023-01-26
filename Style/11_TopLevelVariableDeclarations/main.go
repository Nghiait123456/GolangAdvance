package TopLevelVariableDeclarations

/**
At the top level, use the standard var keyword. Do not specify the type, unless it is not the same type as the expression.
*/

// this is bad
//var _s string = F()
//
//func F() string { return "A" }

// this is good

// this is good
//var _s = F()
//// Since F already states that it returns a string, we don't need to specify
//// the type again.
//
//func F() string { return "A" }
