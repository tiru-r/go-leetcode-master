package word_ladder_ii_126

func findLadders(beginWord, endWord string, wordList []string) [][]string {
	if beginWord == endWord {
		return [][]string{{beginWord}}
	}

	// Use struct{} for memory efficiency in sets
	wordSet := make(map[string]struct{}, len(wordList)+1)
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	wordSet[beginWord] = struct{}{}

	if _, exists := wordSet[endWord]; !exists {
		return [][]string{}
	}

	// BFS to find shortest paths with optimized queue management
	queue := []string{beginWord}
	visited := map[string]int{beginWord: 1}
	parents := make(map[string][]string)
	found := false

	for len(queue) > 0 && !found {
		nextQueue := make([]string, 0, len(queue)*4) // Pre-allocate with expected growth
		levelVisited := make(map[string]struct{})

		for _, word := range queue {
			// Generate neighbors by changing one character
			wordBytes := []byte(word)
			for i := range wordBytes {
				original := wordBytes[i]
				for c := byte('a'); c <= 'z'; c++ {
					if c == original {
						continue
					}
					wordBytes[i] = c
					neighbor := string(wordBytes)

					if _, exists := wordSet[neighbor]; !exists {
						continue
					}

					if neighbor == endWord {
						found = true
					}

					if dist, seen := visited[neighbor]; !seen {
						visited[neighbor] = visited[word] + 1
						parents[neighbor] = []string{word}
						nextQueue = append(nextQueue, neighbor)
						levelVisited[neighbor] = struct{}{}
					} else if dist == visited[word]+1 {
						parents[neighbor] = append(parents[neighbor], word)
					}
				}
				wordBytes[i] = original
			}
		}

		// Remove this level's words to prevent cycles
		for word := range levelVisited {
			delete(wordSet, word)
		}

		queue = nextQueue
	}

	if !found {
		return [][]string{}
	}

	// DFS path reconstruction with optimized memory usage
	result := make([][]string, 0, 4)
	path := make([]string, 0, visited[endWord])

	var dfs func(string)
	dfs = func(word string) {
		path = append(path, word)

		if word == beginWord {
			// Create reversed path efficiently
			reversed := make([]string, len(path))
			for i, j := 0, len(path)-1; i < len(path); i, j = i+1, j-1 {
				reversed[i] = path[j]
			}
			result = append(result, reversed)
		} else {
			for _, parent := range parents[word] {
				dfs(parent)
			}
		}

		path = path[:len(path)-1]
	}

	dfs(endWord)
	return result
}
