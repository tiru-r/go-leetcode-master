package implement_trie_208

// Trie implements an optimized prefix tree using Go 1.24 modern syntax
// Optimized: O(m) operations with lazy allocation and efficient traversal
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// NewTrie creates a new Trie instance (idiomatic Go constructor)
func NewTrie() *Trie {
	return &Trie{}
}

// Insert adds a word to the trie using modern range syntax
// Time: O(m), Space: O(m) worst case with lazy node allocation
func (t *Trie) Insert(word string) {
	curr := t
	for _, char := range []byte(word) {
		idx := char - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

// Search returns true if the word exists in the trie
// Time: O(m) with early termination on missing paths
func (t *Trie) Search(word string) bool {
	if node := t.findPrefix(word); node != nil {
		return node.isEnd
	}
	return false
}

// StartsWith returns true if any word has the given prefix
// Time: O(m) with optimized prefix traversal
func (t *Trie) StartsWith(prefix string) bool {
	return t.findPrefix(prefix) != nil
}

// findPrefix traverses to the end of prefix, returns nil if path doesn't exist
func (t *Trie) findPrefix(prefix string) *Trie {
	curr := t
	for _, char := range []byte(prefix) {
		idx := char - 'a'
		if curr.children[idx] == nil {
			return nil
		}
		curr = curr.children[idx]
	}
	return curr
}
