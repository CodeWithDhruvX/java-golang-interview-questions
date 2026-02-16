# 3. Mathematical & Logical Array Problems

## 1. Find All Pairs With a Given Sum (Two Sum)

### Algorithm
1. Create a `map` to store elements and their indices `val -> index`.
2. Iterate through the array with index `i`.
3. Calculate `complement = target - arr[i]`.
4. Check if `complement` exists in the map.
5. If yes, return `[]int{map[complement], i}`.
6. If no, store `arr[i]` in the map.

### Code
```go
// TwoSum finds indices of two numbers that add up to target
// Time Complexity: O(N)
// Space Complexity: O(N)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
```

---

## 2. Find All Triplets With a Given Sum (Three Sum)

### Algorithm
1. Sort the array.
2. Iterate `i` from `0` to `n-2`.
3. Skip duplicate `arr[i]`.
4. Initialize `left = i + 1`, `right = n - 1`.
5. While `left < right`:
   - Calculate `sum = arr[i] + arr[left] + arr[right]`.
   - If `sum == target` (usually 0), add triplet to result.
     - Skip duplicate `arr[left]` and `arr[right]`.
     - Move `left++`, `right--`.
   - If `sum < target`, `left++`.
   - If `sum > target`, `right--`.

### Code
```go
import (
	"sort"
)

// ThreeSum finds all unique triplets that sum to zero
// Time Complexity: O(N^2)
// Space Complexity: O(1) (excluding output)
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				
				// Skip duplicates for second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
```

---

## 3. Find the Majority Element (> n/2)

### Algorithm (Boyer-Moore Voting Algorithm)
1. Initialize `candidate = 0`, `count = 0`.
2. Iterate through the array:
   - If `count == 0`, set `candidate = x`.
   - If `x == candidate`, increment `count`.
   - Else, decrement `count`.
3. Return `candidate`.

### Code
```go
// MajorityElement finds the element appearing more than n/2 times
// Time Complexity: O(N)
// Space Complexity: O(1)
func MajorityElement(nums []int) int {
	candidate := 0
	count := 0
	
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
```

---

## 4. Find Leaders in an Array

### Algorithm
1. Initialize `maxFromRight` with the last element.
2. Add the last element to the result list.
3. Iterate backwards from `n-2` to `0`.
4. If `arr[i] > maxFromRight`:
   - Add `arr[i]` to result.
   - Update `maxFromRight = arr[i]`.
5. Reverse the result list to maintain original order (optional).

### Code
```go
// Leaders finds elements which are greater than all elements to their right
// Time Complexity: O(N)
// Space Complexity: O(1) (excluding output)
func Leaders(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	
	var leaders []int
	maxRight := arr[n-1]
	leaders = append(leaders, maxRight)
	
	for i := n - 2; i >= 0; i-- {
		if arr[i] > maxRight {
			maxRight = arr[i]
			leaders = append(leaders, maxRight)
		}
	}
	
	// Reverse to get original order
	for i, j := 0, len(leaders)-1; i < j; i, j = i+1, j-1 {
		leaders[i], leaders[j] = leaders[j], leaders[i]
	}
	
	return leaders
}
```

---

## 5. Find the Equilibrium Index

### Algorithm
1. Calculate the total sum of the array.
2. Initialize `leftSum = 0`.
3. Iterate through index `i`:
   - Calculate `rightSum = totalSum - leftSum - arr[i]`.
   - If `leftSum == rightSum`, return `i`.
   - Update `leftSum += arr[i]`.
4. Return -1 if no equilibrium found.

### Code
```go
// EquilibriumIndex finds index where sum of left elements equals sum of right elements
// Time Complexity: O(N)
// Space Complexity: O(1)
func EquilibriumIndex(arr []int) int {
	totalSum := 0
	for _, val := range arr {
		totalSum += val
	}
	
	leftSum := 0
	for i, val := range arr {
		rightSum := totalSum - leftSum - val
		if leftSum == rightSum {
			return i
		}
		leftSum += val
	}
	return -1
}
```

---

## 6. Find the Maximum Difference such that j > i and arr[j] > arr[i]

### Algorithm
1. Initialize `minVal = arr[0]` and `maxDiff = -1` (or 0).
2. Iterate `i` from `1` to `n-1`.
3. Calculate `diff = arr[i] - minVal`.
4. If `diff > maxDiff`, update `maxDiff = diff`.
5. If `arr[i] < minVal`, update `minVal = arr[i]`.
6. Return `maxDiff`.

### Code
```go
// MaximumDifference finds max(arr[j] - arr[i]) such that j > i
// Time Complexity: O(N)
// Space Complexity: O(1)
func MaximumDifference(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	
	minVal := arr[0]
	maxDiff := -1 // Assuming return -1 if no such pair exists (e.g. descending sorted)
	
	for i := 1; i < len(arr); i++ {
		if arr[i] > minVal {
			diff := arr[i] - minVal
			if diff > maxDiff {
				maxDiff = diff
			}
		} else {
			minVal = arr[i]
		}
	}
	return maxDiff
}
```

---

## 7. Check if Array Elements are Consecutive

### Algorithm
1. Find min and max of array.
2. If `max - min + 1 != n` (where n is length), return false.
3. To handle duplicates (if checked):
   - Use a `Set` or `Boolean Array` to mark visited elements.
   - If an element is visited twice, return false.

### Code
```go
import "math"

// AreConsecutive checks if array elements are consecutive integers
// Time Complexity: O(N)
// Space Complexity: O(N) (using map for duplicates check)
func AreConsecutive(arr []int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	
	minVal := math.MaxInt
	maxVal := math.MinInt
	seen := make(map[int]bool)
	
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
		if seen[val] {
			return false // Duplicate found
		}
		seen[val] = true
	}
	
	return maxVal-minVal+1 == n
}
```

---

## 8. Find the Duplicate Number (Floyd's Cycle Detection)

### Algorithm
1. **Phase 1 (Cycle Detection)**:
   - Initialize `slow = arr[0]`, `fast = arr[0]`.
   - `slow = arr[slow]`.
   - `fast = arr[arr[fast]]`.
   - Repeat until `slow == fast`.
2. **Phase 2 (Find Entrance)**:
   - Reset `slow = arr[0]`.
   - Move both `slow` and `fast` one step at a time.
   - When they meet, that value is the duplicate.

### Code
```go
// FindDuplicate finds the duplicate number in array of n+1 integers in range [1, n]
// Time Complexity: O(N)
// Space Complexity: O(1)
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	
	// Phase 1
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	
	// Phase 2
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	
	return slow
}
```
