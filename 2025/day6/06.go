package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("06.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	data := string(fileBytes)

	lines := strings.Split(data, "\n")
	numLines := len(lines) - 1

	lastLine := []byte(lines[len(lines)-1])

	spans := make([]int, 0)

	for _, c := range lastLine {
		if c == ' ' {
			spans[len(spans)-1]++
			continue
		}

		spans = append(spans, 1)
	}

	for i := range spans {
		if i == len(spans)-1 {
			break
		}

		spans[i]--
	}

	fmt.Printf("Spans: %v\n\n", spans)

	for i, line := range lines {
		fmt.Printf("Line %d: %s\n", i, line)
	}

	ops := strings.Fields(lines[len(lines)-1])
	fmt.Printf("Ops: %v\n", ops)

	fmt.Printf("\n")

	totalSpan := 0

	total := 0

	for i, span := range spans {
		numStrs := make([]string, 0)

		for j := span; j >= 0; j-- {
			numStr := ""

			for k := range numLines {
				c := lines[k][totalSpan+j]

				numStr += string(c)
			}

			numStr = strings.TrimSpace(numStr)

			numStrs = append(numStrs, numStr)
		}

		fmt.Printf("numStrs: %v\n", numStrs)

		result := 0

		switch ops[i] {
		case "+":
			sum := 0
			for _, ns := range numStrs {
				n, _ := strconv.Atoi(ns)
				sum += n
			}
			fmt.Printf("Sum: %d", sum)
			fmt.Println()

			result = sum
		case "*":
			prod := 1
			for _, ns := range numStrs {
				n, _ := strconv.Atoi(ns)
				if n > 0 {
					prod *= n
				}
			}
			fmt.Printf("Product: %d", prod)
			fmt.Println()

			result = prod
		default:
			fmt.Printf("Unknown operation: %s", ops[i])
		}

		total += result

		if totalSpan == 0 {
			totalSpan += span
		} else {
			totalSpan += span + 1
		}

		fmt.Printf("\n")
	}

	fmt.Printf("Total: %d\n", total)
}

type Operation struct {
	operands []uint
	op       string
}
