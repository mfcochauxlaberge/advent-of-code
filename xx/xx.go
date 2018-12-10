package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("December xth, 20xx")

	input := []byte(`hello
world`)
	// input := []byte(`test
	// data`)

	res := ""
	lines := strings.Split(string(input), "\n")
	for i := 0; i < len(lines); i++ {
		res += " " + lines[i]
	}

	fmt.Printf("Result (part 1): %v\n", res)
}
