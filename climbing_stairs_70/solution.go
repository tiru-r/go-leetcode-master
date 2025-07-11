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
	// Input validation for negative stairs
	if n < 0 {
		return 0
	}
	// Base cases for n = 0 or n = 1
	if n == 0 || n == 1 {
		return 1
	}

	// Use two variables to store the number of ways to reach the previous two steps
	prev1 := 1 // Ways to reach step i-1
	prev2 := 1 // Ways to reach step i-2

	// Iterate from step 2 to n
	for i := 2; i <= n; i++ {
		// Number of ways to reach step i is the sum of ways to reach step i-1 and i-2
		current := prev1 + prev2
		// Update variables for the next iteration
		prev2 = prev1
		prev1 = current
	}

	// Return the number of ways to reach step n
	return prev1
}
