package max_points_on_a_line_149

import "strconv"

func maxPointsOnLine(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxPoints := 2

	for i := 0; i < n; i++ {
		slopes := make(map[string]int)
		duplicate := 1
		localMax := 0

		for j := i + 1; j < n; j++ {
			deltaX := points[j][0] - points[i][0]
			deltaY := points[j][1] - points[i][1]

			if deltaX == 0 && deltaY == 0 {
				duplicate++
				continue
			}

			gcd := gcd(deltaX, deltaY)
			deltaX /= gcd
			deltaY /= gcd

			if deltaX < 0 || (deltaX == 0 && deltaY < 0) {
				deltaX = -deltaX
				deltaY = -deltaY
			}

			slope := strconv.Itoa(deltaX) + "," + strconv.Itoa(deltaY)
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
