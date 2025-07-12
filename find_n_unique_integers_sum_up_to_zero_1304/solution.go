package find_n_unique_integers_sum_up_to_zero_1304

func sumZero(n int) []int {
	out := make([]int, n)
	for i := 0; i < n/2; i++ {
		out[2*i], out[2*i+1] = i+1, -(i + 1)
	}
	if n&1 == 1 {
		out[n-1] = 0
	}
	return out
}
