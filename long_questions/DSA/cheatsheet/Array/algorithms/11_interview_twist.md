# 11. Interview "Twist" Questions

## 1. Rotate Array Without Using Extra Space

### Algorithm (Reversal - already covered in Basic)
1. Reverse first `n - k` elements.
2. Reverse last `k` elements.
3. Reverse whole array.
   *(Note: This is for Right Rotate. For Left Rotate `k`, reverse `0..k-1`, `k..n-1`, `0..n-1`)*.

### Code
```go
// RotateRight rotates array to the right by k steps
// Time Complexity: O(N)
// Space Complexity: O(1)
func RotateRight(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n
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

---

## 2. Find K-th Smallest Element Without Sorting

### Algorithm (QuickSelect)
1. Choose a pivot (randomly or last element).
2. Partition the array around pivot:
   - Elements smaller than pivot go left.
   - Elements greater go right.
3. If pivot index `p == k-1`, return `arr[p]`.
4. If `p > k-1`, recurse on left part.
5. If `p < k-1`, recurse on right part.

### Code
```go
import "math/rand"

// FindKthSmallest finds the kth smallest element using QuickSelect
// Time Complexity: Avg O(N), Worst O(N^2)
// Space Complexity: O(1) (iterative) or O(log N) stack
func FindKthSmallest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		pivotIndex := partition(nums, left, right)
		if pivotIndex == k-1 {
			return nums[pivotIndex]
		}
		if pivotIndex < k-1 {
			left = pivotIndex + 1
		} else {
			right = pivotIndex - 1
		}
	}
	return -1
}

func partition(nums []int, left, right int) int {
	pivotIdx := left + rand.Intn(right-left+1)
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
```

---

## 3. Find Median of a Stream of Numbers

### Algorithm
1. Use two Heaps:
   - Max-Heap for the lower half.
   - Min-Heap for the upper half.
2. **AddNum**:
   - Add to Max-Heap.
   - Move largest of Max-Heap to Min-Heap.
   - Balance sizes: If Max-Heap has fewer elements than Min-Heap, move smallest of Min-Heap back to Max-Heap.
3. **FindMedian**:
   - If sizes equal, average of tops.
   - If Max-Heap is larger, top of Max-Heap.

### Code
```go
import (
	"container/heap"
)

type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Reverse logic
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	lo *MaxHeap
	hi *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		lo: &MaxHeap{},
		hi: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lo, num)
	heap.Push(this.hi, heap.Pop(this.lo))
	
	if this.lo.Len() < this.hi.Len() {
		heap.Push(this.lo, heap.Pop(this.hi))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64((*this.lo)[0])
	}
	return float64((*this.lo)[0]+(*this.hi)[0]) / 2.0
}
```

---

## 4. Partition Array Into 3 Parts With Equal Sum

### Algorithm
1. Calculate total sum `S`.
2. If `S % 3 != 0`, return false.
3. Target per part `T = S / 3`.
4. Iterate and sum up elements.
5. Identify partitions when `currentSum == T`, `currentSum == 2T`.
6. Ensure we find 3 parts (count >= 3).

### Code
```go
// CanThreePartsEqualSum checks if array can be partitioned into 3 equal sum parts
// Time Complexity: O(N)
// Space Complexity: O(1)
func CanThreePartsEqualSum(arr []int) bool {
	totalSum := 0
	for _, v := range arr {
		totalSum += v
	}
	
	if totalSum%3 != 0 {
		return false
	}
	
	target := totalSum / 3
	currentSum := 0
	partsFound := 0
	
	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		if currentSum == target {
			partsFound++
			currentSum = 0
		}
	}
	
	// We need 3 parts. Note: if target is 0, we can find many parts.
	// But basically if we found at least 3 parts (and last part also sums to target if sum resets), it's true.
	// Correct logic: we need to find 2 cuts.
	// Let's rely on accumulated sum logic which is safer.
	
	return partsFound >= 3
}
```

---

## 5. Minimum Jumps to Reach End of Array (Jump Game II)

### Algorithm (Greedy)
1. Initialize `jumps = 0`, `currentEnd = 0`, `farthest = 0`.
2. Iterate `i` from `0` to `n-2`.
3. `farthest = max(farthest, i + arr[i])`.
4. If `i == currentEnd`:
   - `jumps++`.
   - `currentEnd = farthest`.
   - If `currentEnd >= n-1` break.

### Code
```go
import "math"

// Jump finds min jumps to reach last index
// Time Complexity: O(N)
// Space Complexity: O(1)
func Jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	
	jumps := 0
	currentEnd := 0
	farthest := 0
	
	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		
		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}
```

---

## 6. Check If Array Pairs Are Divisible by K

### Algorithm
1. Count frequencies of `arr[i] % k`.
2. Iterate `i` from `1` to `k/2`.
3. Frequency of remainder `i` must equal frequency of remainder `k-i`.
4. Special case: Remainder `0` count must be even.
5. If `k` is even, Remainder `k/2` count must be even.

### Code
```go
// CanArrange checks if array can be divided into pairs where sum is divisible by k
// Time Complexity: O(N)
// Space Complexity: O(K)
func CanArrange(arr []int, k int) bool {
	freq := make(map[int]int)
	
	for _, val := range arr {
		rem := val % k
		if rem < 0 {
			rem += k
		}
		freq[rem]++
	}
	
	for i := 1; i < k; i++ {
		if freq[i] != freq[k-i] {
			return false
		}
	}
	
	if freq[0]%2 != 0 {
		return false
	}
	
	return true
}
```
