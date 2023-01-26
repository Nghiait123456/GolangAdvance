package InitializingStructReferences

/**
Use &T{} instead of new(T) when initializing struct references so that it is consistent with the struct initialization.
*/

// this is bad

//sval := T{Name: "foo"}
//// inconsistent
//sptr := new(T)
//sptr.Name = "bar"

//  this is good
//sval := T{Name: "foo"}
//
//sptr := &T{Name: "bar"}
