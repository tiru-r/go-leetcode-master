package happy_number_202

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		// 1. LeetCode official
		{"LC 19", 19, true},
		{"LC 2", 2, false},

		// 2. Edge & boundary numbers
		{"1", 1, true},
		{"0", 0, false},
		{"7", 7, true},
		{"10", 10, true},
		{"100", 100, true},
		{"1000", 1000, true},

		// 3. Known unhappy cycle starters
		{"4", 4, false},
		{"16", 16, false},
		{"20", 20, false},
		{"37", 37, false},
		{"58", 58, false},

		// 4. Happy numbers in different ranges
		{"145", 145, true},
		{"1000", 1000, true},
		{"1000000", 1000000, true},
		{"2147483647", 2147483647, false}, // max int32 boundary

		// 5. Large happy (performance smoke)
		{"999999999", 999999999, true},

		// 6. Numbers ending in 4 (cycle trigger)
		{"104", 104, false},
		{"204", 204, false},
		{"304", 304, false},

		// 7. Numbers with repeated digits
		{"111", 111, false},
		{"222", 222, false},
		{"333", 333, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isHappy(tc.n)
			if got != tc.want {
				t.Errorf("isHappy(%d) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
