# 🟡 **221–240: Files, OS, and System Programming**

### 221. How do you read a file line by line in Go?
"I use `bufio.Scanner`.

`scanner := bufio.NewScanner(file)`
`for scanner.Scan() { line := scanner.Text() }`
This is memory efficient because it doesn't load whole file into RAM.
However, I always check `scanner.Err()` after the loop. If the line is too long (over 64KB default), the scanner might error out, so for massive lines, I switch to `bufio.Reader.ReadLine`."

#### Indepth
`bufio.Scanner` strips the newline character (`\n`) automatically. Be wary of this if you are rewriting the file—you must manually append `\n`. Also, the default buffer size is 64KB. If you are reading lines longer than that (e.g., minified JSON), use `scanner.Buffer()` to increase the limit, or use `bufio.Reader` which can grow indefinitely.

---

### 222. How do you write large files efficiently?
"I use `bufio.Writer`.

Writing 1 byte at a time to disk is a syscall, which is slow.
`bufio.NewWriter(file)` buffers the writes in memory (default 4KB) and flush them to disk in fewer, larger chunks."

#### Indepth
On modern SSDs, writing 4KB vs 16KB doesn't change much, but on networked filesystems (NFS, EFS) or rotating rust, buffering is the difference between minutes and seconds. If you are writing critical data (like a WAL), call `file.Sync()` after `Flush()` to force the OS to physically write headers to the platter/NAND.
**Crucial step**: ALWAYS call `writer.Flush()` before closing the file, or the last chunk of data will be lost forever."

---

### 223. How do you watch file system changes in Go?
"I use the **fsnotify** library (cross-platform).

1.  Create a watcher: `fsnotify.NewWatcher()`.
2.  Add directories: `watcher.Add("/config")`.
3.  Loop over events: `select { case event := <-watcher.Events: ... }`.
It hooks into operating system primitives (inotify on Linux, FSEvents on macOS) to detect file creation, modification, or deletion instantly."

#### Indepth
Watching files recursively is hard. On Linux (`inotify`), you must manually add a watch for *every* subdirectory. If a user does `mkdir -p a/b/c`, you need to catch the creation of `a`, add a watch to it, catch `b`, etc. Libraries like `notify` wrap this complexity but `fsnotify` is the low-level primitive.

---

### 224. How to get file metadata like size, mod time?
"I use `os.Stat(filename)`.

It returns a `FileInfo` interface.
`info.Size()` gives me bytes. `info.ModTime()` gives me the timestamp.
If `os.Stat` returns an error, I check `errors.Is(err, os.ErrNotExist)` to handle the 'file not found' case explicitly."

#### Indepth
`os.Stat` follows symlinks. If you want to check the symlink itself (where it points to), use `os.Lstat`. Also, be aware of TOCTOU (Time Of Check, Time Of Use) bugs: checking if a file exists and then opening it is race-prone. It's better to just open it and handle the error.

---

### 225. How do you work with CSV files in Go?
"I use the standard `encoding/csv` package.

To read: `reader := csv.NewReader(file)`. `record, err := reader.Read()`.
To write: `writer := csv.NewWriter(file)`. `writer.Write([]string{"col1", "col2"})`.
I always `defer writer.Flush()` and check `writer.Error()` to ensure all data was physically written to disk."

#### Indepth
`encoding/csv` handles edge cases like quoted fields containing newlines: `"Hello\nWorld"`. If you are processing massive CSVs (GBs), avoid `ReadAll`. Use the streaming `Read()` in a loop to keep memory usage low. You can also set `reader.ReuseRecord = true` to avoid allocating a new slice for every row (zero-allocation parsing).

---

### 226. How do you compress and decompress files in Go?
"I use `compress/gzip`.

It wraps any `io.Writer`.
`zw := gzip.NewWriter(file)`.
Now, anything I write to `zw` is compressed on the fly and sent to the file.
To read, I wrap with `gzip.NewReader(file)`. It acts just like opening a normal file, but the bytes are decompressed transparently as I read them."

#### Indepth
Gzip is not seekable. You cannot jump to the middle of a `gz` file. For random access compressed data, you need block-based compression (like Snappy or Zstd with framing). Also, `gzip.Writer` buffers heavily; ignoring the `Close()` error means you might silently truncate the file footer.

---

### 227. How do you execute shell commands from Go?
"I use `os/exec`.

`cmd := exec.Command("ls", "-la")`.
I capture the output: `out, err := cmd.CombinedOutput()`.
If I need to stream the output (e.g., for a long-running build script), I pipe `cmd.Stdout` to `os.Stdout`.
I am extremely careful **never** to pass user input directly to `exec.Command` to avoid injection attacks."

#### Indepth
`CombinedOutput` captures both stdout and stderr. If you need to process them separately (e.g., log stderr as errors but parse stdout as JSON), you must assign separate `bytes.Buffer`s to `cmd.Stdout` and `cmd.Stderr`. This gives you granular control over the subprocess output.

---

### 228. What is the `os/exec` package used for?
"It lets my Go program run *other* programs.

I use it heavily for automation: invoking git, running docker builds, or calling legacy binaries.
I usually use `CommandContext` so I can kill the process if it hangs:
`ctx, cancel := context.WithTimeout(ctx, 5*time.Second)`
`exec.CommandContext(ctx, "sleep", "10").Run()`."

#### Indepth
When `CommandContext` kills a process, it sends `SIGKILL` (on POSIX). This gives the child no chance to clean up (delete temp files, release locks). You can customize this by setting `cmd.Cancel` to send `SIGTERM` first, wait a bit, and then kill.

---

### 229. How do you set environment variables in Go?
"For the current process: `os.Setenv("PORT", "8080")`.

To read: `port := os.Getenv("PORT")`.
If I need to distinguish between 'empty' and 'unset', I use `val, ok := os.LookupEnv("PORT")`.
When running a *subprocess* (`exec.Cmd`), I set `cmd.Env` explicitly to pass variables to only that child command."

#### Indepth
Values set with `os.Setenv` are process-wide. This is not thread-safe if you have multiple tests running in parallel that expect different env vars. For testing, it's safer to pass config via checks or dependency injection rather than relying on global environment mutations.

---

### 230. How to create and manage temp files/directories?
"I use `os.CreateTemp` (modern replacement for `ioutil`).

`f, err := os.CreateTemp("", "example-*.txt")`.
The `*` is replaced by a random string to prevent collisions.
I immediately `defer os.Remove(f.Name())` to clean up the file when the function exits. If I need a directory, I use `os.MkdirTemp`."

#### Indepth
The operating system automatically cleans up `/tmp` (or `%TEMP%`) effectively on reboot, but long-running servers can fill up the disk with abandoned temp files. Always use `defer` to clean up. In containerized environments, `/tmp` might be an in-memory `tmpfs`, so writing large temp files could OOM the pod.

---

### 231. How do you handle signals like SIGINT in Go?
"I use `os/signal` to implement graceful shutdowns.

`c := make(chan os.Signal, 1)`
`signal.Notify(c, os.Interrupt, syscall.SIGTERM)`
I block on `<-c`. When the user hits Ctrl+C, I catch it, cancel my context, close my database connections, and exit cleanly. If I didn't effectively catch this, the OS would kill my process mid-write."

#### Indepth
If you capture `SIGINT` but don't exit, the user will be annoyed (they have to `kill -9`). The standard pattern is to listen for the signal *once* to trigger graceful shutdown, and if the user hits Ctrl+C *again* (force quit), the default handler takes over and kills the process immediately.

---

### 232. How do you gracefully shut down a CLI app?
"I bind the OS signal to a Context.
`ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)`.

My main loop looks like:
`select { case <-ctx.Done(): cleanUp(); return }`.
Using `NotifyContext` automatically propagates the cancellation to all my child goroutines (HTTP clients, DB queries), ensuring a hierarchical shutdown."

#### Indepth
Graceful shutdown usually involves a timeout. `Select { case <-ctx.Done(): stop; case <-time.After(30*time.Second): forceStop }`. You don't want your deployment to hang for 10 minutes because one goroutine refused to exit. Kubernetes has a `terminationGracePeriodSeconds` (default 30s) after which it sends SIGKILL.

---

### 233. What are file descriptors and how does Go manage them?
"A file descriptor (FD) is an integer handle for an open resource (file, socket, pipe).

The OS has a limit (ulimit -n), usually 1024 or 65535.
In Go, `os.File` wraps the FD.
If I forget to `f.Close()`, I leak the FD. Eventually, I hit 'too many open files' and the app crashes. Go's GC *might* close it with a Finalizer, but I never rely on that."

#### Indepth
`ulimit` is a common production outage cause. The default on many Linux distros is 1024. A high-throughput web server needs tens of thousands. Always check `ulimit -n` in your startup script or Dockerfile (`LimitNOFILE` in systemd). Network sockets count as FDs too!

---

### 234. How to handle large file uploads and streaming?
"I use `io.Pipe` or direct streaming.

I never load the whole file into RAM (`ioutil.ReadAll` is banned).
`io.Copy(dstFile, requestBody)`.
This streams data in 32KB chunks. I can upload a 50GB file using only a few MB of RAM. It’s strictly IO-bound, not Memory-bound."

#### Indepth
For `io.Pipe`, the writer blocks until the reader reads. This is synchronous. It's powerful for transforming streams (e.g., File -> Gzip -> Encrypt -> S3) without ever holding more than a few kilobytes in RAM. It’s the Unix Philosophy applied to Go interfaces.

---

### 235. How do you access OS-specific syscalls in Go?
"I use `golang.org/x/sys/unix`.

It provides low-level access to kernel calls unimplemented in the standard library.
Examples: `unix.Mlock` (prevent swapping), `unix.Mount`, or specific socket flags (`SO_REUSEPORT`).
I guard this code with build tags (`//go:build linux`) because these calls obviously won't work on Windows."

#### Indepth
`golang.org/x/sys/unix` is preferred over the frozen `syscall` package because it's actively maintained and updated with new kernel features (like `io_uring` or `bpf`). It uses generated code to match the exact kernel headers of the target OS.

---

### 236. How do you implement a simple CLI tool in Go?
"I start with the standard `flag` package.

`name := flag.String("name", "World", "name to greet")`
`flag.Parse()`
If the tool grows (needs subcommands like `git commit`), I switch to **Cobra** or **Urfave CLI**.
They handle help text generation, bash completion, and nested commands (`app server start`) automatically."

#### Indepth
When building CLIs, follow the **12-Factor App** principles. Allow configuration via *both* flags and environment variables. Libraries like `viper` facilitate this precedence: Flag > Env > Config File > Default. This makes your tool friendly for both humans (flags) and scripts/Kubernetes (env).

---

### 237. How do you build cross-platform binaries in Go?
"It’s trivial. I set `GOOS` and `GOARCH`.

Check: `GOOS=linux GOARCH=arm64 go build`.
This produces a binary for Raspberry Pi / AWS Graviton on my MacBook.
I don't need Docker or a VM. The Go compiler knows how to generate machine code for almost every architecture out there."

#### Indepth
Cross-compilation disables CGO by default (`CGO_ENABLED=0`). If your app depends on C libraries (like SQLite), cross-compiling becomes much harder—you need a C cross-compiler (like `zig cc` or `musl-gcc`). For pure Go apps, it's seamless.

---

### 238. What is syscall vs os vs exec package difference?
"**syscall**: Raw kernel interface (deprecated, use `x/sys`). Hard to use.
**os**: Platform-independent wrapper (Open, Read, Write). Use this 99% of the time.
**os/exec**: For running *other* programs.
I use `os` to manipulate files, and `exec` to run scripts."

#### Indepth
The `os` package is designed to be POSIX-like but Windows-compatible. However, file permissions (`chmod`) behave very differently on Windows. Go tries to map them (e.g., ReadOnly bit), but don't expect 0777 to mean the same thing on NTFS as it does on ext4.

---

### 239. How do you write to logs with rotation?
"Standard library `log` does **not** support rotation.

I use **Lumberjack** (`gopkg.in/natefinch/lumberjack.v2`).
It implements `io.Writer`.
`log.SetOutput(&lumberjack.Logger{Filename: "app.log", MaxSize: 100})`.
It automatically closes the file when it hits 100MB, renames it to `app-timestamp.log.gz`, and starts a fresh `app.log`."

#### Indepth
In containerized environments (Kubernetes/Docker), **logs should go to stdout**, not files. Log rotation is the responsibility of the infrastructure (Docker logging driver or Fluentd), not the application. File logging is mostly for legacy VM/bare-metal deployments.

---

### 240. What is the use of `ioutil` and its deprecation?
"`ioutil` was a grab-bag of utility functions (`ReadFile`, `WriteFile`).
It is now **deprecated** (since Go 1.16).

I now use:
*   `os.ReadFile` / `os.WriteFile`.
*   `io.ReadAll`.
*   `io.Discard`.
The old code still works, but modern Go code should use the new, more logical locations (`os` and `io`)."

#### Indepth
One big improvement in the migration: `os.ReadFile` returns the file contents. `ioutil.ReadFile` did the same but was just a wrapper. The new structure organizes functions by *what they act on* (`os` = filesystem, `io` = abstract streams), reducing cyclic dependencies in the standard library.
