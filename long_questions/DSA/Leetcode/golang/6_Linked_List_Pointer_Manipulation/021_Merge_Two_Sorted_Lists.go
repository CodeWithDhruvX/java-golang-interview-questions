package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 21. Merge Two Sorted Lists
// Time: O(N+M), Space: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	
	// Attach the remaining elements
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	
	return dummy.Next
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

## 1. ALGORITHM PATTERN: Two-Pointer Merge for Linked Lists
- **Dummy Node Technique**: Use dummy node to simplify head handling
- **Two-Pointer Traversal**: Maintain pointers for both lists
- **Comparative Merging**: Compare values and link smaller node
- **Remainder Attachment**: Attach remaining nodes after main loop

## 2. PROBLEM CHARACTERISTICS
- **Sorted Lists**: Both input lists are sorted in ascending order
- **Linked Structure**: Nodes connected via Next pointers
- **In-Place Merging**: Create new merged list without extra storage
- **Linear Traversal**: Single pass through both lists

## 3. SIMILAR PROBLEMS
- Merge K Sorted Lists (LeetCode 23) - Multiple list merging
- Merge Sorted Array (LeetCode 88) - Array version
- Sort List (LeetCode 148) - Sort using merge technique
- Intersection of Two Linked Lists (LeetCode 160) - List intersection

## 4. KEY OBSERVATIONS
- **Sorted property**: Enables linear merging algorithm
- **Pointer manipulation**: Only Next pointers need modification
- **Dummy node**: Eliminates special case for head initialization
- **Remainder handling**: One list may finish before the other

## 5. VARIATIONS & EXTENSIONS
- **Multiple lists**: Merge K sorted lists using heap
- **Descending order**: Merge in reverse order
- **Custom comparison**: Merge with custom comparator
- **In-place modification**: Modify one of the input lists

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can lists be empty? What about duplicate values?"
- Edge cases: empty lists, single nodes, all equal values
- Time complexity: O(N+M) time, O(1) space (excluding output)
- Pointer manipulation: Be careful with Next pointer assignments

## 7. COMMON MISTAKES
- Not handling empty list cases properly
- Losing reference to remaining nodes
- Creating cycles in the merged list
- Not using dummy node (more complex code)
- Incorrect pointer assignment order

## 8. OPTIMIZATION STRATEGIES
- **Dummy node**: Simplifies head handling
- **Single pass**: Process each node exactly once
- **In-place merging**: No extra storage needed
- **Early termination**: Stop when one list is exhausted

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like merging two sorted decks of cards:**
- You have two decks of cards, each sorted by number
- You want to create a single sorted deck by merging them
- You look at the top card of each deck and take the smaller one
- You repeat this process until one deck is empty
- Then you simply add all remaining cards from the other deck

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two sorted singly-linked lists
2. **Goal**: Merge into one sorted singly-linked list
3. **Constraint**: Must maintain sorted order
4. **Output**: Head of merged list

#### Phase 2: Key Insight Recognition
- **"Two-pointer natural fit"** → Maintain pointers for both lists
- **"Comparative merging"** → Always take smaller value
- **"Dummy node technique"** → Simplify head handling
- **"Remainder attachment"** → Handle remaining nodes

#### Phase 3: Strategy Development
```
Human thought process:
"I need to merge two sorted linked lists.
I'll use a dummy node to make the head handling easier.
I'll maintain two pointers, one for each list.
At each step, I'll compare the current nodes and link the smaller one.
I'll advance the pointer of the list whose node I took.
When one list is empty, I'll attach the remaining nodes from the other list."
```

#### Phase 4: Edge Case Handling
- **Both lists empty**: Return nil
- **One list empty**: Return the other list
- **Single node lists**: Handle correctly with general algorithm
- **Equal values**: Take from either list (consistently)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
list1: 1 → 2 → 4, list2: 1 → 3 → 4

Human thinking:
"I'll start with a dummy node and current pointer:

Initialize: dummy → current, list1 → 1→2→4, list2 → 1→3→4

Step 1: Compare 1 and 1 (list1.Val <= list2.Val)
Take list1's 1: dummy → 1, current → 1
Advance list1: list1 → 2→4, list2 → 1→3→4

Step 2: Compare 2 and 1
Take list2's 1: dummy → 1 → 1, current → 1
Advance list2: list1 → 2→4, list2 → 3→4

Step 3: Compare 2 and 3
Take list1's 2: dummy → 1 → 1 → 2, current → 2
Advance list1: list1 → 4, list2 → 3→4

Step 4: Compare 4 and 3
Take list2's 3: dummy → 1 → 1 → 2 → 3, current → 3
Advance list2: list1 → 4, list2 → 4

Step 5: Compare 4 and 4
Take list1's 4: dummy → 1 → 1 → 2 → 3 → 4, current → 4
Advance list1: list1 → nil, list2 → 4

Step 6: list1 is empty, attach remaining list2
dummy → 1 → 1 → 2 → 3 → 4 → 4

Final merged list: 1 → 1 → 2 → 3 → 4 → 4 ✓"
```

#### Phase 6: Intuition Validation
- **Why two-pointer works**: Maintains position in both lists
- **Why dummy node works**: Eliminates special head handling
- **Why comparative merging works**: Maintains sorted order
- **Why O(N+M) time**: Each node processed exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Can cause stack overflow for large lists
2. **"Should I create new nodes?"** → No, reuse existing nodes
3. **"What about dummy node?"** → Yes, simplifies code significantly
4. **"Can I optimize further?"** → O(N+M) time is already optimal

### Real-World Analogy
**Like merging two sorted queues of people:**
- You have two queues of people, each sorted by height
- You want to form a single queue sorted by height
- You look at the person at the front of each queue
- You always take the shorter person and add them to the merged queue
- When one queue is empty, you add all remaining people from the other queue

### Human-Readable Pseudocode
```
function mergeTwoLists(list1, list2):
    dummy = new ListNode()
    current = dummy
    
    while list1 != null and list2 != null:
        if list1.val <= list2.val:
            current.next = list1
            list1 = list1.next
        else:
            current.next = list2
            list2 = list2.next
        current = current.next
    
    // Attach remaining nodes
    if list1 != null:
        current.next = list1
    else:
        current.next = list2
    
    return dummy.next
```

### Execution Visualization

### Example: list1 = [1,2,4], list2 = [1,3,4]
```
Initial: dummy → current, list1 → 1→2→4, list2 → 1→3→4

Step 1: 1 <= 1, take list1's 1
dummy → 1, current → 1
list1 → 2→4, list2 → 1→3→4

Step 2: 2 > 1, take list2's 1
dummy → 1 → 1, current → 1
list1 → 2→4, list2 → 3→4

Step 3: 2 <= 3, take list1's 2
dummy → 1 → 1 → 2, current → 2
list1 → 4, list2 → 3→4

Step 4: 4 > 3, take list2's 3
dummy → 1 → 1 → 2 → 3, current → 3
list1 → 4, list2 → 4

Step 5: 4 <= 4, take list1's 4
dummy → 1 → 1 → 2 → 3 → 4, current → 4
list1 → nil, list2 → 4

Step 6: list1 is empty, attach list2
dummy → 1 → 1 → 2 → 3 → 4 → 4

Final merged list: [1,1,2,3,4,4]
```

### Key Visualization Points:
- **Dummy node**: Simplifies head handling
- **Two-pointer traversal**: Maintains positions in both lists
- **Comparative linking**: Always links smaller value
- **Remainder attachment**: Handles remaining nodes efficiently

### Memory Layout Visualization:
```
Before merge:
list1: 1 → 2 → 4 → nil
list2: 1 → 3 → 4 → nil

After step 3 (taking 2 from list1):
dummy → 1 → 1 → 2 → current
                    ↑
                    current pointer
list1: 4 → nil
list2: 3 → 4 → nil

Pointer relationships:
current.Next will point to the smaller of list1.Val and list2.Val
```

### Time Complexity Breakdown:
- **Iterative approach**: O(N+M) time, O(1) space
- **Recursive approach**: O(N+M) time, O(N+M) space (stack)
- **Array conversion**: O(N+M) time, O(N+M) space (inefficient)
- **In-place modification**: O(N+M) time, O(1) space

### Alternative Approaches:

#### 1. Recursive Approach (O(N+M) time, O(N+M) space)
```go
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoListsRecursive(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoListsRecursive(list1, list2.Next)
        return list2
    }
}
```
- **Pros**: Elegant, concise code
- **Cons**: Stack overflow for large lists, O(N+M) space

#### 2. Array Conversion (O(N+M) time, O(N+M) space)
```go
func mergeTwoListsArray(list1 *ListNode, list2 *ListNode) *ListNode {
    // Convert to arrays
    arr1 := linkedListToSlice(list1)
    arr2 := linkedListToSlice(list2)
    
    // Merge arrays
    merged := make([]int, 0, len(arr1)+len(arr2))
    i, j := 0, 0
    
    for i < len(arr1) && j < len(arr2) {
        if arr1[i] <= arr2[j] {
            merged = append(merged, arr1[i])
            i++
        } else {
            merged = append(merged, arr2[j])
            j++
        }
    }
    
    merged = append(merged, arr1[i:]...)
    merged = append(merged, arr2[j:]...)
    
    return createLinkedList(merged)
}
```
- **Pros**: Simple to understand
- **Cons**: O(N+M) extra space, inefficient

#### 3. In-Place Modification (O(N+M) time, O(1) space)
```go
func mergeTwoListsInPlace(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    
    // Ensure list1 has smaller head
    if list1.Val > list2.Val {
        list1, list2 = list2, list1
    }
    
    head := list1
    
    for list1.Next != nil && list2 != nil {
        if list1.Next.Val > list2.Val {
            temp := list2.Next
            list2.Next = list1.Next
            list1.Next = list2
            list2 = temp
        }
        list1 = list1.Next
    }
    
    if list2 != nil {
        list1.Next = list2
    }
    
    return head
}
```
- **Pros**: O(1) space, modifies list1 in place
- **Cons**: More complex logic, modifies input

### Extensions for Interviews:
- **Merge K Sorted Lists**: Use min-heap for multiple lists
- **Descending Order**: Merge in reverse order
- **Custom Comparator**: Merge with custom comparison function
- **Circular Lists**: Handle circular linked list inputs
- **Duplicate Handling**: Remove duplicates during merge
*/
func main() {
	// Test cases
	testCases := []struct {
		list1 []int
		list2 []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}},
		{[]int{}, []int{}},
		{[]int{}, []int{0}},
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{1, 3, 5}, []int{2, 4, 6}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{-3, -1, 1}, []int{-2, 0, 2}},
	}
	
	for i, tc := range testCases {
		list1 := createLinkedList(tc.list1)
		list2 := createLinkedList(tc.list2)
		merged := mergeTwoLists(list1, list2)
		result := linkedListToSlice(merged)
		
		fmt.Printf("Test Case %d: list1=%v, list2=%v -> Merged: %v\n", 
			i+1, tc.list1, tc.list2, result)
	}
}
