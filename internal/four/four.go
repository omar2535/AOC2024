package four

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

func PartOne() {
	fileContents := internal.ReadFileIntoArray("res/day4/day4_original.txt")

	// size of the grid
	numRows := len(fileContents)

	// Create the grid
	grid := make([][]string, numRows)
	for i := range grid {
		grid[i] = strings.Split(fileContents[i], "")
	}

	// Iterate through the grid, find XMAS }
	numXMAS := 0
	for y := 0; y < numRows; y++ {
		for x := 0; x < len(grid[y]); x++ {
			foundNumXMAS := getNumXMAS(grid, x, y)
			numXMAS += foundNumXMAS
			if foundNumXMAS > 0 {
				fmt.Printf("%d XMAS found at (%d, %d)\n", foundNumXMAS, x, y)
			}
		}
	}
	fmt.Println(numXMAS)
}

func PartTwo() {
	fileContents := internal.ReadFileIntoArray("res/day4/day4.txt")

	// size of the grid
	numRows := len(fileContents)

	// Create the grid
	grid := make([][]string, numRows)
	for i := range grid {
		grid[i] = strings.Split(fileContents[i], "")
	}

	// Iterate through the grid, find XMAS }
	numXMAS := 0
	for y := 0; y < numRows; y++ {
		for x := 0; x < len(grid[y]); x++ {
			foundNumXMAS := getNumCrossXMas(grid, x, y)
			numXMAS += foundNumXMAS
			if foundNumXMAS > 0 {
				fmt.Printf("%d XMAS found at (%d, %d)\n", foundNumXMAS, x, y)
			}
		}
	}
	fmt.Println(numXMAS)
}

func getNumCrossXMas(grid [][]string, x int, y int) int {
	numCrossXMAS := 0
	if grid[y][x] == "A" {
		fmt.Println("Found A at", x, y)
		// Four different configurations
		// 1. S   S
		//      A
		//    M   M
		if lookUpLeft(grid, x, y, "AS") && lookDownRight(grid, x, y, "AM") && lookUpRight(grid, x, y, "AS") && lookDownLeft(grid, x, y, "AM") {
			numCrossXMAS++
		}
		// 2. M   M
		//      A
		//    S   S
		if lookUpLeft(grid, x, y, "AM") && lookDownRight(grid, x, y, "AS") && lookUpRight(grid, x, y, "AM") && lookDownLeft(grid, x, y, "AS") {
			numCrossXMAS++
		}
		// 3. S   M
		//      A
		//    S   M
		if lookUpLeft(grid, x, y, "AS") && lookDownRight(grid, x, y, "AM") && lookUpRight(grid, x, y, "AM") && lookDownLeft(grid, x, y, "AS") {
			numCrossXMAS++
		}
		// 4. M   S
		//      A
		//    M   S
		if lookUpLeft(grid, x, y, "AM") && lookDownRight(grid, x, y, "AS") && lookUpRight(grid, x, y, "AS") && lookDownLeft(grid, x, y, "AM") {
			numCrossXMAS++
		}
	}
	return numCrossXMAS
}

func getNumXMAS(grid [][]string, x int, y int) int {
	keyword := "XMAS"
	numXMAS := 0
	if grid[y][x] == string(keyword[0]) {
		if lookUp(grid, x, y, keyword) {
			numXMAS++
		}
		if lookDown(grid, x, y, keyword) {
			numXMAS++
		}
		if lookLeft(grid, x, y, keyword) {
			numXMAS++
		}
		if lookRight(grid, x, y, keyword) {
			numXMAS++
		}
		if lookUpLeft(grid, x, y, keyword) {
			numXMAS++
		}
		if lookUpRight(grid, x, y, keyword) {
			numXMAS++
		}
		if lookDownLeft(grid, x, y, keyword) {
			numXMAS++
		}
		if lookDownRight(grid, x, y, keyword) {
			numXMAS++
		}
	}
	return numXMAS
}

func lookUp(grid [][]string, x int, y int, keyword string) bool {
	// impossible if there isn't enough height to look up
	if y < len(keyword)-1 {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y-i][x] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookDown(grid [][]string, x int, y int, keyword string) bool {
	height := len(grid)
	if height-y < len(keyword) {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y+i][x] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookLeft(grid [][]string, x int, y int, keyword string) bool {
	if x < len(keyword)-1 {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y][x-i] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookRight(grid [][]string, x int, y int, keyword string) bool {
	width := len(grid[0])
	if width-x < len(keyword) {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y][x+i] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookUpLeft(grid [][]string, x int, y int, keyword string) bool {
	if y < len(keyword)-1 || x < len(keyword)-1 {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y-i][x-i] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookUpRight(grid [][]string, x int, y int, keyword string) bool {
	width := len(grid[0])
	if y < len(keyword)-1 || width-x < len(keyword) {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y-i][x+i] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookDownLeft(grid [][]string, x int, y int, keyword string) bool {
	height := len(grid)
	if height-y < len(keyword) || x < len(keyword)-1 {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y+i][x-i] != string(keyword[i]) {
			return false
		}
	}
	return true
}

func lookDownRight(grid [][]string, x int, y int, keyword string) bool {
	width := len(grid[0])
	height := len(grid)
	if height-y < len(keyword) || width-x < len(keyword) {
		return false
	}
	for i := 0; i < len(keyword); i++ {
		if grid[y+i][x+i] != string(keyword[i]) {
			return false
		}
	}
	return true
}
