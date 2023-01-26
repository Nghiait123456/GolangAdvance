package FormatStringsOutsidePrintf

/**
If you declare format strings for Printf-style functions outside a string literal, make them const values.

This helps go vet perform static analysis of the format string.
*/

// this is bad
//msg := "unexpected values %v, %v\n"
//fmt.Printf(msg, 1, 2)

// this is good
//const msg = "unexpected values %v, %v\n"
//fmt.Printf(msg, 1, 2)
