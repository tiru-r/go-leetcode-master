package container_with_most_water_11

// maxArea calculates the maximum area of water that can be trapped between two lines,
// given an array of heights representing vertical lines on a graph.
// The area is determined by the width (distance between indices) and the height of the shorter line.
//
// The function uses a two-pointer approach:
// - Start with pointers at both ends of the array.
// - Calculate the area based on the shorter line and the width between pointers.
// - Move the pointer of the shorter line inward to potentially find a taller line.
// - Track the maximum area found.
//
// Time Complexity: O(n), where n is the length of the height array, as each pointer moves at most n steps.
// Space Complexity: O(1), using only a constant amount of extra space.
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0 // Early return for invalid input (less than 2 lines can't form a container).
	}

	left, right := 0, len(height)-1
	maxWater := 0 // Descriptive variable name to store the maximum area.

	for left < right {
		// Calculate width as the distance between the two pointers.
		width := right - left

		// Determine the height of the container (limited by the shorter line).
		var minHeight int
		if height[left] < height[right] {
			minHeight = height[left]
			left++ // Move left pointer inward to try a taller line.
		} else {
			minHeight = height[right]
			right-- // Move right pointer inward to try a taller line.
		}

		// Calculate the current area and update maxWater if larger.
		currentArea := width * minHeight
		maxWater = max(maxWater, currentArea)
	}

	return maxWater
}
