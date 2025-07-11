package first_unique_character_in_a_string_387


func firstUniqChar(s string) int {
	charMap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		idx, ok := charMap[s[i]]

		if ok && idx != -1 {
			charMap[s[i]] = -1
		}

		if !ok {
			charMap[s[i]] = i
		}
	}

	allNonUnique := true
	minUnique := len(s) // Use string length as max possible index
	for _, val := range charMap {
		if val != -1 {
			allNonUnique = false
			minUnique = min(minUnique, val)
		}
	}

	if allNonUnique {
		return -1
	}

	return minUnique
}
