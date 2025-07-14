package insert_interval_57

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		// 1. LeetCode official
		{"LC ex1", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{"LC ex2", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},

		// 2. Empty input
		{"empty intervals", [][]int{}, []int{2, 5}, [][]int{{2, 5}}},

		// 3. Insert before all
		{"before all", [][]int{{5, 7}}, []int{1, 2}, [][]int{{1, 2}, {5, 7}}},

		// 4. Insert after all
		{"after all", [][]int{{1, 2}}, []int{5, 7}, [][]int{{1, 2}, {5, 7}}},

		// 5. Single interval merge
		{"merge single", [][]int{{1, 5}}, []int{2, 7}, [][]int{{1, 7}}},

		// 6. Touching boundary (merges)
		{"touching boundary", [][]int{{1, 2}, {4, 5}}, []int{2, 3}, [][]int{{1, 3}, {4, 5}}},

		// 7. Fully contained
		{"fully contained", [][]int{{1, 10}}, []int{3, 7}, [][]int{{1, 10}}},

		// 8. Fully containing
		{"fully containing", [][]int{{3, 7}}, []int{1, 10}, [][]int{{1, 10}}},

		// 9. Touching edges
		{"touch edges", [][]int{{1, 2}, {4, 5}}, []int{2, 4}, [][]int{{1, 5}}},

		// 10. Large data (performance smoke) - insert at end
		{"large 1000", func() [][]int {
			large := make([][]int, 1000)
			for i := range large {
				large[i] = []int{i * 2, i*2 + 1}
			}
			return large
		}(), []int{3000, 3001}, func() [][]int {
			result := make([][]int, 1001)
			for i := 0; i < 1000; i++ {
				result[i] = []int{i * 2, i*2 + 1}
			}
			result[1000] = []int{3000, 3001}
			return result
		}()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s:\ninput  %v new %v\ngot    %v\nwant   %v",
					tt.name, tt.intervals, tt.newInterval, got, tt.want)
			}
		})
	}
}
