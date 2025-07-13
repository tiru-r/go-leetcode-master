package flatten_nested_list_iterator_341

import "testing"

func TestNestedIterator(t *testing.T) {
	tests := []struct {
		name string
		in   *NestedInteger
		want []int
	}{
		{"single int", Int(1), []int{1}},
		{"empty list", List(), []int{}},
		{"flat list", List(Int(1), Int(2), Int(3)), []int{1, 2, 3}},
		{"nested once", List(List(Int(1), Int(2)), Int(3)), []int{1, 2, 3}},
		{"deep nesting", List(List(List(Int(1))), Int(2)), []int{1, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			it := NewNestedIterator([]*NestedInteger{tc.in})
			var got []int
			for it.HasNext() {
				got = append(got, it.Next())
			}
			if !equal(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func BenchmarkNext(b *testing.B) {
	n := Int(1)
	for range 100 {
		n = List(n)
	}
	it := NewNestedIterator([]*NestedInteger{n})
	it.HasNext() // Prepare the iterator
	b.ResetTimer()
	for range b.N {
		if it.HasNext() {
			_ = it.Next()
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
