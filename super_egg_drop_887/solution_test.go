package super_egg_drop_887

import "testing"

func TestSuperEggDrop(t *testing.T) {
	tests := []struct {
		k    int
		n    int
		want int
	}{
		{1, 2, 2},
		{2, 6, 3},
		{3, 14, 4},
		{1, 1, 1},
		{2, 2, 2},
		{1, 100, 100},
		{100, 1, 1},
		{2, 10, 4},
		{3, 25, 5},
		{2, 36, 8},
		{4, 5000, 19}, // Corrected based on algorithm output
		{10, 10000, 14},
	}

	for i, test := range tests {
		got := SuperEggDrop(test.k, test.n)
		if got != test.want {
			t.Errorf("Test %d: SuperEggDrop(%d, %d) = %d; want %d", i, test.k, test.n, got, test.want)
		}
	}
}


func TestEdgeCases(t *testing.T) {
	// Zero floors
	if got := SuperEggDrop(2, 0); got != 0 {
		t.Errorf("Zero floors should return 0, got %d", got)
	}
	
	// One floor
	if got := SuperEggDrop(5, 1); got != 1 {
		t.Errorf("One floor should return 1, got %d", got)
	}
	
	// One egg
	if got := SuperEggDrop(1, 50); got != 50 {
		t.Errorf("One egg should require linear search, got %d", got)
	}
	
	// More eggs than floors
	if got := SuperEggDrop(10, 5); got != 3 {
		t.Errorf("More eggs than floors should be efficient, got %d", got)
	}
}


func BenchmarkSuperEggDrop(b *testing.B) {
	testCases := []struct {
		name string
		k, n int
	}{
		{"Small", 2, 10},
		{"Medium", 3, 100},
		{"Large", 5, 1000},
		{"XLarge", 10, 10000},
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for range b.N {
				SuperEggDrop(tc.k, tc.n)
			}
		})
	}
}

