# 15. String Manipulation Basics

**Q: Difference between String, StringBuilder, and StringBuffer**
> "This is the classic Java interview question.
>
> **String** is **Immutable**. Once you create `"Hello"`, it can never change. If you do `"Hello" + " World"`, you aren't changing the original string; you are creating a brand new object in memory. This is safe but slow if you do it inside a loop.
>
> **StringBuilder** is **Mutable**. You can modify it (`.append()`) without creating new objects. It's fast and efficient. It is **not** thread-safe, but that's usually what you want for local variables.
>
> **StringBuffer** is the old, legacy version of StringBuilder. It is also mutable, but it is **Synchronized** (thread-safe). This makes it slower. You almost never need it anymore unless you are sharing a string builder between threads (which is rare)."

**Indepth:**
>

---

**Q: String Immutability — Why is it important?**
> "Java made Strings immutable for several key reasons:
>
> 1.  **Security**: Strings are used for everything—database URLs, passwords, file paths. If I pass a filename string to a method, I need to be 100% sure that method can't modify my string and trick me into writing to the wrong file.
> 2.  **Caching (String Pool)**: Because they can't change, Java can safely store one copy of `"Hello"` and let 100 different variables point to it. This saves massive amounts of memory.
> 3.  **Thread Safety**: Immutable objects are automatically thread-safe. You can share a String across threads without any locking."

**Indepth:**
>

---

**Q: Reverse, Palindrome, Anagrams (Logic)**
> "These are the bread-and-butter of coding rounds.
>
> *   **Reverse**: You can use `StringBuilder.reverse()`, but interviewers usually want you to do it manually. Convert to `char[]`, then swap start/end pointers until they meet.
> *   **Palindrome**: A string that reads the same forwards and backwards. Just reuse your reverse logic: does `str.equals(reverse(str))`? Or check pointers: `charAt(0) == charAt(len-1)`, etc.
> *   **Anagrams**: Two strings with the same characters in different orders (e.g., 'listen' and 'silent'). The easiest way? Sort both strings and check if they are equal. The faster way? Use a frequency map (or int[26] array) to count character occurrences."

**Indepth:**
> **Thread Safety**: `StringBuffer` is synchronized (thread-safe) but slow. `StringBuilder` is not synchronized but fast. Since Java 5, `StringBuilder` is the default choice.


---

**Q: First non-repeating character**
> "To find the first unique character (like 'l' in 'google'), you need two passes.
>
> Pass 1: Loop through the string and build a frequency map (`Map<Character, Integer>`) counting how many times each char appears.
> Pass 2: Loop through the string *again* (not the map, because order matters). Check the map for each char. The first one with a count of 1 is your winner."

**Indepth:**
> **UTF-16**: Java Strings use UTF-16 encoding internally. Most characters take 2 bytes, but some rare Unicode characters (emojis) take 4 bytes (surrogate pairs). `length()` returns the number of 2-byte code units, not the number of actual characters!


---

**Q: equals() vs equalsIgnoreCase() vs ==**
> "Never use `==` for Strings!
> `==` checks **Reference Equality**—are these two variables pointing to the exact same object in heap memory? Even if both contain "hello", `==` might return false if one was created with `new String()`.
>
> Always use `.equals()` for **Content Equality**. It checks if the actual characters are the same.
>
> `.equalsIgnoreCase()` is just a convenience wrapper that ignores casing, so 'JAVA' equals 'java'."

**Indepth:**
>

---

**Q: substring() vs subSequence()**
> "Functionally, they do almost the same thing.
>
> `substring()` returns a **String**. It’s what you use 99% of the time.
> `subSequence()` works on the `CharSequence` interface (which `String`, `StringBuilder`, and `StringBuffer` all implement). You usually only see this when using generic APIs that accept any char sequence.
>
> Historically (pre-Java 7), `substring` caused memory leaks because it shared the underlying char array. Modern Java doesn't do that—it copies the data, so it's safe."

**Indepth:**
>

---

**Q: trim() vs strip() (Java 11)**
> "For years, we used `trim()` to remove whitespace. But `trim()` is old—it only removes ASCII characters (space, tab, newline). It doesn't understand newer Unicode whitespace standards.
>
> Java 11 introduced `strip()`. It is 'Unicode-aware'. It removes all kinds of weird whitespace characters that `trim()` misses.
>
> Use `strip()` for modern applications. It also comes with `stripLeading()` and `stripTrailing()` for more control."

**Indepth:**
>

---

**Q: String Constant Pool**
> "This is a special area in the Heap memory.
>
> When you type `String s = "Hello";` (a literal), Java checks the pool. If "Hello" exists, it returns a reference to the existing one. If not, it adds it.
>
> When you type `String s = new String("Hello");`, you force Java to create a **new** object on the heap, bypassing the pool checks (though the internal char array might still be shared). This is generally wasteful and discouraged."

**Indepth:**
>

---

**Q: intern() method**
> "This method manually puts a String into the String Pool.
>
> If you have a String object on the heap (maybe read from a file), calling `.intern()` checks the pool.
> If the pool already has that value, it returns the pool's reference.
> If not, it adds your string to the pool and returns it.
>
> It’s a way to deduplicate strings in memory manually."

**Indepth:**
> **Conversion**: `String.valueOf(10)` vs `Integer.toString(10)`. They do the same thing. `"" + 10` also works but generates extra StringBuilder garbage.

