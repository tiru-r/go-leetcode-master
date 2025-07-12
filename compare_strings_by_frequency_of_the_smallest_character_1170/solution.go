package compare_strings_by_frequency_of_the_smallest_character_1170

// Note: study again.
//
// A good problem for runtime optimization and memory trade-offs.
// By taking memory, we can speed up runtime of the algorithm.

// frequencyOfSmallestChar calculates the frequency of the lexicographically smallest character in a string.
//
// Parameters:
//
//	s: The input string (assumed to contain only lowercase English letters).
//
// Returns:
//
//	The count of the smallest character in the string. Returns 0 if the string is empty.
//
// Example:
//
//	frequencyOfSmallestChar("aabc") returns 2 (smallest char 'a', appears 2 times)
//	frequencyOfSmallestChar("zzza") returns 1 (smallest char 'a', appears 1 time)
func frequencyOfSmallestChar(s string) int {
	if len(s) == 0 {
		return 0
	}
	smallest := s[0]
	frequency := 1
	for i := 1; i < len(s); i++ {
		if s[i] < smallest {
			smallest = s[i]
			frequency = 1
		} else if s[i] == smallest {
			frequency++
		}
	}
	return frequency
}

// numSmallerByFrequency calculates, for each query string, the number of word strings
// whose frequency of the smallest character is strictly greater than the query string's.
//
// This solution uses a counting array to optimize performance, leveraging the fact that
// the maximum frequency is small (typically ≤10 per problem constraints). It avoids sorting
// and binary search, achieving O(n * L1 + m * L2 + maxFreq) time complexity.
//
// Parameters:
//
//	queries: A slice of strings representing the query words.
//	words:   A slice of strings representing the words to compare against.
//
// Returns:
//
//	A slice of integers, where each element ans[i] is the number of words[j]
//	such that f(queries[i]) < f(words[j]), where f(s) is the frequency of the smallest character.
func numSmallerByFrequency(queries []string, words []string) []int {
	// Step 1: Pre-compute frequencies for all query and word strings
	queriesFreq := make([]int, len(queries))
	for i, q := range queries {
		queriesFreq[i] = frequencyOfSmallestChar(q)
	}
	wordsFreq := make([]int, len(words))
	for i, w := range words {
		wordsFreq[i] = frequencyOfSmallestChar(w)
	}

	// Step 2: Use a counting array to store frequency counts of words
	// Since string length is typically ≤10, we use a fixed-size array (0 to 10)
	const maxFreq = 11 // Max frequency is 10, plus 1 for zero-based indexing
	count := make([]int, maxFreq)
	for _, wf := range wordsFreq {
		count[wf]++
	}

	// Step 3: Compute cumulative counts from right to left
	// count[i] will store the number of words with frequency ≥ i
	for i := maxFreq - 2; i >= 0; i-- {
		count[i] += count[i+1]
	}

	// Step 4: For each query frequency, get the count of words with higher frequency
	ans := make([]int, len(queries))
	for i, qf := range queriesFreq {
		if qf+1 < maxFreq {
			ans[i] = count[qf+1] // Number of words with frequency > qf
		} else {
			ans[i] = 0 // No words can have frequency > qf if qf ≥ maxFreq
		}
	}

	return ans
}
