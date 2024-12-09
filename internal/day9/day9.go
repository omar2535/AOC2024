package day9

import (
	"aoc2024/internal"
	"fmt"
	"strconv"
	"strings"
)

func PartOne(isTest bool) {
	fmt.Println("Day 9 part 1")

	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day9/day9_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day9/day9.txt")
	}

	var diskMap []int = internal.StringArrayToIntArray(strings.Split(fileContents[0], ""))
	var blockRepresentation []string = getBlockRepresentationFromDiskMap(diskMap)
	fmt.Println(blockRepresentation)
}

// Returns list of characters of the block representation
func getBlockRepresentationFromDiskMap(dispMap []int) []string {
	builtString := ""
	onFreeSpace := false
	currentIndex := 0
	for _, i := range dispMap {
		if onFreeSpace {
			builtString += strings.Repeat(".", i)
		} else {
			builtString += strings.Repeat(strconv.Itoa(currentIndex), i)
			currentIndex++
		}
		// toggle whether we are on free space or on a file
		onFreeSpace = !onFreeSpace
	}
	builtStringSlice := strings.Split(builtString, "")
	return builtStringSlice
}
