import java.util.*;

public class MaximumDepthOfBinaryTree {
    
    // Definition for a binary tree node.
    public static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        
        TreeNode() {}
        TreeNode(int val) { this.val = val; }
        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    // 104. Maximum Depth of Binary Tree
    // Time: O(N), Space: O(H) where H is tree height (recursion stack)
    public static int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        
        int leftDepth = maxDepth(root.left);
        int rightDepth = maxDepth(root.right);
        
        return Math.max(leftDepth, rightDepth) + 1;
    }

    // Iterative BFS approach
    public static int maxDepthBFS(TreeNode root) {
        if (root == null) {
            return 0;
        }
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        int depth = 0;
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            depth++;
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                
                if (node.left != null) {
                    queue.offer(node.left);
                }
                if (node.right != null) {
                    queue.offer(node.right);
                }
            }
        }
        
        return depth;
    }

    // Helper function to create a binary tree from array
    public static TreeNode createTree(Object[] nums) {
        if (nums.length == 0) {
            return null;
        }
        
        TreeNode[] nodes = new TreeNode[nums.length];
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == null) {
                nodes[i] = null;
            } else {
                nodes[i] = new TreeNode((Integer) nums[i]);
            }
        }
        
        for (int i = 0; i < nums.length; i++) {
            if (nodes[i] != null) {
                int left = 2 * i + 1;
                int right = 2 * i + 2;
                
                if (left < nums.length) {
                    nodes[i].left = nodes[left];
                }
                if (right < nums.length) {
                    nodes[i].right = nodes[right];
                }
            }
        }
        
        return nodes[0];
    }

    public static void main(String[] args) {
        // Test cases
        Object[][] testCases = {
            {3, 9, 20, null, null, 15, 7},
            {1, null, 2},
            {},
            {1},
            {1, 2, 3, 4, 5, null, null, 6, 7, null, null, null, null, 8},
            {1, 2, 3, 4, 5, 6, 7},
            {1, 2, 3, 4, null, null, null, 5},
            {1, null, 2, null, 3, null, 4, null, 5},
            {0},
            {1, 2}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode root = createTree(testCases[i]);
            int result1 = maxDepth(root);
            int result2 = maxDepthBFS(root);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Recursive: %d, BFS: %d\n\n", result1, result2);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Tree Depth-First Search (DFS)
- **Recursive DFS**: Explore tree depth-first using recursion
- **Base Case**: Null node has depth 0
- **Recursive Case**: Depth = 1 + max(leftDepth, rightDepth)
- **Call Stack**: Uses recursion stack for traversal

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Depth Definition**: Number of nodes from root to leaf
- **Longest Path**: Find maximum depth among all root-to-leaf paths
- **Tree Structure**: Can be balanced or unbalanced

## 3. SIMILAR PROBLEMS
- Minimum Depth of Binary Tree
- Balanced Binary Tree
- Diameter of Binary Tree
- Path Sum problems

## 4. KEY OBSERVATIONS
- Depth is 1 + max(leftDepth, rightDepth)
- Base case: null node has depth 0
- Recursive call explores both subtrees
- Maximum depth comes from the longest root-to-leaf path
- Time complexity: O(N) - each node visited once

## 5. VARIATIONS & EXTENSIONS
- Iterative BFS approach
- Find depth of specific node
- Handle N-ary trees
- Return the actual deepest path

## 6. INTERVIEW INSIGHTS
- Clarify: "Is tree guaranteed to be binary?"
- Edge cases: empty tree, single node, skewed tree
- Time complexity: O(N) vs O(N²) naive approach
- Space complexity: O(H) vs O(N) for iterative

## 7. COMMON MISTAKES
- Not handling null base case
- Off-by-one errors in depth calculation
- Stack overflow for very deep trees
- Confusing depth with node count

## 8. OPTIMIZATION STRATEGIES
- Recursive approach is optimal for single query
- Iterative BFS avoids recursion stack limits
- Tail recursion optimization
- Memoization for multiple queries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like exploring a family tree:**
- You're at the root (ancestor)
- You want to find the longest generation chain
- You explore each branch (left and right children)
- For each person, you ask about their descendants
- The longest chain of descendants gives you the depth

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Root of binary tree
2. **Goal**: Find maximum depth (longest root-to-leaf path)
3. **Output**: Integer representing maximum depth

#### Phase 2: Key Insight Recognition
- **"How to measure depth?"** → Count nodes from root to leaf
- **"What's the base case?"** → Null node has depth 0
- **"How to combine results?"** → 1 + max(left, right)
- **"Can I use iteration?"** → Yes, but recursion is natural

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use recursive DFS:
1. If current node is null, return depth 0
2. Otherwise:
   - Recursively find depth of left subtree
   - Recursively find depth of right subtree
   - Return 1 + maximum of these two depths
3. This naturally finds the longest path"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return 0 (no nodes)
- **Single node**: Return 1 (just the root)
- **Skewed tree**: Recursion depth equals number of nodes
- **Balanced tree**: Depth is log₂(N) approximately

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:    3
        /   \
       9   20
      / \   / \
     15   7

Human thinking:
"Start at root (3):
- Left subtree depth?
  - Node 9: left is null (depth 0), right is null (depth 0)
  - So depth of node 9 = 1 + max(0,0) = 1
- Right subtree depth?
  - Node 20: left is 15, right is 7
  - Need depth of node 15:
    - Left is null (0), right is null (0)
    - So depth of 15 = 1 + max(0,0) = 1
  - Need depth of node 7:
    - Both children are null (0)
    - So depth of 7 = 1 + max(0,0) = 1
  - So depth of 20 = 1 + max(1,1) = 2
- Final answer: 1 + max(1,2) = 3

The longest path is 3→9→15 or 3→20→7"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each node contributes 1 to path length
- **Why it's efficient**: Each node visited exactly once
- **Why it's correct**: Recursion naturally explores all paths

### Common Human Pitfalls & How to Avoid Them
1. **"Why not count all paths?"** → Too expensive O(N²)
2. **"What about iterative approach?"** → Possible but less natural
3. **"How to handle null?"** → Always check before accessing children
4. **"What about very deep trees?"** → May cause stack overflow

### Real-World Analogy
**Like finding the longest family generation:**
- You're studying a family tree starting from an ancestor
- You want to find the longest chain of generations
- For each person, you explore their descendants
- The longest chain from ancestor to any descendant gives you the answer
- Each person contributes exactly 1 to their generation level

### Human-Readable Pseudocode
```
function maxDepth(root):
    if root is null:
        return 0
    
    leftDepth = maxDepth(root.left)
    rightDepth = maxDepth(root.right)
    
    return 1 + max(leftDepth, rightDepth)
```

### Execution Visualization

### Example: Tree with root 3, left 9, right 20; 20 has left 15, right 7
```
Tree Structure:
        3
       /   \
      9   20
     / \   / \
    15   7

DFS Traversal:
Start at root (3):
→ Depth = ?

Explore left subtree (node 9):
→ 9 has no children, depth = 1

Explore right subtree (node 20):
→ 20 has children 15 and 7
→ 
→ Explore 15:
→ 15 has no children, depth = 2
→ 
→ Explore 7:
→ 7 has no children, depth = 2
→ 
→ Depth of 20 = 1 + max(2,2) = 3
→ 
→ Final answer: 1 + max(1,3) = 4

Wait, let me recalculate:

Depth of 9 = 1 (no children)
Depth of 15 = 2 (no children)
Depth of 7 = 2 (no children)
Depth of 20 = 1 + max(2,2) = 3
Final answer: 1 + max(1,3) = 4

The longest path is 3→20→15 (length 3) ✓
```

### Key Visualization Points:
- **Recursive DFS** explores each subtree completely
- **Base case** handles null nodes correctly
- **Depth calculation** adds 1 for current node
- **Maximum function** finds longest path

### Memory Layout Visualization:
```
Call Stack Evolution:
maxDepth(3)
├─ maxDepth(9) → returns 1
└─ maxDepth(20)
   ├─ maxDepth(15) → returns 2
   └─ maxDepth(7) → returns 2
   └─ returns 1 + max(2,2) = 3
└─ returns 1 + max(1,3) = 4

Tree Paths:
3→9 (length 2)
3→20→15 (length 3) ← Longest!
3→20→7 (length 3)
```

### Time Complexity Breakdown:
- **Each Node**: Visited exactly once
- **Recursive Calls**: 2 calls per node (left and right)
- **Total**: O(N) time where N is number of nodes
- **Space**: O(H) for recursion stack, H is tree height
- **Optimal**: Cannot do better than O(N) for this problem
*/
