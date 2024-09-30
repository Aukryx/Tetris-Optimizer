package algo

import (
	"fmt"
	"sort"
)

// Solver struct holds the state for the tetromino solving algorithm
type Solver struct {
	tetrominos [][][]string // List of tetrominos to place
	grid       [][]string   // The solution grid
	size       int          // Current size of the grid
	colors     []string     // List of colors (emojis) to represent tetrominos
}

// NewSolver creates and initializes a new Solver instance
func NewSolver(tetrominos [][][]string) *Solver {
	// Sort tetrominos by size (largest first) to optimize placement
	sort.Slice(tetrominos, func(i, j int) bool {
		return len(tetrominos[i])*len(tetrominos[i][0]) > len(tetrominos[j])*len(tetrominos[j][0])
	})

	return &Solver{
		tetrominos: tetrominos,
		size:       CreateMultipleSquares(len(tetrominos)), // Start with the minimum possible square size
		colors:     []string{"🤠", "😱", "🥵", "🍆", "🍑", "🗿", "🥶", "👅", "🐵", "🦶", "🧴", "😎"},
	}
}

// Solve attempts to find a solution for placing all tetrominos
func (s *Solver) Solve() ([][]string, error) {
	maxAttempts := 3 // Limit the number of attempts to prevent infinite loops

	for attempts := 0; attempts < maxAttempts; attempts++ {
		s.grid = CreateSquare(s.size) // Create a new grid for each attempt
		if s.backtrack(0) {
			return s.grid, nil // Solution found
		}
		s.size++ // Increase grid size if no solution found
	}

	return nil, fmt.Errorf("solution not found after %d attempts", maxAttempts)
}

// backtrack implements the recursive backtracking algorithm
func (s *Solver) backtrack(tetrominoIndex int) bool {
	// If all tetrominos are placed, we've found a solution
	if tetrominoIndex == len(s.tetrominos) {
		return true
	}

	// Early termination: if remaining area is less than total area of remaining tetrominos, stop
	if s.remainingArea() < s.totalTetrominoArea(tetrominoIndex) {
		return false
	}

	// Try to place the current tetromino at each position in the grid
	for i := 0; i < s.size; i++ {
		for j := 0; j < s.size; j++ {
			if s.canPlaceTetromino(tetrominoIndex, i, j) {
				s.placeTetromino(tetrominoIndex, i, j)
				if s.backtrack(tetrominoIndex + 1) {
					return true // Solution found
				}
				s.removeTetromino(tetrominoIndex, i, j) // Backtrack
			}
		}
	}

	return false // No valid placement found for this tetromino
}

// canPlaceTetromino checks if a tetromino can be placed at the given position
func (s *Solver) canPlaceTetromino(tetrominoIndex, row, col int) bool {
	tetromino := s.tetrominos[tetrominoIndex]
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] == "#" {
				newRow, newCol := row+i, col+j
				if newRow >= s.size || newCol >= s.size || s.grid[newRow][newCol] != "  " {
					return false // Out of bounds or overlapping
				}
			}
		}
	}
	return true
}

// placeTetromino places a tetromino on the grid
func (s *Solver) placeTetromino(tetrominoIndex, row, col int) {
	tetromino := s.tetrominos[tetrominoIndex]
	color := s.colors[tetrominoIndex%len(s.colors)]
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] == "#" {
				s.grid[row+i][col+j] = color
			}
		}
	}
}

// removeTetromino removes a tetromino from the grid
func (s *Solver) removeTetromino(tetrominoIndex, row, col int) {
	tetromino := s.tetrominos[tetrominoIndex]
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] == "#" {
				s.grid[row+i][col+j] = "  "
			}
		}
	}
}

// remainingArea calculates the number of empty cells in the grid
func (s *Solver) remainingArea() int {
	emptyCount := 0
	for _, row := range s.grid {
		for _, cell := range row {
			if cell == "  " {
				emptyCount++
			}
		}
	}
	return emptyCount
}

// totalTetrominoArea calculates the total area of remaining tetrominos
func (s *Solver) totalTetrominoArea(startIndex int) int {
	total := 0
	for _, tetromino := range s.tetrominos[startIndex:] {
		for _, row := range tetromino {
			for _, cell := range row {
				if cell == "#" {
					total++
				}
			}
		}
	}
	return total
}

// PrintSolution prints the solved grid to the console
func (s *Solver) PrintSolution() {
	for _, row := range s.grid {
		fmt.Println(row)
	}
}
