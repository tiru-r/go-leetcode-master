package implement_trie_208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	// 1. Empty trie
	assert.False(t, trie.Search(""))
	assert.False(t, trie.StartsWith(""))

	// 2. Single character
	trie.Insert("a")
	assert.True(t, trie.Search("a"))
	assert.False(t, trie.Search("b"))
	assert.True(t, trie.StartsWith("a"))
	assert.False(t, trie.StartsWith("b"))

	// 3. Insert & search multiple words
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.False(t, trie.Search("app"))
	assert.True(t, trie.StartsWith("app"))

	trie.Insert("app")
	assert.True(t, trie.Search("app"))

	// 4. Overlapping prefixes
	trie.Insert("bear")
	trie.Insert("bell")
	assert.True(t, trie.StartsWith("be"))
	assert.True(t, trie.Search("bear"))
	assert.True(t, trie.Search("bell"))

	// 5. Case sensitivity (input is lowercase per problem)
	trie.Insert("abc")
	assert.False(t, trie.Search("ABC"))

	// 6. Long word
	trie.Insert("abcdefghijklmnopqrstuvwxyz")
	assert.True(t, trie.Search("abcdefghijklmnopqrstuvwxyz"))
	assert.True(t, trie.StartsWith("abcdefghijklmnopqrstuvwxy"))

	// 7. Repeated insert (idempotent)
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))

	// 8. Not present but prefix exists
	assert.False(t, trie.Search("application"))
	assert.True(t, trie.StartsWith("application")) // prefix exists

	// 9. Empty string after inserts
	assert.True(t, trie.StartsWith("")) // root always matches empty prefix
}

func BenchmarkTrie(b *testing.B) {
	trie := Constructor()
	for i := 0; i < b.N; i++ {
		trie.Insert("apple")
		trie.Search("apple")
		trie.StartsWith("app")
	}
}
