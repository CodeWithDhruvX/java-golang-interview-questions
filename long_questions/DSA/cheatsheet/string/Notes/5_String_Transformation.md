# 5️⃣ String Transformation (Golang Edition)

---

## 1. Reverse Words in a Sentence

```
------------------------------------
| Problem Title -> Reverse Words   |
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
"  hello world  " -> "world hello".

### 2️⃣ Pattern / Category ⭐
**Split & Reverse**

### 3️⃣ Brute Force Idea
Split, reverse array, join.

### 4️⃣ Key Insight (AHA 💡)
Use `strings.Fields(s)` in Go to handle irregular spaces automatically.
Then simple reverse of slice.

### 5️⃣ Algorithm
1. `parts := strings.Fields(s)`.
2. Reverse `parts`.
3. `strings.Join(parts, " ")`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Multiple spaces.
*   Leading/trailing spaces.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N)

---

## 2. String Compression

```
------------------------------------
| Problem Title -> Compression     |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
"aabbccc" -> "a2b2c3". In-place.

### 2️⃣ Pattern / Category ⭐
**Two Pointers (Read/Write)**

### 3️⃣ Brute Force Idea
New string.

### 4️⃣ Key Insight (AHA 💡)
`write` pointer tracks position in new compressed string. `read` scans.
`anchor` remembers start of current char block.

### 5️⃣ Algorithm
1. `write = 0`, `anchor = 0`.
2. Loop `read`:
    - If end of block:
        - Write `chars[anchor]` to `chars[write]`.
        - If `read > anchor` (count > 1): write digits.
        - `anchor = read + 1`.
3. Return `write`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Count > 9 (multiple digits).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 3. Rotate String

```
------------------------------------
| Problem Title -> Rotate Check    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Is "cdeab" rotation of "abcde"?

### 2️⃣ Pattern / Category ⭐
**Concatenation**

### 3️⃣ Brute Force Idea
Simulate rotations.

### 4️⃣ Key Insight (AHA 💡)
A rotation is always a substring of `s + s`.

### 5️⃣ Algorithm
1. Check Lengths.
2. `strings.Contains(s + s, goal)`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N)

---

## 4. Replace Characters with Count

```
------------------------------------
| Problem Title -> Replace Count   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Replace each char with its occurrence count? (Or variations like "abb" -> "a1b2").

### 2️⃣ Pattern / Category ⭐
**Map + Build**

### 3️⃣ Brute Force Idea
Nested Loop.

### 4️⃣ Key Insight (AHA 💡)
Pre-calculate counts in Map.
Build new string.

### 5️⃣ Algorithm
1. Map `counts`.
2. Builder `sb`.
3. Loop string: `sb.WriteString(strconv.Itoa(counts[char]))`.
4. Return string.

### 6️⃣ Edge Cases & Traps ⚠️
*   Single chars.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N)
