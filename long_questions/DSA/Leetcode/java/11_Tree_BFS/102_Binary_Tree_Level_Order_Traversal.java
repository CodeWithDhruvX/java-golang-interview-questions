import java.util.*;

public class BinaryTreeLevelOrderTraversal {
    
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

    // 102. Binary Tree Level Order Traversal
    // Time: O(N), Space: O(N)
    public static List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) {
            return result;
        }
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<Integer> currentLevel = new ArrayList<>();
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                currentLevel.add(node.val);
                
                if (node.left != null) {
                    queue.offer(node.left);
                }
                if (node.right != null) {
                    queue.offer(node.right);
                }
            }
            
            result.add(currentLevel);
        }
        
        return result;
    }

    // Recursive approach
    public static List<List<Integer>> levelOrderRecursive(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) {
            return result;
        }
        
        dfs(root, 0, result);
        return result;
    }

    private static void dfs(TreeNode node, int level, List<List<Integer>> result) {
        if (node == null) {
            return;
        }
        
        // Ensure result has enough levels
        while (result.size() <= level) {
            result.add(new ArrayList<>());
        }
        
        result.get(level).add(node.val);
        
        dfs(node.left, level + 1, result);
        dfs(node.right, level + 1, result);
    }

    // Helper function to create a binary tree from array
    public static TreeNode createTree(Object[] nums) {
        if (nums.length == 0) {
            return null;
        }
        
        TreeNode[] nodes = new TreeNode[nums.length];
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != null) {
                nodes[i] = new TreeNode((Integer) nums[i]);
            }
        }
        
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != null) {
                int left = 2 * i + 1;
                int right = 2 * i + 2;
                
                if (left < nums.length && nums[left] != null) {
                    nodes[i].left = nodes[left];
                }
                if (right < nums.length && nums[right] != null) {
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
            {1},
            {},
            {1, 2, 3, 4, 5, 6, 7},
            {1, null, 2, null, 3, null, 4, null, 5},
            {1, 2, 3, 4, null, null, null, 5},
            {1, 2, 3, null, null, null, 4, null, null, 5},
            {1, 2, null, 3, null, 4, null, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode root = createTree(testCases[i]);
            List<List<Integer>> result1 = levelOrder(root);
            List<List<Integer>> result2 = levelOrderRecursive(root);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Iterative: %s\n", result1);
            System.out.printf("  Recursive: %s\n\n", result2);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Tree Breadth-First Search (BFS)
- **Level Order Traversal**: Visit nodes level by level
- **Queue-Based**: Use queue to track current level
- **Level Processing**: Process all nodes at current depth
- **Next Level**: Add children of current level to queue

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Level Order**: Nodes visited by depth from root
- **BFS Traversal**: Systematic exploration by levels
- **Output Structure**: List of lists, each containing one level

## 3. SIMILAR PROBLEMS
- Maximum Depth of Binary Tree
- Minimum Depth of Binary Tree
- Zigzag Level Order Traversal
- Binary Tree Right Side View

## 4. KEY OBSERVATIONS
- Queue naturally implements BFS order
- Each iteration processes one complete level
- Level size tells us how many nodes at current depth
- Children of current level become next level
- Time complexity: O(N) - each node visited once

## 5. VARIATIONS & EXTENSIONS
- Recursive DFS with level tracking
- Zigzag traversal (alternating directions)
- Bottom-up level order
- Stream processing for very large trees

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I return empty list for empty tree?"
- Edge cases: empty tree, single node, complete binary tree
- Time complexity: O(N) vs O(N log N) for balanced tree
- Space complexity: O(N) for queue + result

## 7. COMMON MISTAKES
- Not handling empty tree case
- Incorrect level boundary management
- Mixing up level order with pre-order
- Not adding children to queue properly

## 8. OPTIMIZATION STRATEGIES
- BFS is optimal for level order traversal
- Pre-allocate result lists
- Use array-based queue for better cache performance
- Iterative approach avoids recursion stack limits

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like reading a book page by page:**
- You have a tree (book with chapters and sections)
- You want to read it level by level (page by page)
- Start with the title page (root)
- Read all sections at the same level (same chapter)
- Then move to the next level (next chapter)
- Continue until you've read the entire book

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Root of binary tree
2. **Goal**: Return nodes in level order (BFS order)
3. **Output**: List of lists, each containing one tree level

#### Phase 2: Key Insight Recognition
- **"How to process by levels?"** → Use queue for BFS
- **"What goes in queue?"** → Nodes to process, in BFS order
- **"How to track levels?"** → Process all nodes at current depth
- **"When to move to next level?"** → After processing current level

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use BFS with a queue:
1. Start with root in queue
2. While queue is not empty:
   - Get current level size (nodes at this depth)
   - Process all nodes at this level
   - Add their children to queue (next level)
   - Collect current level values in result
3. This naturally processes tree level by level"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty list
- **Single node**: Return [[root.value]]
- **Complete tree**: All levels filled completely
- **Skewed tree**: Each level has only one node

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:    3
        /   \
       9   20
      / \   / \
     15   7

Human thinking:
"Initialize: queue = [3], result = [], depth = 0

Level 0 (depth 0):
- Level size = 1 (just root)
- Process node 3, add to result[0] = [3]
- Add children: queue = [9,20]

Level 1 (depth 1):
- Level size = 2 (nodes 9,20)
- Process node 9, add to result[1] = [9]
- Process node 20, add to result[1] = [9,20]
- Add children: queue = [15,7]

Level 2 (depth 2):
- Level size = 2 (nodes 15,7)
- Process node 15, add to result[2] = [15]
- Process node 7, add to result[2] = [15,7]
- No children to add
- Queue is empty

Final result: [[3], [9,20], [15,7]] ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Queue ensures FIFO order for BFS
- **Why it's efficient**: Each node visited exactly once
- **Why it's correct**: Natural level-by-level processing

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DFS?"** → DFS doesn't preserve level order
2. **"What about recursion?"** → Possible but queue is more natural
3. **"How to track levels?"** → Process all nodes at current depth
4. **"What about level boundaries?"** → Use queue size for current level

### Real-World Analogy
**Like organizing company hierarchy by department:**
- You have an organization chart (tree structure)
- You want to list employees by management level
- Start with CEO (root)
- List all executives at same level (first level)
- Then list all managers at next level
- Continue down to individual contributors
- Each level represents a management tier

### Human-Readable Pseudocode
```
function levelOrder(root):
    if root is null:
        return []
    
    queue = [root]
    result = []
    
    while queue is not empty:
        levelSize = queue.length
        currentLevel = []
        
        for i from 1 to levelSize:
            node = queue.dequeue()
            currentLevel.add(node.value)
            
            if node.left is not null:
                queue.enqueue(node.left)
            if node.right is not null:
                queue.enqueue(node.right)
        
        result.add(currentLevel)
    
    return result
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

BFS Traversal:
Initial: queue = [3], result = []

Level 0:
- Process: [3], result = [[3]]
- Queue: [9,20]

Level 1:
- Process: [9,20], result = [[3], [9,20]]
- Queue: [15,7]

Level 2:
- Process: [15,7], result = [[3], [9,20], [15,7]]
- Queue: []

Final: [[3], [9,20], [15,7]] ✓
```

### Key Visualization Points:
- **Queue BFS** naturally processes level by level
- **Level size** determines current processing batch
- **Children addition** prepares next level
- **Result building** accumulates levels in order

### Memory Layout Visualization:
```
Queue Evolution:
Start: [3]                    Process: [3]
Next: [9,20]                  Process: [9,20]
Next: [15,7]                   Process: [15,7]
Next: []                        Process: []

Result Building:
Level 0: [3]
Level 1: [9,20]
Level 2: [15,7]

Final: [[3], [9,20], [15,7]]
```

### Time Complexity Breakdown:
- **Each Node**: Visited exactly once
- **Queue Operations**: O(1) per node (enqueue/dequeue)
- **Level Processing**: O(N) total across all levels
- **Total**: O(N) time where N is number of nodes
- **Space**: O(N) for queue + O(N) for result = O(N)
- **Optimal**: Cannot do better than O(N) for this problem
*/
