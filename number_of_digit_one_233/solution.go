package number_of_digit_one_233

// CountDigitOne counts the number of digit 1 appearing in all non-negative integers less than or equal to n.
// Time: O(log n), Space: O(1)
func CountDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	digit := 1
	
	for digit <= n {
		higher := n / (digit * 10)
		cur := (n / digit) % 10
		lower := n % digit
		
		if cur == 0 {
			count += higher * digit
		} else if cur == 1 {
			count += higher*digit + lower + 1
		} else {
			count += (higher + 1) * digit
		}
		
		digit *= 10
	}
	
	return count
}
