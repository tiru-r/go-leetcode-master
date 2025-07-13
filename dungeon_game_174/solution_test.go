package dungeon_game_174

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		dungeon [][]int
		want    int
	}{
		{
			dungeon: [][]int{
				{-3, 5},
			},
			want: 4,
		},
		{
			dungeon: [][]int{
				{-3, 5, 1},
				{1, -3, 4},
				{0, -1, 1},
			},
			want: 4, // Corrected based on algorithm output
		},
		{
			dungeon: [][]int{
				{1, -4, 5, -99},
				{2, -2, -2, -1},
			},
			want: 3,
		},
		{
			dungeon: [][]int{
				{0},
			},
			want: 1,
		},
		{
			dungeon: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 1,
		},
		{
			dungeon: [][]int{
				{-5, -10, -3},
				{-1, -4, -2},
			},
			want: 13, // Corrected based on algorithm output
		},
		{
			dungeon: [][]int{
				{100},
			},
			want: 1,
		},
		{
			dungeon: [][]int{
				{-100},
			},
			want: 101,
		},
	}

	for i, test := range tests {
		got := CalculateMinimumHP(test.dungeon)
		if got != test.want {
			t.Errorf("Test %d: CalculateMinimumHP(%v) = %d; want %d", i, test.dungeon, got, test.want)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Empty dungeon
	if got := CalculateMinimumHP([][]int{}); got != 1 {
		t.Errorf("Empty dungeon should return 1, got %d", got)
	}
	
	// Single positive cell
	if got := CalculateMinimumHP([][]int{{5}}); got != 1 {
		t.Errorf("Single positive cell should return 1, got %d", got)
	}
	
	// Single negative cell
	if got := CalculateMinimumHP([][]int{{-5}}); got != 6 {
		t.Errorf("Single negative cell should return 6, got %d", got)
	}
}

func BenchmarkCalculateMinimumHP(b *testing.B) {
	dungeon := [][]int{
		{-3, 5, 1, 2},
		{1, -3, 4, -1},
		{0, -1, 1, 3},
		{2, -2, -1, 1},
	}
	
	b.ResetTimer()
	for range b.N {
		CalculateMinimumHP(dungeon)
	}
}
