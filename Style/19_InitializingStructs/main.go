package InitializingStructs

/**
Use Field Names to Initialize Structs
You should almost always specify field names when initializing structs. This is now enforced by go vet.
*/

//this is bad
//k := User{"John", "Doe", true}

// this is good
//k := User{
//FirstName: "John",
//LastName: "Doe",
//Admin: true,
//}

// Exception: Field names may be omitted in test tables when there are 3 or fewer fields
//tests := []struct{
//op Operation
//want string
//}{
//{Add, "add"},
//{Subtract, "subtract"},
//}
