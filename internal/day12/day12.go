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
	positions map[Position]bool
}

func PartOne(isTest bool) {
	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day12/day12_example.txt")
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
			} else {
				connectedRegionsMap := make(map[Position]bool)
				connectedRegions := getConnectedRegions(grid, x, y, grid[y][x], visited)
				// Update the visited with the regions connected
				for _, connectedRegion := range connectedRegions {
					visited[connectedRegion.y][connectedRegion.x] = true
					connectedRegionsMap[connectedRegion] = true
				}
				// Add the regions that we have seen
				perimeter := computePerimiter(grid, connectedRegionsMap)
				regions = append(regions, Region{regionId: grid[y][x], area: len(connectedRegions), perimeter: perimeter, positions: connectedRegionsMap})
			}
		}
	}

	// Print the results
	fmt.Println(regions)
	fmt.Println(grid)
	fmt.Println("Total price: ", computePrice(regions))
}

func computePrice(regions []Region) int {
	total := 0
	for _, region := range regions {
		total += region.area * region.perimeter
	}
	return total
}

func computePerimiter(grid [][]string, connectedRegionsMap map[Position]bool) int {
	perimeter := 0
	for position := range connectedRegionsMap {
		// Check left if it's the same region
		if position.x-1 < 0 || !connectedRegionsMap[Position{x: position.x - 1, y: position.y}] {
			perimeter++
		}

		// Check right if it's the same region
		if position.x+1 >= len(grid[position.y]) || !connectedRegionsMap[Position{x: position.x + 1, y: position.y}] {
			perimeter++
		}

		// Check up if it's the same region
		if position.y-1 < 0 || !connectedRegionsMap[Position{x: position.x, y: position.y - 1}] {
			perimeter++
		}

		// Check down if it's the same region
		if position.y+1 >= len(grid) || !connectedRegionsMap[Position{x: position.x, y: position.y + 1}] {
			perimeter++
		}
	}
	return perimeter
}

func getConnectedRegions(grid [][]string, x int, y int, regionId string, visited [][]bool) []Position {
	// Check if the position is out of bounds
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
		return []Position{}
	}

	if grid[y][x] != regionId {
		return []Position{}
	}

	if visited[y][x] {
		return []Position{}
	}
	// Create a list of connected regions
	connectedRegions := make([]Position, 0)

	// Add the current position
	connectedRegions = append(connectedRegions, Position{x: x, y: y})
	visited[y][x] = true

	// Check the 4 directions
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x+1, y, regionId, visited)...)
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x-1, y, regionId, visited)...)
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x, y+1, regionId, visited)...)
	connectedRegions = append(connectedRegions, getConnectedRegions(grid, x, y-1, regionId, visited)...)

	return connectedRegions
}
