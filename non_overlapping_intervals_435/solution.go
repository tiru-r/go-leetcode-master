package non_overlapping_intervals_435

import (
	"cmp"
	"slices"
)

// DP solution based on sorting by start times.
// Almost identical solution to longest_increasing_subsequence_300, where
// dp[i] stores the max (longest) we *can keep* up to the i-th interval.
//
// Time: O(n^2), Space: O(n)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}

	// sort by interval start times using cmp.Compare for type safety
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// dp[i] stores the max amount of intervals we *can keep*
	// up to the i-th interval. dp[i + 1] is the (max amount from
	// dp[0, i]) + 1 which does not overlap with dp[i + 1].
	dp := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		maxVal := 0
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				maxVal = max(maxVal, dp[j])
			}
		}

		dp[i] = maxVal + 1
	}

	// dp[len(dp)-1] holds the max intervals we can keep, so
	// the difference of all intervals and max we can keep is
	// the min we can erase
	return len(intervals) - dp[len(dp)-1]
}

