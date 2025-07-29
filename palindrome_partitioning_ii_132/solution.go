package palindrome_partitioning_ii_132

// minCut returns the minimum cuts needed to partition s
// so that every substring is a palindrome.
func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}

	// 1. Build pal[i][j] – true if s[i..j] is a palindrome
	pal := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
	}
	for length := 1; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			switch length {
			case 1:
				pal[i][j] = true
			case 2:
				pal[i][j] = s[i] == s[j]
			default:
				pal[i][j] = s[i] == s[j] && pal[i+1][j-1]
			}
		}
	}
	if pal[0][n-1] {
		return 0
	}

	// 2. DP for minimum cuts
	cuts := make([]int, n)
	for i := range cuts {
		cuts[i] = i // worst case: cut every char
	}
	for i := range n {
		if pal[0][i] {
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
