## 🧠 AI, ML & Generative Use Cases in Go (Questions 901-920)

### Question 901: How do you call an OpenAI API using Go?

**Answer:**
(See Q469). HTTP Request (JSON body) to `https://api.openai.com/v1/chat/completions`.
Header: `Authorization: Bearer $KEY`.

### Explanation
OpenAI API calls from Go use HTTP requests with JSON bodies to the chat completions endpoint. The request includes an Authorization header with Bearer token authentication and a JSON payload containing messages and model parameters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you call an OpenAI API using Go?
**Your Response:** "I call the OpenAI API in Go by making HTTP requests to the chat completions endpoint. I create a JSON payload with the model, messages, and parameters, then send it using Go's `net/http` package. The authorization is done with a Bearer token in the header. I typically create structs for the request and response to handle JSON marshaling cleanly. I also handle errors properly and check the HTTP status codes. For production use, I'd add retry logic, timeout handling, and proper error wrapping. The key is understanding the API structure and handling HTTP requests correctly. I might use a client library for convenience, but understanding the raw HTTP approach is important for debugging and customization."

---

### Question 902: How do you stream ChatGPT responses in Go?

**Answer:**
Set `"stream": true` in payload.
The response comes as Server-Sent Events (SSE).
Read response body line-by-line using `bufio.Scanner`. Parse `data: {...}` lines.

### Explanation
Streaming ChatGPT responses in Go uses stream:true in the payload, receiving Server-Sent Events. The response is read line-by-line with bufio.Scanner, parsing data:{...} lines to get incremental response chunks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stream ChatGPT responses in Go?
**Your Response:** "I stream ChatGPT responses by setting `stream: true` in the API payload. Instead of getting one complete response, I receive Server-Sent Events - a stream of data chunks. I read the response body line-by-line using `bufio.Scanner` and parse each line that starts with `data:` to extract the partial responses. Each chunk contains a piece of the generated text that I can display to the user in real-time. This creates a better user experience by showing the response as it's being generated. I handle connection errors and parse the JSON in each chunk. The key is understanding SSE format and processing the stream incrementally rather than waiting for the complete response."

---

### Question 903: How do you build a Telegram AI bot in Go?

**Answer:**
1.  **Telegram Bot API:** `github.com/go-telegram-bot-api/telegram-bot-api`.
2.  **Long Polling:** Fetch updates.
3.  **On Message:** Send text to OpenAI API -> Get Reply -> Send to Telegram.

### Explanation
Telegram AI bot building in Go uses the go-telegram-bot-api library, long polling for message updates, and a pipeline where messages are sent to OpenAI API, responses are received, and replies are sent back to Telegram.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a Telegram AI bot in Go?
**Your Response:** "I build Telegram AI bots using the `go-telegram-bot-api` library. First, I set up the bot with Telegram's Bot API and get a token. Then I use long polling to continuously fetch updates from Telegram. When a message comes in, I extract the text, send it to the OpenAI API, get the AI response, and send it back to the user on Telegram. I handle different message types, commands, and maintain conversation context. The bot runs in a loop, processing messages as they arrive. I also implement error handling, rate limiting, and maybe some basic commands like /help or /reset. The key is understanding the Bot API structure and managing the message flow between Telegram and the AI service."

---

### Question 904: How do you integrate Go with HuggingFace models?

**Answer:**
Use the **Inference API** via HTTP.
Or run local ONNX models exported from HuggingFace (Python) using `onnx-go`.

### Explanation
HuggingFace model integration in Go uses the Inference API via HTTP requests, or runs local ONNX models exported from HuggingFace using the onnx-go library for inference without external dependencies.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate Go with HuggingFace models?
**Your Response:** "I integrate with HuggingFace models in two ways. First, I can use their Inference API via HTTP requests - this is simpler and works like any other REST API. Second, for better performance and offline capability, I can run local ONNX models exported from HuggingFace using the `onnx-go` library. This lets me do inference directly in Go without external API calls. I choose based on the use case - API for simplicity and flexibility, local ONNX for performance and privacy. For the local approach, I export models from Python to ONNX format, then load and run them in Go. The key is understanding that Go can work with ML models either through APIs or by running optimized model formats locally."

---

### Question 905: How do you use TensorFlow models in Go?

**Answer:**
(See Q461). Load `SavedModel` folder. Feed Tensors.

### Explanation
TensorFlow model usage in Go loads SavedModel folders and feeds tensors for inference. Go's TensorFlow bindings allow running pre-trained models by loading the model directory and providing input tensors.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use TensorFlow models in Go?
**Your Response:** "I use TensorFlow models in Go by loading the SavedModel folder and feeding tensors for inference. The Go TensorFlow bindings let me load pre-trained models and run predictions. I create a session, load the model from the SavedModel directory, prepare input tensors with my data, run the inference operation, and extract the output tensors. This approach works for computer vision, NLP, and other ML tasks. The key is understanding the TensorFlow model structure and tensor operations. I handle memory management carefully and ensure proper cleanup. While Python has better ML tooling, Go is excellent for deploying models in production with good performance and concurrency."

---

### Question 906: How do you build a Go app that uses image recognition?

**Answer:**
Use `gocv` (OpenCV bindings).
Or call AWS Rekognition / Google Vision API (Cloud SDK).

### Explanation
Image recognition in Go uses gocv (OpenCV bindings) for local processing, or cloud services like AWS Rekognition or Google Vision API via their SDKs for managed AI services.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a Go app that uses image recognition?
**Your Response:** "I build image recognition apps in Go using either local processing with `gocv` (OpenCV bindings) or cloud services. With `gocv`, I can run computer vision algorithms locally - load images, detect faces, recognize objects, or run custom models. For cloud-based solutions, I use AWS Rekognition or Google Vision API through their Go SDKs. The choice depends on requirements - local processing gives more control and no latency, while cloud services offer pre-trained models and scalability. I handle image preprocessing, model loading, and result parsing. The key is understanding the image processing pipeline and choosing the right approach based on performance, cost, and accuracy requirements."

---

### Question 907: How do you generate code snippets using LLMs in Go?

**Answer:**
Prompt Engineering.
Payload: `{"model": "gpt-4", "messages": [{"role": "system", "content": "You are a standard Go coding assistant."}]}`.

### Explanation
Code generation using LLMs in Go relies on prompt engineering with carefully crafted payloads. The request includes model specification, system messages defining the assistant role, and user messages with coding requirements.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you generate code snippets using LLMs in Go?
**Your Response:** "I generate code snippets using LLMs by crafting effective prompts. I create a JSON payload with the model specification and a system message that defines the role - like 'You are a standard Go coding assistant.' Then I provide a user message with specific coding requirements. The key is prompt engineering - being clear about what I want, providing context, and specifying the output format. I might include examples, constraints, or specific Go patterns I want. The response comes back as generated code that I can parse and use. I handle errors, validate the generated code, and might even compile it to check for syntax errors. The art is in writing prompts that produce high-quality, usable code."

---

### Question 908: How do you do prompt templating in Go?

**Answer:**
`text/template`.
`const prompt = "Summarize this: {{.Text}}"`
Sanitize input to prevent injection attacks (though typical Prompt Injection is handled by the model logic, not string sanitization).

### Explanation
Prompt templating in Go uses text/template for dynamic prompt generation. Templates like "Summarize this: {{.Text}}" allow variable substitution. Input sanitization helps prevent injection, though prompt injection is typically handled by model logic rather than string sanitization.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you do prompt templating in Go?
**Your Response:** "I handle prompt templating in Go using the standard `text/template` package. I create templates with placeholders like `{{.Text}}` that get filled with dynamic data at runtime. For example, a template might be 'Summarize this: {{.Text}}' where Text is filled with user input. I execute the template with the data and send the result to the LLM. While I sanitize inputs to prevent injection attacks, most prompt injection issues are handled by the model logic rather than string sanitization. The key is creating reusable prompt templates that can be dynamically filled with different data. This approach makes prompts consistent and maintainable while allowing for personalization and dynamic content."

---

### Question 909: How do you build a LangChain-style pipeline in Go?

**Answer:**
**LangChainGo**.
Define a `Chain` struct.
`chain.Call(ctx, input)` -> Formats Prompt -> Calls LLM -> Parses Output.

### Explanation
LangChain-style pipelines in Go use LangChainGo library with Chain structs. The chain.Call method orchestrates the pipeline: formatting prompts, calling LLMs, and parsing outputs in a structured workflow.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a LangChain-style pipeline in Go?
**Your Response:** "I build LangChain-style pipelines using the LangChainGo library. I define a `Chain` struct that orchestrates the entire workflow. When I call `chain.Call(ctx, input)`, it automatically formats the prompt with the input, calls the LLM API, and parses the output according to the chain configuration. This creates a clean, reusable pipeline for AI workflows. I can chain multiple steps together - like retrieving documents, formatting context, calling the LLM, and then processing the response. The pattern makes complex AI applications more maintainable by breaking them into discrete, composable steps. The key is understanding the chain concept and how each step transforms the data for the next step."

---

### Question 910: How do you fine-tune prompts using Go templates?

**Answer:**
Store templates in DB/Config.
At runtime, load template, `Execute` with dynamic variables (User Name, History), send to API.

### Explanation
Fine-tuning prompts with Go templates stores templates in database or configuration files. At runtime, templates are loaded and executed with dynamic variables like user names and conversation history before sending to the API.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you fine-tune prompts using Go templates?
**Your Response:** "I fine-tune prompts by storing template variations in a database or configuration files. At runtime, I load the appropriate template and execute it with dynamic variables like the user's name, conversation history, or specific context. This allows me to personalize prompts without hardcoding them. For example, I might have different templates for different user types or conversation stages. The `template.Execute` method fills in the variables, creating a customized prompt that I send to the API. This approach makes prompts maintainable and allows for A/B testing different prompt variations. The key is separating the prompt structure from the dynamic data, which makes the system more flexible and easier to update."

---

### Question 911: How do you handle concurrent API calls to LLMs?

**Answer:**
Use `errgroup` or buffered channels.
Be aware of **Rate Limits** (TPM - Tokens Per Minute).
Implement a token bucket Limiter client-side.

### Explanation
Concurrent LLM API calls use errgroup or buffered channels for goroutine management. Important to respect rate limits (tokens per minute) and implement client-side token bucket limiters to avoid API throttling.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle concurrent API calls to LLMs?
**Your Response:** "I handle concurrent LLM API calls using `errgroup` or buffered channels to manage goroutines. The key challenge is respecting rate limits - most LLM APIs have tokens per minute limits. I implement a client-side token bucket limiter to control the request rate and avoid getting throttled. I might use a semaphore pattern to limit concurrent requests or queue requests when approaching limits. I also handle errors gracefully and implement retry logic for failed requests. The goal is to maximize throughput without hitting API limits. I monitor usage and adjust the concurrency based on the rate limits. This approach ensures I can process multiple requests efficiently while staying within the API constraints."

---

### Question 912: How do you track token usage in LLM APIs from Go?

**Answer:**
Response contains `usage` field: `{"prompt_tokens": 10, "completion_tokens": 20}`.
Log this to Database for billing.
Or use `github.com/pkoukk/tiktoken-go` to count tokens *before* sending.

### Explanation
Token usage tracking in LLM APIs uses the usage field in responses showing prompt and completion tokens. This data is logged to databases for billing, or token counting libraries like tiktoken-go are used to count tokens before sending requests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you track token usage in LLM APIs from Go?
**Your Response:** "I track token usage in two ways. First, the API response includes a `usage` field that breaks down prompt and completion tokens - I log this to a database for billing and analytics. Second, I can count tokens before sending requests using libraries like `tiktoken-go` which implements the same tokenization as the API. This helps me estimate costs and stay within limits. I track usage per user, per request type, and over time to understand patterns. This data is crucial for cost management, rate limiting, and optimizing prompts. The key is understanding both actual usage from responses and estimated usage before requests."

---

### Question 913: How do you stream generation results to a web frontend in Go?

**Answer:**
(See Q687/Q902).
Read SSE from OpenAI -> Write SSE to Browser.
Proxy the stream to hide API Key.

### Explanation
Streaming generation to web frontends reads Server-Sent Events from OpenAI and writes them as SSE to the browser. A proxy pattern hides API keys while streaming real-time responses to the frontend.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stream generation results to a web frontend in Go?
**Your Response:** "I stream generation results to web frontends by acting as a proxy. I read Server-Sent Events from the OpenAI API and immediately write them as SSE to the browser. This creates a real-time streaming experience for users. The proxy approach hides my API keys and allows me to add authentication or rate limiting on my server. I handle the SSE protocol on both ends - parsing incoming chunks and formatting outgoing ones. The key is maintaining the stream connection and handling any errors gracefully. This approach gives users the instant feedback of streaming while keeping my API credentials secure. I can also add logging or modify the stream if needed."

---

### Question 914: How do you handle OpenAI rate limits in Go apps?

**Answer:**
Use Retry middleware with **Exponential Backoff**.
Check for 429 StatusCode.
Wait 1s, 2s, 4s...

### Explanation
OpenAI rate limit handling uses retry middleware with exponential backoff. When receiving 429 status codes, the system waits progressively longer (1s, 2s, 4s...) before retrying requests.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle OpenAI rate limits in Go apps?
**Your Response:** "I handle OpenAI rate limits using retry middleware with exponential backoff. When I get a 429 status code indicating I've hit the rate limit, I don't fail immediately. Instead, I wait progressively longer - 1 second, then 2 seconds, then 4 seconds - before retrying. This exponential backoff helps avoid overwhelming the API and gives it time to reset. I also implement jitter to prevent thundering herd problems. I track the rate limit headers from responses to adjust my request rate dynamically. The key is being resilient to rate limits while maintaining good performance for users. I also provide feedback when requests are delayed so users know what's happening."

---

### Question 915: How do you generate embeddings and store in Go?

**Answer:**
1.  Call `/v1/embeddings` endpoint. Get `[]float32`.
2.  Store in **pgvector** (Postgres extension) or **Weaviate**.

### Explanation
Embedding generation and storage calls the /v1/embeddings endpoint to get float32 vectors, then stores them in vector databases like pgvector (Postgres extension) or Weaviate for similarity search.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you generate embeddings and store in Go?
**Your Response:** "I generate embeddings by calling the `/v1/embeddings` endpoint with my text, which returns a `[]float32` vector representing the semantic meaning. Then I store these vectors in a specialized vector database - either `pgvector` as a Postgres extension or a dedicated vector database like Weaviate. The vector storage allows me to do similarity searches later. I handle the API call, parse the vector response, and store it with metadata about the original text. The key is understanding that embeddings capture semantic meaning as numerical vectors, and storing them properly enables efficient similarity searches. I batch requests when possible and handle the high-dimensional data carefully."

---

### Question 916: How do you integrate Pinecone or Weaviate with Go?

**Answer:**
Use their Go/gRPC clients.
`client.Upsert(id, vector, metadata)`.
`client.Search(vector, limit=10)`.

### Explanation
Pinecone and Weaviate integration uses their Go or gRPC clients for vector operations. Upsert stores vectors with IDs and metadata, while Search performs similarity searches to find similar vectors.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you integrate Pinecone or Weaviate with Go?
**Your Response:** "I integrate vector databases like Pinecone and Weaviate using their Go or gRPC clients. I connect to the service, then use `client.Upsert()` to store vectors with IDs and metadata, and `client.Search()` to find similar vectors. The clients handle the complexity of communicating with these managed services. I prepare my vectors and metadata, make the API calls, and handle the responses. These services are designed for high-performance vector similarity search at scale. I handle connection management, error handling, and batch operations when possible. The key is understanding the vector database concepts and using the client libraries effectively. This allows me to build applications that can search through millions of vectors efficiently."

---

### Question 917: How do you manage vector searches using Go?

**Answer:**
Construct query vector (embedding of search term).
Send to Vector DB.
Receive list of IDs.
Fetch full documents from SQL DB using IDs.

### Explanation
Vector search management constructs query vectors from search terms, sends them to vector databases to get similar document IDs, then fetches full documents from SQL databases using those IDs in a hybrid approach.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you manage vector searches using Go?
**Your Response:** "I manage vector searches using a hybrid approach. First, I convert the search term into an embedding vector. Then I send this query vector to the vector database to find similar documents, which returns a list of document IDs. Finally, I fetch the full documents from my regular SQL database using those IDs. This combines the semantic search capabilities of vectors with the structured data storage of SQL. I handle the embedding generation, vector search, and document retrieval as separate steps. The key is understanding that vector databases are great for similarity matching, but I still need a regular database for the full document content and metadata. This pattern works well for semantic search applications."

---

### Question 918: How do you build a question-answering bot using Go?

**Answer:**
**RAG (Retrieval Augmented Generation).**
1.  User Query -> Embed.
2.  Search Vector DB -> Get Context Chunks.
3.  Fill Prompt: "Answer query based on context: {chunks}".
4.  LLM Generate.

### Explanation
Question-answering bots use RAG (Retrieval Augmented Generation): user queries are embedded, vector database searches retrieve context chunks, prompts are filled with context, and LLM generates answers based on retrieved information.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a question-answering bot using Go?
**Your Response:** "I build question-answering bots using RAG - Retrieval Augmented Generation. First, I take the user's query and convert it to an embedding vector. Then I search my vector database to find relevant context chunks. I create a prompt that includes the context and asks the LLM to answer based on that context. Finally, the LLM generates an answer using the retrieved information. This approach ensures the AI answers are grounded in actual data rather than just its training. I handle the entire pipeline in Go - embedding generation, vector search, prompt construction, and LLM interaction. The key is combining retrieval with generation to get accurate, context-aware answers."

---

### Question 919: How do you evaluate AI responses using Go logic?

**Answer:**
**Heuristics:** Check length, regex match.
**Model-Based:** Send output to another (cheaper) model: "Is this answer safe? Yes/No". Parse result.

### Explanation
AI response evaluation uses heuristics like length checks and regex matching, or model-based approaches where outputs are sent to cheaper models for safety validation with yes/no responses that are parsed for automated quality control.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you evaluate AI responses using Go logic?
**Your Response:** "I evaluate AI responses using two approaches. First, heuristics - I check basic things like response length, format compliance, or use regex patterns to validate specific formats. Second, model-based evaluation - I send the AI's output to another, cheaper model with a prompt like 'Is this answer safe? Yes/No' and parse the result. This automated quality control helps catch inappropriate or incorrect responses before they reach users. I implement both approaches and combine their results for a confidence score. The key is having multiple validation layers and understanding that different evaluation methods catch different types of issues."

---

### Question 920: How do you serialize LLM chat history in Go?

**Answer:**
`[]ChatMessage`. JSON marshal.
Store in Redis/SQL.
Prune old messages to stay within Context Window limit (e.g., keep last 10 turns).

### Explanation
LLM chat history serialization uses []ChatMessage structs with JSON marshaling, storing in Redis or SQL databases. Old messages are pruned to stay within context window limits, typically keeping the last 10 conversation turns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you serialize LLM chat history in Go?
**Your Response:** "I serialize chat history using a `[]ChatMessage` struct that I JSON marshal for storage. I store the conversation in Redis for fast access or SQL for persistence. The key challenge is managing the context window - I prune old messages to stay within the limit, typically keeping just the last 10 turns of conversation. I maintain the message order and roles (user, assistant, system) so the context makes sense to the LLM. When loading history, I reconstruct the conversation in the right format. I might also summarize older conversations to preserve some context while staying within limits. The key is balancing conversation continuity with context window constraints."

---
