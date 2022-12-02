package main

import (
	"errors"
	"fmt"
)

func Layer1Err() error {
	return fmt.Errorf("Layer1Err: failed executing layer1: %w", errors.New("Layer0 -error"))
}

func Layer2Err() error {
	err := Layer1Err()
	return fmt.Errorf("Layer2Err: %w", err)

}

func main() {
	if err := Layer2Err(); err != nil {
		fmt.Printf("trace err: %s \n", err)
		return
	}

	fmt.Println("end")
}
