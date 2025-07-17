package word_search_212

type TrieNode struct {
	children [26]*TrieNode
	word     string
	isEnd    bool
}

func (t *TrieNode) insert(word string) {
	node := t
	for _, ch := range []byte(word) {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.word = word
	node.isEnd = true
}

func (t *TrieNode) hasChildren() bool {
	for _, child := range t.children {
		if child != nil {
			return true
		}
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return []string{}
	}

	root := &TrieNode{}
	for _, word := range words {
		root.insert(word)
	}

	result := make([]string, 0, len(words))

	var dfs func(int, int, *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return
		}

		ch := board[i][j]
		if ch == '#' || ch < 'a' || ch > 'z' {
			return
		}

		idx := ch - 'a'
		next := node.children[idx]
		if next == nil {
			return
		}

		if next.isEnd {
			result = append(result, next.word)
			next.isEnd = false
			next.word = ""
		}

		board[i][j] = '#'

		// Explore all 4 directions
		dfs(i-1, j, next)
		dfs(i+1, j, next)
		dfs(i, j-1, next)
		dfs(i, j+1, next)

		board[i][j] = ch

		// Pruning: remove dead branches
		if !next.isEnd && next.word == "" && !next.hasChildren() {
			node.children[idx] = nil
		}
	}

	for i := range board {
		for j := range board[i] {
			if ch := board[i][j]; ch >= 'a' && ch <= 'z' && root.children[ch-'a'] != nil {
				dfs(i, j, root)
			}
		}
	}

	return result
}
