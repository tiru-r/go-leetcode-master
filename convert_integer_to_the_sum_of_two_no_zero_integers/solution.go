package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"errors"
)

// Split returns two positive integers a and b such that
// a + b == n and neither a nor b contains the digit 0.
// If n <= 1 or no such pair exists, Split reports an error.
func Split(n int) (a, b int, err error) {
	if n <= 1 {
		return 0, 0, errors.New("input must be greater than 1")
	}

	for a = 1; a <= n/2; a++ {
		b = n - a
		if !hasZero(a) && !hasZero(b) {
			return a, b, nil
		}
	}
	return 0, 0, errors.New("no valid pair found")
}

// hasZero reports whether v contains the digit 0.
// It is written so the compiler can inline it and avoid allocations.
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
