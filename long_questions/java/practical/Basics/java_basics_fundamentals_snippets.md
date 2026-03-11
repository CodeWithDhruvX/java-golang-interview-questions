# 100 Pure Code Snippet Interview Questions: Java Basics & Fundamentals

*Each question is a "predict the output / spot the bug / does it compile?" style question.*
*Topics: Variables, Primitives, Type Casting, Control Flow, Arrays, Strings, OOP Basics, Static, Final, Wrapper Classes, Autoboxing, Varargs.*

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** Primitives, Variables & Type Casting (Q1–Q15)
- [ ] **Section 2:** Control Flow (Q16–Q28)
- [ ] **Section 3:** Arrays (Q29–Q38)
- [ ] **Section 4:** Static, Final & Initialization (Q39–Q52)
- [ ] **Section 5:** Wrapper Classes & Autoboxing (Q53–Q65)
- [ ] **Section 6:** Varargs & Overloading (Q66–Q75)
- [ ] **Section 7:** Basic OOP & References (Q76–Q90)
- [ ] **Section 8:** Miscellaneous Gotchas (Q91–Q100)

> 🔖 **Last read:** <!-- e.g. Q15 · Section 1 done -->

---

## Section 1: Primitives, Variables & Type Casting (Q1–Q15)

### 1. Integer Division Truncation
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int a = 7;
        int b = 2;
        System.out.println(a / b);
        System.out.println(a % b);
    }
}
```
**A:**
```
3
1
```
Integer division truncates toward zero. `7/2 = 3`, remainder `7%2 = 1`.

**Code Snippet Internal Behavior:**
- JVM performs integer division using the `idiv` bytecode instruction
- Division truncates toward zero (floor division for positive numbers)
- Remainder calculated using `irem` bytecode instruction
- Both operations work on 32-bit signed integers
- Results are stored in the operand stack before being passed to `System.out.println`

---

### 2. Implicit Widening
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        int i = 100;
        long l = i;     // widening: int → long
        double d = l;   // widening: long → double
        System.out.println(d);
    }
}
```
**A:** **Yes, compiles and prints** `100.0`. Java automatically widens smaller numeric types to larger ones without a cast.

**Code Snippet Internal Behavior:**
- JVM uses `i2l` bytecode to convert int (32-bit) to long (64-bit)
- Then uses `l2d` bytecode to convert long (64-bit) to double (64-bit IEEE 754)
- Widening conversions preserve the numeric value exactly
- No precision loss occurs in this conversion chain
- The double value `100.0` is stored in IEEE 754 format

---

### 3. Narrowing Cast — Data Loss
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        double d = 9.99;
        int i = (int) d;
        System.out.println(i);
    }
}
```
**A:** `9`. Narrowing cast truncates (does NOT round) the decimal part.

**Code Snippet Internal Behavior:**
- JVM uses `d2i` bytecode instruction for double to int conversion
- Conversion drops fractional part (truncates toward zero)
- IEEE 754 double `9.99` becomes integer `9`
- No rounding algorithm is applied - simple truncation
- The result is stored as 32-bit signed integer

---

### 4. Byte Overflow
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        byte b = 127;
        b++;
        System.out.println(b);
    }
}
```
**A:** `-128`. A `byte` holds values -128 to 127. Incrementing past 127 wraps around to -128.

**Code Snippet Internal Behavior:**
- `b++` uses `iinc` bytecode on local variable (treated as int internally)
- Value 127 becomes 128 in int arithmetic
- `i2b` bytecode converts back to byte with signed narrowing
- Binary 10000000 (128) interpreted as signed byte = -128 (two's complement)
- Overflow is silent - no exception thrown

---

### 5. char Arithmetic
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        char c = 'A';
        c += 1;
        System.out.println(c);
        System.out.println((int) c);
    }
}
```
**A:**
```
B
66
```
`char` is an unsigned 16-bit integer. `'A'` = 65, adding 1 gives 66 = `'B'`.

**Code Snippet Internal Behavior:**
- `char` stored as 16-bit unsigned integer (UTF-16 code unit)
- `'A'` = Unicode code point 65 = 0x0041
- `c += 1` uses `iadd` bytecode (char promoted to int)
- Result 66 converted back to char with `i2c` bytecode
- `System.out.println(char)` prints Unicode character
- `(int) c` uses `i2s` bytecode to show numeric value

---

### 6. Integer Literal Overflow
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        int x = 2147483648; // Integer.MAX_VALUE + 1
    }
}
```
**A:** **Compile Error.** `2147483648` exceeds `int` range. Use `long x = 2147483648L;` or cast.

**Code Snippet Internal Behavior:**
- Compiler parses integer literals as 32-bit signed int by default
- Value 2147483648 exceeds Integer.MAX_VALUE (2147483647)
- Compiler detects overflow at compile time (not runtime)
- `L` suffix tells compiler to parse as 64-bit long literal
- No bytecode generated - compilation fails

---

### 7. Short-Circuit Evaluation
**Q: What is the output?**
```java
public class Main {
    static int count = 0;
    static boolean check() { count++; return true; }

    public static void main(String[] args) {
        boolean result = false && check();
        System.out.println(result);
        System.out.println(count);
    }
}
```
**A:**
```
false
0
```
`&&` short-circuits — if the left side is `false`, the right side is never evaluated. `check()` is never called.

**Code Snippet Internal Behavior:**
- JVM uses conditional branch instructions for short-circuit evaluation
- `false && check()` compiles to jump directly to false result
- `check()` method invocation bytecode never generated for right operand
- Static variable `count` remains at initial value 0
- Operand stack contains boolean false without executing check()

---

### 8. Pre vs Post Increment
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int a = 5;
        int b = a++;
        int c = ++a;
        System.out.println(a + " " + b + " " + c);
    }
}
```
**A:** `7 5 7`. `a++` returns 5 then increments (a=6). `++a` increments first (a=7) then returns 7.

**Code Snippet Internal Behavior:**
- `a++` uses `iinc` bytecode after storing original value on stack
- Post-increment: push original value (5), then increment local variable
- `++a` uses `iinc` bytecode before pushing value to stack
- Pre-increment: increment local variable first, then push new value (7)
- String concatenation uses `StringBuilder` internally at compile time

---

### 9. String Concatenation with + Operator
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(1 + 2 + "3");
        System.out.println("1" + 2 + 3);
    }
}
```
**A:**
```
33
123
```
`+` is left-associative. `1 + 2` = `3` (int), then `3 + "3"` = `"33"`. In the second line `"1" + 2` = `"12"`, then `"12" + 3` = `"123"`.

**Code Snippet Internal Behavior:**
- First line: `iadd` bytecode for 1+2=3, then `StringBuilder.append()` for concatenation
- Second line: `StringBuilder` created immediately when String encountered
- Compiler converts `+` chain to `StringBuilder` operations
- Mixed types trigger appropriate `String.valueOf()` calls
- Final result via `StringBuilder.toString()`

---

### 10. Comparing Doubles
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        double a = 0.1 + 0.2;
        System.out.println(a == 0.3);
        System.out.println(a);
    }
}
```
**A:**
```
false
0.30000000000000004
```
Floating-point arithmetic is not exact. Never use `==` to compare doubles; use `Math.abs(a - 0.3) < 1e-9`.

**Code Snippet Internal Behavior:**
- 0.1 and 0.2 cannot be represented exactly in IEEE 754 binary format
- `dadd` bytecode performs binary floating-point addition
- Result is 0.30000000000000004 (closest representable binary value)
- `dcmpg` bytecode for comparison finds values unequal
- Binary floating-point precision limitations cause the discrepancy

---

### 11. Integer Cache (== vs equals)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Integer a = 127;
        Integer b = 127;
        Integer c = 128;
        Integer d = 128;
        System.out.println(a == b);
        System.out.println(c == d);
    }
}
```
**A:**
```
true
false
```
Java caches `Integer` objects for values -128 to 127. `a` and `b` reference the same cached object. `c` and `d` are different heap objects. Always use `.equals()` for object comparison.

**Code Snippet Internal Behavior:**
- `Integer.valueOf()` uses internal cache for -128 to 127
- Cache implemented as static array of Integer objects
- 127 hits cache → same object reference → `==` evaluates true
- 128 exceeds cache → new objects created → different references
- `==` compares object references, `.equals()` compares int values

---

### 12. Variable Shadowing
**Q: What is the output?**
```java
public class Main {
    static int x = 10;

    public static void main(String[] args) {
        int x = 20;
        System.out.println(x);
        System.out.println(Main.x);
    }
}
```
**A:**
```
20
10
```
The local `x` shadows the static field `x`. Use `Main.x` (or `this.x` in instance methods) to access the shadowed field.

**Code Snippet Internal Behavior:**
- Local variables stored in method's local variable array (higher precedence)
- Static fields stored in class's method area memory
- Compiler resolves variable names using lexical scoping rules
- `getstatic` bytecode used for `Main.x` access
- `iload` bytecode used for local variable access

---

### 13. Final Variable Must Be Initialized
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        final int x;
        System.out.println(x);
    }
}
```
**A:** **Compile Error.** A `final` local variable must be definitely assigned before use (blank final). `x` is declared but never assigned.

**Code Snippet Internal Behavior:**
- Compiler performs definite assignment analysis (flow analysis)
- Tracks all possible code paths to ensure final variables are assigned
- Blank final variables have special bytecode verification rules
- `final` modifier stored in method's local variable table
- JVM verifies initialization before any read access

---

### 14. Long Literal Suffix
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        long l1 = 10000000000;   // 10 billion
        long l2 = 10000000000L;  // 10 billion with L suffix
        System.out.println(l2);
    }
}
```
**A:** **Compile Error on l1.** Without the `L` suffix, `10000000000` is treated as an `int` literal, which overflows. `l2` compiles fine.

**Code Snippet Internal Behavior:**
- Integer literals parsed as 32-bit signed int by default
- Value 10000000000 exceeds Integer.MAX_VALUE (2147483647)
- Compiler detects overflow during lexical analysis phase
- `L` suffix triggers parsing as 64-bit long literal
- `lstore` bytecode generated for valid long assignment

---

### 15. Ternary Operator Type Promotion
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int i = 10;
        double result = (i > 5) ? i : 3.14;
        System.out.println(result);
    }
}
```
**A:** `10.0`. When ternary branches have different numeric types, Java promotes both to the wider type (`double` here). The `int` 10 becomes `10.0`.

**Code Snippet Internal Behavior:**
- Ternary operator uses type promotion rules for result type
- Compiler determines common supertype of both branches (double)
- `i2d` bytecode converts int 10 to double 10.0
- Result stored as double regardless of which branch executes
- Type promotion happens at compile time, not runtime

---

## Section 2: Control Flow (Q16–Q28)

### 16. Switch Fall-Through
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int x = 2;
        switch (x) {
            case 1:
                System.out.println("one");
            case 2:
                System.out.println("two");
            case 3:
                System.out.println("three");
                break;
            default:
                System.out.println("other");
        }
    }
}
```
**A:**
```
two
three
```
Without `break`, execution falls through to the next case. `case 2` and `case 3` both execute before `break` halts execution.

**Code Snippet Internal Behavior:**
- Switch compiled to `tableswitch` or `lookupswitch` bytecode
- Each case becomes a jump target in bytecode
- Missing `break` allows execution to continue sequentially
- `tableswitch` used for dense case values, `lookupswitch` for sparse
- Fall-through is intentional language design feature

---

### 17. For Loop Variable Scope
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        for (int i = 0; i < 3; i++) {
            System.out.println(i);
        }
        System.out.println(i); // can we access i here?
    }
}
```
**A:** **Compile Error.** The variable `i` is scoped to the `for` loop block and is not accessible after the closing `}`.

**Code Snippet Internal Behavior:**
- Loop variable stored in method's local variable array
- Variable scope limited to loop's bytecode range
- Compiler enforces lexical scoping during bytecode generation
- Local variable index reused after loop exits
- Access attempt fails bytecode verification

---

### 18. Enhanced For Loop — Can't Modify Array Elements
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3};
        for (int x : arr) {
            x = x * 2; // does this modify the array?
        }
        System.out.println(arr[0]);
    }
}
```
**A:** `1`. The enhanced for loop copies the value into `x`. Modifying `x` does not affect the original array element. `arr[0]` remains `1`.

**Code Snippet Internal Behavior:**
- Enhanced for loop uses iterator internally (`ArrayIterator` for arrays)
- Each iteration loads array element onto operand stack
- `x` is a separate local variable, not a reference to array element
- `istore` bytecode modifies local variable, not array slot
- Array access via `iaload` bytecode is read-only in this context

---

### 19. do-while Runs at Least Once
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int i = 10;
        do {
            System.out.println("ran: " + i);
            i++;
        } while (i < 5);
    }
}
```
**A:** `ran: 10`. The body executes once before the condition is checked. Even though `10 < 5` is false, the body still runs once.

**Code Snippet Internal Behavior:**
- `do-while` uses different bytecode pattern than `while`
- Body executes first, then conditional jump at end
- `if_icmpge` bytecode checks condition after body execution
- Loop control flow differs from pre-test loops
- Guarantees at least one iteration

---

### 20. Labeled Break
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        outer:
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                if (j == 1) break outer;
                System.out.print(i + "" + j + " ");
            }
        }
    }
}
```
**A:** `00 `. `break outer` exits the entire outer loop when `j == 1` on the first `i` iteration.

**Code Snippet Internal Behavior:**
- Labeled break uses `goto` bytecode to jump outside nested loops
- Label becomes a bytecode target address
- Break statement generates unconditional jump to label
- JVM maintains call stack frame for nested loops
- Immediate exit from all nested loop levels

---

### 21. Switch with String (Java 7+)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String day = "MON";
        switch (day) {
            case "SAT": case "SUN":
                System.out.println("Weekend");
                break;
            case "MON": case "TUE": case "WED":
            case "THU": case "FRI":
                System.out.println("Weekday");
                break;
            default:
                System.out.println("Unknown");
        }
    }
}
```
**A:** `Weekday`

**Code Snippet Internal Behavior:**
- String switch compiled to `lookupswitch` with `hashCode()` optimization
- Compiler generates hash-based jump table for String cases
- `String.hashCode()` computed at runtime for case matching
- `equals()` called for hash collisions (rare)
- More efficient than chained if-else for many cases

---

### 22. Unreachable Code
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        return;
        System.out.println("unreachable");
    }
}
```
**A:** **Compile Error.** The statement after `return` is unreachable. Java detects this at compile time.

**Code Snippet Internal Behavior:**
- Compiler performs reachability analysis on all code paths
- `return` generates `return` bytecode with no following code
- Unreachable code detection prevents dead code generation
- Bytecode verification fails for unreachable statements
- Ensures all code paths are executable

---

### 23. while Loop with break
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int i = 0;
        while (true) {
            if (i == 3) break;
            System.out.print(i++ + " ");
        }
    }
}
```
**A:** `0 1 2 `

**Code Snippet Internal Behavior:**
- `while(true)` becomes infinite loop with conditional break
- `if_icmpne` bytecode compares `i` with 3
- `break` generates `goto` bytecode to exit loop
- `iinc` bytecode increments local variable each iteration
- Loop terminates when condition `i == 3` is met

---

### 24. continue vs break
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        for (int i = 0; i < 5; i++) {
            if (i % 2 == 0) continue;
            System.out.print(i + " ");
        }
    }
}
```
**A:** `1 3 `. `continue` skips the rest of the current iteration for even numbers.

**Code Snippet Internal Behavior:**
- `continue` generates `goto` bytecode to loop condition check
- Skips remaining statements in current iteration
- `if_icmpeq` checks if `i % 2 == 0`
- Even numbers trigger jump to next iteration
- Odd numbers proceed to `System.out.print`

---

### 25. Switch Expression (Java 14+)
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int day = 3;
        String name = switch (day) {
            case 1 -> "Monday";
            case 2 -> "Tuesday";
            case 3 -> "Wednesday";
            default -> "Other";
        };
        System.out.println(name);
    }
}
```
**A:** `Wednesday`. Switch expressions with arrow labels don't fall through and directly yield a value.

**Code Snippet Internal Behavior:**
- Switch expression compiled to `tableswitch` with value returns
- Arrow syntax eliminates fall-through behavior
- Each case directly returns value via `lookupswitch`
- No break statements needed - implicit returns
- Type inference determines expression return type

---

### 26. instanceof Check
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Object obj = "hello";
        System.out.println(obj instanceof String);
        System.out.println(obj instanceof Integer);
        System.out.println(null instanceof String);
    }
}
```
**A:**
```
true
false
false
```
`instanceof` returns `false` for `null` — it never throws a NullPointerException.

**Code Snippet Internal Behavior:**
- `instanceof` uses `instanceof` bytecode instruction
- Null check performed before type verification
- Returns false immediately for null references
- No exception thrown for null instanceof checks
- Type verification only occurs for non-null objects

---

### 27. Nested Ternary
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int x = 5;
        String result = x > 10 ? "big" : x > 3 ? "medium" : "small";
        System.out.println(result);
    }
}
```
**A:** `medium`

**Code Snippet Internal Behavior:**
- Nested ternary compiled to series of conditional jumps
- `x > 10` evaluated first (false)
- `x > 3` evaluated second (true)
- Each condition generates `if_icmpgt` bytecode
- Final result selected based on condition chain

---

### 28. for-each on null
**Q: What happens at runtime?**
```java
import java.util.List;
public class Main {
    public static void main(String[] args) {
        List<String> list = null;
        for (String s : list) {
            System.out.println(s);
        }
    }
}
```
**A:** **NullPointerException at runtime.** The enhanced for loop calls `.iterator()` on the collection internally. Calling any method on `null` throws NPE.

**Code Snippet Internal Behavior:**
- Enhanced for loop compiles to iterator pattern
- `list.iterator()` called before loop starts
- Null reference causes `NullPointerException`
- Iterator methods (`hasNext()`, `next()`) never reached
- NPE thrown before loop body execution

---

## Section 3: Arrays (Q29–Q38)

### 29. Array Default Values
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int[] ints = new int[3];
        boolean[] bools = new boolean[2];
        String[] strs = new String[2];
        System.out.println(ints[0]);
        System.out.println(bools[0]);
        System.out.println(strs[0]);
    }
}
```
**A:**
```
0
false
null
```
Arrays are always zero-initialized: numeric types → 0, boolean → false, object references → null.

**Code Snippet Internal Behavior:**
- `newarray` bytecode allocates array and zero-initializes all elements
- Numeric arrays: all bits set to 0 (int = 0, double = 0.0)
- Boolean arrays: all bits set to 0 (false)
- Object arrays: all references set to null (0x0)
- Zero initialization happens at array creation time

---

### 30. Array is an Object
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3};
        System.out.println(arr.length);
        System.out.println(arr instanceof Object);
    }
}
```
**A:**
```
3
true
```
Arrays in Java are objects. `length` is a field (not a method). Arrays extend `Object`.

**Code Snippet Internal Behavior:**
- Arrays are objects with special array class generated by JVM
- `arraylength` bytecode gets array length (not method call)
- `instanceof` bytecode checks if object is array type
- Array objects have class name like `[I` for int array
- Arrays inherit methods from Object class

---

### 31. ArrayIndexOutOfBoundsException
**Q: What happens?**
```java
public class Main {
    public static void main(String[] args) {
        int[] arr = new int[5];
        arr[5] = 10;
    }
}
```
**A:** **ArrayIndexOutOfBoundsException: Index 5 out of bounds for length 5.** Valid indices are 0 to `length-1`.

**Code Snippet Internal Behavior:**
- `iastore` bytecode checks array bounds before storing
- JVM validates index: 0 ≤ index < array.length
- Bounds check happens at runtime (not compile time)
- Exception thrown immediately when check fails
- Array stores remain unchanged after exception

---

### 32. Array Assignment — Reference Copy
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int[] a = {1, 2, 3};
        int[] b = a;       // b points to the same array!
        b[0] = 99;
        System.out.println(a[0]);
    }
}
```
**A:** `99`. `b = a` copies the reference, not the array contents. Both `a` and `b` point to the same array on the heap.

**Code Snippet Internal Behavior:**
- Assignment copies reference value (memory address), not array contents
- Both variables reference same array object
- `astore` bytecode stores same reference in different local variable
- Array modification affects both references (same underlying object)
- No array copying occurs - only reference assignment

---

### 33. 2D Array
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int[][] matrix = new int[3][4];
        System.out.println(matrix.length);
        System.out.println(matrix[0].length);
    }
}
```
**A:**
```
3
4
```
`matrix.length` is the number of rows. `matrix[0].length` is the number of columns in the first row.

**Code Snippet Internal Behavior:**
- 2D arrays are arrays of arrays (jagged arrays supported)
- `matrix.length` gives outer array length (row count)
- `matrix[0].length` gives inner array length (column count)
- `multianewarray` bytecode creates multi-dimensional arrays
- Each row is separate array object

---

### 34. Arrays.sort and Arrays.copyOf
**Q: What is the output?**
```java
import java.util.Arrays;
public class Main {
    public static void main(String[] args) {
        int[] arr = {5, 3, 1, 4, 2};
        int[] copy = Arrays.copyOf(arr, arr.length);
        Arrays.sort(copy);
        System.out.println(Arrays.toString(arr));
        System.out.println(Arrays.toString(copy));
    }
}
```
**A:**
```
[5, 3, 1, 4, 2]
[1, 2, 3, 4, 5]
```
`Arrays.copyOf` creates a true copy; sorting `copy` doesn't affect `arr`.

**Code Snippet Internal Behavior:**
- `Arrays.copyOf` creates new array with `arraycopy` native method
- `Arrays.sort` uses Dual-Pivot Quicksort algorithm
- Original array remains unchanged after copy sorting
- `arraycopy` performs native memory copy operation
- Separate array objects with identical initial content

---

### 35. Varargs as Array
**Q: What is the output?**
```java
public class Main {
    static int sum(int... nums) {
        int total = 0;
        for (int n : nums) total += n;
        return total;
    }

    public static void main(String[] args) {
        System.out.println(sum(1, 2, 3));
        int[] arr = {4, 5, 6};
        System.out.println(sum(arr));
    }
}
```
**A:**
```
6
15
```
Varargs (`int...`) is syntactic sugar for an array parameter. You can pass an array directly.

**Code Snippet Internal Behavior:**
- Varargs compiled to array parameter at bytecode level
- `sum(1,2,3)` creates new int[] array at call site
- `sum(arr)` passes existing array reference directly
- Method receives int[] parameter regardless of call syntax
- Compiler handles varargs to array conversion automatically

---

### 36. Array of Objects — Shallow Reference
**Q: What is the output?**
```java
public class Main {
    static class Box { int val; Box(int v) { val = v; } }

    public static void main(String[] args) {
        Box[] a = { new Box(1), new Box(2) };
        Box[] b = Arrays.copyOf(a, a.length); // import java.util.Arrays
        b[0].val = 99;
        System.out.println(a[0].val);
    }
}
```
**A:** `99`. `Arrays.copyOf` on object arrays is a **shallow copy** — both arrays contain references to the same `Box` objects. Modifying `b[0].val` modifies the same object that `a[0]` references.

**Code Snippet Internal Behavior:**
- Object array copy copies references, not objects themselves
- Both arrays contain references to same Box instances
- `getfield`/`putfield` bytecode operate on same object
- Shallow copy: array structure duplicated, objects shared
- Deep copy would require cloning each object individually

---

### 37. Negative Array Size
**Q: What happens at runtime?**
```java
public class Main {
    public static void main(String[] args) {
        int size = -1;
        int[] arr = new int[size];
    }
}
```
**A:** **NegativeArraySizeException at runtime.** Array sizes must be non-negative.

**Code Snippet Internal Behavior:**
- `newarray` bytecode checks array size before allocation
- Negative size triggers immediate exception
- JVM validates size ≥ 0 before memory allocation
- Exception thrown during array creation, not after
- No array object created when size is negative

---

### 38. Multi-dimensional Array — Ragged
**Q: Does this compile and what is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int[][] ragged = new int[3][];
        ragged[0] = new int[]{1};
        ragged[1] = new int[]{2, 3};
        ragged[2] = new int[]{4, 5, 6};
        System.out.println(ragged[2][2]);
    }
}
```
**A:** **Compiles and prints** `6`. Java supports ragged (jagged) arrays where each row can have a different length.

**Code Snippet Internal Behavior:**
- `new int[3][]` creates outer array with null references
- Each `new int[n]` creates separate inner array
- Ragged arrays stored as array of array references
- `aaload`/`iastore` bytecodes handle nested access
- No requirement for uniform row lengths

---

## Section 4: Static, Final & Initialization (Q39–Q52)

### 39. Static Variable Shared Across Instances
**Q: What is the output?**
```java
public class Main {
    static class Counter {
        static int count = 0;
        Counter() { count++; }
    }

    public static void main(String[] args) {
        new Counter();
        new Counter();
        new Counter();
        System.out.println(Counter.count);
    }
}
```
**A:** `3`. `static` variables belong to the class, not instances. All three constructor calls increment the same `count`.

**Code Snippet Internal Behavior:**
- Static variables stored in method area (class memory), not heap
- All instances share same static variable location
- Constructor bytecode increments same static field each time
- `getstatic`/`putstatic` bytecodes access static variables
- Static initialization happens when class is loaded

---

### 40. Static Initializer Block
**Q: What is the output?**
```java
public class Main {
    static int x;
    static {
        x = 42;
        System.out.println("static block: x = " + x);
    }

    public static void main(String[] args) {
        System.out.println("main: x = " + x);
    }
}
```
**A:**
```
static block: x = 42
main: x = 42
```
Static initializer blocks run once when the class is loaded, before `main`.

**Code Snippet Internal Behavior:**
- Static blocks executed during class initialization phase
- `<clinit>` method contains all static initializers
- Class loader runs static blocks before any method execution
- Static field initialization compiled into `<clinit>`
- Thread-safe class initialization by JVM

---

### 41. Instance Initializer Block Order
**Q: What is the output?**
```java
public class Main {
    int x;
    { x = 5; System.out.println("init block: " + x); }

    Main() {
        System.out.println("constructor: " + x);
    }

    public static void main(String[] args) {
        new Main();
    }
}
```
**A:**
```
init block: 5
constructor: 5
```
Instance initializer blocks run before the constructor body every time an object is created.

**Code Snippet Internal Behavior:**
- Instance initializer code copied into each constructor
- Runs after superclass constructor but before constructor body
- Compiled as part of constructor bytecode sequence
- `this()` constructor chaining includes initializer blocks
- Executes before explicit constructor code

---

### 42. final Method Cannot Be Overridden
**Q: Does this compile?**
```java
class Animal {
    final void speak() { System.out.println("..."); }
}

class Dog extends Animal {
    @Override
    void speak() { System.out.println("Woof"); } // try to override
}
```
**A:** **Compile Error.** A `final` method cannot be overridden in a subclass.

**Code Snippet Internal Behavior:**
- `final` method modifier prevents virtual method table override
- Compiler generates direct method calls (not virtual)
- Subclass attempt to override fails bytecode verification
- Method resolution happens at compile time for final methods
- No polymorphic dispatch for final methods

---

### 43. final Class Cannot Be Extended
**Q: Does this compile?**
```java
final class Immutable {}
class Sub extends Immutable {} // ERROR
```
**A:** **Compile Error.** A `final` class cannot be subclassed. `String` is a famous example of a `final` class.

**Code Snippet Internal Behavior:**
- `final` class modifier prevents inheritance in bytecode
- Compiler rejects extends clause for final classes
- JVM enforces final class restriction at class loading
- Final classes cannot have subclasses in inheritance hierarchy
- Security and immutability reasons for final classes

---

### 44. Static Method Cannot Access Non-Static Members
**Q: Does this compile?**
```java
public class Main {
    int x = 10;

    public static void main(String[] args) {
        System.out.println(x); // access instance field from static context
    }
}
```
**A:** **Compile Error.** Static methods don't have a `this` reference. You cannot access instance fields or methods without an object reference.

**Code Snippet Internal Behavior:**
- Static methods lack implicit `this` parameter in bytecode
- Instance field access requires object reference
- Compiler detects instance access from static context
- `getfield` bytecode requires object on stack
- Static methods called via `invokestatic`, not `invokevirtual`

---

### 45. final Variable — Effectively Immutable Reference
**Q: What is the output?**
```java
import java.util.ArrayList;
public class Main {
    public static void main(String[] args) {
        final ArrayList<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        System.out.println(list.size());
    }
}
```
**A:** `2`. `final` makes the reference immutable (you can't reassign `list`), but the object it points to is still mutable. You can freely modify the list's contents.

**Code Snippet Internal Behavior:**
- `final` applies to variable reference, not object contents
- Reference stored in local variable table cannot be changed
- Object methods (like `add()`) still modify the object
- `astore` bytecode prevents reassignment of final variable
- Object mutability unchanged by final reference

---

### 46. Static Field vs Instance Field
**Q: What is the output?**
```java
public class Main {
    static class A {
        static int s = 1;
        int i = 2;
    }

    public static void main(String[] args) {
        A a1 = new A(); A a2 = new A();
        a1.s = 10; // accessing static via instance (bad style!)
        a1.i = 20;
        System.out.println(a2.s); // what does a2 see?
        System.out.println(a2.i);
    }
}
```
**A:**
```
10
2
```
`s` is static — shared. Changing it via `a1` changes it for all instances. `i` is instance-specific.

**Code Snippet Internal Behavior:**
- Static field accessed via `getstatic` bytecode (ignores instance)
- Instance field accessed via `getfield` bytecode (requires object)
- Static access through instance is compiler convenience
- Static fields stored once per class, not per object
- Instance fields stored separately for each object

---

### 47. Blank Final — Must Be Assigned in Every Constructor
**Q: Does this compile?**
```java
class Broken {
    final int x;
    Broken(boolean flag) {
        if (flag) x = 1; // not assigned if flag is false!
    }
}
```
**A:** **Compile Error.** A blank `final` field must be **definitely assigned** in every constructor path. The compiler detects that `x` might not be assigned when `flag` is `false`.

**Code Snippet Internal Behavior:**
- Compiler performs definite assignment analysis for final fields
- All code paths must assign final fields before constructor end
- Blank finals have special initialization requirements
- Bytecode verification ensures final fields are initialized
- Analysis considers all possible execution paths

---

### 48. Static Method Hiding (Not Overriding)
**Q: What is the output?**
```java
class Parent {
    static void greet() { System.out.println("Parent"); }
}

class Child extends Parent {
    static void greet() { System.out.println("Child"); }
}

public class Main {
    public static void main(String[] args) {
        Parent p = new Child();
        p.greet();
    }
}
```
**A:** `Parent`. Static methods are **hidden**, not overridden. The call is resolved at compile time based on the declared type of `p` (`Parent`), not the runtime type (`Child`).

**Code Snippet Internal Behavior:**
- Static method calls use `invokestatic` bytecode (compile-time binding)
- Method selection based on reference type, not object type
- No virtual method table lookup for static methods
- Hiding vs overriding: static methods don't participate in polymorphism
- Compiler resolves to Parent.greet() at compile time

---

### 49. Initialization Order — Field vs Constructor
**Q: What is the output?**
```java
public class Main {
    int x = 10;
    Main() { x = 20; }

    public static void main(String[] args) {
        System.out.println(new Main().x);
    }
}
```
**A:** `20`. Fields are initialized first (`x = 10`), then the constructor body runs (`x = 20`), so the final value is `20`.

**Code Snippet Internal Behavior:**
- Field initialization compiled into constructor before constructor body
- Instance initializer blocks and field initializers executed first
- Constructor body bytecode executes after field initialization
- Final field assignment can be overridden by constructor
- Order: superclass constructor → field initializers → constructor body

---

### 50. final static Constant Inlining
**Q: What is the output?**
```java
class Constants {
    static final int VALUE = 42;
}

public class Main {
    public static void main(String[] args) {
        System.out.println(Constants.VALUE);
    }
}
```
**A:** `42`. `static final` primitive constants are often inlined by the compiler at the call site.

**Code Snippet Internal Behavior:**
- Constant folding: compile-time evaluation of static final primitives
- Value embedded directly in calling bytecode
- No field access at runtime - value is literal
- Compiler treats as compile-time constant
- Reduces runtime overhead for constant access

---

### 51. Can You Call Overridden Methods from Constructor?
**Q: What is the output?**
```java
class Base {
    Base() { print(); }
    void print() { System.out.println("Base"); }
}

class Derived extends Base {
    int x = 10;
    Derived() { super(); }
    @Override
    void print() { System.out.println("Derived: x = " + x); }
}

public class Main {
    public static void main(String[] args) {
        new Derived();
    }
}
```
**A:** `Derived: x = 0`. **Classic trap.** `super()` → calls `Base()` → calls `print()` which is overridden. `Derived.print()` runs, but at this point `Derived`'s fields haven't been initialized yet, so `x` is `0` (default).

**Code Snippet Internal Behavior:**
- Superclass constructor runs before subclass field initialization
- Virtual method call from constructor uses dynamic dispatch
- Subclass fields still have default values during superclass constructor
- Object initialization order: superclass constructor → subclass fields → subclass constructor
- Dangerous to call overridable methods from constructors

---

### 52. Static Context — No Access to this
**Q: Does this compile?**
```java
public class Main {
    int num = 5;

    static int doubled() {
        return this.num * 2; // using 'this' in static method
    }

    public static void main(String[] args) {
        System.out.println(doubled());
    }
}
```
**A:** **Compile Error.** `this` cannot be used in a static context. Static methods have no instance (`this`) reference.

**Code Snippet Internal Behavior:**
- Static methods lack implicit `this` parameter in method signature
- `this` reference only available in instance methods
- Compiler rejects `this` usage in static context
- Static methods accessed via `invokestatic`, not `invokevirtual`
- No object instance available during static method execution

---

## Section 5: Wrapper Classes & Autoboxing (Q53–Q65)

### 53. Autoboxing/Unboxing — NullPointerException
**Q: What happens at runtime?**
```java
public class Main {
    public static void main(String[] args) {
        Integer i = null;
        int x = i; // unboxing null
        System.out.println(x);
    }
}
```
**A:** **NullPointerException at runtime.** Unboxing `null` throws NPE because Java calls `i.intValue()` on a null reference.

**Code Snippet Internal Behavior:**
- Unboxing compiles to `Integer.intValue()` method call
- Null reference triggers NPE when method invoked
- `invokevirtual` bytecode calls intValue() on null object
- Autoboxing/unboxing is compiler syntactic sugar
- NPE occurs at exact point of unboxing operation

---

### 54. Integer Cache — == Pitfall
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Integer x = 100;
        Integer y = 100;
        Integer a = 200;
        Integer b = 200;
        System.out.println(x == y);
        System.out.println(a == b);
        System.out.println(a.equals(b));
    }
}
```
**A:**
```
true
false
true
```
Integer cache applies for -128 to 127. `100` hits the cache (same object). `200` creates new objects each time. Always use `.equals()` for `Integer` comparison.

**Code Snippet Internal Behavior:**
- `Integer.valueOf()` uses internal cache array for -128 to 127
- Cache size configurable via `-XX:AutoBoxCacheMax`
- Cached objects reused for multiple autoboxing operations
- `==` compares object references, `.equals()` compares int values
- New objects allocated for values outside cache range

---

### 55. Autoboxing in Collections
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(3);
        list.remove(1); // remove by index or value?
        System.out.println(list);
    }
}
```
**A:** `[1, 3]`. `List.remove(int index)` removes by index. `list.remove(1)` removes the element at index 1 (which is `2`). To remove by value, use `list.remove(Integer.valueOf(1))`.

**Code Snippet Internal Behavior:**
- `List.remove(int)` and `List.remove(Object)` cause overload ambiguity
- Compiler chooses primitive int overload for literal 1
- Index-based removal uses element position, not value matching
- `Integer.valueOf(1)` forces Object overload selection
- Method overload resolution happens at compile time

---

### 56. Comparing Wrapper Types with ==
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Double a = 1.5;
        Double b = 1.5;
        System.out.println(a == b);
        System.out.println(a.equals(b));
    }
}
```
**A:**
```
false
true
```
`Double`, `Float`, `Long` (outside -128 to 127), etc., are **never** cached. Always use `.equals()`.

**Code Snippet Internal Behavior:**
- Only Integer (and Byte, Character, Short, Boolean) have caches
- Floating-point types never cached due to NaN/Infinity complexity
- `Double.valueOf()` always creates new objects
- `==` compares different object references
- `.equals()` compares double values using `Double.compare()`

---

### 57. Integer.parseInt vs Integer.valueOf
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int a = Integer.parseInt("42");
        Integer b = Integer.valueOf("42");
        System.out.println(a);
        System.out.println(b);
        System.out.println(a == b); // unboxing here
    }
}
```
**A:**
```
42
42
true
```
`parseInt` returns `int` primitive. `valueOf` returns an `Integer` object. The `==` comparison unboxes `b` to `int`, so it compares primitives: `42 == 42` = `true`.

**Code Snippet Internal Behavior:**
- `parseInt` uses `Integer.parseInt()` static method
- `valueOf` uses `Integer.valueOf()` with caching
- `==` triggers unboxing of Integer to int
- Comparison becomes primitive int equality check
- Compiler inserts `Integer.intValue()` call for unboxing

---

### 58. Autoboxing Performance — ArrayList
**Q: What is the bug in tight-loop code?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        List<Long> list = new ArrayList<>();
        long sum = 0;
        for (long i = 0; i < 1_000_000; i++) {
            list.add(i);      // boxing: long → Long (1 million objects!)
            sum += list.get((int)i); // unboxing: Long → long
        }
        System.out.println(sum);
    }
}
```
**A:** Correct output but **extremely slow** — 1 million `Long` objects are autoboxed/unboxed. In performance-critical code, prefer primitive arrays or `LongStream`.

**Code Snippet Internal Behavior:**
- Each `list.add(i)` calls `Long.valueOf()` (potential new object)
- Each `list.get(i)` calls `Long.longValue()` (unboxing)
- Creates 1 million temporary Long objects on heap
- Garbage collection pressure from object creation
- Primitive operations are much faster than wrapper operations

---

### 59. Boolean.parseBoolean
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(Boolean.parseBoolean("true"));
        System.out.println(Boolean.parseBoolean("True"));
        System.out.println(Boolean.parseBoolean("yes"));
        System.out.println(Boolean.parseBoolean("1"));
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
`Boolean.parseBoolean` is case-insensitive for `"true"` only. Any other string (including `"yes"` and `"1"`) returns `false`.

**Code Snippet Internal Behavior:**
- `parseBoolean` uses `String.equalsIgnoreCase("true")`
- Only exact match for `"true"` (case-insensitive) returns true
- All other inputs return false (no error thrown)
- No parsing of numeric representations like `"1"`
- Simple boolean string recognition algorithm

---

### 60. Unboxing in Arithmetic
**Q: Does this compile and what is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Integer a = 10;
        Integer b = 20;
        System.out.println(a + b);
    }
}
```
**A:** **Compiles and prints** `30`. When using arithmetic operators on `Integer` objects, Java automatically unboxes them to primitives.

**Code Snippet Internal Behavior:**
- `+` operator triggers unboxing of both Integer operands
- Compiler inserts `Integer.intValue()` calls before addition
- `iadd` bytecode performs primitive addition
- Result autoboxed back to Integer for println
- Unboxing happens at compile time, not runtime

---

### 61. Character Wrapper
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        char c = 'a';
        Character ch = c;     // autoboxing
        System.out.println(Character.isLetter(ch));
        System.out.println(Character.toUpperCase(ch));
    }
}
```
**A:**
```
true
A
```

**Code Snippet Internal Behavior:**
- `char` to `Character` autoboxing uses `Character.valueOf()`
- `Character` methods operate on 16-bit Unicode code units
- `isLetter()` checks Unicode character properties
- `toUpperCase()` performs Unicode case conversion
- Static methods don't require Character object instance

---

### 62. Integer.MAX_VALUE Overflow
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(Integer.MAX_VALUE + 1);
    }
}
```
**A:** `-2147483648`. Integer arithmetic silently wraps around. No exception is thrown.

**Code Snippet Internal Behavior:**
- Integer overflow wraps using two's complement arithmetic
- `iadd` bytecode performs modular arithmetic (mod 2^32)
- `Integer.MAX_VALUE + 1` = `Integer.MIN_VALUE`
- CPU overflow flag ignored by JVM for integers
- Silent overflow is specified behavior for Java ints

---

### 63. String to Integer Edge Cases
**Q: What happens?**
```java
public class Main {
    public static void main(String[] args) {
        int a = Integer.parseInt("  42  "); // leading/trailing spaces
    }
}
```
**A:** **NumberFormatException at runtime.** `parseInt` does not trim whitespace. Use `Integer.parseInt("  42  ".trim())`.

**Code Snippet Internal Behavior:**
- `Integer.parseInt()` performs strict numeric parsing
- Whitespace characters cause parsing failure
- Method doesn't automatically trim input string
- `NumberFormatException` thrown at first invalid character
- `String.trim()` removes leading/trailing whitespace before parsing

---

### 64. Wrapper Type in switch
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        Integer x = 1;
        switch (x) {
            case 1: System.out.println("one"); break;
        }
    }
}
```
**A:** **Yes, compiles** (Java 5+). The `Integer` is auto-unboxed to `int` for the switch. But if `x` is `null`, this throws a **NullPointerException** at runtime.

**Code Snippet Internal Behavior:**
- Switch on wrapper types triggers unboxing to primitive
- Compiler inserts `Integer.intValue()` before switch bytecode
- `tableswitch`/`lookupswitch` operate on primitive int
- Null check not performed - NPE possible
- Autoboxing/unboxing happens transparently

---

### 65. Long vs Int Computation
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        int a = 1_000_000;
        int b = 1_000_000;
        long result = a * b;   // computed as int first!
        long safe   = (long)a * b;
        System.out.println(result);
        System.out.println(safe);
    }
}
```
**A:**
```
-727379968
1000000000000
```
`a * b` is computed as `int` (overflows!) then widened to `long`. Cast one operand to `long` first.

**Code Snippet Internal Behavior:**
- `imul` bytecode multiplies ints (32-bit arithmetic)
- Overflow occurs before widening to long
- `i2l` bytecode converts overflowed result to long
- Casting first operand to long uses `lmul` bytecode
- Mixed precision arithmetic follows strict promotion rules

---

## Section 6: Varargs & Overloading (Q66–Q75)

### 66. Method Overloading — Most Specific Match
**Q: What is the output?**
```java
public class Main {
    static void print(int x) { System.out.println("int: " + x); }
    static void print(double x) { System.out.println("double: " + x); }

    public static void main(String[] args) {
        print(5);
        print(5.0);
        print('A');
    }
}
```
**A:**
```
int: 5
double: 5.0
int: 65
```
`'A'` is a `char` which widens to `int` (65) to match the most specific overload.

**Code Snippet Internal Behavior:**
- Method overload resolution uses most specific type matching
- `char` → `int` widening preferred over `char` → `double`
- Compiler selects method with minimal type conversion
- `invokestatic` bytecode calls selected method variant
- Type conversion applied at call site

---

### 67. Varargs Overload Ambiguity
**Q: Does this compile?**
```java
public class Main {
    static void foo(int... x) {}
    static void foo(int x, int y) {}

    public static void main(String[] args) {
        foo(1, 2); // which method?
    }
}
```
**A:** **Compiles.** The two-arg call `foo(1, 2)` prefers the exact match `foo(int x, int y)` over the varargs version. Fixed-arity methods are preferred over varargs.

**Code Snippet Internal Behavior:**
- Compiler prefers exact parameter count match over varargs
- Varargs considered only when no exact match exists
- Method selection follows specificity hierarchy
- Fixed-arity methods have higher priority than varargs
- Call generates `invokestatic` for exact match method

---

### 68. Varargs is Always an Array
**Q: What is the output?**
```java
public class Main {
    static void inspect(Object... args) {
        System.out.println(args.getClass().getSimpleName());
        System.out.println(args.length);
    }

    public static void main(String[] args) {
        inspect("a", "b", "c");
    }
}
```
**A:**
```
Object[]
3
```
Varargs compiles to an array. `args` is literally an `Object[]`.

**Code Snippet Internal Behavior:**
- Varargs parameter compiled to array parameter at bytecode level
- `anewarray` creates Object[] at call site
- Arguments packed into array before method invocation
- Method receives array reference, not individual arguments
- Array length accessible via `arraylength` bytecode

---

### 69. Overloading with Null
**Q: What is the output?**
```java
public class Main {
    static void foo(String s) { System.out.println("String"); }
    static void foo(Object o) { System.out.println("Object"); }

    public static void main(String[] args) {
        foo(null); // which overload is called?
    }
}
```
**A:** `String`. Java picks the most specific type. `String` is more specific than `Object`. If two overloads are equally specific, a compile error occurs.

**Code Snippet Internal Behavior:**
- Null literal matches both reference types
- Compiler selects most specific type in inheritance hierarchy
- `String` is subclass of `Object`, therefore more specific
- Method resolution happens at compile time
- `invokestatic` bytecode calls String version

---

### 70. Return Type Not Part of Method Signature
**Q: Does this compile?**
```java
public class Main {
    static int compute() { return 1; }
    static double compute() { return 1.0; } // ERROR?
}
```
**A:** **Compile Error.** Overloading is based on parameter types, not return type. Two methods with the same name and parameters but different return types are not allowed.

**Code Snippet Internal Behavior:**
- Method signature includes name and parameter types only
- Return type not considered for overload resolution
- Compiler detects duplicate method signatures
- Bytecode generation fails for ambiguous overloads
- JVM method identification uses name+descriptor

---

### 71. Varargs with null
**Q: What is the output?**
```java
public class Main {
    static void greet(String... names) {
        System.out.println(names == null ? "null array" : "length: " + names.length);
    }

    public static void main(String[] args) {
        greet((String[]) null); // explicitly passing null array
        greet();                // passing no args → empty array
    }
}
```
**A:**
```
null array
length: 0
```
Passing `null` gives a null array. Passing no args gives an empty (length 0) array.

**Code Snippet Internal Behavior:**
- Explicit null array passed directly to method
- No arguments creates empty array with `anewarray`
- Varargs array created at call site
- Null array bypasses array creation
- Empty array still allocated (zero-length object)

---

### 72. Overloading and Inheritance
**Q: What is the output?**
```java
public class Main {
    static void describe(Object o) { System.out.println("Object"); }
    static void describe(String s) { System.out.println("String"); }

    public static void main(String[] args) {
        Object o = "hello";   // declared as Object
        describe(o);
    }
}
```
**A:** `Object`. Overloading is resolved at **compile time** based on the **declared type** of the variable (`Object`), not the runtime type (`String`). This is different from overriding.

**Code Snippet Internal Behavior:**
- Overload resolution uses compile-time type information
- Variable declaration type determines method selection
- Runtime object type ignored for overload resolution
- `invokestatic` bytecode calls Object.describe() at compile time
- Different from virtual method dispatch in overriding

---

### 73. Covariant Return Type
**Q: Does this compile?**
```java
class Animal { Animal create() { return new Animal(); } }

class Dog extends Animal {
    @Override
    Dog create() { return new Dog(); } // covariant return
}
```
**A:** **Yes, compiles.** Since Java 5, an overriding method can return a subtype of the original return type. This is called a **covariant return type**.

**Code Snippet Internal Behavior:**
- Covariant returns allowed since Java 5
- Subclass method can return more specific type
- JVM verifies return type compatibility at class loading
- Virtual method table updated with covariant return
- Type checking ensures subtype relationship

---

### 74. Autoboxing Overload Resolution
**Q: What is the output?**
```java
public class Main {
    static void show(long x) { System.out.println("long"); }
    static void show(Integer x) { System.out.println("Integer"); }

    public static void main(String[] args) {
        int i = 5;
        show(i);
    }
}
```
**A:** `long`. Java prefers **widening** (int → long) over **autoboxing** (int → Integer) during overload resolution.

**Code Snippet Internal Behavior:**
- Overload resolution follows strict precedence rules
- Primitive widening preferred over boxing operations
- int → long conversion uses `i2l` bytecode
- Boxing would require object allocation
- Compiler chooses most efficient conversion path

---

### 75. Overriding vs Hiding with return types
**Q: Does this compile?**
```java
class Parent {
    Number compute() { return 1; }
}

class Child extends Parent {
    @Override
    Integer compute() { return 2; } // Integer is a subtype of Number
}
```
**A:** **Yes, compiles.** Covariant return type. `Integer` extends `Number`, so the override is valid.

**Code Snippet Internal Behavior:**
- Return type covariance verified at compile time
- Subclass relationship checked: Integer ≤ Number
- Method descriptor updated in virtual method table
- Type safety maintained through inheritance
- Runtime dispatch returns Integer, assignable to Number

---

## Section 7: Basic OOP & References (Q76–Q90)

### 76. Object Equality — == vs equals
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s1 = new String("hello");
        String s2 = new String("hello");
        System.out.println(s1 == s2);
        System.out.println(s1.equals(s2));
    }
}
```
**A:**
```
false
true
```
`==` compares references (memory addresses). `new String(...)` always creates a new object. `.equals()` compares content.

**Code Snippet Internal Behavior:**
- `==` uses `if_acmpne` bytecode for reference comparison
- `.equals()` uses `invokevirtual` to call String.equals()
- String literals use intern pool, `new String()` creates heap objects
- Different memory addresses for `new String()` objects
- Content comparison uses character-by-character algorithm

---

### 77. Passing Objects to Methods
**Q: What is the output?**
```java
public class Main {
    static class Box { int val; }

    static void modify(Box b) {
        b.val = 99;
    }

    public static void main(String[] args) {
        Box box = new Box();
        box.val = 1;
        modify(box);
        System.out.println(box.val);
    }
}
```
**A:** `99`. Java is pass-by-value, but the value passed for objects is the **reference** (pointer) to the object. The method modifies the object through the same reference, so changes are visible.

**Code Snippet Internal Behavior:**
- Object reference passed by value (copy of reference)
- Both caller and callee have references to same object
- `putfield` bytecode modifies object's field through reference
- Reference value copied, but object is shared
- Changes to object state visible to caller

---

### 78. Reassigning Reference in Method — No Effect Externally
**Q: What is the output?**
```java
public class Main {
    static class Box { int val; }

    static void replace(Box b) {
        b = new Box(); // reassign local reference
        b.val = 99;
    }

    public static void main(String[] args) {
        Box box = new Box();
        box.val = 1;
        replace(box);
        System.out.println(box.val);
    }
}
```
**A:** `1`. Reassigning the local parameter `b` to a new object does NOT affect the original `box` reference in `main`. Java is strictly pass-by-value.

**Code Snippet Internal Behavior:**
- Parameter `b` is separate local variable containing reference copy
- `b = new Box()` changes local variable only
- Original `box` reference in main unchanged
- `new` creates separate object on heap
- Reference assignment doesn't affect caller's variable

---

### 79. Polymorphism — Runtime Method Dispatch
**Q: What is the output?**
```java
class Shape {
    void draw() { System.out.println("Drawing Shape"); }
}
class Circle extends Shape {
    @Override
    void draw() { System.out.println("Drawing Circle"); }
}

public class Main {
    public static void main(String[] args) {
        Shape s = new Circle();
        s.draw();
    }
}
```
**A:** `Drawing Circle`. Java uses **dynamic dispatch** — virtual method calls are resolved at runtime based on the actual object type, not the declared type.

**Code Snippet Internal Behavior:**
- Virtual method call uses `invokevirtual` bytecode
- Method resolution at runtime via virtual method table
- Object's actual class (Circle) determines method implementation
- Polymorphic dispatch overrides compile-time type
- VMT lookup finds Circle.draw() at runtime

---

### 80. Abstract Class vs Interface
**Q: Does this compile?**
```java
abstract class Vehicle {
    abstract void move();
    void refuel() { System.out.println("refueling"); }
}

class Car extends Vehicle {
    @Override
    void move() { System.out.println("car moving"); }
}

public class Main {
    public static void main(String[] args) {
        new Car().refuel();
    }
}
```
**A:** **Yes, compiles and prints** `refueling`. Abstract classes can have concrete methods. Subclasses only need to implement the abstract ones.

**Code Snippet Internal Behavior:**
- Abstract class can have both abstract and concrete methods
- Concrete methods inherited by subclasses
- Abstract methods must be implemented by concrete subclasses
- `invokevirtual` bytecode calls inherited concrete methods
- Abstract methods have no bytecode implementation

---

### 81. Interface Default Method (Java 8+)
**Q: What is the output?**
```java
interface Greeter {
    default String greet(String name) {
        return "Hello, " + name + "!";
    }
}

class EnglishGreeter implements Greeter {}

public class Main {
    public static void main(String[] args) {
        System.out.println(new EnglishGreeter().greet("World"));
    }
}
```
**A:** `Hello, World!`. Default methods in interfaces provide a default implementation that classes can optionally override.

**Code Snippet Internal Behavior:**
- Default methods compiled into interface bytecode
- Classes inherit default implementation if not overridden
- `invokeinterface` bytecode calls default method
- Default method called through interface method table
- Can be overridden by implementing class

---

### 82. Diamond Problem with Default Methods
**Q: Does this compile?**
```java
interface A { default void hello() { System.out.println("A"); } }
interface B { default void hello() { System.out.println("B"); } }

class C implements A, B {
    // no override provided
}
```
**A:** **Compile Error.** When a class implements two interfaces with the same default method, the class must explicitly override the method to resolve the ambiguity.

**Code Snippet Internal Behavior:**
- Multiple inheritance of default methods creates conflict
- Compiler requires explicit method to resolve ambiguity
- Interface method tables conflict for same signature
- Class must provide its own implementation
- Disambiguation required at compile time

---

### 83. super() Call in Constructor
**Q: What is the output?**
```java
class Animal {
    Animal() { System.out.println("Animal created"); }
}
class Dog extends Animal {
    Dog() {
        // super() is implicitly called first
        System.out.println("Dog created");
    }
}

public class Main {
    public static void main(String[] args) { new Dog(); }
}
```
**A:**
```
Animal created
Dog created
```
If you don't call `super()` explicitly, Java inserts an implicit `super()` as the first statement. The parent constructor always runs before the child constructor body.

**Code Snippet Internal Behavior:**
- Compiler inserts `invokespecial Animal.<init>()` at start
- Superclass constructor runs before subclass constructor
- Object initialization follows inheritance hierarchy
- Field initialization happens between constructors
- Ensures proper object initialization order

---

### 84. Object class methods
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        Object o = new Object();
        System.out.println(o instanceof Object);
        System.out.println(o.getClass().getName());
    }
}
```
**A:**
```
true
java.lang.Object
```
Every class in Java extends `Object`. `getClass()` returns the runtime class.

**Code Snippet Internal Behavior:**
- All classes inherit from Object implicitly
- `instanceof` bytecode checks class inheritance
- `getClass()` uses `invokevirtual Object.getClass()`
- Runtime type information stored in class metadata
- Class name available via reflection API

---

### 85. this() Constructor Chaining
**Q: What is the output?**
```java
class Point {
    int x, y;
    Point() { this(0, 0); }
    Point(int x, int y) {
        this.x = x; this.y = y;
        System.out.println("Point(" + x + ", " + y + ")");
    }
}

public class Main {
    public static void main(String[] args) { new Point(); }
}
```
**A:** `Point(0, 0)`. `this(0, 0)` delegates to the two-arg constructor.

**Code Snippet Internal Behavior:**
- Constructor chaining uses `invokespecial this.<init>()`
- First constructor calls second constructor
- Delegation prevents code duplication
- Both constructors share same field initialization
- Only one constructor actually executes field assignments

---

### 86. Interface Cannot Have Constructor
**Q: Does this compile?**
```java
interface MyInterface {
    MyInterface() {} // constructor in interface?
}
```
**A:** **Compile Error.** Interfaces cannot have constructors. They cannot be instantiated directly.

**Code Snippet Internal Behavior:**
- Interfaces cannot have `<init>` methods in bytecode
- Constructor syntax rejected by compiler
- Interfaces define contracts, not implementations
- No object instantiation possible for interfaces
- Only classes can be instantiated

---

### 87. Casting — ClassCastException
**Q: What happens at runtime?**
```java
public class Main {
    public static void main(String[] args) {
        Object obj = "hello";
        Integer num = (Integer) obj; // downcast String to Integer
    }
}
```
**A:** **ClassCastException at runtime.** You can only downcast to the actual runtime type of the object. Use `instanceof` to check before casting.

**Code Snippet Internal Behavior:**
- Cast compiled to `checkcast` bytecode
- Runtime type verification performed
- ClassCastException thrown if types incompatible
- `instanceof` bytecode checks type compatibility
- Safe casting requires type checking first

---

### 88. Multiple Interface Implementation
**Q: Does this compile?**
```java
interface Flyable { void fly(); }
interface Swimmable { void swim(); }

class Duck implements Flyable, Swimmable {
    public void fly() { System.out.println("Duck flying"); }
    public void swim() { System.out.println("Duck swimming"); }
}

public class Main {
    public static void main(String[] args) {
        Duck d = new Duck();
        d.fly(); d.swim();
    }
}
```
**A:** **Compiles and prints:**
```
Duck flying
Duck swimming
```
A class can implement multiple interfaces.

**Code Snippet Internal Behavior:**
- Class implements multiple interface method tables
- `invokeinterface` bytecode resolves method calls
- Multiple interface inheritance supported
- Each interface provides separate method contract
- Class must implement all interface methods

---

### 89. toString() Default
**Q: What is the output?**
```java
public class Main {
    static class Dog {}

    public static void main(String[] args) {
        Dog d = new Dog();
        System.out.println(d);
    }
}
```
**A:** Something like `Main$Dog@6d06d69c`. The default `toString()` from `Object` returns `ClassName@hexHashCode`. Override `toString()` in your class for meaningful output.

**Code Snippet Internal Behavior:**
- Default toString() uses `Object.toString()` implementation
- Format: className + '@' + Integer.toHexString(hashCode())
- hashCode() generates 32-bit integer hash
- Hexadecimal representation of memory address (historical)
- Override toString() for custom string representation

---

### 90. equals() and hashCode() Contract
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    static class Key {
        int val;
        Key(int v) { val = v; }
        // equals() not overridden!
    }

    public static void main(String[] args) {
        Map<Key, String> map = new HashMap<>();
        map.put(new Key(1), "one");
        System.out.println(map.get(new Key(1)));
    }
}
```
**A:** `null`. Since `equals()` and `hashCode()` are not overridden, two `new Key(1)` objects are treated as different keys. Override both `equals()` and `hashCode()` to make HashMap work correctly with custom key objects.

**Code Snippet Internal Behavior:**
- HashMap uses `hashCode()` to find bucket location
- `equals()` called to resolve hash collisions
- Default Object.equals() uses reference equality
- Different objects → different hash codes → different buckets
- Custom key objects require proper equals/hashCode implementation

---

## Section 8: Miscellaneous Gotchas (Q91–Q100)

### 91. main() Can Be Overloaded
**Q: Does this compile?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println("standard main");
    }

    public static void main(int x) {
        System.out.println("overloaded main: " + x);
    }
}
```
**A:** **Yes, compiles.** You can overload `main`. The JVM specifically calls `main(String[] args)` as the entry point; the overloaded version is just a regular method.

**Code Snippet Internal Behavior:**
- Method overloading allowed for any method name including main
- JVM entry point fixed to `public static void main(String[])`
- Overloaded main treated as normal static method
- `invokestatic` bytecode calls specific overloaded version
- Only one main method serves as JVM entry point

---

### 92. Object Comparison in Collections
**Q: What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Set<String> set = new HashSet<>();
        set.add("apple");
        set.add("apple");
        set.add("banana");
        System.out.println(set.size());
    }
}
```
**A:** `2`. `String` properly overrides `equals()` and `hashCode()`, so duplicate strings are rejected by `HashSet`.

**Code Snippet Internal Behavior:**
- HashSet uses `hashCode()` to determine bucket
- `equals()` called to detect duplicates within bucket
- String implements both methods correctly
- Same string content → same hash code → equals() true
- Duplicate detection prevents multiple entries

---

### 93. Stack Overflow
**Q: What happens?**
```java
public class Main {
    static void infinite() { infinite(); }

    public static void main(String[] args) {
        infinite();
    }
}
```
**A:** **StackOverflowError**. Each recursive call pushes a frame onto the call stack. Infinite recursion exhausts the stack. `Error` (not `Exception`) — still catchable but shouldn't be.

**Code Snippet Internal Behavior:**
- Each method call creates new stack frame
- Stack frames stored in JVM stack memory
- Infinite recursion exceeds stack size limit
- StackOverflowError thrown when stack full
- Error type indicates serious JVM problem

---

### 94. NullPointerException — Common Trap
**Q: What happens?**
```java
public class Main {
    public static void main(String[] args) {
        String s = null;
        System.out.println(s.length());
    }
}
```
**A:** **NullPointerException at runtime.** Calling any method on a `null` reference throws NPE.

**Code Snippet Internal Behavior:**
- Method invocation on null reference triggers NPE
- `invokevirtual` bytecode fails with null object
- NPE thrown before method execution
- Null check implicit in virtual method calls
- Most common runtime exception in Java

---

### 95. Checked vs Unchecked Exception
**Q: Does this compile?**
```java
import java.io.*;
public class Main {
    public static void main(String[] args) {
        FileReader fr = new FileReader("test.txt"); // checked exception
    }
}
```
**A:** **Compile Error.** `FileNotFoundException` (a checked exception) is thrown by `FileReader`. You must either wrap in `try-catch` or declare `throws FileNotFoundException` in the method signature.

**Code Snippet Internal Behavior:**
- Checked exceptions require compile-time handling
- Compiler enforces exception handling or declaration
- `throws` clause adds exception to method signature
- `try-catch` block handles exception locally
- Bytecode verification ensures proper exception handling

---

### 96. try-finally Always Executes
**Q: What is the output?**
```java
public class Main {
    static int test() {
        try {
            return 1;
        } finally {
            System.out.println("finally!");
            return 2; // overrides try's return!
        }
    }

    public static void main(String[] args) {
        System.out.println(test());
    }
}
```
**A:**
```
finally!
2
```
`finally` always executes. If `finally` has a `return`, it overrides any `return` in the `try` block.

**Code Snippet Internal Behavior:**
- `finally` block compiled to execute after try/catch
- Return value stored temporarily during try execution
- Finally block can override return value
- JVM guarantees finally execution before method return
- Multiple returns handled via stack manipulation

---

### 97. String Pool
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        String s1 = "hello";
        String s2 = "hello";
        String s3 = new String("hello");
        System.out.println(s1 == s2);     // both from pool
        System.out.println(s1 == s3);     // s3 is a new object
        System.out.println(s1.equals(s3)); // content equal
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

### 98. instanceof with Pattern Matching (Java 16+)
**Q: What is the output?**
```java
public class Main {
    static void describe(Object o) {
        if (o instanceof String s) {
            System.out.println("String of length: " + s.length());
        } else if (o instanceof Integer i) {
            System.out.println("Integer: " + i);
        }
    }

    public static void main(String[] args) {
        describe("hello");
        describe(42);
    }
}
```
**A:**
```
String of length: 5
Integer: 42
```
Pattern matching for `instanceof` (Java 16+) lets you bind a variable in the same line. The variable `s` is in scope only inside the `if` block.

**Code Snippet Internal Behavior:**
- `instanceof` with pattern matching compiled to conditional binding
- Variable scope limited to pattern-matched block
- Type check and variable assignment combined
- Compiler generates efficient bytecode for pattern matching
- Reduces explicit casting boilerplate

---

### 99. Math Class Common Methods
**Q: What is the output?**
```java
public class Main {
    public static void main(String[] args) {
        System.out.println(Math.max(10, 20));
        System.out.println(Math.abs(-15));
        System.out.println(Math.pow(2, 10));
        System.out.println(Math.floor(3.9));
        System.out.println(Math.ceil(3.1));
    }
}
```
**A:**
```
20
15
1024.0
3.0
4.0
```

**Code Snippet Internal Behavior:**
- Math methods compiled to native CPU instructions
- `Math.max()` uses conditional move or CPU max instruction
- `Math.abs()` handles integer overflow correctly
- `Math.pow()` uses floating-point exponentiation
- `Math.floor()`/`ceil()` use IEEE 754 rounding modes

---

### 100. Record (Java 16+)
**Q: What is the output?**
```java
public class Main {
    record Point(int x, int y) {}

    public static void main(String[] args) {
        Point p1 = new Point(1, 2);
        Point p2 = new Point(1, 2);
        System.out.println(p1.x());
        System.out.println(p1.equals(p2));
        System.out.println(p1);
    }
}
```
**A:**
```
1
true
Point[x=1, y=2]
```
Records automatically generate: `equals()`, `hashCode()`, `toString()`, and accessor methods (like `x()`). They are immutable data carriers.

**Code Snippet Internal Behavior:**
- Record compiled to final class with private final fields
- Constructor, accessors, equals/hashCode/toString auto-generated
- Implements `java.lang.Record` interface
- Immutable by design - no setter methods
- Compact constructor syntax for object creation
