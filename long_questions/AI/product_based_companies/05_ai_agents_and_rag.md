# AI Agents and Advanced RAG (Product-Based Companies)

Building Agentic workflows is the current frontier. Top companies expect you to understand how to build systems where the LLM is the reasoning engine controlling tools, rather than just a chatbot.

## 1. What is an AI Agent? Explain the ReAct (Reasoning + Acting) framework.
**Answer:**
An AI Agent is an autonomous or semi-autonomous system where an LLM is given an objective, planning capabilities, memory, and access to external tools (APIs, calculators, web search, execution environments). It continuously loop-processes information until it completes the goal.

**The ReAct Framework:**
ReAct is a prompting structure that interleaves "Chain-of-Thought" reasoning with taking continuous actions in an environment. It forces the LLM to write out its thought process *before* deciding on an action, dramatically reducing hallucinations and getting stuck in loops.

**The Loop:**
1.  **Thought:** The LLM observes the current state and reasons about what to do next. (e.g., "I need to find the population of Paris, then multiply it by 2.")
2.  **Action:** The LLM outputs a structured command to call a specific tool. (e.g., `Action: Search[population of Paris]`)
3.  **Observation:** The system executes the tool in reality and returns the result to the LLM. (e.g., `Observation: 2.16 million`)
4.  *(Repeat)* **Thought:** The LLM analyzes the observation. (e.g., "The population is 2.16 million. Now I use the calculator.") -> **Action:** `Action: Calculator[2.16 * 2]` -> **Observation:** `4.32`
5.  **Finish:** Once the core goal is met, the LLM outputs the final answer.

## 2. When building a multi-agent system (e.g., using frameworks like AutoGen or CrewAI), what are the typical architectures and communication patterns?
**Answer:**
Multi-agent systems break down complex problems by delegating tasks to specialized agents (e.g., a "Researcher Agent," a "Coder Agent," and a "Reviewer Agent").

**Common Architectures:**
1.  **Sequential (Pipeline):** The simplest format. Agent A finishes its task (e.g., Research) and passes the output to Agent B (e.g., Writing), who passes it to Agent C (e.g., Quality Verification).
2.  **Hierarchical (Manager-Worker):** A "Manager/Router Agent" sits at the top. It parses the user request, breaks it into subtasks, and dynamically assigns them to specialized worker agents based on their toolsets. It then synthesizes their results into a final response. This matches corporate structures.
3.  **Joint Collaboration (Group Chat):** All agents are placed in a shared virtual room. When a message is posted, a "Speaker Selection" algorithm determines which agent should speak next based on the persona rules and context. This allows for debate (e.g., Coder pushes code, Reviewer critiques it, Coder fixes it dynamically).

**Challenges:**
*   **Infinite Loops:** Agents agreeing with each other endlessly or continuously re-trying failed tool calls. Requires strict iteration limits and graceful failovers.
*   **Context Window Overflow:** If a group chat gets too long, the context window fills up, breaking the LLM's reasoning. Requires "Summary Agents" to periodically compress the conversation history.

## 3. Beyond vector databases, how does GraphRAG improve retrieval on complex document sets?
**Answer:**
**The limitations of standard Vector RAG:**
Standard RAG relies strictly on semantic similarity to text chunks. It answers "What is topic X?" well, but fails utterly at "Connect the dots" or "global" questions (e.g., "What are the common underlying financial themes across all 40 earnings reports?"). Vector search only retrieves *local* information and often misses the big picture.

**The GraphRAG Approach:**
Developed heavily by Microsoft, GraphRAG builds Knowledge Graphs *before* retrieval.

1.  **Extraction (Offline):** Instead of just chunking text, we pass the entire corpus through an LLM and instruct it to extract **Entities** (e.g., "Company A", "CEO Bob") and **Relationships** (e.g., "Company A [ACQUIRED] Startup B").
2.  **Graph Construction:** We build a massive Knowledge Graph connecting these entities.
3.  **Community Detection (Clustering):** Graph algorithms (like Leiden) group tightly connected nodes into hierarchical "Communities" (Community 1: Tech Mergers, Community 2: Executive changes). We use an LLM to pre-generate narrative summaries of every community at every level.
4.  **Retrieval (Online):**
    *   *Local Queries:* Use standard vector search combined with graph traversal (fetching neighboring nodes of retrieved entities).
    *   *Global Queries:* If a user asks "What are the main risks mentioned across the entire company?", the system bypasses vectors and instead queries the *pre-computed community summaries* across the entire network, using MapReduce with an LLM to generate the final synthetic answer. It provides accurate, comprehensive answers that native vector math simply cannot execute.

## 4. How do you implement "Memory" in AI Agents so they remember vast user histories across multiple sessions?
**Answer:**
LLMs are inherently stateless. "Memory" is a software engineering problem of retrieving and injecting the right context into the prompt.

1.  **Short-Term Memory (Conversation Buffer):**
    Simply appending the last N messages to the prompt. To prevent overflow, use a "Sliding Window" (keep last 10) or a "Summary Buffer" (use an LLM in the background to summarize older messages into a single paragraph and append that to the top of the prompt).
2.  **Long-Term Semantic Memory (Vector Stores):**
    When a user says a fact ("My dog's name is Fido"), the system embeds that sentence and stores it in a Vector Database tied to the User ID. When the user later asks, "What should I feed my dog?", the system queries the DB with the user's intent, retrieves "Fido," and injects it into the system prompt behind the scenes: `<context>User owns a dog named Fido</context>`.
3.  **Knowledge Graph Memory (Entity Repositories):**
    Tools like Mem0 or Zep parse interactions into a continually updating graph format. If the user mentions working at Google, the system creates a node `[User] -> [Works At] -> [Google]`. This allows the agent to query structured relationships rather than just raw text embeddings, enabling highly personalized, logically consistent user profiles over years of usage.
