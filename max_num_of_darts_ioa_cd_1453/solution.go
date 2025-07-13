package maximum_number_of_darts_inside_of_a_circular_dartboard_1453

import (
	"math"
)

const EPS = 1e-7

// Point represents a 2D coordinate
type Point struct {
	x, y float64
}

// NumPointsInCircle finds the maximum number of darts that can be enclosed by a circle of radius r
// Time: O(n³), Space: O(1)
func NumPointsInCircle(darts [][]int, r int) int {
	n := len(darts)
	if n <= 1 {
		return n
	}
	
	points := make([]Point, n)
	for i, dart := range darts {
		points[i] = Point{float64(dart[0]), float64(dart[1])}
	}
	
	maxPoints := 1
	radius := float64(r)
	
	// Try all pairs of points to find circle centers
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			centers := getCircleCenters(points[i], points[j], radius)
			
			for _, center := range centers {
				count := 0
				for k := 0; k < n; k++ {
					if distance(center, points[k]) <= radius+EPS {
						count++
					}
				}
				maxPoints = max(maxPoints, count)
			}
		}
	}
	
	return maxPoints
}

// getCircleCenters finds the centers of circles with given radius passing through two points
func getCircleCenters(p1, p2 Point, r float64) []Point {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	d := math.Sqrt(dx*dx + dy*dy)
	
	// Points are too far apart for a circle of radius r
	if d > 2*r+EPS {
		return []Point{}
	}
	
	// Points are the same or very close
	if d < EPS {
		return []Point{}
	}
	
	// Midpoint between the two points
	mx := (p1.x + p2.x) / 2
	my := (p1.y + p2.y) / 2
	
	// Distance from midpoint to circle centers
	h := math.Sqrt(r*r - (d/2)*(d/2))
	
	// Unit vector perpendicular to the line connecting p1 and p2
	ux := -dy / d
	uy := dx / d
	
	// Two possible circle centers
	c1 := Point{mx + h*ux, my + h*uy}
	c2 := Point{mx - h*ux, my - h*uy}
	
	if math.Abs(h) < EPS {
		return []Point{c1} // Only one center when points are diameter apart
	}
	
	return []Point{c1, c2}
}

// distance calculates Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}
