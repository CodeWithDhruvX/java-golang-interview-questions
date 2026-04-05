import java.util.*;

public class BinaryTreeZigzagLevelOrderTraversal {
    
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

    // 103. Binary Tree Zigzag Level Order Traversal
    // Time: O(N), Space: O(W) where W is maximum width of tree
    public static List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) {
            return result;
        }
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        boolean leftToRight = true;
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<Integer> level = new ArrayList<>();
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                
                if (leftToRight) {
                    level.add(node.val);
                } else {
                    level.add(0, node.val); // Add to front for reverse order
                }
                
                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
            
            result.add(level);
            leftToRight = !leftToRight;
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
            {3, 9, 20, null, null, 15, 7},
            {1},
            {null},
            {1, 2, 3, 4, 5, 6, 7},
            {1, 2, 3, null, null, 4, 5},
            {1, 2, 3, 4, null, null, 5, 6},
            {1, null, 2, null, 3},
            {1, 2, null, 3, null, 4},
            {1, 2, 3, 4, 5, 6, 7, 8, 9},
            {1, 2, 3, 4, 5, null, null, 6}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode root = createTree(testCases[i]);
            List<List<Integer>> result = zigzagLevelOrder(root);
            
            System.out.printf("Test Case %d: Tree=%s -> Zigzag Level Order: %s\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
