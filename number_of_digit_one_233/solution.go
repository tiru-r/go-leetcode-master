package number_of_digit_one_233

func CountDigitOne(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	factor := 1

	for factor <= n {
		higher := n / (factor * 10)
		current := (n / factor) % 10
		lower := n % factor

		switch {
		case current == 0:
			count += higher * factor
		case current == 1:
			count += higher*factor + lower + 1
		default:
			count += (higher + 1) * factor
		}

		factor *= 10
	}

	return count
}
