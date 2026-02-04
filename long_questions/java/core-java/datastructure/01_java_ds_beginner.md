# 🟢 Java Data Structures: Level 1 (Beginner)
*Topics: Arrays, Strings*

## 1️⃣ Arrays (Java Array & Array Manipulation)

### Question 1: How do you declare, initialize, and copy an array in Java?

**Answer:**
**Declaration & Initialization:**
```java
// Declaration
int[] arr; 

// Initialization (Size must be fixed)
arr = new int[5]; // Default values (0)

// Declaration + Initialization + Assignment
int[] numbers = {1, 2, 3, 4, 5};
```
**Copying:**
Use `System.arraycopy()`, `Arrays.copyOf()`, or `clone()`.

---

### Question 2: Difference between `Arrays.copyOf()` and `System.arraycopy()`.

**Answer:**
- **`System.arraycopy(src, srcPos, dest, destPos, length)`**:
  - Native method (fastest).
  - Copies data into an **existing** destination array.
  - Throws exception if destination is too small.

- **`Arrays.copyOf(original, newLength)`**:
  - Internally uses `System.arraycopy`.
  - Creates and **returns a NEW array**.
  - Can be used to resize an array (pad with defaults if larger).

---

### Question 3: Difference between shallow copy and deep copy of arrays.

**Answer:**
- **Shallow Copy:**
  - Copies references.
  - 1D Primitive arrays: Effectively deep (values copied).
  - Object arrays (or 2D arrays): Only references copied. Modifying an object in the copy affects the original.
  - `clone()`, `System.arraycopy()` perform shallow copies.

- **Deep Copy:**
  - Recursively copies objects.
  - Must be implemented manually (looping and strictly cloning elements) or using Serialization.

---

### Question 4: How do you find the maximum or minimum in an array?

**Answer:**
Iterate via loop.
```java
int[] arr = {5, 2, 9, 1};
int max = arr[0];
for (int i : arr) {
    if (i > max) max = i;
}
```
Or use Streams: `Arrays.stream(arr).max().getAsInt()`.

---

### Question 5: How do you reverse an array in place?

**Answer:**
Use Two-Pointer approach. Swap elements at `start` and `end`, then move pointers inward.
Time: O(N), Space: O(1).
```java
int i = 0, j = arr.length - 1;
while (i < j) {
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
    i++; j--;
}
```

---

### Question 6: How do you rotate an array by k positions?

**Answer:**
**Algorithm (Reversal Method):**
To rotate Right by K:
1.  Reverse whole array.
2.  Reverse first K elements.
3.  Reverse remaining N-K elements.
Time: O(N), Space: O(1).

---

### Question 7: How do you remove duplicates from an array?

**Answer:**
Arrays are fixed size, so you cannot physically "remove" an element (length doesn't change).
1.  **Using Set:** Convert to `HashSet` (O(N) space).
2.  **Sort + Scan (In-place):** Sort (O(N log N)), then overwrite unique elements to the front. Return new "logical" length.

---

### Question 8: How do you find the second largest element in an array?

**Answer:**
One pass O(N).
Maintain two variables: `largest` and `secondLargest`.
```java
if (num > largest) {
    secondLargest = largest;
    largest = num;
} else if (num > secondLargest && num != largest) {
    secondLargest = num;
}
```

---

### Question 9: How do you find a missing number in a sequence of integers?

**Answer:**
If sequence is 1 to N:
1.  Calculate expected sum: `N * (N + 1) / 2`.
2.  Calculate actual sum of array elements.
3.  Difference is the missing number.
(Avoid overflow by subtracting as you go, or use XOR method).

---

### Question 10: How do you find duplicate elements in an array?

**Answer:**
1.  **Brute Force:** O(N^2).
2.  **Sorting:** O(N log N) - check neighbors.
3.  **HashSet:** O(N) time, O(N) space. Add elements; if `set.add()` returns false, it's duplicate.
4.  **Negation (if 1 <= x <= N):** Visit index `abs(x)-1`. If value is negative, it's duplicate. Else flip sign.

---

### Question 11: `Arrays.sort()` vs `Collections.sort()`?

**Answer:**
- **`Arrays.sort(primitive[])`**: Uses **Dual-Pivot Quicksort**. O(N log N). Not stable.
- **`Arrays.sort(Object[])`** / **`Collections.sort()`**: Uses **Timsort** (MergeSort + InsertionSort). Stable. O(N log N).

---

### Question 12: `Arrays.binarySearch()` — how does it work? What are the prerequisites?

**Answer:**
Performs Binary Search O(log N).
**Prerequisite:** The array **MUST be sorted**.
**Returns:**
- Index of element if found.
- `-(insertion_point) - 1` if not found.

---

### Question 13: `Arrays.equals()` vs `==` for arrays?

**Answer:**
- **`==`**: Checks reference equality (Are they the same object in memory?).
- **`Arrays.equals(a, b)`**: Checks content equality (Are lengths same and elements equal?).
- **`Arrays.deepEquals(a, b)`**: Required for multidimensional arrays.

---

### Question 14: `Arrays.fill()` — practical use cases?

**Answer:**
Fills an array with a specific value.
`Arrays.fill(arr, -1);` (Useful for initializing DP tables or memoization arrays).
Time: O(N).

---

### Question 15: `Arrays.asList()` — caveats when using with arrays

**Answer:**
Returns a **fixed-size** list backed by the array.
- You **cannot** add/remove elements (Throws `UnsupportedOperationException`).
- Modifying list elements modifies the original array.
- Pass usage: `new ArrayList<>(Arrays.asList(arr))` to get a modifiable list.

---

### Question 16: `Arrays.stream()` — creating streams from arrays

**Answer:**
`IntStream s = Arrays.stream(intArr);`
Supports slicing: `Arrays.stream(arr, 0, 5)`.
Useful for functional operations: `Arrays.stream(arr).filter(x -> x > 10).sum()`.

---

### Question 17: How to convert a primitive array to a list or stream?

**Answer:**
`int[]` cannot be directly preserved in `Arrays.asList()` (it becomes `List<int[]>`).
**Correct way:**
1.  **Stream:** `Arrays.stream(arr).boxed().collect(Collectors.toList())`.
2.  **Apache Commons:** `ArrayUtils.toObject(arr)`.

---

### Question 18: How do you declare, initialize, and traverse a 2D array?

**Answer:**
**Decl:** `int[][] matrix = new int[3][4];` (3 rows, 4 cols).
**Init:** `int[][] m = { {1,2}, {3,4} };`
**Traverse:** Nested loops.
```java
for (int i = 0; i < matrix.length; i++) {
    for (int j = 0; j < matrix[i].length; j++) {
        print(matrix[i][j]);
    }
}
```

---

### Question 19: How do you find the sum of all elements in a 2D array?

**Answer:**
Iterate all elements and accumulate.
Or `Arrays.stream(matrix).flatMapToInt(Arrays::stream).sum()`.

---

### Question 20: How do you rotate a 2D matrix (like image rotation)?

**Answer:**
**Rotate 90 deg Cokwise:**
1.  **Transpose:** Swap `matrix[i][j]` with `matrix[j][i]`.
2.  **Reverse Rows:** Reverse each row individually.

---

### Question 21: How do you check if a matrix is symmetric?

**Answer:**
A matrix is symmetric if `matrix[i][j] == matrix[j][i]` for all i, j.
Must be a square matrix (rows == cols).

---

### Question 22: How to efficiently search in a row-wise & column-wise sorted matrix?

**Answer:**
Start from **Top-Right** corner (or Bottom-Left).
Search `Target`:
- If `current == target`: Found.
- If `current < target`: Move Down (Row++).
- If `current > target`: Move Left (Col--).
Time: O(N + M).

---

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
