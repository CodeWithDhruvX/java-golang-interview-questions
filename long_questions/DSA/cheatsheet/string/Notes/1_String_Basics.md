# 1️⃣ String Basics (Golang Edition)

---

## 1. Reverse a String

```
------------------------------------
| Problem Title -> Reverse String  |
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
Reverse a given string. Must handle **Unicode** characters correctly (e.g., "Hello" -> "olleH", "世" -> "世").

### 2️⃣ Pattern / Category ⭐
**Two Pointers (on rune slice)**

### 3️⃣ Brute Force Idea
Iterate string from end to start and append characters to a new string/builder. Time: O(N), Space: O(N).

### 4️⃣ Key Insight (AHA 💡)
Strings in Go are immutable UTF-8 bytes. To handle multi-byte characters (Unicode), convert to `[]rune` first. Swapping elements from ends towards center reverses it in place (on the slice).

### 5️⃣ Algorithm
1. Convert string to `[]rune`.
2. Initialize `left = 0`, `right = len(runes) - 1`.
3. While `left < right`:
    - Swap `runes[left]` and `runes[right]`.
    - `left++`, `right--`
4. Convert `[]rune` back to string and return.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty string.
*   Single character string.
*   String with Unicode characters (e.g., emojis, Chinese chars).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N) (creating rune slice)

---

## 2. Check if String is Palindrome

```
------------------------------------
| Problem Title -> Is Palindrome   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Check if a string reads the same forward and backward.

### 2️⃣ Pattern / Category ⭐
**Two Pointers**

### 3️⃣ Brute Force Idea
Reverse the string and check if `original == reversed`. Space: O(N).

### 4️⃣ Key Insight (AHA 💡)
Compare characters from both ends moving inwards. If at any point they don't match, it's not a palindrome. No need to create a new string.

### 5️⃣ Algorithm
1. Convert to `[]rune` (to verify unicode safely).
2. `left = 0`, `right = len(runes) - 1`.
3. While `left < right`:
    - If `runes[left] != runes[right]`, return `false`.
    - `left++`, `right--`
4. Return `true`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty string (usually true).
*   Single character (always true).
*   Case sensitivity ("Bob" vs "bob").

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N) (rune slice)

---

## 3. Count Vowels and Consonants

```
------------------------------------
| Problem Title -> Count Vowels    |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Count number of vowels (a, e, i, o, u) and consonants.

### 2️⃣ Pattern / Category ⭐
**Linear Dictionary Check**

### 3️⃣ Brute Force Idea
Multiple `if` statements: `if c == 'a' || c == 'e' ...`

### 4️⃣ Key Insight (AHA 💡)
Use a helper string "aeiouAEIOU" and `strings.ContainsRune` or a switch case for cleaner code. Check `unicode.IsLetter` to ignore numbers/symbols.

### 5️⃣ Algorithm
1. `vowels = 0`, `consonants = 0`.
2. Iterate `char` in string:
    - If `unicode.IsLetter(char)`:
        - If `strings.ContainsRune("aeiouAEIOU", char)`: `vowels++`
        - Else: `consonants++`
3. Return counts.

### 6️⃣ Edge Cases & Traps ⚠️
*   Non-letter characters (numbers, symbols).
*   Case sensitivity.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 4. Remove Spaces from a String

```
------------------------------------
| Problem Title -> Remove Spaces   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Remove all all whitespace characters from a string. " a b c " -> "abc".

### 2️⃣ Pattern / Category ⭐
**StringBuilder / Filter**

### 3️⃣ Brute Force Idea
String concatenation `res += char` (O(N²) in loop).

### 4️⃣ Key Insight (AHA 💡)
Use `strings.Builder` (Go) or `strings.ReplaceAll`. Since strings are immutable, we build a new one efficiently.

### 5️⃣ Algorithm
1. `var sb strings.Builder`.
2. Iterate `char` in `s`:
    - If `!unicode.IsSpace(char)`:
        - `sb.WriteRune(char)`
3. Return `sb.String()`.

### 6️⃣ Edge Cases & Traps ⚠️
*   String with only spaces.
*   Multiple types of spaces (tabs, newlines).

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N)

---

## 5. Find Length Without Using `len()`

```
------------------------------------
| Problem Title -> Manual Length   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Determine number of characters in a string without built-in `len()`.

### 2️⃣ Pattern / Category ⭐
**Iterator**

### 3️⃣ Brute Force Idea
N/A

### 4️⃣ Key Insight (AHA 💡)
A `for range` loop in Go iterates over runes (Unicode code points). Just count the iterations. Note: `len(s)` gives bytes, `utf8.RuneCountInString(s)` gives characters.

### 5️⃣ Algorithm
1. `count = 0`
2. `for _ = range s`:
    - `count++`
3. Return `count`

### 6️⃣ Edge Cases & Traps ⚠️
*   Multi-byte characters (Emoji is 1 char, but 4 bytes). Range loop handles this correctly.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 6. Convert Lowercase to Uppercase (Manual)

```
------------------------------------
| Problem Title -> To Upper/Lower  |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Convert string to uppercase without `strings.ToUpper`.

### 2️⃣ Pattern / Category ⭐
**ASCII Math / Unicode Mapping**

### 3️⃣ Brute Force Idea
Map every char manually.

### 4️⃣ Key Insight (AHA 💡)
For ASCII: 'a' is 97, 'A' is 65. Difference is 32.
`Upper = Lower - 32`.
For Unicode: Use `unicode.ToUpper(rune)`.

### 5️⃣ Algorithm
1. `runes := []rune(s)`
2. For i in runes:
    - If `runes[i] >= 'a' && runes[i] <= 'z'`:
        - `runes[i] -= 32`
3. Return `string(runes)`

### 6️⃣ Edge Cases & Traps ⚠️
*   Already uppercase.
*   Non-alphabetic chars.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(N)

---

## 7. Check Equality Without Built-in `==`

```
------------------------------------
| Problem Title -> String Equals   |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Check if `s1` equals `s2` without `s1 == s2`.

### 2️⃣ Pattern / Category ⭐
**Byte-by-Byte Compare**

### 3️⃣ Brute Force Idea
N/A

### 4️⃣ Key Insight (AHA 💡)
First check lengths. If different -> false. Then check every character at index `i`.

### 5️⃣ Algorithm
1. If `len(s1) != len(s2)` return `false`.
2. Loop `i` from 0 to `len(s1)-1`:
    - If `s1[i] != s2[i]` return `false`.
3. Return `true`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Empty strings.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)

---

## 8. Count Words in a String

```
------------------------------------
| Problem Title -> Count Words     |
------------------------------------
...
```

### 1️⃣ Problem Snapshot
Count words separated by spaces. "Hello   World" -> 2.

### 2️⃣ Pattern / Category ⭐
**State Machine / Split**

### 3️⃣ Brute Force Idea
`strings.Split(s, " ")` creates empty strings for multiple spaces.

### 4️⃣ Key Insight (AHA 💡)
Use `strings.Fields(s)` which handles multiple spaces automatically.
OR Manual: Iterate and count transitions from `space` to `non-space`.

### 5️⃣ Algorithm
1. `count = 0`, `inWord = false`.
2. For `char` in `s`:
    - If `char` is not space:
        - If `!inWord`: `count++`, `inWord = true`.
    - Else:
        - `inWord = false`.
3. Return `count`.

### 6️⃣ Edge Cases & Traps ⚠️
*   Leading/trailing spaces.
*   Multiple spaces between words.

### 7️⃣ Time & Space Complexity
> **Time:** O(N)
> **Space:** O(1)
