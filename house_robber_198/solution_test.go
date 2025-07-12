package house_robber_198

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// 1. LeetCode official samples
		{"LC example 1", []int{1, 2, 3, 1}, 4},
		{"LC example 2", []int{2, 7, 9, 3, 1}, 12},

		// 2. Empty / single house
		{"empty", []int{}, 0},
		{"single house", []int{5}, 5},

		// 3. Two houses
		{"two houses", []int{2, 1}, 2},
		{"two equal", []int{3, 3}, 3},

		// 4. All zeroes
		{"all zeroes", []int{0, 0, 0, 0}, 0},

		// 5. Monotonic increasing
		{"increasing", []int{1, 2, 3, 4, 5}, 9},

		// 6. Monotonic decreasing
		{"decreasing", []int{5, 4, 3, 2, 1}, 9},

		// 7. Alternating high/low
		{"alternating", []int{2, 1, 2, 1, 2}, 6},

		// 8. Large values (performance smoke)
		{"large 100", make([]int, 100), 0}, // all zeroes
		{"large 100 random", func() []int {
			out := make([]int, 100)
			for i := range out {
				out[i] = i + 1
			}
			return out
		}(), 2550}, // 1+3+5+…+99 = 2500, but DP = 2550
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := rob(tc.nums)
			if got != tc.want {
				t.Errorf("%s: rob(%v) = %d, want %d", tc.name, tc.nums, got, tc.want)
			}
		})
	}
}
