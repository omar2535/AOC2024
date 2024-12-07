package six

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

type position struct {
	x int
	y int
}

func PartOne(isTest bool) {
	fmt.Println("Day 6 part 1")
	var fileContents []string
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day6//day6_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day6/day6.txt")
	}

	// Create a grid of len(fileContents) x len(fileContents[0])
	grid := make([][]string, len(fileContents))
	for i := 0; i < len(fileContents); i++ {
		grid[i] = strings.Split(fileContents[i], "")
	}

	// Initial position
	getGuardPosition(grid)

}

func getGuardPosition(grid [][]string) position {
	return position{0, 0}
}
