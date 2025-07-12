package alien_dictionary_269

import (
	"slices"
	"strings"
)

// alienOrder returns the lexicographical order of an alien alphabet
// given a slice of words that are already sorted by that alphabet.
// It returns "" if no valid order exists.
//
// Improvements applied:
//   - sparse adjacency list (map instead of fixed [26][26]bool)
//   - early-cycle check removed (redundant)
//   - prefix check uses strings.HasPrefix
//   - clearer empty-word error handling
func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// ---------- 1. Collect unique characters ----------
	adj := map[byte]map[byte]struct{}{} // sparse adjacency list
	inDegree := map[byte]int{}          // in-degree for Kahn
	exists := map[byte]bool{}

	for _, word := range words {
		if word == "" {
			// Empty string inside the list is explicitly invalid input.
			return ""
		}
		for _, r := range word {
			if r < 'a' || r > 'z' {
				return "" // invalid character
			}
			c := byte(r)
			exists[c] = true
		}
	}

	// ---------- 2. Build the graph ----------
	for i := 0; i < len(words)-1; i++ {
		curr, next := words[i], words[i+1]

		// Invalid case: "abc" comes before "ab"
		if len(curr) > len(next) && strings.HasPrefix(curr, next) {
			return ""
		}

		// Find the first differing character pair
		minLen := min(len(curr), len(next))
		for j := 0; j < minLen; j++ {
			u, v := curr[j], next[j]
			if u == v {
				continue
			}

			if _, ok := adj[u]; !ok {
				adj[u] = map[byte]struct{}{}
			}
			if _, ok := adj[u][v]; !ok {
				adj[u][v] = struct{}{}
				inDegree[v]++
			}
			break // only the first mismatch matters
		}
	}

	// ---------- 3. Topological sort (Kahn) ----------
	zeros := make([]byte, 0, len(exists))
	for c := range exists {
		if inDegree[c] == 0 {
			zeros = append(zeros, c)
		}
	}
	slices.Sort(zeros) // deterministic tie-break for tests

	order := strings.Builder{}
	order.Grow(len(exists))

	for len(zeros) > 0 {
		u := zeros[0]
		zeros = zeros[1:]
		order.WriteByte(u)

		for v := range adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				zeros = append(zeros, v)
			}
		}
	}

	// ---------- 4. Cycle detection ----------
	if order.Len() != len(exists) {
		return ""
	}
	return order.String()
}
