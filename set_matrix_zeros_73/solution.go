package set_matrix_zeros_73

// No imports needed for the optimized solution

func setZeroes(matrix [][]int) {
	// Decide to use the first value of every row and column as the marker
	// for whether the row or column should be set to 0 or not.
	// Since matrix[0][0] can indicate that both 0-th row and column should
	// be set to 0, we need a way to distinguish whether just the 0-th
	// row and column should be set to 0. So, use matrix[0][0] to indicate
	// if row 0 should be set to all zeros, and use an additional variable,
	// setColZero to indicate if col 0 should be set to all zeros.

	setColZero := false
	// Modern range-over-int for ultra-clean matrix traversal
	for r := range len(matrix) {
		for c := range len(matrix[r]) {
			v := matrix[r][c]

			// prevent any column 0 values in the rows from marking matrix[0][0]
			// to 0, which indicates row 0 being marked for 0
			if v == 0 && c == 0 {
				setColZero = true
				continue
			}

			if v == 0 {
				// set the row marker to 0
				matrix[r][0] = 0

				// set the column marker to 0
				matrix[0][c] = 0
			}
		}
	}

	// mark cell to 0 if it's row or column marker is 0, starting
	// at matrix[1][1] and using the markers set earlier
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// mark row 0 with 0's using modern range syntax
	if matrix[0][0] == 0 {
		for c := range len(matrix[0]) {
			matrix[0][c] = 0
		}
	}

	// mark col 0 with 0's using modern range syntax
	if setColZero {
		for r := range len(matrix) {
			matrix[r][0] = 0
		}
	}
}

// REMOVED: Inefficient O(m×n×(m+n)) solution with redundant operations
// This approach had terrible performance because it set entire rows/columns
// immediately upon finding each zero, leading to O(m×n×(m+n)) complexity.
// It also had edge case issues with MaxInt64/MinInt64 values.
// Use the optimized O(m×n) marker-based solution above instead.
