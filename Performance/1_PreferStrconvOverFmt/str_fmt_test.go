package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint(rand.Int())
	}
}

func BenchmarkGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(rand.Int())
	}
}

/**
go test -bench=.

goos: linux
goarch: amd64
pkg: PreferStrconvOverFmt
cpu: Intel(R) Core(TM) i7-5600U CPU @ 2.60GHz
BenchmarkBad-4           6902845               171.3 ns/op
BenchmarkGood-4         13733556                81.37 ns/op
PASS
ok      PreferStrconvOverFmt    2.573s

*/
