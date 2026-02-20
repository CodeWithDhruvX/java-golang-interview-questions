# 19. Java Programs (Arrays Logic)

**Q: Largest/Smallest Element in Array**
> "Initialize `max` (or `min`) to `array[0]`.
> Loop through the array from `i = 1`.
> If `array[i] > max`, update `max = array[i]`.
> If `array[i] < min`, update `min = array[i]`.
>
> Simple O(n) scan. No need to sort the array first (which would be O(n log n))."

**Indepth:**
> **Edge Case**: Always initialize `min/max` to `array[0]` or `Integer.MAX_VALUE / MIN_VALUE`. Initializing to `0` is a bug if the array contains negative numbers (e.g., `[-5, -10]` -> max would wrongly be 0).


---

**Q: Reverse an Array**
> "Don't create a new array unless you have to (that wastes memory). Do it **in-place**.
> Use two pointers: `start = 0` and `end = length - 1`.
> While `start < end`:
> 1.  Swap `array[start]` and `array[end]`.
> 2.  Increment `start`.
> 3.  Decrement `end`."

**Indepth:**
> **Collections**: If you have a `List`, you can just call `Collections.reverse(list)`. It uses the exact same swap login internally but is much more readable.


---

**Q: Bubble Sort Logic**
> "It's the simplest sorting algorithm, but also the slowest (O(n^2)).
> Two nested loops.
> *   Outer loop `i` from 0 to n.
> *   Inner loop `j` from 0 to `n - i - 1`.
> *   If `array[j] > array[j+1]`, swap them.
>
> The largest elements 'bubble' to the top (end) of the array with each pass."

**Indepth:**
> **Optimization**: Bubble sort can be optimized with a boolean flag `swapped`. If a full pass happens with zero swaps, the array is already sorted, and you can break early (best case O(n)).


---

**Q: Linear Search vs Binary Search**
> "**Linear Search**: Scan every element one by one. Works on unsorted arrays. Complexity: O(n).
>
> **Binary Search**: Divide and conquer. **Requires sorted array**.
> 1.  Look at the middle element.
> 2.  If target is smaller, ignore the right half.
> 3.  If target is larger, ignore the left half.
> 4.  Repeat.
> Complexity: O(log n). Much faster for large datasets."

**Indepth:**
> **Overflow**: When calculating mid, using `(low + high) / 2` can overflow if low and high are massive integers. The safer way is `low + (high - low) / 2` (or `(low + high) >>> 1`).


---

**Q: Remove Duplicates from Sorted Array**
> "Since it's **sorted**, duplicates are always adjacent. You can do this in O(n) space or O(1) space.
>
> In-place (O(1) space):
> Use a pointer `j` to track the position of unique elements.
> Loop `i` from 0 to n-1.
> If `array[i] != array[i+1]`, place `array[i]` at `array[j]` and increment `j`.
> Finally, store the last element and return `j` as the new length."

**Indepth:**
> **Sets**: The easiest way to remove duplicates? `new LinkedHashSet<>(list)`. It removes duplicates *and* preserves insertion order.


---

**Q: Second Largest Number**
> "You can sort the array and take the second-to-last element, but that's O(n log n). Better to do it in one pass O(n).
>
> Maintain two variables: `largest` and `secondLargest`.
> Loop through the array:
> 1.  If `current > largest`:
>     *   `secondLargest` becomes `largest`.
>     *   `largest` becomes `current`.
> 2.  Else if `current > secondLargest` AND `current != largest`:
>     *   `secondLargest` becomes `current`.
>
> This handles cases where duplicates exist (like [10, 10, 5] -> second largest is 5)."

**Indepth:**
> **Stream API**: `list.stream().distinct().sorted(Comparator.reverseOrder()).skip(1).findFirst()` is a readable one-liner, but significantly slower than the single-pass loop.


---

**Q: Missing Number in Range 1 to N**
> "Math to the rescue!
> The sum of numbers from 1 to N is `N * (N + 1) / 2`.
>
> 1.  Calculate expected sum using the formula.
> 2.  Iterate through the array and calculate the actual `currentSum`.
> 3.  `Missing Number = expectedSum - currentSum`.
>
> This assumes only *one* number is missing. If multiple are missing, you need a HashSet or a boolean array."

**Indepth:**
> **XOR Method**: A robust way to find the missing number (that avoids integer overflow for large sums) is XOR-ing all indices and all values. `(1^2^...^N) ^ (arr[0]^...^arr[n-1])`. The result is the missing number.


---

**Q: Merge Two Arrays**
> "Create a new array of size `len1 + len2`.
> 1.  Loop through first array and copy items.
> 2.  Loop through second array and copy items (offset index by `len1`).
>
> If you need to merge them **in sorted order** (like Merge Sort):
> Use pointers `i` and `j` for the two arrays. Compare `arr1[i]` and `arr2[j]`. Pick the smaller one, add to result, and move the pointer. Once one array is exhausted, dump the rest of the other array."

**Indepth:**
> **In-Place**: Merging two sorted arrays *in-place* without extra space is a hard problem (requires complex shifting logic or gaps). The standard merge sort uses O(n) auxiliary space.


---

**Q: Find Common Elements (Intersection)**
> "Naive approach: Nested loops. For each element in A, scan B. O(n*m). Slow.
>
> Better approach (use HashSet):
> 1.  Dump array A into a `HashSet`.
> 2.  Loop through array B.
> 3.  If `set.contains(element)`, it's a match! (And remove it from set to avoid duplicates).
> Complexity: O(n + m)."

**Indepth:**
> **RetainAll**: `Set` has a built-in method `setA.retainAll(setB)`, which performs an intersection (keeps only elements present in both sets).


---

**Q: Rotate Array (Left/Right)**
> "To rotate an array by `k` positions efficiently (O(n) time, O(1) space), use the **Value Reversal Algorithm**:
>
> 1.  Reverse the whole array.
> 2.  Reverse the first `k` elements.
> 3.  Reverse the remaining `n - k` elements.
>
> This magically shifts everything into the correct place without needing a temporary array."

**Indepth:**
> **Juggling Algorithm**: Another O(n) approach involves moving elements in cycles computed by the GCD of n and k. It's more complex to implement but conceptually elegant.

