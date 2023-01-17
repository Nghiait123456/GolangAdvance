package subtest

import (
	"testing"
)

func TestSubtests(t *testing.T) {
	// <setup code>
	t.Run("A", func(t *testing.T) {
		t.Log("Test A1 completed")
	})

	t.Run("B", func(t *testing.T) {
		t.Log("Test B completed")
	})

	t.Run("C", func(t *testing.T) {
		t.Log("Test C completed")
	})
	// <teardown code>
}
