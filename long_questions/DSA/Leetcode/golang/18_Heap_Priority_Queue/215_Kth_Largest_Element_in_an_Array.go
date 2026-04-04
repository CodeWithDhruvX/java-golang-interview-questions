package main

import (
	"container/heap"
	"fmt"
)

// 215. Kth Largest Element in an Array
// Time: O(N log K), Space: O(K)
func findKthLargest(nums []int, k int) int {
	// Use a min-heap of size k
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	
	for _, num := range nums {
		heap.Push(minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	
	return (*minHeap)[0]
}

// MinHeap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Alternative solution using QuickSelect (O(N) average case)
func findKthLargestQuickSelect(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, kthSmallest int) int {
	if left == right {
		return nums[left]
	}
	
	pivotIndex := partition(nums, left, right)
	
	if kthSmallest == pivotIndex {
		return nums[pivotIndex]
	} else if kthSmallest < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, kthSmallest)
	} else {
		return quickSelect(nums, pivotIndex+1, right, kthSmallest)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Min-Heap for Kth Largest Selection
- **Min-Heap Structure**: Maintain heap of size K with largest elements
- **Heap Operations**: Insert O(log K), Extract O(log K)
- **Kth Largest**: Root contains Kth largest when heap size = K
- **Space Optimization**: Only store K elements instead of sorting all

## 2. PROBLEM CHARACTERISTICS
- **Selection Problem**: Find Kth largest element in array
- **Order Statistics**: Kth element in sorted order
- **Efficient Selection**: Avoid full sorting O(N log N)
- **Heap Size**: Never exceeds K elements

## 3. SIMILAR PROBLEMS
- Find Median from Data Stream (LeetCode 295) - Dual heap approach
- K Closest Points (LeetCode 973) - Heap for nearest neighbors
- Merge K Sorted Lists (LeetCode 23) - K-way merge with heap
- Top K Frequent Elements (LeetCode 347) - Heap for frequency counting

## 4. KEY OBSERVATIONS
- **Heap Size**: Exactly K elements when fully processed
- **Root Element**: Always Kth largest when heap is full
- **Insertion Logic**: If heap size < K, push directly
- **Maintenance Logic**: If heap size = K and new element > root, replace
- **Final Result**: Root of heap when all elements processed

## 5. VARIATIONS & EXTENSIONS
- **Kth Smallest**: Use max-heap instead of min-heap
- **Streaming Data**: Handle infinite data streams
- **Multiple Queries**: Answer many Kth largest queries
- **Dynamic Updates**: Support insertions and deletions

## 6. INTERVIEW INSIGHTS
- Always clarify: "What if K > array length? Multiple queries?"
- Edge cases: empty array, K=1, K=array length
- Time complexity: O(N log K) time, O(K) space
- Key insight: heap size never exceeds K
- Alternative: QuickSelect algorithm with O(N) average time

## 7. COMMON MISTAKES
- Using max-heap instead of min-heap for Kth largest
- Wrong heap comparison (using < instead of >)
- Not handling heap size correctly
- Off-by-one errors in K handling
- Not returning correct element from heap

## 8. OPTIMIZATION STRATEGIES
- **Min-Heap**: O(N log K) time, O(K) space - standard approach
- **Early Termination**: Not applicable (need all elements)
- **QuickSelect**: O(N) average time, O(1) space - alternative
- **Max-Heap**: For Kth smallest problems

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the Kth tallest person in a line:**
- You have a line of people with different heights
- You want to find the Kth tallest person
- Keep track of the K tallest people seen so far
- Use a min-heap to always know the shortest among the K tallest
- When you see someone taller than the shortest in your K group, replace them

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers and integer K
2. **Goal**: Find Kth largest element
3. **Constraint**: 1 ≤ K ≤ array length
4. **Output**: Kth largest element value

#### Phase 2: Key Insight Recognition
- **"Heap natural fit"** → Need to maintain K largest elements efficiently
- **"Min-heap for largest"** → Counterintuitive but correct
- **"Size maintenance"** → Keep exactly K elements
- **"Root as answer"** → Root is Kth largest when heap is full

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the Kth largest element.
I'll maintain a min-heap of size K containing the K largest seen so far.
For each number:
- If heap has fewer than K elements, just add it
- If heap has K elements and new number > heap root, replace root
- If heap has K elements and new number ≤ heap root, ignore it
At the end, the heap root is the Kth largest!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return error or handle appropriately
- **K = 1**: Return maximum element
- **K = array length**: Return minimum element
- **Invalid K**: Handle K > array length

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [3,2,1,5,6,4], K = 2

Human thinking:
"I'll process each number and maintain a min-heap of size 2:

Initialize empty heap.

Process 3:
- Heap size < 2, push 3
- Heap: [3]

Process 2:
- Heap size < 2, push 2
- Heap: [2,3] (min-heap, root = 2)

Process 1:
- Heap size = 2, 1 ≤ root(2)
- Ignore 1 (not in top 2)
- Heap: [2,3]

Process 5:
- Heap size = 2, 5 > root(2)
- Extract root (2), push 5
- Heap: [3,5]

Process 6:
- Heap size = 2, 6 > root(3)
- Extract root (3), push 6
- Heap: [5,6]

Process 4:
- Heap size = 2, 4 ≤ root(5)
- Ignore 4 (not in top 2)
- Heap: [5,6]

Final heap root: 5 ✓ Kth largest element"
```

#### Phase 6: Intuition Validation
- **Why min-heap works**: Root is smallest among K largest
- **Why O(N log K)**: Each of N operations costs O(log K)
- **Why O(K) space**: Heap never stores more than K elements
- **Why optimal**: Better than O(N log N) full sorting

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N log K), K is usually smaller
2. **"Should I use max-heap?"** → For Kth smallest, not Kth largest
3. **"What about QuickSelect?"** → O(N) average time, but O(N²) worst case
4. **"Can I optimize further?"** → Min-heap is already optimal for this problem

### Real-World Analogy
**Like finding the Kth highest score in a competition:**
- You have contestants with different scores
- You want to find the Kth highest score
- Keep track of the K highest scores seen so far
- Use a ranking system to always know the lowest among the K highest
- When someone scores higher than the lowest in your K group, replace them
- The lowest score in your K group is the Kth highest overall

### Human-Readable Pseudocode
```
function findKthLargest(nums, k):
    if nums is empty or k <= 0:
        return error
    
    minHeap = empty min-heap
    
    for num in nums:
        if minHeap.size() < k:
            heap.push(minHeap, num)
        else if num > minHeap.peek():
            heap.extract(minHeap)
            heap.push(minHeap, num)
    
    return minHeap.peek()
```

### Execution Visualization

### Example: nums = [3,2,1,5,6,4], K = 2
```
Heap Evolution During Processing:
Initial: empty heap

Process 3: heap = [3]
Process 2: heap = [2,3] (root = 2)
Process 1: 1 ≤ root(2), ignore → heap = [2,3]
Process 5: 5 > root(2), replace → heap = [3,5] (root = 3)
Process 6: 6 > root(3), replace → heap = [5,6] (root = 5)
Process 4: 4 ≤ root(5), ignore → heap = [5,6]

Final heap: [5,6], root = 5 ✓ Kth largest
```

### Key Visualization Points:
- **Heap Size**: Never exceeds K elements
- **Root Element**: Always Kth largest when heap is full
- **Replacement Logic**: Only replace if new element > root
- **Final Result**: Root contains Kth largest element

### Memory Layout Visualization:
```
Heap State During Processing:
nums = [3,2,1,5,6,4], K = 2

Step-by-step heap evolution:
[3]           ← Process 3
[2,3]         ← Process 2 (root = 2)
[2,3]         ← Process 1 (1 ≤ 2, ignore)
[3,5]         ← Process 5 (5 > 2, replace)
[5,6]         ← Process 6 (6 > 3, replace)
[5,6]         ← Process 4 (4 ≤ 5, ignore)

Final: [5,6], root = 5 ✓ 2nd largest
```

### Time Complexity Breakdown:
- **Processing**: N elements, each O(log K) heap operation
- **Total Time**: O(N log K) where N is array length
- **Space Complexity**: O(K) for heap storage
- **Optimal**: Better than O(N log N) sorting when K << N

### Alternative Approaches:

#### 1. Full Sorting (O(N log N) time, O(1) space)
```go
func findKthLargestSort(nums []int, k int) int {
    sort.Ints(nums)
    return nums[len(nums)-k]
}
```
- **Pros**: Simple to implement
- **Cons**: O(N log N) time, sorts entire array unnecessarily

#### 2. QuickSelect Algorithm (O(N) average time, O(1) space)
```go
func findKthLargestQuickSelect(nums []int, k int) int {
    // Find (len(nums) - k)th smallest using QuickSelect
    return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, kthSmallest int) int {
    if left == right {
        return nums[left]
    }
    
    pivotIndex := partition(nums, left, right)
    
    if kthSmallest == pivotIndex {
        return nums[pivotIndex]
    } else if kthSmallest < pivotIndex {
        return quickSelect(nums, left, pivotIndex-1, kthSmallest)
    } else {
        return quickSelect(nums, pivotIndex+1, right, kthSmallest)
    }
}

func partition(nums []int, left, right int) int {
    pivot := nums[right]
    i := left
    
    for j := left; j < right; j++ {
        if nums[j] <= pivot {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    
    nums[i], nums[right] = nums[right], nums[i]
    return i
}
```
- **Pros**: O(N) average time, optimal
- **Cons**: O(N²) worst case, more complex

#### 3. Max-Heap Approach (O(N log K) time, O(K) space)
```go
func findKthLargestMaxHeap(nums []int, k int) int {
    // For Kth smallest, use max-heap
    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)
    
    for i := 0; i < len(nums)-k+1; i++ {
        heap.Push(maxHeap, nums[i])
    }
    
    for i := len(nums)-k+1; i < len(nums); i++ {
        if nums[i] < (*maxHeap)[0] {
            heap.Pop(maxHeap)
            heap.Push(maxHeap, nums[i])
        }
    }
    
    return (*maxHeap)[0]
}
```
- **Pros**: Same complexity as min-heap
- **Cons**: More complex logic, less intuitive

### Extensions for Interviews:
- **Kth Smallest**: Use max-heap instead of min-heap
- **Streaming Data**: Handle infinite data streams
- **Multiple Queries**: Answer many Kth largest queries efficiently
- **Dynamic Updates**: Support insertions and deletions
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
		{[]int{1}, 1},
		{[]int{2, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 1},
		{[]int{7, 10, 4, 3, 20, 15}, 3},
		{[]int{-1, -2, -3, -4, -5}, 2},
		{[]int{100, 200, 300, 400, 500}, 4},
	}
	
	for i, tc := range testCases {
		// Make copies for both methods
		nums1 := make([]int, len(tc.nums))
		copy(nums1, tc.nums)
		nums2 := make([]int, len(tc.nums))
		copy(nums2, tc.nums)
		
		result1 := findKthLargest(nums1, tc.k)
		result2 := findKthLargestQuickSelect(nums2, tc.k)
		
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Heap: %d, QuickSelect: %d\n", 
			i+1, tc.nums, tc.k, result1, result2)
	}
}
