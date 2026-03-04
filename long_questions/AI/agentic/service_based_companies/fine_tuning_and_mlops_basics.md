# Fine-Tuning & MLOps Basics for Agentic AI – Interview Q&A (Service-Based Companies)

These questions cover fine-tuning concepts and MLOps practices commonly asked at service-based company interviews (TCS, Infosys, Wipro, Accenture, Capgemini). Focus is on understanding concepts and knowing when to apply them for client projects.

---

## 1. What is Fine-Tuning and When Would You Recommend It to a Client?

**Answer:**

**Fine-tuning** is the process of taking a pre-trained LLM (like GPT-3.5 or Llama) and training it further on domain-specific data so it becomes an expert in that area.

### When to Recommend Fine-Tuning vs. Prompt Engineering

| Situation | Approach |
|---|---|
| Client needs general Q&A from documents | Use **RAG** (cheaper, faster to implement) |
| Client wants consistent output format (e.g., always respond as JSON) | Use **Fine-Tuning** |
| Client's domain has unique vocabulary (medical, legal, finance) | Use **Fine-Tuning** |
| Client is making thousands of API calls and wants to reduce costs | Use **Fine-Tuning** (smaller fine-tuned model = cheaper) |
| Client needs a quick proof of concept | Use **Prompt Engineering** first |

### Simple Analogy
```
Pre-trained model = A fresh MBA graduate — knows a lot broadly
Fine-tuned model  = The same graduate after 6 months at a specific company
                   — knows the company's terminology, style, and processes
```

### OpenAI Fine-Tuning Workflow

```python
import openai
import json

client = openai.OpenAI()

# Step 1: Prepare training data in JSONL format
training_examples = [
    {
        "messages": [
            {"role": "system", "content": "You are a customer support agent for StyleKart fashion store."},
            {"role": "user", "content": "I want to return a shirt I bought."},
            {"role": "assistant", "content": "I'd be happy to help with your return! StyleKart's return policy allows returns within 30 days with original tags. Please share your order number and I'll initiate the process."}
        ]
    },
    # ... minimum 10 examples (recommended: 50-100+)
]

# Save to JSONL file
with open("training_data.jsonl", "w") as f:
    for example in training_examples:
        f.write(json.dumps(example) + "\n")

# Step 2: Upload file to OpenAI
with open("training_data.jsonl", "rb") as f:
    file = client.files.create(file=f, purpose="fine-tune")

# Step 3: Start fine-tuning job
job = client.fine_tuning.jobs.create(
    training_file=file.id,
    model="gpt-3.5-turbo",   # Base model to fine-tune
    suffix="stylekart-support"  # Your model will be named: gpt-3.5-turbo-stylekart-support
)

print(f"Fine-tuning job started: {job.id}")
# Fine-tuning takes 10-60 minutes depending on dataset size

# Step 4: Use the fine-tuned model
response = client.chat.completions.create(
    model=job.fine_tuned_model,   # Use your fine-tuned model
    messages=[{"role": "user", "content": "Can I exchange an item?"}]
)
print(response.choices[0].message.content)
```

---

## 2. What is LoRA and Why Is It Important?

**Answer:**

**LoRA (Low-Rank Adaptation)** is a popular technique for fine-tuning large models efficiently. Instead of changing all model weights (billions of parameters), LoRA adds small trainable matrices and only trains those.

### Why LoRA Matters

```
Full Fine-Tuning (Llama 3-8B):
  → Update all 8,000,000,000 parameters
  → Requires 40+ GB GPU memory
  → Takes hours, costs $200+

LoRA Fine-Tuning (same model):
  → Only update ~0.1% of parameters (8,000,000)
  → Requires 8-12 GB GPU memory (single GPU!)
  → Takes 30-60 minutes, costs $5-20
  
Result quality: nearly identical to full fine-tuning
```

### When to Mention LoRA in an Interview

*"For this client project, since we need domain adaptation without full retraining costs, I'd recommend LoRA fine-tuning on an open-source model like Llama 3. This can be done on a single A100 GPU in about an hour, which fits well within the client's budget and timeline."*

---

## 3. What is MLOps and Why Does a Client Project Need It?

**Answer:**

**MLOps** is the set of practices that ensure AI/ML systems are deployed, monitored, and maintained reliably in production — similar to how DevOps ensures software quality.

### Without MLOps (Common in Client Projects)

```
🚨 Problems that come up without MLOps:
- Agent worked in dev, broken in production (environment mismatch)
- New team member changes a prompt → agent breaks for all users
- Agent costs $500/day instead of expected $50/day (nobody was tracking)
- Model was updated by the provider → agent behavior changed silently
- Bug was introduced → nobody knows which change caused it
```

### With MLOps

```
✅ Benefits:
- Every change tracked (prompts, model versions, tool versions)
- Automated tests run before every deployment
- Real-time monitoring catches issues within minutes
- Rollback to previous version in < 5 minutes
- Clear audit trail (important for regulated industries: BFSI, Healthcare)
```

### Core MLOps Components for Client Projects

```
1. Version Control     → Git for code + prompts
2. Experiment Tracking → MLflow or LangSmith for tracking runs
3. CI/CD Pipeline      → GitHub Actions for automated testing
4. Monitoring          → Grafana + Prometheus for production metrics
5. Model Registry      → Track which model version is deployed where
```

---

## 4. How Do You Test an AI Agent Before Deploying to Client Production?

**Answer:**

### Three Levels of Testing

**Level 1: Unit Testing (Fast, runs in seconds)**
```python
import pytest
from unittest.mock import patch

# Test individual tools work correctly WITHOUT calling the LLM
def test_order_lookup_tool():
    result = get_order_status("ORD-123")
    assert "status" in result
    assert result["status"] in ["Delivered", "Shipped", "Processing", "Cancelled"]

def test_order_not_found():
    result = get_order_status("INVALID-999")
    assert "not found" in result.lower()
```

**Level 2: Integration Testing (Medium, calls real LLM)**
```python
# Test the full agent with real LLM but controlled inputs
def test_agent_handles_order_query():
    result = agent_executor.invoke({
        "input": "What is the status of order ORD-555?"
    })
    # Verify agent called the right tool and gave meaningful answer
    assert "ORD-555" in result["output"]
    assert any(keyword in result["output"].lower() 
               for keyword in ["delivered", "shipped", "processing"])

def test_agent_escalates_complex_complaints():
    result = agent_executor.invoke({
        "input": "Your product damaged my property and I'm going to sue you!"
    })
    # Verify agent escalates instead of trying to handle itself
    assert "escalat" in result["output"].lower() or "manager" in result["output"].lower()
```

**Level 3: Golden Set Evaluation (Comprehensive quality check)**
```python
# Run the agent on 50-100 pre-validated question-answer pairs
golden_set = [
    {"input": "What is your return policy?", "expected_keywords": ["30 days", "original tags"]},
    {"input": "Track order ORD-777", "expected_keywords": ["shipped", "delivery"]},
    # ... 50+ examples
]

def run_golden_evaluation():
    results = []
    for example in golden_set:
        agent_answer = agent_executor.invoke({"input": example["input"]})["output"]
        passed = all(kw.lower() in agent_answer.lower() 
                     for kw in example["expected_keywords"])
        results.append(passed)
    
    success_rate = sum(results) / len(results)
    print(f"Golden Set Pass Rate: {success_rate:.1%}")
    
    # For client demos: anything below 85% needs improvement before deployment
    assert success_rate >= 0.85, f"Agent quality too low: {success_rate:.1%}"
```

### Pre-Launch Checklist for Client Deployments

```markdown
□ Unit tests passing (100%)
□ Integration tests passing (100%)
□ Golden set evaluation > 85% pass rate
□ Load test: agent handles 50 concurrent users without errors
□ Security review: no PII logged, authentication in place
□ Cost estimate reviewed: within client's budget
□ Fallback tested: agent gracefully handles API downtime
□ Escalation path tested: agent hands off to humans correctly
```

---

## 5. How Do You Monitor an Agent After Going Live for a Client?

**Answer:**

### Key Metrics to Show on Client Dashboard

```
HEALTH DASHBOARD (Grafana)
┌────────────────────┬────────────────────┬────────────────────┐
│ Task Success Rate  │ Avg Response Time  │ Daily Cost         │
│   94.2% ✅         │    1.8 seconds ✅  │   ₹4,200 ✅        │
│ (target: > 90%)   │ (target: < 3s)    │ (budget: ₹5,000)   │
└────────────────────┴────────────────────┴────────────────────┘
┌────────────────────┬────────────────────┬────────────────────┐
│ Escalations        │ User Rating        │ Error Rate         │
│   8.3%             │   4.2 / 5.0 ⭐    │   1.1% ✅          │
│ (target: < 15%)   │ (target: > 4.0)   │ (target: < 5%)     │
└────────────────────┴────────────────────┴────────────────────┘
```

### Simple Monitoring Setup

```python
from prometheus_client import Counter, Histogram, Gauge
import time

# Define metrics
task_counter = Counter("agent_tasks_total", "Total agent tasks", ["status"])
latency_histogram = Histogram("agent_latency_seconds", "Agent response time")
cost_gauge = Gauge("agent_daily_cost_rupees", "Daily agent cost in rupees")

def monitored_agent_call(user_message: str) -> str:
    start_time = time.time()
    
    try:
        result = agent_executor.invoke({"input": user_message})
        task_counter.labels(status="success").inc()
        return result["output"]
    
    except Exception as e:
        task_counter.labels(status="error").inc()
        return "Sorry, I'm unable to process your request right now."
    
    finally:
        duration = time.time() - start_time
        latency_histogram.observe(duration)
```

### What to Do When Metrics Drop

| Situation | Immediate Action |
|---|---|
| Success rate drops below 85% | Alert team → investigate LangSmith traces → identify failing query type |
| Response time > 5s | Check if tool APIs are slow → add caching layer |
| Cost doubles overnight | Check for agent loops → review max_iterations limit |
| User rating drops below 3.5 | Read recent conversation logs → identify complaint pattern |

---

## 6. What is the Difference Between Staging and Production for Agent Deployments?

**Answer:**

```
                  DEVELOPMENT          STAGING          PRODUCTION
                      │                   │                  │
Model:           gpt-3.5-turbo      gpt-4o-mini          gpt-4o
                 (cheapest)          (same as prod)      (best quality)

Data:            Fake/test data      Copy of real data    Real user data

Users:           Developers only     QA + Client UAT      All end users

Testing:         Automated tests     Manual UAT           Canary releases

Monitoring:      Basic logs          Full monitoring      Full monitoring
                                                          + Alerts

Cost:            Very low            Medium               Full budget
```

### Key Practices

1. **Never test with prod data in staging** — use anonymized copies for GDPR/data privacy
2. **Match staging to prod model** — using a cheaper model in staging may hide issues
3. **Client UAT on staging** — client acceptance testing happens here, not in production
4. **Automated deployment** — staging should be automatically deployable from CI/CD; production requires manual approval

### A Note for Service-Based Company Interviews

When asked about deployment, always mention:
- *"I'd recommend a staging environment that mirrors production so the client can validate before go-live"*
- *"We should have a rollback plan — if the agent's success rate drops by X% in the first hour, we revert automatically"*
- *"Client sign-off (UAT) should be completed on staging before any production deployment"*
