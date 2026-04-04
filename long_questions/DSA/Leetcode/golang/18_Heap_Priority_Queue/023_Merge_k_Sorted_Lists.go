package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int	
	Next *ListNode
}

// 23. Merge k Sorted Lists
// Time: O(N log K), Space: O(K) where N is total nodes, K is number of lists
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	
	minHeap := &NodeHeap{}
	heap.Init(minHeap)
	
	// Push the head of each list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}
	
	dummy := &ListNode{}
	current := dummy
	
	for minHeap.Len() > 0 {
		// Get the smallest node
		smallest := heap.Pop(minHeap).(*ListNode)
		current.Next = smallest
		current = current.Next
		
		// Push the next node from the same list
		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}
	
	return dummy.Next
}

// NodeHeap implementation for ListNode
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Helper function to create a linked list from slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	
	return head
}

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Min-Heap for K-Way Merge
- **Min-Heap Structure**: Always extract smallest element from K lists
- **K-Way Merge**: Merge K sorted lists efficiently
- **Heap Operations**: Insert O(log K), Extract O(log K)
- **Linked List Traversal**: Process all nodes exactly once

## 2. PROBLEM CHARACTERISTICS
- **Multiple Sorted Lists**: K sorted linked lists to merge
- **Sorted Output**: Result must be sorted
- **Efficient Merging**: Avoid O(NK²) naive approach
- **Optimal Structure**: Min-heap provides O(N log K) solution

## 3. SIMILAR PROBLEMS
- Kth Largest Element (LeetCode 215) - Heap selection
- Find Median from Data Stream (LeetCode 295) - Dual heap approach
- K Closest Points (LeetCode 973) - Heap for nearest neighbors
- Merge K Sorted Arrays (LeetCode 23) - Array version

## 4. KEY OBSERVATIONS
- **Heap Size**: At most K elements in heap at any time
- **Sorted Input**: Each list is already sorted
- **Greedy Selection**: Always pick smallest available element
- **Total Elements**: N = sum of all list lengths
- **Complexity Tradeoff**: O(N log K) vs O(NK) naive approach

## 5. VARIATIONS & EXTENSIONS
- **Different Data Types**: Arrays instead of linked lists
- **Max-Heap**: Merge in descending order
- **Streaming Merge**: Handle infinite or streaming inputs
- **Memory Constraints**: Limited memory for large K

## 6. INTERVIEW INSIGHTS
- Always clarify: "What are constraints on K and N? Memory limits?"
- Edge cases: empty lists, single list, K=1
- Time complexity: O(N log K) where N is total nodes
- Space complexity: O(K) for heap
- Key insight: heap size never exceeds K

## 7. COMMON MISTAKES
- Not handling empty lists correctly
- Wrong heap comparison (using > instead of <)
- Not updating heap when list is exhausted
- Memory leaks in linked list manipulation
- Off-by-one errors in heap operations

## 8. OPTIMIZATION STRATEGIES
- **Min-Heap**: O(N log K) time, O(K) space - optimal
- **Early Termination**: Stop when all lists are exhausted
- **Efficient Comparison**: Compare node values directly
- **Memory Management**: Proper linked list node handling

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like merging K sorted decks of cards:**
- You have K decks of cards, each sorted
- You want to create one sorted deck from all cards
- Always pick the smallest available card from all decks
- Use a min-heap to efficiently find the smallest card
- When you pick a card from a deck, take the next card from that deck

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of K sorted linked lists
2. **Goal**: Merge into one sorted linked list
3. **Constraint**: Input lists are already sorted
4. **Output**: Single sorted linked list containing all nodes

#### Phase 2: Key Insight Recognition
- **"Heap natural fit"** → Need to always find minimum among K elements
- **"K-way merge"** → Classic problem solved with min-heap
- **"Sorted input"** → Each list provides sorted stream
- **"Greedy selection"** → Always pick smallest available element

#### Phase 3: Strategy Development
```
Human thought process:
"I need to merge K sorted lists efficiently.
If I always pick the smallest head element, I'll get sorted output.
A min-heap is perfect for finding the smallest among K elements.
I'll push all list heads into the heap.
Then repeatedly extract min, push next node from same list.
This gives me O(N log K) time instead of O(NK²)."
```

#### Phase 4: Edge Case Handling
- **Empty lists array**: Return nil
- **Some empty lists**: Skip them in initialization
- **Single list**: Return that list directly
- **K=1**: Simple return of the only list

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
lists = [[1,4,5], [1,3,4], [2,6]]

Human thinking:
"I'll push all list heads into min-heap:
Heap contains: 1(from list0), 1(from list1), 2(from list2)

Step 1: Extract 1 (from list0)
- Add to result: [1]
- Push next from list0: 4
- Heap now: 1(list1), 2(list2), 4(list0)

Step 2: Extract 1 (from list1)
- Add to result: [1,1]
- Push next from list1: 3
- Heap now: 2(list2), 3(list1), 4(list0)

Step 3: Extract 2 (from list2)
- Add to result: [1,1,2]
- Push next from list2: 6
- Heap now: 3(list1), 4(list0), 6(list2)

Continue until heap is empty...
Final result: [1,1,2,3,4,4,5,6] ✓"
```

#### Phase 6: Intuition Validation
- **Why heap works**: Always provides smallest available element
- **Why O(N log K)**: Each of N operations costs O(log K)
- **Why O(K) space**: Heap never stores more than K elements
- **Why optimal**: Better than O(NK²) naive approach

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compare all heads?"** → O(K) per element, O(NK²) total
2. **"Should I use priority queue?"** → Yes, that's what a heap is
3. **"What about memory?"** → Use O(K) space, very efficient
4. **"Can I optimize further?"** → Heap approach is already optimal

### Real-World Analogy
**Like merging multiple sorted document streams:**
- You have K document streams, each sorted by timestamp
- You want to create one chronological stream
- Always pick the earliest available document
- Use a priority system to track next available from each stream
- When you take a document, get the next one from that stream

### Human-Readable Pseudocode
```
function mergeKLists(lists):
    if lists is empty:
        return null
    
    minHeap = empty min-heap
    dummy = new ListNode(0)
    current = dummy
    
    // Initialize heap with all list heads
    for list in lists:
        if list is not null:
            heap.push(minHeap, list)
    
    while heap is not empty:
        smallest = heap.extract(minHeap)
        current.next = smallest
        current = current.next
        
        if smallest.next is not null:
            heap.push(minHeap, smallest.next)
    
    return dummy.next
```

### Execution Visualization

### Example: lists = [[1,4,5], [1,3,4], [2,6]]
```
Heap Evolution During Merge:
Initial heap: [1(L0), 1(L1), 2(L2)]

Step 1: Extract 1(L0), push 4(L0)
- Result: [1]
- Heap: [1(L1), 2(L2), 4(L0)]

Step 2: Extract 1(L1), push 3(L1)
- Result: [1,1]
- Heap: [2(L2), 3(L1), 4(L0)]

Step 3: Extract 2(L2), push 6(L2)
- Result: [1,1,2]
- Heap: [3(L1), 4(L0), 6(L2)]

Step 4: Extract 3(L1)
- Result: [1,1,2,3]
- Heap: [4(L0), 6(L2)]

Step 5: Extract 4(L0)
- Result: [1,1,2,3,4]
- Heap: [6(L2)]

Step 6: Extract 6(L2)
- Result: [1,1,2,3,4,6]
- Heap: [4(L0)]

Step 7: Extract 4(L0)
- Result: [1,1,2,3,4,6,4]
- Heap: empty

Final merged list: [1,1,2,3,4,4,5,6] ✓
```

### Key Visualization Points:
- **Heap Operations**: Extract min, push next from same list
- **Sorted Output**: Always maintains sorted order
- **Efficient Selection**: O(log K) per operation
- **Memory Efficiency**: Only O(K) space for heap

### Memory Layout Visualization:
```
Heap State During Merge:
lists = [[1,4,5], [1,3,4], [2,6]]

Heap structure (min-heap):
        1(L1)
      /       \
   2(L2)     4(L0)

After extracting 1(L1):
        2(L2)
      /       \
   3(L1)     4(L0)

Heap always maintains smallest element at root.
```

### Time Complexity Breakdown:
- **Initialization**: O(K) to push all list heads
- **Main Loop**: N extractions and at most N pushes
- **Each Operation**: O(log K) time
- **Total Time**: O(N log K) where N is total nodes
- **Space Complexity**: O(K) for heap

### Alternative Approaches:

#### 1. Naive Comparison (O(NK²) time, O(1) space)
```go
func mergeKListsNaive(lists []*ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    
    for {
        minVal := math.MaxInt32
        minList := -1
        emptyCount := 0
        
        for i, list := range lists {
            if list == nil {
                emptyCount++
                continue
            }
            if list.Val < minVal {
                minVal = list.Val
                minList = i
            }
        }
        
        if emptyCount == len(lists) {
            break
        }
        
        current.Next = &ListNode{Val: minVal}
        current = current.Next
        lists[minList] = lists[minList].Next
    }
    
    return dummy.Next
}
```
- **Pros**: Simple to understand
- **Cons**: O(NK²) time, very inefficient

#### 2. Divide and Conquer (O(N log K) time, O(log K) space)
```go
func mergeKListsDivide(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    
    mid := len(lists) / 2
    left := mergeKListsDivide(lists[:mid])
    right := mergeKListsDivide(lists[mid:])
    
    return mergeTwoLists(left, right)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }
    
    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }
    
    return dummy.Next
}
```
- **Pros**: O(N log K) time, no extra space
- **Cons**: More complex recursion, stack overhead

#### 3. Iterative Two-Way Merge (O(NK) time, O(1) space)
```go
func mergeKListsIterative(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    result := lists[0]
    for i := 1; i < len(lists); i++ {
        result = mergeTwoLists(result, lists[i])
    }
    
    return result
}
```
- **Pros**: Reuses existing merge function
- **Cons**: O(NK²) time, inefficient for large K

### Extensions for Interviews:
- **Different Data Types**: Arrays instead of linked lists
- **Streaming Merge**: Handle infinite or streaming inputs
- **Memory Constraints**: Limited memory for large K
- **Custom Comparators**: Different sorting criteria
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 4, 5}, {1, 3, 4}, {2, 6}},
		{},
		{{}},
		{{1, 2, 3}},
		{{1}, {1}, {1}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{-5, -3, -1}, {-4, -2, 0}, {-6, -4, -2}},
		{{1, 3, 5}, {2, 4, 6}, {7, 8, 9}},
		{{1, 100}, {2, 99}, {3, 98}},
	}
	
	for i, tc := range testCases {
		// Convert to linked lists
		lists := make([]*ListNode, len(tc))
		for j, nums := range tc {
			lists[j] = createLinkedList(nums)
		}
		
		merged := mergeKLists(lists)
		result := linkedListToSlice(merged)
		
		fmt.Printf("Test Case %d: %v -> Merged: %v\n", i+1, tc, result)
	}
}
