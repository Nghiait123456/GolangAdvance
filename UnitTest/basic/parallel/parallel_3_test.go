package parallel

import (
	"fmt"
	"testing"
)

func trace_3(name string) func() {
	fmt.Printf("%s entered\n", name)
	return func() {
		fmt.Printf("%s returned\n", name)
	}

}

func Test_3_Func1(t *testing.T) {
	defer trace_3("Test_Func1")()

	t.Run("Func1_Sub1", func(t *testing.T) {
		defer trace_3("Func1_Sub1")()
		t.Parallel()

		// ...
	})

	t.Run("Func1_Sub2", func(t *testing.T) {
		defer trace_3("Func1_Sub2")()

		t.Parallel()
		// ...
	})

	fmt.Println("continue run other test....")
	// ...
}

func Test_3_Func2(t *testing.T) {
	defer trace_3("Test_Func2")()
	t.Parallel()

	// ...
}

func Test_3_Func3(t *testing.T) {
	defer trace_3("Test_Func3")()
	t.Parallel()

	// ...
}

func Test_3_Func4(t *testing.T) {
	defer trace_3("Test_Func4")()
	t.Parallel()

	// ...
}

func Test_3_Func5(t *testing.T) {
	defer trace_3("Test_Func5")()

	// ...
}
