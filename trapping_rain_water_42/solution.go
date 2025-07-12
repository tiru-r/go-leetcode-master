package trapping_rain_water_42

// trap returns the total units of water trapped after raining.
// height[i] is the non-negative height of the wall at index i.
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]
	water := 0

	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
		}
	}
	return water
}
