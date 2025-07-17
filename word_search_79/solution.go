package word_search_79

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	// Check if board has enough characters to form the word
	charCount := make(map[byte]int)
	for _, row := range board {
		for _, c := range row {
			charCount[c]++
		}
	}
	for _, c := range word {
		if charCount[byte(c)] == 0 {
			return false
		}
		charCount[byte(c)]--
	}

	rows, cols := len(board), len(board[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] && explore(board, word, 0, r, c, rows, cols) {
				return true
			}
		}
	}
	return false
}

func explore(board [][]byte, word string, step, r, c, rows, cols int) bool {
	if step == len(word) {
		return true
	}
	if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != word[step] {
		return false
	}

	originalChar := board[r][c]
	board[r][c] = '#' // Use '#' to mark visited

	found := explore(board, word, step+1, r-1, c, rows, cols) || // Up
		explore(board, word, step+1, r+1, c, rows, cols) || // Down
		explore(board, word, step+1, r, c-1, rows, cols) || // Left
		explore(board, word, step+1, r, c+1, rows, cols) // Right

	board[r][c] = originalChar // Backtrack
	return found
}
