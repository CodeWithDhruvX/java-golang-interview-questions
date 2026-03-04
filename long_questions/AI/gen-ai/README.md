# Gen-AI Interview Questions

This directory contains comprehensive interview preparation material for **Generative AI roles**, organized by company type and depth.

---

## 📁 Directory Structure

```
gen-ai/
├── gen_ai_questins.md               ← General 100+ GenAI interview questions
├── README.md                        ← This file
├── theory/                          ← Foundational concepts by experience level
│   ├── 01_Freshers.md
│   ├── 02_Intermediate.md
│   ├── 03_Advanced.md
│   └── 04_Senior.md
├── product_based_companies/         ← Deep technical questions (FAANG, AI labs)
│   ├── product_based_gen_ai.md      ← LLMs, RAG, fine-tuning, alignment, inference
│   ├── mlops_and_evaluation.md      ← MLOps pipelines, evaluation, cost optimization
│   └── system_design_gen_ai.md     ← System design: chatbots, platforms, VLMs
└── service_based_companies/         ← Applied questions (TCS, Infosys, Accenture)
    ├── service_based_gen_ai.md      ← LLMs, prompt eng, RAG, cloud AI services
    ├── coding_examples_with_answers.md ← Hands-on Python + LangChain code
    └── system_design_basics.md      ← Practical system design for delivery projects
```

---

## 🎯 What's Covered

### Product-Based Companies (Google, OpenAI, Microsoft, Anthropic, Meta AI, Databricks)
- **LLM Architecture:** Transformer internals, RoPE, FlashAttention, Mamba, KV cache
- **Fine-Tuning & Alignment:** LoRA, QLoRA, RLHF, DPO, GRPO, SFT pipelines
- **Advanced RAG:** Hybrid retrieval, reranking, HyDE, RAGAS evaluation
- **Evaluation:** LLM-as-judge, hallucination detection, evaluation harnesses
- **Inference Serving:** vLLM, speculative decoding, quantization (INT4/INT8/NF4), PagedAttention
- **MLOps:** CI/CD for LLMs, drift detection, observability stacks
- **System Design:** Multi-tenant LLM platforms, grounded medical AI, code generation assistants

### Service-Based Companies (TCS, Infosys, Wipro, HCL, Accenture, Cognizant, LTIMindtree)
- **Fundamentals:** How LLMs work, tokenization, embeddings, context windows
- **Prompt Engineering:** Zero-shot, few-shot, CoT, role prompting, structured outputs
- **RAG Pipeline:** Build-by-step architecture, chunking strategies, vector DBs comparison
- **LangChain:** Key components, memory, chains, agents, vs LlamaIndex
- **Cloud AI Services:** AWS Bedrock, Azure OpenAI, GCP Vertex AI — feature comparison
- **Coding Examples:** Ready-to-run Python code for RAG, streaming, structured output, tool calling
- **Enterprise Delivery:** Data isolation, cost control, Teams integration, SAP integration, governance

---

## 💡 How to Use This for Preparation

| Your Target | Start With |
|---|---|
| **Fresher / 0–2 years** | `theory/01_Freshers.md` → `service_based_gen_ai.md` → `coding_examples_with_answers.md` |
| **Mid-level (2–5 years)** | `theory/02_Intermediate.md` → `service_based_gen_ai.md` → `system_design_basics.md` |
| **Senior / Product companies** | `product_based_gen_ai.md` → `mlops_and_evaluation.md` → `system_design_gen_ai.md` |
| **Practice coding** | `coding_examples_with_answers.md` — run and modify all examples |
| **System design rounds** | `system_design_gen_ai.md` (product) or `system_design_basics.md` (service) |

---

## 🔗 Related Sections
- **Agentic AI:** `../agentic/` — Agents, tool calling, LangGraph, AutoGen
- **ML Theory:** `../` parent directory for broader AI/ML interview questions
