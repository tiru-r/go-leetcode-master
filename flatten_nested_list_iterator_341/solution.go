package flatten_nested_list_iterator_341

// ---------- minimal NestedInteger stub so this file compiles ----------
type NestedInteger struct {
	value any
	isInt bool
}

// constructors (not part of LeetCode API, but needed for tests)
func Int(v int) *NestedInteger { return &NestedInteger{value: v, isInt: true} }
func List(list ...*NestedInteger) *NestedInteger {
	return &NestedInteger{value: list, isInt: false}
}

// LeetCode-specified API
func (ni NestedInteger) IsInteger() bool { return ni.isInt }
func (ni NestedInteger) GetInteger() int { return ni.value.(int) }
func (ni NestedInteger) GetList() []*NestedInteger {
	return ni.value.([]*NestedInteger)
}

// ---------- iterator implementation ----------
type NestedIterator struct {
	stk []state
}

type state struct {
	idx  int
	list []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stk: []state{{idx: 0, list: nestedList}},
	}
}

func (it *NestedIterator) Next() int {
	// HasNext guarantees the top is an integer
	top := &it.stk[len(it.stk)-1]
	val := top.list[top.idx].GetInteger()
	top.idx++
	it.shrink()
	return val
}

func (it *NestedIterator) HasNext() bool {
	for len(it.stk) > 0 {
		top := &it.stk[len(it.stk)-1]
		if top.idx == len(top.list) { // finished current list
			it.stk = it.stk[:len(it.stk)-1]
			continue
		}
		node := top.list[top.idx]
		if node.IsInteger() {
			return true
		}
		// descend into nested list
		top.idx++
		it.stk = append(it.stk, state{idx: 0, list: node.GetList()})
	}
	return false
}

// remove exhausted frames after advancing
func (it *NestedIterator) shrink() {
	for len(it.stk) > 0 {
		top := &it.stk[len(it.stk)-1]
		if top.idx < len(top.list) {
			return
		}
		it.stk = it.stk[:len(it.stk)-1]
	}
}
