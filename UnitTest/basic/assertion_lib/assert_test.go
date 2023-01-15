package assertion

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Assertion example with testify
func TestSum(t *testing.T) {

	// Use this for more then one assert
	// assert := assert.New(t)

	got, err := Sum(2, 5)

	assert.Equal(t, 7, got, "they should be equal")

	assert.NotEqual(t, 6, got, "they should not be equal")

	assert.Nil(t, err)

	err = errors.New("different then nil")
	if assert.NotNil(t, err) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "different then nil", err.Error())
	}
}

// Assertion with table test
func TestSumTable(t *testing.T) {
	assert := assert.New(t)

	var table = []struct {
		x    int
		y    int
		want int
	}{
		{2, 2, 4},
		{5, 3, 8},
		{8, 4, 12},
		{12, 5, 17},
	}

	for _, test := range table {
		got, _ := Sum(test.x, test.y)
		assert.Equal(test.want, got)
	}
}

func customAssert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Custom assertion function
func TestSumCustomAssertion(t *testing.T) {
	t.Run("Test 1+3", func(t *testing.T) {
		got, _ := Sum(1, 3)
		want := 4
		customAssert(t, got, want)
	})

	t.Run("Test 4+7", func(t *testing.T) {
		got, _ := Sum(4, 7)
		want := 11
		customAssert(t, got, want)
	})

}
