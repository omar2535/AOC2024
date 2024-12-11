package day11

import (
	"aoc2024/internal"
	"fmt"
	"strconv"
	"strings"
)

func PartOne(isTest bool) {
	fmt.Println("Day 11 part 1")

	// Some initial variables
	var timesToBlink int = 75
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day11/day11_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day11/day11.txt")
	}

	var arrangement []int = internal.StringArrayToIntArray(strings.Split(fileContents[0], " "))

	for i := 0; i < timesToBlink; i++ {
		arrangement = blink(arrangement)
		fmt.Print("Progress: ", i+1, "/", timesToBlink, "\r")
	}

	fmt.Println("Number of pebbles after blinking", len(arrangement))
}

// Blinks once, returns the new arrangement
func blink(arrangement []int) []int {
	var newArrangement []int = make([]int, 0)
	for _, pebble := range arrangement {
		var newPebbles []int
		if pebble == 0 {
			newPebbles = []int{1}
		} else if hasEvenDigits(pebble) {
			firstHalf := strconv.Itoa(pebble)[:len(strconv.Itoa(pebble))/2]
			secondHalf := strconv.Itoa(pebble)[len(strconv.Itoa(pebble))/2:]
			newPebbles = []int{internal.GetNumFromString(firstHalf), internal.GetNumFromString(secondHalf)}
		} else {
			newPebbles = []int{pebble * 2024}
		}
		newArrangement = append(newArrangement, newPebbles...)
	}
	return newArrangement
}

// Checks if a pebble has even digits
func hasEvenDigits(pebble int) bool {
	var pebbleStr string = strconv.Itoa(pebble)
	return len(pebbleStr)%2 == 0
}
