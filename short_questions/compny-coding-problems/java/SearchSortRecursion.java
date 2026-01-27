package java_solutions;

import java.util.*;

public class SearchSortRecursion {

    // 71. Linear search
    public static int linearSearch(int[] arr, int target) {
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == target)
                return i;
        }
        return -1;
    }

    // 72. Binary search (Sorted Array)
    public static int binarySearch(int[] arr, int target) {
        int low = 0, high = arr.length - 1;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (arr[mid] == target)
                return mid;
            else if (arr[mid] < target)
                low = mid + 1;
            else
                high = mid - 1;
        }
        return -1;
    }

    // 73. Bubble sort
    public static void bubbleSort(int[] arr) {
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
    }

    // 74. Selection sort
    public static void selectionSort(int[] arr) {
        int n = arr.length;
        for (int i = 0; i < n - 1; i++) {
            int minIdx = i;
            for (int j = i + 1; j < n; j++) {
                if (arr[j] < arr[minIdx])
                    minIdx = j;
            }
            int temp = arr[minIdx];
            arr[minIdx] = arr[i];
            arr[i] = temp;
        }
    }

    // 75. Insertion sort
    public static void insertionSort(int[] arr) {
        int n = arr.length;
        for (int i = 1; i < n; i++) {
            int key = arr[i];
            int j = i - 1;
            while (j >= 0 && arr[j] > key) {
                arr[j + 1] = arr[j];
                j = j - 1;
            }
            arr[j + 1] = key;
        }
    }

    // 76. Merge sort
    public static void mergeSort(int[] arr, int l, int r) {
        if (l < r) {
            int m = l + (r - l) / 2;
            mergeSort(arr, l, m);
            mergeSort(arr, m + 1, r);
            merge(arr, l, m, r);
        }
    }

    private static void merge(int[] arr, int l, int m, int r) {
        int n1 = m - l + 1;
        int n2 = r - m;
        int[] L = new int[n1];
        int[] R = new int[n2];
        for (int i = 0; i < n1; ++i)
            L[i] = arr[l + i];
        for (int j = 0; j < n2; ++j)
            R[j] = arr[m + 1 + j];
        int i = 0, j = 0;
        int k = l;
        while (i < n1 && j < n2) {
            if (L[i] <= R[j])
                arr[k++] = L[i++];
            else
                arr[k++] = R[j++];
        }
        while (i < n1)
            arr[k++] = L[i++];
        while (j < n2)
            arr[k++] = R[j++];
    }

    // 77. Quick sort
    public static void quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int pi = partition(arr, low, high);
            quickSort(arr, low, pi - 1);
            quickSort(arr, pi + 1, high);
        }
    }

    private static int partition(int[] arr, int low, int high) {
        int pivot = arr[high];
        int i = (low - 1);
        for (int j = low; j < high; j++) {
            if (arr[j] < pivot) {
                i++;
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
        int temp = arr[i + 1];
        arr[i + 1] = arr[high];
        arr[high] = temp;
        return i + 1;
    }

    // 78. Recursion: Factorial
    public static int factorialRec(int N) {
        if (N <= 1)
            return 1;
        return N * factorialRec(N - 1);
    }

    // 79. Recursion: Fibonacci
    public static int fibonacciRec(int N) {
        if (N <= 1)
            return N;
        return fibonacciRec(N - 1) + fibonacciRec(N - 2);
    }

    // 80. Recursion: Sum of digits
    public static int sumDigitsRec(int N) {
        if (N == 0)
            return 0;
        return (N % 10) + sumDigitsRec(N / 10);
    }

    // 81. Recursion: Reverse string
    public static String reverseStringRec(String str) {
        if (str.isEmpty())
            return "";
        return reverseStringRec(str.substring(1)) + str.charAt(0);
    }

    // 82. Power of number (Recursion)
    public static int powerRec(int base, int exp) {
        if (exp == 0)
            return 1;
        return base * powerRec(base, exp - 1);
    }

    // 83. GCD Recursion
    public static int gcdRec(int a, int b) {
        if (b == 0)
            return a;
        return gcdRec(b, a % b);
    }

    // 84. Tower of Hanoi
    public static void towerOfHanoi(int n, char from, char to, char aux) {
        if (n == 1) {
            System.out.println("Move disk 1 from " + from + " to " + to);
            return;
        }
        towerOfHanoi(n - 1, from, aux, to);
        System.out.println("Move disk " + n + " from " + from + " to " + to);
        towerOfHanoi(n - 1, aux, to, from);
    }

    // 85. Permutations of a string
    public static void permute(String str, int l, int r) {
        if (l == r)
            System.out.println(str);
        else {
            for (int i = l; i <= r; i++) {
                str = swap(str, l, i);
                permute(str, l + 1, r);
                str = swap(str, l, i); // backtrack
            }
        }
    }

    private static String swap(String a, int i, int j) {
        char[] charArray = a.toCharArray();
        char temp = charArray[i];
        charArray[i] = charArray[j];
        charArray[j] = temp;
        return String.valueOf(charArray);
    }

    // 87. Binary Search Recursive
    public static int binarySearchRec(int[] arr, int low, int high, int target) {
        if (low <= high) {
            int mid = low + (high - low) / 2;
            if (arr[mid] == target)
                return mid;
            if (arr[mid] > target)
                return binarySearchRec(arr, low, mid - 1, target);
            return binarySearchRec(arr, mid + 1, high, target);
        }
        return -1;
    }

    // 88. Check Palindrome String (Recursive)
    public static boolean isPalindromeRec(String str) {
        if (str.length() <= 1)
            return true;
        if (str.charAt(0) != str.charAt(str.length() - 1))
            return false;
        return isPalindromeRec(str.substring(1, str.length() - 1));
    }

    // 89. Subset Sum Problem (Recursion)
    public static boolean isSubsetSum(int[] arr, int n, int sum) {
        if (sum == 0)
            return true;
        if (n == 0)
            return false;
        if (arr[n - 1] > sum)
            return isSubsetSum(arr, n - 1, sum);
        return isSubsetSum(arr, n - 1, sum) || isSubsetSum(arr, n - 1, sum - arr[n - 1]);
    }

    // 90. Print 1 to N without loops
    public static void printNos(int N) {
        if (N > 0) {
            printNos(N - 1);
            System.out.print(N + " ");
        }
    }

    public static void main(String[] args) {
        System.out.println(
                "71. Linear Search: [10, 20, 30, 40], 30 -> " + linearSearch(new int[] { 10, 20, 30, 40 }, 30));
        System.out.println(
                "72. Binary Search: [10, 20, 30, 40], 30 -> " + binarySearch(new int[] { 10, 20, 30, 40 }, 30));
        int[] arr = { 5, 1, 4, 2, 8 };
        bubbleSort(arr);
        System.out.println("73. Bubble Sort: " + Arrays.toString(arr));

        int[] arr2 = { 5, 1, 4, 2, 8 };
        selectionSort(arr2);
        System.out.println("74. Selection Sort: " + Arrays.toString(arr2));

        int[] arr3 = { 5, 1, 4, 2, 8 };
        insertionSort(arr3);
        System.out.println("75. Insertion Sort: " + Arrays.toString(arr3));

        int[] arrMerge = { 5, 1, 4, 2, 8 };
        mergeSort(arrMerge, 0, arrMerge.length - 1);
        System.out.println("76. Merge Sort: " + Arrays.toString(arrMerge));

        int[] arrQuick = { 10, 7, 8, 9, 1, 5 };
        quickSort(arrQuick, 0, arrQuick.length - 1);
        System.out.println("77. Quick Sort: " + Arrays.toString(arrQuick));

        System.out.println("78. Factorial Rec: 5 -> " + factorialRec(5));
        System.out.println("79. Fibonacci Rec: 5 -> " + fibonacciRec(5));
        System.out.println("80. Sum Digits Rec: 123 -> " + sumDigitsRec(123));
        System.out.println("81. Reverse String Rec: hello -> " + reverseStringRec("hello"));
        System.out.println("82. Power Rec: 2, 3 -> " + powerRec(2, 3));
        System.out.println("83. GCD Rec: 12, 18 -> " + gcdRec(12, 18));

        System.out.println("84. Tower of Hanoi (3):");
        towerOfHanoi(3, 'A', 'C', 'B');
        System.out.println("85. Permutations (ABC):");
        permute("ABC", 0, 2);
        System.out.println("87. Binary Search Rec: [10, 20, 30, 40], 30 -> "
                + binarySearchRec(new int[] { 10, 20, 30, 40 }, 0, 3, 30));
        System.out.println("88. Is Palindrome Rec: madam -> " + isPalindromeRec("madam"));
        System.out.println(
                "89. Subset Sum: [3, 34, 4, 12, 5, 2], 9 -> " + isSubsetSum(new int[] { 3, 34, 4, 12, 5, 2 }, 6, 9));
        System.out.print("90. Print 1 to N: ");
        printNos(10);
        System.out.println();
    }
}
