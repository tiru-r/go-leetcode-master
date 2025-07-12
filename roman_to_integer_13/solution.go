package roman_to_integer_13

func romanToInt(s string) int {
	var values [256]int
	values['I'] = 1
	values['V'] = 5
	values['X'] = 10
	values['L'] = 50
	values['C'] = 100
	values['D'] = 500
	values['M'] = 1000

	total := 0
	prev := 0

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
