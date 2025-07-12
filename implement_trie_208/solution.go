package implement_trie_208

// Trie is exported for LeetCode tests.
type Trie struct{ root *node }

type node struct {
	end      bool
	children [26]*node
}

// Constructor returns an empty Trie.
func Constructor() Trie { return Trie{root: &node{}} }

// Insert adds word to the trie.
func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &node{}
		}
		curr = curr.children[idx]
	}
	curr.end = true
}

// Search returns true if word is stored in the trie.
func (t *Trie) Search(word string) bool { return t.walk(word, true) }

// StartsWith returns true if any stored word starts with prefix.
func (t *Trie) StartsWith(prefix string) bool { return t.walk(prefix, false) }

// walk performs common traversal logic.
func (t *Trie) walk(s string, needEnd bool) bool {
	curr := t.root
	for _, ch := range s {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return !needEnd || curr.end
}
