public class RangeSumQueryMutableFenwickTree {
    
    // 307. Range Sum Query - Mutable - Fenwick Tree (Binary Indexed Tree)
    // Time: O(log N) for update and query, Space: O(N)
    public static class NumArray {
        private int[] tree;
        private int n;
        
        // Constructor builds Fenwick Tree from array
        public NumArray(int[] nums) {
            n = nums.length;
            tree = new int[n + 1]; // 1-based indexing
            
            // Build Fenwick Tree
            for (int i = 0; i < n; i++) {
                updateFenwick(i + 1, nums[i]);
            }
        }
        
        // Update value at index i (0-based)
        public void update(int i, int val) {
            // Calculate the difference
            int delta = val - sumRange(i, i);
            // Update the Fenwick tree
            updateFenwick(i + 1, delta);
        }
        
        // Query sum from left to right (0-based, inclusive)
        public int sumRange(int left, int right) {
            return queryFenwick(right + 1) - queryFenwick(left);
        }
        
        // Update Fenwick tree at position i (1-based) by delta
        private void updateFenwick(int i, int delta) {
            while (i <= n) {
                tree[i] += delta;
                i += i & (-i); // Move to parent
            }
        }
        
        // Query prefix sum from 1 to i (1-based)
        private int queryFenwick(int i) {
            int sum = 0;
            while (i > 0) {
                sum += tree[i];
                i -= i & (-i); // Move to parent
            }
            return sum;
        }
    }

    // Alternative implementation with range updates and point queries
    public static class NumArrayRangeUpdate {
        private int[] tree;
        private int[] original;
        private int n;
        
        public NumArrayRangeUpdate(int[] nums) {
            n = nums.length;
            tree = new int[n + 1];
            original = nums.clone();
            
            // Initialize with differences
            for (int i = 0; i < n; i++) {
                if (i == 0) {
                    updateFenwick(i + 1, nums[i]);
                } else {
                    updateFenwick(i + 1, nums[i] - nums[i - 1]);
                }
            }
        }
        
        public void updateRange(int left, int right, int val) {
            updateFenwick(left + 1, val);
            if (right + 1 < n) {
                updateFenwick(right + 2, -val);
            }
        }
        
        public int get(int i) {
            return original[i] + queryFenwick(i + 1);
        }
        
        private void updateFenwick(int i, int delta) {
            while (i <= n) {
                tree[i] += delta;
                i += i & (-i);
            }
        }
        
        private int queryFenwick(int i) {
            int sum = 0;
            while (i > 0) {
                sum += tree[i];
                i -= i & (-i);
            }
            return sum;
        }
    }

    // Range update, range query Fenwick Tree
    public static class NumArrayRangeUpdateRangeQuery {
        private int[] tree1; // For range updates
        private int[] tree2; // For range queries
        private int n;
        
        public NumArrayRangeUpdateRangeQuery(int[] nums) {
            n = nums.length;
            tree1 = new int[n + 1];
            tree2 = new int[n + 1];
            
            // Initialize with differences
            for (int i = 0; i < n; i++) {
                if (i == 0) {
                    update(tree1, i + 1, nums[i]);
                    update(tree2, i + 1, nums[i]);
                } else {
                    update(tree1, i + 1, nums[i] - nums[i - 1]);
                    update(tree2, i + 1, i * (nums[i] - nums[i - 1]));
                }
            }
        }
        
        public void updateRange(int left, int right, int val) {
            update(tree1, left + 1, val);
            update(tree1, right + 2, -val);
            update(tree2, left + 1, val * left);
            update(tree2, right + 2, -val * (right + 1));
        }
        
        public int sumRange(int left, int right) {
            return prefixSum(right) - prefixSum(left - 1);
        }
        
        private int prefixSum(int i) {
            return query(tree1, i + 1) * (i + 1) - query(tree2, i + 1);
        }
        
        private void update(int[] tree, int i, int delta) {
            while (i <= n) {
                tree[i] += delta;
                i += i & (-i);
            }
        }
        
        private int query(int[] tree, int i) {
            int sum = 0;
            while (i > 0) {
                sum += tree[i];
                i -= i & (-i);
            }
            return sum;
        }
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {1, 3, 5},
            {0, 9, 5, 7, 3},
            {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            {5, 5, 5, 5, 5},
            {100, 200, 300},
            {1},
            {}
        };
        
        for (int i = 0; i < testArrays.length; i++) {
            if (testArrays[i].length == 0) {
                System.out.printf("Test Case %d: Empty array\n", i + 1);
                continue;
            }
            
            System.out.printf("Test Case %d: %s\n", i + 1, java.util.Arrays.toString(testArrays[i]));
            
            // Test basic Fenwick Tree
            NumArray numArray = new NumArray(testArrays[i].clone());
            int sum1 = numArray.sumRange(0, testArrays[i].length - 1);
            System.out.printf("  Basic BIT - Sum(0,%d): %d\n", testArrays[i].length - 1, sum1);
            
            if (testArrays[i].length > 1) {
                numArray.update(1, 10);
                int sum2 = numArray.sumRange(0, testArrays[i].length - 1);
                System.out.printf("  After update(1,10) - Sum(0,%d): %d\n", testArrays[i].length - 1, sum2);
            }
            
            // Test range update version
            NumArrayRangeUpdate rangeUpdate = new NumArrayRangeUpdate(testArrays[i].clone());
            if (testArrays[i].length > 2) {
                rangeUpdate.updateRange(1, 3, 5);
                System.out.printf("  After range update(1,3,5) - get(2): %d\n", rangeUpdate.get(2));
            }
            
            // Test range update, range query version
            NumArrayRangeUpdateRangeQuery advanced = new NumArrayRangeUpdateRangeQuery(testArrays[i].clone());
            if (testArrays[i].length > 3) {
                advanced.updateRange(1, 2, 3);
                int sum3 = advanced.sumRange(0, 3);
                System.out.printf("  Advanced - After range update(1,2,3) - Sum(0,3): %d\n", sum3);
            }
            
            System.out.println();
        }
    }
}
