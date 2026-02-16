# 3️⃣ Substring & Sliding Window (Golang Edition)

---

## 1. Longest Substring Without Repeating Characters

```
------------------------------------
| Problem Title -> Longest Unique  |
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
Find the length of the longest substring with all unique chars.

### 2️⃣ Pattern / Category ⭐
**Sliding Window (Variable)**

### 3️⃣ Brute Force Idea
All substrings O(N²), check uniqueness O(N) -> O(N³).

### 4️⃣ Key Insight (AHA 💡)
Map stores `last_index` of char. If we see a repeat (`s[i]` in map), jump `start` to `map[s[i]] + 1`. This skips the window past the duplicate instantly.

### 5️⃣ Algorithm
1. `lastSeen := map[rune]int`
2. `start = 0`, `maxLen = 0`
3. Loop `i` from 0 to N:
    - If `s[i]` in map AND `map[s[i]] >= start`:
        - `start = map[s[i]] + 1`
    - `lastSeen[s[i]] = i`
    - `maxLen = max(maxLen, i - start + 1)`
4. Return `maxLen`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty string (0).
*   All same characters ("aaaa" -> 1).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 2. Find All Anagrams of a Pattern

```
------------------------------------
| Problem Title -> All Anagrams    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Find start indices of all substrings in `s` that are anagrams of `p`.

### 2️⃣ Pattern / Category ⭐
**Sliding Window (Fixed Size)**

### 3️⃣ Brute Force Idea
Sort every window. O(N * K log K).

### 4️⃣ Key Insight (AHA 💡)
Maintain a `windowFreq` array. Slide window right: add new char, remove old char. Compare `windowFreq` vs `pFreq` (array comparison is O(1) for size 26).

### 5️⃣ Algorithm
1. Count `pFreq`. Init `windowFreq` with first `len(p)` chars.
2. Compare.
3. Loop `i` from `len(p)` to `len(s)`:
    - `windowFreq[s[i]]++`
    - `windowFreq[s[i-len(p)]]--`
    - Compare arrays. If equal, store index.
4. Return indices.

### 6️⃣ Edge Cases & Traps ⚠️
*   `len(s) < len(p)`.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 3. Smallest Window Containing All Characters

```
------------------------------------
| Problem Title -> Min Window Sub  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Find smallest substring in `s` containing all chars of `t`.

### 2️⃣ Pattern / Category ⭐
**Sliding Window (Expand & Shrink)**

### 3️⃣ Brute Force Idea
Check all substrings.

### 4️⃣ Key Insight (AHA 💡)
1. Expand `right` until valid (window has all `t`).
2. Shrink `left` to minimize size while keeping it valid. Global Min tracks answer.

### 5️⃣ Algorithm
1. `need` map from `t`. `have` map.
2. `count = 0`, `req = len(need)`.
3. Loop `right`:
    - Add `s[right]` to `have`. If `have[c] == need[c]`, `count++`.
    - While `count == req` (Valid):
        - Update Min Result.
        - Remove `s[left]`. If `have[c] < need[c]`, `count--`.
        - `left++`.
4. Return Min Result.

### 6️⃣ Edge Cases & Traps ⚠️
*   `t` longer than `s`.
*   No solution.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 4. Count Substrings with All Distinct Characters

```
------------------------------------
| Problem Title -> Count Distinct  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Count how many substrings have no repeating characters.

### 2️⃣ Pattern / Category ⭐
**Sliding Window**

### 3️⃣ Brute Force Idea
Check all O(N²).

### 4️⃣ Key Insight (AHA 💡)
If `s[start...end]` has unique chars, then all substrings ending at `end` (`s[start...end]`, `s[start+1...end]`, ...) are also unique!.
Count added at each step `end` is `end - start + 1`.

### 5️⃣ Algorithm
1. `lastSeen := map`, `start = 0`, `count = 0`.
2. Loop `end` from 0 to N:
    - If `s[end]` seen, `start = max(start, lastSeen[s[end]] + 1)`.
    - `lastSeen[s[end]] = end`.
    - `count += (end - start + 1)`.
3. Return `count`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Single char.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 5. Check if String Contains All Characters of Another (Subsequence vs Subset)

```
------------------------------------
| Problem Title -> Contains All    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Does `s` contain all characters of `t`? (Order doesn't matter = Subset).

### 2️⃣ Pattern / Category ⭐
**Frequency Map**

### 3️⃣ Brute Force Idea
N/A

### 4️⃣ Key Insight (AHA 💡)
Just count freqs of `s`. Check if for every char in `t`, `freq_s[c] >= freq_t[c]`.

### 5️⃣ Algorithm
1. Map `s`.
2. Iterate `t`, decrement map.
3. If value < 0 -> False.

### 6️⃣ Edge Cases & Traps ⚠️
*   Duplicates in `t`.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)
