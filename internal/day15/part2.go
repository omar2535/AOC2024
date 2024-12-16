package day15

import (
	"aoc2024/internal"
	"fmt"
)

type Box struct {
	leftPosition  Position
	rightPosition Position
}

type PositionDetailsNew struct {
	position Position
	hasWall  bool
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
	var grid map[Position]PositionDetailsNew = make(map[Position]PositionDetailsNew)
	var boxes []Box = make([]Box, 0)
	var gridWidth int = len(fileContents[0]) * 2
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
						grid[Position{x * 2, gridHeight}] = PositionDetailsNew{Position{x * 2, len(grid)}, true}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetailsNew{Position{x*2 + 1, len(grid)}, true}
					} else if c == '.' {
						grid[Position{x * 2, gridHeight}] = PositionDetailsNew{Position{x * 2, len(grid)}, false}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetailsNew{Position{x*2 + 1, len(grid)}, false}
					} else if c == 'O' {
						grid[Position{x * 2, gridHeight}] = PositionDetailsNew{Position{x * 2, len(grid)}, false}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetailsNew{Position{x*2 + 1, len(grid)}, false}
						box := Box{Position{x * 2, gridHeight}, Position{x*2 + 1, gridHeight}}
						boxes = append(boxes, box)
					} else if c == '@' {
						grid[Position{x * 2, gridHeight}] = PositionDetailsNew{Position{x * 2, len(grid)}, false}
						grid[Position{x*2 + 1, gridHeight}] = PositionDetailsNew{Position{x*2 + 1, len(grid)}, false}
						startingPosition = Position{x * 2, gridHeight}
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
	fmt.Println("Initial position: ", currentPosition)
	fmt.Println("Initial Boxes: ", boxes)
	printGridNew(grid, boxes, currentPosition, gridWidth, gridHeight)
	for _, move := range moves {
		grid, boxes, currentPosition = doMoveNew(grid, boxes, gridWidth, gridHeight, currentPosition, move)
		// fmt.Println("Move: ", move)
		// fmt.Println("Current position: ", currentPosition)
		// fmt.Println("Boxes: ", boxes)
		// printGridNew(grid, boxes, currentPosition, gridWidth, gridHeight)
	}

	// Calculate the score
	fmt.Println("Score: ", calculateBoxScore(boxes, gridWidth, gridHeight))
}

func calculateBoxScore(boxes []Box, gridWidth int, gridHeight int) int {
	score := 0
	for _, box := range boxes {
		// closest_x := min(box.leftPosition.x, gridWidth-1-box.rightPosition.x)
		// closest_y := min(box.leftPosition.y, gridHeight-1-box.rightPosition.y)
		// score += 100*closest_y + closest_x
		score += 100*box.leftPosition.y + box.leftPosition.x
	}
	return score
}

func doMoveNew(grid map[Position]PositionDetailsNew, boxes []Box, gridWidth int, gridHeight int, currentPosition Position, move string) (map[Position]PositionDetailsNew, []Box, Position) {
	if move == "^" {
		grid, boxes, currentPosition = moveUpNew(grid, boxes, currentPosition)
	} else if move == "v" {
		grid, boxes, currentPosition = moveDownNew(grid, boxes, currentPosition, gridHeight)
	} else if move == "<" {
		grid, boxes, currentPosition = moveLeftNew(grid, boxes, currentPosition)
	} else if move == ">" {
		grid, boxes, currentPosition = moveRightNew(grid, boxes, currentPosition, gridWidth)
	}

	return grid, boxes, currentPosition
}

func moveRightNew(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position, gridWidth int) (map[Position]PositionDetailsNew, []Box, Position) {
	if currentPosition.x+1 >= gridWidth {
		return grid, boxes, currentPosition
	}

	if grid[Position{currentPosition.x + 1, currentPosition.y}].hasWall {
		return grid, boxes, currentPosition
	}

	if getBoxAtPosition(boxes, Position{currentPosition.x + 1, currentPosition.y}) != (Box{}) {
		// Move the box
		affectedBoxes, shouldMove := getAffectedBoxesRight(grid, boxes, Position{currentPosition.x + 1, currentPosition.y}, gridWidth)
		if shouldMove {
			for _, box := range affectedBoxes {
				boxIndex := getIndexOfBox(boxes, box)
				if boxIndex == -1 {
					continue
				}

				// Update the boxes positions
				boxes[boxIndex].leftPosition.x++
				boxes[boxIndex].rightPosition.x++
			}
			// Update robot position
			currentPosition.x++
		}
	} else {
		// Update robot position because there's no box
		currentPosition.x++
	}

	// Return
	return grid, boxes, currentPosition
}

func getAffectedBoxesRight(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position, gridWidth int) ([]Box, bool) {
	affectedBoxes := make([]Box, 0)

	// If there is a wall above, return false
	if grid[Position{currentPosition.x, currentPosition.y}].hasWall || currentPosition.x >= gridWidth-1 {
		return affectedBoxes, false
	}

	// if there is space, return true
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) == (Box{}) {
		return affectedBoxes, true
	}

	// if there is a box, check if there's space behind it
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) != (Box{}) {
		currBox := getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y})
		rightPosition := currBox.rightPosition
		affectedBoxesRight, affectedBoxesRightMoves := getAffectedBoxesRight(grid, boxes, Position{rightPosition.x + 1, rightPosition.y}, gridWidth)

		affectedBoxes = append(affectedBoxes, currBox)
		affectedBoxes = append(affectedBoxes, affectedBoxesRight...)
		affectedBoxes = removeDuplicateBoxes(affectedBoxes)
		return affectedBoxes, affectedBoxesRightMoves
	}

	return affectedBoxes, false
}

func moveLeftNew(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position) (map[Position]PositionDetailsNew, []Box, Position) {
	if currentPosition.x-1 < 0 {
		return grid, boxes, currentPosition
	}

	if grid[Position{currentPosition.x - 1, currentPosition.y}].hasWall {
		return grid, boxes, currentPosition
	}

	if getBoxAtPosition(boxes, Position{currentPosition.x - 1, currentPosition.y}) != (Box{}) {
		// Move the box
		affectedBoxes, shouldMove := getAffectedBoxesLeft(grid, boxes, Position{currentPosition.x - 1, currentPosition.y})
		if shouldMove {
			for _, box := range affectedBoxes {
				boxIndex := getIndexOfBox(boxes, box)
				if boxIndex == -1 {
					continue
				}

				// Update the boxes positions
				boxes[boxIndex].leftPosition.x--
				boxes[boxIndex].rightPosition.x--
			}
			// Update robot position
			currentPosition.x--
		}
	} else {
		// Update robot position because there's no box
		currentPosition.x--
	}

	// Return
	return grid, boxes, currentPosition
}

func getAffectedBoxesLeft(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position) ([]Box, bool) {
	affectedBoxes := make([]Box, 0)

	// If there is a wall above, return false
	if grid[Position{currentPosition.x, currentPosition.y}].hasWall || currentPosition.x <= 0 {
		return affectedBoxes, false
	}

	// if there is space, return true
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) == (Box{}) {
		return affectedBoxes, true
	}

	// if there is a box, check if there's space behind it
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) != (Box{}) {
		currBox := getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y})
		leftPosition := currBox.leftPosition
		affectedBoxesLeft, affectedBoxesLeftMoves := getAffectedBoxesLeft(grid, boxes, Position{leftPosition.x - 1, leftPosition.y})

		affectedBoxes = append(affectedBoxes, currBox)
		affectedBoxes = append(affectedBoxes, affectedBoxesLeft...)
		affectedBoxes = removeDuplicateBoxes(affectedBoxes)
		return affectedBoxes, affectedBoxesLeftMoves
	}

	return affectedBoxes, false
}

func moveDownNew(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position, gridHeight int) (map[Position]PositionDetailsNew, []Box, Position) {
	if currentPosition.y+1 >= gridHeight {
		return grid, boxes, currentPosition
	}

	if grid[Position{currentPosition.x, currentPosition.y + 1}].hasWall {
		return grid, boxes, currentPosition
	}

	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y + 1}) != (Box{}) {
		// Move the box
		affectedBoxes, shouldMove := getAffectedBoxesDown(grid, boxes, Position{currentPosition.x, currentPosition.y + 1}, gridHeight)
		if shouldMove {
			for _, box := range affectedBoxes {
				boxIndex := getIndexOfBox(boxes, box)
				if boxIndex == -1 {
					continue
				}

				// Update the boxes positions
				boxes[boxIndex].leftPosition.y++
				boxes[boxIndex].rightPosition.y++
			}

			// Update robot position
			currentPosition.y++
		}
	} else {
		// Update robot position because there's no box
		currentPosition.y++
	}

	// Return
	return grid, boxes, currentPosition
}

func getAffectedBoxesDown(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position, gridHeight int) ([]Box, bool) {
	affectedBoxes := make([]Box, 0)

	// If there is a wall above, return false
	if grid[Position{currentPosition.x, currentPosition.y}].hasWall || currentPosition.y >= gridHeight-1 {
		return affectedBoxes, false
	}

	// if there is space, return true
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) == (Box{}) {
		return affectedBoxes, true
	}

	// if there is a box, check if there's space behind it
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) != (Box{}) {
		currBox := getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y})
		leftPosition := currBox.leftPosition
		rightPosition := currBox.rightPosition
		affectedBoxesLeft, affectedBoxesLeftMoves := getAffectedBoxesDown(grid, boxes, Position{leftPosition.x, leftPosition.y + 1}, gridHeight)
		affectedBoxesRight, affectedBoxesRightMoves := getAffectedBoxesDown(grid, boxes, Position{rightPosition.x, rightPosition.y + 1}, gridHeight)

		affectedBoxes = append(affectedBoxes, currBox)
		affectedBoxes = append(affectedBoxes, affectedBoxesLeft...)
		affectedBoxes = append(affectedBoxes, affectedBoxesRight...)
		affectedBoxes = removeDuplicateBoxes(affectedBoxes)
		return affectedBoxes, affectedBoxesLeftMoves && affectedBoxesRightMoves
	}

	return affectedBoxes, false
}

func moveUpNew(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position) (map[Position]PositionDetailsNew, []Box, Position) {
	if currentPosition.y-1 < 0 {
		return grid, boxes, currentPosition
	}

	if grid[Position{currentPosition.x, currentPosition.y - 1}].hasWall {
		return grid, boxes, currentPosition
	}

	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y - 1}) != (Box{}) {
		// Move the box
		affectedBoxes, shouldMove := getAffectedBoxesUp(grid, boxes, Position{currentPosition.x, currentPosition.y - 1})
		if shouldMove {
			for _, box := range affectedBoxes {
				boxIndex := getIndexOfBox(boxes, box)
				if boxIndex == -1 {
					continue
				}

				// Update the boxes positions
				boxes[boxIndex].leftPosition.y--
				boxes[boxIndex].rightPosition.y--
			}

			// Update robot position
			currentPosition.y--
		}
	} else {
		// Update robot position because there's no box
		currentPosition.y--
	}

	// Return
	return grid, boxes, currentPosition
}

func getAffectedBoxesUp(grid map[Position]PositionDetailsNew, boxes []Box, currentPosition Position) ([]Box, bool) {
	affectedBoxes := make([]Box, 0)

	// If there is a wall above, return false
	if grid[Position{currentPosition.x, currentPosition.y}].hasWall || currentPosition.y <= 0 {
		return affectedBoxes, false
	}

	// if there is space, return true
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) == (Box{}) {
		return affectedBoxes, true
	}

	// if there is a box, check if there's space behind it
	if getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y}) != (Box{}) {
		currBox := getBoxAtPosition(boxes, Position{currentPosition.x, currentPosition.y})
		leftPosition := currBox.leftPosition
		rightPosition := currBox.rightPosition
		affectedBoxesLeft, affectedBoxesLeftMoves := getAffectedBoxesUp(grid, boxes, Position{leftPosition.x, leftPosition.y - 1})
		affectedBoxesRight, affectedBoxesRightMoves := getAffectedBoxesUp(grid, boxes, Position{rightPosition.x, rightPosition.y - 1})

		affectedBoxes = append(affectedBoxes, currBox)
		affectedBoxes = append(affectedBoxes, affectedBoxesLeft...)
		affectedBoxes = append(affectedBoxes, affectedBoxesRight...)
		affectedBoxes = removeDuplicateBoxes(affectedBoxes)
		return affectedBoxes, affectedBoxesLeftMoves && affectedBoxesRightMoves
	}

	return affectedBoxes, false
}

/*FUNCTIONS FOR PRINTING / UTILS*/
func printGridNew(grid map[Position]PositionDetailsNew, boxes []Box, robotPosition Position, gridWidth int, gridHeight int) {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			if grid[Position{x, y}].hasWall {
				fmt.Print("#")
			} else if getBoxAtPosition(boxes, Position{x, y}) != (Box{}) {
				box := getBoxAtPosition(boxes, Position{x, y})
				if box.leftPosition.x == x {
					fmt.Print("[")
				} else if box.rightPosition.x == x {
					fmt.Print("]")
				}
			} else if (Position{x, y}) == robotPosition {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getBoxAtPosition(boxes []Box, position Position) Box {
	for _, box := range boxes {
		if box.leftPosition == position || box.rightPosition == position {
			return box
		}
	}
	return Box{}
}

func getIndexOfBox(boxes []Box, box Box) int {
	for i, b := range boxes {
		if b.leftPosition == box.leftPosition && b.rightPosition == box.rightPosition {
			return i
		}
	}
	return -1
}

func removeDuplicateBoxes(boxes []Box) []Box {
	uniqueBoxes := make([]Box, 0)
	for _, box := range boxes {
		if !containsBox(uniqueBoxes, box) {
			uniqueBoxes = append(uniqueBoxes, box)
		}
	}
	return uniqueBoxes
}

func containsBox(boxes []Box, box Box) bool {
	for _, b := range boxes {
		if b == box {
			return true
		}
	}
	return false
}
