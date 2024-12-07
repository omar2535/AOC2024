package day7

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

type equation struct {
	targetNumber int
	numbers      []int
}

func PartOne(isTest bool) {
	fmt.Println("Day 7 part 1")

	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day7/day7_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day7/day7.txt")
	}

	// Create slice of equations
	equations := make([]equation, len(fileContents))
	for i := 0; i < len(fileContents); i++ {
		splitEquation := strings.Split(fileContents[i], ": ")
		targetNumber := internal.GetNumFromString(splitEquation[0])
		splitNumbersString := strings.Split(splitEquation[1], " ")
		splitNumbers := internal.StringArrayToIntArray(splitNumbersString)
		currentEquation := equation{targetNumber: targetNumber, numbers: splitNumbers}
		equations[i] = currentEquation
	}

	// Go through each equation
	sum := 0
	for i := 0; i < len(equations); i++ {
		fmt.Println(equations[i])
		currentEquation := equations[i]
		if canHitTargetNumber(currentEquation.numbers, currentEquation.targetNumber) {
			fmt.Println(i, "hit the target number")
			sum += currentEquation.targetNumber
		}
	}
	fmt.Println("Total sum: ", sum)
}

func canHitTargetNumber(remainingNumbers []int, remainder int) bool {
	// base case: only 1 thing left in the array, return it
	if len(remainingNumbers) == 0 {
		return remainder == 0
	} else if remainder < 0 {
		// means it was overshot already
		return false
	} else {
		currentNumber := remainingNumbers[len(remainingNumbers)-1]
		newRemainingNumbers := remainingNumbers[0 : len(remainingNumbers)-1]
		// if not divisible, only do the subtraction
		if remainder%currentNumber != 0 {
			return canHitTargetNumber(newRemainingNumbers, remainder-currentNumber)
		} else {
			return canHitTargetNumber(newRemainingNumbers, remainder/currentNumber) ||
				canHitTargetNumber(newRemainingNumbers, remainder-currentNumber)
		}
	}
}
