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

	fmt.Println(robots)
	fmt.Println(fileContents, width, height)
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
