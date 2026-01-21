## 🟢 CLI Tools, Automation, and Scripting (Questions 441-460)

### Question 441: How do you build an interactive CLI in Go?

**Answer:**
You need to handle `stdin` and `stdout` effectively. Libraries like **Alephier/survey** or **manifoldco/promptui** provide interactive prompts (select, confirm, input) easily.

```go
import "github.com/manifoldco/promptui"

prompt := promptui.Select{
    Label: "Select Day",
    Items: []string{"Monday", "Tuesday", "Wednesday"},
}

_, result, _ := prompt.Run()
fmt.Printf("You choose %q\n", result)
```

---

### Question 442: What libraries do you use for command-line tools in Go?

**Answer:**
1.  **Cobra:** The standard for complex CLIs (nested commands, flags). Used by Kubernetes, Hugo, GitHub CLI.
2.  **Urfave/cli:** Another popular, slightly simpler framework.
3.  **Viper:** For configuration management (often paired with Cobra).
4.  **Bubbletea:** For TUI (Text User Interface) applications (Elm architecture).

---

### Question 443: How do you parse flags and config in CLI?

**Answer:**
- **Simple:** Use the standard `flag` package.
- **Robust:** Use **Cobra** + **Viper**. Cobra handles the command structure and flags (`--port`), while Viper handles reading config files (JSON/YAML), env vars, and binding them to flags.

```go
// Cobra + Viper
var port string
rootCmd.PersistentFlags().StringVar(&port, "port", "8080", "Port to run on")
viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
```

---

### Question 444: How do you implement bash autocompletion for Go CLI?

**Answer:**
If you use **Cobra**, it is built-in.
You just need to add a command to generate the script.

```go
var completionCmd = &cobra.Command{
    Use:   "completion [bash|zsh|fish|powershell]",
    Run: func(cmd *cobra.Command, args []string) {
        switch args[0] {
        case "bash":
            cmd.Root().GenBashCompletion(os.Stdout)
        }
    },
}
```
User runs `source <(myapp completion bash)` to enable it.

---

### Question 445: How would you use Cobra to build a nested command CLI?

**Answer:**
Cobra is designed for this structure: `App -> Command -> SubCommand`.

```go
var rootCmd = &cobra.Command{Use: "app"}
var userCmd = &cobra.Command{Use: "user"}
var createCmd = &cobra.Command{
    Use: "create",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Creating user...")
    },
}

func init() {
    rootCmd.AddCommand(userCmd)
    userCmd.AddCommand(createCmd)
}
// Run: ./app user create
```

---

### Question 446: How do you manage color and styling in terminal output?

**Answer:**
Use libraries like **fatih/color** or **charmbracelet/lipgloss**.
`color` is simple (red text, bold). `lipgloss` is for advanced styling (borders, padding, layouts) similar to CSS.

```go
import "github.com/fatih/color"

color.Red("Error: %s", err)
color.Green("Success!")
```

---

### Question 447: How would you stream CLI output like tail -f?

**Answer:**
You need to continuously read from a source (file/socket) and write to `stdout`, then `Flush`.

```go
func streamFile(path string) {
    t, _ := tail.TailFile(path, tail.Config{Follow: true})
    for line := range t.Lines {
        fmt.Println(line.Text)
    }
}
```
Or simply loop and write, ensuring `os.Stdout.Sync()` is called if buffering issues occur.

---

### Question 448: How do you handle secrets securely in a CLI?

**Answer:**
1.  **Never print secrets to stdout/logs.**
2.  **Input:** Use `golang.org/x/term` to read password without echo.
    ```go
    bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
    ```
3.  **Storage:** Use the OS Keychain (Mac Keychain, Windows Credential Manager) via libraries like `keyring` (zalando/go-keyring), rather than plaintext config files.

---

### Question 449: How do you bundle a CLI as a standalone binary?

**Answer:**
Go compiles to a single static binary by default.
`go build -o myapp main.go`
This binary includes all dependencies and standard libraries. It can be copied to any machine with the same OS/Arch and run without installing Go.

---

### Question 450: How would you version and release CLI with GitHub Actions?

**Answer:**
Use **GoReleaser**.
1.  Add `.goreleaser.yaml` to repo.
2.  Create a GitHub Action that runs `goreleaser release` on a new Git Tag.
This automatically:
- Cross-compiles (Windows, Linux, Mac).
- Creates archives (.tar.gz, .zip).
- Generates checksums.
- Publishes a GitHub Release.
- Can update Homebrew taps.

---

### Question 451: How do you schedule a Go CLI tool with cron?

**Answer:**
The Go tool itself is passive. You use the OS Scheduler.
- **Linux:** Add entry to `/etc/crontab` or `crontab -e`.
  `0 * * * * /usr/local/bin/myapp cleanup`
- **Internal Scheduling:** If the app runs permanently (daemon), use `robfig/cron` to run functions periodically inside the app.

---

### Question 452: How do you use Go as a scripting language?

**Answer:**
While compiled, Go can feel like a script using `go run main.go`.
For "Shebang" support, you can use specialized tools like **Gorun**, or simply write a bash wrapper:

```bash
#!/usr/bin/env go run
package main
import "fmt"
func main() { fmt.Println("Hello Script") }
```
(Note: The shebang trick has nuances, but `go run` is the standard "scripting" mode).

---

### Question 453: How do you embed templates in your Go CLI tool?

**Answer:**
Use `embed` package (Go 1.16+). This allows shipping a single binary that generates complex files (scaffolding).

```go
//go:embed templates/*.html
var content embed.FS

func scaffold() {
    data, _ := content.ReadFile("templates/index.html")
    os.WriteFile("index.html", data, 0644)
}
```

---

### Question 454: How would you create a system daemon in Go?

**Answer:**
1.  **Systemd (Linux):** Write a `.service` file pointing to your Go binary.
2.  **Windows Service:** Use `golang.org/x/sys/windows/svc` to handle Start/Stop signals.
3.  **Library:** Use `github.com/kardianos/service` to abstract OS differences (install/start/stop/restart) programmatically.

---

### Question 455: What are good patterns for CLI testing?

**Answer:**
1.  **Separate Logic from Main:** Do not put logic in `main()`. Put it in `Run(cmd, args) error`.
2.  **Dependency Injection:** Pass `io.Reader` (stdin) and `io.Writer` (stdout) to your commands so you can inject bytes and capture output in tests without real user input.

```go
func RunApp(out io.Writer) {
    fmt.Fprintln(out, "Hello")
}

func TestApp(t *testing.T) {
    buf := new(bytes.Buffer)
    RunApp(buf)
    if buf.String() != "Hello\n" { t.Fail() }
}
```

---

### Question 456: How do you store and manage CLI state/config files?

**Answer:**
Use `os.UserConfigDir()` standard.
- **Linux:** `~/.config/myapp/config.yaml`
- **Mac:** `~/Library/Application Support/myapp/...`
- **Windows:** `%APPDATA%\myapp\...`

Use **Viper** to read/write these files automatically.

---

### Question 457: How do you secure a CLI for local system access?

**Answer:**
1.  **File Permissions:** Create config files with `0600` (Read/Write for owner only).
2.  **Least Privilege:** Do not require `sudo` unless necessary.
3.  **Sanitize Input:** If executing shell commands based on user input, use `exec.Command` (safe) instead of passing strings to `sh -c` (shell injection risk).

---

### Question 458: How do you test CLI tools across multiple OS in CI?

**Answer:**
Use GitHub Actions with a matrix strategy.

```yaml
jobs:
  test:
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
      - run: go test ./...
```

---

### Question 459: How do you expose analytics and usage for a CLI?

**Answer:**
Send events to a remote HTTP endpoint (e.g., Google Analytics, Mixpanel, or custom) asynchronously.
**Critical:**
1.  Must fail silently (don't break the CLI if internet is down).
2.  Must respect `Do Not Track` or have an opt-out config.
3.  Anonymize PII (IP address, Username).

---

### Question 460: How would you build a CLI wrapper for REST APIs?

**Answer:**
1.  **Auto-Generation:** Use `openapi-generator` to create a Go client SDK from the API Swagger spec.
2.  **Cobra Structure:** Map API endpoints to Commands (`app get users` -> `GET /users`).
3.  **Output Formatting:** Allow users to choose output format: `--json`, `--yaml`, or `--table` (human readable tables).

---
