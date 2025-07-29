package maximum_number_of_darts_inside_of_a_circular_dartboard_1453

import "math"

const eps = 1e-7

// numPointsInCircle returns the maximum darts that fit inside (or on)
// a circle of radius r using geometric optimization.
func numPointsInCircle(darts [][2]float64, r float64) int {
	n := len(darts)
	if n <= 1 {
		return n
	}

	radiusSquared := r * r
	maxCount := 1

	// Helper: count points within circle centered at given point
	countPointsInCircle := func(center [2]float64) int {
		count := 0
		for _, dart := range darts {
			dx, dy := center[0]-dart[0], center[1]-dart[1]
			if dx*dx+dy*dy <= radiusSquared+eps {
				count++
			}
		}
		return count
	}

	// Strategy 1: Try circles centered at each dart position
	for _, dart := range darts {
		if pointCount := countPointsInCircle(dart); pointCount > maxCount {
			maxCount = pointCount
		}
	}

	// Strategy 2: Try circles whose boundary passes through two darts
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dart1, dart2 := darts[i], darts[j]
			dx, dy := dart2[0]-dart1[0], dart2[1]-dart1[1]
			distanceSquared := dx*dx + dy*dy
			
			// Skip if darts are too far apart for any circle of radius r
			if distanceSquared > 4*radiusSquared+eps {
				continue
			}
			// Skip coincident points
			if distanceSquared < eps {
				continue
			}

			// Calculate circle centers that pass through both darts
			distance := math.Sqrt(distanceSquared)
			height := math.Sqrt(radiusSquared - distanceSquared/4)
			midX, midY := (dart1[0]+dart2[0])/2, (dart1[1]+dart2[1])/2
			// Perpendicular unit vector
			unitX, unitY := -dy/distance, dx/distance

			// Check both possible circle centers
			for _, direction := range [2]float64{1, -1} {
				center := [2]float64{
					midX + direction*height*unitX,
					midY + direction*height*unitY,
				}
				if pointCount := countPointsInCircle(center); pointCount > maxCount {
					maxCount = pointCount
				}
			}
		}
	}
	return maxCount
}
