package palindrome_partitioning_ii_132

func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	pal := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
	}

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

	if pal[0][n-1] {
		return 0
	}

	cuts := make([]int, n)
	for i := range cuts {
		cuts[i] = i
	}

	for i := range n {
		if pal[0][i] {
			cuts[i] = 0
			continue
		}
		for j := range i {
			if pal[j+1][i] {
				cuts[i] = min(cuts[i], cuts[j]+1)
			}
		}
	}

	return cuts[n-1]
}
