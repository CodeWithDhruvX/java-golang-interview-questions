# 6️⃣ Pattern Matching & Searching (Golang Edition)

---

## 1. Implement strStr()

```
------------------------------------
| Problem Title -> strStr()        |
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
Find index of `needle` in `haystack`.

### 2️⃣ Pattern / Category ⭐
**Sliding Window / KMP**

### 3️⃣ Brute Force Idea
Nested loop. O(N*M).

### 4️⃣ Key Insight (AHA 💡)
Go's `strings.Index` is optimized (Rabin-Karp/Boyer-Moore).
For interviews, Naive is step 1. KMP is step 2 (LPS Array).

### 5️⃣ Algorithm (Naive)
1. Loop `i` 0 to N-M.
2. Loop `j` 0 to M.
3. If mismatch break.
4. If `j==M` return `i`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty needle (0).
*   Needle > Haystack (-1).

### 7️⃣ Time & Space Complexity
> **Time:** O(NM) or O(N+M)
> **Space:** O(1)

---

## 2. Longest Common Prefix

```
------------------------------------
| Problem Title -> Longest Prefix  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Find common prefix of array of strings.

### 2️⃣ Pattern / Category ⭐
**Sorting**

### 3️⃣ Brute Force Idea
Compare all.

### 4️⃣ Key Insight (AHA 💡)
Sort. Compare ONLY first and last string. They differ the most.

### 5️⃣ Algorithm
1. `sort.Strings(strs)`.
2. Compare `strs[0]` and `strs[N-1]`.
3. Return matching part.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty array.

### 7️⃣ Time & Space Complexity
> **Time:** O(N*L log N)
> **Space:** O(1)

---

## 3. Count Occurrences of Pattern

```
------------------------------------
| Problem Title -> Count Patterns  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
How many times `p` appears in `s`.

### 2️⃣ Pattern / Category ⭐
**strings.Count**

### 3️⃣ Brute Force Idea
Sliding window check.

### 4️⃣ Key Insight (AHA 💡)
Use `strings.Count(s, p)`.
Logic: Loop through `s`, match `p`, increment.

### 5️⃣ Algorithm
1. `strings.Count(s, p)`.
2. OR Manual: `count=0`. Loop find next index. `idx += len(p)`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Overlapping patterns ("aa" in "aaa" -> 1 or 2? `strings.Count` is non-overlapping i.e., 1).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)
