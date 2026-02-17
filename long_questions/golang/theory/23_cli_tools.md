# 🟢 Go Theory Questions: 441–460 CLI Tools, Automation, and Scripting

## 441. How do you build an interactive CLI in Go?

**Answer:**
We use `manifoldco/promptui` or `charmbracelet/bubbletea`.

Standard `fmt.Scanln` is too primitive. `promptui` allows:
1.  **Select Lists**: Arrow keys to choose options.
2.  **Confirmations**: "Are you sure? [y/N]".
3.  **Password Masking**: "Enter Password: *****".

The `bubbletea` library takes this further, allowing full TUI (Text User Interface) applications with complex layout models (elm architecture) right in the terminal.

---

## 442. What libraries do you use for command-line tools in Go?

**Answer:**
**Cobra** is the industry standard (used by Kubernetes, Docker, Hugo).
It handles:
*   Subcommands (`git commit`, `git clone`).
*   Global vs Local flags.
*   Help text generation (`-h`).

We pair it with **Viper** for config management.
For smaller, simpler tools, `urfave/cli` is a lightweight alternative, but Cobra is the choice for anything enterprise-grade.

---

## 443. How do you parse flags and config in CLI?

**Answer:**
For flags: `cobra` commands define flags.
`rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")`.

For processing:
1.  **Defaults**: Defined in code.
2.  **Config File**: Viper reads `config.yaml`.
3.  **Env Vars**: Viper reads `MYAPP_PORT`.
4.  **Flags**: Highest priority (overrides everything).

This hierarchy allows a user to have a base config file but override specific settings for a one-off command run.

---

## 444. How do you implement bash autocompletion for Go CLI?

**Answer:**
Cobra allows generating completion scripts automatically.

`rootCmd.GenBashCompletion(os.Stdout)`
`rootCmd.GenZshCompletion(os.Stdout)`

Users perform `source <(myapp completion bash)`.
To make it dynamic (e.g., `myapp deploy <tab>` suggests active docker containers), we define a `ValidArgsFunction` in the Cobra command definition that calls our Go logic to fetch the list of candidates at runtime.

---

## 445. How would you use `cobra` to build a nested command CLI?

**Answer:**
Cobra is a tree structure.

```go
var rootCmd = &cobra.Command{Use: "app"}
var userCmd = &cobra.Command{Use: "user"}
var createCmd = &cobra.Command{Use: "create", Run: func(...) {...}}

func init() {
    rootCmd.AddCommand(userCmd)
    userCmd.AddCommand(createCmd)
}
```

This creates `app user create`. Each command manages its own flags. The `rootCmd` usually holds persistent flags (like `--verbose`) that apply to all children.

---

## 446. How do you manage color and styling in terminal output?

**Answer:**
We use `fatih/color`.

`color.Red("Error: %s", err)`
`color.Green("Success!")`

It automatically detects if the output is a TTY (terminal). If the user pipes output (`myapp | grep foo`), the library automatically disables color codes so the grep doesn't break on ANSI escape sequences. This behavior is crucial for "Good Citizen" CLI tools.

---

## 447. How would you stream CLI output like `tail -f`?

**Answer:**
We read from a channel or reader and print to `stdout` in a loop, often using **Carriage Return (`\r`)** to update simple status bars in place.

For true streaming (logs):
```go
scanner := bufio.NewScanner(reader)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```
If we need to handle Ctrl+C to stop the stream cleanly, we listen for `os.Interrupt` signal and break the loop.

---

## 448. How do you handle secrets securely in a CLI?

**Answer:**
We never store secrets in plain text config files.

1.  **Keyring Integration**: Use `99designs/keyring` to store tokens in the OS Keychain (Mac Keychain, Windows Credential Manager).
2.  **Env Vars**: `export GITHUB_TOKEN=...`.
3.  **Stdin**: `cat token.txt | myapp login --stdin`.

When prompting for a password, we ensure echo is disabled (terminal doesn't show characters) using `term.ReadPassword()`.

---

## 449. How do you bundle a CLI as a standalone binary?

**Answer:**
This is Go's superpower. `go build`.

You get a single file. No Python venv, no node_modules.
We distribute it via:
1.  **Homebrew** (Mac): Write a Formula pointing to the binary tarball.
2.  **Curl | Bash**: A script that detects OS/Arch, downloads the correct binary from GitHub Releases, and `chmod +x` it.
3.  **Scoop/Choco**: For Windows.

---

## 450. How would you version and release CLI with GitHub Actions?

**Answer:**
We tag the repo: `git tag v1.0.0`; `git push --tags`.

The GitHub Action triggers **GoReleaser**.
It builds binaries for generic Linux, RPM/Deb packages, macOS, and Windows.
It uploads them to the GitHub Release page.
It can also update the Homebrew Tap automatically.
The CLI itself checks for updates (by pinging the GitHub API) and can notify the user: "Update available: v1.0.1".

---

## 451. How do you schedule a Go CLI tool with cron?

**Answer:**
The Go tool itself should be ephemeral (run and exit).

We add it to the system crontab: `0 * * * * /usr/local/bin/myapp sync`.
Inside the Go app, we ensure:
1.  **Logging**: Write to stdout/stderr (Cron emails output) or syslog.
2.  **Locking**: Use a file lock (`/var/lock/myapp.lock`) so that if the previous run gets stuck, the next run aborts immediately to prevent piling up processes.

---

## 452. How do you use Go as a scripting language?

**Answer:**
We use **Go Scripts** via `gorun` or simply `go run main.go`.

For heavy automation (replacing Bash), Go is safer because of strict typing and error handling.
We use libraries like `bitfield/script` which mimics pipes:
`script.File("data.txt").Match("Error").CountLines()`.
This gives us the brevity of Bash with the safety and testability of Go.

---

## 453. How do you embed templates in your Go CLI tool?

**Answer:**
Use `//go:embed templates/*`.

When building a CLI that scaffolds projects (like `create-react-app`), we embed the boilerplate code inside the binary.
`var templates embed.FS`
At runtime, we walk `fs.WalkDir(templates, ...)` and copy files to the user's disk. This allows shipping a "Project Generator" as a single `.exe`.

---

## 454. How would you create a system daemon in Go?

**Answer:**
We use the `kardianos/service` library.

It abstracts Systemd (Linux), Launchd (Mac), and Service Control Manager (Windows).
Code: `service.New(prt, config)`.
Command: `myapp install` -> registers the service. `myapp start` -> starts it.
The Go program implements a `Start()` and `Stop()` interface, allowing it to run in the background managed by the OS init system.

---

## 455. What are good patterns for CLI testing?

**Answer:**
1.  **Separate `Main`**: Put logic in `Run(args []string, stdout io.Writer) error`. `main()` just calls `Run(os.Args, os.Stdout)`.
2.  **Test `Run`**: In tests, call `Run` with a `bytes.Buffer` as stdout. Assert the buffer content.
3.  **Golden Files**: For complex output, compare the buffer against a stored golden file.
4.  **Integration**: Use `testscript` (from the Go team) to write shell-like tests that actually execute your binary in a sandbox.

---

## 456. How do you store and manage CLI state/config files?

**Answer:**
We follow the XDG Base Directory specification.

Config: `~/.config/myapp/config.yaml`
Data: `~/.local/share/myapp/`
Cache: `~/.cache/myapp/`

We use `os.UserConfigDir()` to find the correct path per OS (Go handles Windows `%APPDATA%` vs Linux correctly). We creating these directories lazily on first run.

---

## 457. How do you secure a CLI for local system access?

**Answer:**
If the CLI needs `sudo` (e.g., modifying `/etc/hosts`), we check `os.Geteuid() == 0`.

If not, we fail fast: "Please run as root".
However, we prefer to minimize privilege. If only one command needs root, we check it there.
We also guard against **file permission** issues—when writing config files, use `0600` (Read/Write for owner only) so other users on the system cannot steal API tokens.

---

## 458. How do you test CLI tools across multiple OS in CI?

**Answer:**
GitHub Actions Strategy Matrix.

```yaml
strategy:
  matrix:
    os: [ubuntu-latest, windows-latest, macos-latest]
```
We run the build and tests on all three.
Go's `filepath.Join` handles slash differences (`/` vs `\`).
We must be careful with line endings (`\n` vs `\r\n`) in assertions—often we strip `\r` before comparing strings in tests.

---

## 459. How do you expose analytics and usage for a CLI?

**Answer:**
We add an **Opt-In** telemetry hook.

On command completion, we fire a non-blocking UDP packet or a short-timeout HTTP POST to our analytics server (PostHog/Mixpanel).
We count "Command Invocations" (`myapp deploy`).
**Privacy First**: We explicitly ask the user on first run: "Allow anonymous usage stats? [y/N]". If No, we write a config `telemetry: false` and never send data.

---

## 460. How would you build a CLI wrapper for REST APIs?

**Answer:**
We generate the client code from OpenAPI (`oapi-codegen`).

Then we map commands to API calls.
`myapp users list` -> `client.GetUsers()`.
We focus on **UX**:
1.  **Spinners**: Show "Loading..." while the API call is in flight.
2.  **JSON Output**: `--json` flag to print raw response for piping to `jq`.
3.  **Tables**: Print human-readable ASCII tables for default view.
The goal is to make the API explorable from the terminal.
