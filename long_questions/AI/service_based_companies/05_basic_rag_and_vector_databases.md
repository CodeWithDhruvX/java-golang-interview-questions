# Basic RAG and Vector Databases (Service-Based Companies)

Retrieval-Augmented Generation (RAG) is the most common AI use-case you will build for enterprise clients today. You must understand how to ground an LLM in a company's private data securely.

## 1. Explain the step-by-step process of building a Retrieval-Augmented Generation (RAG) system over a client's HR PDF documents.
**Answer:**
Building a standard RAG system involves five main steps, easily implemented using libraries like LangChain or LlamaIndex:

1.  **Document Loading:** Use a document loader (e.g., PyPDFLoader) to extract raw text from the HR PDFs into memory.
2.  **Splitting/Chunking:** An LLM cannot process a 100-page PDF at once due to the context window limit. You split the document into smaller "chunks" (e.g., 1000 characters each), typically adding some "overlap" (e.g., 200 characters) between chunks so you don't accidentally cut a sentence in half, preserving semantic meaning.
3.  **Embedding:** You pass every single chunk through an Embedding Model (like OpenAI's `text-embedding-ada-002` or an open-source Hugging Face model). This model converts the English text into an array of thousands of floating-point numbers (a Vector) that mathematically represents the meaning of the text.
4.  **Vector Store (Insertion):** You insert these vectors, along with the original text chunk and metadata (like the PDF filename), into a Vector Database (like Pinecone, ChromaDB, or FAISS).
5.  **Retrieval & Generation (Querying):**
    *   When an employee asks, "What is the maternity leave policy?", you embed that query into a vector using the *exact same* embedding model.
    *   You query the Vector Database to find the $K$ (e.g., 3) vectors most mathematically similar to the query vector (using Cosine Similarity). The database returns the 3 most relevant text chunks.
    *   Finally, you construct a prompt: *"Answer the user's question based ONLY on the following context. Context: [Insert 3 chunks here]. Question: [User's question]."*. You send this prompt to an LLM (like GPT-4) to generate the final human-readable answer.

## 2. What is an Embedding? How do Vector Databases perform searching?
**Answer:**
**Embeddings:**
An embedding is a numerical representation of data (text, images, audio) as a vector in a high-dimensional mathematical space (often 384, 768, or 1536 dimensions). The crucial feature of embeddings is that *semantically similar concepts are located physically close to each other in this space*.
*   *Example:* The vector for "King" is close to the vector for "Queen," and both are far away from the vector for "Carburetor."

**Vector Database Search:**
Traditional relational databases (SQL) perform exact keyword matching (e.g., `WHERE word = 'leave'`). Vector databases perform **Semantic Search** (Approximate Nearest Neighbor search).
When you query the database, it doesn't look for matching letters. It calculates the geometric distance between the query vector and all the stored document vectors. The most common metric for this distance is **Cosine Similarity**, which measures the angle between the two vectors. An angle of 0 degrees (Cosine = 1) means they point in the exact same direction and are perfectly semantically identical. The database returns the text chunks with the highest cosine similarities to the query.

## 3. Why is LangChain so heavily used in industry right now? What are its core components?
**Answer:**
**Why it's used:**
LangChain acts as "middleware" or "glue" for LLM applications. Instead of writing raw HTTP requests to OpenAI and manually managing prompt strings in Python dictionaries, LangChain provides a standardized, object-oriented framework. It makes it extremely easy to swap out an OpenAI model for an open-source Llama model with just one line of code change.

**Core Components:**
1.  **Models (LLMs/ChatModels):** The wrapper classes for the actual APIs (OpenAI, Anthropic, HuggingFace).
2.  **Prompts:** Templates for creating dynamic prompts. Instead of f-strings (`f"Hello {name}"`), you define formal `PromptTemplates` that ensure inputs are properly escaped and formatted depending on the model type.
3.  **Agents & Tools:** The framework allows you to give the LLM "tools" (like a Google Search function or a SQL connection). The Agent parses the user request, decides *which* tool to use, executes the Python code for that tool, and feeds the result back into the LLM.
4.  **Chains:** The core of LangChain. A chain links multiple components together sequentially. For example, a simple chain combines a `PromptTemplate` -> `LLM` -> `OutputParser`. When you call `chain.invoke(input)`, execution flows automatically.
5.  **Memory:** Since APIs are stateless, Memory classes (like `ConversationBufferMemory`) automatically handle retrieving chat history and injecting it into the prompt.

## 4. In a RAG application, your user complains that the AI frequently provides "hallucinated" answers not found in the PDFs. How do you fix this?
**Answer:**
Hallucinations in RAG usually mean the LLM is ignoring the retrieved context and relying on its pre-trained knowledge.

1.  **Stricter System Prompting:** The prompt is the most critical fix. Instead of "Answer the question based on the context," you must be draconian: *"You are an assistant. You MUST answer the query using ONLY the provided context. If the answer is not explicitly contained in the context, you MUST reply: 'I do not have enough information to answer this based on the provided documents.' DO NOT provide external facts."*
2.  **Citation Requirement:** Force the LLM to cite its sources. Ask it to output JSON: `{"answer": "...", "page_number": "..."}`. If it cannot find a page number in the provided context chunk's metadata, it is less likely to guess.
3.  **Low Temperature:** Ensure the `temperature` parameter via the API is set to `0.0`. You want zero creativity when parsing corporate documents.
4.  **Check the Retrieval:** Often, it's a retrieval failure, not an LLM failure. If the Vector DB returned terrible, irrelevant chunks because the semantic search failed, the LLM will try its best to answer and end up hallucinating. You must debug the retrieved chunks directly to ensure they actually contained the answer in the first place.
5.  **Hybrid Search:** If users are searching for exact part numbers (e.g., "TX-994A"), vector search often fails. Implementing Hybrid Search (combining Vector Embeddings with raw keyword BM25 search) ensures an exact match is retrieved and passed to the context window.

## 5. What are the cost implications of running an LLM application in production, and how do you optimize them?
**Answer:**
LLM API pricing (like OpenAI) is based on **Tokens**. You pay for "Input/Prompt Tokens" (everything you send) and "Output/Completion Tokens" (everything the model generates). Output tokens are usually much more expensive.

**Optimization Strategies:**
1.  **Reduce Prompt Size:** Every RAG request sends the whole System Prompt + Retrieved Context + History.
    *   Only retrieve the absolute minimum chunks needed (e.g., Top 3 instead of Top 10).
    *   Use a "Summary" memory instead of passing the entire raw chat history back every time.
2.  **Cheaper Models for Simpler Tasks:** You don't need GPT-4 to summarize a simple email or classify an intent. Use cheaper, faster models (like GPT-3.5-Turbo, Claude Haiku, or local models) for simple routing tasks, and only invoke the expensive GPT-4 when complex reasoning is required.
3.  **Semantic Caching:** If 50 users ask the bot "How do I reset my password?" in one morning, you shouldn't pay the API 50 times. Use a caching layer (like Redis + a local Embedding Model). Embed the new incoming question; if it's 95% semantically similar to a previous question in the cache, instantly return the cached answer and save the LLM API cost entirely.
4.  **Limit Output:** Output tokens are the most expensive. Instruct the LLM in the system prompt to "Keep responses extremely concise and under 3 sentences." Use `max_tokens` API parameters as a hard cutoff.
