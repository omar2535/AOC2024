package internal

import "strconv"

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
func RemoveFromList(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
