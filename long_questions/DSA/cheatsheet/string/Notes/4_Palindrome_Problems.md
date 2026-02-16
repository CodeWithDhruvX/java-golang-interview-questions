# 4️⃣ Palindrome-Based Problems (Golang Edition)

---

## 1. Longest Palindromic Substring

```
------------------------------------
| Problem Title -> Longest Palin   |
------------------------------------
| 1. Problem Snapshot              |
| 2. Pattern / Category ⭐          |
| 3. Brute Force Idea              |
| 4. Key Insight (AHA) 💡           |
| 5. Algorithm (Steps)             |
| 6. Edge Cases & Traps ⚠️          |
| 7. Complexity                    |
------------------------------------
```

### 1️⃣ Problem Snapshot
Find longest substring that is a palindrome.

### 2️⃣ Pattern / Category ⭐
**Expand Around Center**

### 3️⃣ Brute Force Idea
All substrings + check. O(N³).

### 4️⃣ Key Insight (AHA 💡)
A palindrome mirrors around a center. `2N-1` centers (letters + gaps between).
Expand left/right from each center.

### 5️⃣ Algorithm
1. `maxLen=0`.
2. Loop `i` from 0 to N:
    - `len1 = expand(i, i)` (Odd)
    - `len2 = expand(i, i+1)` (Even)
    - Update max.
3. Return substring.

### 6️⃣ Edge Cases & Traps ⚠️
*   Entire string is palindrome.
*   No palindrome > 1.

### 7️⃣ Time & Space Complexity
> **Time:** O(N²)
> **Space:** O(1)

---

## 2. Count Palindromic Substrings

```
------------------------------------
| Problem Title -> Count Palins    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Total count of palindromic substrings.

### 2️⃣ Pattern / Category ⭐
**Expand Around Center**

### 3️⃣ Brute Force Idea
Check all. O(N³).

### 4️⃣ Key Insight (AHA 💡)
Same as longest, but instead of tracking max, just `count++` for every valid expansion step.

### 5️⃣ Algorithm
1. `ans = 0`.
2. Loop `i`:
    - `ans += countExpand(i, i)`
    - `ans += countExpand(i, i+1)`
3. Return `ans`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Single chars count as 1.

### 7️⃣ Time & Space Complexity
> **Time:** O(N²)
> **Space:** O(1)

---

## 3. Valid Palindrome II (One Deletion)

```
------------------------------------
| Problem Title -> Palin deletion  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Can we make it palindrome by deleting AT MOST 1 char?

### 2️⃣ Pattern / Category ⭐
**Two Pointers + Greedy**

### 3️⃣ Brute Force Idea
Delete every char and check. O(N²).

### 4️⃣ Key Insight (AHA 💡)
Match from outside in.
If `s[L] != s[R]`:
We MUST delete either `L` or `R`.
Check if `s[L+1...R]` OR `s[L...R-1]` is palindrome. If yes -> True.

### 5️⃣ Algorithm
1. `L=0`, `R=N-1`.
2. While `L < R`:
    - If mismatch:
        - Return `isPalin(L+1, R) || isPalin(L, R-1)`
    - `L++`, `R--`
3. Return `true`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Already palindrome.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 4. Minimum Deletions to Make Palindrome

```
------------------------------------
| Problem Title -> Min Deletions   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Minimum chars to remove to make string a palindrome.

### 2️⃣ Pattern / Category ⭐
**LCS (Longest Common Subsequence)**

### 3️⃣ Brute Force Idea
Recursion. Exponential.

### 4️⃣ Key Insight (AHA 💡)
Problem is equivalent to: `N - LongestPalindromicSubsequence`.
LPS is `LCS(s, reverse(s))`.

### 5️⃣ Algorithm
1. `rev = reverse(s)`.
2. `l = LCS(s, rev)`.
3. Return `len(s) - l`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty string.

### 7️⃣ Time & Space Complexity
> **Time:** O(N²) (DP)
> **Space:** O(N²)
