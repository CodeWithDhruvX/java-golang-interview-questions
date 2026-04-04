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
}
