# Production Tooling, Observability & Evaluation for Agentic AI (Product-Based Companies)

These questions test your knowledge of real production tools used to monitor, debug, evaluate, and optimize agentic AI systems. Common at Google, Microsoft, Stripe, Databricks, and AI-native startups.

---

## 1. What is LangSmith and How Do You Use It to Debug an Agent?

**Q: You have an agent in production that occasionally gives wrong answers. How do you use LangSmith to identify the root cause?**

**Answer:**

**LangSmith** is LangChain's observability platform that captures every step of an agent run — including LLM calls, tool invocations, inputs/outputs, and token usage.

### Setup

```python
import os
os.environ["LANGCHAIN_TRACING_V2"] = "true"
os.environ["LANGCHAIN_API_KEY"] = "ls_..."
os.environ["LANGCHAIN_PROJECT"] = "production-support-agent"

# All LangChain/LangGraph calls are now automatically traced — no code change needed
from langchain_openai import ChatOpenAI
llm = ChatOpenAI(model="gpt-4o")
```

### What LangSmith Captures Per Agent Run

```
Run ID: abc-123
├── [LLM] Planning call                    → 450ms, 1200 tokens
│     Input: "Break task X into steps"
│     Output: "Step 1: search DB, Step 2: filter, Step 3: format"
├── [Tool] search_database                 → 120ms
│     Args: {"query": "user ID 4829"}
│     Return: {"rows": 0}                  ← ⚠️ Empty result! Bug starts here.
├── [LLM] Post-tool reasoning              → 380ms, 800 tokens
│     Input: "[tool returned empty rows]"
│     Output: "User not found" → HALLUCINATED "No account exists"
└── Final: "No account exists for user 4829"  ← Wrong (user exists, query was wrong)
```

**Debugging Steps:**
1. Filter runs by `error=True` or low user rating score
2. Click into the failing run → inspect the exact tool arguments
3. Identify: the tool query was malformed (`"user ID 4829"` instead of `{"user_id": 4829}`)
4. Fix: add stronger Pydantic schema validation to the tool

### Adding Custom Metadata for Better Filtering

```python
from langsmith import traceable

@traceable(
    run_type="tool",
    name="database_search",
    metadata={"team": "support", "criticality": "high"}
)
def search_database(query: dict) -> list:
    ...
```

---

## 2. Langfuse vs. LangSmith — Which to Use and When?

**Q: Compare LangSmith and Langfuse. What are the tradeoffs?**

**Answer:**

| Feature | LangSmith | Langfuse |
|---|---|---|
| **Vendor** | LangChain (SaaS) | Open Source (self-host or SaaS) |
| **Framework Lock-in** | LangChain/LangGraph only | Framework-agnostic |
| **Self-hosting** | ❌ Not available | ✅ Full Docker/K8s support |
| **Cost** | Paid after free tier | Free (self-hosted) |
| **Evaluation** | Built-in LLM-as-judge | Built-in + custom evals |
| **Dataset Management** | Built-in | Built-in |
| **Best for** | LangChain-heavy stacks | Multi-framework or privacy-sensitive |

### Langfuse Integration (Framework-Agnostic)

```python
from langfuse import Langfuse
from langfuse.openai import openai  # Drop-in replacement for openai SDK

langfuse = Langfuse()

# Exact same API as openai.OpenAI() — all calls are automatically traced
client = openai.OpenAI()

response = client.chat.completions.create(
    model="gpt-4o",
    messages=[{"role": "user", "content": "What is 2+2?"}],
    # Extra Langfuse metadata:
    metadata={"user_id": "user-123", "session_id": "sess-456"}
)
print(response.choices[0].message.content)

# Manual span for custom code blocks
with langfuse.trace(name="custom-processing") as trace:
    trace.span(name="data-fetch", input={"id": 123})
```

---

## 3. How Do You Run SWE-bench Evaluations?

**Q: What is SWE-bench and how would you use it to evaluate a coding agent?**

**Answer:**

**SWE-bench** is a benchmark of 2,294 real GitHub issues from top Python repositories. An agent must read the issue, navigate the codebase, and produce a git patch that makes the failing tests pass.

### How It Works

```
Input:  GitHub issue description + full repository code
Output: A unified git patch (.diff file)

Evaluation: Run the existing test suite after applying the patch
Success:    All previously failing tests now pass (no regressions)
```

### Running a SWE-bench Evaluation

```bash
# Install
pip install swebench

# Run your agent on 50 instances from the verified subset
python -m swebench.harness.run_evaluation \
    --predictions_path /path/to/your_agent_predictions.jsonl \
    --swe_bench_tasks princeton-nlp/SWE-bench_Verified \
    --split test \
    --log_dir /path/to/logs
```

### Predictions File Format

```json
// your_agent_predictions.jsonl (one JSON per line)
{
  "instance_id": "django__django-11099",
  "model_patch": "diff --git a/django/db/models/sql/compiler.py...",
  "model_name_or_path": "my-coding-agent-v1"
}
```

### Interpreting Results

```
SWE-bench Results:
├── Resolved: 23.4%   (agent fixed the issue + tests pass)
├── Partial:  12.1%   (tests pass but regression introduced)
└── Failed:   64.5%   (patch doesn't compile or tests fail)

Benchmark: Claude 3.5 Sonnet = 49%, GPT-4o = 38%
```

**Key Takeaway:** SWE-bench is the gold-standard for coding agents. If someone asks "how do you know your coding agent works?" — SWE-bench is the answer.

---

## 4. OpenTelemetry for Agentic AI — How to Instrument Agents

**Q: How would you use OpenTelemetry (OTel) to trace a multi-agent system that spans multiple microservices?**

**Answer:**

OpenTelemetry is the industry standard for distributed tracing. For agents spread across microservices (Orchestrator → Coder Agent → Tester Agent), OTel ties all calls into a single trace.

### Instrumentation

```python
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter

# Setup — done once at service startup
provider = TracerProvider()
provider.add_span_processor(
    BatchSpanProcessor(OTLPSpanExporter(endpoint="http://jaeger:4317"))
)
trace.set_tracer_provider(provider)
tracer = trace.get_tracer("coding-agent-service")

# Instrument an agent function
def run_coder_agent(task: str, issue_id: str) -> str:
    with tracer.start_as_current_span("coder-agent-run") as span:
        span.set_attribute("issue.id", issue_id)
        span.set_attribute("task.description", task[:100])

        with tracer.start_as_current_span("llm-planning-call") as llm_span:
            plan = call_llm_for_plan(task)
            llm_span.set_attribute("llm.tokens_used", plan.usage.total_tokens)

        with tracer.start_as_current_span("code-execution") as exec_span:
            result = execute_code_in_sandbox(plan.code)
            exec_span.set_attribute("execution.exit_code", result.exit_code)
            if result.exit_code != 0:
                exec_span.set_status(trace.StatusCode.ERROR, result.stderr)

        return result.output
```

### What You See in Jaeger

```
Trace: issue-11099-fix (3.4s)
├── orchestrator: route-to-coder          200ms
├── coder-agent-service: coder-agent-run  2800ms
│    ├── llm-planning-call                 800ms  [tokens=1200]
│    ├── code-execution                   1200ms  [exit_code=0]
│    └── llm-review-call                   800ms  [tokens=800]
└── github-service: create-pr              400ms
```

**W3C Trace Context Propagation** — the `traceparent` header is automatically passed between microservices, so the entire agent run appears as ONE connected tree.

---

## 5. GAIA Benchmark — What Is It?

**Q: What is the GAIA benchmark and why is it harder than SWE-bench for general agents?**

**Answer:**

**GAIA** (General AI Assistants) is a benchmark testing real-world general task completion requiring multi-step reasoning, tool use, and finding information that can't be answered from memory alone.

### GAIA vs SWE-bench

| Property | SWE-bench | GAIA |
|---|---|---|
| **Domain** | Software engineering | General tasks (research, math, web) |
| **Tools Required** | Code execution, file system | Web search, calculator, file reading |
| **Difficulty** | High (49% SOTA) | Very High (Level 3: ~20% SOTA) |
| **Answer Format** | Git patch | Short exact answer |
| **Multi-hop** | Single codebase | Cross-domain reasoning |

### Example GAIA Task (Level 3)

```
Q: "As of my knowledge cutoff, what is the name of the person who 
    founded the company that makes the GPU used in the system that 
    achieved the highest score on this leaderboard?"

Steps Required:
1. Find GAIA leaderboard
2. Identify the top system
3. Find what GPU it uses
4. Find who manufactures that GPU
5. Find the GPU manufacturer's founder
Answer: "Jensen Huang"
```

### Why It Matters for Interviews

GAIA tests whether agents can autonomously chain 5–10 real-world research steps **without human guidance**. Mentioning this benchmark in a product interview shows you understand the frontier of agent evaluation.

---

## 6. Building an Automated Evaluation Dataset

**Q: How do you build and maintain a golden test set for continuous agent evaluation?**

**Answer:**

### Golden Dataset Construction Pipeline

```python
from langsmith import Client
import json

client = Client()

# Step 1: Capture production runs that users rated positively
good_runs = client.list_runs(
    project_name="production-agent",
    filter="feedback_score > 0.8",
    limit=500
)

# Step 2: Convert to dataset examples
dataset = client.create_dataset(
    dataset_name="Golden Agent Test Set v1",
    description="High-quality production runs for regression testing"
)

for run in good_runs:
    client.create_example(
        inputs=run.inputs,
        outputs=run.outputs,
        dataset_id=dataset.id,
        metadata={
            "source_run_id": str(run.id),
            "user_rating": run.feedback_score,
            "category": classify_task(run.inputs["input"])
        }
    )

print(f"Created dataset with {len(dataset)} examples")
```

### Continuous Evaluation in CI

```python
from langsmith.evaluation import evaluate

def run_agent_for_eval(inputs: dict) -> dict:
    result = agent_executor.invoke(inputs)
    return {"output": result["output"]}

def correctness_evaluator(run, example) -> dict:
    """LLM-as-a-judge evaluator"""
    expected = example.outputs["output"]
    actual = run.outputs["output"]
    score = grade_with_llm(expected, actual)
    return {"score": score, "key": "correctness"}

results = evaluate(
    run_agent_for_eval,
    data="Golden Agent Test Set v1",
    evaluators=[correctness_evaluator],
    experiment_prefix="v2.1-release",
    max_concurrency=5
)

# Compare against baseline
if results.summary_results["correctness"]["mean"] < 0.90:
    raise Exception("Agent quality below 90% threshold — blocking release")
```

### Dataset Maintenance

- **Quarterly audits**: Remove outdated examples, add new edge cases
- **Failure mining**: When production errors occur, add them to the dataset after fixing
- **Stratified sampling**: Ensure dataset covers all intent categories proportionally
- **Anti-hacking**: Rotate dataset quarterly so agents don't overfit to test cases

---

## 7. Prompt Versioning and Management

**Q: How do you manage prompt versions across 10 different agent types in production?**

**Answer:**

### Problems with Unmanaged Prompts
- Prompt changes are not tracked → can't roll back a bad change
- Different team members editing the same prompt → conflicts
- No A/B testing capability

### Solution: LangSmith Prompt Hub or Dedicated Prompt Registry

```python
from langchain import hub

# Pull a specific version of a prompt (pinned by commit hash)
prompt_v2 = hub.pull("myorg/support-agent-prompt:abc123")

# Pull latest (for development only — never in production)
prompt_latest = hub.pull("myorg/support-agent-prompt:latest")
```

### Custom Prompt Registry (if you self-host)

```python
# Store prompts in Postgres with version control
class PromptRegistry:
    def get_prompt(self, name: str, version: str = "latest") -> str:
        query = """
            SELECT content FROM prompts 
            WHERE name = %s AND version = %s
        """
        return db.execute(query, (name, version)).fetchone()

    def publish_prompt(self, name: str, content: str, author: str) -> str:
        """Creates a new version, returns version hash."""
        version = hashlib.sha256(content.encode()).hexdigest()[:8]
        db.execute("INSERT INTO prompts VALUES (%s, %s, %s, NOW())",
                   (name, version, content))
        return version

registry = PromptRegistry()
prompt_version = registry.publish_prompt(
    name="support-agent-system",
    content="You are a helpful customer support agent...",
    author="dhruv@company.com"
)
# Deploy with: PROMPT_VERSION=abc123 in your service config
```

### A/B Testing Prompts

```python
import random

def get_active_prompt(user_id: str) -> tuple[str, str]:
    """Returns (prompt_content, variant_name) for A/B testing."""
    # 50/50 split based on user ID hash
    variant = "control" if hash(user_id) % 2 == 0 else "treatment"

    if variant == "control":
        return registry.get_prompt("support-agent", version="v1.2"), "control"
    else:
        return registry.get_prompt("support-agent", version="v1.3-concise"), "treatment"

prompt, variant = get_active_prompt("user-abc")
# Log variant alongside each trace for statistical analysis
```
