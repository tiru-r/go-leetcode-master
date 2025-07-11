package palindrome_number_9

// Ultra-modern O(log n) solution - reverse only half the number for optimal performance
func isPalindrome(x int) bool {
	// Negative numbers and numbers ending in 0 (except 0 itself) cannot be palindromes
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// For even length: x == reversedHalf
	// For odd length: x == reversedHalf/10 (remove middle digit)
	return x == reversedHalf || x == reversedHalf/10
}

// Alternative string-based approach (simpler but slightly slower due to allocation)
func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	s := []byte{}
	for x > 0 {
		s = append(s, byte('0'+x%10))
		x /= 10
	}

	// Compare from both ends using modern range-over-int
	for i := range len(s) / 2 {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Original complex solution kept for comparison (DO NOT USE - 3x slower)
func isPalindromeOld(x int) bool {
	if x < 0 {
		return false
	}
	n := x

	// count the number of digits in x
	digits := 0
	for n != 0 {
		n = n / 10
		digits++
	}

	// for each digits / 2 times, compare left hand side digits to right hand side digits
	place := 0
	for i := 1; i <= digits/2; i++ {
		hold := x
		// Use integer power calculation instead of floating-point math.Pow
		powerOf10 := 1
		for k := 0; k < place; k++ {
			powerOf10 *= 10
		}
		ld := (hold / powerOf10) % 10

		for j := 0; j < digits-i; j++ {
			hold = hold / 10
		}

		if ld != hold%10 {
			return false
		}

		place++
	}

	return true
}
