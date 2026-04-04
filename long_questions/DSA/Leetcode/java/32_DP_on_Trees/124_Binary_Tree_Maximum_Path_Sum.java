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
}
