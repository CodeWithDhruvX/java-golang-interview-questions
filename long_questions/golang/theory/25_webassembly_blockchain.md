# 🟢 Go Theory Questions: 481–500 WebAssembly, Blockchain, and Experimental Go

## 481. What is WebAssembly and how can Go compile to WASM?

**Answer:**
WebAssembly (Wasm) allows running binary code in the browser at near-native speed.

Go compiles to Wasm natively: `GOOS=js GOARCH=wasm go build -o main.wasm`.
This produces a `.wasm` file and a glue JS file (`wasm_exec.js`).
We can run complex Go logic (image processing, crypto) inside Chrome, communicating with the DOM via the JS object bindings.

---

## 482. How do you share memory between JS and Go in WASM?

**Answer:**
Direct memory sharing is limited for security.
We use the `syscall/js` package.

To pass data:
1.  **Go -> JS**: `js.Global().Set("myVar", "Hello")` (Copying).
2.  **JS -> Go**: `js.Global().Call("myGoFunc", args)`.
For large data (like an image), we copy the bytes into the Go Wasm memory buffer (`uint8Array`), allowing Go to process the raw bits without serialization overhead.

---

## 483. What is TinyGo and what are its limitations?

**Answer:**
Standard Go Wasm binaries are huge (2MB+ "Hello World") because they include the full Runtime and GC.
**TinyGo** is a separate compiler based on LLVM.

It strips out everything unnecessary. "Hello World" becomes 10KB.
**Limitation**: It has a simplified Garbage Collector, slower reflection, and doesn't support the full standard library (though coverage is growing). It is perfect for WASM and Microcontrollers (Arduino/ESP32), but maybe not for a massive backend server app.

---

## 484. How do you write a smart contract simulator in Go?

**Answer:**
A Smart Contract is just a State Machine.
`func (s *State) Transfer(from, to, amount)`.

In Go, we define a struct `Account { Balance uint64 }`.
We write strict logic: "Check balance >= amount. Atomic Decrement. Atomic Increment."
We can simulate millions of transactions per second in pure Go to test the economic game theory (e.g., "Can a whale generate infinite money?") before deploying to the slow, expensive real blockchain.

---

## 485. What is Tendermint and how does Go power it?

**Answer:**
Tendermint (now CometBFT) is the engine behind Cosmos/Terra. It is written in Go.

It handles **Consensus** (BFT) and **P2P Networking**.
It connects to the actual App Logic via an interface called **ABCI** (Application BlockChain Interface).
Developers write the "Chain Logic" in Go (using Cosmos SDK), and Tendermint handles the hard part of ensuring all 100 validator nodes agree on the next block.

---

## 486. How do you use `go-ethereum` to interact with smart contracts?

**Answer:**
We use `abigen`.

1.  Compile Solidity to ABI.
2.  Run `abigen --abi MyContract.abi --pkg main --type MyContract --out bo.go`.
3.  This generates a strict Go binding.
In our code: `instance, _ := NewMyContract(address, backend)`.
`instance.Transfer(&bind.TransactOpts{...}, to, amount)`.
This gives us type-safe access to the blockchain methods, handling the RLP encoding/decoding automatically.

---

## 487. How do you parse blockchain data using Go?

**Answer:**
Blockchain data is a linked list of Blocks containing Merkle Trees of Transactions.

To parse Bitcoin/Ethereum:
We connect to a node (RPC/WS).
We fetch `GetBlockByNumber`.
The efficient way is to not use JSON but purely binary parsing (RLP for Eth, VarInt for BTC).
We iterate the transactions. Since Go handles hex/binary manipulation and concurrency well, Go indexers are standard for ingesting TBs of chain data into SQL for analytics.

---

## 488. How do you generate and verify ECDSA signatures in Go?

**Answer:**
We use `crypto/ecdsa` and `crypto/elliptic` (usually secp256k1 for Bitcoin/Eth).

**Sign**:
`r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hash)`
**Verify**:
`valid := ecdsa.Verify(publicKey, hash, r, s)`

This is the bedrock of crypto ownership. "Not your keys, not your coins." Go's standard library crypto is formally audited and considered one of the safest implementations available.

---

## 489. What is the role of Go in decentralized storage (IPFS)?

**Answer:**
IPFS (InterPlanetary File System) is written in Go (`kubo`).

It breaks files into **Chunks** (256KB).
It hashes them (CID - Content ID).
It uses a **DHT** (Distributed Hash Table, Kademlia) to find who has the chunks.
Go's concurrency allows a node to simultaneously download chunks from 50 peers, verify their hashes, and serve chunks to 50 other peers, saturating the bandwidth efficiently.

---

## 490. How would you implement a Merkle Tree in Go?

**Answer:**
A Merkle Tree verifies data integrity.

Leaf nodes = Hash(Data).
Parent = Hash(ChildL + ChildR).
Root = Top Hash.

Struct: `type Node struct { Left, Right *Node, Hash []byte }`.
We build it recursively.
To generate a **Proof**: Path of sibling hashes from Leaf to Root.
To **Verify**: `Hash(Data + Sibling1 + Sibling2...) == Root`.
This allows verifying one transaction exists in a block of 10,000 without downloading the whole block.

---

## 491. How do you handle base58 and hex encoding/decoding?

**Answer:**
**Hex**: `encoding/hex`. Used for raw bytes <-> string.
**Base58**: (Bitcoin addresses). Standard lib doesn't have it; we use `github.com/btcsuite/btcutil/base58`.

Base58 removes ambiguous chars (0, O, I, l) to avoid human error in typing addresses.
When decoding, strictly handle invalid checksums. A typo in an address must result in an error, effectively preventing sending money to a void.

---

## 492. How do you write a deterministic VM interpreter in Go?

**Answer:**
Deterministic means: "Run code X on any machine, get exact same result Y."

We cannot use `float64` (different precision on different CPUs). We use `int256` or `big.Int`.
We receive bytecode (opcodes).
Switch loop:
`case OP_ADD: stack.Push(a + b)`
`case OP_STORE: state[key] = val`
We enforce **Gas Limits** (count instructions) to prevent infinite loops (Halting Problem), forcing the execution to stop exactly when gas runs out.

---

## 493. How do you simulate a P2P network in Go?

**Answer:**
We use `libp2p`.

We create a **Host**.
`h1, _ := libp2p.New()`
`h2, _ := libp2p.New()`
`h1.Connect(ctx, h2.Addrs())`

This creates a virtual overlay network. We can simulate latency, packet loss, and churn (nodes dropping off).
We implement a **Gossip Protocol**: "I tell 5 friends, they tell 5 friends." Go's lightweight goroutines allow us to run 10,000 "nodes" in a single process for simulation testing.

---

## 494. How do you create a lightweight Go runtime for edge computing?

**Answer:**
We compile to **WASM** or use **TinyGo**.

On the Edge (Cloudflare Workers, Fastly Compute), we cannot run a full Linux binary.
We upload the `.wasm` file.
The Edge provider runs it.
Go's model is perfect here: Single binary, fast cold start.
We design the app to be **Ephemeral**: Handle 1 request, die. No background threads, no persistent connections.

---

## 495. How would you handle offline-first apps in Go?

**Answer:**
Usually, this implies a Mobile/Desktop app (Wails/Flutter+Go).

1.  **Local DB**: SQLite embedded in the Go app.
2.  **Queue**: Actions (POST /api/todo) go to a local "Outbox" table.
3.  **Sync**: A background Go worker watches network status. When online, it reads the Outbox, pushes to Server, and updates local state.
4.  **Conflict Resolution**: Use CRDTs (Conflict-free Replicated Data Types) or "Last Write Wins" timestamp logic.

---

## 496. What is the future of `Generics` in Go (beyond v1.22)?

**Answer:**
Go Generics are evolving towards **Iterators** (`range func`).

Currently, we have `Min[T Number](a, b T)`.
Future is functional patterns: `slices.Map`, `slices.Filter`, `ranges.Pull`.
This moves Go slightly towards functional programming, reducing the boilerplate of `for` loops, but the core team is extremely cautious to not turn Go into Haskell. Simplicity remains the north star.

---

## 497. What is fuzz testing and how do you use it in Go?

**Answer:**
Fuzzing throws random garbage at your function to find crashes.
Go 1.18 added native fuzzing.

```go
func FuzzParse(f *testing.F) {
    f.Add("sample input") // Seed
    f.Fuzz(func(t *testing.T, data []byte) {
        Parse(data) // Should not panic
    })
}
```
The fuzzer uses coverage feedback to mutate inputs intelligently. It finds edge cases (empty bytes, invalid utf8, giant numbers) that humans forget to write unit tests for.

---

## 498. What is the `any` type in Go and how is it different from `interface{}`?

**Answer:**
It is **Not Different**.

`type any = interface{}`
It is a type alias introduced in Go 1.18.
It exists solely for readability. `func Print(v any)` reads better than `func Print(v interface{})`.
We use `any` in new code; we leave legacy code as `interface{}` to avoid unnecessary churn, but they are binary compatible interchangable.

---

## 499. What is the latest experimental feature in Go and why is it important?

**Answer:**
**Range-Over-Functions** (Iterators).

Ideally, we want to `for item := range myCustomCollection { ... }`.
Experimental feature allows defining a standard Iterator protocol.
This allows custom containers (Trees, Linked Lists, database Cursors) to feel like native slices in `for` loops, standardizing iteration across the entire ecosystem and removing the need for `Next()/Scan()` boilerplate.

---

## 500. How do you contribute to the Go open-source project?

**Answer:**
Go uses **Gerrit** (go-review.googlesource.com), not GitHub Pull Requests.

1.  Sign the CLA (Google logic).
2.  Pick an issue from GitHub (labeled `help wanted` or `NeedsInvestigation`).
3.  Design proposal (RFC) for large changes.
4.  Commit code.
5.  Upload with `git codereview mail`.
The review process is rigorous. They care deeply about backward compatibility and simplicity. A contribution is 10% code, 90% discussion and tests.
