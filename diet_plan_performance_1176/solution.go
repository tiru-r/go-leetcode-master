package diet_plan_performance_1176

func dietPlanPerformance(calories []int, k, lower, upper int) int {
	if k <= 0 || k > len(calories) {
		return 0
	}

	// first window
	sum := 0
	for i := 0; i < k; i++ {
		sum += calories[i]
	}

	points := 0
	if sum < lower {
		points--
	} else if sum > upper {
		points++
	}

	// sliding the window
	for i := k; i < len(calories); i++ {
		sum += calories[i] - calories[i-k]
		if sum < lower {
			points--
		} else if sum > upper {
			points++
		}
	}
	return points
}
