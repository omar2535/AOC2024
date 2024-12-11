package day10

import "aoc2024/internal"

// Node type
type node struct {
	value    int
	links    []*node
	position position
}

type position struct {
	x int
	y int
}

// Create nodes of 2d array from 2d array
func create2dNodeArrayFromGrid(grid [][]string) [][]node {
	nodes := make([][]node, len(grid))
	for i := 0; i < len(grid); i++ {
		nodes[i] = make([]node, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			nodes[i][j] = node{value: internal.GetNumFromString(grid[i][j]), position: position{x: j, y: i}}
		}
	}
	return nodes
}

// Links nodes together if the number beside them is within 1 and larger
// Only link them if the immediate left, right, up, or down is within
func linkNodes(nodes [][]node) [][]node {
	newNodes := internal.Clone2dArray(nodes)
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			currentNode := nodes[y][x]
			// Check left
			if x > 0 && nodes[y][x-1].value-currentNode.value == 1 {
				currentNode.links = append(currentNode.links, &newNodes[y][x-1])
			}

			// Check right
			if x < len(nodes[y])-1 && nodes[y][x+1].value-currentNode.value == 1 {
				currentNode.links = append(currentNode.links, &newNodes[y][x+1])
			}

			// Check up
			if y > 0 && nodes[y-1][x].value-currentNode.value == 1 {
				currentNode.links = append(currentNode.links, &newNodes[y-1][x])
			}

			// Check down
			if y < len(nodes)-1 && nodes[y+1][x].value-currentNode.value == 1 {
				currentNode.links = append(currentNode.links, &newNodes[y+1][x])
			}

			newNodes[y][x] = currentNode
		}
	}
	return newNodes
}
