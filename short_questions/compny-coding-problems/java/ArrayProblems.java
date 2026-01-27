package java_solutions;

import java.util.*;

public class ArrayProblems {

    // 21. Find largest element in array
    public static int findMax(int[] arr) {
        int max = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > max) {
                max = arr[i];
            }
        }
        return max;
    }

    // 22. Find smallest element in array
    public static int findMin(int[] arr) {
        int min = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] < min) {
                min = arr[i];
            }
        }
        return min;
    }

    // 23. Find second largest element
    public static int secondLargest(int[] arr) {
        int first = Integer.MIN_VALUE, second = Integer.MIN_VALUE;
        for (int num : arr) {
            if (num > first) {
                second = first;
                first = num;
            } else if (num > second && num != first) {
                second = num;
            }
        }
        return second;
    }

    // 24. Reverse an array
    public static int[] reverseArray(int[] arr) {
        int start = 0, end = arr.length - 1;
        while (start < end) {
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }
        return arr;
    }

    // 25. Rotate array left by K positions
    public static int[] rotateLeft(int[] arr, int k) {
        k = k % arr.length;
        reverse(arr, 0, k - 1);
        reverse(arr, k, arr.length - 1);
        reverse(arr, 0, arr.length - 1);
        return arr;
    }

    // 26. Rotate array right by K positions
    public static int[] rotateRight(int[] arr, int k) {
        k = k % arr.length;
        reverse(arr, 0, arr.length - k - 1);
        reverse(arr, arr.length - k, arr.length - 1);
        reverse(arr, 0, arr.length - 1);
        return arr;
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

    // 27. Remove duplicates from array
    public static List<Integer> removeArrayDuplicates(int[] arr) {
        Set<Integer> seen = new HashSet<>();
        List<Integer> result = new ArrayList<>();
        for (int num : arr) {
            if (!seen.contains(num)) {
                seen.add(num);
                result.add(num);
            }
        }
        return result;
    }

    // 28. Find frequency of elements
    public static void arrayFrequency(int[] arr) {
        Map<Integer, Integer> freqMap = new HashMap<>();
        for (int num : arr) {
            freqMap.put(num, freqMap.getOrDefault(num, 0) + 1);
        }
        System.out.println(freqMap);
    }

    // 29. Find missing number in array (1 to N)
    public static int findMissing(int[] arr, int N) {
        int expectedSum = N * (N + 1) / 2;
        int actualSum = 0;
        for (int num : arr) {
            actualSum += num;
        }
        return expectedSum - actualSum;
    }

    // 30. Find duplicate number
    public static int findDuplicate(int[] arr) {
        Set<Integer> seen = new HashSet<>();
        for (int num : arr) {
            if (seen.contains(num)) {
                return num;
            }
            seen.add(num);
        }
        return -1;
    }

    // 31. Sort array without built-in methods (Bubble Sort)
    public static int[] bubbleSort(int[] arr) {
        int n = arr.length;
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < n - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                }
            }
        }
        return arr;
    }

    // 32. Merge two arrays
    public static int[] mergeArrays(int[] arr1, int[] arr2) {
        int[] result = new int[arr1.length + arr2.length];
        System.arraycopy(arr1, 0, result, 0, arr1.length);
        System.arraycopy(arr2, 0, result, arr1.length, arr2.length);
        return result;
    }

    // 33. Find common elements in two arrays
    public static List<Integer> findCommon(int[] arr1, int[] arr2) {
        Set<Integer> set1 = new HashSet<>();
        for (int num : arr1)
            set1.add(num);
        List<Integer> common = new ArrayList<>();
        for (int num : arr2) {
            if (set1.contains(num)) {
                common.add(num);
                set1.remove(num);
            }
        }
        return common;
    }

    // 34. Move all zeros to end
    public static int[] moveZeros(int[] arr) {
        int count = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] != 0) {
                arr[count++] = arr[i];
            }
        }
        while (count < arr.length) {
            arr[count++] = 0;
        }
        return arr;
    }

    // 35. Find sum of array elements
    public static int sumArray(int[] arr) {
        int total = 0;
        for (int num : arr)
            total += num;
        return total;
    }

    // 36. Find pair with given sum
    public static void findPairWithSum(int[] arr, int target) {
        Set<Integer> seen = new HashSet<>();
        for (int num : arr) {
            int complement = target - num;
            if (seen.contains(complement)) {
                System.out.println("{" + complement + ", " + num + "}");
                return;
            }
            seen.add(num);
        }
        System.out.println("NULL");
    }

    // 37. Find max & min in single loop
    public static void findMaxMin(int[] arr) {
        int max = arr[0], min = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > max)
                max = arr[i];
            if (arr[i] < min)
                min = arr[i];
        }
        System.out.println("{" + max + ", " + min + "}");
    }

    // 38. Print array in reverse order
    public static void printReverse(int[] arr) {
        for (int i = arr.length - 1; i >= 0; i--) {
            System.out.print(arr[i] + " ");
        }
        System.out.println();
    }

    // 39. Check array is sorted or not
    public static boolean isSorted(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] > arr[i + 1])
                return false;
        }
        return true;
    }

    // 40. Count even & odd numbers
    public static void countEvenOdd(int[] arr) {
        int even = 0, odd = 0;
        for (int num : arr) {
            if (num % 2 == 0)
                even++;
            else
                odd++;
        }
        System.out.println("Even: " + even + ", Odd: " + odd);
    }

    public static void main(String[] args) {
        System.out.println("21. Find Max: [1, 5, 3, 9, 2] -> " + findMax(new int[] { 1, 5, 3, 9, 2 }));
        System.out.println("22. Find Min: [1, 5, 3, 9, 2] -> " + findMin(new int[] { 1, 5, 3, 9, 2 }));
        System.out.println(
                "23. Second Largest: [12, 35, 1, 10, 34, 1] -> " + secondLargest(new int[] { 12, 35, 1, 10, 34, 1 }));
        System.out.println(
                "24. Reverse Array: [1, 2, 3, 4] -> " + Arrays.toString(reverseArray(new int[] { 1, 2, 3, 4 })));
        System.out.println("25. Rotate Left: [1, 2, 3, 4, 5], k=2 -> "
                + Arrays.toString(rotateLeft(new int[] { 1, 2, 3, 4, 5 }, 2)));
        System.out.println("26. Rotate Right: [1, 2, 3, 4, 5], k=2 -> "
                + Arrays.toString(rotateRight(new int[] { 1, 2, 3, 4, 5 }, 2)));
        System.out.println("27. Remove Duplicates: [1, 2, 2, 3, 4, 4] -> "
                + removeArrayDuplicates(new int[] { 1, 2, 2, 3, 4, 4 }));
        System.out.print("28. Frequency: ");
        arrayFrequency(new int[] { 1, 2, 2, 3 });
        System.out.println("29. Find Missing: [1, 2, 4, 5], N=5 -> " + findMissing(new int[] { 1, 2, 4, 5 }, 5));
        System.out.println("30. Find Duplicate: [1, 3, 4, 2, 2] -> " + findDuplicate(new int[] { 1, 3, 4, 2, 2 }));
        System.out.println(
                "31. Bubble Sort: [5, 1, 4, 2, 8] -> " + Arrays.toString(bubbleSort(new int[] { 5, 1, 4, 2, 8 })));
        System.out.println("32. Merge Arrays: [1, 2], [3, 4] -> "
                + Arrays.toString(mergeArrays(new int[] { 1, 2 }, new int[] { 3, 4 })));
        System.out.println("33. Common Elements: [1, 2, 3], [2, 3, 4] -> "
                + findCommon(new int[] { 1, 2, 3 }, new int[] { 2, 3, 4 }));
        System.out.println(
                "34. Move Zeros: [0, 1, 0, 3, 12] -> " + Arrays.toString(moveZeros(new int[] { 0, 1, 0, 3, 12 })));
        System.out.println("35. Sum Array: [1, 2, 3, 4] -> " + sumArray(new int[] { 1, 2, 3, 4 }));
        System.out.print("36. Find Pair Sum: [2, 7, 11, 15], 9 -> ");
        findPairWithSum(new int[] { 2, 7, 11, 15 }, 9);
        System.out.print("37. Max & Min: [3, 5, 1, 2, 4, 8] -> ");
        findMaxMin(new int[] { 3, 5, 1, 2, 4, 8 });
        System.out.print("38. Print Reverse: [1, 2, 3, 4] -> ");
        printReverse(new int[] { 1, 2, 3, 4 });
        System.out.println("39. Is Sorted: [1, 2, 3, 4, 5] -> " + isSorted(new int[] { 1, 2, 3, 4, 5 }));
        System.out.print("40. Even Odd: [1, 2, 3, 4, 5] -> ");
        countEvenOdd(new int[] { 1, 2, 3, 4, 5 });
    }
}
