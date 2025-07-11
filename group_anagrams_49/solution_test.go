package group_anagrams_49

import (
	"cmp"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "group anagrams",
			args: args{
				strs: []string{
					"eat", "tea", "tan", "ate", "nat", "bat",
				},
			},
			want: [][]string{
				{
					"bat",
				},
				{
					"eat", "tea", "ate",
				},
				{
					"tan", "nat",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := groupAnagrams(tt.args.strs)
			// Use slices.SortFunc with cmp.Compare for modern slice sorting
			slices.SortFunc(g, func(a, b []string) int {
				return cmp.Compare(a[0], b[0])
			})
			assert.Equal(t, tt.want, g)
		})
	}
}
