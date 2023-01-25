package __2_SpecifyingSliceCapacity

import "testing"

func BenchmarkBad(b *testing.B) {
	size := 1000
	for n := 0; n < b.N; n++ {
		data := make([]int, 0)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

func BenchmarkGood(b *testing.B) {
	size := 1000
	for n := 0; n < b.N; n++ {
		data := make([]int, 0, size)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

/**
cpu: Intel(R) Core(TM) i7-5600U CPU @ 2.60GHz
BenchmarkBad-4            147159              9299 ns/op
BenchmarkGood-4           339157              3715 ns/op

Where possible, provide capacity hints when initializing slices with make(), particularly when appending.

make([]T, length, capacity)
Unlike maps, slice capacity is not a hint: the compiler will allocate enough memory for the capacity of the slice as provided to make(), which means that subsequent append() operations will incur zero allocations (until the length of the slice matches the capacity, after which any appends will require a resize to hold additional elements).

*/
