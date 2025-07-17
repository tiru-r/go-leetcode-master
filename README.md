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

*Arrays*
- **Structure**: Contiguous memory blocks where elements are stored sequentially
- **Access Pattern**: O(1) random access by index due to arithmetic address calculation
- **Cache Performance**: Excellent spatial locality - accessing adjacent elements is very fast
- **Go Implementation**: `[]T` slices with underlying array, length, and capacity
- **Use Cases**: 
  - Matrix problems (2D arrays)
  - Dynamic programming tables
  - Sliding window problems
  - Two pointers technique
- **Memory Layout**: Elements stored consecutively in memory, enabling CPU prefetching
- **Performance**: Best for sequential access patterns and random indexing

*Hash Maps*
- **Structure**: Key-value pairs using hash function to map keys to bucket indices
- **Collision Resolution**: Go uses open addressing with Robin Hood hashing
- **Time Complexity**: O(1) average case for insert/lookup/delete, O(n) worst case
- **Go Implementation**: `map[K]V` with built-in hash functions for basic types
- **Load Factor**: Go automatically resizes when load factor exceeds threshold
- **Use Cases**:
  - Two Sum problem: store complement values
  - Anagram detection: character frequency counting
  - Caching results in dynamic programming
  - Group anagrams: use sorted string as key
- **Memory Overhead**: ~24 bytes per entry plus key/value sizes
- **Performance**: Excellent for lookups, poor for ordered traversal

*Hash Sets*
- **Structure**: Hash table storing only keys, no values
- **Go Implementation**: `map[T]struct{}` using empty struct for zero memory overhead
- **Memory Optimization**: `struct{}` takes 0 bytes, only key storage cost
- **Use Cases**:
  - Duplicate detection (Contains Duplicate problem)
  - Visited node tracking in graph traversal
  - Membership testing in O(1) time
- **Alternative**: `map[T]bool` wastes 1 byte per entry for boolean values
- **Performance**: Same as hash maps but with minimal memory usage

**Trees**

*Binary Trees*
- **Structure**: Hierarchical data structure where each node has at most two children (left, right)
- **Node Definition**: Typically contains data, left pointer, right pointer
- **Properties**: 
  - Height: Maximum path length from root to leaf
  - Depth: Path length from root to specific node
  - Balance: Height difference between left and right subtrees
- **Traversal Methods**:
  - Inorder: Left → Root → Right (gives sorted order for BST)
  - Preorder: Root → Left → Right (useful for copying tree)
  - Postorder: Left → Right → Root (useful for deletion)
  - Level-order: BFS traversal level by level
- **Use Cases**:
  - Expression parsing and evaluation
  - Decision trees in algorithms
  - File system directory structure
  - Hierarchical data representation
- **Memory**: Each node typically 24 bytes (8 bytes data + 16 bytes pointers)

*Binary Search Trees (BST)*
- **Structure**: Binary tree with ordering property: left < root < right
- **Search**: O(log n) average, O(n) worst case (degenerate tree)
- **Insertion**: Maintain BST property while inserting new nodes
- **Deletion**: Three cases: no children, one child, two children
- **Self-Balancing**: AVL, Red-Black trees maintain O(log n) height
- **Use Cases**:
  - Maintain sorted data with efficient insertion/deletion
  - Range queries (find elements between two values)
  - Finding kth smallest/largest element
  - Validate BST problem
- **Degeneration**: Can become linked list if data is sorted
- **In-order Traversal**: Always produces sorted sequence

*Trie (Prefix Tree)*
- **Structure**: Tree where each node represents a character, paths represent strings
- **Node Components**: 
  - Children array/map (usually 26 for lowercase letters)
  - Boolean flag indicating end of word
  - Optional: word count, frequency data
- **Go Implementation**: Array-based `children [26]*TrieNode` vs map-based `children map[rune]*TrieNode`
- **Array vs Map**: Array faster (cache-friendly) but uses more memory
- **Operations**:
  - Insert: O(m) where m is word length
  - Search: O(m) to find if word exists
  - StartsWith: O(m) to find if prefix exists
- **Use Cases**:
  - Word search problems
  - Autocomplete systems
  - IP routing tables
  - Spell checkers
- **Memory**: Can be memory-intensive for sparse datasets
- **Optimization**: Path compression for single-child chains

**Linked Lists**

*Singly Linked Lists*
- **Structure**: Linear collection where each node contains data and pointer to next node
- **Node Definition**: `type ListNode struct { Val int; Next *ListNode }`
- **Properties**:
  - Dynamic size: Can grow/shrink during runtime
  - Sequential access: Must traverse from head to reach specific position
  - No random access: Cannot directly access middle elements
- **Operations**:
  - Insert at head: O(1)
  - Insert at tail: O(n) unless tail pointer maintained
  - Delete: O(1) if have pointer to node, O(n) to find node
  - Search: O(n) linear search
- **Use Cases**:
  - When size is unknown or frequently changing
  - Implementing stacks, queues
  - Representing sparse data
  - Undo functionality in applications
- **Memory**: Each node ~16 bytes (8 bytes data + 8 bytes pointer)
- **Cache Performance**: Poor due to non-contiguous memory layout

*Floyd's Cycle Detection*
- **Algorithm**: Two pointers moving at different speeds (slow=1, fast=2)
- **Principle**: If cycle exists, fast pointer will eventually meet slow pointer
- **Detection**: When fast and slow pointers meet, cycle exists
- **Finding Start**: After detection, reset one pointer to head, move both at same speed
- **Time Complexity**: O(n) - at most 2n iterations
- **Space Complexity**: O(1) - only two pointers used
- **Use Cases**:
  - Detect cycle in linked list
  - Find starting point of cycle
  - Detect infinite loops in function calls
- **Mathematical Proof**: Based on pigeonhole principle and modular arithmetic
- **Variants**: Brent's algorithm, three-pointers technique

**Graphs**

*Adjacency Lists*
- **Structure**: Array/slice of lists where index i contains neighbors of vertex i
- **Go Implementation**: `[][]int` for unweighted, `[][]Edge` for weighted graphs
- **Memory Usage**: O(V + E) where V is vertices, E is edges
- **Advantages**:
  - Memory efficient for sparse graphs
  - Fast neighbor iteration
  - Easy to add/remove edges
- **Disadvantages**:
  - Slower edge existence check O(degree)
  - Poor cache locality for dense graphs
- **Use Cases**:
  - Social networks (sparse connections)
  - Web crawling (follow links)
  - Dependency graphs
  - Transportation networks
- **Alternative**: Adjacency matrix for dense graphs

*Union-Find (Disjoint Set)*
- **Structure**: Forest of trees where each tree represents a connected component
- **Core Operations**:
  - Find: Determine which set element belongs to
  - Union: Merge two sets into one
- **Path Compression**: Make nodes point directly to root during Find
- **Union by Rank**: Attach smaller tree under root of larger tree
- **Time Complexity**: Nearly O(1) amortized with both optimizations
- **Go Implementation**:
  ```go
  type UnionFind struct {
      parent []int  // parent[i] = parent of node i
      rank   []int  // rank[i] = approximate depth of tree rooted at i
  }
  ```
- **Use Cases**:
  - Detect cycles in undirected graphs
  - Find connected components
  - Kruskal's minimum spanning tree algorithm
  - Percolation problems
- **Inverse Ackermann**: Theoretical time complexity involves α(n)

**Stacks & Queues**

*Stacks*
- **Structure**: LIFO (Last In, First Out) data structure
- **Go Implementation**: Use slice with append/pop operations
- **Core Operations**:
  - Push: Add element to top O(1)
  - Pop: Remove element from top O(1)
  - Peek/Top: View top element without removing O(1)
- **Use Cases**:
  - Function call management (call stack)
  - Expression evaluation and parsing
  - Backtracking algorithms
  - Undo operations in applications
  - Depth-first search implementation
- **Problems**:
  - Valid Parentheses: Match opening/closing brackets
  - Evaluate Reverse Polish Notation
  - Daily Temperatures: Monotonic stack
- **Memory**: Dynamic array implementation with amortized O(1) operations

*Queues*
- **Structure**: FIFO (First In, First Out) data structure
- **Go Implementation**: Use slice with append/shift or circular buffer
- **Core Operations**:
  - Enqueue: Add element to rear O(1)
  - Dequeue: Remove element from front O(1) with proper implementation
  - Front: View front element O(1)
- **Circular Buffer**: Efficient implementation avoiding slice shifting
- **Use Cases**:
  - Breadth-first search traversal
  - Level-order tree traversal
  - Job scheduling in operating systems
  - Handling requests in web servers
- **Double-Ended Queue (Deque)**: Allows insertion/deletion at both ends
- **Priority Queue**: Elements have priorities, highest priority dequeued first

### Algorithms Used

**Sorting & Searching**

*Binary Search*
- **Principle**: Divide and conquer on sorted data
- **Algorithm**:
  1. Compare target with middle element
  2. If equal, found; if less, search left half; if greater, search right half
  3. Repeat until found or search space empty
- **Time Complexity**: O(log n) - halves search space each iteration
- **Space Complexity**: O(1) iterative, O(log n) recursive
- **Variants**:
  - Find exact match
  - Find first/last occurrence
  - Find insertion point
  - Search in rotated sorted array
- **Template**:
  ```go
  func binarySearch(nums []int, target int) int {
      left, right := 0, len(nums)-1
      for left <= right {
          mid := left + (right-left)/2  // Avoid overflow
          if nums[mid] == target {
              return mid
          } else if nums[mid] < target {
              left = mid + 1
          } else {
              right = mid - 1
          }
      }
      return -1
  }
  ```
- **Common Mistakes**: Off-by-one errors, integer overflow, infinite loops
- **Applications**: Search in sorted arrays, find peak elements, search 2D matrix

*Merge Sort*
- **Principle**: Divide array into halves, sort each half, merge sorted halves
- **Algorithm**:
  1. Divide array into two halves
  2. Recursively sort both halves
  3. Merge sorted halves into single sorted array
- **Time Complexity**: O(n log n) in all cases
- **Space Complexity**: O(n) for temporary arrays
- **Stability**: Maintains relative order of equal elements
- **Use Cases**:
  - Merge k sorted lists
  - Count inversions in array
  - External sorting for large datasets
- **Merge Process**: Two-pointer technique to combine sorted arrays
- **Advantages**: Consistent performance, stable, parallelizable

**Graph Algorithms**

*DFS (Depth-First Search)*
- **Principle**: Explore as far as possible before backtracking
- **Implementation**: Recursive or iterative with explicit stack
- **Time Complexity**: O(V + E) where V is vertices, E is edges
- **Space Complexity**: O(V) for recursion stack or explicit stack
- **Traversal Order**: Varies based on graph structure and starting point
- **Applications**:
  - Tree traversal (preorder, inorder, postorder)
  - Topological sorting
  - Finding connected components
  - Cycle detection in directed graphs
  - Path finding (not shortest path)
- **Template**:
  ```go
  func dfs(graph [][]int, visited []bool, node int) {
      visited[node] = true
      // Process current node
      for _, neighbor := range graph[node] {
          if !visited[neighbor] {
              dfs(graph, visited, neighbor)
          }
      }
  }
  ```
- **Variants**: DFS with path tracking, DFS with timestamp

*BFS (Breadth-First Search)*
- **Principle**: Explore all neighbors at current depth before moving deeper
- **Implementation**: Iterative with queue data structure
- **Time Complexity**: O(V + E) for complete traversal
- **Space Complexity**: O(V) for queue and visited array
- **Properties**: Finds shortest path in unweighted graphs
- **Applications**:
  - Shortest path in unweighted graphs
  - Level-order tree traversal
  - Finding minimum number of steps
  - Web crawling
  - Social network analysis (degrees of separation)
- **Template**:
  ```go
  func bfs(graph [][]int, start int) {
      queue := []int{start}
      visited := make([]bool, len(graph))
      visited[start] = true
      
      for len(queue) > 0 {
          node := queue[0]
          queue = queue[1:]
          // Process current node
          for _, neighbor := range graph[node] {
              if !visited[neighbor] {
                  visited[neighbor] = true
                  queue = append(queue, neighbor)
              }
          }
      }
  }
  ```
- **Multi-source BFS**: Start from multiple nodes simultaneously

*Topological Sort*
- **Principle**: Linear ordering of vertices where each vertex comes before its dependencies
- **Prerequisites**: Works only on Directed Acyclic Graphs (DAGs)
- **Algorithms**:
  - Kahn's Algorithm: Use in-degree counting and queue
  - DFS-based: Use finishing times in reverse order
- **Time Complexity**: O(V + E)
- **Space Complexity**: O(V) for data structures
- **Applications**:
  - Course scheduling with prerequisites
  - Build system dependency resolution
  - Task scheduling
  - Deadlock detection
- **Cycle Detection**: If topological sort cannot include all vertices, cycle exists
- **Implementation**: 
  ```go
  func topologicalSort(graph [][]int) []int {
      inDegree := make([]int, len(graph))
      // Calculate in-degrees
      for i := 0; i < len(graph); i++ {
          for _, neighbor := range graph[i] {
              inDegree[neighbor]++
          }
      }
      
      queue := []int{}
      for i, degree := range inDegree {
          if degree == 0 {
              queue = append(queue, i)
          }
      }
      
      result := []int{}
      for len(queue) > 0 {
          node := queue[0]
          queue = queue[1:]
          result = append(result, node)
          
          for _, neighbor := range graph[node] {
              inDegree[neighbor]--
              if inDegree[neighbor] == 0 {
                  queue = append(queue, neighbor)
              }
          }
      }
      
      return result
  }
  ```

**Dynamic Programming**

*Memoization (Top-Down)*
- **Principle**: Store results of expensive function calls to avoid recomputation
- **Implementation**: Use hash map or array to cache results
- **When to Use**: When you have overlapping subproblems in recursive solution
- **Advantages**: 
  - Easy to implement from recursive solution
  - Only computes needed subproblems
  - Natural recursion structure
- **Disadvantages**:
  - Function call overhead
  - Stack space usage
  - Potential stack overflow for deep recursion
- **Template**:
  ```go
  func fibMemo(n int, memo map[int]int) int {
      if n <= 1 {
          return n
      }
      if val, exists := memo[n]; exists {
          return val
      }
      memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
      return memo[n]
  }
  ```
- **Use Cases**: Tree problems, string matching, optimization problems

*Bottom-Up DP*
- **Principle**: Build solution iteratively from smaller subproblems
- **Implementation**: Use array/table to store intermediate results
- **Advantages**:
  - No function call overhead
  - Better space complexity control
  - No stack overflow risk
  - Usually faster in practice
- **Disadvantages**:
  - May compute unnecessary subproblems
  - Requires careful order of computation
  - Less intuitive than recursive approach
- **Template**:
  ```go
  func fibBottomUp(n int) int {
      if n <= 1 {
          return n
      }
      dp := make([]int, n+1)
      dp[0], dp[1] = 0, 1
      for i := 2; i <= n; i++ {
          dp[i] = dp[i-1] + dp[i-2]
      }
      return dp[n]
  }
  ```
- **Space Optimization**: Often can reduce from O(n) to O(1) by keeping only needed values

*Space Optimization Techniques*
- **Rolling Array**: For 2D DP, use only current and previous row
- **State Compression**: Use bit manipulation for subset DP
- **In-place DP**: Modify input array when possible
- **Example - House Robber**:
  ```go
  func rob(nums []int) int {
      if len(nums) == 0 {
          return 0
      }
      prev2, prev1 := 0, 0
      for _, num := range nums {
          current := max(prev1, prev2+num)
          prev2, prev1 = prev1, current
      }
      return prev1
  }
  ```
- **Benefits**: Reduces space complexity from O(n) to O(1)

**Two Pointers**

*Left/Right Pointers*
- **Principle**: Two pointers moving from opposite ends of array
- **Movement**: Pointers move toward each other based on conditions
- **Time Complexity**: O(n) - each element visited at most once
- **Space Complexity**: O(1) - only pointer variables
- **Use Cases**:
  - Two Sum in sorted array
  - Valid Palindrome checking
  - Container With Most Water
  - Trapping Rain Water
- **Template**:
  ```go
  func twoPointers(nums []int) {
      left, right := 0, len(nums)-1
      for left < right {
          if condition {
              left++
          } else {
              right--
          }
      }
  }
  ```
- **Advantages**: Efficient, easy to implement, works on sorted data

*Fast/Slow Pointers*
- **Principle**: Two pointers moving at different speeds
- **Common Speeds**: Slow moves 1 step, fast moves 2 steps
- **Applications**:
  - Cycle detection in linked lists
  - Finding middle of linked list
  - Detecting start of cycle
- **Floyd's Algorithm**: Fast and slow pointers will meet if cycle exists
- **Template**:
  ```go
  func hasCycle(head *ListNode) bool {
      if head == nil || head.Next == nil {
          return false
      }
      slow, fast := head, head.Next
      for fast != nil && fast.Next != nil {
          if slow == fast {
              return true
          }
          slow = slow.Next
          fast = fast.Next.Next
      }
      return false
  }
  ```
- **Finding Cycle Start**: Reset one pointer to head after detection

**Sliding Window**

*Variable-Size Window*
- **Principle**: Window expands and contracts based on conditions
- **Expansion**: Add elements to window while condition valid
- **Contraction**: Remove elements when condition violated
- **Use Cases**:
  - Longest substring without repeating characters
  - Minimum window substring
  - Longest subarray with at most k distinct characters
- **Template**:
  ```go
  func slidingWindow(s string) int {
      left, right := 0, 0
      window := make(map[byte]int)
      maxLen := 0
      
      for right < len(s) {
          // Expand window
          window[s[right]]++
          right++
          
          // Contract window when invalid
          for /* window invalid */ {
              window[s[left]]--
              if window[s[left]] == 0 {
                  delete(window, s[left])
              }
              left++
          }
          
          // Update answer
          maxLen = max(maxLen, right-left)
      }
      return maxLen
  }
  ```
- **Time Complexity**: O(n) - each element added and removed at most once

*Fixed-Size Window*
- **Principle**: Window of constant size sliding through array
- **Movement**: Add new element, remove old element
- **Use Cases**:
  - Maximum sum subarray of size k
  - Average of subarrays of size k
  - Find anagrams in string
- **Template**:
  ```go
  func fixedWindow(nums []int, k int) []int {
      if len(nums) < k {
          return []int{}
      }
      
      result := []int{}
      windowSum := 0
      
      // Initialize first window
      for i := 0; i < k; i++ {
          windowSum += nums[i]
      }
      result = append(result, windowSum)
      
      // Slide window
      for i := k; i < len(nums); i++ {
          windowSum = windowSum - nums[i-k] + nums[i]
          result = append(result, windowSum)
      }
      
      return result
  }
  ```
- **Optimization**: Avoid recalculating entire window each time

**Backtracking**

*Recursive Exploration*
- **Principle**: Systematically explore all possibilities by making choices and undoing them
- **Structure**: Make choice → Explore → Undo choice
- **Decision Tree**: Each level represents a choice, leaves are complete solutions
- **Use Cases**:
  - Generate all permutations/combinations
  - N-Queens problem
  - Sudoku solver
  - Word search in grid
- **Template**:
  ```go
  func backtrack(path []int, choices []int, result *[][]int) {
      // Base case: valid solution found
      if len(path) == targetLength {
          *result = append(*result, append([]int{}, path...))
          return
      }
      
      // Try each choice
      for i, choice := range choices {
          // Make choice
          path = append(path, choice)
          
          // Explore with this choice
          backtrack(path, choices[i+1:], result)
          
          // Undo choice
          path = path[:len(path)-1]
      }
  }
  ```
- **State Management**: Careful handling of shared state between recursive calls

*Pruning*
- **Principle**: Skip branches that cannot lead to valid solutions
- **Benefits**: Dramatic performance improvement by avoiding useless computation
- **Techniques**:
  - Bound checking: Skip if partial solution exceeds bounds
  - Constraint violation: Skip if constraints violated
  - Duplicate detection: Skip if same state seen before
- **Example - N-Queens**: Skip if new queen attacks existing queens
- **Implementation**: Add conditions before recursive calls
- **Trade-off**: More complex code but better performance

**Greedy Algorithms**

*Local Optimal Choices*
- **Principle**: Make locally optimal choice at each step
- **Assumption**: Local optimum leads to global optimum
- **When It Works**: Problems with greedy choice property and optimal substructure
- **Use Cases**:
  - Activity selection problem
  - Huffman coding
  - Minimum spanning tree (Kruskal's, Prim's)
  - Fractional knapsack
- **Advantages**: Simple, efficient, easy to implement
- **Disadvantages**: Doesn't always work, requires proof of correctness

*Proof Techniques*
- **Greedy Choice Property**: Optimal solution contains greedy choice
- **Optimal Substructure**: Problem can be broken into subproblems
- **Exchange Argument**: Show greedy solution is at least as good as optimal
- **Inductive Proof**: Prove by induction on problem size
- **Example - Activity Selection**: Earliest finish time gives maximum activities

**Bit Manipulation**

*XOR Operations*
- **Properties**:
  - `a ^ a = 0` (self-cancellation)
  - `a ^ 0 = a` (identity)
  - `a ^ b ^ c = a ^ (b ^ c)` (associativity)
  - `a ^ b = b ^ a` (commutativity)
- **Finding Unique Element**: XOR all elements, duplicates cancel out
- **Swapping**: `a ^= b; b ^= a; a ^= b` swaps without temporary variable
- **Use Cases**:
  - Single Number problem
  - Finding missing number
  - Detecting bit differences
- **Implementation**:
  ```go
  func singleNumber(nums []int) int {
      result := 0
      for _, num := range nums {
          result ^= num
      }
      return result
  }
  ```

*Bit Masks*
- **Principle**: Use integer bits to represent sets
- **Operations**:
  - Set bit: `mask |= (1 << i)`
  - Clear bit: `mask &= ^(1 << i)`
  - Toggle bit: `mask ^= (1 << i)`
  - Check bit: `mask & (1 << i) != 0`
- **Use Cases**:
  - Subset enumeration
  - Dynamic programming on subsets
  - Representing states efficiently
- **Example - Subsets**:
  ```go
  func subsets(nums []int) [][]int {
      n := len(nums)
      result := [][]int{}
      
      for mask := 0; mask < (1 << n); mask++ {
          subset := []int{}
          for i := 0; i < n; i++ {
              if mask & (1 << i) != 0 {
                  subset = append(subset, nums[i])
              }
          }
          result = append(result, subset)
      }
      return result
  }
  ```
- **Advantages**: Compact representation, fast operations

### Time Complexity Analysis

**Complexity Classes**
- **O(1) - Constant**: Hash table operations, array access, arithmetic operations
- **O(log n) - Logarithmic**: Binary search, balanced tree operations, heap operations
- **O(n) - Linear**: Single pass through array, tree traversal, simple graph traversal
- **O(n log n) - Linearithmic**: Efficient sorting algorithms, divide-and-conquer
- **O(n²) - Quadratic**: Nested loops, bubble sort, some dynamic programming
- **O(n³) - Cubic**: Triple nested loops, matrix multiplication
- **O(2ⁿ) - Exponential**: Brute force subset generation, recursive fibonacci
- **O(n!) - Factorial**: Brute force permutation generation

**Amortized Analysis**
- **Principle**: Average time per operation over sequence of operations
- **Techniques**:
  - Aggregate method: Total cost divided by number of operations
  - Accounting method: Assign costs to operations
  - Potential method: Use potential function
- **Examples**:
  - Dynamic array resizing: O(1) amortized insertion
  - Union-Find with path compression: Nearly O(1) amortized

**Space Complexity**
- **Input Space**: Space used by input (not counted in auxiliary space)
- **Auxiliary Space**: Extra space used by algorithm
- **Total Space**: Input space + auxiliary space
- **In-place**: O(1) auxiliary space (may modify input)
- **Examples**:
  - Merge sort: O(n) auxiliary space for temporary arrays
  - Quick sort: O(log n) auxiliary space for recursion stack
  - Heap sort: O(1) auxiliary space (in-place sorting)

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