package day15

import (
	"aoc2024/internal"
	"fmt"
)

type Box struct {
	leftPosition  Position
	rightPosition Position
}

func PartTwo(isTest bool) {

	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day15/day15_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day15/day15.txt")
	}

	// Load into a grid and the moves
	var grid map[Position]PositionDetails = make(map[Position]PositionDetails)
	var gridWidth int = len(fileContents[0])
	var gridHeight int = 0
	var moves []string
	var startingPosition Position

	parsingGrid := true
	for _, line := range fileContents {
		if parsingGrid {
			if line == "" {
				parsingGrid = false
			} else {
				for x, c := range line {
					if c == '#' {
						grid[Position{x * 2, gridHeight}] = PositionDetails{Position{x * 2, len(grid)}, true, false, false}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetails{Position{x*2 + 1, len(grid)}, true, false, false}
					} else if c == '.' {
						grid[Position{x * 2, gridHeight}] = PositionDetails{Position{x * 2, len(grid)}, false, false, false}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetails{Position{x*2 + 1, len(grid)}, false, false, false}
					} else if c == 'O' {
						grid[Position{x * 2, gridHeight}] = PositionDetails{Position{x * 2, len(grid)}, false, true, false}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetails{Position{x*2 + 1, len(grid)}, false, true, false}
					} else if c == '@' {
						grid[Position{x * 2, gridHeight}] = PositionDetails{Position{x * 2, len(grid)}, false, false, true}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetails{Position{x*2 + 1, len(grid)}, false, false, false}
						startingPosition = Position{x, gridHeight}
					}
				}
				gridHeight++
			}
		} else {
			for _, c := range line {
				moves = append(moves, string(c))
			}
		}
	}

	// Perform the moves
	// var currentPosition Position = startingPosition
	printGrid(grid, gridWidth*2, gridHeight)
	// for _, move := range moves {
	// 	grid, currentPosition = doMove(grid, gridWidth, gridHeight, currentPosition, move)
	// 	// printGrid(grid, gridWidth, gridHeight)
	// }

	fmt.Println("Starting position: ", startingPosition)
	fmt.Println("Sum of coordinates: ", sumOfCoordinates(grid))
}
