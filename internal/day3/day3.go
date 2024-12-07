package day3

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

func PartOne() {
	fmt.Println("Part one")
	file_contents := internal.ReadFileIntoArray("res/day3/day3.txt")
	full_file := strings.Join(file_contents, "")

	mul_substring := "mul("
	search_string := strings.Clone(full_file)

	// Set the initial search string to the first occurence of our matching substring
	index_of_first_substring := strings.Index(search_string, mul_substring)
	search_string = search_string[index_of_first_substring:]

	sum := 0

	for strings.Contains(search_string, mul_substring) {
		var index_of_start int = strings.Index(search_string, mul_substring)
		var index_of_starting_bracket int = index_of_start + len(mul_substring) - 1
		var index_of_closing_bracket int = strings.Index(search_string[index_of_starting_bracket:], ")") + index_of_starting_bracket
		var index_of_next_mul int = strings.Index(search_string[index_of_starting_bracket:], mul_substring) + index_of_starting_bracket
		var has_next_mul bool = strings.Contains(search_string[index_of_starting_bracket:], mul_substring)

		// case when the next closing bracket is after the next mul, that means the current mul doesn't have a closing bracket
		if has_next_mul && index_of_next_mul < index_of_closing_bracket {
			search_string = search_string[index_of_next_mul:]
			continue
		}

		// case when there isn't a closing bracket, then just make the search string empty
		if index_of_closing_bracket == -1 {
			search_string = ""
			break
		}

		// When we have the index of the substring and the index of the closing bracket:
		// - extract the inner portion
		// - determine if it's valid by splitting it and int then multiply
		arguments_string := search_string[index_of_starting_bracket+1 : index_of_closing_bracket]
		arguments_array := strings.Split(arguments_string, ",")
		if len(arguments_array) == 2 {
			first_argument := internal.GetNumFromString(arguments_array[0])
			second_argument := internal.GetNumFromString(arguments_array[1])
			multiplied := first_argument * second_argument
			fmt.Printf("%d x %d = %d\n", first_argument, second_argument, multiplied)
			sum += multiplied
		}
		search_string = search_string[index_of_next_mul:]
	}

	fmt.Println("The total is: ", sum)
}

func PartTwo() {
	file_contents := internal.ReadFileIntoArray("res/day3/day3.txt")
	full_file := strings.Join(file_contents, "")

	mul_substring := "mul("
	do_substring := "do()"
	dont_substring := "don't()"
	search_string := strings.Clone(full_file)

	// Set the initial search string to the first occurence of our matching substring
	index_of_first_substring := strings.Index(search_string, mul_substring)
	search_string = search_string[index_of_first_substring:]

	sum := 0
	should_mul := true

	for strings.Contains(search_string, mul_substring) {
		var index_of_start int = strings.Index(search_string, mul_substring)
		var index_of_starting_bracket int = index_of_start + len(mul_substring) - 1
		var index_of_closing_bracket int = strings.Index(search_string[index_of_starting_bracket:], ")") + index_of_starting_bracket
		var index_of_next_mul int = strings.Index(search_string[index_of_starting_bracket:], mul_substring) + index_of_starting_bracket
		var has_next_mul bool = strings.Contains(search_string[index_of_starting_bracket:], mul_substring)
		var new_should_mul bool = should_mul

		// when we see do() that's before the next mul, toggle on do
		// when we see don't() that's before the next mul, toggle off should_mul
		// if we see both, get the larger one
		if has_next_mul {
			index_of_do := strings.Index(search_string[:index_of_next_mul], do_substring)
			index_of_dont := strings.Index(search_string[:index_of_next_mul], dont_substring)
			if index_of_do > index_of_dont {
				new_should_mul = true
			} else if index_of_dont > index_of_do {
				new_should_mul = false
			} else {
				// do nothing
			}
		}

		// case when the next closing bracket is after the next mul, that means the current mul doesn't have a closing bracket
		if has_next_mul && index_of_next_mul < index_of_closing_bracket {
			search_string = search_string[index_of_next_mul:]
			should_mul = new_should_mul
			continue
		}

		// case when there isn't a closing bracket, then just make the search string empty
		if index_of_closing_bracket == -1 {
			search_string = ""
			break
		}

		if should_mul {
			// When we have the index of the substring and the index of the closing bracket:
			// - extract the inner portion
			// - determine if it's valid by splitting it and int then multiply
			arguments_string := search_string[index_of_starting_bracket+1 : index_of_closing_bracket]
			arguments_array := strings.Split(arguments_string, ",")
			if len(arguments_array) == 2 {
				first_argument := internal.GetNumFromString(arguments_array[0])
				second_argument := internal.GetNumFromString(arguments_array[1])
				multiplied := first_argument * second_argument
				fmt.Printf("%d x %d = %d\n", first_argument, second_argument, multiplied)
				sum += multiplied
			}
		}
		search_string = search_string[index_of_next_mul:]
		should_mul = new_should_mul
	}

	fmt.Println("The total is: ", sum)
}
