package climbing_stairs_70

// climbStairs calculates the number of distinct ways to climb to the top of n stairs.
// At each step, you can either climb 1 or 2 steps.
// This solution uses an iterative approach with constant space, tracking only the last two values.
// It also includes input validation for negative values.
//
// Parameters:
//
//	n: The total number of stairs.
//
// Returns:
//
//	The number of distinct ways to climb to the top.
func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}

	// Fibonacci sequence optimization
	prev2, prev1 := 1, 1

	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}

	return prev1
}
