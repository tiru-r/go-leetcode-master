package counting_bits_338

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := range result[1:] {
		result[i+1] = result[(i+1)&i] + 1
	}
	return result
}
