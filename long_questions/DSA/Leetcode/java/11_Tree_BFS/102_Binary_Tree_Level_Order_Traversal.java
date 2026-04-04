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
}
