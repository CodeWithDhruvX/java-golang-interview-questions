# 🟢 Java Data Structures: Level 1 (Beginner) Practice
Contains runnable code examples for Questions 1-42 (Arrays & Strings).

## Question 1: How do you declare, initialize, and copy an array in Java?

### Answer
Declare `[]`, Init `new int[5]`. Copy: `clone` or `System.arraycopy`.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class ArrayBasics {
    public static void main(String[] args) {
        // 1. Declaration & Initialization
        int[] arr = {1, 2, 3, 4, 5};
        
        // 2. Clone (Shallow Copy)
        int[] cloned = arr.clone();
        
        // 3. System.arraycopy
        int[] copied = new int[5];
        System.arraycopy(arr, 0, copied, 0, 5);
        
        System.out.println("Original: " + Arrays.toString(arr));
        System.out.println("Cloned: " + Arrays.toString(cloned));
    }
}
```

---

## Question 2: Difference between `Arrays.copyOf()` and `System.arraycopy()`.

### Answer
`copyOf` creates NEW array. `arraycopy` fills EXISTING array.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class CopyMethods {
    public static void main(String[] args) {
        int[] original = {10, 20, 30};
        
        // Arrays.copyOf (Returns new array, can resize)
        int[] newArr = Arrays.copyOf(original, 5); // Pads with 0
        System.out.println("CopyOf (Resized): " + Arrays.toString(newArr));
        
        // System.arraycopy (Needs existing destination)
        int[] dest = new int[3];
        System.arraycopy(original, 0, dest, 0, 3);
        System.out.println("ArrayCopy: " + Arrays.toString(dest));
    }
}
```

---

## Question 3: Difference between shallow copy and deep copy of arrays.

### Answer
Shallow copies refs. Deep copies objects.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

class Box { int val; Box(int v) { val = v; } @Override public String toString() { return ""+val; } }

public class ShallowVsDeep {
    public static void main(String[] args) {
        Box[] original = { new Box(1) };
        
        // Shallow Copy
        Box[] copy = Arrays.copyOf(original, 1);
        
        copy[0].val = 99; // Modifies object shared by both
        
        System.out.println("Original: " + original[0]); // 99 (Affected)
        System.out.println("Copy: " + copy[0]);       // 99
    }
}
```

---

## Question 4: How do you find the maximum or minimum in an array?

### Answer
Iterate and track max.

### Runnable Code
```java
package datastructures;

public class MinMax {
    public static void main(String[] args) {
        int[] arr = {5, 1, 9, 3, 7};
        
        int min = arr[0];
        int max = arr[0];
        
        for (int i : arr) {
            if (i < min) min = i;
            if (i > max) max = i;
        }
        
        System.out.println("Min: " + min);
        System.out.println("Max: " + max);
    }
}
```

---

## Question 5: How do you reverse an array in place?

### Answer
Two pointers (swap start/end).

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class ReverseArray {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        
        int i = 0, j = arr.length - 1;
        while (i < j) {
            int temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
            i++;
            j--;
        }
        
        System.out.println("Reversed: " + Arrays.toString(arr));
    }
}
```

---

## Question 6: How do you rotate an array by k positions?

### Answer
Reverse whole, Reverse 0-k, Reverse k-n.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class RotateArray {
    static void reverse(int[] arr, int start, int end) {
        while (start < end) {
            int tmp = arr[start];
            arr[start++] = arr[end];
            arr[end--] = tmp;
        }
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        int k = 2; // Rotate right by 2 -> {4, 5, 1, 2, 3}
        k %= arr.length;
        
        reverse(arr, 0, arr.length - 1); // 5 4 3 2 1
        reverse(arr, 0, k - 1);          // 4 5 3 2 1
        reverse(arr, k, arr.length - 1); // 4 5 1 2 3
        
        System.out.println("Rotated: " + Arrays.toString(arr));
    }
}
```

---

## Question 7: How do you remove duplicates from an array?

### Answer
Sort + 2 pointers OR Set.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class RemoveDuplicates {
    public static void main(String[] args) {
        int[] arr = {1, 1, 2, 2, 3};
        
        // Using Set (O(N) space)
        Set<Integer> set = new LinkedHashSet<>();
        for (int i : arr) set.add(i);
        
        System.out.println("Unique: " + set);
        
        // In-place (Sorted Array)
        // int[] arr = {1, 1, 2, 3};
        int uniqueIdx = 0;
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] != arr[uniqueIdx]) {
                uniqueIdx++;
                arr[uniqueIdx] = arr[i];
            }
        }
        // Unique elements are from 0 to uniqueIdx
    }
}
```

---

## Question 8: How do you find the second largest element in an array?

### Answer
Track largest and secondLargest.

### Runnable Code
```java
package datastructures;

public class SecondLargest {
    public static void main(String[] args) {
        int[] arr = {10, 5, 20, 8, 12};
        
        int first = Integer.MIN_VALUE;
        int second = Integer.MIN_VALUE;
        
        for (int n : arr) {
            if (n > first) {
                second = first;
                first = n;
            } else if (n > second && n != first) {
                second = n;
            }
        }
        
        System.out.println("Second Largest: " + second); // 12
    }
}
```

---

## Question 9: How do you find a missing number in a sequence of integers?

### Answer
Sum(1..N) - Sum(Array).

### Runnable Code
```java
package datastructures;

public class MissingNumber {
    public static void main(String[] args) {
        int[] arr = {1, 2, 4, 5}; // Missing 3
        int N = 5; // Expected top
        
        int expectedSum = N * (N + 1) / 2;
        int actualSum = 0;
        for (int i : arr) actualSum += i;
        
        System.out.println("Missing: " + (expectedSum - actualSum));
    }
}
```

---

## Question 10: How do you find duplicate elements in an array?

### Answer
HashSet or Sorting.

### Runnable Code
```java
package datastructures;

import java.util.HashSet;
import java.util.Set;

public class FindDuplicates {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 2, 4, 1};
        Set<Integer> seen = new HashSet<>();
        
        System.out.print("Duplicates: ");
        for (int i : arr) {
            if (!seen.add(i)) { // Returns false if exists
                System.out.print(i + " ");
            }
        }
    }
}
```

---

## Question 11: `Arrays.sort()` vs `Collections.sort()`?

### Answer
`Arrays`: Dual-Pivot Quicksort (Primitives). `Collections`: TimSort (Objects).

### Runnable Code
```java
package datastructures;

import java.util.*;

public class SortDiff {
    public static void main(String[] args) {
        int[] prim = {3, 1, 2};
        Arrays.sort(prim); // Quicksort
        
        List<Integer> list = new ArrayList<>(Arrays.asList(3, 1, 2));
        Collections.sort(list); // TimSort
    }
}
```

---

## Question 12: `Arrays.binarySearch()` — how does it work? What are the prerequisites?

### Answer
O(log N). Array MUST be sorted.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class BinarySearchDemo {
    public static void main(String[] args) {
        int[] arr = {1, 3, 5, 7, 9};
        
        int idx = Arrays.binarySearch(arr, 5);
        System.out.println("Found 5 at: " + idx);
        
        int idx2 = Arrays.binarySearch(arr, 4);
        // Returns -(insertion_point) - 1 => -2 - 1 = -3
        System.out.println("Not Found 4: " + idx2); 
    }
}
```

---

## Question 13: `Arrays.equals()` vs `==` for arrays?

### Answer
`==` Reference. `Arrays.equals` Content.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class ArrayEquals {
    public static void main(String[] args) {
        int[] a = {1, 2};
        int[] b = {1, 2};
        
        System.out.println(a == b);             // false
        System.out.println(Arrays.equals(a, b)); // true
    }
}
```

---

## Question 14: `Arrays.fill()` — practical use cases?

### Answer
Init DP arrays / Reset buffers.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class ArrayFill {
    public static void main(String[] args) {
        int[] dp = new int[5];
        Arrays.fill(dp, -1);
        System.out.println(Arrays.toString(dp));
    }
}
```

---

## Question 15: `Arrays.asList()` — caveats when using with arrays.

### Answer
Fixed size. Backed by array.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class ArraysAsListDemo {
    public static void main(String[] args) {
        String[] arr = {"A", "B"};
        List<String> list = Arrays.asList(arr);
        
        list.set(0, "C"); // Modifies arr too!
        System.out.println("Modified Array: " + arr[0]); // C
        
        // list.add("D"); // Throws UnsupportedOperationException
    }
}
```

---

## Question 16: `Arrays.stream()` — creating streams from arrays.

### Answer
Functional style.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class ArrayStream {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5};
        
        int sum = Arrays.stream(arr)
                        .filter(x -> x % 2 == 0)
                        .sum();
                        
        System.out.println("Sum Evens: " + sum);
    }
}
```

---

## Question 17: How to convert a primitive array to a list or stream?

### Answer
Use Streams to Box it.

### Runnable Code
```java
package datastructures;

import java.util.*;
import java.util.stream.Collectors;

public class PrimitiveToList {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3};
        
        // Arrays.asList(arr) gives List<int[]> NOT List<Integer>
        
        // Correct way:
        List<Integer> list = Arrays.stream(arr)
                                   .boxed()
                                   .collect(Collectors.toList());
        System.out.println(list);
    }
}
```

---

## Question 18: How do you declare, initialize, and traverse a 2D array?

### Answer
`int[][]`. Nested loops.

### Runnable Code
```java
package datastructures;

public class TwoDArray {
    public static void main(String[] args) {
        int[][] matrix = {
            {1, 2, 3},
            {4, 5, 6}
        };
        
        for (int i=0; i < matrix.length; i++) {
            for (int j=0; j < matrix[i].length; j++) {
                System.out.print(matrix[i][j] + " ");
            }
            System.out.println();
        }
    }
}
```

---

## Question 19: How do you find the sum of all elements in a 2D array?

### Answer
Accumulate in nested loop.

### Runnable Code
```java
package datastructures;

public class SumMatrix {
    public static void main(String[] args) {
        int[][] m = {{1, 2}, {3, 4}};
        int sum = 0;
        for (int[] row : m) {
            for (int val : row) {
                sum += val;
            }
        }
        System.out.println("Sum: " + sum);
    }
}
```

---

## Question 20: How do you rotate a 2D matrix (like image rotation)?

### Answer
Transpose + Reverse Rows (Clockwise).

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class RotateMatrix {
    public static void main(String[] args) {
        int[][] m = {
            {1, 2, 3},
            {4, 5, 6},
            {7, 8, 9}
        };
        
        // 1. Transpose
        for(int i=0; i<m.length; i++) {
            for(int j=i; j<m[0].length; j++) {
                int temp = m[i][j];
                m[i][j] = m[j][i];
                m[j][i] = temp;
            }
        }
        
        // 2. Reverse Rows
        for(int i=0; i<m.length; i++) {
            int left=0, right=m.length-1;
            while(left < right) {
                int temp = m[i][left];
                m[i][left] = m[i][right];
                m[i][right] = temp;
                left++; right--;
            }
        }
        
        System.out.println("Rotated 90:");
        for(int[] row : m) System.out.println(Arrays.toString(row));
    }
}
```

---

## Question 21: How do you check if a matrix is symmetric?

### Answer
`m[i][j] == m[j][i]`.

### Runnable Code
```java
package datastructures;

public class SymmetricMatrix {
    public static void main(String[] args) {
        int[][] m = {{1, 2}, {2, 1}};
        boolean sym = true;
        
        for (int i=0; i<m.length; i++) {
            for (int j=0; j<m.length; j++) {
                if (m[i][j] != m[j][i]) sym = false;
            }
        }
        System.out.println("Symmetric: " + sym);
    }
}
```

---

## Question 22: How to efficiently search in a row-wise & column-wise sorted matrix?

### Answer
Start top-right. Move left or down. O(N+M).

### Runnable Code
```java
package datastructures;

public class SearchMatrix {
    public static void main(String[] args) {
        int[][] m = {
            {10, 20, 30},
            {15, 25, 35},
            {27, 29, 37}
        };
        int target = 29;
        
        int row = 0, col = m[0].length - 1;
        boolean found = false;
        
        while (row < m.length && col >= 0) {
            if (m[row][col] == target) {
                found = true; break;
            } else if (m[row][col] > target) {
                col--;
            } else {
                row++;
            }
        }
        System.out.println("Found " + target + ": " + found);
    }
}
```

---

## Question 23: Difference between `String`, `StringBuilder`, and `StringBuffer`.

### Answer
`String` (Immutable), `StringBuilder` (Mutable, Fast), `StringBuffer` (ThreadSafe).

### Runnable Code
```java
package datastructures;

public class StringTypes {
    public static void main(String[] args) {
        String s = "Hello"; // Immutable
        
        StringBuilder sb = new StringBuilder("Hello");
        sb.append(" World"); // Modifies in place
        
        StringBuffer sbf = new StringBuffer("Hello"); // Synchronized
        sbf.append(" World");
        
        System.out.println(sb);
    }
}
```

---

## Question 24: `String` immutability — why is it important?

### Answer
Security, Thread-safety, Caching.

### Runnable Code
*(Conceptual - See Explanation)*

---

## Question 25: How do you reverse a string?

### Answer
`StringBuilder.reverse()`.

### Runnable Code
```java
package datastructures;

public class ReverseString {
    public static void main(String[] args) {
        String s = "Java";
        String r = new StringBuilder(s).reverse().toString();
        System.out.println(r);
    }
}
```

---

## Question 26: How do you check if a string is a palindrome?

### Answer
Reverse equals Original.

### Runnable Code
```java
package datastructures;

public class PalindromeCheck {
    public static void main(String[] args) {
        String s = "madam";
        boolean isPal = s.equals(new StringBuilder(s).reverse().toString());
        System.out.println("Palindrome: " + isPal);
    }
}
```

---

## Question 27: How do you remove duplicate characters from a string?

### Answer
Use `LinkedHashSet` or `distinct` stream.

### Runnable Code
```java
package datastructures;

import java.util.stream.Collectors;

public class RemoveDupChars {
    public static void main(String[] args) {
        String s = "banana";
        
        // Java 8
        String res = s.chars()
                      .distinct()
                      .mapToObj(c -> String.valueOf((char)c))
                      .collect(Collectors.joining());
                      
        System.out.println(res); // ban
    }
}
```

---

## Question 28: How do you count vowels/consonants in a string?

### Answer
Check `aeiou`.

### Runnable Code
```java
package datastructures;

public class CountVowels {
    public static void main(String[] args) {
        String s = "Hello World";
        int v = 0, c = 0;
        
        for (char ch : s.toLowerCase().toCharArray()) {
            if ("aeiou".indexOf(ch) != -1) v++;
            else if (Character.isLetter(ch)) c++;
        }
        System.out.println("Vowels: " + v + ", Cons: " + c);
    }
}
```

---

## Question 29: How do you check anagrams of two strings?

### Answer
Sort and compare.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class AnagramCheck {
    public static void main(String[] args) {
        String s1 = "listen";
        String s2 = "silent";
        
        char[] c1 = s1.toCharArray();
        char[] c2 = s2.toCharArray();
        
        Arrays.sort(c1);
        Arrays.sort(c2);
        
        System.out.println("Anagrams: " + Arrays.equals(c1, c2));
    }
}
```

---

## Question 30: How do you find the first non-repeating character?

### Answer
Frequency Map.

### Runnable Code
```java
package datastructures;

import java.util.*;

public class FirstUnique {
    public static void main(String[] args) {
        String s = "swiss";
        Map<Character, Integer> counts = new LinkedHashMap<>();
        
        for (char c : s.toCharArray()) counts.put(c, counts.getOrDefault(c, 0) + 1);
        
        for (Map.Entry<Character, Integer> entry : counts.entrySet()) {
            if (entry.getValue() == 1) {
                System.out.println("First Unique: " + entry.getKey()); // w
                return;
            }
        }
    }
}
```

---

## Question 31: How do you replace characters or substrings in Java?

### Answer
`replace`.

### Runnable Code
```java
package datastructures;

public class ReplaceDemo {
    public static void main(String[] args) {
        String s = "Hello Java";
        System.out.println(s.replace("Java", "World"));
    }
}
```

---

## Question 32: How do you split a string into an array of substrings?

### Answer
`split(regex)`.

### Runnable Code
```java
package datastructures;

import java.util.Arrays;

public class SplitDemo {
    public static void main(String[] args) {
        String s = "apple,banana,cherry";
        String[] parts = s.split(",");
        System.out.println(Arrays.toString(parts));
    }
}
```

---

## Question 33: Difference between `equals()` and `equalsIgnoreCase()`.

### Answer
Case sensitivity.

### Runnable Code
```java
package datastructures;

public class CaseEqual {
    public static void main(String[] args) {
        System.out.println("java".equalsIgnoreCase("JAVA")); // true
    }
}
```

---

## Question 34: Difference between `==` and `equals()` for strings.

### Answer
Reference vs Content.

### Runnable Code
```java
package datastructures;

public class StringEquality {
    public static void main(String[] args) {
        String s1 = "A";
        String s2 = new String("A");
        System.out.println(s1 == s2);      // false
        System.out.println(s1.equals(s2)); // true
    }
}
```

---

## Question 35: `substring()` vs `subSequence()`.

### Answer
Only return type differs.

### Runnable Code
```java
package datastructures;

public class SubStrSeq {
    public static void main(String[] args) {
        String s = "Hello";
        System.out.println(s.substring(1, 3));    // "el" (String)
        System.out.println(s.subSequence(1, 3));  // "el" (CharSequence)
    }
}
```

---

## Question 36: `charAt()`, `indexOf()`, `lastIndexOf()` — use cases.

### Answer
Search chars.

### Runnable Code
```java
package datastructures;

public class CharSearch {
    public static void main(String[] args) {
        String s = "hello";
        System.out.println(s.charAt(1)); // e
        System.out.println(s.indexOf('l')); // 2
        System.out.println(s.lastIndexOf('l')); // 3
    }
}
```

---

## Question 37: `startsWith()`, `endsWith()`, `contains()` — examples.

### Answer
Boolean checks.

### Runnable Code
```java
package datastructures;

public class StringChecks {
    public static void main(String[] args) {
        String f = "file.txt";
        System.out.println(f.endsWith(".txt")); // true
    }
}
```

---

## Question 38: `trim()`, `strip()`, `stripLeading()`, `stripTrailing()` differences.

### Answer
`strip` handles Unicode.

### Runnable Code
```java
package datastructures;

public class StripTrim {
    public static void main(String[] args) {
        String s = "  hi  ";
        System.out.println("'" + s.trim() + "'"); // 'hi'
        System.out.println("'" + s.strip() + "'"); // 'hi' (Java 11+)
    }
}
```

---

## Question 39: `replace()` vs `replaceAll()` vs `replaceFirst()`.

### Answer
`replaceAll` uses Regex.

### Runnable Code
```java
package datastructures;

public class ReplaceTypes {
    public static void main(String[] args) {
        String s = "foo 123 bar";
        System.out.println(s.replaceAll("\\d+", "#")); // foo # bar
    }
}
```

---

## Question 40: `matches()` and regular expressions for validation.

### Answer
Partial vs Whole match.

### Runnable Code
```java
package datastructures;

public class RegexMatch {
    public static void main(String[] args) {
        System.out.println("12345".matches("\\d+")); // true
    }
}
```

---

## Question 41: Converting string to char array and vice versa.

### Answer
`toCharArray`, `new String`.

### Runnable Code
```java
package datastructures;

public class StrArrConv {
    public static void main(String[] args) {
        char[] arr = "abc".toCharArray();
        String s = new String(arr);
    }
}
```

---

## Question 42: Converting string to uppercase/lowercase and locale considerations.

### Answer
Locale sensitive.

### Runnable Code
```java
package datastructures;

import java.util.Locale;

public class CaseConv {
    public static void main(String[] args) {
        String s = "i";
        System.out.println(s.toUpperCase(Locale.ENGLISH)); // I
        System.out.println(s.toUpperCase(new Locale("tr"))); // İ (Turkish)
    }
}
```
