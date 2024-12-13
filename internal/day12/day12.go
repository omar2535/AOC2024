package day12

import (
	"aoc2024/internal"
	"fmt"
	"slices"
)

type Position struct {
	x int
	y int
}

type Region struct {
	regionId  string
	area      int
	perimeter int
	numSides  int
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
				newRegion := Region{
					regionId:  grid[y][x],
					area:      len(connectedRegions),
					perimeter: perimeter,
					positions: connectedRegionsMap,
					numSides:  0,
				}
				regions = append(regions, newRegion)
			}
		}
	}

	// Print the results
	fmt.Println(regions)
	fmt.Println(grid)
	fmt.Println("Total price: ", computePrice(regions))
}

func PartTwo(isTest bool) {
	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day12/day12_example_3.txt")
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
				regionId := grid[y][x]
				perimeter := computePerimiter(grid, connectedRegionsMap)
				numSides := computeNumberOfSides(grid, connectedRegions, regionId)
				newRegion := Region{
					regionId:  grid[y][x],
					area:      len(connectedRegions),
					perimeter: perimeter,
					positions: connectedRegionsMap,
					numSides:  numSides,
				}
				regions = append(regions, newRegion)
			}
		}
	}

	// Print the results
	fmt.Println(regions)
	fmt.Println(grid)
	fmt.Println("Total price: ", computeNewPrices(regions))
}

func computeNewPrices(regions []Region) int {
	total := 0
	for _, region := range regions {
		total += region.area * region.numSides
	}
	return total
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

// sliding window to find all the edges (# edges = # of sides)
func computeNumberOfSides(grid [][]string, positions []Position, regionId string) int {
	if regionId == "*" {
		return 0
	}
	numEdges := 0
	// Find the borders (min/max x and y)
	minX := getMinX(positions)
	maxX := getMaxX(positions)
	minY := getMinY(positions)
	maxY := getMaxY(positions)

	for y := minY - 1; y <= maxY; y++ {
		for x := minX - 1; x <= maxX; x++ {
			// Check if it's a edge
			numEdges += getNumEdge(grid, x, y, regionId, positions)
		}
	}
	return numEdges
}

// Checks the square starting from the position given
// If the square has 3 regions, then it's an edge
// The square is defined starting from the top left corner
func getNumEdge(grid [][]string, x int, y int, regionId string, regionPositions []Position) int {
	var curr string
	var right string
	var bottomLeft string
	var bottomRight string

	// if position is out of bounds for the grid, make it *
	if x+1 == 0 {
		curr = "*"
		bottomLeft = "*"
	}
	if y+1 == 0 {
		curr = "*"
		right = "*"
	}
	if x+1 >= len(grid) {
		right = "*"
		bottomRight = "*"
	}
	if y+1 >= len(grid) {
		bottomLeft = "*"
		bottomRight = "*"
	}

	// If we weren't out of bounds, then replace it with the correct value
	// In each replacement, make sure it's still part of our region
	if curr != "*" {
		if grid[y][x] == regionId && slices.Contains(regionPositions, Position{x: x, y: y}) {
			curr = grid[y][x]
		} else {
			curr = "*"
		}
	}
	if right != "*" {
		if grid[y][x+1] == regionId && slices.Contains(regionPositions, Position{x: x + 1, y: y}) {
			right = grid[y][x+1]
		} else {
			right = "*"
		}
	}
	if bottomLeft != "*" {
		if grid[y+1][x] == regionId && slices.Contains(regionPositions, Position{x: x, y: y + 1}) {
			bottomLeft = grid[y+1][x]
		} else {
			bottomLeft = "*"
		}
	}
	if bottomRight != "*" {
		if grid[y+1][x+1] == regionId && slices.Contains(regionPositions, Position{x: x + 1, y: y + 1}) {
			bottomRight = grid[y+1][x+1]
		} else {
			bottomRight = "*"
		}
	}
	distinctRegions := make(map[string]int)
	distinctRegions[curr] += 1
	distinctRegions[right] += 1
	distinctRegions[bottomLeft] += 1
	distinctRegions[bottomRight] += 1

	// If the number of regions for our region in this square is 3 or 1, then it's an edge
	// Or if edges are oppposite of each other
	if distinctRegions[regionId] == 3 || distinctRegions[regionId] == 1 {
		return 1
	} else if distinctRegions[regionId] == 2 {
		if curr == bottomRight && right == bottomLeft {
			return 2
		} else {
			return 0
		}
	} else {
		return 0
	}
}

func getMinX(positions []Position) int {
	minX := positions[0].x
	for _, position := range positions {
		if position.x < minX {
			minX = position.x
		}
	}
	return minX
}

func getMaxX(positions []Position) int {
	maxX := positions[0].x
	for _, position := range positions {
		if position.x > maxX {
			maxX = position.x
		}
	}
	return maxX
}

func getMinY(positions []Position) int {
	minY := positions[0].y
	for _, position := range positions {
		if position.y < minY {
			minY = position.y
		}
	}
	return minY
}

func getMaxY(positions []Position) int {
	maxY := positions[0].y
	for _, position := range positions {
		if position.y > maxY {
			maxY = position.y
		}
	}
	return maxY
}
