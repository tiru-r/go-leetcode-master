package word_search_212

const (
	alphabetSize  = 26
	visitedMarker = '#'
)

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type TrieNode struct {
	children [alphabetSize]*TrieNode
	word     string // non-empty only for terminal nodes
	isEnd    bool
}

func (t *TrieNode) insert(word string) {
	n := t
	for _, c := range word {
		idx := c - 'a'
		if c < 'a' || c > 'z' {
			continue // skip invalid characters
		}
		if n.children[idx] == nil {
			n.children[idx] = &TrieNode{}
		}
		n = n.children[idx]
	}
	n.word = word
	n.isEnd = true
}

// hasChildren reports whether node has any living children (used for pruning)
func (t *TrieNode) hasChildren() bool {
	for _, ch := range t.children {
		if ch != nil {
			return true
		}
	}
	return false
}

// findWords returns all dictionary words found on the board.
func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}

	// Build trie
	root := &TrieNode{}
	for _, w := range words {
		root.insert(w)
	}

	m, n := len(board), len(board[0])
	// Optimize capacity: most leetcode problems have < 50% hit rate
	res := make([]string, 0, min(len(words), 10))
	wordsRemaining := len(words)

	var backtrack func(r, c int, node *TrieNode)
	backtrack = func(r, c int, node *TrieNode) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return
		}
		ch := board[r][c]
		if ch == visitedMarker {
			return
		}
		idx := ch - 'a'
		if ch < 'a' || ch > 'z' {
			return // invalid character
		}
		next := node.children[idx]
		if next == nil {
			return
		}

		if next.isEnd {
			res = append(res, next.word)
			next.isEnd = false
			next.word = "" // help GC
			wordsRemaining--
			if wordsRemaining == 0 {
				return // early termination
			}
		}

		board[r][c] = visitedMarker
		for _, d := range directions {
			backtrack(r+d[0], c+d[1], next)
		}
		board[r][c] = ch

		// prune dead branch
		if !next.isEnd && !next.hasChildren() {
			node.children[idx] = nil
		}
	}

	for i := range m {
		if wordsRemaining == 0 {
			break
		}
		for j := range n {
			if wordsRemaining == 0 {
				break
			}
			ch := board[i][j]
			if ch >= 'a' && ch <= 'z' {
				idx := ch - 'a'
				if root.children[idx] != nil {
					backtrack(i, j, root)
				}
			}
		}
	}

	return res
}
