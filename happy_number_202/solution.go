package happy_number_202

// isHappy reports whether n is a happy number.
// Time: O(log n)  Space: O(log n)
func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 {
		if _, ok := seen[n]; ok {
			return false
		}
		seen[n] = struct{}{}
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
	}
	return true
}
