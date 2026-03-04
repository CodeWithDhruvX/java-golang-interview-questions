# Agentic AI Interview Questions for Service-Based Companies

Service-based companies (like TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini, etc.) focus on **rapid development, framework usage, client deliverables, and practical implementation**. The questions here test your knowledge of how to use tools like LangChain, AutoGen, and CrewAI to build Agentic workflows that meet client business requirements.

---

## 1. Core Concepts & Definitions

**Q1: What is the difference between RAG (Retrieval-Augmented Generation) and an Agentic AI workflow?**
*   **Focus Areas:**
    *   **RAG:** Linear flow. Retrieve context -> Augment prompt -> Generate. Has no agency to "decide" the next step.
    *   **Agentic workflow:** Non-linear flow. The LLM acts as the "brain," decides which tools to use, uses them, evaluates the results, and decides what to do next.

**Q2: What is an LLM Agent? Explain its primary components.**
*   **Focus Areas:**
    *   **Brain:** The LLM itself (e.g., GPT-4o, Claude 3.5 Sonnet).
    *   **Memory:** Short-term (chat history) vs. Long-term (Vector Databases like Pinecone, Milvus).
    *   **Tools:** External APIs, Calculators, Web Search, SQL query executors.
    *   **Planning:** Strategies like ReAct (Reasoning and Acting) to break down tasks.

**Q3: Can you define "Function Calling" (or Tool Calling) with respect to LLMs?**
*   **Focus Areas:**
    *   Instead of responding with plain text, the LLM outputs a structured JSON object.
    *   This JSON matches a predefined schema of a function.
    *   The application code intercepts the JSON, runs the real function, and returns the result to the LLM.

---

## 2. Frameworks & Tooling

**Q4: Have you used LangChain or LlamaIndex? Explain how you build an Agent in LangChain.**
*   **Focus Areas:**
    *   Understanding `AgentExecutor`.
    *   Binding predefined tools (like `WikipediaQueryRun`, `SerpAPI`, or custom `@tool` functions).
    *   Setting up the ReAct prompt template and the conversational memory block.

**Q5: What is LangGraph, and why is it preferred over traditional LangChain Agents for complex use cases?**
*   **Focus Areas:**
    *   LangChain Agents are black-box "loops" that are hard to debug or control if they go off track.
    *   LangGraph defines agents as state machines (graphs) with nodes and edges.
    *   It allows "Human-in-the-loop" pauses, persistent checkpoints (saving state), and cyclic routing.

**Q6: Describe a Multi-Agent system (e.g., AutoGen, CrewAI). Why use multiple agents instead of one large agent?**
*   **Focus Areas:**
    *   Different agents can have different system prompts, specializations, and access to different tools.
    *   Prevents context window overflow in a single agent.
    *   Example: A "Researcher Agent" gathers data, hands it to an "Analyst Agent" to process, which hands it to a "Writer Agent" to generate the report.

---

## 3. Practical Implementation Use Cases

**Q7: A client wants a Customer Support bot that can not only answer FAQs but also fetch user order details from a REST API and issue a refund. How would you design this?**
*   **Focus Areas:**
    *   Use a basic API tool for order lookup.
    *   Use a refund tool, but implement *Human-in-the-loop* where the agent drafts the refund and waits for admin approval before executing.
    *   Mention proper authentication (OAuth or API keys) injected at the application layer, not given to the LLM directly.

**Q8: What are some common failure modes for Agents, and how do you handle them?**
*   **Focus Areas:**
    *   **Hallucinating tools:** Requesting a tool that doesn't exist. Fix: Strict system prompting or strongly typed parsers (Pydantic).
    *   **Infinite loops:** The agent repeats the same failing action. Fix: `max_iterations`, timeouts, or explicit prompts telling the agent to "stop if you see this error twice."
    *   **Context window limits:** Fix: Summarize old chat history or drop oldest observations.

**Q9: How do you optimize costs when building agents for clients?**
*   **Focus Areas:**
    *   Using lower-cost models (GPT-3.5, Llama 3) for simple routing/tool-calling.
    *   Implementing a Router layer before the Agent.
    *   Avoiding agents for simple tasks; using RAG or standard API calls where deterministic logic works better.

---

## 4. Basic Prompt Engineering for Agents

**Q10: Explain what a ReAct prompt looks like. Provide a mock example.**
*   **Focus Areas:**
    *   The prompt usually gives instructions like:
    *   *Thought:* I need to find the user's order.
    *   *Action:* GetOrder(123)
    *   *Observation:* Order contains 2 apples.
    *   *Thought:* Now I need to calculate the refund.

**Q11: How do you format tool schemas so the LLM understands them accurately?**
*   **Focus Areas:**
    *   Providing clear descriptions in the JSON schema or Pydantic model (`description` fields are highly attended to by the LLM).
    *   Explaining what each argument does, e.g., instead of just `date`, naming it `booking_date_iso8601` with a descriptor.
