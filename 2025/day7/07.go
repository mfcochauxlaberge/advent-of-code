package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fileBytes, err := os.ReadFile("07.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Lines before:\n%s\n", string(fileBytes))

	lines := bytes.Split(fileBytes, []byte{'\n'})

	total := 0
	timelines := map[[2]int]uint{}

	for i, line := range lines {
		if i == 0 {
			continue
		}

		for j, char := range line {
			if char == '.' && lines[i-1][j] == 'S' {
				lines[i][j] = '|'
				timelines[[2]int{i, j}] += 1
			}

			if char == '.' && lines[i-1][j] == '|' {
				lines[i][j] = '|'
				incoming := timelines[[2]int{i - 1, j}]
				timelines[[2]int{i, j}] += incoming
			}

			if char == '^' && lines[i-1][j] == '|' {
				incoming := timelines[[2]int{i - 1, j}]

				if j > 0 {
					lines[i][j-1] = '|'
					timelines[[2]int{i, j - 1}] += incoming
				}

				if j < len(line)-1 {
					lines[i][j+1] = '|'
					timelines[[2]int{i, j + 1}] += incoming
				}

				total += 1
			}

			if char == '|' && lines[i-1][j] == '|' {
				incoming := timelines[[2]int{i - 1, j}]
				timelines[[2]int{i, j}] += incoming
			}
		}

		lineTimelines := 0

		for coord, num := range timelines {
			if coord[0] < i {
				continue
			}

			lineTimelines += int(num)
		}

		lines[i] = append(lines[i], fmt.Appendf(nil, " (%d timelines)", lineTimelines)...)
	}

	totalTimelines := 0

	for coord, num := range timelines {
		if coord[0] < len(lines)-1 {
			continue
		}

		totalTimelines += int(num)
	}

	fmt.Printf("Lines after:\n")
	for _, line := range lines {
		fmt.Printf("%s\n", string(line))
	}

	fmt.Printf("Total timelines: %d\n", totalTimelines)
}
