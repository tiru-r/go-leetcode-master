package super_egg_drop_887

import "testing"

// ---------- table-driven test cases ----------
var cases = []struct {
	name     string
	k, n     int
	expected int
}{
	{"LeetCode Example 1", 1, 2, 2},  // [^5^]
	{"LeetCode Example 2", 2, 6, 3},  // [^5^]
	{"LeetCode Example 3", 3, 14, 4}, // [^5^]
	{"Single floor", 100, 1, 1},
	{"Single egg, many floors", 1, 10000, 10000},
	{"Eggs >= floors", 100, 100, 7},
	{"k = 2, n = 12", 2, 12, 5},  // [^10^]
	{"k = 3, n = 25", 3, 25, 5},  // [^10^]
	{"k = 2, n = 36", 2, 36, 8},  // [^9^]
	{"k = 1, n = 36", 1, 36, 36}, // [^9^]
	{"k = 2, n = 10", 2, 10, 4},  // [^9^]
	{"Upper limit k", 100, 10000, 9},
	{"Upper limit n", 100, 10000, 9},
}

// ---------- test runner ----------
func TestSuperEggDrop(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := superEggDrop(tc.k, tc.n)
			if got != tc.expected {
				t.Errorf("k=%d n=%d: expected %d, got %d", tc.k, tc.n, tc.expected, got)
			}
		})
	}
}

// ---------- optional benchmarks ----------
func BenchmarkSuperEggDrop(b *testing.B) {
	bench := []struct{ k, n int }{
		{100, 10000},
		{3, 14},
		{2, 1000},
	}
	for _, bb := range bench {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				superEggDrop(bb.k, bb.n)
			}
		})
	}
}
