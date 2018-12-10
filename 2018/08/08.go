package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("December 8th, 2018")

	// input := []byte(``)
	input := []byte(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`)

	nums := strings.Split(string(input), " ")
	root, _ := parseNode(nums, 0)
	summd := countmdentries(root)
	fmt.Printf("Result (part 1): %d\n", summd)
}

type node struct {
	level    int
	children []*node
	metadata []int
}

func countmdentries(n *node) int {
	sum := 0
	for _, node := range n.children {
		sum += countmdentries(node)
	}
	return sum
}

func parseNode(nums []string, level int) (*node, int) {
	fmt.Printf("NUMS: %+v\n", nums)
	if len(nums) < 3 {
		panic("nums is too short")
	}

	currNode := &node{
		level:    level,
		children: []*node{},
	}

	var nn *node
	end := 0

	headerChildren, _ := strconv.ParseInt(nums[0], 10, 64)
	headerMetadata, _ := strconv.ParseInt(nums[0], 10, 64)

	for i := int64(0); i < headerChildren; i++ {
		subnums := make([]string, len(nums[2:]))
		copy(subnums, nums[2:])
		nn, end = parseNode(subnums, level+1)
		currNode.children = append(currNode.children, nn)
	}

	for i := int64(0); i < headerMetadata; i++ {

	}

	return currNode, end + int(headerMetadata)
}
