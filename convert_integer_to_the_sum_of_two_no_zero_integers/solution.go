package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"errors"
	"strconv"
	"strings"
)

// GetNoZeroIntegers finds two positive integers A and B such that A + B = n
// and neither A nor B contains the digit '0'.
// It returns an error if no valid pair is found or if n is invalid (n <= 1).
func GetNoZeroIntegers(n int) ([]int, error) {
	if n <= 1 {
		return nil, errors.New("input must be greater than 1")
	}

	for a := 1; a <= n/2; a++ {
		b := n - a
		if !hasZeroDigit(a) && !hasZeroDigit(b) {
			return []int{a, b}, nil
		}
	}

	return nil, errors.New("no valid pair found")
}

// hasZeroDigit checks if a positive integer contains the digit '0'.
// It uses string conversion for clarity and efficiency.
func hasZeroDigit(n int) bool {
	if n <= 0 {
		return true // Treat 0 or negative as containing zero
	}
	return strings.Contains(strconv.Itoa(n), "0")
}
