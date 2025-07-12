package find_n_unique_integers_sum_up_to_zero_1304

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumZero(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"n = 0", 0},
		{"n = 1", 1},
		{"n = 2", 2},
		{"n = 3", 3},
		{"n = 4", 4},
		{"n = 5", 5},
		{"n = 6", 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sumZero(tc.n)

			// 1. Correct length
			assert.Len(t, got, tc.n)

			// 2. All elements unique
			seen := make(map[int]struct{}, len(got))
			for v := range got {
				seen[v] = struct{}{}
			}
			assert.Len(t, seen, len(got), "slice must contain only unique integers")

			// 3. Sum must be zero
			sum := 0
			for _, v := range got {
				sum += v
			}
			assert.Zero(t, sum, "sum of all elements must be zero")
		})
	}
}
