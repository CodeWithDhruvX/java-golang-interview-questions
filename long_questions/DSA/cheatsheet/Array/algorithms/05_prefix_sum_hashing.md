# 5. Prefix Sum & Hashing Based Problems

## 1. Range Sum Queries Using Prefix Sum

### Algorithm
1. Create a `prefixSum` array where `prefixSum[i] = sum(arr[0]...arr[i])`.
2. For a query `(L, R)`, the sum is `prefixSum[R] - prefixSum[L-1]`.
3. Handle base case where `L=0` separately (`prefixSum[R]`).

### Code
```go
// NumArray handles range sum queries
// Time Complexity: Constructor O(N), SumRange O(1)
// Space Complexity: O(N)
type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefixSum: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}
```

---

## 2. Find Subarray With Sum = 0

### Algorithm
1. Initialize `map` (prefixSum -> index) and `sum = 0`.
2. Iterate through the array.
3. Add `arr[i]` to `sum`.
4. If `sum == 0` or `map` contains `sum`:
   - Found subarray with sum 0. Return true.
5. Add `sum` to map.
6. Return `false` if loop finishes.

### Code
```go
// HasZeroSumSubarray checks if there exists a subarray with 0 sum
// Time Complexity: O(N)
// Space Complexity: O(N)
func HasZeroSumSubarray(arr []int) bool {
	prefixSumSet := make(map[int]bool)
	sum := 0
	
	for _, val := range arr {
		sum += val
		
		if sum == 0 || prefixSumSet[sum] {
			return true
		}
		prefixSumSet[sum] = true
	}
	return false
}
```

---

## 3. Count Subarrays With Sum = K

### Algorithm
1. Initialize `map` (prefixSum -> count) with `{0: 1}`.
2. Initialize `sum = 0`, `count = 0`.
3. Iterate through array:
   - `sum += arr[i]`.
   - If `map` contains `sum - K`, add `map[sum - K]` to `count`.
   - Increment count of `sum` in `map`.
4. Return `count`.

### Code
```go
// SubarraySumEqualsK counts subarrays with sum equals to k
// Time Complexity: O(N)
// Space Complexity: O(N)
func SubarraySumEqualsK(nums []int, k int) int {
	count := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1 // Base case
	
	for _, num := range nums {
		sum += num
		if c, ok := m[sum-k]; ok {
			count += c
		}
		m[sum]++
	}
	return count
}
```

---

## 4. Find Longest Subarray With Distinct Elements

### Algorithm
1. Initialize `left = 0`, `maxLen = 0`.
2. Use a `map` (element -> index) to store last seen position.
3. Iterate with `right` from `0` to `n-1`.
4. If `arr[right]` is in `map` and `map[arr[right]] >= left`:
   - Move `left` to `map[arr[right]] + 1`.
5. Update `map[arr[right]] = right`.
6. Update `maxLen = max(maxLen, right - left + 1)`.

### Code
```go
// LengthOfLongestSubstring finds longest subarray with distinct characters/elements
// Time Complexity: O(N)
// Space Complexity: O(N)
func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left := 0
	maxLen := 0
	
	for right := 0; right < len(s); right++ {
		char := s[right]
		if idx, ok := m[char]; ok && idx >= left {
			left = idx + 1
		}
		m[char] = right
		
		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```

---

## 5. Check If There Exists a Subarray With Given XOR

### Algorithm
1. Similar to "Subarray Sum = K", but use XOR properties.
2. Initialize `map` (prefixXOR -> index/bool), `xorSum = 0`.
3. Iterate through array:
   - `xorSum ^= arr[i]`.
   - If `xorSum == target` or `map` contains `xorSum ^ target`:
     - Return true (found).
   - Add `xorSum` to map.

### Code
```go
// HasSubarrayWithXOR checks if a subarray with given XOR exists
// Time Complexity: O(N)
// Space Complexity: O(N)
func HasSubarrayWithXOR(arr []int, target int) bool {
	prefixXor := make(map[int]bool)
	prefixXor[0] = true
	xorSum := 0
	
	for _, val := range arr {
		xorSum ^= val
		if prefixXor[xorSum^target] {
			return true
		}
		prefixXor[xorSum] = true
	}
	return false
}
```

---

## 6. Find Number of Subarrays With Equal Odd and Even Elements

### Algorithm
1. Transform array: Odd -> 1, Even -> -1.
2. Find number of subarrays with sum = 0.
3. Use the logic from Problem 3 ("Count Subarrays with Sum K") with `K = 0`.

### Code
```go
// CountSubarraysEqualOddEven counts subarrays with equal odd and even numbers
// Time Complexity: O(N)
// Space Complexity: O(N)
func CountSubarraysEqualOddEven(arr []int) int {
	count := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1
	
	for _, val := range arr {
		if val%2 == 0 {
			sum -= 1
		} else {
			sum += 1
		}
		
		if c, ok := m[sum]; ok {
			count += c
		}
		m[sum]++
	}
	return count
}
```
