package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fileBytes, err := os.ReadFile("10_test.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	lines := bytes.Split(fileBytes, []byte{'\n'})

	machines := make([]machine, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		m := machine{}

		// Lights
		lightStart := bytes.IndexByte(line, '[')
		lightEnd := bytes.IndexByte(line, ']')
		if lightStart != -1 && lightEnd != -1 {
			lightStr := line[lightStart+1 : lightEnd]
			m.lights = make([]bool, len(lightStr))
			for i, char := range lightStr {
				m.lights[i] = char == '#'
			}
		}

		// Buttons
		remaining := line[lightEnd+1:]
		for {
			start := bytes.IndexByte(remaining, '(')
			if start == -1 {
				break
			}
			end := bytes.IndexByte(remaining[start:], ')')
			if end == -1 {
				break
			}
			end += start

			buttonStr := remaining[start+1 : end]
			if len(buttonStr) > 0 {
				var button []uint
				nums := bytes.Split(buttonStr, []byte{','})
				for _, numStr := range nums {
					var num uint
					fmt.Sscanf(string(numStr), "%d", &num)
					button = append(button, num)
				}
				m.buttons = append(m.buttons, button)
			}
			remaining = remaining[end+1:]
		}

		// Joltage
		joltStart := bytes.IndexByte(line, '{')
		joltEnd := bytes.IndexByte(line, '}')
		if joltStart != -1 && joltEnd != -1 {
			joltStr := line[joltStart+1 : joltEnd]
			nums := bytes.Split(joltStr, []byte{','})
			m.joltage = make([]uint, 0, len(nums))
			for _, numStr := range nums {
				var num uint
				fmt.Sscanf(string(numStr), "%d", &num)
				m.joltage = append(m.joltage, num)
			}
		}

		machines = append(machines, m)
	}

	for i, m := range machines {
		fmt.Printf("Machine %d:\n", i+1)
		fmt.Printf("  Lights: %v\n", m.lights)
		fmt.Printf("  Buttons: %v\n", m.buttons)
		fmt.Printf("  Joltage: %v\n", m.joltage)
		fmt.Println()
	}

	// smallest := 1000

	// for i, m := range machines {

	// }

	// largest := 0

	// fmt.Printf("Largest rectangle area: %d\n", largest)
}

type machine struct {
	lights  []bool
	buttons [][]uint
	joltage []uint
}
