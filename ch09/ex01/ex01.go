package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func main() {
	name := "not_found"
	err := fileChecker(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("File does not exist: %s\n", name)
		}
	}
	panic("msg")
}
