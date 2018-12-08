package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("December 7th, 2018")

	input := []byte(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`)
	// 	input := []byte(`1, 1
	// 1, 6
	// 8, 3
	// 3, 4
	// 5, 5
	// 8, 9`)

	lines := strings.Split(string(input), "\n")
	for i := 0; i < len(lines); i++ {
		dep := strings.Split(lines[i], " ")[1]
		task := strings.Split(lines[i], " ")[7]
	}
	fmt.Printf("Result (part 1): Unknow (tbd from the areas)\n")
}
