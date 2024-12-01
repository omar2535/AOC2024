package one

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// day1 part 1
func PartOne() {
	file, err := os.ReadFile("res/day1/day1part1.txt")

	if err != nil {
		panic(err)
	}

	// Convert file to string
	var file_contents string = string(file)

	// Split string into individual rows
	var rows []string = strings.Split(file_contents, "\n")

	// Get two arrays of integers
	var first_array []int = make([]int, len(rows))
	var second_array []int = make([]int, len(rows))

	for i := 0; i < len(rows); i++ {
		split_length := len(strings.Split(rows[i], "   "))

		// break out when we encounter the last newline
		if split_length != 2 {
			break
		}

		first_num_s := strings.Split(rows[i], "   ")[0]
		second_num_s := strings.Split(rows[i], "   ")[1]

		first_num, err_first_num := strconv.Atoi(first_num_s)
		second_num, err_second_num := strconv.Atoi(second_num_s)

		if err_first_num != nil || err_second_num != nil {
			panic("Saw not a number")
		}
		first_array[i] = first_num
		second_array[i] = second_num
	}

	// Calculate the differences after sorting
	slices.Sort(first_array)
	slices.Sort(second_array)

	var distance int = 0
	for i := 0; i < len(first_array); i++ {
		first_num := first_array[i]
		second_num := second_array[i]

		if first_num > second_num {
			distance += first_num - second_num
		} else {
			distance += second_num - first_num
		}
	}

	fmt.Println(distance)
}

func SimilarityScore() {
	file, err := os.ReadFile("res/day1/day1part1.txt")

	if err != nil {
		panic(err)
	}

	// Convert file to string
	var file_contents string = string(file)

	// Split string into individual rows
	var rows []string = strings.Split(file_contents, "\n")

	// Get two arrays of integers
	var first_array []int = make([]int, len(rows))
	var second_array []int = make([]int, len(rows))

	for i := 0; i < len(rows); i++ {
		split_length := len(strings.Split(rows[i], "   "))

		// break out when we encounter the last newline
		if split_length != 2 {
			break
		}

		first_num_s := strings.Split(rows[i], "   ")[0]
		second_num_s := strings.Split(rows[i], "   ")[1]

		first_num, err_first_num := strconv.Atoi(first_num_s)
		second_num, err_second_num := strconv.Atoi(second_num_s)

		if err_first_num != nil || err_second_num != nil {
			panic("Saw not a number")
		}
		first_array[i] = first_num
		second_array[i] = second_num
	}

	// Calculate how many times each number in first array appears in second array
	var occurence_map map[int]int = make(map[int]int)
	for i := 0; i < len(second_array); i++ {
		occurence_map[second_array[i]] += 1
	}

	// iterate through first array and check how many times each number appears in second array
	var similarity_score int = 0
	for i := 0; i < len(first_array); i++ {
		similarity_score += first_array[i] * occurence_map[first_array[i]]
	}

	fmt.Println(similarity_score)
}
