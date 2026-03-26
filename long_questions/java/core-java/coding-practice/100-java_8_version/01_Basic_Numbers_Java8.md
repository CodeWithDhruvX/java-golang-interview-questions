# Basic Number Programs with Java 8 Features (1-15)

## 📚 Java 8 Features Demonstrated
- **Lambda Expressions**: Concise anonymous functions
- **Streams API**: Functional data processing pipelines
- **Method References**: Simplified lambda expressions
- **Optional**: Null-safe operations
- **Collectors**: Stream terminal operations
- **Functional Interfaces**: Predicate, Function, Consumer, Supplier
- **Parallel Streams**: Multi-threaded processing

---

## 1. Fibonacci Series
**Java 8 Approach**: Using `Stream.generate()` and `Stream.limit()`

```java
import java.util.*;
import java.util.stream.*;

public class FibonacciJava8 {
    public static void main(String[] args) {
        int n = 10;
        
        // Using Java 8 Stream generate
        System.out.println("Fibonacci Series up to " + n + " terms:");
        Stream.generate(new FibonacciSupplier())
            .limit(n)
            .forEach(num -> System.out.print(num + " "));
        
        System.out.println();
        
        // Alternative: Using IntStream with reduce
        List<Integer> fibonacci = IntStream.range(0, n)
            .collect(ArrayList::new, 
                (list, i) -> {
                    if (list.size() < 2) {
                        list.add(i);
                    } else {
                        list.add(list.get(list.size() - 1) + list.get(list.size() - 2));
                    }
                },
                ArrayList::addAll);
        
        fibonacci.forEach(num -> System.out.print(num + " "));
    }
    
    static class FibonacciSupplier implements Supplier<Integer> {
        private int a = 0, b = 1;
        
        @Override
        public Integer get() {
            int result = a;
            int next = a + b;
            a = b;
            b = next;
            return result;
        }
    }
}
```

## 2. Check Prime Number
**Java 8 Approach**: Using `IntStream.range()` and `noneMatch()`

```java
import java.util.*;
import java.util.stream.*;

public class PrimeCheckJava8 {
    public static void main(String[] args) {
        int num = 29;
        
        // Using Java 8 Streams
        boolean isPrime = num > 1 && IntStream.range(2, (int) Math.sqrt(num) + 1)
            .noneMatch(i -> num % i == 0);
        
        System.out.println(num + " is prime? " + isPrime);
        
        // Alternative: Using parallel stream for large numbers
        boolean isPrimeParallel = num > 1 && IntStream.range(2, (int) Math.sqrt(num) + 1)
            .parallel()
            .noneMatch(i -> num % i == 0);
        
        // Check multiple numbers
        List<Integer> numbers = Arrays.asList(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 17, 19, 23, 29);
        Map<Integer, Boolean> primeResults = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> n > 1 && IntStream.range(2, (int) Math.sqrt(n) + 1)
                    .noneMatch(i -> n % i == 0)
            ));
        
        System.out.println("Prime results: " + primeResults);
    }
}
```

## 3. Factorial of a Number
**Java 8 Approach**: Using `LongStream.rangeClosed()` and `reduce()`

```java
import java.util.*;
import java.util.stream.*;

public class FactorialJava8 {
    public static void main(String[] args) {
        int num = 5;
        
        // Using Java 8 Streams
        long factorial = LongStream.rangeClosed(1, num)
            .reduce(1, (a, b) -> a * b);
        
        System.out.println("Factorial of " + num + " = " + factorial);
        
        // Alternative: Using parallel stream for large numbers
        long factorialParallel = LongStream.rangeClosed(1, num)
            .parallel()
            .reduce(1, (a, b) -> a * b);
        
        // Calculate multiple factorials
        List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8);
        Map<Integer, Long> factorials = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> LongStream.rangeClosed(1, n)
                    .reduce(1, (a, b) -> a * b)
            ));
        
        System.out.println("Factorials: " + factorials);
        
        // Using Optional for null safety
        Optional<Integer> optionalNum = Optional.of(num);
        optionalNum.ifPresent(n -> System.out.println("Factorial of " + n + " = " + 
            LongStream.rangeClosed(1, n).reduce(1, (a, b) -> a * b)));
    }
}
```

## 4. Palindrome Number
**Java 8 Approach**: Using `String.valueOf()` and stream reversal

```java
import java.util.*;
import java.util.stream.*;

public class PalindromeNumberJava8 {
    public static void main(String[] args) {
        int num = 121;
        
        // Using Java 8 Streams with String
        String numStr = String.valueOf(num);
        boolean isPalindrome = IntStream.range(0, numStr.length() / 2)
            .allMatch(i -> numStr.charAt(i) == numStr.charAt(numStr.length() - 1 - i));
        
        System.out.println(num + (isPalindrome ? " is" : " is not") + " a palindrome.");
        
        // Alternative: Using StringBuilder with stream
        String reversed = new StringBuilder(numStr).reverse().toString();
        boolean isPalindromeAlt = numStr.equals(reversed);
        
        // Check multiple numbers
        List<Integer> numbers = Arrays.asList(121, 123, 131, 456, 787, 999);
        Map<Integer, Boolean> palindromeResults = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> {
                    String s = String.valueOf(n);
                    return s.equals(new StringBuilder(s).reverse().toString());
                }
            ));
        
        System.out.println("Palindrome results: " + palindromeResults);
        
        // Using Optional for safe processing
        Optional<Integer> optionalNum = Optional.of(num);
        optionalNum.ifPresent(n -> {
            String s = String.valueOf(n);
            boolean result = s.equals(new StringBuilder(s).reverse().toString());
            System.out.println(n + (result ? " is" : " is not") + " a palindrome.");
        });
    }
}
```

## 5. Armstrong Number
**Java 8 Approach**: Using `String.chars()` and `mapToDouble()`

```java
import java.util.*;
import java.util.stream.*;

public class ArmstrongJava8 {
    public static void main(String[] args) {
        int num = 153;
        
        // Using Java 8 Streams
        String numStr = String.valueOf(num);
        int n = numStr.length();
        
        int result = numStr.chars()
            .map(Character::getNumericValue)
            .map(digit -> (int) Math.pow(digit, n))
            .sum();
        
        System.out.println(num + (result == num ? " is" : " is not") + " an Armstrong number.");
        
        // Alternative: Using parallel stream
        int resultParallel = numStr.chars()
            .parallel()
            .map(Character::getNumericValue)
            .map(digit -> (int) Math.pow(digit, n))
            .sum();
        
        // Check multiple numbers
        List<Integer> numbers = Arrays.asList(153, 370, 371, 407, 1634, 8208, 9474, 100, 200);
        Map<Integer, Boolean> armstrongResults = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> {
                    String s = String.valueOf(n);
                    int power = s.length();
                    return s.chars()
                        .map(Character::getNumericValue)
                        .map(digit -> (int) Math.pow(digit, power))
                        .sum() == n;
                }
            ));
        
        System.out.println("Armstrong results: " + armstrongResults);
        
        // Find all Armstrong numbers in a range
        List<Integer> armstrongNumbers = IntStream.range(1, 10000)
            .filter(n -> {
                String s = String.valueOf(n);
                int power = s.length();
                return s.chars()
                    .map(Character::getNumericValue)
                    .map(digit -> (int) Math.pow(digit, power))
                    .sum() == n;
            })
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("Armstrong numbers (1-9999): " + armstrongNumbers);
    }
}
```

## 6. Reverse a Number
**Java 8 Approach**: Using `StringBuilder` with streams

```java
import java.util.*;
import java.util.stream.*;

public class ReverseNumberJava8 {
    public static void main(String[] args) {
        int num = 1234;
        
        // Using Java 8 Streams with StringBuilder
        String reversed = new StringBuilder(String.valueOf(num))
            .reverse()
            .toString();
        
        System.out.println("Reversed: " + reversed);
        
        // Alternative: Using character stream
        String reversedStream = String.valueOf(num)
            .chars()
            .mapToObj(c -> String.valueOf((char) c))
            .collect(Collectors.collectingAndThen(
                Collectors.toList(),
                list -> {
                    Collections.reverse(list);
                    return String.join("", list);
                }
            ));
        
        System.out.println("Reversed (stream): " + reversedStream);
        
        // Reverse multiple numbers
        List<Integer> numbers = Arrays.asList(1234, 5678, 9012, 3456);
        Map<Integer, String> reversedNumbers = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> new StringBuilder(String.valueOf(n)).reverse().toString()
            ));
        
        System.out.println("Reversed numbers: " + reversedNumbers);
        
        // Using Optional for safe processing
        Optional<Integer> optionalNum = Optional.of(num);
        optionalNum.map(String::valueOf)
                  .map(StringBuilder::new)
                  .map(StringBuilder::reverse)
                  .map(StringBuilder::toString)
                  .ifPresent(reversedNum -> System.out.println("Reversed: " + reversedNum));
    }
}
```

## 7. Swap Two Numbers without Third Variable
**Java 8 Approach**: Using functional interfaces

```java
import java.util.*;
import java.util.function.*;

public class SwapNoTempJava8 {
    public static void main(String[] args) {
        int a = 10, b = 20;
        
        // Using Java 8 functional interfaces
        BiConsumer<int[], int[]> swapArithmetic = (arr1, arr2) -> {
            arr1[0] = arr1[0] + arr2[0];
            arr2[0] = arr1[0] - arr2[0];
            arr1[0] = arr1[0] - arr2[0];
        };
        
        BiConsumer<int[], int[]> swapXor = (arr1, arr2) -> {
            arr1[0] = arr1[0] ^ arr2[0];
            arr2[0] = arr1[0] ^ arr2[0];
            arr1[0] = arr1[0] ^ arr2[0];
        };
        
        // Demonstrate swapping
        int[] arr1 = {a};
        int[] arr2 = {b};
        
        System.out.println("Before: a=" + arr1[0] + ", b=" + arr2[0]);
        swapArithmetic.accept(arr1, arr2);
        System.out.println("After arithmetic swap: a=" + arr1[0] + ", b=" + arr2[0]);
        
        // Reset and try XOR
        arr1[0] = 10; arr2[0] = 20;
        swapXor.accept(arr1, arr2);
        System.out.println("After XOR swap: a=" + arr1[0] + ", b=" + arr2[0]);
        
        // Using streams to demonstrate multiple swaps
        List<int[]> pairs = Arrays.asList(
            new int[]{5, 10},
            new int[]{15, 25},
            new int[]{30, 40}
        );
        
        System.out.println("Before swaps:");
        pairs.forEach(pair -> System.out.println(Arrays.toString(pair)));
        
        // Apply swap to all pairs
        pairs.forEach(pair -> swapArithmetic.accept(
            new int[]{pair[0]}, 
            new int[]{pair[1]}
        ));
        
        System.out.println("After swaps:");
        pairs.forEach(pair -> System.out.println(Arrays.toString(pair)));
    }
}
```

## 8. Check Even or Odd
**Java 8 Approach**: Using `Predicate` and streams

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class EvenOddJava8 {
    public static void main(String[] args) {
        int num = 11;
        
        // Using Java 8 Predicate
        Predicate<Integer> isEven = n -> n % 2 == 0;
        Predicate<Integer> isOdd = n -> n % 2 != 0;
        
        System.out.println(num + " is " + (isEven.test(num) ? "Even" : "Odd"));
        
        // Alternative: Using ternary with lambda
        String result = isEven.test(num) ? "Even" : "Odd";
        System.out.println(num + " is " + result);
        
        // Process multiple numbers
        List<Integer> numbers = Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        
        Map<String, List<Integer>> grouped = numbers.stream()
            .collect(Collectors.groupingBy(
                n -> isEven.test(n) ? "Even" : "Odd"
            ));
        
        System.out.println("Grouped numbers: " + grouped);
        
        // Using parallel stream for large lists
        List<Integer> largeNumbers = IntStream.rangeClosed(1, 1000)
            .boxed()
            .collect(Collectors.toList());
        
        Map<Boolean, List<Integer>> partitioned = largeNumbers.parallelStream()
            .collect(Collectors.partitioningBy(isEven));
        
        System.out.println("Even count: " + partitioned.get(true).size());
        System.out.println("Odd count: " + partitioned.get(false).size());
        
        // Using Optional for safe processing
        Optional<Integer> optionalNum = Optional.of(num);
        optionalNum.map(n -> n + " is " + (isEven.test(n) ? "Even" : "Odd"))
                  .ifPresent(System.out::println);
    }
}
```

## 9. Check Leap Year
**Java 8 Approach**: Using complex `Predicate` composition

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class LeapYearJava8 {
    public static void main(String[] args) {
        int year = 2024;
        
        // Using Java 8 Predicate composition
        Predicate<Integer> divisibleBy4 = n -> n % 4 == 0;
        Predicate<Integer> divisibleBy100 = n -> n % 100 == 0;
        Predicate<Integer> divisibleBy400 = n -> n % 400 == 0;
        
        Predicate<Integer> isLeap = divisibleBy4
            .and(divisibleBy100.negate().or(divisibleBy400));
        
        System.out.println(year + (isLeap.test(year) ? " is" : " is not") + " a leap year.");
        
        // Alternative: Using IntPredicate
        IntPredicate leapYearPredicate = n -> (n % 4 == 0) && 
            ((n % 100 != 0) || (n % 400 == 0));
        
        // Check multiple years
        List<Integer> years = Arrays.asList(2000, 2001, 2004, 1900, 2008, 2012, 2016, 2020, 2024);
        Map<Integer, Boolean> leapResults = years.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                isLeap::test
            ));
        
        System.out.println("Leap year results: " + leapResults);
        
        // Find all leap years in a range
        List<Integer> leapYears = IntStream.range(2000, 2025)
            .filter(leapYearPredicate)
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("Leap years (2000-2024): " + leapYears);
        
        // Using Optional for safe processing
        Optional<Integer> optionalYear = Optional.of(year);
        optionalYear.filter(isLeap)
                   .ifPresent(y -> System.out.println(y + " is a leap year."));
        
        optionalYear.filter(isLeap.negate())
                   .ifPresent(y -> System.out.println(y + " is not a leap year."));
    }
}
```

## 10. Greatest of Three Numbers
**Java 8 Approach**: Using `Stream.max()` and `Comparator`

```java
import java.util.*;
import java.util.stream.*;

public class LargestOfThreeJava8 {
    public static void main(String[] args) {
        int a = 10, b = 20, c = 15;
        
        // Using Java 8 Stream
        int max = Stream.of(a, b, c)
            .max(Integer::compare)
            .orElse(Integer.MIN_VALUE);
        
        System.out.println("Largest: " + max);
        
        // Alternative: Using IntStream
        int maxInt = IntStream.of(a, b, c)
            .max()
            .orElse(Integer.MIN_VALUE);
        
        System.out.println("Largest (IntStream): " + maxInt);
        
        // Find all three numbers in sorted order
        List<Integer> sorted = Stream.of(a, b, c)
            .sorted(Comparator.reverseOrder())
            .collect(Collectors.toList());
        
        System.out.println("Sorted: " + sorted);
        
        // Process multiple triples
        List<int[]> triples = Arrays.asList(
            new int[]{10, 20, 15},
            new int[]{5, 8, 3},
            new int[]{100, 50, 75},
            new int[]{1, 2, 3}
        );
        
        Map<int[], Integer> maxResults = triples.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                triple -> IntStream.of(triple).max().orElse(Integer.MIN_VALUE)
            ));
        
        System.out.println("Maximum values:");
        maxResults.forEach((triple, maxVal) -> 
            System.out.println(Arrays.toString(triple) + " -> " + maxVal));
        
        // Using Optional for safe processing
        OptionalInt maxOptional = IntStream.of(a, b, c).max();
        maxOptional.ifPresent(maxVal -> 
            System.out.println("Largest: " + maxVal));
    }
}
```

## 11. Sum of Digits
**Java 8 Approach**: Using `String.chars()` and `map()`

```java
import java.util.*;
import java.util.stream.*;

public class SumDigitsJava8 {
    public static void main(String[] args) {
        int num = 1234;
        
        // Using Java 8 Streams
        int sum = String.valueOf(num)
            .chars()
            .map(Character::getNumericValue)
            .sum();
        
        System.out.println("Sum: " + sum);
        
        // Alternative: Using IntStream
        int sumAlt = IntStream.of(String.valueOf(num).chars())
            .map(Character::getNumericValue)
            .sum();
        
        // Process multiple numbers
        List<Integer> numbers = Arrays.asList(1234, 5678, 9012, 3456, 7890);
        Map<Integer, Integer> sumResults = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> String.valueOf(n)
                    .chars()
                    .map(Character::getNumericValue)
                    .sum()
            ));
        
        System.out.println("Sum of digits: " + sumResults);
        
        // Find numbers with digit sum equal to target
        int targetSum = 10;
        List<Integer> numbersWithTargetSum = IntStream.range(1, 1000)
            .filter(n -> String.valueOf(n)
                .chars()
                .map(Character::getNumericValue)
                .sum() == targetSum)
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("Numbers with digit sum " + targetSum + ": " + 
            numbersWithTargetSum.subList(0, Math.min(10, numbersWithTargetSum.size())));
        
        // Using parallel stream for large numbers
        List<Integer> largeNumbers = IntStream.rangeClosed(1, 100000)
            .boxed()
            .collect(Collectors.toList());
        
        Map<Integer, Integer> digitSums = largeNumbers.parallelStream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> String.valueOf(n)
                    .chars()
                    .map(Character::getNumericValue)
                    .sum()
            ));
        
        // Find number with maximum digit sum
        int maxDigitSumNumber = digitSums.entrySet().stream()
            .max(Map.Entry.comparingByValue())
            .map(Map.Entry::getKey)
            .orElse(0);
        
        System.out.println("Number with max digit sum (1-100000): " + maxDigitSumNumber);
    }
}
```

## 12. GCD of Two Numbers
**Java 8 Approach**: Using `IntStream` with Euclidean algorithm

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class GCDJava8 {
    public static void main(String[] args) {
        int n1 = 81, n2 = 153;
        
        // Using Java 8 Stream with Euclidean algorithm
        IntBinaryOperator gcd = (a, b) -> {
            while (b != 0) {
                int temp = b;
                b = a % b;
                a = temp;
            }
            return a;
        };
        
        int result = gcd.applyAsInt(n1, n2);
        System.out.println("GCD: " + result);
        
        // Alternative: Using IntStream to find common divisors
        int gcdAlt = IntStream.rangeClosed(1, Math.min(n1, n2))
            .filter(i -> n1 % i == 0 && n2 % i == 0)
            .max()
            .orElse(1);
        
        System.out.println("GCD (alternative): " + gcdAlt);
        
        // Process multiple pairs
        List<int[]> pairs = Arrays.asList(
            new int[]{81, 153},
            new int[]{12, 18},
            new int[]{100, 25},
            new int[]{17, 23}
        );
        
        Map<String, Integer> gcdResults = pairs.stream()
            .collect(Collectors.toMap(
                pair -> Arrays.toString(pair),
                pair -> gcd.applyAsInt(pair[0], pair[1])
            ));
        
        System.out.println("GCD results: " + gcdResults);
        
        // Using parallel stream for large numbers
        List<int[]> largePairs = IntStream.range(1, 100)
            .boxed()
            .flatMap(i -> IntStream.range(i + 1, 101)
                .mapToObj(j -> new int[]{i, j}))
            .collect(Collectors.toList());
        
        Map<String, Integer> largeGcdResults = largePairs.parallelStream()
            .collect(Collectors.toMap(
                pair -> Arrays.toString(pair),
                pair -> gcd.applyAsInt(pair[0], pair[1])
            ));
        
        // Find pairs with GCD = 1 (coprime numbers)
        List<String> coprimePairs = largeGcdResults.entrySet().stream()
            .filter(entry -> entry.getValue() == 1)
            .map(Map.Entry::getKey)
            .limit(10)
            .collect(Collectors.toList());
        
        System.out.println("Coprime pairs (first 10): " + coprimePairs);
    }
}
```

## 13. LCM of Two Numbers
**Java 8 Approach**: Using GCD function and streams

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class LCMJava8 {
    public static void main(String[] args) {
        int n1 = 72, n2 = 120;
        
        // Using Java 8 with GCD function
        IntBinaryOperator gcd = (a, b) -> {
            while (b != 0) {
                int temp = b;
                b = a % b;
                a = temp;
            }
            return a;
        };
        
        IntBinaryOperator lcm = (a, b) -> Math.abs(a * b) / gcd.applyAsInt(a, b);
        
        int result = lcm.applyAsInt(n1, n2);
        System.out.println("LCM: " + result);
        
        // Alternative: Using IntStream to find common multiples
        int lcmAlt = IntStream.rangeClosed(Math.max(n1, n2), n1 * n2)
            .filter(i -> i % n1 == 0 && i % n2 == 0)
            .findFirst()
            .orElse(n1 * n2);
        
        System.out.println("LCM (alternative): " + lcmAlt);
        
        // Process multiple pairs
        List<int[]> pairs = Arrays.asList(
            new int[]{72, 120},
            new int[]{12, 18},
            new int[]{5, 7},
            new int[]{15, 20}
        );
        
        Map<String, Integer> lcmResults = pairs.stream()
            .collect(Collectors.toMap(
                pair -> Arrays.toString(pair),
                pair -> lcm.applyAsInt(pair[0], pair[1])
            ));
        
        System.out.println("LCM results: " + lcmResults);
        
        // Find LCM of multiple numbers
        List<Integer> numbers = Arrays.asList(2, 3, 4, 5);
        int lcmMultiple = numbers.stream()
            .reduce(1, (acc, num) -> lcm.applyAsInt(acc, num));
        
        System.out.println("LCM of " + numbers + " = " + lcmMultiple);
        
        // Using parallel stream for large computations
        List<int[]> largePairs = IntStream.range(1, 50)
            .boxed()
            .flatMap(i -> IntStream.range(i + 1, 51)
                .mapToObj(j -> new int[]{i, j}))
            .collect(Collectors.toList());
        
        Map<String, Integer> largeLcmResults = largePairs.parallelStream()
            .collect(Collectors.toMap(
                pair -> Arrays.toString(pair),
                pair -> lcm.applyAsInt(pair[0], pair[1])
            ));
        
        // Find pairs with smallest LCM
        Optional<Map.Entry<String, Integer>> minLcm = largeLcmResults.entrySet().stream()
            .min(Map.Entry.comparingByValue());
        
        minLcm.ifPresent(entry -> 
            System.out.println("Smallest LCM: " + entry.getKey() + " -> " + entry.getValue()));
    }
}
```

## 14. Perfect Number
**Java 8 Approach**: Using `IntStream.range()` and `filter()`

```java
import java.util.*;
import java.util.stream.*;

public class PerfectNumberJava8 {
    public static void main(String[] args) {
        int num = 28;
        
        // Using Java 8 Streams
        int sum = IntStream.range(1, num)
            .filter(i -> num % i == 0)
            .sum();
        
        boolean isPerfect = sum == num;
        System.out.println(num + (isPerfect ? " is" : " is not") + " a Perfect Number.");
        
        // Alternative: Using parallel stream for large numbers
        int sumParallel = IntStream.range(1, num)
            .parallel()
            .filter(i -> num % i == 0)
            .sum();
        
        // Check multiple numbers
        List<Integer> numbers = Arrays.asList(6, 28, 496, 8128, 33550336, 12, 24, 30);
        Map<Integer, Boolean> perfectResults = numbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> IntStream.range(1, n)
                    .filter(i -> n % i == 0)
                    .sum() == n
            ));
        
        System.out.println("Perfect number results: " + perfectResults);
        
        // Find all perfect numbers in a range
        List<Integer> perfectNumbers = IntStream.range(2, 10000)
            .filter(n -> IntStream.range(1, n)
                .filter(i -> n % i == 0)
                .sum() == n)
            .boxed()
            .collect(Collectors.toList());
        
        System.out.println("Perfect numbers (2-9999): " + perfectNumbers);
        
        // Using Optional for safe processing
        Optional<Integer> optionalNum = Optional.of(num);
        optionalNum.filter(n -> IntStream.range(1, n)
                .filter(i -> n % i == 0)
                .sum() == n)
                   .ifPresent(n -> System.out.println(n + " is a Perfect Number."));
        
        // Find divisors of perfect numbers
        Map<Integer, List<Integer>> perfectDivisors = perfectNumbers.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                n -> IntStream.range(1, n)
                    .filter(i -> n % i == 0)
                    .boxed()
                    .collect(Collectors.toList())
            ));
        
        System.out.println("Perfect number divisors: " + perfectDivisors);
    }
}
```

## 15. Print ASCII Value
**Java 8 Approach**: Using `IntStream.range()` and method references

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class ASCIIJava8 {
    public static void main(String[] args) {
        char ch = 'a';
        
        // Using Java 8 method reference
        System.out.println("ASCII of " + ch + " is " + (int)ch);
        
        // Alternative: Using Character class
        int asciiValue = Character.getNumericValue(ch);
        System.out.println("Numeric value of " + ch + " is " + asciiValue);
        
        // Print ASCII values for multiple characters
        String text = "Hello World";
        Map<Character, Integer> asciiValues = text.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.toMap(
                Function.identity(),
                c -> (int) c
            ));
        
        System.out.println("ASCII values: " + asciiValues);
        
        // Print ASCII table for lowercase letters
        System.out.println("ASCII table for lowercase letters:");
        IntStream.rangeClosed('a', 'z')
            .mapToObj(c -> (char) c + " = " + c)
            .forEach(System.out::println);
        
        // Print ASCII table for uppercase letters
        System.out.println("\nASCII table for uppercase letters:");
        IntStream.rangeClosed('A', 'Z')
            .mapToObj(c -> (char) c + " = " + c)
            .forEach(System.out::println);
        
        // Find characters in specific ASCII range
        IntPredicate printableRange = c -> c >= 32 && c <= 126;
        List<Character> printableChars = IntStream.rangeClosed(32, 126)
            .filter(printableRange)
            .mapToObj(c -> (char) c)
            .collect(Collectors.toList());
        
        System.out.println("\nPrintable ASCII characters: " + printableChars);
        
        // Using Optional for safe processing
        Optional<Character> optionalChar = Optional.of(ch);
        optionalChar.map(c -> "ASCII of " + c + " is " + (int)c)
                   .ifPresent(System.out::println);
        
        // Character classification using streams
        String sample = "Hello123World!@#";
        Map<String, List<Character>> classified = sample.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(c -> {
                if (Character.isLetter(c)) return "Letter";
                else if (Character.isDigit(c)) return "Digit";
                else return "Special";
            }));
        
        System.out.println("\nCharacter classification: " + classified);
        
        // Count character types
        Map<String, Long> charCounts = sample.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                c -> {
                    if (Character.isLetter(c)) return "Letter";
                    else if (Character.isDigit(c)) return "Digit";
                    else return "Special";
                },
                Collectors.counting()
            ));
        
        System.out.println("Character counts: " + charCounts);
    }
}
```

---

## 🎯 Key Java 8 Benefits for Number Processing

1. **Declarative Style**: Express what to do, not how to do it
2. **Parallel Processing**: Easy parallelization for large datasets
3. **Functional Composition**: Chain operations elegantly
4. **Type Safety**: Generic operations with compile-time checking
5. **Null Safety**: Optional for avoiding NullPointerException
6. **Readability**: More expressive and maintainable code

## 📝 Best Practices

1. **Use method references** when lambda expressions just call a method
2. **Prefer parallel streams** for large datasets only
3. **Use Optional** for null-safe operations
4. **Leverage collectors** for complex aggregations
5. **Chain operations** for readable data pipelines
6. **Use primitive streams** (IntStream, LongStream) for better performance

---

*This collection demonstrates how Java 8 features make number processing more elegant, readable, and efficient compared to traditional approaches.*
