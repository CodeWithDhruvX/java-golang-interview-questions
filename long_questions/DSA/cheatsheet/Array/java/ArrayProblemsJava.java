import java.util.*;

public class ArrayProblemsJava {

    // ============= 1. Find Largest and Smallest Element =============
    
    // Brute Force: Sort and pick first/last
    public static int[] findMinMaxBruteForce(int[] arr) {
        if (arr.length == 0) return new int[]{};
        
        int[] sorted = arr.clone();
        Arrays.sort(sorted);
        return new int[]{sorted[0], sorted[sorted.length - 1]};
    }
    
    // Optimized: Single pass with two variables
    public static int[] findMinMaxOptimized(int[] arr) {
        if (arr.length == 0) return new int[]{};
        
        int minVal = arr[0];
        int maxVal = arr[0];
        
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > maxVal) {
                maxVal = arr[i];
            }
            if (arr[i] < minVal) {
                minVal = arr[i];
            }
        }
        return new int[]{minVal, maxVal};
    }

    // ============= 2. Reverse an Array In-Place =============
    
    // Brute Force: Using extra array
    public static int[] reverseBruteForce(int[] arr) {
        int n = arr.length;
        int[] reversed = new int[n];
        
        for (int i = 0; i < n; i++) {
            reversed[i] = arr[n - 1 - i];
        }
        return reversed;
    }
    
    // Optimized: Two pointers swapping
    public static void reverseOptimized(int[] arr) {
        int left = 0;
        int right = arr.length - 1;
        
        while (left < right) {
            // Swap
            int temp = arr[left];
            arr[left] = arr[right];
            arr[right] = temp;
            
            left++;
            right--;
        }
    }

    // ============= 3. Find Second Largest Element =============
    
    // Brute Force: Sort and pick second last
    public static int secondLargestBruteForce(int[] arr) {
        if (arr.length < 2) throw new IllegalArgumentException("Array must have at least 2 elements");
        
        int[] sorted = arr.clone();
        Arrays.sort(sorted);
        
        // Handle duplicates
        for (int i = sorted.length - 2; i >= 0; i--) {
            if (sorted[i] != sorted[sorted.length - 1]) {
                return sorted[i];
            }
        }
        throw new IllegalArgumentException("No second largest element (all elements are equal)");
    }
    
    // Optimized: Single pass with two variables
    public static int secondLargestOptimized(int[] arr) {
        if (arr.length < 2) throw new IllegalArgumentException("Array must have at least 2 elements");
        
        int largest = Integer.MIN_VALUE;
        int secondLargest = Integer.MIN_VALUE;
        
        for (int num : arr) {
            if (num > largest) {
                secondLargest = largest;
                largest = num;
            } else if (num > secondLargest && num != largest) {
                secondLargest = num;
            }
        }
        
        if (secondLargest == Integer.MIN_VALUE) {
            throw new IllegalArgumentException("No second largest element (all elements are equal)");
        }
        return secondLargest;
    }

    // ============= 4. Check if Array is Sorted =============
    
    // Brute Force: Sort and compare
    public static boolean isSortedBruteForce(int[] arr) {
        int[] sorted = arr.clone();
        Arrays.sort(sorted);
        return Arrays.equals(arr, sorted);
    }
    
    // Optimized: Single pass adjacent comparison
    public static boolean isSortedOptimized(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] > arr[i + 1]) {
                return false;
            }
        }
        return true;
    }

    // ============= 5. Count Even and Odd Elements =============
    
    // Brute Force: Same as optimized (already optimal)
    public static int[] countEvenOddBruteForce(int[] arr) {
        int evenCount = 0;
        int oddCount = 0;
        
        for (int num : arr) {
            if (num % 2 == 0) {
                evenCount++;
            } else {
                oddCount++;
            }
        }
        return new int[]{evenCount, oddCount};
    }
    
    // Optimized: Same as brute force (already optimal)
    public static int[] countEvenOddOptimized(int[] arr) {
        return countEvenOddBruteForce(arr);
    }

    // ============= 6. Remove Duplicates from Sorted Array =============
    
    // Brute Force: Using HashSet
    public static int removeDuplicatesBruteForce(int[] arr) {
        Set<Integer> unique = new LinkedHashSet<>();
        for (int num : arr) {
            unique.add(num);
        }
        
        int index = 0;
        for (int num : unique) {
            arr[index++] = num;
        }
        return unique.size();
    }
    
    // Optimized: Two pointers
    public static int removeDuplicatesOptimized(int[] arr) {
        if (arr.length == 0) return 0;
        
        int i = 0;
        for (int j = 1; j < arr.length; j++) {
            if (arr[j] != arr[i]) {
                i++;
                arr[i] = arr[j];
            }
        }
        return i + 1;
    }

    // ============= 7. Left Rotate Array by 1 Position =============
    
    // Brute Force: Using extra array
    public static int[] leftRotateByOneBruteForce(int[] arr) {
        int n = arr.length;
        if (n == 0) return arr;
        
        int[] rotated = new int[n];
        for (int i = 0; i < n - 1; i++) {
            rotated[i] = arr[i + 1];
        }
        rotated[n - 1] = arr[0];
        return rotated;
    }
    
    // Optimized: In-place with temp variable
    public static void leftRotateByOneOptimized(int[] arr) {
        if (arr.length <= 1) return;
        
        int temp = arr[0];
        for (int i = 1; i < arr.length; i++) {
            arr[i - 1] = arr[i];
        }
        arr[arr.length - 1] = temp;
    }

    // ============= 8. Left Rotate Array by K Positions =============
    
    // Brute Force: K times rotate by 1
    public static void leftRotateByKBruteForce(int[] arr, int k) {
        int n = arr.length;
        k = k % n;
        
        for (int i = 0; i < k; i++) {
            leftRotateByOneOptimized(arr);
        }
    }
    
    // Optimized: Reversal algorithm
    public static void leftRotateByKOptimized(int[] arr, int k) {
        int n = arr.length;
        if (n == 0) return;
        
        k = k % n;
        if (k == 0) return;
        
        // Reverse first k elements
        reverse(arr, 0, k - 1);
        // Reverse remaining n-k elements
        reverse(arr, k, n - 1);
        // Reverse entire array
        reverse(arr, 0, n - 1);
    }
    
    private static void reverse(int[] arr, int start, int end) {
        while (start < end) {
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }
    }

    // ============= 9. Find Sum of All Elements =============
    
    // Brute Force: Simple iteration (already optimal)
    public static long sumBruteForce(int[] arr) {
        long sum = 0;
        for (int num : arr) {
            sum += num;
        }
        return sum;
    }
    
    // Optimized: Same as brute force (already optimal)
    public static long sumOptimized(int[] arr) {
        return sumBruteForce(arr);
    }

    // ============= 10. Find Frequency of Each Element =============
    
    // Brute Force: O(N^2) nested loops
    public static Map<Integer, Integer> frequencyBruteForce(int[] arr) {
        Map<Integer, Integer> freq = new HashMap<>();
        
        for (int i = 0; i < arr.length; i++) {
            int count = 0;
            for (int j = 0; j < arr.length; j++) {
                if (arr[i] == arr[j]) {
                    count++;
                }
            }
            freq.put(arr[i], count);
        }
        return freq;
    }
    
    // Optimized: HashMap single pass
    public static Map<Integer, Integer> frequencyOptimized(int[] arr) {
        Map<Integer, Integer> freq = new HashMap<>();
        
        for (int num : arr) {
            freq.put(num, freq.getOrDefault(num, 0) + 1);
        }
        return freq;
    }

    // ============= Helper Methods =============
    
    public static void printArray(int[] arr) {
        System.out.println(Arrays.toString(arr));
    }
    
    public static void printMap(Map<Integer, Integer> map) {
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            System.out.println(entry.getKey() + ": " + entry.getValue());
        }
    }
    
    public static void testAllMethods() {
        int[] testArray = {5, 2, 8, 1, 9, 3, 5, 2};
        
        System.out.println("=== Array Problems Test ===");
        System.out.println("Original Array: " + Arrays.toString(testArray));
        System.out.println();
        
        // Test 1: Min Max
        System.out.println("1. Find Min Max:");
        System.out.println("Brute Force: " + Arrays.toString(findMinMaxBruteForce(testArray)));
        System.out.println("Optimized: " + Arrays.toString(findMinMaxOptimized(testArray)));
        System.out.println();
        
        // Test 2: Reverse
        System.out.println("2. Reverse Array:");
        int[] arrForReverse = testArray.clone();
        System.out.println("Brute Force: " + Arrays.toString(reverseBruteForce(arrForReverse)));
        reverseOptimized(arrForReverse);
        System.out.println("Optimized: " + Arrays.toString(arrForReverse));
        System.out.println();
        
        // Test 3: Second Largest
        System.out.println("3. Second Largest:");
        try {
            System.out.println("Brute Force: " + secondLargestBruteForce(testArray));
            System.out.println("Optimized: " + secondLargestOptimized(testArray));
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
        System.out.println();
        
        // Test 4: Is Sorted
        System.out.println("4. Is Sorted:");
        System.out.println("Brute Force: " + isSortedBruteForce(testArray));
        System.out.println("Optimized: " + isSortedOptimized(testArray));
        System.out.println();
        
        // Test 5: Count Even Odd
        System.out.println("5. Count Even/Odd:");
        System.out.println("Brute Force: " + Arrays.toString(countEvenOddBruteForce(testArray)));
        System.out.println("Optimized: " + Arrays.toString(countEvenOddOptimized(testArray)));
        System.out.println();
        
        // Test 6: Remove Duplicates (sorted array)
        System.out.println("6. Remove Duplicates (Sorted Array):");
        int[] sortedArray = {1, 1, 2, 2, 2, 3, 4, 4, 5};
        int[] arrForRemoveBrute = sortedArray.clone();
        int[] arrForRemoveOpt = sortedArray.clone();
        System.out.println("Original: " + Arrays.toString(sortedArray));
        int bruteLength = removeDuplicatesBruteForce(arrForRemoveBrute);
        System.out.println("Brute Force: " + Arrays.toString(Arrays.copyOf(arrForRemoveBrute, bruteLength)));
        int optLength = removeDuplicatesOptimized(arrForRemoveOpt);
        System.out.println("Optimized: " + Arrays.toString(Arrays.copyOf(arrForRemoveOpt, optLength)));
        System.out.println();
        
        // Test 7: Left Rotate by 1
        System.out.println("7. Left Rotate by 1:");
        int[] arrForRotate1Brute = testArray.clone();
        int[] arrForRotate1Opt = testArray.clone();
        System.out.println("Original: " + Arrays.toString(testArray));
        System.out.println("Brute Force: " + Arrays.toString(leftRotateByOneBruteForce(arrForRotate1Brute)));
        leftRotateByOneOptimized(arrForRotate1Opt);
        System.out.println("Optimized: " + Arrays.toString(arrForRotate1Opt));
        System.out.println();
        
        // Test 8: Left Rotate by K
        System.out.println("8. Left Rotate by K (K=3):");
        int[] arrForRotateKBrute = testArray.clone();
        int[] arrForRotateKOpt = testArray.clone();
        System.out.println("Original: " + Arrays.toString(testArray));
        leftRotateByKBruteForce(arrForRotateKBrute, 3);
        System.out.println("Brute Force: " + Arrays.toString(arrForRotateKBrute));
        leftRotateByKOptimized(arrForRotateKOpt, 3);
        System.out.println("Optimized: " + Arrays.toString(arrForRotateKOpt));
        System.out.println();
        
        // Test 9: Sum
        System.out.println("9. Sum of Elements:");
        System.out.println("Brute Force: " + sumBruteForce(testArray));
        System.out.println("Optimized: " + sumOptimized(testArray));
        System.out.println();
        
        // Test 10: Frequency
        System.out.println("10. Frequency Count:");
        System.out.println("Brute Force:");
        printMap(frequencyBruteForce(testArray));
        System.out.println("Optimized:");
        printMap(frequencyOptimized(testArray));
    }
    
    public static void main(String[] args) {
        testAllMethods();
    }
}
