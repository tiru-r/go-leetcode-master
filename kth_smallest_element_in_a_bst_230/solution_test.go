package kth_smallest_element_in_a_bst_230

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "k=1 in small tree",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{Val: 4},
				},
				k: 1,
			},
			want: 1,
		},
		{
			name: "k=3 in five-node tree",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 1},
						},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 6},
				},
				k: 3,
			},
			want: 3,
		},
		{
			name: "k=4 in five-node tree",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:  2,
							Left: &TreeNode{Val: 1},
						},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 6},
				},
				k: 4,
			},
			want: 4,
		},
		{
			name: "single node, k=1",
			args: args{
				root: &TreeNode{Val: 42},
				k:    1,
			},
			want: 42,
		},
		{
			name: "right-skewed tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				k: 3,
			},
			want: 3,
		},
		{
			name: "k equals size",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, kthSmallest(tt.args.root, tt.args.k))
		})
	}
}
