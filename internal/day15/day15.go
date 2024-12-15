package day15

import (
	"aoc2024/internal"
	"fmt"
)

type Position struct {
	x int
	y int
}

type PositionDetails struct {
	position Position
	hasWall  bool
	hasBox   bool
}

func PartOne(isTest bool) {

	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day15/day15_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day15/day15.txt")
	}

	// Load into a grid and the moves
	var grid map[Position]PositionDetails = make(map[Position]PositionDetails)
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
						grid[Position{x, len(grid)}] = PositionDetails{Position{x, len(grid)}, true, false}
					} else if c == '.' {
						grid[Position{x, len(grid)}] = PositionDetails{Position{x, len(grid)}, false, false}
					} else if c == 'O' {
						grid[Position{x, len(grid)}] = PositionDetails{Position{x, len(grid)}, false, true}
					} else if c == '@' {
						grid[Position{x, len(grid)}] = PositionDetails{Position{x, len(grid)}, false, false}
						startingPosition = Position{x, len(grid)}
					}
				}
			}
		} else {
			for _, c := range line {
				moves = append(moves, string(c))
			}
		}
	}

	fmt.Println(grid)
	fmt.Println(moves)
	fmt.Println(startingPosition)
}
