package flatten_nested_list_iterator_341

// NestedInteger represents either an integer or a list of NestedIntegers
type NestedInteger struct {
	value any
	isInt bool
}

// Factory functions for creating NestedIntegers
func Int(v int) *NestedInteger { 
	return &NestedInteger{value: v, isInt: true} 
}

func List(el ...*NestedInteger) *NestedInteger {
	return &NestedInteger{value: el, isInt: false}
}

// Interface methods for NestedInteger
func (ni NestedInteger) IsInteger() bool           { return ni.isInt }
func (ni NestedInteger) GetInteger() int           { return ni.value.(int) }
func (ni NestedInteger) GetList() []*NestedInteger { return ni.value.([]*NestedInteger) }

// NestedIterator implements lazy flattening using stack-based approach
// Optimized: O(1) amortized Next/HasNext, O(depth) space
type NestedIterator struct {
	stack []frame
}

// frame represents iterator state for each nesting level
type frame struct {
	items []*NestedInteger
	index int
}

// NewNestedIterator creates iterator with optimized initialization
func NewNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	if len(nestedList) == 0 {
		return &NestedIterator{}
	}
	return &NestedIterator{
		stack: []frame{{items: nestedList, index: 0}},
	}
}

// HasNext ensures next integer is available with lazy evaluation
func (it *NestedIterator) HasNext() bool {
	// Flatten until we find an integer or exhaust all elements
	for len(it.stack) > 0 {
		top := &it.stack[len(it.stack)-1]
		
		// Pop completed frames
		if top.index >= len(top.items) {
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}
		
		current := top.items[top.index]
		
		// Found integer - ready to return
		if current.IsInteger() {
			return true
		}
		
		// Expand nested list by pushing new frame
		top.index++
		nestedItems := current.GetList()
		if len(nestedItems) > 0 {
			it.stack = append(it.stack, frame{
				items: nestedItems,
				index: 0,
			})
		}
	}
	
	return false
}

// Next returns the next integer (assumes HasNext() was called)
func (it *NestedIterator) Next() int {
	top := &it.stack[len(it.stack)-1]
	val := top.items[top.index].GetInteger()
	top.index++
	return val
}
