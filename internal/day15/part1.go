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
	hasRobot bool
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
						grid[Position{x, gridHeight}] = PositionDetails{Position{x, len(grid)}, true, false, false}
					} else if c == '.' {
						grid[Position{x, gridHeight}] = PositionDetails{Position{x, len(grid)}, false, false, false}
					} else if c == 'O' {
						grid[Position{x, gridHeight}] = PositionDetails{Position{x, len(grid)}, false, true, false}
					} else if c == '@' {
						grid[Position{x, gridHeight}] = PositionDetails{Position{x, len(grid)}, false, false, true}
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
	var currentPosition Position = startingPosition
	printGrid(grid, gridWidth, gridHeight)
	for _, move := range moves {
		grid, currentPosition = doMove(grid, gridWidth, gridHeight, currentPosition, move)
		// printGrid(grid, gridWidth, gridHeight)
	}

	fmt.Println("Sum of coordinates: ", sumOfCoordinates(grid))
}

// Calculate sum of all cordinates
func sumOfCoordinates(grid map[Position]PositionDetails) int {
	sum := 0
	for k := range grid {
		if grid[k].hasBox {
			sum += 100*k.y + k.x
		}
	}
	return sum
}

// Performs a single move
func doMove(grid map[Position]PositionDetails, gridWidth int, gridHeight int, currentPosition Position, move string) (map[Position]PositionDetails, Position) {
	var newGrid map[Position]PositionDetails
	var newPosition Position
	if move == "^" {
		fmt.Println("Moving up")
		newGrid, newPosition = moveUp(grid, currentPosition)
	} else if move == "v" {
		fmt.Println("Moving down")
		newGrid, newPosition = moveDown(grid, gridHeight, currentPosition)
	} else if move == "<" {
		fmt.Println("Moving left")
		newGrid, newPosition = moveLeft(grid, currentPosition)
	} else if move == ">" {
		fmt.Println("Moving right")
		newGrid, newPosition = moveRight(grid, gridWidth, currentPosition)
	} else {
		panic("Invalid move")
	}
	return newGrid, newPosition
}

func moveUp(grid map[Position]PositionDetails, currentPosition Position) (map[Position]PositionDetails, Position) {
	// If at wall or end, return the same grid
	if currentPosition.y == 0 || grid[Position{currentPosition.x, currentPosition.y - 1}].hasWall {
		return grid, currentPosition
	}

	// If there's a space above, move up (means return the same grid)
	if !grid[Position{currentPosition.x, currentPosition.y - 1}].hasBox && !grid[Position{currentPosition.x, currentPosition.y - 1}].hasWall {
		grid[Position{currentPosition.x, currentPosition.y - 1}] = PositionDetails{Position{currentPosition.x, currentPosition.y - 1}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x, currentPosition.y - 1}
	}

	// If there's a box above, try to move the box and then move up
	if grid[Position{currentPosition.x, currentPosition.y - 1}].hasBox {
		i := 1
		for grid[Position{currentPosition.x, currentPosition.y - i}].hasBox {
			i++
		}
		// There's a wall above all the boxes, return the same grid
		if grid[Position{currentPosition.x, currentPosition.y - i}].hasWall {
			return grid, currentPosition
		}

		// Means there's space above, move the boxes and then move up
		for j := 2; j < i+1; j++ {
			grid[Position{currentPosition.x, currentPosition.y - j}] = PositionDetails{Position{currentPosition.x, currentPosition.y - j}, false, true, false}
		}
		grid[Position{currentPosition.x, currentPosition.y - 1}] = PositionDetails{Position{currentPosition.x, currentPosition.y - 1}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x, currentPosition.y - 1}
	}

	// No box, no wall, no space, return the same grid
	panic("Should never be here")
}

func moveDown(grid map[Position]PositionDetails, gridHeight int, currentPosition Position) (map[Position]PositionDetails, Position) {
	// If at wall or end, return the same grid
	if currentPosition.y == gridHeight-1 || grid[Position{currentPosition.x, currentPosition.y + 1}].hasWall {
		return grid, currentPosition
	}

	// If there's a space below, move down (means return the same grid)
	if !grid[Position{currentPosition.x, currentPosition.y + 1}].hasBox && !grid[Position{currentPosition.x, currentPosition.y + 1}].hasWall {
		grid[Position{currentPosition.x, currentPosition.y + 1}] = PositionDetails{Position{currentPosition.x, currentPosition.y + 1}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x, currentPosition.y + 1}
	}

	// If there's a box below, try to move the box and then move down
	if grid[Position{currentPosition.x, currentPosition.y + 1}].hasBox {
		i := 1
		for grid[Position{currentPosition.x, currentPosition.y + i}].hasBox {
			i++
		}
		// There's a wall below all the boxes, return the same grid
		if grid[Position{currentPosition.x, currentPosition.y + i}].hasWall {
			return grid, currentPosition
		}

		// Means there's space below, move the boxes and then move down
		for j := 2; j < i+1; j++ {
			grid[Position{currentPosition.x, currentPosition.y + j}] = PositionDetails{Position{currentPosition.x, currentPosition.y + j}, false, true, false}
		}
		grid[Position{currentPosition.x, currentPosition.y + 1}] = PositionDetails{Position{currentPosition.x, currentPosition.y + 1}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x, currentPosition.y + 1}
	}

	// No box, no wall, no space, return the same grid
	panic("Should never be here")
}

func moveLeft(grid map[Position]PositionDetails, currentPosition Position) (map[Position]PositionDetails, Position) {
	// If at wall or end, return the same grid
	if currentPosition.x == 0 || grid[Position{currentPosition.x - 1, currentPosition.y}].hasWall {
		return grid, currentPosition
	}

	// If there's a space to the left, move left (means return the same grid)
	if !grid[Position{currentPosition.x - 1, currentPosition.y}].hasBox && !grid[Position{currentPosition.x - 1, currentPosition.y}].hasWall {
		grid[Position{currentPosition.x - 1, currentPosition.y}] = PositionDetails{Position{currentPosition.x - 1, currentPosition.y}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x - 1, currentPosition.y}
	}

	// If there's a box to the left, try to move the box and then move left
	if grid[Position{currentPosition.x - 1, currentPosition.y}].hasBox {
		i := 1
		for grid[Position{currentPosition.x - i, currentPosition.y}].hasBox {
			i++
		}
		// There's a wall to the left of all the boxes, return the same grid
		if grid[Position{currentPosition.x - i, currentPosition.y}].hasWall {
			return grid, currentPosition
		}

		// Means there's space to the left, move the boxes and then move left
		for j := 2; j < i+1; j++ {
			grid[Position{currentPosition.x - j, currentPosition.y}] = PositionDetails{Position{currentPosition.x - j, currentPosition.y}, false, true, false}
		}
		grid[Position{currentPosition.x - 1, currentPosition.y}] = PositionDetails{Position{currentPosition.x - 1, currentPosition.y}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x - 1, currentPosition.y}
	}

	// No box, no wall, no space, return the same grid
	panic("Should never be here")
}

func moveRight(grid map[Position]PositionDetails, gridWidth int, currentPosition Position) (map[Position]PositionDetails, Position) {
	// if at wall or end, return the same grid
	if currentPosition.x == gridWidth-1 || grid[Position{currentPosition.x + 1, currentPosition.y}].hasWall {
		return grid, currentPosition
	}

	// If there's a space to the right, move right (means return the same grid)
	if !grid[Position{currentPosition.x + 1, currentPosition.y}].hasBox && !grid[Position{currentPosition.x + 1, currentPosition.y}].hasWall {
		grid[Position{currentPosition.x + 1, currentPosition.y}] = PositionDetails{Position{currentPosition.x + 1, currentPosition.y}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x + 1, currentPosition.y}
	}

	// If there's a box to the right, try to move the box and then move right
	if grid[Position{currentPosition.x + 1, currentPosition.y}].hasBox {
		i := 1
		for grid[Position{currentPosition.x + i, currentPosition.y}].hasBox {
			i++
		}
		// There's a wall to the right of all the boxes, return the same grid
		if grid[Position{currentPosition.x + i, currentPosition.y}].hasWall {
			return grid, currentPosition
		}

		// Means there's space to the right, move the boxes and then move right
		for j := 2; j <= i; j++ {
			grid[Position{currentPosition.x + j, currentPosition.y}] = PositionDetails{Position{currentPosition.x + j, currentPosition.y}, false, true, false}
		}
		grid[Position{currentPosition.x + 1, currentPosition.y}] = PositionDetails{Position{currentPosition.x + 1, currentPosition.y}, false, false, true}
		grid[currentPosition] = PositionDetails{currentPosition, false, false, false}
		return grid, Position{currentPosition.x + 1, currentPosition.y}
	}

	// No box, no wall, no space, return the same grid
	panic("Should never be here")
}

func printGrid(grid map[Position]PositionDetails, gridWidth int, gridHeight int) {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if grid[Position{x, y}].hasWall {
				fmt.Print("#")
			} else if grid[Position{x, y}].hasBox {
				fmt.Print("O")
			} else if grid[Position{x, y}].hasRobot {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
