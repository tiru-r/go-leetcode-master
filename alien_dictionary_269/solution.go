package alien_dictionary_269

import (
	"slices"
	"strings"
)

// alienOrder returns the lexicographical order of an alien alphabet
// given a slice of words sorted by that alphabet.
// Returns "" if no valid order exists.
func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Handle single-word case
	if len(words) == 1 {
		word := words[0]
		if word == "" {
			return ""
		}
		for _, r := range word {
			if r < 'a' || r > 'z' {
				return ""
			}
		}
		return word
	}

	// ---------- 1. Collect characters and initialize in-degree ----------
	adj := map[byte]map[byte]struct{}{} // sparse adjacency list
	inDegree := map[byte]int{}          // in-degree for Kahn's algorithm

	for _, word := range words {
		if word == "" {
			return ""
		}
		for _, r := range word {
			if r < 'a' || r > 'z' {
				return ""
			}
			inDegree[byte(r)] += 0 // Ensure character is tracked
		}
	}

	// ---------- 2. Build the graph ----------
	for i := 0; i < len(words)-1; i++ {
		curr, next := words[i], words[i+1]

		// Invalid case: longer word is prefix of shorter word
		if len(curr) > len(next) && strings.HasPrefix(curr, next) {
			return ""
		}

		// Find first differing character pair
		minLen := min(len(curr), len(next))
		for j := 0; j < minLen; j++ {
			u, v := curr[j], next[j]
			if u == v {
				continue
			}
			if _, ok := adj[u]; !ok {
				adj[u] = make(map[byte]struct{}, 25)
			}
			if _, ok := adj[u][v]; !ok {
				adj[u][v] = struct{}{}
				inDegree[v]++
			}
			break // Only first mismatch matters
		}
	}

	// ---------- 3. Topological sort (Kahn's algorithm) ----------
	zeros := make([]byte, 0, len(inDegree))
	for c := range inDegree {
		if inDegree[c] == 0 {
			zeros = append(zeros, c)
		}
	}
	slices.Sort(zeros) // Ensure lexicographical order for ties

	order := strings.Builder{}
	order.Grow(len(inDegree))

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
	if order.Len() != len(inDegree) {
		return ""
	}
	return order.String()
}
