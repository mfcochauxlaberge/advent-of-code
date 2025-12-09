package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fileBytes, err := os.ReadFile("09_test.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	lines := bytes.Split(fileBytes, []byte{'\n'})

	fmt.Printf("Total: %d\n", total)
}
