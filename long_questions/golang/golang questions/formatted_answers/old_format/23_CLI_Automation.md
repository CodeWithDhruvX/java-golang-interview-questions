# 🟢 **441–460: CLI Tools, Automation, and Scripting**

### 441. How do you build an interactive CLI in Go?
"I use **Bubble Tea** (Charm).

It implements the **ELM Architecture** for the terminal.
1.  **Model**: My application state.
2.  **View**: Renders the state to a string.
3.  **Update**: Handles keyboard events and modifies the state.
It allows me to build rich, 60fps TUI apps (like lists, spinners, and forms) that feel like native GUI applications."

#### Indepth
Bubble Tea is based on The Elm Architecture (Model-Update-View). This makes it strictly deterministic. However, handling **Async Commands** (like HTTP requests) requires returning a `tea.Cmd` from the `Update` function. This command runs in a separate goroutine and sends a `Msg` back to the `Update` loop when finished.

---

### 442. What libraries do you use for command-line tools in Go?
"**Cobra**: The standard for structure (Flags, Subcommands). used by Kubernetes/Docker.
**Viper**: For configuration (YAML/Env/Flags).
**Bubble Tea / Lip Gloss**: For interactive UI and styling.
**Pterm**: For simple printers (tables, progress bars) if Bubble Tea is overkill."

#### Indepth
Don't use `fmt.Print` in a Bubble Tea app! It corrupts the TUI buffer. Use the `tea.Program`'s output or `log` to a file. If you need to print a final result *after* the program exits (like `jq`), return the string in the Model and print it in `main` after `p.Run()` returns.

---

### 443. How do you parse flags and config in CLI?
"I use **Cobra** binding.

`rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")`.
Then inside `RunE`, I verify arguments.
For validation (e.g., 'port must be number'), I do it explicitly.
Most importantly, Cobra auto-generates `--help`, which is a UX requirement."

#### Indepth
Cobra supports **Persistent Flags** (`rootCmd.PersistentFlags()`) which filter down to *all* subcommands. Use this for global toggles like `--verbose` or `--json`. Also, use `viper.BindPFlag` so that `viper.Get("verbose")` works whether the user set the flag OR the environment variable.

---

### 444. How do you implement bash autocompletion for Go CLI?
"Cobra generates it for free!

`rootCmd.GenBashCompletion(os.Stdout)`.
I verify by running: `source <(my-cli completion bash)`.
Users get tab-completion for subcommands and even flags. I can also add custom dynamic completion (e.g., fetching a list of Kubernetes pods) by implementing the `ValidArgsFunction`."

#### Indepth
For `zsh` users, you need to generate zsh completion scripts (`GenZshCompletion`). A common pattern is to hide a `completion` subcommand that outputs these scripts so users can simply add `source <(my-cli completion zsh)` to their `.zshrc`.

---

### 445. How would you use `cobra` to build a nested command CLI?
"It’s a command tree.

`rootCmd.AddCommand(userCmd)`.
`userCmd.AddCommand(createCmd)`.
This gives me the structure: `my-cli user create`.
Each command is a struct with `Use`, `Short`, and `Run`. This organization keeps `main.go` tiny and separates the logic for 'User' and 'Product' into different packages."

#### Indepth
Structure is key. Put commands in `cmd/root.go`, `cmd/user.go`. Avoid global variables. Pass a `Context` or a `dependencies` struct (DB connection, Logger) to the `RunE` method of your commands via a closure or a struct method receiver, so your CLI remains testable.

---

### 446. How do you manage color and styling in terminal output?
"I use **Lip Gloss**.

It’s 'CSS for the Terminal'.
`var style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA"))`.
`fmt.Println(style.Render("Hello"))`.
It automatically detects if the user's terminal supports TrueColor or just ANSI 16-color, and degrades gracefully. I never hardcode ANSI escape codes (`\033[31m`) anymore."

#### Indepth
Detecting "Is Terminal" is crucial. If the user pipes your output (`my-cli | grep foo`), you should **disable** colors automatically. Check `term.IsTerminal(int(os.Stdout.Fd()))`. Lip Gloss handles this, but if you do manual coloring, use `fatih/color` which respects the `NO_COLOR` standard.

---

### 447. How would you stream CLI output like `tail -f`?
"I read and flush.

`reader := bufio.NewReader(file)`.
I loop:
`line, err := reader.ReadString('\n')`.
If `err == io.EOF`, I sleep for 100ms and try again (polling).
To display it, I write to `os.Stdout`. If I need to overwrite the current line (like a spinner), I print `\r` (Carriage Return) before the new text."

#### Indepth
Be nice to the CPU! A tight loop reading `stdin` or polling a file can hit 100% CPU. always use a `select` with a `time.Ticker` or file watcher (`fsnotify`) to wait for changes efficiently. For `tail -f`, `fsnotify` is far superior to polling.

---

### 448. How do you handle secrets securely in a CLI?
"I never pass secrets as flags (`--password=123`).
That shows up in `bash_history` and `ps aux`.

I accept them via:
1.  **Environment Variables**: `MYAPP_PASSWORD=123 my-cli`.
2.  **Stdin**: `cat pass.txt | my-cli`.
3.  **Keyring**: I use `zalando/go-keyring` to store credentials securely in the OS's native Keychain (Mac) or Credential Manager (Windows)."

#### Indepth
If you must accept a password via flag (e.g. for automation scripts), provide a `--password-stdin` flag (like Docker). This allows `cat pass.txt | my-cli --password-stdin`, which keeps the secret out of the process arguments list and shell history.

---

### 449. How do you bundle a CLI as a standalone binary?
"That's Go's killer feature.

`CGO_ENABLED=0 go build -o my-cli`.
Result: A single, static binary.
No `node_modules`, no Python venv, no shared libraries.
I distribute this file. The user hits `chmod +x` and runs it. It works on Alpine, Debian, Centos—everywhere."

#### Indepth
Embed loop! `CGO_ENABLED=0` creates a static binary. But if you rely on `net` package + DNS, on Linux it *might* still use C library versions if they exist. Use `-tags netgo -ldflags '-extldflags "-static"'` to be 100% sure it's a static binary that runs on a generic Alpine container.

---

### 450. How would you version and release CLI with GitHub Actions?
"I use **GoReleaser**.

I push a tag `v1.0.0`.
GoReleaser detects it and:
1.  Cross-compiles for Linux/Mac/Windows (amd64/arm64).
2.  Creates a GitHub Release with artifacts.
3.  Updates my Homebrew Tap (`brew install my-cli`).
It automates the entire distribution pipeline."

#### Indepth
Sign your binaries! Code Signing is becoming mandatory (Apple Gatekeeper). GoReleaser integrates with `gon` or `cosign`. If you don't sign your Mac binary, users will get a "Developer cannot be verified" warning and likely trash your app.

---

### 451. How do you schedule a Go CLI tool with cron?
"The CLI is just a process.

I add it to system crontab: `0 * * * * /usr/local/bin/my-cli sync`.
If I want the Go app *itself* to be a long-running scheduler (daemon), I use **robfig/cron** library.
`c := cron.New(); c.AddFunc("@hourly", func() { ... })`.
This is useful inside Docker where system cron is missing."

#### Indepth
Distributed Cron? If you run 3 replicas of your app, `robfig/cron` will run the job 3 times! You need a **Distributed Lock** (Redis/Postgres). `if lock.Acquire() { job() }`. Or use a dedicated scheduler like Airflow/Temporal if complexity grows.

---

### 452. How do you use Go as a scripting language?
"I use `go run main.go`.

For single-file scripts, I use a shebang: `///usr/bin/env go run`.
However, if the script grows > 100 lines, I create a proper `go.mod` project. Writing complex automation in Go is safer than Bash (Type Safety, Error Handling, Testing) but more verbose."

#### Indepth
For "Scripting", look at **Go-Script** or **Bar** (build-and-run) tools. But honestly, compile times are so fast that `go run` is usually sufficient. A neat trick: `go run .` works if you have multiple files in `package main` in the current folder.

---

### 453. How do you embed templates in your Go CLI tool?
"I use `//go:embed`.

`//go:embed templates/*`
`var content embed.FS`.
I can generate a starter project for the user:
`my-cli init`.
The CLI reads the files from its own binary and writes them to the user's disk. This makes the CLI a completely self-contained generator without external dependencies."

#### Indepth
`embed` also supports the `http.FileSystem` interface. `http.FileServer(http.FS(content))`. You can embed an entire React/Vue Single Page App into your Go binary and serve it from memory. This is how tools like Grafana or Prometheus are distributed as single binaries.

---

### 454. How would you create a system daemon in Go?
"I use `kardianos/service`.

It abstracts systemd (Linux), Launchd (Mac), and Windows Services.
I define a `Program` struct with `Start()` and `Stop()` methods.
`s, _ := service.New(prg, config); s.Install()`.
This allows my Go app to install itself as a service that starts automatically on boot."

#### Indepth
Daemonizing correctly involves handling OS signals (`SIGTERM`, `SIGINT`). Your `Stop()` method should cancel a global context, allowing all running goroutines to clean up (close DB, flush logs) before the process exits. Hard kills lead to data corruption.

---

### 455. What are good patterns for CLI testing?
"I separate Logic from `main()`.

Instead of putting code in `main`, I have `func Run(args []string, stdout io.Writer) error`.
In tests, I pass a `bytes.Buffer` as stdout.
`err := Run([]string{"--dry-run"}, buf)`.
Then I assert `buf.String()` contains the expected output. This tests the wiring without needing to spawn a child process."

#### Indepth
Golden Files! CLI output often changes (formatting, spaces). Instead of `assert.Equal(t, "expected...", got)`, write the output to `testdata/output.golden`. In the test, compare `got` vs file content. Use a flag `go test -update` to overwrite the file when you intentionally change the format.

---

### 456. How do you store and manage CLI state/config files?
"I use `os.UserConfigDir()` (e.g., `~/.config/my-cli/`).

I save `config.yaml` there.
Viper handles reading it automatically.
I ensure I verify permissions (`0600`) if storing tokens. I use `os.MkdirAll` on startup to ensure the directory exists."

#### Indepth
Check `XDG_CONFIG_HOME`. Linux standards say config goes there, not hardcoded `~/.config`. `os.UserConfigDir()` handles this mostly, but being fully XDG compliant (`CACHE_HOME`, `DATA_HOME`) makes your CLI feel more "Pro" and native to Linux power users.

---

### 457. How do you secure a CLI for local system access?
"**Principle of Least Privilege**.

If my CLI needs to edit `/etc/hosts`, I check `os.Geteuid()`. If not root, I fail gracefully: 'Please run with sudo'.
But I try to avoid needing root.
I also validate all file paths to prevent **Symlink Attacks** (writing to a user-controlled link that points to `/etc/shadow`)."

#### Indepth
If you *must* run as root (e.g., a VPN client), try to **Drop Privileges**. Start as root, open the raw socket, then switch the process effective UID to the user `syscall.Setuid(uid)`. This minimizes the window where a bug could compromise the whole system.

---

### 458. How do you test CLI tools across multiple OS in CI?
"**GitHub Actions Matrix**.

`runs-on: [ubuntu-latest, macos-latest, windows-latest]`.
I run the Go tests on all three.
**Pain Point**: File Paths. I strictly use `filepath.Join` instead of hardcoding `/` or `\` to ensure compatibility."

#### Indepth
Line Endings! Windows uses `\r\n`, Linux `\n`. If your golden files use `\n`, your tests might fail on Windows. Use `strings.ReplaceAll(got, "\r\n", "\n")` in your test helpers to normalize output before comparison.

---

### 459. How do you expose analytics and usage for a CLI?
"I use a **Privacy-Preserving Ping**.

On execution, I fire a 'fire-and-forget' UDP packet or async HTTP POST to my telemetry server: `{'cmd': 'deploy', 'os': 'linux'}`.
I **MUST** ask for opt-in permission on the first run.
I wrap it in a short timeout (500ms) so analytics *never* slow down the user's experience."

#### Indepth
Respect `DO_NOT_TRACK` env var. If `os.Getenv("DO_NOT_TRACK") != ""`, disable analytics silently. Trust is hard to gain and easy to lose. Also, explicitly state what you collect in your `--help` or `README`.

---

### 460. How would you build a CLI wrapper for REST APIs?
"I generate the client using **OpenAPI Generator**.

Then I create Cobra commands for each endpoint.
`my-cli get users` -> `client.GetUsers()`.
I focus heavily on the Output Formatting. I use `tablewriter` to render the JSON response as a pretty ASCII table, which is much nicer for humans than raw JSON."

#### Indepth
JSON output flag is mandatory. `my-cli get users --json`. Power users want to pipe your output into `jq`. If you only output ASCII tables, they have to use `awk`, which is fragile. Always provide a machine-readable bypass for your human-readable output.
