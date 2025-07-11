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

// Alternative O(1) space-optimized solution
func fibSpaceOptimized(N int) int {
	if N < 2 {
		return N
	}

	prev, curr := 0, 1
	for range N - 1 {
		prev, curr = curr, prev+curr
	}
	return curr
}

// High-performance O(log n) matrix exponentiation approach for very large N
func fibMatrix(N int) int {
	if N < 2 {
		return N
	}

	// Matrix multiplication approach: [[1,1],[1,0]]^n
	result := matrixPower([][]int{{1, 1}, {1, 0}}, N)
	return result[0][1]
}

func matrixPower(matrix [][]int, n int) [][]int {
	if n == 1 {
		return matrix
	}

	if n%2 == 0 {
		half := matrixPower(matrix, n/2)
		return matrixMultiply(half, half)
	}

	return matrixMultiply(matrix, matrixPower(matrix, n-1))
}

func matrixMultiply(a, b [][]int) [][]int {
	return [][]int{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

// Original O(2^n) exponential solution kept for comparison (DO NOT USE)
func fibSlow(N int) int {
	if N < 2 {
		return N
	}
	return fibSlow(N-1) + fibSlow(N-2) // EXTREMELY SLOW!
}
