package trapping_rain_water_42

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

// trap2 worked, but ran out of memory on leetcode.
// This was my first pass at the problem.
// I didn't think clearly enough about the memory, which
// could've been a max(height) x len(height) matrix.
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	maxHeight := 0
	for _, h := range height {
		maxHeight = max(maxHeight, h)
	}

	grid := make([][]bool, maxHeight)
	for i := range grid {
		grid[i] = make([]bool, len(height))
	}

	for col, h := range height {
		for row := maxHeight - h; row < maxHeight; row++ {
			grid[row][col] = true
		}
	}

	water := 0
	for row := 0; row < maxHeight; row++ {
		for col := 0; col < len(height); col++ {
			if grid[row][col] {
				for nextCol := col + 1; nextCol < len(height); nextCol++ {
					if grid[row][nextCol] {
						water += nextCol - col - 1
						break
					}
				}
			}
		}
	}

	return water
}
