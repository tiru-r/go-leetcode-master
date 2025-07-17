package wildcard_matching_44

func IsMatch(s, p string) bool {
	m, n := len(s), len(p)
	
	pattern := make([]byte, 0, n)
	for i := range n {
		if p[i] == '*' && len(pattern) > 0 && pattern[len(pattern)-1] == '*' {
			continue
		}
		pattern = append(pattern, p[i])
	}
	n = len(pattern)
	
	prev := make([]bool, n+1)
	curr := make([]bool, n+1)
	
	prev[0] = true
	
	for j := range n {
		prev[j+1] = prev[j] && pattern[j] == '*'
	}
	
	for i := range m {
		curr[0] = false
		
		for j := range n {
			if pattern[j] == '*' {
				curr[j+1] = curr[j] || prev[j+1]
			} else if pattern[j] == '?' || pattern[j] == s[i] {
				curr[j+1] = prev[j]
			} else {
				curr[j+1] = false
			}
		}
		
		prev, curr = curr, prev
	}
	
	return prev[n]
}