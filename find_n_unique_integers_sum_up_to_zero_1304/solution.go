package find_n_unique_integers_sum_up_to_zero_1304

func sumZero(n int) []int {
	result := make([]int, 0, n)
	
	for i := 1; i <= n/2; i++ {
		result = append(result, i, -i)
	}
	
	if n&1 == 1 {
		result = append(result, 0)
	}
	
	return result
}
