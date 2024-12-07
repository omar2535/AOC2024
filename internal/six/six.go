package six

import (
	"aoc2024/internal"
	"fmt"
	"maps"
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

func PartTwo(isTest bool) {
	fmt.Println("Day 6 part 2")
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
	visitedPositions := make(map[position]map[string]bool)
	visitedPositions[currentPosition] = make(map[string]bool)
	visitedPositions[currentPosition][getGuardDirection(grid)] = true

	// Create variable to track where we could add obstacle for cycle
	placedObstacles := make(map[position]bool)

	// Keep ticking until we find a cycle
	for {
		grid = tick(grid, currentPosition)
		currentPosition = getGuardPosition(grid)
		if currentPosition.x == -1 && currentPosition.y == -1 {
			break
		}
		if visitedPositions[currentPosition] == nil {
			visitedPositions[currentPosition] = make(map[string]bool)
		}
		visitedPositions[currentPosition][getGuardDirection(grid)] = true
		if doesTurningRightHaveCycle(grid, visitedPositions, currentPosition) {
			placedObstaclePosition := getPlacedObstacle(currentPosition, getGuardDirection(grid))
			placedObstacles[placedObstaclePosition] = true
			fmt.Println("Cycle found after adding obstacle at:", placedObstaclePosition)
		}
	}
	fmt.Println("Number of visited positions:", len(visitedPositions))
	fmt.Println("Number of obstacles added:", len(placedObstacles))
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

func getGuardDirection(grid [][]string) string {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" || grid[i][j] == "v" || grid[i][j] == "<" || grid[i][j] == ">" {
				return grid[i][j]
			}
		}
	}
	return ""
}

func getPlacedObstacle(currentPosition position, currentDirection string) position {
	if currentDirection == "^" {
		return position{currentPosition.x, currentPosition.y - 1}
	} else if currentDirection == "v" {
		return position{currentPosition.x, currentPosition.y + 1}
	} else if currentDirection == "<" {
		return position{currentPosition.x - 1, currentPosition.y}
	} else if currentDirection == ">" {
		return position{currentPosition.x + 1, currentPosition.y}
	}
	return position{-1, -1}
}

func doesTurningRightHaveCycle(grid [][]string, visitedPositions map[position]map[string]bool, currentPosition position) bool {
	currentDirection := getGuardDirection(grid)
	if visitedPositions[currentPosition] == nil {
		return false
	}
	if currentPosition.x == -1 && currentPosition.y == -1 {
		return false
	}
	if currentDirection == "^" && currentPosition.y == 0 {
		return false
	}
	if currentDirection == "v" && currentPosition.y == len(grid)-1 {
		return false
	}
	if currentDirection == "<" && currentPosition.x == 0 {
		return false
	}
	if currentDirection == ">" && currentPosition.x == len(grid[0])-1 {
		return false
	}
	// Get the turned direction
	if visitedPositions[currentPosition][turnDirection(currentDirection)] {
		return true
	}
	// Check if keeping the same new direction has a cycle
	obstacleLocation := getPlacedObstacle(currentPosition, currentDirection)
	newGrid := internal.Clone2dArray(grid)
	newGrid[obstacleLocation.y][obstacleLocation.x] = "#"
	newVisitedPositions := cloneVisited(visitedPositions)
	return willCycle(newGrid, newVisitedPositions, currentPosition)
}

func willCycle(grid [][]string, visitedPositions map[position]map[string]bool, currentPosition position) bool {
	for {
		// return if we went off the grid
		if currentPosition.x == -1 && currentPosition.y == -1 {
			return false
		}

		newGrid := tick(grid, currentPosition)
		newDirection := getGuardDirection(newGrid)
		newPosition := getGuardPosition(newGrid)

		if visitedPositions[newPosition] == nil {
			visitedPositions[newPosition] = make(map[string]bool)
			visitedPositions[newPosition][newDirection] = true
		} else if visitedPositions[newPosition][newDirection] {
			return true
		}
		grid = newGrid
		currentPosition = newPosition
	}
}

func turnDirection(direction string) string {
	if direction == "^" {
		return ">"
	} else if direction == "v" {
		return "<"
	} else if direction == "<" {
		return "^"
	} else if direction == ">" {
		return "v"
	}
	return ""
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

func cloneVisited(visitedPositions map[position]map[string]bool) map[position]map[string]bool {
	newVisitedPositions := make(map[position]map[string]bool)
	for k, v := range visitedPositions {
		newVisitedPositions[k] = maps.Clone(v)
	}
	return newVisitedPositions
}
