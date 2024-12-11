package day10

import (
	"aoc2024/internal"
	"fmt"
)

func PartOne(isTest bool) {
	fmt.Println("Day 10 part 1")
	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day10/day10_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day10/day10.txt")
	}

	var grid [][]string = internal.ConvertStringListToGrid(fileContents, "")
	var nodeGrid [][]node = create2dNodeArrayFromGrid(grid)
	nodeGrid = linkNodes(nodeGrid)

	// Go through the grid, for each starting node (0 node), compute the number of paths that lead to 9
	var total int = 0
	for y := 0; y < len(nodeGrid); y++ {
		for x := 0; x < len(nodeGrid[y]); x++ {
			if nodeGrid[y][x].value == 0 {
				reached9 := get9sReachedFromNode(nodeGrid[y][x])
				total += getNumDistinctNodes(reached9)
			}
		}
	}

	fmt.Println("Total number of paths that lead to 9:", total)
}

// compute the number of paths that lead to 9
func computeNumPathsThatLeadTo9(node node) int {
	if node.value == 9 {
		return 1
	} else if len(node.links) == 0 {
		return 0
	}

	var total int = 0
	for _, link := range node.links {
		fmt.Println("link: ", link.value)
		total += computeNumPathsThatLeadTo9(*link)
		fmt.Println("total: ", total)
	}
	return total
}

func get9sReachedFromNode(currentNode node) []node {
	if currentNode.value == 9 {
		return []node{currentNode}
	} else if len(currentNode.links) == 0 {
		return []node{}
	}

	var nodes []node
	for _, link := range currentNode.links {
		nodes = append(nodes, get9sReachedFromNode(*link)...)
	}
	return nodes
}

func getNumDistinctNodes(nodes []node) int {
	seenPositions := make(map[position]bool, 0)
	for _, node := range nodes {
		seenPositions[node.position] = true
	}
	return len(seenPositions)
}
