package happy_number_202

// isHappy reports whether n is a happy number using optimized cycle detection
// Optimized: O(log n) time, O(1) space using Floyd's tortoise-hare algorithm
func isHappy(n int) bool {
	// Precomputed lookup for single digits (major optimization)
	if n < 10 {
		return n == 1 || n == 7
	}
	
	// Floyd's cycle detection: O(1) space instead of O(log n) hashmap
	slow, fast := n, n
	for {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(digitSquareSum(fast))
		
		if slow == 1 || fast == 1 {
			return true
		}
		if slow == fast {
			return false // cycle detected
		}
	}
}

// digitSquareSum computes sum of squares of digits efficiently
func digitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
