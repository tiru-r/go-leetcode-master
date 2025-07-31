# 76. Minimum Window Substring

https://leetcode.com/problems/minimum-window-substring/

## Difficulty:

Hard

## Description

Given a string S and a string T, find the minimum window in S 
which will contain all the characters in T in complexity O(n).

Example:
```
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
```

Note:
- If there is no such window in S that covers all characters in T, return the empty string "".
- If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

## Algorithm Explanation

**Goal**: Find the shortest substring in `s` that contains all characters from `t`.

**Strategy**: Use a sliding window that expands and contracts to find the minimum valid window.

### 🎯 Simple Concept
Think of it like a rubber band that stretches and shrinks:
1. **Stretch** the band (expand window) until it covers all required characters
2. **Shrink** the band (contract window) while keeping all characters covered
3. Remember the smallest band size that worked

### 📊 What We Track
```
Given: s = "ADOBECODEBANC", t = "ABC"

need:     A:1, B:1, C:1     ← How many of each char we need
have:     A:0, B:0, C:0     ← How many we currently have in window
required: 3                 ← Total unique chars needed (A, B, C)
formed:   0                 ← How many requirements we've satisfied
```

### 🔄 Complete Example Walkthrough

**Starting Setup:**
```
s = "ADOBECODEBANC"
t = "ABC"

need: A=1, B=1, C=1
required = 3 (we need A, B, and C)
```

**Step-by-Step Process:**
```
Step 1: A D O B E C O D E B A N C    Window: "A"
        ↑
       L,R    have: A=1  formed: 1  (need A✓, need B✗, need C✗)

Step 2: A D O B E C O D E B A N C    Window: "AD" 
        ↑   ↑
        L   R    have: A=1  formed: 1  (D not needed)

Step 3: A D O B E C O D E B A N C    Window: "ADO"
        ↑     ↑
        L     R    have: A=1  formed: 1  (O not needed)

Step 4: A D O B E C O D E B A N C    Window: "ADOB"
        ↑       ↑
        L       R    have: A=1, B=1  formed: 2  (need A✓, need B✓, need C✗)

Step 5: A D O B E C O D E B A N C    Window: "ADOBE"
        ↑         ↑
        L         R    have: A=1, B=1  formed: 2  (E not needed)

Step 6: A D O B E C O D E B A N C    Window: "ADOBEC"  ← FIRST VALID!
        ↑           ↑
        L           R    have: A=1, B=1, C=1  formed: 3  ✅ ALL SATISFIED!
                         minWindow = "ADOBEC" (length 6)
```

**Now Contract the Window (try to make it smaller):**
```
Step 7: A D O B E C O D E B A N C    Window: "DOBEC"
          ↑         ↑
          L         R    have: A=0, B=1, C=1  formed: 2  ❌ Lost A!
                         Window invalid, need to expand again...

Continue expanding right pointer until we find A again...

Step 10: A D O B E C O D E B A N C   Window: "DOBECODEBANC"
          ↑                   ↑
          L                   R   have: A=1, B=2, C=2  formed: 3  ✅ Valid again!
                              But length=10 > 6, so not better

Contract again...

Step 11: A D O B E C O D E B A N C   Window: "BANC"  ← BETTER!
                          ↑   ↑
                          L   R   have: A=1, B=1, C=1  formed: 3  ✅
                              minWindow = "BANC" (length 4)  🎉
```

### 🎮 Algorithm Flow
```
START
  ↓
┌─────────────────────┐
│  Setup counters     │
│  need[], have[]     │
│  left=0, right=0    │
└──────────┬──────────┘
           ↓
    ┌──────────────────┐ ← Main outer loop
    │  right < len(s)? │
    └──────┬───────────┘
           │ YES
           ↓
    ┌──────────────────┐
    │ Add s[right]     │ ← Expand window
    │ have[s[right]]++ │
    │ Update formed    │
    └──────┬───────────┘
           ↓
    ┌──────────────────┐ ← Inner contraction loop
    │ formed==required │
    │ AND left<=right? │
    └──────┬───────────┘
           │ YES
           ↓
    ┌──────────────────┐
 ┌─►│ Update minimum   │ ← Contract while valid
 │  │ Remove s[left]   │
 │  │ have[s[left]]--  │
 │  │ Update formed    │
 │  │ left++           │
 │  └──────┬───────────┘
 │         ↓
 │  ┌──────────────────┐
 │  │ Still valid?     │
 │  │ formed==required │
 │  │ AND left<=right? │
 │  └──────┬───────────┘
 │         │ YES
 └─────────┘
           │ NO (window invalid or left > right)
           ↓
    ┌──────────────────┐
    │ right++          │ ← Move to next character
    └──────┬───────────┘
           ↓
    Back to "right < len(s)?" check
           │
           ↓ NO (right >= len(s))
    ┌──────────────────┐
    │      END         │
    │ Return result    │
    │ (empty if none)  │
    └──────────────────┘
```

### ✅ Key Rules
1. **Window is VALID** when `formed == required`
   - This means we have enough of every character from `t`

2. **Expand Phase**: Move `right` pointer to include more characters
   - Add characters until window becomes valid

3. **Contract Phase**: Move `left` pointer to exclude characters  
   - Remove characters while window stays valid
   - Update minimum if current window is smaller

4. **Character Counting**:
   - `have[char]++` when adding a character
   - `have[char]--` when removing a character
   - `formed++` when `have[char]` reaches `need[char]`
   - `formed--` when `have[char]` drops below `need[char]`

### 🏁 Final Result
For `s = "ADOBECODEBANC"` and `t = "ABC"`:
- Answer: `"BANC"` (length 4)
- This is the shortest substring containing all A, B, C

### Complexity
- **Time**: O(|s| + |t|) - each character visited at most twice
- **Space**: O(1) - fixed 128-character arrays for ASCII
