package diet_plan_performance_1176

func dietPlanPerformance(calories []int, k, lower, upper int) int {
	n := len(calories)
	if k > n {
		return 0
	}
	
	sum, points := 0, 0
	
	for i := 0; i < n; i++ {
		sum += calories[i]
		if i >= k {
			sum -= calories[i-k]
		}
		if i >= k-1 {
			if sum < lower {
				points--
			} else if sum > upper {
				points++
			}
		}
	}
	
	return points
}
