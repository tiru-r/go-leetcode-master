package the_skyline_problem_218

import (
	"reflect"
	"testing"
)

func TestGetSkyline(t *testing.T) {
	tests := []struct {
		name      string
		buildings [][]int
		want      [][]int
	}{
		{"empty", [][]int{}, [][]int{}},
		{"single", [][]int{{0, 5, 7}}, [][]int{{0, 7}, {5, 0}}},
		{"classic", [][]int{
			{2, 9, 10}, {3, 7, 15}, {5, 12, 12},
			{15, 20, 10}, {19, 24, 8},
		}, [][]int{
			{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0},
		}},
		{"no overlap", [][]int{{1, 2, 1}, {3, 4, 2}}, [][]int{{1, 1}, {2, 0}, {3, 2}, {4, 0}}},
		{"nested", [][]int{{1, 10, 3}, {2, 5, 8}, {6, 8, 5}}, [][]int{{1, 3}, {2, 8}, {5, 3}, {6, 5}, {8, 3}, {10, 0}}},
		{"all same boundaries", [][]int{{0, 5, 3}, {0, 5, 4}}, [][]int{{0, 4}, {5, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetSkyline(tc.buildings)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("getSkyline(%v) = %v; want %v", tc.buildings, got, tc.want)
			}
		})
	}
}

func BenchmarkGetSkyline(b *testing.B) {
	buildings := [][]int{
		{2, 9, 10}, {3, 7, 15}, {5, 12, 12},
		{15, 20, 10}, {19, 24, 8}, {1, 30, 5},
		{25, 35, 20}, {30, 40, 15}, {35, 50, 12},
		{45, 55, 18}, {50, 60, 25}, {55, 65, 8},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetSkyline(buildings)
	}
}
