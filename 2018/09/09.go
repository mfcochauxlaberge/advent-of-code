package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("December 9th, 2018")

	input := []byte(``)
	// 	input := []byte(`1, 1
	// 1, 6
	// 8, 3
	// 3, 4
	// 5, 5
	// 8, 9`)

	res := ""
	lines := strings.Split(string(input), "\n")
	for i := 0; i < len(lines); i++ {

	}
	fmt.Printf("Result (part 1): %v\n", res)

}
