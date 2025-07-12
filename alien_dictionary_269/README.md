# 269. Alien Dictionary

https://leetcode.com/problems/alien-dictionary/

## Difficulty:

Hard

## Description

There is a new alien language which uses the latin alphabet. 
However, the order among letters are unknown to you. You receive a 
list of non-empty words from the dictionary, where words are sorted 
lexicographically by the rules of this new language. Derive the order 
of letters in this language.

Example 1:
```
Input:
[
  "wrt",
  "wrf",
  "er",
  "ett",
  "rftt"
]

Output: "wertf"
```

Example 2:
```
Input:
[
  "z",
  "x"
]

Output: "zx"
```

Example 3:
```
Input:
[
  "z",
  "x",
  "z"
] 

Output: "" 
Explanation: The order is invalid, so return "".
```

Note:
- You may assume all letters are in lowercase.
- You may assume that if a is a prefix of b, then a must appear before b in 
the given dictionary.
- If the order is invalid, return an empty string.
- There may be multiple valid order of letters, return any one of them is fine.

## Explanation

This Go function, `alienOrder`, figures out the correct alphabetical order of characters in an alien language. You give it a list of `words` that are already sorted according to this alien alphabet. The function then returns a string representing that order. If it can't determine a valid order (for example, if there are contradictory rules or a character dependency loop), it returns an empty string.

---

### How it Works

The function breaks down into four main steps:

### 1. Collect Unique Characters

First, the function gathers all the unique characters from the input words.

* It initializes three maps:
    * `adj`: This will be our **adjacency list**, storing the relationships between characters (e.g., if 'a' comes before 'b', there's an edge from 'a' to 'b').
    * `inDegree`: This map tracks how many other characters *must* come before a given character. This is crucial for the sorting algorithm.
    * `exists`: A simple map to note down every character found in the words.
* It then loops through each word:
    * It **validates** the input, ensuring no empty words or characters outside the 'a' through 'z' range. If it finds any, it immediately returns an empty string.
    * Every valid character found is marked in the `exists` map.

---

### 2. Build the Graph

This is where the function establishes the rules for character order by comparing adjacent words in the provided sorted list.

* The function iterates through pairs of consecutive words (`curr` and `next`).
* **Invalid Prefix Check**: A critical rule is checked here: If a longer word (`curr`) comes *before* a shorter word (`next`) that is its prefix (e.g., "apple" before "app"), it's an impossible scenario for a sorted list, so the function returns an empty string.
* **Finding Differences**: It then compares `curr` and `next` character by character until it finds the first difference.
    * If `u` is the character from `curr` and `v` is the character from `next` at the first differing position, it means `u` **must come before** `v` in the alien alphabet.
    * An **edge is added** from `u` to `v` in the `adj` list, and the `inDegree` of `v` is incremented. This process stops after the very first differing pair is found, as that's all we need from that word pair.

---

### 3. Topological Sort (Kahn's Algorithm)

Once all character dependencies are mapped out, the function uses **Kahn's algorithm** to determine a valid linear ordering.

* **Starting Points**: It first identifies all characters that have an `inDegree` of `0` (meaning no other character needs to come before them). These are our starting points for the order, and they are collected in a slice called `zeros`. For consistent results in testing, `zeros` is then sorted alphabetically.
* **Building the Order**:
    * The function uses a `strings.Builder` to efficiently construct the final `order` string.
    * It repeatedly takes a character `u` from the `zeros` list (a character with no dependencies).
    * `u` is appended to our `order` string.
    * Then, for all characters `v` that `u` directly precedes (its "neighbors" in the graph):
        * The `inDegree` of `v` is decremented. This signifies that one of `v`'s prerequisites (`u`) has now been placed in the order.
        * If `v`'s `inDegree` drops to `0`, it means all characters that needed to come before `v` are now accounted for. So, `v` is added to the `zeros` list, making it a candidate for the next spot in the alien alphabet.

---

### 4. Cycle Detection

Finally, the function checks if a complete, valid order was found.

* After the sorting process, if the length of the `order` string is not equal to the total number of unique characters initially found (`len(exists)`), it means some characters were left with non-zero in-degrees. This indicates a **cycle** in the graph (e.g., 'a' comes before 'b', 'b' comes before 'c', and 'c' comes before 'a'). A cycle means no valid linear order exists, so the function returns an empty string.
* If all characters were successfully placed, the function returns the constructed `order` string.

In essence, `alienOrder` builds a **directed acyclic graph (DAG)** from your sorted words, where the arrows show character dependencies. Then, it uses **Kahn's algorithm** to flatten this DAG into a linear sequence, which is your alien alphabet.