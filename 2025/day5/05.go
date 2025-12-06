package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("05.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	data := string(fileBytes)

	parts := strings.Split(data, "\n\n")
	rangesRaw := parts[0]

	ranges := [][2]uint{}
	for _, line := range strings.Split(rangesRaw, "\n") {
		if len(line) == 0 {
			continue
		}

		bounds := strings.Split(line, "-")

		start, _ := strconv.ParseUint(bounds[0], 10, 0)
		end, _ := strconv.ParseUint(bounds[1], 10, 0)

		ranges = append(ranges, [2]uint{uint(start), uint(end)})
	}

	moved := true
	for moved {
		moved = false

		newRanges := map[[2]uint]struct{}{}

		count := 0
		for i := 0; i < len(ranges); i++ {
			merged := false

			for j := 0; j < len(ranges); j++ {
				count++

				if i == j {
					continue
				}

				a := ranges[i]
				b := ranges[j]

				if a[1] >= b[0] && b[1] >= a[0] {
					newStart := a[0]
					if b[0] < newStart {
						newStart = b[0]
					}

					newEnd := a[1]
					if b[1] > newEnd {
						newEnd = b[1]
					}

					newRanges[[2]uint{newStart, newEnd}] = struct{}{}

					merged = true

					break
				}
			}

			if merged {
				moved = true
			} else {
				newRanges[ranges[i]] = struct{}{}
			}
		}

		ranges = slices.Collect(maps.Keys(newRanges))
	}

	totalIds := 0
	for _, r := range ranges {
		totalIds += int(r[1] - r[0] + 1)
	}

	fmt.Printf("Total IDs covered: %d\n", totalIds)
}
