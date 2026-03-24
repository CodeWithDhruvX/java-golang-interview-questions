package main

import (
	"math"
	"sort"
)

// 1. Find All Pairs with a Given Sum (Two Sum)

// Brute Force Approach: Check all possible pairs
// Time: O(N^2), Space: O(1)
func TwoSumBruteForce(arr []int, target int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] + arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Optimized Approach: Hash map for complement lookup
// Time: O(N), Space: O(N)
func TwoSumOptimized(arr []int, target int) []int {
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

// Legacy function for backward compatibility
func TwoSum(arr []int, target int) []int {
	return TwoSumOptimized(arr, target)
}

// 2. Find All Triplets with a Given Sum (3Sum)

// Brute Force Approach: Check all possible triplets
// Time: O(N^3), Space: O(1)
func ThreeSumBruteForce(arr []int) [][]int {
	n := len(arr)
	var res [][]int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i] + arr[j] + arr[k] == 0 {
					triplet := []int{arr[i], arr[j], arr[k]}
					// Sort triplet to avoid duplicates
					sort.Ints(triplet)
					res = append(res, triplet)
				}
			}
		}
	}
	return res
}

// Optimized Approach: Sort + two pointers
// Time: O(N^2), Space: O(1) (ignoring output)
func ThreeSumOptimized(arr []int) [][]int {
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

// Legacy function for backward compatibility
func ThreeSum(arr []int) [][]int {
	return ThreeSumOptimized(arr)
}

// 3. Find Majority Element (> N/2 times)

// Brute Force Approach: Count frequency of each element
// Time: O(N^2), Space: O(1)
func MajorityElementBruteForce(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count > n/2 {
			return arr[i]
		}
	}
	return -1 // No majority element
}

// Optimized Approach: Boyer-Moore Voting Algorithm
// Time: O(N), Space: O(1)
func MajorityElementOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func MajorityElement(arr []int) int {
	return MajorityElementOptimized(arr)
}

// 4. Find Leaders in an Array

// Brute Force Approach: For each element, check if it's greater than all elements to its right
// Time: O(N^2), Space: O(1)
func FindLeadersBruteForce(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	var leaders []int
	for i := 0; i < n; i++ {
		isLeader := true
		for j := i + 1; j < n; j++ {
			if arr[i] < arr[j] {
				isLeader = false
				break
			}
		}
		if isLeader {
			leaders = append(leaders, arr[i])
		}
	}
	return leaders
}

// Optimized Approach: Scan from right to left
// Time: O(N), Space: O(1)
func FindLeadersOptimized(arr []int) []int {
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

// Legacy function for backward compatibility
func FindLeaders(arr []int) []int {
	return FindLeadersOptimized(arr)
}

// 5. Find Equilibrium Index

// Brute Force Approach: For each index, calculate left and right sum
// Time: O(N^2), Space: O(1)
func EquilibriumIndexBruteForce(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		leftSum := 0
		rightSum := 0
		for j := 0; j < i; j++ {
			leftSum += arr[j]
		}
		for j := i + 1; j < n; j++ {
			rightSum += arr[j]
		}
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

// Optimized Approach: Use total sum and running left sum
// Time: O(N), Space: O(1)
func EquilibriumIndexOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func EquilibriumIndex(arr []int) int {
	return EquilibriumIndexOptimized(arr)
}

// 6. Max Difference such that j > i

// Brute Force Approach: Check all possible pairs
// Time: O(N^2), Space: O(1)
func MaxDifferenceBruteForce(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	maxDiff := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			diff := arr[j] - arr[i]
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	return maxDiff
}

// Optimized Approach: Track minimum so far while scanning
// Time: O(N), Space: O(1)
func MaxDifferenceOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func MaxDifference(arr []int) int {
	return MaxDifferenceOptimized(arr)
}

// 7. Check if Array Elements are Consecutive

// Brute Force Approach: For each element, check if all consecutive numbers exist
// Time: O(N^2), Space: O(1)
func CheckConsecutiveBruteForce(arr []int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	// Find min and max
	minVal, maxVal := arr[0], arr[0]
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}
	// Check if range matches array size
	if maxVal-minVal+1 != n {
		return false
	}
	// Check for duplicates
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}

// Optimized Approach: Use hash set
// Time: O(N), Space: O(N) using Set
func CheckConsecutiveOptimized(arr []int) bool {
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

// Legacy function for backward compatibility
func CheckConsecutive(arr []int) bool {
	return CheckConsecutiveOptimized(arr)
}

// 8. Find the Duplicate Number

// Brute Force Approach: Use hash set to track seen numbers
// Time: O(N), Space: O(N)
func FindDuplicateNumberBruteForce(arr []int) int {
	seen := make(map[int]bool)
	for _, val := range arr {
		if seen[val] {
			return val
		}
		seen[val] = true
	}
	return -1 // Should never reach here for valid input
}

// Optimized Approach: Floyd's Cycle Detection
// Time: O(N), Space: O(1)
func FindDuplicateNumberOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func FindDuplicateNumber(arr []int) int {
	return FindDuplicateNumberOptimized(arr)
}
