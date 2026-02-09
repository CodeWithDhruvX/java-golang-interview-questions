# 6. Sorting-Based Interview Problems

---

## 1. Merge Two Sorted Arrays

```
------------------------------------
| Problem Title -> Merge Sorted    |
```

### 1️⃣ Problem Snapshot
Merge two sorted arrays into one sorted array.
`[1, 3, 5]` & `[2, 4, 6]` -> `[1, 2, 3, 4, 5, 6]`.

### 2️⃣ Pattern / Category ⭐
**Two Pointers (Merge Step)**

### 3️⃣ Brute Force Idea
Combine into one array, then sort. O((N+M) log (N+M)).

### 4️⃣ Key Insight (AHA 💡)
Since arrays are already sorted, looking at the heads of both arrays gives the next smallest element.
Comparing `A[i]` and `B[j]` is sufficient.

### 5️⃣ Algorithm
1. `i = 0`, `j = 0`, `k = 0`.
2. `result` array of size `N+M`.
3. While `i < N` & `j < M`:
   - If `A[i] <= B[j]`: `result[k++] = A[i++]`.
   - Else: `result[k++] = B[j++]`.
4. Copy remaining elements from A or B.

### 6️⃣ Edge Cases & Traps ⚠️
* One array empty.
* Arrays of different sizes.
* Merge in-place (Harder variation - GAP method/Shell sort logic).

### 7️⃣ Time & Space Complexity
> Time: O(N + M)
> Space: O(N + M)

---

## 2. Median of Two Sorted Arrays

```
------------------------------------
| Problem Title -> Median Sorted   |
```

### 1️⃣ Problem Snapshot
Find median of two combined sorted arrays of size N and M.
`[1, 3]`, `[2]` -> Median 2.

### 2️⃣ Pattern / Category ⭐
**Binary Search on Partition**

### 3️⃣ Brute Force Idea
Merge arrays (O(N+M)), then pick middle. Space O(N+M).

### 4️⃣ Key Insight (AHA 💡)
We effectively want to partition both arrays such that:
1. `LeftPart` size == `RightPart` size.
2. `Max(LeftPart) <= Min(RightPart)`.
Binary search on the smaller array's cut position can find this partition.

### 5️⃣ Algorithm
1. BS on smaller array `A` (size N).
2. `partitionA = mid`.
3. `partitionB = (N + M + 1)/2 - partitionA`.
4. Check valid partition conditions:
   - `maxLeftA <= minRightB` && `maxLeftB <= minRightA`.
5. If valid: Calculate median from boundary values.
6. Else: Adjust BS range.

### 6️⃣ Edge Cases & Traps ⚠️
* Partition at index 0 or N (use `-INF` / `+INF`).
* Total size odd/even logic.

### 7️⃣ Time & Space Complexity
> Time: O(log(min(N, M)))
> Space: O(1)

---

## 3. Sort Array by Frequency

```
------------------------------------
| Problem Title -> Frequency Sort  |
```

### 1️⃣ Problem Snapshot
Sort elements based on frequency (descending). If frequency same, sort by value (ascending or original order).
`[1, 1, 2, 2, 2, 3]` -> `[2, 2, 2, 1, 1, 3]`.

### 2️⃣ Pattern / Category ⭐
**Hashing + Custom Sorting**

### 3️⃣ Brute Force Idea
Count freq, then bubble sort with custom check. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
Precompute frequencies. Then sort using a custom comparator that checks Frequency first, then Value.

### 5️⃣ Algorithm
1. Count `freq` using Map.
2. Defines slice `nums`.
3. `sort.Slice(nums, func(i, j int) {`
   - `if freq[nums[i]] == freq[nums[j]] return nums[i] < nums[j]`
   - `return freq[nums[i]] > freq[nums[j]]`
   `})`

### 6️⃣ Edge Cases & Traps ⚠️
* All frequencies same (behaves like normal sort).
* Unstable sort might change relative order of same values (Go's `sort.Slice` is not stable, `sort.SliceStable` is).

### 7️⃣ Time & Space Complexity
> Time: O(N log N)
> Space: O(N) (Map)

---

## 4. Minimum Swaps to Sort Array

```
------------------------------------
| Problem Title -> Min Swaps       |
```

### 1️⃣ Problem Snapshot
Find min swaps required to sort an array.
`[4, 3, 2, 1]` -> 2.

### 2️⃣ Pattern / Category ⭐
**Cycle Decomposition**

### 3️⃣ Brute Force Idea
Simulate sorting (Selection Sort count swaps). O(N^2).

### 4️⃣ Key Insight (AHA 💡)
The permutation can be decomposed into disjoint cycles.
To resolve a cycle of `k` nodes, we need `k-1` swaps.
Answer = `Sum(CycleLength - 1)` for all cycles.

### 5️⃣ Algorithm
1. Create `pairs` of `(val, original_index)`. Sort `pairs` by value.
2. `visited` array false.
3. `swaps = 0`.
4. Loop `i` from 0 to N-1:
   - If `visited[i]` or `pairs[i].index == i` continue.
   - Trace cycle: `j = i`, `cycle_len = 0`.
   - While `!visited[j]`:
     - `visited[j] = true`
     - `j = pairs[j].index`
     - `cycle_len++`
   - `swaps += cycle_len - 1`.

### 6️⃣ Edge Cases & Traps ⚠️
* Already sorted (0 swaps).

### 7️⃣ Time & Space Complexity
> Time: O(N log N)
> Space: O(N)

---

## 5. Count Inversions in Array

```
------------------------------------
| Problem Title -> Inversion Count |
```

### 1️⃣ Problem Snapshot
Count pairs `(i, j)` such that `i < j` and `arr[i] > arr[j]`.
Measure of "unsortedness".

### 2️⃣ Pattern / Category ⭐
**Merge Sort Modification**

### 3️⃣ Brute Force Idea
Nested loops. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
During Merge step of Merge Sort:
If `LeftArr[i] > RightArr[j]`, then `LeftArr[i]` IS GREATER than `RightArr[j]`.
AND... since `LeftArr` is sorted, ALL elements after `LeftArr[i]` are ALSO greater than `RightArr[j]`.
Add `(mid - i + 1)` to count.

### 5️⃣ Algorithm
1. `count = 0`.
2. `mergeSort(arr, low, high)`:
   - `mid = ...`
   - `count += mergeSort(low, mid)`
   - `count += mergeSort(mid+1, high)`
   - `count += merge(low, mid, high)`
3. Inside `merge`:
   - If `arr[i] > arr[j]`:
     - `inv_count += (mid - i + 1)`
     - `arr[k] = arr[j], j++`

### 6️⃣ Edge Cases & Traps ⚠️
* Reverse sorted array (max inversions N*(N-1)/2).

### 7️⃣ Time & Space Complexity
> Time: O(N log N)
> Space: O(N)

---

## 6. Chocolate Distribution Problem

```
------------------------------------
| Problem Title -> Min Diff Packet |
```

### 1️⃣ Problem Snapshot
Given N packets of chocolates (values array), distribute to M students such that:
1. Each gets 1 packet.
2. Diff between max packet and min packet is minimized.

### 2️⃣ Pattern / Category ⭐
**Sort + Sliding Window (Fixed K)**

### 3️⃣ Brute Force Idea
Check all combinations of size M. Exponential.

### 4️⃣ Key Insight (AHA 💡)
If we pick a subset of size M, to minimize `Max - Min`, the subset elements MUST be close to each other in value.
Wait! This means if we SORT the array, the best subset must be a contiguous subarray of size M.
`Diff = arr[i+M-1] - arr[i]`.

### 5️⃣ Algorithm
1. Sort `arr`.
2. `minDiff = INF`.
3. Loop `i` from 0 to `N-M`:
   - `diff = arr[i+M-1] - arr[i]`.
   - `minDiff = min(minDiff, diff)`.

### 6️⃣ Edge Cases & Traps ⚠️
* M > N (Invalid).
* M = 0.

### 7️⃣ Time & Space Complexity
> Time: O(N log N)
> Space: O(1)

---

## 7. Largest Number from Array

```
------------------------------------
| Problem Title -> Largest Number  |
```

### 1️⃣ Problem Snapshot
Arrange numbers to form the largest possible string.
`[3, 30, 34, 5, 9]` -> `"9534330"`.

### 2️⃣ Pattern / Category ⭐
**Custom Sorting Comparator**

### 3️⃣ Brute Force Idea
Permutations. O(N!).

### 4️⃣ Key Insight (AHA 💡)
This is not standard sort.
Example `3` vs `30`: `330` > `303`, so `3` comes before `30`.
Rule: Sort `A` before `B` if `A+B > B+A` (string concatenation).

### 5️⃣ Algorithm
1. Convert integers to strings.
2. Sort strings with comparator:
   - `return (str(a) + str(b)) > (str(b) + str(a))`
3. Join sorted strings.
4. Handle leading "0" case (if result is "000", return "0").

### 6️⃣ Edge Cases & Traps ⚠️
* `[0, 0]` -> Should return "0", not "00".
* Large numbers (Use string comparisons).

### 7️⃣ Time & Space Complexity
> Time: O(N log N) (assuming string constant length)
> Space: O(N)
