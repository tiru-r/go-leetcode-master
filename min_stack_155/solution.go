package min_stack_155

type MinStack struct {
	stack []int
	mins  []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack: make([]int, 0, 16),
		mins:  make([]int, 0, 16),
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.mins) == 0 || val <= ms.mins[len(ms.mins)-1] {
		ms.mins = append(ms.mins, val)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if top == ms.mins[len(ms.mins)-1] {
		ms.mins = ms.mins[:len(ms.mins)-1]
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		panic("empty stack")
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.mins) == 0 {
		panic("empty stack")
	}
	return ms.mins[len(ms.mins)-1]
}
