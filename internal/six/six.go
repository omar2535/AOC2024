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

	// Initial currentPosition
	currentPosition := getGuardPosition(grid)
	fmt.Println("Guard position:", currentPosition)

	// Initialize map of visited positions
	visitedPositions := make(map[position]bool)
	for {
		grid = tick(grid, currentPosition)
		currentPosition = getGuardPosition(grid)
		visitedPositions[currentPosition] = true
		if currentPosition.x == -1 && currentPosition.y == -1 {
			break
		}
	}
	fmt.Println("Number of visited positions:", len(visitedPositions))
}

func getGuardPosition(grid [][]string) position {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" || grid[i][j] == "v" || grid[i][j] == "<" || grid[i][j] == ">" {
				return position{j, i}
			}
		}
	}
	return position{-1, -1}
}

// Ticks the guard grid
// either face up, down, left or right
func tick(grid [][]string, guardPosition position) [][]string {
	// If the guard is not on the grid, return the grid as is
	if guardPosition.x == -1 && guardPosition.y == -1 {
		return grid
	}
	// If the guard is on the grid, move the guard
	guardDirection := string(grid[guardPosition.y][guardPosition.x])
	if guardDirection == "^" {
		if guardPosition.y > 0 && grid[guardPosition.y-1][guardPosition.x] != "#" {
			grid[guardPosition.y][guardPosition.x] = "."
			grid[guardPosition.y-1][guardPosition.x] = "^"
		} else if guardPosition.y > 0 && grid[guardPosition.y-1][guardPosition.x] == "#" {
			grid[guardPosition.y][guardPosition.x] = ">" // Turn right cause of obstacle
		} else {
			grid[guardPosition.y][guardPosition.x] = "." // Fell off the screen
		}
	} else if guardDirection == "v" {
		if guardPosition.y < len(grid)-1 && grid[guardPosition.y+1][guardPosition.x] != "#" {
			grid[guardPosition.y][guardPosition.x] = "."
			grid[guardPosition.y+1][guardPosition.x] = "v"
		} else if guardPosition.y < len(grid)-1 && grid[guardPosition.y+1][guardPosition.x] == "#" {
			grid[guardPosition.y][guardPosition.x] = "<" // Turn left cause of obstacle
		} else {
			grid[guardPosition.y][guardPosition.x] = "." // Fell off the screen
		}
	} else if guardDirection == "<" {
		if guardPosition.x > 0 && grid[guardPosition.y][guardPosition.x-1] != "#" {
			grid[guardPosition.y][guardPosition.x] = "."
			grid[guardPosition.y][guardPosition.x-1] = "<"
		} else if guardPosition.x > 0 && grid[guardPosition.y][guardPosition.x-1] == "#" {
			grid[guardPosition.y][guardPosition.x] = "^" // Turn up cause of obstacle
		} else {
			grid[guardPosition.y][guardPosition.x] = "." // Fell off the screen
		}
	} else if guardDirection == ">" {
		if guardPosition.x < len(grid[0])-1 && grid[guardPosition.y][guardPosition.x+1] != "#" {
			grid[guardPosition.y][guardPosition.x] = "."
			grid[guardPosition.y][guardPosition.x+1] = ">"
		} else if guardPosition.x < len(grid[0])-1 && grid[guardPosition.y][guardPosition.x+1] == "#" {
			grid[guardPosition.y][guardPosition.x] = "v" // Turn down cause of obstacle
		} else {
			grid[guardPosition.y][guardPosition.x] = "." // Fell off the screen
		}
	}
	return grid
}
