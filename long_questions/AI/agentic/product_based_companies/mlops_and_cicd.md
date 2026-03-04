# MLOps & CI/CD for Agentic AI – Interview Q&A (Product-Based Companies)

These questions test your understanding of productionizing agent systems: versioning, deployment pipelines, model registry, data drift, and continuous improvement loops.

---

## 1. What is MLOps and How Does It Apply Specifically to Agentic AI Systems?

**Answer:**

MLOps (ML Operations) is the practice of applying DevOps principles to machine learning systems. For agentic AI, it requires **additional layers** beyond standard MLOps because agents are non-deterministic, use external tools, and change behavior based on both the model AND the prompt.

### Standard MLOps vs. Agentic MLOps

| Concern | Standard MLOps | Agentic MLOps |
|---|---|---|
| What to version | Model weights | Model weights + prompts + tools |
| What to test | `predict(X) == expected_Y` | Entire trajectory: was the task completed correctly? |
| What to monitor | Accuracy, data drift | Task success rate, tool error rate, cost per task |
| Rollback trigger | Performance drop on test set | Task success rate drop, hallucination spike |
| Feedback loop | Human labels on model outputs | User ratings + LLM-as-judge on full agent sessions |

### The Agentic MLOps Stack

```
┌────────────────────────────────────────────────────────────┐
│                    AGENTIC MLOps STACK                     │
├────────────────┬───────────────────────┬───────────────────┤
│   DEVELOPMENT  │     CI/CD PIPELINE    │   PRODUCTION      │
│                │                       │                   │
│ - LangChain/   │ - GitHub Actions      │ - Kubernetes      │
│   LangGraph    │ - Golden Set Eval     │ - FastAPI         │
│ - LangSmith    │ - Prompt Testing      │ - LangSmith       │
│   (tracing)    │ - Integration Tests   │   (monitoring)    │
│ - Langfuse     │ - Docker build        │ - Prometheus      │
│   (eval)       │ - Model registry      │ - Grafana         │
└────────────────┴───────────────────────┴───────────────────┘
```

---

## 2. How Do You Version an Agent? (Not Just the Model — Everything)

**Answer:**

An agent's behavior is determined by 4 components — all must be versioned:

```
Agent Version = Model Version + Prompt Version + Tool Version + Config Version
                    │                │               │              │
                  "gpt-4o-0513"   "v2.1.3"      "v1.4.0"    "max_iter=5"
```

### Implementation: Agent Version Manifest

```yaml
# agent_manifest.yaml — committed to Git, referenced by deployments
agent_name: "support-agent"
version: "2.3.1"
released_at: "2026-03-04"

model:
  provider: "openai"
  model_id: "gpt-4o-2024-08-06"   # Always pin to a specific snapshot, not "latest"
  temperature: 0
  max_tokens: 2048

prompts:
  system: "v4.2"                    # References entry in Prompt Registry
  tool_descriptions: "v1.8"

tools:
  - name: "get_order_status"
    version: "v3.1.0"               # Semver for each tool's API contract
  - name: "process_refund"
    version: "v2.0.1"

config:
  max_iterations: 7
  timeout_seconds: 30
  semantic_cache_enabled: true
```

### Why Pinning Model Version Matters

```python
# BAD — "gpt-4o-latest" can change overnight and break your agent
llm = ChatOpenAI(model="gpt-4o")

# GOOD — pinned to a specific release snapshot
llm = ChatOpenAI(model="gpt-4o-2024-08-06")

# When a new model version is available:
# 1. Create a new branch
# 2. Update manifest to "gpt-4o-2025-01-15"
# 3. Run full golden-set evaluation
# 4. If metrics improve or stay stable → merge & deploy
```

---

## 3. Design a CI/CD Pipeline for an AI Agent

**Question:** Walk me through how you would set up a CI/CD pipeline for an agent where every PR must pass evaluation before merging.

**Answer:**

### Pipeline Stages

```yaml
# .github/workflows/agent_cicd.yml

name: Agent CI/CD Pipeline

on:
  pull_request:
    branches: [main]
  push:
    branches: [main]

jobs:
  # ─── Stage 1: Unit Tests (fast, ~2 min) ───────────────────────────────
  unit_tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: pip install -r requirements.txt
      - run: pytest tests/unit/ -v --timeout=30
        # Tests individual tools, parsers, and utility functions in isolation

  # ─── Stage 2: Agent Evaluation (golden set, ~10 min) ──────────────────
  agent_evaluation:
    needs: unit_tests
    runs-on: ubuntu-latest
    steps:
      - run: python eval/run_golden_set.py \
               --agent-config agent_manifest.yaml \
               --dataset golden_test_set_v3.jsonl \
               --output eval_results.json

      - name: Check quality gate
        run: |
          python eval/quality_gate.py \
            --results eval_results.json \
            --min-success-rate 0.90 \        # Fail if < 90% task success
            --max-cost-per-task 0.05 \       # Fail if avg cost > $0.05
            --max-p99-latency 3000           # Fail if P99 > 3000ms

      - name: Upload results to LangSmith
        run: python eval/upload_to_langsmith.py --results eval_results.json

  # ─── Stage 3: Integration Tests (live API, ~5 min) ─────────────────────
  integration_tests:
    needs: agent_evaluation
    runs-on: ubuntu-latest
    steps:
      - run: pytest tests/integration/ -v \
               -k "not slow" \
               --env staging           # Tests against staging environment

  # ─── Stage 4: Docker Build & Push ──────────────────────────────────────
  build_and_push:
    needs: integration_tests
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - run: docker build -t agent-service:${{ github.sha }} .
      - run: docker push gcr.io/mycompany/agent-service:${{ github.sha }}

  # ─── Stage 5: Deploy to Staging (auto) / Production (manual approval) ──
  deploy_staging:
    needs: build_and_push
    runs-on: ubuntu-latest
    steps:
      - run: kubectl set image deployment/agent-service \
               agent=gcr.io/mycompany/agent-service:${{ github.sha }}
      - run: kubectl rollout status deployment/agent-service --timeout=5m

  deploy_production:
    needs: deploy_staging
    environment:
      name: production
      url: https://agent.mycompany.com
    # 🔑 Requires manual approval in GitHub before this runs
    steps:
      - run: ./scripts/blue_green_deploy.sh ${{ github.sha }}
```

---

## 4. How Do You Detect and Handle Data Drift in Production Agents?

**Answer:**

**Input drift** = the distribution of user queries changes. **Output drift** = the agent's response quality changes. Both can silently degrade agent performance.

### Types of Drift in Agents

| Drift Type | Example | Detection |
|---|---|---|
| **Input drift** | Users start asking about a new product feature the agent wasn't trained on | Monitor query topic distribution via embedding clustering |
| **Model drift** | LLM provider silently updates underlying model | Pin model versions, run nightly regression |
| **Tool drift** | External API your agent calls changes its response schema | Schema validation on every tool response |
| **Prompt drift** | A prompt that worked well starts returning off-format responses | Monitor output parse success rate |

### Detecting Input Drift Using Embeddings

```python
import numpy as np
from sklearn.decomposition import PCA
from langchain_openai import OpenAIEmbeddings

embeddings_model = OpenAIEmbeddings()

# Step 1: Establish baseline distribution (first 2 weeks of production)
baseline_queries = load_from_db(start="2026-01-01", end="2026-01-14", limit=5000)
baseline_embeddings = embeddings_model.embed_documents(baseline_queries)
baseline_mean = np.mean(baseline_embeddings, axis=0)

# Step 2: Daily drift check
def check_input_drift(recent_queries: list[str], threshold: float = 0.15) -> bool:
    recent_embeddings = embeddings_model.embed_documents(recent_queries)
    recent_mean = np.mean(recent_embeddings, axis=0)
    
    # Cosine distance between distributions
    cosine_sim = np.dot(baseline_mean, recent_mean) / (
        np.linalg.norm(baseline_mean) * np.linalg.norm(recent_mean)
    )
    drift_score = 1 - cosine_sim
    
    if drift_score > threshold:
        alert_team(f"⚠️ Input drift detected! Score: {drift_score:.3f}")
        return True
    return False

# Step 3: When drift is detected
# → Investigate clustered topics in the drifted embeddings
# → Update the agent's system prompt or RAG knowledge base
# → Generate new training examples for those topics
```

### Schema Validation for Tool Drift

```python
from pydantic import BaseModel, ValidationError

class OrderStatusResponse(BaseModel):
    order_id: str
    status: str
    estimated_delivery: str    # Validates this field will exist

def call_order_api(order_id: str) -> dict:
    raw_response = requests.get(f"https://orders-api/orders/{order_id}").json()
    
    try:
        validated = OrderStatusResponse(**raw_response)
        return validated.dict()
    except ValidationError as e:
        # API changed its response schema without notice!
        alert_team(f"🚨 Tool schema drift: {e}")
        raise ToolError("Order API response format changed unexpectedly")
```

---

## 5. How Do You Implement Blue-Green Deployment for an Agent?

**Answer:**

Blue-green deployment ensures zero-downtime agent updates by running two identical environments.

```
CURRENT STATE:
  Load Balancer → 100% traffic → Blue (v2.2 agent)

DEPLOYMENT STEP 1: Deploy new version alongside
  Blue (v2.2 agent)  ← 100% traffic (current)
  Green (v2.3 agent) ← 0% traffic (being prepared)

DEPLOYMENT STEP 2: Run shadow testing (5% canary)
  Blue (v2.2 agent)  ← 95% traffic
  Green (v2.3 agent) ← 5% traffic
  → Monitor: task success rate, latency, error rate for 30 min

DEPLOYMENT STEP 3: Gradual traffic shift (if metrics are good)
  Blue (v2.2 agent)  ← 50% → 10% → 0%
  Green (v2.3 agent) ← 50% → 90% → 100%

ROLLBACK (if metrics degrade at any step):
  Load Balancer → 100% traffic → Blue (instantly, <1 min)
```

### Kubernetes Implementation

```yaml
# blue-green-deploy.yaml
apiVersion: v1
kind: Service
metadata:
  name: agent-service
spec:
  selector:
    app: agent
    slot: blue   # Change to "green" to fully switch over
  ports:
    - port: 80

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: agent-green    # New version
spec:
  replicas: 3
  selector:
    matchLabels:
      app: agent
      slot: green
  template:
    metadata:
      labels:
        app: agent
        slot: green
    spec:
      containers:
        - name: agent
          image: gcr.io/mycompany/agent-service:v2.3.1
          env:
            - name: AGENT_VERSION
              value: "2.3.1"
```

```bash
# Switch traffic (zero downtime)
kubectl patch service agent-service \
  -p '{"spec":{"selector":{"slot":"green"}}}'

# Instant rollback if something goes wrong
kubectl patch service agent-service \
  -p '{"spec":{"selector":{"slot":"blue"}}}'
```

---

## 6. How Do You Track Costs and Set Budgets for Agents in Production?

**Answer:**

Uncontrolled agent costs can become catastrophic. A single runaway loop calling GPT-4o 100 times costs ~$15 — at scale, this destroys budget.

### Cost Tracking Per Agent Call

```python
import tiktoken
from dataclasses import dataclass
from datetime import datetime

@dataclass
class AgentCostTracker:
    model: str
    input_tokens: int = 0
    output_tokens: int = 0
    
    # OpenAI pricing (per 1M tokens)
    PRICES = {
        "gpt-4o": {"input": 2.50, "output": 10.00},
        "gpt-4o-mini": {"input": 0.15, "output": 0.60},
        "gpt-3.5-turbo": {"input": 0.50, "output": 1.50},
    }
    
    def add_call(self, input_tokens: int, output_tokens: int):
        self.input_tokens += input_tokens
        self.output_tokens += output_tokens
    
    @property
    def total_cost_usd(self) -> float:
        prices = self.PRICES[self.model]
        return (
            self.input_tokens * prices["input"] / 1_000_000 +
            self.output_tokens * prices["output"] / 1_000_000
        )

# Usage in agent
tracker = AgentCostTracker(model="gpt-4o")

def budget_controlled_agent(user_id: str, task: str, budget_usd: float = 0.10):
    for step in range(max_iterations):
        response = llm.invoke(...)
        tracker.add_call(
            input_tokens=response.usage.prompt_tokens,
            output_tokens=response.usage.completion_tokens
        )
        
        # Hard budget cap — kills the agent before it overspends
        if tracker.total_cost_usd > budget_usd:
            log_cost_exceeded(user_id, tracker.total_cost_usd)
            return "I was unable to complete this task within the budget. Please break it into smaller parts."
        
        if task_complete:
            break
    
    # Log to cost dashboard
    save_cost_metric(user_id=user_id, cost=tracker.total_cost_usd, task_type=classify(task))
    return final_answer
```

### Budget Alerting

```yaml
# Prometheus alert: avg cost per task spike
- alert: AgentCostSpike
  expr: avg(agent_cost_per_task_dollars) > 0.15
  for: 5m
  labels:
    severity: warning
  annotations:
    summary: "Agent costs are 3x normal — check for loops or expensive tasks"
```
