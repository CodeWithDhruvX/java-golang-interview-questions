# Agentic AI Interview Questions for Product-Based Companies

Product-based companies (like Google, Microsoft, Meta, Amazon, Stripe, Databricks, etc.) focus on **deep technical understanding, system design, scalability, robustness, and productionizing Agentic AI**. The questions here will test your fundamental understanding of Agent architectures, evaluation, reliability, cost management, and latency optimization.

---

## 1. System Design & Architecture of AI Agents

**Q1: Design an autonomous Coding Agent (similar to GitHub Copilot Workspace or Devin) that can take an issue description, navigate a codebase, write code, and submit a PR. Walk me through the architecture.**
*   **Focus Areas:**
    *   **Orchestration layer:** What framework/pattern? (ReAct, Plan-and-Solve, LLM Compiler).
    *   **Component breakdown:** Planner, Coder, Reviewer, Environment simulator.
    *   **Memory management:** Short-term (context window) vs. Long-term (vector DBs, AST parsing strategies).
    *   **Action Space:** How does the agent execute terminal commands safely? Docker containers, sandboxing, gRPC interfaces.
    *   **Failure handling:** How does it recover from compilation errors? Self-reflection loops.

**Q2: How do you design an Agentic System that requires querying large-scale, live relational databases (Natural Language to SQL) while maintaining low latency and avoiding hallucinations in the schema?**
*   **Focus Areas:**
    *   Schema pruning (injecting only relevant tables/schemas into the prompt).
    *   Few-shot prompting with golden SQL examples.
    *   Agent loop that verifies the query format before executing against a read replica.
    *   Handling query timeouts or "table not found" errors via an error-correction loop.

**Q3: Describe the ReAct (Reasoning + Acting) pattern. What are its limitations in production at scale, and how would you overcome them?**
*   **Focus Areas:**
    *   ReAct traces: `Thought -> Action -> Observation -> ... -> Result`.
    *   *Limitations:* High token consumption, latency due to sequential LLM calls, getting stuck in infinite loops.
    *   *Solutions:* Multi-agent patterns, caching observations, early exit heuristics, fine-tuning a smaller model to predict actions without full ReAct reasoning.

---

## 2. Productionization, Reliability & Scale

**Q4: Your autonomous agent has a 15% rate of getting stuck in infinite loops (e.g., repeating the same failed API call). How do you detect, mitigate, and resolve this in a production system?**
*   **Focus Areas:**
    *   **Detection:** Implementing a max-step threshold, monitoring duplicate action/observation sequences, tracking transition entropy.
    *   **Mitigation:** Circuit breakers, fallback to human-in-the-loop (HITL), explicitly prompting the agent with its history to force a new approach.
    *   **Resolution:** Fine-tuning base models on negative trajectories so they learn to avoid loops.

**Q5: How do you handle "State" in multi-agent workflows (e.g., using frameworks like LangGraph, AutoGen, or Temporal)?**
*   **Focus Areas:**
    *   Stateless stateless conversational turns vs. persistent graphs.
    *   Saving agent states in Redis/Postgres for resume-ability.
    *   Checkpointing in LangGraph.
    *   Handling distributed state if agents are running as separate microservices.

**Q6: What is the cost-latency tradeoff when designing an Agentic workflow, and how do you optimize it?**
*   **Focus Areas:**
    *   Using Large/Expensive models (GPT-4) for the "Planner/Router" agent.
    *   Using Small/Fast models (Llama-3-8B, GPT-3.5) for "Worker" tools (e.g., parsing JSON, summarizing).
    *   Semantic caching (RedisVL) to avoid re-computing identical agent paths.
    *   Prompt caching features in modern API providers.

---

## 3. Tool Calling & External Integrations

**Q7: How does an LLM actually execute a tool under the hood? (Function calling / Tool use mechanisms)**
*   **Focus Areas:**
    *   Explain how tools are represented in the prompt (JSON schema, OpenAPI spec).
    *   How the LLM outputs unstructured text vs. structured JSON.
    *   The role of system prompts in enforcing tool schemas.
    *   Fine-tuning for tool-use (e.g., Gorilla, ToolLLaMA).

**Q8: Explain how you would implement "Graceful Degradation" if an external API tool that your agent relies on goes down.**
*   **Focus Areas:**
    *   Standard software engineering practices (retries, backoff, circuit breakers).
    *   Agent-specific practices: Providing the agent with alternative tools dynamically.
    *   Prompting the agent to "skip" non-essential steps and provide a partial but accurate answer.

---

## 4. Evaluation & Guardrails

**Q9: How do you evaluate an Agentic system? Standard RAG metrics (like RAGAS) don't apply directly to agents making sequential decisions.**
*   **Focus Areas:**
    *   **Trajectory Evaluation:** Evaluating the *path* the agent took, not just the final answer.
    *   **Reward Modeling:** Using an "Evaluator LLM" (LLM-as-a-judge) to score the efficiency and safety of the steps.
    *   **Deterministic Benchmarks:** WebArena, SWE-bench – how they work conceptually.
    *   Measuring Action Success Rate, Token Cost per Task, and Latency per Task.

**Q10: Design the Security Guardrails for an Agent that has access to write to a company's internal Confluence and Slack.**
*   **Focus Areas:**
    *   **Input Guardrails:** Prompt injection detection (Llama-Guard, NeMo Guardrails).
    *   **Output/Action Guardrails:** Rule-based verification before a tool allows a write action.
    *   **Human-in-the-loop (HITL):** For any destructive or highly visible action, the agent drafts the action and pauses for human approval.
    *   Principle of Least Privilege (PoLP) in agent service accounts.

---

## 5. Advanced / Cutting-Edge Concepts

**Q11: What is the difference between a Plan-and-Execute agent and an iterative routing agent? When would you use which?**
*   *Plan-and-Execute:* Creates a full DAG of tasks up front, then executes. Good for deterministic, known workflows.
*   *Iterative/ReAct:* Decides the next step based on the previous observation. Good for exploratory tasks (like debugging).

**Q12: How would you approach fine-tuning an open-source model (like Llama 3) to be a better domain-specific agent?**
*   *Focus Areas:* Generating synthetic trajectories using a strong teacher model (GPT-4). Formatting data as multi-turn conversations with tool calls. Using DPO (Direct Preference Optimization) to penalize hallucinated tool parameters.

**Q13: Explain "Self-Reflection" (or Reflexion) in the context of Agents.**
*   *Focus Areas:* An agent makes an attempt, receives a failure (e.g., compile error), and uses a separate prompt to analyze *why* it failed before retrying. Tradeoffs: greatly increases token cost and latency for improved accuracy.
