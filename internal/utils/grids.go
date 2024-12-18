package utils

// Find the element in the grid
func FindElementInGrid(grid [][]string, element string) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == element {
				return j, i
			}
		}
	}
	return -1, -1
}
