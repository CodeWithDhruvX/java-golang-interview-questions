# Gen-AI System Design — Product-Based Companies

These are **system design questions focused on GenAI products**, asked at companies like Google, Microsoft, Meta, Amazon Bedrock, Cohere, Anthropic, and AI-first product startups. Expect to whiteboard full end-to-end architectures covering LLM inference, RAG pipelines, agent systems, and multi-tenant platform design.

---

## 1. Core GenAI System Design Questions

**Q1: Design a production-grade AI-powered customer support chatbot that handles 100,000 conversations/day. Include LLM selection, RAG pipeline, escalation logic, and cost controls.**

**Architecture Walkthrough:**
```
User → API Gateway → Intent Classifier
                           ↓
              [Simple FAQ] → Fast retrieval (BM25 + Haiku)
              [Complex]   → Full RAG → Claude Sonnet/GPT-4o
              [Escalate]  → Human handoff queue (Zendesk API)
```
- **LLM tiering strategy:**
  - `gpt-3.5-turbo` / `claude-haiku` for greeting + simple intent.
  - `gpt-4o` / `claude-sonnet` for multi-turn troubleshooting.
- **RAG pipeline:**
  - Knowledge base: product docs, FAQs, policy documents.
  - Hybrid retrieval (BM25 + dense) → cross-encoder reranker → top-3 chunks injected.
- **Escalation logic:** If model outputs "I'm not sure" or confidence < threshold → route to human agent with conversation summary.
- **Cost controls:** Semantic caching with Redis + vector similarity for repeated queries; prompt caching for stable system prompt prefix.
- **Evaluation:** CSAT (customer satisfaction) as primary signal; LLM-as-judge for faithfulness sampling.

**Q2: Design a multi-tenant GenAI platform (like Azure OpenAI) that serves 500 enterprise customers with data isolation, model versioning, and usage-based billing.**

- **Tenant isolation:**
  - Separate vector DB namespaces per tenant for RAG.
  - Row-level security in PostgreSQL for conversation history.
  - API key scoped to tenant_id; all requests signed with tenant key.
- **Model versioning:**
  - Model registry (MLflow) with semantic versioning; tenants pin to a model version.
  - Canary slot: allows tenants to opt in to beta model versions.
- **Billing:** Metering service tracks tokens (input/output) per API call, stores in time-series DB. Billing pipeline aggregates daily and exports to Stripe / ERP.
- **Storage tiers:** Hot (Redis cache), Warm (Postgres), Cold (S3 + lifecycle policies for conversation history > 90 days).

**Q3: Design a real-time document intelligence system that allows users to "chat with" uploaded PDFs (like ChatPDF or NotebookLM).**

- **Ingestion pipeline (async):**
  - PDF upload → S3 → SQS trigger → Lambda parses PDF (pypdf2/unstructured) → chunk (512 tokens, 64 overlap) → embed (`text-embedding-3-small`) → upsert to Pinecone with `doc_id` metadata.
- **Query pipeline (sync, < 2s p95):**
  - User question → embed → ANN search in Pinecone (filter by `doc_id`) → hybrid with BM25 → rerank → inject top-5 chunks → LLM generates grounded answer with citations.
- **Citation handling:** Each chunk tagged with page number. Answer includes `[Page 12]` style references.
- **Multi-doc support:** Each user session scoped to selected doc IDs; vector search filtered accordingly.
- **Streaming:** Stream LLM response tokens to frontend via SSE (Server-Sent Events).

**Q4: Design a code generation assistant (like GitHub Copilot) for an enterprise IDE plugin. Cover model serving, context retrieval, latency, and privacy.**

- **Context retrieval (repo-aware RAG):**
  - Index the codebase: AST-based chunking (function-level, class-level) → embed → store in local or cloud vector DB.
  - At completion time: current file context + related functions (via vector search on function name/signature) → inject into prompt.
- **Latency target:** FIM (Fill-In-the-Middle) completion < 200ms perceived latency using speculative decoding and streaming.
- **Model options:**
  - Cloud: GPT-4o, Claude Sonnet – higher quality but latency + privacy concern.
  - On-device / self-hosted: CodeLlama-7B, Qwen2.5-Coder – lower latency, air-gapped enterprise option.
- **Privacy:** PII stripping from code snippets before sending to cloud APIs; customer option for on-prem VPC-deployed model.
- **Telemetry:** Track acceptance rate of completions, edit distance after acceptance – primary metrics for model evaluation.

---

## 2. Scalability & Reliability

**Q5: How do you handle burst traffic for an LLM API where p99 latency must stay under 3 seconds even at 10x normal load?**
- **Pre-warming:** Kubernetes HPA with custom metrics (GPU queue depth). Warm pool of model server replicas that can activate within 60s (vs. cold-start > 5 min).
- **Request prioritization:** Priority queues: paid tier → free tier. Paid users served first during overload.
- **Graceful degradation:** Under extreme load, route to smaller/faster model (Haiku instead of Sonnet) transparently, with a UI notice.
- **Circuit breaker:** If the LLM provider is slow, fail fast (HTTP 503) rather than holding connections open (snowball effect).
- **Caching + deflection:** Semantic cache deflects up to 30% of traffic on common queries.

**Q6: How do you design a GenAI system that guarantees grounded, hallucination-free outputs for a medical information product?**
- **Mandatory RAG:** Never generate without a retrieved source. System prompt explicitly instructs: "Only answer from the provided context."
- **Faithfulness guard:** After generation, run a lightweight faithfulness classifier (NLI-based or GPT-4o-mini as judge) to check if every claim in the answer appears in the retrieved context. If score < threshold → reject and respond with "I can only answer based on verified sources."
- **Source citation:** Every sentence linked to the specific document snippet it came from. UI shows original text on hover.
- **Human review queue:** Flag low-confidence answers (where faithfulness scores are borderline) for medical expert review before displaying.
- **Regulatory:** Maintain an immutable audit log of every prompt + retrieved context + response for compliance (HIPAA-aligned logging).

---

## 3. Multimodal Systems

**Q7: Design an AI system that accepts image + text inputs and outputs structured reports (e.g., insurance claim processing from photos).**

- **Input handling:** Accept JPEG/PNG; resize to max 2048px; convert to base64 for VLM API call.
- **Model choice:** GPT-4o Vision / Claude 3.5 Sonnet Vision – strong instruction following + structure output.
- **Prompt design:** System prompt specifies the JSON schema to output (damage type, severity, estimated cost range, affected parts list).
- **Validation:** Pydantic model validates the JSON output; if schema mismatch, retry with error context.
- **Fallback:** If image quality is too low (blurry detector pre-check), prompt user to re-upload before calling VLM (saves cost).
- **Audit trail:** Store raw image (S3), VLM prompt, raw response, validated JSON in DynamoDB per claim ID.
