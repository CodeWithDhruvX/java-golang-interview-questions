# 10. Constraint-Driven Problems

## 1. Find Duplicate in Array of Size N+1 with Elements 1..N

### Algorithm (Floyd's Cycle Detection)
1. Initialize `slow = arr[0]`, `fast = arr[0]`.
2. Move `slow` by 1 step (`arr[slow]`) and `fast` by 2 steps (`arr[arr[fast]]`).
3. Repeat until they meet.
4. Reset `slow` to `arr[0]`.
5. Move both `slow` and `fast` by 1 step until they meet.
6. The meeting point is the duplicate.

### Code
```go
// FindDuplicateNumber finds the duplicate number without modifying array (O(N) time, O(1) space)
func FindDuplicateNumber(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	
	return slow
}
```

---

## 2. Find Missing and Repeating Number

### Algorithm (Math Equation)
1. Let missing be `X`, repeating be `Y`.
2. `Sum_N` = sum of 1 to N. `Sum_Arr` = sum of array elements.
3. `Sum_N - Sum_Arr = X - Y`.
4. `SumSq_N` = sum of squares 1 to N. `SumSq_Arr` = sum of squares of array.
5. `SumSq_N - SumSq_Arr = X^2 - Y^2 = (X - Y)(X + Y)`.
6. Solve two equations to get `X` and `Y`.

### Code
```go
// FindMissingAndRepeating finds missing and repeating numbers using math
// Time Complexity: O(N)
// Space Complexity: O(1)
func FindMissingAndRepeating(arr []int) (int, int) {
	n := len(arr)
	s := int64(n * (n + 1) / 2)
	sSq := int64(n * (n + 1) * (2*n + 1) / 6)
	
	actualSum := int64(0)
	actualSumSq := int64(0)
	
	for _, val := range arr {
		actualSum += int64(val)
		actualSumSq += int64(val) * int64(val)
	}
	
	diff := s - actualSum       // X - Y
	diffSq := sSq - actualSumSq // X^2 - Y^2
	
	sum := diffSq / diff // X + Y
	
	missing := (diff + sum) / 2
	repeating := sum - missing
	
	return int(missing), int(repeating)
}
```

---

## 3. Find Maximum Sum Subarray With At Least K Elements

### Algorithm
1. Compute `prefixSum`.
2. Use a sliding window of size `K`.
3. Track minimum prefix sum seen so far `minPrefix` for index `i - K`.
4. `maxSum = max(maxSum, prefixSum[i] - minPrefix)`.

### Code
```go
import "math"

// MaxSumSubarrayAtLeastK finds max sum subarray with length >= k
// Time Complexity: O(N)
// Space Complexity: O(N)
func MaxSumSubarrayAtLeastK(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return 0
	}
	
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	
	maxSum := prefixSum[k-1]
	minPrefix := 0 
	
	// Consider windows ending at i (where i >= k)
	// Window sum = prefixSum[i] - prefixSum[i-len]
	// We want to maximize this.
	// This is equivalent to maximizing prefixSum[i] - MIN(prefixSum[j]) where j <= i-k
	// Base case handling needs care for index -1 (which effectively has prefix sum 0)
	
	// Re-evaluating loop range
	// We need to track min prefix sum in range [-1, i-k]
	
	minPrefixArr := make([]int, n)
	// Precompute min prefix sum? Not strictly needed if we iterate.
	
	// Sliding window logic fix:
	// Iterate through array. At index i, we can form a valid subarray ending at i starting at j <= i-k+1.
	// Sum = prefixSum[i] - prefixSum[j-1].
	// We want to minimize prefixSum[j-1] where j-1 <= i-k.
	
	currentMin := 0 // Represents min prefix sum before the window starts
	
	for i := k; i < n; i++ {
		// Update min prefix sum considering the element that just became available to be "left out"
		// The element at `i-k` is now a valid end-point for the "excluded prefix"
		if prefixSum[i-k] < currentMin {
			currentMin = prefixSum[i-k]
		}
		
		val := prefixSum[i] - currentMin
		if val > maxSum {
			maxSum = val
		}
	}
	
	// Also check initial K window against potentially initialized maxSum
	// (handled by initialization)
	
	// One edge case: if K=N.
	// Handled.
	
	return maxSum
}
```

---

## 4. Subarray With Equal Number of 0s, 1s, and 2s

### Algorithm
1. Iterate through array, counting `c0`, `c1`, `c2`.
2. Compute differences: `d1 = c0 - c1`, `d2 = c1 - c2`.
3. Store pair `(d1, d2)` in a map `key -> first_index`.
4. If pair repeats, it means subarray between occurrences has equal counts.
5. `maxLen = max(maxLen, i - map[key])`.

### Code
```go
import "fmt"

// LongestEqual012 finds longest substring with equal 0s, 1s, and 2s
// Time Complexity: O(N)
// Space Complexity: O(N)
func LongestEqual012(arr []int) int {
	type Key struct {
		d1, d2 int
	}
	m := make(map[Key]int)
	m[Key{0, 0}] = -1 // Base case
	
	c0, c1, c2 := 0, 0, 0
	maxLen := 0
	
	for i, val := range arr {
		if val == 0 {
			c0++
		} else if val == 1 {
			c1++
		} else if val == 2 {
			c2++
		}
		
		key := Key{c0 - c1, c1 - c2}
		
		if idx, ok := m[key]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		} else {
			m[key] = i
		}
	}
	return maxLen
}
```

---

## 5. Find Count of Subarrays Whose Product < K

### Algorithm (Sliding Window)
1. Initialize `left = 0`, `prod = 1`, `count = 0`.
2. Iterate `right` from `0` to `n-1`.
3. `prod *= arr[right]`.
4. While `prod >= K` and `left <= right`:
   - `prod /= arr[left]`.
   - `left++`.
5. `count += right - left + 1`.

### Code
```go
// NumSubarrayProductLessThanK counts subarrays with product strictly less than k
// Time Complexity: O(N)
// Space Complexity: O(1)
func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	
	prod := 1
	ans := 0
	left := 0
	
	for right, val := range nums {
		prod *= val
		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
```
