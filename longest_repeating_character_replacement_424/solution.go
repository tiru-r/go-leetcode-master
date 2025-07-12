package longest_repeating_character_replacement_424

func characterReplacement(s string, k int) int {
	const alphabet = 26
	cnt := [alphabet]int{}

	start, maxCnt, maxLen := 0, 0, 0
	for end := 0; end < len(s); end++ {
		idx := s[end] - 'A'
		cnt[idx]++
		maxCnt = max(maxCnt, cnt[idx])

		if end-start+1-maxCnt > k {
			cnt[s[start]-'A']--
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}
