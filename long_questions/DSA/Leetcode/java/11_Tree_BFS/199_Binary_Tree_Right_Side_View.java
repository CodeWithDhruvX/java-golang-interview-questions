import java.util.*;

public class BinaryTreeRightSideView {
    
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

    // 199. Binary Tree Right Side View
    // Time: O(N), Space: O(W) where W is maximum width of tree
    public static List<Integer> rightSideView(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        if (root == null) {
            return result;
        }
        
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                
                // Add the rightmost node of each level
                if (i == levelSize - 1) {
                    result.add(node.val);
                }
                
                if (node.left != null) queue.offer(node.left);
                if (node.right != null) queue.offer(node.right);
            }
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
            {1, 2, 3, null, 5, null, 4},
            {1, null, 3},
            {null},
            {1},
            {1, 2, 3, 4, 5, 6, 7},
            {1, 2, null, 3, null, 4},
            {1, 2, 3, 4, null, null, 5},
            {1, 2, 3, null, 4, null, 5},
            {1, 2, 3, 4, 5, 6, 7, 8},
            {1, 2, 3, null, null, 4, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode root = createTree(testCases[i]);
            List<Integer> result = rightSideView(root);
            
            System.out.printf("Test Case %d: Tree=%s -> Right Side View: %s\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
