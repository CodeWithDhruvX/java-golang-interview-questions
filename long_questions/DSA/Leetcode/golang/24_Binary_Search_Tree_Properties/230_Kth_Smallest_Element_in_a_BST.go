package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230. Kth Smallest Element in a BST
// Time: O(H + K) where H is tree height, Space: O(H)
func kthSmallest(root *TreeNode, k int) int {
	var result int
	count := 0
	
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		
		// Traverse left subtree
		inorder(node.Left)
		
		// Process current node
		count++
		if count == k {
			result = node.Val
			return
		}
		
		// Traverse right subtree
		inorder(node.Right)
	}
	
	inorder(root)
	return result
}

// Iterative approach using stack
func kthSmallestIterative(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	
	stack := []*TreeNode{}
	current := root
	count := 0
	
	for len(stack) > 0 || current != nil {
		// Go as far left as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		count++
		if count == k {
			return current.Val
		}
		
		// Move to right subtree
		current = current.Right
	}
	
	return -1 // k is larger than number of nodes
}

// Optimized approach with early termination
func kthSmallestOptimized(root *TreeNode, k int) int {
	var result int
	var count int
	
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		
		// Check if we can skip entire left subtree
		leftSize := countNodes(node.Left)
		if count+leftSize+1 >= k {
			// The kth element is in this subtree
			if count+leftSize+1 == k {
				result = node.Val
				count = k // Stop further processing
				return
			}
			
			// Search in left subtree
			inorder(node.Left)
			if count >= k {
				return
			}
			
			// Process current node
			count++
			if count == k {
				result = node.Val
				return
			}
			
			// Search in right subtree
			inorder(node.Right)
		} else {
			// Skip entire left subtree and current node
			count += leftSize + 1
			// Search in right subtree
			inorder(node.Right)
		}
	}
	
	inorder(root)
	return result
}

// Helper function to count nodes in subtree
func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

// Approach using reverse inorder for kth largest
func kthLargest(root *TreeNode, k int) int {
	var result int
	count := 0
	
	var reverseInorder func(*TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		
		// Traverse right subtree first
		reverseInorder(node.Right)
		
		// Process current node
		count++
		if count == k {
			result = node.Val
			return
		}
		
		// Traverse left subtree
		reverseInorder(node.Left)
	}
	
	reverseInorder(root)
	return result
}

// Approach that returns both kth smallest and largest
func kthSmallestAndLargest(root *TreeNode, k int) (int, int) {
	var smallResult, largeResult int
	var smallCount, largeCount int
	
	// For kth smallest
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || smallCount >= k {
			return
		}
		
		inorder(node.Left)
		smallCount++
		if smallCount == k {
			smallResult = node.Val
			return
		}
		inorder(node.Right)
	}
	
	// For kth largest
	var reverseInorder func(*TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil || largeCount >= k {
			return
		}
		
		reverseInorder(node.Right)
		largeCount++
		if largeCount == k {
			largeResult = node.Val
			return
		}
		reverseInorder(node.Left)
	}
	
	inorder(root)
	reverseInorder(root)
	
	return smallResult, largeResult
}

// Approach using augmented BST (with node counts)
type AugmentedNode struct {
	Val   int
	Left  *AugmentedNode
	Right *AugmentedNode
	Size  int // Size of subtree including this node
}

func kthSmallestAugmented(root *AugmentedNode, k int) int {
	return kthSmallestAugmentedHelper(root, k)
}

func kthSmallestAugmentedHelper(node *AugmentedNode, k int) int {
	if node == nil {
		return -1
	}
	
	leftSize := 0
	if node.Left != nil {
		leftSize = node.Left.Size
	}
	
	if k <= leftSize {
		// kth element is in left subtree
		return kthSmallestAugmentedHelper(node.Left, k)
	} else if k == leftSize+1 {
		// Current node is the kth element
		return node.Val
	} else {
		// kth element is in right subtree
		return kthSmallestAugmentedHelper(node.Right, k-leftSize-1)
	}
}

// Helper function to build augmented BST
func buildAugmentedBST(nums []int) *AugmentedNode {
	if len(nums) == 0 {
		return nil
	}
	
	root := &AugmentedNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		insertAugmentedBST(root, nums[i])
	}
	updateSize(root)
	return root
}

func insertAugmentedBST(root *AugmentedNode, val int) {
	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &AugmentedNode{Val: val}
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &AugmentedNode{Val: val}
				return
			}
			current = current.Right
		}
	}
}

func updateSize(node *AugmentedNode) {
	if node == nil {
		return
	}
	
	updateSize(node.Left)
	updateSize(node.Right)
	
	leftSize := 0
	rightSize := 0
	if node.Left != nil {
		leftSize = node.Left.Size
	}
	if node.Right != nil {
		rightSize = node.Right.Size
	}
	
	node.Size = 1 + leftSize + rightSize
}

// Helper function to create BST from array
func createBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	root := &TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		insertBST(root, nums[i])
	}
	return root
}

func insertBST(root *TreeNode, val int) {
	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &TreeNode{Val: val}
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{Val: val}
				return
			}
			current = current.Right
		}
	}
}

// Helper function to convert tree to array (in-order traversal)
func treeToArray(root *TreeNode) []int {
	var result []int
	inorderTraversal(root, &result)
	return result
}

func inorderTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	
	inorderTraversal(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversal(node.Right, result)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Kth Element in BST using In-order Traversal
- **In-order Traversal**: BST produces sorted sequence (ascending)
- **Early Termination**: Stop when kth element found
- **Count Tracking**: Keep count of visited elements
- **Multiple Approaches**: Recursive, iterative, optimized, augmented

## 2. PROBLEM CHARACTERISTICS
- **BST Property**: Left subtree < root < right subtree
- **Sorted Order**: In-order traversal gives ascending sequence
- **Kth Element**: Find kth smallest (1-indexed)
- **Efficiency**: Can leverage BST structure for optimization

## 3. SIMILAR PROBLEMS
- Kth Largest Element in BST (LeetCode 230) - Reverse in-order
- Binary Search Tree Iterator (LeetCode 173) - Next/prev operations
- Convert Sorted Array to BST (LeetCode 108) - BST construction
- Find Median in BST - Use kth element logic

## 4. KEY OBSERVATIONS
- **In-order Property**: BST in-order = sorted ascending
- **Early Termination**: No need to traverse entire tree
- **Optimization Opportunity**: Skip entire subtrees using subtree sizes
- **Multiple Solutions**: Different trade-offs between simplicity and efficiency

## 5. VARIATIONS & EXTENSIONS
- **Kth Largest**: Use reverse in-order traversal
- **Multiple Queries**: Handle repeated kth queries efficiently
- **Augmented BST**: Store subtree sizes for O(log N) queries
- **Dynamic Updates**: Support insert/delete operations

## 6. INTERVIEW INSIGHTS
- Always clarify: "1-indexed or 0-indexed? Multiple queries? Tree size?"
- Edge cases: empty tree, k > tree size, k = 1, k = tree size
- Time complexity: O(H + K) for basic, O(log N) for augmented
- Space complexity: O(H) for recursion stack
- Key insight: in-order traversal gives sorted sequence

## 7. COMMON MISTAKES
- Wrong indexing (0 vs 1 indexed)
- Not handling k > tree size
- Not terminating early when kth found
- Off-by-one errors in counting
- Not handling empty tree correctly

## 8. OPTIMIZATION STRATEGIES
- **Basic In-order**: O(H + K) time, O(H) space - simple
- **Iterative In-order**: O(H + K) time, O(H) space - no recursion
- **Optimized Counting**: O(H + K) time, O(H) space - skip subtrees
- **Augmented BST**: O(log N) time, O(N) space - best for multiple queries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the kth person in a sorted line:**
- BST is like a sorted list with hierarchical structure
- In-order traversal visits people in ascending order
- Like counting people in line until you reach the kth person
- Can skip entire groups if you know their sizes
- Like having a directory that tells you group sizes

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: BST root, integer k (1-indexed)
2. **Goal**: Find kth smallest element
3. **Property**: BST in-order = sorted ascending
4. **Output**: Value of kth smallest element

#### Phase 2: Key Insight Recognition
- **"In-order natural fit"** → BST produces sorted sequence
- **"Early termination"** → Stop when kth element found
- **"Counting approach"** → Track visited elements
- **"Optimization opportunity"** → Skip subtrees using sizes

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find kth smallest in BST.
This is perfect for in-order traversal:

Basic Approach:
1. Perform in-order traversal (left, root, right)
2. Count elements as I visit them
3. When count reaches k, return current element
4. Stop early - no need to traverse entire tree

Optimized Approach:
1. At each node, count left subtree size
2. If k <= leftSize: kth is in left subtree
3. If k == leftSize + 1: current node is kth
4. If k > leftSize + 1: search right subtree with adjusted k

Augmented BST Approach:
1. Store subtree sizes in nodes
2. Use sizes to navigate directly to kth element
3. O(log N) time per query

Each approach has different trade-offs!"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return appropriate value (often -1 or error)
- **k > tree size**: Handle gracefully
- **k = 1**: Return minimum element
- **k = tree size**: Return maximum element

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
BST: [5, 3, 7, 2, 4, 6, 8], k = 4

Human thinking:
"Basic In-order Approach:
Traverse in-order: 2, 3, 4, 5, 6, 7, 8
Count: 1, 2, 3, 4 ✓ Found 4th = 4

Optimized Counting Approach:
At root (5):
- Left subtree size = 3 (nodes: 2,3,4)
- k=4 > 3+1, so search right with k=4-3-1=0
- Wait, k=4 > 3+1, search right with k=4-3-1=0
- Actually k=4 > 3+1, search right with k=4-3-1=0
- k=4 > leftSize+1, so kth is in right subtree
- Search right subtree with k=4-leftSize-1=4-3-1=0
- k=0 means we want smallest in right subtree
- Right subtree: [6,8], smallest is 6

Wait, let me recalculate:
k=4, leftSize=3
k > leftSize+1 (4 > 4)? No, k == leftSize+1
So current node (5) is the 4th element ✓

Augmented BST Approach:
At root (5), leftSize=3:
k=4 == leftSize+1, so root is answer ✓"
```

#### Phase 6: Intuition Validation
- **Why in-order works**: BST property guarantees sorted order
- **Why early termination works**: No need to visit all nodes
- **Why optimization works**: Can skip entire subtrees
- **Why augmented works**: Subtree sizes enable direct navigation

### Common Human Pitfalls & How to Avoid Them
1. **"Why not convert to array?"** → Inefficient for large trees
2. **"Should I use BFS?"** → BFS doesn't give sorted order
3. **"What about multiple queries?"** → Consider augmented BST
4. **"Can I optimize further?"** → O(log N) with augmentation
5. **"What about kth largest?"** → Use reverse in-order

### Real-World Analogy
**Like finding the kth ranked student in a school:**
- Students are organized in a hierarchical class structure
- Each class has students sorted by grade
- You need to find the kth highest scoring student
- Can skip entire classes if you know their size and ranking ranges
- Like having a class directory with student counts

### Human-Readable Pseudocode
```
function kthSmallest(root, k):
    count = 0
    result = null
    
    function inorder(node):
        if node is null or count >= k:
            return
        
        inorder(node.left)
        
        count++
        if count == k:
            result = node.val
            return
        
        inorder(node.right)
    
    inorder(root)
    return result

// Optimized version
function kthSmallestOptimized(root, k):
    return kthSmallestHelper(root, k)

function kthSmallestHelper(node, k):
    if node is null:
        return null
    
    leftSize = countNodes(node.left)
    
    if k <= leftSize:
        return kthSmallestHelper(node.left, k)
    else if k == leftSize + 1:
        return node.val
    else:
        return kthSmallestHelper(node.right, k - leftSize - 1)
```

### Execution Visualization

### Example: BST = [5, 3, 7, 2, 4, 6, 8], k = 4
```
Tree Structure:
        5
       /   \
      3     7
     / \   / \
    2   4 6   8

Basic In-order Traversal:
Visit order: 2, 3, 4, 5, 6, 7, 8
Count: 1, 2, 3, 4 ✓ Found 4th = 5

Optimized Counting:
At root (5):
- Left subtree size = 3 (nodes: 2,3,4)
- k=4 == leftSize+1 (4 == 3+1) ✓
- Current node (5) is the 4th element

Augmented BST:
Node sizes: 5(7), 3(3), 7(3), 2(1), 4(1), 6(1), 8(1)
At root (5), leftSize=3:
k=4 == leftSize+1, so root is answer ✓
```

### Key Visualization Points:
- **In-order Property**: BST produces sorted sequence
- **Counting Logic**: Track visited elements until kth reached
- **Optimization**: Use subtree sizes to skip entire branches
- **Early Termination**: Stop when answer found

### Memory Layout Visualization:
```
Recursive Stack During Basic In-order:
Call Stack:          Count:    Current Node:
inorder(5)           0          5
├── inorder(3)        0          3
│   ├── inorder(2)    0          2
│   │   ├── inorder(null) ✓
│   │   └── inorder(null) ✓
│   ├── Process 2     1          2 (count=1)
│   ├── inorder(4)    1          4
│   │   ├── inorder(null) ✓
│   │   └── inorder(null) ✓
│   ├── Process 4     2          4 (count=2)
│   └── inorder(null) ✓
├── Process 5         3          5 (count=3)
└── inorder(7)        3          7
    ├── inorder(6)    3          6
    │   ├── inorder(null) ✓
    │   └── inorder(null) ✓
    ├── Process 6     4          6 (count=4) ✓ FOUND!
    └── return (early termination)
```

### Time Complexity Breakdown:
- **Basic In-order**: O(H + K) time, O(H) space
- **Optimized Counting**: O(H + K) time, O(H) space (better average case)
- **Augmented BST**: O(log N) time, O(N) space (best for multiple queries)
- **Iterative**: O(H + K) time, O(H) space (no recursion)

### Alternative Approaches:

#### 1. Iterative In-order (O(H + K) time, O(H) space)
```go
func kthSmallestIterative(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    current := root
    count := 0
    
    for len(stack) > 0 || current != nil {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        count++
        if count == k {
            return current.Val
        }
        
        current = current.Right
    }
    
    return -1
}
```
- **Pros**: No recursion, memory efficient
- **Cons**: More complex implementation

#### 2. Augmented BST (O(log N) time, O(N) space)
```go
type AugmentedNode struct {
    Val   int
    Left  *AugmentedNode
    Right *AugmentedNode
    Size  int // Size of subtree
}

func kthSmallestAugmented(root *AugmentedNode, k int) int {
    return kthSmallestAugmentedHelper(root, k)
}

func kthSmallestAugmentedHelper(node *AugmentedNode, k int) int {
    if node == nil {
        return -1
    }
    
    leftSize := 0
    if node.Left != nil {
        leftSize = node.Left.Size
    }
    
    if k <= leftSize {
        return kthSmallestAugmentedHelper(node.Left, k)
    } else if k == leftSize + 1 {
        return node.Val
    } else {
        return kthSmallestAugmentedHelper(node.Right, k-leftSize-1)
    }
}
```
- **Pros**: O(log N) time, excellent for multiple queries
- **Cons**: Requires tree augmentation, extra space

#### 3. Convert to Array (O(N) time, O(N) space)
```go
func kthSmallestArray(root *TreeNode, k int) int {
    arr := treeToArray(root)
    if k > 0 && k <= len(arr) {
        return arr[k-1]
    }
    return -1
}
```
- **Pros**: Simple implementation
- **Cons**: O(N) space, inefficient for single query

### Extensions for Interviews:
- **Kth Largest**: Use reverse in-order traversal
- **Multiple Queries**: Consider augmented BST approach
- **Dynamic Updates**: Support insert/delete with size updates
- **Range Queries**: Find elements in value ranges
- **Performance Analysis**: Discuss trade-offs between approaches
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Kth Smallest Element in BST ===")
	
	testCases := []struct {
		initialTree []interface{}
		k          int
		description string
	}{
		{[]interface{}{3, 1, 4, nil, 2}, 1, "Find smallest element"},
		{[]interface{}{3, 1, 4, nil, 2}, 2, "Find second smallest"},
		{[]interface{}{3, 1, 4, nil, 2}, 3, "Find third smallest"},
		{[]interface{}{5, 3, 6, 2, 4, nil, nil, 1}, 4, "Find 4th smallest"},
		{[]interface{}{5, 3, 6, 2, 4, nil, nil, 1}, 7, "Find largest element"},
		{[]interface{}{}, 1, "Empty tree"},
		{[]interface{}{10}, 1, "Single node"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 5, "Medium tree"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 3, "Find third in medium tree"},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.initialTree)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Tree: %v\n", tc.initialTree)
		fmt.Printf("  k = %d\n", tc.k)
		
		result1 := kthSmallest(root, tc.k)
		result2 := kthSmallestIterative(root, tc.k)
		result3 := kthSmallestOptimized(root, tc.k)
		
		fmt.Printf("  Recursive: %d\n", result1)
		fmt.Printf("  Iterative: %d\n", result2)
		fmt.Printf("  Optimized: %d\n\n", result3)
	}
	
	// Test kth largest
	fmt.Println("=== Testing Kth Largest Element ===")
	largeRoot := createBST([]int{10, 5, 15, 2, 7, 12, 18, 1, 3, 6, 8, 11, 13, 16, 20})
	fmt.Printf("BST: %v\n", treeToArray(largeRoot))
	
	for k := 1; k <= 5; k++ {
		largest := kthLargest(largeRoot, k)
		fmt.Printf("kth largest (k=%d): %d\n", k, largest)
	}
	
	// Test both kth smallest and largest
	fmt.Println("\n=== Testing Both Kth Smallest and Largest ===")
	for k := 1; k <= 3; k++ {
		small, large := kthSmallestAndLargest(largeRoot, k)
		fmt.Printf("k=%d: smallest=%d, largest=%d\n", k, small, large)
	}
	
	// Test augmented BST
	fmt.Println("\n=== Testing Augmented BST ===")
	augmentedRoot := buildAugmentedBST([]int{10, 5, 15, 2, 7, 12, 18, 1, 3, 6, 8, 11, 13, 16, 20})
	
	for k := 1; k <= 5; k++ {
		result := kthSmallestAugmented(augmentedRoot, k)
		fmt.Printf("Augmented BST kth smallest (k=%d): %d\n", k, result)
	}
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	// Very large k
	singleRoot := &TreeNode{Val: 10}
	fmt.Printf("Single node, k=5: %d\n", kthSmallest(singleRoot, 5))
	
	// Large tree
	largeBST := createBST([]int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85})
	fmt.Printf("Large tree, k=8: %d\n", kthSmallest(largeBST, 8))
	
	// Test with negative values
	negRoot := createBST([]int{0, -10, 10, -20, -5, 5, 15})
	fmt.Printf("Tree with negatives: %v\n", treeToArray(negRoot))
	fmt.Printf("kth smallest (k=3): %d\n", kthSmallest(negRoot, 3))
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	perfRoot := createBST([]int{100, 50, 150, 25, 75, 125, 175, 12, 37, 62, 87, 112, 137, 162, 187})
	
	start := 0
	for k := 1; k <= 20; k++ {
		start += kthSmallest(perfRoot, k)
	}
	
	fmt.Printf("Performance test completed with 20 queries\n")
}
