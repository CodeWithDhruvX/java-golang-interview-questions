# 🔴 **461–480: AI, Machine Learning & Data Processing**

### 461. How do you use TensorFlow or ONNX models in Go?
"I use the **Go bindings for TensorFlow** or `onnx-go`.
I typically train in Python (PyTorch/TF) and export to **ONNX**.
In Go, I load the `.onnx` model and run inference.
This gives me the best of both worlds: Python's ecosystem for training, and Go's raw speed and concurrency for the serving API."

#### Indepth
For production inference, I avoid the overhead of Python-Go CGO bridges if possible. Instead, I use **Triton Inference Server** (which runs the model) and call it from Go via gRPC. This decouples the heavy ML runtime from my lightweight Go business logic and allows independent scaling.

---

### 462. What is `gorgonia` and when would you use it?
"**Gorgonia** is 'TensorFlow for Go'.
It creates computation graphs and handles auto-differentiation in pure Go.
I use it when I need to build/train simple models *without* Cgo or Python dependencies (e.g., a standalone binary that learns on the fly).
However, for production LLMs/Vision, I prefer bindings to C++ runtimes."

#### Indepth
Gorgonia is great for learning "Computational Graphs" but has a steep curve. It uses a Lisp-like `Let` expression system. If you just need matrix math, use **Gonum**. If you need full autograd for a custom neural net, use Gorgonia. For standard models, stick to ONNX.

---

### 463. How do you implement cosine similarity in Go?
"It’s the dot product divided by the magnitudes.
`Sim(A, B) = (A . B) / (||A|| * ||B||)`.
I iterate over the slices of `float64` to compute this.
For high-performance search (vectors of 1536 dims), I use **SIMD** instructions (via `gonum` assembly) or offload it to a Vector DB."

#### Indepth
Cosine Similarity is sensitive to **Magnitude**. Always normalize your vectors (L2 norm = 1) *before* storing them. If vectors are normalized, Cosine Similarity simplifies to just the Dot Product, which is much faster to compute (no division or square roots needed during the search query).

---

### 464. How would you stream CSV → transform → JSON using pipelines?
"I use a **Pipeline** of channels.
1.  **Reader**: Reads lines, sends to `chan []string`.
2.  **Worker**: Parses and Transforms (normalizes data), sends to `chan Struct`.
3.  **Writer**: Marshals to JSON, writes to `io.Writer`.
This allows me to convert a 100GB CSV file using only 10MB of RAM."

#### Indepth
Error handling in pipelines is tricky. If the "Transform" stage fails for row 500, should the whole pipeline crash? I usually have a separate `errChan` where workers send non-fatal errors. The main loop logs them and continues, ensuring one bad row doesn't kill the bulk job.

---

### 465. How do you process large datasets using goroutines?
"I use the **Worker Pool** pattern.
I don't spawn 1 million goroutines for 1 million rows (that kills the GC).
I spawn `runtime.NumCPU()` workers.
I feed them jobs via a buffered channel.
This keeps the CPU saturated at 100% without thrashing memory."

#### Indepth
Tuning the pool size: `runtime.NumCPU()` is a good default for CPU-bound tasks (math). For IO-bound tasks (fetching URLs), you can go much higher (`100 * NumCPU`). Use a semaphore (weighted channel) to control concurrency if your "jobs" vary wildly in cost.

---

### 466. How do you implement TF-IDF in Go?
"**Term Frequency - Inverse Document Frequency**.
1.  **Map**: I tokenize docs and count words (TF) in parallel using goroutines/sharded maps.
2.  **Reduce**: I aggregate counts to compute IDF (global rarity).
3.  **Score**: Multiply TF * IDF.
Go's concurrency makes the 'Map' phase incredibly fast compared to single-threaded Python scripts."

#### Indepth
Be careful with the memory layout. Storing `map[string]map[string]int` (word -> docID -> count) overhead is massive due to pointers. Use **Integer IDs** for words (start with a dictionary mapping) to change the problem to `map[uint32]map[uint32]uint16`, which is much more cache-friendly and compact.

---

### 467. How do you parse and tokenize text in Go?
"For simple english, `bufio.Scanner` with `ScanWords`.
For NLP, I use **segmentation libraries** (like `prose`) that handle punctuation better.
If performance is critical (log parsing), I write a zero-allocation lexer that operates on `[]byte` without creating new strings."

#### Indepth
Use `GOEXPERIMENT=arenas` (Go 1.20+) for bulk text processing. You can allocate all the parse nodes for a document in a single memory arena and free them all at once when done. This eliminates the GC overhead of tracking millions of tiny abstract syntax tree nodes.

---

### 468. How would you embed a local LLM into a Go app?
"I use **llama.cpp** bindings (`go-llama.cpp`).
I load a **GGUF** quantized model (e.g., Llama-3-8B-Q4) into memory.
Inference runs locally on CPU/GPU.
This is perfect for privacy-focused apps (GDPR) where data cannot leave the server."

#### Indepth
Go + llama.cpp allows specific tricks like **Grammar Constrained Sampling**. You can force the LLM to output valid JSON by providing a grammar file. The inference engine will zero-out probabilities for any token that would break the JSON syntax, guaranteeing perfect structure every time.

---

### 469. How do you integrate OpenAI API in Go?
"I use `sashabaranov/go-openai`.
It handles the JSON boilerplate and Context.
`resp, err := client.CreateChatCompletionStream(...)`.
I **always stream** the response. Waiting 5s for a full paragraph feels broken; seeing tokens appear instantly feels magic."

#### Indepth
Handling Stream disconnects: `stream.Recv()` will return `io.EOF` when done. But if the context is canceled (user closed tab), you get `context.Canceled`. You must handle this to stop processing and save tokens. Also, the `[DONE]` message from OpenAI is a special case to handle in the loop.

---

### 470. How do you do prompt engineering for AI from Go?
"I treat prompts as **Go Templates**.
`const tpl = "Summarize this: {{.Text}}"`
I execute the template with the user's data to generate the final string.
This separates the 'Prompt Logic' from the 'Code Logic', allowing me to simplify updating prompts without recompiling if I load them from a file."

#### Indepth
Defense against **Prompt Injection**: Treat user input as untrusted. Never just `{{.Input}}`. Wrap it in delimiters like XML tags `<user_input>{{.Input}}</user_input>` and tell the model to "Only answer based on content inside the tags". Go templates helps structure this defensively.

---

### 471. How do you use a local vector database with Go?
"For small scale (<100k items), I keep vectors in memory (`[]float32`) and brute-force the dot product. Go is fast enough.
For large scale, I use **Weaviate** or **Chroma**.
I send the vector to the DB, and it returns the top-K IDs. Go just acts as the orchestrator."

#### Indepth
Latency killer: **Serialization**. Sending 1000 floats as JSON `[0.123, 0.456, ...]` is slow. Use binary protocols (gRPC/Protobuf) or raw bytes when sending vectors to the DB. Most Vector DBs support a binary interface or Arrow format for bulk insertion.

---

### 472. How would you implement semantic search using Go?
"1.  **Embed**: Send user query to OpenAI/Local model -> Get Vector.
2.  **Search**: Query Vector DB with that vector.
3.  **Retrieve**: Fetch full documents from Postgres using returned IDs.
I build this as a microservice where Go handles the concurrent fan-out to these APIs."

#### Indepth
**Hybrid Search** (RRF - Reciprocal Rank Fusion). Pure vector search misses exact keyword matches (e.g., searching for a specific product SKU). Best practice is to run a Keyword Search (Elastic) AND a Vector Search (Weaviate), then merge the results in Go using a weighted scoring algorithm.

---

### 473. How would you extract entities using regex or AI?
"**Structured (IDs, Emails)**: Regex. Fast, deterministic.
**Unstructured (Names, Intent)**: LLM.
I use a **Hybrid Chain**:
Run Regex first. If no match, call LLM: 'Extract the order ID from this text: ...'.
This balances cost and accuracy."

#### Indepth
For locally extracting PII (Emails, Phones) *before* sending data to an LLM (for privacy), use `gliderlabs/ssh`'s pattern or google's `re2`. `re2` guarantees linear time execution, preventing ReDoS (Regex Denial of Service) attacks if a user inputs a malicious string against your regex.

---

### 474. How do you manage model input/output formats in Go?
"I use strictly typed **Structs**.
`type AIResponse struct { Answer string `json:"answer"` }`.
I define the expected JSON schema in the struct tags.
When calling the LLM, I often ask it to 'Response in JSON', and then I unmarshal it directly into my Go struct. If unmarshal fails, I retry."

#### Indepth
**Function Calling** (Tool Use). Instead of begging for JSON, define a "Tools" schema in the Go OpenAI client. The model will return a structured function call argument (valid JSON by design) which you can unmarshal strictly. This is far more reliable than "Prompt Engineering" for output formatting.

---

### 475. How would you create a chatbot backend with Go?
"**WebSockets** + **Redis**.
1.  Client connects via WS.
2.  I fetch conversation history from Redis.
3.  I call LLM (streaming).
4.  I push tokens to WS as they arrive.
5.  I append the new exchange to Redis.
Go's concurrency handles thousands of active WS connections easily."

#### Indepth
Be careful with **Concurrency Limits**. If 1000 users chat at once, you can't spawn 1000 concurrent requests to OpenAI (you will hit rate limits). Use a **Job Queue** (pgueue/Asynq) or a bounded semaphore to limit active inferences, queuing the rest with a "Thinking..." status.

---

### 476. How do you build a recommendation engine with Go?
"I use Go for the **Serving Layer**.
I pre-compute the User-Item matrix or embeddings offline (Python).
I upload the results to Redis (`SET user:1:recs [5, 12, 99]`).
The Go API just fetches from Redis. It delivers sub-millisecond responses, which is critical for the homepage load time."

#### Indepth
**Bloom Filters**. To avoid showing items the user has already seen/bought, I fetch their history. But if history is huge, I use a Bloom Filter (probabilistic set) stored in Redis. I check `Assuming not seen` in O(1) before recommending. Go's `bits` package helps implement efficient bitsets for this.

---

### 477. How would you integrate LangChain-like logic in Go?
"I use **LangChainGo** or custom interfaces.
I define a `Chain` interface: `Run(ctx, input) output`.
I chain steps: `Prompt -> LLM -> Parser`.
Go's static typing makes these chains much easier to debug than Python's dictionary-passing chaos."

#### Indepth
LangChainGo is evolving fast. The core value is the **Interface abstraction**. You can swap "OpenAI" for "Claude" or "Local Llama" just by changing the struct implementation, and the rest of your chain (Memory, parsing, tools) remains identical. This prevents vendor lock-in.

---

### 478. How would you cache AI model outputs in Go?
"**Semantic Caching**.
I don't cache by exact string match.
I embed the query.
I check my vector cache for a query with **>0.99 similarity**.
If found, I return the cached answer.
This drastically cuts API costs for meaningful duplicates (e.g., 'Hello' vs 'Hi there')."

#### Indepth
For the cache key, don't use the raw string. Tokenize, sort, and normalize. Or better, use the **Embedding Vector** itself as the key (using Locality Sensitive Hashing - LSH). This creates a "Fuzzy Cache" where "How do I reset password?" and "Reset password instructions" hit the same cache entry.

---

### 479. What is the role of concurrency in AI inference in Go?
"**Batching**.
GPUs hate single requests. They love batches.
I use a Go channel to buffer incoming requests.
Every 50ms (or when I have 32 items), I send a **Batch** to the model.
Then I distribute the answers back to the waiting clients. This increases throughput by 10x-100x."

#### Indepth
Dynamic Batching requires a strict timeout. `select { case req := <-ch: batch = append(batch, req); case <-time.After(10 * time.Millisecond): send(batch) }`. You must balance latency (waiting for batch to fill) vs throughput. 10ms is usually the sweet spot for real-time apps.

---

### 480. How do you monitor and scale AI pipelines in Go?
"I track **Token Usage** and **Latency** (Time to First Token).
I use **KEDA** (Kubernetes Event-Driven Autoscaling).
I scale my pods based on the **Queue Depth** (pending prompts).
If the queue fills up, KEDA adds more GPU nodes. Go's role is to efficiently feed that queue."

#### Indepth
Cost Monitoring! AI is expensive. I log `metadata["total_tokens"]` from every response. I wrap the `openai.Client` to count tokens per Tenant/User. If a user burns $50 in 1 hour, I trigger a Circuit Breaker to block them. Go's atomic counters make this tracking virtually free.
