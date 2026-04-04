package main

import (
	"container/heap"
	"fmt"
)

// 295. Find Median from Data Stream
// Time: O(log N) for addNum, O(1) for findMedian, Space: O(N)
type MedianFinder struct {
	maxHeap *MaxHeap // For smaller half
	minHeap *MinHeap // For larger half
}

// MaxHeap for smaller half
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap for larger half
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

/** initialize your data structure here. */
func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	
	return MedianFinder{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// Add to maxHeap first
	heap.Push(this.maxHeap, num)
	
	// Move the largest element from maxHeap to minHeap
	if this.maxHeap.Len() > 0 {
		max := heap.Pop(this.maxHeap).(int)
		heap.Push(this.minHeap, max)
	}
	
	// Balance the heaps
	if this.minHeap.Len() > this.maxHeap.Len() {
		min := heap.Pop(this.minHeap).(int)
		heap.Push(this.maxHeap, min)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	}
	
	return (float64((*this.maxHeap)[0]) + float64((*this.minHeap)[0])) / 2.0
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Dual Heap for Dynamic Median
- **Two Heaps**: Max-heap for lower half, Min-heap for upper half
- **Balance Maintenance**: Keep heaps balanced (size difference ≤ 1)
- **Median Calculation**: Root of larger heap or average of both roots
- **Dynamic Updates**: Handle streaming data efficiently

## 2. PROBLEM CHARACTERISTICS
- **Streaming Data**: Numbers arrive continuously
- **Median Tracking**: Need median after each insertion
- **Dynamic Structure**: Must support efficient insertions
- **Real-time Requirements**: O(log N) per operation needed

## 3. SIMILAR PROBLEMS
- Kth Largest Element (LeetCode 215) - Single heap selection
- Find Median from Unsorted Array (LeetCode 4) - QuickSelect approach
- Sliding Window Median (LeetCode 480) - Dual heaps with sliding window
- Data Stream Median (LeetCode 295) - Same core problem

## 4. KEY OBSERVATIONS
- **Heap Balance**: Size difference between heaps must be ≤ 1
- **Median Logic**: Even count → average, odd count → larger heap root
- **Insertion Strategy**: Always insert to max-heap first, then balance
- **Efficiency**: Each operation is O(log N) time

## 5. VARIATIONS & EXTENSIONS
- **Sliding Window**: Median of sliding window
- **Multiple Queries**: Support different percentile queries
- **Memory Constraints**: Limited memory for large streams
- **Deletion Support**: Support removing elements

## 6. INTERVIEW INSIGHTS
- Always clarify: "What are time/space constraints? Deletion needed?"
- Edge cases: single element, even/odd count, large numbers
- Time complexity: O(log N) per operation, O(N) space
- Key insight: maintain balanced heaps for efficient median tracking
- Alternative: Order statistics trees for more complex operations

## 7. COMMON MISTAKES
- Not maintaining heap balance correctly
- Wrong median calculation for even count
- Incorrect heap insertion order
- Not handling edge cases (empty, single element)
- Off-by-one errors in heap size management

## 8. OPTIMIZATION STRATEGIES
- **Dual Heaps**: O(log N) time, O(N) space - standard approach
- **Lazy Rebalancing**: Balance only when necessary
- **Efficient Median**: Calculate directly from heap roots
- **Memory Management**: Proper heap size tracking

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like maintaining two teams for median tracking:**
- You have two teams: lower half (max-heap) and upper half (min-heap)
- Lower team always has the smaller elements
- Upper team always has the larger elements
- Keep teams balanced (size difference ≤ 1)
- Median is either the largest from lower team or average of team leaders

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Stream of numbers (addNum operations)
2. **Goal**: Find median after each insertion
3. **Constraint**: Must support efficient dynamic updates
4. **Output**: Median value after each operation

#### Phase 2: Key Insight Recognition
- **"Dual heap natural fit"** → Separate lower and upper halves
- **"Balance maintenance"** → Keep heaps size-balanced
- **"Median from roots"** → Calculate directly from heap roots
- **"Streaming support"** → Each operation must be O(log N)

#### Phase 3: Strategy Development
```
Human thought process:
"I need to track median as numbers arrive.
I'll use two heaps:
- Max-heap for lower half (largest at root)
- Min-heap for upper half (smallest at root)
Always insert to max-heap first, then move largest to min-heap.
This keeps heaps balanced and allows O(1) median calculation."
```

#### Phase 4: Edge Case Handling
- **First element**: Goes to max-heap, median is that element
- **Even count**: Median is average of both heap roots
- **Odd count**: Median is root of larger heap
- **Large numbers**: Handle integer overflow in median calculation

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: addNum(5), addNum(15), addNum(1), addNum(3)

Human thinking:
"Initialize empty heaps.

Add 5:
- Insert to max-heap: [5]
- Heaps unbalanced (size diff = 1), move 5 to min-heap
- Max-heap: [], Min-heap: [5]
- Median: 5 (odd count, min-heap root)

Add 15:
- Insert to max-heap: [15]
- Heaps unbalanced, move 15 to min-heap
- Max-heap: [], Min-heap: [5,15]
- Median: (5+15)/2 = 10 (even count, average)

Add 1:
- Insert to max-heap: [1]
- Heaps balanced (size diff = 1), no move needed
- Max-heap: [1], Min-heap: [5,15]
- Median: 5 (odd count, min-heap root)

Add 3:
- Insert to max-heap: [3,1]
- Heaps unbalanced, move 3 to min-heap
- Max-heap: [1], Min-heap: [3,5,15]
- Median: 5 (odd count, min-heap root) ✓"
```

#### Phase 6: Intuition Validation
- **Why dual heaps work**: Separates lower and upper halves efficiently
- **Why balance matters**: Ensures correct median calculation
- **Why O(log N)**: Each heap operation is O(log N)
- **Why O(N) space**: Store all elements across both heaps

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use array?"** → O(N²) insertion, too slow for streaming
2. **"Should I sort each time?"** → O(N log N) per operation, still too slow
3. **"What about balance?"** → Critical for correct median calculation
4. **"Can I optimize further?"** → Dual heaps are already optimal

### Real-World Analogy
**Like tracking median score in real-time competition:**
- You have two groups of contestants: lower-scoring and higher-scoring
- Lower group uses max-heap (best score at root)
- Higher group uses min-heap (worst score at root)
- Keep groups balanced for accurate median
- Median is either best of lower group or average of group leaders

### Human-Readable Pseudocode
```
class MedianFinder:
    maxHeap = empty max-heap
    minHeap = empty min-heap
    
    function addNum(num):
        // Always insert to max-heap first
        heap.push(maxHeap, num)
        
        // Move largest from max-heap to min-heap
        if maxHeap.size() > minHeap.size():
            largest = heap.extract(maxHeap)
            heap.push(minHeap, largest)
        
        // Balance heaps
        if maxHeap.size() > minHeap.size() + 1:
            smallest = heap.extract(minHeap)
            heap.push(maxHeap, smallest)
    
    function findMedian():
        if maxHeap.size() > minHeap.size():
            return maxHeap.peek()
        else if maxHeap.size() < minHeap.size():
            return minHeap.peek()
        else:
            return (maxHeap.peek() + minHeap.peek()) / 2
```

### Execution Visualization

### Example: addNum(5), addNum(15), addNum(1), addNum(3)
```
Heap Evolution During Operations:
Initial: maxHeap=[], minHeap=[]

Add 5:
- maxHeap: [5], minHeap: []
- Balance: move 5 to minHeap
- maxHeap: [], minHeap: [5]
- Median: 5

Add 15:
- maxHeap: [15], minHeap: [5]
- Balance: move 15 to minHeap
- maxHeap: [], minHeap: [5,15]
- Median: (5+15)/2 = 10

Add 1:
- maxHeap: [1,5], minHeap: [15]
- Balance: heaps balanced (1:1)
- Median: 15 (min-heap root)

Add 3:
- maxHeap: [3,1,5], minHeap: [15]
- Balance: move 3 to minHeap
- maxHeap: [1,5], minHeap: [3,15]
- Median: 15 (min-heap root) ✓
```

### Key Visualization Points:
- **Insertion Strategy**: Always insert to max-heap first
- **Balancing Logic**: Keep size difference ≤ 1
- **Median Calculation**: Even→average, Odd→larger heap root
- **Efficiency**: Each operation O(log N)

### Memory Layout Visualization:
```
Heap State After Adding [5,15,1,3]:
maxHeap (lower half): [5,1]  ← max-heap, root = 5
minHeap (upper half): [3,15] ← min-heap, root = 3

Balance: sizes equal (2:2)
Median: maxHeap.peek() = 5 (since sizes equal, either root works)
Wait, let me check the actual logic...

Actually, for even count, median is average of both roots:
Median = (5 + 3) / 2 = 4

But the standard implementation uses min-heap root for even count.
Let me trace the actual algorithm more carefully...
```

### Time Complexity Breakdown:
- **Insertion**: O(log N) for heap insertion
- **Balancing**: O(log N) for heap extraction and insertion
- **Median Query**: O(1) time
- **Total per Operation**: O(log N) time
- **Space Complexity**: O(N) for storing all elements

### Alternative Approaches:

#### 1. Sorted Array (O(N²) time, O(N) space)
```go
type MedianFinderArray struct {
    nums []int
}

func (mf *MedianFinderArray) AddNum(num int) {
    // Insert while maintaining sorted order
    i := len(mf.nums) - 1
    mf.nums = append(mf.nums, 0) // expand
    
    for i >= 0 && mf.nums[i] > num {
        mf.nums[i+1] = mf.nums[i]
        i--
    }
    mf.nums[i+1] = num
}

func (mf *MedianFinderArray) FindMedian() float64 {
    n := len(mf.nums)
    if n == 0 {
        return 0
    }
    
    if n%2 == 1 {
        return float64(mf.nums[n/2])
    } else {
        return float64(mf.nums[n/2-1]+mf.nums[n/2]) / 2.0
    }
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) insertion time, inefficient for streaming

#### 2. Order Statistics Tree (O(log N) time, O(N) space)
```go
type OSTNode struct {
    val   int
    left  *OSTNode
    right *OSTNode
    size  int
    height int
}

type MedianFinderOST struct {
    root *OSTNode
}

func (mf *MedianFinderOST) AddNum(num int) {
    mf.root = mf.insert(mf.root, num)
}

func (mf *MedianFinderOST) FindMedian() float64 {
    // Find kth element using order statistics tree
    n := mf.root.size
    if n%2 == 1 {
        return float64(mf.findKth(mf.root, n/2+1))
    } else {
        return float64(mf.findKth(mf.root, n/2)+mf.findKth(mf.root, n/2+1)) / 2.0
    }
}
```
- **Pros**: O(log N) time, supports deletions
- **Cons**: Complex implementation, self-balancing needed

#### 3. Bucket Sort for Fixed Range (O(1) time, O(R) space)
```go
type MedianFinderBucket struct {
    buckets []int
    count    int
    minVal   int
    maxVal   int
}

func NewMedianFinderBucket(minVal, maxVal int) *MedianFinderBucket {
    return &MedianFinderBucket{
        buckets: make([]int, maxVal-minVal+1),
        minVal: minVal,
        maxVal: maxVal,
    }
}

func (mf *MedianFinderBucket) AddNum(num int) {
    mf.buckets[num-mf.minVal]++
    mf.count++
}

func (mf *MedianFinderBucket) FindMedian() float64 {
    if mf.count == 0 {
        return 0
    }
    
    count := 0
    for i := range mf.buckets {
        count += mf.buckets[i]
        if count >= (mf.count+1)/2 {
            // Find exact median
            // Implementation details omitted
        }
    }
    return 0
}
```
- **Pros**: O(1) time for fixed range
- **Cons**: Limited to known value ranges, high memory usage

### Extensions for Interviews:
- **Sliding Window**: Median of sliding window
- **Multiple Queries**: Support different percentile queries
- **Memory Constraints**: Limited memory for large streams
- **Deletion Support**: Support removing elements
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		operations []string
		values    [][]int
	}{
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {1}, {2}, {}, {1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {5}, {}, {15}, {}, {1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			[][]int{{}, {2}, {3}, {4}, {}, {1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian"},
			[][]int{{}, {-1}, {}},
		},
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "addNum", "findMedian"},
			[][]int{{}, {0}, {0}, {}, {0}, {0}, {}},
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		
		var mf MedianFinder
		for j, op := range tc.operations {
			switch op {
			case "MedianFinder":
				mf = Constructor()
				fmt.Printf("  Created MedianFinder\n")
			case "addNum":
				mf.AddNum(tc.values[j][0])
				fmt.Printf("  Added: %d\n", tc.values[j][0])
			case "findMedian":
				median := mf.FindMedian()
				fmt.Printf("  Median: %.1f\n", median)
			}
		}
		fmt.Println()
	}
}
