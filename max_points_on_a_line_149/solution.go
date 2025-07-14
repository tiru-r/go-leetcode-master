package max_points_on_a_line_149


func maxPointsOnLine(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxPoints := 2

	for i := range n {
		slopes := make(map[[2]int]int)
		duplicate := 1
		localMax := 0

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			g := gcd(dx, dy)
			dx /= g
			dy /= g

			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			slope := [2]int{dx, dy}
			slopes[slope]++
			localMax = max(localMax, slopes[slope])
		}

		maxPoints = max(maxPoints, localMax+duplicate)
	}

	return maxPoints
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
