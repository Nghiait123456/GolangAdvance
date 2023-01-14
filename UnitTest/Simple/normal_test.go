package normal

import "testing"

func TestSum(t *testing.T) {

	want := 8
	got := Sum(3, 5)

	if got != want {
		t.Errorf("Test fail! want: '%d', got: '%d'", want, got)
	}

}
