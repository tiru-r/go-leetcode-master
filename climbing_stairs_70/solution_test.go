package climbing_stairs_70

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "climbing stairs n=1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "climbing stairs n=2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "climbing stairs n=3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "climbing stairs n=44",
			args: args{
				n: 44,
			},
			want: 1134903170,
		},
		{
			name: "climbing stairs n=0", // Added test case for n=0
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "climbing stairs n=negative", // Added test case for negative n
			args: args{
				n: -5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, climbStairs(tt.args.n))
		})
	}
}
