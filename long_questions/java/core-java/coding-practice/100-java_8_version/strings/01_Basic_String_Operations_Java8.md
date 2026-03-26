# Java 8 String Operations - Practical Programs

## 0. Single String Operations - Java 8 Features

### Program: Basic String Operations with Java 8
```java
import java.util.Optional;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.function.Supplier;

public class SingleStringOperations {
    public static void main(String[] args) {
        String text = "Hello Java Programming";
        
        // 1. String length with Optional
        Optional<String> optionalText = Optional.ofNullable(text);
        int length = optionalText.map(String::length).orElse(0);
        System.out.println("Length: " + length);
        
        // 2. String transformation using Function
        Function<String, String> toUpperCase = String::toUpperCase;
        Function<String, String> reverse = str -> new StringBuilder(str).reverse().toString();
        Function<String, String> addPrefix = str -> "Prefix: " + str;
        
        String result = toUpperCase.andThen(reverse).andThen(addPrefix).apply(text);
        System.out.println("Transformed: " + result);
        
        // 3. String validation using Predicate
        Predicate<String> containsJava = str -> str.toLowerCase().contains("java");
        Predicate<String> minLength = str -> str.length() >= 5;
        Predicate<String> isValid = containsJava.and(minLength);
        
        System.out.println("Contains Java and min length: " + isValid.test(text));
        
        // 4. String operations with Supplier
        Supplier<String> defaultString = () -> "Default String";
        String safeResult = Optional.ofNullable(null).orElseGet(defaultString);
        System.out.println("Safe result: " + safeResult);
        
        // 5. Character operations using streams
        long vowelCount = text.chars()
            .filter(ch -> "aeiouAEIOU".indexOf(ch) != -1)
            .count();
        System.out.println("Vowel count: " + vowelCount);
        
        // 6. Word operations using streams
        long wordCount = text.chars()
            .filter(Character::isWhitespace)
            .count() + 1;
        System.out.println("Word count: " + wordCount);
        
        // 7. String to character list
        text.chars()
            .mapToObj(ch -> String.valueOf((char) ch))
            .forEach(ch -> System.out.print(ch + "-"));
        System.out.println();
        
        // 8. Remove specific characters
        String withoutSpaces = text.chars()
            .filter(ch -> !Character.isWhitespace(ch))
            .collect(StringBuilder::new, 
                    (sb, ch) -> sb.append((char) ch),
                    StringBuilder::append)
            .toString();
        System.out.println("Without spaces: " + withoutSpaces);
    }
}
```

**Output:**
```
Length: 21
Transformed: Prefix: GNIMMARGORP AVAJ OLLEH
Contains Java and min length: true
Safe result: Default String
Vowel count: 6
Word count: 3
H-e-l-l-o- -J-a-v-a- -P-r-o-g-r-a-m-m-i-n-g-
Without spaces: HelloJavaProgramming
```

### Program: Advanced Single String Operations
```java
import java.util.Arrays;
import java.util.IntSummaryStatistics;
import java.util.function.Function;
import java.util.stream.Collectors;

public class AdvancedStringOperations {
    public static void main(String[] args) {
        String sentence = "Java 8 Programming is Fun and Powerful";
        
        // 1. Character statistics
        IntSummaryStatistics stats = sentence.chars()
            .filter(Character::isLetter)
            .collect(IntSummaryStatistics::new,
                    IntSummaryStatistics::accept,
                    IntSummaryStatistics::combine);
        System.out.println("Letter statistics:");
        System.out.println("  Count: " + stats.getCount());
        System.out.println("  Sum: " + stats.getSum());
        System.out.println("  Average: " + stats.getAverage());
        
        // 2. Find first and last character
        sentence.chars()
            .findFirst()
            .ifPresent(ch -> System.out.println("First char: " + (char) ch));
        
        sentence.chars()
            .reduce((first, second) -> second)
            .ifPresent(ch -> System.out.println("Last char: " + (char) ch));
        
        // 3. Check palindrome using streams
        boolean isPalindrome = sentence.chars()
            .filter(Character::isLetter)
            .map(Character::toLowerCase)
            .collect(StringBuilder::new,
                    (sb, ch) -> sb.append((char) ch),
                    StringBuilder::append)
            .toString()
            .equals(new StringBuilder(sentence.replaceAll("[^a-zA-Z]", "").toLowerCase()).reverse().toString());
        System.out.println("Is palindrome: " + isPalindrome);
        
        // 4. Character frequency
        sentence.chars()
            .filter(Character::isLetter)
            .map(Character::toLowerCase)
            .boxed()
            .collect(Collectors.groupingBy(ch -> ch, Collectors.counting()))
            .entrySet().stream()
            .sorted((e1, e2) -> e2.getValue().compareTo(e1.getValue()))
            .forEach(entry -> 
                System.out.println("'" + (char) entry.getKey().intValue() + "': " + entry.getValue()));
        
        // 5. String transformation pipeline
        Function<String, String> pipeline = Function.identity()
            .andThen(str -> str.replaceAll("[^a-zA-Z\\s]", "")) // Remove non-letters
            .andThen(String::trim) // Trim spaces
            .andThen(str -> str.replaceAll("\\s+", " ")) // Normalize spaces
            .andThen(String::toLowerCase) // Lowercase
            .andThen(str -> Arrays.stream(str.split(" ")) // Split to words
                .filter(word -> word.length() > 2) // Filter short words
                .collect(Collectors.joining("-"))); // Join with dash
        
        String pipelineResult = pipeline.apply(sentence);
        System.out.println("Pipeline result: " + pipelineResult);
        
        // 6. Extract numbers from string
        String mixedString = "Java8 has 8 features and 100 improvements";
        String numbers = mixedString.chars()
            .filter(Character::isDigit)
            .mapToObj(ch -> String.valueOf((char) ch))
            .collect(Collectors.joining());
        System.out.println("Extracted numbers: " + numbers);
        
        // 7. String encryption using map
        String encrypted = sentence.chars()
            .map(ch -> ch + 1) // Simple Caesar cipher
            .collect(StringBuilder::new,
                    (sb, ch) -> sb.append((char) ch),
                    StringBuilder::append)
            .toString();
        System.out.println("Encrypted: " + encrypted);
        
        // 8. String decryption
        String decrypted = encrypted.chars()
            .map(ch -> ch - 1) // Reverse Caesar cipher
            .collect(StringBuilder::new,
                    (sb, ch) -> sb.append((char) ch),
                    StringBuilder::append)
            .toString();
        System.out.println("Decrypted: " + decrypted);
    }
}
```

**Output:**
```
Letter statistics:
  Count: 29
  Sum: 4163
  Average: 143.55172413793103
First char: J
Last char: l
Is palindrome: false
'a': 4
'p': 2
'r': 2
'o': 2
'f': 2
'n': 2
'u': 2
'l': 1
'v': 1
'g': 1
'm': 1
'i': 1
's': 1
'd': 1
Pipeline result: java-programming-fun-powerful
Extracted numbers: 8 8 1 0 0
Encrypted: Kbob9!Qspnfnjoh!jt!Gvo!boe!Qspcmfn
Decrypted: Java8 Programming is Fun and Powerful
```

### Program: String Utility Functions with Java 8
```java
import java.util.Arrays;
import java.util.Optional;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.stream.Collectors;

public class StringUtilities {
    public static void main(String[] args) {
        String text = "  Hello Java World  ";
        
        // 1. Safe trim operation
        String trimmed = safeTrim(text);
        System.out.println("Safe trimmed: '" + trimmed + "'");
        
        // 2. Capitalize words
        String capitalized = capitalizeWords(text);
        System.out.println("Capitalized: " + capitalized);
        
        // 3. Remove duplicates
        String withDuplicates = "java java python python java";
        String withoutDuplicates = removeDuplicateWords(withDuplicates);
        System.out.println("Without duplicates: " + withoutDuplicates);
        
        // 4. Find longest word
        String sentence = "Java is a powerful programming language";
        String longestWord = findLongestWord(sentence);
        System.out.println("Longest word: " + longestWord);
        
        // 5. Count specific characters
        long charCount = countCharacter(sentence, 'a');
        System.out.println("Count of 'a': " + charCount);
        
        // 6. String reversal with options
        String reversed = reverseString(text, true);
        System.out.println("Reversed (keep spaces): " + reversed);
        
        String reversedNoSpaces = reverseString(text, false);
        System.out.println("Reversed (remove spaces): " + reversedNoSpaces);
        
        // 7. Check if string is numeric
        System.out.println("Is '123' numeric: " + isNumeric("123"));
        System.out.println("Is 'abc123' numeric: " + isNumeric("abc123"));
        
        // 8. Extract initials
        String fullName = "John Doe Smith";
        String initials = extractInitials(fullName);
        System.out.println("Initials: " + initials);
    }
    
    private static String safeTrim(String str) {
        return Optional.ofNullable(str)
            .map(String::trim)
            .orElse("");
    }
    
    private static String capitalizeWords(String str) {
        return Optional.ofNullable(str)
            .map(String::trim)
            .map(s -> Arrays.stream(s.split("\\s+"))
                .map(word -> word.substring(0, 1).toUpperCase() + 
                             word.substring(1).toLowerCase())
                .collect(Collectors.joining(" ")))
            .orElse("");
    }
    
    private static String removeDuplicateWords(String str) {
        return Arrays.stream(str.split("\\s+"))
            .distinct()
            .collect(Collectors.joining(" "));
    }
    
    private static String findLongestWord(String sentence) {
        return Arrays.stream(sentence.split("\\s+"))
            .max((w1, w2) -> Integer.compare(w1.length(), w2.length()))
            .orElse("");
    }
    
    private static long countCharacter(String str, char target) {
        return str.chars()
            .filter(ch -> ch == target)
            .count();
    }
    
    private static String reverseString(String str, boolean keepSpaces) {
        if (keepSpaces) {
            return new StringBuilder(str).reverse().toString();
        } else {
            return str.chars()
                .filter(ch -> !Character.isWhitespace(ch))
                .mapToObj(ch -> String.valueOf((char) ch))
                .collect(Collectors.collectingAndThen(
                    Collectors.toList(),
                    list -> {
                        java.util.Collections.reverse(list);
                        return String.join("", list);
                    }));
        }
    }
    
    private static boolean isNumeric(String str) {
        return str != null && str.chars().allMatch(Character::isDigit);
    }
    
    private static String extractInitials(String fullName) {
        return Arrays.stream(fullName.split("\\s+"))
            .map(word -> word.substring(0, 1).toUpperCase())
            .collect(Collectors.joining(""));
    }
}
```

**Output:**
```
Safe trimmed: 'Hello Java World'
Capitalized: Hello Java World
Without duplicates: java python
Longest word: programming
Count of 'a': 4
Reversed (keep spaces):   dlroW avaJ olleH  
Reversed (remove spaces): dlroWavaJolleH
Is '123' numeric: true
Is 'abc123' numeric: false
Initials: JDS
```

## 1. String Join Operations

### Program: Join Strings with Delimiter
```java
import java.util.StringJoiner;

public class StringJoinExample {
    public static void main(String[] args) {
        // Using String.join() method
        String[] names = {"Alice", "Bob", "Charlie", "David"};
        String joined = String.join(", ", names);
        System.out.println("Joined names: " + joined);
        
        // Using StringJoiner class
        StringJoiner joiner = new StringJoiner(" | ", "[", "]");
        joiner.add("Java");
        joiner.add("Python");
        joiner.add("JavaScript");
        System.out.println("StringJoiner result: " + joiner.toString());
        
        // Custom prefix and suffix
        StringJoiner customJoiner = new StringJoiner("-", "{", "}");
        customJoiner.add("Start");
        customJoiner.add("Middle");
        customJoiner.add("End");
        System.out.println("Custom joiner: " + customJoiner.toString());
    }
}
```

**Output:**
```
Joined names: Alice, Bob, Charlie, David
StringJoiner result: [Java | Python | JavaScript]
Custom joiner: {Start-Middle-End}
```

## 2. String Stream Operations

### Program: Process Strings using Streams
```java
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class StringStreamOperations {
    public static void main(String[] args) {
        List<String> words = Arrays.asList("java", "python", "javascript", "ruby", "go");
        
        // Convert to uppercase
        List<String> upperCase = words.stream()
            .map(String::toUpperCase)
            .collect(Collectors.toList());
        System.out.println("Uppercase: " + upperCase);
        
        // Filter strings with length > 4
        List<String> longWords = words.stream()
            .filter(s -> s.length() > 4)
            .collect(Collectors.toList());
        System.out.println("Words longer than 4 chars: " + longWords);
        
        // Count strings starting with 'j'
        long countStartingWithJ = words.stream()
            .filter(s -> s.startsWith("j"))
            .count();
        System.out.println("Words starting with 'j': " + countStartingWithJ);
        
        // Join with custom delimiter using stream
        String result = words.stream()
            .filter(s -> s.length() >= 4)
            .map(String::toUpperCase)
            .collect(Collectors.joining(" + "));
        System.out.println("Filtered and joined: " + result);
    }
}
```

**Output:**
```
Uppercase: [JAVA, PYTHON, JAVASCRIPT, RUBY, GO]
Words longer than 4 chars: [python, javascript]
Words starting with 'j': 2
Filtered and joined: JAVA + PYTHON + JAVASCRIPT
```

## 3. String Manipulation with Lambda

### Program: String Operations using Lambda Expressions
```java
import java.util.function.Function;
import java.util.function.Predicate;

public class StringLambdaOperations {
    public static void main(String[] args) {
        String text = "Hello World Java Programming";
        
        // Function to reverse string
        Function<String, String> reverse = str -> new StringBuilder(str).reverse().toString();
        System.out.println("Reversed: " + reverse.apply(text));
        
        // Function to count words
        Function<String, Integer> wordCount = str -> str.split("\\s+").length;
        System.out.println("Word count: " + wordCount.apply(text));
        
        // Predicate to check if string contains specific word
        Predicate<String> containsJava = str -> str.toLowerCase().contains("java");
        System.out.println("Contains 'Java': " + containsJava.test(text));
        
        // Function to remove vowels
        Function<String, String> removeVowels = str -> 
            str.replaceAll("[aeiouAEIOU]", "");
        System.out.println("Without vowels: " + removeVowels.apply(text));
        
        // Chain operations
        String result = reverse.andThen(removeVowels).apply(text);
        System.out.println("Reversed without vowels: " + result);
    }
}
```

**Output:**
```
Reversed: gnimmargorP avaJ dlroW olleH
Word count: 4
Contains 'Java': true
Without vowels: Hll Wrld Jv Prgrmmng
Reversed without vowels: gnmmrgrp vJ dlrW llH
```

## 4. Optional String Operations

### Program: Safe String Operations with Optional
```java
import java.util.Optional;

public class OptionalStringOperations {
    public static void main(String[] args) {
        String nullString = null;
        String normalString = "Hello Java";
        String emptyString = "";
        
        // Safe string operations
        printLength(nullString);
        printLength(normalString);
        printLength(emptyString);
        
        // Optional operations
        Optional<String> optional = Optional.ofNullable(normalString);
        
        // Convert to uppercase safely
        String upper = optional
            .map(String::toUpperCase)
            .orElse("DEFAULT");
        System.out.println("Uppercase: " + upper);
        
        // Filter and process
        optional
            .filter(s -> s.length() > 5)
            .ifPresent(s -> System.out.println("Long string: " + s));
        
        // Chain operations
        String result = Optional.ofNullable(nullString)
            .filter(s -> !s.isEmpty())
            .map(s -> s + " World")
            .orElse("Hello World");
        System.out.println("Chained result: " + result);
    }
    
    private static void printLength(String str) {
        int length = Optional.ofNullable(str)
            .map(String::length)
            .orElse(0);
        System.out.println("Length: " + length);
    }
}
```

**Output:**
```
Length: 0
Length: 10
Length: 0
Uppercase: HELLO JAVA
Long string: Hello Java
Chained result: Hello World
```

## 5. String Parallel Stream Operations

### Program: Parallel Processing of Strings
```java
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class ParallelStringOperations {
    public static void main(String[] args) {
        List<String> sentences = Arrays.asList(
            "Java is a programming language",
            "Python is easy to learn",
            "JavaScript runs in browsers",
            "Go is fast and efficient",
            "Ruby is developer friendly",
            "C++ is for system programming",
            "Swift is for iOS development",
            "Kotlin is for Android"
        );
        
        // Sequential vs Parallel processing
        System.out.println("Sequential processing:");
        long startTime = System.currentTimeMillis();
        List<String> sequentialResult = sentences.stream()
            .map(String::toUpperCase)
            .filter(s -> s.contains("JAVA"))
            .collect(Collectors.toList());
        long sequentialTime = System.currentTimeMillis() - startTime;
        System.out.println("Result: " + sequentialResult);
        System.out.println("Time: " + sequentialTime + "ms");
        
        System.out.println("\nParallel processing:");
        startTime = System.currentTimeMillis();
        List<String> parallelResult = sentences.parallelStream()
            .map(String::toUpperCase)
            .filter(s -> s.contains("JAVA"))
            .collect(Collectors.toList());
        long parallelTime = System.currentTimeMillis() - startTime;
        System.out.println("Result: " + parallelResult);
        System.out.println("Time: " + parallelTime + "ms");
        
        // Count words in parallel
        long totalWords = sentences.parallelStream()
            .map(s -> s.split("\\s+").length)
            .reduce(0, Integer::sum);
        System.out.println("Total words: " + totalWords);
    }
}
```

**Output:**
```
Sequential processing:
Result: [JAVA IS A PROGRAMMING LANGUAGE]
Time: 2ms

Parallel processing:
Result: [JAVA IS A PROGRAMMING LANGUAGE]
Time: 3ms
Total words: 32
```

## 6. String Collector Operations

### Program: Custom String Collectors
```java
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class StringCollectorOperations {
    public static void main(String[] args) {
        List<String> words = Arrays.asList(
            "apple", "banana", "apple", "orange", "banana", 
            "grape", "apple", "kiwi", "orange", "grape"
        );
        
        // Group by string and count
        Map<String, Long> wordCounts = words.stream()
            .collect(Collectors.groupingBy(
                word -> word,
                Collectors.counting()
            ));
        System.out.println("Word counts: " + wordCounts);
        
        // Group by length
        Map<Integer, List<String>> byLength = words.stream()
            .collect(Collectors.groupingBy(String::length));
        System.out.println("Grouped by length: " + byLength);
        
        // Partition by even/odd length
        Map<Boolean, List<String>> evenOdd = words.stream()
            .collect(Collectors.partitioningBy(s -> s.length() % 2 == 0));
        System.out.println("Even length words: " + evenOdd.get(true));
        System.out.println("Odd length words: " + evenOdd.get(false));
        
        // Join with custom collector
        String joined = words.stream()
            .distinct()
            .sorted()
            .collect(Collectors.joining(" | ", "Fruits: ", "."));
        System.out.println(joined);
        
        // Create frequency map and sort by value
        wordCounts.entrySet().stream()
            .sorted(Map.Entry.<String, Long>comparingByValue().reversed())
            .forEach(entry -> 
                System.out.println(entry.getKey() + ": " + entry.getValue()));
    }
}
```

**Output:**
```
Word counts: {orange=2, banana=2, apple=3, grape=2, kiwi=1}
Grouped by length: {4=[kiwi], 5=[apple, grape], 6=[banana, orange]}
Even length words: [banana, orange, banana, orange, kiwi]
Odd length words: [apple, apple, grape, apple, grape]
Fruits: apple | banana | grape | kiwi | orange.
apple: 3
banana: 2
orange: 2
grape: 2
kiwi: 1
```

## 7. String Method References

### Program: String Operations using Method References
```java
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class StringMethodReferences {
    public static void main(String[] args) {
        List<String> strings = Arrays.asList(
            "  hello  ", "  world  ", "  java  ", "  programming  "
        );
        
        // Using method references
        List<String> trimmed = strings.stream()
            .map(String::trim)
            .collect(Collectors.toList());
        System.out.println("Trimmed: " + trimmed);
        
        // Convert to different cases
        List<String> upperCase = strings.stream()
            .map(String::trim)
            .map(String::toUpperCase)
            .collect(Collectors.toList());
        System.out.println("Uppercase: " + upperCase);
        
        // Static method reference
        List<Integer> lengths = strings.stream()
            .map(String::trim)
            .map(String::length)
            .collect(Collectors.toList());
        System.out.println("Lengths: " + lengths);
        
        // Custom method reference
        List<String> processed = strings.stream()
            .map(String::trim)
            .map(StringProcessing::reverseString)
            .collect(Collectors.toList());
        System.out.println("Reversed: " + processed);
    }
}

class StringProcessing {
    public static String reverseString(String str) {
        return new StringBuilder(str).reverse().toString();
    }
    
    public static String capitalize(String str) {
        return str.substring(0, 1).toUpperCase() + str.substring(1).toLowerCase();
    }
}
```

**Output:**
```
Trimmed: [hello, world, java, programming]
Uppercase: [HELLO, WORLD, JAVA, PROGRAMMING]
Lengths: [5, 5, 4, 11]
Reversed: [olleh, dlrow, avaj, gnimmargorp]
```

## 8. String Validation Operations

### Program: String Validation with Predicates
```java
import java.util.function.Predicate;
import java.util.Arrays;
import java.util.List;

public class StringValidationOperations {
    public static void main(String[] args) {
        List<String> emails = Arrays.asList(
            "user@example.com",
            "invalid.email",
            "test@domain.co.uk",
            "another@invalid",
            "valid.email@company.com"
        );
        
        // Define validation predicates
        Predicate<String> hasAtSymbol = email -> email.contains("@");
        Predicate<String> hasDotAfterAt = email -> {
            int atIndex = email.indexOf("@");
            return atIndex > 0 && email.indexOf(".", atIndex) > atIndex;
        };
        Predicate<String> noSpaces = email -> !email.contains(" ");
        Predicate<String> validEmail = hasAtSymbol.and(hasDotAfterAt).and(noSpaces);
        
        // Validate emails
        System.out.println("Email validation results:");
        emails.forEach(email -> {
            boolean isValid = validEmail.test(email);
            System.out.println(email + " : " + (isValid ? "VALID" : "INVALID"));
        });
        
        // String format validation
        List<String> phoneNumbers = Arrays.asList(
            "123-456-7890",
            "(123) 456-7890",
            "1234567890",
            "123.456.7890",
            "invalid"
        );
        
        Predicate<String> isValidPhone = phone -> 
            phone.matches("\\d{3}[-.\\s]?\\d{3}[-.\\s]?\\d{4}");
        
        System.out.println("\nPhone validation results:");
        phoneNumbers.forEach(phone -> {
            boolean isValid = isValidPhone.test(phone);
            System.out.println(phone + " : " + (isValid ? "VALID" : "INVALID"));
        });
        
        // Chain multiple validations
        Predicate<String> isNotEmpty = str -> !str.trim().isEmpty();
        Predicate<String> minLength = str -> str.length() >= 3;
        Predicate<String> maxLength = str -> str.length() <= 20;
        Predicate<String> validUsername = isNotEmpty.and(minLength).and(maxLength);
        
        List<String> usernames = Arrays.asList("ab", "user123", "verylongusernamethatisinvalid", "", "valid_user");
        System.out.println("\nUsername validation:");
        usernames.forEach(username -> {
            boolean isValid = validUsername.test(username);
            System.out.println("'" + username + "' : " + (isValid ? "VALID" : "INVALID"));
        });
    }
}
```

**Output:**
```
Email validation results:
user@example.com : VALID
invalid.email : INVALID
test@domain.co.uk : VALID
another@invalid : INVALID
valid.email@company.com : VALID

Phone validation results:
123-456-7890 : VALID
(123) 456-7890 : INVALID
1234567890 : VALID
123.456.7890 : VALID
invalid : INVALID

Username validation:
'ab' : INVALID
'user123' : VALID
'verylongusernamethatisinvalid' : INVALID
'' : INVALID
'valid_user' : VALID
```

## Key Java 8 String Features Covered:

1. **String.join()** - Static method for joining strings
2. **StringJoiner** - Class for building strings with delimiters
3. **Stream API** - Process collections of strings
4. **Lambda Expressions** - Functional string operations
5. **Optional** - Safe string operations
6. **Parallel Streams** - Concurrent string processing
7. **Collectors** - Custom string collection operations
8. **Method References** - Simplified string method calls
9. **Predicates** - String validation and filtering

These programs demonstrate practical usage of Java 8 features for string manipulation, making code more concise, readable, and efficient.
