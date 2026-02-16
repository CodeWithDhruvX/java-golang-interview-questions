# 🧵 String Problem Notes (Golang Edition)

> **Format:** One-Page Notes for Quick Revision
> **Focus:** Indian Interview Questions (Service & Product Companies)

---

# 1️⃣ String Basics

---

## Problem: Reverse a String

### 1. Snapshot
Reverse a given string. Must handle **Unicode** characters correctly (e.g., "Hello" -> "olleH", "世" -> "世").

### 2. Pattern ⭐
**Two Pointers** (on rune slice)

### 3. Brute Force
Iterate string from end to start and append characters to a new string/builder. (Time: O(N), Space: O(N))

### 4. Key Insight 💡
Strings in Go are immutable UTF-8 bytes. To handle multi-byte characters (Unicode), convert to `[]rune` first. Swapping elements from ends towards center reverses it in place (on the slice).

### 5. Algorithm
1. Convert string to `[]rune` (to handle Unicode).
2. Initialize `left = 0`, `right = len(runes) - 1`.
3. While `left < right`:
    - Swap `runes[left]` and `runes[right]`.
    - `left++`, `right--`
4. Convert `[]rune` back to string and return.

### 6. Edge Cases ⚠️
- Empty string.
- Single character string.
- String with Unicode characters (e.g., emojis, Chinese chars).

### 7. Complexity
- **Time:** O(N) (where N is number of characters/runes)
- **Space:** O(N) (creating rune slice)

---

## Problem: Check if String is Palindrome

### 1. Snapshot
Check if a string reads the same forward and backward.

### 2. Pattern ⭐
**Two Pointers**

### 3. Brute Force
Reverse the string and check if `original == reversed`. (Space: O(N))

### 4. Key Insight 💡
We compare characters from both ends moving inwards. If at any point they don't match, it's not a palindrome. No need to create a new string.

### 5. Algorithm
1. Convert to `[]rune`.
2. Initialize `left = 0`, `right = len(runes) - 1`.
3. While `left < right`:
    - If `runes[left] != runes[right]`, return `false`.
    - `left++`, `right--`
4. Return `true`.

### 6. Edge Cases ⚠️
- Empty string (usually true).
- Single character (always true).
- Case sensitivity (ask interviewer).

### 7. Complexity
- **Time:** O(N/2) ≈ O(N)
- **Space:** O(N) (for rune slice) or O(1) if accessing bytes directly (ASCII only).

---

## Problem: Count Vowels and Consonants

### 1. Snapshot
Count number of vowels (a, e, i, o, u) and consonants in a string.

### 2. Pattern ⭐
**Linear Scanners / Set Lookup**

### 3. Brute Force
Iterate and checking multiple `if` conditions: `if char == 'a' || char == 'e' ...`

### 4. Key Insight 💡
Use a simple helper function or a switch statement for cleaner code. Handle case insensitivity by converting char to lower/upper.

### 5. Algorithm
1. Initialize `vowels = 0`, `consonants = 0`.
2. Iterate through string (as runes).
3. Check if char is letter.
4. If yes, check if it's in "aeiouAEIOU".
    - If yes, `vowels++`.
    - Else, `consonants++`.
5. Return counts.

### 6. Edge Cases ⚠️
- String with numbers/symbols (ignore them).
- Mixed case "Hello".

### 7. Complexity
- **Time:** O(N)
- **Space:** O(1)

---

# 2️⃣ Frequency & Hashing (Map Based)

---

## Problem: First Non-Repeating Character

### 1. Snapshot
Find the *first* character in a string that occurs only once.

### 2. Pattern ⭐
**Hash Map (Frequency Array)**

### 3. Brute Force
For each char, loop through rest of string to check duplicates. (Time: O(N²))

### 4. Key Insight 💡
Two passes are needed.
1. Count frequency of *every* character.
2. Scan string again: check if current char's count is 1. The first one you find is the answer.

### 5. Algorithm
1. Create a map `freq := make(map[rune]int)`.
2. **Pass 1:** Iterate string, `freq[char]++`.
3. **Pass 2:** Iterate string again.
    - If `freq[char] == 1`, return `char`.
4. If loop finishes, return "no such char".

### 6. Edge Cases ⚠️
- All characters repeat ("aabb").
- Empty string.

### 7. Complexity
- **Time:** O(N) (Two passes)
- **Space:** O(1) (Fixed alphabet size, e.g., 26 or 256)

---

## Problem: Check if Two Strings are Anagrams

### 1. Snapshot
Check if two strings contain the same characters with the same frequencies (e.g., "listen", "silent")

### 2. Pattern ⭐
**Hash Map / Sort**

### 3. Brute Force
Sort both strings and compare. (Time: O(N log N))

### 4. Key Insight 💡
Use a frequency map (or fixed array array for ASCII). Increment counts for string A, decrement for string B. All counts must be zero at the end.

### 5. Algorithm
1. If `len(A) != len(B)`, return `false`.
2. Create map or array `counts`.
3. Loop `i` from 0 to `len(A)`:
    - `counts[A[i]]++`
    - `counts[B[i]]--`
4. Loop through `counts`:
    - If any value != 0, return `false`.
5. Return `true`.

### 6. Edge Cases ⚠️
- Different lengths (FALSE immediately).
- Unicode characters (use map, not array).

### 7. Complexity
- **Time:** O(N)
- **Space:** O(1) (or O(U) where U is unique chars)

---

## Problem: Find All Duplicate Characters

### 1. Snapshot
Print all characters that appear more than once in the string.

### 2. Pattern ⭐
**Hash Map**

### 3. Brute Force
Nested loops to compare every pair. (Time: O(N²))

### 4. Key Insight 💡
Store counts in a map. Iterate the map (not the string) to print chars with `count > 1`.

### 5. Algorithm
1. Create `countMap`.
2. Iterate string: `countMap[char]++`.
3. Iterate `countMap`:
    - If `val > 1`, print `key`.

### 6. Edge Cases ⚠️
- No duplicates.
- All duplicates.

### 7. Complexity
- **Time:** O(N)
- **Space:** O(N)

---

## Problem: Most Frequent Character

### 1. Snapshot
Find the character that appears the most number of times.

### 2. Pattern ⭐
**Hash Map & Tracking Max**

### 3. Brute Force
Count each char one by one. O(N^2).

### 4. Key Insight 💡
While building the frequency map (or in a second pass), keep track of a `maxCount` and `maxChar` variable. Update whenever current `count > maxCount`.

### 5. Algorithm
1. `counts := make(map[rune]int)`
2. `maxChar`, `maxVal` = 0, 0
3. Iterate string:
    - `counts[char]++`
    - If `counts[char] > maxVal`:
        - `maxVal = counts[char]`
        - `maxChar = char`
4. Return `maxChar`.

### 6. Edge Cases ⚠️
- Multiple characters with same max frequency (clarify if return any/all).
- Empty string.

### 7. Complexity
- **Time:** O(N)
- **Space:** O(1) (Alphabet size limit)

---

# 3️⃣ Substring & Sliding Window

---

## Problem: Longest Substring Without Repeating Characters

### 1. Snapshot
Find the length of the longest substring that has all unique characters.

### 2. Pattern ⭐
**Sliding Window (Variable Size) + Hash Map**

### 3. Brute Force
Check all substrings `O(N^2)` and check uniqueness `O(N)`. Total `O(N^3)`.

### 4. Key Insight 💡
Use a generic `map` to store the **last index** where a character was seen. If we encounter a duplicate character (i.e., it's already in map and its index >= current start), we must move the `start` pointer right past the previous occurrence.

### 5. Algorithm
1. `lastSeen := make(map[rune]int)`
2. `start = 0`, `maxLength = 0`
3. Loop `i` from 0 to `len(str)`:
    - If `char` in `lastSeen` AND `lastSeen[char] >= start`:
        - `start = lastSeen[char] + 1`
    - `lastSeen[char] = i`
    - `currentLen = i - start + 1`
    - `maxLength = max(maxLength, currentLen)`
4. Return `maxLength`.

### 6. Edge Cases ⚠️
- Empty string (return 0).
- String with all non-repeating chars (return N).
- String with all same chars (return 1).

### 7. Complexity
- **Time:** O(N) (One pass)
- **Space:** O(U) (U = unique characters table size)

---

## Problem: Longest Palindromic Substring

### 1. Snapshot
Find the longest substring which is a palindrome.

### 2. Pattern ⭐
**Expand Around Center**

### 3. Brute Force
Check all substrings `O(N^2)`, verify palindrome `O(N)`. Total `O(N^3)`.

### 4. Key Insight 💡
A palindrome mirrors around a center. The center can be a character (e.g., "aba" center 'b') or a gap between characters (e.g., "abba" center gap). There are `2N-1` such centers. Expanding from each center takes linear time.

### 5. Algorithm
1. `start = 0`, `end = 0` (to track max palindrome indices).
2. Loop `i` from 0 to `len`:
    - `len1 = expandAround(i, i)` (Odd length)
    - `len2 = expandAround(i, i+1)` (Even length)
    - `len = max(len1, len2)`
    - If `len > (end - start)`:
        - Update `start` and `end` based on `i` and `len`.
3. Return substring from `start` to `end`.

### 6. Edge Cases ⚠️
- Single character.
- Entire string is palindrome.
- No palindrome > 1 char (return first char).

### 7. Complexity
- **Time:** O(N²) (Better than brute force O(N^3))
- **Space:** O(1)

---

## Problem: Find All Anagrams in a String

### 1. Snapshot
Find all start indices in string `s` where the substring is an anagram of string `p`.

### 2. Pattern ⭐
**Sliding Window (Fixed Size)**

### 3. Brute Force
Take every substring of length `p` from `s`, sort it, and compare with sorted `p`. (Time: O(N * M log M))

### 4. Key Insight 💡
Maintain a window of size `len(p)` on `s`. Use a frequency array (hash). When sliding the window one step right:
- Add new character (increment freq)
- Remove old character (decrement freq)
- Compare partial freq array with `p`'s freq array.

### 5. Algorithm
1. Create `pFreq` map/array.
2. Initialize `windowFreq`. Populate for first `len(p)` chars.
3. If match, add index 0 to results.
4. Loop `i` from `len(p)` to `len(s)`:
    - Add `s[i]`, remove `s[i - len(p)]` in `windowFreq`.
    - If `windowFreq == pFreq`, append `(i - len(p) + 1)` to results.
5. Return results.

### 6. Edge Cases ⚠️
- `len(s) < len(p)` (return empty).
- No anagrams found.

### 7. Complexity
- **Time:** O(N) (Comparing 26-size array is O(1))
- **Space:** O(1)

---

## Problem: Smallest Window Containing All Characters

### 1. Snapshot
Find the smallest substring in `s` that contains all characters of `t` (including duplicates).

### 2. Pattern ⭐
**Sliding Window (Variable) + Two Hash Maps**

### 3. Brute Force
Check all substrings.

### 4. Key Insight 💡
1. Expand `right` until window is "valid" (contains all chars of `t`).
2. Once valid, increment `left` to shrink window (making it smaller) while it remains valid. Always track the minimum size seen.

### 5. Algorithm
1. Count chars of `t` in `tFreq`. `required = len(tFreq)`.
2. `left = 0`, `formed = 0`.
3. Loop `right` from 0 to `len(s)`:
    - Add `s[right]` to `windowFreq`.
    - If `s[right]` condition satisfies `tFreq`, `formed++`.
    - While `formed == required` (Valid Window):
        - Update `minLen` and `minStart` if current is smaller.
        - Remove `s[left]` from `windowFreq`, update `formed` if needed.
        - `left++`
4. Return `minSubString`.

### 6. Edge Cases ⚠️
- `t` strictly larger than `s`.
- Character in `t` doesn't exist in `s`.

### 7. Complexity
- **Time:** O(N + M)
- **Space:** O(1) (Alphabet size)

---

# 4️⃣ Palindrome-Based Problems

---

## Problem: Count Palindromic Substrings

### 1. Snapshot
Count how many contiguous substrings are palindromes.

### 2. Pattern ⭐
**Expand Around Center**

### 3. Brute Force
O(N^3).

### 4. Key Insight 💡
Same logic as "Longest Palindromic Substring", but instead of finding max, we increment a counter every time we successfully expand.

### 5. Algorithm
1. `count = 0`
2. Loop `i` from 0 to `len`:
    - `count += countExpand(i, i)` (Odd centers)
    - `count += countExpand(i, i+1)` (Even centers)
3. Return `count`.
- `countExpand(left, right)`:
    - While `left >= 0` and `right < len` and `s[left] == s[right]`:
        - `res++`, `left--`, `right++`
    - Return `res`

### 6. Edge Cases ⚠️
- Single char string = 1.
- All same chars "aaa" = 6 (a, a, a, aa, aa, aaa).

### 7. Complexity
- **Time:** O(N²)
- **Space:** O(1)

---

## Problem: Valid Palindrome II (One Deletion)

### 1. Snapshot
Return true if string can be palindrome after deleting **at most one** character.

### 2. Pattern ⭐
**Two Pointers + Greedy**

### 3. Brute Force
Try deleting every character and check if remaining is palindrome. O(N^2).

### 4. Key Insight 💡
Use standard two pointers. If `s[left] != s[right]`, we have two choices:
1. Delete `s[left]` (check if `s[left+1...right]` is palindrome).
2. Delete `s[right]` (check if `s[left...right-1]` is palindrome).
If either is true, return true.

### 5. Algorithm
1. `left = 0`, `right = len - 1`.
2. While `left < right`:
    - If `s[left] != s[right]`:
        - Return `isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)`
    - `left++`, `right--`
3. Return `true`.

### 6. Edge Cases ⚠️
- Already a palindrome.
- Deleting middle char.

### 7. Complexity
- **Time:** O(N)
- **Space:** O(1)

---

# 5️⃣ String Transformation & Manipulation

---

## Problem: Reverse Words in a Sentence

### 1. Snapshot
Reverse the order of words in a string. Remove extra spaces. "  hello world  " -> "world hello".

### 2. Pattern ⭐
**Built-in Split/Join OR Two-Pass Reverse**

### 3. Brute Force
Split by space, reverse array, join. (Easy in high-level langs).

### 4. Key Insight 💡
**In-place approach (O(1) extra space):**
1. Reverse the **entire string**.
2. Reverse each **individual word**.
3. Clean up spaces.

### 5. Algorithm (Go Logic)
1. `parts := strings.Fields(s)` (Handles multiple spaces automatically).
2. Reverse the `parts` slice (Two pointers).
3. `return strings.Join(parts, " ")`.

### 6. Edge Cases ⚠️
- Leading/trailing spaces.
- Multiple spaces between words.
- Single word.

### 7. Complexity
- **Time:** O(N)
- **Space:** O(N) (For parts slice)

---

## Problem: String Compression

### 1. Snapshot
Compress string by replacing repeated chars with count. "aabbccc" -> "a2b2c3". If compressed is longer, return original (or as per problem variation).

### 2. Pattern ⭐
**Two Pointers (Read & Write)**

### 3. Brute Force
Use a new string builder.

### 4. Key Insight 💡
Use two pointers:
- `read`: iterates through the string.
- `write`: overwrites the original array (if in-place) or appends to result.
- `anchor`: marks start of current group of identical characters.

### 5. Algorithm
1. `write = 0`, `anchor = 0`.
2. Loop `read` from 0 to `len`:
    - If `read + 1 == len` OR `s[read + 1] != s[read]`:
        - Write `s[anchor]` to `s[write]`. `write++`.
        - If `read > anchor` (count > 1):
            - Convert count to string/chars. Write each digit to `s[write]`.
        - `anchor = read + 1`.
3. Return `write` (new length).

### 6. Edge Cases ⚠️
- Single characters (no count written).
- Count >= 10 (multiple digits).

### 7. Complexity
- **Time:** O(N)
- **Space:** O(1)

---

## Problem: Rotate String

### 1. Snapshot
Check if string `A` is a rotation of `B` (e.g., "abcde", "cdeab").

### 2. Pattern ⭐
**String Concatenation**

### 3. Brute Force
Generate all rotations of A and check equality with B.

### 4. Key Insight 💡
If `A` is a rotation of `B`, then `A` must be a substring of `B + B`.
(e.g., `cdeab` is in `abcdeabcde`).

### 5. Algorithm
1. If `len(A) != len(B)`, return `false`.
2. `doubled := B + B`.
3. Check `strings.Contains(doubled, A)`.

### 6. Edge Cases ⚠️
- Different lengths.
- Empty strings.

### 7. Complexity
- **Time:** O(N) (Depends on `Contains` implementation, usually efficient).
- **Space:** O(N) (For `B+B`).

---

# 6️⃣ Pattern Matching & Searching

---

## Problem: Implement strStr() (Find Substring)

### 1. Snapshot
Find the index of the first occurrence of needle in haystack.

### 2. Pattern ⭐
**Sliding Window / KMP**

### 3. Brute Force
Nested loops: Check every position in haystack. O(N*M).

### 4. Key Insight 💡
For interviews, O(N*M) is often the first step (Naive). For optimization, mention **KMP** (Knuth-Morris-Pratt).
**KMP Insight:** When mismatch happens, use the "LPS Array" (Longest Prefix Suffix) to skip unnecessary comparisons in the pattern.

### 5. Algorithm (Naive - often sufficient unless specified)
1. Loop `i` from 0 to `len(hay)-len(needle)`.
2. Loop `j` from 0 to `len(needle)`:
    - If `hay[i+j] != needle[j]`: break.
    - If `j == len(needle)-1`: return `i`.
3. Return -1.

### 6. Edge Cases ⚠️
- Empty needle (return 0).
- Needle longer than haystack.

### 7. Complexity
- **Time:** O(N*M) (Naive), O(N+M) (KMP).

---

## Problem: Longest Common Prefix

### 1. Snapshot
Find the longest string that is a prefix of all strings in an array.

### 2. Pattern ⭐
**Horizontal Scanning / Sorting**

### 3. Brute Force
Compare character by character for all strings.

### 4. Key Insight 💡
**Sorting Trick:** Sort the array of strings. The longest common prefix must be a prefix of the **first** and the **last** string (since they are most different).

### 5. Algorithm
1. Sort the string array `strs`. (Lexicographically).
2. Take `s1 = strs[0]`, `s2 = strs[len-1]`.
3. Loop `i` while `i < len(s1)` and `i < len(s2)`:
    - If `s1[i] == s2[i]`: continue.
    - Else: break.
4. Return `s1[0:i]`.

### 6. Edge Cases ⚠️
- Empty array.
- No common prefix.

### 7. Complexity
- **Time:** O(N*L*log(N)) (Sorting takes dominant time).
- **Space:** O(1) or O(L).

---

# 7️⃣ Tricky & Go-Specific (Important!)

---

## Problem: Group Anagrams

### 1. Snapshot
Group an array of strings into anagrams. `["eat", "tea", "tan", "ate", "nat", "bat"]` -> `[["eat","tea","ate"], ...]`.

### 2. Pattern ⭐
**Hash Map with Sorted Key**

### 3. Brute Force
Compare every string with every other string. O(N^2 * M log M).

### 4. Key Insight 💡
Anagrams share the same "signature".
**Option A (Sorting):** Sort each string ("eat" -> "aet"). Use "aet" as key in map.
**Option B (Frequency Count):** Use an array `[26]int` as key (Go allows arrays as map keys!).

### 5. Algorithm (Option B - Faster)
1. `groups := make(map[[26]int][]string)`
2. For each string `s`:
    - Create count array `key` of size 26.
    - Count chars of `s` into `key`.
    - `groups[key] = append(groups[key], s)`
3. Collect values from map.

### 6. Edge Cases ⚠️
- Empty strings.
- Unicode (use sorting method or larger array).

### 7. Complexity
- **Time:** O(N * M) (No sorting needed).
- **Space:** O(N * M).

---

## Concept: Strings vs []byte in Go 🐹

### 1. Snapshot
Why does Go have both? When to use which?

### 2. Pattern ⭐
**Immutability vs Mutability**

### 3. Explanation
- **`string`**: Read-only (Immutable). Underlying array cannot be changed. Cheap to copy (just pointer + length).
- **`[]byte`**: Mutable slice of bytes. Can be modified in place.

### 4. Key Insight 💡
- Use **`string`** for map keys, struct fields, and function arguments (safe).
- Use **`[]byte`** when you need to manipulate data (e.g., IO buffers, decoding, in-place string modification).

### 5. Conversion Cost
- `string(bytes)` and `[]byte(str)` usually cause **memory allocation** and **copying**.
- **Optimization:** In rare hot paths, `unsafe` can cast without copy (but be careful!).

---

## Concept: Efficient String Concatenation

### 1. Snapshot
How to concatenate strings efficiently in a loop?

### 2. Pattern ⭐
**`strings.Builder`**

### 3. Anti-Pattern ❌
`s += "next"` inside a loop.
**Why bad?** Strings are immutable. Every `+=` allocates a NEW string and copies old content. O(N^2).

### 4. Key Insight 💡
Use **`strings.Builder`**. It maintains a growing buffer and only allocates the final string once when `.String()` is called.

### 5. Usage
1. `var sb strings.Builder`
2. `sb.Grow(n)` (Optional optimization if size known).
3. `sb.WriteString("hello")`
4. `res := sb.String()`

### 6. Complexity
- **Time:** O(N)
- **Space:** O(N)

---

## Concept: Rune vs Byte

### 1. Snapshot
What is a `rune`?

### 2. Pattern ⭐
**Unicode Code Point**

### 3. Explanation
- **`byte`**: `uint8` (Alias). Represents ASCII char (1 byte).
- **`rune`**: `int32` (Alias). Represents a Unicode Code Point (can be 1-4 bytes in UTF-8).

### 4. Key Interview Check 💡
- `len("你好")` is **6** (3 bytes each).
- `len([]rune("你好"))` is **2**.
- **Always** convert to `[]rune` if iterating over user input that might contain Emoji or non-English text.

---

# 🚀 Quick Practice Strategy
1. **Read Pattern** for 5 problems (2 mins).
2. **Visualize Algorithm** mentally.
3. **Write Code** for 1 complex problem (e.g., Longest Substring).
4. **Explain Complexity** out loud.

**You got this!** 👊
