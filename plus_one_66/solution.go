package plus_one_66

import "slices"

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		incr := digits[i] + carry

		if incr == 10 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] = incr
			carry = 0
		}
	}

	if carry == 1 {
		// Use slices.Insert for efficient prepending
		result := make([]int, 0, len(digits)+1)
		result = slices.Insert(result, 0, 1)
		result = append(result, digits...)
		return result
	}

	return digits
}
