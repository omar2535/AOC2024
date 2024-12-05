package five

import (
	"aoc2024/internal"
	"fmt"
	"slices"
	"strings"
)

func PartOne(isTest bool) {
	fmt.Println("Day 5 Part 1")
	var rules []string
	var updates []string
	if isTest {
		rules = internal.ReadFileIntoArray("res/day5/day5_example_rules.txt")
		updates = internal.ReadFileIntoArray("res/day5/day5_example_updates.txt")
	} else {
		rules = internal.ReadFileIntoArray("res/day5/day5_rules.txt")
		updates = internal.ReadFileIntoArray("res/day5/day5_updates.txt")
	}

	// Create map of number and the numbers that should have come before it
	// This is a map of number -> []int
	rulesMap := make(map[int][]int)
	for i := 0; i < len(rules); i++ {
		ruleSplit := strings.Split(rules[i], "|")
		val := internal.GetNumFromString(ruleSplit[0])
		key := internal.GetNumFromString(ruleSplit[1])
		if rulesMap[key] != nil {
			rulesMap[key] = append(rulesMap[key], val)
		} else {
			rulesMap[key] = []int{val}
		}
	}

	// iterate through the updates to find updates that don't follow the rules
	validUpdates := make([]string, 0)
	for i := 0; i < len(updates); i++ {
		update := updates[i]
		updateNums := strings.Split(update, ",")
		isCorrect := true
		for j := 0; j < len(updateNums); j++ {
			currentNum := internal.GetNumFromString(updateNums[j])

			// last number, don't do anything
			if j == len(updateNums)-1 {
				continue
			}
			// go through each following number and check if it follows the rules
			for k := j + 1; k < len(updateNums); k++ {
				checkNum := internal.GetNumFromString(updateNums[k])
				if rulesMap[checkNum] == nil || !slices.Contains(rulesMap[checkNum], currentNum) {
					isCorrect = false
					break
				}
			}
			if !isCorrect {
				break
			}
		}
		if isCorrect {
			validUpdates = append(validUpdates, update)
		}
	}

	fmt.Println("Number of valid updates:", len(validUpdates))

	// Calculate the number of updates that are valid
	sum := 0
	for i := 0; i < len(validUpdates); i++ {
		updateNums := strings.Split(validUpdates[i], ",")
		var middleNum int
		if len(updateNums)%2 == 0 {
			middleNum = internal.GetNumFromString(updateNums[(len(updateNums) / 2)])
		} else {
			middleNum = internal.GetNumFromString(updateNums[int((float64(len(updateNums))/2.0)-0.5)])
		}
		fmt.Println("Middle number:", middleNum)
		sum += middleNum
	}

	fmt.Println("Sum of middle numbers:", sum)
}
