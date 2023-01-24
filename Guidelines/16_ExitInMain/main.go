package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// this is bad case, dont use log.Fatal in main
//func main() {
//	body := readFile(path)
//	fmt.Println(body)
//}
//
//func readFile(path string) string {
//	f, err := os.Open(path)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	b, err := io.ReadAll(f)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return string(b)
//}

func main() {
	fmt.Println("Go programs use os.Exit or log.Fatal* to exit immediately. (Panicking is not a good way to exit programs, please don't panic.)\n\nCall one of os.Exit or log.Fatal* only in main(). All other functions should return errors to signal failure." +
		"Go programs use os.Exit or log.Fatal* to exit immediately. (Panicking is not a good way to exit programs, please don't panic.)\n\nCall one of os.Exit or log.Fatal* only in main(). All other functions should return errors to signal failure.",
	)
	body, err := readFile("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)
}

func readFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
