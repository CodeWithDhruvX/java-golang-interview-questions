# 🟢 Go Theory Questions: 461–480 AI, Machine Learning & Data Processing in Go

## 461. How do you use TensorFlow or ONNX models in Go?

**Answer:**
We typically use the **C Bindings**.

Go is not a native ML language like Python. We use the `tensorflow/go` package which wraps `libtensorflow.so` (C++).
However, for production inference, we prefer **ONNX Runtime**.
We export the model from PyTorch/TensorFlow to `.onnx`.
Then we use a pure Go ONNX runner (or CGO wrapper) to load the graph and run `Session.Run(inputTensor)`. This keeps the Go binary small and fast without requiring a full Python environment.

---

## 462. What is `gorgonia` and when would you use it?

**Answer:**
`gorgonia` is "TensorFlow for Go".

It is a native graph computation library. You define a graph (`z = x * y + b`) and it can perform automatic differentiation (backpropagation).
You use it if you need to **train** a model directly in Go or build a custom neural network from scratch without CGO.
However, for standard tasks, it's widely considered less mature than the Python ecosystem, so we mostly use it for niche, high-performance edge cases.

---

## 463. How do you implement cosine similarity in Go?

**Answer:**
Cosine similarity measures how close two vectors are (Used in RAG/Search).
Formula: `dot(A, B) / (norm(A) * norm(B))`.

In Go, we write a tight loop:
```go
func key(a, b []float64) float64 {
    var dot, normA, normB float64
    for i := range a {
        dot += a[i] * b[i]
        normA += a[i] * a[i]
        normB += b[i] * b[i]
    }
    return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}
```
For huge vectors, we use SIMD-optimized assembly versions (like `gonum`) to make this 10x faster.

---

## 464. How would you stream CSV → transform → JSON using pipelines?

**Answer:**
We build a **3-Stage Pipeline** connected by channels.

1.  **Reader**: Reads CSV line-by-line using `csv.NewReader`, sends `[]string` to `chan RawData`.
2.  **Transformer**: `N` concurrent workers read `RawData`, parse types, apply business logic, and send structs to `chan Result`.
3.  **Writer**: Reads `Result`, uses `json.NewEncoder(file).Encode()`.
This allows us to process a 100GB CSV file using only 10MB of RAM, as data flows through the stream constantly.

---

## 465. How do you process large datasets using goroutines?

**Answer:**
We use the **Worker Pool** pattern.

Don't spawn a goroutine per row (overhead kills you).
Spawn `runtime.NumCPU()` workers.
Feed input data into a shared channel.
Each worker processes a batch.
Critically, if order matters, we attach an ID to every chunk and re-assemble them at the end. If order doesn't matter (like "Resize all images"), it's a perfect parallel problem.

---

## 466. How do you implement TF-IDF in Go?

**Answer:**
TF (Term Frequency) * IDF (Inverse Document Frequency).

1.  **Map Phase**: Iterate all docs. Tokenize. Count words per doc.
2.  **Global Count**: Track "In how many docs does word X appear?"
3.  **Compute**: `Score = (CountInDoc / WordsInDoc) * log(TotalDocs / DocsWithWord)`.

We use `map[string]int` for the counters. Since this is memory-intensive for large corpora, we might stream the docs twice (once to count global frequencies, once to compute scores).

---

## 467. How do you parse and tokenize text in Go?

**Answer:**
For simple needs, `strings.Fields()` splits by whitespace.
For complex NLP, we use `github.com/jdkato/prose` or `bleve/analysis`.

We need a chain:
**Tokenizer**: Split "Hello, world!" -> ["Hello", ",", "world", "!"]
**Normalizer**: Lowercase, remove accents.
**Stopword Filter**: Remove "the", "is", "at".
**Stemmer**: "running" -> "run".

Go's strict UTF-8 support makes handling multi-byte languages (like Emoji or Chinese) straightforward compared to C++.

---

## 468. How would you embed a local LLM into a Go app?

**Answer:**
We use **binding to llama.cpp**.

There are Go wrappers (like `go-llama.cpp`) that link against the C++ library.
We load a GGUF quantized model (4GB file) into RAM.
We call `model.Predict("Hello")`.
The Go app effectively becomes the inference engine. This is popular for "Private AI" tools running on user laptops where sending data to OpenAI is forbidden.

---

## 469. How do you integrate OpenAI API in Go?

**Answer:**
We use `sashabaranov/go-openai` (the community standard SDK).

`client := openai.NewClient(token)`
`resp, err := client.CreateChatCompletion(...)`

The critical part is handling **Context Windows**. We must count tokens (using `tiktoken-go`) before sending the request. If the conversation history is too long, we must truncate the oldest messages or summarize them, otherwise the API returns a 400 error.

---

## 470. How do you do prompt engineering for AI from Go?

**Answer:**
We treat prompts as **Go Templates** (`text/template`).

`const promptTmpl = "Summarize this email from {{.Sender}}: {{.Body}}"`

We define a struct `Data { Sender, Body string }`.
We execute the template to buffer.
This ensures user input is inserted into the prompt structure cleanly. We also validate the input *before* template execution to prevent "Prompt Injection" attacks (e.g., ensuring the body doesn't contain "Ignore previous instructions").

---

## 471. How do you use a local vector database with Go?

**Answer:**
For local/embedded usage, we use libraries like **Chromem-go** or simply a specialized in-memory structure.

We store vectors in a `[]float32`.
we use a KD-Tree or HNSW (Hierarchical Navigable Small World) index implementation in pure Go to speed up the "Nearest Neighbor" search.
Unlike calling a remote Pinot/Milvus, this runs in-process, offering microsecond latency for small datasets (< 1M vectors).

---

## 472. How would you implement semantic search using Go?

**Answer:**
1.  **Ingest**: User types query "Laptop".
2.  **Embed**: Call OpenAI `text-embedding-3-small` to get a `[]float32` vector.
3.  **Search**: Query our Vector DB (like Weaviate or Postgres `pgvector`) using the vector.
4.  **Rank**: Return top 5 matches (e.g., "MacBook", "Dell XPS") even if they don't contain the exact string "Laptop".

Go orchestrates these 3 calls concurrently to keep latency under 200ms.

---

## 473. How would you extract entities using regex or AI?

**Answer:**
**Regex**: Good for structured formats (Emails, Phones). `regexp.MustCompile(`[a-z]+@[a-z]+\.com`)`. Fast, deterministic.
**AI (NER)**: Good for unstructured data ("Meeting with Bob tomorrow").
We send the text to an LLM with the prompt: "Extract names and dates as JSON."
In Go, we unmarshal that JSON response into a strong struct `type Entity struct { Name, Date string }`. We prefer AI for recall, but Regex for precision/speed.

---

## 474. How do you manage model input/output formats in Go?

**Answer:**
Data Science uses Python/Pandas (DataFrames).
Go uses Structs.

We use **Apache Arrow** or **Parquet** as the interchange format.
Go has libraries (`parquet-go`) to read these binary columnar formats efficiently.
This allows Python to train the model and save to Parquet. Go reads the Parquet file, loads specific columns into structs, and runs business logic. This avoids the slowness of CSV parsing.

---

## 475. How would you create a chatbot backend with Go?

**Answer:**
We need **State Management**.

HTTP is stateless. The Bot must remember "User is in the middle of booking a flight."
We store `SessionID -> ConversationState` in Redis.
Go Handler:
1.  Get Msg.
2.  Fetch History from Redis.
3.  Append Msg.
4.  Send to LLM.
5.  Receive Answer.
6.  Append Answer.
7.  Save to Redis.
8.  Push to WebSocket/Slack.

---

## 476. How do you build a recommendation engine with Go?

**Answer:**
We generally implement **Collaborative Filtering** (Matrix Factorization) or simple Item-Item similarity.

In Go, we load the "User-Item Interaction Matrix" (sparse) into memory.
When User A visits, we find "Similar Users" (concurrency helps here to scan millions of users).
We verify items those similar users bought.
We sort and return.
Go's performance allows doing this "Online" (real-time) rather than pre-computing everything offline, enabling "Session-based Recommendations" (reacting to what I just clicked 5 seconds ago).

---

## 477. How would you integrate LangChain-like logic in Go?

**Answer:**
"LangChain" is just a chain of API calls.
In Go, we write this imperatively (which is often cleaner).

```go
summary := CallLLM("Summarize", text)
keywords := CallLLM("Extract Keywords", summary)
sql := CallLLM("Generate SQL for", keywords)
db.Exec(sql)
```
We handle the error checking explicitly between steps. We use `langchaingo` library if we want pre-built abstractions (Agents, Chains, Memory), but raw Go code is often easier to debug.

---

## 478. How would you cache AI model outputs in Go?

**Answer:**
LLM calls are expensive ($) and slow (sec).
We use **Semantic Caching**.

Key = Vector Embedding of the Query.
Val = The LLM Response.

When a query comes in, we embed it. We query Redis Vector Search for "Similar queries" (> 0.95 cosine similarity).
If we find a match ("How reset pass?" ~= "How run password recovery?"), we return the cached answer. This saves 90% of costs for repetitive questions.

---

## 479. What is the role of concurrency in AI inference in Go?

**Answer:**
Models are usually CPU/GPU bound (single threaded per inference).
Go handles the **Serving Layer**.

We can accept 10,000 HTTP requests.
We have a pool of 5 GPU Workers (Channels).
Go queues the 10,000 requests and feeds the 5 workers.
This **Batching** is critical. We might wait 10ms to accumulate 32 requests, stack them into one Tensor, and send it to the GPU at once. This increases Throughput massively compared to 1-by-1 processing.

---

## 480. How do you monitor and scale AI pipelines in Go?

**Answer:**
We track **Token Latency** (Time per token).
We use Prometheus Histograms.

Scaling Trigger: **Queue Depth**.
If the channel `chan InferenceRequest` has > 50 items pending, we spin up more generic worker pods (if CPU bound) or alert that we need more GPU nodes.
We also monitor **Drift**: If the output distribution changes significantly (e.g., model starts predicting "False" 99% of time), we trigger an alert to retrain.
