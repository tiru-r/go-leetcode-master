package compare_strings_by_frequency_of_the_smallest_character_1170

// Note: study again.
//
// A good problem for runtime optimization and memory trade-offs.
// By taking memory, we can speed up runtime of the algorithm.

func numSmallerByFrequency(queries, words []string) []int {
	const maxF = 10        // problem guarantees length ≤ 10
	cnt := [maxF + 2]int{} // indices 0..10 + sentinel

	// 1. one-pass frequency histogram for words
	for _, w := range words {
		minC := byte('z' + 1)
		count := 0
		for i := 0; i < len(w); i++ {
			c := w[i]
			switch {
			case c < minC:
				minC, count = c, 1
			case c == minC:
				count++
			}
		}
		if count > 0 { // empty string already skipped by problem
			cnt[count]++
		}
	}

	// 2. suffix sum (prefix of suffix) – O(11)
	for i := maxF; i >= 0; i-- {
		cnt[i] += cnt[i+1]
	}

	// 3. answer queries – O(M)
	ans := make([]int, len(queries))
	for i, q := range queries {
		minC := byte('z' + 1)
		count := 0
		for j := 0; j < len(q); j++ {
			c := q[j]
			switch {
			case c < minC:
				minC, count = c, 1
			case c == minC:
				count++
			}
		}
		if count <= maxF {
			ans[i] = cnt[count+1]
		}
	}
	return ans
}
