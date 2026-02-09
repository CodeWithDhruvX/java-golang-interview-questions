# 7. Advanced / Product Company Level Problems

---

## 1. Maximum Product Subarray

```
------------------------------------
| Problem Title -> Max Product     |
```

### 1️⃣ Problem Snapshot
Find contiguous subarray with largest product.
`[2, 3, -2, 4]` -> `6` (`2*3`).
`[-2, 0, -1]` -> `0`.

### 2️⃣ Pattern / Category ⭐
**Modified Kadane (Track Max & Min)**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
A negative number flips max to min and min to max.
We must maintain BOTH `currentMax` and `currentMin` because a large negative `currentMin` multiplied by a negative number becomes a huge positive `currentMax`.

### 5️⃣ Algorithm
1. `res = max(nums)`. `curMax = 1`, `curMin = 1`.
2. Loop `n` in `nums`:
   - If `n == 0`: reset `curMax=1`, `curMin=1`.
   - `tmp = curMax * n`.
   - `curMax = max(n, n*curMax, n*curMin)`.
   - `curMin = min(n, tmp, n*curMin)`.
   - `res = max(res, curMax)`.

### 6️⃣ Edge Cases & Traps ⚠️
* Zeros in array (reset product).
* Negative numbers.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 2. Trapping Rain Water

```
------------------------------------
| Problem Title -> Trap Rain Water |
```

### 1️⃣ Problem Snapshot
Given elevation map (array), calculate how much water it can trap after raining.
`[0,1,0,2,1,0,1,3,2,1,2,1]` -> 6.

### 2️⃣ Pattern / Category ⭐
**Precompute MaxLeft/MaxRight OR Two Pointers**

### 3️⃣ Brute Force Idea
For each `i`, find max left and max right. `min(maxL, maxR) - height[i]`. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
Water at `i` is determined by the bottleneck: `min(LeftMax, RightMax) - Height[i]`.
We can compute `LeftMax` and `RightMax` arrays in O(N).
Better: Two Pointers. Maintain `leftMax` and `rightMax` and move the smaller pointer.

### 5️⃣ Algorithm (Two Pointers)
1. `l=0, r=N-1`, `maxL=0, maxR=0`.
2. While `l < r`:
   - If `height[l] < height[r]`:
     - If `height[l] >= maxL`: `maxL = height[l]`.
     - Else: `res += maxL - height[l]`.
     - `l++`.
   - Else:
     - (Symmetric logic for R).

### 6️⃣ Edge Cases & Traps ⚠️
* Ascending/Descending/V-shape (No water).
* Fewer than 3 bars.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1) (Two Pointers)

---

## 3. Container With Most Water

```
------------------------------------
| Problem Title -> Max Water Area  |
```

### 1️⃣ Problem Snapshot
Find two lines that together with the x-axis form a container, such that the container contains the most water. `Area = min(h1, h2) * width`.

### 2️⃣ Pattern / Category ⭐
**Two Pointers (Greedy)**

### 3️⃣ Brute Force Idea
Check all pairs. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
Start with max width (Left=0, Right=N-1).
To find a potentially larger area, we MUST overcome the width reduction by finding a taller line.
The shorter line limits the height. Discard the shorter line (move pointer inward).

### 5️⃣ Algorithm
1. `l=0`, `r=N-1`. `maxArea=0`.
2. While `l < r`:
   - `area = (r - l) * min(height[l], height[r])`.
   - `maxArea = max(maxArea, area)`.
   - If `height[l] < height[r]`: `l++`.
   - Else: `r--`.

### 6️⃣ Edge Cases & Traps ⚠️
* All heights equal.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 4. Subarray Sum Divisible by K

```
------------------------------------
| Problem Title -> Sum Div by K    |
```

### 1️⃣ Problem Snapshot
Count subarrays where sum is divisible by K.
`[4, 5, 0, -2, -3, 1]`, K=5 -> 7.

### 2️⃣ Pattern / Category ⭐
**Prefix Sum % K + HashMap**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
`Sum(i, j)` is divisible by K if `PrefixSum[j] % K == PrefixSum[i] % K`.
Use Map to count occurrences of remainders.
Handle negative remainders: `rem = (sum % K + K) % K`.

### 5️⃣ Algorithm
1. `map = {0: 1}`.
2. `sum = 0`, `count = 0`.
3. Loop `x` in `nums`:
   - `sum += x`.
   - `rem = (sum % K + K) % K`.
   - `count += map[rem]`.
   - `map[rem]++`.

### 6️⃣ Edge Cases & Traps ⚠️
* Negative values (Must correct modulo).
* K=1 (All subarrays divisible).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(K)

---

## 5. Maximum Circular Subarray Sum

```
------------------------------------
| Problem Title -> Circular Max Sum|
```

### 1️⃣ Problem Snapshot
Find max subarray sum in a **circular** array (end connects to start).

### 2️⃣ Pattern / Category ⭐
**Kadane’s (Max & Min)**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
Two cases:
1. Max subarray is non-circular (Normal Kadane).
2. Max subarray wraps around. This means the *Minimum* subarray is in the middle (non-circular).
   `MaxCircular = TotalSum - MinSubarraySum`.
Note: If all numbers are negative, `MaxCircular` logic returns 0 (Total - Total). Handle this by returning normal Max.

### 5️⃣ Algorithm
1. `max_kadane = KadaneMax(arr)`.
2. `min_kadane = KadaneMin(arr)`.
3. `total = Sum(arr)`.
4. If `max_kadane < 0`: return `max_kadane`. (All negatives).
5. Return `max(max_kadane, total - min_kadane)`.

### 6️⃣ Edge Cases & Traps ⚠️
* All negative array (Total - Min = 0, which is wrong, should return Max element).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 6. Longest Consecutive Sequence

```
------------------------------------
| Problem Title -> Longest Seq     |
```

### 1️⃣ Problem Snapshot
Find length of longest elements sequence (`1, 2, 3...`) in unsorted array.
`[100, 4, 200, 1, 3, 2]` -> 4 (`1, 2, 3, 4`). O(N) required.

### 2️⃣ Pattern / Category ⭐
**HashSet**

### 3️⃣ Brute Force Idea
Sort O(N log N).

### 4️⃣ Key Insight (AHA 💡)
Put all in Set.
For each `x`: check if `x-1` exists in Set.
- If YES: `x` is NOT the start of a sequence. Skip.
- If NO: `x` IS the start. Count `x+1, x+2...` sequence.

### 5️⃣ Algorithm
1. `set = Set(arr)`.
2. `maxLen = 0`.
3. Loop `x` in `set`:
   - If `x-1` not in set:
     - `currentNum = x`.
     - `curLen = 1`.
     - While `currentNum + 1` in set: `currentNum++`, `curLen++`.
     - `maxLen = max(maxLen, curLen)`.

### 6️⃣ Edge Cases & Traps ⚠️
* Empy array.
* Duplicates (Set handles this).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 7. Count Smaller Elements on Right

```
------------------------------------
| Problem Title -> Count Smaller   |
```

### 1️⃣ Problem Snapshot
For each element, count how many smaller elements are to its right.
`[5, 2, 6, 1]` -> `[2, 1, 1, 0]`.

### 2️⃣ Pattern / Category ⭐
**Merge Sort / Fenwick Tree**

### 3️⃣ Brute Force Idea
N^2.

### 4️⃣ Key Insight (AHA 💡)
Use Merge Sort. When merging, if `Left[i] > Right[j]`, it means `Right[j]` is smaller than `Left[i]` (and all subsequent Lefts). But to count for specific indices, we need to track indices.
Alternatively, iterate from right to left, insert into sorted structure (Fenwick/BST) and query "Sum of elements < x".

### 5️⃣ Algorithm (Merge Sort)
1. Store pairs `(val, index)`.
2. Perform Merge Sort.
3. In Merge step: While `right[j] < left[i]`: `j++`.
4. `count[left[i].index] += j` (number of elements jumped over).
(Requires careful implementation).

### 6️⃣ Edge Cases & Traps ⚠️
* Duplicates.

### 7️⃣ Time & Space Complexity
> Time: O(N log N)
> Space: O(N)

---

## 8. Stock Buy and Sell (Multiple Transactions)

```
------------------------------------
| Problem Title -> Stock Buy/Sell  |
```

### 1️⃣ Problem Snapshot
Prices array. Buy and sell as many times as you want to maximize profit.
`[7, 1, 5, 3, 6, 4]` -> 7 (1->5, 3->6).

### 2️⃣ Pattern / Category ⭐
**Greedy Valley-Peak**

### 3️⃣ Brute Force Idea
DP O(N) or Recursive O(2^N).

### 4️⃣ Key Insight (AHA 💡)
If `price[i] > price[i-1]`, we capture that profit.
`Profit = Sum(max(0, price[i] - price[i-1]))`.
We just need to capture every upslope.

### 5️⃣ Algorithm
1. `maxProfit = 0`.
2. Loop `i` from 1 to N-1:
   - If `prices[i] > prices[i-1]`:
     - `maxProfit += prices[i] - prices[i-1]`.
3. Return `maxProfit`.

### 6️⃣ Edge Cases & Traps ⚠️
* Always decreasing (Profit 0).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)
