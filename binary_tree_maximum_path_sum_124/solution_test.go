package binary_tree_maximum_path_sum_124

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Basic positive tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 6, // Path: 2 -> 1 -> 3
		},
		{
			name: "Test Case 2: Tree with a negative child",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 4, // Path: 1 -> 3 (as -2 is ignored due to negative contribution)
		},
		{
			name: "Test Case 3: Larger tree with negative root and positive path in subtree",
			args: args{
				root: &TreeNode{
					Val: -10,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 42, // Path: 15 -> 20 -> 7 (doesn't involve -10)
		},
		{
			name: "Test Case 4: Single node (positive)",
			args: args{
				root: &TreeNode{Val: 5},
			},
			want: 5,
		},
		{
			name: "Test Case 5: Single node (negative)",
			args: args{
				root: &TreeNode{Val: -3},
			},
			want: -3, // The path must contain at least one node, so the node itself is the max path.
		},
		{
			name: "Test Case 6: All negative nodes, max is the least negative",
			args: args{
				root: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: -3,
					},
				},
			},
			want: -1, // Path: -1 (or -2 if it was the only node, etc. Here -1 is the least negative)
		},
		{
			name: "Test Case 7: All negative nodes, deeper",
			args: args{
				root: &TreeNode{
					Val: -10,
					Left: &TreeNode{
						Val: -5,
						Left: &TreeNode{
							Val: -1,
						},
					},
					Right: &TreeNode{
						Val: -20,
					},
				},
			},
			want: -1, // Path: -1 (the maximum value among all nodes)
		},
		{
			name: "Test Case 8: Left skewed tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: 10, // Path: 1 -> 2 -> 3 -> 4
		},
		{
			name: "Test Case 9: Right skewed tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
							Right: &TreeNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: 10, // Path: 1 -> 2 -> 3 -> 4
		},
		{
			name: "Test Case 10: Path only in left subtree",
			args: args{
				root: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: -2,
					},
				},
			},
			want: 11, // Path: 5 -> 6 (doesn't involve -1 or -2)
		},
		{
			name: "Test Case 11: Path only in right subtree",
			args: args{
				root: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -2,
					},
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: 11, // Path: 5 -> 6 (doesn't involve -1 or -2)
		},
		{
			name: "Test Case 12: Tree with zeros",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			want: 0, // Path: 0 -> 0 -> 0 (or any single 0)
		},
		{
			name: "Test Case 13: Mixed positive and negative, complex",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
			},
			want: 48, // Path: 7 -> 11 -> 4 -> 5 -> 8 -> 13
		},
		{
			name: "Test Case 14: Negative values leading to positive path",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: -1,
					},
				},
			},
			want: 2, // Path: 2 (as -1 is ignored)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxPathSum(tt.args.root))
		})
	}
}
