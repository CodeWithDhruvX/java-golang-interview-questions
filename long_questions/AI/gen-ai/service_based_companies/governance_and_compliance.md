# GenAI Governance, Compliance & Cloud Strategy — Service-Based Companies

Service-based companies (Accenture, Cognizant, Wipro, TCS, HCL, Capgemini) regularly deal with **enterprise compliance, regulatory requirements, and cloud platform decisions**. This file covers the governance questions that are increasingly asked in delivery and practice lead interviews.

---

## 1. AI Regulatory Landscape

### EU AI Act (2024 — First Comprehensive AI Law)

**Q1: What is the EU AI Act? How does it categorize risks?**

The EU AI Act (effective August 2024, enforcement from 2025–2026) classifies AI systems by risk level:

| Risk Level | Category | Examples | Obligation |
|---|---|---|---|
| **Unacceptable Risk** | Banned outright | Social scoring, biometric mass surveillance, subliminal manipulation | **Prohibited** |
| **High Risk** | Regulated strictly | Medical devices, recruitment AI, credit scoring, law enforcement, critical infrastructure | Conformity assessment, human oversight, transparency, data governance |
| **Limited Risk** | Transparency obligations | Chatbots, deepfake generators, emotion recognition | Must disclose AI nature to users |
| **Minimal Risk** | No obligations | Email spam filters, AI in video games | Voluntary codes of conduct |

**GPAI (General Purpose AI) models (like GPT-4, Claude, Gemini):**
- Systems with >10^25 FLOPs training compute classified as "systemic risk" GPAI → additional obligations.
- Must conduct adversarial testing, report serious incidents to EU AI Office, publish transparency reports.

**Key obligations for service companies implementing AI for EU clients:**
- Maintain a **Technical Documentation** file showing compliance.
- Register high-risk AI in the **EU AI Act database**.
- Implement human oversight mechanisms.
- Ensure data used for high-risk AI respects GDPR.

---

### India's Digital Personal Data Protection (DPDP) Act, 2023

**Q2: What is India's DPDP Act and how does it affect GenAI deployments in India?**

**Key provisions relevant to GenAI:**

| Provision | Implication for GenAI |
|---|---|
| **Consent requirement** | If GenAI processes personal data of Indian users, explicit consent must be obtained |
| **Data Fiduciary obligation** | Companies using GenAI on personal data must appoint a Data Protection Officer (DPO) |
| **Data minimization** | Only collect/process data necessary for the stated purpose |
| **Right to erasure** | Users can request deletion — model fine-tuned on their data creates compliance challenge |
| **Data localization (upcoming)** | Certain categories of personal data may be required to stay in India |
| **Cross-border transfers** | Can only transfer personal data to countries approved by the Indian government |

**Practical implementation for service companies:**
```
Client scenario: Building an HR analytics GenAI for an Indian enterprise
Compliance checklist:
☐ Get explicit consent from employees before processing their HR data with LLMs
☐ Use Azure OpenAI India region (Central India) or GCP Vertex AI (Mumbai) 
☐ Log all data processing activities in a Record of Processing Activities (RoPA)
☐ Implement right-to-erasure workflow (challenge: embedding model may memorize data)
☐ DPA (Data Processing Agreement) with cloud provider
☐ No sending employee personal data to US-based public OpenAI API without SCCs
```

---

### GDPR & AI (Europe)

**Q3: How does GDPR interact with GenAI systems?**

| GDPR Article | GenAI Relevance |
|---|---|
| **Art. 22 — Automated Decision Making** | Decisions with "significant effect" on individuals (loans, hiring) must have human review; right to explanation |
| **Art. 5 — Data minimization** | Don't send more personal data than necessary to the LLM API |
| **Art. 17 — Right to erasure** | If a user's data was used in fine-tuning, deletion is technically complex (machine unlearning) |
| **Art. 28 — Data Processor** | OpenAI, Azure, GCP are data processors; enterprise is the data controller; DPA required |
| **Art. 35 — DPIA** | Data Protection Impact Assessment required before deploying high-risk AI processing personal data |

**Practical guidance:**
- Use **Azure OpenAI** or **AWS Bedrock** (US/EU deployments) which offer Data Processing Agreements compliant with GDPR.
- **Never send PII to the public ChatGPT API** in a business context without a DPA.
- Implement **PII masking pre-processing**: detect names, emails, IDs → substitute with tokens → send to LLM → re-substitute in response.

---

## 2. Azure AI Foundry / AI Studio (2025 Update)

**Q4: What is Azure AI Foundry? How is it different from Azure Machine Learning Studio?**

**Azure AI Foundry** (rebranded from Azure AI Studio in late 2024) is Microsoft's unified platform for building enterprise AI applications.

| Feature | Azure AI Foundry | Azure Machine Learning |
|---|---|---|
| **Primary focus** | LLM/GenAI app development | Traditional ML training & MLOps |
| **Model catalog** | GPT-4o, Llama 3, Mistral, Phi-3, Command R | Custom models |
| **Prompt Flow** | Built-in visual prompt engineering + testing | Limited |
| **RAG pipeline** | Native Azure AI Search integration | Manual |
| **Evaluation** | Built-in GenAI evaluation (faithfulness, groundedness) | Custom metrics only |
| **Agents** | Azure AI Agent Service (preview 2025) | Not supported |

**Key Azure AI Foundry services (2025):**

```
Azure AI Foundry Hub
├── Projects (team workspaces)
│   ├── Deployments (model endpoints: GPT-4o, Llama, Phi-3)
│   ├── Prompt Flow (visual LLM pipeline orchestration)
│   ├── Evaluation (automated faithfulness/relevance scoring)
│   └── Fine-tuning (supervised fine-tuning of supported models)
├── Azure AI Search (vector + BM25 hybrid search)
├── Azure AI Document Intelligence (OCR + layout extraction)
├── Azure AI Content Safety (input/output moderation)
└── Azure AI Agent Service (tool-calling, code interpreter, file search)
```

**Q5: How do you implement a safe, enterprise RAG pipeline using Azure AI Foundry?**

```
Architecture:
1. Documents → Azure Blob Storage (private endpoint)
2. Azure AI Document Intelligence → extract text + layout from PDFs
3. Azure AI Foundry → chunk + embed → Azure AI Search (vector index)
4. Azure AI Foundry Prompt Flow:
   a. User query → embed query → AI Search hybrid retrieval
   b. Retrieved chunks → Azure AI Content Safety (input check)
   c. Prompt assembly → GPT-4o deployment
   d. Response → Azure AI Content Safety (output check)
   e. Return grounded response with citations
5. Monitoring → Azure Monitor + Application Insights
   - Log every request/response (PII-masked)
   - Alert on safety filter hits
   - Cost tracking per department

Security:
- Private endpoints for ALL services (no public internet)
- Azure AD + RBAC (role-based access)
- Customer Managed Keys (CMK) for encryption at rest
- Network isolation via Azure Virtual Network
```

---

## 3. Open-Source Model Selection Guide (Service Company Perspective)

**Q6: As a service company architect, how do you recommend an open-source model for a client?**

**Decision framework:**

```
Step 1: What is the task?
  → Pure text understanding (classification, extraction) → BERT family (smaller, cheaper)
  → Text generation, RAG, chat → Instruct LLM

Step 2: What are the constraints?
  → Must run on CPU or tiny GPU → Phi-3 mini (3.8B), Gemma 2B
  → Single consumer GPU (24GB VRAM) → Llama 3.1 8B, Mistral 7B, Phi-3 Medium
  → On-premise data center (2×A100 80GB) → Llama 3.1 70B, Mistral 8x22B
  → Cloud (unlimited budget) → GPT-4o, Claude 3.5 Sonnet

Step 3: What is the language?
  → English-primary → Llama 3, Mistral, Phi-3
  → English + Hindi/Indic → Llama 3 (decent), Sarvam-1 (specialist Indic)
  → English + Chinese → Qwen 2.5
  → Multilingual EU → Aya (Cohere), Mistral (good EU language support)

Step 4: License requirements?
  → Internal use, no redistribution → Any license works
  → Building a product / SaaS → Apache 2.0 (Mistral) or MIT (Phi-3) preferred
  → Worried about Llama license? → Apache 2.0 alternatives: Mistral, Falcon, OLMo
```

**Model comparison (2025 tier list for enterprise):**

| Tier | Model | Size | License | Best Use Case |
|---|---|---|---|---|
| **S** | GPT-4o, Claude 3.5 Sonnet | API | Proprietary | Max quality, no infra |
| **S** | Llama 3.3 70B | 70B | Meta License | Best open-weight general |
| **A** | Qwen 2.5 72B | 72B | Apache 2.0 | Code + multilingual |
| **A** | DeepSeek-R1 | 671B MoE | MIT | Complex reasoning (self-host) |
| **B** | Mistral 7B / Mixtral 8x7B | 7B / 46B | Apache 2.0 | Efficient, permissive license |
| **B** | Phi-4 (14B) | 14B | MIT | Small but strong reasoning |
| **C** | Gemma 2 9B / 27B | 9B/27B | Google License | Fine-tuning experiments |

---

## 4. GenAI Governance Framework for Enterprise Delivery

**Q7: A client wants to establish an AI governance framework before deploying GenAI. What do you recommend?**

**4-layer governance model (for Indian/global enterprise delivery):**

### Layer 1: Policy & Ethics
- **AI Use Policy:** Define approved/prohibited AI use cases. Who approves new AI projects?
- **Acceptable Use:** Can employees use ChatGPT at work? Must they use only approved tools?
- **Ethics board:** For high-impact AI (HR, lending, healthcare), require ethics review.

### Layer 2: Data Governance
- **Data classification:** What data can be sent to LLM APIs? (Public ✅ / Internal ⚠️ / Confidential ❌)
- **PII handling SOP:** Standard operating procedure for masking PII before LLM calls.
- **Retention policy:** How long are AI chat logs retained? Who can access them?

### Layer 3: Technical Controls
- **Approved model list:** Pre-vetted models/APIs that meet security standards.
- **Guardrails:** Input/output safety filtering using Azure Content Safety or AWS Guardrails.
- **Audit logging:** Every LLM request logged (prompt hash, tokens, cost, user ID).
- **Cost controls:** Per-department token budgets with hard stops.

### Layer 4: Monitoring & Improvement
- **Quality metrics dashboard:** Faithfulness, answer quality, user satisfaction over time.
- **Incident response:** Process for handling AI-caused harm incidents (wrong medical advice, discriminatory output).
- **Red-teaming schedule:** Quarterly adversarial testing of production AI systems.
- **Model update governance:** Freeze testing before upgrading model versions; canary deployment.

---

## 5. Cost Governance & FinOps for GenAI

**Q8: A client's GenAI project is 3x over the expected API cost. How do you address it?**

**Diagnosis steps:**

```python
# Analyze cost drivers — things to check in logs
cost_drivers = {
    "1. Large system prompts": "A 2000-token system prompt sent on EVERY request = massive waste",
    "2. No prompt caching": "Anthropic/OpenAI support prefix caching — repeated system prompt costs 90% less",
    "3. Wrong model for task": "Using GPT-4o for simple FAQ = 20x cost vs GPT-4o-mini",
    "4. No semantic caching": "Same question asked repeatedly — should cache response",
    "5. Excessive max_tokens": "Setting max_tokens=4000 for responses that average 200 tokens",
    "6. RAG chunk injection": "Injecting 10 chunks when 3 would suffice — wasted context",
    "7. Embedding recalculation": "Re-embedding the same documents on every request",
}
```

**Cost reduction playbook:**

| Technique | Typical Savings |
|---|---|
| **Prompt caching** (Anthropic/OpenAI) | 50–90% on system prompt tokens |
| **Semantic caching** (Redis + vector similarity) | 20–40% of requests served from cache |
| **Model tiering** (GPT-4o-mini for simple, GPT-4o for complex) | 60–80% cost reduction on routed requests |
| **Reduce max_tokens** to practical limit | 10–30% savings |
| **Reduce RAG chunks** from 10 → 3 | 20–40% prompt token savings |
| **Batch async workloads** | 50% throughput improvement = cost efficiency |

---

## 6. Interview Quick Reference

| Topic | For Which Interviews |
|---|---|
| EU AI Act risk tiers | Accenture, Capgemini, Big 4 AI practices |
| India DPDP Act | TCS, Infosys, HCL (India-focused delivery) |
| Azure AI Foundry | Microsoft ISV partners, Accenture Microsoft practice |
| Open-source model selection | Any delivery architect role |
| Governance framework | Practice lead, AI CoE lead, solution architect roles |
| FinOps for GenAI | Delivery manager, cloud architect roles |
