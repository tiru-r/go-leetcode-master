package reverse_integer_7

import "math"

func reverse(x int) int {
	result := 0
	
	for x != 0 {
		digit := x % 10
		x /= 10
		
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}
		
		result = result*10 + digit
	}
	
	return result
}
