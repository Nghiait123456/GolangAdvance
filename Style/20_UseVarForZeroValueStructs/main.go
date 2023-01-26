package UseVarForZeroValueStructs

/**
When all the fields of a struct are omitted in a declaration, use the var form to declare the struct.
This differentiates zero valued structs from those with non-zero fields similar to the distinction created for map initialization, and matches how we prefer to declare empty slices.
*/

// this is bad
//user := User{}

// this is good
//var user User
