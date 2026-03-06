# 🗣️ Theory — DevOps & Deployment in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are the main build commands in Go and when do you use each?"

> *"There are three main ones. `go run` compiles and runs in one step — great for quick experiments but produces no binary. `go build` compiles and produces a binary. You can specify the output name: `go build -o server ./cmd/main.go`. In production, I'd add `-ldflags='-s -w'` to strip debug symbols and reduce binary size. `go install` compiles and installs the binary to your `$GOPATH/bin`. Cross-compilation is built-in — for Linux from Windows you just set `GOOS=linux GOARCH=amd64`. No toolchain needed. That's one of Go's killer features."*

---

## Q: "How do you Dockerize a Go application? What is a multi-stage build?"

> *"The best practice for Go Docker images is a multi-stage build. In the first stage — the builder — you use the full golang image, copy the source, and compile. In the second stage, you start from a minimal base image like `alpine` or even `scratch`, and only copy the compiled binary. This way your final image doesn't contain the Go toolchain, source code, or build dependencies — just the binary and what it needs to run. The resulting image can be as small as 5–20 MB versus 800+ MB for a full golang image. One important flag: `CGO_ENABLED=0` ensures a statically linked binary that works in minimal containers."*

---

## Q: "How do Go modules work? What are `go.mod` and `go.sum`?"

> *"Go modules are Go's dependency management system, introduced in Go 1.11. `go.mod` is the module manifest — it declares the module's name and the minimum version of each dependency. `go.sum` contains cryptographic hashes of every module used, ensuring that what you download matches what was committed. You never edit `go.sum` manually. Key commands: `go mod init` to create a new module, `go get` to add or update a dependency, `go mod tidy` to remove unused dependencies and update both files. In CI/CD, you usually run `go mod download` once to cache dependencies, then build."*

---

## Q: "What is `go vet` and `golangci-lint`? Why are they important?"

> *"`go vet` is built into Go — it's a static analysis tool that catches common mistakes the compiler doesn't catch, like calling `Printf` with the wrong format verbs, or closing over a loop variable incorrectly. Run it as part of every CI pipeline. `golangci-lint` is a community tool that aggregates 50+ different linters in one tool — it includes `go vet` plus things like `errcheck` which catches unchecked errors, `gosec` which finds security issues, `staticcheck` which finds bugs and dead code, and many more. Teams define a `.golangci.yml` config to pick which linters to enforce. It's the standard linting tool in the Go ecosystem."*

---

## Q: "What is `go generate` and what is it used for?"

> *"`go generate` scans source files for `//go:generate` directives and runs the specified commands. It's not a build step — you run it explicitly when you need to regenerate code. Common uses: generating mock implementations from interfaces using tools like `mockery`, generating the `String()` method for enum-like types with `stringer`, compiling protobuf files, or generating boilerplate code from templates. The generated code is checked into version control, not regenerated on every build. It's Go's answer to the question of code generation without a separate build tool."*
