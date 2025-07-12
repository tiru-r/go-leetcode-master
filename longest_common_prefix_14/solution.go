package longest_common_prefix_14

func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	// find the shortest string once
	min := len(strs[0])
	for _, s := range strs[1:] {
		if len(s) < min {
			min = len(s)
		}
		if min == 0 {
			return ""
		}
	}

	// byte-by-byte scan (cache-friendly)
	for i := 0; i < min; i++ {
		c := strs[0][i]
		for _, s := range strs[1:] {
			if s[i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:min]
}
