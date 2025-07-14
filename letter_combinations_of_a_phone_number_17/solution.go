package letter_combinations_of_a_phone_number_17

// letterCombinations generates all letter combinations for phone number digits
// Optimized: O(3^m * 4^n) time, O(3^m * 4^n) space with Go 1.24 modern syntax
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	// Constant lookup table with Go 1.24 syntax
	digitToLetters := [10]string{
		2: "abc",
		3: "def", 
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	// Calculate exact capacity and validate
	capacity := 1
	for _, digit := range digits {
		d := int(digit - '0')
		if d < 2 || d > 9 {
			return nil
		}
		capacity *= len(digitToLetters[d])
	}

	// Pre-allocate with exact capacity
	result := make([]string, 0, capacity)
	combination := make([]byte, len(digits))

	// Backtracking with early termination
	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(digits) {
			result = append(result, string(combination))
			return
		}
		
		letters := digitToLetters[digits[index]-'0']
		for i := range letters {
			combination[index] = letters[i]
			backtrack(index + 1)
		}
	}

	backtrack(0)
	return result
}
