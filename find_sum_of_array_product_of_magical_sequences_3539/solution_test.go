package find_sum_of_array_product_of_magical_sequences_3539

import "testing"

func TestFindSumOfArrayProductOfMagicalSequences(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1, 2, 3},
			k:    1,
			want: 6, // 1 + 2 + 3
		},
		{
			nums: []int{1, 2, 3},
			k:    2,
			want: 11, // 1*2 + 1*3 + 2*3 = 2 + 3 + 6 = 11
		},
		{
			nums: []int{1, 2, 3},
			k:    3,
			want: 6, // 1*2*3 = 6
		},
		{
			nums: []int{2, 1, 3},
			k:    2,
			want: 11, // Same as above (order doesn't matter for magical sequences)
		},
		{
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			nums: []int{5, 3, 1, 4},
			k:    2,
			want: 59, // 1*3 + 1*4 + 1*5 + 3*4 + 3*5 + 4*5 = 3+4+5+12+15+20 = 59
		},
		{
			nums: []int{},
			k:    1,
			want: 0,
		},
		{
			nums: []int{1, 2, 3},
			k:    0,
			want: 0,
		},
		{
			nums: []int{10, 20, 30},
			k:    1,
			want: 60, // 10 + 20 + 30
		},
	}

	for i, test := range tests {
		got := FindSumOfArrayProductOfMagicalSequences(test.nums, test.k)
		if got != test.want {
			t.Errorf("Test %d: FindSumOfArrayProductOfMagicalSequences(%v, %d) = %d; want %d", 
				i, test.nums, test.k, got, test.want)
		}
	}
}

func TestIsValidExtension(t *testing.T) {
	tests := []struct {
		prevVal int
		num     int
		want    bool
	}{
		{1, 2, true},
		{2, 1, false},
		{5, 5, false},
		{3, 10, true},
	}

	for i, test := range tests {
		got := isValidExtension(test.prevVal, test.num)
		if got != test.want {
			t.Errorf("Test %d: isValidExtension(%d, %d) = %v; want %v", 
				i, test.prevVal, test.num, got, test.want)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Large k
	if got := FindSumOfArrayProductOfMagicalSequences([]int{1, 2}, 5); got != 0 {
		t.Errorf("Large k should return 0, got %d", got)
	}
	
	// Single element
	if got := FindSumOfArrayProductOfMagicalSequences([]int{42}, 1); got != 42 {
		t.Errorf("Single element should return element value, got %d", got)
	}
	
	// Decreasing sequence - we can still form increasing subsequences
	if got := FindSumOfArrayProductOfMagicalSequences([]int{5, 4, 3, 2, 1}, 2); got != 85 {
		t.Errorf("Decreasing sequence should have magical subsequences, got %d", got)
	}
}

func BenchmarkFindSumOfArrayProductOfMagicalSequences(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 3
	
	b.ResetTimer()
	for range b.N {
		FindSumOfArrayProductOfMagicalSequences(nums, k)
	}
}
