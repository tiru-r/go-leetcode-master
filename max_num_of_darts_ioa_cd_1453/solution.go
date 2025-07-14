package maximum_number_of_darts_inside_of_a_circular_dartboard_1453

import "math"

const eps = 1e-7

type point struct {
	x, y float64
}

// numPointsInCircle finds the maximum number of darts that can be enclosed by a circle of radius r.
// Uses geometric approach: for optimal circle, at least 2 points must be on the boundary.
// Time: O(n³), Space: O(n)
func numPointsInCircle(darts [][]int, r int) int {
	n := len(darts)
	if n <= 1 {
		return n
	}

	points := make([]point, n)
	for i, dart := range darts {
		points[i] = point{float64(dart[0]), float64(dart[1])}
	}

	maxCount := 1
	radius := float64(r)

	// Try all pairs of points as potential boundary points
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			centers := getCircleCenters(points[i], points[j], radius)
			for _, center := range centers {
				maxCount = max(maxCount, countPointsInCircle(points, center, radius))
			}
		}

		// Also try circles centered at each point
		maxCount = max(maxCount, countPointsInCircle(points, points[i], radius))
	}

	return maxCount
}

// getCircleCenters finds centers of circles with given radius passing through two points
func getCircleCenters(p1, p2 point, r float64) []point {
	dx, dy := p2.x-p1.x, p2.y-p1.y
	d := math.Sqrt(dx*dx + dy*dy)

	// Points too far apart or coincident
	if d > 2*r+eps || d < eps {
		return nil
	}

	// Midpoint and perpendicular distance
	mx, my := (p1.x+p2.x)/2, (p1.y+p2.y)/2
	h := math.Sqrt(r*r - (d/2)*(d/2))

	// Perpendicular unit vector
	ux, uy := -dy/d, dx/d

	centers := []point{
		{mx + h*ux, my + h*uy},
		{mx - h*ux, my - h*uy},
	}

	// If points are diameter apart, return only one center
	if math.Abs(h) < eps {
		return centers[:1]
	}

	return centers
}

// countPointsInCircle counts points within or on the circle boundary
func countPointsInCircle(points []point, center point, radius float64) int {
	count := 0
	for _, p := range points {
		if distanceSquared(center, p) <= radius*radius+eps {
			count++
		}
	}
	return count
}

// distanceSquared calculates squared Euclidean distance (avoids sqrt for performance)
func distanceSquared(p1, p2 point) float64 {
	dx, dy := p1.x-p2.x, p1.y-p2.y
	return dx*dx + dy*dy
}
