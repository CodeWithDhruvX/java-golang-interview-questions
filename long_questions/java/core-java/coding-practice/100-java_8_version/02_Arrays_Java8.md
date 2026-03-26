# Array Programs with Java 8 Features (16-35)

## 📚 Java 8 Features Demonstrated
- **Lambda Expressions**: Concise array operations
- **Streams API**: Functional array processing
- **Method References**: Simplified operations
- **Collectors**: Array aggregations
- **Functional Interfaces**: Predicate, Function, Consumer
- **Parallel Streams**: Multi-threaded array processing
- **Optional**: Null-safe array operations

---

## 16. Largest Element in Array
**Java 8 Approach**: Using `Arrays.stream()` and `max()`

```java
import java.util.*;
import java.util.stream.*;

public class LargestInArrayJava8 {
    public static void main(String[] args) {
        int[] arr = {10, 50, 20, 90, 40};
        
        // Using Java 8 Streams
        int max = Arrays.stream(arr)
            .max()
            .orElse(Integer.MIN_VALUE);
        
        System.out.println("Largest: " + max);
        
        // Alternative: Using IntStream directly
        int maxInt = IntStream.of(arr)
            .max()
            .orElse(Integer.MIN_VALUE);
        
        System.out.println("Largest (IntStream): " + maxInt);
        
        // Find index of largest element
        int maxIndex = IntStream.range(0, arr.length)
            .reduce((i, j) -> arr[j] > arr[i] ? j : i)
            .orElse(-1);
        
        System.out.println("Largest element at index: " + maxIndex);
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{10, 50, 20, 90, 40},
            new int[]{5, 15, 25, 35},
            new int[]{100, 200, 300, 400, 500}
        );
        
        Map<String, Integer> maxResults = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array).max().orElse(Integer.MIN_VALUE)
            ));
        
        System.out.println("Maximum values: " + maxResults);
        
        // Using parallel stream for large arrays
        int[] largeArray = IntStream.range(1, 1000001).toArray();
        int maxParallel = Arrays.stream(largeArray)
            .parallel()
            .max()
            .orElse(Integer.MIN_VALUE);
        
        System.out.println("Max in large array (parallel): " + maxParallel);
    }
}
```

## 17. Smallest Element in Array
**Java 8 Approach**: Using `Arrays.stream()` and `min()`

```java
import java.util.*;
import java.util.stream.*;

public class SmallestInArrayJava8 {
    public static void main(String[] args) {
        int[] arr = {10, 5, 20, 90, 40};
        
        // Using Java 8 Streams
        int min = Arrays.stream(arr)
            .min()
            .orElse(Integer.MAX_VALUE);
        
        System.out.println("Smallest: " + min);
        
        // Find index of smallest element
        int minIndex = IntStream.range(0, arr.length)
            .reduce((i, j) -> arr[j] < arr[i] ? j : i)
            .orElse(-1);
        
        System.out.println("Smallest element at index: " + minIndex);
        
        // Find both min and max in one pass
        IntSummaryStatistics stats = Arrays.stream(arr)
            .summaryStatistics();
        
        System.out.println("Min: " + stats.getMin() + ", Max: " + stats.getMax());
        System.out.println("Average: " + stats.getAverage());
        System.out.println("Sum: " + stats.getSum());
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{10, 5, 20, 90, 40},
            new int[]{15, 25, 35, 45},
            new int[]{500, 400, 300, 200, 100}
        );
        
        Map<String, IntSummaryStatistics> arrayStats = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array).summaryStatistics()
            ));
        
        System.out.println("Array statistics: " + arrayStats);
        
        // Find arrays with minimum element
        Optional<int[]> arrayWithMin = arrays.stream()
            .min(Comparator.comparingInt(array -> Arrays.stream(array).min().orElse(Integer.MAX_VALUE)));
        
        arrayWithMin.ifPresent(array -> 
            System.out.println("Array with minimum element: " + Arrays.toString(array)));
    }
}
```

## 18. Reverse an Array
**Java 8 Approach**: Using `IntStream` and collectors

```java
import java.util.*;
import java.util.stream.*;

public class ReverseArrayJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        
        // Using Java 8 Streams
        int[] reversed = IntStream.range(0, arr.length)
            .map(i -> arr[arr.length - 1 - i])
            .toArray();
        
        System.out.println("Reversed: " + Arrays.toString(reversed));
        
        // Alternative: Using Collections.reverse with boxed stream
        Integer[] reversedBoxed = Arrays.stream(arr)
            .boxed()
            .collect(Collectors.collectingAndThen(
                Collectors.toList(),
                list -> {
                    Collections.reverse(list);
                    return list.toArray(new Integer[0]);
                }
            ));
        
        System.out.println("Reversed (boxed): " + Arrays.toString(reversedBoxed));
        
        // Reverse multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3, 4, 5},
            new int[]{10, 20, 30},
            new int[]{100, 200, 300, 400}
        );
        
        Map<String, int[]> reversedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> IntStream.range(0, array.length)
                    .map(i -> array[array.length - 1 - i])
                    .toArray()
            ));
        
        System.out.println("Reversed arrays:");
        reversedArrays.forEach((original, reversedArr) -> 
            System.out.println(original + " -> " + Arrays.toString(reversedArr)));
        
        // Check if array is palindrome
        boolean isPalindrome = IntStream.range(0, arr.length / 2)
            .allMatch(i -> arr[i] == arr[arr.length - 1 - i]);
        
        System.out.println("Is palindrome: " + isPalindrome);
        
        // Find all palindrome arrays
        List<int[]> palindromeArrays = arrays.stream()
            .filter(array -> IntStream.range(0, array.length / 2)
                .allMatch(i -> array[i] == array[array.length - 1 - i]))
            .collect(Collectors.toList());
        
        System.out.println("Palindrome arrays: " + 
            palindromeArrays.stream().map(Arrays::toString).collect(Collectors.toList()));
    }
}
```

## 19. Sort Array (Bubble Sort with Java 8)
**Java 8 Approach**: Using streams for comparison and swapping

```java
import java.util.*;
import java.util.stream.*;

public class BubbleSortJava8 {
    public static void main(String[] args) {
        int[] arr = {64, 34, 25, 12, 22, 11, 90};
        
        // Traditional bubble sort with Java 8 style
        IntStream.range(0, arr.length - 1)
            .forEach(i -> 
                IntStream.range(0, arr.length - i - 1)
                    .filter(j -> arr[j] > arr[j + 1])
                    .forEach(j -> {
                        int temp = arr[j];
                        arr[j] = arr[j + 1];
                        arr[j + 1] = temp;
                    })
            );
        
        System.out.println("Sorted: " + Arrays.toString(arr));
        
        // Using Java 8 built-in sort (more efficient)
        int[] arr2 = {64, 34, 25, 12, 22, 11, 90};
        Arrays.sort(arr2);
        System.out.println("Sorted (Arrays.sort): " + Arrays.toString(arr2));
        
        // Sort using streams (creates new array)
        int[] arr3 = {64, 34, 25, 12, 22, 11, 90};
        int[] sortedStream = Arrays.stream(arr3).sorted().toArray();
        System.out.println("Sorted (stream): " + Arrays.toString(sortedStream));
        
        // Sort in descending order
        Integer[] arr4 = {64, 34, 25, 12, 22, 11, 90};
        Arrays.sort(arr4, Collections.reverseOrder());
        System.out.println("Sorted descending: " + Arrays.toString(arr4));
        
        // Sort multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{64, 34, 25, 12, 22, 11, 90},
            new int[]{5, 2, 8, 1, 9},
            new int[]{100, 50, 75, 25}
        );
        
        Map<String, int[]> sortedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array).sorted().toArray()
            ));
        
        System.out.println("Sorted arrays:");
        sortedArrays.forEach((original, sorted) -> 
            System.out.println(original + " -> " + Arrays.toString(sorted)));
        
        // Parallel sort for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 1000))
            .limit(100000)
            .toArray();
        
        long startTime = System.currentTimeMillis();
        Arrays.parallelSort(largeArray);
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel sort time for 100k elements: " + parallelTime + "ms");
    }
}
```

## 20. Linear Search
**Java 8 Approach**: Using `IntStream` and `anyMatch()`

```java
import java.util.*;
import java.util.stream.*;

public class LinearSearchJava8 {
    public static void main(String[] args) {
        int[] arr = {10, 20, 30, 40, 50};
        int key = 30;
        
        // Using Java 8 Streams
        OptionalInt index = IntStream.range(0, arr.length)
            .filter(i -> arr[i] == key)
            .findFirst();
        
        System.out.println(key + " found at index: " + index.orElse(-1));
        
        // Alternative: Using anyMatch
        boolean found = Arrays.stream(arr)
            .anyMatch(num -> num == key);
        
        System.out.println("Found " + key + "? " + found);
        
        // Find all occurrences
        List<Integer> indices = IntStream.range(0, arr.length)
            .filter(i -> arr[i] == key)
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("All indices of " + key + ": " + indices);
        
        // Search multiple keys
        List<Integer> keys = Arrays.asList(10, 20, 30, 40, 50, 60);
        Map<Integer, OptionalInt> searchResults = keys.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                k -> IntStream.range(0, arr.length)
                    .filter(i -> arr[i] == k)
                    .findFirst()
            ));
        
        System.out.println("Search results:");
        searchResults.forEach((k, idx) -> 
            System.out.println(k + " -> " + (idx.isPresent() ? idx.getAsInt() : "Not found")));
        
        // Search in multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{10, 20, 30, 40, 50},
            new int[]{5, 15, 25, 35},
            new int[]{100, 200, 300}
        );
        
        Map<String, Boolean> foundInArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array).anyMatch(num -> num == key)
            ));
        
        System.out.println("Arrays containing " + key + ": " + foundInArrays);
        
        // Parallel search for large arrays
        int[] largeArray = IntStream.range(1, 1000001).toArray();
        int searchKey = 999999;
        
        long startTime = System.currentTimeMillis();
        OptionalInt parallelIndex = IntStream.range(0, largeArray.length)
            .parallel()
            .filter(i -> largeArray[i] == searchKey)
            .findFirst();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel search time: " + parallelTime + "ms");
        System.out.println("Found " + searchKey + " at index: " + parallelIndex.orElse(-1));
    }
}
```

## 21. Binary Search
**Java 8 Approach**: Using `Arrays.binarySearch()` and streams

```java
import java.util.*;
import java.util.stream.*;

public class BinarySearchJava8 {
    public static void main(String[] args) {
        int[] arr = {10, 20, 30, 40, 50};
        int key = 40;
        
        // Using Java 8 built-in binary search
        int index = Arrays.binarySearch(arr, key);
        System.out.println(key + " found at index: " + (index >= 0 ? index : "Not found"));
        
        // Custom binary search using streams
        OptionalInt customIndex = IntStream.range(0, arr.length)
            .filter(i -> arr[i] == key)
            .findFirst();
        
        System.out.println(key + " found at index (custom): " + customIndex.orElse(-1));
        
        // Search multiple keys
        List<Integer> keys = Arrays.asList(10, 20, 30, 40, 50, 60);
        Map<Integer, Integer> searchResults = keys.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                k -> Arrays.binarySearch(arr, k)
            ));
        
        System.out.println("Binary search results:");
        searchResults.forEach((k, idx) -> 
            System.out.println(k + " -> " + (idx >= 0 ? idx : "Not found")));
        
        // Check if array is sorted
        boolean isSorted = IntStream.range(0, arr.length - 1)
            .allMatch(i -> arr[i] <= arr[i + 1]);
        
        System.out.println("Is array sorted: " + isSorted);
        
        // Find insertion point for non-existent elements
        List<Integer> nonExistentKeys = Arrays.asList(25, 35, 45);
        Map<Integer, Integer> insertionPoints = nonExistentKeys.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                k -> {
                    int idx = Arrays.binarySearch(arr, k);
                    return idx >= 0 ? idx : -idx - 1; // Convert to insertion point
                }
            ));
        
        System.out.println("Insertion points: " + insertionPoints);
        
        // Binary search in multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{10, 20, 30, 40, 50},
            new int[]{5, 15, 25, 35, 45},
            new int[]{100, 200, 300, 400, 500}
        );
        
        Map<String, Integer> resultsInArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> {
                    int idx = Arrays.binarySearch(array, key);
                    return idx >= 0 ? idx : -1;
                }
            ));
        
        System.out.println("Arrays containing " + key + ":");
        resultsInArrays.forEach((arrayStr, idx) -> 
            System.out.println(arrayStr + " -> " + (idx >= 0 ? "Index " + idx : "Not found")));
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.range(0, 1000000).toArray();
        int searchKeyLarge = 999999;
        
        long startTime = System.currentTimeMillis();
        int binaryResult = Arrays.binarySearch(largeArray, searchKeyLarge);
        long binaryTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        OptionalInt linearResult = IntStream.range(0, largeArray.length)
            .parallel()
            .filter(i -> largeArray[i] == searchKeyLarge)
            .findFirst();
        long linearTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Binary search time: " + binaryTime + "ms");
        System.out.println("Parallel linear search time: " + linearTime + "ms");
    }
}
```

## 22. Remove Duplicates from Sorted Array
**Java 8 Approach**: Using `IntStream` and `distinct()`

```java
import java.util.*;
import java.util.stream.*;

public class RemoveDuplicatesJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 1, 2, 2, 3, 4, 4, 5};
        
        // Using Java 8 Streams distinct()
        int[] unique = Arrays.stream(arr).distinct().toArray();
        System.out.println("After removing duplicates: " + Arrays.toString(unique));
        
        // Traditional approach with Java 8 style
        int j = 0;
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] != arr[i + 1]) {
                arr[j++] = arr[i];
            }
        }
        arr[j++] = arr[arr.length - 1];
        
        System.out.println("In-place result (first " + j + " elements): " + 
            Arrays.toString(Arrays.copyOf(arr, j)));
        
        // Count duplicates
        long duplicateCount = IntStream.range(0, arr.length - 1)
            .filter(i -> arr[i] == arr[i + 1])
            .count();
        
        System.out.println("Number of duplicates: " + duplicateCount);
        
        // Find unique elements and their counts
        Map<Integer, Long> elementCounts = Arrays.stream(arr)
            .boxed()
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        
        System.out.println("Element counts: " + elementCounts);
        
        // Remove duplicates from multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 1, 2, 2, 3, 4, 4, 5},
            new int[]{10, 10, 20, 30, 30, 40},
            new int[]{5, 5, 5, 5, 5}
        );
        
        Map<String, int[]> uniqueArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array).distinct().toArray()
            ));
        
        System.out.println("Unique arrays:");
        uniqueArrays.forEach((original, uniqueArr) -> 
            System.out.println(original + " -> " + Arrays.toString(uniqueArr)));
        
        // Find arrays with no duplicates
        List<int[]> noDuplicateArrays = arrays.stream()
            .filter(array -> Arrays.stream(array).distinct().count() == array.length)
            .collect(Collectors.toList());
        
        System.out.println("Arrays with no duplicates: " + 
            noDuplicateArrays.stream().map(Arrays::toString).collect(Collectors.toList()));
        
        // Parallel processing for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 100))
            .limit(1000000)
            .sorted()
            .toArray();
        
        long startTime = System.currentTimeMillis();
        int[] largeUnique = Arrays.stream(largeArray).parallel().distinct().toArray();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel distinct processing time: " + parallelTime + "ms");
        System.out.println("Original size: " + largeArray.length + ", Unique size: " + largeUnique.length);
    }
}
```

## 23. Second Largest Number
**Java 8 Approach**: Using `Stream.distinct()` and `skip()`

```java
import java.util.*;
import java.util.stream.*;

public class SecondLargestJava8 {
    public static void main(String[] args) {
        int[] arr = {12, 35, 1, 10, 34, 1};
        
        // Using Java 8 Streams
        OptionalInt secondLargest = Arrays.stream(arr)
            .distinct()
            .sorted()
            .skip(arr.length - 2)
            .findFirst();
        
        System.out.println("Second Largest: " + secondLargest.orElse(-1));
        
        // Alternative: Using reverse order
        OptionalInt secondLargestDesc = Arrays.stream(arr)
            .distinct()
            .boxed()
            .sorted(Comparator.reverseOrder())
            .skip(1)
            .mapToInt(Integer::intValue)
            .findFirst();
        
        System.out.println("Second Largest (desc): " + secondLargestDesc.orElse(-1));
        
        // Find top 3 largest
        int[] top3 = Arrays.stream(arr)
            .distinct()
            .boxed()
            .sorted(Comparator.reverseOrder())
            .limit(3)
            .mapToInt(Integer::intValue)
            .toArray();
        
        System.out.println("Top 3 largest: " + Arrays.toString(top3));
        
        // Find second smallest
        OptionalInt secondSmallest = Arrays.stream(arr)
            .distinct()
            .sorted()
            .skip(1)
            .findFirst();
        
        System.out.println("Second Smallest: " + secondSmallest.orElse(-1));
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{12, 35, 1, 10, 34, 1},
            new int[]{5, 8, 3, 9, 1},
            new int[]{100, 50, 75, 25, 125}
        );
        
        Map<String, Integer> secondLargestResults = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array)
                    .distinct()
                    .boxed()
                    .sorted(Comparator.reverseOrder())
                    .skip(1)
                    .findFirst()
                    .orElse(-1)
            ));
        
        System.out.println("Second largest results:");
        secondLargestResults.forEach((arrayStr, secondLargestVal) -> 
            System.out.println(arrayStr + " -> " + secondLargestVal));
        
        // Find arrays with insufficient unique elements
        List<String> insufficientArrays = arrays.stream()
            .filter(array -> Arrays.stream(array).distinct().count() < 2)
            .map(Arrays::toString)
            .collect(Collectors.toList());
        
        System.out.println("Arrays with insufficient unique elements: " + insufficientArrays);
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 1000000))
            .limit(100000)
            .toArray();
        
        long startTime = System.currentTimeMillis();
        OptionalInt largeSecondLargest = Arrays.stream(largeArray)
            .parallel()
            .distinct()
            .boxed()
            .sorted(Comparator.reverseOrder())
            .skip(1)
            .findFirst();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel second largest time: " + parallelTime + "ms");
        System.out.println("Second largest in large array: " + largeSecondLargest.orElse(-1));
    }
}
```

## 24. Missing Number in Array
**Java 8 Approach**: Using `IntStream` arithmetic operations

```java
import java.util.*;
import java.util.stream.*;

public class MissingNumberJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 4, 5, 6};
        int n = 6; // Max number
        
        // Using Java 8 Streams - arithmetic method
        int expectedSum = IntStream.rangeClosed(1, n).sum();
        int actualSum = Arrays.stream(arr).sum();
        int missing = expectedSum - actualSum;
        
        System.out.println("Missing Number: " + missing);
        
        // Alternative: Using XOR method
        int xorAll = IntStream.rangeClosed(1, n).reduce(0, (a, b) -> a ^ b);
        int xorArr = Arrays.stream(arr).reduce(0, (a, b) -> a ^ b);
        int missingXor = xorAll ^ xorArr;
        
        System.out.println("Missing Number (XOR): " + missingXor);
        
        // Find missing numbers in multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 4, 5, 6},
            new int[]{2, 3, 5, 6},
            new int[]{1, 3, 4, 5}
        );
        
        Map<String, Integer> missingNumbers = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> {
                    int max = array.length + 1;
                    return IntStream.rangeClosed(1, max).sum() - Arrays.stream(array).sum();
                }
            ));
        
        System.out.println("Missing numbers:");
        missingNumbers.forEach((arrayStr, missingNum) -> 
            System.out.println(arrayStr + " -> " + missingNum));
        
        // Find all missing numbers (when multiple are missing)
        int[] arrWithMultipleMissing = {1, 2, 5, 7, 8};
        int maxNum = 8;
        
        List<Integer> allMissing = IntStream.rangeClosed(1, maxNum)
            .filter(num -> Arrays.stream(arrWithMultipleMissing).noneMatch(arrNum -> arrNum == num))
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("All missing numbers: " + allMissing);
        
        // Check if array is complete (no missing numbers)
        boolean isComplete = IntStream.rangeClosed(1, arr.length)
            .allMatch(num -> Arrays.stream(arr).anyMatch(arrNum -> arrNum == num));
        
        System.out.println("Is array complete: " + isComplete);
        
        // Find missing number in unsorted array
        int[] unsortedArr = {5, 3, 1, 2, 6};
        int missingUnsorted = IntStream.rangeClosed(1, unsortedArr.length + 1)
            .filter(num -> Arrays.stream(unsortedArr).noneMatch(arrNum -> arrNum == num))
            .findFirst()
            .orElse(-1);
        
        System.out.println("Missing in unsorted array: " + missingUnsorted);
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.range(1, 100001)
            .filter(num -> num != 50000) // Remove one number
            .toArray();
        
        long startTime = System.currentTimeMillis();
        int largeMissing = IntStream.rangeClosed(1, 100001).sum() - Arrays.stream(largeArray).sum();
        long arithmeticTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        int largeMissingParallel = IntStream.rangeClosed(1, 100001).parallel().sum() - 
                                 Arrays.stream(largeArray).parallel().sum();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Arithmetic method time: " + arithmeticTime + "ms");
        System.out.println("Parallel method time: " + parallelTime + "ms");
        System.out.println("Missing in large array: " + largeMissing);
    }
}
```

## 25. Merge Two Arrays
**Java 8 Approach**: Using `Stream.concat()` and collectors

```java
import java.util.*;
import java.util.stream.*;

public class MergeArraysJava8 {
    public static void main(String[] args) {
        int[] a = {1, 2, 3};
        int[] b = {4, 5, 6};
        
        // Using Java 8 Streams
        int[] merged = IntStream.concat(Arrays.stream(a), Arrays.stream(b))
            .toArray();
        
        System.out.println("Merged: " + Arrays.toString(merged));
        
        // Alternative: Using flatMap
        int[] mergedFlatMap = IntStream.of(a, b)
            .flatMapToInt(Arrays::stream)
            .toArray();
        
        System.out.println("Merged (flatMap): " + Arrays.toString(mergedFlatMap));
        
        // Merge and sort
        int[] mergedSorted = IntStream.concat(Arrays.stream(a), Arrays.stream(b))
            .sorted()
            .toArray();
        
        System.out.println("Merged and sorted: " + Arrays.toString(mergedSorted));
        
        // Merge and remove duplicates
        int[] mergedDistinct = IntStream.concat(Arrays.stream(a), Arrays.stream(b))
            .distinct()
            .toArray();
        
        System.out.println("Merged and distinct: " + Arrays.toString(mergedDistinct));
        
        // Merge multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3},
            new int[]{4, 5, 6},
            new int[]{7, 8, 9}
        );
        
        int[] multiMerged = arrays.stream()
            .flatMapToInt(Arrays::stream)
            .toArray();
        
        System.out.println("Multiple arrays merged: " + Arrays.toString(multiMerged));
        
        // Merge with interleaving
        int[] interleaved = IntStream.range(0, Math.max(a.length, b.length))
            .flatMap(i -> IntStream.of(
                i < a.length ? a[i] : -1,
                i < b.length ? b[i] : -1
            ))
            .filter(val -> val != -1)
            .toArray();
        
        System.out.println("Interleaved: " + Arrays.toString(interleaved));
        
        // Merge arrays of different sizes and pad with zeros
        int maxSize = Math.max(a.length, b.length);
        int[] paddedMerge = IntStream.range(0, maxSize)
            .flatMap(i -> IntStream.of(
                i < a.length ? a[i] : 0,
                i < b.length ? b[i] : 0
            ))
            .toArray();
        
        System.out.println("Padded merge: " + Arrays.toString(paddedMerge));
        
        // Calculate statistics of merged array
        IntSummaryStatistics stats = IntStream.concat(Arrays.stream(a), Arrays.stream(b))
            .summaryStatistics();
        
        System.out.println("Merged array statistics:");
        System.out.println("Count: " + stats.getCount());
        System.out.println("Sum: " + stats.getSum());
        System.out.println("Average: " + stats.getAverage());
        System.out.println("Min: " + stats.getMin());
        System.out.println("Max: " + stats.getMax());
        
        // Parallel merge for large arrays
        int[] largeA = IntStream.range(0, 50000).toArray();
        int[] largeB = IntStream.range(50000, 100000).toArray();
        
        long startTime = System.currentTimeMillis();
        int[] largeMerged = IntStream.concat(
            Arrays.stream(largeA).parallel(),
            Arrays.stream(largeB).parallel()
        ).toArray();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel merge time for 100k elements: " + parallelTime + "ms");
        System.out.println("Large merged array size: " + largeMerged.length);
    }
}
```

## 26. Check if Arrays are Equal
**Java 8 Approach**: Using `Arrays.equals()` and stream comparison

```java
import java.util.*;
import java.util.stream.*;

public class ArrayEqualityJava8 {
    public static void main(String[] args) {
        int[] a = {1, 2, 3};
        int[] b = {1, 2, 3};
        int[] c = {1, 2, 4};
        
        // Using Java 8 built-in method
        boolean equalAB = Arrays.equals(a, b);
        boolean equalAC = Arrays.equals(a, c);
        
        System.out.println("a equals b? " + equalAB);
        System.out.println("a equals c? " + equalAC);
        
        // Custom comparison using streams
        boolean customEqual = a.length == b.length && 
            IntStream.range(0, a.length)
                .allMatch(i -> a[i] == b[i]);
        
        System.out.println("Custom comparison a equals b? " + customEqual);
        
        // Compare multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3},
            new int[]{1, 2, 3},
            new int[]{1, 2, 4},
            new int[]{1, 2, 3, 4}
        );
        
        Map<String, Boolean> equalityResults = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.equals(array, a)
            ));
        
        System.out.println("Equality with array a:");
        equalityResults.forEach((arrayStr, isEqual) -> 
            System.out.println(arrayStr + " -> " + isEqual));
        
        // Find all equal arrays
        List<int[]> equalArrays = arrays.stream()
            .filter(array -> Arrays.equals(array, a))
            .collect(Collectors.toList());
        
        System.out.println("Arrays equal to a: " + 
            equalArrays.stream().map(Arrays::toString).collect(Collectors.toList()));
        
        // Compare arrays with different orders
        Integer[] aBoxed = {1, 2, 3};
        Integer[] bBoxed = {3, 2, 1};
        
        boolean equalUnordered = Arrays.asList(aBoxed).containsAll(Arrays.asList(bBoxed)) &&
                               Arrays.asList(bBoxed).containsAll(Arrays.asList(aBoxed));
        
        System.out.println("a equals b (unordered)? " + equalUnordered);
        
        // Compare using sorted streams
        boolean equalSorted = Arrays.stream(aBoxed).sorted().toArray().length ==
                               Arrays.stream(bBoxed).sorted().toArray().length &&
                               IntStream.range(0, aBoxed.length)
                                   .allMatch(i -> Arrays.stream(aBoxed).sorted().toArray()[i] == 
                                                  Arrays.stream(bBoxed).sorted().toArray()[i]);
        
        System.out.println("a equals b (sorted comparison)? " + equalSorted);
        
        // Deep equality for 2D arrays
        int[][] a2D = {{1, 2}, {3, 4}};
        int[][] b2D = {{1, 2}, {3, 4}};
        int[][] c2D = {{1, 2}, {4, 3}};
        
        boolean deepEqualAB = Arrays.deepEquals(a2D, b2D);
        boolean deepEqualAC = Arrays.deepEquals(a2D, c2D);
        
        System.out.println("2D arrays a equals b? " + deepEqualAB);
        System.out.println("2D arrays a equals c? " + deepEqualAC);
        
        // Performance comparison for large arrays
        int[] largeA = IntStream.range(0, 1000000).toArray();
        int[] largeB = IntStream.range(0, 1000000).toArray();
        int[] largeC = IntStream.range(0, 1000000).toArray();
        largeC[999999] = 9999999; // Make one different
        
        long startTime = System.currentTimeMillis();
        boolean largeEqual = Arrays.equals(largeA, largeB);
        long builtInTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        boolean largeEqualStream = largeA.length == largeB.length && 
            IntStream.range(0, largeA.length).parallel()
                .allMatch(i -> largeA[i] == largeB[i]);
        long streamTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Built-in equals time: " + builtInTime + "ms");
        System.out.println("Stream equals time: " + streamTime + "ms");
        System.out.println("Large arrays equal? " + largeEqual);
    }
}
```

## 27. Find Common Elements
**Java 8 Approach**: Using `Stream.filter()` and sets

```java
import java.util.*;
import java.util.stream.*;

public class CommonElementsJava8 {
    public static void main(String[] args) {
        int[] arr1 = {1, 2, 3, 4};
        int[] arr2 = {3, 4, 5, 6};
        
        // Using Java 8 Streams with Set
        Set<Integer> set1 = Arrays.stream(arr1).boxed().collect(Collectors.toSet());
        int[] common = Arrays.stream(arr2)
            .filter(set1::contains)
            .distinct()
            .toArray();
        
        System.out.println("Common elements: " + Arrays.toString(common));
        
        // Alternative: Using streams without Set
        int[] commonAlt = Arrays.stream(arr1)
            .flatMap(num1 -> Arrays.stream(arr2)
                .filter(num2 -> num2 == num1)
                .distinct())
            .distinct()
            .toArray();
        
        System.out.println("Common elements (alternative): " + Arrays.toString(commonAlt));
        
        // Find common elements with counts
        Map<Integer, Long> arr1Counts = Arrays.stream(arr1).boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        Map<Integer, Long> arr2Counts = Arrays.stream(arr2).boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        
        Map<Integer, Long> commonWithCounts = arr1Counts.entrySet().stream()
            .filter(entry -> arr2Counts.containsKey(entry.getKey()))
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                entry -> Math.min(entry.getValue(), arr2Counts.get(entry.getKey()))
            ));
        
        System.out.println("Common elements with counts: " + commonWithCounts);
        
        // Find common elements in multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3, 4, 5},
            new int[]{3, 4, 5, 6, 7},
            new int[]{4, 5, 6, 7, 8},
            new int[]{5, 6, 7, 8, 9}
        );
        
        List<Integer> commonInAll = arrays.stream()
            .flatMap(Arrays::stream)
            .boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()))
            .entrySet().stream()
            .filter(entry -> entry.getValue() == arrays.size())
            .map(Map.Entry::getKey)
            .collect(Collectors.toList());
        
        System.out.println("Common in all arrays: " + commonInAll);
        
        // Find elements common to exactly two arrays
        Map<Integer, Long> frequencyMap = arrays.stream()
            .flatMap(Arrays::stream)
            .boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        
        List<Integer> commonInTwo = frequencyMap.entrySet().stream()
            .filter(entry -> entry.getValue() == 2)
            .map(Map.Entry::getKey)
            .collect(Collectors.toList());
        
        System.out.println("Common in exactly two arrays: " + commonInTwo);
        
        // Find intersection using parallel streams
        int[] largeArr1 = IntStream.range(0, 100000).toArray();
        int[] largeArr2 = IntStream.range(50000, 150000).toArray();
        
        long startTime = System.currentTimeMillis();
        Set<Integer> largeSet1 = Arrays.stream(largeArr1).parallel().boxed()
            .collect(Collectors.toSet());
        int[] largeCommon = Arrays.stream(largeArr2).parallel()
            .filter(largeSet1::contains)
            .distinct()
            .toArray();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel intersection time: " + parallelTime + "ms");
        System.out.println("Large common count: " + largeCommon.length);
        
        // Find symmetric difference (elements in one but not both)
        Set<Integer> set2 = Arrays.stream(arr2).boxed().collect(Collectors.toSet());
        List<Integer> symmetricDifference = Stream.concat(
                Arrays.stream(arr1).boxed().filter(num -> !set2.contains(num)),
                Arrays.stream(arr2).boxed().filter(num -> !set1.contains(num))
            )
            .distinct()
            .collect(Collectors.toList());
        
        System.out.println("Symmetric difference: " + symmetricDifference);
        
        // Check if arrays have any common elements
        boolean hasCommon = Arrays.stream(arr1)
            .anyMatch(num1 -> Arrays.stream(arr2).anyMatch(num2 -> num2 == num1));
        
        System.out.println("Arrays have common elements? " + hasCommon);
    }
}
```

## 28. Left Rotate Array
**Java 8 Approach**: Using `IntStream` range manipulation

```java
import java.util.*;
import java.util.stream.*;

public class LeftRotateJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        
        // Using Java 8 Streams for single rotation
        int[] rotated = IntStream.range(0, arr.length)
            .map(i -> arr[(i + 1) % arr.length])
            .toArray();
        
        System.out.println("Left rotated by 1: " + Arrays.toString(rotated));
        
        // Rotate by k positions
        int k = 2;
        int[] rotatedByK = IntStream.range(0, arr.length)
            .map(i -> arr[(i + k) % arr.length])
            .toArray();
        
        System.out.println("Left rotated by " + k + ": " + Arrays.toString(rotatedByK));
        
        // Alternative using concat
        int[] rotatedConcat = IntStream.concat(
            Arrays.stream(arr).skip(k),
            Arrays.stream(arr).limit(k)
        ).toArray();
        
        System.out.println("Rotated (concat): " + Arrays.toString(rotatedConcat));
        
        // Rotate multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3, 4, 5},
            new int[]{10, 20, 30, 40},
            new int[]{100, 200, 300, 400, 500, 600}
        );
        
        Map<String, int[]> rotatedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> IntStream.range(0, array.length)
                    .map(i -> array[(i + 1) % array.length])
                    .toArray()
            ));
        
        System.out.println("Rotated arrays:");
        rotatedArrays.forEach((original, rotatedArr) -> 
            System.out.println(original + " -> " + Arrays.toString(rotatedArr)));
        
        // Rotate by different amounts
        List<Integer> rotations = Arrays.asList(1, 2, 3, 4);
        Map<Integer, int[]> rotationResults = rotations.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                rot -> IntStream.range(0, arr.length)
                    .map(i -> arr[(i + rot) % arr.length])
                    .toArray()
            ));
        
        System.out.println("Array rotated by different amounts:");
        rotationResults.forEach((rot, rotatedArr) -> 
            System.out.println("By " + rot + ": " + Arrays.toString(rotatedArr)));
        
        // Check if array is rotation of another
        int[] original = {1, 2, 3, 4, 5};
        int[] candidate = {3, 4, 5, 1, 2};
        
        boolean isRotation = IntStream.range(0, original.length)
            .anyMatch(i -> {
                return IntStream.range(0, original.length)
                    .allMatch(j -> original[j] == candidate[(i + j) % candidate.length]);
            });
        
        System.out.println("Candidate is rotation of original? " + isRotation);
        
        // Find all rotations of array
        List<int[]> allRotations = IntStream.range(0, arr.length)
            .mapToObj(i -> IntStream.range(0, arr.length)
                .map(j -> arr[(i + j) % arr.length])
                .toArray())
            .collect(Collectors.toList());
        
        System.out.println("All rotations:");
        allRotations.forEach(rotation -> System.out.println(Arrays.toString(rotation)));
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.range(0, 100000).toArray();
        int largeK = 1000;
        
        long startTime = System.currentTimeMillis();
        int[] largeRotated = IntStream.range(0, largeArray.length)
            .parallel()
            .map(i -> largeArray[(i + largeK) % largeArray.length])
            .toArray();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel rotation time for 100k elements: " + parallelTime + "ms");
        
        // Verify rotation correctness
        boolean rotationCorrect = IntStream.range(0, Math.min(10, largeArray.length))
            .allMatch(i -> largeRotated[i] == largeArray[(i + largeK) % largeArray.length]);
        
        System.out.println("Large rotation correct? " + rotationCorrect);
    }
}
```

## 29. Right Rotate Array
**Java 8 Approach**: Using `IntStream` with modular arithmetic

```java
import java.util.*;
import java.util.stream.*;

public class RightRotateJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        
        // Using Java 8 Streams for single rotation
        int[] rotated = IntStream.range(0, arr.length)
            .map(i -> arr[(i - 1 + arr.length) % arr.length])
            .toArray();
        
        System.out.println("Right rotated by 1: " + Arrays.toString(rotated));
        
        // Rotate by k positions
        int k = 2;
        int[] rotatedByK = IntStream.range(0, arr.length)
            .map(i -> arr[(i - k + arr.length) % arr.length])
            .toArray();
        
        System.out.println("Right rotated by " + k + ": " + Arrays.toString(rotatedByK));
        
        // Alternative using concat
        int[] rotatedConcat = IntStream.concat(
            Arrays.stream(arr).skip(arr.length - k),
            Arrays.stream(arr).limit(arr.length - k)
        ).toArray();
        
        System.out.println("Rotated (concat): " + Arrays.toString(rotatedConcat));
        
        // Rotate multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3, 4, 5},
            new int[]{10, 20, 30, 40},
            new int[]{100, 200, 300, 400, 500, 600}
        );
        
        Map<String, int[]> rotatedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> IntStream.range(0, array.length)
                    .map(i -> array[(i - 1 + array.length) % array.length])
                    .toArray()
            ));
        
        System.out.println("Right rotated arrays:");
        rotatedArrays.forEach((original, rotatedArr) -> 
            System.out.println(original + " -> " + Arrays.toString(rotatedArr)));
        
        // Rotate by different amounts
        List<Integer> rotations = Arrays.asList(1, 2, 3, 4);
        Map<Integer, int[]> rotationResults = rotations.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                rot -> IntStream.range(0, arr.length)
                    .map(i -> arr[(i - rot + arr.length) % arr.length])
                    .toArray()
            ));
        
        System.out.println("Array right rotated by different amounts:");
        rotationResults.forEach((rot, rotatedArr) -> 
            System.out.println("By " + rot + ": " + Arrays.toString(rotatedArr)));
        
        // Compare left and right rotation
        int[] leftRotated = IntStream.range(0, arr.length)
            .map(i -> arr[(i + 1) % arr.length])
            .toArray();
        
        int[] rightRotated = IntStream.range(0, arr.length)
            .map(i -> arr[(i - 1 + arr.length) % arr.length])
            .toArray();
        
        boolean leftEqualsRight = Arrays.equals(leftRotated, rightRotated);
        System.out.println("Left rotation equals right rotation? " + leftEqualsRight);
        
        // Find array that when right rotated equals left rotated original
        int[] testArr = {1, 2, 3, 4, 5};
        int[] leftTest = IntStream.range(0, testArr.length)
            .map(i -> testArr[(i + 1) % testArr.length])
            .toArray();
        
        OptionalInt matchingRotation = IntStream.range(1, testArr.length)
            .filter(i -> {
                int[] rightTest = IntStream.range(0, testArr.length)
                    .map(j -> testArr[(j - i + testArr.length) % testArr.length])
                    .toArray();
                return Arrays.equals(leftTest, rightTest);
            })
            .findFirst();
        
        System.out.println("Right rotation that matches left rotation: " + 
            matchingRotation.orElse(-1));
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.range(0, 100000).toArray();
        int largeK = 1000;
        
        long startTime = System.currentTimeMillis();
        int[] largeRotated = IntStream.range(0, largeArray.length)
            .parallel()
            .map(i -> largeArray[(i - largeK + largeArray.length) % largeArray.length])
            .toArray();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel right rotation time for 100k elements: " + parallelTime + "ms");
        
        // Verify rotation correctness
        boolean rotationCorrect = IntStream.range(0, Math.min(10, largeArray.length))
            .allMatch(i -> largeRotated[i] == largeArray[(i - largeK + largeArray.length) % largeArray.length]);
        
        System.out.println("Large rotation correct? " + rotationCorrect);
    }
}
```

## 30. Move Zeros to End
**Java 8 Approach**: Using `IntStream` filtering and concatenation

```java
import java.util.*;
import java.util.stream.*;

public class MoveZerosJava8 {
    public static void main(String[] args) {
        int[] arr = {0, 1, 0, 3, 12};
        
        // Using Java 8 Streams
        int[] nonZeros = Arrays.stream(arr).filter(num -> num != 0).toArray();
        int[] zeros = new int[arr.length - nonZeros.length];
        int[] result = IntStream.concat(Arrays.stream(nonZeros), Arrays.stream(zeros)).toArray();
        
        System.out.println("After moving zeros: " + Arrays.toString(result));
        
        // Alternative: Single pass with counting
        int[] resultAlt = new int[arr.length];
        int[] index = {0};
        
        Arrays.stream(arr).forEach(num -> {
            if (num != 0) resultAlt[index[0]++] = num;
        });
        
        System.out.println("Alternative result: " + Arrays.toString(resultAlt));
        
        // Count zeros
        long zeroCount = Arrays.stream(arr).filter(num -> num == 0).count();
        long nonZeroCount = Arrays.stream(arr).filter(num -> num != 0).count();
        
        System.out.println("Zeros: " + zeroCount + ", Non-zeros: " + nonZeroCount);
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{0, 1, 0, 3, 12},
            new int[]{0, 0, 0, 1, 2},
            new int[]{1, 2, 3, 0, 0, 0},
            new int[]{0, 0, 0, 0, 0}
        );
        
        Map<String, int[]> processedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> {
                    int[] nonZero = Arrays.stream(array).filter(num -> num != 0).toArray();
                    int[] zeroArray = new int[array.length - nonZero.length];
                    return IntStream.concat(Arrays.stream(nonZero), Arrays.stream(zeroArray)).toArray();
                }
            ));
        
        System.out.println("Processed arrays:");
        processedArrays.forEach((original, processed) -> 
            System.out.println(original + " -> " + Arrays.toString(processed)));
        
        // Find arrays with no zeros
        List<int[]> noZeroArrays = arrays.stream()
            .filter(array -> Arrays.stream(array).noneMatch(num -> num == 0))
            .collect(Collectors.toList());
        
        System.out.println("Arrays with no zeros: " + 
            noZeroArrays.stream().map(Arrays::toString).collect(Collectors.toList()));
        
        // Find arrays with only zeros
        List<int[]> onlyZeroArrays = arrays.stream()
            .filter(array -> Arrays.stream(array).allMatch(num -> num == 0))
            .collect(Collectors.toList());
        
        System.out.println("Arrays with only zeros: " + 
            onlyZeroArrays.stream().map(Arrays::toString).collect(Collectors.toList()));
        
        // Calculate zero percentage
        Map<String, Double> zeroPercentages = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array)
                    .filter(num -> num == 0)
                    .count() * 100.0 / array.length
            ));
        
        System.out.println("Zero percentages:");
        zeroPercentages.forEach((arrayStr, percentage) -> 
            System.out.println(arrayStr + " -> " + String.format("%.1f%%", percentage)));
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 10))
            .limit(1000000)
            .toArray();
        
        long startTime = System.currentTimeMillis();
        int[] largeNonZeros = Arrays.stream(largeArray).parallel()
            .filter(num -> num != 0)
            .toArray();
        int[] largeZeros = new int[largeArray.length - largeNonZeros.length];
        int[] largeResult = IntStream.concat(Arrays.stream(largeNonZeros), Arrays.stream(largeZeros)).toArray();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel processing time for 1M elements: " + parallelTime + "ms");
        System.out.println("Large array zeros moved: " + largeResult.length + " elements");
    }
}
```

## 31. Find Duplicate Elements
**Java 8 Approach**: Using `Collectors.groupingBy()` and `counting()`

```java
import java.util.*;
import java.util.stream.*;

public class FindDuplicatesJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 1, 4, 2};
        
        // Using Java 8 Streams with groupingBy
        Map<Integer, Long> frequency = Arrays.stream(arr)
            .boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        
        List<Integer> duplicates = frequency.entrySet().stream()
            .filter(entry -> entry.getValue() > 1)
            .map(Map.Entry::getKey)
            .collect(Collectors.toList());
        
        System.out.println("Duplicates: " + duplicates);
        
        // Alternative: Using Set
        Set<Integer> seen = new HashSet<>();
        Set<Integer> duplicatesSet = Arrays.stream(arr)
            .boxed()
            .filter(num -> !seen.add(num))
            .collect(Collectors.toSet());
        
        System.out.println("Duplicates (Set): " + duplicatesSet);
        
        // Find duplicates with counts
        Map<Integer, Long> duplicatesWithCounts = frequency.entrySet().stream()
            .filter(entry -> entry.getValue() > 1)
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                Map.Entry::getValue
            ));
        
        System.out.println("Duplicates with counts: " + duplicatesWithCounts);
        
        // Find all duplicates (including multiple occurrences)
        List<Integer> allDuplicates = Arrays.stream(arr)
            .boxed()
            .collect(Collectors.groupingBy(Function.identity()))
            .entrySet().stream()
            .filter(entry -> entry.getValue().size() > 1)
            .flatMap(entry -> entry.getValue().stream().skip(1))
            .collect(Collectors.toList());
        
        System.out.println("All duplicate occurrences: " + allDuplicates);
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3, 1, 4, 2},
            new int[]{5, 5, 5, 6, 7},
            new int[]{8, 9, 10, 11},
            new int[]{12, 13, 12, 14, 13, 15}
        );
        
        Map<String, List<Integer>> duplicateResults = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array)
                    .boxed()
                    .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()))
                    .entrySet().stream()
                    .filter(entry -> entry.getValue() > 1)
                    .map(Map.Entry::getKey)
                    .collect(Collectors.toList())
            ));
        
        System.out.println("Duplicates in multiple arrays:");
        duplicateResults.forEach((arrayStr, duplicatesList) -> 
            System.out.println(arrayStr + " -> " + duplicatesList));
        
        // Find arrays with no duplicates
        List<int[]> uniqueArrays = arrays.stream()
            .filter(array -> Arrays.stream(array)
                .boxed()
                .collect(Collectors.toSet())
                .size() == array.length)
            .collect(Collectors.toList());
        
        System.out.println("Arrays with no duplicates: " + 
            uniqueArrays.stream().map(Arrays::toString).collect(Collectors.toList()));
        
        // Find most frequent element
        Optional<Map.Entry<Integer, Long>> mostFrequent = frequency.entrySet().stream()
            .max(Map.Entry.comparingByValue());
        
        mostFrequent.ifPresent(entry -> 
            System.out.println("Most frequent: " + entry.getKey() + " (count: " + entry.getValue() + ")"));
        
        // Count total duplicates
        long totalDuplicates = frequency.values().stream()
            .mapToLong(count -> count > 1 ? count - 1 : 0)
            .sum();
        
        System.out.println("Total duplicate occurrences: " + totalDuplicates);
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 100))
            .limit(1000000)
            .toArray();
        
        long startTime = System.currentTimeMillis();
        Map<Integer, Long> largeFrequency = Arrays.stream(largeArray).parallel()
            .boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        long parallelTime = System.currentTimeMillis() - startTime;
        
        long largeDuplicates = largeFrequency.values().stream()
            .mapToLong(count -> count > 1 ? count - 1 : 0)
            .sum();
        
        System.out.println("Parallel duplicate detection time: " + parallelTime + "ms");
        System.out.println("Large array duplicates: " + largeDuplicates);
    }
}
```

## 32. Frequency of Each Element
**Java 8 Approach**: Using `Collectors.groupingBy()` and `counting()`

```java
import java.util.*;
import java.util.stream.*;

public class FrequencyJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 2, 3, 1, 4, 2};
        
        // Using Java 8 Streams
        Map<Integer, Long> frequency = Arrays.stream(arr)
            .boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        
        System.out.println("Frequency: " + frequency);
        
        // Sort by frequency
        Map<Integer, Long> sortedByFrequency = frequency.entrySet().stream()
            .sorted(Map.Entry.comparingByValue())
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                Map.Entry::getValue,
                (oldVal, newVal) -> oldVal,
                LinkedHashMap::new
            ));
        
        System.out.println("Sorted by frequency: " + sortedByFrequency);
        
        // Sort by value
        Map<Integer, Long> sortedByValue = frequency.entrySet().stream()
            .sorted(Map.Entry.comparingByKey())
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                Map.Entry::getValue,
                (oldVal, newVal) -> oldVal,
                LinkedHashMap::new
            ));
        
        System.out.println("Sorted by value: " + sortedByValue);
        
        // Find most and least frequent
        Optional<Map.Entry<Integer, Long>> mostFrequent = frequency.entrySet().stream()
            .max(Map.Entry.comparingByValue());
        
        Optional<Map.Entry<Integer, Long>> leastFrequent = frequency.entrySet().stream()
            .min(Map.Entry.comparingByValue());
        
        mostFrequent.ifPresent(entry -> 
            System.out.println("Most frequent: " + entry.getKey() + " (count: " + entry.getValue() + ")"));
        
        leastFrequent.ifPresent(entry -> 
            System.out.println("Least frequent: " + entry.getKey() + " (count: " + entry.getValue() + ")"));
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 2, 3, 1, 4, 2},
            new int[]{5, 5, 6, 7, 6, 5},
            new int[]{8, 9, 10, 8, 9, 8}
        );
        
        Map<String, Map<Integer, Long>> frequencyMaps = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array)
                    .boxed()
                    .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()))
            ));
        
        System.out.println("Frequency maps:");
        frequencyMaps.forEach((arrayStr, freqMap) -> 
            System.out.println(arrayStr + " -> " + freqMap));
        
        // Find elements with specific frequency
        long targetFrequency = 2;
        List<Integer> elementsWithTargetFreq = frequency.entrySet().stream()
            .filter(entry -> entry.getValue() == targetFrequency)
            .map(Map.Entry::getKey)
            .collect(Collectors.toList());
        
        System.out.println("Elements appearing " + targetFrequency + " times: " + elementsWithTargetFreq);
        
        // Calculate frequency statistics
        DoubleSummaryStatistics freqStats = frequency.values().stream()
            .mapToDouble(Long::doubleValue)
            .summaryStatistics();
        
        System.out.println("Frequency statistics:");
        System.out.println("Average frequency: " + freqStats.getAverage());
        System.out.println("Max frequency: " + freqStats.getMax());
        System.out.println("Min frequency: " + freqStats.getMin());
        
        // Find frequency distribution
        Map<Long, List<Integer>> frequencyDistribution = frequency.entrySet().stream()
            .collect(Collectors.groupingBy(
                Map.Entry::getValue,
                Collectors.mapping(Map.Entry::getKey, Collectors.toList())
            ));
        
        System.out.println("Frequency distribution: " + frequencyDistribution);
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 100))
            .limit(1000000)
            .toArray();
        
        long startTime = System.currentTimeMillis();
        Map<Integer, Long> largeFrequency = Arrays.stream(largeArray).parallel()
            .boxed()
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel frequency calculation time: " + parallelTime + "ms");
        System.out.println("Unique elements in large array: " + largeFrequency.size());
        
        // Find top 10 most frequent elements in large array
        List<Map.Entry<Integer, Long>> top10 = largeFrequency.entrySet().stream()
            .sorted(Map.Entry.<Integer, Long>comparingByValue().reversed())
            .limit(10)
            .collect(Collectors.toList());
        
        System.out.println("Top 10 most frequent elements:");
        top10.forEach(entry -> System.out.println(entry.getKey() + ": " + entry.getValue()));
    }
}
```

## 33. Odd and Even Numbers in Array
**Java 8 Approach**: Using `Collectors.partitioningBy()`

```java
import java.util.*;
import java.util.stream.*;

public class OddEvenArrayJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 5, 6, 3, 2};
        
        // Using Java 8 Streams with partitioningBy
        Map<Boolean, List<Integer>> partitioned = Arrays.stream(arr)
            .boxed()
            .collect(Collectors.partitioningBy(num -> num % 2 == 0));
        
        System.out.println("Even: " + partitioned.get(true));
        System.out.println("Odd: " + partitioned.get(false));
        
        // Alternative: Using separate filters
        List<Integer> evens = Arrays.stream(arr)
            .filter(num -> num % 2 == 0)
            .boxed()
            .collect(Collectors.toList());
        
        List<Integer> odds = Arrays.stream(arr)
            .filter(num -> num % 2 != 0)
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("Even (filter): " + evens);
        System.out.println("Odd (filter): " + odds);
        
        // Count odds and evens
        long evenCount = Arrays.stream(arr).filter(num -> num % 2 == 0).count();
        long oddCount = Arrays.stream(arr).filter(num -> num % 2 != 0).count();
        
        System.out.println("Even count: " + evenCount);
        System.out.println("Odd count: " + oddCount);
        
        // Sum of odds and evens
        int evenSum = Arrays.stream(arr).filter(num -> num % 2 == 0).sum();
        int oddSum = Arrays.stream(arr).filter(num -> num % 2 != 0).sum();
        
        System.out.println("Even sum: " + evenSum);
        System.out.println("Odd sum: " + oddSum);
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 5, 6, 3, 2},
            new int[]{10, 15, 20, 25},
            new int[]{1, 3, 5, 7, 9},
            new int[]{2, 4, 6, 8, 10}
        );
        
        Map<String, Map<Boolean, List<Integer>>> partitionedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array)
                    .boxed()
                    .collect(Collectors.partitioningBy(num -> num % 2 == 0))
            ));
        
        System.out.println("Partitioned arrays:");
        partitionedArrays.forEach((arrayStr, partition) -> 
            System.out.println(arrayStr + " -> Even: " + partition.get(true) + ", Odd: " + partition.get(false)));
        
        // Find arrays with only evens or only odds
        List<String> onlyEvenArrays = arrays.stream()
            .filter(array -> Arrays.stream(array).allMatch(num -> num % 2 == 0))
            .map(Arrays::toString)
            .collect(Collectors.toList());
        
        List<String> onlyOddArrays = arrays.stream()
            .filter(array -> Arrays.stream(array).allMatch(num -> num % 2 != 0))
            .map(Arrays::toString)
            .collect(Collectors.toList());
        
        System.out.println("Only even arrays: " + onlyEvenArrays);
        System.out.println("Only odd arrays: " + onlyOddArrays);
        
        // Calculate statistics for odds and evens
        IntSummaryStatistics evenStats = Arrays.stream(arr)
            .filter(num -> num % 2 == 0)
            .summaryStatistics();
        
        IntSummaryStatistics oddStats = Arrays.stream(arr)
            .filter(num -> num % 2 != 0)
            .summaryStatistics();
        
        System.out.println("Even statistics: " + evenStats);
        System.out.println("Odd statistics: " + oddStats);
        
        // Find maximum and minimum in each category
        OptionalInt maxEven = Arrays.stream(arr)
            .filter(num -> num % 2 == 0)
            .max();
        
        OptionalInt minEven = Arrays.stream(arr)
            .filter(num -> num % 2 == 0)
            .min();
        
        OptionalInt maxOdd = Arrays.stream(arr)
            .filter(num -> num % 2 != 0)
            .max();
        
        OptionalInt minOdd = Arrays.stream(arr)
            .filter(num -> num % 2 != 0)
            .min();
        
        System.out.println("Max even: " + maxEven.orElse(-1));
        System.out.println("Min even: " + minEven.orElse(-1));
        System.out.println("Max odd: " + maxOdd.orElse(-1));
        System.out.println("Min odd: " + minOdd.orElse(-1));
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.generate(() -> (int)(Math.random() * 1000))
            .limit(1000000)
            .toArray();
        
        long startTime = System.currentTimeMillis();
        Map<Boolean, Long> largePartitioned = Arrays.stream(largeArray).parallel()
            .boxed()
            .collect(Collectors.partitioningBy(num -> num % 2 == 0, Collectors.counting()));
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel partitioning time for 1M elements: " + parallelTime + "ms");
        System.out.println("Large array evens: " + largePartitioned.get(true));
        System.out.println("Large array odds: " + largePartitioned.get(false));
    }
}
```

## 34. Sum of Array Elements
**Java 8 Approach**: Using `Arrays.stream().sum()`

```java
import java.util.*;
import java.util.stream.*;

public class ArraySumJava8 {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        
        // Using Java 8 Streams
        int sum = Arrays.stream(arr).sum();
        System.out.println("Sum: " + sum);
        
        // Alternative: Using reduce
        int sumReduce = Arrays.stream(arr).reduce(0, Integer::sum);
        System.out.println("Sum (reduce): " + sumReduce);
        
        // Using IntStream directly
        int sumIntStream = IntStream.of(arr).sum();
        System.out.println("Sum (IntStream): " + sumIntStream);
        
        // Calculate sum with statistics
        IntSummaryStatistics stats = Arrays.stream(arr).summaryStatistics();
        System.out.println("Statistics: " + stats);
        
        // Process multiple arrays
        List<int[]> arrays = Arrays.asList(
            new int[]{1, 2, 3, 4, 5},
            new int[]{10, 20, 30},
            new int[]{100, 200, 300, 400},
            new int[]{-1, -2, -3, -4, -5}
        );
        
        Map<String, Integer> sumResults = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array).sum()
            ));
        
        System.out.println("Sum results:");
        sumResults.forEach((arrayStr, sumVal) -> 
            System.out.println(arrayStr + " -> " + sumVal));
        
        // Find array with maximum sum
        Optional<int[]> maxSumArray = arrays.stream()
            .max(Comparator.comparingInt(array -> Arrays.stream(array).sum()));
        
        maxSumArray.ifPresent(array -> 
            System.out.println("Array with max sum: " + Arrays.toString(array) + 
                             " (sum: " + Arrays.stream(array).sum() + ")"));
        
        // Find array with minimum sum
        Optional<int[]> minSumArray = arrays.stream()
            .min(Comparator.comparingInt(array -> Arrays.stream(array).sum()));
        
        minSumArray.ifPresent(array -> 
            System.out.println("Array with min sum: " + Arrays.toString(array) + 
                             " (sum: " + Arrays.stream(array).sum() + ")"));
        
        // Calculate sum of positive and negative numbers separately
        int positiveSum = Arrays.stream(arr).filter(num -> num > 0).sum();
        int negativeSum = Arrays.stream(arr).filter(num -> num < 0).sum();
        
        System.out.println("Positive sum: " + positiveSum);
        System.out.println("Negative sum: " + negativeSum);
        
        // Calculate cumulative sum
        int[] cumulativeSum = new int[arr.length];
        Arrays.stream(arr).reduce(0, (acc, num) -> {
            cumulativeSum[acc] = num;
            return acc + 1;
        });
        
        // Better cumulative sum calculation
        AtomicInteger index = new AtomicInteger(0);
        int[] cumulativeSumAlt = Arrays.stream(arr)
            .map(num -> {
                int idx = index.getAndIncrement();
                return IntStream.rangeClosed(0, idx).map(i -> arr[i]).sum();
            })
            .toArray();
        
        System.out.println("Cumulative sum: " + Arrays.toString(cumulativeSumAlt));
        
        // Calculate sum of squares
        int sumOfSquares = Arrays.stream(arr)
            .map(num -> num * num)
            .sum();
        
        System.out.println("Sum of squares: " + sumOfSquares);
        
        // Calculate sum of even and odd numbers separately
        int evenSum = Arrays.stream(arr).filter(num -> num % 2 == 0).sum();
        int oddSum = Arrays.stream(arr).filter(num -> num % 2 != 0).sum();
        
        System.out.println("Even sum: " + evenSum);
        System.out.println("Odd sum: " + oddSum);
        
        // Performance comparison for large arrays
        int[] largeArray = IntStream.range(1, 1000001).toArray();
        
        long startTime = System.currentTimeMillis();
        int largeSum = Arrays.stream(largeArray).parallel().sum();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel sum time for 1M elements: " + parallelTime + "ms");
        System.out.println("Large array sum: " + largeSum);
        
        // Verify with formula
        int expectedSum = 1000000 * 1000001 / 2;
        System.out.println("Expected sum (formula): " + expectedSum);
        System.out.println("Sum correct? " + (largeSum == expectedSum));
    }
}
```

## 35. Sort Array in Descending Order
**Java 8 Approach**: Using `Arrays.stream()` with `sorted(Comparator.reverseOrder())`

```java
import java.util.*;
import java.util.stream.*;

public class DescendingSortJava8 {
    public static void main(String[] args) {
        Integer[] arr = {5, 2, 9, 1, 6};
        
        // Using Java 8 Streams with reverse order
        Integer[] sorted = Arrays.stream(arr)
            .sorted(Comparator.reverseOrder())
            .toArray(Integer[]::new);
        
        System.out.println("Sorted descending: " + Arrays.toString(sorted));
        
        // Alternative: Using Collections.reverseOrder
        Integer[] sortedAlt = arr.clone();
        Arrays.sort(sortedAlt, Collections.reverseOrder());
        
        System.out.println("Sorted (Collections): " + Arrays.toString(sortedAlt));
        
        // For primitive arrays, need to box and unbox
        int[] primitiveArr = {5, 2, 9, 1, 6};
        int[] sortedPrimitive = Arrays.stream(primitiveArr)
            .boxed()
            .sorted(Comparator.reverseOrder())
            .mapToInt(Integer::intValue)
            .toArray();
        
        System.out.println("Primitive sorted: " + Arrays.toString(sortedPrimitive));
        
        // Sort multiple arrays
        List<Integer[]> arrays = Arrays.asList(
            new Integer[]{5, 2, 9, 1, 6},
            new Integer[]{10, 30, 20, 40},
            new Integer[]{100, 50, 75, 25, 125}
        );
        
        Map<String, Integer[]> sortedArrays = arrays.stream()
            .collect(Collectors.toMap(
                array -> Arrays.toString(array),
                array -> Arrays.stream(array)
                    .sorted(Comparator.reverseOrder())
                    .toArray(Integer[]::new)
            ));
        
        System.out.println("Sorted arrays:");
        sortedArrays.forEach((original, sortedArr) -> 
            System.out.println(original + " -> " + Arrays.toString(sortedArr)));
        
        // Sort and get top N elements
        int n = 3;
        Integer[] topN = Arrays.stream(arr)
            .sorted(Comparator.reverseOrder())
            .limit(n)
            .toArray(Integer[]::new);
        
        System.out.println("Top " + n + " elements: " + Arrays.toString(topN));
        
        // Sort and get bottom N elements
        Integer[] bottomN = Arrays.stream(arr)
            .sorted()
            .limit(n)
            .toArray(Integer[]::new);
        
        System.out.println("Bottom " + n + " elements: " + Arrays.toString(bottomN));
        
        // Sort with custom comparator (by absolute value descending)
        Integer[] arrWithNegatives = {5, -2, 9, -1, 6, -8};
        Integer[] sortedByAbsolute = Arrays.stream(arrWithNegatives)
            .sorted(Comparator.comparingInt(Math::abs).reversed())
            .toArray(Integer[]::new);
        
        System.out.println("Sorted by absolute value: " + Arrays.toString(sortedByAbsolute));
        
        // Sort strings by length descending
        String[] strings = {"Apple", "Banana", "Kiwi", "Strawberry", "Fig"};
        String[] sortedByLength = Arrays.stream(strings)
            .sorted(Comparator.comparingInt(String::length).reversed())
            .toArray(String[]::new);
        
        System.out.println("Strings sorted by length: " + Arrays.toString(sortedByLength));
        
        // Sort objects by multiple criteria
        class Person {
            String name;
            int age;
            Person(String name, int age) { this.name = name; this.age = age; }
            @Override public String toString() { return name + "(" + age + ")"; }
        }
        
        Person[] people = {
            new Person("Alice", 25),
            new Person("Bob", 30),
            new Person("Charlie", 25),
            new Person("David", 35)
        };
        
        Person[] sortedPeople = Arrays.stream(people)
            .sorted(Comparator.comparingInt((Person p) -> p.age).reversed()
                     .thenComparing(p -> p.name))
            .toArray(Person[]::new);
        
        System.out.println("People sorted by age desc, then name: " + Arrays.toString(sortedPeople));
        
        // Performance comparison for large arrays
        Integer[] largeArray = IntStream.generate(() -> (int)(Math.random() * 1000000))
            .limit(100000)
            .boxed()
            .toArray(Integer[]::new);
        
        long startTime = System.currentTimeMillis();
        Integer[] largeSorted = Arrays.stream(largeArray)
            .parallel()
            .sorted(Comparator.reverseOrder())
            .toArray(Integer[]::new);
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Parallel sort time for 100k elements: " + parallelTime + "ms");
        System.out.println("First 5 elements of sorted large array: " + 
            Arrays.toString(Arrays.copyOf(largeSorted, 5)));
        
        // Verify sorting correctness
        boolean isSortedCorrectly = IntStream.range(0, largeSorted.length - 1)
            .allMatch(i -> largeSorted[i] >= largeSorted[i + 1]);
        
        System.out.println("Large array sorted correctly? " + isSortedCorrectly);
    }
}
```

---

## 🎯 Key Java 8 Benefits for Array Processing

1. **Declarative Processing**: Express what operations to perform
2. **Parallel Processing**: Easy parallelization for large arrays
3. **Functional Operations**: Map, filter, reduce patterns
4. **Type Safety**: Generic operations with compile-time checking
5. **Built-in Statistics**: SummaryStatistics for quick analysis
6. **Collection Integration**: Seamless conversion between arrays and collections

## 📝 Best Practices

1. **Use primitive streams** (`IntStream`, `LongStream`) for better performance
2. **Prefer parallel streams** for large datasets only
3. **Use method references** when lambda expressions are simple
4. **Leverage collectors** for complex aggregations
5. **Use `Optional`** for safe operations that might not return values
6. **Consider memory usage** when working with large arrays

---

*This collection demonstrates how Java 8 features make array processing more elegant, readable, and efficient compared to traditional approaches.*
