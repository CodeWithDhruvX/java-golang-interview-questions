import java.util.*;

public class SpiralMatrix {
    
    // 54. Spiral Matrix
    // Time: O(M*N), Space: O(1) (excluding output)
    public static List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return result;
        }
        
        int m = matrix.length, n = matrix[0].length;
        int top = 0, bottom = m - 1;
        int left = 0, right = n - 1;
        
        while (top <= bottom && left <= right) {
            // Traverse from left to right (top row)
            for (int col = left; col <= right; col++) {
                result.add(matrix[top][col]);
            }
            top++;
            
            // Traverse from top to bottom (right column)
            for (int row = top; row <= bottom; row++) {
                result.add(matrix[row][right]);
            }
            right--;
            
            // Traverse from right to left (bottom row) if still valid
            if (top <= bottom) {
                for (int col = right; col >= left; col--) {
                    result.add(matrix[bottom][col]);
                }
                bottom--;
            }
            
            // Traverse from bottom to top (left column) if still valid
            if (left <= right) {
                for (int row = bottom; row >= top; row--) {
                    result.add(matrix[row][left]);
                }
                left++;
            }
        }
        
        return result;
    }

    // Alternative approach using direction vectors
    public static List<Integer> spiralOrderDirections(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return result;
        }
        
        int m = matrix.length, n = matrix[0].length;
        
        // Directions: right, down, left, up
        int[][] directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
        int dirIndex = 0;
        
        // Boundaries
        int top = 0, bottom = m - 1;
        int left = 0, right = n - 1;
        
        int row = 0, col = 0;
        
        while (result.size() < m * n) {
            result.add(matrix[row][col]);
            
            // Calculate next position
            int nextRow = row + directions[dirIndex][0];
            int nextCol = col + directions[dirIndex][1];
            
            // Check if we need to change direction
            if (nextRow < top || nextRow > bottom || nextCol < left || nextCol > right) {
                // Update boundaries and change direction
                switch (dirIndex) {
                    case 0: // Moving right
                        top++;
                        break;
                    case 1: // Moving down
                        right--;
                        break;
                    case 2: // Moving left
                        bottom--;
                        break;
                    case 3: // Moving up
                        left++;
                        break;
                }
                
                dirIndex = (dirIndex + 1) % 4;
                nextRow = row + directions[dirIndex][0];
                nextCol = col + directions[dirIndex][1];
            }
            
            row = nextRow;
            col = nextCol;
        }
        
        return result;
    }

    // Recursive approach
    public static List<Integer> spiralOrderRecursive(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return result;
        }
        
        spiralHelper(matrix, 0, matrix.length - 1, 0, matrix[0].length - 1, result);
        return result;
    }

    private static void spiralHelper(int[][] matrix, int top, int bottom, int left, int right, List<Integer> result) {
        if (top > bottom || left > right) {
            return;
        }
        
        // Traverse from left to right (top row)
        for (int col = left; col <= right; col++) {
            result.add(matrix[top][col]);
        }
        
        // Traverse from top to bottom (right column)
        for (int row = top + 1; row <= bottom; row++) {
            result.add(matrix[row][right]);
        }
        
        // Traverse from right to left (bottom row) if still valid
        if (top < bottom) {
            for (int col = right - 1; col >= left; col--) {
                result.add(matrix[bottom][col]);
            }
        }
        
        // Traverse from bottom to top (left column) if still valid
        if (left < right) {
            for (int row = bottom - 1; row > top; row--) {
                result.add(matrix[row][left]);
            }
        }
        
        // Recursively process inner matrix
        spiralHelper(matrix, top + 1, bottom - 1, left + 1, right - 1, result);
    }

    public static void main(String[] args) {
        // Test cases
        int[][][] testCases = {
            {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
            {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
            {{1, 2, 3}, {4, 5, 6}},
            {{1, 2}, {3, 4}, {5, 6}, {7, 8}},
            {{1}},
            {{1, 2, 3, 4, 5}},
            {{1}, {2}, {3}, {4}, {5}},
            {{1, 2}, {3, 4}},
            {{1, 2, 3}, {4, 5}, {6, 7, 8}},
            {}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.deepToString(testCases[i]));
            
            List<Integer> result1 = spiralOrder(testCases[i]);
            List<Integer> result2 = spiralOrderDirections(testCases[i]);
            List<Integer> result3 = spiralOrderRecursive(testCases[i]);
            
            System.out.printf("  Standard: %s\n", result1);
            System.out.printf("  Directions: %s\n", result2);
            System.out.printf("  Recursive: %s\n\n", result3);
        }
    }
}
