# Go LeetCode Solutions

High-performance Go solutions to LeetCode problems with optimal algorithms and Go best practices.

## Table of Contents

- [Features](#features)
- [Quick Start](#quick-start)
- [Problem Categories](#problem-categories)
  - [Arrays & Hashing](#arrays--hashing)
  - [Two Pointers](#two-pointers)
  - [Sliding Window](#sliding-window)
  - [Stack](#stack)
  - [Binary Search](#binary-search)
  - [Linked Lists](#linked-lists)
  - [Trees](#trees)
  - [Tries](#tries)
  - [Heap / Priority Queue](#heap--priority-queue)
  - [Backtracking](#backtracking)
  - [Graphs](#graphs)
  - [Advanced Graphs](#advanced-graphs)
  - [Dynamic Programming](#dynamic-programming)
  - [Greedy](#greedy)
  - [Intervals](#intervals)
  - [Math & Geometry](#math--geometry)
  - [Bit Manipulation](#bit-manipulation)
  - [Concurrency](#concurrency)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Optimal Algorithms**: Best time/space complexity solutions
- **Go Best Practices**: Modern Go 1.24+ patterns and idioms
- **Performance Focused**: Zero-allocation patterns and optimizations
- **Complete Testing**: 100% test coverage with benchmarks
- **Comprehensive Documentation**: Detailed explanations for each solution

## Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/go-leetcode-master.git
cd go-leetcode-master

# Run tests for a specific problem
cd two_sum_1
go test -v -bench=. -benchmem

# Run all tests
make test

# Run benchmarks
make bench
```

## Problem Categories

### Arrays & Hashing

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [1. Two Sum](./two_sum_1/) | Easy | [Go](./two_sum_1/solution.go) | Hash Map, Single Pass |
| [49. Group Anagrams](./group_anagrams_49/) | Medium | [Go](./group_anagrams_49/solution.go) | Hash Map, Sorting |
| [128. Longest Consecutive Sequence](./longest_consecutive_sequence_128/) | Medium | [Go](./longest_consecutive_sequence_128/solution.go) | Hash Set, Union Find |
| [217. Contains Duplicate](./contains_duplicate_217/) | Easy | [Go](./contains_duplicate_217/solution.go) | Hash Set |
| [238. Product of Array Except Self](./product_of_array_except_self_238/) | Medium | [Go](./product_of_array_except_self_238/solution.go) | Prefix/Suffix Products |
| [242. Valid Anagram](./valid_anagram_242/) | Easy | [Go](./valid_anagram_242/solution.go) | Hash Map, Sorting |
| [347. Top K Frequent Elements](./top_k_frequent_elements_347/) | Medium | [Go](./top_k_frequent_elements_347/solution.go) | Hash Map, Heap |
| [350. Intersection of Two Arrays II](./intersection_of_two_arrays_ii_350/) | Easy | [Go](./intersection_of_two_arrays_ii_350/solution.go) | Hash Map, Two Pointers |

### Two Pointers

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [11. Container With Most Water](./container_with_most_water_11/) | Medium | [Go](./container_with_most_water_11/solution.go) | Two Pointers, Greedy |
| [15. 3Sum](./three_sum_15/) | Medium | [Go](./three_sum_15/solution.go) | Two Pointers, Sorting |
| [125. Valid Palindrome](./valid_palindrome_125/) | Easy | [Go](./valid_palindrome_125/solution.go) | Two Pointers, String Processing |
| [141. Linked List Cycle](./linked_list_cycle_141/) | Easy | [Go](./linked_list_cycle_141/solution.go) | Floyd's Cycle Detection |
| [160. Intersection of Two Linked Lists](./intersection_of_two_linked_lists_160/) | Easy | [Go](./intersection_of_two_linked_lists_160/solution.go) | Two Pointers |

### Sliding Window

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [3. Longest Substring Without Repeating Characters](./longest_substring_without_repeating_characters_3/) | Medium | [Go](./longest_substring_without_repeating_characters_3/solution.go) | Sliding Window, Hash Set |
| [76. Minimum Window Substring](./minimum_window_substring_76/) | Hard | [Go](./minimum_window_substring_76/solution.go) | Sliding Window, Hash Map |
| [424. Longest Repeating Character Replacement](./longest_repeating_character_replacement_424/) | Medium | [Go](./longest_repeating_character_replacement_424/solution.go) | Sliding Window, Hash Map |
| [643. Maximum Average Subarray I](./maximum_average_subarray_i_643/) | Easy | [Go](./maximum_average_subarray_i_643/solution.go) | Sliding Window |

### Stack

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [20. Valid Parentheses](./valid_parentheses_20/) | Easy | [Go](./valid_parentheses_20/solution.go) | Stack, Hash Map |
| [155. Min Stack](./min_stack_155/) | Medium | [Go](./min_stack_155/solution.go) | Stack, Design |

### Binary Search

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [33. Search in Rotated Sorted Array](./search_in_rotated_sorted_array_33/) | Medium | [Go](./search_in_rotated_sorted_array_33/solution.go) | Binary Search, Array Rotation |
| [153. Find Minimum in Rotated Sorted Array](./find_minimum_in_rotated_sorted_array_153/) | Medium | [Go](./find_minimum_in_rotated_sorted_array_153/solution.go) | Binary Search, Array Rotation |
| [981. Time Based Key-Value Store](./time_based_key_value_store_981/) | Medium | [Go](./time_based_key_value_store_981/solution.go) | Binary Search, Hash Map, Design |

### Linked Lists

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [2. Add Two Numbers](./add_two_numbers_2/) | Medium | [Go](./add_two_numbers_2/solution.go) | Linked List, Math |
| [19. Remove Nth Node From End of List](./remove_nth_node_from_end_of_list_19/) | Medium | [Go](./remove_nth_node_from_end_of_list_19/solution.go) | Two Pointers, Linked List |
| [21. Merge Two Sorted Lists](./merge_two_sorted_lists_21/) | Easy | [Go](./merge_two_sorted_lists_21/solution.go) | Linked List, Recursion |
| [23. Merge k Sorted Lists](./merge_k_sorted_lists_23/) | Hard | [Go](./merge_k_sorted_lists_23/solution.go) | Linked List, Heap, Divide & Conquer |
| [143. Reorder List](./reorder_list_143/) | Medium | [Go](./reorder_list_143/solution.go) | Linked List, Two Pointers |
| [206. Reverse Linked List](./reverse_linked_list_206/) | Easy | [Go](./reverse_linked_list_206/solution.go) | Linked List, Recursion |
| [234. Palindrome Linked List](./palindrome_linked_list_234/) | Easy | [Go](./palindrome_linked_list_234/solution.go) | Linked List, Two Pointers |
| [237. Delete Node in a Linked List](./delete_node_in_a_linked_list_237/) | Medium | [Go](./delete_node_in_a_linked_list_237/solution.go) | Linked List |

### Trees

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [98. Validate Binary Search Tree](./validate_binary_search_tree_98/) | Medium | [Go](./validate_binary_search_tree_98/solution.go) | BST, DFS, Recursion |
| [100. Same Tree](./same_tree_100/) | Easy | [Go](./same_tree_100/solution.go) | Binary Tree, DFS |
| [101. Symmetric Tree](./symmetric_tree_101/) | Easy | [Go](./symmetric_tree_101/solution.go) | Binary Tree, DFS |
| [102. Binary Tree Level Order Traversal](./binary_tree_level_order_traversal_102/) | Medium | [Go](./binary_tree_level_order_traversal_102/solution.go) | Binary Tree, BFS |
| [104. Maximum Depth of Binary Tree](./maximum_depth_of_binary_tree_104/) | Easy | [Go](./maximum_depth_of_binary_tree_104/solution.go) | Binary Tree, DFS |
| [105. Construct Binary Tree from Preorder and Inorder Traversal](./construct_binary_tree_from_preorder_and_inorder_traversal_105/) | Medium | [Go](./construct_binary_tree_from_preorder_and_inorder_traversal_105/solution.go) | Binary Tree, DFS |
| [124. Binary Tree Maximum Path Sum](./binary_tree_maximum_path_sum_124/) | Hard | [Go](./binary_tree_maximum_path_sum_124/solution.go) | Binary Tree, DFS |
| [226. Invert Binary Tree](./invert_binary_tree_226/) | Easy | [Go](./invert_binary_tree_226/solution.go) | Binary Tree, DFS |
| [230. Kth Smallest Element in a BST](./kth_smallest_element_in_a_bst_230/) | Medium | [Go](./kth_smallest_element_in_a_bst_230/solution.go) | BST, DFS |
| [235. Lowest Common Ancestor of a Binary Search Tree](./lowest_common_ancestor_of_a_binary_search_tree_235/) | Medium | [Go](./lowest_common_ancestor_of_a_binary_search_tree_235/solution.go) | BST, DFS |
| [572. Subtree of Another Tree](./subtree_of_another_tree_572/) | Easy | [Go](./subtree_of_another_tree_572/solution.go) | Binary Tree, DFS |

### Tries

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [208. Implement Trie](./implement_trie_208/) | Medium | [Go](./implement_trie_208/solution.go) | Trie, Design |
| [212. Word Search II](./word_search_212/) | Hard | [Go](./word_search_212/solution.go) | Trie, Backtracking |

### Heap / Priority Queue

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [295. Find Median from Data Stream](./find_median_from_data_stream_295/) | Hard | [Go](./find_median_from_data_stream_295/solution.go) | Heap, Design |

### Backtracking

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [17. Letter Combinations of a Phone Number](./letter_combinations_of_a_phone_number_17/) | Medium | [Go](./letter_combinations_of_a_phone_number_17/solution.go) | Backtracking, String |
| [22. Generate Parentheses](./generate_parentheses_22/) | Medium | [Go](./generate_parentheses_22/solution.go) | Backtracking, String |
| [37. Sudoku Solver](./soduku_solver_37/) | Hard | [Go](./soduku_solver_37/solution.go) | Backtracking, Matrix |
| [79. Word Search](./word_search_79/) | Medium | [Go](./word_search_79/solution.go) | Backtracking, Matrix |
| [698. Partition to K Equal Sum Subsets](./partition_to_k_equal_sum_subsets_698/) | Medium | [Go](./partition_to_k_equal_sum_subsets_698/solution.go) | Backtracking, Bit Manipulation |
| [1239. Maximum Length of a Concatenated String with Unique Characters](./maximum_length_of_a_concatenated_string_with_unique_characters_1239/) | Medium | [Go](./maximum_length_of_a_concatenated_string_with_unique_characters_1239/solution.go) | Backtracking, Bit Manipulation |

### Graphs

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [200. Number of Islands](./number_of_islands_200/) | Medium | [Go](./number_of_islands_200/solution.go) | DFS, BFS, Union Find |
| [207. Course Schedule](./course_schedule_207/) | Medium | [Go](./course_schedule_207/solution.go) | Topological Sort, DFS |
| [261. Graph Valid Tree](./graph_valid_tree_261/) | Medium | [Go](./graph_valid_tree_261/solution.go) | DFS, BFS, Union Find |
| [269. Alien Dictionary](./alien_dictionary_269/) | Hard | [Go](./alien_dictionary_269/solution.go) | Topological Sort, DFS |
| [417. Pacific Atlantic Water Flow](./pacific_atlantic_water_flow_417/) | Medium | [Go](./pacific_atlantic_water_flow_417/solution.go) | DFS, BFS |
| [847. Shortest Path Visiting All Nodes](./shortest_path_visiting_all_nodes_847/) | Hard | [Go](./shortest_path_visiting_all_nodes_847/solution.go) | BFS, Bit Manipulation |
| [953. Verifying an Alien Dictionary](./verifying_an_alien_dictionary_953/) | Easy | [Go](./verifying_an_alien_dictionary_953/solution.go) | Hash Map, String |

### Advanced Graphs

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [1192. Critical Connections in a Network](./critical_connections_in_a_network_1192/) | Hard | [Go](./critical_connections_in_a_network_1192/solution.go) | Tarjan's Algorithm, DFS |
| [1568. Minimum Number of Days to Disconnect Island](./minimum_number_of_days_to_disconnect_island_1568/) | Hard | [Go](./minimum_number_of_days_to_disconnect_island_1568/solution.go) | DFS, Articulation Points |

### Dynamic Programming

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [5. Longest Palindromic Substring](./longest_palindromic_substring_5/) | Medium | [Go](./longest_palindromic_substring_5/solution.go) | DP, String |
| [53. Maximum Subarray](./maximum_subarray_53/) | Medium | [Go](./maximum_subarray_53/solution.go) | DP, Kadane's Algorithm |
| [70. Climbing Stairs](./climbing_stairs_70/) | Easy | [Go](./climbing_stairs_70/solution.go) | DP, Fibonacci |
| [121. Best Time to Buy and Sell Stock](./best_time_to_buy_and_sell_stock_121/) | Easy | [Go](./best_time_to_buy_and_sell_stock_121/solution.go) | DP, Array |
| [122. Best Time to Buy and Sell Stock II](./best_time_to_buy_and_sell_stock_ii_122/) | Medium | [Go](./best_time_to_buy_and_sell_stock_ii_122/solution.go) | DP, Greedy |
| [132. Palindrome Partitioning II](./palindrome_partitioning_ii_132/) | Hard | [Go](./palindrome_partitioning_ii_132/solution.go) | DP, String |
| [139. Word Break](./word_break_139/) | Medium | [Go](./word_break_139/solution.go) | DP, String |
| [152. Maximum Product Subarray](./maximum_product_subarray_152/) | Medium | [Go](./maximum_product_subarray_152/solution.go) | DP, Array |
| [174. Dungeon Game](./dungeon_game_174/) | Hard | [Go](./dungeon_game_174/solution.go) | DP, Matrix |
| [198. House Robber](./house_robber_198/) | Medium | [Go](./house_robber_198/solution.go) | DP, Array |
| [213. House Robber II](./house_robber_ii_213/) | Medium | [Go](./house_robber_ii_213/solution.go) | DP, Array |
| [300. Longest Increasing Subsequence](./longest_increasing_subsequence_300/) | Medium | [Go](./longest_increasing_subsequence_300/solution.go) | DP, Binary Search |
| [322. Coin Change](./coin_change_322/) | Medium | [Go](./coin_change_322/solution.go) | DP, BFS |
| [338. Counting Bits](./counting_bits_338/) | Easy | [Go](./counting_bits_338/solution.go) | DP, Bit Manipulation |
| [354. Russian Doll Envelopes](./russian_doll_envelopes_354/) | Hard | [Go](./russian_doll_envelopes_354/solution.go) | DP, Binary Search |
| [377. Combination Sum IV](./combination_sum_iv_377/) | Medium | [Go](./combination_sum_iv_377/solution.go) | DP, Backtracking |
| [509. Fibonacci Number](./fibonacci_number_509/) | Easy | [Go](./fibonacci_number_509/solution.go) | DP, Math |
| [647. Palindromic Substrings](./palindromic_substrings_647/) | Medium | [Go](./palindromic_substrings_647/solution.go) | DP, String |
| [887. Super Egg Drop](./super_egg_drop_887/) | Hard | [Go](./super_egg_drop_887/solution.go) | DP, Binary Search |
| [1143. Longest Common Subsequence](./longest_common_subsequence_1143/) | Medium | [Go](./longest_common_subsequence_1143/solution.go) | DP, String |
| [1622. Fancy Sequence](./fancy_sequence_1622/) | Hard | [Go](./fancy_sequence_1622/solution.go) | DP, Math |
| [2518. Number of Great Partitions](./number_of_great_partitions_2518/) | Hard | [Go](./number_of_great_partitions_2518/solution.go) | DP, Combinatorics |

### Greedy

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [435. Non-overlapping Intervals](./non_overlapping_intervals_435/) | Medium | [Go](./non_overlapping_intervals_435/solution.go) | Greedy, Sorting |

### Intervals

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [56. Merge Intervals](./merge_intervals_56/) | Medium | [Go](./merge_intervals_56/solution.go) | Sorting, Array |
| [57. Insert Interval](./insert_interval_57/) | Medium | [Go](./insert_interval_57/solution.go) | Array, Intervals |
| [252. Meeting Rooms](./meeting_rooms_252/) | Easy | [Go](./meeting_rooms_252/solution.go) | Sorting, Array |
| [253. Meeting Rooms II](./meeting_rooms_ii_253/) | Medium | [Go](./meeting_rooms_ii_253/solution.go) | Heap, Sorting |

### Math & Geometry

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [48. Rotate Image](./rotate_image_48/) | Medium | [Go](./rotate_image_48/solution.go) | Matrix, Math |
| [54. Spiral Matrix](./spiral_matrix_54/) | Medium | [Go](./spiral_matrix_54/solution.go) | Matrix, Simulation |
| [73. Set Matrix Zeroes](./set_matrix_zeros_73/) | Medium | [Go](./set_matrix_zeros_73/solution.go) | Matrix, Array |
| [149. Max Points on a Line](./max_points_on_a_line_149/) | Hard | [Go](./max_points_on_a_line_149/solution.go) | Math, Geometry |
| [335. Self Crossing](./self_crossing_335/) | Hard | [Go](./self_crossing_335/solution.go) | Math, Geometry |
| [587. Erect the Fence](./erect_the_fence_587/) | Hard | [Go](./erect_the_fence_587/solution.go) | Geometry, Convex Hull |
| [1453. Maximum Number of Darts Inside of a Circular Dartboard](./max_num_of_darts_ioa_cd_1453/) | Hard | [Go](./max_num_of_darts_ioa_cd_1453/solution.go) | Geometry, Math |

### Bit Manipulation

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [136. Single Number](./single_number_136/) | Easy | [Go](./single_number_136/solution.go) | Bit Manipulation, XOR |
| [190. Reverse Bits](./reverse_bits_190/) | Easy | [Go](./reverse_bits_190/solution.go) | Bit Manipulation |
| [191. Number of 1 Bits](./number_of_1_bits_191/) | Easy | [Go](./number_of_1_bits_191/solution.go) | Bit Manipulation |
| [268. Missing Number](./missing_number_268/) | Easy | [Go](./missing_number_268/solution.go) | Bit Manipulation, Math |
| [371. Sum of Two Integers](./sum_of_two_integers_371/) | Medium | [Go](./sum_of_two_integers_371/solution.go) | Bit Manipulation |

### Concurrency

| Problem | Difficulty | Solution | Key Concepts |
|---------|------------|----------|--------------|
| [1117. Building H2O](./building_h2o_1117/) | Medium | [Go](./building_h2o_1117/solution.go) | Concurrency, Synchronization |
| [1226. The Dining Philosophers](./dining_philosophers_1226/) | Medium | [Go](./dining_philosophers_1226/solution.go) | Concurrency, Deadlock Prevention |

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-solution`)
3. Add your solution following the existing structure:
   - Create a directory named `problem_name_number/`
   - Add `solution.go` with the implementation
   - Add `solution_test.go` with comprehensive tests
   - Add `README.md` with problem description and approach
4. Ensure all tests pass and benchmarks are included
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.