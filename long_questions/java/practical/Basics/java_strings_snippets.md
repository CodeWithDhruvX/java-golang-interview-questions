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
