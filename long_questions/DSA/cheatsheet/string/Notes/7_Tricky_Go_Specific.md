# 7️⃣ Tricky & Go-Specific (Golang Edition)

---

## 1. Group Anagrams

```
------------------------------------
| Problem Title -> Group Anagrams  |
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
Group strings that are anagrams. `["eat", "tea"]`.

### 2️⃣ Pattern / Category ⭐
**Map with Array Key**

### 3️⃣ Brute Force Idea
Compare all pairs.

### 4️⃣ Key Insight (AHA 💡)
In Go, `[26]int` (array) can be a map key! Slice `[]int` cannot.
Count chars of each string into an array, use that array as the key.

### 5️⃣ Algorithm
1. `m := make(map[[26]int][]string)`
2. For `s` in `strs`:
    - `count = [26]int`
    - Fill count.
    - `m[count] = append(m[count], s)`
3. Return map values.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty input.

### 7️⃣ Time & Space Complexity
> **Time:** O(N * L)
> **Space:** O(N * L)

---

## 2. Print All Permutations of a String

```
------------------------------------
| Problem Title -> Permutations    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
"ABC" -> "ABC", "ACB", "BAC"...

### 2️⃣ Pattern / Category ⭐
**Backtracking**

### 3️⃣ Brute Force Idea
Recursive swapping. O(N!).

### 4️⃣ Key Insight (AHA 💡)
Swap `s[i]` with `s[start]`. Recurse. Swap back (backtrack).

### 5️⃣ Algorithm
1. `func permute(arr, start)`:
    - If `start == len`: print.
    - Loop `i` from `start` to len:
        - Swap `arr[start], arr[i]`
        - `permute(arr, start+1)`
        - Swap back.

### 6️⃣ Edge Cases & Traps ⚠️
*   Duplicates (need Set).

### 7️⃣ Time & Space Complexity
> **Time:** O(N * N!)
> **Space:** O(N)

---

## 3. String vs []byte vs rune

```
------------------------------------
| Problem Title -> Go Types        |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Difference?

### 2️⃣ Pattern / Category ⭐
**Concept**

### 3️⃣ Brute Force Idea
N/A

### 4️⃣ Key Insight (AHA 💡)
*   **string**: Immutable, UTF-8.
*   **[]byte**: Mutable, raw bytes (ASCII).
*   **rune**: Unicode Code Point (int32).

### 5️⃣ Algorithm
*   Use `string` for keys/passing.
*   Use `[]byte` for IO/modification.
*   Use `[]rune` for reversing/iterating chars.

### 6️⃣ Edge Cases & Traps ⚠️
*   Conversions create copies (O(N)).

### 7️⃣ Time & Space Complexity
> **Time:** Conversion O(N)
> **Space:** O(N)
