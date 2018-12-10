package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("December 8th, 2018")

	// input := []byte(``)
	input := []byte(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`)

	nums := strings.Split(string(input), "\n")
	root := node{
		children: []node{},
	}
	for i := 0; i < len(nums); i++ {
		header:=
	}
	fmt.Printf("Result (part 1): Unknow (tbd from the areas)\n")
}

type node struct {
	children []node
}
