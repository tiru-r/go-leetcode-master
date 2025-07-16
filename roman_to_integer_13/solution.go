package roman_to_integer_13

func romanToInt(s string) int {
	values := [256]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	total, prev := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		curr := values[s[i]]
		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}
	return total
}
