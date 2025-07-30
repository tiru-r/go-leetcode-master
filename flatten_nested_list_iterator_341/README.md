# 341. Flatten Nested List Iterator

## Problem Description

You are given a nested list of integers `nestedList`. Each element is either an integer or a list whose elements may also be integers or other lists. Implement an iterator to flatten it.

Implement the `NestedIterator` class:

- `NestedIterator(List<NestedInteger> nestedList)` Initializes the iterator with the nested list `nestedList`.
- `int next()` Returns the next integer in the nested list.
- `boolean hasNext()` Returns `true` if there are still some integers in the nested list and `false` otherwise.

Your code will be tested with the following pseudocode:

```
initialize iterator with nestedList
res = []
while iterator.hasNext()
    append iterator.next() to the end of res
return res
```

## Examples

**Example 1:**
```
Input: nestedList = [[1,1],2,[1,1]]
Output: [1,1,2,1,1]
Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,1,2,1,1].
```

**Example 2:**
```
Input: nestedList = [1,[4,[6]]]
Output: [1,4,6]
Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,4,6].
```

## Constraints

- `1 <= nestedList.length <= 500`
- The values of the integers in the nested list is in the range `[-10^6, 10^6]`.

## Approach

This solution uses a **stack-based approach** with lazy evaluation:

### Key Ideas:
1. **Stack of Lists**: Maintain a stack of nested lists being processed
2. **Position Tracking**: Track the current position in each list level
3. **Lazy Evaluation**: Only flatten when `hasNext()` is called, not during initialization
4. **Efficient Navigation**: Use stack operations to handle nested structures

### Algorithm:
1. **Initialization**: Push the root list onto the stack with position 0
2. **HasNext()**: 
   - Navigate through the stack to find the next integer
   - Pop exhausted lists from the stack
   - Push nested lists onto the stack when encountered
   - Return true if an integer is found
3. **Next()**: Return the current integer and advance the position

### Time Complexity:
- **HasNext()**: O(1) amortized - each element is processed exactly once
- **Next()**: O(1) - simple array access and position increment
- **Overall**: O(n) where n is the total number of integers

### Space Complexity:
- O(d) where d is the maximum nesting depth (for the stack)
- In the worst case, O(n) if the structure is deeply nested

## Solution Features

- **Memory Efficient**: Pre-allocates stack capacity to reduce allocations
- **Lazy Evaluation**: Only processes elements when needed
- **Clean Interface**: Follows the iterator pattern with clear separation of concerns
- **Optimal Performance**: O(1) amortized time complexity for both operations