package container_with_most_water_11

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	
	left, right := 0, len(height)-1
	maxWater := 0
	
	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxWater = max(maxWater, area)
		
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	
	return maxWater
}
