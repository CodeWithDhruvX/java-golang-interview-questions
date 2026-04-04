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
}
