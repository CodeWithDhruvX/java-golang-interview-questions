package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 876. Middle of the Linked List
// Time: O(N), Space: O(1) - Tortoise and Hare Algorithm
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	return slow
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

## 1. ALGORITHM PATTERN: Tortoise and Hare (Fast and Slow Pointers)
- **Two Pointers**: Slow pointer moves 1 step, fast pointer moves 2 steps
- **Middle Detection**: When fast reaches end, slow is at middle
- **Space Efficiency**: O(1) space, no extra storage needed
- **Linear Time**: O(N) time complexity

## 2. PROBLEM CHARACTERISTICS
- **Middle Finding**: Locate middle node(s) of linked list
- **Even Length Handling**: Return second middle for even length
- **Pointer Manipulation**: Navigate through linked list structure
- **Single Pass**: Find middle in one traversal

## 3. SIMILAR PROBLEMS
- Find Start of Cycle (LeetCode 142) - Similar fast/slow technique
- Linked List Cycle (LeetCode 141) - Same pointer technique
- Palindrome Linked List (LeetCode 234) - Use middle for palindrome check
- Remove Nth Node From End (LeetCode 19) - Two-pointer technique

## 4. KEY OBSERVATIONS
- **Speed Ratio**: Fast moves twice as fast as slow
- **Middle Guarantee**: When fast reaches end, slow is at middle
- **Even Length**: Fast reaches nil, slow at second middle
- **Odd Length**: Fast reaches last node, slow at exact middle

## 5. VARIATIONS & EXTENSIONS
- **First Middle**: Return first middle for even length
- **Both Middles**: Return both middle nodes for even length
- **Nth from End**: Find Nth node from end using similar technique
- **Multiple Lists**: Find middle of multiple linked lists

## 6. INTERVIEW INSIGHTS
- Always clarify: "For even length, which middle should I return?"
- Edge cases: empty list, single node, even/odd length
- Time complexity: O(N) time, O(1) space
- Alternative approaches: Two-pass with length counting

## 7. COMMON MISTAKES
- Not handling empty list case
- Starting fast pointer incorrectly
- Not checking both fast and fast.Next for nil
- Returning wrong middle for even length
- Using array conversion (inefficient)

## 8. OPTIMIZATION STRATEGIES
- **Fast/Slow Pointers**: Optimal O(N) time, O(1) space
- **Two-Pass Approach**: O(N) time, O(1) space (but two traversals)
- **Array Conversion**: O(N) time, O(N) space (inefficient)
- **Length Counting**: O(N) time, O(1) space (two passes)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like two runners on a track:**
- You have two runners starting at the same point
- One runner (slow) runs at normal speed
- The other runner (fast) runs at twice the speed
- When the fast runner reaches the finish line, the slow runner will be exactly halfway
- This works because the fast runner covers twice the distance in the same time

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Head of a singly-linked list
2. **Goal**: Find middle node of the list
3. **Constraint**: For even length, return second middle
4. **Output**: Middle node pointer

#### Phase 2: Key Insight Recognition
- **"Two-pointer natural fit"** → Fast and slow pointers
- **"Speed ratio"** → Fast moves 2x speed of slow
- **"Middle guarantee"** → Fast at end means slow at middle
- **"Single pass"** → Find middle in one traversal

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the middle of the linked list efficiently.
I'll use two pointers: slow moves one step, fast moves two steps.
When fast reaches the end, slow will be at the middle.
For even length, fast will reach nil, slow at second middle.
For odd length, fast will reach last node, slow at exact middle."
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return nil (no middle)
- **Single node**: Return the only node (it's the middle)
- **Two nodes**: Return second node (as per problem requirement)
- **Large list**: General algorithm handles all cases

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
List: 1 → 2 → 3 → 4 → 5 (odd length)

Human thinking:
"I'll start with both pointers at head:
slow = 1, fast = 1

Step 1: slow moves to 2, fast moves to 3
slow = 2, fast = 3

Step 2: slow moves to 3, fast moves to 5
slow = 3, fast = 5

Step 3: slow moves to 4, fast moves to nil (past end)
slow = 4, fast = nil

Fast is nil, so stop and return slow = 4
Wait, that's wrong! Let me recheck...

Actually, let me be more careful:
Initial: slow=1, fast=1

Step 1: Check fast != nil and fast.Next != nil (1 != nil and 2 != nil) ✓
slow = slow.Next = 2
fast = fast.Next.Next = 3

Step 2: Check fast != nil and fast.Next != nil (3 != nil and 4 != nil) ✓
slow = slow.Next = 3
fast = fast.Next.Next = 5

Step 3: Check fast != nil and fast.Next != nil (5 != nil but 5.Next = nil) ✗
Stop here, return slow = 3

Perfect! Node 3 is the middle ✓"

List: 1 → 2 → 3 → 4 (even length)

Human thinking:
"Initial: slow=1, fast=1

Step 1: Check fast != nil and fast.Next != nil (1 != nil and 2 != nil) ✓
slow = slow.Next = 2
fast = fast.Next.Next = 3

Step 2: Check fast != nil and fast.Next != nil (3 != nil and 4 != nil) ✓
slow = slow.Next = 3
fast = fast.Next.Next = nil

Step 3: Check fast != nil and fast.Next != nil (fast = nil) ✗
Stop here, return slow = 3

Perfect! Node 3 is the second middle ✓"
```

#### Phase 6: Intuition Validation
- **Why two pointers work**: Fast covers twice the distance
- **Why O(1) space**: Only two pointers needed
- **Why O(N) time**: Each pointer traverses at most N nodes
- **Why second middle**: Fast moves 2 steps, slow moves 1 step

### Common Human Pitfalls & How to Avoid Them
1. **"Why not count length first?"** → Works but requires two passes
2. **"Should I start fast at head.Next?"** → No, both start at head
3. **"What about pointer speed?"** → Fast should move 2x speed
4. **"Can I optimize further?"** → Fast/slow is already optimal

### Real-World Analogy
**Like finding the middle of a line of people:**
- You have a line of people and want to find the middle person
- You could count everyone first, then go to the middle (two passes)
- Or you could have two people walk from the start, one twice as fast
- When the fast person reaches the end, the slow person will be at the middle
- This is much more efficient as you only walk through once

### Human-Readable Pseudocode
```
function middleNode(head):
    if head == null:
        return null
    
    slow = head
    fast = head
    
    while fast != null and fast.next != null:
        slow = slow.next
        fast = fast.next.next
    
    return slow
```

### Execution Visualization

### Example 1: 1 → 2 → 3 → 4 → 5 (odd length)
```
Pointer Movement:
Initial: slow=1, fast=1

Step 1: slow=2, fast=3
Step 2: slow=3, fast=5
Step 3: fast.next = nil, stop

Final: slow=3 (middle node)
```

### Example 2: 1 → 2 → 3 → 4 (even length)
```
Pointer Movement:
Initial: slow=1, fast=1

Step 1: slow=2, fast=3
Step 2: slow=3, fast=nil
Stop here

Final: slow=3 (second middle node)
```

### Key Visualization Points:
- **Speed difference**: Fast moves twice as fast as slow
- **Middle guarantee**: Fast at end means slow at middle
- **Even length**: Fast reaches nil, slow at second middle
- **Odd length**: Fast reaches last node, slow at exact middle

### Memory Layout Visualization:
```
Odd length: 1 → 2 → 3 → 4 → 5
           slow    fast
           (step 2)

Even length: 1 → 2 → 3 → 4
             slow    fast
             (step 2)

The fast pointer "outruns" the slow pointer,
leaving slow exactly at the middle when fast reaches the end.
```

### Time Complexity Breakdown:
- **Fast/Slow Pointers**: O(N) time, O(1) space
- **Two-Pass Approach**: O(N) time, O(1) space (but two traversals)
- **Array Conversion**: O(N) time, O(N) space (inefficient)
- **Recursive Approach**: O(N) time, O(N) space (stack)

### Alternative Approaches:

#### 1. Two-Pass Approach (O(N) time, O(1) space)
```go
func middleNodeTwoPass(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    // First pass: count length
    length := 0
    current := head
    for current != nil {
        length++
        current = current.Next
    }
    
    // Second pass: find middle
    middleIndex := length / 2
    current = head
    for i := 0; i < middleIndex; i++ {
        current = current.Next
    }
    
    return current
}
```
- **Pros**: Simple to understand
- **Cons**: Two traversals, less efficient

#### 2. Array Conversion (O(N) time, O(N) space)
```go
func middleNodeArray(head *ListNode) *ListNode {
    nodes := []*ListNode{}
    current := head
    
    for current != nil {
        nodes = append(nodes, current)
        current = current.Next
    }
    
    if len(nodes) == 0 {
        return nil
    }
    
    return nodes[len(nodes)/2]
}
```
- **Pros**: Very simple concept
- **Cons**: O(N) space, inefficient

#### 3. Recursive Approach (O(N) time, O(N) space)
```go
func middleNodeRecursive(head *ListNode) *ListNode {
    _, middle := middleNodeHelper(head, 0)
    return middle
}

func middleNodeHelper(node *ListNode, index int) (int, *ListNode) {
    if node == nil {
        return index - 1, nil
    }
    
    length, middle := middleNodeHelper(node.Next, index+1)
    
    if length/2 == index {
        return length, node
    }
    
    return length, middle
}
```
- **Pros**: Elegant recursive solution
- **Cons**: O(N) space due to recursion stack

### Extensions for Interviews:
- **First Middle**: Return first middle for even length
- **Both Middles**: Return both middle nodes for even length
- **Nth from End**: Find Nth node from end using similar technique
- **Middle Index**: Return index of middle node instead of node
- **Multiple Lists**: Find middle of multiple linked lists
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4, 5},     // Odd length
		{1, 2, 3, 4, 5, 6},  // Even length
		{1},                  // Single node
		{},                   // Empty list
		{1, 2},               // Two nodes
		{1, 2, 3},            // Three nodes
		{1, 2, 3, 4},         // Four nodes
		{10, 20, 30, 40, 50, 60, 70}, // Seven nodes
	}
	
	for i, nums := range testCases {
		head := createLinkedList(nums)
		middle := middleNode(head)
		result := linkedListToSlice(middle)
		
		fmt.Printf("Test Case %d: %v -> Middle node(s): %v\n", i+1, nums, result)
	}
}
