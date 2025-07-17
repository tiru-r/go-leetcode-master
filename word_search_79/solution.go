package word_search_79

func exist(board [][]byte, word string) bool {
	for r := range len(board) {
		for c := range len(board[r]) {
			if word[0] == board[r][c] && explore(board, word, 0, r, c) {
				return true
			}
		}
	}

	return false
}

func explore(board [][]byte, word string, step, r, c int) bool {
	if step == len(word) {
		return true
	}

	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}

	if word[step] != board[r][c] {
		return false
	}

	if board[r][c] == 0 {
		return false
	}

	original := board[r][c]
	board[r][c] = 0

	exists := explore(board, word, step+1, r-1, c) ||
		explore(board, word, step+1, r+1, c) ||
		explore(board, word, step+1, r, c-1) ||
		explore(board, word, step+1, r, c+1)

	board[r][c] = original

	return exists
}
