package day10

import (
	"aoc2024/internal"
	"fmt"
)

func PartTwo(isTest bool) {
	fmt.Println("Day 10 Part 2")
	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day10/day10_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day10/day10.txt")
	}

	fmt.Println(fileContents)
}
