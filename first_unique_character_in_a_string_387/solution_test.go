package first_unique_character_in_a_string_387

import "testing"

func Test_firstUniqChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"leetcode", "leetcode", 0},
		{"loveleetcode", "loveleetcode", 2},
		{"aabb", "aabb", -1},
		{"single", "z", 0},
		{"empty", "", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
