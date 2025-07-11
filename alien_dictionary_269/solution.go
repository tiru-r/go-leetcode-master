package alien_dictionary_269

import (
	"slices"
	"strings"
)

// alienOrder determines the order of characters in an alien language based on a list of lexicographically sorted words.
// Returns an empty string if no valid order exists.
// This version uses array-based graph representation for optimal performance with fixed-size alphabet,
// and includes modern Go 1.21+ features and robust input validation.
func alienOrder(words []string) string {
	// --- Phase 1: Initialize Structures and Collect Unique Characters ---

	// Handle edge case of no words.
	if len(words) == 0 {
		return ""
	}

	// Fixed-size arrays for graph representation, assuming lowercase English letters ('a' through 'z').
	// Index i corresponds to character 'a' + i.
	adjList := [26][26]bool{} // adjList[u-'a'][v-'a'] = true if u -> v
	inDegrees := [26]int{}    // inDegrees[c-'a'] = count of incoming edges for char c
	exists := [26]bool{}      // exists[c-'a'] = true if char c is present in any word
	totalUniqueChars := 0     // Count of unique characters

	// Validate words and populate 'exists' array and 'totalUniqueChars'.
	for _, word := range words {
		if len(word) == 0 {
			// This indicates an empty string in the input list, which is an invalid condition.
			// Example: ["abc", "", "def"]
			return ""
		}
		for _, r := range word {
			// Validate characters are lowercase English letters.
			if r < 'a' || r > 'z' {
				return "" // Invalid character detected
			}
			if !exists[r-'a'] {
				exists[r-'a'] = true
				totalUniqueChars++
			}
		}
	}

	// --- Phase 2: Build the Graph (Edges and In-Degrees) ---

	for i := 0; i < len(words)-1; i++ {
		wordCurr, wordNext := words[i], words[i+1]
		minLen := min(len(wordCurr), len(wordNext)) // Use built-in min (Go 1.21+)

		// Critical check for invalid prefix ordering: "abc" before "ab" is a contradiction.
		// Iterate manually to avoid slice allocation for sub-strings, maintaining performance.
		if len(wordCurr) > len(wordNext) {
			isPrefix := true
			for j := 0; j < minLen; j++ {
				if wordCurr[j] != wordNext[j] {
					isPrefix = false
					break
				}
			}
			if isPrefix {
				return "" // Invalid order detected
			}
		}

		// Find the first differing character to establish an ordering rule (charCurr -> charNext).
		for r := 0; r < minLen; r++ {
			charCurrIdx, charNextIdx := wordCurr[r]-'a', wordNext[r]-'a'
			if charCurrIdx != charNextIdx {
				// Add edge charCurr -> charNext if it doesn't already exist.
				if !adjList[charCurrIdx][charNextIdx] {
					adjList[charCurrIdx][charNextIdx] = true
					inDegrees[charNextIdx]++ // Increment in-degree of the character that comes later.
				}
				break // Only the first differing pair matters for a given word pair.
			}
		}
	}

	// --- Phase 3: Topological Sort (Kahn's Algorithm) ---

	// Initialize the queue ('zeros' slice) with all characters having an in-degree of 0.
	zeros := make([]byte, 0, 26) // Pre-allocate capacity for efficiency
	for i := 0; i < 26; i++ {
		if exists[i] && inDegrees[i] == 0 { // Only consider characters that were actually in the input words
			zeros = append(zeros, byte('a'+i))
		}
	}

	// Sort the initial `zeros` queue for deterministic output in tests.
	// This does not affect correctness, as any valid topological sort is acceptable, but aids testing.
	slices.Sort(zeros) // Requires Go 1.21+

	// If there are unique characters but no character has an in-degree of 0,
	// it means all existing characters are part of a cycle (or multiple cycles).
	if len(zeros) == 0 && totalUniqueChars > 0 {
		return ""
	}

	// Build the result string using strings.Builder for efficiency.
	var ordering strings.Builder
	ordering.Grow(totalUniqueChars) // Pre-allocate capacity for the final string

	for len(zeros) > 0 {
		v := zeros[0]         // Dequeue the current character
		zeros = zeros[1:]     // Efficiently re-slice for small 'zeros' (max 26 elements)
		ordering.WriteByte(v) // Add to the result

		// For each possible neighbor 'k' of 'v' (represented by its index)
		for k := 0; k < 26; k++ {
			if adjList[v-'a'][k] { // If there's an edge from v to ('a'+k)
				inDegrees[k]-- // Decrement in-degree of the neighbor
				if inDegrees[k] == 0 {
					zeros = append(zeros, byte('a'+k)) // If in-degree becomes 0, enqueue it
				}
			}
		}
	}

	// --- Phase 4: Cycle Detection and Final Result ---

	// If the length of the built ordering is not equal to the total number
	// of unique characters found, it implies a cycle in the graph,
	// meaning not all characters could be topologically sorted.
	if ordering.Len() != totalUniqueChars {
		return ""
	}

	return ordering.String()
}
