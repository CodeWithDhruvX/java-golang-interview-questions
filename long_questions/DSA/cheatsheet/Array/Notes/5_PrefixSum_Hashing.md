# 5. Prefix Sum & Hashing Based Problems

---

## 1. Range Sum Queries (Immutable Array)

```
------------------------------------
| Problem Title -> Range Sum Query |
```

### 1️⃣ Problem Snapshot
Given an array, handle multiple queries: `Sum(i, j)` = sum of elements from index `i` to `j`.

### 2️⃣ Pattern / Category ⭐
**Prefix Sum Array**

### 3️⃣ Brute Force Idea
Loop from `i` to `j` for every query. O(N) per query.

### 4️⃣ Key Insight (AHA 💡)
Precompute cumulative sums.
`P[x] = Sum(0...x)`.
`Sum(i, j) = P[j] - P[i-1]`.
Query becomes O(1).

### 5️⃣ Algorithm
1. Create `P` of size `N`. `P[0] = arr[0]`.
2. Loop 1 to N-1: `P[i] = P[i-1] + arr[i]`.
3. For query `(i, j)`:
   - If `i == 0`: return `P[j]`.
   - Else: return `P[j] - P[i-1]`.

### 6️⃣ Edge Cases & Traps ⚠️
* `i = 0` (Formula `P[i-1]` would crash).
* Updates to array? (Segment Tree required if immutable constraint removed).

### 7️⃣ Time & Space Complexity
> Time: O(N) build, O(1) query
> Space: O(N)

---

## 2. Find Subarray with Sum = 0

```
------------------------------------
| Problem Title -> Zero Sum Subarr |
```

### 1️⃣ Problem Snapshot
Check if any subarray has sum 0.
`[4, 2, -3, 1, 6]` -> Yes (`2, -3, 1`).

### 2️⃣ Pattern / Category ⭐
**Prefix Sum + Hashing**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
If `PrefixSum` repeats (e.g., `P[x] == P[y]`), then the sum of elements between `x` and `y` must be 0.
Also, if any `P[x] == 0`, subarray `0...x` is the answer.

### 5️⃣ Algorithm
1. `Set visited`.
2. `sum = 0`.
3. Loop `x` in `arr`:
   - `sum += x`.
   - If `sum == 0` or `sum in visited`: return `True`.
   - `visited.add(sum)`.
4. Return `False`.

### 6️⃣ Edge Cases & Traps ⚠️
* Single element 0.
* Entire array sum 0.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 3. Count Subarrays with Sum = K

```
------------------------------------
| Problem Title -> Count Sum=K     |
```

### 1️⃣ Problem Snapshot
Find the total number of subarrays whose sum equals `K`.
`[1, 1, 1]`, K=2 -> 2.

### 2️⃣ Pattern / Category ⭐
**Prefix Sum + Hash Map Frequency**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
We need `CurrentSum - OldSum = K` -> `OldSum = CurrentSum - K`.
Check how many times `CurrentSum - K` has appeared before. Add that count to result.

### 5️⃣ Algorithm
1. `Map freq = {0: 1}` (Base case).
2. `sum = 0`, `count = 0`.
3. Loop `x` in `arr`:
   - `sum += x`.
   - `count += freq[sum - K]` (if exists).
   - `freq[sum]++`.
4. Return `count`.

### 6️⃣ Edge Cases & Traps ⚠️
* `K = 0` (Count zeroes properly).
* `freq` should be initialized with `{0: 1}` to handle subarrays starting at index 0.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 4. Longest Subarray with Distinct Elements

```
------------------------------------
| Problem Title -> Longest Unique  |
```

### 1️⃣ Problem Snapshot
Find length of longest subarray with all unique elements.
`[1, 2, 3, 1, 4]` -> 4 (`2, 3, 1, 4`).

### 2️⃣ Pattern / Category ⭐
**Sliding Window + Last Index Map**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
Expand `right`. If `arr[right]` was seen at `prevIndex` and `prevIndex >= left`, move `left` to `prevIndex + 1` to exclude the duplicate.
Update `maxLen` at every step.

### 5️⃣ Algorithm
1. `Map lastSeen`.
2. `left = 0`, `maxLen = 0`.
3. Loop `right` from 0 to N-1:
   - If `arr[right]` in `lastSeen`:
     - `left = max(left, lastSeen[arr[right]] + 1)`.
   - `maxLen = max(maxLen, right - left + 1)`.
   - `lastSeen[arr[right]] = right`.

### 6️⃣ Edge Cases & Traps ⚠️
* `left` should not move backwards (use `max`).
* Empty array.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 5. Subarray with Given XOR

```
------------------------------------
| Problem Title -> Subarray XOR    |
```

### 1️⃣ Problem Snapshot
Count subarrays where XOR of elements = `B`.
`[4, 2, 2, 6, 4]`, B=6 -> 4.

### 2️⃣ Pattern / Category ⭐
**Prefix XOR + Hash Map**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
Identical to "Sum = K" but with XOR.
`PrefixXOR[i] ^ PrefixXOR[j] = B`
=> `PrefixXOR[i] ^ B = PrefixXOR[j]`.
Look for `CurrentXOR ^ B` in map.

### 5️⃣ Algorithm
1. `Map freq = {0: 1}`.
2. `xorSum = 0`, `count = 0`.
3. Loop `x` in `arr`:
   - `xorSum ^= x`.
   - `count += freq[xorSum ^ B]`.
   - `freq[xorSum]++`.

### 6️⃣ Edge Cases & Traps ⚠️
* `B = 0` (Subarrays with XOR 0).
* Large numbers (XOR doesn't overflow like Sum).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 6. Count Subarrays with Equal Odd and Even Numbers

```
------------------------------------
| Problem Title -> Odd/Even Equal  |
```

### 1️⃣ Problem Snapshot
Find number of subarrays with equal count of odd and even integers.

### 2️⃣ Pattern / Category ⭐
**Transformation + Zero Sum Subarray**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
Transform array: `Odd -> 1`, `Even -> -1`.
Now problem is "Count Subarrays with Sum = 0".
Using Map as seen in Problem 2 & 3.

### 5️⃣ Algorithm
1. `Map freq = {0: 1}`.
2. `sum = 0`, `res = 0`.
3. Loop `x` in `arr`:
   - `val = (x % 2 != 0) ? 1 : -1`.
   - `sum += val`.
   - `res += freq[sum]`.
   - `freq[sum]++`.

### 6️⃣ Edge Cases & Traps ⚠️
* `Sum` can be negative (Map handles this).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)
