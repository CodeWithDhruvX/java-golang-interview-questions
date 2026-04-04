import java.util.*;

public class LongestConsecutiveSequence {
    
    // 128. Longest Consecutive Sequence
    // Time: O(N), Space: O(N)
    public static int longestConsecutive(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        Set<Integer> numSet = new HashSet<>();
        for (int num : nums) {
            numSet.add(num);
        }
        
        int maxLength = 0;
        
        for (int num : numSet) {
            // Only start counting from the beginning of a sequence
            if (!numSet.contains(num - 1)) {
                int currentNum = num;
                int currentLength = 1;
                
                // Count the length of consecutive sequence
                while (numSet.contains(currentNum + 1)) {
                    currentNum++;
                    currentLength++;
                }
                
                maxLength = Math.max(maxLength, currentLength);
            }
        }
        
        return maxLength;
    }

    // Union-Find approach
    public static int longestConsecutiveUnionFind(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        // Initialize Union-Find
        Map<Integer, Integer> parent = new HashMap<>();
        Map<Integer, Integer> rank = new HashMap<>();
        
        for (int num : nums) {
            parent.put(num, num);
            rank.put(num, 0);
        }
        
        // Find function with path compression
        class UnionFind {
            int find(int x) {
                if (parent.get(x) != x) {
                    parent.put(x, find(parent.get(x)));
                }
                return parent.get(x);
            }
            
            void union(int x, int y) {
                int rootX = find(x);
                int rootY = find(y);
                
                if (rootX == rootY) {
                    return;
                }
                
                if (rank.get(rootX) < rank.get(rootY)) {
                    parent.put(rootX, rootY);
                } else if (rank.get(rootX) > rank.get(rootY)) {
                    parent.put(rootY, rootX);
                } else {
                    parent.put(rootY, rootX);
                    rank.put(rootX, rank.get(rootX) + 1);
                }
            }
        }
        
        UnionFind uf = new UnionFind();
        
        // Union consecutive numbers
        for (int num : nums) {
            if (parent.containsKey(num - 1)) {
                uf.union(num, num - 1);
            }
            if (parent.containsKey(num + 1)) {
                uf.union(num, num + 1);
            }
        }
        
        // Count the size of each component
        Map<Integer, Integer> componentSize = new HashMap<>();
        for (int num : nums) {
            int root = uf.find(num);
            componentSize.put(root, componentSize.getOrDefault(root, 0) + 1);
        }
        
        int maxLength = 0;
        for (int size : componentSize.values()) {
            maxLength = Math.max(maxLength, size);
        }
        
        return maxLength;
    }

    // Sorting approach (O(N log N))
    public static int longestConsecutiveSorting(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        // Remove duplicates and sort
        Set<Integer> unique = new HashSet<>();
        for (int num : nums) {
            unique.add(num);
        }
        
        int[] sorted = new int[unique.size()];
        int index = 0;
        for (int num : unique) {
            sorted[index++] = num;
        }
        
        // Simple bubble sort for demonstration (in practice, use Arrays.sort)
        for (int i = 0; i < sorted.length - 1; i++) {
            for (int j = 0; j < sorted.length - i - 1; j++) {
                if (sorted[j] > sorted[j + 1]) {
                    int temp = sorted[j];
                    sorted[j] = sorted[j + 1];
                    sorted[j + 1] = temp;
                }
            }
        }
        
        int maxLength = 1;
        int currentLength = 1;
        
        for (int i = 1; i < sorted.length; i++) {
            if (sorted[i] == sorted[i - 1] + 1) {
                currentLength++;
            } else {
                maxLength = Math.max(maxLength, currentLength);
                currentLength = 1;
            }
        }
        
        maxLength = Math.max(maxLength, currentLength);
        
        return maxLength;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {100, 4, 200, 1, 3, 2},
            {0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
            {},
            {1},
            {1, 2, 0, 1},
            {9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
            {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
            {1, 3, 5, 7, 9},
            {-1, -2, -3, 0, 1, 2, 3}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result1 = longestConsecutive(testCases[i]);
            int result2 = longestConsecutiveUnionFind(testCases[i]);
            int result3 = longestConsecutiveSorting(testCases[i]);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  HashSet: %d, Union-Find: %d, Sorting: %d\n\n", 
                result1, result2, result3);
        }
    }
}
