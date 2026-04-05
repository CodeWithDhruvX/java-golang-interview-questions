import java.util.*;

public class SameTree {
    
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

    // 100. Same Tree
    // Time: O(N), Space: O(H) where H is tree height (recursion stack)
    public static boolean isSameTree(TreeNode p, TreeNode q) {
        // Both null -> same
        if (p == null && q == null) {
            return true;
        }
        
        // One null, one not null -> different
        if (p == null || q == null) {
            return false;
        }
        
        // Both not null, check values and subtrees
        return p.val == q.val && 
               isSameTree(p.left, q.left) && 
               isSameTree(p.right, q.right);
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
        Object[][] testCases = {
            {new Integer[]{1, 2, 3}, new Integer[]{1, 2, 3}},
            {new Integer[]{1, 2}, new Integer[]{1, null, 2}},
            {new Integer[]{1, 2, 1}, new Integer[]{1, 1, 2}},
            {new Integer[]{}, new Integer[]{}},
            {new Integer[]{1}, new Integer[]{1}},
            {new Integer[]{1}, new Integer[]{2}},
            {new Integer[]{1, 2, 3, 4, 5}, new Integer[]{1, 2, 3, 4, 5}},
            {new Integer[]{1, null, 2, null, 3}, new Integer[]{1, null, 2, null, 3}},
            {new Integer[]{1, 2, null, 3}, new Integer[]{1, 2, null, 3}},
            {new Integer[]{1, 2, 3}, new Integer[]{1, 3, 2}}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            TreeNode p = createTree((Integer[]) testCases[i][0]);
            TreeNode q = createTree((Integer[]) testCases[i][1]);
            
            boolean result = isSameTree(p, q);
            System.out.printf("Test Case %d: Tree1=%s, Tree2=%s -> Same: %b\n", 
                i + 1, Arrays.toString((Integer[]) testCases[i][0]), 
                Arrays.toString((Integer[]) testCases[i][1]), result);
        }
    }
}
