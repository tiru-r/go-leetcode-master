package maximum_length_of_a_concatenated_string_with_unique_characters_1239

// Note: Study again. Good problem with bit manipulation to detect
// if a string has unique characters.
func maxLength(arr []string) int {
	return explore(arr, 0, 0)
}

func explore(arr []string, index int, usedChar int) int {
	if index == len(arr) {
		return 0
	}

	// does arr[index] have any characters with an occurrence greater than 1?
	currentUsedChar := 0
	s := arr[index]
	valid := true
	for i := 0; i < len(s); i++ {
		// currentlyUsedChar is a bit mask which represents which characters
		// have been used. In the mask, each used character (flipped on by shifting
		// (1 << (s[i] - 'a') the distance from char to 'a'.

		// If there is an overlapping character '00000100' & '00000100', the
		// bitwise AND will *not* return 0, indicating that we've seen a duplicate.
		exists := currentUsedChar & (1 << (s[i] - 'a'))
		if exists == 0 {
			// update currentUsedChar with bitwise OR, which will keep
			// old 1 bits and add new 1 bits to the mask.
			currentUsedChar |= 1 << (s[i] - 'a')
		} else {
			valid = false
			break
		}
	}

	output := explore(arr, index+1, usedChar)
	if valid && (usedChar&currentUsedChar) == 0 {
		output = max(output, len(s)+explore(arr, index+1, usedChar|currentUsedChar))
	}

	return output
}
