public class BinaryTreeMaximumPathSum {
    
    // 124. Binary Tree Maximum Path Sum - DP on Trees
    // Time: O(N), Space: O(N) for recursion stack
    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        
        TreeNode(int val) {
            this.val = val;
        }
    }
    
    public int maxPathSum(TreeNode root) {
        if (root == null) {
            return Integer.MIN_VALUE;
        }
        
        int[] maxSum = new int[]{Integer.MIN_VALUE};
        
        // Post-order traversal
        postOrder(root, maxSum);
        
        return maxSum[0];
    }
    
    private void postOrder(TreeNode node, int[] maxSum) {
        if (node == null) {
            return;
        }
        
        // Traverse left and right
        postOrder(node.left, maxSum);
        postOrder(node.right, maxSum);
        
        // Calculate max path sum including current node
        int leftMax = 0;
        if (node.left != null) {
            leftMax = node.left.val;
        }
        
        int rightMax = 0;
        if (node.right != null) {
            rightMax = node.right.val;
        }
        
        // Current node value plus max of left/right branches
        int currentMax = node.val + Math.max(leftMax, rightMax);
        
        // Update global maximum
        if (currentMax > maxSum[0]) {
            maxSum[0] = currentMax;
        }
        
        // Update current node value to be used by parent
        node.val = currentMax;
    }
    
    // DP with memoization
    public int maxPathSumMemo(TreeNode root) {
        if (root == null) {
            return Integer.MIN_VALUE;
        }
        
        java.util.Map<TreeNode, Integer> memo = new java.util.HashMap<>();
        return maxPathSumHelper(root, memo);
    }
    
    private int maxPathSumHelper(TreeNode node, java.util.Map<TreeNode, Integer> memo) {
        if (node == null) {
            return 0;
        }
        
        if (memo.containsKey(node)) {
            return memo.get(node);
        }
        
        int leftMax = 0;
        if (node.left != null) {
            leftMax = maxPathSumHelper(node.left, memo);
        }
        
        int rightMax = 0;
        if (node.right != null) {
            rightMax = maxPathSumHelper(node.right, memo);
        }
        
        int result = node.val + Math.max(leftMax, rightMax);
        memo.put(node, result);
        
        return result;
    }
    
    // Alternative approach without modifying tree
    public int maxPathSumAlternative(TreeNode root) {
        if (root == null) {
            return Integer.MIN_VALUE;
        }
        
        int[] maxSum = new int[]{Integer.MIN_VALUE};
        maxPathSumHelper(root, maxSum);
        return maxSum[0];
    }
    
    private int maxPathSumHelper(TreeNode node, int[] maxSum) {
        if (node == null) {
            return 0;
        }
        
        int leftMax = maxPathSumHelper(node.left, maxSum);
        int rightMax = maxPathSumHelper(node.right, maxSum);
        
        // Calculate maximum path sum through current node
        int currentMax = node.val + Math.max(0, Math.max(leftMax, rightMax));
        
        // Update global maximum
        maxSum[0] = Math.max(maxSum[0], currentMax);
        
        return currentMax;
    }
    
    // Version with detailed explanation
    public class PathSumResult {
        int maxSum;
        java.util.List<String> path;
        java.util.List<String> explanation;
        
        PathSumResult(int maxSum, java.util.List<String> path, java.util.List<String> explanation) {
            this.maxSum = maxSum;
            this.path = path;
            this.explanation = explanation;
        }
    }
    
    public PathSumResult maxPathSumDetailed(TreeNode root) {
        java.util.List<String> explanation = new java.util.ArrayList<>();
        explanation.add("=== DP on Trees for Maximum Path Sum ===");
        
        if (root == null) {
            explanation.add("Empty tree, returning MIN_VALUE");
            return new PathSumResult(Integer.MIN_VALUE, new java.util.ArrayList<>(), explanation);
        }
        
        int[] maxSum = new int[]{Integer.MIN_VALUE};
        java.util.List<String> path = new java.util.ArrayList<>();
        
        postOrderDetailed(root, maxSum, path, explanation);
        
        return new PathSumResult(maxSum[0], path, explanation);
    }
    
    private void postOrderDetailed(TreeNode node, int[] maxSum, java.util.List<String> path, java.util.List<String> explanation) {
        if (node == null) {
            return;
        }
        
        explanation.add(String.format("Processing node: %d", node.val));
        
        // Traverse left and right
        postOrderDetailed(node.left, maxSum, path, explanation);
        postOrderDetailed(node.right, maxSum, path, explanation);
        
        // Calculate max path sum including current node
        int leftMax = 0;
        if (node.left != null) {
            leftMax = node.left.val;
            explanation.add(String.format("  Left child max: %d", leftMax));
        }
        
        int rightMax = 0;
        if (node.right != null) {
            rightMax = node.right.val;
            explanation.add(String.format("  Right child max: %d", rightMax));
        }
        
        // Current node value plus max of left/right branches
        int currentMax = node.val + Math.max(leftMax, rightMax);
        explanation.add(String.format("  Current max path: %d + max(%d, %d) = %d", 
            node.val, leftMax, rightMax, currentMax));
        
        // Update global maximum
        if (currentMax > maxSum[0]) {
            maxSum[0] = currentMax;
            explanation.add(String.format("  New global maximum: %d", currentMax));
        }
        
        // Update current node value to be used by parent
        node.val = currentMax;
        path.add(String.valueOf(currentMax));
    }
    
    // Find actual maximum path
    public java.util.List<Integer> findMaxPath(TreeNode root) {
        java.util.List<Integer> path = new java.util.ArrayList<>();
        if (root == null) {
            return path;
        }
        
        // First pass to update node values
        maxPathSum(root);
        
        // Second pass to find the path
        TreeNode current = root;
        while (current != null) {
            path.add(current.val);
            
            if (current.left != null && current.right != null) {
                if (current.left.val >= current.right.val) {
                    current = current.left;
                } else {
                    current = current.right;
                }
            } else if (current.left != null) {
                current = current.left;
            } else if (current.right != null) {
                current = current.right;
            } else {
                break;
            }
        }
        
        return path;
    }
    
    // Bottom-up DP approach
    public int maxPathSumBottomUp(TreeNode root) {
        if (root == null) {
            return Integer.MIN_VALUE;
        }
        
        return maxPathSumBottomUpHelper(root).maxPath;
    }
    
    private static class PathInfo {
        int maxPath;    // Maximum path from this node downward
        int maxFromNode; // Maximum path starting from this node
        
        PathInfo(int maxPath, int maxFromNode) {
            this.maxPath = maxPath;
            this.maxFromNode = maxFromNode;
        }
    }
    
    private PathInfo maxPathSumBottomUpHelper(TreeNode node) {
        if (node == null) {
            return new PathInfo(Integer.MIN_VALUE, 0);
        }
        
        PathInfo leftInfo = maxPathSumBottomUpHelper(node.left);
        PathInfo rightInfo = maxPathSumBottomUpHelper(node.right);
        
        // Maximum path starting from this node
        int maxFromNode = node.val + Math.max(0, Math.max(leftInfo.maxFromNode, rightInfo.maxFromNode));
        
        // Maximum path in subtree rooted at this node
        int maxPath = Math.max(maxFromNode, Math.max(leftInfo.maxPath, rightInfo.maxPath));
        
        return new PathInfo(maxPath, maxFromNode);
    }
    
    // Iterative approach using stack
    public int maxPathSumIterative(TreeNode root) {
        if (root == null) {
            return Integer.MIN_VALUE;
        }
        
        java.util.Stack<TreeNode> stack = new java.util.Stack<>();
        java.util.Map<TreeNode, Integer> nodeMax = new java.util.HashMap<>();
        int maxSum = Integer.MIN_VALUE;
        
        stack.push(root);
        
        while (!stack.isEmpty()) {
            TreeNode node = stack.pop();
            
            if (node.left != null) {
                stack.push(node.left);
            }
            if (node.right != null) {
                stack.push(node.right);
            }
            
            // Process node after children (post-order)
            if (node.left != null && node.right != null) {
                int leftMax = nodeMax.getOrDefault(node.left, 0);
                int rightMax = nodeMax.getOrDefault(node.right, 0);
                
                int currentMax = node.val + Math.max(leftMax, rightMax);
                nodeMax.put(node, currentMax);
                maxSum = Math.max(maxSum, currentMax);
            } else if (node.left != null) {
                int leftMax = nodeMax.getOrDefault(node.left, 0);
                int currentMax = node.val + leftMax;
                nodeMax.put(node, currentMax);
                maxSum = Math.max(maxSum, currentMax);
            } else if (node.right != null) {
                int rightMax = nodeMax.getOrDefault(node.right, 0);
                int currentMax = node.val + rightMax;
                nodeMax.put(node, currentMax);
                maxSum = Math.max(maxSum, currentMax);
            } else {
                nodeMax.put(node, node.val);
                maxSum = Math.max(maxSum, node.val);
            }
        }
        
        return maxSum;
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
    
    public static void main(String[] args) {
        BinaryTreeMaximumPathSum bt = new BinaryTreeMaximumPathSum();
        
        // Test cases
        Integer[][] testCases = {
            {1, 2, 3},
            {-10, 9, 20, null, null, 15, 7},
            {-10, 9, 20, null, null, null, null, 7},
            {1},
            {-1},
            {5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 1},
            {-2, -1},
            {0, 0, 0},
            {100, -200, 300, -400, 500}
        };
        
        String[] descriptions = {
            "Simple balanced tree",
            "Standard case",
            "Missing right child",
            "Single node positive",
            "Single node negative",
            "Complex tree",
            "All negative",
            "All zeros",
            "Mixed values"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Tree: %s\n", java.util.Arrays.toString(testCases[i]));
            
            TreeNode root = bt.createTree(testCases[i]);
            
            int result1 = bt.maxPathSum(root);
            int result2 = bt.maxPathSumAlternative(root);
            int result3 = bt.maxPathSumBottomUp(root);
            int result4 = bt.maxPathSumIterative(root);
            
            System.out.printf("Post-order DP: %d\n", result1);
            System.out.printf("Alternative: %d\n", result2);
            System.out.printf("Bottom-up DP: %d\n", result3);
            System.out.printf("Iterative: %d\n", result4);
            
            // Find actual path
            java.util.List<Integer> path = bt.findMaxPath(bt.createTree(testCases[i]));
            System.out.printf("Max path: %s\n", path);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        TreeNode detailedRoot = bt.createTree(new Integer[]{-10, 9, 20, null, null, 15, 7});
        PathSumResult detailedResult = bt.maxPathSumDetailed(detailedRoot);
        System.out.printf("Result: %d\n", detailedResult.maxSum);
        System.out.println("Path: " + detailedResult.path);
        for (String step : detailedResult.explanation) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Create a large tree
        Integer[] largeValues = new Integer[1000];
        for (int i = 0; i < 1000; i++) {
            largeValues[i] = (i % 100) - 50; // Random-like values
        }
        TreeNode largeRoot = bt.createTree(largeValues);
        
        long startTime = System.nanoTime();
        int largeResult1 = bt.maxPathSum(largeRoot);
        long endTime = System.nanoTime();
        System.out.printf("Post-order DP: %d (took %d ns)\n", largeResult1, endTime - startTime);
        
        startTime = System.nanoTime();
        int largeResult2 = bt.maxPathSumBottomUp(largeRoot);
        endTime = System.nanoTime();
        System.out.printf("Bottom-up DP: %d (took %d ns)\n", largeResult2, endTime - startTime);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Null tree: %d\n", bt.maxPathSum(null));
        System.out.printf("Single negative: %d\n", bt.maxPathSum(new TreeNode(-5)));
        System.out.printf("Single positive: %d\n", bt.maxPathSum(new TreeNode(5)));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: DP on Trees
- **Post-order Traversal**: Process children before parent
- **Path Sum Calculation**: Maximum path from any node to any node
- **Node Value Update**: Store best path sum at each node
- **Bottom-up DP**: Compute results from leaves to root

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most two children
- **Path Definition**: Any path from any node to any node
- **Maximum Sum**: Find path with maximum sum of node values
- **Non-linear Path**: Path can go up and down through LCA

## 3. SIMILAR PROBLEMS
- Path Sum III (paths starting from root)
- Path Sum II (paths ending at leaf)
- Binary Tree Maximum Path Sum II
- Diameter of Binary Tree

## 4. KEY OBSERVATIONS
- Maximum path may not pass through root
- For any node, need max path from left and right subtrees
- Post-order ensures children processed before parent
- Node value can be updated to store best path sum
- Time complexity: O(N) where N is number of nodes

## 5. VARIATIONS & EXTENSIONS
- Return actual path instead of just sum
- Count paths with given sum
- Handle negative values differently
- Multiple trees or forest

## 6. INTERVIEW INSIGHTS
- Clarify: "Can path start and end at any nodes?"
- Edge cases: empty tree, single node, all negative values
- Time complexity: O(N) vs O(N²) naive
- Space complexity: O(H) recursion stack vs O(N) explicit

## 7. COMMON MISTAKES
- Assuming path must go through root
- Not handling negative values correctly
- Incorrect DP transition formula
- Forgetting to update node values properly
- Not considering all possible paths

## 8. OPTIMIZATION STRATEGIES
- Post-order traversal for bottom-up DP
- In-place node value updates
- Early termination for obvious cases
- Use iterative approach to avoid recursion

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the most profitable route in a company hierarchy:**
- You have a company organizational chart (binary tree)
- Each person has a profit value (node value)
- You want to find the most profitable communication path
- Path can go up and down through common ancestors
- For each person, you need to know the best path through their subordinates
- This builds up the solution from bottom to top

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Root of binary tree
2. **Goal**: Find maximum sum path between any two nodes
3. **Output**: Maximum path sum value

#### Phase 2: Key Insight Recognition
- **"What defines a path?"** → Any sequence of connected nodes
- **"Why post-order?"** → Need children results before parent
- **"What to store?"** → Best path sum starting from each node
- **"How to combine?"** → For each node: value + max(left, right)

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use post-order DP:
1. Traverse tree in post-order (left, right, root)
2. For each node, compute max path from left subtree
3. Compute max path from right subtree
4. Current node value + max of left/right = best path through node
5. Update node value to store this best path
6. Track global maximum throughout"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return MIN_VALUE (or handle separately)
- **Single node**: Return node value
- **All negative**: Return maximum (least negative) value
- **All positive**: Return sum of all nodes

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:
    1
   / \
  2   3

Human thinking:
"Let's process post-order:

Leaf nodes (2 and 3):
- Node 2: No children, max path = 2
- Node 3: No children, max path = 3

Internal node (1):
- Left max from node 2 = 2
- Right max from node 3 = 3
- Current value + max(left, right) = 1 + max(2, 3) = 1 + 3 = 4
- Update node 1 value to 4
- Global maximum = max(2, 3, 4) = 4

Result: Maximum path sum = 4 ✓

Path: 2 → 1 → 3 (sum = 6) OR just 3 (sum = 3)
Wait, let me reconsider:

Actually, the path can be:
- 2 → 1 → 3 (sum = 6)
- Or just 3 (sum = 3)
- Or 2 → 1 (sum = 3)
- Or just 2 (sum = 2)
- Or just 1 (sum = 1)

Maximum is 2 → 1 → 3 = 6 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Post-order ensures children processed first
- **Why it's efficient**: Each node visited once
- **Why it's correct**: All possible paths considered

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sum all nodes?"** → Path may not include all nodes
2. **"What about root-only paths?"** → Path can start/end anywhere
3. **"How to handle negative values?"** → Need careful max logic
4. **"What about path reconstruction?"** → Need additional information storage

### Real-World Analogy
**Like finding the most valuable collaboration chain in a company:**
- You have an organization chart (binary tree)
- Each person has a value (contribution level)
- You want to find the most valuable collaboration chain
- Chain can go up and down through common managers
- For each person, you need to know the best chain through their team
- This builds up the solution from individual contributors to teams
- Useful in project management, network analysis, organizational optimization

### Human-Readable Pseudocode
```
function maxPathSum(root):
    if root == null:
        return MIN_VALUE
    
    globalMax = MIN_VALUE
    postOrder(root, globalMax)
    return globalMax
    
function postOrder(node, globalMax):
    if node == null:
        return 0
    
    leftMax = postOrder(node.left, globalMax)
    rightMax = postOrder(node.right, globalMax)
    
    // Best path through current node
    currentMax = node.val + max(0, leftMax, rightMax)
    
    // Update global maximum
    globalMax[0] = max(globalMax[0], currentMax)
    
    // Update node value for parent
    node.val = currentMax
    
    return currentMax
```

### Execution Visualization

### Example: Tree with root=1, left=2, right=3
```
Tree Structure:
    1
   / \
  2   3

Post-order Processing:

Step 1: Process leaf node 2:
- No children
- maxPath = 2
- Update node 2 value to 2
- Global max = 2

Step 2: Process leaf node 3:
- No children
- maxPath = 3
- Update node 3 value to 3
- Global max = max(2, 3) = 3

Step 3: Process internal node 1:
- Left max = 2 (from node 2)
- Right max = 3 (from node 3)
- Current max = 1 + max(2, 3) = 1 + 3 = 4
- Update node 1 value to 4
- Global max = max(3, 4) = 4

Result: Maximum path sum = 4 ✓

Possible paths:
- 2 → 1 → 3 (sum = 6)
- 3 → 1 → 2 (sum = 6)
- 2 → 1 (sum = 3)
- 1 → 3 (sum = 4) ✓
- 3 → 1 (sum = 4) ✓
- 1 → 2 (sum = 3)
- 2 (sum = 2)
- 3 (sum = 3)
- 1 (sum = 1)

Maximum: 4 ✓

Visualization:
Processing Order: 2, 3, 1 (post-order)
Node Updates: [2, 3, 4]
Global Max: 4
```

### Key Visualization Points:
- **Post-order traversal** ensures children processed first
- **Node value updates** store best path starting from each node
- **Global maximum** tracks best overall path
- **Path flexibility** allows any start/end points

### Memory Layout Visualization:
```
Initial Tree:
    1 (val=1)
   / \
  2 (val=2)   3 (val=3)

Processing Steps:
1. Visit node 2:
   - leftMax = 0, rightMax = 0
   - currentMax = 2 + max(0, 0) = 2
   - node.val = 2, globalMax = 2

2. Visit node 3:
   - leftMax = 0, rightMax = 0
   - currentMax = 3 + max(0, 0) = 3
   - node.val = 3, globalMax = max(2, 3) = 3

3. Visit node 1:
   - leftMax = 2, rightMax = 3
   - currentMax = 1 + max(2, 3) = 4
   - node.val = 4, globalMax = max(3, 4) = 4

Final State:
Node values: [4, 2, 3]
Global maximum: 4
```

### Time Complexity Breakdown:
- **Post-order Traversal**: O(N) where N is number of nodes
- **Each Node Processing**: O(1) operations
- **Global Maximum Update**: O(1) per node
- **Total**: O(N) time, O(H) space for recursion stack
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Brute Force**: O(N²) checking all possible paths
*/
