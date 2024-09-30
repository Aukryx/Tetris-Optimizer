package algo

// CreateSquare generates a 2D slice representing an empty square grid
func CreateSquare(size int) [][]string {
	grid := make([][]string, size)
	// Initialize each cell in the grid with empty spaces
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "  "
		}
	}
	return grid
}

// CreateMultipleSquares calculates the minimum square size needed to fit n tetrominos
func CreateMultipleSquares(n int) int {
	// Start from size 2 and increase until we find a square that can fit all tetrominos
	for i := 2; i <= 10; i++ {
		if n*4 <= i*i {
			return i
		}
	}
	return 0 // This should never happen if n is a valid number of tetrominos
}
