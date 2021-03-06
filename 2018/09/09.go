package main

import (
	"fmt"
)

func main() {
	fmt.Println("December 9th, 2018")

	numPlayers := 462
	lastMarblePoints := 7193800

	currentPlayer := 1
	marbles := &circle{
		marbles: make([]int, lastMarblePoints, lastMarblePoints),
		current: 0,
		length:  1,
	}
	scores := map[int]int{}
	for i := 0; i <= numPlayers; i++ {
		scores[i] = 0
	}

	// fmt.Printf("[-] %v\n", marbles)
	for m := 1; m <= lastMarblePoints; m++ {
		points := marbles.addMarble(m)
		scores[currentPlayer] += points
		// fmt.Printf("[%d] %v\n", currentPlayer, marbles)
		currentPlayer++
		if currentPlayer > numPlayers {
			currentPlayer = 1
		}
		// if m%100000 == 0 {
		// 	fmt.Printf("%7d\n", m)
		// }
	}
	bestPlayer := 0
	bestPlayerScore := 0
	for player, score := range scores {
		if score > bestPlayerScore {
			bestPlayer = player
			bestPlayerScore = score
		}
	}
	fmt.Printf("Result (part 1): %d (%d)\n", bestPlayer, bestPlayerScore)
	// For part 2, just run again with different lastMarblePoints
}

type circle struct {
	marbles []int
	current int // position
	length  int
}

func (c *circle) addMarble(n int) int {
	if n%23 == 0 {
		points := n

		newPos := c.current - 7
		if newPos < 0 {
			newPos += c.length
		}
		points += c.marbles[newPos]
		for i := newPos; i < c.length; i++ {
			c.marbles[i] = c.marbles[i+1]
		}
		c.current = newPos
		c.length--

		return points
	}

	newPos := c.current + 2
	if newPos >= c.length+1 {
		newPos = 1
	}
	for i := c.length; i > newPos; i-- {
		c.marbles[i] = c.marbles[i-1]
	}
	c.marbles[newPos] = n

	c.current = newPos
	c.length++

	return 0
}

func (c *circle) String() string {
	out := ""
	for i, m := range c.marbles {
		if i == c.current {
			out += fmt.Sprintf("(%02d)", m)
		} else {
			out += fmt.Sprintf(" %02d ", m)
		}
	}
	return out
}
