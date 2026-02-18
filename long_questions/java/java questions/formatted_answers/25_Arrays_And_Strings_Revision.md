# 25. Data Structures Revision (Arrays & Strings)

**Q: Arrays.copyOf() vs System.arraycopy()**
> "**Arrays.copyOf()** is the easy, high-level way.
> You pass the original array and the new length, and it returns a **new** array. Use it when you just want a copy or need to resize.
>
> **System.arraycopy()** is the low-level, high-performance way.
> You must create the target array **first**. It copies data from source to destination at a specific index. It’s a native method (written in C/C++), making it extremely fast. Use it for complex shifting or merging."

**Indepth:**
> **Under the Hood**: `Arrays.copyOf()` checks the new length. If it's larger, it pads with default values (null/0). If smaller, it truncates. It then calls `System.arraycopy` internally.


---

**Q: Shallow Copy vs Deep Copy (Arrays)**
> "If you have an `int[]` array, a copy is just numbers. Easy.
>
> But if you have an `Employee[]` array:
> *   **Shallow Copy** (Default): The new array points to the **same** Employee objects. If you change Employee #1's name in the copy, the original array sees the change.
> *   **Deep Copy**: You must iterate through the array and manually create `new Employee()` for every single element. It’s the only way to have truly independent data."

**Indepth:**
> **Serialization**: One way to achieve a deep copy without writing manual code is to Serialize the object to a Byte Stream and immediately Deserialize it. It's slower but foolproof for complex graphs of objects.


---

**Q: Arrays.asList() Caveats**
> "Stop treating `Arrays.asList()` like a normal ArrayList.
> It creates a wrapper backed by the **original array**.
>
> 1.  **You cannot add/remove**: It throws `UnsupportedOperationException` because the underlying array size is fixed.
> 2.  **Changes reflect**: If you change an element in the list (`list.set(0, "X")`), the original array **also changes**.
>
> Always wrap it: `new ArrayList<>(Arrays.asList(...))` to be safe."

**Indepth:**
> **View**: Think of `Arrays.asList` as a "View" or "Window" onto the array. It doesn't own the data; the array does. Any change to the view passes through to the backing array.


---

**Q: Arrays.equals() vs ==**
> "Never use `==` on arrays unless you are checking if they are the exact same object in memory.
>
> `int[] a = {1, 2}; int[] b = {1, 2};`
> *   `a == b` is **false**.
> *   `a.equals(b)` is **false** (Arrays don't override `equals` from Object class!).
>
> **The Solution**: Use `Arrays.equals(a, b)` (for 1D arrays) or `Arrays.deepEquals(a, b)` (for 2D arrays)."

**Indepth:**
> **Multidimensional**: `Arrays.equals` only checks the first layer. If you have an array of arrays `int[][]`, the "elements" are array objects. `equals` checks if those array object references are the same (which they aren't). `deepEquals` recursively checks the contents.


---

**Q: String vs StringBuilder vs StringBuffer**
> "**String** is Immutable. Every time you say `str + "A"`, you create a brand new object. Slow for loops.
>
> **StringBuilder** is Mutable. It modifies the existing buffer. It is **Not Thread Safe**, but very fast. Use this 99% of the time.
>
> **StringBuffer** is the legacy version. It is **Synchronized** (Thread Safe), which makes it slower. Only use it if multiple threads are editing the string simultaneously (rare)."

**Indepth:**
> **Capacity**: `StringBuilder` has a capacity. When you append past the limit, it has to resize (usually doubles) which involves copying the old array to a new one. Setting an initial capacity close to expected size improves performance.


---

**Q: String Pool & Immutability**
> "Java saves memory by keeping a 'Pool' of unique strings.
> If you write `String s1 = "Hello"` and `String s2 = "Hello"`, they point to the exact same spot in memory.
>
> **Immutability** is crucial for security and safety.
> If Strings were mutable, I could pass a Database Connection URL to a function, and that function could change the URL to 'MaliciousSite.com', affecting everyone else using that string. Immutability prevents this."

**Indepth:**
> **GC**: String Deduplication (G1GC feature) allows the Garbage Collector to inspect the heap for duplicate strings. If it finds two identical string objects, it makes them share the same underlying `char[]` array to save RAM.


---

**Q: substring() Memory Leak (Historical)**
> "In older Java versions (pre-Java 7), `substring()` didn't create a new string. It just pointed to the original massive string array with a different start/end offset.
> This meant if you loaded a 10MB text file and blindly took a 5-byte substring `str.substring(0, 5)`, Java kept the **entire 10MB** in memory just for those 5 bytes.
>
> Modern Java (JDK 7u6+) fixed this: `substring()` now creates a fresh array copy. But it's good to know the history."

**Indepth:**
> **Offset**: The `String` class used to have `offset` and `count` fields to share `char[]`. This was removed to prevent memory leaks where a tiny substring prevents a massive collection from being GC'd.


---

**Q: equals() vs equalsIgnoreCase()**
> "Simple but standard.
> *   `"A".equals("a")` is **false**.
> *   `"A".equalsIgnoreCase("a")` is **true**.
>
> **Tip**: Always put the constant on the left to avoid NullPointers.
> Don't write `userInput.equals("ADMIN")`.
> Write `"ADMIN".equals(userInput)`. This works even if `userInput` is null."

**Indepth:**
> **Null Safety**: `Objects.equals(a, b)` is the safest way. It handles null checks for both `a` and `b` automatically. `Objects.equals(null, "A")` returns false, not a crash.


---

**Q: split() vs StringTokenizer**
> "**StringTokenizer** is legacy. It was there before Regex existed. Do not use it.
>
> use `str.split(regex)`.
> *   `"a,b,,c".split(",")` gives `["a", "b", "", "c"]`.
> *   **Watch out**: Trailing empty strings are discarded by default. Use `split(regex, -1)` to keep them."

**Indepth:**
> **Performance**: `StringTokenizer` is faster than `split` for simple delimiters because it doesn't use Regex. However, it's considered deprecated for new code due to its limited API and confusing behavior with empty tokens. Use `split` or `Guava Splitter`.


---

**Q: First Non-Repeating Character**
> "The logic:
> 1.  Loop once to populate a Frequency Map (`LinkedHashMap` preserves order, or just `int[256]` for ASCII).
> 2.  Loop through the **string** again.
> 3.  The first character with count 1 is the winner."

**Indepth:**
> **Optimization**: If the string only contains standard ASCII (0-127), a `boolean[128]` array is enough. If Extended ASCII, `boolean[256]`. If Unicode, you effectively need a `HashMap` or a sparse array.

