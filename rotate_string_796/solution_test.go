package rotate_string_796

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateString(t *testing.T) {
	type args struct {
		s string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"basic rotation", args{"abcde", "cdeab"}, true},
		{"no rotation", args{"abcde", "abced"}, false},
		{"same string", args{"abcde", "abcde"}, true},
		{"empty strings", args{"", ""}, true},
		{"different lengths", args{"abc", "abcd"}, false},
		{"single char", args{"a", "a"}, true},
		{"complex case", args{"bbbacddceeb", "ceebbbbacdd"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rotateString(tt.args.s, tt.args.goal))
		})
	}
}
