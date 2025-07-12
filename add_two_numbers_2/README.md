# 2. Add Two Numbers

https://leetcode.com/problems/add-two-numbers/

# Description

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```
# Explanation

No problem\! Here's the explanation of the `add_two_numbers_2` package, formatted for a Markdown file.

-----

# Adding Two Numbers Represented as Linked Lists

This document explains the Go code for adding two numbers stored in a unique way: as **linked lists**, where each node contains a single digit, and the digits are stored in **reverse order**.

For instance:

  * The number `342` is represented as: `2 -> 4 -> 3`
  * The number `465` is represented as: `5 -> 6 -> 4`

When added together, `342 + 465 = 807`. The resulting linked list should be: `7 -> 0 -> 8`.

-----

## `ListNode` Structure

The fundamental building block of our linked lists is defined as:

```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

  * `Val`: An integer representing the digit stored in this node (0-9).
  * `Next`: A pointer to the next `ListNode` in the sequence. If `nil`, it signifies the end of the list.

-----

## `addTwoNumbers` Function

This is the core function that performs the addition.

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Initialize dummy head for the result list to simplify construction.
    dummyHead := &ListNode{}
    current := dummyHead // 'current' pointer will traverse our new list
    carry := 0           // Stores any carry-over digit

    // Loop until both lists are exhausted and no carry remains.
    for l1 != nil || l2 != nil || carry > 0 {
        sum := carry // Start sum with carry from previous iteration

        // Add l1's digit if available, then advance l1.
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        // Add l2's digit if available, then advance l2.
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        // Calculate the new carry and the digit for the current position.
        carry = sum / 10           // e.g., if sum is 17, carry is 1
        current.Next = &ListNode{Val: sum % 10} // e.g., if sum is 17, digit is 7

        // Move 'current' to the newly added node for the next iteration.
        current = current.Next
    }

    // Return the result list, starting after the dummy head.
    return dummyHead.Next
}
```

### How it Works:

1.  **Initialization**:

      * `dummyHead := &ListNode{}`: A **dummy head node** is created. This simplifies adding the first node to our result list, as we can always append to `current.Next`. The actual result list will begin from `dummyHead.Next`.
      * `current := dummyHead`: This pointer will move along our new result list, always pointing to the last node added.
      * `carry := 0`: This variable holds any carry-over from adding digits (e.g., when `5 + 6 = 11`, `1` is carried over).

2.  **Main Loop (`for l1 != nil || l2 != nil || carry > 0`)**:

      * This loop continues as long as there are digits left in either input list (`l1` or `l2`) or a `carry` digit needs to be processed.
      * `sum := carry`: The `sum` for the current position starts with any `carry` from the previous position.
      * **Adding Digits**: If `l1` (or `l2`) still has nodes, its `Val` (digit) is added to `sum`, and the respective list pointer is moved to its `Next` node.
      * **Calculating Carry and New Digit**:
          * `carry = sum / 10`: The new `carry` is determined (e.g., `17 / 10 = 1`).
          * `current.Next = &ListNode{Val: sum % 10}`: A new node is created for the result list. Its `Val` is the current digit (e.g., `17 % 10 = 7`).
      * **Advancing `current`**: `current = current.Next` moves our `current` pointer to this newly added node, ready for the next digit.

3.  **Return Value (`return dummyHead.Next`)**:

      * Once the loop finishes (meaning all digits from both lists have been added, and any final carry has been processed), `dummyHead.Next` points to the first actual node of our sum list. This is returned as the final result.

-----

## Helper Functions (for Testing)

These functions are typically used for testing and convenience, not part of the core solution for platforms like LeetCode.

### `createList` Function

```go
func createList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    dummy := &ListNode{}
    current := dummy
    for _, val := range arr {
        current.Next = &ListNode{Val: val}
        current = current.Next
    }
    return dummy.Next
}
```

  * This function takes an array of integers (e.g., `[2, 4, 3]`) and converts it into a `ListNode` representation.
  * It uses a similar dummy node pattern to simplify list construction.

### `listToArray` Function

```go
func listToArray(head *ListNode) []int {
    var result []int
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    return result
}
```

  * This function performs the opposite of `createList`: it takes a `ListNode` (the head of a list) and converts it back into an array of integers. This is useful for verifying the output of `addTwoNumbers` in tests.

-----

## Example Walkthrough

Let's trace `addTwoNumbers` with `l1 = [2,4,3]` (342) and `l2 = [5,6,4]` (465):

**Initial State**: `dummyHead = {}`, `current = {}`, `carry = 0`

1.  **Units Place (2 + 5)**:

      * `sum = 0 + 2 + 5 = 7`
      * `carry = 7 / 10 = 0`
      * New node: `Val: 7`
      * Result list: `7`
      * `l1` moves to `4`, `l2` moves to `6`

2.  **Tens Place (4 + 6 + carry 0)**:

      * `sum = 0 + 4 + 6 = 10`
      * `carry = 10 / 10 = 1`
      * New node: `Val: 0`
      * Result list: `7 -> 0`
      * `l1` moves to `3`, `l2` moves to `4`

3.  **Hundreds Place (3 + 4 + carry 1)**:

      * `sum = 1 + 3 + 4 = 8`
      * `carry = 8 / 10 = 0`
      * New node: `Val: 8`
      * Result list: `7 -> 0 -> 8`
      * `l1` becomes `nil`, `l2` becomes `nil`

4.  **Loop End**: `l1`, `l2`, and `carry` are all `0`. The loop terminates.

**Final Result**: `dummyHead.Next` points to the list `7 -> 0 -> 8`, correctly representing `807`.

-----

## Constraints Handled

The `addTwoNumbers` function is designed to handle various constraints and edge cases:

  * **Lists of Different Lengths**: The `for` loop condition (`l1 != nil || l2 != nil || carry > 0`) gracefully handles situations where one list is shorter than the other by simply adding `0` for the missing digits.
  * **Carry Propagation**: A final `carry` (e.g., `99 + 1 = 100` would result in `0 -> 0 -> 1`) is correctly appended to the end of the list because the loop continues as long as `carry > 0`.
  * **Single-Digit Numbers**: Adding `0 + 0` correctly results in `[0]`.
  * **No Leading Zeros**: Except for the number `0` itself, the algorithm ensures the result list doesn't have unnecessary leading zeros.

This solution provides a robust and efficient way to add numbers represented as linked lists in reverse order.