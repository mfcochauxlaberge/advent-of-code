package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("December 6th, 2018")

	input := []byte(`108, 324
46, 91
356, 216
209, 169
170, 331
332, 215
217, 104
75, 153
110, 207
185, 102
61, 273
233, 301
278, 151
333, 349
236, 249
93, 155
186, 321
203, 138
103, 292
47, 178
178, 212
253, 174
348, 272
83, 65
264, 227
239, 52
243, 61
290, 325
135, 96
165, 339
236, 132
84, 185
94, 248
164, 82
325, 202
345, 323
45, 42
292, 214
349, 148
80, 180
314, 335
210, 264
302, 108
235, 273
253, 170
150, 303
249, 279
255, 159
273, 356
275, 244`)
	// 	input := []byte(`1, 1
	// 1, 6
	// 8, 3
	// 3, 4
	// 5, 5
	// 8, 9`)

	lines := strings.Split(string(input), "\n")
	coords := make([][2]int, len(lines))
	for i := 0; i < len(lines); i++ {
		coord := strings.Split(lines[i], ", ")
		x, _ := strconv.ParseInt(coord[1], 10, 64)
		y, _ := strconv.ParseInt(coord[0], 10, 64)
		coords[i] = [2]int{int(x) + 1000, int(y) + 1000}
	}
	areas := make([]int, len(coords))
	cells := make([][]cell, 2000)
	for i := range cells {
		cells[i] = make([]cell, 2000)
		for j := range cells[i] {
			cells[i][j] = newcell()
		}
	}
	for i := 0; i < len(coords); i++ {
		cells[coords[i][0]][coords[i][1]].coord = i
		areas[i] = 0
	}
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if cells[i][j].coord == -1 {
				dist := float64(999999)
				dist2 := float64(999999)
				closest := -1
				for c := 0; c < len(coords); c++ {
					_dist := distance([2]int{i, j}, coords[c])
					if _dist <= dist {
						dist2 = dist
						dist = _dist
						closest = c
					}
				}
				if dist != dist2 {
					cells[i][j].closest = closest
					areas[closest]++
				}
			}
		}
	}
	sort.Ints(areas)
	for i := 0; i < len(areas); i++ {
		fmt.Printf("Area %v: %v squares\n", i, areas[i]+1)
	}
	fmt.Printf("Result (part 1): Unknow (tbd from the areas)\n")

	regionsize := 0
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			// if cells[i][j].coord == -1 {
			sum := float64(0)
			for c := 0; c < len(coords); c++ {
				sum += distance([2]int{i, j}, coords[c])
			}
			if sum < 10000 {
				regionsize++
			}
			// }
		}
	}
	fmt.Printf("Result (part 2): %v\n", regionsize)

}

type cell struct {
	coord   int
	closest int
}

func newcell() cell {
	return cell{coord: -1, closest: -1}
}

func distance(a [2]int, b [2]int) float64 {
	var s float64
	s += math.Abs(float64(b[0]) - float64(a[0]))
	s += math.Abs(float64(b[1]) - float64(a[1]))
	return s
}
