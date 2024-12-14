package day14

import (
	"aoc2024/internal"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

// 6586 showed me one
func PartTwo(isTest bool) {
	// Some initial variables
	var fileContents []string
	var width int
	var height int
	var numSeconds int = 10000

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
	for second := 0; second < numSeconds; second++ {
		for j := 0; j < len(robots); j++ {
			robots[j] = updateRobotPosition(robots[j], width, height)
			drawRobots(robots, width, height, second)
		}
	}
}

func drawRobots(robots []Robot, width int, height int, second int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < len(robots); i++ {
		img.Set(robots[i].position.X, robots[i].position.Y, color.RGBA{255, 0, 0, 255})
	}

	file, err := os.Create("res/day14/iteration_" + strconv.Itoa(second) + ".png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
