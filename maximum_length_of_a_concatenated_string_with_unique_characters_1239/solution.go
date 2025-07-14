package maximum_length_of_a_concatenated_string_with_unique_characters_1239

func maxLength(arr []string) int {
	return dfs(arr, 0, 0)
}

func dfs(arr []string, idx int, used int) int {
	if idx == len(arr) {
		return 0
	}

	skip := dfs(arr, idx+1, used)
	
	mask := charMask(arr[idx])
	if mask == -1 || used&mask != 0 {
		return skip
	}
	
	take := len(arr[idx]) + dfs(arr, idx+1, used|mask)
	return max(skip, take)
}

func charMask(s string) int {
	mask := 0
	for _, c := range s {
		bit := 1 << (c - 'a')
		if mask&bit != 0 {
			return -1
		}
		mask |= bit
	}
	return mask
}
