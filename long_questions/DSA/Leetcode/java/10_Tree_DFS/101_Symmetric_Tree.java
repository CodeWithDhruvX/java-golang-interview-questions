import java.util.*;

public class SymmetricTree {
    
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

    // 101. Symmetric Tree
    // Time: O(N), Space: O(H) where H is tree height (recursion stack)
    public static boolean isSymmetric(TreeNode root) {
        if (root == null) {
            return true;
        }
        return isMirror(root.left, root.right);
    }
    
    private static boolean isMirror(TreeNode left, TreeNode right) {
        // Both null -> mirror
        if (left == null && right == null) {
            return true;
        }
        
        // One null, one not null -> not mirror
        if (left == null || right == null) {
            return false;
        }
        
        // Both not null, check values and subtrees
        return left.val == right.val && 
               isMirror(left.left, right.right) && 
               isMirror(left.right, right.left);
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
            {1, 2, 2, 3, 4, 4, 3},
            {1, 2, 2, null, 3, null, 3},
            {1},
            {null},
            {1, 2, 2, 3, null, null, 3},
            {1, 2, 2, null, 3, 3, null},
            {1, 2, 3},
            {1, null, 1},
            {1, 2, 2, 3, 4, 4, 3, 5, 6, 6, 5},
            {1, 2, 2, null, null, 2, 2}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode root = createTree(testCases[i]);
            
            boolean result = isSymmetric(root);
            System.out.printf("Test Case %d: Tree=%s -> Symmetric: %b\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
