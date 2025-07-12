package palindrome_partitioning_ii_132

func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	// 1. pal[i][j] == true => s[i:j+1] is palindrome
	pal := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
	}

	// 2. Fill palindrome table
	for length := 1; length <= n; length++ {
		for i := 0; i+length <= n; i++ {
			j := i + length - 1
			if length == 1 {
				pal[i][j] = true
			} else if length == 2 {
				pal[i][j] = s[i] == s[j]
			} else {
				pal[i][j] = s[i] == s[j] && pal[i+1][j-1]
			}
		}
	}

	// Early termination: if entire string is palindrome
	if pal[0][n-1] {
		return 0
	}

	// 3. cuts[i] = minimum cuts for prefix s[0:i+1]
	cuts := make([]int, n)
	for i := range cuts {
		cuts[i] = i // Max cuts = i (cut after each character)
	}

	for i := 0; i < n; i++ {
		if pal[0][i] { // Prefix s[0:i+1] is palindrome
			cuts[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if pal[j+1][i] {
				cuts[i] = min(cuts[i], cuts[j]+1)
			}
		}
	}
	return cuts[n-1]
}
