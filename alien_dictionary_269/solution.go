package alien_dictionary_269

import (
	"slices"
	"strings"
)

// AlienOrder finds lexicographical order of alien alphabet from sorted words
// Time: O(N*M + V + E), Space: O(V + E) where N=words, M=avg length, V=chars, E=edges
func AlienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Collect all characters and initialize graph
	adj := make(map[byte][]byte)
	inDegree := make(map[byte]int)

	for _, word := range words {
		for i := range len(word) {
			c := word[i]
			if _, exists := inDegree[c]; !exists {
				inDegree[c] = 0
			}
		}
	}

	// Build dependency graph
	for i := range len(words) - 1 {
		curr, next := words[i], words[i+1]
		
		// Invalid: longer word prefixes shorter
		if len(curr) > len(next) && strings.HasPrefix(curr, next) {
			return ""
		}

		// Find first difference
		for j := range min(len(curr), len(next)) {
			if curr[j] != next[j] {
				u, v := curr[j], next[j]
				// Avoid duplicate edges
				if !slices.Contains(adj[u], v) {
					adj[u] = append(adj[u], v)
					inDegree[v]++
				}
				break
			}
		}
	}

	// Kahn's algorithm for topological sort
	queue := make([]byte, 0, len(inDegree))
	for c, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, c)
		}
	}
	slices.Sort(queue)

	var result strings.Builder
	result.Grow(len(inDegree))

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result.WriteByte(u)

		for _, v := range adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// Check for cycles
	if result.Len() != len(inDegree) {
		return ""
	}
	return result.String()
}
