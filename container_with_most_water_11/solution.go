package container_with_most_water_11

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		
		if height[left] < height[right] {
			maxArea = max(maxArea, width*height[left])
			left++
		} else {
			maxArea = max(maxArea, width*height[right])
			right--
		}
	}

	return maxArea
}
