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

## Data Structures & Algorithms

### Data Structures Used

**Arrays & Hash Tables**
- **Arrays**: Contiguous memory blocks providing O(1) random access by index. Used for problems requiring fast lookups and cache-friendly sequential access patterns.
- **Hash Maps**: Key-value pairs with O(1) average lookup time. Essential for problems like Two Sum where we need to find complements quickly.
- **Hash Sets**: Unique element storage using `map[T]struct{}` in Go for zero memory overhead. Perfect for duplicate detection and membership testing.

**Trees**
- **Binary Trees**: Hierarchical structure where each node has at most two children. Used for problems involving tree traversal, path finding, and hierarchical data representation.
- **Binary Search Trees (BST)**: Ordered binary trees where left subtree < root < right subtree. Provides O(log n) search, insert, and delete operations.
- **Trie (Prefix Tree)**: Tree-like structure for string storage where each node represents a character. Excellent for word search, autocomplete, and prefix matching problems.

**Linked Lists**
- **Singly Linked Lists**: Linear data structure where elements point to the next node. Used when you need dynamic size and don't require random access.
- **Floyd's Cycle Detection**: Two-pointer technique (tortoise and hare) to detect cycles in linked lists in O(n) time with O(1) space.

**Graphs**
- **Adjacency Lists**: Graph representation using arrays/slices of neighbors. Memory-efficient for sparse graphs and enables fast neighbor iteration.
- **Union-Find (Disjoint Set)**: Data structure for tracking connected components with path compression and union by rank for nearly O(1) operations.

**Stacks & Queues**
- **Stacks**: LIFO (Last In, First Out) structure perfect for problems involving matching parentheses, function calls, and backtracking.
- **Queues**: FIFO (First In, First Out) structure essential for BFS traversal and level-order processing.

### Algorithms Used

**Sorting & Searching**
- **Binary Search**: O(log n) algorithm for finding elements in sorted arrays. Works by repeatedly dividing search space in half.
- **Merge Sort**: Divide-and-conquer sorting algorithm with O(n log n) time complexity. Concepts used in problems requiring stable sorting or merging operations.

**Graph Algorithms**
- **DFS (Depth-First Search)**: Explores as far as possible along each branch before backtracking. Used for tree traversal, topological sorting, and finding connected components.
- **BFS (Breadth-First Search)**: Explores all neighbors at current depth before moving to next level. Perfect for shortest path in unweighted graphs and level-order traversal.
- **Topological Sort**: Ordering vertices in directed acyclic graph where each vertex comes before its dependencies. Used in course scheduling and dependency resolution.

**Dynamic Programming**
- **Memoization**: Top-down approach storing results of expensive function calls to avoid recomputation. Used in recursive problems with overlapping subproblems.
- **Bottom-up DP**: Building solution iteratively from smaller subproblems to larger ones. More space-efficient than memoization.
- **Space Optimization**: Techniques like rolling arrays to reduce space complexity from O(n²) to O(n) in 2D DP problems.

**Two Pointers**
- **Left/Right Pointers**: Two pointers moving from opposite ends of array. Used for problems like palindrome checking, two sum in sorted array, and container problems.
- **Fast/Slow Pointers**: Pointers moving at different speeds. Classic technique for cycle detection and finding middle elements.

**Sliding Window**
- **Variable-Size Window**: Window that expands and contracts based on conditions. Used for substring problems like longest substring without repeating characters.
- **Fixed-Size Window**: Window of constant size sliding through array. Used for problems involving subarrays of specific length.

**Backtracking**
- **Recursive Exploration**: Systematic way of exploring all possible solutions by making choices and undoing them. Used for generating permutations, combinations, and solving constraint satisfaction problems.
- **Pruning**: Optimization technique to avoid exploring branches that cannot lead to valid solutions.

**Greedy Algorithms**
- **Local Optimal Choices**: Making locally optimal choice at each step hoping to find global optimum. Used in interval scheduling, activity selection, and some shortest path problems.
- **Proof of Correctness**: Greedy algorithms require mathematical proof that local choices lead to global optimum.

**Bit Manipulation**
- **XOR Operations**: Bitwise exclusive OR useful for finding unique elements (x ^ x = 0, x ^ 0 = x). Used in problems involving pairs and duplicates.
- **Bit Masks**: Using integers to represent sets where each bit represents presence/absence of element. Enables efficient subset enumeration and dynamic programming on subsets.

### Time Complexity Analysis

**Common Complexities**
- **O(1)**: Constant time - hash table operations, array access
- **O(log n)**: Logarithmic time - binary search, balanced tree operations
- **O(n)**: Linear time - single pass through array, tree traversal
- **O(n log n)**: Linearithmic time - efficient sorting, divide-and-conquer
- **O(n²)**: Quadratic time - nested loops, some dynamic programming
- **O(2ⁿ)**: Exponential time - brute force subset generation, naive recursion

**Space Complexity Considerations**
- **In-place algorithms**: O(1) space by modifying input
- **Auxiliary space**: Additional space needed beyond input
- **Recursive space**: Stack space used by recursive calls

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