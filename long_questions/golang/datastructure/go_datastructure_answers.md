## 🟤 1–27: Strings & Text Processing

### Question 1: How are strings represented in memory in Go?

**Answer:**
A string in Go is a read-only slice of bytes. It is represented internally as a lightweight structure containing two fields:

1.  A pointer to the backing byte array.
2.  The length of the string (number of bytes).

**Internal Structure:**
```go
type StringHeader struct {
    Data uintptr
    Len  int
}
```

**Memory Layout:**
- Since strings are immutable, multiple strings can share the same underlying memory if one is a substring of the other, without reallocation.
- The `Len` field purely indicates how many bytes the string contains, not necessarily the number of characters (due to UTF-8 encoding).

**Example:**
```go
s := "Hello"
// Pointer -> [ 'H', 'e', 'l', 'l', 'o' ]
// Len    -> 5
```

---

### Question 2: What is the difference between string and []byte in Go?

**Answer:**
While both store sequences of bytes, they have fundamental differences in mutability and usage:

**Key Differences:**
1.  **Mutability:** `string` is immutable (cannot be changed once created). `[]byte` is mutable (elements can be modified).
2.  **Memory:** Strings are often optimized for sharing backing arrays. `[]byte` requires separate allocation if you want independent copies.
3.  **Usage:** Use `string` for text processing, keys in maps, and constant text. Use `[]byte` for I/O operations, mutable buffers, and binary data manipulation.

**Conversion:**
```go
s := "hello"
b := []byte(s) // Allocates new memory and copies bytes
b[0] = 'H'     // OK
// s[0] = 'H'  // Compilation Error
```

---

### Question 3: Why are strings immutable in Go?

**Answer:**
String immutability provides several benefits for performance and safety:

**Reasons:**
1.  **Thread Safety:** Immutable strings are inherently safe to access concurrently without locks.
2.  **Memory Efficiency:** Substrings can share the same underlying byte array as the original string without copying memory.
3.  **Map Keys:** Strings can be safely used as hash map keys because their hash value remains constant.
4.  **Compiler Optimization:** The compiler can optimize string handling knowing the contents won't change.

**Example of Sharing:**
```go
s := "hello world"
sub := s[0:5] // "hello"
// 'sub' shares the same memory address for data as 's'
```

---

### Question 4: What is the most efficient way to concatenate strings?

**Answer:**
The efficiency depends on the use case, but `strings.Builder` is generally the best for building strings in a loop.

**Comparison:**
1.  **`+` Operator:** Good for small, fixed numbers of strings.
    ```go
    s := "Hello" + " " + "World" // Simple and readable
    ```
2.  **`strings.Builder`:** Best for loops or building complex strings. It minimizes memory allocations.
    ```go
    var sb strings.Builder
    for i := 0; i < 100; i++ {
        sb.WriteString("data")
    }
    result := sb.String()
    ```
3.  **`fmt.Sprintf`:** Slower due to reflection and parsing overhead; use for complex formatting.
4.  **`bytes.Buffer`:** Similar to `strings.Builder` but creates a `[]byte` first; slightly less efficient for pure string building.

---

### Question 5: What is strings.Builder and how does it work?

**Answer:**
`strings.Builder` is a type in the `strings` package designed to efficiently build strings by minimizing memory copying.

**How it works:**
1.  It maintains an internal buffer (`[]byte`).
2.  Methods like `WriteString` append data to this buffer.
3.  The `String()` method returns the final string. Crucially, it uses `unsafe` pointers to convert the internal `[]byte` to `string` *without* a memory allocation/copy, unlike `bytes.Buffer`.

**Code Example:**
```go
import "strings"

func buildString() string {
    var sb strings.Builder
    sb.Grow(32) // Optional: Pre-allocate optimization
    sb.WriteString("Hello")
    sb.WriteString(", ")
    sb.WriteString("World")
    return sb.String() // No copy performed here
}
```

---

### Question 6: How do you iterate over a string containing multibyte characters?

**Answer:**
To correctly iterate over UTF-8 strings (which may contain multibyte characters like emojis or non-Latin letters), use the `range` loop.

**Mechanism:**
- A standard `for i := 0; i < len(s); i++` iterates by **byte**.
- A `for i, r := range s` iterates by **rune** (Unicode code point).

**Example:**
```go
s := "Hello 🌍"

// Correct way (by Rune)
for i, r := range s {
    fmt.Printf("%d: %c\n", i, r)
}
// Output:
// 0: H
// ...
// 6: 🌍 (takes 4 bytes)

// Incorrect way (by Byte) for characters
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}
```

---

### Question 7: What is a rune in Go?

**Answer:**
A `rune` in Go is an alias for `int32`. It strictly represents a Unicode Code Point.

**Key Definition:**
- **Type:** `int32`
- **Purpose:** To handle individual distinct characters, which might be composed of 1 to 4 bytes in UTF-8.
- **Literal:** Enclosed in single quotes, e.g., `'A'`, `'界'`, `'😊'`.

**Example:**
```go
var r rune = 'A' // 65
var emoji rune = '😎' // 128526
fmt.Printf("%T\n", r) // int32
```

---

### Question 8: What is the difference between len(str) and utf8.RuneCountInString(str)?

**Answer:**
They measure different properties of the string.

**Differences:**
1.  **`len(str)`:** Returns the number of **bytes** in the string. This is O(1).
2.  **`utf8.RuneCountInString(str)`:** Returns the number of **runes** (logical characters) in the string. This is O(n) because it must decode the UTF-8 sequences.

**Example:**
```go
import "unicode/utf8"

s := "José"
fmt.Println(len(s))                    // 5 (because 'é' is 2 bytes)
fmt.Println(utf8.RuneCountInString(s)) // 4 (J, o, s, é)
```

---

### Question 9: How does the strings package differ from the bytes package?

**Answer:**
Both packages provide similar utility functions (split, join, replace, index), but they target different data types.

**Comparison:**
- **`strings` Package:** specialized for `string` manipulation. It generally takes `string` arguments and returns `string` results.
- **`bytes` Package:** specialized for `[]byte` manipulation. It is useful when manipulating raw data streams or mutable buffers.

**When to use:**
- Use `strings` when working with text.
- Use `bytes` when processing I/O streams, binary protocols, or when you need mutability during the process.

---

### Question 10: How do you efficiently replace a substring multiple times?

**Answer:**
Use `strings.Replace` or `strings.ReplaceAll` (for all occurrences). For more complex pattern-based replacements, use `regexp` (though slower).

**Optimized Replacement:**
To utilize specific replacement counts:
```go
import "strings"

s := "apple banana apple orange"

// Replace first 'apple' only
res1 := strings.Replace(s, "apple", "mango", 1) 
// "mango banana apple orange"

// Replace all 'apple'
res2 := strings.ReplaceAll(s, "apple", "mango")
// "mango banana mango orange"
```

For highly frequent replacements in a loop, consider using `strings.Replacer`:
```go
r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
result := r.Replace("<div>Content</div>")
// "&lt;div&gt;Content&lt;/div&gt;"
```
`strings.Replacer` pre-compiles the replacement strategy, making it faster for multiple replacements.

---

### Question 11: How do you check if a string contains a substring?

**Answer:**
Use the `strings.Contains` function for a boolean check.

**Code:**
```go
import "strings"

s := "Golang is awesome"
if strings.Contains(s, "awesome") {
    fmt.Println("Found!")
}
```
**Other Variations:**
- `strings.HasPrefix(s, prefix)`: Check start.
- `strings.HasSuffix(s, suffix)`: Check end.
- `strings.Index(s, sub)`: Returns index or -1 if not found.

---

### Question 12: How do you split a string by whitespace or custom separators?

**Answer:**
Use `strings.Split` for custom separators and `strings.Fields` for whitespace.

**Methods:**
1.  **Split by specific separator:**
    ```go
    s := "a,b,c"
    parts := strings.Split(s, ",") // ["a", "b", "c"]
    ```
2.  **Split by whitespace (multiple spaces treated as one):**
    ```go
    s := "foo   bar  baz"
    parts := strings.Fields(s) // ["foo", "bar", "baz"]
    ```

---

### Question 13: How do you convert a string to an integer?

**Answer:**
Use the `strconv` package, specifically `strconv.Atoi` (ASCII to Integer) or `strconv.ParseInt` for more control.

**Code:**
```go
import "strconv"

s := "123"

// Simple int conversion
i, err := strconv.Atoi(s)
if err != nil {
    // handle error
}

// Control base and bit size (e.g., base 10, int64)
i64, err := strconv.ParseInt(s, 10, 64)
```

---

### Question 14: What happens if you cast a negative integer to a string?

**Answer:**
If you directly convert a negative integer to a string using `string()`, Go interprets the integer as a Unicode code point. Since negative values are invalid code points, it returns the replacement character `\uFFFD` ().

**Correct approach:** Use `strconv.Itoa()` or `fmt.Sprint()`.

**Example:**
```go
val := -42
// Incorrect
fmt.Println(string(val)) // Output: "" (or garbage)

// Correct
fmt.Println(strconv.Itoa(val)) // Output: "-42"
```

---

### Question 15: How do you use Raw String Literals?

**Answer:**
Raw string literals are enclosed in backticks (\`). They do not process escape characters (like `\n`, `\t`) and can span multiple lines.

**Usage:**
- Regex patterns (avoids double escaping backslashes).
- JSON blobs.
- HTML templates.

**Example:**
```go
// Interpreted string
path := "C:\\Users\\Name"

// Raw string
rawPath := `C:\Users\Name`

json := `{
    "name": "John"
}`
```

---

### Question 16: How to compare two strings in a case-insensitive manner?

**Answer:**
Use `strings.EqualFold`. It uses Unicode case folding, which is more robust than simply converting both to lowercase.

**Code:**
```go
s1 := "GoLang"
s2 := "golang"

if strings.EqualFold(s1, s2) {
    fmt.Println("Match!")
}
```

---

### Question 17: How do you reverse a string correctly in Go?

**Answer:**
To reverse a string correctly, you must handle multibyte characters (runes). You cannot just reverse the bytes.

**Algorithm:**
1.  Convert string to `[]rune`.
2.  Swap elements from ends towards the center.
3.  Convert back to `string`.

**Code:**
```go
func reverse(s string) string {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    return string(runes)
}
```

---

### Question 18: How do you check if two strings are anagrams?

**Answer:**
Two strings are anagrams if they contain the same characters with the same frequencies.

**Optimized Approach (using array for ASCII or map for Unicode):**
```go
func isAnagram(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    
    // Assuming ASCII for simplicity (use map[rune]int for Unicode)
    var charCount [26]int
    
    for _, char := range s1 {
        charCount[char-'a']++
    }
    for _, char := range s2 {
        charCount[char-'a']--
        if charCount[char-'a'] < 0 {
            return false
        }
    }
    return true
}
```

---

### Question 19: How to check if a string is a palindrome?

**Answer:**
A palindrome reads the same forwards and backwards.

**Approach:**
Iterate from the start and end simultaneously, comparing runes.

**Code:**
```go
func isPalindrome(s string) bool {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        if runes[i] != runes[n-1-i] {
            return false
        }
    }
    return true
}
```

---

### Question 20: What is string interning and does Go support it?

**Answer:**
String interning is a method of storing only one copy of each distinct string value, which must be immutable.

**Go Support:**
Go does **not** automatically intern all strings (unlike Java's constant pool). However, you can manually implement interning or rely on compiler optimizations for string constants.
- String literals in code *are* effectively interned (they point to the same data in the read-only data section).
- Runtime strings are separate allocations unless manually managed using a "pool" (e.g., a `map[string]string`).

**Manual Interning Example:**
```go
var pool = map[string]string{}

func intern(s string) string {
    if val, ok := pool[s]; ok {
        return val
    }
    pool[s] = s
    return s
}
```

---

### Question 21: How to format strings using fmt.Sprintf?

**Answer:**
`fmt.Sprintf` formats a string according to a format specifier and returns the resulting string instead of printing it.

**Common Verbs:**
- `%s`: String
- `%d`: Integer
- `%f`: Float
- `%v`: Default format
- `%+v`: Struct with field names
- `%T`: Type of value

**Example:**
```go
name := "Alice"
age := 30
// "User Alice is 30 years old"
res := fmt.Sprintf("User %s is %d years old", name, age)
```

---

### Question 22: How to trim whitespace from a string?

**Answer:**
Use `strings.TrimSpace` to remove leading and trailing whitespace.

**Functions:**
- `strings.TrimSpace(s)`: Removes all leading/trailing whitespace.
- `strings.Trim(s, cutset)`: Removes specific characters from ends.
- `strings.TrimLeft(s, cutset)` / `strings.TrimRight(s, cutset)`.

**Example:**
```go
s := "  hello  "
trimmed := strings.TrimSpace(s) // "hello"
```

---

### Question 23: How do you convert a struct to a string?

**Answer:**
You can convert a struct to a string representation using `fmt` verbs or `json.Marshal` for a structured string.

**Methods:**
1.  **`fmt.Sprintf("%+v", s)`:** Prints field names and values.
2.  **JSON Marshaling:** Good for logging or data transfer.

**Code:**
```go
type User struct {
    ID   int
    Name string
}

u := User{ID: 1, Name: "Bob"}

// Method 1:
str1 := fmt.Sprintf("%+v", u) // "{ID:1 Name:Bob}"

// Method 2 (JSON):
b, _ := json.Marshal(u)
str2 := string(b) // "{\"ID\":1,\"Name\":\"Bob\"}"
```

---

### Question 24: What is the difference between Sprint, Sprintln, and Sprintf?

**Answer:**
They are all formatting functions that return a string.

**Differences:**
1.  **`Sprint`:** Concatenates arguments using default formats without adding spaces (unless operands are both non-strings) and no newline.
2.  **`Sprintln`:** Concatenates arguments using default formats, always adding spaces between operands, and appends a newline at the end.
3.  **`Sprintf`:** Formats arguments according to a specific format string.

**Example:**
```go
fmt.Sprint("A", "B")   // "AB"
fmt.Sprintln("A", "B") // "A B\n"
fmt.Sprintf("A%s", "B") // "AB"
```

---

### Question 25: How to parse a URL string in Go?

**Answer:**
Use the `net/url` package, specifically `url.Parse`.

**Key Components:**
- Scheme, Host, Path, RawQuery (Query Params).

**Code:**
```go
import (
    "fmt"
    "net/url"
)

func main() {
    s := "https://example.com/path?name=ferret&color=purple"
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme) // "https"
    fmt.Println(u.Host)   // "example.com"
    
    // Parse query params
    q := u.Query()
    fmt.Println(q.Get("name")) // "ferret"
}
```

---

### Question 26: How to use regular expressions (regexp) in Go?

**Answer:**
Use the `regexp` package. Compile the regex first for performance if it's reused.

**Methods:**
- `MatchString`: Simple boolean check.
- `Compile`: Returns a `Regexp` object for repeated use.
- `FindString`: Returns the first match.

**Code:**
```go
import "regexp"

func main() {
    // 1. Simple check
    matched, _ := regexp.MatchString(`foo.*`, "seafood")
    
    // 2. Pre-compile (better for loops/reuse)
    re := regexp.MustCompile(`\w+`)
    fmt.Println(re.FindString("Hello World")) // "Hello"
}
```

---

### Question 27: How do you validate an email format using regex?

**Answer:**
Use a robust regex pattern with the `regexp` package. Note that a perfect email regex is extremely complex, but standard checks work for most cases.

**Code:**
```go
import "regexp"

func validateEmail(email string) bool {
    // Simple regex for demonstration:
    const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re := regexp.MustCompile(emailRegex)
    return re.MatchString(email)
}

func main() {
    fmt.Println(validateEmail("test@example.com")) // true
    fmt.Println(validateEmail("invalid-email"))    // false
}
```

## ⚫ 28–47: Advanced Slices & Built-in Containers

### Question 28: How do you delete an element from a slice (at index i)?

**Answer:**
There is no built-in "delete" function for slices. You must manipulate the slice using `append` to exclude the element at index `i`.

**Optimized Approaches:**

1.  **Preserve Order (Standard):**
    Append elements *after* `i` to the elements *before* `i`.
    ```go
    s := []int{1, 2, 3, 4, 5}
    i := 2 // Remove 3
    s = append(s[:i], s[i+1:]...) 
    // s is now [1, 2, 4, 5]
    ```

2.  **Fast Delete (Order Not Preserved):**
    Swap the element to be removed with the last element, then truncate the slice. This operates in O(1) time.
    ```go
    s := []int{1, 2, 3, 4, 5}
    i := 2
    s[i] = s[len(s)-1] // Copy last element to index i
    s = s[:len(s)-1]   // Truncate
    // s is now [1, 2, 5, 4] (Order changed)
    ```

---

### Question 29: How do you insert an element at a specific index in a slice?

**Answer:**
You need to expand the slice and shift elements to the right.

**Code:**
```go
func insert(s []int, index int, value int) []int {
    if len(s) == index { // nil or empty slice or after last element
        return append(s, value)
    }
    s = append(s[:index+1], s[index:]...) // Expand slice
    s[index] = value
    return s
}

// Usage
s := []int{1, 2, 3, 4}
s = insert(s, 2, 99) 
// [1, 2, 99, 3, 4]
```
*Note:* This operation can be costly (O(N)) because it involves copying elements.

---

### Question 30: How do you reverse a slice in place?

**Answer:**
Iterate from the start and end indices towards the middle, swapping elements. This avoids allocating a new slice.

**Code:**
```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
```

---

### Question 31: How do you remove duplicate elements from a slice?

**Answer:**
Use a map to track seen elements.

**Code:**
```go
func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{} // Or make([]int, 0, len(intSlice))
    
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}
```
*Complexity:* O(n) time, O(n) space.

---

### Question 32: How do you check if two slices are equal?

**Answer:**
Slices cannot be compared directly using `==` (except against `nil`). You must use `reflect.DeepEqual` (slow) or iterate manually/use `slices.Equal` (Go 1.21+).

**Approaches:**

1.  **Go 1.21+ (Recommended):**
    ```go
    import "slices"
    equal := slices.Equal(s1, s2)
    ```

2.  **Manual Comparison (Fastest for older Go):**
    ```go
    func equal(a, b []int) bool {
        if len(a) != len(b) {
            return false
        }
        for i, v := range a {
            if v != b[i] {
                return false
            }
        }
        return true
    }
    ```

---

### Question 33: How do you perform a deep copy of a slice?

**Answer:**
Use the built-in `copy()` function. Simply assigning `s2 := s1` creates a shallow copy (both share the same backing array).

**Code:**
```go
src := []int{1, 2, 3}

// 1. Allocate destination with correct size
dst := make([]int, len(src))

// 2. Perform copy
numCopied := copy(dst, src)

// dst is now {1, 2, 3} and has its own backing array
```

---

### Question 34: How does copy() work with overlapping slices?

**Answer:**
The built-in `copy()` function handles overlapping slices correctly. It copies data as if the source were first copied to a temporary location.

**Example:**
```go
s := []int{1, 2, 3, 4, 5}
// Shift [2, 3, 4, 5] down to index 0
copy(s[0:], s[1:]) 
// s is now [2, 3, 4, 5, 5]
```

---

### Question 35: How do you merge two sorted slices?

**Answer:**
Use the "two pointers" technique to merge them in O(n+m) time without sorting the combined result from scratch.

**Code:**
```go
func mergeSorted(a, b []int) []int {
    result := make([]int, 0, len(a)+len(b))
    i, j := 0, 0
    
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            result = append(result, a[i])
            i++
        } else {
            result = append(result, b[j])
            j++
        }
    }
    // Append remaining
    result = append(result, a[i:]...)
    result = append(result, b[j:]...)
    
    return result
}
```

---

### Question 36: How do you implement a generic slice filter function?

**Answer:**
With Go 1.18+, you can use generics.

**Code:**
```go
func Filter[T any](s []T, keep func(T) bool) []T {
    var result []T
    for _, v := range s {
        if keep(v) {
            result = append(result, v)
        }
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4}
evens := Filter(nums, func(n int) bool { return n%2 == 0 })
// [2, 4]
```

---

### Question 37: Does Go have a built-in Set data structure?

**Answer:**
No, Go does not have a dedicated `Set` type in the standard library.

**Workaround:**
Sets are idiomatically implemented using maps with empty structs as values: `map[KeyType]struct{}`.
- Using `struct{}` takes 0 bytes of storage for the value, making it memory efficient.

---

### Question 38: How do you implement a Set efficiently in Go?

**Answer:**
Use a map.

**Code:**
```go
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
    return make(Set[T])
}

func (s Set[T]) Add(val T) {
    s[val] = struct{}{}
}

func (s Set[T]) Exists(val T) bool {
    _, ok := s[val]
    return ok
}

func (s Set[T]) Remove(val T) {
    delete(s, val)
}
```

---

### Question 39: What is the container/list package used for?

**Answer:**
It implements a **Doubly Linked List**.

**Usage:**
- When you need frequent insertions/deletions at both ends or the middle without reallocating arrays (like in slices).
- Rarely used in Go compared to slices due to poor cache locality and type assertions (it stores `interface{}`).

**Code:**
```go
import "container/list"

l := list.New()
l.PushBack(1)
l.PushFront("Start")

for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
}
```

---

### Question 40: What is container/ring?

**Answer:**
It implements a **Circular List** (Ring).

**Characteristics:**
- Every element points to the next; the last points back to the first.
- Useful for round-robin scheduling, fixed-size buffers, or recurring operations.

**Code:**
```go
import "container/ring"

r := ring.New(5) // Ring of size 5
for i := 0; i < r.Len(); i++ {
    r.Value = i
    r = r.Next()
}
```

---

### Question 41: How do you use container/heap to implement a Priority Queue?

**Answer:**
The `container/heap` package provides heap operations (push, pop) for any type that implements `heap.Interface`.

**Steps:**
1.  Define a type (e.g., slice).
2.  Implement `sort.Interface` (`Len`, `Less`, `Swap`).
3.  Implement `Push` (append) and `Pop` (remove last).

**Code:**
```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Usage
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
fmt.Println(heap.Pop(h)) // 1 (Min value)
```

---

### Question 42: How does Go's sort.Slice work under the hood?

**Answer:**
`sort.Slice` uses a hybrid sorting algorithm called **pdqsort** (Pattern-Defeating Quicksort).

**Algorithm:**
1.  **Quicksort:** Used for large slices.
2.  **Heap Sort:** Fallback if Quicksort recursion depth becomes too high (prevents O(n²) worst case).
3.  **Shell/Insertion Sort:** Used for small slices (<= 12 elements) for speed.

It is generally O(n log n) and is not stable by default (use `sort.SliceStable` for stability).

---

### Question 43: How do you sort a slice of custom structs?

**Answer:**
Use `sort.Slice` and provide a comparator callback.

**Code:**
```go
type User struct {
    Name string
    Age  int
}

users := []User{
    {"Bob", 30},
    {"Alice", 25},
}

sort.Slice(users, func(i, j int) bool {
    return users[i].Age < users[j].Age // Sort by Age asc
})
```

---

### Question 44: How do you search in a sorted slice?

**Answer:**
Use `sort.Search` (or `sort.SearchInts`, etc.). It uses binary search, providing O(log n) performance.

**Code:**
```go
nums := []int{10, 20, 30, 40}
target := 30

i := sort.Search(len(nums), func(i int) bool {
    return nums[i] >= target 
})

if i < len(nums) && nums[i] == target {
    fmt.Printf("Found %d at index %d\n", target, i)
} else {
    fmt.Println("Not found")
}
```

---

### Question 45: What is the difference between sort.Sort and sort.Slice?

**Answer:**
1.  **`sort.Sort`**: Requires the collection to implement the `sort.Interface` (Len, Less, Swap). It is more verbose but type-safe (no reflection or closures).
2.  **`sort.Slice`**: Introduced in Go 1.8. It accepts a `Less` function closure. It is easier to use but uses reflection internally (though heavily optimized) and panics if the first argument isn't a slice.

---

### Question 46: How do you handle slice out of bounds panics?

**Answer:**
Go panics at runtime if you access an index explicitly outside the bounds.

**Prevention:**
Always check `len(s)` before accessing `s[i]`.

**Recovery:**
If you must handle it (e.g., processing user input scripts), use `defer` and `recover`.

```go
func safeGet(s []int, i int) (val int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("out of bounds: %v", r)
        }
    }()
    return s[i], nil
}
```

---

### Question 47: How efficient is appending to a nil slice?

**Answer:**
It is perfectly safe and efficient. A `nil` slice acts exactly like an empty slice with 0 capacity.

**Initial Allocation:**
When you append to a `nil` slice, Go allocates a new underlying array.
```go
var s []int // nil
s = append(s, 1) // Allocates array, returns slice
```
*Performance Tip:* If you know the size, use `make([]int, 0, capacity)` to avoid the initial allocations during growth.

## ⚪ 48–67: Classic Data Structures & Algorithms

### Question 48: How do you implement a Stack (LIFO) in Go?

**Answer:**
You can just use a slice.

**Code:**
```go
type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
    *s = append(*s, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(*s) == 0 {
        var zero T
        return zero, false
    }
    index := len(*s) - 1
    element := (*s)[index]
    *s = (*s)[:index]
    return element, true
}
```

---

### Question 49: How do you implement a Queue (FIFO) in Go?

**Answer:**
You can use a slice, but removing from the front is O(n) due to shifting. For efficiency, use a linked list or wait for the circular buffer resizing.

**Simple Implementation (Slice):**
```go
type Queue[T any] []T

func (q *Queue[T]) Enqueue(v T) {
    *q = append(*q, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(*q) == 0 {
        var zero T
        return zero, false
    }
    element := (*q)[0]
    *q = (*q)[1:] // This leaks memory if not careful (the array persists)
    return element, true
}
```

---

### Question 50: How do you implement a Linked List in Go?

**Answer:**
Define a node struct with a value and a pointer to the next node.

**Code:**
```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T any] struct {
    Head *Node[T]
}

func (l *LinkedList[T]) Prepend(value T) {
    newNode := &Node[T]{Value: value, Next: l.Head}
    l.Head = newNode
}
```

---

### Question 51: How do you reverse a Linked List?

**Answer:**
Use three pointers: `prev`, `current`, and `next`.

**Code:**
```go
func ReverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        nextTemp := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }
    return prev
}
```

---

### Question 52: How do you detect a cycle in a Linked List?

**Answer:**
Use Floyd’s Cycle-Finding Algorithm (Tortoise and Hare).

**Code:**
```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

---

### Question 53: How do you implement a Binary Search Tree (BST) in Go?

**Answer:**
A tree where the left child is smaller and the right child is larger.

**Code:**
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func (t *TreeNode) Insert(val int) {
    if val < t.Val {
        if t.Left == nil {
            t.Left = &TreeNode{Val: val}
        } else {
            t.Left.Insert(val)
        }
    } else {
        if t.Right == nil {
            t.Right = &TreeNode{Val: val}
        } else {
            t.Right.Insert(val)
        }
    }
}
```

---

### Question 54: How do you implement Tree traversals (Inorder, Preorder, Postorder)?

**Answer:**
Use recursion.

**Code:**
```go
func Inorder(root *TreeNode) {
    if root != nil {
        Inorder(root.Left)
        fmt.Println(root.Val)
        Inorder(root.Right)
    }
}

func Preorder(root *TreeNode) {
    if root != nil {
        fmt.Println(root.Val)
        Preorder(root.Left)
        Preorder(root.Right)
    }
}
```

---

### Question 55: How do you find the max depth of a binary tree?

**Answer:**
Recursive Depth-First Search (DFS).

**Code:**
```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    if left > right {
        return left + 1
    }
    return right + 1
}
```

---

### Question 56: How do you implement a graph using an adjacency list?

**Answer:**
Use a map where the key is the vertex and the value is a slice of neighbors.

**Code:**
```go
type Graph struct {
    adjList map[int][]int
}

func NewGraph() *Graph {
    return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
    g.adjList[u] = append(g.adjList[u], v)
    g.adjList[v] = append(g.adjList[v], u) // For undirected
}
```

---

### Question 57: How do you implement BFS (Breadth-First Search) in Go?

**Answer:**
Use a queue.

**Code:**
```go
func BFS(graph map[int][]int, start int) {
    queue := []int{start}
    visited := make(map[int]bool)
    visited[start] = true

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:] // Dequeue
        fmt.Println(curr)

        for _, neighbor := range graph[curr] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}
```

---

### Question 58: How do you implement DFS (Depth-First Search) in Go?

**Answer:**
Use recursion or a stack.

**Code:**
```go
func DFS(graph map[int][]int, curr int, visited map[int]bool) {
    if visited[curr] {
        return
    }
    visited[curr] = true
    fmt.Println(curr)

    for _, neighbor := range graph[curr] {
        DFS(graph, neighbor, visited)
    }
}
```

---

### Question 59: How do you implement an LRU Cache in Go?

**Answer:**
Use a hash map for O(1) access and a doubly linked list (`container/list`) for O(1) eviction logic.

**Structure:**
```go
import "container/list"

type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}
// 1. Get: If exists, move element to front (most recently used).
// 2. Put: If exists, update & move to front. If new, push front. If over capacity, remove back.
```

---

### Question 60: How do you implement a Trie (Prefix Tree)?

**Answer:**
A tree where each node represents a character.

**Code:**
```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func (t *Trie) Insert(word string) {
    curr := t.root
    for _, ch := range word {
        if curr.children[ch] == nil {
            curr.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        curr = curr.children[ch]
    }
    curr.isEnd = true
}
```

---

### Question 61: How do you implement a Hash Map from scratch (conceptually)?

**Answer:**
Use an array of "buckets." Use a hash function to map keys to an index (`hash(key) % array_size`).

**Collision Handling:**
Use Chaining (Linked List at each bucket) or Open Addressing.

---

### Question 62: How do you handle hash collisions in Go maps?

**Answer:**
Go uses **Chaining** with bucket arrays. Each bucket holds up to 8 key/value pairs. If a bucket overflows, it chains to an overflow bucket.

---

### Question 63: How do you implement a Min-Stack (get min in O(1))?

**Answer:**
Maintain two stacks: one for data, one for minimums.

**Logic:**
- **Push(x):** Push to data stack. If x <= current min, push to min stack.
- **Pop():** Pop from data stack. If value == top of min stack, pop min stack too.
- **GetMin():** Peek top of min stack.

---

### Question 64: How to find the 'nth' Fibonacci number efficiently?

**Answer:**
1.  **Recursion:** O(2^n) - Bad.
2.  **Memoization/Iterative:** O(n).
3.  **Matrix Exponentiation:** O(log n).

**Iterative Example:**
```go
func fib(n int) int {
    if n <= 1 { return n }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```

---

### Question 65: How to validate balanced parentheses in a string?

**Answer:**
Use a Stack.

**Algorithm:**
1.  Push opening brackets `(`, `{`, `[` onto stack.
2.  When encountering a closing bracket, check if stack is empty or top doesn't match. If so, invalid. Pop the matching opener.
3.  At end, stack must be empty.

---

### Question 66: How to find the first non-repeating character in a string?

**Answer:**
Use a map (frequency counter) or an array (for ASCII).

**Algorithm:**
1.  Pass 1: Count frequency of each character.
2.  Pass 2: Iterate string again, check usage count. First with count 1 is the answer.

---

### Question 67: How to implement Merge Sort in Go?

**Answer:**
Divide and Conquer algorithm.

**Code:**
```go
func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2
    left := MergeSort(arr[:mid])
    right := MergeSort(arr[mid:])
    return merge(left, right)
}

func merge(left, right []int) []int {
    res := make([]int, 0, len(left)+len(right))
    // ... merge logic (two pointers) ...
    return res
}
```

##  68�97: Basics of Arrays, Maps, Structs & Loops

### Question 68: What is the difference between an Array and a Slice in Go?

**Answer:**
- **Array:** Fixed size. Value type (assigning copies the whole array). Size is part of the type (e.g., [4]int != [5]int).
- **Slice:** Dynamic size (view into an array). Reference type (assigning copies the header). Much more common.

---

### Question 69: How do you declare an array of fixed size?

**Answer:**
Specify the length in brackets.
`go
var arr [5]int           // Zero-initialized
arr2 := [3]int{1, 2, 3} // Initialized
arr3 := [...]int{4, 5}  // Compiler infers length (2)
`

---

### Question 70: What is the zero value of a slice vs. an array?

**Answer:**
- **Slice:** 
il. It has no underlying array.
- **Array:** An array of zero values of the element type (e.g., [0, 0, 0]). It is never nil.

---

### Question 71: How does the append function work internally?

**Answer:**
1. Checks if the slice has enough capacity (cap - len).
2. If yes, it extends the length and places the element.
3. If no, it allocates a new, larger array (usually double size), copies existing elements, places the new element, and returns the updated slice header.

---

### Question 72: What are the components of a slice header?

**Answer:**
A slice is a struct with 3 fields:
1. **Pointer:** To the first element of the slice in the backing array.
2. **Len:** Number of elements in the slice.
3. **Cap:** Number of elements in the backing array starting from the pointer.
`go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
`

---

### Question 73: How do you initialize a map with values?

**Answer:**
Use the map literal syntax.
`go
m := map[string]int{
    "one": 1,
    "two": 2,
}
`

---

### Question 74: What happens if you read from a nil map?

**Answer:**
It is safe. It returns the zero value of the value type. It does not panic.

---

### Question 75: What happens if you write to a nil map?

**Answer:**
It **panics** (panic: assignment to entry in nil map). You must initialize it with make() or a literal first.

---

### Question 76: How do you check if a key exists in a map?

**Answer:**
Use the "comma-ok" idiom.
`go
val, ok := m["key"]
if ok {
    // exists
}
`

---

### Question 77: Is the iteration order of a map guaranteed?

**Answer:**
**No.** It is essentially random. Go actively randomizes the iteration order to prevent developers from relying on it.

---

### Question 78: How do you delete a key from a map?

**Answer:**
Use the built-in delete function.
`go
delete(m, "key")
`
If the key doesn't exist, it does nothing (no-op).

---

### Question 79: Can you use a slice as a map key? Why or why not?

**Answer:**
**No.** Map keys must be **comparable** (support ==). Slices do not support ==. Use arrays or structs (if fields are comparable) instead.

---

### Question 80: How do you define a struct in Go?

**Answer:**
`go
type Person struct {
    Name string
    Age  int
}
`

---

### Question 81: What represent anonymous structs?

**Answer:**
Structs defined without a type name, usually for one-off use (e.g., inside a function or for JSON unmarshaling).
`go
p := struct {
    Name string
}{
    Name: "John",
}
`

---

### Question 82: How do you access fields of a struct?

**Answer:**
Use the dot operator ..
`go
p.Name = "Alice"
fmt.Println(p.Age)
`
This works for both struct values and pointers to structs (automatic dereferencing).

---

### Question 83: What are promoted fields in embedded structs?

**Answer:**
When a struct is embedded anonymously, its fields are accessible directly on the outer struct.
`go
type Address struct { City string }
type User struct {
    Address // Embedded
}
u := User{Address: Address{City: "NY"}}
fmt.Println(u.City) // Promoted field
`

---

### Question 84: How to compare two structs?

**Answer:**
- By default, you can use == if **all fields are comparable**.
- If fields contain non-comparable types (slices, maps), you must use eflect.DeepEqual (slow) or write a custom comparison method.

---

### Question 85: What is the only loop construct in Go?

**Answer:**
The or loop. There is no while or do-while.

---

### Question 86: How do you simulate a while-loop in Go?

**Answer:**
Use or with just a condition.
`go
for condition {
    // ...
}
`

---

### Question 87: How do you create an infinite loop?

**Answer:**
Use or with no condition (or select{} if waiting).
`go
for {
    // ...
}
`

---

### Question 88: How does the range keyword work?

**Answer:**
It iterates over elements in a variety of data structures (slice, array, map, string, channel). It returns one or two values depending on context.

---

### Question 89: What are the two values returned by range for a slice?

**Answer:**
1. **Index** (int)
2. **Value** (copy of the element at that index)

---

### Question 90: What are the two values returned by range for a map?

**Answer:**
1. **Key**
2. **Value**

---

### Question 91: How to ignore index or value in a range loop?

**Answer:**
Use the blank identifier _.
`go
for _, val := range slice { ... } // Ignore index
for idx, _ := range slice { ... } // Ignore value (or just "for idx := ...")
`

---

### Question 92: What is the problem with capturing loop variables in closures?

**Answer:**
Before Go 1.22, loop variables were reused per iteration. Capturing them in a closure (goroutine) would capture the **reference**, so strictly all closures would see the *last* value.
*Fix:*  := v (shadowing) inside the loop.
*Note:* Fixed in **Go 1.22+** (loop variables are per-iteration).

---

### Question 93: How do you break out of a nested loop with execution labels?

**Answer:**
Define a label before the outer loop and reak to it.
`go
OuterLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if condition {
            break OuterLoop
        }
    }
}
`

---

### Question 94: What is the goto statement and when should you use it?

**Answer:**
It jumps execution to a label.
*Use Cases (Rare):* Breaking out of deep loops, centralizing error handling/cleanup in a function (though defer is preferred).

---

### Question 95: How do you iterate over a channel?

**Answer:**
Use ange. It keeps reading until the channel is closed.
`go
for msg := range ch {
    fmt.Println(msg)
}
`
If the channel is never closed, this blocks forever (deadlock).

---

### Question 96: What is the difference between break and continue?

**Answer:**
- **reak**: Exits the innermost loop (or switch/select) immediately.
- **continue**: Skips the rest of the current iteration and starts the next one.

---

### Question 97: How do you loop through a string (byte vs rune)?

**Answer:**
- **Bytes:** Standard or loop with index.
  `go
  for i := 0; i < len(s); i++ { fmt.Println(s[i]) }
  `
- **Runes (Characters):** ange loop.
  `go
  for _, r := range s { fmt.Println(r) }
  `

##  98�117: Time & Date Handling

### Question 98: How do you get the current time in Go?

**Answer:**
Use 	ime.Now().
`go
now := time.Now()
`

---

### Question 99: What does time.Time represent?

**Answer:**
It represents an instant in time with nanosecond precision. It includes both the wall clock time and the location (timezone).

---

### Question 100: How do you format a date string in Go (e.g., YYYY-MM-DD)?

**Answer:**
Use Format with the reference layout Mon Jan 2 15:04:05 MST 2006.
`go
t := time.Now()
fmt.Println(t.Format("2006-01-02"))
`

---

### Question 101: Why is the reference date "Mon Jan 2 15:04:05 MST 2006"?

**Answer:**
It corresponds to the sequence 1, 2, 3, 4, 5, 6, 7.
- Month: 1 (Jan)
- Day: 2
- Hour: 3 (15)
- Minute: 4
- Second: 5
- Year: 2006
- Timezone: -7 (MST is GMT-7)

---

### Question 102: How do you parse a string into a time.Time object?

**Answer:**
Use 	ime.Parse.
`go
layout := "2006-01-02"
str := "2023-10-25"
t, err := time.Parse(layout, str)
`

---

### Question 103: How do you calculate the difference between two times?

**Answer:**
Use the Sub method, which returns a 	ime.Duration.
`go
diff := t2.Sub(t1)
`

---

### Question 104: What is time.Duration?

**Answer:**
It is an int64 representing the elapsed time in **nanoseconds**.
Useful constants: 	ime.Second, 	ime.Minute, 	ime.Hour.

---

### Question 105: How do you add or subtract time from a date?

**Answer:**
- **Add:** 	.Add(time.Hour * 2)
- **Subtract:** 	.Add(-time.Hour * 2)
- **Add Date:** 	.AddDate(years, months, days)

---

### Question 106: How do you convert a Unix timestamp to time.Time?

**Answer:**
Use 	ime.Unix.
`go
t := time.Unix(1698256000, 0) // seconds, nanoseconds
`

---

### Question 107: How do you get the Unix timestamp from a time.Time object?

**Answer:**
`go
sec := t.Unix()
nano := t.UnixNano()
`

---

### Question 108: What is the difference between time.NewTicker and time.NewTimer?

**Answer:**
- **Ticker:** Fires repeatedly at intervals. (Must be stopped with .Stop()).
- **Timer:** Fires once after a duration.

---

### Question 109: How do you implement a simple timeout using time.After?

**Answer:**
Use inside a select.
`go
select {
case res := <-ch:
    fmt.Println(res)
case <-time.After(2 * time.Second):
    fmt.Println("Timed out")
}
`

---

### Question 110: How do you compare if one time is before or after another?

**Answer:**
Use methods for clarity (handles monotonic clock correctly).
`go
if t1.After(t2) { ... }
if t1.Before(t2) { ... }
if t1.Equal(t2) { ... }
`

---

### Question 111: How do you handle time zones in Go?

**Answer:**
Use 	ime.Location.
`go
loc, _ := time.LoadLocation("America/New_York")
t := time.Now().In(loc)
`

---

### Question 112: How to load a specific location (Timezone)?

**Answer:**
	ime.LoadLocation("Region/City"). It looks up the system's zoneinfo database or the ZONEINFO env var.

---

### Question 113: What is time.Sleep and how does it work?

**Answer:**
	ime.Sleep(d) pauses the current goroutine for at least the duration d. It yields the processor to other goroutines.

---

### Question 114: How to measure execution time of a function?

**Answer:**
Use defer and 	ime.Since.
`go
func measure() {
    start := time.Now()
    defer func() {
        fmt.Println("Took:", time.Since(start))
    }()
    // ... work ...
}
`

---

### Question 115: How strictly does time.Parse validate input?

**Answer:**
Strictly. If the string format doesn't match the layout exactly, it returns an error.

---

### Question 116: How do you reset a timer?

**Answer:**
Use 	.Reset(d).
*Important:* Ensure the channel is drained if you haven't received from it yet, to avoid race conditions.
`go
if !timer.Stop() {
    <-timer.C
}
timer.Reset(newDuration)
`

---

### Question 117: How do you serialize time.Time to JSON?

**Answer:**
	ime.Time implements json.Marshaler interface. It serializes to an RFC 3339 formatted string (e.g., "2023-10-25T08:00:00Z").
