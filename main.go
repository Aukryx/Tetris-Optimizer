package main

import (
	"fmt"
	"os"
	"tetris-optimizer/algo"
	"tetris-optimizer/format"
)

func main() {
	// Start the timer to measure execution time
	t := algo.NewTimer()
	t.Start()

	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("You must have 1 argument exactly")
		return
	}

	// Validate the input file format
	err := format.IsValidFile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid File")
		return
	}

	// Get the tetrominos from the input file
	data := format.GetTetrominos(os.Args[1])
	if len(data) < 1 {
		fmt.Println("Invalid tetromino(s)")
		return
	}

	// Validate and format the tetrominos
	formattedTetrominos, err := format.IsValidTetrominos(data)
	if err != nil {
		fmt.Println("Invalid tetromino(s)")
		return
	}

	// Remove empty columns from the tetrominos
	formattedTetrominos, err = format.IsValidColumnTetrominos(formattedTetrominos)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	// Create a new solver with the formatted tetrominos
	solver := algo.NewSolver(formattedTetrominos)

	// Attempt to solve the puzzle
	solution, err := solver.Solve()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Print the solution
	for _, row := range solution {
		for _, cell := range row {
			if cell == "  " {
				fmt.Print(".")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}

	// Print the execution time
	elapsed := t.ElapsedSeconds()
	fmt.Printf("Program took %.2f seconds to finish\n", elapsed)
}
