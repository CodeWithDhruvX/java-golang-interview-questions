import java.util.*;

public class InvertBinaryTree {
    
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

    // 226. Invert Binary Tree
    // Time: O(N), Space: O(H) where H is tree height (recursion stack)
    public static TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return null;
        }
        
        // Swap left and right children
        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;
        
        // Recursively invert subtrees
        invertTree(root.left);
        invertTree(root.right);
        
        return root;
    }

    // Helper method to print tree (level order)
    public static List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) return result;
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<Integer> level = new ArrayList<>();
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                level.add(node.val);
                
                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
            
            result.add(level);
        }
        
        return result;
    }

    // Helper method to create tree from array
    public static TreeNode createTree(Integer[] arr) {
        if (arr == null || arr.length == 0 || arr[0] == null) {
            return null;
        }
        
        TreeNode root = new TreeNode(arr[0]);
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        int i = 1;
        while (!queue.isEmpty() && i < arr.length) {
            TreeNode node = queue.poll();
            
            if (i < arr.length && arr[i] != null) {
                node.left = new TreeNode(arr[i]);
                queue.offer(node.left);
            }
            i++;
            
            if (i < arr.length && arr[i] != null) {
                node.right = new TreeNode(arr[i]);
                queue.offer(node.right);
            }
            i++;
        }
        
        return root;
    }

    public static void main(String[] args) {
        Integer[][] testCases = {
            {4, 2, 7, 1, 3, 6, 9},
            {2, 1, 3},
            {null},
            {1},
            {1, 2, 3, 4, 5, 6, 7},
            {1, null, 2, null, 3},
            {1, 2, null, 3},
            {1, 2, 3, null, null, 4, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode root = createTree(testCases[i]);
            List<List<Integer>> original = levelOrder(root);
            
            invertTree(root);
            List<List<Integer>> inverted = levelOrder(root);
            
            System.out.printf("Test Case %d:\n", i + 1);
            System.out.printf("Original: %s\n", original);
            System.out.printf("Inverted: %s\n\n", inverted);
        }
    }
}
