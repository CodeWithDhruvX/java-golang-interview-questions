# 🟡 **481–500: WebAssembly, Blockchain, and Experimental**

### 481. What is WebAssembly and how can Go compile to WASM?
"**WebAssembly (WASM)** is a binary instruction format for a stack-based virtual machine. It runs in browsers at near-native speed.
I compile Go to it natively: `GOOS=js GOARCH=wasm go build -o main.wasm`.
It requires a small JS glue file (`wasm_exec.js`) to load. It allows me to share backend logic (like validation rules) with the frontend."

#### Indepth
WASM is a 32-bit architecture (mostly). Pointers are `uint32`. Be careful passing 64-bit integers from JS to Go; they might get truncated or split into high/low words depending on the binder. Go's `js.Value` handles this abstraction, but for raw performance, use direct memory access.

---

### 482. How do you share memory between JS and Go in WASM?
"Direct sharing is tricky due to isolation.
I use `syscall/js`.
Go -> JS: `js.Global().Set("result", value)`.
JS -> Go: `js.Global().Get("userInput")`.
For large data (images), I write to the WASM linear memory and pass the pointer + length to JS, which reads it as a `Uint8Array`. This avoids serialization overhead."

#### Indepth
Memory in WASM is a single linear array buffer. Go's GC lives *inside* this buffer. If your WASM module runs out of memory, it crashes the tab (OOM). Use `debug.SetGCPercent` to tune aggressively for low memory environments, or use TinyGo which has a simpler allocator.

---

### 483. What is TinyGo and what are its limitations?
"**TinyGo** is an LLVM-based compiler for embedded systems and WASM.
**Pros**: Produces tiny binaries (10KB vs standard Go's 2MB).
**Cons**: Limited standard library (no `net/http` server), simpler GC, partial reflection support.
I use it for IoT devices (Arduino) or ultra-lightweight WASM modules where download size is critical."

#### Indepth
Reflect is the enemy of TinyGo. `encoding/json` relies heavily on reflection. If you use `json.Unmarshal` in TinyGo, your binary size explodes. Use `easyjson` or `fastjson` (code generation) to keep the WASM binary under 500KB.

---

### 484. How do you write a smart contract simulator in Go?
"I define a **State Machine**.
`type Contract struct { Balance map[string]int }`.
`func (c *Contract) Transfer(from, to string, amount int)`.
I wrap this in a loop that processes batches of transactions ('Blocks').
Since blockchain logic is deterministic, I can unit test my smart contracts in pure Go without ever spinning up a real node."

#### Indepth
Fuzzing smart contracts is non-negotiable. Use Go's native fuzzer to bombard your `Transfer` function with random edge cases (negative amounts, overflows, self-transfers). Logic bugs in contracts are immutable and catastrophic; Go's type safety + fuzzing is a strong defense.

---

### 485. What is Tendermint and how does Go power it?
"**Tendermint** (CometBFT) is a state-machine replication engine written in Go.
It handles P2P networking and Consensus (PBFT).
I verify to write the **Application Logic** (ABCI app) in Go.
Tendermint ensures that my app executes the same transactions in the same order on every machine. It powers the Cosmos ecosystem."

#### Indepth
Tendermint uses **ABCI** (Application BlockChain Interface). It's just a socket protocol (like HTTP). You can implement your blockchain in *any* language, but Go is the "native" tongue. The critical rule: **Determinism**. Never use `map` iteration order or random numbers in your state machine, or the chain will fork.

---

### 486. How do you use `go-ethereum` to interact with smart contracts?
"I use `abigen`.
1.  Compile Solidity to ABI.
2.  Run `abigen --abi=MySource.abi --pkg=main --out=MySource.go`.
3.  This generates a Go struct with methods (`contract.Transfer`).
I connect via `ethclient.Dial`, and call the methods. It handles RLP encoding, signing, and broadcasting automatically."

#### Indepth
Gas estimation is tricky. `client.EstimateGas` simulates the transaction. Always add a buffer (start with +10%) to `GasLimit` to account for state changes between simulation and execution. A simplified "out of gas" error error is the most common support ticket for dApps.

---

### 487. How do you parse blockchain data using Go?
"I connect to an RPC node (Geth).
`client.BlockByNumber(ctx, nil)`.
I iterate over `block.Transactions()`.
To decode logs (events), I use the generated ABI bindings: `contract.ParseMyEvent(log)`.
For high-performance indexers, I sometimes decode the raw RLP bytes manually to skip the overhead of the standard RPC structs."

#### Indepth
Reorgs happen. Your indexer *must* handle "Chain Reorganizations" (where block 100 changes hash). Keep a pointer to the "Last Verified Block". If the new block 100 doesn't match your DB, roll back your DB to block 99. Ethereum finality is probabilistic (unless using PoS checkpoints).

---

### 488. How do you generate and verify ECDSA signatures in Go?
"I use `crypto/ecdsa` and `crypto/elliptic`.
**Sign**: `ecdsa.Sign(rand.Reader, priv, hash)`.
**Verify**: `ecdsa.Verify(pub, hash, r, s)`.
For Blockchain (Bitcoin/Ethereum), I specifically use the `secp256k1` curve (available in `github.com/ethereum/go-ethereum/crypto`), not the standard P256."

#### Indepth
Performance matters. `crypto/ecdsa` in Go standard lib is constant-time (safe against timing attacks) but slower than C-bindings (libsecp256k1). For a high-frequency trading bot, use the C-binding wrapper (`github.com/ethereum/go-ethereum/crypto/secp256k1`) to sign transactions 10x faster.

---

### 489. What is the role of Go in decentralized storage (IPFS)?
"**IPFS** (InterPlanetary File System) is written in Go (`kubo`).
It uses a DHT (Kademlia) for routing and Bitswap for exchange.
Go’s concurrency model is perfect for maintaining thousands of peer connections in the swarm. I verify to add files programmatically using the `go-ipfs-api`."

#### Indepth
IPFS content addressing (`QmHash`) makes data immutable. If you edit a file, the hash changes. To support mutable data (like a "User Profile"), use **IPNS** (InterPlanetary Name System) which points a static PeerID to a dynamic IPFS hash, acting like a DNS record for the decentralized web.

---

### 490. How would you implement a Merkle Tree in Go?
"A Merkle Tree is a hash tree.
I hash data blocks: `H1 = sha256(Block1)`.
I pair them: `H12 = sha256(H1 + H2)`.
I repeat until I get the **Root Hash**.
In Go, I define a `Node` struct. If I change one byte in a block, the Root Hash changes completely. I use this to verify data integrity efficiently."

#### Indepth
Bitcoin uses Double-SHA256. Ethereum uses Keccak-256. When implementing Merkle Trees, beware of **Preimage Attacks**. Always prefix leaf nodes with `0x00` and internal nodes with `0x01` before hashing to ensure a leaf can never be interpreted as an internal node.

---

### 491. How do you handle base58 and hex encoding/decoding?
"**Hex**: `hex.EncodeToString(bytes)`. Standard for debugging.
**Base58**: Used in Bitcoin/IPFS to avoid ambiguous chars (0 vs O).
Go doesn't have it in stdlib. I use `github.com/btcsuite/btcutil/base58`.
`base58.Encode(input)`. It’s essentially a base conversion algorithm."

#### Indepth
Base58 checksums prevent typos. Bitcoin addresses include a 4-byte checksum (double hash) at the end. When you decode, verify the checksum. If you send crypto to a typo-address without checksum validation, the money is burned forever. Go's strong typing (`Type Address string`) helps, but validation is key.

---

### 492. How do you write a deterministic VM interpreter in Go?
"Determinism allows zero randomness.
I implement a loop: `Fetch -> Decode -> Execute`.
I strictly **avoid** map iteration (random order) and `time.Now()`.
If I use maps, I gather keys, sort them, then iterate.
This ensures that every node running my code reaches the exact same state for the same input."

#### Indepth
Floating Point math is **non-deterministic** across different CPUs/architectures (IEEE 754 handling varies). Never use `float32/64` in a blockchain VM. Use fixed-point arithmetic libraries (`big.Int` or custom decimal types) to ensure `1.0 + 2.0` is exactly `3.0` on every node on Earth.

---

### 493. How do you simulate a P2P network in Go?
"I use **libp2p** (written in Go).
`host, _ := libp2p.New()`.
It handles NAT traversal, transport upgrades (QUIC), and protocol negotiation.
I create a simulation by starting 100 in-process hosts (goroutines) and connecting them in a mesh. I can flood messages and measure propagation delay without leaving `localhost`."

#### Indepth
GossipSub (part of libp2p) is the standard for message propagation. It uses a "Mesh" for stability (high reliability) and "Gossip" for metadata (low bandwidth). Tuning the heartbeat and mesh degree (`D`, `D_low`, `D_high`) is the difference between a 1s block time and a stalled network.

---

### 494. How do you create a lightweight Go runtime for edge computing?
"I use **Wazero**.
It’s a zero-dependency WASM runtime written in pure Go.
`r := wazero.NewRuntime(ctx)`.
I compile untrusted user code to WASM and run it inside Wazero.
It gives me a secure sandbox with millisecond startup times, perfect for 'Functions as a Service'."

#### Indepth
Security Isolation: Wazero is safer than Cgo because it accesses memory safely. It doesn't use `unsafe` pointers to cross the boundary. This means a malicious WASM module cannot SEGFAULT the host Go process, making it ideal for running user-submitted plugins.

---

### 495. How would you handle offline-first apps in Go?
"I use a local embedded DB (**SQLite** via `modernc.org/sqlite` or **BadgerDB**).
Reads come from local DB.
Writes go to a 'Sync Queue' table.
When network is available, a background goroutine drains the queue, POSTs to the API, and marks items as synced. I use 'Last-Write-Wins' for conflict resolution."

#### Indepth
CRDTs (Conflict-free Replicated Data Types) are the robust answer to "Last-Write-Wins". Use a Go library like `delta-crdt`. Instead of "State", you sync "Operations". This allows two users to edit the same Todo list offline and merge perfectly without data loss when they come online.

---

### 496. What is the future of `Generics` in Go (beyond v1.22)?
"We have Type Parameters.
The community wants **Iterators** (standard `Yield` pattern) to make generic `Map/Filter/Reduce` ergonomic.
We also want **Generic Methods on Structs** (currently methods can't have their *own* type params). This would allow truly expressive fluent APIs like LINQ."

#### Indepth
The "Monomorphization" (Stenciling) strategy of Go generics means `List[int]` and `List[string]` are compiled to two different code paths. This is faster (no boxing) but increases binary size. Be mindful of instantiating huge generic structs with many different types in embedded environments.

---

### 497. What is fuzz testing and how do you use it in Go?
"Go 1.18 added native fuzzing.
`func FuzzParse(f *testing.F)`.
I seed it with `f.Add("valid_input")`.
Then `f.Fuzz(func(t *testing.T, input string) { ... })`.
The runtime generates millions of random mutations. I verify my code doesn't panic or hang. It routinely finds edge cases (invalid UTF-8, huge integers) I missed."

#### Indepth
Fuzzing shines at **Parsers** (JSON, YAML, Protocol Buffers). If your app accepts a file upload, Fuzz it. `f.Fuzz(func(t *testing.T, data []byte) { Parse(data) })`. It will find the one byte sequence that causes your parser to index out of range or allocate 10GB of RAM.

---

### 498. What is the `any` type in Go and how is it different from `interface{}`?
"It is **not** different.
`type any = interface{}`.
It’s an alias introduced in Go 1.18.
It makes code readable: `func Print(v any)` vs `func Print(v interface{})`.
However, `any` is still a static type—it means 'box that can hold anything', not dynamic typing."

#### Indepth
Go interfaces are implemented as `{type, value}` pointers. `any` is just an empty interface. Assigning a concrete type to `any` involves an allocation (boxing) if the value is not a pointer. Frequent interfaces conversions in hot loops will generate garbage. Use concrete types where possible.

---

### 499. What is the latest experimental feature in Go and why is it important?
"**Range over func** (Iterators) in Go 1.23.
`for v := range mySeq`.
It standardizes custom collection iteration.
Any function matching `func(yield func(T) bool)` works in a `for-range` loop. This unifies how we iterate over SQL rows, HTTP streams, and slices."

#### Indepth
This is a paradigm shift. `iter.Seq[V]` allows "Push" iteration. It simplifies resource cleanup. The iterator function can `defer file.Close()`, and it keeps the file open as long as the loop runs, closing it automatically when the loop breaks. No more leaking resources in complex loops.

---

### 500. How do you contribute to the Go open-source project?
"I use **Gerrit** (`go-review.googlesource.com`).
GitHub is just a mirror.
1.  Sign the CLA.
2.  Install `git-codereview`.
3.  Discuss on the Issue Tracker.
4.  `git change` -> `git mail`.
The review process is rigorous but fair. I start with doc fixes or small bug reports."

#### Indepth
Running `all.bash` (the full Go test suite) takes time. For your first contribution, target `golang/tools` or `golang/website` repositories. They are smaller and have faster review cycles. Read `CONTRIBUTING.md` twice—Go maintainers are strict about commit message formats (`package: description`).
