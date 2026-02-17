# 🧰 Go Theory Questions: 881–900 Refactoring, CLI, WebAssembly & Design

## 881. How do you build CLI apps with Cobra?

**Answer:**
**Cobra** is the standard.
Structure:
- `rootCmd`: The base entry point.
- `AddCommand`: `root.AddCommand(deployCmd)`.
- Flags: `cmd.Flags().StringVar(&region, "region", "us-east-1", "AWS Region")`.
- Run: `if err := rootCmd.Execute(); err != nil { os.Exit(1) }`.
It handles `--help`, subcommands, and POSIX flags automatically.

---

## 882. How do you add auto-completion to CLI tools?

**Answer:**
(See Q 798).
Cobra generates it.
`completionCmd`.
Bash: `source <(myapp completion bash)`.
Zsh: `source <(myapp completion zsh)`.
To make it dynamic (e.g., complete server names from API):
`cmd.RegisterFlagCompletionFunc("server", func(...) { return ListServers(), nil })`.

---

## 883. How do you handle subcommands in CLI tools?

**Answer:**
Nesting.
`kubectl get pods`.
`getCmd` is added to `rootCmd`.
`podsCmd` is added to `getCmd`.
Cobra executes the "leaf" command's Run function.
Persistent Flags (global flags like `--verbose`) defined on Root are inherited by all children.

---

## 884. How do you package Go binaries for multiple platforms?

**Answer:**
**GoReleaser** (See Q 782).
Or manual:
Make a `Makefile`.
```makefile
build:
    GOOS=linux GOARCH=amd64 go build -o bin/app-linux .
    GOOS=windows GOARCH=amd64 go build -o bin/app.exe .
    GOOS=darwin GOARCH=arm64 go build -o bin/app-mac .
```
Zip them up and upload to S3/GitHub Releases.

---

## 885. How do you write a Wasm frontend in Go?

**Answer:**
1.  **Build**: `GOOS=js GOARCH=wasm go build -o main.wasm`.
2.  **Glue**: Copy `$(go env GOROOT)/misc/wasm/wasm_exec.js`.
3.  **HTML**: Load `wasm_exec.js`, then `WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(go.run)`.
Go code starts in `main()` and can block forever using `select{}` to keep the browser app running.

---

## 886. How do you expose Go functions to JS using Wasm?

**Answer:**
`syscall/js`.
`js.Global().Set("myGoFunc", js.FuncOf(wrapper))`
Wrapper: `func(this js.Value, args []js.Value) any`.
Now in Browser Console: `myGoFunc("hello")`.
**Caveat**: `js.FuncOf` creates a goroutine/resource. You must release it if you unload, otherwise memory leak.

---

## 887. How do you reduce Wasm binary size?

**Answer:**
Standard Go Wasm is huge (~10MB) because it includes the entire Runtime (GC, Scheduler).
1.  **Strip**: `-ldflags="-s -w"` (-2MB).
2.  **Zip**: Serve with Gzip/Brotli (2MB -> 500KB).
3.  **TinyGo**: Replaces standard runtime.
`tinygo build -o main.wasm -target=wasm .`.
Result: 10KB - 100KB. Perfect for simple web modules.

---

## 888. How do you interact with DOM from Go Wasm?

**Answer:**
`doc := js.Global().Get("document")`.
`btn := doc.Call("getElementById", "submitBtn")`.
`btn.Set("innerText", "Processing...")`.
It's verbose. Wrapper libraries like **go-app** or **vecty** provide a React-like component model to abstract raw DOM calls.

---

## 889. How do you debug Go WebAssembly apps?

**Answer:**
Hard. No Delve support in browser.
1.  **Console**: `fmt.Println` goes to Browser Console.
2.  **Panic**: Stack traces appear in Console.
3.  **Logic**: Test the business logic in standard Go unit tests (`_test.go`) first. Only the UI binding layer needs browser testing.

---

## 890. How do you build a WebAssembly module loader?

**Answer:**
If you want to load Go Wasm as a library (not main app).
Go's Wasm is designed to "Take over the world" (runs as main).
To run as a module/worker:
Use `TinyGo`. Export functions using `//export add`.
In JS: `instance.exports.add(1, 2)`.
Standard Go doesn't support generic Wasm exports well yet (WASI is the future standard here).

---

## 891. How do you manage state in Go WebAssembly apps?

**Answer:**
Global struct or Central Store (Redux pattern).
Since Wasm is single-threaded (in UI thread), we don't need Mutexes for state accessed only by UI callbacks.
We render output based on state changes.
`func update() { render(state) }`.

---

## 892. How do you integrate Go Wasm with JS promises?

**Answer:**
Go is blocking/sync. JS is async/Promise.
To await a JS Promise in Go:
Make a channel.
`ch := make(chan js.Value)`
`promise.Call("then", js.FuncOf(func(_,_ args) { ch <- args[0] }))`
`result := <-ch`
This bridges the gap, allowing Go to write linear code while JS executes the async task.

---

## 893. How do you decide between Go CLI and REST tool?

**Answer:**
**Context**.
- **CLI**: Human operator, scripting, local files, interactive prompts.
- **REST**: Machine-to-machine, remote access, browser clients.
Often we build **Both**.
Core Logic -> Library.
CLI -> Wrapper around Library.
REST API -> Wrapper around Library.

---

## 894. How do you document CLI help and usage info?

**Answer:**
Cobra does it.
`Short`, `Long`, `Example` fields in `cobra.Command`.
`Long` supports Markdown-like text.
`Example: "  myapp deploy --prod"`
Running `myapp --help` formats this beautifully.
We can also generate Man Pages: `doc.GenManTree(rootCmd, header, "./man")`.

---

## 895. How do you call an OpenAI API using Go?

**Answer:**
JSON HTTP Client.
`sashabaranov/go-openai` library.
`resp, err := client.CreateChatCompletion(...)`.
It handles the JSON marshaling of messages `[]Message{{Role: "user", Content: "..."}}`.
Context handling is key (`ctx`) to cancel long-running LLM generation requests.

---

## 896. How do you stream ChatGPT responses in Go?

**Answer:**
**Server-Sent Events (SSE)**.
OpenAI API returns a stream.
`stream, _ := client.CreateChatCompletionStream(...)`.
Loop:
`resp, err := stream.Recv()`.
In our Go HTTP Handler:
`w.Header().Set("Content-Type", "text/event-stream")`.
`fmt.Fprintf(w, "data: %s\n\n", resp.Choices[0].Delta.Content)`.
`flusher.Flush()`.

---

## 897. How do you build a Telegram AI bot in Go?

**Answer:**
Library: `go-telegram-bot-api/telegram-bot-api`.
Long Polling (easiest for local dev):
`u := tgbotapi.NewUpdate(0); u.Timeout = 60`.
`updates := bot.GetUpdatesChan(u)`.
Loop `for update := range updates`.
Send text to OpenAI.
Send reply: `msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)`.

---

## 898. How do you integrate Go with HuggingFace models?

**Answer:**
1.  **Inference API**: HTTP Call (easier).
2.  **Local Execution**: Bindings for `onnxruntime` or `tensorflow`.
    We load the `.onnx` model exported from PyTorch.
    `session, _ := onnxruntime.NewSession(modelPath)`.
    `output, _ := session.Run(inputTensor)`.
    This allows Go to run models locally without Python dependencies (Deployment friendly).

---

## 899. How do you implement Feature Flags in Go?

**Answer:**
Simple: Env Var. `if os.Getenv("FEATURE_X") == "true"`.
Advanced: **LaunchDarkly** or **Go-Feature-Flag**.
Evaluate flags per user context.
`variation, _ := client.BoolVariation("new-ui", user, false)`.
This allows canary rollouts (10% of users) and instant kill-switches.

---

## 900. How do you handle configuration hot-reloading in Go?

**Answer:**
**Viper**.
`viper.WatchConfig()`.
`viper.OnConfigChange(func(e fsnotify.Event) { ... })`.
When `config.yaml` changes, Viper detects it.
However, updating running structs is thread-sensitive. We usually protect config access with `RWMutex` or use `atomic.Value` to swap the entire config struct pointer atomically.
