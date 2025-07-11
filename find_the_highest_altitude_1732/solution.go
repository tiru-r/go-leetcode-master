package find_the_highest_altitude_1732

// -4 -3 -2 -1   4  3  2
//
//	^
//
// -4 -7 -9 -10 -6 -3  -1
// Ultra-efficient O(n) time, O(1) space - direct processing without allocation
func largestAltitude(gain []int) int {
	highest := 0
	currentAltitude := 0
	
	// Process gains directly without intermediate storage
	for _, g := range gain {
		currentAltitude += g
		highest = max(highest, currentAltitude)
	}
	
	return highest
}
