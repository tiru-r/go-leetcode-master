package implement_trie_208

// Trie is a generic prefix tree for Unicode runes.
// Supports O(m) insert, search, and prefix operations where m is word length.
type Trie struct {
	next  map[rune]*Trie
	isEnd bool
}

// NewTrie returns an empty prefix tree.
func NewTrie() *Trie { 
	return &Trie{next: make(map[rune]*Trie, 4)} // hint for typical branching factor
}

// Insert adds word to the trie (O(m) time, O(m) space in worst case).
func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		if next := curr.next[ch]; next == nil {
			curr.next[ch] = &Trie{next: make(map[rune]*Trie, 2)} // smaller hint for leaf nodes
			curr = curr.next[ch]
		} else {
			curr = next
		}
	}
	curr.isEnd = true
}

// Search reports whether word exists in the trie (O(m) time, O(1) space).
func (t *Trie) Search(word string) bool {
	curr := t
	for _, ch := range word {
		if next := curr.next[ch]; next != nil {
			curr = next
		} else {
			return false
		}
	}
	return curr.isEnd
}

// StartsWith reports whether any stored word has the given prefix (O(m) time, O(1) space).
func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for _, ch := range prefix {
		if next := curr.next[ch]; next != nil {
			curr = next
		} else {
			return false
		}
	}
	return true
}
