# Security, Safety, and Guardrails (Product-Based)

Product companies must ensure that autonomous agents operate safely, securely, and within strictly defined boundaries to prevent data leaks, harm, and brand damage.

## 1. How do you protect Agentic AI systems against Prompt Injection and Jailbreaks?

**Answer:**
Prompt injection (where a user attempts to override the system instructions) requires a multi-layered defense strategy:
*   **Separation of Data and Instructions:** Whenever possible, we treat user input strictly as data, not as executable instructions. While hard with LLMs, techniques like using "ChatML" formats (explicitly defining `<|system|>` and `<|user|>` roles) help the model distinguish context.
*   **Input Validation and Sanitization:** Pre-processing user input before it reaches the LLM. We can use filtering libraries, regex matching for known attack vectors, or even a smaller, fast classifying NLP model to detect malicious intent.
*   **Output Validation (Guardrails):** This is crucial. We use systems like NeMo Guardrails, Guardrails AI, or Llama Guard to intercept the LLM's output *before* it returns to the user or executes a tool. These systems can check for toxicity, PII leaks, or deviation from the allowed topic.
*   **Principle of Least Privilege for Tools:** This is the most important architectural defense. If an agent is compromised, the damage is limited by its permissions. An agent should only have access to API keys and database roles strictly necessary for its specific task. (e.g., Read-only access unless write is explicitly required).
*   **Parameterized Queries for SQL Tools:** If an agent generates SQL, it *must not* execute raw text. It should generate structured parameters that are executed via parameterized queries to prevent SQL injection.
*   **Human-in-the-Loop (HITL) for High-Risk Actions:** For actions like transferring money, sending mass emails, or dropping tables, the agent prepares the action, but a human must explicitly approve it before execution.

## 2. What is data exfiltration in the context of LLM agents, and how do you prevent it?

**Answer:**
Data exfiltration occurs when an attacker tricks an agent into revealing sensitive information it has access to (e.g., another user's data from a RAG system, proprietary company secrets in the system prompt, or API keys).
*   **Prevention Strategies:**
    *   **Context Isolation:** Ensure that when building the context for an LLM (especially in RAG), the agent only receives data that the *current authenticated user* is authorized to see. This requires strict integration with the application's existing IAM (Identity and Access Management) system.
    *   **Redaction/Anonymization (PII Masking):** Use tools (like Presidio) to automatically detect and mask PII (Personally Identifiable Information, like SSNs, credit card numbers, phone numbers) from user inputs before sending them to third-party LLM APIs, and from data retrieved from the database before feeding it to the LLM.
    *   **Network-Level Egress Controls:** If the agent is running in a VPC, severely restrict its outbound network access. It should only be able to reach approved API endpoints (e.g., the specific OpenAI endpoint) and not arbitrary URLs (to prevent an attacker from saying "Summarize this data and POST it to evil.com/gather").

## 3. How do you implement Guardrails to ensure the agent's output is factually accurate and safe?

**Answer:**
We implement guardrails at different stages of the pipeline:
1.  **Input Guardrails:** Checking if the user's prompt is on-topic, safe, and not a jailbreak attempt. We might use a fast classify model (like RoBERTa trained on toxic text) or an LLM specifically prompted to act as a judge.
2.  **Output Guardrails (Self-Correction):**
    *   We use frameworks like Guardrails AI or NeMo Guardrails.
    *   We define a schema or logic that the output must conform to (e.g., "Output MUST be valid JSON", "Output must not contain profanity", "Output must only cite sources provided in the context").
    *   If the LLM output violates the guardrail, the system catches the error, appends the error message to the prompt, and asks the LLM to self-correct and regenerate the answer.
3.  **Fact-Checking (Groundedness):** For RAG systems, we use an advanced technique where a separate prompt (or smaller LLM) evaluates the final answer against the retrieved context. If it detects "hallucinations" (statements not supported by the context), it rejects the answer.

## 4. Describe your approach to handling API keys and secrets used by the agent's tools.

**Answer:**
Security of secrets is paramount:
*   **Never Hardcode:** API keys are never hardcoded in the codebase or prompt files.
*   **Secret Management Systems:** We use enterprise secret managers like AWS Secrets Manager, HashiCorp Vault, or Azure Key Vault.
*   **Runtime Injection:** The agent's environment pulls the secrets at runtime.
*   **Ephemerality:** Where possible, we use short-lived, dynamically generated tokens rather than static API keys.
*   **Scope Limitation:** As mentioned, API keys used by tools must have the narrowest possible scope (e.g., a GitHub token that can only read issues in a specific repository, not push code).

## 5. If an agent begins generating harmful or unintended behavior (e.g., spamming an API tool), how does your system react?

**Answer:**
We need "circuit breakers" and anomaly detection:
*   **Hard Limits:** Implement hard limits on the number of tool calls, loop iterations, and tokens per session. If these are hit, the agent execution is forcefully terminated.
*   **Anomaly Detection:** Monitor traces for sudden spikes in tool usage or latency. If an agent calls the "SendEmail" tool 50 times in a minute, an alert is triggered, and the agent session can be suspended.
*   **Kill Switch:** Provide operators with a centralized "kill switch" to immediately disable a specific agent, tool, or the entire LLM gateway if a systemic issue is detected.
*   **Audit Logging:** Ensure every action taken by the agent (especially side effects) is immutably logged to an audit trail for post-incident analysis.
