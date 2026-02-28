# Spring AI, LLMs, and RAG - Interview Questions and Answers

## 1. What is Spring AI, and what are its primary goals?
**Answer:**
**Spring AI** is an application framework that simplifies the development of AI-powered applications in the Java ecosystem.

**Primary Goals:**
1. **Portability / Abstraction:** Just as Spring Data abstracts away the differences between Hibernate and EclipseLink, and Spring Cloud abstracts away AWS vs. Azure, Spring AI provides a unified `ChatClient` interface. You can write your logic once and seamlessly switch between AI providers (OpenAI, Azure, Google Vertex AI, Anthropic, Hugging Face, Ollama) by just changing a configuration property, without rewriting code.
2. **Integration:** It seamlessly integrates AI models into traditional enterprise Java applications (e.g., Spring Boot, security, data access).
3. **Structured Outputs:** It provides robust mechanisms to take free-flowing text output from an LLM and map it directly into strongly-typed Java POJOs (Records/Classes) using Output Parsers.

## 2. Explain the concept of Prompts and Prompt Templates in Spring AI.
**Answer:**
- **Prompt:** The actual text instructions sent to an LLM to elicit a response. In Spring AI, a `Prompt` object is more than just a string; it encapsulates a list of `Message` objects.
- **Messages:** Different models support different roles (e.g., `SystemMessage` to set the AI's persona, `UserMessage` for the human's input, `AssistantMessage` for the AI's reply).
- **Prompt Template (`PromptTemplate`):** It is an engine that allows you to parameterize prompts. Instead of hardcoding "Tell me a joke about dogs", you write "Tell me a joke about {topic}". At runtime, you pass a `Map<String, Object>` containing `"topic" -> "dogs"`, and the template engine replaces the placeholders before sending the final string to the `ChatClient`. This is similar to rendering a Thymeleaf HTML document, but for AI prompts.

## 3. What is RAG (Retrieval-Augmented Generation), and why is it essential for Enterprise AI?
**Answer:**
LLMs (like ChatGPT) are frozen in time; they only know information up to their last training date. Crucially, they know absolutely *nothing* about your company's proprietary data (internal wikis, private databases, customer records). Furthermore, they can "hallucinate" (make up facts).

**RAG (Retrieval-Augmented Generation)** solves this by augmenting the LLM's prompt with dynamically retrieved, relevant context *before* asking it to answer.

**How it works (The RAG Pipeline):**
1. **The Question:** The user asks: "What is our company's refund policy?"
2. **Retrieval (The 'R'):** The application takes the question, converts it to a mathematical vector (an Embedding), and searches a Vector Database containing all the company's internal documents. It retrieves the top 3 most relevant paragraphs about "refund policies".
3. **Augmentation (The 'A'):** The application constructs a new prompt: "Answer the user's question using *only* the following context: [insert the 3 retrieved paragraphs here]. Question: What is our company's refund policy?"
4. **Generation (The 'G'):** The LLM receives the prompt, reads the context, and generates an accurate, factual response based purely on the provided proprietary data, eliminating hallucinations.

## 4. How do Vector Databases and Embeddings work in the context of Spring AI?
**Answer:**
To make RAG work, you need a way to mathematically compare the meaning of sentences.

- **Embeddings:** An embedding model (like OpenAI's `text-embedding-ada-002`) takes a piece of text (a word, a sentence, or a document) and converts it into a high-dimensional array of floating-point numbers (a Vector). The magical part is that texts with similar *meanings* end up with vectors that are physically close to each other in this multidimensional space. (e.g., "Dog" and "Puppy" are close; "Dog" and "Car" are far).
- **Vector Database:** A specialized database (like Chroma, Pinecone, PgVector, Milvus) built specifically to store, index, and query these high-dimensional vectors extremely fast.
- **Spring AI `VectorStore` Interface:** Spring AI provides the unified `VectorStore` abstraction.
    1. During data ingestion, you use a `DocumentReader` to parse internal files (PDFs, text files).
    2. You use a `DocumentSplitter` (TokenTextSplitter) to break large files into smaller chunks.
    3. The `VectorStore` automatically calls an `EmbeddingModel` to generate vectors for those chunks and saves them.
    4. At query time, `vectorStore.similaritySearch("refund policy")` converts the query to a vector, calculates the distance (e.g., cosine similarity) against all stored vectors, and returns the mathematically closest Documents.

## 5. What is Function Calling (or Tool Calling) in LLMs, and how does Spring AI support it?
**Answer:**
LLMs are essentially text-prediction engines. They cannot interact with the outside world (they can't query your SQL database, check live weather, or send an email). **Function Calling** bridges this gap.

**Concept:**
Instead of the LLM trying to answer a question it lacks data for, you provide it with a list of "Tools" (Functions) it is allowed to use. If you ask "What is the weather in London?", the LLM recognizes it needs live data. It pauses generation and replies, "Execute the `getWeather(location="London")` function." The Spring Boot application executes the Java method, gets the result ("15°C and raining"), and sends that result back to the LLM. The LLM then resumes generation: "The weather in London is currently 15°C and raining."

**Spring AI Implementation:**
1. You annotate a standard Java method returning a value (e.g., `WeatherResponse getWeather(WeatherRequest request)`) with `@Bean` and `@Description` (crucial so the LLM knows what the tool does).
2. You define the Bean as a `java.util.function.Function`.
3. When interacting with the `ChatClient`, you simply provide the bean name in the prompt options: `PromptOptionsBuilder.builder().withFunction("getWeatherFunction").build()`.
4. Spring AI handles the entire complex lifecycle of passing the tool schema to the LLM, catching the function execution request, calling the Java method, and returning the result to the model transparently.

## 6. What is the Model Context Protocol (MCP), and how does it relate to tools/agents?
**Answer:**
**Model Context Protocol (MCP)** is an emerging open standard championed by companies like Anthropic.

**The Problem:** Currently, if you build an AI application and want it to talk to Slack, GitHub, and Jira, you have to write custom integration code (API clients, authentication) for all three tools directly within your application (tight coupling).

**The MCP Solution:**
MCP standardizes how AI models access data sources and tools.
- It introduces an architecture consisting of **MCP Hosts** (your Spring AI application/agent) and **MCP Servers** (standalone, lightweight servers that wrap APIs).
- Instead of writing a Slack integration in your code, you run an off-the-shelf "Slack MCP Server." Your application (the Host) connects to that server using the standard MCP protocol.
- The MCP Server exposes its capabilities (e.g., "I can read Slack channels," "I can post messages"). The Host seamlessly passes these capabilities to the underlying LLM.
- **Benefits:** Massive ecosystem reuse. Developers can plug in hundreds of existing community-built MCP servers (for SQL databases, file systems, code repositories) into their Spring AI applications instantly without writing custom API integration code for each one.
