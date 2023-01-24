package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func mainBad() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("missing file")
	}
	name := args[0]

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// If we call log.Fatal after this line,
	// f.Close will not be called.

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
	// ...
}

func main() {
	fmt.Println("this is main good, if have many exit in main, move them to other fc to help main clear and beauty")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	args := os.Args[1:]
	if len(args) != 1 {
		return errors.New("missing file")
	}
	name := args[0]

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	fmt.Println(b)

	// ...
	return nil
}
