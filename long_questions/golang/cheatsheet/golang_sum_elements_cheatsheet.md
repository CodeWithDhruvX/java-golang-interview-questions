# Golang Sum of Elements Cheatsheet

A comprehensive collection of "Sum" related interview questions with concise Golang solutions.

## 🔹 Basic / Beginner

### 1. Find the sum of all elements in an array
```go
func sumArray(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}
```

### 2. Sum of even / odd elements in an array
```go
func sumEvenOdd(arr []int) (int, int) {
    evenSum, oddSum := 0, 0
    for _, v := range arr {
        if v%2 == 0 {
            evenSum += v
        } else {
            oddSum += v
        }
    }
    return evenSum, oddSum
}
```

### 3. Sum of digits of a number
```go
func sumDigits(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n /= 10
    }
    return sum
}
```

### 4. Sum of first N natural numbers
```go
func sumN(n int) int {
    return n * (n + 1) / 2 // Formula: O(1)
}
```

### 5. Sum of elements at even/odd positions
```go
func sumEvenOddPositions(arr []int) (int, int) {
    evenIdxSum, oddIdxSum := 0, 0
    for i, v := range arr {
        if i%2 == 0 {
            evenIdxSum += v
        } else {
            oddIdxSum += v
        }
    }
    return evenIdxSum, oddIdxSum
}
```

---

## 🔹 Intermediate (Very Common)

### 6. Find a pair with a given sum (Two Sum)
```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, v := range nums {
        if idx, ok := m[target-v]; ok {
            return []int{idx, i}
        }
        m[v] = i
    }
    return nil
}
```

### 7. Find all pairs with a given sum
```go
func allPairs(arr []int, target int) [][]int {
    var result [][]int
    seen := make(map[int]int)
    for _, num := range arr {
        diff := target - num
        if count, ok := seen[diff]; ok && count > 0 {
             result = append(result, []int{num, diff})
             seen[diff]-- // simple handling for duplicates logic varies
        } else {
            seen[num]++
        }
    }
    return result
}
```

### 8. Find triplets with a given sum
```go
import "sort"

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] { continue }
        l, r := i+1, len(nums)-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++; r--
                for l < r && nums[l] == nums[l-1] { l++ }
                for l < r && nums[r] == nums[r+1] { r-- }
            } else if sum < 0 {
                l++
            } else {
                r--
            }
        }
    }
    return res
}
```

### 9. Maximum subarray sum (Kadane’s Algorithm)
```go
import "math"

func maxSubArray(nums []int) int {
    maxSum, currentSum := math.MinInt, 0
    for _, v := range nums {
        currentSum += v
        if currentSum > maxSum {
            maxSum = currentSum
        }
        if currentSum < 0 {
            currentSum = 0
        }
    }
    return maxSum
}
```

### 10. Subarray with given sum (Non-negative check)
```go
func subarraySum(nums []int, k int) []int {
    start, currentSum := 0, 0
    for end := 0; end < len(nums); end++ {
        currentSum += nums[end]
        for currentSum > k && start <= end {
            currentSum -= nums[start]
            start++
        }
        if currentSum == k {
            return []int{start, end}
        }
    }
    return nil
}
```

### 11. Check if any subarray has sum = K (Handles negatives)
```go
func checkSubarraySum(nums []int, k int) bool {
    sum := 0
    m := map[int]bool{0: true}
    for _, v := range nums {
        sum += v
        if m[sum-k] {
            return true
        }
        m[sum] = true
    }
    return false
}
```

---

## 🔹 Array / Prefix Sum Based

### 12. Prefix sum array implementation
```go
func prefixSum(arr []int) []int {
    prefix := make([]int, len(arr))
    prefix[0] = arr[0]
    for i := 1; i < len(arr); i++ {
        prefix[i] = prefix[i-1] + arr[i]
    }
    return prefix
}
```

### 13. Range sum queries
```go
type NumArray struct {
    prefix []int
}
func Constructor(nums []int) NumArray {
    p := make([]int, len(nums)+1)
    for i, v := range nums {
        p[i+1] = p[i] + v
    }
    return NumArray{prefix: p}
}
func (this *NumArray) SumRange(left int, right int) int {
    return this.prefix[right+1] - this.prefix[left]
}
```

### 14. Find equilibrium index (left sum = right sum)
```go
func pivotIndex(nums []int) int {
    totalSum, leftSum := 0, 0
    for _, v := range nums { totalSum += v }
    for i, v := range nums {
        if leftSum == totalSum - leftSum - v {
            return i
        }
        leftSum += v
    }
    return -1
}
```

### 15. Count subarrays with sum equal to K
```go
func subarraySumCount(nums []int, k int) int {
    count, sum := 0, 0
    m := map[int]int{0: 1}
    for _, v := range nums {
        sum += v
        if c, ok := m[sum-k]; ok {
            count += c
        }
        m[sum]++
    }
    return count
}
```

---

## 🔹 Advanced / Tricky

### 16. Closest sum to a given value (Three Sum Closest)
```go
import ("sort"; "math")

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    closest := nums[0] + nums[1] + nums[2]
    for i := 0; i < len(nums)-2; i++ {
        l, r := i+1, len(nums)-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if math.Abs(float64(target-sum)) < math.Abs(float64(target-closest)) {
                closest = sum
            }
            if sum < target { l++ } else { r-- }
        }
    }
    return closest
}
```

### 17. Minimum subarray sum (size) ≥ K
```go
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    minLen := n + 1
    left, sum := 0, 0
    for right := 0; right < n; right++ {
        sum += nums[right]
        for sum >= target {
            if right-left+1 < minLen {
                minLen = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }
    if minLen == n+1 { return 0 }
    return minLen
}
```

### 18. Circular subarray maximum sum
```go
func maxSubarraySumCircular(nums []int) int {
    total, maxSum, curMax, minSum, curMin := 0, nums[0], 0, nums[0], 0
    for _, a := range nums {
        curMax = max(curMax+a, a)
        maxSum = max(maxSum, curMax)
        curMin = min(curMin+a, a)
        minSum = min(minSum, curMin)
        total += a
    }
    if maxSum > 0 { return max(maxSum, total-minSum) }
    return maxSum
}
func max(a, b int) int { if a > b { return a }; return b }
func min(a, b int) int { if a < b { return a }; return b }
```

### 19. Partition array into two parts with equal sum
```go
func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums { sum += num }
    if sum%2 != 0 { return false }
    target := sum / 2
    dp := make([]bool, target+1)
    dp[0] = true
    for _, num := range nums {
        for i := target; i >= num; i-- {
            dp[i] = dp[i] || dp[i-num]
        }
    }
    return dp[target]
}
```

### 20. Subset sum problem (DP)
```go
// Same logic as Partition Equal Subset Sum, checking for specific target
func isSubsetSum(arr []int, sum int) bool {
    dp := make([]bool, sum+1)
    dp[0] = true
    for _, val := range arr {
        for j := sum; j >= val; j-- {
            if dp[j-val] {
                dp[j] = true
            }
        }
    }
    return dp[sum]
}
```

---

## 🔹 Math / Logic Based

### 21. Sum without using loops
```go
func sumRecursive(arr []int) int {
    if len(arr) == 0 { return 0 }
    return arr[0] + sumRecursive(arr[1:])
}
```

### 22. Sum without using ‘+’ operator
```go
func getSum(a int, b int) int {
    for b != 0 {
        carry := (a & b) << 1
        a = a ^ b
        b = carry
    }
    return a
}
```

### 23. Find missing number using sum
```go
func missingNumber(nums []int) int {
    n := len(nums)
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    for _, num := range nums {
        actualSum += num
    }
    return expectedSum - actualSum
}
```

### 24. Find duplicate number using sum (1 to N)
```go
// Assume array contains n+1 integers where each is between 1 and n
func findDuplicate(nums []int) int {
    // Only works if there's exactly one duplicate and values 1..n
    // Otherwise use Floyd's Tortoise and Hare
    // Simple sum diff approach:
    n := len(nums) - 1
    expected := n * (n + 1) / 2
    actual := 0
    for _, v := range nums { actual += v }
    return actual - expected
}
```

---

## 🔹 Additional Sum-Related Questions

### 25. Range Sum 2D (Matrix)
```go
type NumMatrix struct {
    dp [][]int
}
func Constructor2D(matrix [][]int) NumMatrix {
    if len(matrix) == 0 { return NumMatrix{} }
    r, c := len(matrix), len(matrix[0])
    dp := make([][]int, r+1)
    for i := range dp { dp[i] = make([]int, c+1) }
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + matrix[i][j]
        }
    }
    return NumMatrix{dp: dp}
}
func (this *NumMatrix) SumRegion(r1, c1, r2, c2 int) int {
    return this.dp[r2+1][c2+1] - this.dp[r1][c2+1] - this.dp[r2+1][c1] + this.dp[r1][c1]
}
```

### 26. Maximum sum of non-adjacent elements (House Robber)
```go
func rob(nums []int) int {
    prev1, prev2 := 0, 0
    for _, num := range nums {
        tmp := prev1
        prev1 = max(prev2 + num, prev1)
        prev2 = tmp
    }
    return prev1
}
```

### 27. Largest subarray with sum = 0
```go
func maxLen(arr []int) int {
    m := make(map[int]int)
    sum, maxL := 0, 0
    for i, v := range arr {
        sum += v
        if sum == 0 {
            maxL = i + 1
        } else if idx, ok := m[sum]; ok {
            if i - idx > maxL { maxL = i - idx }
        } else {
            m[sum] = i
        }
    }
    return maxL
}
```

### 28. Subarray Product Less Than K
```go
func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 { return 0 }
    prod, res, left := 1, 0, 0
    for right, val := range nums {
        prod *= val
        for prod >= k {
            prod /= nums[left]
            left++
        }
        res += right - left + 1
    }
    return res
}
```

### 29. Maximum Sum of K Consecutive Elements (Sliding Window)
```go
func maxSumK(arr []int, k int) int {
    if len(arr) < k { return -1 }
    currSum := 0
    for i := 0; i < k; i++ { currSum += arr[i] }
    maxSum := currSum
    for i := k; i < len(arr); i++ {
        currSum += arr[i] - arr[i-k]
        if currSum > maxSum { maxSum = currSum }
    }
    return maxSum
}
```
