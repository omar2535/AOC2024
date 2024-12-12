package day12

import (
	"aoc2024/internal"
	"fmt"
)

type Position struct {
	x int
	y int
}

type Region struct {
	regionId  string
	area      int
	perimeter int
	positions []Position
}

func PartOne(isTest bool) {
	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day12/day12_example_small.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day12/day12.txt")
	}

	// Convert to grid
	grid := internal.ConvertStringListToGrid(fileContents, "")

	// Create a list of regions and the visited
	regions := make([]Region, 0)
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	// Go througy the grid and add regions
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visited[y][x] {
				continue
			}
		}
	}

	fmt.Println(grid)
}

func getConnectedRegions(grid [][]string, x int, y int, regionId string) []Position {
	// Check if the position is out of bounds
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return []Position{}
	}

	// Create a list of connected regions
	connectedRegions := make([]Position, 0)

	// Add the current position
	connectedRegions = append(connectedRegions, Position{x: x, y: y})

	// Check the 4 directions
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x+1, y, regionId)...)
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x-1, y, regionId)...)
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x, y+1, regionId)...)
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x, y-1, regionId)...)

	return connectedRegions
}
