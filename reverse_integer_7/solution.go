package reverse_integer_7

// Modern solution using built-in constants (Go 1.24+)

func reverse(x int) int {
	var out int64
	for x != 0 {
		out = (out * 10) + int64(x%10)
		x = x / 10
	}

	// if out overflows positive (2^32 - 1) or negative (-2^31) 32bit integer
	// Use built-in 32-bit integer limits instead of math package
	if out > (1<<31-1) || out < (-1<<31) {
		return 0
	}

	return int(out)
}
