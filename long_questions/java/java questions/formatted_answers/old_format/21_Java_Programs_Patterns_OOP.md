# 21. Java Programs (Patterns & OOP Concepts)

**Q: Star Patterns (Triangle, Pivot, Diamond)**
> "Pattern printing is all about nested loops.
>
> 1.  **Right Triangle**:
>     *   Outer loop `i` from 1 to n (rows).
>     *   Inner loop `j` from 1 to `i` (columns).
>     *   Print `*`.
>
> 2.  **Pyramid (Centered)**:
>     *   Outer loop `i`.
>     *   Inner loop 1: Print spaces (`n - i`).
>     *   Inner loop 2: Print stars (`2*i - 1`).
>
> 3.  **Diamond**:
>     *   Print top half (Pyramid).
>     *   Print bottom half (Inverted Pyramid).
>
> **Tip**: Focus on the relation between Row number and Count of stars/spaces. Don't memorize code; derive the formula."

**Indepth:**
> **Formatting**: Use `System.out.printf()` or `String.format()` for complex patterns to align columns perfectly (e.g., %4d) instead of manually calculating spaces.


---

**Q: Floyd's Triangle (0-1 Pattern)**
> "1
> 0 1
> 1 0 1
>
> Logic:
> Loop `i` for rows, `j` for cols.
> If `(i + j)` is Even, print 1.
> If `(i + j)` is Odd, print 0.
> Alternatively, use a boolean flag and flip it `!flag` every time."

**Indepth:**
> **Parity**: This pattern is essentially checking checking the parity (even/odd) of the grid coordinates. It's equivalent to a chessboard coloring problem.


---

**Q: Pascal's Triangle**
> "1
> 1 1
> 1 2 1
>
> Each number is the sum of the two numbers directly above it.
>
> **Efficient Way**:
> Use the Combination formula `nCr`.
> `Value = (Previous_Value * (Row - Col + 1)) / Col`.
> This allows you to calculate the next number in the row using the previous number, without calculating factorials every time."

**Indepth:**
> **Binomial Coefficients**: Each entry is `nCr` (n choose r). The sum of numbers in row `n` is `2^n`. This is a powerful property used in probability theory.


---

**Q: Spiral Pattern (Number Grid)**
> "Input: 3. Output:
> 3 3 3 3 3
> 3 2 2 2 3
> 3 2 1 2 3
> 3 2 2 2 3
> 3 3 3 3 3
>
> This looks hard but has a trick.
> For every cell (i, j), the value is `n - min(distance to any wall)`.
> Distance to Top: `i`
> Distance to Bottom: `size - 1 - i`
> Distance to Left: `j`
> Distance to Right: `size - 1 - j`
> Take the minimum of these 4 distances, and subtract from `n`."

**Indepth:**
> **Generalization**: This "Distance Transform" logic works for any grid filling problem. Instead of simulating the movement (right, down, left, up), you calculate the value based on coordinates `(i, j)`.


---

**Q: Singleton Class (Implementation)**
> "To limit a class to exactly one instance:
>
> 1.  Private Static Variable: `private static Singleton instance;`
> 2.  Private Constructor: `private Singleton() {}` (stops `new S()`).
> 3.  Public Static Method:
> ```java
> public static Singleton getInstance() {
>     if (instance == null) {
>         instance = new Singleton();
>     }
>     return instance;
> }
> ```
> **Note**: For thread safety, use 'Double-Checked Locking' or simpler: `private static final Singleton INSTANCE = new Singleton();` (Eager initialization)."

**Indepth:**
> **Bill Pugh**: The thread-safe, efficient, and clean way is the **Bill Pugh Singleton Implementation**:
> `private static class Holder { static final Singleton INSTANCE = new Singleton(); }`
> `public static Singleton getInstance() { return Holder.INSTANCE; }`
> It uses the classloader guarantees for thread safety and is lazy-loaded.


---

**Q: Immutable Class**
> "How to make a class like `String`:
>
> 1.  Make class `final` (so no one can extend and override it).
> 2.  Make all fields `private` and `final`.
> 3.  No Setters.
> 4.  Initialize via Constructor.
> 5.  **Crucial**: If a field is mutable (like `Date` or `List`), don't return the original reference in the Getter. Return a copy (clone). Otherwise, the caller can modify the internal state."

**Indepth:**
> **Defensive Copying**: Both in the constructor (when accepting a mutable object) and in the getter (when returning it). If you miss either one, immutability is broken.


---

**Q: Method Overloading vs Overriding**
> "**Overloading** happens in the **same class**.
> *   Same method name.
> *   **Different parameters** (type or count).
> *   Return type doesn't matter.
> *   Compile-time polymorphism.
>
> **Overriding** happens in **Child classes**.
> *   Same method name.
> *   **Same parameters**.
> *   Runtime polymorphism.
> *   Used to change behavior inherited from Parent."

**Indepth:**
> **Binding**: Overloading is **Static Binding** (Early Binding) - Compiler decides. Overriding is **Dynamic Binding** (Late Binding) - JVM decides.


---

**Q: Abstract Class vs Interface**
> "**Interface**: Describes **Capabilities** (what it *can* do). 'I can Fly, I can Swim'. A class can implement multiple capabilities. Use for defining contracts.
>
> **Abstract Class**: Describes **Identity** (what it *is*). 'I am a Vehicle'. A class can only be ONE thing (single inheritance). Use when you have shared state (variables) or concrete code that children should reuse."

**Indepth:**
> **Evolution**: Interfaces can now have `default` methods (behavior) and `static` constants (state-ish). Note: They still can't have instance variables (true state). This blurs the line, but the "Is-A" vs "Can-Do" distinction remains.


---

**Q: Deep Copy vs Shallow Copy**
> "**Shallow Copy** (Default `clone()`):
> Copies the object, but the fields still point to the same references.
> `NewObject.list == OldObject.list`. Adding to one affects both.
>
> **Deep Copy**:
> Recursively copies everything.
> `NewObject.list = new ArrayList(OldObject.list)`. The objects are completely independent."

**Indepth:**
> **Copy Constructor**: The best way to implement Deep Copy is using a **Copy Constructor**: `public User(User other)`. It's explicit, doesn't throw exceptions, and avoids the weirdness of `Cloneable`.


---

**Q: Static Block vs Instance Block**
> "**static { ... }**:
> *   Runs **once** when the Class is loaded by ClassLoader.
> *   Used to initialize static variables.
>
> **{ ... }** (Instance Block):
> *   Runs **every time** you create a new Object (`new`).
> *   Runs *before* the constructor.
> *   Rarely used; usually we just put this code in the Constructor."

**Indepth:**
> **Bytecode**: Static blocks are compiled into a special method called `<clinit>` (Class Init). Instance blocks are copied into every `<init>` (Constructor) method.


---

**Q: Comparator vs Comparable**
> "**Comparable** (Internal): The class defines its own sort order (`compareTo`). 'I compare myself to others'.
>
> **Comparator** (External): A separate judge class (`compare`). 'I compare these two objects'.
>
> Use `Comparable` for default sorting (ID, Name). Use `Comparator` for ad-hoc sorting (Sort by Salary, Sort by Age)."

**Indepth:**
> **Contract**: `compareTo` returns negative (less), zero (equal), or positive (greater). A common trick `return this.val - other.val` is dangerous because of integer overflow! Use `Integer.compare(this.val, other.val)` instead.

