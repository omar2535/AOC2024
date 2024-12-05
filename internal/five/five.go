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
		updateNumsInt := internal.StringArrayToIntArray(updateNums)
		isCorrect := isCorrect(rulesMap, updateNumsInt)
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

func PartTwo(isTest bool) {
	fmt.Println("Day 5 Part 2")
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
	validUpdates := make([][]int, 0)
	invalidUpdates := make([][]int, 0)
	for i := 0; i < len(updates); i++ {
		update := updates[i]
		updateNums := strings.Split(update, ",")
		updateNumsInt := internal.StringArrayToIntArray(updateNums)
		isCorrect := isCorrect(rulesMap, updateNumsInt)
		if isCorrect {
			validUpdates = append(validUpdates, updateNumsInt)
		} else {
			invalidUpdates = append(invalidUpdates, updateNumsInt)
		}
	}

	fmt.Println("Number of valid updates:", len(validUpdates))
	fmt.Println("Number of invalid updates:", len(invalidUpdates))

	reorderedUpdates := make([][]int, 0)
	for i := 0; i < len(invalidUpdates); i++ {
		reorderedUpdate := getReorderedUpdate(rulesMap, invalidUpdates[i])
		reorderedUpdates = append(reorderedUpdates, reorderedUpdate)
	}

	// Calculate the number of updates that are valid
	sum := 0
	for i := 0; i < len(reorderedUpdates); i++ {
		updateNums := reorderedUpdates[i]
		var middleNum int
		if len(updateNums)%2 == 0 {
			middleNum = updateNums[(len(updateNums) / 2)]
		} else {
			middleNum = updateNums[int((float64(len(updateNums))/2.0)-0.5)]
		}
		fmt.Println("Middle number:", middleNum)
		sum += middleNum
	}

	fmt.Println("Sum of middle numbers:", sum)
}

func getReorderedUpdate(rulesMap map[int][]int, updates []int) []int {
	newArray := make([]int, len(updates))
	for i := 0; i < len(updates); i++ {
		numDependants := getNumDependants(rulesMap, updates, updates[i])
		newArray[len(updates)-numDependants-1] = updates[i]
	}
	return newArray
}

func getNumDependants(rulesMap map[int][]int, updates []int, num int) int {
	numDependants := 0
	for i := 0; i < len(updates); i++ {
		if slices.Contains(rulesMap[updates[i]], num) {
			numDependants++
		}
	}
	return numDependants
}

func isCorrect(rulesMap map[int][]int, update []int) bool {
	isCorrect := true
	for j := 0; j < len(update); j++ {
		currentNum := update[j]

		// last number, don't do anything
		if j == len(update)-1 {
			continue
		}
		// go through each following number and check if it follows the rules
		for k := j + 1; k < len(update); k++ {
			checkNum := update[k]
			if rulesMap[checkNum] == nil || !slices.Contains(rulesMap[checkNum], currentNum) {
				isCorrect = false
				break
			}
		}
		if !isCorrect {
			break
		}
	}
	return isCorrect
}
