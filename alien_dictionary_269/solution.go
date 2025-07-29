package alien_dictionary_269

import (
	"slices"
	"strings"
)

// AlienOrder finds lexicographical order of alien alphabet from sorted words
// Time: O(N*M + V + E), Space: O(V + E) where N=words, M=avg length, V=chars, E=edges
func AlienOrder(words []string) string {
	// Return empty string if no words provided
	if len(words) == 0 {
		return ""
	}

	// Collect all characters and initialize graph
	adj := make(map[byte][]byte)
	inDegree := make(map[byte]int)

	// Initialize inDegree for all characters found in words
	for _, word := range words {
		for i := range len(word) {
			c := word[i]
			if _, exists := inDegree[c]; !exists {
				inDegree[c] = 0
			}
		}
	}

	// Build dependency graph by comparing adjacent words
	for i := range len(words) - 1 {
		curr, next := words[i], words[i+1]

		// Invalid case: longer word is a prefix of shorter word
		if len(curr) > len(next) && strings.HasPrefix(curr, next) {
			return ""
		}

		// Find first differing character between current and next word
		minLen := min(len(curr), len(next))
		for j := range minLen {
			if curr[j] != next[j] {
				u, v := curr[j], next[j]
				// Check if this would create a cycle
				if dfsHasPath(adj, v, u) {
					return "" // Cycle detected
				}
				// Avoid adding duplicate edges
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

	// Check if all characters were processed (no cycles)
	if result.Len() != len(inDegree) {
		return ""
	}
	return result.String()
}

// Helper function to check if there's a path from src to dest
func dfsHasPath(graph map[byte][]byte, src, dest byte) bool {
	if src == dest {
		return true
	}
	visited := make(map[byte]bool)
	return dfsHasPathHelper(graph, src, dest, visited)
}

func dfsHasPathHelper(graph map[byte][]byte, current, dest byte, visited map[byte]bool) bool {
	if current == dest {
		return true
	}
	if visited[current] {
		return false
	}
	visited[current] = true

	for _, neighbor := range graph[current] {
		if dfsHasPathHelper(graph, neighbor, dest, visited) {
			return true
		}
	}
	return false
}
