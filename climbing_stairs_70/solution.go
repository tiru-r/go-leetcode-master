package climbing_stairs_70

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	
	a, b := 1, 1
	for range n - 1 {
		a, b = b, a+b
	}
	
	return b
}
