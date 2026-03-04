# Client Estimation, Sizing & Scoping for AI Agent Projects (Service-Based Companies)

This is a uniquely **consulting-specific** skill that barely appears in product company interviews but is almost guaranteed at Accenture, TCS, Infosys, Wipro, and Capgemini. The interviewer wants to know: *Can you walk into a client meeting and intelligently discuss cost, timeline, and feasibility?*

---

## 1. The Discovery Framework — Questions to Ask a Client First

Before proposing any architecture, a good consultant asks these scoping questions:

```
BUSINESS
├── What is the exact problem to solve? (Don't assume)
├── What does success look like? (Metric: reduce support tickets by 30%?)
├── What is the monthly budget for this solution?
└── Who are the end users? (Internal employees vs. B2C customers)

DATA & KNOWLEDGE
├── Where does the agent's knowledge come from? (Docs, DBs, APIs)
├── How frequently does this data change? (Daily, weekly, monthly)
├── How many documents/records? (100 PDFs vs. 10M records changes everything)
└── Any data residency requirements? (Must stay in India?)

VOLUME
├── How many queries/day do you expect? (100/day vs. 100,000/day)
├── What are peak hours? (Sale events? Office hours only?)
└── What is the acceptable response time? (Real-time <2s vs. async <60s)

INTEGRATION
├── What systems must the agent connect to? (SAP, Salesforce, custom APIs)
├── Do you have existing APIs or do we build them too?
└── Authentication? (SSO, OAuth, API keys)

GOVERNANCE
├── What actions can the agent do autonomously vs. requiring human approval?
├── Any regulatory requirements? (SEBI, RBI, HIPAA, GDPR)
└── Who is responsible if the agent makes an error?
```

---

## 2. Effort Estimation Template

**Q: A client asks for an estimate to build a customer support agent. How do you estimate the effort?**

**Answer:**

Break the work into standard components and estimate each:

| Component | Simple (days) | Medium (days) | Complex (days) |
|---|---|---|---|
| **Core agent setup** (LangChain + LLM + basic tools) | 3 | 5 | 8 |
| **RAG / Knowledge Base** (ingest docs, vector DB) | 2 | 5 | 10 |
| **Custom tools** (per tool, with testing) | 1–2 each | 3–4 each | 5–7 each |
| **Channel integration** (WhatsApp/Teams/Slack) | 2 | 3 | 5 |
| **Human-in-the-loop workflow** | 2 | 4 | 7 |
| **Authentication & Security** (PII scrubbing, RBAC) | 1 | 3 | 5 |
| **Testing** (unit + integration + UAT) | 3 | 5 | 8 |
| **Monitoring setup** (LangSmith/Prometheus/dashboards) | 1 | 3 | 5 |
| **Deployment + CI/CD** | 1 | 2 | 4 |
| **Documentation + training** | 1 | 2 | 3 |

### Example: Customer Support Agent Estimate

**Client scenario:** E-commerce client, WhatsApp integration, 4 tools (order status, returns, FAQ RAG, escalation), 500 users/day.

```
Component                           Days
──────────────────────────────────────────
Core agent + LLM setup              3
FAQ RAG (500 company docs)          4
Tool: Order Status API              2
Tool: Returns Processing            3
Tool: Escalation to human           2
WhatsApp integration                3
PII scrubbing (customer data)       2
Unit + integration testing          4
Staging + prod deployment           2
Monitoring dashboard                2
Client UAT support                  3
Documentation                       2
──────────────────────────────────────────
Total (1 developer)                32 days

With 2 developers working in parallel: ~18 working days ≈ 3.5 weeks
Add 20% buffer for client feedback rounds: ~4.5 weeks
```

**Interview Talking Point:**
> *"My estimates always include a 15–20% buffer for client change requests — in my experience, this is always needed. I also split the delivery into a 2-phase approach: Phase 1 (2.5 weeks) delivers a working prototype with the 2 highest-impact tools for UAT; Phase 2 (2 weeks) adds remaining tools after client feedback. This gives clients visibility early and reduces rework risk."*

---

## 3. Cloud Cost Estimation

**Q: How do you estimate the monthly cloud cost for an AI agent for a client?**

**Answer:**

Break costs into three buckets: **LLM API costs**, **Infrastructure costs**, **Tooling costs**.

### LLM API Cost Calculator

```python
def estimate_monthly_llm_cost(
    queries_per_day: int,
    avg_input_tokens: int,
    avg_output_tokens: int,
    model: str = "gpt-4o-mini",
    cache_hit_rate: float = 0.3  # 30% cache hits = no LLM call
) -> dict:
    """
    Estimate monthly LLM API costs.
    
    Token pricing (per 1M tokens, as of early 2026):
    - gpt-4o:      $2.50 input, $10.00 output
    - gpt-4o-mini: $0.15 input,  $0.60 output
    - gpt-3.5:     $0.50 input,  $1.50 output
    """
    pricing = {
        "gpt-4o":      {"input": 2.50,  "output": 10.00},
        "gpt-4o-mini": {"input": 0.15,  "output": 0.60},
        "gpt-3.5":     {"input": 0.50,  "output": 1.50},
    }
    
    rates = pricing[model]
    effective_queries = queries_per_day * (1 - cache_hit_rate)  # Cache saves this %
    
    # Monthly totals
    monthly_queries = effective_queries * 30
    monthly_input_tokens  = monthly_queries * avg_input_tokens
    monthly_output_tokens = monthly_queries * avg_output_tokens
    
    input_cost  = (monthly_input_tokens  / 1_000_000) * rates["input"]
    output_cost = (monthly_output_tokens / 1_000_000) * rates["output"]
    total_usd   = input_cost + output_cost
    total_inr   = total_usd * 83  # Approximate USD → INR
    
    return {
        "model": model,
        "monthly_queries": int(monthly_queries),
        "monthly_input_tokens":  int(monthly_input_tokens),
        "monthly_output_tokens": int(monthly_output_tokens),
        "monthly_cost_usd": round(total_usd, 2),
        "monthly_cost_inr": round(total_inr, 0),
        "with_cache_savings_pct": f"{cache_hit_rate*100:.0f}%"
    }

# Example: Customer support agent with 500 queries/day
print(estimate_monthly_llm_cost(
    queries_per_day=500,
    avg_input_tokens=1500,   # System prompt + history + user message
    avg_output_tokens=300,   # Agent response
    model="gpt-4o-mini",
    cache_hit_rate=0.35      # 35% FAQ hits cached
))
# Output:
# monthly_queries: 9,750
# monthly_cost_usd: $6.27
# monthly_cost_inr: ₹520
```

### Full Infrastructure Cost Breakdown (Azure, 500 queries/day)

| Item | Service | Est. Monthly (INR) |
|---|---|---|
| LLM API (gpt-4o-mini, 500 q/day, 35% cache) | Azure OpenAI | ₹520 |
| Vector DB for RAG | Azure AI Search (Basic tier) | ₹1,700 |
| App hosting | Azure App Service B2 (2 vCPU, 3.5GB) | ₹3,400 |
| Redis cache | Azure Cache for Redis C0 | ₹700 |
| PostgreSQL (state + logs) | Azure Database PostgreSQL Flexible B1ms | ₹1,200 |
| Monitoring + logs | Azure Monitor (10GB/month) | ₹400 |
| WhatsApp API (Meta) | Meta Cloud API (1000 free/month, then $0.005/msg) | ₹0–₹2,000 |
| **Total** | | **₹7,920 – ₹9,920/month** |

**Always round up to ₹10,000–₹12,000 for client proposals** to account for spikes and support.

---

## 4. Pricing Models for Delivering to Clients

**Q: How do you structure your billing to the client for an AI agent project?**

**Answer:**

| Model | When to Use | Pros | Cons |
|---|---|---|---|
| **Fixed Price** | Well-defined scope, known requirements | Client loves certainty | Risk on us if scope creeps |
| **Time & Materials (T&M)** | Ambiguous requirements, iterative builds | Flexible, we're protected | Client has budget anxiety |
| **Retainer** | Ongoing support, monitoring, improvements | Recurring revenue | Scope can balloon |
| **Outcome-based** | Client pays per resolved ticket / saved call | Aligns incentives | Hard if agent fails |

### Recommending the Right Model

```
Well-defined scope + tight deadline    → Fixed Price
Client wants to iterate/experiment     → T&M
Maintenance + model upgrades post-go-live → Monthly Retainer
Very mature client, high volume, trust → Outcome-based
```

**Interview Answer:**
> *"For first-time AI projects with a client, I recommend T&M for the build phase and transition to a monthly retainer for support. Fixed price is risky with AI because model behavior can change unexpectedly and new use cases always emerge during UAT. Once we've done one project together and the client trusts us, we can discuss fixed price for subsequent modules."*

---

## 5. Red Flags to Identify in Client Requirements

Experienced consultants know when to push back. Here are problems to catch early:

| Client Says | The Problem | Your Response |
|---|---|---|
| *"Make the agent autonomous with no approvals"* | Legal/reputational risk for write actions | *"We can automate reads fully. For writes, we recommend HITL for the first 3 months, then relax as accuracy is proven."* |
| *"We don't have any documentation yet"* | RAG needs documents to exist | *"Without a knowledge base, RAG won't work. We'd need 3–4 weeks to create and digitize the content first."* |
| *"Can it handle 10,000 users on day 1?"* | Scale testing not built into original plan | *"Load testing for that scale requires dedicated infra work. Let's add 5 days for load testing and autoscaling configuration."* |
| *"We need it in 1 week"* | Unrealistic timeline | *"In 1 week we can build a demo prototype. Production-grade with proper testing is minimum 4–5 weeks."* |
| *"Our customer data is on 20 different systems"* | Integration complexity explosion | *"Let's phase: start with the 2–3 systems that cover 80% of queries. Add others in Phase 2."* |

---

## 6. Quick 60-Second Estimation Formula (For Whiteboard Interviews)

When an interviewer throws a vague scenario at you, use this:

```
Step 1: Clarify scope in 3 questions
  → "How many users/queries per day?"
  → "How many tools need to be built?"
  → "Any compliance requirements?"

Step 2: Apply the formula
  Effort = (Base: 3-5 days) 
         + (Tools: 2 days × N tools)
         + (RAG: 3-5 days if docs needed)
         + (Integration: 2-3 days per channel)
         + (Testing: 25% of above)
         + (Buffer: 20%)

Step 3: Propose phased delivery
  → Phase 1 (MVP): 2 most valuable tools → demo in 2 weeks
  → Phase 2 (Full): remaining features → 2 more weeks
  → Phase 3 (Ongoing): monitoring, improvements → monthly retainer

Step 4: State the cost range
  → Small (< 1000 q/day, 3-4 tools): ₹8,000–15,000/month infra
  → Medium (1000–10,000 q/day): ₹20,000–60,000/month infra
  → Large (10,000+ q/day): ₹1L+/month, custom architecture needed
```
