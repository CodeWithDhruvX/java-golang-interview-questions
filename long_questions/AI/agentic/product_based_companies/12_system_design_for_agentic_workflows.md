# System Design for Agentic Workflows (Product-Based)

System design interviews for AI roles at product companies focus on how to architect robust, scalable, and intelligent systems combining traditional backend infrastructure with LLM-driven components.

## 1. Design an automated, intelligent customer support triage and resolution system.

**Candidate Approach:**

**1. Requirements Gathering:**
*   **Functional:** System receives support tickets (email, chat). It must categorize the ticket, attempt to resolve it automatically using knowledge bases and internal APIs (billing, order status), and route to a human agent if unresolved, providing a summary.
*   **Non-Functional:** High availability, low latency (especially for chat), secure handling of customer PII, scalable to thousands of concurrent tickets.

**2. High-Level Architecture (Multi-Agent System):**
I would design a multi-agent system using a routing architecture.
*   **Ingestion API:** Webhook receiver for emails and chat messages. Placed behind an API Gateway (e.g., AWS API Gateway) and Load Balancer.
*   **Message Queue:** Incoming tickets are dropped onto a Kafka/SQS topic to decouple ingestion from processing and handle traffic spikes.
*   **Orchestrator Agent (The Router):** An LLM agent consumes messages from the queue. Its sole job is classification and routing.
    *   **Prompt:** Analyzes text, extracts intent (e.g., "refund", "technical issue", "login problem") and urgency.
    *   **Action:** Routes the ticket to the appropriate specialized worker agent via secondary queues.
*   **Specialized Worker Agents:**
    *   **Billing Agent:** Has tools to query the billing database (SQL) and process refunds via Stripe API.
    *   **Tech Support Agent:** Has RAG tools connected to the technical documentation vector database (e.g., Pinecone, Milvus) and tools to query system status APIs.
*   **Human Handoff Layer:** A service that manages the transition to a human. Pushes the conversation history and the agent's summary to a CRM (like Zendesk or Salesforce).

**3. Deep Dive: The Tech Support Agent (RAG + Tools)**
*   When a ticket arrives, the Tech Support agent executes a ReAct loop.
*   *Action 1:* Queries the Vector DB with the user's problem.
*   *Action 2:* Evaluates if the retrieved chunks solve the problem. If yes, generate a response.
*   *Action 3:* If the issue is related to a server being down, call the `CheckSystemStatusBoard` API tool.
*   *Action 4:* Construct the final reply based on RAG and tool data.

**4. Data Management & State:**
*   **Session State:** Chat interactions require short-term memory. We use Redis to store the conversation history for the active session, retrieved and appended to the prompt for each turn.
*   **Long-Term Storage:** Final resolutions and structured data (ticket category, resolution time) are stored in a relational database (PostgreSQL) for analytics.

**5. Observability and Guardrails:**
*   **Output Guardrails:** Before sending a reply to the customer, the response passes through a NeMo Guardrails layer to ensure no abusive language or hallucinated policies.
*   **Tracing:** LangSmith/Langfuse is integrated to trace the Orchestrator's routing decisions and the Worker agents' tool calls to debug "why did this ticket not get resolved?"

## 2. Design a system for an AI Coding Assistant (like GitHub Copilot or Cursor) that operates on an entire enterprise codebase.

**Candidate Approach:**

**1. Requirements:**
*   **Core Function:** User asks a question ("How is authentication handled?") or requests a feature ("Add a retry mechanism to the database client"). The system must analyze the entire relevant codebase and provide accurate answers or code diffs.
*   **Challenges:** Context window limits (cannot fit millions of lines of code into a prompt), latency, access control (users can only query code they have permissions for).

**2. Architectural Flow (Advanced RAG for Code):**
The challenge is retrieval. Standard text RAG performs poorly on code structure.

*   **Ingestion Pipeline (Asynchronous):**
    *   Hook into the Git repository (e.g., via webhooks on commit).
    *   **Parsing:** Use tools like Tree-sitter to parse the code into Abstract Syntax Trees (ASTs). This allows us to extract meaningful chunks: complete functions, classes, and interfaces, rather than arbitrary text splits.
    *   **Embedding & Storage:** Embed these chunks using a code-optimized embedding model (e.g., Voyage Code or OpenAI `text-embedding-3`). Store them in a Vector Database.
    *   **Graph Database (Crucial):** Store the *relationships* between components in a Graph Database (like Neo4j). Example Node: `Function A`, Edge: `Calls`, Node: `Function B`.

*   **Query Processing:**
    *   User asks: "Where is the password hashed in the auth module?"
    *   **Query Reformulation Layer:** An LLM refines the user query into precise search terms.
    *   **Hybrid Retrieval:**
        1.  *Vector Search:* Find semantically similar code snippets (e.g., looking for "cryptography", "hash", "bcrypt").
        2.  *BM25 (Keyword) Search:* Find exact matches for terms like "auth module".
        3.  *Graph Traversal:* If the vector search finds the `LoginController`, use the Graph DB to trace what services it calls to find the actual hashing function.
    *   **Reranking:** Use a Cross-Encoder model to rerank the combined results based on relevance to the user's specific query.

*   **Generation (The Agent Phase):**
    *   The top, most relevant context chunks (functions, docstrings) are provided to the generation LLM (e.g., Claude 3.5 Sonnet or GPT-4o).
    *   If the user requested a change, the LLM acts as an agent, outputting specific edit commands (e.g., line-by-line diffs) rather than just dumping whole files.

**3. Enhancing Accuracy (Agentic Loop):**
For complex tasks, we don't just generate an answer immediately. We implement a sub-agent loop:
*   Retrieve initial context -> Let the agent review it -> If the agent realizes a crucial file is missing based on imports, it uses a `SearchFile` tool to retrieve more context -> Synthesize final answer.

## 3. How would you design a cost-management infrastructure for a wildly popular generative AI application that uses expensive APIs (GPT-4)?

**Candidate Approach:**

When scaling, API costs can become prohibitive. The design must focus on caching, routing, and tiering.

**1. The Semantic Caching Layer:**
*   Place a Semantic Cache (e.g., Redis with vector search capabilities + an embedding model) immediately behind the API Gateway.
*   When a request comes in ("Summarize World War 2"), embed the query.
*   Search the cache. If a query with high cosine similarity (e.g., "Give me a summary of WWII") exists, return the cached answer immediately. This costs $0 in LLM API fees and has near-zero latency.

**2. LLM Gateway and Routing (The "Router Model"):**
*   All LLM calls go through an internal LLM Gateway (e.g., LiteLLM or an internal service).
*   **Dynamic Routing based on complexity:** We train a very small, fast classifier model (or use an LLM specifically prompted to act as a judge).
*   *Simple tasks* (e.g., translating a sentence, formatting text, basic chit-chat) are routed to a cheap, fast model (e.g., Llama 3 8B, GPT-3.5-Turbo, Claude Haiku).
*   *Complex tasks* (e.g., complex coding, advanced reasoning, tool calling) are routed to the expensive, highly capable model (e.g., GPT-4, Claude 3 Opus).

**3. Asynchronous Processing and Batching:**
*   For non-real-time tasks (e.g., generating nightly reports, bulk summarization of reviews), do not use synchronous API calls.
*   Queue these requests and use the Batch APIs provided by OpenAI/Anthropic (often 50% cheaper) to process them asynchronously during off-peak hours.

**4. Prompt Economics:**
*   Implement systems to minify prompts before sending them.
*   Remove unnecessary whitespace, use shorter variable names in JSON tool constraints, and ensure conversation history is aggressively truncated or summarized as it grows.

**5. Visibility and Quotas:**
*   Attach metadata (Tags/Labels) to every API call via the Gateway, indicating the user tier, internal feature ID, and team.
*   Implement hard quotas per user tier.
*   Build dashboards tracking cost-per-feature and cost-per-user to identify anomalies (e.g., users finding ways to burn compute or inefficient prompt structures).
