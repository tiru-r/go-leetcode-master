package fibonacci_number_509

// Modern O(n) iterative solution with Go 1.24+ range-over-int
func fib(N int) int {
	if N < 2 {
		return N
	}

	a, b := 0, 1
	for range N - 1 { // Modern range-over-int pattern
		a, b = b, a+b
	}
	return b
}

