## 🟡 WebAssembly, Blockchain, and Experimental Go (Questions 481-500)

### Question 481: What is WebAssembly and how can Go compile to WASM?

**Answer:**
WebAssembly (WASM) is a binary instruction format for a stack-based virtual machine, designed as a portable target for compiling high-level languages like C++/Rust/Go for deployment on the web.

To compile Go to WASM:
```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```
You also need the `wasm_exec.js` glue code (provided in the Go installation) to run it in the browser.

---

### Question 482: How do you share memory between JS and Go in WASM?

**Answer:**
Direct memory sharing is limited. Go (via `syscall/js`) provides `CopyBytesToGo` and `CopyBytesToJS`.
You move data by copying it across the boundary.

```go
import "syscall/js"

func myFunc(this js.Value, args []js.Value) interface{} {
    input := make([]byte, args[0].Length())
    js.CopyBytesToGo(input, args[0])
    // process input...
    return nil
}
```

---

### Question 483: What is TinyGo and what are its limitations?

**Answer:**
**TinyGo** is an alternate Go compiler used for:
1.  **Microcontrollers:** (Arduino, ESP32) where standard Go runtime is too heavy (GC, Scheduling).
2.  **WASM:** Produces much smaller binaries (100KB vs 2MB+ for standard Go).

**Limitations:**
- No complete standard library (some `net/http` features missing).
- GC is simpler (slower allocation).
- Recovering from panics might not work everywhere.

---

### Question 484: How do you write a smart contract simulator in Go?

**Answer:**
Define a state machine struct and transaction methods that mutate it.

```go
type Contract struct {
    Balances map[string]int
}

func (c *Contract) Transfer(from, to string, amount int) error {
    if c.Balances[from] < amount { return fmt.Errorf("insufficient funds") }
    c.Balances[from] -= amount
    c.Balances[to] += amount
    return nil
}
```
This mimics Ethereum's state transition function ($$\sigma' = \Upsilon(\sigma, T)$$) without the network overhead.

---

### Question 485: What is Tendermint and how does Go power it?

**Answer:**
**Tendermint (now CometBFT)** is a state-of-the-art BFT (Byzantine Fault Tolerant) engine written in Go.
- It handles p2p networking and consensus.
- Developers write the application logic in any language (often Go) via **ABCI** (Application BlockChain Interface).
- Go's concurrency is crucial for handling thousands of peer connections and consensus votes.

---

### Question 486: How do you use go-ethereum to interact with smart contracts?

**Answer:**
1.  **abigen:** Generate Go bindings from the Solidity ABI.
    `abigen --abi=token.abi --pkg=token --out=token.go`
2.  **Connect:** Use `ethclient`.
3.  **Transact:**

```go
conn, _ := ethclient.Dial("https://mainnet.infura.io")
token, _ := NewToken(address, conn)
bal, _ := token.BalanceOf(nil, userAddress)
```

---

### Question 487: How do you parse blockchain data using Go?

**Answer:**
Blockchain data is usually RLP (Recursive Length Prefix) or Protocol Buffers.
`go-ethereum` provides types for handling Blocks, Transactions, and Receipts.

```go
header, _ := client.HeaderByNumber(context.Background(), nil)
block, _ := client.BlockByNumber(context.Background(), header.Number)

for _, tx := range block.Transactions() {
    fmt.Println(tx.Hash().Hex(), tx.Value())
}
```

---

### Question 488: How do you generate and verify ECDSA signatures in Go?

**Answer:**
Use `crypto/ecdsa` and `crypto/elliptic` (usually secp256k1 for Bitcoin/Eth).

```go
// Sign
privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
hash := sha256.Sum256([]byte("data"))
r, s, _ := ecdsa.Sign(rand.Reader, privKey, hash[:])

// Verify
valid := ecdsa.Verify(&privKey.PublicKey, hash[:], r, s)
```

---

### Question 489: What is the role of Go in decentralized storage (IPFS)?

**Answer:**
**IPFS (InterPlanetary File System)** reference implementation (`kubo`) is written in Go.
Go was chosen for:
- **Concurrency:** Handling massive peer/swarm connections.
- **Performance:** CPU-bound crypto hashing (hashing file chunks).
- **Portability:** Running on servers, desktops, and IoT.

---

### Question 490: How would you implement a Merkle Tree in Go?

**Answer:**
A Merkle Tree is a binary tree where every node is the hash of its children. Root is the "fingerprint" of all data.

```go
func NewMerkleNode(left, right *Node, data []byte) *Node {
    node := &Node{}
    if left == nil && right == nil {
        hash := sha256.Sum256(data)
        node.Data = hash[:]
    } else {
        prevHashes := append(left.Data, right.Data...)
        hash := sha256.Sum256(prevHashes)
        node.Data = hash[:]
    }
    node.Left = left
    node.Right = right
    return node
}
```

---

### Question 491: How do you handle base58 and hex encoding/decoding?

**Answer:**
- **Hex:** `encoding/hex`.
- **Base58:** (Used in Bitcoin addresses to avoid look-alike chars 0OIl). Not in std lib. Use `github.com/btcsuite/btcutil/base58`.

```go
encoded := base58.Encode([]byte("hello"))
decoded := base58.Decode(encoded)
```

---

### Question 492: How do you write a deterministic VM interpreter in Go?

**Answer:**
Determinism means same input + same code = same output, always (no `time.Now()`, no `random`, no map iteration order dependency).
1.  **Bytecode:** Define opcodes (`ADD`, `PUSH`, `STORE`).
2.  **Stack:** Implement a `[]uint64` stack.
3.  **Loop:** `for pc < len(code) { switch code[pc] { ... } }`.
4.  **Strictness:** Error immediately on underflow across all nodes.

---

### Question 493: How do you simulate a P2P network in Go?

**Answer:**
Use **libp2p** (the networking stack of IPFS).
1.  Create a **Host** (Identity + Keys).
2.  Connect to peers using **Multiaddr** (`/ip4/127.0.0.1/tcp/4001/p2p/QmHash`).
3.  Set up **StreamHandlers** (like HTTP handlers but for bidirectional binary streams).

---

### Question 494: How do you create a lightweight Go runtime for edge computing?

**Answer:**
1.  **Strip DWARF:** `go build -ldflags="-s -w"`.
2.  **Use TinyGo**.
3.  **Disable CGO:** `CGO_ENABLED=0` for pure static binaries.
4.  **Compress:** Use `upx` on the binary.
This reduces a 50MB app to <2MB, suitable for Edge/Lambda@Edge.

---

### Question 495: How would you handle offline-first apps in Go?

**Answer:**
1.  **Local DB:** Use embedded SQLite (`modernc.org/sqlite` - no CGO) or BoltDB.
2.  **Sync Queue:** writes go to Local DB + Queue table.
3.  **Reconciler:** When online, background worker pushes Queue to Server.
4.  **Conflict Resolution:** Use "Last Writer Wins" or CRDTs (Conflict-free Replicated Data Types).

---

### Question 496: What is the future of Generics in Go (beyond v1.22)?

**Answer:**
(Theoretical)
- **Generic Allocations:** Optimizing compiler to monomorphize types better (less overhead).
- **Methods on Generic Types:** Allowing methods to introduce *new* type parameters (currently not allowed).
- **Iterator Pattern:** `func(yield func(V) bool)` standard (Added in Go 1.23 as `iter` package).

---

### Question 497: What is fuzz testing and how do you use it in Go?

**Answer:**
Automated testing that inputs random data to find crashes.
Built-in since Go 1.18.

```go
func FuzzParser(f *testing.F) {
    f.Add("sample input") // Corpus
    f.Fuzz(func(t *testing.T, data string) {
        // Did we crash?
        Parse(data)
    })
}
```
Run `go test -fuzz=FuzzParser`.

---

### Question 498: What is the `any` type in Go and how is it different from `interface{}`?

**Answer:**
`any` is literally an alias for `interface{}`.
```go
type any = interface{}
```
It was introduced in Go 1.18 to make code with Generics more readable. `[T any]` reads better than `[T interface{}]`. Behavior is identical.

---

### Question 499: What is the latest experimental feature in Go?

**Answer:**
(As of Go 1.22/1.23 context):
- **Range over Function (Iterators):** `for v := range mySeq`.
- **Loop Variable Fix:** `for i, v := range list` now creates a *new* `v` per iteration (fixing the common `&v` bug).
- **Profile Guided Optimization (PGO):** Compiler uses prod profile data to optimize hot paths.

---

### Question 500: How do you contribute to the Go open-source project?

**Answer:**
1.  **Gerrit:** Go does not use GitHub PRs for core. It uses Gerrit.
2.  **Steps:**
    - Sign CLA.
    - `git codereview change`.
    - Mail to `go.googlesource.com`.
3.  **Proposals:** For new features, open a GitHub Issue titled `proposal: ...`.

---
