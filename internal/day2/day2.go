package day2

import (
	"aoc2024/internal"
	"fmt"
	"os"
	"strings"
)

func PartOne() {
	file, err := os.ReadFile("res/day2/day2.txt")
	// file, err := os.ReadFile("res/day2/day2_example.txt")

	if err != nil {
		panic(err)
	}

	// Convert file to string
	var file_contents string = string(file)

	// Split string into individual rows
	var rows []string = strings.Split(file_contents, "\n")

	// Initialize the number of safe levels
	num_safe := 0

	// parse each line
	for i := 0; i < len(rows); i++ {
		split_nums_s := strings.Split(rows[i], " ")
		split_length := len(split_nums_s)

		// break out when we encounter the last newline
		if split_length < 2 {
			break
		}

		// special cases: length of row is 2, only check diff
		//                length of row is 1, just return true
		if split_length == 2 {
			num1 := internal.GetNumFromString(split_nums_s[0])
			num2 := internal.GetNumFromString(split_nums_s[1])

			if internal.Abs(num1, num2) <= 3 && internal.Abs(num1, num2) >= 1 {
				num_safe++
				continue
			}
		}

		// normal case: length of row is 3 or more
		// 				go through each number in the row
		is_safe := isSafe(split_nums_s, split_length)

		// if safe, increment the number of safe levels
		if is_safe {
			fmt.Println("Row: ", i+1, " is safe")
			num_safe++
		}
	}
	fmt.Println("Number safe: ", num_safe)
}

func PartTwo() {
	file, err := os.ReadFile("res/day2/day2.txt")
	// file, err := os.ReadFile("res/day2/day2_example.txt")

	if err != nil {
		panic(err)
	}

	// Convert file to string
	var file_contents string = string(file)

	// Split string into individual rows
	var rows []string = strings.Split(file_contents, "\n")

	// Initialize the number of safe levels
	num_safe := 0

	// parse each line
	for i := 0; i < len(rows); i++ {
		split_nums_s := strings.Split(rows[i], " ")
		split_length := len(split_nums_s)

		// break out when we encounter the last newline
		if split_length < 2 {
			break
		}

		// special cases: length of row is 2, only check diff
		//                length of row is 1, just return true
		if split_length == 2 {
			num1 := internal.GetNumFromString(split_nums_s[0])
			num2 := internal.GetNumFromString(split_nums_s[1])

			if internal.Abs(num1, num2) <= 3 && internal.Abs(num1, num2) >= 1 {
				num_safe++
				continue
			}
		}

		// normal case: length of row is 3 or more
		// First, check if it's safe with the default row
		is_safe := isSafe(split_nums_s, split_length)

		// fmt.Println("split_num: ", split_nums_s)
		if !is_safe {
			for j := 0; j < split_length; j++ {
				removed_split_nums_s := internal.RemoveFromList(split_nums_s, j)
				// fmt.Println("split_num_removed: ", removed_split_nums_s, "index: ", j)
				if isSafe(removed_split_nums_s, len(removed_split_nums_s)) {
					is_safe = true
					break
				}
			}
		}

		// if safe, increment the number of safe levels
		if is_safe {
			fmt.Println("Row: ", i+1, " is safe")
			num_safe++
		} else {
			fmt.Println("Row: ", i+1, " is not safe")
		}
	}
	fmt.Println("Number safe: ", num_safe)
}

func isSafe(split_nums_s []string, split_length int) bool {
	for j := 2; j < split_length; j++ {
		current_num := internal.GetNumFromString(split_nums_s[j])
		previous_num := internal.GetNumFromString(split_nums_s[j-1])
		previous_previous_num := internal.GetNumFromString(split_nums_s[j-2])
		// fmt.Println("current_num: ", current_num, " previous_num: ", previous_num, " previous_previous_num: ", previous_previous_num)

		// check if the difference between curr_num and previous_num is safe
		if internal.Abs(current_num, previous_num) > 3 || internal.Abs(current_num, previous_num) < 1 {
			return false
		} else if internal.Abs(previous_num, previous_previous_num) > 3 || internal.Abs(previous_num, previous_previous_num) < 1 {
			return false
		}

		// check ioncrease/decrease/same
		previous_pattern := previous_num - previous_previous_num
		current_pattern := current_num - previous_num
		if previous_pattern > 0 && current_pattern < 0 {
			return false
		} else if previous_pattern < 0 && current_pattern > 0 {
			return false
		} else if previous_pattern == 0 || current_pattern == 0 {
			return false
		}
	}
	return true
}
