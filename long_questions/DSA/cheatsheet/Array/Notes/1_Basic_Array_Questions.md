# 1. Basic Array Problems (Warm-up)

---

## 1. Find Largest and Smallest Element

```
------------------------------------
| Problem Title                    |
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
Find the maximum and minimum values in an array in a single pass.
Return both values.

### 2️⃣ Pattern / Category ⭐
**Single Pass / Two Variables Tracking**

### 3️⃣ Brute Force Idea
Sort the array ($O(N \log N)$).
Smallest is at index 0, largest at index N-1.

### 4️⃣ Key Insight (AHA 💡)
We can assume the first element is both min and max, then compare every subsequent element against these two running values.

### 5️⃣ Algorithm
1. If array empty, return error.
2. Initialize `minVal` and `maxVal` to `arr[0]`.
3. Loop from index 1 to end:
   - If `arr[i] > maxVal`, update `maxVal`.
   - If `arr[i] < minVal`, update `minVal`.
4. Return `minVal`, `maxVal`.

### 6️⃣ Edge Cases & Traps ⚠️
* Empty array
* Array with 1 element
* Array with all duplicate values

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 2. Reverse an Array In-Place

```
------------------------------------
| Problem Title -> Reverse Array   |
...
```

### 1️⃣ Problem Snapshot
Reverse the given array without using a new array (modify original).
Example: `[1, 2, 3]` -> `[3, 2, 1]`

### 2️⃣ Pattern / Category ⭐
**Two Pointers (Start & End)**

### 3️⃣ Brute Force Idea
Create a new array given size. Copy element at `N-1` to `0`, `N-2` to `1`... then replace original. Space O(N).

### 4️⃣ Key Insight (AHA 💡)
Swapping the first element with the last, the second with the second-last, effectively reverses the array as we move towards the center.

### 5️⃣ Algorithm
1. Initialize `left = 0`, `right = n-1`.
2. While `left < right`:
   - Swap `arr[left]` and `arr[right]`.
   - Increment `left`, decrement `right`.
3. Stop when pointers meet or cross.

### 6️⃣ Edge Cases & Traps ⚠️
* Even number of elements (pointers cross)
* Odd number of elements (pointers meet at center)
* Empty or single-element array

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 3. Find Second Largest Element

```
------------------------------------
| Problem Title -> 2nd Largest     |
...
```

### 1️⃣ Problem Snapshot
Find the second highest value in the array.
Note: It must be strictly smaller than the largest if distinct, or just the second rank value if duplicates logic allows (usually distinct).

### 2️⃣ Pattern / Category ⭐
**Two Variables / Tournament Logic**

### 3️⃣ Brute Force Idea
Sort array descending. The element at index 1 is the answer (if distinct). $O(N \log N)$.

### 4️⃣ Key Insight (AHA 💡)
Maintain `largest` and `secondLargest`. Provide a clear update rule: if `x > largest`, `largest` becomes `second`, `x` becomes `largest`. If `x` is between them, update `second`.

### 5️⃣ Algorithm
1. Initialize `largest = -INF`, `second = -INF`.
2. Iterate through array:
   - If `arr[i] > largest`: update `second = largest`, `largest = arr[i]`.
   - Else if `arr[i] > second` AND `arr[i] != largest`: update `second = arr[i]`.
3. If `second` is still `-INF`, no second largest exists.

### 6️⃣ Edge Cases & Traps ⚠️
* Array logic with duplicates (e.g., `[10, 10, 5]` -> Is 2nd largest 10 or 5? usually 5).
* Array with < 2 elements.
* All elements equal.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 4. Check if Array is Sorted

```
------------------------------------
| Problem Title -> Is Sorted?      |
...
```

### 1️⃣ Problem Snapshot
Return `true` if array is currently sorted in non-decreasing order, `false` otherwise.

### 2️⃣ Pattern / Category ⭐
**Linear Scann / Adjacent Comparison**

### 3️⃣ Brute Force Idea
Copy array, sort the copy, compare with original. O(N log N).

### 4️⃣ Key Insight (AHA 💡)
A sorted array must satisfy `arr[i] <= arr[i+1]` for all valid `i`. One violation is enough to prove it false.

### 5️⃣ Algorithm
1. Loop `i` from 0 to `N-2`.
2. check if `arr[i] > arr[i+1]`.
3. If condition meets, return `false` immediately.
4. If loop finishes, return `true`.

### 6️⃣ Edge Cases & Traps ⚠️
* Single element (always sorted)
* Empty array (usually true)
* Duplicates (`[1, 1, 2]` is sorted)

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 5. Count Even and Odd Elements

```
------------------------------------
| Problem Title -> Count Even/Odd  |
...
```

### 1️⃣ Problem Snapshot
Count how many integers are even and how many are odd in an array.

### 2️⃣ Pattern / Category ⭐
**Modulo Arithmetic / Counter**

### 3️⃣ Brute Force Idea
Iterate and check. (This is already the optimal way).

### 4️⃣ Key Insight (AHA 💡)
`number % 2 == 0` check is sufficient.

### 5️⃣ Algorithm
1. Init `evenCount = 0`, `oddCount = 0`.
2. Loop through `arr`.
3. If `arr[i] % 2 == 0` increment `evenCount`, else increment `oddCount`.
4. Return counts.

### 6️⃣ Edge Cases & Traps ⚠️
* Negative numbers (in some languages modulo behaves differently, but for parity check `%2` usually works checking 0 or !0).
* 0 is even.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 6. Remove Duplicates from Sorted Array

```
------------------------------------
| Problem Title -> Remove Dups     |
...
```

### 1️⃣ Problem Snapshot
Given a **sorted** array, remove duplicates **in-place** such that each element appears only once and return the new length.
`[1,1,2]` -> `[1,2]`, return 2.

### 2️⃣ Pattern / Category ⭐
**Two Pointers (Reader & Writer)**

### 3️⃣ Brute Force Idea
Use a HashSet to store unique elements, then put back into array. Space O(N).

### 4️⃣ Key Insight (AHA 💡)
Since sorted, duplicates are grouped. We can use a slow pointer (`i`) to track the "end of unique section" and a fast pointer (`j`) to find new unique elements.

### 5️⃣ Algorithm
1. If empty, return 0.
2. `i = 0`.
3. Loop `j` from 1 to `N-1`.
4. If `arr[j] != arr[i]`:
   - Increment `i`.
   - `arr[i] = arr[j]`.
5. Return `i + 1` (length).

### 6️⃣ Edge Cases & Traps ⚠️
* Array with all same elements.
* Array with no duplicates.
* Empty array.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 7. Left Rotate Array by 1 Position

```
------------------------------------
| Problem Title -> Shift Left      |
...
```

### 1️⃣ Problem Snapshot
Shift all elements to the left by 1. The first element moves to the last position.
`[1, 2, 3, 4, 5]` -> `[2, 3, 4, 5, 1]`

### 2️⃣ Pattern / Category ⭐
**Store First & Shift**

### 3️⃣ Brute Force Idea
Create new array, fill indices `i` with `i+1` value, handle wrap around.

### 4️⃣ Key Insight (AHA 💡)
Store the first element in a temporary variable. Move everyone else one step left. Put temp at the end.

### 5️⃣ Algorithm
1. Store `temp = arr[0]`.
2. Loop `i` from 1 to `N-1`.
3. `arr[i-1] = arr[i]`.
4. `arr[N-1] = temp`.

### 6️⃣ Edge Cases & Traps ⚠️
* Single element (no change).
* Empty array.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 8. Left Rotate Array by K Positions

```
------------------------------------
| Problem Title -> Rotate K        |
...
```

### 1️⃣ Problem Snapshot
Rotate array to left by `K` steps.
`[1,2,3,4,5]`, K=2 -> `[3,4,5,1,2]`.

### 2️⃣ Pattern / Category ⭐
**Reversal Algorithm**

### 3️⃣ Brute Force Idea
Call "Rotate by 1" K times. Time O(N*K) -> Too slow.
Or use extra array O(N) space.

### 4️⃣ Key Insight (AHA 💡)
Reverse parts of the array to achieve rotation:
1. Reverse first `K` elements.
2. Reverse remaining `N-K` elements.
3. Reverse entire array.
(Or reverse whole, then parts, depending on left/right direction logic).
**For Left Rotate:**
Reverse(0, K-1)
Reverse(K, N-1)
Reverse(0, N-1)

### 5️⃣ Algorithm
1. `K = K % N` (handle K > N).
2. `reverse(arr, 0, K-1)`.
3. `reverse(arr, K, N-1)`.
4. `reverse(arr, 0, N-1)`.

### 6️⃣ Edge Cases & Traps ⚠️
* `K > N` (use modulo).
* `K = 0`.
* `N = 0`.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 9. Find Sum of All Elements

```
------------------------------------
| Problem Title -> Array Sum       |
...
```

### 1️⃣ Problem Snapshot
Calculate the sum of all integers in the array.

### 2️⃣ Pattern / Category ⭐
**Accumulator**

### 3️⃣ Brute Force Idea
Iterate and add.

### 4️⃣ Key Insight (AHA 💡)
Beware of integer overflow if numbers are large.

### 5️⃣ Algorithm
1. `sum = 0`.
2. Foreach `x` in `arr`, `sum += x`.
3. Return `sum`.

### 6️⃣ Edge Cases & Traps ⚠️
* Integer Overflow (use `int64` or `BigInt` if needed).
* Empty array (sum 0).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 10. Find Frequency of Each Element

```
------------------------------------
| Problem Title -> Frequency Count |
...
```

### 1️⃣ Problem Snapshot
Count how many times each element appears.
`[1, 1, 2, 3]` -> `{1: 2, 2: 1, 3: 1}`

### 2️⃣ Pattern / Category ⭐
**Hash Map / Dictionary**

### 3️⃣ Brute Force Idea
For each element, count its occurrences in rest of array. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
A Hash Map gives O(1) lookup. Pass through array once to populate counts.

### 5️⃣ Algorithm
1. Create empty Map `freq`.
2. Loop through `arr`:
   - `freq[arr[i]]++`
3. Print or return Map.

### 6️⃣ Edge Cases & Traps ⚠️
* Negative numbers (Hash Map handles this, array-based generic frequency table might fail).
* Large numbers.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)
