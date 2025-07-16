package roman_to_integer_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single I", args{"I"}, 1},
		{"single V", args{"V"}, 5},
		{"single X", args{"X"}, 10},
		{"single L", args{"L"}, 50},
		{"single C", args{"C"}, 100},
		{"single D", args{"D"}, 500},
		{"single M", args{"M"}, 1000},
		{"III", args{"III"}, 3},
		{"IV", args{"IV"}, 4},
		{"IX", args{"IX"}, 9},
		{"XL", args{"XL"}, 40},
		{"CD", args{"CD"}, 400},
		{"CM", args{"CM"}, 900},
		{"LVIII", args{"LVIII"}, 58},
		{"MCMXCIV", args{"MCMXCIV"}, 1994},
		{"MCDXLIV", args{"MCDXLIV"}, 1444},
		{"MMMDCCCXC", args{"MMMDCCCXC"}, 3890},
		{"MMMDCCCLXXXVIII", args{"MMMDCCCLXXXVIII"}, 3888},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, romanToInt(tt.args.s))
		})
	}
}
