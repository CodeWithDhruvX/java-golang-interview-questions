# Mock Interview Practice – Agentic AI (Product & Service-Based Companies)

Practice answering these questions out loud or in writing. Each question has a sample strong answer and key points to cover. This simulates the actual interview experience.

---

## ROUND 1: INTRODUCTORY SCREENING (Both Company Types)

---

**Q1. What is an AI Agent? Explain it to a non-technical person.**

✅ **Strong Answer:**
*"An AI agent is like a smart assistant that can complete tasks on its own — not just answer questions, but actually take steps to get things done. Imagine you ask a regular ChatGPT 'book me a flight to Mumbai' — it might tell you how to do it. An AI agent would actually go to the booking website, search for flights, pick the best one based on your preferences, and confirm the booking — all without you doing each step manually. It breaks goals into smaller steps, uses tools like websites and databases, and keeps working until it finishes the task."*

**Key Points:** Autonomy, tool use, multi-step goal completion. Avoid jargon.

---

**Q2. What's the difference between LangChain and LangGraph?**

✅ **Strong Answer:**
*"LangChain is a framework for building LLM-powered applications and agents. It gives you tools, memory, and chains for connecting components. LangGraph is built on top of LangChain and adds the ability to define agent workflows as state machines — meaning you explicitly draw out nodes (steps) and edges (transitions between steps), including loops. The key advantage of LangGraph is control: you can add conditional branching, human-in-the-loop pauses, and checkpointing that lets you resume a workflow from where it left off. For simple agents, LangChain's AgentExecutor is fine. For complex, production workflows with multiple decision points, LangGraph is preferred."*

---

**Q3. What is RAG? How is it different from an agent?**

✅ **Strong Answer:**
*"RAG — Retrieval-Augmented Generation — is a technique where, before generating a response, you first retrieve relevant documents from a knowledge base and include them in the prompt so the LLM can answer based on real, up-to-date information rather than just its training data. It's a one-shot, linear pipeline: retrieve → augment → generate. An agent is more dynamic — it decides for itself what to do next based on the previous result. An agent might use RAG as one of its tools, but it can also call APIs, run code, or take other actions. RAG is a component; an agent is an orchestrator that might use RAG as one capability."*

---

## ROUND 2: TECHNICAL DEPTH (Product-Based Focus)

---

**Q4. I'm going to describe a failure scenario. Tell me how you'd debug it.**

*Scenario: "Your production coding agent successfully generates code 90% of the time, but for 10% of tasks it just keeps retrying the same failing approach over and over until it hits max_iterations."*

✅ **Strong Answer:**
*"My first step is diagnosis, not guessing fixes. I'd pull all 'max_iterations exceeded' traces from LangSmith and look for patterns — are these failures on specific file types? Specific error messages? Specific LLM models? I'd bet the agent is not reading the tool error output properly, just repeating because it doesn't understand the failure.*

*To confirm: I'd look at one full trace. If I see `Thought: I need to write to file.py → Action: write_code → Observation: [Error: syntax error on line 12] → Thought: I need to write to file.py` — the agent is ignoring the error. Fix: modify the system prompt to explicitly instruct the agent to analyze error messages before retrying. Also add a deduplication check: if the same (action, args) pair appears twice in the agent's trajectory, inject 'You have tried this approach twice and it failed. Take a completely different approach.' Finally, add a Prometheus alert for step count > 8 so we catch this early in future."*

---

**Q5. Walk me through the system design of a multi-agent pipeline to process insurance claims.**

✅ **Strong Answer:**
*"I'd design a pipeline with specialized agents and clear handoffs. First, a Document Intake Agent that reads the claim PDF, extracts structured information (claimant name, policy number, incident date, claimed amount) using a structured output parser — Pydantic models, not free text. Second, a Validation Agent that checks the policy database — is the policy active? Is the claim within the coverage period? Is the claimed amount within limits? It calls internal APIs with strict auth.*

*Third, based on validation: if straightforward and below ₹50,000 — auto-approve with an Approval Agent that writes to the claims DB. If above ₹50,000 or flagged as suspicious — route to a Fraud Detection Agent that checks historical claim patterns. If fraud score is high — escalate to a human reviewer queue.*

*For the infrastructure: LangGraph state machine with checkpointing, each step's decision is logged for auditability (required for IRDAI compliance). Maximum latency: document intake (10s), validation (2s), approval (1s). Human escalation within 24 hours. All PII is stripped from LLM prompts — agents work with claim IDs, not names or policy numbers directly."*

---

**Q6. What is DPO and how would you use it to improve your agent's tool-calling accuracy?**

✅ **Strong Answer:**
*"DPO — Direct Preference Optimization — is a fine-tuning technique where you train the model on pairs of (chosen, rejected) responses. The model learns to prefer chosen over rejected. For tool-calling accuracy, I'd collect examples where the agent called a tool with wrong arguments — like passing a date as 'March 10' instead of '2026-03-10' — and pair them with a corrected version. My dataset would have entries like: prompt = [user query + tool schema], chosen = agent correctly calls `get_order(order_id='ORD-123', date='2026-03-10')`, rejected = agent wrongly calls `get_order(order_id='March 10', date='ORD-123')` — confused args. After DPO training on 500-1000 such pairs, the model learns the pattern of correct tool parameter usage. I'd validate using my golden test set and measure the tool error rate before and after."*

---

## ROUND 3: PRACTICAL CODING (Both — Expect Live Coding)

---

**Q7. Write a function that detects if an agent is stuck in a loop.**

```python
from collections import Counter
from typing import List, Tuple

def detect_agent_loop(trajectory: List[Tuple[str, dict]]) -> dict:
    """
    Detects if an agent is stuck in a loop by checking for repeated (action, args) pairs.

    Args:
        trajectory: List of (action_name, arguments) tuples from agent history
    
    Returns:
        dict with 'is_looping', 'repeated_action', 'count'

    Example trajectory:
        [("search_db", {"query": "user_123"}),
         ("search_db", {"query": "user_123"}),   ← same action repeated
         ("search_db", {"query": "user_123"})]
    """
    # Serialize each action+args to a hashable string
    action_strings = []
    for action_name, args in trajectory:
        # Sort args dict to ensure consistent string representation
        sorted_args = dict(sorted(args.items()))
        action_str = f"{action_name}::{sorted_args}"
        action_strings.append(action_str)

    # Count occurrences of each action
    counts = Counter(action_strings)

    # Find any action repeated 2+ times
    for action_str, count in counts.items():
        if count >= 2:
            action_name = action_str.split("::")[0]
            return {
                "is_looping": True,
                "repeated_action": action_name,
                "count": count,
                "message": f"Agent repeated '{action_name}' {count} times — possible loop"
            }

    return {"is_looping": False, "repeated_action": None, "count": 0, "message": "No loop detected"}


# Test
trajectory = [
    ("get_order", {"order_id": "ORD-123"}),
    ("search_web", {"query": "refund policy"}),
    ("get_order", {"order_id": "ORD-123"}),   # ← repeated
    ("get_order", {"order_id": "ORD-123"}),   # ← repeated again
]

result = detect_agent_loop(trajectory)
print(result)
# {'is_looping': True, 'repeated_action': 'get_order', 'count': 3, ...}
```

---

**Q8. Write code to implement a basic token/cost tracker for agent calls.**

```python
from dataclasses import dataclass, field
from typing import Optional

# OpenAI pricing per 1 million tokens (as of early 2026)
PRICING = {
    "gpt-4o":           {"input": 2.50,  "output": 10.00},
    "gpt-4o-mini":      {"input": 0.15,  "output": 0.60},
    "gpt-3.5-turbo":    {"input": 0.50,  "output": 1.50},
    "claude-3-5-sonnet":{"input": 3.00,  "output": 15.00},
}

@dataclass
class AgentCostTracker:
    model: str
    user_id: str
    session_id: str
    input_tokens: int = 0
    output_tokens: int = 0
    tool_calls: int = 0
    steps: int = 0
    
    def record_llm_call(self, input_tokens: int, output_tokens: int):
        """Call this after every LLM interaction."""
        self.input_tokens += input_tokens
        self.output_tokens += output_tokens
        self.steps += 1
    
    def record_tool_call(self):
        """Call this after every tool execution."""
        self.tool_calls += 1
    
    @property
    def total_cost_usd(self) -> float:
        if self.model not in PRICING:
            return 0.0
        prices = PRICING[self.model]
        return (
            self.input_tokens  * prices["input"]  / 1_000_000 +
            self.output_tokens * prices["output"] / 1_000_000
        )
    
    def total_cost_inr(self, usd_to_inr: float = 83.5) -> float:
        return self.total_cost_usd * usd_to_inr
    
    def is_over_budget(self, budget_usd: float) -> bool:
        return self.total_cost_usd > budget_usd
    
    def summary(self) -> dict:
        return {
            "user_id": self.user_id,
            "model": self.model,
            "steps": self.steps,
            "tool_calls": self.tool_calls,
            "total_tokens": self.input_tokens + self.output_tokens,
            "cost_usd": round(self.total_cost_usd, 4),
            "cost_inr": round(self.total_cost_inr(), 2)
        }


# Usage in an agent wrapper
def run_tracked_agent(user_id: str, task: str, budget_usd: float = 0.10):
    tracker = AgentCostTracker(
        model="gpt-4o-mini",
        user_id=user_id,
        session_id=f"sess-{user_id}-{int(time.time())}"
    )
    
    for step in range(10):  # max 10 steps
        # After each LLM call:
        response = llm.invoke(task)
        tracker.record_llm_call(
            input_tokens=response.usage.prompt_tokens,
            output_tokens=response.usage.completion_tokens
        )
        
        if tracker.is_over_budget(budget_usd):
            print(f"⚠️ Budget exceeded: {tracker.summary()}")
            return "Task stopped — budget limit reached."
        
        if task_complete(response):
            break
    
    print(f"✅ Completed: {tracker.summary()}")
    save_to_db(tracker.summary())  # For billing and analytics
    return response.content
```

---

## ROUND 4: SYSTEM DESIGN (30-45 min, Product-Based)

**Q9. Design a system where 1 million developers use an AI coding agent daily.**

**Key Dimensions to Cover (examiner looks for these):**
```
1. Scale math:
   1M users × 10 tasks/day × avg 5 LLM calls/task = 50M LLM calls/day
   At $0.002/call (gpt-4o-mini) = $100,000/day → need caching!
   With 40% semantic cache hit rate → $60,000/day (still a lot!)
   
2. Architecture:
   - API Gateway (Nginx/Kong) → rate limiting per user
   - Kafka → decouple request intake from processing
   - Agent microservices (autoscaled Kubernetes pods)
   - Redis cluster → semantic cache (this is your biggest cost lever)
   - Vector DB (Pinecone/Weaviate) → code context retrieval
   - PostgreSQL → audit logs, task history, user preferences

3. Latency strategy:
   - Streaming tokens: user sees output in <100ms
   - Parallel tool calls: search + lint simultaneously
   - Tiered models: GPT-4o-mini for 80% of tasks, GPT-4o for complex ones

4. Reliability:
   - Circuit breakers on OpenAI API calls
   - Fallback to self-hosted Llama 3-70B if OpenAI is down
   - Agent checkpointing: resume incomplete tasks on failure
   
5. Security:
   - Code execution in isolated Docker containers (no network access)
   - Sensitive data (API keys in code) detected and masked before LLM
   - Each developer's code isolated — no cross-tenant data access

6. Monitoring:
   - Task success rate, P99 latency, cost per task, cache hit rate
   - Alert if any metric degrades by >10% in a 5-minute window
```

---

## QUICK-FIRE ROUND (Either Company)

| Question | Expected Answer |
|---|---|
| Name 3 LangChain agent types | ReAct, OpenAI Tools Agent, Structured Chat Agent |
| What is `max_iterations` for? | Prevents infinite agent loops |
| What tool prevents prompt injection? | Llama-Guard, NeMo Guardrails, or input classifier |
| What is semantic caching? | Cache LLM responses by semantic similarity, not exact match |
| Diff between SFT and RLHF? | SFT learns from examples; RLHF uses human preference rankings + RL |
| What is LangSmith? | LangChain's observability/tracing platform for debugging agents |
| What is `handle_parsing_errors=True`? | Makes agent retry instead of crash on malformed LLM output |
| What benchmark tests coding agents? | SWE-bench |
| What is LoRA? | Fine-tune only low-rank weight matrices (~1%), saves GPU memory |
| What is the ReAct loop? | Thought → Action → Observation → Thought → ... → Final Answer |
