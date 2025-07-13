package self_crossing_335

import "testing"

func TestIsSelfCrossing(t *testing.T) {
	tests := []struct {
		distance []int
		want     bool
	}{
		{
			distance: []int{2, 1, 1, 2},
			want:     true,
		},
		{
			distance: []int{1, 2, 3, 4},
			want:     false,
		},
		{
			distance: []int{1, 1, 2, 1, 1},
			want:     true,
		},
		{
			distance: []int{3, 3, 4, 2, 2},
			want:     false,
		},
		{
			distance: []int{1, 1, 1, 1},
			want:     true,
		},
		{
			distance: []int{1, 2, 2, 1},
			want:     false,
		},
		{
			distance: []int{2, 1, 2, 1},
			want:     true, // Corrected: this path does cross
		},
		{
			distance: []int{1, 1, 2, 2, 1, 1},
			want:     true,
		},
		{
			distance: []int{3, 3, 3, 2, 1, 1},
			want:     false,
		},
		{
			distance: []int{1, 1, 2, 2, 3, 3},
			want:     false,
		},
		{
			distance: []int{},
			want:     false,
		},
		{
			distance: []int{1},
			want:     false,
		},
		{
			distance: []int{1, 2},
			want:     false,
		},
		{
			distance: []int{1, 2, 3},
			want:     false,
		},
	}

	for i, test := range tests {
		got := IsSelfCrossing(test.distance)
		if got != test.want {
			t.Errorf("Test %d: IsSelfCrossing(%v) = %v; want %v", i, test.distance, got, test.want)
		}
	}
}

func TestSelfCrossingPatterns(t *testing.T) {
	// Test Case 1: Fourth line crosses first line
	if !IsSelfCrossing([]int{2, 1, 1, 2}) {
		t.Error("Should detect fourth line crossing first line")
	}
	
	// Test Case 2: Fifth line crosses first line
	if !IsSelfCrossing([]int{1, 1, 2, 1, 1}) {
		t.Error("Should detect fifth line crossing first line")
	}
	
	// Test Case 3: Sixth line crosses first line
	if !IsSelfCrossing([]int{1, 1, 2, 2, 1, 1}) {
		t.Error("Should detect sixth line crossing first line")
	}
	
	// No crossing - spiral outward
	if IsSelfCrossing([]int{1, 2, 3, 4}) {
		t.Error("Outward spiral should not cross")
	}
	
	// No crossing - spiral inward
	if IsSelfCrossing([]int{4, 3, 2, 1}) {
		t.Error("Inward spiral should not cross")
	}
}

func BenchmarkIsSelfCrossing(b *testing.B) {
	distance := []int{1, 1, 2, 2, 1, 1, 3, 3, 2, 2}
	
	b.ResetTimer()
	for range b.N {
		IsSelfCrossing(distance)
	}
}
