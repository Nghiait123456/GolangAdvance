package assertion_self

import (
	"testing"
)

// Assertion example with testify
func TestSum(t *testing.T) {

	// Use this for more then one assert
	// assert := assert.New(t)

	got, _ := Sum(2, 5)

	AssertEqual(t, got, 7)
	AssertNotEqual(t, got, 6)
}
