package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("December 11th, 2018")

	input := int64(5153)

	grid := [301][301]int64{}
	for y := int64(0); int(y) < len(grid); y++ {
		for x := int64(0); int(x) < len(grid); x++ {
			rackID := x + 10
			powerLevel := rackID * y
			powerLevel += input
			powerLevel *= rackID
			powerLevel, _ = strconv.ParseInt(string(fmt.Sprintf("%09d", powerLevel)[6]), 10, 64)
			powerLevel -= 5
			grid[x][y] = powerLevel
		}
	}
	largestX := int64(0)
	largestY := int64(0)
	largestFuel := int64(0)
	for y := int64(1); int(y) < len(grid)-2; y++ {
		for x := int64(1); int(x) < len(grid)-2; x++ {
			fuel := grid[x+0][y+0]
			fuel += grid[x+0][y+1]
			fuel += grid[x+0][y+2]
			fuel += grid[x+1][y+0]
			fuel += grid[x+1][y+1]
			fuel += grid[x+1][y+2]
			fuel += grid[x+2][y+0]
			fuel += grid[x+2][y+1]
			fuel += grid[x+2][y+2]
			if fuel > largestFuel {
				largestFuel = fuel
				largestX = x
				largestY = y
			}
		}
	}
	fmt.Printf("Result (part 1): %d,%d\n", largestX, largestY)

	largestX = int64(0)
	largestY = int64(0)
	largestFuel = int64(0)
	largestSquare := int64(0)
	fuel := int64(0)
	var (
		f, f2, y, x, s int64
	)
	for s = 1; int(s) < len(grid); s++ {
		for y = 1; int(y) < len(grid)-int(s)-1; y++ {
			for x = 1; int(x) < len(grid)-int(s)-1; x++ {
				fuel = 0
				for f = 0; f < s; f++ {
					for f2 = 0; f2 < s; f2++ {
						fuel += grid[x+f][y+f2]
						fuel += grid[x+f][y+f2]
						fuel += grid[x+f][y+f2]
					}
				}
				if fuel > largestFuel {
					largestFuel = fuel
					largestX = x
					largestY = y
					largestSquare = s
				}
			}
		}
	}
	fmt.Printf("Result (part 2): %d,%d,%d\n", largestX, largestY, largestSquare)
}
