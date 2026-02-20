# 20. Java Programs (String Algorithms)

**Q: Reverse a String (Logic)**
> "Don't use `StringBuilder.reverse()` in an interview unless allowed.
> Convert to char array: `char[] chars = str.toCharArray()`.
> Use two pointers: `left = 0`, `right = length - 1`.
> While `left < right`:
> 1.  Swap `chars[left]` and `chars[right]`.
> 2.  Move `left` forward, `right` backward.
> Return `new String(chars)`."

**Indepth:**
> **Surrogates**: `reverse()` in StringBuilder is actually quite complex because it has to handle Surrogate Pairs (Unicode characters that take 2 chars). If you reverse blindly, you might split an Emoji into two invalid characters. The built-in method handles this.


---

**Q: Check Palindrome String**
> "A palindrome reads the same backwards (e.g., 'MADAM').
> Loop `i` from 0 to `length / 2`.
> If `str.charAt(i) != str.charAt(length - 1 - i)`, return false.
> If you finish the loop, it's true."

**Indepth:**
> **Optimization**: You don't need to check the whole string. `i < length / 2` is enough. Checking `i < length` does double the work (checking everything twice) but gives the same result.


---

**Q: Anagram Check**
> "Anagrams have the same characters in different orders (e.g., 'Silent' vs 'Listen').
>
> **Approach 1 (Sorting)**: Clean strings (lowercase, remove spaces). Convert to char arrays. Sort them. `Arrays.equals(arr1, arr2)`. Complexity: O(n log n).
>
> **Approach 2 (Frequency Array - Faster)**:
> Create an int array `counts` of size 26.
> Loop through string 1: `counts[char - 'a']++`.
> Loop through string 2: `counts[char - 'a']--`.
> Finally, check if every element in `counts` is 0. Complexity: O(n)."

**Indepth:**
> **HashMap vs Array**: Use the array approach (`int[26]`) whenever possible. A HashMap has overhead for hashing, resizing, and boxing integers. The array is pure memory access and is order of magnitude faster.


---

**Q: Count Vowels and Consonants**
> "Iterate through the string.
> Convert char to lowercase.
> If it's 'a', 'e', 'i', 'o', 'u', increment `vowels`.
> Else if it's between 'a' and 'z', increment `consonants`.
> Ignore numbers and symbols."

**Indepth:**
> **Regex**: You can also use `str.replaceAll("[^aeiouAEIOU]", "").length()` to count vowels. It's shorter code but much slower due to compiling the regex.


---

**Q: Find First Non-Repeated Character**
> "This requires two passes.
>
> 1.  **Count Frequencies**: Use a `HashMap<Character, Integer>` (or `int[256]` array for ASCII). Loop through string and fill the map.
> 2.  **Check Order**: Loop through the **String** again (not the map!).
> 3.  The first character where `map.get(char) == 1` is your answer."

**Indepth:**
> **LinkedHashMap**: Alternatively, use a `LinkedHashMap` to store counts. Since it preserves insertion order, you can just iterate through the *Map* afterwards and pick the first one with count 1.


---

**Q: String to Integer (atoi)**
> "Converting '123' to `int` manually.
>
> 1.  Initialize `result = 0`.
> 2.  Loop through digits.
> 3.  `int digit = char - '0';` (ASCII magic).
> 4.  `result = result * 10 + digit;`
>
> **Edge Cases**:
> *   Handle negative sign at index 0.
> *   Check for non-digit characters (throw exception).
> *   Check for integer overflow (if result exceeds `Integer.MAX_VALUE`)."

**Indepth:**
> **Overflow Logic**: To check for overflow *before* it happens: `if (result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE / 10 && digit > 7))`.


---

**Q: Integer to String**
> "The easy way: `String.valueOf(123)`.
>
> The usage way (if asked logic):
> Use a `StringBuilder`.
> While `num > 0`:
> 1.  `digit = num % 10`.
> 2.  Append digit to builder.
> 3.  `num /= 10`.
> Finally, reverse the builder (because you extracted digits backwards)."

**Indepth:**
> **Log10**: You can determine the number of digits using `Math.log10(num) + 1` to pre-allocate the StringBuilder capacity, avoiding resizing.


---

**Q: Reverse Words in a Sentence**
> "Input: 'Hello World Java'. Output: 'Java World Hello'.
>
> 1.  `String[] words = str.split(\" \");`
> 2.  Use a `StringBuilder`.
> 3.  Loop through `words` array **backwards** (from `len-1` down to 0).
> 4.  Append word + space.
> 5.  Trim the result."

**Indepth:**
> **Whitespace**: `split(" ")` leaves empty strings if there are multiple spaces ("Hello   World"). Use `split("\\s+")` to treat multiple spaces as a single delimiter.


---

**Q: Check for Subsequence**
> "Is 'ace' a subsequence of 'abcde'? Yes.
> Use two pointers: `i` for small string, `j` for big string.
> While `i < smallLen` and `j < bigLen`:
> *   If `small[i] == big[j]`, increment `i` (found a match).
> *   Always increment `j` (keep moving in big string).
>
> If `i == smallLen`, you found all characters in order."

**Indepth:**
> **Iterators**: A cleaner way (if allowed) is to use `Iterator` on the big string and advance it only when a match is found. It's the same logic but more abstract.


---

**Q: Rotation Check**
> "Is 'BCDA' a rotation of 'ABCD'?
>
> The trick:
> 1.  Check if lengths are equal.
> 2.  Concatenate original string with itself: `DoubleStr = "ABCD" + "ABCD" -> "ABCDABCD"`.
> 3.  Check if 'BCDA' is a substring of `DoubleStr`.
> `return (str1.length() == str2.length()) && (str1 + str1).contains(str2);`"

**Indepth:**
> **KMP Algorithm**: `contains()` is naïve (O(n*m)). For massive strings, KMP (Knuth-Morris-Pratt) is O(n+m) because it avoids backtracking by creating a prefix table.


---

**Q: Permutations of a String**
> "This is a recursive backtracking problem.
> Method `permute(String str, String answer)`:
> 1.  **Base Case**: If `str` is empty, print `answer`.
> 2.  **Recursive Step**: Loop `i` from 0 to length.
>     *   Pick character at `i`.
>     *   Rest of string = `substring(0, i) + substring(i+1)`.
>     *   Call `permute(rest, answer + char)`."

**Indepth:**
> **Complexity**: There are `n!` permutations. This algorithm is O(n * n!). Even for a small string like "12characters", this will run forever. It's only feasible for inputs of length <= 10.

