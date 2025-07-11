package reverse_linked_list_206

import (
	"fmt"
	"slices"
	"testing"
)

// Test modern Go 1.24 features
func TestIterators(t *testing.T) {
	list := FromSlice([]int{1, 2, 3, 4, 5})
	
	// Test All() iterator
	var values []int
	for val := range list.All() {
		values = append(values, val)
	}
	expected := []int{1, 2, 3, 4, 5}
	if !slices.Equal(values, expected) {
		t.Errorf("All() iterator = %v, want %v", values, expected)
	}
	
	// Test AllNodes() iterator
	var nodeValues []int
	for node := range list.AllNodes() {
		nodeValues = append(nodeValues, node.Val)
	}
	if !slices.Equal(nodeValues, expected) {
		t.Errorf("AllNodes() iterator = %v, want %v", nodeValues, expected)
	}
}

func TestGenerics(t *testing.T) {
	// Test with int
	intList := FromSliceGeneric([]int{1, 2, 3})
	reversedInt := ReverseListGeneric(intList)
	expectedInt := []int{3, 2, 1}
	actualInt := ToSliceGeneric(reversedInt)
	if !slices.Equal(actualInt, expectedInt) {
		t.Errorf("Generic int reverse = %v, want %v", actualInt, expectedInt)
	}
	
	// Test with string
	stringList := FromSliceGeneric([]string{"a", "b", "c"})
	reversedString := ReverseListGeneric(stringList)
	expectedString := []string{"c", "b", "a"}
	actualString := ToSliceGeneric(reversedString)
	if !slices.Equal(actualString, expectedString) {
		t.Errorf("Generic string reverse = %v, want %v", actualString, expectedString)
	}
}

func TestFunctionalOperations(t *testing.T) {
	list := FromSlice([]int{1, 2, 3, 4, 5, 6})
	
	// Test Filter
	even := Filter(list, func(x int) bool { return x%2 == 0 })
	expectedEven := []int{2, 4, 6}
	actualEven := ToSlice(even)
	if !slices.Equal(actualEven, expectedEven) {
		t.Errorf("Filter even = %v, want %v", actualEven, expectedEven)
	}
	
	// Test Map
	doubled := Map(list, func(x int) int { return x * 2 })
	expectedDoubled := []int{2, 4, 6, 8, 10, 12}
	actualDoubled := ToSlice(doubled)
	if !slices.Equal(actualDoubled, expectedDoubled) {
		t.Errorf("Map double = %v, want %v", actualDoubled, expectedDoubled)
	}
}

func TestAllReverseImplementations(t *testing.T) {
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
			// Test all implementations
			funcs := []struct {
				name string
				fn   func(*ListNode) *ListNode
			}{
				{"ReverseList", ReverseList},
				{"ReverseListRecursive", ReverseListRecursive},
				{"ReverseListUnsafe", ReverseListUnsafe},
				{"ReverseIteratively", ReverseIteratively},
			}

			for _, f := range funcs {
				input := FromSlice(tt.input)
				result := f.fn(input)
				actual := ToSlice(result)

				if !slices.Equal(actual, tt.expected) {
					t.Errorf("%s: got %v, want %v", f.name, actual, tt.expected)
				}
			}
		})
	}
}

// Modern Go 1.24 benchmarks
func BenchmarkReverseListModern(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := range size {
				vals[i] = i
			}
			
			b.ResetTimer()
			for range b.N {
				list := FromSlice(vals)
				ReverseList(list)
			}
		})
	}
}

func BenchmarkReverseListGeneric(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := range size {
				vals[i] = i
			}
			
			b.ResetTimer()
			for range b.N {
				list := FromSliceGeneric(vals)
				ReverseListGeneric(list)
			}
		})
	}
}

func BenchmarkReverseListUnsafe(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := range size {
				vals[i] = i
			}
			
			b.ResetTimer()
			for range b.N {
				list := FromSlice(vals)
				ReverseListUnsafe(list)
			}
		})
	}
}

func BenchmarkReverseIteratively(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			vals := make([]int, size)
			for i := range size {
				vals[i] = i
			}
			
			b.ResetTimer()
			for range b.N {
				list := FromSlice(vals)
				ReverseIteratively(list)
			}
		})
	}
}