package main

import (
	"fmt"
	"time"
)

var _timeNow = time.Now

func sign(msg string) string {
	now := _timeNow()
	return now.String() + msg
}

type signer struct {
	now time.Time
}

func newSigner() *signer {
	return &signer{
		now: time.Now(),
	}
}

func (s *signer) Sign(msg string) string {
	return s.now.String() + msg
}

func main() {
	fmt.Println("Avoid mutating global variables, instead opting for dependency injection. This applies to function pointers as well as other kinds of values.")
	fmt.Println(sign("---- test global bad"))
	fmt.Println(newSigner().Sign("--- test good case "))
}
