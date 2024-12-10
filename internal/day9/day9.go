package day9

import (
	"aoc2024/internal"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type block struct {
	blockId    string
	blockSize  int
	blockValue int
}

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
	var blockRepresentation []string = getBlockStringRepresentationFromDiskMap(diskMap)

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

func PartTwo(isTest bool) {
	fmt.Println("Day 9 part 2")

	// file contents
	var fileContents []string

	// Read the file
	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day9/day9_example.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day9/day9.txt")
	}

	// Initialize initial variables we'll need later
	var diskMap []int = internal.StringArrayToIntArray(strings.Split(fileContents[0], ""))
	var blockRepresentation []block = getBlockRepresentationFromDiskMap(diskMap)
	var maxBlockID int = getMaxBlockID(blockRepresentation)

	fmt.Print("Initial block string representation: ")
	printBlockRepresentationn(blockRepresentation)
	fmt.Println("initial block representation", blockRepresentation)
	fmt.Println("max block ID:", maxBlockID)

	// Iterate over the blocks to move, attempt to move it if possible, otherwise just skp
	for currBlockID := maxBlockID; currBlockID >= 0; currBlockID-- {
		fileBlockIndex := getIndexOfBlockID(blockRepresentation, strconv.Itoa(currBlockID))
		fileBlock := blockRepresentation[fileBlockIndex]
		emptySpaceIndex := getNextEmptySpaceForBlockSize(blockRepresentation, fileBlock.blockSize)

		// If the empty space occurs on top of or after the file block index, don't swap
		if fileBlockIndex <= emptySpaceIndex || emptySpaceIndex == -1 {
			// do nothing
		} else {
			blockRepresentation = moveBlockToEmptySpace(blockRepresentation, fileBlockIndex, emptySpaceIndex)
		}
		// printBlockRepresentationn(blockRepresentation)
	}

	fmt.Print("Final block representation: ")
	printBlockRepresentationn(blockRepresentation)
	fmt.Println("Checksum: ", computeChecksumBlocks(blockRepresentation))
}

func printBlockRepresentationn(blockRepresentation []block) {
	builtString := ""
	for _, block := range blockRepresentation {
		builtString += strings.Repeat(block.blockId, block.blockSize)
	}
	fmt.Println(builtString)
}

func getBlockRepresentationFromDiskMap(diskMap []int) []block {
	blockRepresentation := make([]block, 0)
	onFreeSpace := false
	currentBlockID := 0
	for _, numBlocks := range diskMap {
		var currentBlock block
		if onFreeSpace {
			currentBlock = block{
				blockId:    ".",
				blockSize:  numBlocks,
				blockValue: 0,
			}
		} else {
			currentBlock = block{
				blockId:    strconv.Itoa(currentBlockID),
				blockSize:  numBlocks,
				blockValue: currentBlockID,
			}
			currentBlockID++
		}
		if currentBlock.blockSize != 0 {
			blockRepresentation = append(blockRepresentation, currentBlock)
		}
		// toggle whether we are on free space or on a file
		onFreeSpace = !onFreeSpace
	}
	return blockRepresentation
}

// Get the index of a block based on its ID
// panics if the block ID isn't found
func getIndexOfBlockID(blockRepresentation []block, blockID string) int {
	for index, block := range blockRepresentation {
		if block.blockId != "." && block.blockId == blockID {
			return index
		}
	}
	panic("Couldn't find block ID")
}

// get the index of the maximum block ID
func getMaxBlockID(blockRepresentation []block) int {
	max := 0
	for _, block := range blockRepresentation {
		if block.blockId != "." {
			if block.blockValue >= max {
				max = block.blockValue
			}
		}
	}
	return max
}

// Moves block from the blockToMove index to the empty space
// If there is remaining empty space left, splits into an used block and an empty space block
// The original block will turn into empty space
// Returns the new block representation
func moveBlockToEmptySpace(blockRepresentation []block, fileBlockIndex int, emptySpaceIndex int) []block {
	emptySpaceBlock := blockRepresentation[emptySpaceIndex]
	fileBlock := blockRepresentation[fileBlockIndex]

	// If we can't move it, then just return the original block representation
	if emptySpaceBlock.blockSize < fileBlock.blockSize {
		return blockRepresentation
	}

	// Move it
	// Make a new block representation first (so we aren't modifying the original - easier to debug)
	var newBlockRepresentation []block = make([]block, len(blockRepresentation))
	copy(newBlockRepresentation, blockRepresentation)

	if emptySpaceBlock.blockSize == fileBlock.blockSize {
		newBlockRepresentation[emptySpaceIndex] = fileBlock
		newBlockRepresentation[fileBlockIndex] = emptySpaceBlock
	} else {
		remainingEmptySpace := emptySpaceBlock.blockSize - fileBlock.blockSize
		newEmptyBlock := block{blockId: ".", blockSize: remainingEmptySpace, blockValue: 0}

		// Replace the original file block with empty space
		newBlockRepresentation[fileBlockIndex].blockId = "."
		newBlockRepresentation[fileBlockIndex].blockValue = 0

		// Remove the original, then add the two blocks (fileBlock + remainingEmptySpace)
		newBlockRepresentation = internal.RemoveFromList(newBlockRepresentation, emptySpaceIndex)
		newBlockRepresentation = slices.Insert(newBlockRepresentation, emptySpaceIndex, newEmptyBlock)
		newBlockRepresentation = slices.Insert(newBlockRepresentation, emptySpaceIndex, fileBlock)
	}

	return newBlockRepresentation
}

// Get index of next empty space forthe size
// Returns -1 if no block found
func getNextEmptySpaceForBlockSize(blockRepresentation []block, size int) int {
	for index, currentBlock := range blockRepresentation {
		if currentBlock.blockSize >= size && currentBlock.blockId == "." {
			return index
		}
	}
	return -1
}

// Returns list of characters of the block representation
func getBlockStringRepresentationFromDiskMap(dispMap []int) []string {
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

// compute checksum
func computeChecksumBlocks(blockRepresentation []block) int {

	blockRepresentationStringSlice := make([]string, 0)
	for _, block := range blockRepresentation {
		for i := 0; i < block.blockSize; i++ {
			blockRepresentationStringSlice = append(blockRepresentationStringSlice, block.blockId)
		}
	}
	return computeChecksum(blockRepresentationStringSlice)
}
