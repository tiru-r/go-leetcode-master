package word_ladder_ii_126

import "slices"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return [][]string{}
	}
	
	level := map[string]bool{beginWord: true}
	parents := make(map[string][]string)
	found := false
	
	for len(level) > 0 && !found {
		nextLevel := make(map[string]bool)
		for word := range level {
			for i := 0; i < len(word); i++ {
				for c := byte('a'); c <= 'z'; c++ {
					if c == word[i] {
						continue
					}
					
					newWord := word[:i] + string(c) + word[i+1:]
					if !wordSet[newWord] {
						continue
					}
					
					if newWord == endWord {
						found = true
					}
					
					if !found {
						if _, exists := level[newWord]; !exists {
							nextLevel[newWord] = true
						}
					}
					
					parents[newWord] = append(parents[newWord], word)
				}
			}
		}
		
		for word := range nextLevel {
			delete(wordSet, word)
		}
		level = nextLevel
	}
	
	if !found {
		return [][]string{}
	}
	
	var result [][]string
	var dfs func(word string, path []string)
	dfs = func(word string, path []string) {
		if word == beginWord {
			reversed := make([]string, len(path))
			copy(reversed, path)
			slices.Reverse(reversed)
			result = append(result, reversed)
			return
		}
		
		for _, parent := range parents[word] {
			dfs(parent, append(path, parent))
		}
	}
	
	dfs(endWord, []string{endWord})
	return result
}
