public class ValidateBinarySearchTree {
    
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

    // 98. Validate Binary Search Tree
    // Time: O(N), Space: O(H) where H is tree height (recursion stack)
    public static boolean isValidBST(TreeNode root) {
        return validateBST(root, null, null);
    }

    private static boolean validateBST(TreeNode node, Integer min, Integer max) {
        if (node == null) {
            return true;
        }
        
        // Check current node value against bounds
        if (min != null && node.val <= min) {
            return false;
        }
        if (max != null && node.val >= max) {
            return false;
        }
        
        // Recursively check left and right subtrees
        return validateBST(node.left, min, node.val) && 
               validateBST(node.right, node.val, max);
    }

    // Alternative approach using in-order traversal
    public static boolean isValidBSTInorder(TreeNode root) {
        TreeNodeWrapper prev = new TreeNodeWrapper();
        return inorder(root, prev);
    }
    
    private static class TreeNodeWrapper {
        TreeNode node;
    }
    
    private static boolean inorder(TreeNode node, TreeNodeWrapper prev) {
        if (node == null) {
            return true;
        }
        
        // Check left subtree
        if (!inorder(node.left, prev)) {
            return false;
        }
        
        // Check current node
        if (prev.node != null && node.val <= prev.node.val) {
            return false;
        }
        prev.node = node;
        
        // Check right subtree
        return inorder(node.right, prev);
    }

    // Iterative approach using stack
    public static boolean isValidBSTIterative(TreeNode root) {
        if (root == null) {
            return true;
        }
        
        java.util.Stack<TreeNode> stack = new java.util.Stack<>();
        TreeNode prev = null;
        TreeNode current = root;
        
        while (!stack.isEmpty() || current != null) {
            // Go as far left as possible
            while (current != null) {
                stack.push(current);
                current = current.left;
            }
            
            // Process node
            current = stack.pop();
            
            if (prev != null && current.val <= prev.val) {
                return false;
            }
            prev = current;
            
            // Move to right subtree
            current = current.right;
        }
        
        return true;
    }

    // Range-based validation
    public static boolean isValidBSTRange(TreeNode root) {
        return isValidBSTRangeHelper(root, Integer.MIN_VALUE, Integer.MAX_VALUE);
    }
    
    private static boolean isValidBSTRangeHelper(TreeNode node, long min, long max) {
        if (node == null) {
            return true;
        }
        
        if (node.val <= min || node.val >= max) {
            return false;
        }
        
        return isValidBSTRangeHelper(node.left, min, node.val) &&
               isValidBSTRangeHelper(node.right, node.val, max);
    }

    // Helper function to create a BST from array
    public static TreeNode createBST(Integer[] nums) {
        return createBSTHelper(nums, 0, nums.length - 1);
    }
    
    private static TreeNode createBSTHelper(Integer[] nums, int left, int right) {
        if (left > right) {
            return null;
        }
        
        int mid = left + (right - left) / 2;
        TreeNode node = new TreeNode(nums[mid]);
        
        node.left = createBSTHelper(nums, left, mid - 1);
        node.right = createBSTHelper(nums, mid + 1, right);
        
        return node;
    }

    // Helper function to create an invalid BST
    public static TreeNode createInvalidBST() {
        TreeNode root = new TreeNode(5);
        root.left = new TreeNode(1);
        root.right = new TreeNode(4);
        root.right.left = new TreeNode(3);
        root.right.right = new TreeNode(6);
        return root;
    }

    public static void main(String[] args) {
        // Test cases
        System.out.println("=== Testing Valid BST ===");
        Integer[] nums1 = {2, 1, 3, 6, 9, 4, 7};
        TreeNode validBST = createBST(nums1);
        
        boolean result1 = isValidBST(validBST);
        boolean result2 = isValidBSTInorder(validBST);
        boolean result3 = isValidBSTIterative(validBST);
        boolean result4 = isValidBSTRange(validBST);
        
        System.out.printf("Valid BST - Recursive: %b, Inorder: %b, Iterative: %b, Range: %b\n", 
            result1, result2, result3, result4);
        
        System.out.println("\n=== Testing Invalid BST ===");
        TreeNode invalidBST = createInvalidBST();
        
        boolean result5 = isValidBST(invalidBST);
        boolean result6 = isValidBSTInorder(invalidBST);
        boolean result7 = isValidBSTIterative(invalidBST);
        boolean result8 = isValidBSTRange(invalidBST);
        
        System.out.printf("Invalid BST - Recursive: %b, Inorder: %b, Iterative: %b, Range: %b\n", 
            result5, result6, result7, result8);
        
        System.out.println("\n=== Testing Edge Cases ===");
        TreeNode nullTree = null;
        TreeNode singleNode = new TreeNode(1);
        
        System.out.printf("Null tree: %b\n", isValidBST(nullTree));
        System.out.printf("Single node: %b\n", isValidBST(singleNode));
        
        // Test with extreme values
        TreeNode extremeTree = new TreeNode(Integer.MAX_VALUE);
        extremeTree.left = new TreeNode(Integer.MIN_VALUE);
        System.out.printf("Extreme values: %b\n", isValidBST(extremeTree));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search Tree Properties
- **BST Validation**: Check if tree follows BST ordering rules
- **Range Validation**: Each node must be within valid range
- **Recursive Traversal**: Validate left and right subtrees
- **Multiple Approaches**: Range-based, inorder, iterative stack

## 2. PROBLEM CHARACTERISTICS
- **Binary Search Tree**: Left subtree < node < right subtree
- **Validation**: Check if given tree is valid BST
- **Ordering Constraint**: All nodes must follow BST property
- **Recursive Structure**: Tree traversal with bounds checking

## 3. SIMILAR PROBLEMS
- Recover Binary Search Tree
- Kth Smallest Element in BST
- Validate Binary Tree
- Convert Sorted Array to BST

## 4. KEY OBSERVATIONS
- BST property: left < node < right for all nodes
- Range validation: pass min/max bounds down recursively
- Inorder traversal of BST yields sorted sequence
- Iterative approach avoids recursion stack overflow
- Time complexity: O(N) where N is number of nodes

## 5. VARIATIONS & EXTENSIONS
- Find minimum value in BST
- Check if tree is balanced
- Validate with duplicate values allowed
- Convert between BST and array representations

## 6. INTERVIEW INSIGHTS
- Clarify: "Are duplicate values allowed?"
- Edge cases: empty tree, single node, extreme values
- Time complexity: O(N) vs O(N²) naive
- Space complexity: O(H) for recursion vs O(1) iterative

## 7. COMMON MISTAKES
- Not handling null nodes properly
- Incorrect boundary conditions (<= vs <)
- Integer overflow with Integer.MIN/MAX_VALUE
- Not updating bounds correctly in recursion

## 8. OPTIMIZATION STRATEGIES
- Use long for bounds to avoid overflow
- Iterative approach for very deep trees
- Early termination on obvious violations
- Inorder traversal for sorted sequence check

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like validating a family tree:**
- You have a family tree where each person has descendants
- BST rule: All left descendants must be younger, right descendants older
- For each person, check if this rule holds
- Pass down age constraints to descendants
- If any violation found, tree is invalid

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Root of binary tree
2. **Goal**: Check if it's a valid BST
3. **Output**: Boolean indicating validity

#### Phase 2: Key Insight Recognition
- **"What is BST property?"** → left < node < right for all nodes
- **"How to validate efficiently?"** → Pass bounds down recursively
- **"What bounds to pass?"** → min, node, max for each subtree
- **"Why multiple approaches?"** → Different tradeoffs for different scenarios

#### Phase 3: Strategy Development
```
Human thought process:
"I'll validate using range bounds:
1. Start with root, bounds = (-∞, +∞)
2. For left child: bounds = (min, node.val)
3. For right child: bounds = (node.val, max)
4. Recursively validate subtrees
5. If any violation, return false"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return true (no violations)
- **Single node**: Always valid BST
- **Extreme values**: Use long bounds to avoid overflow
- **Unbalanced tree**: Still valid if BST property holds

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:
    5
   / \
  1   9
 / \ / \
3   4   7

Human thinking:
"Let's validate this tree step by step:

Root node: 5
- Bounds: (-∞, +∞)
- Check: 5 is within bounds ✓

Left child: 1
- Bounds: (-∞, 5)
- Check: 1 < 5 ✓
- Check: 1 > -∞ ✓

Right child: 9
- Bounds: (5, +∞)
- Check: 9 > 5 ✓
- Check: 9 < +∞ ✓

Left subtree (node 1):
- Left child: 3
- Bounds: (-∞, 1)
- Check: 3 < 1 ✓
- Check: 3 > -∞ ✓

Right child: 4
- Bounds: (1, +∞)
- Check: 4 > 1 ✗ VIOLATION!
- 4 should be < 1, but 4 > 1

Found violation: tree is invalid ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Range bounds ensure BST property at each level
- **Why it's efficient**: Each node visited once
- **Why it's correct**: All possible violations are checked

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just check immediate children?"** → Only checks local, not global
2. **"What about inorder traversal?"** → Works but O(N) space
3. **"How to handle bounds?"** → Use long to avoid overflow
4. **"What about duplicates?"** → Need to clarify allowed positions

### Real-World Analogy
**Like validating an organizational hierarchy:**
- You have a company hierarchy (tree structure)
- BST rule: Managers must have higher rank than subordinates
- For each employee, check if this rule holds throughout their team
- Pass down rank constraints to their subordinates
- If any violation found, hierarchy is invalid
- This ensures proper organizational structure

### Human-Readable Pseudocode
```
function isValidBST(root):
    return validateBST(root, -∞, +∞)
    
function validateBST(node, min, max):
    if node == null:
        return true
    
    // Check current node against bounds
    if node.val <= min or node.val >= max:
        return false
    
    // Recursively validate subtrees
    return validateBST(node.left, min, node.val) &&
           validateBST(node.right, node.val, max)
```

### Execution Visualization

### Example: Invalid BST
```
Tree:
    5
   / \
  1   9
 / \ / \
3   4   7

Range Validation Process:
Root(5): bounds=(-∞,+∞) ✓
Left(1): bounds=(-∞,5) ✓
Right(9): bounds=(5,+∞) ✓

Left subtree:
Node(1): bounds=(-∞,1) ✓
Left(3): bounds=(-∞,1) ✓
Right(4): bounds=(1,+∞)
VIOLATION: 4 > 1 ✗

Result: Invalid BST ✓

Visualization:
    5
   / \
  1   9
 / \ / \
3   4   7  ← 4 violates BST property
```

### Key Visualization Points:
- **Range bounds** ensure BST property at each level
- **Recursive validation** checks all nodes
- **Boundary conditions** prevent overflow
- **Multiple approaches** for different scenarios

### Memory Layout Visualization:
```
Recursive Call Stack:
validateBST(root, -∞, +∞)
├─ validateBST(node=5, -∞, +∞)
│  ├─ validateBST(node=1, -∞, 5)
│  │  ├─ validateBST(node=3, -∞, 1)
│  │  └─ validateBST(null, -∞, 3)
│  └─ validateBST(node=4, 1, +∞)
│     └─ VIOLATION: 4 > 1
└─ Return false

Boundaries Passed:
Root: (-∞, +∞)
Node 1: (-∞, 5)
Node 3: (-∞, 1)
Node 4: (1, +∞) ← VIOLATION
```

### Time Complexity Breakdown:
- **Each Node**: Visited exactly once
- **Validation**: O(1) per node (comparisons)
- **Total**: O(N) time where N is number of nodes
- **Space**: O(H) recursion stack where H is tree height
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Inorder**: O(N) time but O(N) extra space
*/
