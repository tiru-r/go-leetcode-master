package longest_common_prefix_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// Find minimum length with early termination
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
		if minLen == 0 {
			return ""
		}
	}

	// Binary search for optimal length
	left, right := 0, minLen
	for left < right {
		mid := (left + right + 1) / 2
		if hasCommonPrefix(strs, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	
	return strs[0][:left]
}

func hasCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if strs[i][:length] != prefix {
			return false
		}
	}
	return true
}
