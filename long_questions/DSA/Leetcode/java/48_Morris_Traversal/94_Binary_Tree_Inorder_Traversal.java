import java.util.*;

public class BinaryTreeInorderTraversal {
    
    // 94. Binary Tree Inorder Traversal - Morris Traversal
    // Time: O(N), Space: O(1)
    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        
        TreeNode(int val) {
            this.val = val;
        }
    }
    
    // Morris Traversal for inorder
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        
        while (current != null) {
            if (current.left == null) {
                // No left child, visit current and move to right
                result.add(current.val);
                current = current.right;
            } else {
                // Find predecessor
                TreeNode predecessor = current.left;
                while (predecessor.right != null && predecessor.right != current) {
                    predecessor = predecessor.right;
                }
                
                if (predecessor.right == null) {
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Revert the changes and visit current
                    predecessor.right = null;
                    result.add(current.val);
                    current = current.right;
                }
            }
        }
        
        return result;
    }
    
    // Morris Traversal for preorder
    public List<Integer> preorderTraversalMorris(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        
        while (current != null) {
            if (current.left == null) {
                // No left child, visit current and move to right
                result.add(current.val);
                current = current.right;
            } else {
                // Find predecessor
                TreeNode predecessor = current.left;
                while (predecessor.right != null && predecessor.right != current) {
                    predecessor = predecessor.right;
                }
                
                if (predecessor.right == null) {
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Revert the changes and visit current
                    predecessor.right = null;
                    result.add(current.val); // Visit before going left
                    current = current.right;
                }
            }
        }
        
        return result;
    }
    
    // Morris Traversal for postorder
    public List<Integer> postorderTraversalMorris(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        TreeNode dummy = new TreeNode(0);
        dummy.left = root;
        current = dummy;
        
        while (current != null) {
            if (current.left == null) {
                // No left child, move to right
                current = current.right;
            } else {
                // Find predecessor
                TreeNode predecessor = current.left;
                while (predecessor.right != null && predecessor.right != current) {
                    predecessor = predecessor.right;
                }
                
                if (predecessor.right == null) {
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                } else {
                    // Process current and revert changes
                    TreeNode temp = predecessor.right;
                    predecessor.right = null;
                    current = temp;
                }
            }
        }
        
        // Reverse the result for postorder
        Collections.reverse(result);
        return result;
    }
    
    // Standard recursive inorder for comparison
    public List<Integer> inorderTraversalRecursive(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        inorderHelper(root, result);
        return result;
    }
    
    private void inorderHelper(TreeNode node, List<Integer> result) {
        if (node == null) {
            return;
        }
        
        inorderHelper(node.left, result);
        result.add(node.val);
        inorderHelper(node.right, result);
    }
    
    // Iterative inorder using stack for comparison
    public List<Integer> inorderTraversalIterative(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        Deque<TreeNode> stack = new ArrayDeque<>();
        TreeNode current = root;
        
        while (current != null || !stack.isEmpty()) {
            while (current != null) {
                stack.push(current);
                current = current.left;
            }
            
            current = stack.pop();
            result.add(current.val);
            current = current.right;
        }
        
        return result;
    }
    
    // Version with detailed explanation
    public class MorrisTraversalResult {
        List<Integer> result;
        List<String> explanation;
        
        MorrisTraversalResult(List<Integer> result, List<String> explanation) {
            this.result = result;
            this.explanation = explanation;
        }
    }
    
    public MorrisTraversalResult inorderTraversalDetailed(TreeNode root) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== Morris Traversal for Inorder ===");
        explanation.add("Using O(1) space algorithm without recursion or stack");
        
        List<Integer> result = new ArrayList<>();
        TreeNode current = root;
        int step = 1;
        
        while (current != null) {
            explanation.add(String.format("Step %d: Current node = %d", step++, current.val));
            
            if (current.left == null) {
                explanation.add("  No left child - visiting current and moving right");
                result.add(current.val);
                explanation.add(String.format("  Added %d to result", current.val));
                current = current.right;
                explanation.add(String.format("  Moved to right child: %s", 
                    current.right != null ? String.valueOf(current.right.val) : "null"));
            } else {
                explanation.add("  Finding predecessor in left subtree");
                
                // Find predecessor
                TreeNode predecessor = current.left;
                int predecessorSteps = 0;
                
                while (predecessor.right != null && predecessor.right != current) {
                    predecessorSteps++;
                    predecessor = predecessor.right;
                    explanation.add(String.format("    Predecessor step %d: moved to %d", 
                        predecessorSteps, predecessor.val));
                }
                
                if (predecessor.right == null) {
                    explanation.add("  Predecessor.right is null - creating temporary link");
                    explanation.add(String.format("  Making %d the right child of %d", 
                        current.val, predecessor.val));
                    
                    // Make current the right child of predecessor
                    predecessor.right = current;
                    current = current.left;
                    explanation.add(String.format("  Moved to left child: %d", current.val));
                } else {
                    explanation.add("  Predecessor.right is not null - visiting current and moving right");
                    explanation.add(String.format("  Reverting temporary link, visiting %d", current.val));
                    
                    // Revert the changes and visit current
                    predecessor.right = null;
                    result.add(current.val);
                    explanation.add(String.format("  Added %d to result", current.val));
                    current = current.right;
                    explanation.add(String.format("  Moved to right child: %s", 
                        current.right != null ? String.valueOf(current.right.val) : "null"));
                }
            }
        }
        
        return new MorrisTraversalResult(result, explanation);
    }
    
    // Performance comparison
    public void comparePerformance(TreeNode root, int trials) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Trials: " + trials);
        
        // Morris Traversal
        long startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            inorderTraversal(root);
        }
        long endTime = System.nanoTime();
        System.out.printf("Morris Traversal: took %d ns\n", endTime - startTime);
        
        // Recursive Traversal
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            inorderTraversalRecursive(root);
        }
        endTime = System.nanoTime();
        System.out.printf("Recursive Traversal: took %d ns\n", endTime - startTime);
        
        // Iterative Traversal
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            inorderTraversalIterative(root);
        }
        endTime = System.nanoTime();
        System.out.printf("Iterative Traversal: took %d ns\n", endTime - startTime);
    }
    
    // Memory usage analysis
    public void memoryUsageAnalysis(TreeNode root) {
        System.out.println("=== Memory Usage Analysis ===");
        
        // Morris Traversal - O(1) space
        long morrisStartTime = System.nanoTime();
        List<Integer> morrisResult = inorderTraversal(root);
        long morrisEndTime = System.nanoTime();
        
        // Recursive Traversal - O(H) space (recursion stack)
        long recursiveStartTime = System.nanoTime();
        List<Integer> recursiveResult = inorderTraversalRecursive(root);
        long recursiveEndTime = System.nanoTime();
        
        // Iterative Traversal - O(H) space (explicit stack)
        long iterativeStartTime = System.nanoTime();
        List<Integer> iterativeResult = inorderTraversalIterative(root);
        long iterativeEndTime = System.nanoTime();
        
        System.out.println("Results verification:");
        System.out.println("Morris: " + morrisResult);
        System.out.println("Recursive: " + recursiveResult);
        System.out.println("Iterative: " + iterativeResult);
        
        System.out.println("\nPerformance comparison:");
        System.out.printf("Morris: %d ns\n", morrisEndTime - morrisStartTime);
        System.out.printf("Recursive: %d ns\n", recursiveEndTime - recursiveStartTime);
        System.out.printf("Iterative: %d ns\n", iterativeEndTime - iterativeStartTime);
        
        System.out.println("\nSpace complexity:");
        System.out.println("Morris: O(1) - no recursion stack or explicit stack");
        System.out.println("Recursive: O(H) - recursion stack depth");
        System.out.println("Iterative: O(H) - explicit stack size");
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
    
    // Helper method to print tree
    public void printTree(TreeNode root) {
        if (root == null) {
            System.out.println("Empty tree");
            return;
        }
        
        List<List<String>> levels = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            List<String> level = new ArrayList<>();
            
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                if (node != null) {
                    level.add(String.valueOf(node.val));
                    queue.offer(node.left);
                    queue.offer(node.right);
                } else {
                    level.add("null");
                    queue.offer(null);
                    queue.offer(null);
                }
            }
            
            levels.add(level);
        }
        
        for (int i = 0; i < levels.size(); i++) {
            System.out.println("Level " + i + ": " + levels.get(i));
        }
    }
    
    public static void main(String[] args) {
        BinaryTreeInorderTraversal bt = new BinaryTreeInorderTraversal();
        
        // Test cases
        Integer[][] testCases = {
            {1, null, 2, 3},
            {1, 2, 3, 4, 5},
            {1},
            {1, null, 2, null, 3},
            {5, 3, 1, null, null, null, 4, 6},
            {1, 2, null, 3},
            {10, 5, 15, null, null, 20},
            {1, null, null, 2},
            {1, 2, 3, 4, 5, 6, 7}
        };
        
        String[] descriptions = {
            "Simple tree",
            "Complete tree",
            "Single node",
            "Sparse tree",
            "Unbalanced tree",
            "Missing nodes",
            "Large values",
            "Left heavy",
            "Perfect tree"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Tree structure: " + Arrays.toString(testCases[i]));
            
            TreeNode root = bt.createTree(testCases[i]);
            bt.printTree(root);
            
            // Test all traversal methods
            List<Integer> morrisResult = bt.inorderTraversal(root);
            List<Integer> recursiveResult = bt.inorderTraversalRecursive(root);
            List<Integer> iterativeResult = bt.inorderTraversalIterative(root);
            
            System.out.println("Traversal results:");
            System.out.println("Morris: " + morrisResult);
            System.out.println("Recursive: " + recursiveResult);
            System.out.println("Iterative: " + iterativeResult);
            
            // Test preorder
            List<Integer> preorderResult = bt.preorderTraversalMorris(root);
            System.out.println("Preorder Morris: " + preorderResult);
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        TreeNode detailedRoot = bt.createTree(new Integer[]{4, 2, 5, 1, 3});
        MorrisTraversalResult detailedResult = bt.inorderTraversalDetailed(detailedRoot);
        
        System.out.println("Result: " + detailedResult.result);
        for (String step : detailedResult.explanation) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        TreeNode performanceRoot = bt.createTree(new Integer[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10});
        bt.comparePerformance(performanceRoot, 10000);
        
        // Memory usage analysis
        System.out.println("\n=== Memory Usage Analysis ===");
        TreeNode analysisRoot = bt.createTree(new Integer[]{1, 2, 3, 4, 5, 6, 7});
        bt.memoryUsageAnalysis(analysisRoot);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty tree
        List<Integer> emptyResult = bt.inorderTraversal(null);
        System.out.println("Empty tree: " + emptyResult);
        
        // Single node
        TreeNode singleNode = new TreeNode(42);
        List<Integer> singleResult = bt.inorderTraversal(singleNode);
        System.out.println("Single node: " + singleResult);
        
        // Large tree
        Integer[] largeValues = new Integer[1000];
        for (int i = 0; i < 1000; i++) {
            largeValues[i] = i + 1;
        }
        
        TreeNode largeRoot = bt.createTree(largeValues);
        long startTime = System.nanoTime();
        bt.inorderTraversal(largeRoot);
        long endTime = System.nanoTime();
        System.out.printf("Large tree (1000 nodes): took %d ns\n", endTime - startTime);
    }
}
