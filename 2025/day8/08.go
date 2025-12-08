package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	fileBytes, err := os.ReadFile("08.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	lines := bytes.Split(fileBytes, []byte{'\n'})

	boxes := make(map[Box]struct{})

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var x, y, z int
		parts := bytes.Split(line, []byte{','})
		if len(parts) != 3 {
			err = fmt.Errorf("invalid format")
		} else {
			x, err = strconv.Atoi(string(parts[0]))
			if err == nil {
				y, err = strconv.Atoi(string(parts[1]))
			}
			if err == nil {
				z, err = strconv.Atoi(string(parts[2]))
			}
		}
		if err != nil {
			fmt.Printf("Failed to parse line '%s': %s\n", string(line), err)
			continue
		}

		box := Box{X: x, Y: y, Z: z}
		boxes[box] = struct{}{}
	}

	boxSlice := make([]Box, 0, len(boxes))
	for box := range boxes {
		boxSlice = append(boxSlice, box)
	}

	var pairs []Pair
	for i := 0; i < len(boxSlice); i++ {
		for j := i + 1; j < len(boxSlice); j++ {
			box1 := boxSlice[i]
			box2 := boxSlice[j]

			distance := int(math.Sqrt(
				float64((box1.X-box2.X)*(box1.X-box2.X) +
					(box1.Y-box2.Y)*(box1.Y-box2.Y) +
					(box1.Z-box2.Z)*(box1.Z-box2.Z)),
			))

			pairs = append(pairs, Pair{
				Box1:     box1,
				Box2:     box2,
				Distance: distance,
			})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Distance < pairs[j].Distance
	})

	for _, pair := range pairs {
		fmt.Printf("Distance %d: (%d,%d,%d) -> (%d,%d,%d)\n",
			pair.Distance,
			pair.Box1.X, pair.Box1.Y, pair.Box1.Z,
			pair.Box2.X, pair.Box2.Y, pair.Box2.Z)
	}

	circuits := make([]Circuit, 0)

	for _, box := range boxSlice {
		circuits = append(circuits, Circuit{
			Boxes: map[Box]struct{}{
				box: {},
			},
		})
	}

	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]

		var circuit1, circuit2 *Circuit
		var idx1, idx2 int

		for j, circuit := range circuits {
			if _, exists := circuit.Boxes[pair.Box1]; exists {
				circuit1 = &circuits[j]
				idx1 = j
			}
			if _, exists := circuit.Boxes[pair.Box2]; exists {
				circuit2 = &circuits[j]
				idx2 = j
			}
		}

		if circuit1 != nil && circuit2 != nil && idx1 != idx2 {
			for box := range circuit2.Boxes {
				circuit1.Boxes[box] = struct{}{}
			}

			circuits = append(circuits[:idx2], circuits[idx2+1:]...)
		}

		if len(circuits) == 1 {
			fmt.Printf(
				"Box %v connected with pair %v.\n",
				pair.Box1, pair.Box2,
			)

			fmt.Printf(
				"All boxes connected into one circuit with %d boxes.\n",
				len(circuits[0].Boxes),
			)

			result := pair.Box1.X * pair.Box2.X

			fmt.Printf("Result: %d\n", result)

			break
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].Boxes) > len(circuits[j].Boxes)
	})

	fmt.Printf("Circuit 1: %d boxes\n", len(circuits[0].Boxes))
	// fmt.Printf("Circuit 2: %d boxes\n", len(circuits[1].Boxes))
	// fmt.Printf("Circuit 3: %d boxes\n", len(circuits[2].Boxes))

	// total := len(circuits[0].Boxes) * len(circuits[1].Boxes) * len(circuits[2].Boxes)
	total := 0

	// for box := range boxes {
	// 	fmt.Printf("Box: %d,%d,%d\n", box.X, box.Y, box.Z)
	// 	total++
	// }

	fmt.Printf("Total: %d\n", total)
}

type Box struct {
	X int
	Y int
	Z int
}

type Pair struct {
	Box1     Box
	Box2     Box
	Distance int
}

type Circuit struct {
	Boxes map[Box]struct{}
}
