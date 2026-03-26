# String Programs with Java 8 Features (Lambda, Streams & Collections)

## 📚 Java 8 Features Demonstrated
- **Lambda Expressions**: Concise anonymous functions
- **Streams API**: Functional data processing pipelines
- **Method References**: Simplified lambda expressions
- **Optional**: Null-safe operations
- **Collectors**: Stream terminal operations
- **Functional Interfaces**: Predicate, Function, Consumer
- **Parallel Streams**: Multi-threaded processing

---

## 🚀 Beginner Level (1-10) - Java 8 Style

### 1. Count Occurrence of Each Character
**Java 8 Approach**: Using `chars()` stream and `Collectors.groupingBy()`

```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class CharFrequencyJava8 {
    public static void main(String[] args) {
        String str = "programming";
        
        // Using Java 8 Streams
        Map<Character, Long> frequency = str.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                LinkedHashMap::new,
                Collectors.counting()
            ));
        
        // Print using lambda
        frequency.forEach((ch, count) -> 
            System.out.println(ch + "=" + count));
        
        // Alternative: Using parallel stream for large strings
        Map<Character, Long> parallelFreq = str.parallelChars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
    }
}
```

### 2. Find Length Without length()
**Java 8 Approach**: Using `chars().count()`

```java
public class StringLengthJava8 {
    public static void main(String[] args) {
        String str = "Hello";
        
        // Using Java 8 Stream count
        long length = str.chars().count();
        System.out.println("Length: " + length);
        
        // Alternative: Using reduce
        int lengthReduce = str.chars()
            .reduce(0, (acc, c) -> acc + 1);
        System.out.println("Length (reduce): " + lengthReduce);
        
        // Using Optional for null safety
        Optional.ofNullable(str)
            .ifPresent(s -> System.out.println("Length: " + s.chars().count()));
    }
}
```

### 3. First Repeated Character
**Java 8 Approach**: Using `Stream.filter()` and `findFirst()`

```java
import java.util.*;
import java.util.stream.*;

public class FirstRepeatedJava8 {
    public static void main(String[] args) {
        String str = "swiss";
        
        // Using Java 8 Streams
        Optional<Character> firstRepeated = str.chars()
            .mapToObj(c -> (char) c)
            .filter(ch -> str.indexOf(ch) != str.lastIndexOf(ch))
            .findFirst();
        
        firstRepeated.ifPresentOrElse(
            ch -> System.out.println("First repeated: " + ch),
            () -> System.out.println("No repeated character")
        );
        
        // Alternative: Using Set with streams
        Set<Character> seen = new HashSet<>();
        Optional<Character> firstRepeatedSet = str.chars()
            .mapToObj(c -> (char) c)
            .filter(ch -> !seen.add(ch))
            .findFirst();
        
        firstRepeatedSet.ifPresent(ch -> 
            System.out.println("First repeated (Set): " + ch));
    }
}
```

### 4. Remove Duplicate Characters
**Java 8 Approach**: Using `distinct()` collector

```java
import java.util.*;
import java.util.stream.*;

public class RemoveDuplicatesJava8 {
    public static void main(String[] args) {
        String str = "programming";
        
        // Using distinct() with streams
        String result = str.chars()
            .mapToObj(c -> String.valueOf((char) c))
            .distinct()
            .collect(Collectors.joining());
        
        System.out.println("After removing duplicates: " + result);
        
        // Alternative: Using groupingBy and keeping first occurrence
        String resultOrdered = str.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                LinkedHashMap::new,
                Collectors.counting()
            ))
            .keySet()
            .stream()
            .map(String::valueOf)
            .collect(Collectors.joining());
        
        System.out.println("Ordered result: " + resultOrdered);
    }
}
```

---

## 🔥 Intermediate Level (11-20) - Advanced Java 8

### 5. Longest Substring Without Repeating Characters
**Java 8 Approach**: Using streams with custom collector

```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class LongestUniqueSubstringJava8 {
    public static void main(String[] args) {
        String str = "abcabcbb";
        
        // Using Java 8 Streams with sliding window
        int maxLen = IntStream.range(0, str.length())
            .map(i -> {
                Set<Character> seen = new HashSet<>();
                int j = i;
                while (j < str.length() && seen.add(str.charAt(j))) {
                    j++;
                }
                return j - i;
            })
            .max()
            .orElse(0);
        
        System.out.println("Longest substring length: " + maxLen);
        
        // Alternative: Using parallel stream for large strings
        int maxLenParallel = IntStream.range(0, str.length())
            .parallel()
            .map(i -> {
                Set<Character> seen = new HashSet<>();
                int j = i;
                while (j < str.length() && seen.add(str.charAt(j))) {
                    j++;
                }
                return j - i;
            })
            .max()
            .orElse(0);
    }
}
```

### 6. String Compression
**Java 8 Approach**: Using `Collectors.groupingBy()` with consecutive grouping

```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class StringCompressionJava8 {
    public static void main(String[] args) {
        String str = "aaabbc";
        
        // Using Java 8 Streams with custom grouping
        String compressed = IntStream.range(0, str.length())
            .collect(StringBuilder::new,
                (sb, i) -> {
                    if (i == 0 || str.charAt(i) != str.charAt(i - 1)) {
                        sb.append(str.charAt(i));
                    }
                },
                StringBuilder::append)
            .toString();
        
        // Count consecutive characters
        Map<Character, Long> counts = IntStream.range(0, str.length())
            .mapToObj(i -> str.charAt(i))
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        
        // Build compressed string
        String result = str.chars()
            .mapToObj(c -> (char) c)
            .distinct()
            .map(ch -> ch + String.valueOf(counts.get(ch)))
            .collect(Collectors.joining());
        
        System.out.println("Compressed: " + result);
    }
}
```

### 7. Print All Substrings
**Java 8 Approach**: Using nested streams with `flatMap`

```java
import java.util.stream.*;

public class AllSubstringsJava8 {
    public static void main(String[] args) {
        String str = "ABC";
        
        // Using Java 8 Streams with flatMap
        IntStream.range(0, str.length())
            .boxed()
            .flatMap(start -> 
                IntStream.range(start + 1, str.length() + 1)
                    .mapToObj(end -> str.substring(start, end)))
            .forEach(System.out::println);
        
        // Alternative: Collect to list
        List<String> substrings = IntStream.range(0, str.length())
            .boxed()
            .flatMap(start -> 
                IntStream.range(start + 1, str.length() + 1)
                    .mapToObj(end -> str.substring(start, end)))
            .collect(Collectors.toList());
        
        System.out.println("All substrings: " + substrings);
    }
}
```

### 8. Check Balanced Parentheses
**Java 8 Approach**: Using streams with stack operations

```java
import java.util.*;
import java.util.stream.*;

public class BalancedParenthesesJava8 {
    public static void main(String[] args) {
        String str = "{[()]}";
        
        // Using Java 8 Streams with Stack
        boolean isBalanced = str.chars()
            .mapToObj(c -> (char) c)
            .reduce(
                new Stack<Character>(),
                (stack, ch) -> {
                    if (ch == '(' || ch == '[' || ch == '{') {
                        stack.push(ch);
                    } else if (ch == ')' || ch == ']' || ch == '}') {
                        if (!stack.isEmpty() && isMatching(stack.peek(), ch)) {
                            stack.pop();
                        } else {
                            stack.push(ch); // Mark as unbalanced
                        }
                    }
                    return stack;
                },
                (stack1, stack2) -> stack1
            )
            .isEmpty();
        
        System.out.println("Balanced? " + isBalanced);
    }
    
    private static boolean isMatching(char open, char close) {
        return (open == '(' && close == ')') ||
               (open == '[' && close == ']') ||
               (open == '{' && close == '}');
    }
}
```

---

## 🚀 Advanced Level (21-30) - FAANG Style with Java 8

### 9. Longest Palindrome Substring
**Java 8 Approach**: Using streams with center expansion

```java
import java.util.*;
import java.util.stream.*;

public class LongestPalindromeJava8 {
    public static void main(String[] args) {
        String str = "babad";
        
        // Using Java 8 Streams for center expansion
        String longest = IntStream.range(0, str.length())
            .mapToObj(i -> Arrays.asList(
                expand(str, i, i),      // Odd length
                expand(str, i, i + 1)   // Even length
            ))
            .flatMap(List::stream)
            .max(Comparator.comparingInt(String::length))
            .orElse("");
        
        System.out.println("Longest palindrome: " + longest);
        
        // Alternative: Using parallel stream
        String longestParallel = IntStream.range(0, str.length())
            .parallel()
            .mapToObj(i -> Arrays.asList(
                expand(str, i, i),
                expand(str, i, i + 1)
            ))
            .flatMap(List::stream)
            .max(Comparator.comparingInt(String::length))
            .orElse("");
    }
    
    private static String expand(String str, int left, int right) {
        while (left >= 0 && right < str.length() && 
               str.charAt(left) == str.charAt(right)) {
            left--;
            right++;
        }
        return str.substring(left + 1, right);
    }
}
```

### 10. Remove Adjacent Duplicates
**Java 8 Approach**: Using `Collectors.reducing()` with stack logic

```java
import java.util.*;
import java.util.stream.*;

public class RemoveAdjacentDuplicatesJava8 {
    public static void main(String[] args) {
        String str = "abbaca";
        
        // Using Java 8 Streams with custom reduction
        String result = str.chars()
            .mapToObj(c -> (char) c)
            .collect(Collector.of(
                StringBuilder::new,
                (sb, ch) -> {
                    int len = sb.length();
                    if (len > 0 && sb.charAt(len - 1) == ch) {
                        sb.deleteCharAt(len - 1);
                    } else {
                        sb.append(ch);
                    }
                },
                StringBuilder::append,
                StringBuilder::toString
            ));
        
        System.out.println("Result: " + result);
        
        // Alternative: Using reduce with stack
        String resultReduce = str.chars()
            .mapToObj(c -> (char) c)
            .reduce(
                new Stack<Character>(),
                (stack, ch) -> {
                    if (!stack.isEmpty() && stack.peek() == ch) {
                        stack.pop();
                    } else {
                        stack.push(ch);
                    }
                    return stack;
                },
                (s1, s2) -> s1
            )
            .stream()
            .map(String::valueOf)
            .collect(Collectors.joining());
        
        System.out.println("Result (reduce): " + resultReduce);
    }
}
```

### 11. Check if Strings are Isomorphic
**Java 8 Approach**: Using streams with mapping logic

```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class IsomorphicJava8 {
    public static void main(String[] args) {
        String s1 = "egg", s2 = "add";
        
        // Using Java 8 Streams
        boolean isIsomorphic = s1.length() == s2.length() &&
            IntStream.range(0, s1.length())
                .allMatch(i -> {
                    Map<Character, Character> map = new HashMap<>();
                    Set<Character> used = new HashSet<>();
                    
                    return IntStream.rangeClosed(0, i)
                        .allMatch(j -> {
                            char ch1 = s1.charAt(j), ch2 = s2.charAt(j);
                            if (map.containsKey(ch1)) {
                                return map.get(ch1) == ch2;
                            } else {
                                if (used.contains(ch2)) return false;
                                map.put(ch1, ch2);
                                used.add(ch2);
                                return true;
                            }
                        });
                });
        
        System.out.println("Isomorphic? " + isIsomorphic);
        
        // More elegant approach using collectors
        Map<Character, List<Integer>> s1Positions = IntStream.range(0, s1.length())
            .boxed()
            .collect(Collectors.groupingBy(s1::charAt));
        
        Map<Character, List<Integer>> s2Positions = IntStream.range(0, s2.length())
            .boxed()
            .collect(Collectors.groupingBy(s2::charAt));
        
        boolean isIsomorphicElegant = s1Positions.values().equals(s2Positions.values());
        System.out.println("Isomorphic (elegant)? " + isIsomorphicElegant);
    }
}
```

### 12. Check if Strings are One Edit Away
**Java 8 Approach**: Using streams with edit detection

```java
import java.util.stream.*;

public class OneEditAwayJava8 {
    public static void main(String[] args) {
        String s1 = "pale", s2 = "ple";
        
        // Using Java 8 Streams
        boolean isOneEditAway = s1.equals(s2) ||
            Math.abs(s1.length() - s2.length()) <= 1 &&
            IntStream.range(0, Math.max(s1.length(), s2.length()))
                .filter(i -> {
                    char c1 = i < s1.length() ? s1.charAt(i) : '\0';
                    char c2 = i < s2.length() ? s2.charAt(i) : '\0';
                    return c1 != c2;
                })
                .count() <= 1;
        
        System.out.println("One edit away? " + isOneEditAway);
        
        // More precise approach
        boolean isOneEditPrecise = checkOneEdit(s1, s2);
        System.out.println("One edit away (precise)? " + isOneEditPrecise);
    }
    
    private static boolean checkOneEdit(String s1, String s2) {
        if (s1.equals(s2)) return true;
        
        int len1 = s1.length(), len2 = s2.length();
        if (Math.abs(len1 - len2) > 1) return false;
        
        String shorter = len1 < len2 ? s1 : s2;
        String longer = len1 < len2 ? s2 : s1;
        
        long differences = IntStream.range(0, shorter.length())
            .filter(i -> shorter.charAt(i) != longer.charAt(i))
            .count();
        
        return differences <= 1;
    }
}
```

---

## 🔥 FAANG/Product-Based Questions - Java 8 Style

### 13. Group Anagrams
**Java 8 Approach**: Using `Collectors.groupingBy()` with sorted key

```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class GroupAnagramsJava8 {
    public static void main(String[] args) {
        String[] words = {"eat", "tea", "tan", "ate", "nat", "bat"};
        
        // Using Java 8 Streams
        Map<String, List<String>> anagramGroups = Arrays.stream(words)
            .collect(Collectors.groupingBy(
                word -> word.chars()
                    .sorted()
                    .mapToObj(c -> String.valueOf((char) c))
                    .collect(Collectors.joining())
            ));
        
        // Print groups using lambda
        anagramGroups.forEach((key, group) -> 
            System.out.println(key + ": " + group));
        
        // Alternative: Using parallel stream
        Map<String, List<String>> parallelGroups = Arrays.stream(words)
            .parallel()
            .collect(Collectors.groupingBy(
                word -> word.chars()
                    .sorted()
                    .mapToObj(c -> String.valueOf((char) c))
                    .collect(Collectors.joining())
            ));
        
        // Convert to list format
        List<List<String>> result = new ArrayList<>(anagramGroups.values());
        System.out.println("Grouped anagrams: " + result);
    }
}
```

### 14. Minimum Window Substring
**Java 8 Approach**: Using streams with sliding window

```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class MinWindowSubstringJava8 {
    public static void main(String[] args) {
        String s = "ADOBECODEBANC", t = "ABC";
        
        // Using Java 8 Streams for character frequency
        Map<Character, Long> need = t.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        
        // Sliding window with streams
        String result = IntStream.range(0, s.length())
            .boxed()
            .flatMap(start -> 
                IntStream.range(start + 1, s.length() + 1)
                    .mapToObj(end -> new AbstractMap.SimpleEntry<>(
                        s.substring(start, end), 
                        start
                    ))
            )
            .filter(entry -> containsAllChars(entry.getKey(), need))
            .min(Comparator.comparingInt(entry -> entry.getKey().length()))
            .map(AbstractMap.SimpleEntry::getKey)
            .orElse("");
        
        System.out.println("Minimum window: " + result);
    }
    
    private static boolean containsAllChars(String window, Map<Character, Long> need) {
        Map<Character, Long> windowCount = window.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        
        return need.entrySet().stream()
            .allMatch(entry -> 
                windowCount.getOrDefault(entry.getKey(), 0L) >= entry.getValue());
    }
}
```

### 15. Decode String
**Java 8 Approach**: Using streams with stack operations

```java
import java.util.*;
import java.util.stream.*;

public class DecodeStringJava8 {
    public static void main(String[] args) {
        String str = "3[a2[c]]";
        
        // Using Java 8 Streams with stack processing
        String result = str.chars()
            .mapToObj(c -> (char) c)
            .reduce(
                new Stack<String>(),
                (stack, ch) -> {
                    if (Character.isDigit(ch)) {
                        stack.push(String.valueOf(ch));
                    } else if (ch == '[') {
                        stack.push("[");
                    } else if (ch == ']') {
                        StringBuilder sb = new StringBuilder();
                        while (!stack.isEmpty() && !stack.peek().equals("[")) {
                            sb.insert(0, stack.pop());
                        }
                        stack.pop(); // Remove '['
                        
                        if (!stack.isEmpty() && stack.peek().matches("\\d+")) {
                            int count = Integer.parseInt(stack.pop());
                            String repeated = sb.toString().repeat(count);
                            stack.push(repeated);
                        }
                    } else {
                        if (!stack.isEmpty() && stack.peek().matches("[a-zA-Z]+")) {
                            stack.push(stack.pop() + ch);
                        } else {
                            stack.push(String.valueOf(ch));
                        }
                    }
                    return stack;
                },
                (s1, s2) -> s1
            )
            .stream()
            .collect(Collectors.joining());
        
        System.out.println("Decoded: " + result);
    }
}
```

### 16. Multiply Strings
**Java 8 Approach**: Using streams for digit processing

```java
import java.util.*;
import java.util.stream.*;

public class MultiplyStringsJava8 {
    public static void main(String[] args) {
        String num1 = "123", num2 = "456";
        
        // Using Java 8 Streams for multiplication
        String result = IntStream.range(0, num1.length())
            .mapToObj(i -> {
                int digit1 = num1.charAt(num1.length() - 1 - i) - '0';
                return IntStream.range(0, num2.length())
                    .mapToObj(j -> {
                        int digit2 = num2.charAt(num2.length() - 1 - j) - '0';
                        int product = digit1 * digit2;
                        int position = i + j;
                        return new AbstractMap.SimpleEntry<>(product, position);
                    });
            })
            .flatMap(Function.identity())
            .collect(Collectors.groupingBy(
                AbstractMap.SimpleEntry::getValue,
                TreeMap::new,
                Collectors.summingInt(AbstractMap.SimpleEntry::getKey)
            ))
            .entrySet().stream()
            .sorted(Map.Entry.<Integer, Integer>comparingByKey().reversed())
            .reduce(
                new StringBuilder(),
                (sb, entry) -> {
                    int sum = entry.getValue() + (sb.length() > 0 ? sb.charAt(0) - '0' : 0);
                    sb.insert(0, (char)('0' + (sum % 10)));
                    if (sum >= 10 && entry.getKey() == 0) {
                        sb.insert(0, (char)('0' + (sum / 10)));
                    }
                    return sb;
                },
                StringBuilder::append
            )
            .toString()
            .replaceFirst("^0+", "");
        
        System.out.println("Product: " + result);
    }
}
```

---

## 🎯 Special Java 8 Features Showcase

### 17. String Processing with Optional
```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class StringOptionalJava8 {
    public static void main(String[] args) {
        List<String> strings = Arrays.asList("apple", "banana", "cherry", null, "date");
        
        // Safe processing with Optional
        List<Integer> lengths = strings.stream()
            .map(Optional::ofNullable)
            .map(opt -> opt.map(String::length).orElse(0))
            .collect(Collectors.toList());
        
        System.out.println("Lengths: " + lengths);
        
        // Filter non-null and transform
        List<String> processed = strings.stream()
            .filter(Objects::nonNull)
            .map(String::toUpperCase)
            .filter(s -> s.length() > 4)
            .collect(Collectors.toList());
        
        System.out.println("Processed: " + processed);
        
        // Find first matching element
        Optional<String> firstLong = strings.stream()
            .filter(Objects::nonNull)
            .filter(s -> s.contains("a"))
            .findFirst();
        
        firstLong.ifPresentOrElse(
            s -> System.out.println("Found: " + s),
            () -> System.out.println("Not found")
        );
    }
}
```

### 18. Parallel String Processing
```java
import java.util.*;
import java.util.stream.*;

public class ParallelStringJava8 {
    public static void main(String[] args) {
        String longString = "a".repeat(1000000) + "b".repeat(1000000) + 
                           "c".repeat(1000000);
        
        // Sequential vs Parallel processing
        long startTime = System.currentTimeMillis();
        long sequentialCount = longString.chars()
            .filter(ch -> ch == 'a')
            .count();
        long sequentialTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        long parallelCount = longString.parallelChars()
            .filter(ch -> ch == 'a')
            .count();
        long parallelTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Sequential: " + sequentialCount + " in " + sequentialTime + "ms");
        System.out.println("Parallel: " + parallelCount + " in " + parallelTime + "ms");
        
        // Parallel palindrome check
        boolean isPalindrome = IntStream.range(0, longString.length() / 2)
            .parallel()
            .allMatch(i -> longString.charAt(i) == longString.charAt(longString.length() - 1 - i));
        
        System.out.println("Is palindrome: " + isPalindrome);
    }
}
```

### 19. Custom Collectors for String Processing
```java
import java.util.*;
import java.util.stream.*;
import java.util.function.*;

public class CustomCollectorJava8 {
    public static void main(String[] args) {
        String str = "Hello World Java 8 Streams";
        
        // Custom collector for word frequency
        Map<String, Long> wordFrequency = Arrays.stream(str.split("\\s+"))
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        
        System.out.println("Word frequency: " + wordFrequency);
        
        // Custom collector for character statistics
        CharStats stats = str.chars()
            .mapToObj(c -> (char) c)
            .filter(Character::isLetter)
            .collect(Collector.of(
                CharStats::new,
                (cs, ch) -> {
                    cs.totalChars++;
                    if (Character.isUpperCase(ch)) cs.upperCase++;
                    else cs.lowerCase++;
                    cs.vowels.add(ch);
                },
                (cs1, cs2) -> {
                    cs1.totalChars += cs2.totalChars;
                    cs1.upperCase += cs2.upperCase;
                    cs1.lowerCase += cs2.lowerCase;
                    cs1.vowels.addAll(cs2.vowels);
                    return cs1;
                }
            ));
        
        System.out.println("Stats: " + stats);
    }
    
    static class CharStats {
        int totalChars = 0;
        int upperCase = 0;
        int lowerCase = 0;
        Set<Character> vowels = new HashSet<>();
        
        @Override
        public String toString() {
            return String.format("Total: %d, Upper: %d, Lower: %d, Vowels: %d",
                totalChars, upperCase, lowerCase, 
                vowels.stream().filter(ch -> "AEIOUaeiou".indexOf(ch) >= 0).count());
        }
    }
}
```

### 20. Functional String Operations
```java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class FunctionalStringJava8 {
    public static void main(String[] args) {
        List<String> words = Arrays.asList("java", "stream", "lambda", "functional");
        
        // Function composition
        Function<String, String> toUpper = String::toUpperCase;
        Function<String, String> addPrefix = s -> "PREFIX_" + s;
        Function<String, String> pipeline = toUpper.andThen(addPrefix);
        
        List<String> transformed = words.stream()
            .map(pipeline)
            .collect(Collectors.toList());
        
        System.out.println("Transformed: " + transformed);
        
        // Predicate chaining
        Predicate<String> startsWithJ = s -> s.startsWith("j");
        Predicate<String> lengthGreaterThan4 = s -> s.length() > 4;
        Predicate<String> combined = startsWithJ.and(lengthGreaterThan4);
        
        List<String> filtered = words.stream()
            .filter(combined)
            .collect(Collectors.toList());
        
        System.out.println("Filtered: " + filtered);
        
        // Consumer operations
        Consumer<String> printer = System.out::println;
        Consumer<String> logger = s -> System.err.println("Log: " + s);
        Consumer<String> combinedConsumer = printer.andThen(logger);
        
        words.forEach(combinedConsumer);
        
        // Supplier for string generation
        Supplier<String> randomString = () -> UUID.randomUUID().toString().substring(0, 8);
        List<String> randomStrings = Stream.generate(randomString)
            .limit(5)
            .collect(Collectors.toList());
        
        System.out.println("Random strings: " + randomStrings);
    }
}
```

---

## 📊 Performance Comparison

### Java 7 vs Java 8 Performance
```java
import java.util.*;
import java.util.stream.*;

public class PerformanceComparison {
    public static void main(String[] args) {
        String largeString = "a".repeat(100000) + "b".repeat(100000) + "c".repeat(100000);
        
        // Traditional approach
        long start = System.currentTimeMillis();
        Map<Character, Integer> traditional = new HashMap<>();
        for (char ch : largeString.toCharArray()) {
            traditional.put(ch, traditional.getOrDefault(ch, 0) + 1);
        }
        long traditionalTime = System.currentTimeMillis() - start;
        
        // Java 8 Sequential Stream
        start = System.currentTimeMillis();
        Map<Character, Long> sequential = largeString.chars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        long sequentialTime = System.currentTimeMillis() - start;
        
        // Java 8 Parallel Stream
        start = System.currentTimeMillis();
        Map<Character, Long> parallel = largeString.parallelChars()
            .mapToObj(c -> (char) c)
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        long parallelTime = System.currentTimeMillis() - start;
        
        System.out.println("Traditional: " + traditionalTime + "ms");
        System.out.println("Sequential Stream: " + sequentialTime + "ms");
        System.out.println("Parallel Stream: " + parallelTime + "ms");
    }
}
```

---

## 🎯 Key Java 8 Benefits for String Processing

1. **Readability**: Declarative style vs imperative loops
2. **Conciseness**: Less boilerplate code
3. **Parallelism**: Easy parallel processing
4. **Null Safety**: Optional and method references
5. **Composability**: Function composition and chaining
6. **Type Safety**: Generic collectors and operations

## 📝 Best Practices

1. **Use `Optional`** for null-safe operations
2. **Prefer method references** over lambdas when possible
3. **Use parallel streams** for large datasets only
4. **Custom collectors** for complex aggregations
5. **Function composition** for reusable transformations
6. **Stream pipelines** for data processing workflows

---

*This collection demonstrates how Java 8 features make string processing more elegant, readable, and efficient compared to traditional approaches.*
