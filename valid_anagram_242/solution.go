package valid_anagram_242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count [26]int
	for i := range len(s) {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
