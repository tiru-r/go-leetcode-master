package convert_integer_to_the_sum_of_two_no_zero_integers_1317

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return nil
}

func hasZero(v int) bool {
	if v <= 0 {
		return true
	}
	for v > 0 {
		if v%10 == 0 {
			return true
		}
		v /= 10
	}
	return false
}

// Split returns two positive integers a and b such that
// a + b == n and neither a nor b contains the digit 0.
func Split(n int) (a, b int, err error) {
	if n <= 1 {
		return 0, 0, &ValidationError{"input must be greater than 1"}
	}

	result := getNoZeroIntegers(n)
	if result == nil {
		return 0, 0, &ValidationError{"no valid pair found"}
	}
	return result[0], result[1], nil
}

type ValidationError struct {
	msg string
}

func (e *ValidationError) Error() string {
	return e.msg
}
