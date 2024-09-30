// Package format provides functions for validating and formatting tetromino files and shapes.
package format

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// IsValidFile checks if the given file contains valid tetromino representations.
// It reads the file line by line and validates the characters and format.
func IsValidFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("couldn't open file")
	}
	defer file.Close()

	compteur := 0 // Counter to keep track of lines within each tetromino
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		compteur++
		for _, ch := range line {
			if compteur == 5 {
				// The 5th line should be empty (newline)
				if ch == '.' || ch == '#' {
					return errors.New("bad formatting")
				}
				compteur = 0 // Reset counter for next tetromino
			}
			// Check if character is valid ('#' or '.')
			if ch != '.' && ch != '#' {
				return errors.New("a character is not valid")
			}
		}
	}
	return nil
}

// IsValidTetrominos validates a list of tetrominos and returns formatted versions.
// It checks each tetromino for validity and removes any empty rows.
func IsValidTetrominos(ListofTetrominos [][][]string) ([][][]string, error) {
	FormatedTetrominos := [][][]string{}

	for _, tetromino := range ListofTetrominos {
		valid, filteredTetromino := isValidTetromino(tetromino)
		if !valid {
			return nil, fmt.Errorf("invalid tetromino found")
		}

		FormatedTetrominos = append(FormatedTetrominos, filteredTetromino)
	}

	return FormatedTetrominos, nil
}

// isValidTetromino checks if a single tetromino is valid and removes empty rows.
// It returns a boolean indicating validity and the filtered tetromino.
func isValidTetromino(tetromino [][]string) (bool, [][]string) {
	if len(tetromino) == 0 || len(tetromino) > 4 {
		return false, nil
	}

	blocks := 0 // Count of '#' blocks
	links := 0  // Count of connections between blocks
	validTetromino := [][]string{}

	for i, row := range tetromino {
		if len(row) != 4 {
			return false, nil
		}
		hasHash := false
		for j, ch := range row {
			if ch == "#" {
				hasHash = true
				blocks++
				// Check for vertical connection
				if i > 0 && j < len(tetromino[i-1]) && tetromino[i-1][j] == "#" {
					links++
				}
				// Check for horizontal connection
				if j > 0 && tetromino[i][j-1] == "#" {
					links++
				}
			} else if ch != "." {
				return false, nil
			}
		}
		if hasHash {
			validTetromino = append(validTetromino, row)
		}
	}
	// A valid tetromino must have 4 blocks and at least 3 connections
	if blocks != 4 || links < 3 {
		return false, nil
	}

	return true, validTetromino
}

// RemoveEmptyColumns removes empty columns from a tetromino.
// It returns a boolean (always true in this implementation) and the formatted tetromino.
func RemoveEmptyColumns(tetromino [][]string) (bool, [][]string) {
	formattedTetro := [][]string{}
	colCounts := make([]int, len(tetromino[0])) // count of non-empty cells in each column

	// Count non-empty cells in each column
	for _, row := range tetromino {
		for j, ch := range row {
			if ch != "." {
				colCounts[j]++
			}
		}
	}

	// Remove empty columns
	for _, row := range tetromino {
		newRow := []string{}
		for j, ch := range row {
			if colCounts[j] > 0 {
				newRow = append(newRow, ch)
			}
		}
		formattedTetro = append(formattedTetro, newRow)
	}

	return true, formattedTetro
}

// IsValidColumnTetrominos validates a list of tetrominos and removes empty columns.
// It returns the formatted list of tetrominos and any error encountered.
func IsValidColumnTetrominos(ListofTetrominos [][][]string) ([][][]string, error) {
	FormatedTetrominos := [][][]string{}

	for _, tetromino := range ListofTetrominos {
		valid, filteredTetromino := RemoveEmptyColumns(tetromino)
		if !valid {
			return nil, fmt.Errorf("invalid tetromino found")
		}

		FormatedTetrominos = append(FormatedTetrominos, filteredTetromino)
	}

	return FormatedTetrominos, nil
}
