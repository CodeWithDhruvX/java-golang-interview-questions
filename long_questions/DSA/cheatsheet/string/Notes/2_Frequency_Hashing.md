# 2️⃣ Frequency & Hashing (Golang Edition)

---

## 1. First Non-Repeating Character

```
------------------------------------
| Problem Title -> First Unique    |
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
Find the *first* character in a string that occurs only once.

### 2️⃣ Pattern / Category ⭐
**Hash Map (Frequency Table)**

### 3️⃣ Brute Force Idea
For each char, loop rest of string to check duplicates. O(N²).

### 4️⃣ Key Insight (AHA 💡)
We need Two Passes.
1. Fill a map with counts.
2. Re-read the string (order matters!) and check who has count 1.

### 5️⃣ Algorithm
1. `freq := make(map[rune]int)`
2. Loop `char` in `s`: `freq[char]++`
3. Loop `char` in `s`:
    - If `freq[char] == 1` return `char`.
4. Return -1 or 0 (not found).

### 6️⃣ Edge Cases & Traps ⚠️
*   All duplicates (return error/0).
*   Empty string.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1) (Max 256 for ASCII)

---

## 2. Find All Duplicate Characters

```
------------------------------------
| Problem Title -> Find Dups       |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Print/Return all characters appearing > 1 time.

### 2️⃣ Pattern / Category ⭐
**Hash Map / Counting Sort**

### 3️⃣ Brute Force Idea
Nested loops. O(N²).

### 4️⃣ Key Insight (AHA 💡)
Use a map to store counts. Iterate the map to find `val > 1`.
Or sort string and check `s[i] == s[i-1]`.

### 5️⃣ Algorithm
1. `counts := map[rune]int`
2. Fill counts from `s`.
3. Iterate map keys:
    - If `val > 1`, append key to result.
4. Return result.

### 6️⃣ Edge Cases & Traps ⚠️
*   No duplicates.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N)

---

## 3. Count Frequency of Each Character

```
------------------------------------
| Problem Title -> Char Frequency  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Return a map or list of how many times each character appears.

### 2️⃣ Pattern / Category ⭐
**Hash Map**

### 3️⃣ Brute Force Idea
Iterate and count.

### 4️⃣ Key Insight (AHA 💡)
Simple map populating. This is the base for 90% of string problems.

### 5️⃣ Algorithm
1. `freq := make(map[rune]int)`
2. For `c` in `s`: `freq[c]++`
3. Return `freq`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Case sensitivity.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 4. Check if Two Strings are Anagrams

```
------------------------------------
| Problem Title -> Anagram Check   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Check if s1 and s2 use exact same characters with exact same frequencies.

### 2️⃣ Pattern / Category ⭐
**Frequency Map / Sorting**

### 3️⃣ Brute Force Idea
Sort both strings and compare. O(N log N).

### 4️⃣ Key Insight (AHA 💡)
Use one frequency map. Increment for s1, decrement for s2.
If map has all ZEROs at end -> Anagrams.
Also, if lengths differ -> False immediately.

### 5️⃣ Algorithm
1. If `len(s1) != len(s2)` return `false`.
2. `freq := make(map[rune]int)`
3. Loop `i` from 0 to N:
    - `freq[s1[i]]++`
    - `freq[s2[i]]--`
4. Loop map values:
    - If `val != 0` return `false`.
5. Return `true`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Unicode (emojis).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 5. Most Frequent Character

```
------------------------------------
| Problem Title -> Max Freq Char   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Find the char that appears most.

### 2️⃣ Pattern / Category ⭐
**Tracking Max**

### 3️⃣ Brute Force Idea
Count each.

### 4️⃣ Key Insight (AHA 💡)
Track `maxVal` and `maxChar` while populating the map (or in 2nd pass).

### 5️⃣ Algorithm
1. Populate `freq` map.
2. `maxC, maxV = 0, 0`
3. For `char, count` in `freq`:
    - If `count > maxV`:
        - `maxV = count`
        - `maxC = char`
4. Return `maxC`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Multiple chars with same max frequency (ask requirement).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 6. Remove Duplicate Characters

```
------------------------------------
| Problem Title -> Unique Chars    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
"banana" -> "ban". Keep only first occurrence of each char.

### 2️⃣ Pattern / Category ⭐
**Set (Seen Map)**

### 3️⃣ Brute Force Idea
Check if char exists in result string before appending. O(N²).

### 4️⃣ Key Insight (AHA 💡)
Use a `seen` map (boolean). Only append to result if `!seen[char]`.

### 5️⃣ Algorithm
1. `seen := make(map[rune]bool)`
2. `var sb strings.Builder`
3. For `c` in `s`:
    - If `!seen[c]`:
        - `seen[c] = true`
        - `sb.WriteRune(c)`
4. Return `sb.String()`

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty string.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 7. Print Characters with Frequencies in Sorted Order

```
------------------------------------
| Problem Title -> Sorted Freqs    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
"tree" -> "e:2, r:1, t:1" (Sorted alphabetically).

### 2️⃣ Pattern / Category ⭐
**Map + Sorting**

### 3️⃣ Brute Force Idea
N/A

### 4️⃣ Key Insight (AHA 💡)
Maps in Go are **unordered**. You MUST extract keys to a slice, sort the slice, then iterate.

### 5️⃣ Algorithm
1. Fill `freq` map.
2. `keys := []rune{}`
3. For `k` in `freq`: append `k` to `keys`.
4. `sort.Slice(keys, ...)`
5. Iterate `keys` and print `key, freq[key]`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Printing order (Ascending/Descending).

### 7️⃣ Time & Space Complexity
> **Time:** O(N log N) (Sorting keys)
> **Space:** O(N)
