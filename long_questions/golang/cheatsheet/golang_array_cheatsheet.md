# Golang Arrays & Slices Cheatsheet

Quick reference for interview-frequent basic programs involving Arrays and Slices.

---

## 🟢 Generic Helpers

### Max & Min (Integers)
Go 1.21+ has `max()` and `min()`, but for interviews you might need to write them or use them in logic.
```go
func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}
```

---

## 🔵 Array & Slice Algorithms

### 1. Find Max Element in Slice
```go
func FindMax(nums []int) int {
    if len(nums) == 0 {
        return 0 // or error
    }
    maxVal := nums[0]
    for _, v := range nums {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
```

### 2. Reverse a Slice
In-place reversal.
```go
func ReverseSlice(nums []int) {
    for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
```

### 3. Remove Duplicates
Returns a new slice with unique elements.
```go
func RemoveDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    
    for _, num := range nums {
        if !seen[num] {
            seen[num] = true
            result = append(result, num)
        }
    }
    return result
}
```

### 4. Rotate Array (Right by K steps)
Using generic reverse helper. Step 1: Reverse whole. Step 2: Reverse first k. Step 3: Reverse rest.
```go
func Rotate(nums []int, k int) {
    n := len(nums)
    k = k % n
    
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```

### 5. Two Sum (Target Sum)
Find indices of two numbers that add up to target. O(n) using map.
```go
func TwoSum(nums []int, target int) []int {
    seen := make(map[int]int) // val -> index
    
    for i, num := range nums {
        complement := target - num
        if idx, ok := seen[complement]; ok {
            return []int{idx, i}
        }
        seen[num] = i
    }
    return nil
}
```

### 6. Merge Two Sorted Arrays
Merge `nums2` into `nums1` (assuming nums1 has space).
```go
func MergeSorted(nums1 []int, m int, nums2 []int, n int) {
    // Start from end
    i, j, k := m-1, n-1, m+n-1
    
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
    
    // Fill remaining from nums2 (nums1 remainder is already in place)
    for j >= 0 {
        nums1[k] = nums2[j]
        j--
        k--
    }
}
```

### 7. Intersection of Two Arrays
Using map for O(n+m) time.
```go
func Intersection(nums1, nums2 []int) []int {
    set := make(map[int]bool)
    for _, n := range nums1 {
        set[n] = true
    }
    
    var result []int
    for _, n := range nums2 {
        if set[n] {
            result = append(result, n)
            delete(set, n) // ensure unique in result
        }
    }
    return result
}
```
