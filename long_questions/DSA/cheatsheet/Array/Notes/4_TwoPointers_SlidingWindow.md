# 4. Two Pointer & Sliding Window Problems

---

## 1. Move All Zeros to End

```
------------------------------------
| Problem Title -> Move Zeros      |
```

### 1️⃣ Problem Snapshot
Move all `0`s to the end of array while maintaining relative order of non-zero elements.
`[0, 1, 0, 3, 12]` -> `[1, 3, 12, 0, 0]`.

### 2️⃣ Pattern / Category ⭐
**Two Pointers (Reader & Writer)**

### 3️⃣ Brute Force Idea
Create new array, copy non-zeros, fill rest with zeros. Space O(N).

### 4️⃣ Key Insight (AHA 💡)
We can maintain a `writeIndex`. Whenever we see a non-zero, we place it at `writeIndex` and increment `writeIndex`. After loop, fill remaining from `writeIndex` to end with 0s.

### 5️⃣ Algorithm
1. `write = 0`.
2. Loop `read` from 0 to `N-1`:
   - If `arr[read] != 0`: `arr[write] = arr[read]`, `write++`.
3. Loop from `write` to `N-1`:
   - `arr[i] = 0`.

### 6️⃣ Edge Cases & Traps ⚠️
* No zeros.
* All zeros.
* Empty array.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 2. Sort Array of 0s, 1s, and 2s

```
------------------------------------
| Problem Title -> Sort Colors     |
```

### 1️⃣ Problem Snapshot
Sort an array containing only 0, 1, and 2 in-place (DNF Problem).
`[2, 0, 2, 1, 1, 0]` -> `[0, 0, 1, 1, 2, 2]`.

### 2️⃣ Pattern / Category ⭐
**Dutch National Flag (3 Pointers)**

### 3️⃣ Brute Force Idea
Sort O(N log N) or Count Sort (2 passes) O(N).

### 4️⃣ Key Insight (AHA 💡)
Use 3 pointers: `low` (end of 0s), `mid` (current element), `high` (start of 2s).
- 0: swap `arr[low]`, `arr[mid]`. `low++`, `mid++`.
- 1: `mid++`.
- 2: swap `arr[mid]`, `arr[high]`. `high--`.

### 5️⃣ Algorithm
1. `low = 0`, `mid = 0`, `high = N-1`.
2. While `mid <= high`:
   - `if arr[mid] == 0`: swap(low, mid), low++, mid++.
   - `if arr[mid] == 1`: mid++.
   - `if arr[mid] == 2`: swap(mid, high), high--.

### 6️⃣ Edge Cases & Traps ⚠️
* `mid <= high` is the loop condition (don't forget equals).
* Swapping 2 doesn't increment `mid` because the swapped value from `high` could be anything (0, 1, or 2) and needs re-checking.

### 7️⃣ Time & Space Complexity
> Time: O(N) (One pass)
> Space: O(1)

---

## 3. Find Subarray with Given Sum (positive numbers)

```
------------------------------------
| Problem Title -> Subarray Sum    |
```

### 1️⃣ Problem Snapshot
Given positive integers, find a contiguous subarray that sums to `target`.

### 2️⃣ Pattern / Category ⭐
**Variable Size Sliding Window**

### 3️⃣ Brute Force Idea
Nested loops O(N^2).

### 4️⃣ Key Insight (AHA 💡)
Since numbers are positive, extending window strictly increases sum, shrinking strictly decreases sum.
`current < target` -> Expand Right.
`current > target` -> Shrink Left.

### 5️⃣ Algorithm
1. `left = 0`, `currentSum = 0`.
2. Loop `right` from 0 to `N-1`:
   - `currentSum += arr[right]`.
   - While `currentSum > target` & `left <= right`:
     - `currentSum -= arr[left]`, `left++`.
   - If `currentSum == target`: return `[left, right]`.

### 6️⃣ Edge Cases & Traps ⚠️
* Only positives (this fails with negatives).
* No solution.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 4. Find Maximum Sum Subarray

```
------------------------------------
| Problem Title -> Kadane's Algo   |
```

### 1️⃣ Problem Snapshot
Find the contiguous subarray with the largest sum. (Can contain negatives).

### 2️⃣ Pattern / Category ⭐
**Kadane’s Algorithm (Greedy/DP)**

### 3️⃣ Brute Force Idea
Check all subarrays O(N^2).

### 4️⃣ Key Insight (AHA 💡)
If `currentRunningSum` becomes negative, it resets to 0 because carrying a negative sum only hurts future subarrays. But if all numbers are negative, return the max element.

### 5️⃣ Algorithm
1. `maxSum = arr[0]`, `currentSum = 0`.
2. Loop `x` in `arr`:
   - `currentSum += x`.
   - `maxSum = max(maxSum, currentSum)`.
   - If `currentSum < 0`: `currentSum = 0`.
3. Return `maxSum`.

### 6️⃣ Edge Cases & Traps ⚠️
* All negative numbers (Need to ensure `maxSum` isn't stuck at 0 if initialized to 0. Initialize with `arr[0]` or `-INF`).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 5. Find Longest Subarray with Sum = K (Positives & Negatives)

```
------------------------------------
| Problem Title -> Longest k-Sum   |
```

### 1️⃣ Problem Snapshot
Find length of longest subarray summing to K. Handles negative numbers.

### 2️⃣ Pattern / Category ⭐
**Prefix Sum + Hash Map**

### 3️⃣ Brute Force Idea
O(N^2).

### 4️⃣ Key Insight (AHA 💡)
`Sum(i, j) = PrefixSum[j] - PrefixSum[i]`.
We need `PrefixSum[j] - PrefixSum[i] = K`
=> `PrefixSum[i] = PrefixSum[j] - K`.
Store `PrefixSum` in Map. For each `currentSum`, check if `currentSum - K` exists in map.

### 5️⃣ Algorithm
1. `map = {0: -1}` (Sum 0 at index -1).
2. `sum = 0`, `maxLen = 0`.
3. Loop `i` from 0 to `N-1`:
   - `sum += arr[i]`.
   - If `(sum - K)` in map: `maxLen = max(maxLen, i - map[sum-K])`.
   - If `sum` NOT in map: `map[sum] = i` (Keep earliest index for longest length).

### 6️⃣ Edge Cases & Traps ⚠️
* `K = 0`.
* Subarray starting from index 0 (handled by `map[0] = -1`).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 6. Smallest Subarray with Sum > X

```
------------------------------------
| Problem Title -> Min Subarray > X|
```

### 1️⃣ Problem Snapshot
Find minimal length subarray where sum > X. (Positive integers).

### 2️⃣ Pattern / Category ⭐
**Sliding Window (Shrinkable)**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
Expand window until sum > X. Then shrink from the left as much as possible while keeping sum > X to find minimum length.

### 5️⃣ Algorithm
1. `left = 0`, `sum = 0`, `minLen = INF`.
2. Loop `right` from 0 to `N-1`:
   - `sum += arr[right]`.
   - While `sum > X`:
     - `minLen = min(minLen, right - left + 1)`.
     - `sum -= arr[left]`.
     - `left++`.

### 6️⃣ Edge Cases & Traps ⚠️
* No such subarray (return 0).
* Single element > X.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 7. Longest Subarray with Equal 0s and 1s

```
------------------------------------
| Problem Title -> Equal 0s & 1s   |
```

### 1️⃣ Problem Snapshot
Binary array (0s and 1s). Find max length with equal number of 0s and 1s.

### 2️⃣ Pattern / Category ⭐
**Prefix Sum (Transform 0 to -1)**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
If we treat 0 as -1, then "equal 0s and 1s" means "Sum = 0".
Problem reduces to "Longest Subarray with Sum = 0".

### 5️⃣ Algorithm
1. Replace 0 with -1 virtually.
2. Use "Longest Subarray with Sum K" logic where `K=0`.
3. `map = {0: -1}`, `sum = 0`, `maxLen = 0`.
4. Loop `i` from 0 to `N-1`:
   - `val = arr[i] == 0 ? -1 : 1`.
   - `sum += val`.
   - If `sum` in map: `maxLen = max(maxLen, i - map[sum])`.
   - Else: `map[sum] = i`.

### 6️⃣ Edge Cases & Traps ⚠️
* No 0s or No 1s.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)
