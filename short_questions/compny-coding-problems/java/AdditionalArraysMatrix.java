package java_solutions;

import java.util.*;

public class AdditionalArraysMatrix {

    // 16. Find leaders in array
    public static void findLeaders(int[] arr) {
        int n = arr.length;
        if (n == 0)
            return;
        int maxRight = arr[n - 1];
        System.out.print(maxRight + " ");
        for (int i = n - 2; i >= 0; i--) {
            if (arr[i] > maxRight) {
                maxRight = arr[i];
                System.out.print(maxRight + " ");
            }
        }
        System.out.println();
    }

    // 17. Find equilibrium index
    public static int equilibriumPoint(int[] arr) {
        int totalSum = Arrays.stream(arr).sum();
        int leftSum = 0;
        for (int i = 0; i < arr.length; i++) {
            totalSum -= arr[i];
            if (leftSum == totalSum)
                return i;
            leftSum += arr[i];
        }
        return -1;
    }

    // 18. Subarray with given sum
    public static void subArraySum(int[] arr, int sum) {
        int currSum = arr[0], start = 0;
        for (int i = 1; i <= arr.length; i++) {
            while (currSum > sum && start < i - 1) {
                currSum -= arr[start++];
            }
            if (currSum == sum) {
                System.out.println("Sum found between " + start + " and " + (i - 1));
                return;
            }
            if (i < arr.length)
                currSum += arr[i];
        }
        System.out.println("No subarray found");
    }

    // 19. Kadane’s algorithm
    public static int kadanes(int[] arr) {
        if (arr.length == 0)
            return 0;
        int maxSoFar = Integer.MIN_VALUE, currMax = 0;
        for (int num : arr) {
            currMax += num;
            if (maxSoFar < currMax)
                maxSoFar = currMax;
            if (currMax < 0)
                currMax = 0;
        }
        return maxSoFar;
    }

    // 20. Find majority element
    public static int majorityElement(int[] arr) {
        int candidate = -1, count = 0;
        for (int num : arr) {
            if (count == 0)
                candidate = num;
            if (num == candidate)
                count++;
            else
                count--;
        }
        return candidate;
    }

    // 21. Rearrange array alternately
    public static void rearrangeAlternate(int[] arr) {
        Arrays.sort(arr);
        int i = 0, j = arr.length - 1;
        while (i < j) {
            System.out.print(arr[j--] + " ");
            System.out.print(arr[i++] + " ");
        }
        if (i == j)
            System.out.print(arr[i]);
        System.out.println();
    }

    // 23. Find union of two arrays
    public static List<Integer> unionArrays(int[] arr1, int[] arr2) {
        Set<Integer> set = new HashSet<>();
        for (int v : arr1)
            set.add(v);
        for (int v : arr2)
            set.add(v);
        List<Integer> res = new ArrayList<>(set);
        Collections.sort(res);
        return res;
    }

    // 24. Find intersection of two arrays
    public static List<Integer> intersectionArrays(int[] arr1, int[] arr2) {
        Set<Integer> set1 = new HashSet<>();
        for (int v : arr1)
            set1.add(v);
        List<Integer> res = new ArrayList<>();
        for (int v : arr2) {
            if (set1.contains(v)) {
                res.add(v);
                set1.remove(v);
            }
        }
        Collections.sort(res);
        return res;
    }

    // 25. Count pairs with given difference
    public static int countDiffPairs(int[] arr, int k) {
        int count = 0;
        for (int i = 0; i < arr.length; i++) {
            for (int j = i + 1; j < arr.length; j++) {
                if (Math.abs(arr[i] - arr[j]) == k)
                    count++;
            }
        }
        return count;
    }

    // 26. Find peak element
    public static int findPeak(int[] arr) {
        int n = arr.length;
        if (n == 1)
            return 0;
        if (arr[0] >= arr[1])
            return 0;
        if (arr[n - 1] >= arr[n - 2])
            return n - 1;
        for (int i = 1; i < n - 1; i++) {
            if (arr[i] >= arr[i - 1] && arr[i] >= arr[i + 1])
                return i;
        }
        return -1;
    }

    // 27. Left rotate by 1
    public static void leftRotateOne(int[] arr) {
        if (arr.length == 0)
            return;
        int temp = arr[0];
        for (int i = 0; i < arr.length - 1; i++)
            arr[i] = arr[i + 1];
        arr[arr.length - 1] = temp;
        System.out.println(Arrays.toString(arr));
    }

    // 28. Find minimum difference pair
    public static int minDiffPair(int[] arr) {
        Arrays.sort(arr);
        int minDiff = Integer.MAX_VALUE;
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i + 1] - arr[i] < minDiff)
                minDiff = arr[i + 1] - arr[i];
        }
        return minDiff;
    }

    // 29. Product of array except self
    public static int[] productExceptSelf(int[] arr) {
        int n = arr.length;
        int[] left = new int[n];
        int[] right = new int[n];
        int[] prod = new int[n];
        left[0] = 1;
        right[n - 1] = 1;
        for (int i = 1; i < n; i++)
            left[i] = arr[i - 1] * left[i - 1];
        for (int i = n - 2; i >= 0; i--)
            right[i] = arr[i + 1] * right[i + 1];
        for (int i = 0; i < n; i++)
            prod[i] = left[i] * right[i];
        return prod;
    }

    // 30. Find subarray with max product
    public static int maxProductSubarray(int[] arr) {
        if (arr.length == 0)
            return 0;
        int maxSoFar = arr[0], minSoFar = arr[0], result = maxSoFar;
        for (int i = 1; i < arr.length; i++) {
            int curr = arr[i];
            int tempMax = Math.max(curr, Math.max(curr * maxSoFar, curr * minSoFar));
            minSoFar = Math.min(curr, Math.min(curr * maxSoFar, curr * minSoFar));
            maxSoFar = tempMax;
            result = Math.max(maxSoFar, result);
        }
        return result;
    }

    // 32. Separate positive and negative numbers
    public static void separatePosNeg(int[] arr) {
        int j = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] < 0) {
                if (i != j) {
                    int temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
                j++;
            }
        }
        System.out.println(Arrays.toString(arr));
    }

    // 33. Count distinct elements
    public static int countDistinct(int[] arr) {
        Set<Integer> set = new HashSet<>();
        for (int v : arr)
            set.add(v);
        return set.size();
    }

    // 34. Replace element with next greatest
    public static void replaceNextGreatest(int[] arr) {
        int maxFromRight = -1;
        for (int i = arr.length - 1; i >= 0; i--) {
            int temp = arr[i];
            arr[i] = maxFromRight;
            if (temp > maxFromRight)
                maxFromRight = temp;
        }
        System.out.println(Arrays.toString(arr));
    }

    // 35. Find smallest subarray with sum > X
    public static int smallestSubWithSum(int[] arr, int x) {
        int minLen = arr.length + 1;
        int currSum = 0, start = 0, end = 0;
        while (end < arr.length) {
            currSum += arr[end++];
            while (currSum > x && start < end) {
                minLen = Math.min(minLen, end - start);
                currSum -= arr[start++];
            }
        }
        return minLen;
    }

    // --- MATRIX ---

    // 36. Matrix addition
    public static int[][] addMatrices(int[][] A, int[][] B) {
        int rows = A.length, cols = A[0].length;
        int[][] C = new int[rows][cols];
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++)
                C[i][j] = A[i][j] + B[i][j];
        }
        return C;
    }

    // 37. Matrix multiplication
    public static int[][] multiplyMatrices(int[][] A, int[][] B) {
        int rowsA = A.length, colsA = A[0].length, colsB = B[0].length;
        int[][] C = new int[rowsA][colsB];
        for (int i = 0; i < rowsA; i++) {
            for (int j = 0; j < colsB; j++) {
                for (int k = 0; k < colsA; k++)
                    C[i][j] += A[i][k] * B[k][j];
            }
        }
        return C;
    }

    // 38. Transpose of matrix
    public static int[][] transpose(int[][] A) {
        int rows = A.length, cols = A[0].length;
        int[][] T = new int[cols][rows];
        for (int i = 0; i < cols; i++) {
            for (int j = 0; j < rows; j++)
                T[i][j] = A[j][i];
        }
        return T;
    }

    // 39. Rotate matrix 90 degrees
    public static int[][] rotateMatrix(int[][] A) {
        int[][] T = transpose(A);
        for (int i = 0; i < T.length; i++) {
            int left = 0, right = T[i].length - 1;
            while (left < right) {
                int temp = T[i][left];
                T[i][left] = T[i][right];
                T[i][right] = temp;
                left++;
                right--;
            }
        }
        return T;
    }

    // 40. Print matrix in spiral order
    public static void spiralOrder(int[][] matrix) {
        int top = 0, bottom = matrix.length - 1;
        int left = 0, right = matrix[0].length - 1;

        while (top <= bottom && left <= right) {
            for (int i = left; i <= right; i++)
                System.out.print(matrix[top][i] + " ");
            top++;
            for (int i = top; i <= bottom; i++)
                System.out.print(matrix[i][right] + " ");
            right--;
            if (top <= bottom) {
                for (int i = right; i >= left; i--)
                    System.out.print(matrix[bottom][i] + " ");
                bottom--;
            }
            if (left <= right) {
                for (int i = bottom; i >= top; i--)
                    System.out.print(matrix[i][left] + " ");
                left++;
            }
        }
        System.out.println();
    }

    // 41. Search element in sorted matrix
    public static boolean searchMatrix(int[][] matrix, int target) {
        int row = 0, col = matrix[0].length - 1;
        while (row < matrix.length && col >= 0) {
            if (matrix[row][col] == target)
                return true;
            if (matrix[row][col] > target)
                col--;
            else
                row++;
        }
        return false;
    }

    // 42. Diagonal sum (Primary & Secondary)
    public static int diagonalSum(int[][] matrix) {
        int n = matrix.length;
        int sum = 0;
        for (int i = 0; i < n; i++) {
            sum += matrix[i][i];
            sum += matrix[i][n - i - 1];
        }
        if (n % 2 != 0)
            sum -= matrix[n / 2][n / 2];
        return sum;
    }

    // 43. Print boundary elements
    public static void printBoundary(int[][] matrix) {
        int rows = matrix.length, cols = matrix[0].length;
        for (int col = 0; col < cols; col++)
            System.out.print(matrix[0][col] + " ");
        for (int row = 1; row < rows; row++)
            System.out.print(matrix[row][cols - 1] + " ");
        for (int col = cols - 2; col >= 0; col--)
            System.out.print(matrix[rows - 1][col] + " ");
        for (int row = rows - 2; row > 0; row--)
            System.out.print(matrix[row][0] + " ");
        System.out.println();
    }

    // 44. Check symmetric matrix
    public static boolean isSymmetric(int[][] matrix) {
        int rows = matrix.length, cols = matrix[0].length;
        if (rows != cols)
            return false;
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (matrix[i][j] != matrix[j][i])
                    return false;
            }
        }
        return true;
    }

    // 46. Count zeros and ones
    public static void countZeroOne(int[][] matrix) {
        int zeros = 0, ones = 0;
        for (int[] row : matrix) {
            for (int val : row) {
                if (val == 0)
                    zeros++;
                else if (val == 1)
                    ones++;
            }
        }
        System.out.println("Zeros: " + zeros + ", Ones: " + ones);
    }

    // 47. Row with maximum 1s
    public static int rowMaxOnes(int[][] matrix) {
        int maxOnes = 0, rowIndex = -1;
        for (int i = 0; i < matrix.length; i++) {
            int count = 0;
            for (int val : matrix[i])
                if (val == 1)
                    count++;
            if (count > maxOnes) {
                maxOnes = count;
                rowIndex = i;
            }
        }
        return rowIndex;
    }

    // 49. Snake pattern printing
    public static void snakePattern(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            if (i % 2 == 0) {
                for (int j = 0; j < matrix[0].length; j++)
                    System.out.print(matrix[i][j] + " ");
            } else {
                for (int j = matrix[0].length - 1; j >= 0; j--)
                    System.out.print(matrix[i][j] + " ");
            }
        }
        System.out.println();
    }

    // 50. Identity matrix check
    public static boolean isIdentity(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[0].length; j++) {
                if (i == j && matrix[i][j] != 1)
                    return false;
                if (i != j && matrix[i][j] != 0)
                    return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        System.out.println("16. Find Leaders:");
        findLeaders(new int[] { 16, 17, 4, 3, 5, 2 });
        System.out.println("17. Equilibrium Point: " + equilibriumPoint(new int[] { 1, 3, 5, 2, 2 }));
        System.out.print("18. Subarray Sum: ");
        subArraySum(new int[] { 1, 2, 3, 7, 5 }, 12);
        System.out.println("19. Kadanes: " + kadanes(new int[] { -2, 1, -3, 4, -1, 2, 1, -5, 4 }));
        System.out.println("20. Majority Element: " + majorityElement(new int[] { 3, 2, 3 }));
        System.out.print("21. Rearrange Alternate: ");
        rearrangeAlternate(new int[] { 1, 2, 3, 4, 5, 6 });
        System.out.println("23. Union Arrays: " + unionArrays(new int[] { 1, 2, 3 }, new int[] { 2, 3, 4 }));
        System.out.println(
                "24. Intersection Arrays: " + intersectionArrays(new int[] { 1, 2, 3 }, new int[] { 2, 3, 4 }));
        System.out.println("25. Count Diff Pairs: " + countDiffPairs(new int[] { 1, 5, 3, 4, 2 }, 3));
        System.out.println("26. Find Peak: " + findPeak(new int[] { 1, 2, 3, 1 }));
        System.out.print("27. Left Rotate By 1: ");
        leftRotateOne(new int[] { 1, 2, 3, 4, 5 });
        System.out.println("28. Min Diff Pair: " + minDiffPair(new int[] { 2, 4, 5, 9, 7 }));
        System.out.println("29. Product Except Self: " + Arrays.toString(productExceptSelf(new int[] { 1, 2, 3, 4 })));
        System.out.println("30. Max Product Subarray: " + maxProductSubarray(new int[] { 2, 3, -2, 4 }));
        System.out.print("32. Separate Pos Neg: ");
        separatePosNeg(new int[] { -1, 2, -3, 4, 5, 6, -7, 8, 9 });
        System.out.println("33. Count Distinct: " + countDistinct(new int[] { 10, 20, 20, 10, 30 }));
        System.out.print("34. Replace Next Greatest: ");
        replaceNextGreatest(new int[] { 16, 17, 4, 3, 5, 2 });
        System.out.println("35. Smallest Sub Sum > X: " + smallestSubWithSum(new int[] { 1, 4, 45, 6, 0, 19 }, 51));

        System.out.println("\n--- MATRIX ---");
        int[][] A = { { 1, 2 }, { 3, 4 } };
        int[][] B = { { 5, 6 }, { 7, 8 } };
        System.out.println("36. Add Matrices: " + Arrays.deepToString(addMatrices(A, B)));
        System.out.println("37. Multiply Matrices: "
                + Arrays.deepToString(multiplyMatrices(A, new int[][] { { 1, 0 }, { 0, 1 } })));
        System.out.println("38. Transpose: " + Arrays.deepToString(transpose(A)));
        System.out.println("39. Rotate 90: "
                + Arrays.deepToString(rotateMatrix(new int[][] { { 1, 2, 3 }, { 4, 5, 6 }, { 7, 8, 9 } })));
        System.out.print("40. Spiral Order: ");
        spiralOrder(new int[][] { { 1, 2, 3 }, { 4, 5, 6 }, { 7, 8, 9 } });
        System.out.println("41. Search Matrix: " + searchMatrix(new int[][] { { 1, 3, 5 }, { 7, 9, 11 } }, 9));
        System.out.println("42. Diagonal Sum: " + diagonalSum(new int[][] { { 1, 2, 3 }, { 4, 5, 6 }, { 7, 8, 9 } }));
        System.out.print("43. Print Boundary: ");
        printBoundary(new int[][] { { 1, 2, 3 }, { 4, 5, 6 }, { 7, 8, 9 } });
        System.out.println("44. Is Symmetric: " + isSymmetric(new int[][] { { 1, 2 }, { 2, 1 } }));
        System.out.print("46. Count Zero One: ");
        countZeroOne(new int[][] { { 0, 1 }, { 1, 0 } });
        System.out.println("47. Row Max Ones: " + rowMaxOnes(new int[][] { { 0, 1, 1 }, { 0, 0, 1 } }));
        System.out.print("49. Snake Pattern: ");
        snakePattern(new int[][] { { 1, 2 }, { 3, 4 } });
        System.out.println("50. Is Identity: " + isIdentity(new int[][] { { 1, 0 }, { 0, 1 } }));
    }
}
