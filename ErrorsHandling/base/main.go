package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("/tmp/dat")
	if err != nil {
		fmt.Println("error: ", err.Error())
	}

	fmt.Println("read file success: ")
}
