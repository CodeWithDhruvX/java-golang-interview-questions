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

### Question 118: What is the difference between passing by value and passing by pointer in Go?

**Answer:**
In Go, everything is passed by value. When you pass a variable by value, a copy of the data is created and passed to the function. Modifying the copy inside the function does not affect the original variable.

When you pass a pointer (which is also passed by value), you pass a copy of the memory address. This allows the function to modify the value at that memory address, effectively modifying the original variable.

```go
func modifyVal(x int) { x = 10 }
func modifyPtr(x *int) { *x = 10 }

func main() {
    a := 5
    modifyVal(a) // a is still 5
    modifyPtr(&a) // a becomes 10
}
```

---

### Question 119: How do you define a method on a struct type?

**Answer:**
A method is a function with a special receiver argument. You define it by specifying the receiver type between the `func` keyword and the method name.

```go
type User struct {
    Name string
}

// Method on User struct
func (u User) Greet() string {
    return "Hello " + u.Name
}
```

---

### Question 120: Can you define methods on non-struct types (e.g., `type MyInt int`)?

**Answer:**
Yes, you can define methods on any type you define in the same package, except for pointer types or interfaces. This includes aliases for built-in types like `int`, `string`, slice, map, etc.

```go
type MyInt int

func (m MyInt) IsEven() bool {
    return m%2 == 0
}
```

---

### Question 121: What is the difference between a Value Receiver and a Pointer Receiver?

**Answer:**
*   **Value Receiver (`func (t T) ...`):** The method operates on a copy of the value. Changes made to the receiver inside the method are not visible to the caller. It is generally safer for concurrency but can be expensive if the struct is large.
*   **Pointer Receiver (`func (t *T) ...`):** The method operates on the actual value (via its address). Changes are visible to the caller. It avoids copying the value.

---

### Question 122: When should you use a Pointer Receiver?

**Answer:**
Use a pointer receiver when:
1.  **Modifying state:** The method needs to modify the receiver.
2.  **Performance:** The struct is large, and copying it is expensive.
3.  **Consistency:** If some methods of the type have pointer receivers, usually all methods should have pointer receivers for consistency.
4.  **Sync primitives:** If the struct contains a Mutex or similar synchronization primitive, it must not be copied.

---

### Question 123: Can you call a pointer receiver method on a value variable?

**Answer:**
Yes, Go automatically takes the address of the value if it is addressable.

```go
type Data struct {
    Val int
}
func (d *Data) Set(v int) { d.Val = v }

d := Data{}
d.Set(5) // Compiler interprets as (&d).Set(5)
```

---

### Question 124: Can you call a value receiver method on a pointer variable?

**Answer:**
Yes, Go automatically dereferences the pointer to get the value.

```go
func (d Data) Get() int { return d.Val }

ptr := &Data{Val: 10}
val := ptr.Get() // Compiler interprets as (*ptr).Get()
```
*Note: This works only if the pointer is not nil; calling a value method on a nil pointer causes a panic.*

---

### Question 125: What is a "Method Set" in Go?

**Answer:**
The method set determines which interfaces a type implements.
*   The method set of a type `T` consists of all methods declared with receiver `T`.
*   The method set of a type `*T` consists of all methods declared with receiver `*T` **AND** receiver `T`.

---

### Question 126: What methods belong to the method set of type `T` vs `*T`?

**Answer:**
*   **T:** Contains only methods with value receivers (`(t T)`).
*   **T:** Contains methods with pointer receivers (`(t *T)`) AND methods with value receivers (due to automatic dereferencing capabilities, but strictly speaking, the interface satisfaction rule says *T has both).

This distinction is crucial for interface implementation. If an interface requires a method defined with a pointer receiver, only `*T` implements that interface, not `T`.

---

### Question 127: How does `new()` differ from `make()` exactly?

**Answer:**
*   **`new(T)`**: Allocates zeroed storage for a new item of type `T` and returns its address, a value of type `*T`. It applies to value types like structs, ints, arrays.
*   **`make(T, args)`**: Creates an initialized (not zeroed) value of type `T` (not `*T`). It is used **only** for slices, maps, and channels. These types require internal data structure initialization (e.g., allocating an underlying array for a slice) before use.

---

### Question 128: What types can be created using `make()`?

**Answer:**
`make()` is restricted to three built-in reference types:
1.  **Slices** (`make([]int, len, cap)`)
2.  **Maps** (`make(map[string]int, hint)`)
3.  **Channels** (`make(chan int, bufSize)`)

---

### Question 129: What is the return value of `new(T)`?

**Answer:**
`new(T)` returns a **pointer** to the newly allocated zero value of type `T`. So, the return type is `*T`.

---

### Question 130: How are interfaces represented in memory (itab and data)?

**Answer:**
An interface value is conceptually a pair of words (pointers):
1.  **itab (Interface Table) pointer:** Points to a table containing type information about the concrete type and a list of function pointers for the methods required by the interface.
2.  **data pointer:** Points to the actual copy of the concrete data (if it's a value type) or the pointer to the data (if it's a pointer type) that implements the interface.

If an interface is nil, both words are zero (nil).

---

### Question 131: What is a Type Assertion and how is it used?

**Answer:**
A type assertion provides access to the underlying concrete value of an interface.
Syntax: `t := i.(T)`

*   If `i` holds a `T`, `t` will be the underlying value.
*   If `i` does not hold a `T`, it panics.

To avoid panic, use the comma-ok idiom:
```go
t, ok := i.(T)
if ok {
    // success
} else {
    // failure, no panic
}
```

---

### Question 132: What is a Type Switch?

**Answer:**
A type switch is a construct that permits several type assertions in series. It looks like a regular switch statement but uses `.(type)`.

```go
switch v := i.(type) {
case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
default:
    fmt.Printf("I don't know about type %T!\n", v)
}
```

---

### Question 133: How do you check if an interface value is `nil`?

**Answer:**
An interface value is `nil` only if both its value and dynamic type are nil. You can check:
```go
if i == nil { ... }
```
However, beware of the "nil pointer inside interface" trap.

---

### Question 134: Can an interface holding a nil concrete pointer be nil?

**Answer:**
**No.** If you store a nil pointer of a concrete type (e.g., `*int(nil)`) in an interface, the interface variable itself is **not nil**.
It has a concrete type (`*int`) but a nil value.

```go
var p *int = nil
var i interface{} = p
fmt.Println(i == nil) // False
```

---

### Question 135: What are the methods required to implement `sort.Interface`?

**Answer:**
To sort a collection using `sort.Sort`, the collection must implement `sort.Interface` which requires three methods:
1.  `Len() int`: The number of elements in the collection.
2.  `Less(i, j int) bool`: Reports whether the element at index `i` should sort before the element at index `j`.
3.  `Swap(i, j int)`: Swaps the elements with indexes `i` and `j`.

---

### Question 136: How do you get the capacity (`cap`) and length (`len`) of a channel?

**Answer:**
You use the built-in functions `cap()` and `len()`.
*   `cap(ch)`: Returns the buffer size (capacity) of the channel.
*   `len(ch)`: Returns the number of elements currently queued in the channel buffer.

---

### Question 137: What happens if you send to a closed channel?

**Answer:**
Sending to a closed channel causes a **runtime panic**.

```go
ch := make(chan int)
close(ch)
ch <- 1 // Panic: send on closed channel
```

---

### Question 138: What happens if you receive from a closed channel?

**Answer:**
Receiving from a closed channel returns the **zero value** of the channel's element type immediately, without blocking. It does not panic.
You can detect if it's closed using the second return value: `v, ok := <-ch`. If `ok` is false, the channel is closed.

---

### Question 139: How do you check if a channel is closed ensuring no panic?

**Answer:**
You cannot check if a channel is closed without receiving from it (or trying to send, which might panic). The idiomatic way to "check" is to receive with the comma-ok idiom.
Since sending to a closed channel panics, the sender should be the one responsible for closing it, so they naturally know it is closed.

---

### Question 140: What is the zero value of a function type?

**Answer:**
The zero value of a function type is `nil`. Calling a nil function value causes a panic.

---

### Question 141: Can functions be used as map keys?

**Answer:**
**No**, functions are not comparable in Go (you cannot check `func1 == func2`), so they cannot be used as map keys.

---

### Question 142: How do anonymous functions (closures) capture variables?

**Answer:**
Closures capture variables by **reference**, not by value. If the captured variable is modified outside the closure (or inside), the change is reflected everywhere.
Loop variable capture is a common pitfall (fixed in Go 1.22), where all strict closures in a loop used to capture the same loop variable instance.

---

### Question 143: What is a variadic function and how do you pass a slice to it?

**Answer:**
A variadic function accepts a variable number of arguments (e.g., `func sum(nums ...int)`).
To pass a slice to a variadic function, you must unpack it using the `...` suffix.

```go
nums := []int{1, 2, 3}
sum(nums...)
```

---

### Question 144: How does `defer` work with method evaluation (arguments vs execution)?

**Answer:**
When a `defer` statement is executed, the **arguments** to the deferred function are evaluated immediately (at the time of the defer call). However, the **function body** is executed only when the surrounding function returns.

```go
func example() {
    i := 0
    defer fmt.Println(i) // i evaluates to 0 here
    i++
    return
} // prints 0
```

---

### Question 145: What is `unsafe.Pointer` and when is it used?

**Answer:**
`unsafe.Pointer` is a special pointer type that bypasses Go's type safety.
*   It can convert any pointer type to `unsafe.Pointer` and vice versa.
*   It allows pointer arithmetic when converted to `uintptr`.
*   **Usage:** It is used for interacting with C code (cgo), low-level memory manipulation, or system calls where type safety must be explicitly overridden. It is dangerous and should be avoided unless absolutely necessary.

---

### Question 146: How does `uintptr` differ from `*int`?

**Answer:**
*   `*int`: A standard pointer to an integer. The Garbage Collector tracks it and knows it points to memory.
*   `uintptr`: An integer type properly sized to hold a pointer address. However, for the GC, it is just an integer. The GC does **not** track the memory it points to. If the object pointed to by `uintptr` is moved or collected, the `uintptr` becomes invalid. It is primarily used with `unsafe.Pointer` for arithmetic.

---

### Question 147: How do you manually manage memory alignment (padding)?

**Answer:**
Go structs are aligned based on the field with the largest alignment requirement. To optimize memory (reduce padding), you should order struct fields from largest (in bytes) to smallest.
For manual alignment, you can use blank fields `_ [N]byte` to force padding, but usually, simple reordering is sufficient.

```go
// Improved alignment
type Good struct {
    a int64 // 8 bytes
    b int32 // 4 bytes
    c int32 // 4 bytes
}
```

---

### Question 148: How do you define a generic Stack data structure in Go?

**Answer:**
You use type parameters (Generics, introduced in Go 1.18).

```go
type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(v T) {
    s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    index := len(s.elements) - 1
    val := s.elements[index]
    s.elements = s.elements[:index]
    return val, true
}
```

---

### Question 149: How do you use the `comparable` constraint in generic maps?

**Answer:**
The keys of a map must be comparable (support `==` and `!=`). In generic functions or types involving maps as keys, you must use the `comparable` constraint for the key type.

```go
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
```

---

### Question 150: Can you use a generic type as a receiver for a method?

**Answer:**
Yes, methods on a generic type simply declare the type parameter in the receiver.

```go
type Box[T any] struct {
    val T
}

// Method on generic type Box
func (b *Box[T]) Set(v T) {
    b.val = v
}
```
*Note: You cannot introduce **new** type parameters in a method declaration; they must come from the type definition.*

---

### Question 151: How to implement a generic Linked List?

**Answer:**

```go
type Node[T any] struct {
    Val  T
    Next *Node[T]
}

type LinkedList[T any] struct {
    Head *Node[T]
}

func (ll *LinkedList[T]) Push(v T) {
    newNode := &Node[T]{Val: v, Next: ll.Head}
    ll.Head = newNode
}
```

---

### Question 152: What is `sync.Map` and when should you use it over a regular map?

**Answer:**
`sync.Map` is a specialized concurrent-safe map implementation provided by the standard library. You should use it **only** in specific cases:
1.  When an entry is written once but read many times (append-only caches).
2.  When multiple goroutines read, write, and overwrite entries for disjoint sets of keys.

For standard use cases (frequent reads and writes to the same keys), a regular `map` protected by `sync.RWMutex` is often faster and type-safe.

---

### Question 153: How does `sync.Pool` work and what is it used for?

**Answer:**
`sync.Pool` caches allocated but unused objects for later reuse, relieving pressure on the garbage collector. It is thread-safe.
It is commonly used for frequently allocated short-lived objects like buffers (`bytes.Buffer`).
*Items in the pool may be deallocated by the runtime at any time (usually during GC).*

---

### Question 154: What is the downside of using `sync.Map` for all concurrent map needs?

**Answer:**
1.  **Type Safety:** `sync.Map` stores `interface{}`, so you lose compile-time type safety and incur the overhead of type assertions and interface wrapping/unwrapping.
2.  **Performance:** For general R/W workloads, `sync.Map` can be slower than a `map` + `RWMutex` due to its complex internal logic aimed at specialized disjoint access patterns.

---

### Question 155: How do you use `sync.WaitGroup` to wait for multiple goroutines?

**Answer:**
A `WaitGroup` waits for a collection of goroutines to finish.
1.  `wg.Add(n)` increments the counter.
2.  `wg.Done()` decrements the counter (call in the goroutine).
3.  `wg.Wait()` blocks until the counter is zero.

```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // work
    }()
}
wg.Wait()
```

---

### Question 156: What is `sync.Cond` and how do you use it for signaling?

**Answer:**
`sync.Cond` implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.
It requires a `Locker` (usually a `Mutex`).
*   `Wait()`: Unlocks the mutex and suspends execution until signaled.
*   `Signal()`: Wakes one waiting goroutine.
*   `Broadcast()`: Wakes all waiting goroutines.

```go
c := sync.NewCond(&sync.Mutex{})
// Wait
c.L.Lock()
for !condition {
    c.Wait()
}
// act on condition
c.L.Unlock()
```

---

### Question 157: How does `sync.Once` ensure a function is called exactly once?

**Answer:**
`sync.Once` uses an internal atomic counter (and a mutex for the slow path) to track if the function has been executed.
If `once.Do(f)` is called multiple times, only the first call executes `f`, even if called concurrently.

```go
var once sync.Once
once.Do(func() { fmt.Println("Only once") })
```

---

### Question 158: What is the difference between `sync.Mutex` and `sync.RWMutex`?

**Answer:**
*   **`sync.Mutex` (Mutual Exclusion):** Allows only one goroutine to access the critical section at a time.
*   **`sync.RWMutex` (Read-Write Mutual Exclusion):** Allows multiple readers **OR** one writer.
    *   `RLock()`: Read lock (multiple allowed).
    *   `Lock()`: Write lock (exclusive).
    *   Use `RWMutex` when you have many reads and few writes to improve concurrency.

---

### Question 159: What defines an "Atomic Operation" in Go (`sync/atomic`)?

**Answer:**
An atomic operation is indivisible; it completes entirely or not at all, without interference from other goroutines. The `sync/atomic` package provides low-level atomic memory primitives (Swap, CAS, Add, Load, Store) useful for implementing synchronization algorithms without mutexes.

---

### Question 160: How do you strictly type atomic values using `atomic.Pointer[T]` (Go 1.19+)?

**Answer:**
`atomic.Pointer[T]` provides a type-safe atomic pointer. You don't need `unsafe.Pointer`.

```go
var ptr atomic.Pointer[User]
u := &User{Name: "Alice"}
ptr.Store(u)
readUser := ptr.Load()
```

---

### Question 161: How to implement a semaphore using a buffered channel?

**Answer:**
A buffered channel of capacity `N` can act as a semaphore allowing `N` concurrent operations.

```go
sem := make(chan struct{}, 5) // Max 5 concurrent

func work() {
    sem <- struct{}{} // Acquire token
    defer func() { <-sem }() // Release token
    // Critical work
}
```

---

### Question 162: What is `errgroup.Group` and how does it help with structured concurrency?

**Answer:**
`errgroup.Group` (from `golang.org/x/sync/errgroup`) provides synchronization, error propagation, and Context cancellation for groups of goroutines working on a common task.
*   If any goroutine returns an error, the Group's `Wait()` returns that error.
*   It simplifies managing multiple goroutines where failure of one should cancel the others (using `WithContext`).

---

### Question 163: How to implement a Thread-Safe Queue using Mutex?

**Answer:**
Wrap a slice (or list) and a Mutex in a struct. Lock the mutex for every Push and Pop operation.

```go
type SafeQueue struct {
    q   []int
    mu  sync.Mutex
}

func (s *SafeQueue) Push(v int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.q = append(s.q, v)
}

func (s *SafeQueue) Pop() int {
    s.mu.Lock()
    defer s.mu.Unlock()
    // handle empty check
    v := s.q[0]
    s.q = s.q[1:]
    return v
}
```

---

### Question 164: How to implement a Worker Pool using channels?

**Answer:**
Create a pool of worker goroutines that listen on a "jobs" channel.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    // Collect results
}
```

---

### Question 165: What is the `context.Context` structure used for?

**Answer:**
`context.Context` carries:
1.  **Deadlines/Timeouts:** Signals when work should stop.
2.  **Cancellation Signals:** Propagates cancellation down the call graph.
3.  **Request-scoped values:** Passes data (like RequestID, UserAuth) through API boundaries and between processes.

It is essential for controlling the lifecycle of concurrent operations.

---

### Question 166: How to implement a generic "Set" using a map [T]struct{}?

**Answer:**
Since Go doesn't have a built-in Set, use a map where the key is the element and the value is an empty struct (consuming 0 bytes).

```go
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
    return make(Set[T])
}

func (s Set[T]) Add(v T) {
    s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
    _, exists := s[v]
    return exists
}
```

---

### Question 167: How to implement a generic "Option Pattern" for struct initialization?

**Answer:**
The Functional Option Pattern uses functions to modify a configuration struct. Generics allow one set of logic for strict typing if needed, but often the config struct itself is specific.

```go
type Server struct {
    Port int
    Host string
}

type Option func(*Server)

func WithPort(p int) Option {
    return func(s *Server) {
        s.Port = p
    }
}

func NewServer(opts ...Option) *Server {
    s := &Server{Host: "localhost", Port: 8080} // defaults
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```
