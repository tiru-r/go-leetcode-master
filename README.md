# Go LeetCode Solutions

High-performance Go solutions to LeetCode problems with optimal algorithms and Go best practices.

## Features

- **Optimal Algorithms**: Best time/space complexity solutions
- **Go Best Practices**: Modern Go 1.24+ patterns and idioms
- **Performance Focused**: Zero-allocation patterns and optimizations
- **Complete Testing**: 100% test coverage with benchmarks
- **Anti-patterns**: Common mistakes and how to avoid them

## Quick Start

```bash
# Clone and test a specific problem
git clone https://github.com/yourusername/go-leetcode-master.git
cd go-leetcode-master/two_sum_1
go test -v -bench=. -benchmem
```

## Problem Categories

### Arrays & Hashing
- ✅ [Two Sum](two_sum_1/) - O(n) HashMap lookup
- ✅ [Contains Duplicate](contains_duplicate_217/) - O(n) HashSet
- ✅ [Valid Anagram](valid_anagram_242/) - O(n) frequency counting
- ✅ [Group Anagrams](group_anagrams_49/) - O(n×k log k) sorting
- ✅ [Product of Array Except Self](product_of_array_except_self_238/) - O(n) prefix products

### Two Pointers
- ✅ [Valid Palindrome](valid_palindrome_125/) - O(n)
- ✅ [3Sum](three_sum_15/) - O(n²)
- ✅ [Container With Most Water](container_with_most_water_11/) - O(n)
- ✅ [Trapping Rain Water](trapping_rain_water_42/) - O(n)

### Sliding Window
- ✅ [Longest Substring Without Repeating Characters](longest_substring_without_repeating_characters_3/) - O(n)
- ✅ [Best Time to Buy and Sell Stock](best_time_to_buy_and_sell_stock_121/) - O(n)
- ✅ [Minimum Window Substring](minimum_window_substring_76/) - O(n)

### Trees
- ✅ [Invert Binary Tree](invert_binary_tree_226/) - O(n)
- ✅ [Maximum Depth of Binary Tree](maximum_depth_of_binary_tree_104/) - O(n)
- ✅ [Validate Binary Search Tree](validate_binary_search_tree_98/) - O(n)
- ✅ [Binary Tree Level Order Traversal](binary_tree_level_order_traversal_102/) - O(n)

### Graphs
- ✅ [Number of Islands](number_of_islands_200/) - O(n×m) DFS
- ✅ [Course Schedule](course_schedule_207/) - O(V+E) topological sort
- ✅ [Pacific Atlantic Water Flow](pacific_atlantic_water_flow_417/) - O(n×m) BFS

### Dynamic Programming
- ✅ [Climbing Stairs](climbing_stairs_70/) - O(n)
- ✅ [House Robber](house_robber_198/) - O(n) space O(1)
- ✅ [Coin Change](coin_change_322/) - O(n×amount)
- ✅ [Longest Increasing Subsequence](longest_increasing_subsequence_300/) - O(n log n)

### Linked Lists
- ✅ [Reverse Linked List](reverse_linked_list_206/) - O(n)
- ✅ [Merge Two Sorted Lists](merge_two_sorted_lists_21/) - O(n)
- ✅ [Linked List Cycle](linked_list_cycle_141/) - O(n) Floyd's algorithm

## Go Best Practices

### Memory Optimization
```go
// ✅ Pre-allocate with known capacity
result := make([]int, 0, expectedSize)

// ✅ Use struct{} for sets (zero memory)
visited := make(map[string]struct{})

// ✅ Reuse slices to avoid allocations
buffer = buffer[:0]  // Reset length, keep capacity
```

### Performance Patterns
```go
// ✅ Array-based data structures for cache locality
type TrieNode struct {
    children [26]*TrieNode  // Array > map for performance
    isEnd    bool
}

// ✅ Iterative over recursive to avoid stack overflow
func inorderTraversal(root *TreeNode) []int {
    // Iterative implementation
}
```

### Anti-patterns to Avoid
```go
// ❌ Allocation in loops
for i := 0; i < n; i++ {
    temp := make([]int, 10)  // Expensive!
}

// ❌ String concatenation in loops
var result string
for _, word := range words {
    result += word  // O(n²) complexity
}

// ❌ Using interface{} unnecessarily
func process(data []interface{}) {  // Type assertions hurt performance
}
```

## Algorithm Patterns

### Sliding Window Template
```go
func slidingWindow(s string) int {
    left, right := 0, 0
    window := make(map[byte]int)
    
    for right < len(s) {
        // Expand window
        window[s[right]]++
        right++
        
        // Contract window when invalid
        for /* window invalid */ {
            window[s[left]]--
            left++
        }
        
        // Update answer
    }
    return answer
}
```

### Two Pointers Template
```go
func twoPointers(nums []int) {
    left, right := 0, len(nums)-1
    
    for left < right {
        if /* condition */ {
            left++
        } else {
            right--
        }
    }
}
```

### DFS Template
```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }
    
    // Process current node
    
    dfs(node.Left)
    dfs(node.Right)
}
```

## Data Structures

### Trie Implementation
```go
type Trie struct {
    children [26]*Trie
    isEnd    bool
}

func (t *Trie) Insert(word string) {
    curr := t
    for _, ch := range word {
        idx := ch - 'a'
        if curr.children[idx] == nil {
            curr.children[idx] = &Trie{}
        }
        curr = curr.children[idx]
    }
    curr.isEnd = true
}
```

### Union-Find (Disjoint Set)
```go
type UnionFind struct {
    parent []int
    rank   []int
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])  // Path compression
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX != rootY {
        // Union by rank
        if uf.rank[rootX] < uf.rank[rootY] {
            uf.parent[rootX] = rootY
        } else if uf.rank[rootX] > uf.rank[rootY] {
            uf.parent[rootY] = rootX
        } else {
            uf.parent[rootY] = rootX
            uf.rank[rootX]++
        }
    }
}
```

## Testing & Benchmarking

### Benchmark Example
```go
func BenchmarkTwoSum(b *testing.B) {
    nums := []int{2, 7, 11, 15}
    target := 9
    
    b.ResetTimer()
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        result := twoSum(nums, target)
        _ = result  // Prevent optimization
    }
}
```

### Running Tests
```bash
# Run tests with coverage
go test -v -cover

# Run benchmarks with memory stats
go test -bench=. -benchmem

# Profile CPU usage
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

## Time & Space Complexity

### Big O Reference
- **O(1)**: Hash table lookup, array access
- **O(log n)**: Binary search, balanced tree operations
- **O(n)**: Linear search, tree traversal
- **O(n log n)**: Efficient sorting, divide & conquer
- **O(n²)**: Nested loops, bubble sort
- **O(2ⁿ)**: Recursive fibonacci, subset generation

### Space Optimization Techniques
- **In-place algorithms**: Modify input array directly
- **Rolling arrays**: Use O(k) space instead of O(n×k) for DP
- **Bit manipulation**: Use integers to represent sets
- **Two pointers**: O(1) space for array problems

## Performance Tips

### Memory Management
- Pre-allocate slices with known capacity
- Use object pooling for frequent allocations
- Avoid creating unnecessary temporary objects
- Use value receivers when possible

### CPU Optimization
- Prefer arrays over maps for small collections
- Use iterative solutions over recursive when possible
- Minimize function call overhead in hot paths
- Consider bit manipulation for set operations

### Profiling Tools
```bash
# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# CPU profiling  
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Trace analysis
go test -trace=trace.out -bench=.
go tool trace trace.out
```

## Contributing

1. Follow Go best practices and style guide
2. Include comprehensive tests with edge cases
3. Add benchmarks for performance-critical code
4. Document time/space complexity
5. Avoid common anti-patterns

## License

MIT License