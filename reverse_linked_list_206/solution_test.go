package reverse_linked_list_206

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"duplicate elements", []int{1, 1, 2, 2}, []int{2, 2, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := FromSlice(tt.input)
			result := ReverseList(input)
			actual := ToSlice(result)

			if len(actual) != len(tt.expected) {
				t.Errorf("ReverseList() length = %v, want %v", len(actual), len(tt.expected))
				return
			}

			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("ReverseList() = %v, want %v", actual, tt.expected)
					break
				}
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := FromSlice(tt.input)
			result := ReverseListRecursive(input)
			actual := ToSlice(result)

			if len(actual) != len(tt.expected) {
				t.Errorf("ReverseListRecursive() length = %v, want %v", len(actual), len(tt.expected))
				return
			}

			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("ReverseListRecursive() = %v, want %v", actual, tt.expected)
					break
				}
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected bool
	}{
		{"both empty", []int{}, []int{}, true},
		{"same lists", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"different lists", []int{1, 2, 3}, []int{3, 2, 1}, false},
		{"different lengths", []int{1, 2}, []int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listA := FromSlice(tt.a)
			listB := FromSlice(tt.b)
			result := Equal(listA, listB)

			if result != tt.expected {
				t.Errorf("Equal() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"empty list", []int{}, true},
		{"single element", []int{1}, true},
		{"sorted ascending", []int{1, 2, 3, 4, 5}, true},
		{"sorted with duplicates", []int{1, 1, 2, 3}, true},
		{"not sorted", []int{1, 3, 2, 4}, false},
		{"descending", []int{5, 4, 3, 2, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := FromSlice(tt.input)
			result := IsSorted(list)

			if result != tt.expected {
				t.Errorf("IsSorted() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkReverseList(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := 0; i < size; i++ {
				vals[i] = i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				list := FromSlice(vals)
				ReverseList(list)
			}
		})
	}
}

func BenchmarkReverseListRecursive(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := 0; i < size; i++ {
				vals[i] = i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				list := FromSlice(vals)
				ReverseListRecursive(list)
			}
		})
	}
}

func BenchmarkReverseListWithStdLib(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := 0; i < size; i++ {
				vals[i] = i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				list := FromSlice(vals)
				ReverseListWithStdLib(list)
			}
		})
	}
}
