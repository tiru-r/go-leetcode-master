package first_unique_character_in_a_string_387

func firstUniqChar(s string) int {
	var charCount [26]int

	for i := range len(s) {
		charCount[s[i]-'a']++
	}

	for i := range len(s) {
		if charCount[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
