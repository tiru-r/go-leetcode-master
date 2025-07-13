package flatten_nested_list_iterator_341

type NestedInteger struct {
	value any
	isInt bool
}

func Int(v int) *NestedInteger { return &NestedInteger{value: v, isInt: true} }
func List(el ...*NestedInteger) *NestedInteger {
	return &NestedInteger{value: el, isInt: false}
}

func (ni NestedInteger) IsInteger() bool           { return ni.isInt }
func (ni NestedInteger) GetInteger() int           { return ni.value.(int) }
func (ni NestedInteger) GetList() []*NestedInteger { return ni.value.([]*NestedInteger) }

type iteratorState struct {
	index int
	items []*NestedInteger
}

type NestedIterator struct {
	stack []iteratorState
}

func NewNestedIterator(list []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: []iteratorState{{index: 0, items: list}}}
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		top := &it.stack[len(it.stack)-1]
		if top.index == len(top.items) {
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}
		current := top.items[top.index]
		if current.IsInteger() {
			return true
		}
		top.index++
		it.stack = append(it.stack, iteratorState{index: 0, items: current.GetList()})
	}
	return false
}

func (it *NestedIterator) Next() int {
	top := &it.stack[len(it.stack)-1]
	val := top.items[top.index].GetInteger()
	top.index++
	return val
}
