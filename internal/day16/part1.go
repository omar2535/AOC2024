package day16

import (
	"aoc2024/internal"
	"aoc2024/internal/utils"
	"fmt"
	"os"
)

const FILEPATH = "res/day16/maze_progress.txt"

type position struct {
	x int
	y int
}

type node struct {
	value          string
	position       position
	links          []*node
	isWall         bool
	isStartingNode bool
	isEndingNode   bool
}

func PartOne(isTest bool) {

	// Some initial variables
	var fileContents []string

	if isTest {
		fileContents = internal.ReadFileIntoArray("res/day16/day16_example2.txt")
	} else {
		fileContents = internal.ReadFileIntoArray("res/day16/day16.txt")
	}

	// Create a 2d array of nodes & link the nodes
	var gridString [][]string = internal.ConvertStringListToGrid(fileContents, "")
	var nodes [][]node = create2dnodearrayfromgrid(gridString)
	nodes = linkNodes(nodes)
	var startingNode node = findStartingNode(nodes)
	var endingNode node = findEndingNode(nodes)

	// Do DFS
	visited_nodes, shortestPath := findShortestPath(startingNode, startingNode, endingNode, "", []node{}, gridString)
	fmt.Print("Visited nodes: ")
	printVisitedNodes(visited_nodes)
	fmt.Println("Shortest path:", shortestPath)
	fmt.Println("Starting node: ", startingNode.position)
	fmt.Println("Ending node: ", endingNode.position)
}

// Find the shortest path
func findShortestPath(
	currentNode node,
	startingNode node,
	endingNode node,
	direction string,
	visitedNodes []node,
	gridString [][]string) ([]node, int) {

	// fmt.Println("Current node:", currentNode.position)
	// printGrid(gridString, visitedNodes, currentNode, startingNode, endingNode, direction)
	printGridToFile(gridString, visitedNodes, currentNode, startingNode, endingNode, direction, FILEPATH)

	// Check if the node is the ending node
	if currentNode.position == endingNode.position {
		return append(visitedNodes, currentNode), 0
	}

	// Visit the other nodes
	// Check if the node has been visited
	var possibleCosts []int = make([]int, 0)
	var possiblePaths [][]node = make([][]node, 0)
	visitedNodes = append(visitedNodes, currentNode)
	for i := 0; i < len(currentNode.links); i++ {
		if !isNodeInNodes(*currentNode.links[i], visitedNodes) {
			// Check if direction is up, down, left, or right
			var nextNodeDirection string = getNextNodeDirection(currentNode, *currentNode.links[i])
			var costFactor int = 1

			if nextNodeDirection != direction {
				costFactor = 1001
			}

			// Add the cost of the path
			path, cost := findShortestPath(*currentNode.links[i], startingNode, endingNode, nextNodeDirection, visitedNodes, gridString)
			possibleCosts = append(possibleCosts, cost+costFactor)
			possiblePaths = append(possiblePaths, path)
		}
	}

	// No path found, just return a big number
	if len(possibleCosts) == 0 {
		return visitedNodes, 100000000000000000
	}
	minPath := possibleCosts[0]
	minPathIndex := 0
	for index, path := range possibleCosts {
		if path < minPath {
			minPath = path
			minPathIndex = index
		}
	}
	return possiblePaths[minPathIndex], minPath
}

// Links nodes together only if the besides are .'s as well
func linkNodes(nodes [][]node) [][]node {
	newNodes := internal.Clone2dArray(nodes)
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			currentNode := nodes[y][x]
			// Check left
			if x > 0 && !nodes[y][x-1].isWall {
				currentNode.links = append(currentNode.links, &newNodes[y][x-1])
			}

			// Check right
			if x < len(nodes[y])-1 && !nodes[y][x+1].isWall {
				currentNode.links = append(currentNode.links, &newNodes[y][x+1])
			}

			// Check up
			if y > 0 && !nodes[y-1][x].isWall {
				currentNode.links = append(currentNode.links, &newNodes[y-1][x])
			}

			// Check down
			if y < len(nodes)-1 && !nodes[y+1][x].isWall {
				currentNode.links = append(currentNode.links, &newNodes[y+1][x])
			}

			newNodes[y][x] = currentNode
		}
	}
	return newNodes
}

// create nodes of 2d array from 2d array
func create2dnodearrayfromgrid(grid [][]string) [][]node {
	nodes := make([][]node, len(grid))
	for i := 0; i < len(grid); i++ {
		nodes[i] = make([]node, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			nodes[i][j] = node{
				value:          grid[i][j],
				isWall:         grid[i][j] == "#",
				isStartingNode: grid[i][j] == "S",
				isEndingNode:   grid[i][j] == "E",
				position:       position{x: j, y: i},
			}
		}
	}
	return nodes
}

func findStartingNode(nodes [][]node) node {
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			if nodes[y][x].isStartingNode {
				return nodes[y][x]
			}
		}
	}
	panic("No starting node found")
}

func findEndingNode(nodes [][]node) node {
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			if nodes[y][x].isEndingNode {
				return nodes[y][x]
			}
		}
	}
	panic("No ending node found")
}

func isNodeInNodes(node node, nodes []node) bool {
	for i := 0; i < len(nodes); i++ {
		if node.position == nodes[i].position {
			return true
		}
	}
	return false
}

func getNextNodeDirection(currentNode node, nextNode node) string {
	if currentNode.position.x == nextNode.position.x {
		if currentNode.position.y > nextNode.position.y {
			return "up"
		} else {
			return "down"
		}
	} else {
		if currentNode.position.x > nextNode.position.x {
			return "left"
		} else {
			return "right"
		}
	}
}

func printVisitedNodes(visitedNodes []node) {
	for i := 0; i < len(visitedNodes); i++ {
		fmt.Println(visitedNodes[i].position)
	}
}

func isVisitedNodeForPosition(position position, visitedNodes []node) bool {
	for i := 0; i < len(visitedNodes); i++ {
		if visitedNodes[i].position == position {
			return true
		}
	}
	return false
}

func printGrid(gridString [][]string, visitedNodes []node, currentNode node, startingNode node, endingNode node, direction string) {
	utils.CallClear()
	for y := 0; y < len(gridString); y++ {
		for x := 0; x < len(gridString[y]); x++ {
			if currentNode.position.x == x && currentNode.position.y == y {
				if direction == "up" {
					fmt.Print("^")
				} else if direction == "down" {
					fmt.Print("v")
				} else if direction == "left" {
					fmt.Print("<")
				} else if direction == "right" {
					fmt.Print(">")
				} else {
					fmt.Print("X")
				}
			} else if startingNode.position == (position{x: x, y: y}) {
				fmt.Print("S")
			} else if endingNode.position == (position{x: x, y: y}) {
				fmt.Print("E")
			} else if isVisitedNodeForPosition(position{x: x, y: y}, visitedNodes) {
				fmt.Print("o")
			} else {
				fmt.Print(gridString[y][x])
			}
		}
		fmt.Println()
	}
}

func printGridToFile(gridString [][]string, visitedNodes []node, currentNode node, startingNode node, endingNode node, direction string, filePath string) {
	// Open the file for writing and truncate it each time
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for y := 0; y < len(gridString); y++ {
		for x := 0; x < len(gridString[y]); x++ {
			if currentNode.position.x == x && currentNode.position.y == y {
				if direction == "up" {
					fmt.Fprint(file, "^")
				} else if direction == "down" {
					fmt.Fprint(file, "v")
				} else if direction == "left" {
					fmt.Fprint(file, "<")
				} else if direction == "right" {
					fmt.Fprint(file, ">")
				} else {
					fmt.Fprint(file, "X")
				}
			} else if startingNode.position == (position{x: x, y: y}) {
				fmt.Fprint(file, "S")
			} else if endingNode.position == (position{x: x, y: y}) {
				fmt.Fprint(file, "E")
			} else if isVisitedNodeForPosition(position{x: x, y: y}, visitedNodes) {
				fmt.Fprint(file, "o")
			} else {
				fmt.Fprint(file, gridString[y][x])
			}
		}
		fmt.Fprintln(file) // Newline after each row
	}
}
