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
	// adj stores adjacency list: character -> characters that come after it
	adj := make(map[byte][]byte)
	// inDegree tracks how many characters come before each character
	inDegree := make(map[byte]int)

	// Initialize inDegree for all characters found in words
	for _, word := range words {
		// Iterate through each character in current word
		for i := range len(word) {
			c := word[i]  // Get current character
			// Add character to graph if not already present
			if _, exists := inDegree[c]; !exists {
				inDegree[c] = 0  // Initialize with 0 incoming edges
			}
		}
	}

	// Build dependency graph by comparing adjacent words
	for i := range len(words) - 1 {
		curr, next := words[i], words[i+1]  // Get current and next word
		
		// Invalid case: longer word is a prefix of shorter word
		// This violates lexicographical ordering rules
		if len(curr) > len(next) && strings.HasPrefix(curr, next) {
			return ""  // No valid ordering exists
		}

		// Find first differing character between current and next word
		for j := range min(len(curr), len(next)) {
			// Compare characters at same position
			if curr[j] != next[j] {
				u, v := curr[j], next[j]  // u comes before v in alien alphabet
				// Avoid adding duplicate edges in adjacency list
				if !slices.Contains(adj[u], v) {
					adj[u] = append(adj[u], v)  // Add edge u -> v
					inDegree[v]++               // Increment incoming degree of v
				}
				break  // Only need first difference between words
			}
		}
	}

	// Kahn's algorithm for topological sort
	// Initialize queue with characters having no incoming edges
	queue := make([]byte, 0, len(inDegree))
	for c, deg := range inDegree {
		// Characters with inDegree 0 can be placed first
		if deg == 0 {
			queue = append(queue, c)
		}
	}
	slices.Sort(queue)  // Sort for consistent ordering

	// Build result string using topological sort
	var result strings.Builder
	result.Grow(len(inDegree))  // Pre-allocate capacity for efficiency

	// Process characters in topological order
	for len(queue) > 0 {
		u := queue[0]      // Get character with no dependencies
		queue = queue[1:]  // Remove from queue (dequeue)
		result.WriteByte(u)  // Add to result ordering

		// Process all characters that depend on current character
		for _, v := range adj[u] {
			inDegree[v]--  // Decrease incoming degree (remove dependency)
			// If character has no more dependencies, add to queue
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// Check for cycles in dependency graph
	// If cycle exists, not all characters will be processed
	if result.Len() != len(inDegree) {
		return ""  // Cycle detected, no valid ordering
	}
	return result.String()  // Return the alien alphabet ordering
}
