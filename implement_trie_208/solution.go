package implement_trie_208

type Trie struct {
	root *trieNode
}

type trieNode struct {
	isEnd    bool
	children [26]*trieNode
}

func NewTrie() *Trie { return &Trie{root: &trieNode{}} }

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		index := char - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool { return t.traverse(word, true) }

func (t *Trie) StartsWith(prefix string) bool { return t.traverse(prefix, false) }

func (t *Trie) traverse(s string, requireEnd bool) bool {
	current := t.root
	for _, char := range s {
		index := char - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return !requireEnd || current.isEnd
}
