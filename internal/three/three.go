package three

import (
	"aoc2024/internal"
	"fmt"
	"strings"
)

func PartOne() {
	fmt.Println("Part one")
	file_contents := internal.ReadFileIntoArray("res/day3/day3.txt")
	full_file := strings.Join(file_contents, "")

	for i := 0; i < len(full_file); i++ {
	}
}
