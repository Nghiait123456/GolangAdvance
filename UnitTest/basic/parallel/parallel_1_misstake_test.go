package parallel

import (
	"fmt"
	"testing"
)

func Sum(x, y int) (int, error) {
	z := x + y
	return z, nil
}

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

func TestSumParalelMisstake(t *testing.T) {
	t.Parallel()
	for _, row := range table {
		testName := fmt.Sprintf("Test %d+%d", row.x, row.y)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got, _ := Sum(row.x, row.y)
			if got != row.want {
				t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
			}
			fmt.Printf("Test pass:want %d, got %d \n", row.want, got)
		})
	}
}

func TestSumParalelFixMisstake(t *testing.T) {
	t.Parallel()
	for _, row := range table {
		testName := fmt.Sprintf("Test %d+%d", row.x, row.y)
		row := row
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got, _ := Sum(row.x, row.y)
			if got != row.want {
				t.Errorf("Test fail! want: '%d', got: '%d'", row.want, got)
			}
			fmt.Printf("Test pass:want %d, got %d \n", row.want, got)
		})
	}
}
