package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

func main() {
	fileBytes, err := os.ReadFile("09.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	lines := bytes.Split(fileBytes, []byte{'\n'})

	maxX, maxY := 0, 0
	coords := make([][2]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var x, y int
		parts := bytes.Split(line, []byte{','})
		x, _ = strconv.Atoi(string(parts[0]))
		y, _ = strconv.Atoi(string(parts[1]))

		coords = append(coords, [2]int{x, y})

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	fmt.Printf("Coordinates parsed (%d).\n", len(coords))

	grid := make([][]byte, maxY+1)
	for i := range grid {
		grid[i] = make([]byte, maxX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	fmt.Printf("Grid initialized (%dx%d).\n", len(grid), len(grid[0]))

	for _, coord := range coords {
		x, y := coord[0], coord[1]
		grid[y][x] = '#'
	}

	fmt.Printf("Coordinates marked on grid.\n")

	// rectSizes := make(map[[2]int]int)

	hashPositions := make([][2]int, 0)
	for y := range grid {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				hashPositions = append(hashPositions, [2]int{x, y})
			}
		}
	}

	fmt.Printf("Hash positions calculated (%d).\n", len(hashPositions))

	rects := make([]rect, 0)

	for i := 0; i < len(hashPositions); i++ {
		for j := i + 1; j < len(hashPositions); j++ {
			pos1 := hashPositions[i]
			pos2 := hashPositions[j]

			minX, maxX := pos1[0], pos2[0]
			if minX > maxX {
				minX, maxX = maxX, minX
			}

			minY, maxY := pos1[1], pos2[1]
			if minY > maxY {
				minY, maxY = maxY, minY
			}

			area := (maxX - minX + 1) * (maxY - minY + 1)

			rects = append(rects, rect{
				x1:   minX,
				y1:   minY,
				x2:   maxX,
				y2:   maxY,
				area: area,
			})
		}
	}

	sort.Slice(rects, func(i, j int) bool {
		return rects[i].area > rects[j].area
	})

	largest := 0

	for i := 0; i < len(hashPositions); i++ {
		for j := i + 1; j < len(hashPositions); j++ {
			pos1 := hashPositions[i]
			pos2 := hashPositions[j]

			// Both tiles align vertically.
			if pos1[0] == pos2[0] {
				minY, maxY := pos1[1], pos2[1]
				if minY > maxY {
					minY, maxY = maxY, minY
				}

				for y := minY; y <= maxY; y++ {
					if grid[y][pos1[0]] == '#' {
						continue
					}

					grid[y][pos1[0]] = 'X'
				}
			}

			// Both tiles align horizontally.
			if pos1[1] == pos2[1] {
				minX, maxX := pos1[0], pos2[0]
				if minX > maxX {
					minX, maxX = maxX, minX
				}

				for x := minX; x <= maxX; x++ {
					if grid[pos1[1]][x] == '#' {
						continue
					}

					grid[pos1[1]][x] = 'X'
				}
			}
		}
	}

	fmt.Printf("Edges done.\n")

	for y := range grid {
		start := 0
		end := len(grid[y]) - 1

		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' || grid[y][x] == 'X' {
				start = x
				break
			}
		}

		for x := len(grid[y]) - 1; x >= 0; x-- {
			if grid[y][x] == '#' || grid[y][x] == 'X' {
				end = x
				break
			}
		}

		if start == 0 && end == len(grid[y])-1 {
			continue
		}

		for x := start; x <= end; x++ {
			if grid[y][x] != '.' {
				continue
			}

			grid[y][x] = 'X'
		}
	}

	fmt.Printf("Area filled.\n")

	// edgedPositions := make([][2]int, 0)

	// Probably useless optimization since all red carpets are
	// not fully surrounded.
	// for i := 0; i < len(hashPositions); i++ {
	// 	pos := hashPositions[i]

	// 	surrounded := true

	// 	x1, y1 := pos[0], pos[1]
	// 	if x1 > 0 && grid[y1][x1-1] != '#' && grid[y1][x1-1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if x1 < len(grid[0])-1 && grid[y1][x1+1] != '#' && grid[y1][x1+1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if y1 > 0 && grid[y1-1][x1] != '#' && grid[y1-1][x1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if y1 < len(grid)-1 && grid[y1+1][x1] != '#' && grid[y1+1][x1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if x1 > 0 && y1 > 0 && grid[y1-1][x1-1] != '#' && grid[y1-1][x1-1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if x1 < len(grid[0])-1 && y1 > 0 && grid[y1-1][x1+1] != '#' && grid[y1-1][x1+1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if x1 > 0 && y1 < len(grid)-1 && grid[y1+1][x1-1] != '#' && grid[y1+1][x1-1] != 'X' {
	// 		surrounded = false
	// 	}
	// 	if x1 < len(grid[0])-1 && y1 < len(grid)-1 && grid[y1+1][x1+1] != '#' && grid[y1+1][x1+1] != 'X' {
	// 		surrounded = false
	// 	}

	// 	if surrounded {
	// 		continue
	// 	}

	// 	edgedPositions = append(edgedPositions, hashPositions[i])
	// }

	// hashPositions = edgedPositions

	fmt.Printf("Edged hash positions calculated (%d).\n", len(hashPositions))

	// largest = 0
	// count := 0
	// totalCount := (len(hashPositions) * (len(hashPositions))) / 2
	// lock := sync.RWMutex{}
	funcs := make(chan func(int) int, 300_000)

	for _, rect := range rects {
		f := func(largest int) int {
			// fmt.Printf(
			// 	"Checking rectangle between (%d,%d) and (%d,%d) with area %d (current largest is %d)\n",
			// 	rect.x1, rect.y1, rect.x2, rect.y2, rect.area, largest,
			// )

			// perc := (float64(i) / float64(len(hashPositions))) * 100
			// fmt.Printf("Checking position pair %d and %d... (%.2f%%)\n", i, j, perc)
			// defer func() {
			// 	lock.Lock()
			// 	count++
			// 	if count%1000 == 0 {
			// 		fmt.Printf(
			// 			"Processed %d position pairs... (%.2f%%)\n",
			// 			count, (float64(count)/float64(totalCount))*100,
			// 		)
			// 	}
			// 	lock.Unlock()
			// }()

			if largest > 0 {
				// fmt.Printf("Current largest area is %d\n", largest)
			} else {
				// fmt.Printf("No largest area found yet.\n")
			}

			x := abs(rect.x2-rect.x1) + 1
			y := abs(rect.y2-rect.y1) + 1

			area := x * y

			// rectSizes[[2]int{pos1[0], pos1[1]}] = area

			// lock.RLock()
			// largestLocal := largest
			// lock.RUnlock()

			// if area > largestLocal {
			if area > 0 {
				// fmt.Printf("Potential new largest area %d found between points (%d,%d) and (%d,%d)\n", area, pos1[0], pos1[1], pos2[0], pos2[1])
				// fmt.Printf("Checking if covered in red or green tiles...\n")

				minX, maxX := rect.x1, rect.x2
				if minX > maxX {
					minX, maxX = maxX, minX
				}

				minY, maxY := rect.y1, rect.y2
				if minY > maxY {
					minY, maxY = maxY, minY
				}

				allCovered := true
				for y := minY; y <= maxY; y++ {
					for x := minX; x <= maxX; x++ {
						if grid[y][x] != '#' && grid[y][x] != 'X' {
							allCovered = false
							break
						}
					}

					if !allCovered {
						break
					}
				}

				if !allCovered {
					// fmt.Printf("Area not fully covered, skipping.\n")
					// fmt.Println()
					return 0
				}

				// lock.Lock()
				// largest = area
				// lock.Unlock()

				// fmt.Printf("Fully covered. Updating largest area to %d\n", largest)
				// fmt.Printf("dx = %d, dy = %d, side1 = %d, side2 = %d\n", x, y, x, y)
				// fmt.Printf("New largest area %d found between points (%d,%d) and (%d,%d)\n", area, pos1[0], pos1[1], pos2[0], pos2[1])
				// fmt.Println()

				return largest
			}

			fmt.Printf("Area %d from (%d,%d) to (%d,%d) is not covered in red or green tiles.\n", area, rect.x1, rect.y1, rect.x2, rect.y2)

			return 0
		}

		funcs <- f
	}

	fmt.Printf("Function channel populated.\n")

	const numWorkers = 12
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for range numWorkers {
		go func() {
			defer wg.Done()
			for f := range funcs {
				result := f(largest)
				if result > 0 {
					fmt.Printf("Largest area is %d\n", result)
					return
				}
			}
		}()
	}

	wg.Wait()

	close(funcs)

	// largest := 0
	// for _, size := range rectSizes {
	// 	if size > largest {
	// 		largest = size
	// 	}
	// }

	for _, row := range grid {
		fmt.Printf("%s\n", string(row))
	}

	fmt.Printf("Largest rectangle area: %d\n", largest)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type rect struct {
	x1, y1, x2, y2 int
	area           int
}
