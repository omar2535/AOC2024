package day8

import (
	"aoc2024/internal"
	"fmt"
)

type position struct {
	x int
	y int
}

func PartOne(isTest bool) {
	fmt.Println("Day 8 part 1")

	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day8/day8_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day8/day8.txt")
	}

	// Go through the file and convert to 2d array of strings
	var grid [][]string = internal.ConvertStringListToGrid(fileContents, "")

	// Variables we'll need to store positions
	var antinodePositionsMap map[position]bool = make(map[position]bool)
	var nodePositionsMap map[string][]position = make(map[string][]position)

	// First iteration: go through the grid and find all the node positions
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != "." {
				nodeType := grid[y][x]
				currentPosition := position{x: x, y: y}
				if _, ok := nodePositionsMap[nodeType]; !ok {
					nodePositionsMap[nodeType] = make([]position, 0)
				}
				nodePositionsMap[nodeType] = append(nodePositionsMap[nodeType], currentPosition)
			}
		}
	}

	// Second iteration: go through each node type and find their antinode maps
	for _, nodePositions := range nodePositionsMap {
		antinodePositionsMapFound := getAntinodePositions(grid, nodePositions)
		fmt.Println(antinodePositionsMapFound)
		for antinodePosition := range antinodePositionsMapFound {
			antinodePositionsMap[antinodePosition] = true
		}
	}

	fmt.Println("Number of antinode positions:", len(antinodePositionsMap))
}

// get the antiNode positions for the node
func getAntinodePositions(grid [][]string, nodePositions []position) map[position]bool {
	antinodePositionsMap := make(map[position]bool)
	for i := 0; i < len(nodePositions); i++ {
		firstNodePosition := nodePositions[i]
		for j := i; j < len(nodePositions); j++ {
			// skip if it's just the node on top of itself
			if i == j {
				continue
			}
			comparedNodePosition := nodePositions[j]
			antiNodesList := getAntinodesForPositionPair(grid, firstNodePosition, comparedNodePosition)
			for _, antinodePosition := range antiNodesList {
				antinodePositionsMap[antinodePosition] = true
			}
		}
	}
	return antinodePositionsMap
}

// Get a list of anitnodes for pair of positions
// IE. (2, 2,), (3, 3) -> [(1, 1), (4, 4)]
// IE: (3, 3), (5, 4) -> [(1, 2), (7, 5)]
func getAntinodesForPositionPair(grid [][]string, position1 position, position2 position) []position {
	p1XDiff := position1.x - position2.x
	p1YDiff := position1.y - position2.y
	p2XDiff := position2.x - position1.x
	p2YDiff := position2.y - position1.y

	antinode1 := position{x: position1.x + p1XDiff, y: position1.y + p1YDiff}
	antinode2 := position{x: position2.x + p2XDiff, y: position2.y + p2YDiff}

	result := make([]position, 0)
	if isPositionInBounds(len(grid[0]), len(grid), antinode1) {
		result = append(result, antinode1)
	}
	if isPositionInBounds(len(grid[0]), len(grid), antinode2) {
		result = append(result, antinode2)
	}
	return result
}

// Checks if the position is in the bounds
func isPositionInBounds(gridWidth int, grideHeight int, currentPosition position) bool {
	if currentPosition.x < 0 || currentPosition.y < 0 {
		return false
	} else if currentPosition.x >= gridWidth || currentPosition.y >= grideHeight {
		return false
	}
	return true
}
