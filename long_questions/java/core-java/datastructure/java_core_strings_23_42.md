## 2️⃣ Strings (Java String & Related Classes)

### Question 23: Difference between `String`, `StringBuilder`, and `StringBuffer`.

**Answer:**
- **`String`**: Immutable. Stored in String Pool. Slow for concatenations (creates new objects). Thread-safe (due to immutability).
- **`StringBuilder`**: Mutable. Fast. **Not** Thread-safe. Use for string manipulation in single thread.
- **`StringBuffer`**: Mutable. Slower than Builder (has synchronized methods). **Thread-safe**. Legacy.

---

### Question 24: `String` immutability — why is it important?

**Answer:**
1.  **Security:** String params (DB URL, passwords) can't be changed after creation.
2.  **Concurrency:** Safe to share across threads without locks.
3.  **Caching:** HashCode is cached (great for HashMap keys).
4.  **String Pool:** Saves heap memory by reusing literals.

---

### Question 25: How do you reverse a string?

**Answer:**
1.  `new StringBuilder(str).reverse().toString()`.
2.  Convert to `char[]`, swap elements (Two-pointer), create new String.

---

### Question 26: How do you check if a string is a palindrome?

**Answer:**
1.  **StringBuilder:** `str.equals(new StringBuilder(str).reverse().toString())`.
2.  **Two Pointers:** Check `charAt(start) == charAt(end)` while `start < end`. If mismatch, return false.

---

### Question 27: How do you remove duplicate characters from a string?

**Answer:**
1.  **LinkedHashSet:** Ensures unique + Insertion Order. Loop chars, add to Set, build string.
2.  **Stream:** `str.chars().distinct()...`.
3.  **boolean[] seen:** Iterate, check if `seen[char]`, if not append and mark true.

---

### Question 28: How do you count vowels/consonants in a string?

**Answer:**
Iterate `char` array.
Check if `ch` is in "aeiouAEIOU". Increment `vowels`, else `consonants` (if it's a letter).

---

### Question 29: How do you check anagrams of two strings?

**Answer:**
Two strings are anagrams if they contain same characters with same frequencies.
1.  **Sort:** `Arrays.sort(arr1); Arrays.sort(arr2); return Arrays.equals(arr1, arr2);`. Time: O(N log N).
2.  **Frequency Array:** Count char frequencies of A (++), then B (--). Check if all counts are 0. Time: O(N).

---

### Question 30: How do you find the first non-repeating character?

**Answer:**
1.  **Map:** Store counts in `LinkedHashMap<Character, Integer>`. Iterate map, find first with count 1.
2.  **Array:** Frequency array (size 256 for ASCII). First pass count. Second pass (on string) check count.

---

### Question 31: How do you replace characters or substrings in Java?

**Answer:**
- `str.replace('a', 'b')` (Char replacement).
- `str.replace("foo", "bar")` (Exact substring replacement).
- `str.replaceAll("\\d", "#")` (Regex replacement).

---

### Question 32: How do you split a string into an array of substrings?

**Answer:**
`str.split(regex)`.
`"a,b,c".split(",")` -> `["a", "b", "c"]`.
Note: Takes a **Regex**. To split by dot, use `split("\\.")`.

---

### Question 33: Difference between `equals()` and `equalsIgnoreCase()`.

**Answer:**
- `equals()`: Case sensitive. "Java" != "java".
- `equalsIgnoreCase()`: Case insensitive.

---

### Question 34: Difference between `==` and `equals()` for strings.

**Answer:**
- **`==`**: Checks reference (Do they point to same String Pool object?).
- **`equals()`**: Checks character content.
Always use `equals` for logical comparison.

---

### Question 35: `substring()` vs `subSequence()`.

**Answer:**
- **`substring(start, end)`**: Returns a `String`.
- **`subSequence(start, end)`**: Returns `CharSequence`. (Interface implemented by String, StringBuilder). Used for generalization.

---

### Question 36: `charAt()`, `indexOf()`, `lastIndexOf()` — use cases.

**Answer:**
- `charAt(i)`: Get char at index.
- `indexOf(char/str)`: Find first occurrence index. (Returns -1 if not found).
- `lastIndexOf()`: Find last occurrence.

---

### Question 37: `startsWith()`, `endsWith()`, `contains()` — examples.

**Answer:**
Return `boolean`.
`"hello".startsWith("he")` (true).
`"file.txt".endsWith(".txt")` (true).
`"hello".contains("ll")` (true).
These avoids writing manual loops.

---

### Question 38: `trim()`, `strip()`, `stripLeading()`, `stripTrailing()` differences.

**Answer:**
- `trim()`: Removes whitespace (ASCII <= 32). Legacy.
- `strip()`: (Java 11+) Unicode aware. Removes all Unicode whitespace. Recommended.
- `stripLeading()/Trailing()`: Removes only from one side.

---

### Question 39: `replace()` vs `replaceAll()` vs `replaceFirst()`.

**Answer:**
- `replace(target, replacement)`: Replaces **All** occurrences of **literal** string.
- `replaceAll(regex, replacement)`: Replaces **All** matches of **regex**.
- `replaceFirst(regex, replacement)`: Replaces **First** match of regex.

---

### Question 40: `matches()` and regular expressions for validation.

**Answer:**
`str.matches(regex)` returns true if **entire** string matches regex.
Example: Email validation, Phone validation.
Pattern is compiled every time (slow for tight loops, use compiled Pattern instead).

---

### Question 41: Converting string to char array and vice versa.

**Answer:**
- String to Array: `char[] arr = str.toCharArray();`
- Array to String: `String s = new String(arr);` or `String.valueOf(arr);`

---

### Question 42: Converting string to uppercase/lowercase and locale considerations.

**Answer:**
`str.toUpperCase()` / `str.toLowerCase()`.
**Caveat:** Without locale, it uses Default Locale.
Better: `str.toUpperCase(Locale.ENGLISH)` or `Locale.ROOT` to avoid surprises (e.g., Turkish 'i' issue).
