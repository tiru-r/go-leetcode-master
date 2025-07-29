package select_cells_grid_3276

import (
	"math"
	"slices"
)

// maxScore returns the maximum score achievable by selecting cells with
// distinct values from different rows using bitmask DP.
func maxScore(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// value -> distinct rows that contain it
	valueToRows := make(map[int][]int)
	for r, row := range grid {
		seen := make(map[int]bool)
		for _, v := range row {
			if !seen[v] {
				seen[v] = true
				valueToRows[v] = append(valueToRows[v], r)
			}
		}
	}

	// unique values sorted descending
	vals := make([]int, 0, len(valueToRows))
	for v := range valueToRows {
		vals = append(vals, v)
	}
	slices.SortFunc(vals, func(a, b int) int { return b - a })

	memo := make(map[uint64]int)

	// key = (idx<<33) | (mask<<1) | hasSel
	key := func(idx int, mask uint32, hasSel bool) uint64 {
		k := uint64(idx)<<33 | uint64(mask)<<1
		if hasSel {
			k |= 1
		}
		return k
	}

	var dfs func(idx int, mask uint32, hasSel bool) int
	dfs = func(idx int, mask uint32, hasSel bool) int {
		if idx == len(vals) {
			if hasSel {
				return 0
			}
			return math.MinInt32
		}
		if v, ok := memo[key(idx, mask, hasSel)]; ok {
			return v
		}

		// skip current value
		res := dfs(idx+1, mask, hasSel)

		// take current value from any free row
		v := vals[idx]
		for _, r := range valueToRows[v] {
			if mask&(1<<r) == 0 {
				res = max(res, v+dfs(idx+1, mask|(1<<r), true))
			}
		}
		memo[key(idx, mask, hasSel)] = res
		return res
	}

	return dfs(0, 0, false)
}
