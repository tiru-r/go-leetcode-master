package select_cells_grid_3276

import (
	"math"
	"slices"
)

func maxScore(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	valueToRows := make(map[int][]int)
	for r := range len(grid) {
		for c := range len(grid[r]) {
			v := grid[r][c]
			if !slices.Contains(valueToRows[v], r) {
				valueToRows[v] = append(valueToRows[v], r)
			}
		}
	}

	values := make([]int, 0, len(valueToRows))
	for v := range valueToRows {
		values = append(values, v)
	}
	slices.SortFunc(values, func(a, b int) int { return b - a })

	memo := make(map[uint64]int)
	var solve func(valueIdx int, usedRows uint32, hasSelected bool) int
	solve = func(valueIdx int, usedRows uint32, hasSelected bool) int {
		if valueIdx >= len(values) {
			if hasSelected {
				return 0
			}
			return math.MinInt32
		}

		key := uint64(valueIdx)<<32 | uint64(usedRows)<<1
		if hasSelected {
			key |= 1
		}
		if result, exists := memo[key]; exists {
			return result
		}

		value := values[valueIdx]
		result := solve(valueIdx+1, usedRows, hasSelected)

		for _, row := range valueToRows[value] {
			if (usedRows & (1 << row)) == 0 {
				takeValue := value + solve(valueIdx+1, usedRows|(1<<row), true)
				result = max(result, takeValue)
				break
			}
		}

		memo[key] = result
		return result
	}

	return solve(0, 0, false)
}
