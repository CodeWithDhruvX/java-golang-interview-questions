# 9. Observation-Heavy Problems

## 1. Find the Element With Maximum Repeating Count in O(n) Time

### Algorithm
1. Constraint: `0 <= arr[i] < n`.
2. Iterate through array. For every index `i`, increment `arr[arr[i]%n]` by `n`.
3. The element with maximum frequency will have the maximum value at its corresponding index after updates.
4. Iterate again to find index with max value.
5. Restore array if needed (modulo `n`).

### Code
```go
// MaxRepeatingElement finds the max repeating element in O(N) time and O(1) extra space
// Assumption: 0 <= arr[i] < n
func MaxRepeatingElement(arr []int, k int) int {
	// k is the range/size. If not given, assume n=len(arr) and values < n.
	n := len(arr)
	
	// Increment index
	for i := 0; i < n; i++ {
		idx := arr[i] % k
		arr[idx] += k
	}
	
	maxVal := arr[0]
	result := 0
	
	for i := 1; i < n; i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
			result = i
		}
	}
	
	// Optional: restore array
	// for i := 0; i < n; i++ {
	// 	arr[i] = arr[i] % k
	// }
	
	return result
}
```

---

## 2. Rearrange Array Such That arr[i] = i

### Algorithm
1. Iterate through the array.
2. For each element `i` where `arr[i] != -1` and `arr[i] != i`:
   - While `arr[i] != -1` and `arr[i] != i`:
     - Swap `arr[i]` with `arr[arr[i]]`.
   - Wait, simpler logic:
   - If `arr[i] >= 0` and `arr[i] != i`:
     - Place `arr[i]` to its correct position `arr[arr[i]]`.
     - Careful with infinite loops or overwriting.
     - Swap `arr[i]` and `arr[arr[i]]`.
3. Iterate again, if `arr[i] != i`, set `arr[i] = -1` (if problem implies missing elements are -1).

### Code
```go
// RearrangeArray rearranges such that arr[i] = i return modified array
// Time Complexity: O(N)
// Space Complexity: O(1)
func RearrangeArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for arr[i] != -1 && arr[i] != i {
			correctPos := arr[i]
			
			// If correct position already has the correct value
			if arr[correctPos] == correctPos {
				break
			}
			
			// Swap
			arr[i], arr[correctPos] = arr[correctPos], arr[i]
		}
	}
	return arr
}
```

---

## 3. Minimum Operations to Make Array Equal

### Algorithm
1. Target is usually the median.
2. If `arr[i] = (2 * i) + 1` (as in LC 1551), then:
   - Target is `n`.
   - Operations = Sum of `(n - arr[i])` for first `n/2` elements.
   - Formula simplifies to `n * n / 4` (integer division).

### Code
```go
// MinOperations makes array equal where arr[i] = (2 * i) + 1
// Time Complexity: O(1) via formula
// Space Complexity: O(1)
func MinOperations(n int) int {
	return (n * n) / 4
}
```

---

## 4. Maximum Length Bitonic Subarray

### Algorithm
1. Compute `inc[i]`: Length of increasing subarray ending at `i`.
2. Compute `dec[i]`: Length of decreasing subarray starting at `i`.
3. `maxLen = max(inc[i] + dec[i] - 1)` for all `i`.

### Code
```go
// LongestBitonicSubarray finds the max length of a subarray that is first increasing then decreasing
// Time Complexity: O(N)
// Space Complexity: O(N)
func LongestBitonicSubarray(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	
	inc := make([]int, n)
	dec := make([]int, n)
	
	inc[0] = 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			inc[i] = inc[i-1] + 1
		} else {
			inc[i] = 1
		}
	}
	
	dec[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			dec[i] = dec[i+1] + 1
		} else {
			dec[i] = 1
		}
	}
	
	maxLen := 0
	for i := 0; i < n; i++ {
		if inc[i]+dec[i]-1 > maxLen {
			maxLen = inc[i] + dec[i] - 1
		}
	}
	return maxLen
}
```

---

## 5. Minimum Operations to Make Array Palindrome

### Algorithm
1. Use distinct merging operations.
2. Initialize `left = 0`, `right = n-1`, `ops = 0`.
3. While `left < right`:
   - If `arr[left] == arr[right]`: `left++`, `right--`.
   - If `arr[left] < arr[right]`: Merge `arr[left]` with `arr[left+1]`, `ops++`, `left++`.
   - If `arr[left] > arr[right]`: Merge `arr[right]` with `arr[right-1]`, `ops++`, `right--`.

### Code
```go
// MinMergeOperations returns minimum merge operations to make array palindrome
// Time Complexity: O(N)
// Space Complexity: O(1)
func MinMergeOperations(arr []int) int {
	ans := 0
	i, j := 0, len(arr)-1
	
	for i <= j {
		if arr[i] == arr[j] {
			i++
			j--
		} else if arr[i] < arr[j] {
			// Merge left
			i++
			arr[i] += arr[i-1]
			ans++
		} else {
			// Merge right
			j--
			arr[j] += arr[j+1]
			ans++
		}
	}
	return ans
}
```

---

## 6. Find the Longest Alternating Even-Odd Subarray

### Algorithm
1. Initialize `maxLen = 1`, `currLen = 1`.
2. Iterate `i` from `1` to `n-1`.
3. If `(arr[i] % 2 != arr[i-1] % 2)` (one even, one odd):
   - `currLen++`.
   - `maxLen = max(maxLen, currLen)`.
4. Else:
   - `currLen = 1`.
5. Return `maxLen`.

### Code
```go
// LongestAlternatingSubarray finds max length of subarray with alternating even-odd elements
// Time Complexity: O(N)
// Space Complexity: O(1)
func LongestAlternatingSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	
	maxLen := 1
	currLen := 1
	
	for i := 1; i < len(arr); i++ {
		if (arr[i]%2 == 0 && arr[i-1]%2 != 0) || (arr[i]%2 != 0 && arr[i-1]%2 == 0) {
			currLen++
		} else {
			currLen = 1
		}
		
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
```
