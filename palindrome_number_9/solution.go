package palindrome_number_9

// Modern solution avoiding floating-point math.Pow

func isPalindrome(x int) bool {
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
