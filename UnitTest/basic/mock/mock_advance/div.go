package main

import (
	"math/rand"
)

func divByRand(numerator int) int {
	return numerator / int(rand.Intn(10))
}

type randNumberGenerator interface {
	randomInt(max int) int
}

type standardRand struct{}

func (s standardRand) randomInt(max int) int {
	return rand.Intn(max)
}
