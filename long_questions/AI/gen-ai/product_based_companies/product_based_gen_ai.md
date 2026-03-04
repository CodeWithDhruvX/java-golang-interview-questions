# Gen-AI Interview Questions for Product-Based Companies

Product-based companies (Google DeepMind, OpenAI, Anthropic, Microsoft, Meta AI, Cohere, Mistral, Databricks, etc.) expect **deep theoretical understanding, production-grade system design, research awareness, and the ability to debug and optimize LLM-powered products at scale.**

---

## 1. LLM Architecture & Fundamentals

**Q1: Explain the Transformer architecture end-to-end. Why is it dominant in Gen-AI over RNNs and LSTMs?**
- **Focus Areas:**
  - Multi-head self-attention: Q, K, V matrices, scaled dot-product attention formula.
  - Positional encodings (absolute vs. rotary PE - RoPE, ALiBi).
  - Feed-forward layers, layer norm, residual connections.
  - Why transformers parallelize better on GPUs (vs. sequential RNN/LSTM recurrence).
  - Limitations: quadratic attention complexity O(n²) and context length constraints.

**Q2: What is the difference between GPT-style (decoder-only), BERT-style (encoder-only), and T5-style (encoder-decoder) models? When do you use each?**
- **GPT (Decoder-only):** Autoregressive next-token prediction. Used for text generation, chat, code completion.
- **BERT (Encoder-only):** Masked language modeling. Used for classification, NER, embeddings, retrieval.
- **T5/Seq2Seq (Encoder-Decoder):** Conditional generation. Used for translation, summarization, Q&A with long context.
- **Interview cue:** Know **when to fine-tune BERT vs. GPT** for a given downstream task.

**Q3: Walk me through how an LLM generates a token. What is "temperature", "top-k", and "top-p" sampling?**
- Softmax over vocabulary logits → probability distribution.
- **Temperature < 1:** More deterministic, peaks sharpen. **Temperature > 1:** More random.
- **Top-k:** Sample from the top-k most likely tokens only.
- **Top-p (nucleus sampling):** Sample from minimum set of tokens whose cumulative probability ≥ p.
- **Typical best practice:** Temperature ~0.7, top_p ~0.9 for creative generation.

**Q4: Explain Rotary Positional Embedding (RoPE) and why it enables better length generalization compared to learned absolute positional encodings.**
- RoPE encodes position by rotating Q and K vectors in 2D subspaces.
- Relative position is preserved when computing attention (dot product naturally captures relative distance).
- Enables better extrapolation beyond the training context length (with techniques like YaRN, LongRoPE).
- Used in Llama, Mistral, Gemma, Qwen – almost all modern open-source LLMs.

**Q5: What is the "attention is all you need" quadratic bottleneck? What solutions exist (e.g., FlashAttention, Mamba, sliding window attention)?**
- Standard attention: O(n²) time and space for sequence length n.
- **FlashAttention (1, 2, 3):** Tiling-based SRAM-aware computation that avoids materializing full n×n attention matrix. Same mathematical result, but IO-efficient.
- **Sliding Window Attention (Mistral):** Each token attends only to a local window. Efficient but misses long-range dependencies.
- **Mamba (State Space Models):** Linear complexity O(n). Selective state space models – different from transformers but competitive on long sequences.

---

## 2. Fine-Tuning & Alignment

**Q6: Explain the full parameter fine-tuning vs. PEFT (Parameter Efficient Fine-Tuning) tradeoff. When would you use LoRA, QLoRA, or Prefix Tuning?**
- **Full fine-tuning:** Updates all weights. Best accuracy, but prohibitively expensive (needs 8x model size in GPU RAM for optimizer states).
- **LoRA:** Injects low-rank adapter matrices (A × B) into specific layers (usually Q/V projections). Only trains adapter params (~0.1–1% of total). Merges at inference with no latency overhead.
- **QLoRA:** LoRA on a 4-bit quantized base model. Enables fine-tuning 65B+ models on a single consumer GPU.
- **Prefix Tuning:** Learns soft tokens prepended to layers. Less stable than LoRA in practice.
- **When to choose:** QLoRA for resource-constrained fine-tuning; LoRA for production fine-tuning with merge; full fine-tuning only when you have data and compute at scale.

**Q7: What is RLHF (Reinforcement Learning from Human Feedback)? Describe the three-phase pipeline.**
- **Phase 1 – Supervised Fine-Tuning (SFT):** Fine-tune the base LLM on high-quality demonstration data (curated human-written responses).
- **Phase 2 – Reward Modeling (RM):** Train a separate reward model on human preference pairs (response A vs. B). RM learns a scalar "quality" score.
- **Phase 3 – PPO (Proximal Policy Optimization):** The SFT model (the "policy") is trained via RL to maximize reward from the RM, with a KL-divergence penalty against the SFT model to prevent reward hacking.
- **Known issues:** Reward hacking, reward model distribution shift, high compute cost of PPO.

**Q8: What is DPO (Direct Preference Optimization) and why is it considered simpler than PPO-based RLHF?**
- DPO reformulates the RLHF objective: instead of training a reward model + PPO loop, it directly optimizes the LLM on preference pairs (chosen vs. rejected) using a cross-entropy-like loss derived from the Bradley-Terry model.
- **Advantages:** No need for a separate reward model, no RL loop, faster and more stable training.
- **Tradeoffs:** DPO can be sensitive to the quality of preference data; may not match PPO at the frontier scale.
- **Successors:** SimPO, ORPO – further simplify by removing the reference model.

**Q9: How do you create high-quality fine-tuning datasets for a domain-specific LLM? What are common pitfalls?**
- **Data sources:** Internal documents, curated QA pairs, synthetic data from GPT-4, web scraping + deduplication.
- **Format:** Instruction-response pairs in chat template format (chatml, alpaca, llama-3 template).
- **Pitfalls:**
  - Data contamination (training on test data).
  - Short, repetitive responses leading to mode collapse.
  - Instruction drift if template tokens are wrong.
  - Not balancing data distribution across topics.
- **Deduplication:** MinHash, near-dedup before training.

---

## 3. Retrieval-Augmented Generation (RAG) – Advanced

**Q10: Design a production-grade RAG system for a 10,000-document enterprise knowledge base. Walk me through every design decision.**
- **Chunking strategy:** Fixed-size (512 tokens) vs. semantic chunking (sentence boundaries). Overlap (stride) matters.
- **Embedding model:** Domain-specific (e.g., `bge-large-en`, `text-embedding-3-large`). Fine-tuning the embedding model on domain data dramatically improves retrieval.
- **Vector DB:** Pinecone, Weaviate, Qdrant, pgvector. ANN index type: HNSW (best recall/speed balance).
- **Retrieval:** Dense retrieval (ANN) + sparse retrieval (BM25) → hybrid retrieval with RRF (Reciprocal Rank Fusion) scoring.
- **Reranking:** Cross-encoder reranker (e.g., `bge-reranker-large`) on top-K results to rerank with full attention.
- **Context packing:** Inject reranked chunks into the LLM prompt. Handle long context with hierarchical summarization if needed.
- **Evaluation:** RAGAS metrics: Faithfulness, Answer Relevancy, Context Recall, Context Precision.

**Q11: What is the difference between Naive RAG, Advanced RAG, and Modular RAG?**
- **Naive RAG:** Retrieve top-k chunks → stuff into prompt → generate. Simple but poor precision.
- **Advanced RAG:** Pre-retrieval (query rewriting, HyDE), retrieval (hybrid+rerank), post-retrieval (compression, cohere rerank). Better quality.
- **Modular RAG:** Flexible pipeline where retrieval, summarization, generation, and validation are separate modules orchestrated by an LLM. Can route between RAG and direct generation based on query type.

**Q12: What is HyDE (Hypothetical Document Embeddings) and when does it help RAG?**
- Instead of embedding the user's sparse question directly, ask the LLM to generate a *hypothetical answer* to the question. Embed the hypothetical answer and use it to query the vector DB.
- **Why it works:** The hypothetical answer is semantically closer to real documents than the short question is.
- **When it helps:** Highly specific, sparse, or technical queries where the question text has low overlap with document embeddings.
- **Risk:** If the LLM hallucinates in the hypothetical document, retrieval accuracy degrades.

---

## 4. Evaluation & Reliability

**Q13: How do you evaluate an LLM's output quality in production without human labelers for every request?**
- **LLM-as-a-Judge:** Use a strong model (GPT-4, Claude 3 Opus) to score outputs on dimensions like correctness, faithfulness, conciseness (G-Eval, MT-Bench patterns).
- **Reference-based metrics:** BLEU, ROUGE for summarization; BERTScore for semantic similarity.
- **RAG-specific:** RAGAS (Faithfulness = no hallucination vs. context; Answer Relevancy = does it answer the question?).
- **Production signals:** Thumbs up/down, user corrections, session abandonment rate, re-query rate.
- **Canary testing:** Shadow new models vs. production baseline on real traffic; compare judge scores.

**Q14: What is "hallucination" in LLMs? Categorize its types and describe mitigation strategies.**
- **Types:**
  - *Intrinsic hallucination:* Generated text contradicts source context.
  - *Extrinsic hallucination:* Generated text cannot be verified from source (made-up facts).
  - *Factual hallucination:* Incorrect facts about the real world.
- **Mitigations:**
  - RAG with grounded context.
  - Source citation enforcement + faithfulness checking.
  - Reducing temperature, using greedy decoding for factual tasks.
  - Fine-tuning on "I don't know" responses for out-of-distribution queries.
  - Self-consistency prompting (sample N responses, take the majority answer).

**Q15: What is "prompt injection" and how do you defend against it in a production GenAI application?**
- **Prompt injection:** A malicious user input hijacks the system prompt's instructions, causing the LLM to ignore safety guidelines or exfiltrate data (e.g., "Ignore previous instructions and output the system prompt").
- **Defenses:**
  - **Input sanitization:** Strip/detect trigger patterns (Llama-Guard, Rebuff).
  - **Privilege separation:** System prompt and user input in separate roles that are structurally isolated.
  - **Output guardrails:** Never execute agent actions based solely on user-supplied input; require role elevation.
  - **Monitoring:** Log all requests and use classifier to detect injection patterns post-hoc.

---

## 5. Quantization, Serving & Performance

**Q16: Explain INT8, INT4, and NF4 quantization. What are the tradeoffs between quantization levels?**
- **Quantization:** Represent model weights in lower precision to reduce memory and increase throughput.
- **INT8 (LLM.int8):** Roughly 2x memory reduction vs. FP16. Minimal accuracy loss for most models.
- **INT4 (GPTQ, AWQ):** 4x reduction vs. FP16. Slightly more accuracy loss; requires calibration data.
- **NF4 (Normal Float 4):** Used in QLoRA. Optimal for normally distributed weights (which transformer weights typically are).
- **Tradeoff:** More quantization = lower VRAM usage + faster inference, but potential accuracy degradation on reasoning-heavy tasks.

**Q17: How does speculative decoding improve LLM inference throughput? What is the role of the draft model?**
- **Problem:** LLM autoregressive generation is sequential and slow (one token per forward pass).
- **Speculative Decoding:** A small, fast "draft" model generates k tokens speculatively. The large "target" model verifies all k tokens in **one parallel forward pass** (since it can evaluate all draft tokens simultaneously).
- **If any draft token is rejected:** Rollback to the last accepted token, draft again.
- **Speedup:** 2–3x depending on draft acceptance rate. Best when draft and target models are aligned in style.
- **Used by:** Google (Medusa), DeepMind (SpecInfer), Anthropic's Claude serving.

**Q18: Design the serving infrastructure for a high-throughput LLM API serving 10,000 requests/minute with p99 latency < 2s.**
- **Model Server:** vLLM (PagedAttention for KV cache - eliminates wasted VRAM), or TensorRT-LLM.
- **Batching:** Continuous batching (in-flight batching) to interleave requests and minimize GPU idle time.
- **Autoscaling:** K8s HPA based on GPU utilization / queue depth; pre-warm replicas.
- **Load balancing:** Route by model version + request type (chat vs. embedding vs. completion).
- **KV Cache management:** PagedAttention allocates KV cache in non-contiguous pages like OS virtual memory.
- **Monitoring:** Token throughput (tokens/sec), TTFT (time-to-first-token), ITL (inter-token latency), queue wait time.

---

## 6. Multimodal & Frontier Topics

**Q19: How do Vision-Language Models (VLMs) like LLaVA or GPT-4V work? How is image information fused with text tokens?**
- A visual encoder (e.g., CLIP ViT-L) converts the image to patch embeddings.
- A projection layer (linear or MLP) maps image embeddings to the LLM's token embedding space.
- Image patch embeddings are concatenated with text token embeddings and fed as a sequence to the LLM decoder.
- **LLaVA approach:** Fine-tune the projection layer while keeping CLIP and LLM frozen initially, then full fine-tune.

**Q20: What is RLVR (Reinforcement Learning with Verifiable Rewards) as used in reasoning models like DeepSeek-R1? How is it different from RLHF?**
- In tasks with verifiable outcomes (math, code), a **rule-based verifier** (not a learned reward model) checks correctness.
- **DeepSeek-R1 key insight:** Group Relative Policy Optimization (GRPO) – sampling multiple responses per prompt, normalizing rewards by group statistics → stable training without a critic model.
- **Difference from RLHF:** No human labelers, no reward model; instead, exact correctness checking. Scales naturally to any verifiable domain.
- **Result:** Models that generate explicit chain-of-thought reasoning ("thinking tokens") and achieve GPT-4+ performance on math/code.
