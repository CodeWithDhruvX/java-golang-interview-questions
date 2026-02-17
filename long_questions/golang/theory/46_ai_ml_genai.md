# 🧠 Go Theory Questions: 901–920 AI, ML & Generative Use Cases

## 901. How do you generate code snippets using LLMs in Go?

**Answer:**
We construct a prompt:
`prompt := "Write a Go function to reverse a string."`
Send to OpenAI/Anthropic API (`Completion` endpoint).
Parse `response.Choices[0].Text`.
For structured code, we can prompt with: "Return ONLY the code block, no markdown".
We can then write this string to a `.go` file and run `go fmt` on it programmatically.

---

## 902. How do you do prompt templating in Go?

**Answer:**
`text/template`.
`const promptTpl = "Translating {{.Text}} to {{.Language}}."`
`t := template.Must(template.New("p").Parse(promptTpl))`
`t.Execute(&buf, data)`.
This allows dynamic injection of user input into system prompts while preventing basic injection attacks (if using specialized escaping logic, though rarely needed for LLM prompts).

---

## 903. How do you build a LangChain-style pipeline in Go?

**Answer:**
Chain pattern.
`type Chain interface { Run(ctx, input) output }`.
Structs: `PromptTemplate`, `LLMChain`, `VectorDBChain`.
We compose them:
`chain := NewSequentialChain(PromptTpl, OpenAIModel, OutputParser)`.
Libraries like `tmc/langchaingo` provide these primitives out of the box.

---

## 904. How do you fine-tune prompts using Go templates?

**Answer:**
We use **Few-Shot Prompting**.
Template:
```text
Classify the sentiment:
"I love usage": Positive
"I hate it": Negative
"{{.Input}}":
```
We iterate on this template in code. We can load different template files (`v1.tpl`, `v2.tpl`) at runtime to A/B test which prompt yields better JSON output from the model.

---

## 905. How do you handle concurrent API calls to LLMs?

**Answer:**
(See Q 346).
Launch goroutines.
`var wg sync.WaitGroup`.
`for _, prompt := range prompts { wg.Add(1); go func() { callLLM(prompt) }() }`.
**Critical**: Rate Limiting.
LLM APIs have strict TPM (Tokens Per Minute).
We must use a **Rate Limiter** (`golang.org/x/time/rate`) to throttle requests: `limiter.Wait(ctx)` before calling API.

---

## 906. How do you track token usage in LLM APIs from Go?

**Answer:**
The API response includes usage metadata.
`resp.Usage.TotalTokens`.
We log this metric to Prometheus.
`tokenCounter.Add(resp.Usage.TotalTokens)`.
To estimate *before* sending (for budgeting), we use a tokenizer library in Go (like `tiktoken-go`) to count BPE tokens in the prompt string.

---

## 907. How do you stream generation results to a web frontend in Go?

**Answer:**
(See Q 896).
1.  Frontend: `EventSource("/stream")`.
2.  Backend: `stream := openai.CreateCompletionStream(...)`.
3.  Loop: `recv()`.
4.  Write: `fmt.Fprintf(w, "data: %s\n\n", text)`.
5.  Flush: `w.(http.Flusher).Flush()`.
This gives the "typing effect" to the user.

---

## 908. How do you handle OpenAI rate limits in Go apps?

**Answer:**
**Exponential Backoff**.
API returns `429 Too Many Requests`.
We catch this error.
Sleep `base * 2^attempt`.
Libraries like `go-retryablehttp` handle this automatically.
For TPM limits, we implement a client-side Token Bucket that replenishes at the rate allowed by our tier.

---

## 909. How do you generate embeddings and store in Go?

**Answer:**
1.  Call Embedding API: `client.CreateEmbeddings(text)`. Returns `[]float32`.
2.  Store: If small, in-memory slice.
3.  If large, use **pgvector** (Postgres extension).
    `db.Exec("INSERT INTO items (vec) VALUES (?)", pgvector.NewVector(vec))`.

---

## 910. How do you integrate Pinecone or Weaviate with Go?

**Answer:**
They are just REST/gRPC APIs.
Weaviate: `client.Schema().ClassCreator().WithClass(class).Do(ctx)`.
Insert: `client.Data().Creator().WithClassName("Document").WithProperties(props).WithVector(vec).Do(ctx)`.
Search: `client.GraphQL().Get().WithNearVector(vec).Do(ctx)`.
We map our Go structs to their schema properties.

---

## 911. How do you manage vector searches using Go?

**Answer:**
Query: "Find similar documents".
1.  Embed query string -> `qVec`.
2.  Send `qVec` to Vector DB.
3.  Receive `[]ID` and `Scores`.
4.  Fetch full docs from SQL/Redis using `ID`s (Hybrid Search).
    Go acts as the orchestrator between the Embedding Provider, Vector DB, and Primary DB.

---

## 912. How do you build a question-answering bot using Go?

**Answer:**
**RAG (Retrieval Augmented Generation)**.
1.  User asks Q.
2.  Go searches Vector DB for relevant context chunks.
3.  Go constructs Prompt: "Context: [chunks]... Question: [Q]... Answer:".
4.  Go sends to LLM.
5.  Returns answer.

---

## 913. How do you evaluate AI responses using Go logic?

**Answer:**
**Deterministic Guardrails**.
LLM output: `{"action": "delete"}`.
Go validation:
`if action == "delete" && user.Role != "admin" { return Error("Unsafe action") }`.
If output is code, we can try `go build` on it. If it fails, feeding the error back to the LLM to "Fix the code" is a common pattern (Self-Healing).

---

## 914. How do you serialize LLM chat history in Go?

**Answer:**
History is a slice `[]Message`.
Structure: `[{Role: "user", Content: "Hi"}, {Role: "assistant", Content: "Hello"}]`.
We store this in Redis or SQL as a JSON blob.
On new request:
1.  Fetch History.
2.  Append User Msg.
3.  Truncate old messages (if Token Limit exceeded).
4.  Send to API.

---

## 915. How do you use `database/sql` in Go?

**Answer:**
Standard library abstraction.
1.  Open: `db, _ := sql.Open("driver", "dsn")`.
2.  Query: `rows, _ := db.Query("SELECT ...")`.
3.  Scan: `rows.Scan(&dest)`.
It manages the connection pool automatically. It is driver-agnostic (works for Postgres, MySQL, SQLite).

---

## 916. What are connection pools and how to manage them?

**Answer:**
`sql.DB` *IS* a pool.
Config:
- `SetMaxOpenConns(100)`: Max active queries.
- `SetMaxIdleConns(10)`: Keep warm.
- `SetConnMaxLifetime(1 * time.Hour)`: Restart connection periodically (avoids Firewall timeouts).
If not configured, MaxOpen is infinite, which can crash the DB server under load.

---

## 917. How do you write raw queries using `sqlx`?

**Answer:**
`sqlx` extends `database/sql` with Struct Mapping.
`db.Get(&user, "SELECT * FROM users WHERE id=$1", id)`. (Single row).
`db.Select(&users, "SELECT * FROM users")`. (Multiple rows).
It removes the boilerplate of `rows.Scan(&u.ID, &u.Name, ...)`.

---

## 918. How do you use GORM with PostgreSQL?

**Answer:**
ORM (Object Relational Mapper).
`db.AutoMigrate(&User{})`.
`db.Create(&User{Name: "John"})`.
`db.Where("name = ?", "John").First(&user)`.
It handles hooks (BeforeSave), soft deletes, and associations (HasMany) automatically.

---

## 919. How do you handle transactions in Go?

**Answer:**
`tx, err := db.Begin()`.
Defer Rollback: `defer tx.Rollback()`.
Execute queries on `tx` (Not `db`).
`tx.Exec(...)`.
At end: `err = tx.Commit()`.
If `Commit()` is not called, the deferred `Rollback()` ensures no partial data is saved.

---

## 920. How do you create database migrations in Go?

**Answer:**
Tools: **golang-migrate** or **Goose**.
Files: `001_init.up.sql`, `001_init.down.sql`.
Go code embeds these files.
On startup: `migrate.Up()`.
It tracks the "schema_version" table to know which files to run.
This ensures the DB schema matches the application code capabilities.
