package compare_strings_by_frequency_of_the_smallest_character_1170

// f returns the frequency of the lexicographically smallest character in s.
func f(s string) int {
	if len(s) == 0 {
		return 0
	}

	minChar := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		switch c := s[i]; {
		case c < minChar:
			minChar, count = c, 1
		case c == minChar:
			count++
		}
	}
	return count
}

// numSmallerByFrequency returns, for every query string q, how many words w
// satisfy f(q) < f(w).
//
//   - Time   : O(|queries| + |words|)
//   - Memory : O(1)
func numSmallerByFrequency(queries []string, words []string) []int {
	const maxF = 10 // max possible value of f for any string with len <= 10

	// cnt[i] = #words whose f equals i
	cnt := [maxF + 1]int{}
	for _, w := range words {
		cnt[f(w)]++
	}

	// cnt[i] = #words whose f is >= i
	for i := maxF - 1; i >= 0; i-- {
		cnt[i] += cnt[i+1]
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		fq := f(q)
		if fq < maxF {
			res[i] = cnt[fq+1]
		}
		// if fq == maxF, no word can have a larger f, so res[i] remains 0.
	}
	return res
}
