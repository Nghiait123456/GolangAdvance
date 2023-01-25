package __AvoidStringToByteConversion

import (
	"log"
	"testing"
)

func BenchmarkBad(b *testing.B) {
	var a []byte
	for i := 0; i < b.N; i++ {
		a = []byte("Hello world")
	}

	log.Println(a)
}

func BenchmarkGood(b *testing.B) {
	var a []byte
	data := []byte("Hello world")
	for i := 0; i < b.N; i++ {
		a = data
	}
	log.Println(a)
}

/**
BenchmarkBad-4  44844873                29.76 ns/op
BenchmarkGood-4 1000000000               0.5447 ns/op
*/
