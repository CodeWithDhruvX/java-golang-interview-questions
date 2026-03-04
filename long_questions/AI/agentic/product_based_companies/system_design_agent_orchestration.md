# System Design for Agentic AI Orchestration at Scale (Product-Based Companies)

These questions are asked at Google, Microsoft, Meta, Stripe, and similar product companies. They test your ability to design robust, scalable, and production-grade agentic systems — not just prototype code.

---

## 1. Design an AI Coding Agent (like Devin or GitHub Copilot Workspace)

**Question:** Design an autonomous coding agent that accepts a GitHub issue, navigates the codebase, writes a fix, runs tests, and submits a PR. Handle reliability and scale.

**Answer:**

### High-Level Architecture

```
User / GitHub Issue
        │
        ▼
┌──────────────────┐
│  Orchestrator    │  ← LangGraph / Temporal Workflow
│  (Planner LLM)  │
└──────┬───────────┘
       │ Delegates to
  ┌────┴────┐
  │ Agents  │
  ├─────────┤
  │ Coder   │ → Reads/writes files in sandboxed Docker container
  │ Tester  │ → Runs pytest/jest/go test inside sandbox
  │ Reviewer│ → Reads git diff and checks for regressions
  └────┬────┘
       │
  ┌────▼─────────────┐
  │ Persistent Layer │
  │  - Redis (state) │
  │  - Postgres      │
  │    (task logs)   │
  │  - S3 (artifacts)│
  └──────────────────┘
       │
  ┌────▼────────────┐
  │ GitHub API Tool │ → Creates branch, commits, opens PR
  └─────────────────┘
```

### Component Breakdown

| Component | Responsibility | Tech |
|---|---|---|
| **Orchestrator** | Reads issue, creates a high-level plan | GPT-4o (expensive, used once) |
| **Coder Agent** | Implements specific file changes | GPT-4o or Claude 3.5 |
| **Tester Agent** | Runs test suite, parses results | GPT-4o-mini + shell tool |
| **Reviewer Agent** | Validates the diff for style, regressions | GPT-4o |
| **Sandboxed Env** | Isolated Docker container per task | Docker-in-Docker |
| **State Store** | Checkpoints the workflow at every step | Redis + LangGraph MemorySaver |

### Key Design Decisions

**1. Sandboxing for Security**
```
Every code execution tool runs inside:
  - A Docker container with no internet access
  - A read-only mount of the repo + a writable /tmp workspace
  - CPU/Memory limits (e.g., 2 CPU, 4GB RAM)
  - A 5-minute execution timeout
```

**2. Hierarchical LLM Strategy (Cost Control)**
```
Planner   → GPT-4o        ($15/1M tokens) — Only 1 call per issue
Coder     → Claude 3.5    ($3/1M tokens)  — Multiple calls
Tester    → GPT-4o-mini   ($0.15/1M tokens) — Many small calls
Reviewer  → GPT-4o-mini   ($0.15/1M tokens) — Diff analysis
```

**3. Failure Recovery**
- **LangGraph Checkpoints** save state after each node → resume from last successful step
- If the Tester returns compile errors, the Coder agent is called again with the error as context (max 3 retries)
- After 3 failures, escalate to a human via PagerDuty

**4. Context Window Management**
- Large codebases don't fit in context → Use **AST-based file selection** (TreeSitter) to extract only relevant files/functions
- Maintain a `recently_modified_files` list in state to keep the agent focused

**Scale Consideration:** 1000 concurrent coding sessions → each runs in its own Kubernetes pod with isolated state. Orchestrator is a horizontally scalable FastAPI service.

---

## 2. Design a Multi-Agent Customer Support System at 1M QPS

**Question:** Design a multi-agent customer support system for a company like Flipkart or Amazon that handles 1 million queries per second during a sale event.

**Answer:**

### Architecture Diagram

```
                    Load Balancer
                         │
              ┌──────────┴──────────┐
              ▼                     ▼
        ┌──────────┐         ┌──────────┐
        │ Kafka    │         │ Kafka    │  ← Event-driven fanout
        │ Topic:   │         │ Topic:   │
        │ queries  │         │ priority │
        └────┬─────┘         └────┬─────┘
             │                   │
    ┌────────▼────────────────────▼────────┐
    │           ROUTER AGENT               │  ← Cheap fast model (GPT-4o-mini)
    │   Classifies: FAQ / Order / Billing  │
    └────────┬────────────────────────────┘
             │
    ┌────────┴─────────────────────────────┐
    │   Specialized Agent Pool             │
    ├──────────────────────────────────────┤
    │  FAQ Agent   → Redis-cached answers  │
    │  Order Agent → Order DB + API tools  │
    │  Billing Agent → Finance system      │
    │  Escalation Agent → HITL queue       │
    └──────────────────────────────────────┘
             │
    ┌────────▼────────────────────────────┐
    │     Response Cache (RedisVL)        │  ← Semantic cache, 60% hit rate
    └─────────────────────────────────────┘
```

### Scaling Strategy

**1. Router Agent (Cheap + Fast)**
- Use `gpt-4o-mini` or a locally hosted `Llama 3-8B` for routing
- Classifies intent in <100ms using a fine-tuned classifier
- Routes to specialized queues (Kafka topics per intent)

**2. Semantic Caching**
```
FAQ queries: ~60% are semantically identical
→ Use RedisVL with OpenAI embeddings
→ Cache TTL: 1 hour for FAQs, 10 seconds for order status
→ Saves ~60% of LLM costs on peak days
```

**3. Specialized Agent Pools**
```
Each agent type runs in its own autoscaling group:
  - FAQ Agent: 100 replicas, stateless, Redis-backed
  - Order Agent: 50 replicas, reads order DB
  - Billing Agent: 10 replicas, restricted access
→ Scale independently based on traffic patterns
```

**4. Circuit Breakers**
```
If Order DB is slow (>500ms):
  → Circuit breaker trips
  → Agent responds: "Order status temporarily unavailable"
  → Fallback to cached last-known status
```

**5. Rate Limiting**
```
Per-user rate limit: 10 queries/min
Implemented at the Load Balancer + Redis (token bucket algorithm)
Prevents a single user from exhausting agent capacity
```

### Numbers
| Tier | Models | Latency | Cost/1K queries |
|---|---|---|---|
| Router | GPT-4o-mini | 80ms | $0.015 |
| FAQ Agent (cached) | Redis | 5ms | $0.001 |
| FAQ Agent (uncached) | GPT-4o-mini | 150ms | $0.15 |
| Order Agent | GPT-4o | 400ms | $1.50 |

---

## 3. Design an Agent Evaluation Pipeline

**Question:** How do you build a continuous evaluation pipeline for agents deployed in production to detect regressions?

**Answer:**

### Evaluation Architecture

```
─────────────────────────────────────────────────────────
                  OFFLINE EVALUATION
─────────────────────────────────────────────────────────
Golden Test Set (100–1000 labeled trajectories)
         │
         ▼
┌─────────────────┐    ┌───────────────────┐
│  Agent Under    │───▶│  Evaluator LLM    │
│  Test (AUT)     │    │  (LLM-as-a-Judge) │
└─────────────────┘    └────────┬──────────┘
                                │
                         ┌──────▼──────────┐
                         │  Metrics Store  │
                         │  (Postgres)     │
                         └─────────────────┘
                                │
                         ┌──────▼──────────┐
                         │  Dashboard      │
                         │  (Grafana)      │
                         └─────────────────┘

─────────────────────────────────────────────────────────
                  ONLINE EVALUATION  
─────────────────────────────────────────────────────────
Production Traffic
         │ (10% shadow traffic)
         ▼
[Shadow Agent] → LangSmith traces → Anomaly Detector
                                          │
                                    Alert if metrics drop
```

### Key Metrics to Track

| Metric | Definition | Target |
|---|---|---|
| **Task Success Rate** | % of tasks completed correctly | > 90% |
| **Trajectory Efficiency** | Avg steps taken vs. optimal | < 1.5x optimal |
| **Tool Error Rate** | % of tool calls that fail/throw | < 5% |
| **Hallucination Rate** | % outputs with factual errors (graded by LLM-judge) | < 2% |
| **Latency P99** | 99th percentile end-to-end latency | < 3s |
| **Cost per Task** | Average LLM token cost per completed task | < $0.05 |

### Evaluator LLM (LLM-as-a-Judge)

```python
EVALUATOR_PROMPT = """
You are an expert evaluator. Score the following agent trajectory:

Task: {task}
Agent Steps: {trajectory}
Final Answer: {final_answer}
Expected Answer: {expected_answer}

Score on:
1. Correctness (0-10): Is the final answer correct?
2. Efficiency (0-10): Did the agent use the minimum necessary steps?
3. Safety (0-10): Did the agent avoid any unsafe actions?

Respond as JSON: {{"correctness": X, "efficiency": X, "safety": X, "reasoning": "..."}}
"""
```

### CI/CD Integration
```yaml
# .github/workflows/agent_eval.yml
on: [pull_request]
jobs:
  evaluate_agent:
    steps:
      - run: python eval/run_golden_set.py --agent-version ${{ github.sha }}
      - run: python eval/compare_baseline.py --threshold 0.95
      # Fails the PR if task success rate drops below 95% of baseline
```

---

## 4. Design a RAG + Agent Hybrid for Enterprise Document Q&A

**Question:** Design a system at a company where agents can answer complex, multi-hop questions across 10 million internal documents. Handle accuracy, cost, and latency.

**Answer:**

### System Architecture

```
User Question
     │
     ▼
┌─────────────────────┐
│   Query Planner     │  ← Decides: Simple RAG vs Agent Loop
│   (GPT-4o-mini)     │
└──────┬──────────────┘
       │
  ┌────┴────────────────┐
  │ Simple?             │ Complex?
  ▼                     ▼
RAG Pipeline      Agent Loop
  │                    │
  │               ┌────┴───────────────────┐
  │               │ Sub-question Generator │  → Breaks "multi-hop" into sub-queries
  │               └────┬───────────────────┘
  │                    │ (3-5 sub-questions)
  ▼                    ▼
┌──────────────────────────────┐
│   Vector Store (Milvus)       │
│   - 10M docs, 1536-dim       │
│   - HNSW index, ~10ms lookup │
└──────────────────────────────┘
       │
  Retrieved chunks
       │
  ┌────▼──────────────┐
  │  Re-ranker         │  ← Cross-encoder re-ranks top K=100 → top 5
  │  (Cohere Rerank)   │
  └────┬──────────────┘
       │
  ┌────▼──────────────┐
  │  Answer Synthesizer│  ← GPT-4o generates grounded answer with citations
  └────────────────────┘
```

### Multi-Hop Query Decomposition

**Query:** "What was the revenue growth of the mobile division in Q3, and how did it compare to the PC division's performance after the 2023 acquisition?"

**Decomposed into:**
1. "What was the mobile division revenue in Q3?"
2. "What was the mobile division revenue in Q2 (for growth calculation)?"
3. "When was the PC division acquired in 2023?"
4. "What was the PC division revenue in Q3?"

Each sub-question runs independently (parallel retrieval), then results are synthesized.

### Key Optimizations

**1. Sparse + Dense Hybrid Retrieval**
```
Dense: OpenAI embeddings (semantic similarity)
Sparse: BM25 (keyword match for specific terms like "Q3 2023")
Fusion: Reciprocal Rank Fusion (RRF)
→ 15-20% accuracy improvement over dense-only
```

**2. Query Routing**
```python
def route_query(question: str) -> str:
    # Single-hop: "What is the refund policy?"
    # → Simple RAG, 200ms, $0.002
    
    # Multi-hop: "Compare Q3 revenue across divisions post-acquisition"  
    # → Agent loop, 2s, $0.05
    
    complexity_score = classify_complexity(question)
    return "rag" if complexity_score < 0.5 else "agent"
```

**3. Citation Grounding**
- Every answer must include source document IDs
- Post-processing step validates cited chunks actually contain the claimed information
- If validation fails → re-generate with stricter prompt

---

## 5. Design Agent Observability & Monitoring Infrastructure

**Question:** How do you build an observability system for 50 production agents serving millions of users?

**Answer:**

### Three Pillars of Agent Observability

```
┌─────────────────────────────────────────┐
│           OBSERVABILITY STACK           │
├────────────────┬───────────────┬────────┤
│    TRACES      │    METRICS    │  LOGS  │
│  (LangSmith /  │  (Prometheus/ │ (ELK   │
│   Langfuse)    │   Grafana)    │ Stack) │
└────────────────┴───────────────┴────────┘
```

### Distributed Tracing for Agents

```python
from langfuse import Langfuse
from langfuse.decorators import observe

langfuse = Langfuse()

@observe()  # Auto-traces this function as a "span"
def run_agent_step(step_name: str, input_data: dict) -> dict:
    # Every LLM call, tool call, and step is automatically traced
    ...

@observe(name="full-agent-run")
def run_agent(task: str) -> str:
    # Creates a parent trace that contains all child spans
    plan = run_agent_step("planning", {"task": task})
    result = run_agent_step("execution", {"plan": plan})
    return result
```

**What Gets Captured Per Trace:**
- Agent input & output
- Each LLM call: model, tokens in/out, latency, cost
- Each tool call: tool name, args, return value, duration
- Full trajectory as a linked list of spans

### Alerting Rules

```yaml
# Prometheus alerting rules
groups:
  - name: agent_alerts
    rules:
      - alert: HighHallucinationRate
        expr: agent_hallucination_rate > 0.05
        for: 5m
        annotations:
          summary: "Agent hallucination rate exceeded 5%"

      - alert: AgentLoopDetected
        expr: agent_step_count > 20
        for: 0m
        annotations:
          summary: "Agent potentially stuck in loop"

      - alert: HighCostPerTask
        expr: agent_cost_per_task_dollars > 0.10
        for: 10m
        annotations:
          summary: "Agent cost exceeding budget threshold"
```

### Dashboard Metrics

| Dashboard Panel | Metric | Visualization |
|---|---|---|
| Success Rate | % tasks completed successfully | Time-series line |
| Avg Steps per Task | Mean LLM calls per task | Histogram |
| Cost Breakdown | Cost by model/agent type | Stacked bar |
| P99 Latency | 99th percentile latency | Heatmap |
| Error Types | Breakdown of failure reasons | Pie chart |
| Cache Hit Rate | % of calls served from cache | Gauge |
