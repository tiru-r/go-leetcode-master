package flatten_nested_list_iterator_341

// ---------- NestedInteger stub (LeetCode provided) ----------
type NestedInteger struct {
	value any
	isInt bool
}

func Int(v int) *NestedInteger { return &NestedInteger{v, true} }
func List(el ...*NestedInteger) *NestedInteger {
	return &NestedInteger{el, false}
}
func (ni NestedInteger) IsInteger() bool           { return ni.isInt }
func (ni NestedInteger) GetInteger() int           { return ni.value.(int) }
func (ni NestedInteger) GetList() []*NestedInteger { return ni.value.([]*NestedInteger) }

// ---------- Iterator ----------
type NestedIterator struct {
	stk [][]*NestedInteger // stack of lists
	pos []int              // current index inside each list
}

func NewNestedIterator(list []*NestedInteger) *NestedIterator {
	// Pre-allocate capacity for typical nesting depth
	stk := make([][]*NestedInteger, 1, 8)
	pos := make([]int, 1, 8)
	stk[0] = list
	pos[0] = 0
	return &NestedIterator{
		stk: stk,
		pos: pos,
	}
}

// HasNext lazily moves to the next integer (O(1) amortised).
func (it *NestedIterator) HasNext() bool {
	for len(it.stk) > 0 {
		// Cache frequently used values
		top := len(it.stk) - 1
		curList := it.stk[top]
		curIdx := it.pos[top]

		if curIdx == len(curList) { // list exhausted → pop
			it.stk = it.stk[:top]
			it.pos = it.pos[:top]
			if top > 0 {
				it.pos[top-1]++
			}
			continue
		}

		item := curList[curIdx]
		if item.IsInteger() {
			return true
		}

		// Dive into sub-list
		it.stk = append(it.stk, item.GetList())
		it.pos = append(it.pos, 0)
	}
	return false
}

// Next returns the next integer.  Must be preceded by HasNext().
func (it *NestedIterator) Next() int {
	top := len(it.stk) - 1
	list := it.stk[top]
	idx := it.pos[top]
	val := list[idx].GetInteger()
	it.pos[top]++
	return val
}
