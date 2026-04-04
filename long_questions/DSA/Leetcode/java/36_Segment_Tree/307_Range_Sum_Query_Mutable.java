public class RangeSumQueryMutable {
    
    // 307. Range Sum Query - Mutable - Segment Tree
    // Time: O(log N) for update and query, Space: O(N)
    public static class NumArray {
        private int[] tree;
        private int n;
        
        // Constructor builds segment tree from array
        public NumArray(int[] nums) {
            n = nums.length;
            tree = new int[2 * n];
            
            // Build leaf nodes
            for (int i = 0; i < n; i++) {
                tree[i + n] = nums[i];
            }
            
            // Build internal nodes
            for (int i = n - 1; i > 0; i--) {
                tree[i] = tree[2 * i] + tree[2 * i + 1];
            }
        }
        
        // Update value at index i
        public void update(int i, int val) {
            // Update leaf node
            int pos = i + n;
            tree[pos] = val;
            
            // Update internal nodes
            pos /= 2;
            while (pos > 0) {
                tree[pos] = tree[2 * pos] + tree[2 * pos + 1];
                pos /= 2;
            }
        }
        
        // Sum range query
        public int sumRange(int left, int right) {
            // Convert to leaf positions
            left += n;
            right += n;
            
            int sum = 0;
            
            // Query from both ends
            while (left <= right) {
                if (left % 2 == 1) {
                    sum += tree[left];
                    left++;
                }
                if (right % 2 == 0) {
                    sum += tree[right];
                    right--;
                }
                left /= 2;
                right /= 2;
            }
            
            return sum;
        }
    }

    // Alternative implementation with recursive segment tree
    public static class NumArrayRecursive {
        private int[] tree;
        private int n;
        
        public NumArrayRecursive(int[] nums) {
            n = nums.length;
            tree = new int[4 * n]; // 4*n is safe upper bound
            buildTree(nums, tree, 1, 0, n - 1);
        }
        
        private void buildTree(int[] nums, int[] tree, int node, int start, int end) {
            if (start == end) {
                tree[node] = nums[start];
            } else {
                int mid = start + (end - start) / 2;
                buildTree(nums, tree, 2 * node, start, mid);
                buildTree(nums, tree, 2 * node + 1, mid + 1, end);
                tree[node] = tree[2 * node] + tree[2 * node + 1];
            }
        }
        
        public void update(int i, int val) {
            updateTree(1, 0, n - 1, i, val);
        }
        
        private void updateTree(int node, int start, int end, int idx, int val) {
            if (start == end) {
                tree[node] = val;
            } else {
                int mid = start + (end - start) / 2;
                if (idx <= mid) {
                    updateTree(2 * node, start, mid, idx, val);
                } else {
                    updateTree(2 * node + 1, mid + 1, end, idx, val);
                }
                tree[node] = tree[2 * node] + tree[2 * node + 1];
            }
        }
        
        public int sumRange(int left, int right) {
            return queryTree(1, 0, n - 1, left, right);
        }
        
        private int queryTree(int node, int start, int end, int left, int right) {
            if (right < start || end < left) {
                return 0; // No overlap
            }
            if (left <= start && end <= right) {
                return tree[node]; // Total overlap
            }
            
            int mid = start + (end - start) / 2;
            int leftSum = queryTree(2 * node, start, mid, left, right);
            int rightSum = queryTree(2 * node + 1, mid + 1, end, left, right);
            return leftSum + rightSum;
        }
    }

    // Alternative implementation using Binary Indexed Tree (Fenwick Tree)
    public static class NumArrayBIT {
        private int[] bit;
        private int n;
        
        public NumArrayBIT(int[] nums) {
            n = nums.length;
            bit = new int[n + 1];
            
            // Build BIT
            for (int i = 0; i < n; i++) {
                update(i, nums[i]);
            }
        }
        
        public void update(int i, int val) {
            int current = sumRange(i, i);
            int diff = val - current;
            
            i++; // BIT is 1-indexed
            while (i <= n) {
                bit[i] += diff;
                i += i & (-i);
            }
        }
        
        public int sumRange(int left, int right) {
            return prefixSum(right) - prefixSum(left - 1);
        }
        
        private int prefixSum(int i) {
            int sum = 0;
            i++; // BIT is 1-indexed
            while (i > 0) {
                sum += bit[i];
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
            
            // Test iterative segment tree
            NumArray numArray = new NumArray(testArrays[i].clone());
            int sum1 = numArray.sumRange(0, testArrays[i].length - 1);
            System.out.printf("  Iterative - Sum(0,%d): %d\n", testArrays[i].length - 1, sum1);
            
            if (testArrays[i].length > 1) {
                numArray.update(1, 10);
                int sum2 = numArray.sumRange(0, testArrays[i].length - 1);
                System.out.printf("  After update(1,10) - Sum(0,%d): %d\n", testArrays[i].length - 1, sum2);
            }
            
            // Test recursive segment tree
            NumArrayRecursive numArrayRec = new NumArrayRecursive(testArrays[i].clone());
            int sum3 = numArrayRec.sumRange(0, testArrays[i].length - 1);
            System.out.printf("  Recursive - Sum(0,%d): %d\n", testArrays[i].length - 1, sum3);
            
            // Test BIT
            NumArrayBIT numArrayBIT = new NumArrayBIT(testArrays[i].clone());
            int sum4 = numArrayBIT.sumRange(0, testArrays[i].length - 1);
            System.out.printf("  BIT - Sum(0,%d): %d\n", testArrays[i].length - 1, sum4);
            
            System.out.println();
        }
    }
}
