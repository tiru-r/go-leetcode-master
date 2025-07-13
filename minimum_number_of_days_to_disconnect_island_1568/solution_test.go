package minimum_number_of_days_to_disconnect_island_1568

import "testing"

func TestMinDays(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{1, 1},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{1, 0, 1, 0},
			},
			want: 0,
		},
		{
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{1},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{0, 0, 0, 0},
			},
			want: 0,
		},
	}

	for i, test := range tests {
		// Create a copy to avoid modifying original
		gridCopy := make([][]int, len(test.grid))
		for j := range test.grid {
			gridCopy[j] = make([]int, len(test.grid[j]))
			copy(gridCopy[j], test.grid[j])
		}
		
		got := MinDays(gridCopy)
		if got != test.want {
			t.Errorf("Test %d: MinDays(%v) = %d; want %d", i, test.grid, got, test.want)
		}
	}
}

func TestCountIslands(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			grid: [][]int{
				{1, 1, 1},
				{0, 1, 0},
				{1, 1, 1},
			},
			want: 1,
		},
		{
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			want: 4,
		},
	}

	for i, test := range tests {
		got := countIslands(test.grid)
		if got != test.want {
			t.Errorf("Test %d: countIslands(%v) = %d; want %d", i, test.grid, got, test.want)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Single cell with land
	grid := [][]int{{1}}
	if got := MinDays(grid); got != 1 {
		t.Errorf("Single land cell should return 1, got %d", got)
	}
	
	// Single cell with water
	grid = [][]int{{0}}
	if got := MinDays(grid); got != 0 {
		t.Errorf("Single water cell should return 0, got %d", got)
	}
	
	// Already disconnected
	grid = [][]int{
		{1, 0, 1},
		{0, 0, 0},
	}
	if got := MinDays(grid); got != 0 {
		t.Errorf("Already disconnected should return 0, got %d", got)
	}
	
	// Articulation point (bridge)
	grid = [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	if got := MinDays(grid); got != 1 {
		t.Errorf("Articulation point should return 1, got %d", got)
	}
}

func BenchmarkMinDays(b *testing.B) {
	grid := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	
	b.ResetTimer()
	for range b.N {
		// Create copy for each iteration
		gridCopy := make([][]int, len(grid))
		for i := range grid {
			gridCopy[i] = make([]int, len(grid[i]))
			copy(gridCopy[i], grid[i])
		}
		MinDays(gridCopy)
	}
}
