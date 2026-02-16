# 7. Advanced / Product Company Level

## 1. Maximum Product Subarray

### Algorithm
1. Initialize `maxProd = arr[0]`, `minProd = arr[0]`, `result = arr[0]`.
2. Iterate from `1` to `n-1`.
3. If `arr[i] < 0`, swap `maxProd` and `minProd`.
4. `maxProd = max(arr[i], maxProd * arr[i])`.
5. `minProd = min(arr[i], minProd * arr[i])`.
6. Update `result = max(result, maxProd)`.

### Code
```go
import "math"

// MaxProductSubarray finds the contiguous subarray with the largest product
// Time Complexity: O(N)
// Space Complexity: O(1)
func MaxProductSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]
	
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		
		if curr < 0 {
			maxProd, minProd = minProd, maxProd
		}
		
		maxProd = max(curr, maxProd*curr)
		minProd = min(curr, minProd*curr)
		
		if maxProd > result {
			result = maxProd
		}
	}
	return result
}
```

---

## 2. Trapping Rain Water

### Algorithm
1. Initialize `left = 0`, `right = n-1`.
2. Initialize `leftMax = 0`, `rightMax = 0`, `res = 0`.
3. Loop while `left <= right`:
   - If `arr[left] <= arr[right]`:
     - If `arr[left] >= leftMax`: `leftMax = arr[left]`.
     - Else: `res += leftMax - arr[left]`.
     - `left++`.
   - Else:
     - If `arr[right] >= rightMax`: `rightMax = arr[right]`.
     - Else: `res += rightMax - arr[right]`.
     - `right--`.

### Code
```go
// TropRainWater computes how much water it can trap after raining
// Time Complexity: O(N)
// Space Complexity: O(1)
func TrapRainWater(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0
	
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
```

---

## 3. Container With Most Water

### Algorithm
1. Initialize `left = 0`, `right = n-1`, `maxArea = 0`.
2. Loop while `left < right`:
   - `width = right - left`.
   - `height = min(arr[left], arr[right])`.
   - `maxArea = max(maxArea, width * height)`.
   - If `arr[left] < arr[right]`: `left++`.
   - Else: `right--`.

### Code
```go
// MaxArea finds two lines that together with the x-axis form a container containing the most water
// Time Complexity: O(N)
// Space Complexity: O(1)
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	
	for left < right {
		h := min(height[left], height[right])
		width := right - left
		area := h * width
		if area > maxArea {
			maxArea = area
		}
		
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
```

---

## 4. Subarray Sums Divisible by K

### Algorithm
1. Initialize `map` (remainder -> count) with `{0: 1}`.
2. Initialize `prefixSum = 0`, `count = 0`.
3. Iterate through array:
   - `prefixSum += arr[i]`.
   - `remainder = prefixSum % K`.
   - If `remainder < 0`: `remainder += K` (handle negative sums).
   - `count += map[remainder]`.
   - `map[remainder]++`.

### Code
```go
// SubarraysDivByK counts subarrays whose sum is divisible by k
// Time Complexity: O(N)
// Space Complexity: O(K)
func SubarraysDivByK(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	count := 0
	
	for _, num := range nums {
		sum += num
		rem := sum % k
		if rem < 0 {
			rem += k
		}
		
		if c, ok := m[rem]; ok {
			count += c
		}
		m[rem]++
	}
	return count
}
```

---

## 5. Maximum Circular Subarray Sum

### Algorithm
1. Calculate `totalSum`.
2. Find `maxSubarraySum` (Kadane).
3. Find `minSubarraySum` (Kadane with inverted reasoning or signs).
4. If `maxSubarraySum < 0` (all numbers negative), return `maxSubarraySum`.
5. Return `max(maxSubarraySum, totalSum - minSubarraySum)`.

### Code
```go
// MaxSubarraySumCircular finds the maximum sum of a circular subarray
// Time Complexity: O(N)
// Space Complexity: O(1)
func MaxSubarraySumCircular(nums []int) int {
	totalSum := 0
	maxSum := nums[0]
	currMax := 0
	minSum := nums[0]
	currMin := 0
	
	for _, num := range nums {
		totalSum += num
		
		currMax = max(currMax+num, num)
		maxSum = max(maxSum, currMax)
		
		currMin = min(currMin+num, num)
		minSum = min(minSum, currMin)
	}
	
	if maxSum < 0 {
		return maxSum
	}
	
	return max(maxSum, totalSum-minSum)
}
```

---

## 6. Longest Consecutive Sequence

### Algorithm
1. Put all elements in a `HashSet`.
2. Iterate through array.
3. If `num - 1` is NOT in set (it's the start of a sequence):
   - Check `num + 1`, `num + 2`... exist in set.
   - Count length.
   - Update `maxLength`.

### Code
```go
// LongestConsecutive finds the length of the longest consecutive elements sequence
// Time Complexity: O(N)
// Space Complexity: O(N)
func LongestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	
	longest := 0
	for num := range set { // Iterate over map to avoid duplicates check
		if !set[num-1] {
			currentNum := num
			currentStreak := 1
			
			for set[currentNum+1] {
				currentNum++
				currentStreak++
			}
			
			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}
```

---

## 7. Count Smaller Elements on Right Side

### Algorithm (Merge Sort based)
1. Use Merge Sort logic.
2. Store indices to track original positions.
3. During merge, if `arr[left] > arr[right]`:
   - It means `arr[left]` is greater than all elements from `right` to end of right half.
   - Or, simply count how many elements jump from right to left side.
   - Actually for "smaller on right", when we pick `arr[left]`, we add count of elements currently picked from right half.
   - Wait, better logic:
     - Merge two sorted halves.
     - When we pick `left[i]`, subsequent elements from `right` are greater? No.
     - When `right[j] < left[i]`, `right[j]` is smaller than `left[i]`. So for `left[i]`, we found a smaller element on right.
     - Count jumps.

### Code
```go
// CountSmaller counts the number of smaller elements to the right of each element
// Time Complexity: O(N log N)
// Space Complexity: O(N)
type Pair struct {
	val, idx int
}

func CountSmaller(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	pairs := make([]Pair, n)
	for i, v := range nums {
		pairs[i] = Pair{v, i}
	}
	
	mergeSort(pairs, result)
	return result
}

func mergeSort(pairs []Pair, result []int) []Pair {
	if len(pairs) <= 1 {
		return pairs
	}
	
	mid := len(pairs) / 2
	left := mergeSort(pairs[:mid], result)
	right := mergeSort(pairs[mid:], result)
	
	merged := make([]Pair, len(pairs))
	i, j, k := 0, 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i].val <= right[j].val {
			// When taking from left, we know j elements from right were smaller
			result[left[i].idx] += j
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}
	
	for i < len(left) {
		result[left[i].idx] += j
		merged[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		merged[k] = right[j]
		j++
		k++
	}
	
	return merged
}
```

---

## 8. Stock Buy and Sell (Best Time to Buy and Sell Stock)

### Algorithm
1. Initialize `minPrice = INF`, `maxProfit = 0`.
2. Iterate through prices.
3. If `price < minPrice`, `minPrice = price`.
4. Else if `price - minPrice > maxProfit`, `maxProfit = price - minPrice`.
5. Return `maxProfit`.

### Code
```go
import "math"

// MaxProfit finds maximum profit from buying and selling stock once
// Time Complexity: O(N)
// Space Complexity: O(1)
func MaxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
```
