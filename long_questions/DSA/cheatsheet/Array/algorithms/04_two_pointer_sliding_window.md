# 4. Two Pointer & Sliding Window (High Frequency)

## 1. Move All Zeros to the End

### Algorithm
1. Initialize `insertPos = 0`.
2. Iterate through the array.
3. If current element != 0:
   - Place it at `insertPos`.
   - Increment `insertPos`.
4. After the loop, fill remaining positions from `insertPos` to end with `0`.

### Code
```go
// MoveZeroes moves all 0s to the end while maintaining order of non-zero elements
// Time Complexity: O(N)
// Space Complexity: O(1)
func MoveZeroes(nums []int) {
	insertPos := 0
	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}
	
	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
```

---

## 2. Sort Array of 0s, 1s, and 2s (Dutch National Flag)

### Algorithm
1. Initialize `low = 0`, `mid = 0`, `high = n - 1`.
2. Loop while `mid <= high`:
   - If `arr[mid] == 0`:
     - Swap `arr[low]` and `arr[mid]`.
     - `low++`, `mid++`.
   - If `arr[mid] == 1`:
     - `mid++`.
   - If `arr[mid] == 2`:
     - Swap `arr[mid]` and `arr[high]`.
     - `high--`.

### Code
```go
// SortColors sorts an array of 0s, 1s, and 2s in-place
// Time Complexity: O(N)
// Space Complexity: O(1)
func SortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
```

---

## 3. Find Subarray With Given Sum (Non-negative Integers)

### Algorithm
1. Initialize `start = 0`, `currentSum = 0`.
2. Iterate `end` from `0` to `n-1`.
3. Add `arr[end]` to `currentSum`.
4. While `currentSum > target` and `start <= end`:
   - Subtract `arr[start]` from `currentSum`.
   - `start++`.
5. If `currentSum == target`, return `[start, end]`.
6. Return empty slice if not found.

### Code
```go
// SubarraySum finds a continuous subarray which adds up to a given number
// Time Complexity: O(N)
// Space Complexity: O(1)
func SubarraySum(arr []int, target int) []int {
	currentSum := 0
	start := 0
	
	for end, val := range arr {
		currentSum += val
		
		for currentSum > target && start <= end {
			currentSum -= arr[start]
			start++
		}
		
		if currentSum == target {
			return []int{start, end}
		}
	}
	return nil
}
```

---

## 4. Find Maximum Sum Subarray (Kadane’s Algorithm)

### Algorithm
1. Initialize `maxSoFar = arr[0]`, `currentMax = arr[0]`.
2. Iterate from `1` to `n-1`.
3. `currentMax = max(arr[i], currentMax + arr[i])`.
4. `maxSoFar = max(maxSoFar, currentMax)`.
5. Return `maxSoFar`.

### Code
```go
import "math"

// MaxSubArrays returns the contiguous subarray with the largest sum
// Time Complexity: O(N)
// Space Complexity: O(1)
func MaxSubArray(nums []int) int {
	maxSoFar := nums[0]
	currentMax := nums[0]
	
	for i := 1; i < len(nums); i++ {
		if currentMax < 0 {
			currentMax = nums[i]
		} else {
			currentMax += nums[i]
		}
		
		if currentMax > maxSoFar {
			maxSoFar = currentMax
		}
	}
	return maxSoFar
}
```

---

## 5. Find Longest Subarray With Sum = K (Handles Negatives)

### Algorithm
1. Initialize `map` (prefixSum -> index), `sum = 0`, `maxLen = 0`.
2. Iterate through array.
3. `sum += arr[i]`.
4. If `sum == K`, update `maxLen = i + 1`.
5. If `sum - K` exists in map, update `maxLen = max(maxLen, i - map[sum-K])`.
6. If `sum` not in map, store `map[sum] = i`.

### Code
```go
// LongestSubarraySumK finds the length of the longest subarray with sum equal to k
// Time Complexity: O(N)
// Space Complexity: O(N)
func LongestSubarraySumK(arr []int, k int) int {
	sumMap := make(map[int]int)
	sum := 0
	maxLen := 0
	
	for i, val := range arr {
		sum += val
		
		if sum == k {
			maxLen = i + 1
		}
		
		if idx, found := sumMap[sum-k]; found {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		}
		
		if _, found := sumMap[sum]; !found {
			sumMap[sum] = i
		}
	}
	return maxLen
}
```

---

## 6. Find Smallest Subarray With Sum > X

### Algorithm
1. Initialize `start = 0`, `currentSum = 0`, `minLen = n + 1`.
2. Iterate `end` from `0` to `n-1`.
3. `currentSum += arr[end]`.
4. While `currentSum > x`:
   - Update `minLen = min(minLen, end - start + 1)`.
   - `currentSum -= arr[start]`.
   - `start++`.
5. If `minLen > n`, return 0.

### Code
```go
import "math"

// MinSubArrayLen finds the minimal length of a contiguous subarray of which the sum is greater than or equal to target
// Time Complexity: O(N)
// Space Complexity: O(1)
func MinSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	sum := 0
	left := 0
	
	for right, val := range nums {
		sum += val
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
```

---

## 7. Find Longest Subarray With Equal 0s and 1s

### Algorithm
1. Treat 0s as -1.
2. The problem becomes "Longest Subarray with Sum = 0".
3. Use Map (prefixSum -> index).
4. `sum = 0`, `maxLen = 0`.
5. Iterate through array:
   - If `arr[i] == 0`, sum += -1. Else sum += 1.
   - If `sum == 0`, `maxLen = i + 1`.
   - If `sum` seen in map, `maxLen = max(maxLen, i - map[sum])`.
   - If `sum` not seen, `map[sum] = i`.

### Code
```go
// FindMaxLength finds the maximum length of a contiguous subarray with equal number of 0 and 1
// Time Complexity: O(N)
// Space Complexity: O(N)
func FindMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1 // Base case for sum 0 at start
	sum := 0
	maxLen := 0
	
	for i, val := range nums {
		if val == 0 {
			sum--
		} else {
			sum++
		}
		
		if idx, ok := m[sum]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		} else {
			m[sum] = i
		}
	}
	return maxLen
}
```
