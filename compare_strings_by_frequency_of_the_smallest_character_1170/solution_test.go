package compare_strings_by_frequency_of_the_smallest_character_1170

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSmallerByFrequency(t *testing.T) {
	type args struct {
		queries []string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1: Basic example",
			args: args{
				queries: []string{
					"cbd",
				},
				words: []string{
					"zaaaz",
				},
			},
			want: []int{
				1,
			},
		},
		{
			name: "Test Case 2: Multiple queries and words",
			args: args{
				queries: []string{
					"bbb", "cc",
				},
				words: []string{
					"a", "aa", "aaa", "aaaa",
				},
			},
			want: []int{
				1, 2,
			},
		},
		{
			name: "Test Case 3: No words with greater frequency",
			args: args{
				queries: []string{
					"g", "h",
				},
				words: []string{
					"a", "b", "c",
				},
			},
			want: []int{
				0, 0,
			},
		},
		{
			name: "Test Case 4: All words with greater frequency",
			args: args{
				queries: []string{
					"z",
				},
				words: []string{
					"aa", "bb", "cc", "dd",
				},
			},
			want: []int{
				4,
			},
		},
		{
			name: "Test Case 5: Empty queries or words",
			args: args{
				queries: []string{
					"a",
				},
				words: []string{},
			},
			want: []int{
				0,
			},
		},
		{
			name: "Test Case 6: Queries with same frequency as words",
			args: args{
				queries: []string{
					"a", "aa",
				},
				words: []string{
					"b", "bb", "bbb",
				},
			},
			want: []int{
				2, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numSmallerByFrequency(tt.args.queries, tt.args.words))
		})
	}
}
