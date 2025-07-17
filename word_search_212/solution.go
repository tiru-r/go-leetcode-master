package word_search_212

type TrieNode struct {
	children map[rune]*TrieNode
	word     string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *TrieNode) insert(word string) {
	cur := t
	for _, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = NewTrieNode()
		}
		cur = cur.children[ch]
	}
	cur.word = word
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}

	root := NewTrieNode()
	for _, w := range words {
		root.insert(w)
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	var res []string

	var dfs func(int, int, *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
			return
		}

		ch := rune(board[i][j])
		next := node.children[ch]
		if next == nil {
			return
		}

		visited[i][j] = true

		if next.word != "" {
			res = append(res, next.word)
			next.word = ""
		}

		for _, d := range dirs {
			dfs(i+d[0], j+d[1], next)
		}

		visited[i][j] = false

		if len(next.children) == 0 {
			delete(node.children, ch)
		}
	}

	for i := range m {
		for j := range n {
			dfs(i, j, root)
		}
	}
	return res
}
