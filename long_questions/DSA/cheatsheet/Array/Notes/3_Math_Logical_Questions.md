# 3. Mathematical & Logical Array Problems

---

## 1. Find All Pairs with a Given Sum (Two Sum)

```
------------------------------------
| Problem Title -> Two Sum         |
```

### 1️⃣ Problem Snapshot
Find two numbers in the array that add up to a specific `target`. Return their indices.

### 2️⃣ Pattern / Category ⭐
**Hashing (Unsorted) OR Two Pointers (Sorted)**

### 3️⃣ Brute Force Idea
Nested loops: check `arr[i] + arr[j] == target`. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
If we embrace Space, we save Time.
`y = target - x`. As we iterate `x`, check if `y` exists in the hash map.

### 5️⃣ Algorithm (Hash Map)
1. Create `map` (value -> index).
2. Loop `i` from 0 to `N-1`:
   - `complement = target - arr[i]`.
   - If `map` has `complement`, return `[map[complement], i]`.
   - `map[arr[i]] = i`.

### 6️⃣ Edge Cases & Traps ⚠️
* Same element used twice (e.g., target 6, array `[3]`, returning `[0, 0]` is wrong).
* No solution.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N)

---

## 2. Find All Triplets with a Given Sum (3Sum)

```
------------------------------------
| Problem Title -> 3Sum            |
```

### 1️⃣ Problem Snapshot
Find all unique triplets `[a, b, c]` such that `a + b + c = 0` (or target).

### 2️⃣ Pattern / Category ⭐
**Sort + Two Pointers**

### 3️⃣ Brute Force Idea
Three nested loops. O(N^3).

### 4️⃣ Key Insight (AHA 💡)
Sort the array first. Fix one element `arr[i]`, and the problem reduces to "Two Sum" on the remaining array (`target = -arr[i]`).

### 5️⃣ Algorithm
1. Sort `arr`.
2. Loop `i` from 0 to `N-3`:
   - If `i > 0` and `arr[i] == arr[i-1]`, skip (avoid duplicates).
   - Run Two Pointers on `[i+1...N-1]` for target `-arr[i]`.
   - If found, add to result, and skip duplicates for `left` and `right`.

### 6️⃣ Edge Cases & Traps ⚠️
* Duplicate triplets (must handle skipping logic carefully).
* Array size < 3.

### 7️⃣ Time & Space Complexity
> Time: O(N^2)
> Space: O(1) (ignoring output)

---

## 3. Find Majority Element (> N/2 times)

```
------------------------------------
| Problem Title -> Majority Elem   |
```

### 1️⃣ Problem Snapshot
Find the element that appears more than `N/2` times. Assumed to always exist.

### 2️⃣ Pattern / Category ⭐
**Moore’s Voting Algorithm**

### 3️⃣ Brute Force Idea
Count frequency of each element. O(N*N) or O(N) space.

### 4️⃣ Key Insight (AHA 💡)
If we cancel out every pair of distinct elements, the majority element will remain. It dominates the count.

### 5️⃣ Algorithm
1. `candidate = None`, `count = 0`.
2. Loop through `arr`:
   - If `count == 0`: `candidate = x`.
   - If `x == candidate`: `count++`.
   - Else: `count--`.
3. Return `candidate`.

### 6️⃣ Edge Cases & Traps ⚠️
* No majority element (count check at end required if not guaranteed).
* N=1.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 4. Find Leaders in an Array

```
------------------------------------
| Problem Title -> Array Leaders   |
```

### 1️⃣ Problem Snapshot
An element is a leader if it is strictly greater than all elements to its right. The rightmost element is always a leader.

### 2️⃣ Pattern / Category ⭐
**Scan from Right (Suffix Max)**

### 3️⃣ Brute Force Idea
For each element, scan all elements to its right. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
If we process from Right to Left, we only need to track the "Max so far". If current > Max, it's a leader.

### 5️⃣ Algorithm
1. `maxRight = -INF`.
2. Loop `i` from `N-1` down to 0:
   - If `arr[i] > maxRight`:
     - Add `arr[i]` to leaders.
     - `maxRight = arr[i]`.
3. Reverse leaders list (if left-to-right order needed).

### 6️⃣ Edge Cases & Traps ⚠️
* Ascending sorted array (only last element is leader).
* Duplicates (strictly greater means equal is not leader).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 5. Find Equilibrium Index

```
------------------------------------
| Problem Title -> Equilibrium Idx |
```

### 1️⃣ Problem Snapshot
Find an index such that `Sum(Left) == Sum(Right)`.

### 2️⃣ Pattern / Category ⭐
**Prefix Sum vs Total Sum**

### 3️⃣ Brute Force Idea
Calculate sum left and right for every index. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
`RightSum` depends on `TotalSum` and `LeftSum`.
`RightSum = TotalSum - LeftSum - arr[i]`.
We need `LeftSum == RightSum`.

### 5️⃣ Algorithm
1. Calculate `totalSum`.
2. `leftSum = 0`.
3. Loop `i` from 0 to `N-1`:
   - `rightSum = totalSum - leftSum - arr[i]`.
   - If `leftSum == rightSum`, return `i`.
   - `leftSum += arr[i]`.
4. Return -1.

### 6️⃣ Edge Cases & Traps ⚠️
* Negative numbers.
* No equilibrium index. (e.g. `[1, 2, 3]`).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 6. Max Difference such that j > i

```
------------------------------------
| Problem Title -> Buy/Sell Stock  |
```

### 1️⃣ Problem Snapshot
Find max value of `arr[j] - arr[i]` where `j > i`. (Equivalent to Best Time to Buy/Sell Stock).

### 2️⃣ Pattern / Category ⭐
**Track Minimum So Far**

### 3️⃣ Brute Force Idea
Nested Loop `j > i`. O(N^2).

### 4️⃣ Key Insight (AHA 💡)
To maximize difference at `j`, we want to subtract the *smallest* value seen *before* `j`.

### 5️⃣ Algorithm
1. `minSoFar = arr[0]`.
2. `maxDiff = 0` (or -1 if negative diffs allowed).
3. Loop `i` from 1 to `N-1`:
   - `diff = arr[i] - minSoFar`.
   - `maxDiff = max(maxDiff, diff)`.
   - `minSoFar = min(minSoFar, arr[i])`.

### 6️⃣ Edge Cases & Traps ⚠️
* Decreasing array (MaxDiff usually 0 or negative).
* Single element.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)

---

## 7. Check if Array Elements are Consecutive

```
------------------------------------
| Problem Title -> Consecutivity   |
```

### 1️⃣ Problem Snapshot
Given an array, check if it contains consecutive integers (unsorted allowed).
`[5, 4, 2, 3]` -> Yes (`2, 3, 4, 5`).

### 2️⃣ Pattern / Category ⭐
**Min/Max Range + Formula**

### 3️⃣ Brute Force Idea
Sort and check `arr[i] == arr[i-1] + 1`: O(N log N).

### 4️⃣ Key Insight (AHA 💡)
If elements are consecutive:
1. `Max - Min + 1 == N`
2. All elements must be distinct.

### 5️⃣ Algorithm
1. Find `max` and `min`.
2. If `max - min + 1 != N`, return `false`.
3. Check for duplicates (use Set or modify array if allowed).
   - E.g. Visited array logic using `index = arr[i] - min`.

### 6️⃣ Edge Cases & Traps ⚠️
* Duplicates with correct range (e.g., `[1, 2, 2]` -> range is 2, N=3 -> fails range check). But `[1, 2, 2]` vs `[1, 2, 3]` logic needs distinct check.

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(N) (Set) or O(1) (In-place sort/negation)

---

## 8. Find the Duplicate Number (Floyd's Cycle)

```
------------------------------------
| Problem Title -> Duplicate Num   |
```

### 1️⃣ Problem Snapshot
Given array of `N+1` integers in range `[1, N]`, find the one repeated number.
**Constraint:** Must not modify array, O(1) space.

### 2️⃣ Pattern / Category ⭐
**Linked List Cycle Method (Tortoise & Hare)**

### 3️⃣ Brute Force Idea
HashMap O(N) space. Sorting O(N log N).

### 4️⃣ Key Insight (AHA 💡)
Since values are indices (`1..N`), we can treat `arr[i]` as a `next` pointer. Duplicate means two indices point to same value -> Cycle!

### 5️⃣ Algorithm
1. `slow = arr[0]`, `fast = arr[0]`.
2. Phase 1: `slow = arr[slow]`, `fast = arr[arr[fast]]`. Stop when `slow == fast`.
3. Phase 2: `slow = arr[0]`. Move both one step at a time.
4. Stop when `slow == fast`. That is the start of structure (duplicate).

### 6️⃣ Edge Cases & Traps ⚠️
* Values must be within bounds (1 to N).
* Exactly one duplicate exists (pigeons).

### 7️⃣ Time & Space Complexity
> Time: O(N)
> Space: O(1)
