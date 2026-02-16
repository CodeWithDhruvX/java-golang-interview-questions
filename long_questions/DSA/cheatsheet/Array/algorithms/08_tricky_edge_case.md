# 8. Tricky Edge-Case Arrays

## 1. Find the Shortest Unsorted Continuous Subarray

### Algorithm
1. Find `min` value from left that is smaller than max seen so far -> `end` index logic?
   - Actually:
   - Find the first element from left which is smaller than `max_seen_so_far`. Wait, no.
   - Find the first element from left that breaks increasing order?
   - Correct Logic:
     - Iterate left to right, keep track of `max`. If `arr[i] < max`, it means `arr[i]` is out of order. Mark `end = i`.
     - Iterate right to left, keep track of `min`. If `arr[i] > min`, it means `arr[i]` is out of order. Mark `start = i`.
2. Return `end - start + 1`.

### Code
```go
import "math"

// FindUnsortedSubarray finds the shortest subarray that needs to be sorted
// Time Complexity: O(N)
// Space Complexity: O(1)
func FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	
	maxVal := math.MinInt
	end := -2
	
	// Find end
	for i := 0; i < n; i++ {
		if nums[i] < maxVal {
			end = i
		} else {
			maxVal = nums[i]
		}
	}
	
	minVal := math.MaxInt
	start := -1
	
	// Find start
	for i := n - 1; i >= 0; i-- {
		if nums[i] > minVal {
			start = i
		} else {
			minVal = nums[i]
		}
	}
	
	if end == -2 {
		return 0 // Already sorted
	}
	
	return end - start + 1
}
```

---

## 2. Check if Array Can Be Sorted by Only One Swap

### Algorithm
1. Create a sorted copy of the array.
2. Compare the original array with the sorted copy.
3. Count mismatches.
4. If mismatches == 0 (already sorted) or mismatches == 2 (one swap fixes two positions), return true.
5. Else return false.

### Code
```go
import "sort"

// CanBeSortedByOneSwap checks if array can be sorted by swapping two elements
// Time Complexity: O(N log N) (due to sorting for comparison)
// Space Complexity: O(N)
func CanBeSortedByOneSwap(nums []int) bool {
	n := len(nums)
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	
	diffCount := 0
	for i := 0; i < n; i++ {
		if nums[i] != sortedNums[i] {
			diffCount++
		}
	}
	
	return diffCount == 0 || diffCount == 2
}
```

---

## 3. Check if Array Can Be Sorted by Reversing One Subarray

### Algorithm
1. Find the first index `start` where order breaks (`arr[i] > arr[i+1]`).
2. If no break, return true.
3. Find the last index `end` where order is broken.
4. Reverse the subarray `arr[start...end]`.
5. Check if the entire array is now sorted.

### Code
```go
// CanSortByReversingSubarray checks if reversing one subarray sorts the array
// Time Complexity: O(N)
// Space Complexity: O(1) (modifies array temporarily or works on indices)
func CanSortByReversingSubarray(arr []int) bool {
	n := len(arr)
	if n < 2 {
		return true
	}
	
	start := -1
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			start = i
			break
		}
	}
	
	if start == -1 {
		return true // Already sorted
	}
	
	end := -1
	for i := n - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			end = i
			break
		}
	}
	
	// Reverse candidate subarray
	// For checking, we can verify if reversing restores order without actual modification
	// 1. arr[start-1] <= arr[end]
	// 2. arr[start] <= arr[end+1]
	// 3. The subarray itself must be decreasing
	
	// Check connections
	if start > 0 && arr[start-1] > arr[end] {
		return false
	}
	if end < n-1 && arr[start] > arr[end+1] {
		return false
	}
	
	// Check if subarray is strictly decreasing
	for i := start; i < end; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	
	return true
}
```

---

## 4. Find the First Missing Positive Integer

### Algorithm
1. Ideal range of positive integers is `1` to `N`.
2. Ignore non-positive numbers and numbers `> N`.
3. Use Cyclic Sort / Index Mapping:
   - Place `nums[i]` at index `nums[i] - 1` if possible.
   - Swap `nums[i]` with `nums[nums[i]-1]`.
4. Iterate again. First index `i` where `nums[i] != i+1` is the answer.
5. If all match, return `N + 1`.

### Code
```go
// FirstMissingPositive finds the smallest missing positive integer
// Time Complexity: O(N)
// Space Complexity: O(1)
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// Swap nums[i] with nums[nums[i]-1]
			targetIdx := nums[i] - 1
			nums[i], nums[targetIdx] = nums[targetIdx], nums[i]
		}
	}
	
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
```

---

## 5. Find the Element Which Appears More Than N/K Times (Majority Element II)

### Algorithm (Boyer-Moore Voting Extended)
1. For `> N/3`, maintain **two** candidates and two counters.
2. Iterate `x`:
   - If `x == cand1`, `count1++`.
   - If `x == cand2`, `count2++`.
   - If `count1 == 0`, `cand1 = x`, `count1 = 1`.
   - If `count2 == 0`, `cand2 = x`, `count2 = 1`.
   - Else, `count1--`, `count2--`.
3. Verify counts of `cand1` and `cand2` in a second pass.

### Code
```go
// MajorityElementII finds elements appearing more than n/3 times
// Time Complexity: O(N)
// Space Complexity: O(1)
func MajorityElementII(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	
	cand1, cand2 := 0, 0
	cnt1, cnt2 := 0, 0
	
	for _, n := range nums {
		if n == cand1 {
			cnt1++
		} else if n == cand2 {
			cnt2++
		} else if cnt1 == 0 {
			cand1 = n
			cnt1 = 1
		} else if cnt2 == 0 {
			cand2 = n
			cnt2 = 1
		} else {
			cnt1--
			cnt2--
		}
	}
	
	count1, count2 := 0, 0
	for _, n := range nums {
		if n == cand1 {
			count1++
		} else if n == cand2 {
			count2++
		}
	}
	
	var res []int
	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}
```

---

## 6. House Robber (Max Sum No Two Adjacent)

### Algorithm
1. Initialize `prev1 = 0` (max sum ending at i-1), `prev2 = 0` (max sum ending at i-2).
2. Iterate `num` in array:
   - `newVal = max(prev1, prev2 + num)`.
   - `prev2 = prev1`.
   - `prev1 = newVal`.
3. Return `prev1`.

### Code
```go
// Rob finds maximum amount of money you can rob avoiding adjacent houses
// Time Complexity: O(N)
// Space Complexity: O(1)
func Rob(nums []int) int {
	prev1 := 0
	prev2 := 0
	
	for _, num := range nums {
		tmp := prev1
		if prev2+num > prev1 {
			prev1 = prev2 + num
		}
		prev2 = tmp
	}
	return prev1
}
```

---

## 7. Find Subarray With Sum = 0 (No Extra Space Attempt)

### Algorithm
1. "No extra space" usually means O(1) auxiliary space beyond input.
2. Standard approach uses HashMap (O(N) space).
3. If we can't use space, we might sort? Sorting takes O(N log N) and allows finding duplicate prefix sums.
   - Calculate Prefix Sums.
   - Sort Prefix Sums.
   - Check if any adjacent elements in sorted prefix sums are equal.
   - Also check if any prefix sum is 0.
4. Total Time: O(N log N). Space: Depends on sorting (O(log N) or O(1)).

### Code
```go
import "sort"

// HasZeroSumSubarrayNoSpace checks for zero sum subarray using sorting (O(N log N))
// Time Complexity: O(N log N)
// Space Complexity: O(1) assuming heapsort or ignoring stack space
func HasZeroSumSubarrayNoSpace(arr []int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	
	// Create prefix sums
	// Note: We might need to handle the case where a prefix sum itself is 0
	// Modify array in-place to store prefix sums?
	// Risk: Overflow. Assuming logic holds.
	
	// If allowed to modify array:
	// 1. Calculate prefix sums in-place
	// 2. Sort the array
	// 3. Check duplicates or 0
	
	currSum := 0
	for i := 0; i < n; i++ {
		currSum += arr[i]
		if currSum == 0 {
			return true
		}
		arr[i] = currSum
	}
	
	sort.Ints(arr)
	
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	
	return false
}
```
