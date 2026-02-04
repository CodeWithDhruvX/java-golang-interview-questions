## 🧠 AI, ML & Generative Use Cases in Go (Questions 901-920)

### Question 901: How do you call an OpenAI API using Go?

**Answer:**
(See Q469). HTTP Request (JSON body) to `https://api.openai.com/v1/chat/completions`.
Header: `Authorization: Bearer $KEY`.

---

### Question 902: How do you stream ChatGPT responses in Go?

**Answer:**
Set `"stream": true` in payload.
The response comes as Server-Sent Events (SSE).
Read response body line-by-line using `bufio.Scanner`. Parse `data: {...}` lines.

---

### Question 903: How do you build a Telegram AI bot in Go?

**Answer:**
1.  **Telegram Bot API:** `github.com/go-telegram-bot-api/telegram-bot-api`.
2.  **Long Polling:** Fetch updates.
3.  **On Message:** Send text to OpenAI API -> Get Reply -> Send to Telegram.

---

### Question 904: How do you integrate Go with HuggingFace models?

**Answer:**
Use the **Inference API** via HTTP.
Or run local ONNX models exported from HuggingFace (Python) using `onnx-go`.

---

### Question 905: How do you use TensorFlow models in Go?

**Answer:**
(See Q461). Load `SavedModel` folder. Feed Tensors.

---

### Question 906: How do you build a Go app that uses image recognition?

**Answer:**
Use `gocv` (OpenCV bindings).
Or call AWS Rekognition / Google Vision API (Cloud SDK).

---

### Question 907: How do you generate code snippets using LLMs in Go?

**Answer:**
Prompt Engineering.
Payload: `{"model": "gpt-4", "messages": [{"role": "system", "content": "You are a standard Go coding assistant."}]}`.

---

### Question 908: How do you do prompt templating in Go?

**Answer:**
`text/template`.
`const prompt = "Summarize this: {{.Text}}"`
Sanitize input to prevent injection attacks (though typical Prompt Injection is handled by the model logic, not string sanitization).

---

### Question 909: How do you build a LangChain-style pipeline in Go?

**Answer:**
**LangChainGo**.
Define a `Chain` struct.
`chain.Call(ctx, input)` -> Formats Prompt -> Calls LLM -> Parses Output.

---

### Question 910: How do you fine-tune prompts using Go templates?

**Answer:**
Store templates in DB/Config.
At runtime, load template, `Execute` with dynamic variables (User Name, History), send to API.

---

### Question 911: How do you handle concurrent API calls to LLMs?

**Answer:**
Use `errgroup` or buffered channels.
Be aware of **Rate Limits** (TPM - Tokens Per Minute).
Implement a token bucket Limiter client-side.

---

### Question 912: How do you track token usage in LLM APIs from Go?

**Answer:**
Response contains `usage` field: `{"prompt_tokens": 10, "completion_tokens": 20}`.
Log this to Database for billing.
Or use `github.com/pkoukk/tiktoken-go` to count tokens *before* sending.

---

### Question 913: How do you stream generation results to a web frontend in Go?

**Answer:**
(See Q687/Q902).
Read SSE from OpenAI -> Write SSE to Browser.
Proxy the stream to hide API Key.

---

### Question 914: How do you handle OpenAI rate limits in Go apps?

**Answer:**
Use Retry middleware with **Exponential Backoff**.
Check for 429 StatusCode.
Wait 1s, 2s, 4s...

---

### Question 915: How do you generate embeddings and store in Go?

**Answer:**
1.  Call `/v1/embeddings` endpoint. Get `[]float32`.
2.  Store in **pgvector** (Postgres extension) or **Weaviate**.

---

### Question 916: How do you integrate Pinecone or Weaviate with Go?

**Answer:**
Use their Go/gRPC clients.
`client.Upsert(id, vector, metadata)`.
`client.Search(vector, limit=10)`.

---

### Question 917: How do you manage vector searches using Go?

**Answer:**
Construct query vector (embedding of search term).
Send to Vector DB.
Receive list of IDs.
Fetch full documents from SQL DB using IDs.

---

### Question 918: How do you build a question-answering bot using Go?

**Answer:**
**RAG (Retrieval Augmented Generation).**
1.  User Query -> Embed.
2.  Search Vector DB -> Get Context Chunks.
3.  Fill Prompt: "Answer query based on context: {chunks}".
4.  LLM Generate.

---

### Question 919: How do you evaluate AI responses using Go logic?

**Answer:**
**Heuristics:** Check length, regex match.
**Model-Based:** Send output to another (cheaper) model: "Is this answer safe? Yes/No". Parse result.

---

### Question 920: How do you serialize LLM chat history in Go?

**Answer:**
`[]ChatMessage`. JSON marshal.
Store in Redis/SQL.
Prune old messages to stay within Context Window limit (e.g., keep last 10 turns).

---
