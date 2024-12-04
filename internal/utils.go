package internal

import (
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// get absoulte value of a - b
func Abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// Get number from string, panics if it isn't a number
func GetNumFromString(num_s string) int {
	num, err := strconv.Atoi(num_s)
	Check(err)
	return num
}

// Remove element from array
func RemoveFromList[T any](s []T, index int) []T {
	// make a new slice so that it doesn't modify the original
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// Reads file into array of strings
func ReadFileIntoArray(filepath string) []string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var file_contents string = string(file)
	var rows []string = strings.Split(file_contents, "\n")
	return rows
}
