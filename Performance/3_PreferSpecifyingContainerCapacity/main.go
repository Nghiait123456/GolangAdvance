package __PreferSpecifyingContainerCapacity

/**
Specify container capacity where possible in order to allocate memory for the container up front. This minimizes subsequent allocations (by copying and resizing of the container) as elements are added.
*/

var a [4]int                      // array with zero values
var b [4]int = [4]int{0, 1, 2}    // partially initialized array
var c [4]int = [4]int{1, 2, 3, 4} // array initialization
var d = [...]int{5, 6, 7, 0}

func main() {
}
