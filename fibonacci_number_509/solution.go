package fibonacci_number_509

func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for range n - 1 {
		a, b = b, a+b
	}
	return b
}
