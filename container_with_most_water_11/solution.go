package container_with_most_water_11

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		currentArea := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, currentArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
