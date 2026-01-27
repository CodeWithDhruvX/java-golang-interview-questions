# Problem Analysis & Approaches

This document provides a detailed breakdown of coding problems, including understanding, intuition, dry runs, and solution approaches.

---

## 1. STRING PROBLEMS (Q1 - Q20)

### 1. Reverse a String
**1. Understanding the Problem**
- Input: A string (e.g., "Hello").
- Output: The string with characters in reverse order (e.g., "olleH").

**2. Ways to Solve**
- **Method 1 (Brute Force):** Create a new string/array. Loop from the end of the original string to the start, appending characters to the new string. Time: O(N), Space: O(N).
- **Method 2 (Two Pointers - Optimized):** Swap characters at the start and end indices, moving towards the center. Time: O(N), Space: O(1) (in-place for mutable strings).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** "Copying backwards" is the most natural way to think. Take the last char, put it first.
- **Optimized:** Reversing is symmetric. The first char swaps with the last, second with second-to-last. We only need to iterate half the length.

**4. Dry Run**
- Input: `str = "Code"`
- **Step 1:** `i=0` ('C'), `j=3` ('e'). Swap. `str = "eodC"`
- **Step 2:** `i=1` ('o'), `j=2` ('d'). Swap. `str = "edoC"`
- **Step 3:** `i=2`, `j=1`. Stop (i >= j).
- Output: `"edoC"`

---

### 2. Reverse Words in a Sentence
**1. Understanding the Problem**
- Input: "Hello World"
- Output: "World Hello" (Words order reversed, not characters within words).

**2. Ways to Solve**
- **Method 1 (Split & Swap):** Split string by space into an array of words. Reverse the array. Join back. Time: O(N), Space: O(N).
- **Method 2 (Double Reverse):** Reverse the entire string first ("dlroW olleH"). Then reverse each individual word ("World Hello"). Time: O(N), Space: O(1) (if mutable).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** Treat words as tokens. Just reorder the tokens.
- **Optimized:** If you view the sentence as a stack of characters, reversing the whole thing flips the order of words *and* chars. Removing the char flip (by re-reversing words) leaves just the word order flipped.

**4. Dry Run**
- Input: `"Hi Bye"`
- **Double Reverse Method:**
    1. Reverse entire: `"eyB iH"`
    2. Reverse "eyB" -> `"Bye"` (Current: `"Bye iH"`)
    3. Reverse "iH" -> `"Hi"` (Current: `"Bye Hi"`)

---

### 3. Check Palindrome
**1. Understanding the Problem**
- Check if a string reads the same forward and backward.
- Example: "madam" -> true.

**2. Ways to Solve**
- **Method 1 (Reverse & Compare):** Generate reversed string. Check `original == reversed`. Time: O(N), Space: O(N).
- **Method 2 (Two Pointers - Optimized):** Compare `str[start]` vs `str[end]` moving inwards. Mismatch = false. Time: O(N), Space: O(1).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** A palindrome is equal to its reverse.
- **Optimized:** We don't need to generate the full reverse. As soon as we find a mismatch at symmetrical positions, it's not a palindrome.

**4. Dry Run**
- Input: `"aba"`
- `i=0` ('a'), `j=2` ('a'). Match.
- `i=1` ('b'), `j=1` ('b'). Match.
- i > j. End. Return True.

---

### 4. Check Anagrams
**1. Understanding the Problem**
- Two strings are anagrams if they contain the same characters with the same frequencies.
- Example: "listen", "silent" -> true.

**2. Ways to Solve**
- **Method 1 (Sort):** Sort both strings. Compare `sortedStr1 == sortedStr2`. Time: O(N log N).
- **Method 2 (Frequency Count - Optimized):** Use a HashMap or Array[256] to count chars in `str1`. Decrement for `str2`. All counts should be 0. Time: O(N), Space: O(1) (fixed alphabet).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force (Sort):** Canonical form. If they are the same "stuff", they look identical when sorted.
- **Optimized:** We don't care about order, just quantity. "Net zero" balance sheet for characters.

**4. Dry Run**
- Input: `"rat"`, `"art"`
- Counts: `{r:1, a:1, t:1}`
- Process "art": 'a' (count `r:1, a:0, t:1`), 'r' (count `r:0, a:0, t:1`), 't' (count `r:0, a:0, t:0`).
- All zero. True.

---

### 5. Count Vowels & Consonants
**1. Understanding the Problem**
- Count letters that are vowels (a,e,i,o,u) vs others.

**2. Ways to Solve**
- **Method 1 (Iterate):** Loop through char array. Check if char in "aeiou". Increment counters. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Simple categorization task. Iterate and classify each character.

**4. Dry Run**
- Input: `"Hi"`
- 'H' -> Consonant (C=1)
- 'i' -> Vowel (V=1)

---

### 6. Character Frequency
**1. Understanding the Problem**
- Count how many times each character appears.
- Input: "banana" -> b:1, a:3, n:2.

**2. Ways to Solve**
- **Method 1 (Map):** HashMap<Char, Int>. Iterate and update count. Time: O(N).
- **Method 2 (Array):** If ASCII data, use `int[256]`. Index = char ASCII value. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Use a lookup table (Map or Array) to store the state (count) associated with each unique item (char).

---

### 7. First Non-Repeating Character
**1. Understanding the Problem**
- Find the first char that appears only once.
- Input: "swiss" -> 'w' ('s' repeats).

**2. Ways to Solve**
- **Method 1 (Nested Loop):** For each char, check if it appears elsewhere. Time: O(N^2).
- **Method 2 (Two Pass - Optimized):** 
    1. Count frequencies (Map/Array). 
    2. Iterate string again, check if `count[char] == 1`. Return first match. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** Pick one, scan the rest. Repeat. Slow.
- **Optimized:** We need global knowledge (total counts) to make a local decision (is *this* one unique?). Pre-calculate global knowledge.

**4. Dry Run**
- Input: `"lol"`
- Map: `{l:2, o:1}`
- Scan "lol":
    1. 'l': count is 2. Skip.
    2. 'o': count is 1. Return 'o'.

---

### 8. Remove Duplicates
**1. Understanding the Problem**
- Input: "banana" -> "ban" (keep unique chars).

**2. Ways to Solve**
- **Method 1 (Set):** Use a HashSet to track seen characters. Append to result only if not in Set. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- "Seen list". As you process, check "Have I seen this?". If no, keep it and add to "Seen list".

---

### 9. Replace Spaces
**1. Understanding the Problem**
- Replace ' ' with a specific char/string (e.g., "%20").

**2. Ways to Solve**
- **Method 1 (StringBuilder):** Create new buffer. Copy chars. If space, append replacement. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- String reconstruction. Strings are often immutable, so building a new one is standard.

---

### 10. Convert Lowercase to Uppercase
**1. Understanding the Problem**
- "abc" -> "ABC" without `toUpperCase()`.

**2. Ways to Solve**
- **Method 1 (ASCII Math):** `a` is 97, `A` is 65. Difference is 32. If char is 'a'-'z', subtract 32. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Characters are just numbers internally. Shifting the number shifts the casing.

---

### 11. Longest Word in Sentence
**1. Understanding the Problem**
- Input: "I love code" -> "love" or "code" (max length).

**2. Ways to Solve**
- **Method 1 (Split & Scan):** Split by space. Iterate words, keep track of `maxLen` and `maxWord`. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Tokenize and measure. Classic "find max in array" logic applied to string lengths.

---

### 12. Count Words
**1. Understanding the Problem**
- Count whitespace-separated tokens.

**2. Ways to Solve**
- **Method 1 (Split):** `str.trim().split("\\s+")`. Length of array. Time: O(N).
- **Method 2 (Iterate):** Count spaces where previous char was not a space (handling multiple spaces).

**3. Intuition (Brute Force & Optimized)**
- A word ends when a space begins (or string ends).

---

### 13. Substring Check
**1. Understanding the Problem**
- Does "Hello World" contain "World"?

**2. Ways to Solve**
- **Method 1 (Built-in):** `.contains()` or `.indexOf()`.
- **Method 2 (Sliding Window/KMP - Optimized):** Check for match at every index. Simple loops O(N*M). KMP O(N+M).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** Try to match the pattern starting at index 0, then 1, then 2...
- **Optimized (KMP):** If we fail a match, use previous success info to skip ahead intelligently.

---

### 14. Remove Vowels
**1. Understanding the Problem**
- "Apple" -> "ppl".

**2. Ways to Solve**
- **Method 1 (Filter):** Append chars to new string if NOT in "aeiou". Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Filtering stream. Pass through filter, keep what passes.

---

### 15. Sort Characters
**1. Understanding the Problem**
- "cba" -> "abc".

**2. Ways to Solve**
- **Method 1 (Arrays.sort):** Convert to char array, sort, convert back. Time: O(N log N).
- **Method 2 (Counting Sort - Optimized):** Count freq of a-z. Rebuild string based on counts. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** Comparisons.
- **Optimized:** Since alphabet is small (26 chars), just count them and print 'a' count times, then 'b'...

---

### 16. Find Duplicates
**1. Understanding the Problem**
- Print chars that appear > 1 time.

**2. Ways to Solve**
- **Method 1 (Map/Count):** Same as logic for Frequency Count. Iterate map, print if value > 1.

---

### 17. Reverse Recursively
**1. Understanding the Problem**
- Reverse string without loops.

**2. Ways to Solve**
- **Method 1 (Recursion):** `reverse(str) = reverse(substring(1)) + charAt(0)`.
- Base Case: Empty string returns empty.

**3. Intuition (Brute Force & Optimized)**
- Take the first item, reverse the rest of the pile, then put the first item at the end.

**4. Dry Run**
- `rev("ab")` -> `rev("b") + 'a'`
- `rev("b")` -> `rev("") + 'b'`
- `rev("")` -> `""`
- Result: `"" + 'b' + 'a'` = "ba".

---

### 18. Zig-Zag String
**1. Understanding the Problem**
- Print string in a sinusoidal wave pattern over K rows.

**2. Ways to Solve**
- **Method 1 (Row Simulation):** Use K strings (one for each row). Iterate chars, processing "down" then "up". Append char to current row. Join rows. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Simulate the pen movement. Move index variable `row` up and down (`row++`, `row--`) and place characters.

---

### 19. String Rotation Check
**1. Understanding the Problem**
- Is "CDAB" a rotation of "ABCD"?

**2. Ways to Solve**
- **Method 1 (Concatenation):** `(s1 + s1).contains(s2)`.
- **Logic:** `ABCD` + `ABCD` = `ABCDABCD`. This allows checking any wrap-around. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Doubling the string simulates the circle. Any rotation must appear as a substring in the doubled version.

**4. Dry Run**
- s1="AB", s2="BA"
- s1+s1 = "ABAB".
- Contains "BA"? Yes.

---

### 20. Compare Strings
**1. Understanding the Problem**
- Check equality without `equals()`.

**2. Ways to Solve**
- **Method 1 (Char by Char):** Check length. Loop `i` from 0 to N. If `s1[i] != s2[i]`, false. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Definition of equality: Same length, same stuff at same place.

---

## 2. ARRAY PROBLEMS (Q21 - Q40)

### 21. Find Largest Element
**1. Understanding the Problem**
- Find the maximum value in an array.
- Input: [1, 5, 3] -> Output: 5.

**2. Ways to Solve**
- **Method 1 (Sort):** Sort array. Return `arr[n-1]`. Time: O(N log N).
- **Method 2 (Iterate - Optimized):** Assume `max = arr[0]`. Loop `i=1` to `N`. If `arr[i] > max`, update `max`. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- King of the hill. A challenger (`arr[i]`) only takes the throne (`max`) if they are bigger.

**4. Dry Run**
- Arr: [1, 5, 2]
- `max=1`.
- `5 > 1`? Yes. `max=5`.
- `2 > 5`? No.
- Result: 5.

---

### 22. Find Smallest Element
**1. Understanding the Problem**
- Find the minimum value.

**2. Ways to Solve**
- **Method 1 (Iterate):** Similar to finding Max. Initialize `min = arr[0]`. Update if `arr[i] < min`. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Same as Largest, just checking for explicitly smaller values.

---

### 23. Find Second Largest
**1. Understanding the Problem**
- Find the 2nd highest unique value.
- Input: [10, 20, 20, 5] -> Output: 10.

**2. Ways to Solve**
- **Method 1 (Sort):** Sort Logic. Traverse from end, picking first element `< max`. Time: O(N log N).
- **Method 2 (One Pass - Optimized):** Maintain `first` and `second`. If `curr > first`: update second=first, first=curr. If `first > curr > second`: update second=curr. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** Sort it, then look at the second to last item.
- **Optimized:** Podium places. If a new winner comes in, the old winner drops to specific, and the old second drops off.

**4. Dry Run**
- Arr: [10, 20, 5]
- `first=10`, `second=-inf`
- Process 20: 20 > 10. `second=10`, `first=20`.
- Process 5: 5 < 10. No change.
- Result: 10.

---

### 24. Reverse Array
**1. Understanding the Problem**
- Flip the array order. [1, 2, 3] -> [3, 2, 1].

**2. Ways to Solve**
- **Method 1 (Two Pointers):** Swap `arr[start]` and `arr[end]`. Increment `start`, decrement `end`. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Identical to Reversing a String. Swap ends moving inwards.

---

### 25. Rotate Array Left by K
**1. Understanding the Problem**
- Shift elements left. Wrap around.
- [1, 2, 3, 4, 5], k=1 -> [2, 3, 4, 5, 1].

**2. Ways to Solve**
- **Method 1 (Temp Array):** Store first k elements. Shift rest. Append stored. Time: O(N), Space: O(k).
- **Method 2 (Reversal Algo - Optimized):** 
    1. Reverse 0 to k-1.
    2. Reverse k to n-1.
    3. Reverse 0 to n-1. 
    Time: O(N), Space: O(1).

**3. Intuition (Brute Force & Optimized)**
- **Brute Force:** Pick up the chunk to move, slide the rest over, put the chunk down.
- **Optimized:** Rotating is mathematically equivalent to reversing parts and then reversing the whole. 

**4. Dry Run**
- Input: [1, 2, 3, 4, 5], k=2.
- Rev(0,1): [2, 1, 3, 4, 5]
- Rev(2,4): [2, 1, 5, 4, 3]
- Rev(0,4): [3, 4, 5, 1, 2] (Correct).

---

### 26. Rotate Array Right by K
**1. Understanding the Problem**
- [1, 2, 3, 4, 5], k=1 -> [5, 1, 2, 3, 4].

**2. Ways to Solve**
- **Method 1 (Reversal Algo):** 
    1. Reverse n-k to n-1 (tail).
    2. Reverse 0 to n-k-1 (head).
    3. Reverse all.

---

### 27. Remove Duplicates (Sorted Array)
**1. Understanding the Problem**
- In-place removal. [1, 1, 2] -> [1, 2].

**2. Ways to Solve**
- **Method 1 (Two Pointers):** `i` tracks unique position. `j` scans. If `arr[j] != arr[i]`, increment `i`, `arr[i] = arr[j]`. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Construction crew. `j` finds new material. `i` builds the solid wall of unique/good material.

**4. Dry Run**
- [1, 1, 2]. `i=0`, `j=1`.
- `arr[1] == arr[0]`. Skip.
- `j=2`. `arr[2] (2) != arr[0] (1)`. `i=1`. `arr[1] = 2`.
- Result: [1, 2].

---

### 28. Frequency of Elements
**1. Understanding the Problem**
- Count occurrences.

**2. Ways to Solve**
- **Method 1 (Map):** Same as String Frequency. Time: O(N).

---

### 29. Find Missing Number (1 to N)
**1. Understanding the Problem**
- Array of size N-1 has numbers 1..N. Find missing.

**2. Ways to Solve**
- **Method 1 (Sum Formula):** Sum(1..N) = `N*(N+1)/2`. Difference between expected sum and actual array sum is the missing number. Time: O(N).
- **Method 2 (XOR):** XOR all elements 1..N. XOR with array elements. Result is missing No. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Math magic. If you know the total weight of a full cart, and you weigh your cart, the difference is the missing item.

**4. Dry Run**
- [1, 3], N=3.
- Expected Sum (1+2+3) = 6.
- Actual Sum (1+3) = 4.
- Missing: 2.

---

### 30. Find Duplicate Number
**1. Understanding the Problem**
- Find the number repeated.

**2. Ways to Solve**
- **Method 1 (Set):** Store seen. If seen, return. Time: O(N).
- **Method 2 (Floyd's Cycle - Optimized):** Treat indices as linked list pointers. `arr[i]` points to next index. Use Slow/Fast pointers to find cycle. Time: O(N), Space: O(1).

**3. Intuition (Brute Force & Optimized)**
- **Set:** simple check list.
- **Floyd's:** Because values are within 1..N range, they form a valid pointer graph. A duplicate value means two nodes point to the same next node (cycle entry).

---

### 31. Sort Array (Bubble Sort)
**1. Understanding the Problem**
- Sort without library.

**2. Ways to Solve**
- **Method 1 (Bubble Sort):** Repeatedly swap adjacent elements if wrong order. Largest bubbles to top. Time: O(N^2).

---

### 32. Merge Two Arrays
**1. Understanding the Problem**
- [1, 2] + [3, 4] -> [1, 2, 3, 4].

**2. Ways to Solve**
- **Method 1 (New Array):** Create array of size `n1+n2`. Copy first, copy second. Time: O(N+M).

---

### 33. Common Elements (Intersection)
**1. Understanding the Problem**
- Elements present in both arrays.

**2. Ways to Solve**
- **Method 1 (Set):** Put arr1 in Set. Iterate arr2, check if in Set. Time: O(N+M).

---

### 34. Move Zeros to End
**1. Understanding the Problem**
- [0, 1, 0, 3] -> [1, 3, 0, 0].

**2. Ways to Solve**
- **Method 1 (Count & Fill):** Count non-zeros, place them at start. Fill rest with 0. Time: O(N).
- **Method 2 (Swap - Optimized):** `count` tracks position of next non-zero. If `arr[i] != 0`, `swap(arr[count], arr[i])`, `count++`. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Snowplow. Push all the real cars (non-zeros) to the front, leave the snow (zeros) behind.

**4. Dry Run**
- [0, 1, 0]. `count=0`.
- `i=0`: 0. Skip.
- `i=1`: 1. Swap `arr[0]` and `arr[1]`. `count=1`. Arr: [1, 0, 0].
- Result: [1, 0, 0].

---

### 35. Sum of Array
**1. Understanding the Problem**
- Sum all elements.

**2. Ways to Solve**
- **Method 1 (Iterate):** `total += arr[i]`.

---

### 36. Pair with Given Sum (Two Sum)
**1. Understanding the Problem**
- Find a, b such that a + b = Target.

**2. Ways to Solve**
- **Method 1 (Nested Loop):** Check all pairs. Time: O(N^2).
- **Method 2 (Map - Optimized):** Map stores `value -> index`. For each `x`, check if `Target - x` is in Map. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- **Map:** "Hey, I'm a 7, looking for a 3. Has a 3 passed by yet?"

---

### 37. Max & Min in Single Loop
**1. Understanding the Problem**
- Find range.

**2. Ways to Solve**
- **Method 1 (Iterate):** Update `min` and `max` in the same for-loop iteration.

---

### 38. Print Reverse
**1. Understanding the Problem**
- Just print, don't modify.

**2. Ways to Solve**
- **Method 1 (Loop Down):** `for i = n-1 to 0`. Use print.

---

### 39. Check Sorted
**1. Understanding the Problem**
- Is array sorted ascending?

**2. Ways to Solve**
- **Method 1 (Iterate):** If `arr[i] > arr[i+1]` at any point, False. If loop finishes, True. Time: O(N).

---

### 40. Count Even & Odd
**1. Understanding the Problem**
- Count numbers divisible by 2.

**2. Ways to Solve**
- **Method 1 (Modulus):** `if num % 2 == 0` even++ else odd++.

---

## 3. PATTERN PROBLEMS (Q41 - Q55)

### 41. Right Triangle Star Pattern
**1. Understanding the Problem**
- Print distinct rows of stars increasing in count.
- *
- **
- ***

**2. Ways to Solve**
- **Method 1 (Nested Loop):** Outer loop `i` from 1 to N. Inner loop `j` from 1 to `i`. Print star.

**3. Intuition (Brute Force & Optimized)**
- Validating the grid structure. Row 1 has 1 item, Row 2 has 2... `row` controls the count.

---

### 42. Left Triangle Star Pattern
**1. Understanding the Problem**
- Stars aligned to the right.
-   *
-  **
- ***

**2. Ways to Solve**
- **Method 1 (Spaces + Stars):** Inner loop 1: Print `N-i` spaces. Inner loop 2: Print `i` stars.

**3. Intuition (Brute Force & Optimized)**
- You need to push the stars. The 'push' is invisible characters (spaces). The number of spaces decreases as stars increase.

---

### 43. Pyramid Star Pattern
**1. Understanding the Problem**
- Centered triangle.
-   *
-  * *
- * * *

**2. Ways to Solve**
- **Method 1 (Spaces + Star-Space):** Print `N-i` leading spaces. Then print `i` stars, but append a space after each star to spread them out.

---

### 44. Inverted Pyramid
**1. Understanding the Problem**
- Upside down pyramid.

**2. Ways to Solve**
- **Method 1 (Reverse Loop):** Loop `i` from N down to 1. Same logic as Pyramid.

---

### 45. Diamond Pattern
**1. Understanding the Problem**
- Pyramid on top of Inverted Pyramid.

**2. Ways to Solve**
- **Method 1 (Combine):** Run Pyramid logic for 1 to N. Then Run Inverted Pyramid logic for N-1 down to 1.

---

### 46. Number Pyramid
**1. Understanding the Problem**
-   1
-  1 2
- 1 2 3

**2. Ways to Solve**
- **Method 1 (Nested Loop):** Similar to star pyramid, but print `j` instead of `*`.

---

### 47. Floyd’s Triangle
**1. Understanding the Problem**
- 1
- 2 3
- 4 5 6

**2. Ways to Solve**
- **Method 1 (Counter):** Maintain a global counter `num=1`. In nested loop, print `num` and increment it.

**3. Intuition (Brute Force & Optimized)**
- Don't reset numbers per row. Just keep counting.

---

### 48. Hollow Square
**1. Understanding the Problem**
- Border of stars only.

**2. Ways to Solve**
- **Method 1 (Condition Check):** Iterate valid grid `N x N`. If `row==1` or `row==N` or `col==1` or `col==N`, print `*`, else print ` `.

**3. Intuition (Brute Force & Optimized)**
- Boundary check logic. "Am on the edge?"

---

### 49. Hollow Pyramid
**1. Understanding the Problem**
- Triangle border.

**2. Ways to Solve**
- **Method 1 (Math Pattern):** Check if `col==1` (left slope) or `col==2*i-1` (right slope) or `row==N` (base).

---

### 50. Zig-Zag Star Pattern
**1. Understanding the Problem**
-    *   *
-   * * * *
-  *   *   *

**2. Ways to Solve**
- **Method 1 (Modulus Logic):** The pattern repeats every 4 columns. Check `(row+col)%4 == 0` or similar conditions for specific rows.

---

### 51. Butterfly Pattern
**1. Understanding the Problem**
- *    *
- **  **
- ******
- **  **
- *    *

**2. Ways to Solve**
- **Method 1 (Mirror):** Top half: Stars = `i`, Spaces = `2*(N-i)`. Bottom half: Reverse order.

---

### 52. Pascal’s Triangle
**1. Understanding the Problem**
- Coefficients of binomial expansion. Each number is sum of two above it.

**2. Ways to Solve**
- **Method 1 (Formula):** `C(n,k) = C(n,k-1) * (n-k+1) / k`. Compute next value using previous.

**3. Intuition (Brute Force & Optimized)**
- Use the mathematical property rather than building a 2D array and summing.

---

### 53. Number Increasing Pattern
**1. Understanding the Problem**
- 1
- 1 2
- 1 2 3

**2. Ways to Solve**
- **Method 1:** Standard Reset loop. `j` goes 1 to `i`.

---

### 54. Number Increasing Reverse
**1. Understanding the Problem**
- 1 2 3 4
- 1 2 3
- 1 2

**2. Ways to Solve**
- **Method 1:** Outer loop `i` from N down to 1.

---

### 55. Number Changing Pyramid
**1. Understanding the Problem**
- 1
- 2 3
- 4 5 6

**2. Ways to Solve**
- **Method 1:** Same as Floyd's Triangle.

---

## 4. MATH PROBLEMS (Q56 - Q70)

### 56. Check Prime Number
**1. Understanding the Problem**
- Number > 1 divisible only by 1 and itself.

**2. Ways to Solve**
- **Method 1 (Loop N):** Check mod from 2 to N-1. Time: O(N).
- **Method 2 (Sqrt - Optimized):** Check mod from 2 to sqrt(N). Time: O(sqrt(N)).

**3. Intuition (Brute Force & Optimized)**
- **Optimized:** Factors come in pairs (a*b = N). One of them must be <= sqrt(N). If we don't find a factor by sqrt(N), there isn't one (except N).

---

### 57. Print Primes in Range
**1. Understanding the Problem**
- Print all primes between A and B.

**2. Ways to Solve**
- **Method 1 (Nested Loop):** For each number, run `isPrime`. Time: O(N * sqrt(N)).

---

### 58. Fibonacci Series
**1. Understanding the Problem**
- 0, 1, 1, 2, 3, 5... (Next = sum of prev two).

**2. Ways to Solve**
- **Method 1 (Iterative):** `a=0, b=1`. Loop: `c=a+b`, `a=b`, `b=c`. Time: O(N).
- **Method 2 (Recursion):** `fib(n) = fib(n-1) + fib(n-2)`. Time: O(2^N) - Very Slow.

**3. Intuition (Brute Force & Optimized)**
- **Iterative:** Just keep a running total of the last two numbers. Efficient.

---

### 59. Factorial
**1. Understanding the Problem**
- N! = 1 * 2 * ... * N.

**2. Ways to Solve**
- **Method 1 (Loop):** `res *= i`.

---

### 60. Armstrong Number
**1. Understanding the Problem**
- Sum of digits raised to power of count of digits equals number.
- 153: 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153.

**2. Ways to Solve**
- **Method 1 (Digit Extraction):** Count digits. Then loop: extract last digit (`%10`), add `pow(digit, count)`, remove last digit (`/10`).

**3. Intuition (Brute Force & Optimized)**
- Number property check. Requires breaking number into components.

---

### 61. Sum of Digits
**1. Understanding the Problem**
- 123 -> 1+2+3 = 6.

**2. Ways to Solve**
- **Method 1 (Modulo):** `sum += n % 10`, `n /= 10`.

---

### 62. Reverse Number
**1. Understanding the Problem**
- 123 -> 321.

**2. Ways to Solve**
- **Method 1 (Modulo Math):** `rev = rev*10 + digit`.

**3. Intuition (Brute Force & Optimized)**
- Shift current reverse to the left (multiply by 10) to make room for the new ones digit.

---

### 63. Palindrome Number
**1. Understanding the Problem**
- 121 -> True.

**2. Ways to Solve**
- **Method 1 (Reverse):** Calculate `reverse(N)`. Check `N == reverse`.

---

### 64. Swap Two Numbers (No Temp)
**1. Understanding the Problem**
- a=5, b=10 -> a=10, b=5 without `c`.

**2. Ways to Solve**
- **Method 1 (Arithmetic):** `a = a+b` (Total), `b = a-b` (Total - b = original a), `a = a-b` (Total - new b = original b).

**3. Intuition (Brute Force & Optimized)**
- Mix the colors, then extract one color to verify the other. Warning: Potential overflow.

---

### 65. LCM
**1. Understanding the Problem**
- Least Common Multiple.

**2. Ways to Solve**
- **Method 1 (Formula):** `LCM(a, b) = (a * b) / GCD(a, b)`.

---

### 66. GCD (HCF)
**1. Understanding the Problem**
- Greatest Common Divisor.

**2. Ways to Solve**
- **Method 1 (Euclidean Algo - Optimized):** `gcd(a, b) = gcd(b, a % b)`. Base: `gcd(a, 0) = a`.

**3. Intuition (Brute Force & Optimized)**
- If `a` divides `b` and `c`, it divides `b-c`. Repeat until remainder is 0.

---

### 67. Check Power of 2
**1. Understanding the Problem**
- Is N = 2^k?

**2. Ways to Solve**
- **Method 1 (Loop):** Keep dividing by 2. If remainder != 0, false.
- **Method 2 (Bitwise - Optimized):** `N & (N-1) == 0`. (and N > 0).

**3. Intuition (Brute Force & Optimized)**
- **Bitwise:** Powers of 2 are `100...0`. Subtracting 1 gives `011...1`. ANDing them always gives 0.

---

### 68. Square Root
**1. Understanding the Problem**
- Calc sqrt(N).

**2. Ways to Solve**
- **Method 1 (Built-in):** `Math.sqrt`.
- **Method 2 (Newton's Method):** Approximation loop. `root = 0.5 * (X + N/X)`.

---

### 69. Perfect Number
**1. Understanding the Problem**
- Sum of divisors (excluding self) = Number. 6 (1+2+3).

**2. Ways to Solve**
- **Method 1 (Loop):** Loop 1 to N/2. Sum factors.

---

### 70. Count Digits
**1. Understanding the Problem**
- Number of digits.

**2. Ways to Solve**
- **Method 1 (Loop):** `while(n>0) n/=10; count++`.
- **Method 2 (String):** `String(n).length()`.


---


---

## 5. SEARCHING & SORTING (Q71 - Q90)

### 71. Linear Search
**1. Understanding the Problem**
- Find the index of a target element in an array. Return -1 if not found.
- Input: [10, 50, 30], Target: 30 -> Output: 2.

**2. Ways to Solve**
- **Method 1 (Iteration):** Traverse from 0 to N-1. If arr[i] == target, return i. Time: O(N).

**3. Intuition (Brute Force & Optimized)**
- Simple scan. Look at each item one by one until you find what you want.

**4. Dry Run**
- Input: [10, 20, 30], Target=20.
- i=0 (10 != 20).
- i=1 (20 == 20). Return 1.

---

### 72. Binary Search
**1. Understanding the Problem**
- Find target in a **sorted** array efficiently.

**2. Ways to Solve**
- **Method 1 (Linear):** O(N) scan.
- **Method 2 (Divide & Conquer - Optimized):** Check middle. If target < mid, search left. Else search right. Time: O(log N).

**3. Intuition (Brute Force & Optimized)**
- **Optimized:** Dictionary search. You don't read every word. You open the middle, see if the word is before or after, and cut the problem in half repeatedly.

**4. Dry Run**
- Arr: [10, 20, 30, 40], T=30.
- L=0, R=3, Mid=1 (20). 30 > 20. L=2.
- L=2, R=3, Mid=2 (30). 30 == 30. Return 2.

---

### 73. Bubble Sort
**1. Understanding the Problem**
- Sort array by repeatedly swapping adjacent elements.

**2. Ways to Solve**
- **Method 1 (Bubble):** Nested loop. Inner loop swaps arr[j] and arr[j+1] if out of order. Largest element bubbles to the end in each pass. Time: O(N^2).

**3. Intuition (Brute Force & Optimized)**
- Like heavy bubbles rising in water (or sinking stones). Push the biggest current item to the far right steps by step.

**4. Dry Run**
- [3, 1, 2]
- Pass 1: Swap 3,1->[1,3,2]. Swap 3,2->[1,2,3].
- Pass 2: Swap 1,2 (No change).

---

### 74. Selection Sort
**1. Understanding the Problem**
- Sort by finding the minimum element and placing it at the beginning.

**2. Ways to Solve**
- **Method 1 (Selection):** Loop i from 0 to N-1. Find min in i to N-1. Swap arr[i] with min. Time: O(N^2).

**3. Intuition (Brute Force & Optimized)**
- Select the best (smallest) candidate from the remaining pool and put it in its correct spot.

**4. Dry Run**
- [3, 1, 2]
- i=0: Find min in [3,1,2] -> 1. Swap 3,1 -> [1, 3, 2].
- i=1: Find min in [3,2] -> 2. Swap 3,2 -> [1, 2, 3].

---

### 75. Insertion Sort
**1. Understanding the Problem**
- Sort by building a sorted subarray one item at a time.

**2. Ways to Solve**
- **Method 1 (Insertion):** Take arr[i]. Compare with i-1, i-2.... Shift larger elements right. Insert arr[i] into valid position. Time: O(N^2).

**3. Intuition (Brute Force & Optimized)**
- Card game hand. You pick a new card and slide it into the correct position among the cards you already hold.

**4. Dry Run**
- [2, 1, 3]
- i=1 (Val 1): 2 > 1. Shift 2. [2, 2, 3]. Place 1. [1, 2, 3].
- i=2 (Val 3): 3 > 2. No shift.

---

### 76. Sort Array of 0s, 1s, 2s
**1. Understanding the Problem**
- Input: [0, 1, 2, 0, 1, 2] -> [0, 0, 1, 1, 2, 2].

**2. Ways to Solve**
- **Method 1 (Sort):** O(N log N).
- **Method 2 (Count):** Count 0s, 1s, 2s. Overwrite array. Time: O(N), Space: O(1).
- **Method 3 (Dutch National Flag - Optimized):** 3 pointers. Low (0s), Mid (1s), High (2s). If 0, swap Low/Mid. If 2, swap Mid/High. Time: O(N), One pass.

**3. Intuition (Brute Force & Optimized)**
- **Optimized:** Organize into three regions: Left (0), Middle (1), Right (2). Mid pointer scans specific items and throws them to Left or Right bucket.

**4. Dry Run**
- [1, 0, 2]
- Mid=1. No swap. Mid++.
- Mid=0. Swap with Low. [0, 1, 2]. Low++, Mid++.
- Mid=2. Swap with High.

---

### 77. Find Kth Largest Element
**1. Understanding the Problem**
- Input: [3, 2, 1, 5, 6, 4], k=2 -> Output: 5.

**2. Ways to Solve**
- **Method 1 (Sort):** Arrays.sort. Return arr[N-k]. Time: O(N log N).
- **Method 2 (Min-Heap - Optimized):** Maintain Heap of size K. Min element is top. If new > top, replace. Top is Nth largest. Time: O(N log K).

**3. Intuition (Brute Force & Optimized)**
- **Sort:** Simplest.
- **Heap:** Keep a Hall of Fame of size K. Only the best K survive. The smallest in the Hall of Fame is the Kth largest overall.

**4. Dry Run**
- Heap (Size 2): []
- Add 3: [3]
- Add 2: [2, 3] (Sorted for visualize)
- Add 1: 1 < 2. Ignore.
- Add 5: 5 > 2. Pop 2, Add 5. [3, 5].
- Result: Top is 3 (Wait, K=2. 3rd largest in heap size 2?). 
- *Correction:* For Kth Largest, use Min Heap of size K.
- Heap [3, 5]. Top is 3. Correct.

---

### 78. Find Median of Array
**1. Understanding the Problem**
- Middle element of sorted array.

**2. Ways to Solve**
- **Method 1 (Sort):** Sort -> Pick middle. Time: O(N log N).

**3. Intuition (Brute Force & Optimized)**
- The center of data.

**4. Dry Run**
- [3, 1, 2] -> Sort [1, 2, 3] -> Median 2.

---

### 79. Counting Sort
**1. Understanding the Problem**
- Sort numbers in specific range (e.g., 0-100).

**2. Ways to Solve**
- **Method 1 (Freq Array):** Count occurrences. Reconstruct array. Time: O(N+Range).

**3. Intuition (Brute Force & Optimized)**
- If you know values are limited (e.g. grades 0-100), just tally them up. 2 students got 90, 5 got 80...

**4. Dry Run**
- [1, 0, 1]
- Counts: 0->1, 1->2.
- Rebuild: 0 (Dec count), 1 (Dec count), 1 (Dec count). -> [0, 1, 1].

---

### 80. Sort String
**1. Understanding the Problem**
- edcba -> abcde.

**2. Ways to Solve**
- **Method 1 (Char Array):** str.toCharArray() -> Arrays.sort() -> new String.

**3. Intuition (Brute Force & Optimized)**
- Convert to sortable format (array), sort, convert back.

**4. Dry Run**
- edc -> [e, d, c] -> Sort [c, d, e] -> cde.

---

## 6. RECURSION & BASIC LOGIC

### 81. Reverse String using Recursion
**1. Understanding the Problem**
- Reverse without loops.

**2. Ways to Solve**
- **Recursion:** rev(s) = rev(tail) + head. Base case: length <= 1 return s.

**3. Intuition**
- Break down the string. The reverse of the whole is the reverse of the rest + the first part.

**4. Dry Run**
- rev(ab) -> rev(b) + a -> b + a -> ba.

---

### 82. Factorial using Recursion
**1. Understanding the Problem**
- 5! = 5 * 4 * 3 * 2 * 1.

**2. Ways to Solve**
- **Recursion:** fact(n) = n * fact(n-1). Base: n<=1 return 1.

**3. Intuition**
- A number times the factorial of the one below it.

**4. Dry Run**
- fact(3) -> 3 * fact(2) -> 3 * 2 * fact(1) -> 3 * 2 * 1 = 6.

---

### 83. Fibonacci using Recursion
**1. Understanding the Problem**
- Nth Fib number.

**2. Ways to Solve**
- **Recursion:** fib(n) = fib(n-1) + fib(n-2). (Inefficient O(2^N) without memoization).

**3. Intuition**
- Assume you know the previous two values. Just add them.

**4. Dry Run**
- fib(3) -> fib(2) + fib(1) -> (fib(1)+fib(0)) + 1 -> (1+0) + 1 = 2.

---

### 84. Sum of Array using Recursion
**1. Understanding the Problem**
- Sum elements.

**2. Ways to Solve**
- **Recursion:** sum(arr, n) = arr[n-1] + sum(arr, n-1).

**3. Intuition**
- Current element + Sum of the rest.

**4. Dry Run**
- sum([1,2], 2) -> 2 + sum([1,2], 1) -> 2 + 1 + sum(0) -> 3 + 0 = 3.

---

### 85. Print numbers 1 to N Recursively
**1. Understanding the Problem**
- Print distinct lines.

**2. Ways to Solve**
- **Recursion:** printN(n-1) THEN print(n). (Head recursion).

**3. Intuition**
- To print up to N, first print up to N-1.

**4. Dry Run**
- P(2) -> P(1) -> P(0)(Ret) -> Print 1 -> Print 2.

---

### 86. Check Palindrome Recursively
**1. Understanding the Problem**
- madam.

**2. Ways to Solve**
- **Recursion:** isPal(str, start, end) = (str[start] == str[end]) AND isPal(start+1, end-1).

**3. Intuition**
- Are ends same? If yes, check the inner part.

**4. Dry Run**
- aba -> a==a? Yes. Check b? Yes. True.

---

### 87. Count Digits Recursively
**1. Understanding the Problem**
- 123 -> 3.

**2. Ways to Solve**
- **Recursion:** 1 + count(n/10). Base: n==0 return 0.

**3. Intuition**
- One digit + count of the rest.

**4. Dry Run**
- C(12) -> 1 + C(1) -> 1 + 1 + C(0) -> 2.

---

### 88. Power of Number Recursively
**1. Understanding the Problem**
- 2^3 = 8.

**2. Ways to Solve**
- **Recursion:** base * power(base, exp-1).

**3. Intuition**
- Multiply base by result of power-1.

**4. Dry Run**
- P(2,2) -> 2 * P(2,1) -> 2 * 2 * P(2,0) -> 4 * 1 = 4.

---

### 89. GCD Recursively
**1. Understanding the Problem**
- GCD(48, 18).

**2. Ways to Solve**
- **Recursion:** gcd(b, a % b) if b != 0 else a.

**3. Intuition**
- Euclidean algorithm reduction.

**4. Dry Run**
- gcd(10, 5) -> gcd(5, 0) -> return 5.

---

### 90. Remove Duplicates Recursively
**1. Understanding the Problem**
- aaabb -> ab.

**2. Ways to Solve**
- **Recursion:** If s[0] == s[1] return rec(s.substring(1)). Else return s[0] + rec(s.substring(1)).

**3. Intuition**
- If neighbors same, skip one.

**4. Dry Run**
- rem(aa) -> a==a -> rem(a) -> a.

---

## 7. SCENARIO / OUTPUT-BASED (Q91 - Q100)

### 91. Predict output of loop
**1. Understanding the Problem**
- Analyze loop flow.
- Loop: i=0; i<3; i++. Print i.

**3. Intuition**
- Step variables manually line by line.

**4. Dry Run**
- i=0. Print 0.
- i=1. Print 1.
- i=2. Print 2.
- Stop.

### 92. Predict output of string operations
**1. Understanding the Problem**
- String + Integer behavior.

**3. Intuition**
- String dominance in concatenation.

**4. Dry Run**
- 5 + 2 -> 7.
- 7 + 2 -> 72.

### 93. Find error in given code
**1. Understanding the Problem**
- Look for syntax/logic bugs.

**3. Intuition**
- Compilation rules vs Runtime logic.

**4. Example**
- while(true); -> Infinite loop.

### 94. What will be printed (Array Index)
**1. Understanding the Problem**
- Access invalid index.

**3. Intuition**
- Arrays are fixed size 0 to N-1.

**4. Example**
- arr[3] (size 3) -> Error.

### 95. Time complexity of simple loop
**1. Understanding the Problem**
- O(N) estimation.

**3. Intuition**
- How many times does the statement run relative to N?

**4. Example**
- for(i=0; i<N; i++) -> N times.

### 96. Swap two numbers without temp
**1. Understanding the Problem**
- a=5, b=10. Swap.

**2. Ways to Solve**
- **Method 1:** a = a+b; b = a-b; a = a-b.

**3. Intuition**
- Sum contains both. Subtracting one gives the other.

**4. Dry Run**
- a=15. b=5. a=10.

### 97. Difference between == and equals()
**1. Understanding the Problem**
- Comparing Strings/Objects.

**2. Ways to Solve**
- **==**: Address compare.
- **equals**: Content compare.

**3. Intuition**
- Do you mean the same physical object or the same data value?

**4. Example**
- s1=new String(A). s2=new String(A).
- s1==s2 (False). s1.equals(s2) (True).

### 98. Difference between Array and ArrayList
**1. Understanding the Problem**
- Collection choice.

**3. Intuition**
- Static vs Dynamic memory.

**4. Example**
- Array: int[5]. ArrayList: .add(1).

### 99. Call by Value vs Reference
**1. Understanding the Problem**
- How functions receive data.

**3. Intuition**
- Java copies value. For objects, it copies the reference address value.

**4. Example**
- func(x). x doesnt change outside unless member fields modified.

### 100. Static vs Non-static
**1. Understanding the Problem**
- Global vs Instance scope.

**3. Intuition**
- Static = Blueprint data. Non-static = Build house data.

**4. Example**
- Static count = 10 (Shared). Instance Name = Bob (Unique).


---

## 7. ADDITIONAL 100 QUESTIONS

### Strings (Add'l Q1-15)

### 1. Find Longest Palindrome Substring
**1. Understanding the Problem**
- Find the longest substring that is a palindrome.
- Input: babad -> bab (or aba).

**2. Ways to Solve**
- **Method 1 (Expand Around Center):** Treat every character (and space between chars) as a center. Expand outwards while chars match. Time: O(N^2).
- **Method 2 (Brute Force):** Check all substrings. O(N^3).

**3. Intuition (Brute Force & Optimized)**
- A palindrome mirrors around a center. Why generate all? Just grow from every possible center.

**4. Dry Run**
- babad
- Center a (i=1): Expand left(b) right(b). Match. Expand b, d. Mismatch.
- Longest found: bab.

### 2. Find Longest Common Prefix
**1. Understanding the Problem**
- Input: [flower, flow, flight] -> fl.

**2. Ways to Solve**
- **Method 1 (Sort):** Sort array. Compare first and last string. Time: O(N log N * M).

**3. Intuition (Brute Force & Optimized)**
- If strings are sorted, the most different ones are at ends. If they share a prefix, everyone in between does too.

**4. Dry Run**
- Sorted: [flight, flow, flower]
- Compare flight & flower:
- f(Match), l(Match), i!=o.
- Prefix: fl.

### 3. Check if String contains only digits
**1. Understanding the Problem**
- 123 -> True. 12a -> False.

**2. Ways to Solve**
- **Method 1 (Regex):** s.matches([0-9]+).
- **Method 2 (Loop):** Check if each char is between 0 and 9.

**3. Intuition**
- Validate character class membership.

**4. Dry Run**
- 1a
- 1 is digit.
- a is not digit. Return False.

### 4. Count Uppercase & Lowercase
**1. Understanding the Problem**
- Count based on case.

**3. Intuition**
- ASCII ranges. A-Z vs a-z.

**4. Dry Run**
- Ab
- A: Upper++
- b: Lower++

### 5. Remove Special Characters
**1. Understanding the Problem**
- Remove anything not a-z, A-Z, 0-9.

**2. Ways to Solve**
- **Method 1 (Regex):** replaceAll([^a-zA-Z0-9], ).

**4. Dry Run**
- A#B -> Regex removes # -> AB.

### 6. Find all Permutations
**1. Understanding the Problem**
- ABC -> ABC, ACB, BAC, BCA...

**2. Ways to Solve**
- **Method 1 (Backtracking):** Swap char at index i with match, recurse, swap back.

**3. Intuition**
- Fix one char, permute the rest.

**4. Dry Run**
- P(ABC, 0)
- Swap A,A -> P(BC, 1) -> ... -> ABC
- Swap A,B -> P(AC, 1) -> ... -> BAC

### 7. Check Valid Parentheses
**1. Understanding the Problem**
- ()[]{} -> True.

**2. Ways to Solve**
- **Method 1 (Stack):** Push opening brackets. Pop matching closing brackets. Empty stack at end = Valid.

**3. Intuition**
- Last Opened, First Closed (LIFO).

**4. Dry Run**
- ([])
- Push ( -> Stack: [(]
- Push [ -> Stack: [(, []
- Pop ] matches [ -> Stack: [(]
- Pop ) matches ( -> Stack: []
- Empty -> Valid.

### 8. Find Duplicate Words
**1. Understanding the Problem**
- Count word frequencies > 1.

**2. Ways to Solve**
- **Method 1 (Map):** Split -> Count in Map -> Print > 1.

**4. Dry Run**
- hi hi bye
- Map: {hi:2, bye:1}
- Print hi.

### 9. Reverse each word in place
**1. Understanding the Problem**
- Hello World -> olleH dlroW.

**2. Ways to Solve**
- **Method 1 (Split):** Split -> Reverse each string -> Join.

**4. Dry Run**
- Hello -> olleH.
- World -> dlroW.
- Join -> olleH dlroW.

### 10. Check Isomorphic Strings
**1. Understanding the Problem**
- egg, add -> True (e->a, g->d).

**2. Ways to Solve**
- **Method 1 (Two Maps):** Map s->t and t->s. Ensure 1:1 consistent mapping.

**3. Intuition**
- Consistency check. If X maps to Y, X cannot map to Z, and W cannot map to Y.

**4. Dry Run**
- egg, add
- e->a.
- g->d.
- g->d (Consistent).

### 11. Check Pangram
**1. Understanding the Problem**
- Contains all 26 letters.

**2. Ways to Solve**
- **Method 1 (Set):** Add all chars to Set. Check size == 26.

**4. Dry Run**
- abc...
- Set adds a, b, c...
- Size at end 26? Yes.

### 12. Print all Substrings
**1. Understanding the Problem**
- Nested loops from i=0..n, j=i..n.

**3. Intuition**
- All continuous segments.

**4. Dry Run**
- ab
- i=0: a, ab
- i=1: b

### 13. Remove Consecutive Duplicates
**1. Understanding the Problem**
- aaabb -> ab. (Same as Q90 logic).

**2. Ways to Solve**
- **Method 1 (Loop):** If curr != prev, append.

**4. Dry Run**
- a a b
- a != null -> append a.
- a == a -> skip.
- b != a -> append b. -> ab.

### 14. Check if strings differ by one character
**1. Understanding the Problem**
- abc, abd -> True.

**2. Ways to Solve**
- **Method 1 (Loop):** Count mismatches at same index (if len equal) or shift index (if len diff).

**4. Dry Run**
- abc, abd
- a=a, b=b, c!=d (1 mismatch). True.

### 15. Find smallest & largest word
**1. Understanding the Problem**
- Find by length.

**2. Ways to Solve**
- **Method 1 (Loop):** Track minLen and maxLen.

**4. Dry Run**
- a bb
- a (len 1) -> min.
- bb (len 2) -> max.

### Arrays (Add'l Q16-35)

### 16. Find Leaders in Array
**1. Understanding the Problem**
- Element > all elements to its right.

**2. Ways to Solve**
- **Method 1 (Right Scan):** Start from right (n-1). Max from right is always a leader. Update max as you go left.

**3. Intuition**
- Looking forward is hard (requires loop). Looking backward from end means you already know the champion.

**4. Dry Run**
- [16, 17, 4, 3, 5, 2]
- max=2 (Leader).
- 5 > 2 -> Leader. max=5.
- 3 < 5.
- 4 < 5.
- 17 > 5 -> Leader. max=17.
- 16 < 17.

### 17. Find Equilibrium Index
**1. Understanding the Problem**
- Sum Left == Sum Right.

**2. Ways to Solve**
- **Method 1 (Total Sum):** totalSum. Iterate: rightSum = totalSum - leftSum - curr. Check leftSum == rightSum. Update leftSum.

**3. Intuition**
- Balance scale. As you move pointer right, weight shifts from right pan to left pan.

**4. Dry Run**
- [-7, 1, 5, 2, -4, 3, 0]
- Sum=0. Left=0.
- i=0 (-7): R = 0 - 0 - (-7) = 7. L != R. L += -7.
- ...

### 18. Subarray with Given Sum
**1. Understanding the Problem**
- Contiguous subarray sums to X.

**2. Ways to Solve**
- **Method 1 (Sliding Window):** Add to window. If > sum, shrink from left. (Works for non-negative).

**3. Intuition**
- Caterpillar method. Head eats (adds), Tail excretes (removes).

**4. Dry Run**
- [1, 4, 20], X=33.
- ...

### 19. Kadanes Algorithm
**1. Understanding the Problem**
- Max contiguous subarray sum.

**2. Ways to Solve**
- **Method 1 (Dynamic Prog):** curr = max(num, curr+num); global = max(global, curr).

**3. Intuition**
- If carrying a burden (negative sum) drags you down below zero, drop it and start fresh.

**4. Dry Run**
- [-2, 1, -3]
- -2. curr=-2.
- 1. curr = max(1, -2+1) = 1.
- -3. curr = max(-3, 1-3) = -2.
- Max was 1.

### 20. Majority Element
**1. Understanding the Problem**
- Element appears > N/2 times.

**2. Ways to Solve**
- **Method 1 (Moores Voting):** If count 0, set candidate. Match inc, mismatch dec.

**3. Intuition**
- War of attrition. Additional Q: If every soldier kills one enemy, who is left standing?

**4. Dry Run**
- [2, 2, 1, 1, 1, 2, 2]
- Cand: 2. Count 1.
- 2: Count 2.
- 1: Count 1.
- 1: Count 0.
- 1: Cand 1. Count 1.
- 2: Count 0.
- 2: Cand 2. Count 1.
- Result 2.

### 21. Rearrange Array Alternately
**1. Understanding the Problem**
- Max, Min, 2nd Max, 2nd Min...

**2. Ways to Solve**
- **Method 1 (Two Pointers):** Sorted array. Take from last (max), take from first (min).

**4. Dry Run**
- [1, 2, 3, 4] -> 4, 1, 3, 2.

### 22. Rotate Array (Reversal)
**1. Understanding the Problem**
- See Q25.

**3. Intuition**
- Reversing parts reverses the whole rotation logic.

### 23. Find Union of Arrays
**1. Understanding the Problem**
- All distinct elements.

**2. Ways to Solve**
- **Method 1 (Set):** Add all to set.

**4. Dry Run**
- [1, 2] U [2, 3] -> {1, 2, 3}.

### 24. Find Intersection of Arrays
**1. Understanding the Problem**
- Common elements.

**2. Ways to Solve**
- **Method 1 (Set):** Add A to Set. Check B against Set.

**4. Dry Run**
- [1, 2] n [2, 3] -> {2}.

### 25. Count Pairs with Difference
**1. Understanding the Problem**
- arr[i] - arr[j] = k.

**2. Ways to Solve**
- **Method 1 (Map):** Similar to Two Sum.

**3. Intuition**
- Look for x + k or x - k.


### 26. Find Peak Element
**1. Understanding the Problem**
- Greater than neighbors.

**2. Ways to Solve**
- **Method 1 (Binary Search):** If mid < mid+1, peak extends to right. Else left.

**3. Intuition**
- Climbing a hill. If the next step is up, go that way. If down, go the other way or stay.

**4. Dry Run**
- [1, 2, 3, 1]
- Mid=2 (Val 3). 3 > 1. Slope down. Look left.
- Found peak 3.

### 27. Left Rotate by 1
**1. Understanding the Problem**
- [1,2,3] -> [2,3,1].

**2. Ways to Solve**
- **Method 1:** Store arr[0]. Shift rest left. Place stored at end.

**4. Dry Run**
- Temp=1. Shift [2, 3, 3]. Place 1 -> [2, 3, 1].

### 28. Find Minimum Difference Pair
**1. Understanding the Problem**
- Min abs difference.

**2. Ways to Solve**
- **Method 1 (Sort):** Sort. Compare adjacent elements.

**3. Intuition**
- Closest values are neighbors in a sorted list.

**4. Dry Run**
- [4, 1, 8] -> Sort [1, 4, 8]
- 4-1=3. 8-4=4. Min=3.

### 29. Product of Array Except Self
**1. Understanding the Problem**
- res[i] = product of all except arr[i]. No division.

**2. Ways to Solve**
- **Method 1 (Prefix/Suffix):** LeftProd array and RightProd array. result = L * R.

**3. Intuition**
- Prefix product accumulates everything before. Suffix product accumulates everything after. Combine them.

**4. Dry Run**
- [1, 2]
- Left: [1, 1]
- Right: [2, 1]
- Res: [2, 1]

### 30. Max Product Subarray
**1. Understanding the Problem**
- Max product (negatives flip signs).

**2. Ways to Solve**
- **Method 1 (Swap Max/Min):** Maintain maxSoFar and minSoFar. If negative, swap max/min.

**3. Intuition**
- Negatives are tricky. A negative times a negative is a big positive. So track the biggest negative (minimum) too.

**4. Dry Run**
- [2, 3, -2, 4]
- 2: max=2.
- 3: max=6.
- -2: max=-2 (min=-12).
- 4: max=4.
- Global Max 6.

### 31. Check Circular Rotation
**1. Understanding the Problem**
- Same as String Rotation Q19.

### 32. Separate Positive and Negative
**1. Understanding the Problem**
- Move negatives to one side.

**2. Ways to Solve**
- **Method 1 (Partition):** QuickSort partition logic (pivot 0).

**4. Dry Run**
- [-1, 2, -3]
- Pivot 0. [-1, -3, 2].

### 33. Count Distinct Elements
**1. Understanding the Problem**
- Set.size().

### 34. Replace with Next Greatest
**1. Understanding the Problem**
- Replace arr[i] with max(arr[i+1]...end).

**2. Ways to Solve**
- **Method 1 (Right Scan):** Track max from right. Update arr[i] with current max. Update max with arr[i] (original).

**3. Intuition**
- Right-to-left pass carries the future knowledge.

**4. Dry Run**
- [16, 17, 4, 3, 5, 2]
- Start from 2. Max=2.
- 5: Replace with 2. Max=5.
- ...

### 35. Smallest Subarray with Sum > X
**1. Understanding the Problem**
- Min length.

**2. Ways to Solve**
- **Method 1 (Sliding Window):** Expand end. While sum > x, update minLen and shrink start.

**3. Intuition**
- Expand until valid, then shrink to minimize.

**4. Dry Run**
- Sum > 5. [1, 2, 2].
- Acc 1+2+2 = 5 (Not >). Wait.

### Matrix (Add'l Q36-50)

### 36. Matrix Addition
**1. Understanding the Problem**
- Add cell by cell.

**4. Dry Run**
- A[0][0] + B[0][0] -> C[0][0].

### 37. Matrix Multiplication
**1. Understanding the Problem**
- Row * Col dot product.
- **Complexity:** O(N^3).

**3. Intuition**
- Dot product of Row i from A and Col j from B.

**4. Dry Run**
- [1 2], [3 4]
- 1*3 + 2*...

### 38. Transpose of Matrix
**1. Understanding the Problem**
- Swap A[i][j] with A[j][i].

**3. Intuition**
- Flip over the diagonal.

### 39. Rotate Matrix 90 Degrees
**1. Understanding the Problem**
- Clockwise.

**2. Ways to Solve**
- **Method 1:** Transpose -> Reverse each Row.

**3. Intuition**
- Mathematical trick. Tranpose gets x,y to y,x. Reverse gets the right orientation.

**4. Dry Run**
- [1 2] -> T [1 3] -> Rev [3 1]
- [3 4]    [2 4]        [4 2]

### 40. Spiral Order Matrix
**1. Understanding the Problem**
- Print in spiral.

**2. Ways to Solve**
- **Method 1 (Boundaries):** Maintain Top, Bottom, Left, Right. Loop and shrink boundaries.

**3. Intuition**
- Walk until wall, turn right. Wall moves in.

### 41. Search Element in Sorted Matrix
**1. Understanding the Problem**
- Row sorted, Col sorted.

**2. Ways to Solve**
- **Method 1 (Staircase):** Start Top-Right. If < target, go Down. If > target, go Left. Time: O(M+N).

**3. Intuition**
- From top-right, left is smaller, down is bigger. Like a BST.

**4. Dry Run**
- Start TR. 15. Target 10. Left. 10 == 10. Found.

### 42. Diagonal Sum
**1. Understanding the Problem**
- Sum diag (i,i) and anti-diag (i, n-i-1). Handle center if odd.

### 43. Print Boundary Elements
**1. Understanding the Problem**
- Print 1st row, last col, last row, 1st col.

**3. Intuition**
- Just 4 loops.

### 44. Check Symmetric Matrix
**1. Understanding the Problem**
- A == Transpose(A).

**4. Dry Run**
- check A[i][j] == A[j][i].

### 45. Interchange Rows and Cols
**1. Understanding the Problem**
- Same as Transpose.

### 46. Count Zeros and Ones
**1. Understanding the Problem**
- Iterate and count.

### 47. Row with Maximum 1s
**1. Understanding the Problem**
- Find row text has most 1s.

**2. Ways to Solve**
- **Method 1:** Sum each row. Max aggregate.

### 48. Matrix Palindrome
**1. Understanding the Problem**
- Check symmetry logic.

### 49. Snake Pattern Printing
**1. Understanding the Problem**
- Even rows Left->Right. Odd rows Right->Left.

**3. Intuition**
- If i%2==0 loop j 0->N. Else loop j N->0.

### 50. Identity Matrix Check
**1. Understanding the Problem**
- 1s on diagonal, 0s elsewhere.

**3. Intuition**
- if i==j check 1. else check 0.


### Linked List (Add'l Q51-60)

### 51. Create Linked List
**1. Understanding the Problem**
- Define Node structure.

**2. Ways to Solve**
- **Method:** Class Node { int data; Node next; }. head = new Node(10).

**4. Example**
- Node n1 = new Node(10).

### 52. Traverse Linked List
**2. Ways to Solve**
- **Method:** while(curr != null) { print(curr.data); curr = curr.next; }

**4. Dry Run**
- [10|->20|->null]
- Curr=10. Print 10. Curr=20.
- Curr=20. Print 20. Curr=null.

### 53. Insert at Begin
**2. Ways to Solve**
- **Method:** newNode.next = head; head = newNode.

**3. Intuition**
- New guy holds the old boss's hand, then becomes the new boss.

### 54. Insert at End
**2. Ways to Solve**
- **Method:** Traverse to last. last.next = newNode.

### 55. Delete Node
**2. Ways to Solve**
- **Method:** prev.next = curr.next.

**3. Intuition**
- Bypass surgery. Connect A to C, skipping B.

### 56. Reverse Linked List
**1. Understanding the Problem**
- Reverse links.

**2. Ways to Solve**
- **Method 1 (Iterative):** 3 Pointers (Prev, Curr, Next).

**4. Dry Run**
- 1->2->3
- Prev=null, Curr=1.
- Next=2. 1->null. Prev=1, Curr=2.
- Next=3. 2->1. Prev=2, Curr=3.
- Next=null. 3->2. Prev=3. (Head).

### 57. Find Middle Element
**1. Understanding the Problem**
- One pass.

**2. Ways to Solve**
- **Method 1 (Slow/Fast):** Slow moves 1 step, Fast moves 2 steps. When Fast ends, Slow is at middle.

**3. Intuition**
- If you run twice as fast, you reach the end when I reach the middle.

**4. Dry Run**
- 1-2-3-4-5
- S=1, F=1.
- S=2, F=3.
- S=3, F=5 (End). Middle 3.

### 58. Detect Loop
**1. Understanding the Problem**
- Do links cycle?

**2. Ways to Solve**
- **Method 1 (Floyds Cycle):** If Slow == Fast, loop exists.

**3. Intuition**
- On a circular track, the fast runner will eventually lap the slow runner.

### 59. Count Nodes
**Method:** Traversal counter.

### 60. Merge Two Linked Lists
**Method:** Compare heads. Attach smaller to result. Recurse or Loop.

### Stack/Queue/Hashing/OOPs (Add'l Q61-100)

### 61. Stack Implementation
**1. Understanding the Problem**
- LIFO (Last In First Out).

**2. Ways to Solve**
- **Method:** Array (arr[++top] = val) or Linked List.

### 62. Queue Implementation
**1. Understanding the Problem**
- FIFO (First In First Out).

**2. Ways to Solve**
- **Method:** Array (Circular) or Linked List (Front/Rear).

### 63. Reverse String using Stack
**1. Understanding the Problem**
- abc -> cba using Stack.

**2. Ways to Solve**
- **Method:** Push all chars. Pop all chars (LIFO reverses order).

**4. Dry Run**
- Push a, b, c. Stack: [a, b, c] (Top is c).
- Pop c, b, a. Result cba.

### 64. Balanced Parentheses
**Method:** Stack (See Q7).

### 65. Stack using Queue
**1. Understanding the Problem**
- Use FIFO to make LIFO.

**2. Ways to Solve**
- **Method:** 2 Queues. Push: Enqueue. Pop: Move N-1 elements to other queue, dequeue last.

### 66. Queue using Stack
**2. Ways to Solve**
- **Method:** 2 Stacks (Inbox, Outbox). Push to Inbox. Pop from Outbox (if empty, pour Inbox to Outbox).

**3. Intuition**
- Inverting the order twice gives original order.

### 67. Next Greater Element
**1. Understanding the Problem**
- Closest greater element on right.

**2. Ways to Solve**
- **Method:** Monotonic Stack (Decreasing).

**3. Intuition**
- Keep a waiting list of people looking for a richer person. When a rich person comes, they clear the waiting list.

**4. Dry Run**
- [2, 1, 5]
- Stack: [2].
- 1 < 2. Stack: [2, 1].
- 5 > 1. Pop 1 -> NGE 5.
- 5 > 2. Pop 2 -> NGE 5.
- Push 5.

### 68. Evaluate Postfix Expression
**1. Understanding the Problem**
- 2 3 + -> 5.

**2. Ways to Solve**
- **Method:** Stack operands. Operator pops 2, pushes result.

**4. Dry Run**
- 2, 3, +
- Push 2, 3.
- +: Pop 3, 2. Add 2+3=5. Push 5.

### 69. Reverse Stack
**Method:** Recursion insertAtBottom.

### 70. Circular Queue
**Method:** (index + 1) % capacity.

### 71. Frequency of Elements
**Method:** Map<Element, Count>.

### 72. First Repeating Element
**Method:** Use Set to track seen.

### 73. First Non-Repeating Element
**Method:** Frequency Map.

### 74. Two Sum
**Method:** Map<Target - Curr, Index>.

### 75. Group Anagrams
**Method:** Key = SortedString. Value = List<Strings>.


### 76. Count Distinct Characters
**1. Understanding the Problem**
- Number of unique chars.

**2. Ways to Solve**
- **Method:** Set.size().

**4. Example**
- aab -> Set{a, b} -> Size 2.

### 77. Find Majority Element
**Method:** Map counts check if > N/2.

### 78. Check Subset
**1. Understanding the Problem**
- Is Array B a subset of A?

**2. Ways to Solve**
- **Method:** Put A into Map. Check B against Map.

**3. Intuition**
- Check every item in your bag (B) against the master inventory (A).

### 79. Common Elements
**Method:** Intersection (Set).

### 80. Longest Substring Without Repeating
**1. Understanding the Problem**
- abcabcbb -> abc (3).

**2. Ways to Solve**
- **Method:** Sliding Window + Map<Char, Index>. start = max(start, map.get(char) + 1).

**3. Intuition**
- Expand window. If duplicate found inside window, jump start past the previous occurrence.

**4. Dry Run**
- a b c a
- Map {a:0}. Start 0. Len 1.
- Map {a:0, b:1}. Start 0. Len 2.
- Map {a:0, b:1, c:2}. Start 0. Len 3.
- a found at 0. Start = max(0, 0+1) = 1. Window bca.

### OOPs Concepts (Add'l Q81-90)

### 81. Method Overloading
**Definition:** Same name, different parameters (Compile Time Poly).
**Example:** add(int a, int b) vs add(double a, double b).

### 82. Method Overriding
**Definition:** Subclass redefines Superclass method (Runtime Poly).
**Example:** Animal.speak() vs Dog.speak().

### 83. Encapsulation
**Definition:** Hiding data (private var, public getter).
**Intuition:** Shield internal gears from outside tampering.

### 84. Inheritance
**Definition:** extends. Is-A relationship.
**Intuition:** Don't reinvent the wheel. Car extends Vehicle.

### 85. Polymorphism
**Definition:** Many forms. Overloading vs Overriding.
**Intuition:** The same button (method) does different things depending on what object you push it on.

### 86. Abstract Class vs Interface
**Definition:** Abstract (Partial implementation) vs Interface (Contract/Multiple Implementation).

### 87. Final Keyword
**Definition:** Constant variable, Non-overridable method, Non-inheritable class.

### 88. Exception Handling
**Definition:** try, catch, finally, throw, throws.

### 89. Singleton Pattern
**Definition:** One instance per JVM. Private constructor.
**Code:** private static Instance; public static getInstance().

### 90. Immutable Class
**Definition:** Cannot change state. final class, final fields, no setters.
**Example:** String.

### SQL & Basic CS (Add'l Q91-100)

### 91. Second Highest Salary
**Query:** SELECT MAX(Salary) FROM Emp WHERE Salary < (SELECT MAX(Salary) FROM Emp).

### 92. Delete vs Truncate vs Drop
**Diff:** Delete (Rows, Logged), Truncate (All Rows, Fast), Drop (Schema).

### 93. Inner vs Left Join
**Diff:** Inner (Matches only), Left (All Left + Matches).

### 94. Find Duplicate Rows
**Query:** GROUP BY col HAVING COUNT(*) > 1.

### 95. Primary vs Foreign Key
**Diff:** PK (Unique ID), FK (Reference to PK).

### 96. Normalization
**Definition:** Organizing data to reduce redundancy (1NF, 2NF, 3NF).

### 97. Indexing
**Definition:** B-Tree structure for faster search (Scan avoidance).

### 98. ACID Properties
**Definition:** Atomicity, Consistency, Isolation, Durability.

### 99. Deadlock
**Definition:** Cycle of dependencies waiting for resources.
**Intuition:** I have Key A, need Key B. You have Key B, need Key A. We wait forever.

### 100. GET vs POST
**Definition:** GET (Retrieve, Idempotent), POST (Submit, Side Effects).

