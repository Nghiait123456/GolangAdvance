package UseRawStringLiteralsToAvoidEscaping

/**
Go supports raw string literals, which can span multiple lines and include quotes. Use these to avoid hand-escaped strings which are much harder to read
*/

// this is bad
//wantError := "unknown name:\"test\""

// this is good
//wantError := `unknown error:"test"`
