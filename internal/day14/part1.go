package day14

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

type Robot struct {
	position  internal.Position
	velocityX int
	velocityY int
}

func PartOne(isTest bool) {
	// Some initial variables
	var fileContents []string
	var width int
	var height int
	var numSeconds int = 100

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day14/day14_example.txt")
		width = 11
		height = 7
	} else {
		fileContents = internal.ReadFileIntoArray("res/day14/day14.txt")
		width = 101
		height = 103
	}

	// Go through the file and add robots
	robots := make([]Robot, 0)
	for i := 0; i < len(fileContents); i++ {
		robots = append(robots, getRobot(fileContents[i]))
	}

	// Tick 100 seconds
	for i := 0; i < numSeconds; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j] = updateRobotPosition(robots[j], width, height)
		}
	}

	// Count robots in each quadrant
	quadrantCounts := countRobotsInEachQuadrant(robots, width, height)

	fmt.Println("Safety factor:", quadrantCounts[0]*quadrantCounts[1]*quadrantCounts[2]*quadrantCounts[3])
}

func countRobotsInEachQuadrant(robots []Robot, gridWidth int, gridHeight int) []int {
	quadrantCounts := make([]int, 4)
	for i := 0; i < len(robots); i++ {
		for j := 0; j < 4; j++ {
			if isRobotInQuadrant(robots[i], gridWidth, gridHeight, j+1) {
				quadrantCounts[j]++
			}
		}
	}
	return quadrantCounts
}

func getRobot(line string) Robot {
	positionString := strings.Split(strings.Split(line, " ")[0], "=")[1]
	velocityString := strings.Split(strings.Split(line, " ")[1], "=")[1]
	positionX := internal.GetNumFromString(strings.Split(positionString, ",")[0])
	positionY := internal.GetNumFromString(strings.Split(positionString, ",")[1])
	velocityX := internal.GetNumFromString(strings.Split(velocityString, ",")[0])
	velocityY := internal.GetNumFromString(strings.Split(velocityString, ",")[1])
	return Robot{internal.Position{X: positionX, Y: positionY}, velocityX, velocityY}
}

func updateRobotPosition(robot Robot, gridWidth int, gridHeight int) Robot {
	newRobotPositionX := robot.position.X + robot.velocityX
	newRobotPositionY := robot.position.Y + robot.velocityY

	// if new position X is greater than grid width, wrap around
	if newRobotPositionX >= gridWidth {
		newRobotPositionX = newRobotPositionX - gridWidth
	} else if newRobotPositionX < 0 {
		newRobotPositionX = gridWidth + newRobotPositionX
	}

	// if new position Y is greater than grid height, wrap around
	if newRobotPositionY >= gridHeight {
		newRobotPositionY = newRobotPositionY - gridHeight
	} else if newRobotPositionY < 0 {
		newRobotPositionY = gridHeight + newRobotPositionY
	}

	return Robot{internal.Position{X: newRobotPositionX, Y: newRobotPositionY}, robot.velocityX, robot.velocityY}
}

// If width is even, split in half
// If width is odd, remove middle row and split in half
func isRobotInQuadrant(robot Robot, gridWidth int, gridHeight int, quadrantNum int) bool {
	if quadrantNum < 1 || quadrantNum > 4 {
		panic("Invalid quadrant number")
	}

	var minX int
	var maxX int
	var minY int
	var maxY int
	var includeMiddleRow bool = gridWidth%2 == 0
	var includeMiddleColumn bool = gridHeight%2 == 0

	if quadrantNum == 1 {
		minX = 0
		minY = 0
		if includeMiddleColumn {
			maxX = gridWidth / 2
		} else {
			maxX = gridWidth/2 - 1
		}
		if includeMiddleRow {
			maxY = gridHeight / 2
		} else {
			maxY = gridHeight/2 - 1
		}
	} else if quadrantNum == 2 {
		maxX = gridWidth - 1
		minY = 0
		if includeMiddleColumn {
			minX = gridWidth / 2
		} else {
			minX = gridWidth/2 + 1
		}
		if includeMiddleRow {
			maxY = gridHeight / 2
		} else {
			maxY = gridHeight/2 - 1
		}
	} else if quadrantNum == 3 {
		minX = 0
		maxY = gridHeight - 1
		if includeMiddleColumn {
			maxX = gridWidth / 2
		} else {
			maxX = gridWidth/2 - 1
		}
		if includeMiddleRow {
			minY = gridHeight / 2
		} else {
			minY = gridHeight/2 + 1
		}
	} else {
		maxX = gridWidth
		maxY = gridHeight
		if includeMiddleColumn {
			minX = gridWidth / 2
		} else {
			minX = gridWidth/2 + 1
		}
		if includeMiddleRow {
			minY = gridHeight / 2
		} else {
			minY = gridHeight/2 + 1
		}
	}

	return robot.position.X >= minX && robot.position.X <= maxX && robot.position.Y >= minY && robot.position.Y <= maxY
}
