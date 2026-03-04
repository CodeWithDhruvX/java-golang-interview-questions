# Gen-AI MLOps, Deployment & Evaluation — Product-Based Companies

These questions are asked by companies like **Databricks, Scale AI, Cohere, Mistral, Anyscale, and Hugging Face** where you're expected to own the full ML lifecycle — training, evaluation, serving, and continuous improvement.

---

## 1. MLOps for LLMs

**Q1: How do you build a CI/CD pipeline for an LLM application? What stages does it have?**
- **Stage 1 – Pre-PR checks:** Lint prompt templates, unit test tool schemas, static code analysis.
- **Stage 2 – Evaluation gate:** Run the change against an evaluation harness (RAGAS, Evals-based framework). Any regression in faithfulness/answer relevancy blocks the merge.
- **Stage 3 – Shadow deployment:** Route 5% of production traffic to the new model/prompt version. Compare LLM-as-judge scores vs. baseline.
- **Stage 4 – Canary rollout:** Gradual rollout from 5% → 20% → 100% with auto-rollback if error rate spikes.
- **Tooling:** LangSmith / Weights & Biases / Opik / Braintrust for eval tracking.

**Q2: How do you detect and handle LLM drift in production?**
- **What is LLM drift?** The model's output distribution shifts over time due to: updated API model versions (e.g., GPT-4 upgrades), input distribution shift, or prompt sensitivity to new content.
- **Detection:**
  - Track LLM-as-judge scores over a rolling 7-day window. Alert if score drops > 5%.
  - Monitor output length distribution, vocabulary entropy.
  - A/B compare with a pinned frozen model (reference model) daily.
- **Response:** Trigger an eval harness run; if regression confirmed, rollback or update prompts.

**Q3: What observability stack would you build for a production GenAI application?**
- **Tracing:** Log every LLM call with: prompt, completion, model version, latency, token counts, cost (use OpenTelemetry + LangSmith or Langfuse).
- **Metrics:** TTFT (time-to-first-token), ITL (inter-token latency), throughput (tokens/sec), p95/p99 latency, hallucination rate (sampled).
- **Alerts:** Spike in output refusals, error rate > threshold, cost per request exceeding budget.
- **Dashboards:** Grafana boards for GPU utilization, KV cache hit rate, continuous batching queue depth.

---

## 2. Evaluation Frameworks

**Q4: Describe the RAGAS evaluation framework. What does each metric measure?**

| Metric | Formula Concept | What it Catches |
|---|---|---|
| **Faithfulness** | Claims in answer supported by retrieved context? | Hallucination |
| **Answer Relevancy** | Does the answer actually address the question? | Off-topic answers |
| **Context Recall** | Were the relevant ground-truth facts retrieved? | Retrieval gaps |
| **Context Precision** | How much of retrieved context is actually useful? | Noise in retrieval |
| **Answer Correctness** | End-to-end: Is the answer factually correct? | Overall quality |

**Q5: How do you build a custom evaluation harness for a domain-specific LLM?**
- **Step 1:** Curate a golden dataset: 200+ (question, ground-truth answer, optional source) pairs.
- **Step 2:** Define binary/rubric metrics: exact match, contains-key-fact, safety classification.
- **Step 3:** Run LLM-as-judge with a carefully designed scoring rubric (G-Eval style, 1–5 scale) for open-ended questions.
- **Step 4:** Store eval results in a database keyed by model version + prompt hash for trend analysis.
- **Step 5:** Automate in CI/CD; gate model promotions on minimum score thresholds.

---

## 3. Cost Optimization

**Q6: How would you reduce inference costs for an LLM-powered product serving 1M requests/day?**
- **Prompt caching:** Use provider-side prompt caching (Anthropic, OpenAI support prefix caching). For stable system prompts, this yields 50–70% cache hit rates.
- **Semantic caching:** Embed the user query; if a semantically similar query was answered within the past 24h, return the cached answer (GPTCache, Redis + vector similarity).
- **Model tiering:** Route simple queries (greetings, factual lookups) to smaller/cheaper models (GPT-3.5, Haiku). Route complex reasoning to flagship models.
- **Output length control:** Enforce `max_tokens` limits; use structured output to avoid verbose preamble.
- **Batching:** For async workloads (nightly report generation), batch requests to maximize throughput per dollar.

**Q7: How do you choose between self-hosted open-source LLMs vs. closed-source API providers in a product context?**

| Factor | Open-Source (Llama, Mistral) | Closed-Source (GPT-4, Claude) |
|---|---|---|
| **Cost at scale** | Lower (fixed GPU cost) | High per-token pricing |
| **Data privacy** | Full control | Data leaves your infra |
| **Customization** | Fine-tunable | Prompt-only |
| **Capability frontier** | 1–2 generations behind | SOTA |
| **Ops burden** | High (infra team needed) | Zero |

- **Rule of thumb:** If > 10M tokens/day, self-hosting typically breaks even vs. API costs around 6–12 months.

---

## 4. Prompt Engineering at Scale

**Q8: What is chain-of-thought (CoT) prompting? What are its variants and when do you use each?**
- **Standard CoT:** Add "Let's think step by step" to elicit multi-step reasoning. Improves accuracy on arithmetic, logic, and multi-step tasks.
- **Zero-shot CoT:** Just append the instruction; no examples needed (works for GPT-4 class models).
- **Few-shot CoT:** Prepend 3–5 worked examples with reasoning chains. More reliable for smaller models.
- **Self-consistency:** Sample N CoT completions, take the majority answer. Reduces variance significantly.
- **Tree of Thoughts (ToT):** Explore multiple reasoning branches in parallel and evaluate paths. For complex planning.

**Q9: What are "structured outputs" and why are they important for GenAI applications? How do you enforce them?**
- LLMs naturally produce unstructured text. Applications need JSON, function arguments, or domain schema-constrained output.
- **Enforcement methods:**
  - **JSON Mode (OpenAI):** Guarantees the output is valid JSON.
  - **Function Calling / Tool Use:** LLM outputs a structured function call with typed arguments.
  - **Instructor / Outlines library:** Constrained sampling using a JSON schema grammar (guarantees schema compliance, no post-processing needed).
  - **Pydantic + retry loop:** Parse output; if invalid, send error back to LLM for self-correction.
