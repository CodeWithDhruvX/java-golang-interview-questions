package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 973. K Closest Points to Origin
// Time: O(N log K), Space: O(K)
func kClosest(points [][]int, k int) [][]int {
	// Use a max-heap of size k to keep k closest points
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	
	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		heap.Push(maxHeap, Point{point, distance})
		
		if maxHeap.Len() > k {
			heap.Pop(maxHeap)
		}
	}
	
	// Extract points from heap
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(maxHeap).(Point).coords
	}
	
	return result
}

type Point struct {
	coords   []int
	distance int
}

// MaxHeap implementation for Point
type MaxHeap []Point

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Alternative solution using QuickSelect (O(N) average case)
func kClosestQuickSelect(points [][]int, k int) [][]int {
	distances := make([]int, len(points))
	for i, point := range points {
		distances[i] = point[0]*point[0] + point[1]*point[1]
	}
	
	kth := quickSelect(distances, 0, len(distances)-1, k)
	
	result := make([][]int, 0, k)
	for i, point := range points {
		if distances[i] <= kth {
			result = append(result, point)
		}
	}
	
	return result
}

func quickSelect(distances []int, left, right, k int) int {
	if left == right {
		return distances[left]
	}
	
	pivotIndex := partition(distances, left, right)
	
	if k == pivotIndex {
		return distances[pivotIndex]
	} else if k < pivotIndex {
		return quickSelect(distances, left, pivotIndex-1, k)
	} else {
		return quickSelect(distances, pivotIndex+1, right, k)
	}
}

func partition(distances []int, left, right int) int {
	pivot := distances[right]
	i := left
	
	for j := left; j < right; j++ {
		if distances[j] <= pivot {
			distances[i], distances[j] = distances[j], distances[i]
			i++
		}
	}
	
	distances[i], distances[right] = distances[right], distances[i]
	return i
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Max-Heap for K Nearest Neighbors
- **Max-Heap Structure**: Maintain heap of size K with closest points
- **Distance Calculation**: Euclidean distance from origin (0,0)
- **Heap Operations**: Insert O(log K), Extract O(log K)
- **K Selection**: Keep K closest points in heap

## 2. PROBLEM CHARACTERISTICS
- **Spatial Problem**: Find K points closest to origin
- **Distance Metric**: Euclidean distance sqrt(x² + y²)
- **Selection Problem**: Choose K out of N points
- **Heap Size**: Never exceeds K elements

## 3. SIMILAR PROBLEMS
- Kth Largest Element (LeetCode 215) - Heap selection
- Find Median from Data Stream (LeetCode 295) - Dual heap approach
- Merge K Sorted Lists (LeetCode 23) - K-way merge with heap
- Top K Frequent Elements (LeetCode 347) - Heap for frequency counting

## 4. KEY OBSERVATIONS
- **Distance Comparison**: Can compare squared distances to avoid sqrt
- **Heap Size**: Exactly K elements when fully processed
- **Root Element**: Always farthest among K closest when heap is full
- **Insertion Logic**: If heap size < K, push directly
- **Maintenance Logic**: If heap size = K and new point is closer, replace

## 5. VARIATIONS & EXTENSIONS
- **Different Metrics**: Manhattan distance, custom distance functions
- **3D Points**: Extend to 3D coordinates
- **Multiple Queries**: Answer many K closest queries
- **Streaming Points**: Handle infinite point streams

## 6. INTERVIEW INSIGHTS
- Always clarify: "What distance metric? Ties? K > points?"
- Edge cases: empty array, K=1, K=points length
- Time complexity: O(N log K) time, O(K) space
- Key insight: heap size never exceeds K
- Optimization: compare squared distances to avoid sqrt

## 7. COMMON MISTAKES
- Using min-heap instead of max-heap for K closest
- Not handling ties correctly
- Wrong distance calculation
- Off-by-one errors in K handling
- Not returning correct points from heap

## 8. OPTIMIZATION STRATEGIES
- **Max-Heap**: O(N log K) time, O(K) space - standard approach
- **Squared Distances**: Avoid sqrt for comparison
- **QuickSelect**: O(N) average time, O(1) space - alternative
- **Early Termination**: Not applicable (need all points)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding K nearest neighbors in a neighborhood:**
- You have houses (points) in a city
- You want to find K houses closest to city center (origin)
- Keep track of K closest houses seen so far
- Use a max-heap to always know the farthest among your K closest
- When you see a house closer than your farthest in K group, replace it

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of points [x,y] and integer K
2. **Goal**: Find K points closest to origin (0,0)
3. **Distance**: Euclidean distance sqrt(x² + y²)
4. **Output**: Array of K closest points

#### Phase 2: Key Insight Recognition
- **"Heap natural fit"** → Need to maintain K closest points efficiently
- **"Max-heap for closest"** → Counterintuitive but correct
- **"Size maintenance"** → Keep exactly K elements
- **"Root as farthest"** → Root is farthest among K closest

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find K closest points to origin.
I'll maintain a max-heap of size K containing the K closest seen so far.
For each point:
- Calculate its distance from origin
- If heap has fewer than K elements, just add it
- If heap has K elements and new point is closer than root, replace root
- If heap has K elements and new point is farther, ignore it
At the end, heap contains K closest points!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **K = 1**: Return single closest point
- **K = array length**: Return all points
- **Invalid K**: Handle K > array length

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
points = [[1,3], [-2,2]], K = 2

Human thinking:
"I'll process each point and maintain a max-heap of size 2:

Initialize empty heap.

Process [1,3]:
- Distance = sqrt(1² + 3²) = sqrt(10) ≈ 3.16
- Heap size < 2, push (3.16, [1,3])
- Heap: [(3.16, [1,3])]

Process [-2,2]:
- Distance = sqrt((-2)² + 2²) = sqrt(8) ≈ 2.83
- Heap size < 2, push (2.83, [-2,2])
- Heap: [(2.83, [-2,2]), (3.16, [1,3])] (max-heap, root = 3.16)

Final heap contains 2 closest points ✓"
```

#### Phase 6: Intuition Validation
- **Why max-heap works**: Root is farthest among K closest
- **Why O(N log K)**: Each of N operations costs O(log K)
- **Why O(K) space**: Heap never stores more than K elements
- **Why optimal**: Better than O(N log N) full sorting

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort all?"** → O(N log N) vs O(N log K), K is usually smaller
2. **"Should I use min-heap?"** → For K farthest, not K closest
3. **"What about QuickSelect?"** → O(N) average time, but O(N²) worst case
4. **"Can I optimize further?"** → Max-heap is already optimal for this problem

### Real-World Analogy
**Like finding K nearest stores to a customer location:**
- You have store locations in a city
- Customer wants to find K nearest stores to their location
- Keep track of K closest stores seen so far
- Use a ranking system to always know the farthest among K closest
- When you find a store closer than the farthest in your K group, replace them
- The K stores in your group are the K nearest to the customer

### Human-Readable Pseudocode
```
function kClosest(points, k):
    if points is empty or k <= 0:
        return []
    
    maxHeap = empty max-heap
    
    for point in points:
        distance = point.x² + point.y²  // squared distance
        
        if maxHeap.size() < k:
            heap.push(maxHeap, (distance, point))
        else if distance < maxHeap.peek().distance:
            heap.extract(maxHeap)
            heap.push(maxHeap, (distance, point))
    
    result = []
    while maxHeap is not empty:
        result.append(heap.extract(maxHeap).point)
    
    return result
```

### Execution Visualization

### Example: points = [[1,3], [-2,2]], K = 2
```
Heap Evolution During Processing:
Initial: empty heap

Process [1,3]:
- Distance = 1² + 3² = 10
- Heap: [(10, [1,3])]

Process [-2,2]:
- Distance = (-2)² + 2² = 8
- Heap: [(8, [-2,2]), (10, [1,3])] (max-heap, root = 10)

Final heap contains 2 closest points ✓
```

### Key Visualization Points:
- **Distance Calculation**: Use squared distance to avoid sqrt
- **Heap Size**: Never exceeds K elements
- **Root Element**: Always farthest among K closest
- **Replacement Logic**: Only replace if new point is closer than root

### Memory Layout Visualization:
```
Heap State During Processing:
points = [[1,3], [-2,2]], K = 2

Step-by-step heap evolution:
[(10, [1,3])]           ← Process [1,3]
[(8, [-2,2]), (10, [1,3])] ← Process [-2,2]

Final: [(8, [-2,2]), (10, [1,3])]
Root = 10 (farthest among 2 closest)
Extract in reverse order: [[1,3], [-2,2]] ✓
```

### Time Complexity Breakdown:
- **Processing**: N points, each O(log K) heap operation
- **Distance Calculation**: O(1) per point
- **Total Time**: O(N log K) where N is number of points
- **Space Complexity**: O(K) for heap storage
- **Optimal**: Better than O(N log N) sorting when K << N

### Alternative Approaches:

#### 1. Full Sorting (O(N log N) time, O(1) space)
```go
func kClosestSort(points [][]int, k int) [][]int {
    sort.Slice(points, func(i, j int) bool {
        distI := points[i][0]*points[i][0] + points[i][1]*points[i][1]
        distJ := points[j][0]*points[j][0] + points[j][1]*points[j][1]
        return distI < distJ
    })
    
    return points[:k]
}
```
- **Pros**: Simple to implement
- **Cons**: O(N log N) time, sorts entire array unnecessarily

#### 2. QuickSelect Algorithm (O(N) average time, O(1) space)
```go
func kClosestQuickSelect(points [][]int, k int) [][]int {
    if len(points) <= k {
        return points
    }
    
    // Find kth closest using QuickSelect on distances
    kthDistance := quickSelect(points, 0, len(points)-1, k)
    
    // Partition points around kth distance
    result := make([][]int, 0, k)
    for _, point := range points {
        dist := point[0]*point[0] + point[1]*point[1]
        if dist <= kthDistance {
            result = append(result, point)
            if len(result) == k {
                break
            }
        }
    }
    
    return result
}

func quickSelect(points [][]int, left, right, kth int) int {
    // QuickSelect implementation for kth smallest distance
    if left == right {
        dist := points[left][0]*points[left][0] + points[left][1]*points[left][1]
        return dist
    }
    
    pivotIndex := partition(points, left, right)
    
    if kth == pivotIndex {
        dist := points[pivotIndex][0]*points[pivotIndex][0] + points[pivotIndex][1]*points[pivotIndex][1]
        return dist
    } else if kth < pivotIndex {
        return quickSelect(points, left, pivotIndex-1, kth)
    } else {
        return quickSelect(points, pivotIndex+1, right, kth)
    }
}
```
- **Pros**: O(N) average time, optimal
- **Cons**: O(N²) worst case, more complex

#### 3. Min-Heap with Size Limit (O(N log K) time, O(K) space)
```go
func kClosestMinHeap(points [][]int, k int) [][]int {
    if len(points) <= k {
        return points
    }
    
    // Use min-heap but keep only K elements
    minHeap := &MinHeap{}
    heap.Init(minHeap)
    
    for _, point := range points {
        dist := point[0]*point[0] + point[1]*point[1]
        heap.Push(minHeap, dist)
        
        if minHeap.Len() > k {
            heap.Pop(minHeap) // Remove farthest
        }
    }
    
    result := make([][]int, k)
    for i := 0; i < k; i++ {
        result[i] = heap.Pop(minHeap)
    }
    
    return result
}
```
- **Pros**: Same complexity as max-heap
- **Cons**: More complex logic, less intuitive

### Extensions for Interviews:
- **Different Metrics**: Manhattan distance, custom distance functions
- **3D Points**: Extend to 3D coordinates
- **Multiple Queries**: Answer many K closest queries efficiently
- **Streaming Points**: Handle infinite point streams
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		points [][]int
		k      int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
		{[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 0}, {0, 1}, {1, 1}}, 2},
		{[][]int{{-5, -5}, {-4, -4}, {-3, -3}, {-2, -2}, {-1, -1}}, 3},
		{[][]int{{100, 100}, {-100, -100}, {100, -100}, {-100, 100}}, 2},
		{[][]int{{0, 1}, {1, 0}}, 1},
		{[][]int{{2, 2}, {2, 2}, {2, 2}, {2, 2}}, 3},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}, 4},
	}
	
	for i, tc := range testCases {
		// Make copies for both methods
		points1 := make([][]int, len(tc.points))
		copy(points1, tc.points)
		points2 := make([][]int, len(tc.points))
		copy(points2, tc.points)
		
		result1 := kClosest(points1, tc.k)
		result2 := kClosestQuickSelect(points2, tc.k)
		
		fmt.Printf("Test Case %d: points=%v, k=%d\n", i+1, tc.points, tc.k)
		fmt.Printf("  Heap: %v\n", result1)
		fmt.Printf("  QuickSelect: %v\n\n", result2)
	}
}
