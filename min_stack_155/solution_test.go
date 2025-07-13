package min_stack_155

import "testing"

func TestMinStack(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		vals     []int
		expected []int
	}{
		{
			name:     "basic",
			ops:      []string{"push", "push", "push", "getMin", "pop", "top", "getMin"},
			vals:     []int{-2, 0, -3, 0, 0, 0, 0},
			expected: []int{-1, -1, -1, -3, -1, 0, -2},
		},
		{
			name:     "single element",
			ops:      []string{"push", "getMin", "top"},
			vals:     []int{42, 0, 0},
			expected: []int{-1, 42, 42},
		},
		{
			name:     "overwrite min",
			ops:      []string{"push", "push", "pop", "getMin"},
			vals:     []int{1, 2, 0, 0},
			expected: []int{-1, -1, -1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMinStack()
			outIdx := 0
			for i, op := range tt.ops {
				switch op {
				case "push":
					ms.Push(tt.vals[i])
				case "pop":
					ms.Pop()
				case "top":
					got := ms.Top()
					if exp := tt.expected[outIdx]; exp != -1 && got != exp {
						t.Fatalf("top: want %d, got %d", exp, got)
					}
					outIdx++
				case "getMin":
					got := ms.GetMin()
					if exp := tt.expected[outIdx]; exp != -1 && got != exp {
						t.Fatalf("getMin: want %d, got %d", exp, got)
					}
					outIdx++
				}
			}
		})
	}
}

func BenchmarkMinStack(b *testing.B) {
	ms := NewMinStack()
	for range b.N {
		ms.Push(b.N)
		_ = ms.GetMin()
	}
}
