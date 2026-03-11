# Java Strings & StringBuilder — Practical Code Snippets

> **Topics:** String immutability, String pool, String methods, StringBuilder, StringJoiner, String formatting, Regular expressions, String comparison

---

## 📋 Reading Progress

- [ ] **Section 1:** String Immutability & Pool (Q1–Q12)
- [ ] **Section 2:** String Methods (Q13–Q28)
- [ ] **Section 3:** StringBuilder & StringBuffer (Q29–Q42)
- [ ] **Section 4:** String Formatting & Parsing (Q43–Q52)
- [ ] **Section 5:** String & Collections (Q53–Q60)

> 🔖 **Last read:** <!-- -->

---

## Section 1: String Immutability & Pool (Q1–Q12)

### 1. String is Immutable
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "hello";
        s.toUpperCase(); // returns a NEW string, doesn't modify s
        System.out.println(s);
        String upper = s.toUpperCase();
        System.out.println(upper);
    }
}
```
**A:**
```
hello
HELLO
```
String methods never modify the original — they return new String objects. String is immutable.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why doesn't the first string change?"

**Your Response:** "The output shows 'hello' then 'HELLO'. This demonstrates String immutability - one of Java's fundamental concepts. When we call `s.toUpperCase()`, it doesn't modify the original string 's'. Instead, it returns a completely new String object with the uppercase version. The original 's' remains unchanged because String objects are immutable - their state cannot be modified after creation. This is why we need to assign the result to a new variable 'upper'. Immutability makes Strings thread-safe and allows for optimizations like string pooling and substring sharing."

---

### 2. String Pool — Literals vs new

**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s1 = "hello";
        String s2 = "hello";          // same pool object
        String s3 = new String("hello"); // new heap object
        System.out.println(s1 == s2);
        System.out.println(s1 == s3);
        System.out.println(s1.equals(s3));
    }
}
```
**A:**
```
true
false
true
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the String pool?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This demonstrates the String pool and the difference between string literals and objects created with 'new'. 's1' and 's2' both reference the same object in the string pool because they're identical literals - Java optimizes by reusing the same object. 's3' is created with 'new String()', so it gets a separate object on the heap, hence 's1 == s3' is false. However, 's1.equals(s3)' is true because equals() compares the actual character sequences, not object references. The string pool is a memory optimization for string literals."

---

### 3. String Concatenation Creates New Object
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "Hello";
        String s2 = s + " World"; // creates a new String object
        System.out.println(System.identityHashCode(s));
        System.out.println(System.identityHashCode(s2));
        System.out.println(s2);
    }
}
```
**A:** Two different hash codes and `Hello World`. Concatenation always creates a new `String` object.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why are the hash codes different?"

**Your Response:** "The output shows two different identity hash codes and 'Hello World'. This proves that string concatenation creates a new String object. When we do `s + " World"`, Java doesn't modify the original string 's' - it creates a completely new String object containing the concatenated result. The `System.identityHashCode()` shows us that 's' and 's2' are different objects in memory. This is a direct consequence of String immutability - every operation that appears to modify a string actually creates a new one. For frequent string modifications, StringBuilder is more efficient."

---

### 4. intern() — Pool Lookup
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s1 = new String("hello").intern(); // put in pool
        String s2 = "hello";
        System.out.println(s1 == s2);
    }
}
```
**A:** `true`. `intern()` returns the pooled canonical representation.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does intern() work?"

**Your Response:** "The output is 'true'. This shows how the `intern()` method works with the string pool. When we create `new String("hello")`, it creates a new object on the heap. But when we call `intern()` on it, Java checks if an equivalent string already exists in the string pool. Since 'hello' is already in the pool (from the literal), `intern()` returns a reference to that pooled object instead of the heap object. So both 's1' and 's2' end up referencing the same pooled object, making '==' return true. This is useful for memory optimization and when you need string equality with reference comparison."

---

### 5. Compile-Time Constant Folding
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s1 = "hello" + " " + "world"; // compile-time constant → pooled
        String s2 = "hello world";
        System.out.println(s1 == s2);
    }
}
```
**A:** `true`. The compiler folds compile-time string concatenations into a single literal, which is pooled.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why does this happen at compile time?"

**Your Response:** "The output is 'true'. This demonstrates compile-time constant folding by the Java compiler. When we have `"hello" + " " + "world"`, the compiler recognizes these as compile-time constants and combines them into a single string literal `"hello world"` during compilation. This literal is then placed in the string pool. Since both 's1' and 's2' end up referencing the same pooled literal, '==' returns true. This optimization only works with compile-time constants - if any part involves variables, the concatenation happens at runtime and creates new objects."

---

### 6. String Comparison — Always Use equals()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "HELLO";
        System.out.println(s.equals("hello"));
        System.out.println(s.equalsIgnoreCase("hello"));
        System.out.println("hello".compareTo("HELLO")); // case-sensitive comparison
    }
}
```
**A:**
```
false
true
32
```
`compareTo` returns the Unicode difference of first differing characters ('h' - 'H' = 104 - 72 = 32).

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between these comparison methods?"

**Your Response:** "The output shows 'false', 'true', then '32'. This demonstrates different string comparison approaches in Java. `equals()` performs exact case-sensitive comparison, so 'HELLO' doesn't equal 'hello'. `equalsIgnoreCase()` ignores case differences and returns true. `compareTo()` performs lexicographical comparison and returns the Unicode difference between the first differing characters - in this case, 'h' (104) minus 'H' (72) equals 32. This is useful for sorting strings. Always use `equals()` for equality checks and `compareTo()` for ordering/sorting."

---

### 7. String is a CharSequence
**Q: Does this compile?**
```java
public class Main {
    static void print(CharSequence cs) { System.out.println(cs.length()); }
    public static void main(String[] args) {
        print("hello");               // String implements CharSequence
        print(new StringBuilder("world")); // StringBuilder implements CharSequence
    }
}
```
**A:** **Yes, compiles and prints:**
```
5
5
```
`String`, `StringBuilder`, `StringBuffer`, and `CharBuffer` all implement `CharSequence`.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Does this compile and what's the significance of CharSequence?"

**Your Response:** "Yes, this compiles and prints '5' and '5'. This demonstrates the `CharSequence` interface, which is a common abstraction for character sequences in Java. `String`, `StringBuilder`, `StringBuffer`, and even `CharBuffer` all implement `CharSequence`, making them interchangeable in many APIs. This interface provides methods like `length()`, `charAt()`, and `subSequence()`. It's useful for writing flexible methods that can work with different character sequence types without being tied to a specific implementation. This is a great example of programming to an interface rather than an implementation."

---

### 8. String.isEmpty() vs String.isBlank()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String empty = "";
        String blank = "   "; // only whitespace
        System.out.println(empty.isEmpty());
        System.out.println(blank.isEmpty());
        System.out.println(blank.isBlank()); // Java 11+
    }
}
```
**A:**
```
true
false
true
```
`isEmpty()` checks length == 0. `isBlank()` (Java 11+) also returns true for strings containing only whitespace.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between these methods?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This illustrates the difference between `isEmpty()` and `isBlank()` methods. `isEmpty()` checks if the string length is exactly zero - so only an empty string returns true. `isBlank()` was introduced in Java 11 and checks if the string is empty OR contains only whitespace characters. So the string with three spaces returns false for `isEmpty()` but true for `isBlank()`. This is particularly useful for input validation where you want to treat strings with only whitespace as empty."

---

### 9. String Null Safety
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = null;
        System.out.println("Value: " + s);    // null concatenation is safe
        try { System.out.println(s.length()); } // NullPointerException!
        catch (NullPointerException e) { System.out.println("NPE!"); }
    }
}
```
**A:**
```
Value: null
NPE!
```
`null` can be concatenated (becomes the string "null") but you cannot call methods on it.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why does concatenation work but method calls don't?"

**Your Response:** "The output shows 'Value: null' then 'NPE!'. This demonstrates Java's null handling in string operations. When concatenating a null string using the '+' operator, Java safely converts it to the literal string 'null'. However, when you try to call methods like `length()` on a null reference, you get a NullPointerException. This is because the '+' operator has built-in null safety - it internally uses `String.valueOf()` which handles null gracefully. But direct method calls don't have this protection. This is why you should always check for null before calling string methods, or use utility methods like `String.valueOf()` for safety."

---

### 10. String.valueOf() for Null Safety
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Object obj = null;
        System.out.println(String.valueOf(obj)); // safe null → "null"
        System.out.println(obj.toString());      // NullPointerException!
    }
}
```
**A:** `null` then **NullPointerException**. Use `String.valueOf()` when the object might be null.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why should you use String.valueOf()?"

**Your Response:** "The output is 'null' followed by a NullPointerException. This shows the importance of null-safe string conversion. `String.valueOf(obj)` safely handles null values and converts them to the literal string 'null'. However, calling `toString()` directly on a null reference throws an NPE. This is why `String.valueOf()` is preferred in production code, especially when dealing with potentially null objects from external sources or databases. It's a defensive programming practice that prevents unexpected crashes. Many frameworks and libraries use this pattern for robust null handling."

---

### 11. Character Case Methods
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        char c = 'a';
        System.out.println(Character.isLetter(c));
        System.out.println(Character.isDigit(c));
        System.out.println(Character.toUpperCase(c));
        System.out.println(Character.isWhitespace(' '));
    }
}
```
**A:**
```
true
false
A
true
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these Character methods do?"

**Your Response:** "The output shows 'true', 'false', 'A', then 'true'. This demonstrates the Character wrapper class utility methods. `isLetter('a')` returns true because 'a' is a letter. `isDigit('a')` returns false because it's not a digit. `toUpperCase('a')` returns 'A' - the uppercase version. `isWhitespace(' ')` returns true because space is considered whitespace. These static methods are essential for character-level validation and processing. They handle Unicode properly and are more reliable than manual ASCII comparisons. Common use cases include input validation, parsing, and text processing."

---

### 12. String Constant Pool — Impact on Memory
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String a = "abc";
        String b = "ab" + "c";   // compile-time folding → same pool entry
        String c = "ab";
        String d = c + "c";      // runtime concatenation → new heap object
        System.out.println(a == b);
        System.out.println(a == d);
        System.out.println(a.equals(d));
    }
}
```
**A:**
```
true
false
true
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does this show about compile-time vs runtime string creation?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This demonstrates the difference between compile-time constant folding and runtime string concatenation. 'a == b' is true because 'ab' + 'c' gets folded into 'abc' at compile time, so both reference the same pool entry. 'a == d' is false because 'c + \"c\"' happens at runtime, creating a new heap object. However, 'a.equals(d)' is true because equals() compares content, not references. This is a crucial concept for understanding Java's string optimization and memory usage. It shows why string literals are efficient but runtime concatenation creates new objects."

---

## Section 2: String Methods (Q13–Q28)

### 13. substring() — Indices
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "Hello World";
        System.out.println(s.substring(6));     // from index 6 to end
        System.out.println(s.substring(0, 5));  // from 0, exclusive end 5
        System.out.println(s.substring(6, 11)); // "World"
    }
}
```
**A:**
```
World
Hello
World
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does substring work with indices?"

**Your Response:** "The output shows 'World', 'Hello', then 'World'. This demonstrates the `substring()` method behavior with different index parameters. `substring(6)` returns characters from index 6 to the end of the string. `substring(0, 5)` returns characters from index 0 up to, but not including, index 5. `substring(6, 11)` returns characters from index 6 to 10 (exclusive of 11). The key point is that the second parameter is exclusive, not inclusive. This is a common source of off-by-one errors. Substring is zero-indexed, so 'H' is at index 0, 'e' at index 1, and so on."

---

### 14. indexOf and lastIndexOf
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "abcabc";
        System.out.println(s.indexOf('c'));      // first 'c'
        System.out.println(s.lastIndexOf('c'));  // last 'c'
        System.out.println(s.indexOf("abc", 1)); // search from index 1
        System.out.println(s.indexOf("xyz"));    // not found
    }
}
```
**A:**
```
2
5
3
-1
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between indexOf and lastIndexOf?"

**Your Response:** "The output shows '2', '5', '3', then '-1'. This demonstrates string searching methods. `indexOf('c')` returns the first occurrence of 'c' at index 2. `lastIndexOf('c')` returns the last occurrence at index 5. `indexOf("abc", 1)` searches for "abc" starting from index 1, finding it at index 3. `indexOf("xyz")` returns -1 because the substring doesn't exist. These methods are fundamental for string parsing and searching. The fromIndex parameter is useful for finding multiple occurrences or implementing search algorithms."

---

### 15. replace() vs replaceAll()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "aaa.bbb.ccc";
        System.out.println(s.replace('.', '-'));       // literal char replace
        System.out.println(s.replace(".", "!"));       // literal string replace
        System.out.println(s.replaceAll("\\.", "#")); // regex replace
        System.out.println("abc123".replaceAll("[0-9]", "*")); // regex
    }
}
```
**A:**
```
aaa-bbb-ccc
aaa!bbb!ccc
aaa#bbb#ccc
abc***
```
`replaceAll()` uses regex. `replace()` uses literal strings. The `.` in regex means "any char" so it must be escaped as `\\.`.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the critical difference between replace and replaceAll?"

**Your Response:** "The output shows 'aaa-bbb-ccc', 'aaa!bbb!ccc', 'aaa#bbb#ccc', then 'abc***'. This demonstrates the crucial difference between `replace()` and `replaceAll()` methods. `replace()` works with literal strings - it replaces exact character or string matches. `replaceAll()` uses regular expressions, so `.` means 'any character' and needs to be escaped as `\\.` to match a literal dot. This is a common interview pitfall - many developers mistakenly use `replaceAll()` when they mean `replace()`. Always use `replace()` for literal replacements and `replaceAll()` only when you need regex functionality."

---

### 16. split()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String csv = "a,b,c,d,";
        String[] parts = csv.split(",");
        System.out.println(parts.length);
        // Trailing empty strings are discarded by default!
        String[] parts2 = csv.split(",", -1); // -1 keeps trailing empties
        System.out.println(parts2.length);
    }
}
```
**A:**
```
4
5
```
`split()` discards trailing empty strings by default. Use `split(",", -1)` to keep them.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why are the lengths different?"

**Your Response:** "The output shows '4' then '5'. This demonstrates a subtle but important behavior of the `split()` method. By default, `split()` discards trailing empty strings, so 'a,b,c,d,' splits into only 4 parts, ignoring the empty string after the last comma. However, when we pass `-1` as the second parameter, it tells the method to keep all trailing empty strings, resulting in 5 parts. This behavior surprises many developers and is a common source of bugs in CSV parsing. The limit parameter controls how many splits occur and whether to keep trailing empties."

---

### 17. join()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String result = String.join(", ", "apple", "banana", "cherry");
        System.out.println(result);

        java.util.List<String> list = java.util.List.of("x", "y", "z");
        System.out.println(String.join("-", list));
    }
}
```
**A:**
```
apple, banana, cherry
x-y-z
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does the join method do?"

**Your Response:** "The output shows 'apple, banana, cherry' then 'x-y-z'. This demonstrates the `String.join()` method introduced in Java 8. It's a convenient utility for joining multiple strings with a delimiter. The first overload takes a delimiter and varargs of strings. The second overload takes a delimiter and an Iterable (like a List). This method is much cleaner than manually building strings with loops or StringBuilder. It's particularly useful for creating CSV strings, path construction, and any scenario where you need to concatenate strings with separators."

---

### 18. trim() vs strip()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "  hello  ";
        System.out.println("[" + s.trim() + "]");
        System.out.println("[" + s.strip() + "]");      // Java 11+ — Unicode-aware
        System.out.println("[" + s.stripLeading() + "]");
        System.out.println("[" + s.stripTrailing() + "]");
    }
}
```
**A:**
```
[hello]
[hello]
[hello  ]
[  hello]
```
`strip()` (Java 11+) handles Unicode whitespace. `trim()` only handles ASCII whitespace (≤ '\u0020').

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between trim and strip?"

**Your Response:** "The output shows '[hello]', '[hello]', '[hello  ]', then '[  hello]'. This demonstrates the difference between `trim()` and `strip()` methods. Both remove leading/trailing whitespace, but `strip()` (Java 11+) is Unicode-aware and handles all Unicode whitespace characters, while `trim()` only handles ASCII whitespace up to '\u0020'. `stripLeading()` and `stripTrailing()` are new methods that allow selective removal. For international applications, `strip()` is preferred as it handles various Unicode spaces, non-breaking spaces, and other whitespace characters properly."

---

### 19. startsWith() and endsWith()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String url = "https://www.example.com";
        System.out.println(url.startsWith("https"));
        System.out.println(url.endsWith(".com"));
        System.out.println(url.startsWith("www", 8)); // start checking at index 8
    }
}
```
**A:**
```
true
true
true
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these string prefix/suffix methods do?"

**Your Response:** "The output shows 'true', 'true', then 'true'. This demonstrates string prefix and suffix checking methods. `startsWith("https")` checks if the string begins with 'https'. `endsWith(".com")` checks if it ends with '.com'. `startsWith("www", 8)` starts checking from index 8, effectively checking if 'www' appears at that position. These methods are essential for URL validation, file extension checking, and protocol detection. They're more readable and efficient than using `substring()` or `indexOf()` for the same purpose. The offset parameter is useful for checking substrings at specific positions."

---

### 20. contains() and matches()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String email = "user@example.com";
        System.out.println(email.contains("@"));
        System.out.println(email.matches("[a-z]+@[a-z]+\\.[a-z]+")); // full match
        System.out.println("abc123".matches("[a-z]+")); // must match ENTIRE string
    }
}
```
**A:**
```
true
true
false
```
`matches()` requires the **entire string** to match the regex, unlike `find()` in Matcher which finds partial matches.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the key difference between contains and matches?"

**Your Response:** "The output shows 'true', 'true', then 'false'. This demonstrates different string matching approaches. `contains('@')` simply checks if the substring exists anywhere in the string. `matches()` requires the ENTIRE string to match the regex pattern - that's why the email validation passes but 'abc123'.matches('[a-z]+')' fails because the regex only allows lowercase letters, not digits. This is a common mistake - developers expect `matches()` to work like `contains()` but with regex. For partial regex matching, you'd need to use Pattern and Matcher with `find()`."

---

### 21. chars() Stream
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        long count = "Hello World".chars()
                .filter(Character::isUpperCase)
                .count();
        System.out.println(count);
    }
}
```
**A:** `2`. (`H` and `W`). `chars()` returns an `IntStream` of char values.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does the chars() method do?"

**Your Response:** "The output is '2'. This demonstrates the `chars()` method introduced in Java 8, which returns an `IntStream` of character values from the string. We filter this stream to count only uppercase characters using `Character::isUpperCase`. The result is 2 because 'Hello World' contains two uppercase letters: 'H' and 'W'. This method is incredibly useful for functional-style string processing, allowing you to use all the stream operations like filter, map, and reduce on characters. It's much cleaner than traditional for-loops for character-level analysis."

---

### 22. String.format() vs formatted() (Java 15+)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String name = "Alice"; int age = 30;
        System.out.println(String.format("Name: %s, Age: %d", name, age));
        System.out.println("Pi = %.4f".formatted(Math.PI)); // Java 15+
    }
}
```
**A:**
```
Name: Alice, Age: 30
Pi = 3.1416
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between these formatting approaches?"

**Your Response:** "The output shows formatted name/age string and Pi rounded to 4 decimal places. This demonstrates two string formatting approaches in Java. `String.format()` is the traditional static method that's been available since Java 5. The `formatted()` method was added in Java 15 as an instance method that makes the code more readable - you call it directly on the format string. Both use the same format specifiers like %s for strings, %d for integers, and %.4f for floating-point numbers with 4 decimal places. The `formatted()` method is more fluent and often preferred in newer Java code."

---

### 23. String to char Array and Back
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s = "hello";
        char[] chars = s.toCharArray();
        chars[0] = 'H'; // modify the char array
        System.out.println(s);             // original unchanged
        System.out.println(new String(chars)); // new string from chars
    }
}
```
**A:**
```
hello
Hello
```
`toCharArray()` returns a copy of the underlying char array. Modifying it doesn't affect the original String.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why doesn't changing the char array affect the original string?"

**Your Response:** "The output shows 'hello' then 'Hello'. This demonstrates String immutability and the `toCharArray()` method. When we call `toCharArray()`, it returns a copy of the string's internal character array, not the original array. When we modify the first character to 'H', we're only changing the copy. The original string remains unchanged because strings are immutable - their internal state cannot be modified after creation. This is why we need to create a new string with `new String(chars)` to see the modification. This behavior ensures thread safety and allows for internal optimizations like string sharing."

---

### 24. String.valueOf() for Primitives
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(String.valueOf(42));
        System.out.println(String.valueOf(3.14));
        System.out.println(String.valueOf(true));
        System.out.println(String.valueOf(new char[]{'h','i'}));
    }
}
```
**A:**
```
42
3.14
true
hi
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does String.valueOf() do for different types?"

**Your Response:** "The output shows '42', '3.14', 'true', then 'hi'. This demonstrates the versatility of `String.valueOf()` method, which has overloads for all primitive types and objects. It converts integers, doubles, booleans, and even char arrays to their string representations. For char arrays, it concatenates all characters. This method is null-safe - if you pass null, it returns the string 'null' instead of throwing an exception. It's the preferred way to convert primitives to strings when you need explicit control over the conversion process."

---

### 25. repeat() (Java 11+)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println("ab".repeat(3));
        System.out.println("-".repeat(10));
    }
}
```
**A:**
```
ababab
----------
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does the repeat() method do?"

**Your Response:** "The output shows 'ababab' then '----------'. This demonstrates the `repeat()` method introduced in Java 11, which repeats a string a specified number of times. 'ab'.repeat(3) creates 'ababab' and '-'.repeat(10) creates a line of 10 dashes. This method is much cleaner and more efficient than using loops or StringBuilder for simple string repetition. It's commonly used for creating separators, padding, or repeating patterns in text output. The method throws IllegalArgumentException if the count is negative, which is a sensible safety check."

---

### 26. lines() (Java 11+)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String text = "line1\nline2\nline3";
        text.lines().forEach(System.out::println);
        System.out.println(text.lines().count());
    }
}
```
**A:**
```
line1
line2
line3
3
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does the lines() method do?"

**Your Response:** "The output shows each line on a separate line followed by '3'. This demonstrates the `lines()` method introduced in Java 11, which returns a `Stream<String>` of lines from a multiline string. It automatically handles different line separators (\n, \r, \r\n) and provides a clean way to process multiline text. Each line is printed separately, and the count shows there are 3 lines. This method is much cleaner than manually splitting on line separators and handles edge cases properly. It's particularly useful for processing text files, configuration files, or any multiline input."

---

### 27. Text Blocks (Java 15+)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String json = """
                {
                    "name": "Alice",
                    "age": 30
                }
                """;
        System.out.println(json.strip());
    }
}
```
**A:**
```
{
    "name": "Alice",
    "age": 30
}
```
Text blocks (Java 15+) allow multiline strings without escape sequences.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what are text blocks?"

**Your Response:** "The output shows a nicely formatted JSON structure. This demonstrates text blocks, a major feature added in Java 15 for creating multiline strings. Using triple quotes (\"\"\"\"), you can create strings that span multiple lines without needing escape sequences for newlines or quotes. The `strip()` method removes the incidental leading whitespace. Text blocks make code much more readable, especially for JSON, SQL queries, HTML templates, or any multiline text. They preserve the formatting while eliminating the need for concatenation or escape sequences."

---

### 28. String Interning — When to Use
**Q: What is the performance consideration?**
```java
public class Main {
    public static void main(String[] args) {
        // Reading from a file might create millions of duplicate strings
        // Interning deduplicates them in the pool
        String s1 = new String("common").intern();
        String s2 = new String("common").intern();
        System.out.println(s1 == s2); // now the same pooled object
        // Useful for high-volume deduplication, but pool is limited in size
    }
}
```
**A:** `true`. `intern()` deduplicates strings but the pool is stored in the heap (Java 7+). Overusing it can cause memory pressure.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and when should you use string interning?"

**Your Response:** "The output is 'true'. This shows string interning for deduplication. When we call `intern()` on strings, Java ensures that equivalent strings share the same reference in the string pool. This is useful for memory optimization when dealing with millions of duplicate strings, like when processing large text files or database results. However, the string pool is stored in heap memory (since Java 7), so overusing `intern()` can cause memory pressure and GC issues. It's best used selectively for strings that are known to be duplicated frequently, not as a general optimization technique."

---

## Section 3: StringBuilder & StringBuffer (Q29–Q42)

### 29. StringBuilder — Mutable String
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder("Hello");
        sb.append(" World");
        sb.append("!");
        System.out.println(sb);
        System.out.println(sb.length());
    }
}
```
**A:**
```
Hello World!
12
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what makes StringBuilder different from String?"

**Your Response:** "The output shows 'Hello World!' and '12'. This demonstrates StringBuilder, which is mutable unlike String. When we call `append()` methods, they modify the internal buffer directly rather than creating new objects. The final string is 'Hello World!' with length 12. StringBuilder is essential for efficient string manipulation, especially in loops where string concatenation would create many temporary objects. It's the go-to choice when you need to build strings dynamically. The key difference is mutability - StringBuilder changes in place, String creates new objects."

---

### 30. StringBuilder vs Concatenation in Loop
**Q: Which is faster and why?**
```java
public class Main {
    public static void main(String[] args) {
        // SLOW: creates a new String object on each iteration
        String result = "";
        for (int i = 0; i < 5; i++) result += i;
        System.out.println(result);

        // FAST: O(n) amortized, single mutable buffer
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < 5; i++) sb.append(i);
        System.out.println(sb);
    }
}
```
**A:**
```
01234
01234
```
String concatenation in a loop is O(n²). `StringBuilder` is O(n). Always use `StringBuilder` in loops.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "Which is faster and why does performance differ so much?"

**Your Response:** "Both produce the same output '01234', but the performance difference is huge. String concatenation in a loop is O(n²) because each '+' creates a new String object and copies all previous characters. StringBuilder is O(n) because it maintains a mutable internal buffer and just appends to it. For 5 iterations, the difference is small, but for thousands of iterations, String concatenation becomes dramatically slower. This is a classic interview optimization question - always use StringBuilder in loops for string building. It's one of the most important performance optimizations in Java."

---

### 31. StringBuilder reverse()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder("hello");
        System.out.println(sb.reverse());
        // reverse is in-place — sb is now modified
        System.out.println(sb);
    }
}
```
**A:**
```
olleh
olleh
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why does the same StringBuilder print twice?"

**Your Response:** "The output shows 'olleh' twice. This demonstrates that StringBuilder's `reverse()` method modifies the StringBuilder in-place and also returns the same StringBuilder reference for chaining. When we call `sb.reverse()`, it reverses the internal buffer. The first `println` prints the reversed result. Since the StringBuilder remains reversed, the second `println` prints the same thing. This in-place modification is key to StringBuilder's efficiency - it doesn't create new objects. This behavior is different from String methods which always return new objects."

---

### 32. StringBuilder insert() and delete()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder("Hello World");
        sb.insert(5, ",");     // insert at index 5
        System.out.println(sb);
        sb.delete(0, 6);       // delete from 0 to 5 (exclusive 6)
        System.out.println(sb);
        sb.deleteCharAt(0);    // delete char at index 0
        System.out.println(sb);
    }
}
```
**A:**
```
Hello, World
World
orld
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do the insert and delete methods work?"

**Your Response:** "The output shows 'Hello, World', 'World', then 'orld'. This demonstrates StringBuilder's modification methods. `insert(5, ',')` inserts a comma at index 5, pushing the rest of the string to the right. `delete(0, 6)` removes characters from index 0 up to (but not including) index 6, effectively removing 'Hello, '. `deleteCharAt(0)` removes just the character at index 0. These methods modify the StringBuilder in-place, making them very efficient for string manipulation. The indices follow standard Java conventions - inclusive start, exclusive end for ranges."

---

### 33. StringBuilder indexOf and replace
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder("Hello World");
        int idx = sb.indexOf("World");
        sb.replace(idx, idx + 5, "Java");
        System.out.println(sb);
    }
}
```
**A:** `Hello Java`

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the replace method work?"

**Your Response:** "The output is 'Hello Java'. This demonstrates StringBuilder's `replace()` method, which works with indices rather than string patterns. First, we find the index of 'World' using `indexOf()`. Then `replace(idx, idx + 5, 'Java')` replaces characters from that start index to the end index (exclusive) with 'Java'. This is different from String's replace() which works with substrings or regex. StringBuilder's replace() is more efficient for large strings since it modifies the internal buffer directly. It's commonly used for template processing or when you need to replace specific regions."

---

### 34. StringBuffer vs StringBuilder
**Q: What is the key difference?**
```java
public class Main {
    public static void main(String[] args) {
        // StringBuffer — thread-safe (synchronized methods), slower
        StringBuffer sb1 = new StringBuffer("hello");
        sb1.append(" world");

        // StringBuilder — NOT thread-safe, faster (no synchronization)
        StringBuilder sb2 = new StringBuilder("hello");
        sb2.append(" world");

        System.out.println(sb1.equals(sb2)); // both extend AbstractStringBuilder, but different classes
    }
}
```
**A:** `false`. `StringBuffer.equals()` compares by reference (from Object), not content. Use `sb1.toString().equals(sb2.toString())` for content comparison. Prefer `StringBuilder` in single-threaded code.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the key difference and why does equals() return false?"

**Your Response:** "The output is 'false'. This demonstrates the key difference between StringBuffer and StringBuilder. StringBuffer is thread-safe with synchronized methods, making it slower. StringBuilder is not thread-safe but faster. The surprising part is that `equals()` returns false even though both contain 'hello world'. This is because neither class overrides `equals()` - they use Object's default reference comparison. To compare content, you must call `toString().equals()`. This is a classic gotcha that catches many developers off guard. In single-threaded code, always prefer StringBuilder for better performance."

---

### 35. StringJoiner (Java 8+)
**Q: What is the output?**
```java
import java.util.StringJoiner;
public class Main {
    public static void main(String[] args) {
        StringJoiner sj = new StringJoiner(", ", "[", "]");
        sj.add("apple");
        sj.add("banana");
        sj.add("cherry");
        System.out.println(sj);
        System.out.println(sj.length());
    }
}
```
**A:**
```
[apple, banana, cherry]
22
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does StringJoiner do?"

**Your Response:** "The output shows '[apple, banana, cherry]' and '22'. This demonstrates StringJoiner, introduced in Java 8 for building delimited strings. The constructor takes the delimiter and optional prefix/suffix. When we add elements, it automatically handles the delimiter placement. The length includes all characters including delimiters and brackets. StringJoiner is cleaner than manual StringBuilder concatenation for CSV-like output. It's particularly useful when you need consistent delimiter placement and optional prefix/suffix handling. It's the underlying implementation for `String.join()` method."

---

### 36. StringJoiner with Empty Check
**Q: What is the output?**
```java
import java.util.StringJoiner;
public class Main {
    public static void main(String[] args) {
        StringJoiner sj = new StringJoiner(", ", "[", "]");
        sj.setEmptyValue("EMPTY");
        System.out.println(sj); // no elements added

        sj.add("x");
        System.out.println(sj);
    }
}
```
**A:**
```
EMPTY
[x]
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what does setEmptyValue do?"

**Your Response:** "The output shows 'EMPTY' then '[x]'. This demonstrates StringJoiner's `setEmptyValue()` method, which provides a default value when no elements have been added. When the StringJoiner is empty, it returns the specified empty value instead of just the prefix/suffix. Once we add an element 'x', it behaves normally, showing '[x]'. This is useful for providing meaningful defaults in empty scenarios, like displaying 'No data available' instead of an empty list representation. It's a thoughtful API design that handles edge cases gracefully."

---

### 37. StringBuilder — Chaining Methods
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String result = new StringBuilder()
                .append("Hello")
                .append(", ")
                .append("World")
                .append("!")
                .reverse()
                .toString();
        System.out.println(result);
    }
}
```
**A:** `!dlroW ,olleH`. StringBuilder methods return `this` — enabling fluent chaining. Then `reverse()` is applied to the joined string.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does method chaining work?"

**Your Response:** "The output is '!dlroW ,olleH'. This demonstrates StringBuilder's fluent API design where most methods return 'this' for chaining. We start with an empty StringBuilder, append 'Hello', ', ', 'World', and '!' in sequence, then call `reverse()` on the entire result. The chaining makes the code readable and eliminates temporary variables. This pattern is common in builder classes and functional programming. The key insight is that each append modifies the same StringBuilder instance, and reverse() operates on the final accumulated string."

---

### 38. charAt() and setCharAt()
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder("Hello");
        System.out.println(sb.charAt(1));
        sb.setCharAt(0, 'h');
        System.out.println(sb);
    }
}
```
**A:**
```
e
hello
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between these methods?"

**Your Response:** "The output shows 'e' then 'hello'. This demonstrates StringBuilder's character access methods. `charAt(1)` returns the character at index 1 without modifying the StringBuilder, so it shows 'e'. `setCharAt(0, 'h')` modifies the character at index 0 directly in the internal buffer, changing 'H' to 'h'. The key difference is that `charAt()` is read-only while `setCharAt()` modifies the StringBuilder. These methods are useful for character-by-character manipulation when you need more control than append/insert provides."

---

### 39. StringBuilder capacity vs length
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder(); // default capacity 16
        System.out.println(sb.capacity());
        sb.append("hello");
        System.out.println(sb.length());
        System.out.println(sb.capacity()); // still 16 (not exceeded)
    }
}
```
**A:**
```
16
5
16
```
`length()` = actual characters. `capacity()` = internal buffer size. When capacity is exceeded, it doubles (+ 2).

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's the difference between length and capacity?"

**Your Response:** "The output shows '16', '5', then '16'. This demonstrates the difference between StringBuilder's `length()` and `capacity()`. `length()` returns the actual number of characters stored (5). `capacity()` returns the size of the internal buffer (16). When we create a StringBuilder without arguments, it starts with default capacity of 16. Since our string 'hello' is only 5 characters, the capacity remains 16. This pre-allocation strategy avoids frequent buffer resizing. When capacity is exceeded, it typically doubles to maintain O(n) performance. Understanding this helps optimize memory usage for large string building."

---

### 40. Palindrome Check with StringBuilder
**Q: What is the output?**
```java
public class Main {
    static boolean isPalindrome(String s) {
        String clean = s.replaceAll("[^a-zA-Z0-9]", "").toLowerCase();
        return clean.equals(new StringBuilder(clean).reverse().toString());
    }
    public static void main(String[] args) {
        System.out.println(isPalindrome("A man, a plan, a canal: Panama"));
        System.out.println(isPalindrome("race a car"));
    }
}
```
**A:**
```
true
false
```
Classic interview problem using StringBuilder's `reverse()`.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this palindrome check work?"

**Your Response:** "The output shows 'true' then 'false'. This demonstrates a classic palindrome checking algorithm using StringBuilder. The method first cleans the input by removing non-alphanumeric characters and converting to lowercase. Then it creates a reversed version using `new StringBuilder(clean).reverse().toString()`. Finally, it compares the cleaned string with its reversed version. The first string 'A man, a plan, a canal: Panama' is a palindrome when ignoring case and punctuation. The second string 'race a car' is not. This approach is efficient because StringBuilder's reverse() is optimized and works in linear time."

---

### 41. String.format With Various Formats
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.printf("%-10s | %5d | %.2f%n", "Alice", 30, 95.678);
        System.out.printf("%-10s | %5d | %.2f%n", "Bob", 25, 87.5);
    }
}
```
**A:**
```
Alice      |    30 | 95.68
Bob        |    25 | 87.50
```
`%-10s` = left-aligned, width 10. `%5d` = right-aligned, width 5. `%.2f` = 2 decimal places.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these format specifiers mean?"

**Your Response:** "The output shows formatted table with aligned columns. This demonstrates printf format specifiers. `%-10s` means string with width 10, left-aligned (the minus sign). `%5d` means integer with width 10, right-aligned (default for numbers). `%.2f` means floating-point with exactly 2 decimal places. The `|` character creates visual column separators. This formatting is essential for creating readable console output and reports. Understanding these specifiers is crucial for data presentation in command-line applications."

---

### 42. StringBuilder as Char Stack — Interview Pattern
**Q: What is the output?**
```java
public class Main {
    // Remove all adjacent duplicates: "abbaca" → "ca"
    static String removeDuplicates(String s) {
        StringBuilder sb = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (sb.length() > 0 && sb.charAt(sb.length() - 1) == c) {
                sb.deleteCharAt(sb.length() - 1); // pop
            } else {
                sb.append(c); // push
            }
        }
        return sb.toString();
    }
    public static void main(String[] args) {
        System.out.println(removeDuplicates("abbaca"));
        System.out.println(removeDuplicates("azxxzy"));
    }
}
```
**A:**
```
ca
ay
```
`StringBuilder` used as a stack for character processing — a common interview pattern.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this algorithm work?"

**Your Response:** "The output shows 'ca' then 'ay'. This demonstrates a clever use of StringBuilder as a stack to remove adjacent duplicates. The algorithm iterates through each character, checking if it matches the last character in the StringBuilder. If it matches, we remove (pop) the last character using `deleteCharAt()`. If it doesn't match, we append (push) the current character. For 'abbaca', this process removes 'bb' and 'aa' pairs, leaving 'ca'. For 'azxxzy', it removes 'xx' leaving 'ay'. This is a common interview pattern that tests understanding of stack-like behavior and string manipulation. It's O(n) time and O(n) space."

---

## Section 4: String Formatting & Parsing (Q43–Q52)

### 43. Integer.parseInt Edge Cases
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(Integer.parseInt("42"));
        System.out.println(Integer.parseInt("-100"));
        System.out.println(Integer.parseInt("FF", 16)); // parse hex
        System.out.println(Integer.parseInt("1010", 2)); // parse binary
        try { Integer.parseInt("abc"); }
        catch (NumberFormatException e) { System.out.println("NFE: " + e.getMessage()); }
    }
}
```
**A:**
```
42
-100
255
10
NFE: For input string: "abc"
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what edge cases does parseInt handle?"

**Your Response:** "The output shows '42', '-100', '255', '10', then an error message. This demonstrates `Integer.parseInt()` capabilities and error handling. It parses decimal strings, handles negative numbers, and supports different bases like hexadecimal (16) and binary (2). When parsing fails with invalid input like 'abc', it throws `NumberFormatException` with a descriptive message. This method is fundamental for converting string input to integers. Always wrap it in try-catch blocks for production code to handle invalid input gracefully."

---

### 44. Double.parseDouble and Number Formatting
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        double d = Double.parseDouble("3.14159");
        System.out.printf("%.2f%n", d);
        System.out.printf("%e%n", d);   // scientific notation
        System.out.printf("%10.3f%n", d); // width 10, 3 decimals
    }
}
```
**A:**
```
3.14
3.141590e+00
     3.142
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these format specifiers do?"

**Your Response:** "The output shows the same number formatted in different ways. This demonstrates floating-point formatting with `printf()`. `%.2f` formats to exactly 2 decimal places. `%e` displays in scientific notation with 6 decimal places by default. `%10.3f` creates width 10, right-aligned with 3 decimal places, padding with spaces. These format specifiers are essential for creating consistent numerical output. Understanding them helps in data reporting, scientific computing, and financial applications where precision formatting matters."

---

### 45. String to boolean
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(Boolean.parseBoolean("true"));
        System.out.println(Boolean.parseBoolean("TRUE"));
        System.out.println(Boolean.parseBoolean("yes"));  // false!
        System.out.println(Boolean.parseBoolean("1"));    // false!
    }
}
```
**A:**
```
true
true
false
false
```
Only "true" (case-insensitive) returns `true`. Everything else returns `false`.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what's surprising about boolean parsing?"

**Your Response:** "The output shows 'true', 'true', 'false', then 'false'. This demonstrates the very specific behavior of `Boolean.parseBoolean()`. Only the exact string 'true' (case-insensitive) returns true. Everything else, including 'TRUE', 'yes', '1', or 'false', returns false. This is surprisingly strict - many developers expect 'yes'/'no' or '1'/'0' to work. This strictness prevents ambiguity but requires careful input validation. For user interfaces, you often need to pre-process input to map common affirmative responses to 'true'."

---

### 46. String.format() with Padding
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(String.format("%05d", 42));   // zero-padded
        System.out.println(String.format("%+d", 42));    // always show sign
        System.out.println(String.format("%x", 255));    // hex lowercase
        System.out.println(String.format("%X", 255));    // hex uppercase
        System.out.println(String.format("%o", 8));      // octal
    }
}
```
**A:**
```
00042
+42
ff
FF
10
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and what do these format specifiers do?"

**Your Response:** "The output shows various number formatting options. This demonstrates `String.format()` with different format specifiers. `%05d` zero-pads to width 5, creating '00042'. `%+d` always shows the sign, creating '+42'. `%x` formats as lowercase hexadecimal, showing 'ff' for 255. `%X` formats as uppercase hexadecimal, showing 'FF'. `%o` formats as octal, showing '10' for decimal 8. These format specifiers are essential for creating consistent numeric output, especially in low-level programming, debugging, or when working with different number systems."

---

### 47. Regex — Pattern and Matcher
**Q: What is the output?**
```java
import java.util.regex.*;
public class Main {
    public static void main(String[] args) {
        String email = "user123@example.com";
        Pattern pattern = Pattern.compile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$");
        Matcher matcher = pattern.matcher(email);
        System.out.println(matcher.matches());

        String text = "Call 123-456-7890 or 987-654-3210";
        Pattern phones = Pattern.compile("\\d{3}-\\d{3}-\\d{4}");
        Matcher m = phones.matcher(text);
        while (m.find()) System.out.println(m.group());
    }
}
```
**A:**
```
true
123-456-7890
987-654-3210
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does regex matching work?"

**Your Response:** "The output shows 'true' then two phone numbers. This demonstrates Java's regex capabilities with Pattern and Matcher. First, we validate an email using `matcher.matches()` which checks if the entire string matches the pattern. The email regex validates the structure with alphanumeric characters, @ symbol, domain, and TLD. Then we extract phone numbers using a different pattern - `\d{3}-\d{3}-\d{4}` matches exactly 3-3-4 digit format. The `while(m.find())` loop finds all matches in the text. This is fundamental for input validation and data extraction."

---

### 48. Regex Groups
**Q: What is the output?**
```java
import java.util.regex.*;
public class Main {
    public static void main(String[] args) {
        Pattern p = Pattern.compile("(\\d{4})-(\\d{2})-(\\d{2})");
        Matcher m = p.matcher("Date: 2024-01-15");
        if (m.find()) {
            System.out.println("Year: " + m.group(1));
            System.out.println("Month: " + m.group(2));
            System.out.println("Day: " + m.group(3));
        }
    }
}
```
**A:**
```
Year: 2024
Month: 01
Day: 15
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how do regex groups work?"

**Your Response:** "The output shows extracted year, month, and day. This demonstrates regex groups for data extraction. The pattern `(\\d{4})-(\\d{2})-(\\d{2})` uses parentheses to create three capturing groups. `group(1)` returns the first group (year), `group(2)` the second (month), and `group(3)` the third (day). Groups are numbered starting from 1, with group(0) being the entire match. This is incredibly useful for parsing structured text like dates, log entries, or custom formats. Regex groups are more powerful than string splitting for complex patterns."

---

### 49. String.chars() to Count Characters
**Q: What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        String s = "hello world";
        Map<Character, Long> freq = s.chars()
                .filter(c -> c != ' ')
                .mapToObj(c -> (char)c)
                .collect(java.util.stream.Collectors.groupingBy(c -> c, Collectors.counting()));
        System.out.println(new TreeMap<>(freq));
    }
}
```
**A:** `{d=1, e=1, h=1, l=3, o=2, r=1, w=1}`

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this stream processing work?"

**Your Response:** "The output shows a character frequency map. This demonstrates functional string processing with streams. The `chars()` method returns an IntStream of character values. We filter out spaces, convert each int back to char, then collect using `groupingBy()` with `counting()` to build a frequency map. The `TreeMap` constructor sorts the keys alphabetically. This approach is much cleaner than manual counting with loops and maps. It showcases Java 8+ functional programming capabilities and is very efficient for text analysis, word frequency counting, or building histograms."

---

### 50. Scanner for Parsing
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner("10 3.14 hello");
        int i = sc.nextInt();
        double d = sc.nextDouble();
        String s = sc.next();
        System.out.println(i + " " + d + " " + s);
        sc.close();
    }
}
```
**A:** `10 3.14 hello`

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Scanner work?"

**Your Response:** "The output is '10 3.14 hello'. This demonstrates Scanner for parsing input strings. Scanner can read from strings, files, or console input. It automatically tokenizes based on whitespace and provides type-specific methods like `nextInt()`, `nextDouble()`, and `next()`. Each method reads the next token and converts it to the appropriate type. Scanner throws `InputMismatchException` if the token doesn't match the expected type. It's very convenient for parsing structured input or CSV-like data, though for complex parsing, regex or dedicated libraries might be better."

---

### 51. String.chars() — Anagram Check
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static boolean isAnagram(String s1, String s2) {
        char[] a = s1.toCharArray(), b = s2.toCharArray();
        Arrays.sort(a); Arrays.sort(b);
        return Arrays.equals(a, b);
    }
    public static void main(String[] args) {
        System.out.println(isAnagram("listen", "silent"));
        System.out.println(isAnagram("hello", "world"));
    }
}
```
**A:**
```
true
false
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this numeric check work?"

**Your Response:** "The output shows 'true' then 'false'. This demonstrates checking if a string contains only numeric characters. The method first handles edge cases - null or empty strings return false. Then it iterates through each character using `toCharArray()` and checks if it's a digit using `Character.isDigit()`. If any character is not a digit, it immediately returns false. Only if all characters pass the digit test does it return true. This is a common utility function used in input validation. The approach is straightforward and efficient, running in O(n) time where n is the string length."

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this anagram check work?"

**Your Response:** "The output shows 'true' then 'false'. This demonstrates an anagram checking algorithm. The method converts both strings to character arrays, sorts them using `Arrays.sort()`, then compares with `Arrays.equals()`. If they're anagrams, the sorted arrays will be identical. 'listen' and 'silent' are anagrams, so they return true. 'hello' and 'world' are not, so they return false. This approach is efficient with O(n log n) time complexity due to sorting. It's a classic interview problem that tests understanding of arrays, sorting, and string manipulation."

---

### 52. toUpperCase / toLowerCase with Locale
**Q: What is the output?**
```java
import java.util.Locale;
public class Main {
    public static void main(String[] args) {
        String s = "istanbul";
        // Turkish locale: dotted-i issue
        System.out.println(s.toUpperCase(Locale.ENGLISH));
        System.out.println(s.toUpperCase(Locale.forLanguageTag("tr")));
    }
}
```
**A:**
```
ISTANBUL
İSTANBUL
```
Turkish locale has unique uppercase rules for 'i'. Always specify `Locale.ENGLISH` for programmatic string processing to avoid locale surprises.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and why do the results differ?"

**Your Response:** "The output shows 'ISTANBUL' then 'İSTANBUL'. This demonstrates the critical importance of locale in string case conversion. In Turkish, the lowercase 'i' converts to a dotted uppercase 'İ', not the regular 'I'. This can cause bugs in string comparisons, hashing, or database lookups. The English locale produces the expected 'ISTANBUL'. This is why you should always specify `Locale.ENGLISH` for programmatic string processing, especially for passwords, usernames, or identifiers where consistency matters. It's a classic internationalization gotcha."

---

## Section 5: String Common Interview Patterns (Q53–Q60)

### 53. Find First Non-Repeating Character
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static char firstUnique(String s) {
        Map<Character, Integer> freq = new LinkedHashMap<>();
        for (char c : s.toCharArray()) freq.merge(c, 1, Integer::sum);
        for (Map.Entry<Character, Integer> e : freq.entrySet())
            if (e.getValue() == 1) return e.getKey();
        return '_';
    }
    public static void main(String[] args) {
        System.out.println(firstUnique("leetcode"));
        System.out.println(firstUnique("aabb"));
    }
}
```
**A:**
```
l
_
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this algorithm work?"

**Your Response:** "The output shows 'l' then '_'. This demonstrates finding the first non-repeating character using a LinkedHashMap. The algorithm first counts character frequencies using `merge()` for conciseness. Since LinkedHashMap preserves insertion order, the second loop finds the first character with count 1. For 'leetcode', 'l' appears once and comes first, so it's returned. For 'aabb', all characters repeat, so we return '_' as a sentinel value. This is O(n) time and O(k) space where k is the number of unique characters. It's a classic interview problem that tests hash map usage and order preservation."

---

### 54. Reverse Words in a String
**Q: What is the output?**
```java
public class Main {
    static String reverseWords(String s) {
        String[] words = s.trim().split("\\s+");
        StringBuilder sb = new StringBuilder();
        for (int i = words.length - 1; i >= 0; i--) {
            sb.append(words[i]);
            if (i > 0) sb.append(" ");
        }
        return sb.toString();
    }
    public static void main(String[] args) {
        System.out.println(reverseWords("the sky is blue"));
        System.out.println(reverseWords("  hello world  "));
    }
}
```
**A:**
```
blue is sky the
world hello
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this reversal work?"

**Your Response:** "The output shows 'blue is sky the' then 'world hello'. This demonstrates reversing words in a string while handling extra spaces. The algorithm first trims leading/trailing spaces, then splits on one or more whitespace characters using `\\s+`. This handles multiple spaces between words. Then it iterates backwards through the words array, building the result with a StringBuilder. The key insight is processing from end to start and adding spaces between words. This is a common interview problem that tests string manipulation, array handling, and edge cases like extra spaces."

---

### 55. Count Vowels
**Q: What is the output?**
```java
import java.util.Set;
public class Main {
    static long countVowels(String s) {
        Set<Character> vowels = Set.of('a','e','i','o','u','A','E','I','O','U');
        return s.chars().filter(c -> vowels.contains((char)c)).count();
    }
    public static void main(String[] args) {
        System.out.println(countVowels("Hello World"));
        System.out.println(countVowels("aeiou"));
    }
}
```
**A:**
```
3
5
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this counting work?"

**Your Response:** "The output shows '3' then '5'. This demonstrates counting vowels using a functional approach. We create a Set of vowel characters (both cases) for O(1) lookup. The `chars()` method returns an IntStream of characters, which we filter to keep only vowels, then count them. For 'Hello World', the vowels are e, o, o (3 total). For 'aeiou', all 5 characters are vowels. This approach is concise and leverages Java 8+ streams. It's much cleaner than manual iteration and character-by-character checking."

---

### 56. Longest Common Prefix
**Q: What is the output?**
```java
public class Main {
    static String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) return "";
        String prefix = strs[0];
        for (String s : strs) {
            while (!s.startsWith(prefix)) prefix = prefix.substring(0, prefix.length() - 1);
        }
        return prefix;
    }
    public static void main(String[] args) {
        System.out.println(longestCommonPrefix(new String[]{"flower","flow","flight"}));
        System.out.println(longestCommonPrefix(new String[]{"dog","racecar","car"}));
    }
}
```
**A:**
```
fl

```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this algorithm work?"

**Your Response:** "The output shows 'fl' then an empty line. This demonstrates finding the longest common prefix algorithm. We start with the first string as the initial prefix, then iteratively shrink it until all strings start with it. For each string, we use `startsWith()` to check if it begins with the current prefix. If not, we remove the last character using `substring(0, length - 1)`. For 'flower', 'flow', and 'flight', the common prefix is 'fl'. For 'dog', 'racecar', and 'car', there's no common prefix, so it becomes empty. This is O(n*m) time where n is number of strings and m is average length."

---

### 57. String Compression
**Q: What is the output?**
```java
public class Main {
    static String compress(String s) {
        StringBuilder sb = new StringBuilder();
        int i = 0;
        while (i < s.length()) {
            char c = s.charAt(i);
            int count = 0;
            while (i < s.length() && s.charAt(i) == c) { i++; count++; }
            sb.append(c).append(count);
        }
        return sb.length() < s.length() ? sb.toString() : s;
    }
    public static void main(String[] args) {
        System.out.println(compress("aabcccccaaa"));
        System.out.println(compress("abc")); // no compression benefit
    }
}
```
**A:**
```
a2b1c5a3
abc
```

---

### How to Explain in Interview (Spoken style format)

**Your Response:** "The output shows 'a2b1c5a3' then 'abc'. This demonstrates string compression with run-length encoding. The algorithm iterates through the string, counting consecutive identical characters. When the character changes or we reach the end, we append the character and its count. For 'aabcccccaaa', this produces 'a2b1c5a3'. For 'abc', there are no consecutive duplicates, so the compressed version would be longer than the original, so we return the original string. This is a classic interview problem that tests string processing, counting logic, and space-time tradeoff analysis."

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does this compression algorithm work?"

**Your Response:** "The output shows 'a2b1c5a3' then 'abc'. This demonstrates a simple string compression algorithm. The `compress()` method counts consecutive identical characters and replaces them with the character followed by the count. For 'aabcccccaaa', it compresses to 'a2b1c5a3'. For 'abc', there are no consecutive duplicates, so it returns unchanged. This is a classic interview problem that tests string manipulation and algorithmic thinking. The algorithm uses StringBuilder for efficiency since strings are immutable."

---

### 58. Check if String is Numeric
**Q: What is the output?**
```java
public class Main {
    static boolean isNumeric(String s) {
        if (s == null || s.isEmpty()) return false;
        for (char c : s.toCharArray()) if (!Character.isDigit(c)) return false;
        return true;
    }
    public static void main(String[] args) {
        System.out.println(isNumeric("12345"));
        System.out.println(isNumeric("123.45"));
        System.out.println(isNumeric("123abc"));
    }
}
```
**A:**
```
true
false
false
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the numeric check work?"

**Your Response:** "The output is 'true', 'false', then 'false'. This shows how to check if a string contains only numeric characters. The `isNumeric()` method first handles edge cases - null or empty strings return false. Then it iterates through each character using `toCharArray()` and checks if each one is a digit using `Character.isDigit()`. If any character is not a digit, it immediately returns false. Only if all characters pass the digit test does it return true. This is a common utility function used in input validation."

---

### 59. String Reversal

---

### 59. Roman to Integer (Common Interview Q)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static int romanToInt(String s) {
        Map<Character, Integer> map = Map.of('I',1,'V',5,'X',10,'L',50,'C',100,'D',500,'M',1000);
        int result = 0;
        for (int i = 0; i < s.length(); i++) {
            int curr = map.get(s.charAt(i));
            int next = i + 1 < s.length() ? map.get(s.charAt(i+1)) : 0;
            result += (curr < next) ? -curr : curr;
        }
        return result;
    }
    public static void main(String[] args) {
        System.out.println(romanToInt("III"));
        System.out.println(romanToInt("IV"));
        System.out.println(romanToInt("LVIII"));
        System.out.println(romanToInt("MCMXCIV"));
    }
}
```
**A:**
```
3
4
58
1994
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does Roman numeral conversion work?"

**Your Response:** "The output shows '3', '4', '58', then '1994'. This demonstrates Roman numeral to integer conversion. The algorithm uses a Map for symbol values. The key insight is handling subtractive notation - when a smaller value appears before a larger one, we subtract it (like IV = 5-1 = 4). For each character, we look ahead to the next character. If the current value is less than the next, we subtract; otherwise we add. This handles all Roman numeral rules including subtractive pairs like IV, IX, XL, XC, CM. It's a classic interview problem that tests algorithmic thinking and handling of special cases."

---

### 60. Valid Parentheses (Stack + String)
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static boolean isValid(String s) {
        Deque<Character> stack = new ArrayDeque<>();
        for (char c : s.toCharArray()) {
            if (c == '(' || c == '[' || c == '{') { stack.push(c); }
            else if (stack.isEmpty()) return false;
            else if (c == ')' && stack.pop() != '(') return false;
            else if (c == ']' && stack.pop() != '[') return false;
            else if (c == '}' && stack.pop() != '{') return false;
        }
        return stack.isEmpty();
    }
    public static void main(String[] args) {
        System.out.println(isValid("()[]{}"));
        System.out.println(isValid("(]"));
        System.out.println(isValid("{[]}"));
    }
}
```
**A:**
```
true
false
true
```

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** "What's the output and how does the stack-based validation work?"

**Your Response:** "The output shows 'true', 'false', then 'true'. This demonstrates the classic valid parentheses algorithm using a stack. We push opening brackets onto the stack. When we encounter a closing bracket, we check if the stack is empty (invalid) or if the top doesn't match the corresponding opening bracket. If either condition fails, we return false. If they match, we pop the opening bracket. At the end, the string is valid only if the stack is empty (all brackets matched). This algorithm works for any type of brackets and is O(n) time and O(n) space. It's a fundamental computer science problem that tests understanding of stacks and string parsing."
