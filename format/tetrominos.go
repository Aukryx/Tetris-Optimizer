package format

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetTetrominos(filePath string) [][][]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return [][][]string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ListofTetrominos [][][]string
	var currentTetromino [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentTetromino) > 0 {
				ListofTetrominos = append(ListofTetrominos, currentTetromino)
				currentTetromino = nil
			}
		} else {
			currentTetromino = append(currentTetromino, strings.Split(line, ""))
			if len(currentTetromino) == 4 {
				ListofTetrominos = append(ListofTetrominos, currentTetromino)
				currentTetromino = nil
			}
		}
	}

	if len(currentTetromino) > 0 {
		ListofTetrominos = append(ListofTetrominos, currentTetromino)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return ListofTetrominos
}
