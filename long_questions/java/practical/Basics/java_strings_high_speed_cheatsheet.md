# Java Strings & StringBuilder — High-Speed Retrieval Cheatsheet

> **Interview-Ready Framework**: 5-Step Mental Model for Instant Recall

---

## 🔧 The 'Underlying Engine' Map

### **Category 1: String Fundamentals & Memory**
**Logic Pattern**: String Creation → Memory Management → Optimization
- **Immutability**: Fixed state → Thread safety, optimization
- **String Pool**: Literal reuse → Memory efficiency
- **Object Creation**: `new String()` vs literals → Heap vs pool
- **intern()**: Pool management → Deduplication

### **Category 2: String Operations & Methods**
**Logic Pattern**: String Input → Transformation → Output
- **Searching**: `indexOf()`, `lastIndexOf()`, `contains()` → Position finding
- **Extraction**: `substring()`, `split()`, `join()` → Content manipulation
- **Modification**: `replace()`, `replaceAll()`, `trim()`, `strip()` → Content change
- **Validation**: `startsWith()`, `endsWith()`, `matches()` → Pattern checking

### **Category 3: Mutable String Builders**
**Logic Pattern**: Buffer Creation → In-place Modification → Result Extraction
- **StringBuilder**: Single-threaded mutable → Fast string building
- **StringBuffer**: Thread-safe mutable → Synchronized operations
- **StringJoiner**: Delimited building → CSV-like output
- **Method Chaining**: Fluent API → Readable operations

### **Category 4: Advanced String Processing**
**Logic Pattern**: Complex Input → Algorithm → Result
- **Parsing**: `parseInt()`, `parseDouble()`, `Scanner` → Type conversion
- **Formatting**: `String.format()`, `formatted()` → Output presentation
- **Regex**: `Pattern`, `Matcher` → Pattern matching
- **Functional**: `chars()`, `lines()`, `repeat()` → Stream processing

---

## 🚨 The 'Red-Flag' Failure Section

### **Critical Runtime Errors**

| **Error Type** | **Trigger** | **Example** | **Fix** |
|---|---|---|---|
| **NullPointerException** | Method on null string | `null.length()` | Check null first |
| **NumberFormatException** | Invalid numeric string | `Integer.parseInt("abc")` | Try-catch block |
| **StringIndexOutOfBoundsException** | Invalid indices | `substring(10, 5)` | Validate bounds |
| **PatternSyntaxException** | Invalid regex pattern | `Pattern.compile("[")` | Test regex |
| **IllegalStateException** | Empty Optional | `optional.get()` | Use `orElse()` |

### **Logic Failure Patterns**

- **String vs Object Comparison**: Using `==` instead of `equals()` → Wrong results
- **Immutable Misunderstanding**: Expecting `toUpperCase()` to modify original → No change
- **Concatenation Performance**: String `+` in loops → O(n²) performance
- **Split Trailing Empties**: `split()` discarding empty strings → Data loss
- **Replace vs ReplaceAll**: Using regex when literal needed → Unexpected behavior
- **Locale Issues**: Case conversion with different locales → Inconsistent results

### **Performance Killers**

- **Loop Concatenation**: `result += part` → Creates many objects
- **Substring Memory**: Pre-Java 7 substring sharing → Memory leaks
- **Regex Overuse**: `replaceAll()` for simple cases → Performance overhead
- **String Creation**: Unnecessary `new String()` → Object waste

---

## ⚡ The 'Performance & Complexity' Table

| **Operation** | **Time Complexity** | **Memory Usage** | **Best Use Case** |
|---|---|---|---|
| **String concatenation** | O(n²) in loops | O(n²) objects | Simple one-time ops |
| **StringBuilder append** | O(1) amortized | O(n) buffer | Loop string building |
| **substring()** | O(1) | O(1) reference | Extracting portions |
| **indexOf()** | O(n) | O(1) | Finding characters |
| **split()** | O(n) | O(n) array | Tokenizing strings |
| **replaceAll()** | O(n) | O(n) | Pattern replacement |
| **toLowerCase()** | O(n) | O(n) new string | Case conversion |
| **trim()** | O(n) | O(n) new string | Whitespace removal |

### **Performance Rules**
- **Loop Concatenation**: Always use StringBuilder → O(n) vs O(n²)
- **StringBuilder vs StringBuffer**: Use StringBuilder unless thread-safe → 2-3x faster
- **String Pool**: Literals reuse memory → Use for constants
- **Regex Compilation**: Compile once, reuse → Avoid Pattern.compile() in loops
- **charAt() vs substring()**: Use charAt() for single characters → No object creation

---

## 🛡️ The 'Safe vs. Risky' Comparison

### **Standard/Safe Methods**
```java
// ✅ SAFE: Modern, recommended approaches
StringBuilder sb = new StringBuilder();
sb.append("text").append(" more");
String result = sb.toString();

String.join(", ", items); // Java 8+
"hello".strip(); // Java 11+ Unicode-aware
"pattern".matches("regex"); // Full string match
```

### **Legacy/Dangerous Methods**
```java
// ❌ RISKY: Avoid in modern code
String result = "";
for (String s : list) result += s; // O(n²) performance

"hello".replaceAll(".", "!"); // Uses regex, not literal
new String("literal"); // Unnecessary object creation
trim() vs strip() // ASCII-only vs Unicode-aware
```

### **Why Use X Over Y?**

| **Safe Choice** | **Risky Alternative** | **Reason** |
|---|---|---|
| `StringBuilder` | String `+` in loops | O(n) vs O(n²) performance |
| `String.join()` | Manual concatenation | Cleaner, handles delimiters |
| `strip()` | `trim()` | Unicode-aware vs ASCII-only |
| `replace()` | `replaceAll()` | Literal vs regex |
| `equals()` | `==` | Content vs reference comparison |
| `String.valueOf()` | `toString()` on null | Null-safe vs NPE |

---

## 🎯 The 'Interview Logic' Column

### **Core Concepts with Analogies & Golden Rules**

| **Concept** | **Real-World Analogy** | **Golden Rule** |
|---|---|---|
| **String Immutability** | **Carved Stone**: Once written, cannot be changed | "String methods never modify - they always return new objects" |
| **String Pool** | **Library**: Same book title, one copy shared | "Literals reuse memory - new String() creates new objects" |
| **StringBuilder** | **Whiteboard**: Write, erase, rewrite freely | "Use StringBuilder in loops - it's mutable and efficient" |
| **Regex Patterns** | **Search Pattern**: Find matching text fragments | "Use replace() for literals, replaceAll() only for regex" |
| **String Split** | **Scissors**: Cut text into pieces | "split() discards trailing empties - use limit -1 to keep" |
| **Case Conversion** | **Language Rules**: Different alphabets, different rules | "Always specify Locale.ENGLISH for consistent results" |
| **String Formatting** | **Template**: Fill in the blanks consistently | "Use printf/format for aligned output - specify width and precision" |

### **Quick Interview Decision Tree**

1. **Need to build strings in loop?** → `StringBuilder` (never String `+`)
2. **Need to compare content?** → `equals()` (never `==`)
3. **Need literal replacement?** → `replace()` (not `replaceAll()`)
4. **Need to split by delimiter?** → `split()` (check trailing empties)
5. **Need to handle whitespace?** → `strip()` (not `trim()` for Unicode)
6. **Need to validate input?** → `matches()` or `Pattern/Matcher`

---

## 📚 Mental Index Cards for Rapid Recall

### **Card 1: String Fundamentals**
```
// Immutability
String s = "hello";
s.toUpperCase(); // s unchanged
String upper = s.toUpperCase(); // new object

// String Pool
String a = "hello";
String b = "hello"; // same object
String c = new String("hello"); // different object
```

### **Card 2: StringBuilder Pattern**
```
StringBuilder sb = new StringBuilder();
for (String part : parts) {
    sb.append(part).append(" ");
}
String result = sb.toString().trim();
```

### **Card 3: Common Operations**
```
// Search
int index = str.indexOf("pattern");
boolean found = str.contains("pattern");

// Extract
String sub = str.substring(start, end);
String[] parts = str.split(",");
```

### **Card 4: Validation & Formatting**
```
// Validation
boolean valid = str.matches("[a-z]+");
boolean starts = str.startsWith("prefix");

// Formatting
String formatted = String.format("%-10s | %5d", name, count);
```

---

## 🔥 Top 10 Interview Patterns

1. **String Reversal**: `new StringBuilder(s).reverse().toString()`
2. **Palindrome Check**: `s.equals(new StringBuilder(s).reverse().toString())`
3. **First Unique Character**: `LinkedHashMap` frequency counting
4. **Word Reversal**: `split("\\s+")` + reverse iteration
5. **String Compression**: Run-length encoding with StringBuilder
6. **Roman to Integer**: Map lookup with subtractive logic
7. **Valid Parentheses**: Stack-based bracket matching
8. **Anagram Check**: Sort character arrays and compare
9. **Longest Common Prefix**: Iterative prefix shrinking
10. **Character Frequency**: `chars().collect(groupingBy(), counting())`

---

## 🚨 Critical Interview Red Flags to Avoid

### **Never Say These in Interviews**
- ❌ "I use `==` to compare strings"
- ❌ "String concatenation with `+` is efficient in loops"
- ❌ "I use `replaceAll()` for simple replacements"
- ❌ "String and StringBuilder are the same"
- ❌ "I don't worry about null strings"

### **Always Mention These**
- ✅ String immutability and thread safety
- ✅ String pool memory optimization
- ✅ StringBuilder vs StringBuffer performance
- ✅ Null-safe string handling with `String.valueOf()`
- ✅ Locale-aware case conversion

---

## 🎯 Quick Reference: Method Selection Guide

| **Task** | **Best Method** | **Alternative** | **When to Use** |
|---|---|---|---|
| **Build string in loop** | `StringBuilder.append()` | `String.concat()` | Always use StringBuilder |
| **Compare content** | `equals()` | `contentEquals()` | Content comparison |
| **Check prefix/suffix** | `startsWith()/endsWith()` | `substring().equals()` | Prefix/suffix checking |
| **Remove whitespace** | `strip()` (Java 11+) | `trim()` | Unicode text |
| **Join strings** | `String.join()` | `StringBuilder` | Delimited output |
| **Split string** | `split()` | `StringTokenizer` | Tokenization |
| **Replace literal** | `replace()` | `replaceAll()` | Simple replacement |
| **Case conversion** | `toLowerCase(Locale.ENGLISH)` | `toLowerCase()` | Consistent results |

---

**Interview Strategy**: Start with the immutability analogy, identify if you need mutable (StringBuilder) or immutable operations, choose the right method based on the task (search, extract, replace), and always consider performance implications and null safety.
