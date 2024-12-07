package day6

import (
	"aoc2024/internal"
	"fmt"
	"strings"
	"sync"
)

func PartTwoParallel(isTest bool) {
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

	// Initial startingPosition
	startingPosition := getGuardPosition(grid)
	fmt.Println("Starting guard position:", startingPosition)

	// Initialize map of visited positions
	visitedPositions := make(map[position]map[string]bool)
	visitedPositions[startingPosition] = make(map[string]bool)
	visitedPositions[startingPosition][getGuardDirection(grid)] = true

	// Initialize variables for concurrency
	channel := make(chan int)
	waitGroup := &sync.WaitGroup{}

	// Add obstacle to different variations of the grid
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			go checkCycleRoutine(x, y, grid, visitedPositions, waitGroup, channel)
		}
	}

	// Get all the values in the channel and add them together
	total := 0
	for result := range channel {
		total += result
	}
	fmt.Print("\n")
	fmt.Println("Number of obstacles added:", total)
}

// Check for cycles in the grid in a goroutine
func checkCycleRoutine(x int, y int, grid [][]string, visitedPositions map[position]map[string]bool, waitGroup *sync.WaitGroup, channel chan int) {
	newGrid := internal.Clone2dArray(grid)
	newGrid[y][x] = "#"
	guardPosition := getGuardPosition(newGrid)
	if willCycle(newGrid, visitedPositions, guardPosition) {
		fmt.Println("Cycle at: ", x, y)
		channel <- 1
	} else {
		fmt.Println("No cycle at: ", x, y)
		channel <- 0
	}
}
