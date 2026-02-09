# 2. Searching & Index-Based Problems

---

## 1. Linear Search

```
------------------------------------
| Problem Title -> Linear Search   |
```

### 1️⃣ Problem Snapshot
Find the index of a target element `x` in an array. Return -1 if not found.

### 2️⃣ Pattern / Category ⭐
**Sequential Scan**

### 3️⃣ Brute Force Idea
This IS the brute force. Check one by one.

### 4️⃣ Key Insight (AHA 💡)
No special property (like sorting) is known, so we must check every element in the worst case.

### 5️⃣ Algorithm
1. Loop `i` from 0 to `N-1`.
2. If `arr[i] == target`, return `i`.
3. If loop ends, return -1.

### 6️⃣ Edge Cases & Traps ⚠️
* Element present multiple times (return first index).
* Element not present.
* Empty array.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 2. Binary Search (Iterative & Recursive)

```
------------------------------------
| Problem Title -> Binary Search   |
```

### 1️⃣ Problem Snapshot
Find the index of `target` in a **sorted** array. O(log N) expected.

### 2️⃣ Pattern / Category ⭐
**Divide and Conquer (Reduce Search Space)**

### 3️⃣ Brute Force Idea
Linear Search O(N).

### 4️⃣ Key Insight (AHA 💡)
If `target > mid`, it MUST be in the right half. If `target < mid`, it MUST be in the left half. We discard half the array each step.

### 5️⃣ Algorithm (Iterative)
1. `low = 0, high = N-1`.
2. While `low <= high`:
   - `mid = low + (high - low) / 2`.
   - If `arr[mid] == target`, return `mid`.
   - If `arr[mid] < target`, `low = mid + 1`.
   - Else `high = mid - 1`.
3. Return -1.

### 6️⃣ Edge Cases & Traps ⚠️
* `mid` calculation overflow: use `low + (high-low)/2` instead of `(low+high)/2`.
* Target smaller/larger than all elements.

### 7️⃣ Time & Space Complexity
> Time: O(log N)
> Space: O(1) (Recursive is O(log N))

---

## 3. Find First and Last Occurrence

```
------------------------------------
| Problem Title -> First/Last Pos  |
```

### 1️⃣ Problem Snapshot
Given a sorted array with duplicates, find the starting and ending position of a given `target`.
`[5, 7, 7, 8, 8, 10]`, target=8 -> `[3, 4]`

### 2️⃣ Pattern / Category ⭐
**Modified Binary Search**

### 3️⃣ Brute Force Idea
Linear scan to find first occurence, continue to find last. O(N).

### 4️⃣ Key Insight (AHA 💡)
Use Binary Search twice.
1. `FindFirst`: When `arr[mid] == target`, don't stop. Store `mid`, moving `high` to `mid-1` to find *earlier* occurrences.
2. `FindLast`: When `arr[mid] == target`, store `mid`, move `low` to `mid+1` to find *later* occurrences.

### 5️⃣ Algorithm
1. Run `BS_First`:
   - If `arr[mid] == target`: `res = mid`, `high = mid - 1`.
2. Run `BS_Last`:
   - If `arr[mid] == target`: `res = mid`, `low = mid + 1`.
3. Return `{first, last}`.

### 6️⃣ Edge Cases & Traps ⚠️
* Target not found (First=-1, Last=-1).
* Array with all same elements.

### 7️⃣ Time & Space Complexity
> Time: O(log N)
> Space: O(1)

---

## 4. Count Occurrences of a Number in Sorted Array

```
------------------------------------
| Problem Title -> Count Occur     |
```

### 1️⃣ Problem Snapshot
Count how many times `x` appears in a sorted array.
`[1, 1, 2, 2, 2, 3]`, x=2 -> Count is 3.

### 2️⃣ Pattern / Category ⭐
**First & Last Occurrence (Binary Search)**

### 3️⃣ Brute Force Idea
Linear count O(N).

### 4️⃣ Key Insight (AHA 💡)
Since it's sorted, all `x` are adjacent. Count = `LastIndex(x) - FirstIndex(x) + 1`.

### 5️⃣ Algorithm
1. Find `first` index using modified BS.
2. If `first == -1`, return 0.
3. Find `last` index using modified BS.
4. Return `last - first + 1`.

### 6️⃣ Edge Cases & Traps ⚠️
* Element not present (0).
* Single element.

### 7️⃣ Time & Space Complexity
> Time: O(log N)
> Space: O(1)

---

## 5. Find Missing Number (1..N)

```
------------------------------------
| Problem Title -> Missing Num     |
```

### 1️⃣ Problem Snapshot
Given an array containing `N` distinct numbers taken from `0, 1, 2, ..., N`, find the one that is missing.
`[3, 0, 1]` -> Missing is 2.

### 2️⃣ Pattern / Category ⭐
**Math Sumation OR XOR**

### 3️⃣ Brute Force Idea
Sort and check indices. O(N log N).
Cycle Sort can do O(N).

### 4️⃣ Key Insight (AHA 💡)
**Sum Formula:** The sum of `0..N` is `N*(N+1)/2`. The missing number is `ExpectedSum - ActualArraySum`.
**XOR:** `XOR(all 0..N) ^ XOR(arr)` cancels out duplicates, leaving the missing number. Prevents overflow.

### 5️⃣ Algorithm (XOR)
1. `xorAll = 0`, `xorArr = 0`.
2. Loop `0` to `N`: `xorAll ^= i`.
3. Loop `x` in `arr`: `xorArr ^= x`.
4. Return `xorAll ^ xorArr`.

### 6️⃣ Edge Cases & Traps ⚠️
* Integer Overflow with Sum method (if N is large).
* Values not starting from 0 (e.g., 1..N+1), adjust formula.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 6. Find Element That Appears Only Once

```
------------------------------------
| Problem Title -> Single Number   |
```

### 1️⃣ Problem Snapshot
Non-empty array where every element appears **twice** except for one. Find that single one.
`[4, 1, 2, 1, 2]` -> 4.

### 2️⃣ Pattern / Category ⭐
**XOR Properties**

### 3️⃣ Brute Force Idea
HashMap to count frequencies. O(N) space.

### 4️⃣ Key Insight (AHA 💡)
XOR property: `A ^ A = 0` and `A ^ 0 = A`.
XORing all numbers together will "cancel out" the pairs, leaving only the unique number.

### 5️⃣ Algorithm
1. `res = 0`.
2. Foreach `num` in `arr`:
   - `res ^= num`.
3. Return `res`.

### 6️⃣ Edge Cases & Traps ⚠️
* All elements unique (problem constraint usually says ONLY ONE unique).
* Negative numbers (XOR handles them fine).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 7. Find Peak Element

```
------------------------------------
| Problem Title -> Peak Element    |
```

### 1️⃣ Problem Snapshot
A peak element is one that is strictly greater than its neighbors. Find index of any peak.
`[1, 2, 3, 1]` -> 3 is a peak (index 2).

### 2️⃣ Pattern / Category ⭐
**Binary Search on Answer**

### 3️⃣ Brute Force Idea
Linear scan: find element where `arr[i-1] < arr[i] > arr[i+1]`. O(N).

### 4️⃣ Key Insight (AHA 💡)
If `arr[mid] < arr[mid+1]`, it means we are on an "uphill" slope, so a peak MUST exist to the right.
If `arr[mid] > arr[mid+1]`, we are on a "downhill" slope, a peak is to the left (or `mid` itself).

### 5️⃣ Algorithm
1. `low = 0, high = N-1`.
2. While `low < high`:
   - `mid = low + (high - low) / 2`.
   - If `arr[mid] < arr[mid+1]`: `low = mid + 1`.
   - Else: `high = mid`.
3. Return `low`.

### 6️⃣ Edge Cases & Traps ⚠️
* Peak at ends (index 0 or N-1).
* Strictly increasing/decreasing array.

### 7️⃣ Time & Space Complexity
> Time: O(log N)
> Space: O(1)
