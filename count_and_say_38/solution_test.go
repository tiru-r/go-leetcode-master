package count_and_say_38

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"n=1 returns initial term", 1, "1"},
		{"n=2 returns one 1", 2, "11"},
		{"n=3 returns two 1s", 3, "21"},
		{"n=4 returns one 2 one 1", 4, "1211"},
		{"n=5 returns one 1 one 2 two 1s", 5, "111221"},
		{"n=6 returns complex sequence", 6, "312211"},
		{"n=0 returns empty string", 0, ""},
		{"n=-1 returns empty string", -1, ""},
		{"n=31 exceeds constraint returns empty string", 31, ""},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countAndSay(tt.n))
		})
	}
}
