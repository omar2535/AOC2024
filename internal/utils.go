package internal

import (
	"fmt"
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

// Get number from string, returns 0 if not a number
func GetNumFromString(num_s string) int {
	num, _ := strconv.Atoi(num_s)
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

func StringArrayToIntArray(arr []string) []int {
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = GetNumFromString(arr[i])
	}
	return ret
}

func Clone2dArray[T any](grid [][]T) [][]T {
	clonedGrid := make([][]T, len(grid))
	for i := 0; i < len(grid); i++ {
		newRow := make([]T, len(grid[i]))
		copy(newRow, grid[i])
		clonedGrid[i] = newRow
	}
	return clonedGrid
}

func ConvertStringListToGrid(inputList []string, delminator string) [][]string {
	grid := make([][]string, len(inputList))
	for i := 0; i < len(inputList); i++ {
		currentRow := inputList[i]
		grid[i] = strings.Split(currentRow, delminator)
	}
	return grid
}

// Print functions
func Print2dArray[T any](grid [][]T) {
	for i := 0; i < len(grid); i++ {
		fmt.Printf("%v ", grid[i])
		println()
	}
}
