package parallel

import (
	"fmt"
	"testing"
)

func trace(name string) func() {
	fmt.Printf("%s entered\n", name)
	return func() {
		fmt.Printf("%s returned\n", name)
	}

}

func Test_2_Func1(t *testing.T) {
	defer trace("Test_Func1")()

	// ...
}

func Test_2_Func2(t *testing.T) {
	defer trace("Test_Func2")()
	t.Parallel()

	// ...
}

func Test_2_Func3(t *testing.T) {
	defer trace("Test_Func3")()

	// ...
}

func Test_2_Func4(t *testing.T) {
	defer trace("Test_Func4")()
	t.Parallel()

	// ...
}

func Test_2_Func5(t *testing.T) {
	defer trace("Test_Func5")()

	// ...
}
