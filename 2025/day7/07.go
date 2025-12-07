package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fileBytes, err := os.ReadFile("07_test.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Lines before:\n%s\n", string(fileBytes))

	lines := bytes.Split(fileBytes, []byte{'\n'})

	total := 0

	for i, line := range lines {
		if i == 0 {
			continue
		}

		for j, char := range line {
			if char == '.' && (lines[i-1][j] == 'S' || lines[i-1][j] == '|') {
				lines[i][j] = '|'
			}

			if char == '^' && lines[i-1][j] == '|' {
				if j > 0 {
					lines[i][j-1] = '|'
				}
				if j < len(line)-1 {
					lines[i][j+1] = '|'
				}

				total += 1
			}
		}
	}

	fmt.Printf("Lines after:\n%s\n", string(fileBytes))

	fmt.Printf("Total: %d\n", total)
}
