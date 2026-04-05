import java.util.*;

public class BinaryTreeInorderTraversal {
    
    // 94. Binary Tree Inorder Traversal - Morris Traversal
    // Time: O(N), Space: O(1)
    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        
        TreeNode(int val) {
            this.val = val;
        }
    }
    
    // Morris Traversal for inorder
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        
        while (current != null) {
            if (current.left == null) {
                // No left child, visit current and move to right
                result.add(current.val);
                current = current.right;
            } else {
                // Find predecessor
                TreeNode predecessor = current.left;
                while (predecessor.right != null && predecessor.right != current) {
                    predecessor = predecessor.right;
                }
                
                if (predecessor.right == null) {
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Revert the changes and visit current
                    predecessor.right = null;
                    result.add(current.val);
                    current = current.right;
                }
            }
        }
        
        return result;
    }
    
    // Morris Traversal for preorder
    public List<Integer> preorderTraversalMorris(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        
        while (current != null) {
            if (current.left == null) {
                // No left child, visit current and move to right
                result.add(current.val);
                current = current.right;
            } else {
                // Find predecessor
                TreeNode predecessor = current.left;
                while (predecessor.right != null && predecessor.right != current) {
                    predecessor = predecessor.right;
                }
                
                if (predecessor.right == null) {
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Revert the changes and visit current
                    predecessor.right = null;
                    result.add(current.val); // Visit before going left
                    current = current.right;
                }
            }
        }
        
        return result;
    }
    
    // Morris Traversal for postorder
    public List<Integer> postorderTraversalMorris(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        TreeNode dummy = new TreeNode(0);
        dummy.left = root;
        current = dummy;
        
        while (current != null) {
            if (current.left == null) {
                // No left child, move to right
                current = current.right;
            } else {
                // Find predecessor
                TreeNode predecessor = current.left;
                while (predecessor.right != null && predecessor.right != current) {
                    predecessor = predecessor.right;
                }
                
                if (predecessor.right == null) {
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Process current and revert changes
                    TreeNode temp = predecessor.right;
                    predecessor.right = null;
                    current = temp;
                }
            }
        }
        
        // Reverse the result for postorder
        Collections.reverse(result);
        return result;
    }
    
    // Standard recursive inorder for comparison
    public List<Integer> inorderTraversalRecursive(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        inorderHelper(root, result);
        return result;
    }
    
    private void inorderHelper(TreeNode node, List<Integer> result) {
        if (node == null) {
            return;
        }
        
        inorderHelper(node.left, result);
        result.add(node.val);
        inorderHelper(node.right, result);
    }
    
    // Iterative inorder using stack for comparison
    public List<Integer> inorderTraversalIterative(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        Deque<TreeNode> stack = new ArrayDeque<>();
        TreeNode current = root;
        
        while (current != null || !stack.isEmpty()) {
            while (current != null) {
                stack.push(current);
                current = current.left;
            }
            
            current = stack.pop();
            result.add(current.val);
            current = current.right;
        }
        
        return result;
    }
    
    // Version with detailed explanation
    public class MorrisTraversalResult {
        List<Integer> result;
        List<String> explanation;
        
        MorrisTraversalResult(List<Integer> result, List<String> explanation) {
            this.result = result;
            this.explanation = explanation;
        }
    }
    
    public MorrisTraversalResult inorderTraversalDetailed(TreeNode root) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== Morris Traversal for Inorder ===");
        explanation.add("Using O(1) space algorithm without recursion or stack");
        
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        int step = 1;
        
        while (current != null) {
            explanation.add(String.format("Step %d: Current node = %d", step++, current.val));
            
            if (current.left == null) {
                explanation.add("  No left child - visiting current and moving right");
                result.add(current.val);
                explanation.add(String.format("  Added %d to result", current.val));
                current = current.right;
                explanation.add(String.format("  Moved to right child: %s", 
                    current.right != null ? String.valueOf(current.right.val) : "null"));
            } else {
                explanation.add("  Finding predecessor in left subtree");
                
                // Find predecessor
                TreeNode predecessor = current.left;
                int predecessorSteps = 0;
                
                while (predecessor.right != null && predecessor.right != current) {
                    predecessorSteps++;
                    predecessor = predecessor.right;
                    explanation.add(String.format("    Predecessor step %d: moved to %d", 
                        predecessorSteps, predecessor.val));
                }
                
                if (predecessor.right == null) {
                    explanation.add("  Predecessor.right is null - creating temporary link");
                    explanation.add(String.format("  Making %d the right child of %d", 
                        current.val, predecessor.val));
                    
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                    explanation.add(String.format("  Moved to left child: %d", current.val));
                } else {
                    explanation.add("  Predecessor.right is not null - visiting current and moving right");
                    explanation.add(String.format("  Reverting temporary link, visiting %d", current.val));
                    
                    // Revert the changes and visit current
                    predecessor.right = null;
                    result.add(current.val);
                    explanation.add(String.format("  Added %d to result", current.val));
                    current = current.right;
                    explanation.add(String.format("  Moved to right child: %s", 
                        current.right != null ? String.valueOf(current.right.val) : "null"));
                }
            }
        }
        
        return new MorrisTraversalResult(result, explanation);
    }
    
    // Performance comparison
    public void comparePerformance(TreeNode root, int trials) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Trials: " + trials);
        
        // Morris Traversal
        long startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            inorderTraversal(root);
        }
        long endTime = System.nanoTime();
        System.out.printf("Morris Traversal: took %d ns\n", endTime - startTime);
        
        // Recursive Traversal
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            inorderTraversalRecursive(root);
        }
        endTime = System.nanoTime();
        System.out.printf("Recursive Traversal: took %d ns\n", endTime - startTime);
        
        // Iterative Traversal
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            inorderTraversalIterative(root);
        }
        endTime = System.nanoTime();
        System.out.printf("Iterative Traversal: took %d ns\n", endTime - startTime);
    }
    
    // Memory usage analysis
    public void memoryUsageAnalysis(TreeNode root) {
        System.out.println("=== Memory Usage Analysis ===");
        
        // Morris Traversal - O(1) space
        long morrisStartTime = System.nanoTime();
        List<Integer> morrisResult = inorderTraversal(root);
        long morrisEndTime = System.nanoTime();
        
        // Recursive Traversal - O(H) space (recursion stack)
        long recursiveStartTime = System.nanoTime();
        List<Integer> recursiveResult = inorderTraversalRecursive(root);
        long recursiveEndTime = System.nanoTime();
        
        // Iterative Traversal - O(H) space (explicit stack)
        long iterativeStartTime = System.nanoTime();
        List<Integer> iterativeResult = inorderTraversalIterative(root);
        long iterativeEndTime = System.nanoTime();
        
        System.out.println("Results verification:");
        System.out.println("Morris: " + morrisResult);
        System.out.println("Recursive: " + recursiveResult);
        System.out.println("Iterative: " + iterativeResult);
        
        System.out.println("\nPerformance comparison:");
        System.out.printf("Morris: %d ns\n", morrisEndTime - morrisStartTime);
        System.out.printf("Recursive: %d ns\n", recursiveEndTime - recursiveStartTime);
        System.out.printf("Iterative: %d ns\n", iterativeEndTime - iterativeStartTime);
        
        System.out.println("\nSpace complexity:");
        System.out.println("Morris: O(1) - no recursion stack or explicit stack");
        System.out.println("Recursive: O(H) - recursion stack depth");
        System.out.println("Iterative: O(H) - explicit stack size");
    }
    
    // Helper method to create tree
    public TreeNode createTree(Integer[] values) {
        if (values == null || values.length == 0) {
            return null;
        }
        
        return createTreeHelper(values, 0);
    }
    
    private TreeNode createTreeHelper(Integer[] values, int index) {
        if (index >= values.length || values[index] == null) {
            return null;
        }
        
        TreeNode node = new TreeNode(values[index]);
        node.left = createTreeHelper(values, 2 * index + 1);
        node.right = createTreeHelper(values, 2 * index + 2);
        
        return node;
    }
    
    // Helper method to print tree
    public void printTree(TreeNode root) {
        if (root == null) {
            System.out.println("Empty tree");
            return;
        }
        
        List<List<String>> levels = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<String> level = new ArrayList<>();
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                if (node != null) {
                    level.add(String.valueOf(node.val));
                    queue.offer(node.left);
                    queue.offer(node.right);
                } else {
                    level.add("null");
                    queue.offer(null);
                    queue.offer(null);
                }
            }
            
            levels.add(level);
        }
        
        for (int i = 0; i < levels.size(); i++) {
            System.out.println("Level " + i + ": " + levels.get(i));
        }
    }
    
    public static void main(String[] args) {
        BinaryTreeInorderTraversal bt = new BinaryTreeInorderTraversal();
        
        // Test cases
        Integer[][] testCases = {
            {1, null, 2, 3},
            {1, 2, 3, 4, 5},
            {1},
            {1, null, 2, null, 3},
            {5, 3, 1, null, null, null, 4, 6},
            {1, 2, null, 3},
            {10, 5, 15, null, null, 20},
            {1, null, null, 2},
            {1, 2, 3, 4, 5, 6, 7}
        };
        
        String[] descriptions = {
            "Simple tree",
            "Complete tree",
            "Single node",
            "Sparse tree",
            "Unbalanced tree",
            "Missing nodes",
            "Large values",
            "Left heavy",
            "Perfect tree"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Tree structure: " + Arrays.toString(testCases[i]));
            
            TreeNode root = bt.createTree(testCases[i]);
            bt.printTree(root);
            
            // Test all traversal methods
            List<Integer> morrisResult = bt.inorderTraversal(root);
            List<Integer> recursiveResult = bt.inorderTraversalRecursive(root);
            List<Integer> iterativeResult = bt.inorderTraversalIterative(root);
            
            System.out.println("Traversal results:");
            System.out.println("Morris: " + morrisResult);
            System.out.println("Recursive: " + recursiveResult);
            System.out.println("Iterative: " + iterativeResult);
            
            // Test preorder
            List<Integer> preorderResult = bt.preorderTraversalMorris(root);
            System.out.println("Preorder Morris: " + preorderResult);
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        TreeNode detailedRoot = bt.createTree(new Integer[]{4, 2, 5, 1, 3});
        MorrisTraversalResult detailedResult = bt.inorderTraversalDetailed(detailedRoot);
        
        System.out.println("Result: " + detailedResult.result);
        for (String step : detailedResult.explanation) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        TreeNode performanceRoot = bt.createTree(new Integer[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10});
        bt.comparePerformance(performanceRoot, 10000);
        
        // Memory usage analysis
        System.out.println("\n=== Memory Usage Analysis ===");
        TreeNode analysisRoot = bt.createTree(new Integer[]{1, 2, 3, 4, 5, 6, 7});
        bt.memoryUsageAnalysis(analysisRoot);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty tree
        List<Integer> emptyResult = bt.inorderTraversal(null);
        System.out.println("Empty tree: " + emptyResult);
        
        // Single node
        TreeNode singleNode = new TreeNode(42);
        List<Integer> singleResult = bt.inorderTraversal(singleNode);
        System.out.println("Single node: " + singleResult);
        
        // Large tree
        Integer[] largeValues = new Integer[1000];
        for (int i = 0; i < 1000; i++) {
            largeValues[i] = i + 1;
        }
        
        TreeNode largeRoot = bt.createTree(largeValues);
        long startTime = System.nanoTime();
        bt.inorderTraversal(largeRoot);
        long endTime = System.nanoTime();
        System.out.printf("Large tree (1000 nodes): took %d ns\n", endTime - startTime);
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Morris Traversal
- **Threaded Binary Tree**: Uses temporary links for traversal
- **O(1) Space**: No recursion stack or explicit stack needed
- **Inorder Traversal**: Efficient tree traversal without extra space
- **Tree Modification**: Temporarily modifies tree structure

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree Traversal**: Visit nodes in specific order
- **Space Optimization**: Avoid O(H) stack space
- **Tree Modification**: Can temporarily modify tree structure
- **Multiple Orders**: Inorder, preorder, postorder support

## 3. SIMILAR PROBLEMS
- Binary Tree Preorder Traversal
- Binary Tree Postorder Traversal
- Flatten Binary Tree to Linked List
- Validate Binary Search Tree

## 4. KEY OBSERVATIONS
- Morris traversal eliminates need for recursion stack
- Uses predecessor links to navigate back up
- Time complexity: O(N) for N nodes
- Space complexity: O(1) vs O(H) for recursion
- Restores original tree structure after traversal

## 5. VARIATIONS & EXTENSIONS
- Different traversal orders
- Threaded binary trees
- Tree balancing applications
- Iterator implementations

## 6. INTERVIEW INSIGHTS
- Clarify: "Can we modify the tree temporarily?"
- Edge cases: empty tree, single node, skewed tree
- Time complexity: O(N) vs O(N) for all methods
- Space complexity: O(1) vs O(H) for recursion

## 7. COMMON MISTAKES
- Incorrect predecessor finding
- Not restoring tree structure properly
- Wrong temporary link creation
- Infinite loops in traversal
- Missing null pointer checks

## 8. OPTIMIZATION STRATEGIES
- Efficient predecessor finding
- Proper link restoration
- Minimal tree modifications
- Clear traversal logic

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like adding temporary ropes to climb a tree:**
- You have a binary tree you need to traverse
- Normally you'd use a rope (stack) to climb back up
- Morris traversal adds temporary ropes (links) to climb back
- When you visit a node, you add a rope from its predecessor
- This lets you climb back up without using your own rope
- After traversal, you remove all temporary ropes
- This is like a self-guided tour with temporary markers

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root
2. **Goal**: Traverse tree in inorder without extra space
3. **Output**: List of node values in inorder

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(H) space for recursion/stack
- **"How to optimize?"** → Use predecessor links to navigate back
- **"Why Morris traversal?"** → Eliminates need for stack
- **"How to navigate back?"** → Use rightmost node as predecessor

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Morris traversal:
1. Start at root
2. If current has left child:
   - Find predecessor (rightmost node in left subtree)
   - If predecessor.right is null:
     - Make predecessor.right point to current
     - Move current to current.left
   - If predecessor.right points to current:
     - Visit current
     - Move current to current.right
3. If current has no left child:
   - Visit current
   - Move current to current.right
4. Continue until current is null
5. Tree structure is restored automatically"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty list
- **Single node**: Return that node's value
- **Skewed tree**: Handle correctly with minimal operations
- **Large trees**: Ensure O(N) time complexity

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:    4
        /   \
       2     5
      / \
     1   3

Human thinking:
"Let's apply Morris traversal:

Step 1: current = 4
- current.left is not null (2)
- Find predecessor of 4: rightmost in left subtree = 3
- predecessor.right is null
- Make predecessor.right = current (3.right = 4)
- current = current.left (current = 2)

Step 2: current = 2
- current.left is not null (1)
- Find predecessor of 2: rightmost in left subtree = 1
- predecessor.right is null
- Make predecessor.right = current (1.right = 2)
- current = current.left (current = 1)

Step 3: current = 1
- current.left is null
- Visit current: add 1 to result
- current = current.right (current = 2) [via temporary link]

Step 4: current = 2
- current.left is not null (1)
- Find predecessor of 2: rightmost in left subtree = 1
- predecessor.right points to current (1.right = 2)
- Visit current: add 2 to result
- Restore: predecessor.right = null (1.right = null)
- current = current.right (current = 4) [via original right]

Step 5: current = 4
- current.left is not null (2)
- Find predecessor of 4: rightmost in left subtree = 3
- predecessor.right points to current (3.right = 4)
- Visit current: add 4 to result
- Restore: predecessor.right = null (3.right = null)
- current = current.right (current = 5)

Continue...

Final result: [1, 2, 3, 4, 5] ✓

Manual verification:
Inorder traversal visits left subtree, root, right subtree ✓
Tree structure is restored ✓
```

#### Phase 6: Intuition Validation
- **Why it works**: Predecessor links provide back navigation
- **Why it's efficient**: O(1) space vs O(H) for stack
- **Why it's correct**: Visits nodes in correct order

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Uses O(H) stack space
2. **"What about tree modification?"** → Temporary, restored after traversal
3. **"How to find predecessor?"** → Rightmost node in left subtree
4. **"What about infinite loops?"** → Proper link restoration prevents loops

### Real-World Analogy
**Like exploring a maze with temporary markers:**
- You have a maze (tree) you need to explore systematically
- Normally you'd use breadcrumbs (stack) to find your way back
- Morris traversal adds temporary markers (links) in the maze
- When you visit a room, you mark how to get back to previous room
- This lets you explore without leaving your own trail
- After exploration, you remove all temporary markers
- Useful in memory-constrained environments, embedded systems
- Like a self-guided tour that leaves no trace

### Human-Readable Pseudocode
```
function morrisInorder(root):
    result = []
    current = root
    
    while current != null:
        if current.left == null:
            // No left subtree, visit current
            result.add(current.val)
            current = current.right
        else:
            // Find predecessor
            predecessor = current.left
            while predecessor.right != null and predecessor.right != current:
                predecessor = predecessor.right
            
            if predecessor.right == null:
                // Create temporary link
                predecessor.right = current
                current = current.left
            else:
                // Temporary link exists, visit current
                predecessor.right = null  // Restore
                result.add(current.val)
                current = current.right
    
    return result
```

### Execution Visualization

### Example: Tree with root=4, left=2, right=5
```
Morris Traversal Process:

Initial tree:
    4
   / \
  2   5
 / \
1   3

Step 1: current=4
- predecessor of 4 is 3
- Create link: 3.right = 4
- current = 2

Step 2: current=2
- predecessor of 2 is 1
- Create link: 1.right = 2
- current = 1

Step 3: current=1
- current.left is null
- Visit 1, add to result
- current = 1.right = 2 [via temporary link]

Step 4: current=2
- predecessor of 2 is 1
- 1.right points to current (temporary link)
- Visit 2, add to result
- Restore: 1.right = null
- current = 2.right = 4 [via original right]

Step 5: current=4
- predecessor of 4 is 3
- 3.right points to current (temporary link)
- Visit 4, add to result
- Restore: 3.right = null
- current = 4.right = 5

Continue...

Final result: [1, 2, 3, 4, 5] ✓

Visualization:
Temporary links enable back navigation
Tree structure is restored after traversal
Space complexity: O(1) ✓
```

### Key Visualization Points:
- **Predecessor Finding**: Rightmost node in left subtree
- **Temporary Links**: predecessor.right = current
- **Link Restoration**: predecessor.right = null
- **Back Navigation**: Use temporary links to climb back up

### Memory Layout Visualization:
```
Tree Evolution During Morris Traversal:

Initial:
    4
   / \
  2   5
 / \
1   3

After Step 1 (create link 3→4):
    4
   / \
  2   5
 / \
1---3 (temporary link from 3 to 4)

After Step 2 (create link 1→2):
    4
   / \
  2   5
 / \
1---2 (temporary link from 1 to 2)

After restoration:
    4
   / \
  2   5
 / \
1   3 (all temporary links removed)

Morris traversal uses O(1) extra space
Temporary links enable efficient back navigation
```

### Time Complexity Breakdown:
- **Morris Traversal**: O(N) time, O(1) space
- **Recursive Traversal**: O(N) time, O(H) space
- **Iterative Traversal**: O(N) time, O(H) space
- **Optimal**: Best possible space complexity
- **vs Standard**: O(1) vs O(H) space trade-off
*/
