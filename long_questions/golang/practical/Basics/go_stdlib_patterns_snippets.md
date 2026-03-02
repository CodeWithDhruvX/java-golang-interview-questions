# Go — stdlib Patterns Snippets: flag · regexp · runtime · filepath · unicode · embed · net/url

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Topics**: `flag` package · `regexp` basics · `runtime` package · `path/filepath` · `unicode/utf8` · `embed` · `net/url` · `os/exec` · `sort` advanced · `math/rand`

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** flag Package (Q1–Q12)
- [ ] **Section 2:** regexp Basics (Q13–Q24)
- [ ] **Section 3:** runtime Package (Q25–Q34)
- [ ] **Section 4:** path/filepath Package (Q35–Q44)
- [ ] **Section 5:** unicode/utf8 Package (Q45–Q54)
- [ ] **Section 6:** embed Package (Q55–Q60)
- [ ] **Section 7:** net/url & os/exec (Q61–Q72)

> 🔖 **Last read:** <!-- e.g. Q24 · Section 2 done -->

---

## Section 1: flag Package (Q1–Q12)

### 1. Basic Flag Definition and Parse
**Q: What is the output when run with `go run main.go -name=Alice -count=3`?**
```go
package main
import (
    "flag"
    "fmt"
)

func main() {
    name  := flag.String("name", "World", "a name to greet")
    count := flag.Int("count", 1, "number of greetings")
    flag.Parse()

    for i := 0; i < *count; i++ {
        fmt.Printf("Hello, %s!\n", *name)
    }
}
```
**A:**
```
Hello, Alice!
Hello, Alice!
Hello, Alice!
```
`flag.String`/`flag.Int` return *pointers*. Always dereference after `flag.Parse()`.

---

### 2. Flag Default Values
**Q: What happens when run with no arguments?**
```go
name  := flag.String("name", "World", "usage")
count := flag.Int("count", 1, "usage")
flag.Parse()
fmt.Printf("Hello, %s! (x%d)\n", *name, *count)
```
**A:** `Hello, World! (x1)`. If a flag is not provided, its default value is used.

---

### 3. flag.BoolVar — Bind to Existing Variable
**Q: What is the output with `-verbose`?**
```go
package main
import (
    "flag"
    "fmt"
)

var verbose bool

func main() {
    flag.BoolVar(&verbose, "verbose", false, "enable verbose output")
    flag.Parse()
    fmt.Println("verbose:", verbose)
}
```
**A:** `verbose: true`. `BoolVar` (and `StringVar`, `IntVar`) binds a flag to an *existing variable* instead of returning a pointer — useful for package-level vars.

---

### 4. flag.Args — Non-Flag Arguments
**Q: What is the output with `go run main.go -n=2 file1.txt file2.txt`?**
```go
flag.Int("n", 1, "")
flag.Parse()
fmt.Println(flag.Args())
fmt.Println(flag.NArg())
```
**A:**
```
[file1.txt file2.txt]
2
```
`flag.Args()` returns the non-flag positional arguments after parsing. `flag.NArg()` returns their count.

---

### 5. flag.Visit — Inspect Only Set Flags
**Q: What is the output with `-name=Go`?**
```go
flag.String("name", "", "")
flag.String("debug", "", "")
flag.Parse()
flag.Visit(func(f *flag.Flag) {
    fmt.Println("set:", f.Name, "=", f.Value)
})
```
**A:** `set: name = Go`. `flag.Visit` only iterates flags that were *explicitly set* by the user. `flag.VisitAll` would include `debug` with its default.

---

### 6. Custom Flag Type — Implementing flag.Value
**Q: What is the output?**
```go
package main
import (
    "flag"
    "fmt"
    "strings"
)

type csvFlag []string

func (c *csvFlag) String() string  { return strings.Join(*c, ",") }
func (c *csvFlag) Set(v string) error {
    *c = strings.Split(v, ",")
    return nil
}

func main() {
    var tags csvFlag
    flag.Var(&tags, "tags", "comma-separated tags")
    flag.CommandLine.Parse([]string{"-tags=go,web,api"})
    fmt.Println(tags)
}
```
**A:** `[go web api]`. Implement `flag.Value` (two methods: `String()`, `Set(string) error`) for custom flag types.

---

### 7. flag.FlagSet — Sub-commands
**Q: What is the pattern?**
```go
package main
import (
    "flag"
    "fmt"
    "os"
)

func main() {
    serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
    port := serveCmd.Int("port", 8080, "port to listen on")

    if len(os.Args) < 2 {
        fmt.Println("expected subcommand")
        os.Exit(1)
    }
    switch os.Args[1] {
    case "serve":
        serveCmd.Parse(os.Args[2:])
        fmt.Printf("serving on :%d\n", *port)
    }
}
```
**A:** `go run main.go serve -port=9090` → `serving on :9090`. `flag.FlagSet` enables sub-command patterns (like `git commit`, `git push`).

---

### 8. flag.Usage — Custom Help Text
**Q: What does overriding flag.Usage do?**
```go
flag.Usage = func() {
    fmt.Fprintf(flag.CommandLine.Output(), "Usage: myapp [options]\n")
    flag.PrintDefaults()
}
```
**A:** When an invalid flag is passed or `-help` is requested, `flag.Usage` is called. Override it to show custom help text before the auto-generated defaults.

---

### 9. `-h` / `-help` — Auto-Generated
**Q: What happens when you run with `-help`?**
```go
flag.String("name", "World", "Name to greet")
flag.Parse()
```
**A:** The flag package automatically handles `-h`/`-help`, printing:
```
  -name string
        Name to greet (default "World")
```
Then exits with status 2.

---

### 10. flag.Duration — Parsing time.Duration
**Q: What is the output with `-timeout=5s`?**
```go
timeout := flag.Duration("timeout", 30*time.Second, "request timeout")
flag.Parse()
fmt.Println(*timeout)
fmt.Println((*timeout).Seconds())
```
**A:**
```
5s
5
```
`flag.Duration` parses strings like `"5s"`, `"2m30s"`, `"1h"` into `time.Duration`. No manual parsing needed.

---

### 11. Bool Flag — Difference Between `-verbose` and `-verbose=true`
**Q: Are these equivalent?**
```go
flag.Bool("verbose", false, "")
```
**A:** **Yes.** `-verbose` alone sets the bool to `true`. `-verbose=true` and `-verbose=false` are also valid. `-verbose false` is **NOT** valid — `false` would be interpreted as a positional argument.

---

### 12. flag.Parse Must Be Called Before Accessing Values
**Q: What is the bug?**
```go
name := flag.String("name", "World", "")
fmt.Println(*name) // BUG: flag.Parse() not called yet
flag.Parse()
```
**A:** `World` (the default). Flags haven't been parsed from `os.Args` yet, so `*name` still holds the default. Always call `flag.Parse()` *before* accessing flag values.

---

## Section 2: regexp Basics (Q13–Q24)

### 13. regexp.MatchString — Quick Match
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "regexp"
)

func main() {
    matched, err := regexp.MatchString(`\d+`, "abc123")
    fmt.Println(matched, err)

    matched2, _ := regexp.MatchString(`^\d+$`, "abc123")
    fmt.Println(matched2)
}
```
**A:**
```
true <nil>
false
```
`\d+` matches because `123` is in the string. `^\d+$` requires the *entire* string to be digits.

---

### 14. regexp.Compile vs regexp.MustCompile
**Q: When should you use each?**
```go
// MustCompile — panics on invalid pattern; use for package-level regexps
var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Compile — returns error; use for user-supplied patterns
re, err := regexp.Compile(userInput)
if err != nil { return fmt.Errorf("invalid pattern: %w", err) }
```
**A:** `MustCompile` is for compile-time-known patterns (panics on bad pattern — safe since it's your code). `Compile` is for runtime patterns from user input — handle the error.

---

### 15. Compiled Regexp — FindString
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`\d+`)
    fmt.Println(re.FindString("abc 42 def 99"))
    fmt.Println(re.FindAllString("abc 42 def 99", -1))
}
```
**A:**
```
42
[42 99]
```
`FindString` returns the first match. `FindAllString(s, -1)` returns all matches (`-1` = no limit).

---

### 16. FindStringIndex — Match Position
**Q: What is the output?**
```go
re := regexp.MustCompile(`\d+`)
loc := re.FindStringIndex("price: 42 items")
fmt.Println(loc)
fmt.Println("price: 42 items"[loc[0]:loc[1]])
```
**A:**
```
[7 9]
42
```
`FindStringIndex` returns `[start, end]` byte indices of the first match.

---

### 17. Subgroups — FindStringSubmatch
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
    m := re.FindStringSubmatch("user@example.com")
    fmt.Println(m)
}
```
**A:** `[user@example.com user example com]`. Index 0 = full match, indices 1+ = capture groups.

---

### 18. ReplaceAllString
**Q: What is the output?**
```go
re := regexp.MustCompile(`\s+`)
fmt.Println(re.ReplaceAllString("Hello   World\t!", " "))
```
**A:** `Hello World !`. `ReplaceAllString` replaces all matches of the pattern with the replacement string. `\s+` matches one or more whitespace characters.

---

### 19. ReplaceAllStringFunc — Dynamic Replacement
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    re := regexp.MustCompile(`\b\w+\b`)
    result := re.ReplaceAllStringFunc("hello world", strings.ToUpper)
    fmt.Println(result)
}
```
**A:** `HELLO WORLD`. `ReplaceAllStringFunc` calls a function for each match to compute the replacement dynamically.

---

### 20. Named Capture Groups
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
    m := re.FindStringSubmatch("date: 2024-03-15")
    names := re.SubexpNames()
    for i, name := range names {
        if name != "" {
            fmt.Printf("%s: %s\n", name, m[i])
        }
    }
}
```
**A:**
```
year: 2024
month: 03
day: 15
```
`(?P<name>...)` syntax names a capture group. `SubexpNames()` returns all group names in order.

---

### 21. regexp.Split
**Q: What is the output?**
```go
re := regexp.MustCompile(`[,\s]+`)
parts := re.Split("one, two,   three", -1)
fmt.Println(parts)
```
**A:** `[one two three]`. `Split` works like `strings.Split` but with a regexp delimiter.

---

### 22. Precompile Regexps at Package Level
**Q: Why is this better?**
```go
// BAD: compiles the regexp on every function call
func isEmail(s string) bool {
    matched, _ := regexp.MatchString(`^[\w.]+@[\w.]+$`, s)
    return matched
}

// GOOD: compile once
var emailRe = regexp.MustCompile(`^[\w.]+@[\w.]+$`)
func isEmail(s string) bool { return emailRe.MatchString(s) }
```
**A:** `regexp.MustCompile` compiles the regex once at program startup. Calling `regexp.MatchString` in a hot path recompiles every time. Precompiled regexps are also thread-safe for concurrent use.

---

### 23. Literal Dot vs Escaped Dot
**Q: What is the output?**
```go
re1 := regexp.MustCompile(`go.lang`)  // . matches ANY character
re2 := regexp.MustCompile(`go\.lang`) // \. matches literal dot

fmt.Println(re1.MatchString("goxlang")) // true  — . matches 'x'
fmt.Println(re2.MatchString("goxlang")) // false — needs literal '.'
fmt.Println(re2.MatchString("go.lang")) // true
```
**A:**
```
true
false
true
```

---

### 24. FindAllStringSubmatch — Multiple Matches with Groups
**Q: What is the output?**
```go
re := regexp.MustCompile(`(\w+)=(\w+)`)
matches := re.FindAllStringSubmatch("a=1 b=2 c=3", -1)
for _, m := range matches {
    fmt.Printf("key=%s value=%s\n", m[1], m[2])
}
```
**A:**
```
key=a value=1
key=b value=2
key=c value=3
```

---

## Section 3: runtime Package (Q25–Q34)

### 25. runtime.GOMAXPROCS
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // 0 = query current value
    runtime.GOMAXPROCS(2)
    fmt.Println("After set:", runtime.GOMAXPROCS(0))
}
```
**A:** (example on 8-core machine)
```
CPUs: 8
GOMAXPROCS: 8
After set: 2
```
`GOMAXPROCS` sets the max number of OS threads that can execute Go code simultaneously. Default = `NumCPU()` since Go 1.5.

---

### 26. runtime.NumGoroutine
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    fmt.Println("before:", runtime.NumGoroutine())
    for i := 0; i < 5; i++ {
        go func() { time.Sleep(time.Second) }()
    }
    fmt.Println("after:", runtime.NumGoroutine())
}
```
**A:**
```
before: 1
after: 6
```
`NumGoroutine` returns the number of currently running goroutines. Useful for detecting goroutine leaks in tests.

---

### 27. runtime.Gosched — Yield the Processor
**Q: What does runtime.Gosched() do?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    go func() { fmt.Println("goroutine") }()
    runtime.Gosched() // yield: allow other goroutines to run
    fmt.Println("main")
}
```
**A:** `goroutine` then `main` (likely). `Gosched` yields the processor, giving other goroutines a chance to run. Not a guarantee — use `sync.WaitGroup` for proper synchronisation.

---

### 28. runtime.GC — Trigger Garbage Collection
**Q: When should you call this?**
```go
runtime.GC() // trigger a GC cycle right now
```
**A:** Almost never in production — the GC is tuned automatically. Useful in **benchmarks** to start from a clean state (`b.ResetTimer()` + `runtime.GC()`) and in **tests** to verify finalizers or detect memory leaks.

---

### 29. runtime.ReadMemStats — Memory Usage
**Q: What does this print?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc: %d KB\n", m.Alloc/1024)
    fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
    fmt.Printf("NumGC: %d\n", m.NumGC)
}
```
**A:** Prints current heap allocation, total lifetime allocation, and GC count (values vary). `MemStats` is useful for profiling memory in tests and benchmarks.

---

### 30. runtime.Caller — Stack Frame Info
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
)

func current() string {
    _, file, line, ok := runtime.Caller(1) // 1 = caller of current()
    if !ok { return "unknown" }
    return fmt.Sprintf("%s:%d", file, line)
}

func main() {
    fmt.Println(current())
}
```
**A:** Something like `main.go:14`. `runtime.Caller(n)` returns the file/line of the nth call frame. Used in logging libraries to print the caller's location.

---

### 31. runtime.Stack — Goroutine Stack Dump
**Q: What is the pattern?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    buf := make([]byte, 4096)
    n := runtime.Stack(buf, false) // false = current goroutine only; true = all
    fmt.Printf("%s", buf[:n])
}
```
**A:** Prints the stack trace of the current goroutine — similar to what you see in a panic. Pass `true` to dump all goroutine stacks (like `SIGQUIT` on Unix). Used in signal handlers for crash reporting.

---

### 32. GOOS and GOARCH
**Q: What is the output on a 64-bit Linux system?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.GOOS)
    fmt.Println(runtime.GOARCH)
    fmt.Println(runtime.Version())
}
```
**A:**
```
linux
amd64
go1.xx.x
```
`GOOS`/`GOARCH` are the target OS/architecture at compile time. Useful for conditional code paths (though build tags are preferred for larger blocks).

---

### 33. runtime.SetFinalizer — Called Before GC Collects Object
**Q: What is the pattern?**
```go
type Resource struct{ id int }

func NewResource(id int) *Resource {
    r := &Resource{id: id}
    runtime.SetFinalizer(r, func(r *Resource) {
        fmt.Printf("finalizer: cleaning up resource %d\n", r.id)
    })
    return r
}
```
**A:** Finalizers run before GC collects the object. Use sparingly — for resources that MUST be cleaned up even if `Close()` is forgotten (e.g., open file handles, C memory). They add GC overhead and their timing is non-deterministic. Prefer explicit `Close()`/`defer`.

---

### 34. GOGC Environment Variable
**Q: What does setting `GOGC=200` do?**
**A:** `GOGC` controls the GC target percentage. Default is `100` — trigger GC when heap size doubles. `GOGC=200` triggers GC when heap is 3× the size after last collection — less frequent GC but higher memory. `GOGC=off` disables GC entirely. Tune with `GODEBUG=gctrace=1` to observe.

---

## Section 4: path/filepath Package (Q35–Q44)

### 35. filepath.Join — OS-Aware Path Building
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "path/filepath"
)

func main() {
    p := filepath.Join("home", "user", "docs", "file.txt")
    fmt.Println(p)
    fmt.Println(filepath.Join("/var", "../etc", "hosts"))
}
```
**A:** (on Linux/Mac)
```
home/user/docs/file.txt
/etc/hosts
```
`filepath.Join` uses the OS path separator and cleans `..`. Never build paths with string concatenation — use `Join`.

---

### 36. filepath.Dir and filepath.Base
**Q: What is the output?**
```go
p := "/home/user/docs/file.txt"
fmt.Println(filepath.Dir(p))
fmt.Println(filepath.Base(p))
fmt.Println(filepath.Ext(p))
```
**A:**
```
/home/user/docs
file.txt
.txt
```
`Dir` = directory portion. `Base` = filename. `Ext` = extension (including the dot).

---

### 37. filepath.Abs — Resolve to Absolute Path
**Q: What is the output?**
```go
abs, err := filepath.Abs("config.yaml")
fmt.Println(abs, err)
```
**A:** Something like `/home/user/myapp/config.yaml <nil>`. `Abs` converts a relative path to absolute using the current working directory. Always use absolute paths when configuring file watchers, loggers, or serving static files.

---

### 38. filepath.Walk
**Q: What does this do?**
```go
filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil { return err }
    if !info.IsDir() {
        fmt.Println(path)
    }
    return nil
})
```
**A:** Recursively walks the directory tree rooted at `"."`, calling the function for every file and directory. Use `filepath.WalkDir` (Go 1.16+) instead — it's more efficient as it doesn't call `os.Lstat` on every entry.

---

### 39. filepath.Glob — Pattern Matching
**Q: What is the output?**
```go
files, _ := filepath.Glob("*.go")
fmt.Println(files)
```
**A:** A slice of all `.go` files in the current directory, e.g. `[main.go utils.go]`. `Glob` returns matches in alphabetical order.

---

### 40. filepath.Match — Single Path Matching
**Q: What is the output?**
```go
matched, _ := filepath.Match("*.txt", "readme.txt")
fmt.Println(matched)

matched2, _ := filepath.Match("*.txt", "docs/readme.txt")
fmt.Println(matched2)
```
**A:**
```
true
false
```
`Match` matches only against a single component — `*` doesn't match path separators. Use `**` with `WalkDir` for recursive globbing.

---

### 41. filepath.Rel — Relative Path
**Q: What is the output?**
```go
rel, err := filepath.Rel("/home/user", "/home/user/docs/file.txt")
fmt.Println(rel, err)
```
**A:** `docs/file.txt <nil>`. `Rel` computes a relative path from `basepath` to `targpath`.

---

### 42. filepath.Clean — Normalise a Path
**Q: What is the output?**
```go
fmt.Println(filepath.Clean("/foo//bar/../baz"))
fmt.Println(filepath.Clean("./a/./b/../c"))
```
**A:**
```
/foo/baz
a/c
```
`Clean` removes duplicate separators, resolves `.` and `..`. `Join` calls `Clean` internally.

---

### 43. path vs path/filepath — When to Use Each
**Q: Which should you use for URL paths?**
```go
import "path"          // always uses '/' — for URL paths, zip entries
import "path/filepath" // uses OS separator — for OS file paths
```
**A:** Use `path` for URL segments, ZIP file entries, HTTP routes (always `/`). Use `path/filepath` for OS filesystem paths (OS-aware separator). Never use `filepath` for URLs or you'll get backslashes on Windows.

---

### 44. filepath.SplitList — PATH-Like String
**Q: What is the output?**
```go
paths := filepath.SplitList("/usr/bin:/usr/local/bin:/home/user/bin")
fmt.Println(paths)
```
**A:** `[/usr/bin /usr/local/bin /home/user/bin]`. `SplitList` splits a PATH-like colon-separated (Linux/Mac) or semicolon-separated (Windows) list — useful for parsing `$PATH`, `$GOPATH`, etc.

---

## Section 5: unicode/utf8 Package (Q45–Q54)

### 45. String Length vs Rune Count
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "Hello, 世界"
    fmt.Println(len(s))                    // byte length
    fmt.Println(utf8.RuneCountInString(s)) // rune (character) count
}
```
**A:**
```
13
9
```
`len(s)` = byte length. The two CJK characters `世界` are 3 bytes each (UTF-8). `RuneCountInString` counts Unicode code points.

---

### 46. Range Over String — Yields Runes
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i, r := range "Hi🙂" {
        fmt.Printf("index=%d rune=%c\n", i, r)
    }
}
```
**A:**
```
index=0 rune=H
index=1 rune=i
index=2 rune=🙂
```
`range` over a string decodes UTF-8 runes. The emoji `🙂` is 4 bytes — `i` jumps from 2 to 6 on the next iteration showing the byte index.

---

### 47. Indexing a String — Bytes, Not Runes
**Q: What is the output?**
```go
s := "世界"
fmt.Println(s[0])         // first byte
fmt.Printf("%c\n", s[0]) // first byte as char — garbage
r, _ := utf8.DecodeRuneInString(s)
fmt.Printf("%c\n", r)    // first rune — correct
```
**A:**
```
228        (first byte of 世 in UTF-8)
ä          (misinterpretation of single byte)
世          (correct rune)
```
`s[i]` gives a raw byte. Use `utf8.DecodeRuneInString` or `[]rune(s)` conversion for characters.

---

### 48. []rune vs []byte Conversion
**Q: What is the output?**
```go
s := "Hello, 世界"
runes := []rune(s)
fmt.Println(len(runes))     // rune count
fmt.Println(string(runes[7])) // 8th rune
```
**A:**
```
9
世
```
`[]rune(s)` converts to a slice of Unicode code points — safe for character-level indexing. Expensive for long strings because it allocates. Prefer `range` when iterating.

---

### 49. utf8.Valid — Validate UTF-8
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    valid   := []byte("Hello, 世界")
    invalid := []byte{0xff, 0xfe, 0xfd} // invalid UTF-8
    fmt.Println(utf8.Valid(valid))
    fmt.Println(utf8.Valid(invalid))
}
```
**A:**
```
true
false
```
`utf8.Valid` checks if a byte slice is valid UTF-8. Important when accepting bytes from external sources before converting to `string`.

---

### 50. utf8.RuneLen — Bytes Needed for a Rune
**Q: What is the output?**
```go
fmt.Println(utf8.RuneLen('A'))  // ASCII
fmt.Println(utf8.RuneLen('é'))  // Latin extended
fmt.Println(utf8.RuneLen('世')) // CJK
fmt.Println(utf8.RuneLen('🙂')) // emoji
```
**A:**
```
1
2
3
4
```
UTF-8 uses 1–4 bytes per rune. ASCII = 1 byte; most Western European = 2 bytes; CJK = 3 bytes; emoji = 4 bytes.

---

### 51. Reverse a UTF-8 String Correctly
**Q: Which version is correct?**
```go
// WRONG: reverses bytes, corrupts multi-byte runes
func reverseBytes(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

// CORRECT: reverses runes, preserves characters
func reverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```
**A:**
- `reverseBytes("世界")` → corrupted bytes (invalid UTF-8)
- `reverseRunes("世界")` → `"界世"` ✅

Always operate on `[]rune` when you need character-level manipulation.

---

### 52. unicode.IsLetter / IsDigit / IsSpace
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unicode"
)

func main() {
    fmt.Println(unicode.IsLetter('A'))
    fmt.Println(unicode.IsLetter('世'))
    fmt.Println(unicode.IsDigit('5'))
    fmt.Println(unicode.IsSpace('\t'))
    fmt.Println(unicode.IsUpper('a'))
}
```
**A:**
```
true
true
true
true
false
```
`unicode.IsLetter` covers ALL Unicode letters — not just ASCII. Essential for international text processing.

---

### 53. strings.Map with unicode Functions
**Q: What is the output?**
```go
import (
    "fmt"
    "strings"
    "unicode"
)

s := "Hello, 世界 123!"
onlyLetters := strings.Map(func(r rune) rune {
    if unicode.IsLetter(r) { return r }
    return -1 // -1 drops the rune
}, s)
fmt.Println(onlyLetters)
```
**A:** `Hello世界`. `strings.Map` with a function returning `-1` removes that rune from the output.

---

### 54. utf8.AppendRune — Encode a Rune to Bytes
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    var buf []byte
    buf = utf8.AppendRune(buf, 'H')
    buf = utf8.AppendRune(buf, '世')
    fmt.Println(string(buf))
    fmt.Println(len(buf)) // 1 + 3 = 4 bytes
}
```
**A:**
```
H世
4
```
`utf8.AppendRune` (Go 1.18+) appends the UTF-8 encoding of a rune to a byte slice — useful for building UTF-8 byte slices without `[]byte` conversions.

---

## Section 6: embed Package (Q55–Q60)

### 55. go:embed — Embed a File
**Q: What is the output (given `hello.txt` contains `"Hello, embed!"`)? **
```go
package main
import (
    _ "embed"
    "fmt"
)

//go:embed hello.txt
var content string

func main() {
    fmt.Println(content)
}
```
**A:** `Hello, embed!`. The `//go:embed` directive embeds the file content into the binary at compile time. The variable is set before `main()` runs.

---

### 56. go:embed — Embed as []byte
**Q: What are the valid variable types for go:embed?**
```go
//go:embed logo.png
var logo []byte  // binary file as bytes

//go:embed config.json
var config string  // text file as string

//go:embed static/
var staticFS embed.FS  // directory as virtual filesystem
```
**A:** Three supported types: `string` (for text), `[]byte` (for binary), `embed.FS` (for directories and multiple files). All other types cause a compile error.

---

### 57. embed.FS — Directory Embedding
**Q: What is the output?**
```go
package main
import (
    "embed"
    "fmt"
)

//go:embed testdata
var files embed.FS

func main() {
    data, err := files.ReadFile("testdata/greeting.txt")
    fmt.Println(string(data), err)

    entries, _ := files.ReadDir("testdata")
    for _, e := range entries {
        fmt.Println(e.Name())
    }
}
```
**A:** Prints contents of `greeting.txt`, then names of all files in `testdata/`. `embed.FS` implements `fs.FS` so it works with `http.FileServer`, `text/template`, etc.

---

### 58. go:embed — Files Embedded at Build Time
**Q: What is the key advantage?**
```go
//go:embed migrations/
var migrationsFS embed.FS
// Use to run DB migrations without shipping separate files
```
**A:** Files are bundled into the binary — no external files needed at runtime. Perfect for: SQL migrations, HTML templates, static web assets, certificate files, config defaults. The binary becomes fully self-contained.

---

### 59. embed.FS as http.FileServer
**Q: What is the pattern?**
```go
//go:embed static
var staticFiles embed.FS

func main() {
    sub, _ := fs.Sub(staticFiles, "static")
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(sub))))
    http.ListenAndServe(":8080", nil)
}
```
**A:** `http.FS(sub)` adapts `embed.FS` to `http.FileSystem`. `fs.Sub` removes the prefix directory. Combined result: embed a `static/` folder and serve it over HTTP as a self-contained binary.

---

### 60. go:embed Restrictions
**Q: Which of these are valid embed directives?**
```go
//go:embed *.txt              // pattern — valid
//go:embed data/              // directory — valid  
//go:embed ../secret.txt      // path outside module — INVALID
//go:embed .git/              // hidden directory — INVALID (unless -trimpath)
//go:embed nonexistent.txt    // missing file — COMPILE ERROR
```
**A:**
- Patterns and directories ✅
- Paths outside the module root ❌
- `.dot` directories (hidden) ❌ by default (except `.well-known`)
- Missing files ❌ — compile error (not runtime!)

---

## Section 7: net/url & os/exec (Q61–Q72)

### 61. url.Parse — Breaking Down a URL
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "net/url"
)

func main() {
    u, _ := url.Parse("https://user:pass@example.com:8080/path?key=val#frag")
    fmt.Println(u.Scheme)
    fmt.Println(u.Host)
    fmt.Println(u.Path)
    fmt.Println(u.RawQuery)
    fmt.Println(u.Fragment)
    fmt.Println(u.User.Username())
}
```
**A:**
```
https
example.com:8080
/path
key=val
frag
user
```

---

### 62. url.Values — Building Query Strings
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "net/url"
)

func main() {
    params := url.Values{}
    params.Set("q", "golang generics")
    params.Set("page", "2")
    params.Add("tag", "go")
    params.Add("tag", "programming")
    fmt.Println(params.Encode())
}
```
**A:** `page=2&q=golang+generics&tag=go&tag=programming` (sorted by key). `url.Values` handles URL encoding and multi-value keys. `Encode()` produces a properly escaped query string.

---

### 63. url.QueryEscape / QueryUnescape
**Q: What is the output?**
```go
s := url.QueryEscape("Hello World! 你好")
fmt.Println(s)
fmt.Println(url.QueryUnescape(s))
```
**A:**
```
Hello+World%21+%E4%BD%A0%E5%A5%BD
Hello World! 你好 <nil>
```
Spaces become `+` or `%20`. Special chars are percent-encoded. Unicode characters are UTF-8 encoded then percent-encoded.

---

### 64. Parsing Query Parameters from Request
**Q: What is the pattern?**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()           // returns url.Values
    name := q.Get("name")        // first value for "name"
    tags := q["tag"]             // all values for "tag" (slice)
    fmt.Fprintf(w, "name=%s tags=%v", name, tags)
}
```
**A:** `r.URL.Query()` parses the query string into `url.Values`. `Get` returns the first value for a key. Direct map access returns all values as a slice.

---

### 65. url.URL.String() — Reconstruct URL
**Q: What is the output?**
```go
u := &url.URL{
    Scheme:   "https",
    Host:     "api.example.com",
    Path:     "/v1/users",
    RawQuery: url.Values{"limit": {"10"}, "offset": {"0"}}.Encode(),
}
fmt.Println(u.String())
```
**A:** `https://api.example.com/v1/users?limit=10&offset=0`. Build URLs programmatically rather than string formatting — handles encoding correctly.

---

### 66. os/exec — Run a Command
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os/exec"
)

func main() {
    out, err := exec.Command("echo", "Hello from exec").Output()
    fmt.Println(string(out), err)
}
```
**A:** `Hello from exec <nil>`. `exec.Command` creates a `Cmd`. `.Output()` runs it and captures stdout. On Windows, use `exec.Command("cmd", "/C", "echo", "Hello")`.

---

### 67. exec.Command — Never Use Shell
**Q: What is the security issue?**
```go
// DANGEROUS: shell injection vulnerability
name := userInput // e.g., "; rm -rf /"
exec.Command("sh", "-c", "echo "+name).Run()

// SAFE: args are passed directly, no shell interpretation
exec.Command("echo", name).Run()
```
**A:** Never pass user input to `sh -c`. Each argument to `exec.Command` is passed directly to the OS — no shell expansion, no injection. Shell metacharacters (`; & | >`) are treated as literal strings.

---

### 68. exec.Cmd — Capture Stderr Separately
**Q: What is the pattern?**
```go
cmd := exec.Command("go", "build", "./...")
var stdout, stderr bytes.Buffer
cmd.Stdout = &stdout
cmd.Stderr = &stderr
err := cmd.Run()
if err != nil {
    fmt.Println("stderr:", stderr.String())
}
```
**A:** Assigning `Stdout`/`Stderr` to buffers captures each stream independently. If you want combined output, assign the same writer to both, or use `cmd.CombinedOutput()`.

---

### 69. exec.LookPath — Find Executable in PATH
**Q: What is the output?**
```go
path, err := exec.LookPath("go")
fmt.Println(path, err)

path2, err2 := exec.LookPath("nonexistent")
fmt.Println(path2, err2)
```
**A:** (example)
```
/usr/local/go/bin/go <nil>
 exec: "nonexistent": executable file not found in $PATH
```
`LookPath` checks if a binary exists in `$PATH` — use before `exec.Command` when you need to verify availability.

---

### 70. exec.CommandContext — Cancellable Command
**Q: What is the pattern?**
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

cmd := exec.CommandContext(ctx, "sleep", "30")
err := cmd.Run()
// After 5s: err = "signal: killed"
```
**A:** `CommandContext` ties the command lifetime to a context. When the context is cancelled/times out, the process is killed (SIGKILL on Unix). The idiomatic way to add timeouts to external commands.

---

### 71. sort.Slice — Sort Any Slice
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    people := []struct {
        Name string
        Age  int
    }{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })
    for _, p := range people {
        fmt.Printf("%s(%d) ", p.Name, p.Age)
    }
}
```
**A:** `Bob(25) Alice(30) Charlie(35) `. `sort.Slice` sorts in-place using a less function. Not guaranteed stable — use `sort.SliceStable` if equal elements must preserve order.

---

### 72. math/rand — Seeded Random Numbers (Go 1.20+)
**Q: What is the output and what changed in Go 1.20?**
```go
package main
import (
    "fmt"
    "math/rand"
)

func main() {
    // Go 1.20+: global source is automatically seeded — no need for rand.Seed()
    fmt.Println(rand.Intn(100))
    fmt.Println(rand.Float64())
}
```
**A:** Random values. Before Go 1.20, `rand.Seed(time.Now().UnixNano())` was needed — otherwise the global source was deterministic (seed=1). Since Go 1.20, the global source is **automatically randomly seeded**. `rand.Seed` is deprecated.

---
