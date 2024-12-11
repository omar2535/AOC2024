package day10

import (
	"aoc2024/internal"
	"fmt"
)

func PartTwo(isTest bool) {
	fmt.Println("Day 10 Part 2")
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
				total += computeNumPathsThatLeadTo9(nodeGrid[y][x])
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
		total += computeNumPathsThatLeadTo9(*link)
	}
	return total
}
