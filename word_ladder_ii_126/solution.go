package word_ladder_ii_126

// findLadders returns every shortest transformation sequence
// from beginWord to endWord, changing one letter at a time,
// with each intermediate word present in wordList.
func findLadders(beginWord, endWord string, wordList []string) [][]string {
	if beginWord == endWord {
		return [][]string{{beginWord}}
	}

	// 1. Build set; ensure endWord exists
	wordSet := make(map[string]struct{}, len(wordList))
	for _, w := range wordList {
		wordSet[w] = struct{}{}
	}
	if _, ok := wordSet[endWord]; !ok {
		return nil
	}

	// 2. BFS: shortest distance + parent graph
	visited := map[string]int{beginWord: 1}
	parents := make(map[string][]string, len(wordList)/2)

	queue := []string{beginWord}

	for len(queue) > 0 {
		nextQ := make([]string, 0, len(queue)*2)
		levelUsed := make(map[string]struct{})

		for _, word := range queue {
			// generate 1-char neighbors
			runes := []rune(word)
			for i := range runes {
				orig := runes[i]
				for c := 'a'; c <= 'z'; c++ {
					if c == orig {
						continue
					}
					runes[i] = c
					nbr := string(runes)

					if _, ok := wordSet[nbr]; !ok {
						continue
					}

					if d, ok := visited[nbr]; !ok {
						visited[nbr] = visited[word] + 1
						parents[nbr] = []string{word}
						nextQ = append(nextQ, nbr)
						levelUsed[nbr] = struct{}{}
					} else if d == visited[word]+1 {
						parents[nbr] = append(parents[nbr], word)
					}
				}
				runes[i] = orig
			}
		}

		// prune words discovered this level
		for w := range levelUsed {
			delete(wordSet, w)
		}

		// early exit if endWord reached
		if visited[endWord] != 0 {
			break
		}
		queue = nextQ
	}

	// 3. DFS reconstruction from endWord
	if visited[endWord] == 0 {
		return nil
	}

	var res [][]string
	path := make([]string, 0, visited[endWord])

	reversePath := func(p []string) []string {
		rev := make([]string, len(p))
		for i, s := range p {
			rev[len(p)-1-i] = s
		}
		return rev
	}

	var dfs func(string)
	dfs = func(word string) {
		path = append(path, word)
		if word == beginWord {
			res = append(res, reversePath(path))
		} else {
			for _, p := range parents[word] {
				dfs(p)
			}
		}
		path = path[:len(path)-1]
	}

	dfs(endWord)
	return res
}
