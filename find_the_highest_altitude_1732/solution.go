package find_the_highest_altitude_1732

// largestAltitude finds the highest altitude reached during the journey
// Optimized: O(n) time, O(1) space using single-pass prefix sum
func largestAltitude(gain []int) int {
	maxAltitude, altitude := 0, 0
	
	// Single pass with modern range iteration
	for _, delta := range gain {
		altitude += delta
		maxAltitude = max(maxAltitude, altitude)
	}
	
	return maxAltitude
}
