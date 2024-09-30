package format

// ColorTetrominos applies emoji colors to the tetrominos
func ColorTetrominos(ListofTetrominos [][][]string) [][][]string {
	// Define a list of emoji colors
	colors := []string{
		"🤠", "😱", "🥵", "🍆", "🍑", "🗿", "🥶", "👅", "🐵", "🦶", "🧴", "😎",
	}
	index := 0

	// Iterate through each tetromino and apply colors
	for i, tetromino := range ListofTetrominos {
		for j, row := range tetromino {
			for k, ch := range row {
				if ch == "#" {
					// Replace "#" with the current color emoji
					ListofTetrominos[i][j][k] = colors[index]
				}
				if ch == "." {
					// Replace "." with empty space
					ListofTetrominos[i][j][k] = "  "
				}
			}
		}
		// Move to the next color for the next tetromino
		index++
	}
	return ListofTetrominos
}
