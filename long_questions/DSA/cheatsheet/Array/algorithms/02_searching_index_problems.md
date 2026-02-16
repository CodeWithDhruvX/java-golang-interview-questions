# 2. Searching & Index-Based Problems

## 1. Linear Search

### Algorithm
1. Iterate through each element of the array.
2. If the current element matches the target, return its index.
3. If the loop completes without finding the target, return -1.

### Code
```go
// LinearSearch search for a target value in the array
// Time Complexity: O(N)
// Space Complexity: O(1)
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}
```

---

## 2. Binary Search (Iterative & Recursive)

### Algorithm (Iterative)
1. Initialize `left = 0` and `right = n-1`.
2. Loop while `left <= right`:
   - Calculate `mid = left + (right - left) / 2`.
   - If `arr[mid] == target`, return `mid`.
   - If `arr[mid] < target`, discard left half: `left = mid + 1`.
   - If `arr[mid] > target`, discard right half: `right = mid - 1`.
3. Return -1 if not found.

### Code
```go
// BinarySearchIterative searches for a target in a sorted array
// Time Complexity: O(log N)
// Space Complexity: O(1)
func BinarySearchIterative(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchRecursive searches for a target recursively
// Time Complexity: O(log N)
// Space Complexity: O(log N) (stack space)
func BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	}
	
	if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}
	return BinarySearchRecursive(arr, target, left, mid-1)
}
```

---

## 3. Find First and Last Occurrence of an Element

### Algorithm
1. **First Occurrence**:
   - Use Binary Search.
   - If `arr[mid] == target`, store `mid` as potential answer and move left: `right = mid - 1`.
   - Continue until `left > right`.
2. **Last Occurrence**:
   - Use Binary Search.
   - If `arr[mid] == target`, store `mid` as potential answer and move right: `left = mid + 1`.
   - Continue until `left > right`.
3. Return both indices.

### Code
```go
// FindFirstAndLast returns the first and last position of element in sorted array
// Time Complexity: O(log N)
// Space Complexity: O(1)
func FindFirstAndLast(arr []int, target int) (int, int) {
	first := findBound(arr, target, true)
	if first == -1 {
		return -1, -1
	}
	last := findBound(arr, target, false)
	return first, last
}

func findBound(arr []int, target int, isFirst bool) int {
	left, right := 0, len(arr)-1
	bound := -1
	
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			bound = mid
			if isFirst {
				right = mid - 1 // Search in left half
			} else {
				left = mid + 1 // Search in right half
			}
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return bound
}
```

---

## 4. Count Occurrences of a Number in a Sorted Array

### Algorithm
1. Find the **First Occurrence** index (`first`).
2. If `first == -1`, return 0.
3. Find the **Last Occurrence** index (`last`).
4. Return `last - first + 1`.

### Code
```go
// CountOccurrences returns the number of times target appears in sorted array
// Time Complexity: O(log N)
// Space Complexity: O(1)
func CountOccurrences(arr []int, target int) int {
	first := findBound(arr, target, true) // Reusing helper from above
	if first == -1 {
		return 0
	}
	last := findBound(arr, target, false)
	return last - first + 1
}
```

---

## 5. Find a Missing Number in Range 1...N

### Algorithm
1. Calculate expected sum of `1` to `N` using formula `S = N * (N + 1) / 2`.
2. Calculate actual sum of array elements.
3. `Missing Number = Expected Sum - Actual Sum`.

### Code
```go
// FindMissingNumber finds the missing number in range 0..n
// Time Complexity: O(N)
// Space Complexity: O(1)
func FindMissingNumber(arr []int) int {
	n := len(arr)
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for _, num := range arr {
		actualSum += num
	}
	return expectedSum - actualSum
}
```

---

## 6. Find the Element That Appears Only Once

### Algorithm (XOR Approach)
1. Initialize `result = 0`.
2. XOR each element of the array with `result`.
3. Since duplicates cancel out (`A ^ A = 0`) and `0 ^ B = B`, result will hold the unique element.

### Code
```go
// SingleNumber finds the element that appears only once
// Time Complexity: O(N)
// Space Complexity: O(1)
func SingleNumber(arr []int) int {
	result := 0
	for _, num := range arr {
		result ^= num
	}
	return result
}
```

---

## 7. Find the Peak Element

### Algorithm (Binary Search)
1. Initialize `left = 0`, `right = n - 1`.
2. While `left < right`:
   - Calculate `mid`.
   - If `arr[mid] > arr[mid+1]`, the peak is on the left side (including `mid`). Set `right = mid`.
   - Else, the peak is on the right side. Set `left = mid + 1`.
3. Return `arr[left]`.

### Code
```go
// FindPeakElement finds a peak element where arr[i] > arr[i+1] and arr[i] > arr[i-1]
// Time Complexity: O(log N)
// Space Complexity: O(1)
func FindPeakElement(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left // Returns index of peak
}
```
