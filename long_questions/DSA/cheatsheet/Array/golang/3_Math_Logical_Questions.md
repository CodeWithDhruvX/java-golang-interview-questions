```go
package main

import (
	"math"
	"sort"
)

// 1. Find All Pairs with a Given Sum (Two Sum)
// Time: O(N), Space: O(N)
func TwoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i, val := range arr {
		complement := target - val
		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[val] = i
	}
	return nil
}

// 2. Find All Triplets with a Given Sum (3Sum)
// Time: O(N^2), Space: O(1) (ignoring output)
func ThreeSum(arr []int) [][]int {
	sort.Ints(arr)
	var res [][]int
	n := len(arr)

	for i := 0; i < n-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		target := -arr[i]
		left, right := i+1, n-1
		for left < right {
			sum := arr[left] + arr[right]
			if sum == target {
				res = append(res, []int{arr[i], arr[left], arr[right]})
				// Skip duplicates for second and third elements
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// 3. Find Majority Element (> N/2 times)
// Time: O(N), Space: O(1)
func MajorityElement(arr []int) int {
	candidate := 0
	count := 0
	for _, num := range arr {
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

// 4. Find Leaders in an Array
// Time: O(N), Space: O(1)
func FindLeaders(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	var leaders []int
	maxRight := math.MinInt
	for i := n - 1; i >= 0; i-- {
		if arr[i] > maxRight {
			leaders = append(leaders, arr[i])
			maxRight = arr[i]
		}
	}
	// Reverse to maintain left-to-right order finding (optional but common)
	for i, j := 0, len(leaders)-1; i < j; i, j = i+1, j-1 {
		leaders[i], leaders[j] = leaders[j], leaders[i]
	}
	return leaders
}

// 5. Find Equilibrium Index
// Time: O(N), Space: O(1)
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

// 6. Max Difference such that j > i
// Time: O(N), Space: O(1)
func MaxDifference(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	minSoFar := arr[0]
	maxDiff := 0 // Assuming non-negative difference required, else init to -1 or MinInt
	for i := 1; i < len(arr); i++ {
		if arr[i]-minSoFar > maxDiff {
			maxDiff = arr[i] - minSoFar
		}
		if arr[i] < minSoFar {
			minSoFar = arr[i]
		}
	}
	return maxDiff
}

// 7. Check if Array Elements are Consecutive
// Time: O(N), Space: O(N) using Set
func CheckConsecutive(arr []int) bool {
	n := len(arr)
	if n == 0 {
		return false // debatable
	}
	minVal, maxVal := arr[0], arr[0]
	seen := make(map[int]bool)

	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
		if seen[val] {
			return false
		}
		seen[val] = true
	}
	return maxVal-minVal+1 == n
}

// 8. Find the Duplicate Number (Floyd's Cycle)
// Time: O(N), Space: O(1)
func FindDuplicateNumber(arr []int) int {
	slow := arr[0]
	fast := arr[0]

	// Phase 1: Meet at intersection
	for {
		slow = arr[slow]
		fast = arr[arr[fast]]
		if slow == fast {
			break
		}
	}

	// Phase 2: Find entrance to cycle
	slow = arr[0]
	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
	}
	return slow
}
```
