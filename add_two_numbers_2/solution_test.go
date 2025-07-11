package add_two_numbers_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// createNinesList is a helper function to create a linked list of 'n' nines.
func createNinesList(n int) *ListNode {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 9
	}
	return createList(arr)
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example case: 342 + 465 = 807",
			args: args{
				l1: createList([]int{2, 4, 3}),
				l2: createList([]int{5, 6, 4}),
			},
			want: createList([]int{7, 0, 8}),
		},
		{
			name: "Different lengths: 1001 + 465 = 1466",
			args: args{
				l1: createList([]int{1, 0, 0, 1}),
				l2: createList([]int{5, 6, 4}),
			},
			want: createList([]int{6, 6, 4, 1}),
		},
		{
			name: "Carry propagation: 999 + 1 = 1000",
			args: args{
				l1: createList([]int{9, 9, 9}),
				l2: createList([]int{1}),
			},
			want: createList([]int{0, 0, 0, 1}),
		},
		{
			name: "Single-digit lists: 0 + 0 = 0",
			args: args{
				l1: createList([]int{0}),
				l2: createList([]int{0}),
			},
			want: createList([]int{0}),
		},
		{
			name: "Single-digit with carry: 9 + 9 = 18",
			args: args{
				l1: createList([]int{9}),
				l2: createList([]int{9}),
			},
			want: createList([]int{8, 1}),
		},
		{
			name: "Maximum digits with carry: 999...9 (100 nines) + 1 = 1000...0 (99 zeros, 1)",
			args: args{
				l1: createNinesList(100), // Using helper for clarity
				l2: createList([]int{1}),
			},
			// Expected output: 99 zeros followed by 1
			want: createList(append(make([]int, 100), 1)), // 100 zeros, then a 1
		},
		{
			name: "L1 is nil (robustness check, though not strictly per constraints)",
			args: args{
				l1: nil,
				l2: createList([]int{1, 2, 3}),
			},
			want: createList([]int{1, 2, 3}),
		},
		{
			name: "L2 is nil (robustness check, though not strictly per constraints)",
			args: args{
				l1: createList([]int{4, 5, 6}),
				l2: nil,
			},
			want: createList([]int{4, 5, 6}),
		},
		{
			name: "Both L1 and L2 are nil (robustness check, though not strictly per constraints)",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Compare the result linked list with the expected linked list using listToArray
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			assert.Equal(t, listToArray(tt.want), listToArray(got), "Linked lists should match")
		})
	}
}
