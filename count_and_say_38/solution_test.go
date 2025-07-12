package count_and_say_38

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "n=1 returns initial term",
			args: args{n: 1},
			want: "1",
		},
		{
			name: "n=2 returns one 1",
			args: args{n: 2},
			want: "11",
		},
		{
			name: "n=3 returns two 1s",
			args: args{n: 3},
			want: "21",
		},
		{
			name: "n=4 returns one 2 one 1",
			args: args{n: 4},
			want: "1211",
		},
		{
			name: "n=5 returns one 1 one 2 two 1s",
			args: args{n: 5},
			want: "111221",
		},
		{
			name: "n=6 returns complex sequence",
			args: args{n: 6},
			want: "312211",
		},
		{
			name: "n=0 returns empty string",
			args: args{n: 0},
			want: "",
		},
		{
			name: "n=-1 returns empty string",
			args: args{n: -1},
			want: "",
		},
		{
			name: "n=31 exceeds constraint returns empty string",
			args: args{n: 31},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countAndSay(tt.args.n))
		})
	}
}
