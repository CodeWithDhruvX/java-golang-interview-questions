# Gen-AI System Design Basics — Service-Based Companies

This file covers **system design questions for GenAI applications** at a level appropriate for service-based company interviews (TCS, Infosys, HCL, Wipro, Accenture, Cognizant, LTIMindtree, Capgemini). These questions focus on practical architecture, vendor cloud services, and delivery-oriented design.

---

## 1. System Design Fundamentals for GenAI

**Q1: How do you design a simple GenAI FAQ chatbot for a government website? Walk through all components.**

```
Architecture:
┌─────────────────────────────────────────────┐
│ User Browser / Mobile App                   │
└──────────────┬──────────────────────────────┘
               │ HTTPS
┌──────────────▼──────────────────────────────┐
│ API Gateway (rate limiting, auth)           │
└──────────────┬──────────────────────────────┘
               │
┌──────────────▼──────────────────────────────┐
│  Chat Service (FastAPI / Node.js)           │
│  - Session management                       │
│  - Conversation history (Redis)             │
└────────┬─────────────────────────────────────┘
         │
┌────────▼────────────────────────────────────┐
│  RAG Pipeline                               │
│  1. Embed query → Azure AI Search           │
│  2. Retrieve top-5 relevant FAQ chunks      │
│  3. Inject into prompt + call Azure OpenAI  │
└─────────────────────────────────────────────┘
         │
┌────────▼────────────────────────────────────┐
│  Response with citations                    │
└─────────────────────────────────────────────┘
```

**Components:**
- **Knowledge base:** Scanned government PDF circulars → chunked → embedded → Azure AI Search.
- **LLM:** Azure OpenAI GPT-4o (data stays in Azure, suitable for government).
- **Session memory:** Redis stores last 5 turns per user session (30-min TTL).
- **Languages:** System prompt supports Hindi and English; model responds in user's language.
- **Fallback:** If confidence is low, show "Contact 1800-XXX-XXXX helpline."

---

**Q2: A retail company wants to add an AI product recommendation chatbot to their e-commerce site. Design the system.**

- **Product catalog ingestion:** Nightly export from database → embed product names + descriptions + categories → upsert to Pinecone vector DB.
- **User intent understanding:** LLM classifies query as "search", "compare", or "recommend."
- **Search path:** Hybrid (vector + keyword) search in Pinecone → return top-8 products → LLM formats into conversational response.
- **Compare path:** User mentions 2+ products → fetch from catalog DB → LLM generates structured comparison table.
- **Recommend path:** Fetch user's purchase history from DB → inject as context → LLM suggests complementary or improved products.
- **Session:** Store conversation in PostgreSQL (conversation_id, turn_id, role, content).
- **A/B testing:** Compare chatbot click-through rate vs. traditional search for conversion metrics.

---

## 2. RAG Architecture Design

**Q3: A law firm has 10,000 case files (PDFs). Design a RAG system that lets lawyers query them.**

**Ingestion Pipeline (run once + incremental):**
```
New Case File → Azure Blob Storage
      ↓
Azure Document Intelligence (form recognizer) → extracts text + tables
      ↓
Chunk by section (Introduction, Arguments, Verdict) — ~500 tokens
      ↓
Embed via text-embedding-3-small
      ↓
Store in Azure AI Search (with metadata: case_id, date, judge, category)
```

**Query Pipeline:**
```
Lawyer types query
      ↓
Filter intent: "case lookup" vs "legal precedent research" vs "clause search"
      ↓
Hybrid retrieval (BM25 + vector) with metadata filters (e.g., category=criminal, year>2020)
      ↓
Cross-encoder reranker (top 10 → top 5)
      ↓
Inject into GPT-4o → grounded answer with case citations
```

**Access Control:** Row-level security — each lawyer's search filtered to their authorized case IDs via metadata filter.

**Security:** Azure Private Link, Customer Managed Keys (CMK), audit log of every query + retrieved docs.

---

**Q4: What is hybrid retrieval in RAG and why is it better than pure vector search?**

| Approach | How it Works | Strength | Weakness |
|---|---|---|---|
| **BM25 (keyword)** | Exact keyword frequency scoring | Handles specific terms (product codes, names, IDs) | Misses semantic meaning |
| **Dense (vector)** | ANN cosine similarity | Semantic understanding | Misses exact term matches |
| **Hybrid** | Score both + RRF fusion | Best of both worlds | Slightly more complex setup |

- **RRF (Reciprocal Rank Fusion):** `score = Σ 1/(k + rank_i)` across BM25 and dense rankings. Simple, effective, no tuning needed.
- **Best practice:** For enterprise RAG, always use hybrid. BM25 catches employee IDs, model numbers; vector catches "how to," "what is" type queries.

---

## 3. Latency & Cost Design

**Q5: A client complains their AI chatbot takes 8 seconds to respond. How do you diagnose and fix it?**

**Diagnosis checklist:**
1. Break down the 8 seconds: network round trip + retrieval time + LLM inference time.
2. Log each step with timestamps in middleware.

**Optimization actions:**
- **Retrieval latency (> 1s):** Switch to managed vector DB with pre-warmed index. Check if embedding call is batched.
- **LLM latency (> 5s):** 
  - Enable streaming → user sees first tokens within 300–500ms even if full response takes 4s.
  - Switch to a faster/smaller model for simple queries (GPT-4o-mini, Claude Haiku).
  - Enable prompt caching for fixed system prompt (saves 200–400ms).
- **Reduce context size:** Don't inject 10 chunks if 3 are sufficient. Fewer tokens = faster generation.
- **Async embedding:** Pre-embed the user query while retrieval is running (pipeline the steps).
- **Result:** Typical 8s → 2–3s with these changes.

**Q6: Design a cost-monitoring system for a GenAI application used by 50,000 employees of a company.**

- **Track per request:** Store (user_id, session_id, prompt_tokens, completion_tokens, model, timestamp) in a time-series store (ClickHouse or BigQuery).
- **Budget alerts:** Set daily/monthly token budget per department; alert at 80% and hard-stop at 100%.
- **Cost dashboards:** Grafana dashboard showing: cost by department, cost per user, top consuming queries, model-wise breakdown.
- **Optimization insights:** Surface top 10 high-token prompts — these are candidates for prompt optimization.
- **Chargeback:** Monthly cost allocation report by cost center for finance team.

---

## 4. Integration Patterns

**Q7: How do you integrate a GenAI chatbot with a Microsoft Teams deployment for an enterprise client?**

- **Bot Framework SDK** (C# or Node.js) as the Teams bot backend.
- **Conversation logic:** Receive message → call internal GenAI API (FastAPI) → return response to Teams.
- **Auth:** Azure AD Bot Service handles OAuth. The bot service account gets an Azure AD application registration.
- **State:** Azure Cosmos DB stores conversation history keyed by `conversation_id` (Teams provides this).
- **SSO:** Teams user identity flows through to the bot; user's role used to filter RAG metadata.
- **File upload support:** Users can upload PDFs in Teams → bot processes and answers questions about the file (using Azure Document Intelligence + per-session vector chunk).
- **Deployment:** Azure App Service (always-on), Azure Bot Service registration.

**Q8: A client uses SAP S/4HANA. How would you add a GenAI layer that lets users query HR and finance data in natural language?**

- **NL2SQL approach (for structured data):**
  - Extract SAP schema (table names, column names, relationships) → include relevant schema in prompt.
  - LLM generates SAP-compatible SQL/ABAP query.
  - Execute on read-replica (never on production OLTP), validate results, format in plain language.
- **RAG approach (for policy docs, reports):**
  - Export SAP BI reports as PDFs → ingest into RAG pipeline.
  - Answer policy questions from indexed report content.
- **Schema security:** Only inject schema for tables the user's SAP role has access to (principle of least privilege).
- **Guardrails:** Validate generated SQL is SELECT-only; no INSERT/UPDATE/DELETE.

---

## 5. GenAI Governance & Ethics (Client Conversations)

**Q9: A client asks how to ensure their GenAI chatbot is "safe." What do you recommend?**

**Safety framework (practical for service delivery):**

1. **Input guardrails:**
   - Prompt injection detection (Llama-Guard, or rule-based pattern matching).
   - PII detection before sending to external LLM API (use AWS Comprehend or Azure PII detection to mask names, AADHAAR, phone numbers).

2. **Output guardrails:**
   - Toxicity/profanity filter on model output before displaying.
   - Faithfulness check: flag if response seems to contradict retrieved context (simple cross-check).

3. **Scope enforcement:**
   - System prompt explicitly limits the bot to its domain.
   - Add a classifier to detect out-of-scope queries before calling the LLM.

4. **Audit and logging:**
   - Log all interactions (user input + LLM output) with timestamps.
   - Regular red-teaming: test with adversarial inputs every sprint.

5. **Model governance:**
   - Pin the model version; don't auto-update. Evaluate new versions before upgrading.
   - Maintain an immutable record of the system prompt version used per interaction.

**Q10: What are the key challenges when delivering GenAI projects for Indian enterprise clients?**
- **Data privacy:** Ensuring sensitive data (customer info, internal docs) doesn't leave the country → prefer Azure OpenAI (India region) or GCP Vertex AI (Mumbai) or on-prem Ollama deployment.
- **Language support:** Many clients need Hindi, Tamil, Marathi support. Validate model performance for vernacular languages (Aya, Sarvam AI, or GPT-4 with language-specific prompts).
- **Cost expectations:** Indian enterprise clients often expect "ROI in 3 months." Design with cost efficiency as a top constraint: semantic caching, smaller models for simple tasks.
- **Change management:** End-user adoption is often harder than the tech. Budget for training, demos, and feedback loops.
- **Connectivity:** Some tier-2 city offices have bandwidth constraints → design for async, low-payload interaction patterns.
