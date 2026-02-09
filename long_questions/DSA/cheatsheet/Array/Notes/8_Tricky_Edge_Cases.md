# 8. Tricky & Edge-Case Problems (The Final 5%)

---

## 1. Shortest Unsorted Continuous Subarray

```
------------------------------------
| Problem Title -> Shortest Unsort |
```

### 1️⃣ Problem Snapshot
Find the length of the shortest subarray that, if sorted, makes the whole array sorted.
`[2, 6, 4, 8, 10, 9, 15]` -> 5 (`[6, 4, 8, 10, 9]`).

### 2️⃣ Pattern / Category ⭐
**Sort Comparison OR Two Pass Scan**

### 3️⃣ Brute Force Idea
Copy array, sort it, comparing indices. First mismatch is `start`, last mismatch is `end`. O(N log N).

### 4️⃣ Key Insight (AHA 💡)
The "unsorted" part is bounded by the first element (from left) that is smaller than `max_seen_so_far` and the first element (from right) that is larger than `min_seen_so_far`.
(Actually: Scan Left->Right to find `end` (where `curr < max`). Scan Right->Left to find `start` (where `curr > min`)).

### 5️⃣ Algorithm
1. `maxVal = -INF`, `end = -1`.
2. Loop `i` from 0 to N-1:
   - `maxVal = max(maxVal, arr[i])`.
   - If `arr[i] < maxVal`: `end = i`.
3. `minVal = INF`, `start = 0`.
4. Loop `i` from N-1 to 0:
   - `minVal = min(minVal, arr[i])`.
   - If `arr[i] > minVal`: `start = i`.
5. If `end == -1`: return 0.
6. Return `end - start + 1`.

### 6️⃣ Edge Cases & Traps ⚠️
* Already sorted (end remains -1).
* Reverse sorted.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 2. Check if Sorted by One Swap

```
------------------------------------
| Problem Title -> One Swap Sort   |
```

### 1️⃣ Problem Snapshot
Can you make the array sorted by swapping exactly two elements?
`[1, 5, 3, 3, 7]` -> true (Swap 5 and 3).

### 2️⃣ Pattern / Category ⭐
**Identify Inversions**

### 3️⃣ Brute Force Idea
Try all swaps. O(N^3).

### 4️⃣ Key Insight (AHA 💡)
Find all pairs `(i, i+1)` where order breaks (`arr[i] > arr[i+1]`).
If more than 2 breaks, impossible.
If 1 break: Swap elements.
If 2 breaks: Swap first element of 1st break and second element of 2nd break.
Check if sorted after swap.

### 5️⃣ Algorithm
1. Count inversions `arr[i] > arr[i+1]`.
2. If count > 2: return False.
3. If count == 0: return True.
4. Perform the logical swap based on inversion indices.
5. Verify if sorted.

### 6️⃣ Edge Cases & Traps ⚠️
* `[1, 3, 2, 4]` (1 break).
* `[4, 3, 2, 1]` (3 breaks -> False).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 3. Check if Sorted by Reversing One Subarray

```
------------------------------------
| Problem Title -> Reverse to Sort |
```

### 1️⃣ Problem Snapshot
Can you sort the array by reversing exactly one contiguous subarray?
`[1, 2, 5, 4, 3, 6]` -> True (Reverse `5, 4, 3`).

### 2️⃣ Pattern / Category ⭐
**Locate Decreasing Segment**

### 3️⃣ Brute Force Idea
N^3.

### 4️⃣ Key Insight (AHA 💡)
A sorted array increases. If we reverse a subarray, we'll see an increasing part, then a "decreasing" part (the reversed bit), then increasing again.
There should be exactly one decreasing subarray. We identify it, reverse it in a copy/logically, and check if sorted.

### 5️⃣ Algorithm
1. Find first index `start` where `arr[i] > arr[i+1]`.
2. If no such index, True.
3. Find last index `end` where `arr[j] < arr[j-1]`.
4. Reverse `arr[start...end]`.
5. Check if whole array sorted.

### 6️⃣ Edge Cases & Traps ⚠️
* Multiple peaks/valleys (False).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1) (if in-place check)

---

## 4. First Missing Positive Integer

```
------------------------------------
| Problem Title -> First Missing + |
```

### 1️⃣ Problem Snapshot
Unsorted integer array. Find smallest missing positive integer.
`[3, 4, -1, 1]` -> 2. `[1, 2, 0]` -> 3.

### 2️⃣ Pattern / Category ⭐
**Cyclic Sort / Index Mapping**

### 3️⃣ Brute Force Idea
Sort O(N log N).

### 4️⃣ Key Insight (AHA 💡)
We only care about numbers `1` to `N`.
We can place `x` at index `x-1` (i.e., `1` at `0`, `2` at `1`).
Ignore negatives and numbers > N.

### 5️⃣ Algorithm
1. Loop `i` from 0 to N-1:
   - While `arr[i]` in `[1, N]` AND `arr[arr[i]-1] != arr[i]`:
     - Swap `arr[i]` with `arr[arr[i]-1]`.
2. Loop `i` from 0 to N-1:
   - If `arr[i] != i + 1`, return `i + 1`.
3. Return `N + 1`.

### 6️⃣ Edge Cases & Traps ⚠️
* Duplicates (logic `arr[arr[i]-1] != arr[i]` handles infinite loop).
* All negatives (return 1).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1) (In-place)

---

## 5. Elements Appearing > N/K Times

```
------------------------------------
| Problem Title -> N/K Frequency   |
```

### 1️⃣ Problem Snapshot
Find all elements appearing more than `N/K` times. (Generalized Majority).

### 2️⃣ Pattern / Category ⭐
**Tetris / Generalized Moore Voting**

### 3️⃣ Brute Force Idea
Hash Map O(N) space.

### 4️⃣ Key Insight (AHA 💡)
There can be at most `K-1` such elements.
Maintain `K-1` candidates and their counts.
If a new element matches a candidate, `count++`.
Start new candidate if slot empty.
If all slots full, decrement ALL counts (cancel out).

### 5️⃣ Algorithm
1. `Map candidates` (size K-1).
2. Process array:
   - Match? `count++`.
   - Empty slot? `Insert`.
   - Full? `Decrement All`. Remove 0 counts.
3. Verification Pass: Count actual occurrences of candidates.

### 6️⃣ Edge Cases & Traps ⚠️
* K is large (Use HashMap, then it becomes trivial O(N) space).
* Candidates might be false positives (Verification pass is mandatory).

### 7️⃣ Time & Space Complexity
> Time: O(N * K)
> Space: O(K)

---

## 6. Max Sum Such That No Two Elements Are Adjacent

```
------------------------------------
| Problem Title -> House Robber    |
```

### 1️⃣ Problem Snapshot
Find max sum of subsequence where no two indices are adjacent.
`[1, 2, 3, 1]` -> 4 (`1 + 3`).

### 2️⃣ Pattern / Category ⭐
**Dynamic Programming**

### 3️⃣ Brute Force Idea
Recursion O(2^N).

### 4️⃣ Key Insight (AHA 💡)
At index `i`, we have two choices:
1. Include `arr[i]`: Then we limit sum to `prev2 + arr[i]`.
2. Exclude `arr[i]`: Then we take `prev1`.
`dp[i] = max(dp[i-1], arr[i] + dp[i-2])`.

### 5️⃣ Algorithm
1. `prev2 = 0`, `prev1 = arr[0]`.
2. Loop `i` from 1 to N-1:
   - `pick = arr[i] + prev2`.
   - `skip = prev1`.
   - `curr = max(pick, skip)`.
   - `prev2 = prev1`, `prev1 = curr`.
3. Return `prev1`.

### 6️⃣ Edge Cases & Traps ⚠️
* Single element.
* Empty array.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 7. Subarray Sum = 0 (Space Optimized?)

```
------------------------------------
| Problem Title -> Zero Sum Subarr |
```

### 1️⃣ Problem Snapshot
Does a subarray with sum 0 exist?
Try to solve with O(1) space if possible (Wait... usually O(N) space map is the answer).
*Special Case:* If elements are small integers, or if we modify array.
(Assuming standard interview Context: "Can we do better than Map O(N)?")

### 2️⃣ Pattern / Category ⭐
**Prefix Sums & Sorting**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
If we cannot use Map, valid approach is:
1. Calculate Prefix Sums (Modify array or new array).
2. sort Prefix Sums.
3. If `P[i] == P[i+1]`, then sum between them is 0.

### 5️⃣ Algorithm
1. Loop `i=1` to N-1: `arr[i] += arr[i-1]` (Prefix Sum).
2. Sort `arr`.
3. Check `arr[i] == arr[i+1]` OR `arr[i] == 0`.
4. If yes, return True.

### 6️⃣ Edge Cases & Traps ⚠️
* Modifies array.
* Iterate to check duplicates.

### 7️⃣ Time & Space Complexity
> Time: O(N log N)
> Space: O(1) (if in-place sort)

---

## 8. Rearrange such that arr[i] = i

```
------------------------------------
| Problem Title -> Index Mapping   |
```

### 1️⃣ Problem Snapshot
Given array with values `-1` to `N-1`.
If `x` exists in array, place it at `arr[x]`. Put `-1` elsewhere.

### 2️⃣ Pattern / Category ⭐
**Placement / Cyclic Swap**

### 3️⃣ Brute Force Idea
New array O(N).

### 4️⃣ Key Insight (AHA 💡)
Iterate through. If `arr[i] != -1` and `arr[i] != i`:
Swap `arr[i]` with `arr[arr[i]]`.
Repeat until `arr[i]` is correct or `-1`.

### 5️⃣ Algorithm
1. Loop `i` from 0 to N-1:
   - While `arr[i] != -1` AND `arr[i] != i`:
     - Swap `x = arr[i]` with `arr[x]`.

### 6️⃣ Edge Cases & Traps ⚠️
* Cycles (handled by logic).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)
