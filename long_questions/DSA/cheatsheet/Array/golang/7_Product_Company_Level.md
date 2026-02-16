```go
package main

// 1. Maximum Product Subarray
// Time: O(N), Space: O(1)
func MaxProductSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	// We need to keep track of both min and max because a negative * negative = positive
	curMax, curMin := 1, 1

	for _, n := range nums {
		// If n is zero, the chain breaks. Reset to 1.
		// But before that, update res with 0 (implicitly handled by algo if res was negative,
		// strictly speaking we might want to ensure res >= 0 if 0 is in array, but max(res, curMax) handles it if curMax becomes 0 or we handle n=0 logic).
		// Actually, standard algo:
		if n == 0 {
			curMax, curMin = 1, 1
			if 0 > res {
				res = 0
			}
			continue
		}

		tmp := curMax * n
		curMax = max3(n, n*curMax, n*curMin)
		curMin = min3(n, tmp, n*curMin)

		if curMax > res {
			res = curMax
		}
	}
	return res
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

// 2. Trapping Rain Water
// Time: O(N), Space: O(1)
func TrapRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	maxL, maxR := 0, 0
	ans := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= maxL {
				maxL = height[l]
			} else {
				ans += maxL - height[l]
			}
			l++
		} else {
			if height[r] >= maxR {
				maxR = height[r]
			} else {
				ans += maxR - height[r]
			}
			r--
		}
	}
	return ans
}

// 3. Container With Most Water
// Time: O(N), Space: O(1)
func ContainerMostWater(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		h := min(height[l], height[r])
		area := (r - l) * h
		if area > maxArea {
			maxArea = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

// 4. Subarray Sum Divisible by K
// Time: O(N), Space: O(K)
func SubarrayDivByK(nums []int, k int) int {
	remMap := make(map[int]int)
	remMap[0] = 1
	sum := 0
	count := 0

	for _, n := range nums {
		sum += n
		rem := sum % k
		if rem < 0 {
			rem += k
		}
		count += remMap[rem]
		remMap[rem]++
	}
	return count
}

// 5. Maximum Circular Subarray Sum
// Time: O(N), Space: O(1)
func MaxCircularSum(nums []int) int {
	totalSum := 0
	maxSum := nums[0]
	curMax := 0
	minSum := nums[0]
	curMin := 0

	for _, n := range nums {
		totalSum += n

		// Kadane Max
		curMax += n
		if curMax > maxSum {
			maxSum = curMax
		}
		if curMax < 0 {
			curMax = 0
		}

		// Kadane Min
		curMin += n
		if curMin < minSum {
			minSum = curMin
		}
		if curMin > 0 {
			curMin = 0
		}
	}

	// If all numbers are negative, maxSum is the answer (max single element)
	// totalSum - minSum would be 0, which is incorrect.
	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, totalSum-minSum)
}

// 6. Longest Consecutive Sequence
// Time: O(N), Space: O(N)
func LongestConsecutiveSeq(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}

	maxLen := 0
	for n := range set {
		// Only check start of sequence
		if !set[n-1] {
			currNum := n
			currLen := 1
			for set[currNum+1] {
				currNum++
				currLen++
			}
			if currLen > maxLen {
				maxLen = currLen
			}
		}
	}
	return maxLen
}

// 7. Count Smaller Elements on Right (Using Merge Sort type logic detailed helper needed, omitted full merge sort implementation for brevity, implementing simpler version or placeholder if complex)
// Ideally usually done with Fenwick Tree or Merge Sort.
// Here providing a placeholder or the actual Merge Sort logic if required.
// Let's implement a simpler O(N^2) for now or the actual Merge Sort one which is verbose.
// Given strict one-page constraint, I will implement the efficient Merge Sort approach.

type pair struct {
	val, idx int
}

func CountSmallerRight(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	mergeSortCount(nums, indices, result, 0, n-1)
	return result
}

func mergeSortCount(nums, indices, result []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSortCount(nums, indices, result, left, mid)
	mergeSortCount(nums, indices, result, mid+1, right)

	// Merge
	merge(nums, indices, result, left, mid, right)
}

func merge(nums, indices, result []int, left, mid, right int) {
	// Optimization: copy only needed range
	tempIndices := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	// Logic: We sort INDICES based on values in nums
	// When moving an element from RIGHT subarray to Temp, it means it is smaller than
	// remaining elements in LEFT subarray.
	// Actually, easier logic:
	// When moving element from LEFT subarray, we add count of elements already moved from RIGHT subarray?
	// Or: Standard approach: Count jumps.
	// Let's stick to standard:
	// If nums[indices[i]] <= nums[indices[j]]: NO increment.
	// We want "smaller on right".

	// Better logic for "Smaller on Right":
	// Iterate. If left[i] > right[j], right[j] is smaller. But we don't know how many more.
	// Standard Inversion count logic counts TOTAL.
	// For specific element count, we usually do:
	// When popping from LEFT, we add (j - (mid+1)) to result[indices[i]].
	// meaningful: j is current ptr in right array. (j - (mid+1)) is count of elements from right array already merged (which were smaller).

	i = left
	j = mid + 1
	k = 0
	countRightSmaller := 0

	for i <= mid && j <= right {
		if nums[indices[j]] < nums[indices[i]] {
			// Right element is smaller. Move it to temp.
			tempIndices[k] = indices[j]
			countRightSmaller++
			j++
		} else {
			// Left element is smaller or equal. Move it.
			// It is larger than 'countRightSmaller' elements from right part seen so far.
			result[indices[i]] += countRightSmaller
			tempIndices[k] = indices[i]
			i++
		}
		k++
	}

	for i <= mid {
		result[indices[i]] += countRightSmaller
		tempIndices[k] = indices[i]
		i++
		k++
	}
	for j <= right {
		tempIndices[k] = indices[j]
		j++
		k++
	}

	for i := 0; i < len(tempIndices); i++ {
		indices[left+i] = tempIndices[i]
	}
}

// 8. Stock Buy and Sell (Multiple Transactions)
// Time: O(N), Space: O(1)
func StockBuySell(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
```
