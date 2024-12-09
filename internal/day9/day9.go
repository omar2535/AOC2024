package day9

import (
	"aoc2024/internal"
	"fmt"
	"strconv"
	"strings"
)

func PartOne(isTest bool) {
	fmt.Println("Day 9 part 1")

	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day9/day9_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day9/day9.txt")
	}

	var diskMap []int = internal.StringArrayToIntArray(strings.Split(fileContents[0], ""))
	var blockRepresentation []string = getBlockRepresentationFromDiskMap(diskMap)

	// Create 2 queues:
	//   1 for the position of empty space
	//   1 for the position of file blocks in reverse
	var emptySpaceQueue []int = getPositionsOfEmptySpace(blockRepresentation)
	var fileBlockReverseQueue []int = getPositionsOfFileBlocksReversed(blockRepresentation)

	fmt.Println("initial block representation", blockRepresentation)

	// Rearranges until the emptySpace index >= fileBlock index
	currEmptySpaceIndex := emptySpaceQueue[0]
	currFileBlockIndex := fileBlockReverseQueue[0]
	for currEmptySpaceIndex < currFileBlockIndex {
		blockRepresentation = rearrangeOnce(currEmptySpaceIndex, currFileBlockIndex, blockRepresentation)
		//fmt.Println(blockRepresentation)
		emptySpaceQueue = emptySpaceQueue[1:]
		fileBlockReverseQueue = fileBlockReverseQueue[1:]
		if len(emptySpaceQueue) == 0 || len(fileBlockReverseQueue) == 0 {
			break
		}
		currEmptySpaceIndex = emptySpaceQueue[0]
		currFileBlockIndex = fileBlockReverseQueue[0]
	}

	fmt.Println("final block representation", blockRepresentation)
	fmt.Println("Checksum", computeChecksum(blockRepresentation))
}

// Returns list of characters of the block representation
func getBlockRepresentationFromDiskMap(dispMap []int) []string {
	blockRepresentation := make([]string, 0)
	onFreeSpace := false
	currentIndex := 0
	for _, i := range dispMap {
		for j := 0; j < i; j++ {
			if onFreeSpace {
				blockRepresentation = append(blockRepresentation, ".")
			} else {
				blockRepresentation = append(blockRepresentation, strconv.Itoa(currentIndex))
			}
		}
		// toggle whether we are on free space or on a file
		if !onFreeSpace {
			currentIndex++
		}
		onFreeSpace = !onFreeSpace
	}
	return blockRepresentation
}

// get all the positions of the empty spaces in a slice
func getPositionsOfEmptySpace(blockRepresentation []string) []int {
	emptySpaceSlice := make([]int, 0)
	for index, block := range blockRepresentation {
		if block == "." {
			emptySpaceSlice = append(emptySpaceSlice, index)
		}
	}
	return emptySpaceSlice
}

// get positions of numbers in reverse
func getPositionsOfFileBlocksReversed(blockRepresentation []string) []int {
	fileBlockPositionsReversed := make([]int, 0)
	for i := len(blockRepresentation) - 1; i >= 0; i-- {
		if blockRepresentation[i] != "." {
			fileBlockPositionsReversed = append(fileBlockPositionsReversed, i)
		}
	}
	return fileBlockPositionsReversed
}

// Rearrange once, return the block representation
// Does this by
func rearrangeOnce(emptySpaceIndex int, fileBlockindex int, blockRepresentation []string) []string {
	var newBlockRepresentation []string = make([]string, len(blockRepresentation))
	copy(newBlockRepresentation, blockRepresentation)
	newBlockRepresentation[emptySpaceIndex] = strings.Clone(blockRepresentation[fileBlockindex])
	newBlockRepresentation[fileBlockindex] = strings.Clone(blockRepresentation[emptySpaceIndex])
	return newBlockRepresentation
}

// Computes the checksum
func computeChecksum(blockRepresentation []string) int {
	total := 0
	for index, num := range blockRepresentation {
		if num != "." {
			total += index * internal.GetNumFromString(num)
		}
	}
	return total
}
