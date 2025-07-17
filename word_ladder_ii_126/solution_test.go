package word_ladder_ii_126

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindLadders(t *testing.T) {
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  [][]string
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}},
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  [][]string{},
		},
		{
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			expected:  [][]string{{"a", "c"}},
		},
	}

	for i, test := range tests {
		result := findLadders(test.beginWord, test.endWord, test.wordList)
		
		sort.Slice(result, func(i, j int) bool {
			if len(result[i]) != len(result[j]) {
				return len(result[i]) < len(result[j])
			}
			for k := 0; k < len(result[i]); k++ {
				if result[i][k] != result[j][k] {
					return result[i][k] < result[j][k]
				}
			}
			return false
		})
		sort.Slice(test.expected, func(i, j int) bool {
			if len(test.expected[i]) != len(test.expected[j]) {
				return len(test.expected[i]) < len(test.expected[j])
			}
			for k := 0; k < len(test.expected[i]); k++ {
				if test.expected[i][k] != test.expected[j][k] {
					return test.expected[i][k] < test.expected[j][k]
				}
			}
			return false
		})

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d failed: expected %v, got %v", i+1, test.expected, result)
		}
	}
}

func BenchmarkFindLadders(b *testing.B) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findLadders(beginWord, endWord, wordList)
	}
}

func BenchmarkFindLaddersLarge(b *testing.B) {
	beginWord := "qa"
	endWord := "sq"
	wordList := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findLadders(beginWord, endWord, wordList)
	}
}
