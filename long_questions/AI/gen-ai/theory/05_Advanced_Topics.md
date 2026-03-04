# Advanced GenAI Topics — Senior / Research Level

This file covers **frontier and advanced GenAI topics** asked at top-tier product companies, AI research labs (DeepMind, OpenAI, Anthropic, Meta FAIR), and senior engineering roles. These topics go beyond commodity LLM usage into architecture innovation, training methodologies, and modality expansion.

---

## 1. Mixture of Experts (MoE)

**Q1: How does Mixture of Experts (MoE) architecture work? Why does it reduce inference cost vs a dense model of equivalent capability?**

### Architecture Overview

```
Standard Dense Transformer FFN:
  Input → Linear(d_model → d_ff) → GELU → Linear(d_ff → d_model) → Output
  ALL parameters active for EVERY token

MoE Transformer FFN:
  Input → Router(d_model → N experts) → Top-K Expert Selection
        → Only K of N expert FFNs are run for each token
        → Weighted sum of K expert outputs → Output
  Only K/N fraction of FFN parameters active per token
```

**Key formulas:**
- **Router:** Softmax over N expert scores → select top-K (usually K=2)
- **Output:** `Σ(gate_score_i × expert_i(x))` for the K selected experts
- **Sparse activation:** With N=8 experts, K=2 → 25% of FFN params used per token

### Real-World MoE Models

| Model | Experts | Active/Token | Total Params | Active Params |
|---|---|---|---|---|
| **Mixtral 8x7B** | 8 experts × 7B | 2 experts | ~46B | ~12B |
| **Mistral 8x22B** | 8 experts × 22B | 2 experts | ~141B | ~39B |
| **GPT-4** (rumored) | ~16 experts | 2 experts | ~1.8T | ~220B |
| **Gemini 1.5 Pro** | MoE (details undisclosed) | — | — | — |

**Why MoE is efficient:**
- Same *capability* as a dense model of equivalent total parameter count.
- But *inference cost* ≈ cost of the active parameter count (much smaller).
- Trade-off: Higher memory bandwidth needed (all expert weights must fit in VRAM even if most are idle).

**Q2: What is the "load balancing" problem in MoE training?**

- **Expert collapse:** Without intervention, the router always routes to the same 1-2 experts → other experts never train → wasted capacity.
- **Fix — Auxiliary load balancing loss:**
  - Add a differentiable penalty when expert utilization is uneven.
  - `L_balance = α × N × Σ(f_i × P_i)` where f_i = fraction of tokens routed to expert i, P_i = average router probability for expert i.
- **Switch Transformer (Fedus et al., 2022):** Used this technique; first to scale MoE transformers effectively.
- **Expert Choice routing (Zhou et al., 2022):** Instead of tokens choosing experts, experts choose their top-K tokens → perfect load balancing by construction.

---

## 2. Constitutional AI (Anthropic)

**Q3: What is Constitutional AI (CAI) and how does it differ from RLHF?**

- **Paper:** Bai et al., Anthropic (2022) — "Constitutional AI: Harmlessness from AI Feedback"

### RLHF vs CAI

| Aspect | RLHF | Constitutional AI |
|---|---|---|
| **Feedback source** | Human labelers rate responses | AI model evaluates against a written Constitution |
| **Scalability** | Limited by human labeler bandwidth | Nearly fully automated |
| **Consistency** | Varies by labeler values | Consistent with explicit Constitutional principles |
| **Transparency** | Principles implicit in human preferences | Principles explicitly written and auditable |

### CAI Pipeline

**Phase 1 — Supervised Learning from AI Feedback (SL-CAF):**
```
1. Red-teaming: Prompt the initial model with harmful queries → get harmful responses
2. Critique: Ask the same model to critique its response against a specific principle:
   "Identify specific ways in which the assistant's response is harmful, unethical..."
3. Revision: Ask the model to revise the response to fix the critique
4. Repeat for multiple principles: honesty, non-maliciousness, respectfulness
5. Fine-tune on the (original prompt, final revised response) pairs
```

**Phase 2 — Reinforcement Learning from AI Feedback (RLAIF):**
```
1. Generate response pairs from the SL-CAF model
2. Ask the AI (Claude) to compare pairs using the Constitution as a rubric
3. Collect AI-generated preference labels (chosen / rejected)
4. Train a reward model on AI-labeled preferences
5. PPO (or DPO) using this AI-trained reward model
```

**The Constitution (example principles Anthropic used):**
- "Choose the response that is least likely to contain harmful or unethical content."
- "Choose the response that is most honest and avoids deception."
- "Choose the response that demonstrates more concern for the well-being of the user."

**Result:** Harmless+helpful models without needing human preference labels for safety data.

---

## 3. Audio & Speech Models

**Q4: How does OpenAI's Whisper work? What is its architecture?**

### Whisper Architecture
```
Audio Input (raw waveform)
    ↓
Mel Spectrogram (80-band, 25ms windows, 10ms stride)
    ↓
CNN Feature Encoder (2 conv layers)
    ↓
Transformer Encoder (standard, like BERT for audio)
    ↓
Transformer Decoder (autoregressive, like GPT)
    ↓
Output: Transcribed text tokens (with special tokens for language, timestamps, task)
```

- **Training:** 680,000 hours of multilingual web audio (supervised: audio → transcript pairs).
- **Multi-task:** One model does transcription, translation, language ID, and voice activity detection via special task tokens.
- **Key strength:** Extremely robust to background noise, accents, domain variation — trained on diversity.
- **Weakness:** Not real-time capable in large model sizes; Whisper Large v3 ~10 RTF (real-time factor).

**Q5: What is AudioPaLM? How does it go beyond Whisper?**

- Google's **AudioPaLM** (2023): A unified model handling both speech and text seamlessly.
- Built on PaLM-2 backbone — extends text tokens with audio tokens.
- Audio is tokenized using SoundStream (neural audio codec → discrete tokens).
- Can do: speech-to-text, text-to-speech, speech-to-speech translation, and voice continuation in **one model**.
- **Key innovation:** The same transformer processes audio tokens and text tokens together — enabling cross-modal reasoning.

**Proprietary speech models (2025 context):**
| Model | Capability | Notes |
|---|---|---|
| **GPT-4o Audio** | Native speech I/O | Real-time, emotion-aware |
| **Gemini 1.5** | Audio understanding | Long-form audio (hours) |
| **ElevenLabs** | TTS | Ultra-realistic, voice cloning |
| **Suno / Udio** | Music generation | Lyrics + composition |

---

## 4. Open-Source LLM Landscape (2025)

**Q6: Compare the leading open-source LLMs. When would you recommend each?**

| Model Family | Creator | Key Strengths | Best For |
|---|---|---|---|
| **Llama 3.1 / 3.3** | Meta | Best overall open-source at 70B/405B; multilingual | General purpose, fine-tuning baseline |
| **Mistral / Mixtral** | Mistral AI | Efficient MoE; Mistral 7B punches above weight | Resource-constrained inference |
| **Phi-3 / Phi-4** | Microsoft | Tiny but powerful (3.8B, 14B); trained on "textbooks" | Edge deployment, mobile, IoT |
| **Gemma 2 / 3** | Google | Strong reasoning; 2B–27B sizes | Fine-tuning for specific domains |
| **Qwen 2.5** | Alibaba | Excellent coding + multilingual (Chinese+English) | Asian language contexts, coding tasks |
| **DeepSeek-R1** | DeepSeek | RLVR reasoning; GPT-o1 quality at lower cost | Math, code, complex reasoning |
| **Command R** | Cohere | RAG-optimized; RAG+tool calling built in | Enterprise RAG applications |
| **Falcon** | TII | Apache 2.0 license | Permissively licensed commercial use |

**Licensing nuance (interview critical):**
- **Llama 3 License:** Free for commercial use (<700M MAU), but restricted for certain uses.
- **Mistral:** Apache 2.0 — truly permissive, no restrictions.
- **Phi-3:** MIT License — fully open.
- **Gemma:** Google's custom license — commercial use allowed with conditions.

---

## 5. Reasoning Models (o1, DeepSeek-R1, etc.)

**Q7: What makes o1/DeepSeek-R1 style reasoning models architecturally different from GPT-4?**

| Aspect | Standard LLM (GPT-4) | Reasoning Model (o1, R1) |
|---|---|---|
| **Training approach** | SFT + RLHF | RLVR (Reinforcement Learning w/ Verifiable Rewards) |
| **Inference** | Single forward pass per token | Extended "thinking" chain before answering |
| **Internal reasoning** | Implicit (hidden in weights) | Explicit chain-of-thought tokens ("thinking tokens") |
| **Strength** | Speed, general knowledge | Multi-step math, code, logic |
| **Cost** | Standard per-token cost | Higher — thinking tokens are extra |

**GRPO (Group Relative Policy Optimization) — DeepSeek's key innovation:**
```
Standard PPO needs:
  - A separate critic/value network
  - Complex reward normalization

GRPO instead:
  - For each prompt, sample G responses (e.g., G=8)
  - For math/code: check each response correctness (verifiable reward)
  - Normalize rewards by group mean and std: r_normalized = (r - mean) / std
  - Use normalized rewards as advantage estimate — no critic needed
  - Simpler, more stable training than PPO
```

---

## 6. Diffusion Models vs. Autoregressive Models for Images

**Q8: How do diffusion models work? Why are they preferred over GANs for image generation?**

### Diffusion Model Process
```
Forward process (training):
  x_0 (clean image) → add Gaussian noise for T steps → x_T (pure noise)
  This is a fixed Markov chain; q(x_t | x_{t-1}) = N(x_t; √(1-β_t)x_{t-1}, β_t I)

Reverse process (inference):
  x_T (pure noise) → neural network predicts noise at each step → x_0 (clean image)
  Network is trained to predict ε (the noise) at each timestep: L = ||ε - ε_θ(x_t, t)||²
```

**Why better than GANs:**
| Aspect | GAN | Diffusion |
|---|---|---|
| **Training stability** | Unstable (mode collapse) | Stable — no adversarial game |
| **Mode coverage** | Mode collapse — limited diversity | Excellent diversity |
| **Image quality** | Sharp but artifacts | SOTA quality |
| **Inference speed** | Fast (single forward pass) | Slow (T denoising steps, ~20–50) |
| **Controllability** | Hard to condition | Excellent — CFG, ControlNet |

**SDXL / FLUX / Stable Diffusion architecture (interview depth):**
- **U-Net** backbone (for older SD) or **DiT (Diffusion Transformer)** for FLUX/SD3.
- **CLIP** text encoder converts prompt to embeddings that condition the denoising.
- **VAE** compresses images to latent space (8x downsampling) — denoising happens in latent space, not pixel space (Latent Diffusion Models).
- **CFG (Classifier-Free Guidance):** Run model twice — conditioned + unconditioned → interpolate: `output = uncond + scale * (cond - uncond)`. Higher scale = stronger prompt adherence.

---

## 7. Multimodal Architecture (Beyond Vision)

**Q9: How is document understanding (OCR + Layout + Semantics) handled in modern multimodal models?**

**Traditional pipeline:**
```
PDF → OCR (Tesseract) → Extract text → Feed to LLM
Problem: Loses layout, tables, charts, reading order
```

**Modern approaches:**
1. **Document AI (Azure Form Recognizer, Google Document AI):** Combines OCR + layout analysis + named entity extraction. Returns structured JSON with bounding boxes.
2. **Nougat (Meta):** End-to-end transformer that reads PDFs directly as images → outputs LaTeX (preserves math, tables).
3. **GPT-4V / Gemini 1.5:** Feed page images directly → VLM understands layout natively. Best for complex documents but expensive at scale.
4. **ColPali:** CLIP-style retrieval where entire document pages are embedded as images for retrieval (bypasses OCR entirely).

**When to use which:**
- `High accuracy + structured output` → Azure Document Intelligence
- `Research papers with math` → Nougat
- `General document QA at low volume` → GPT-4V page images
- `Document retrieval at scale` → ColPali

---

## Summary: Advanced Topics Cheat Sheet

| Topic | Key Terms | One-Liner |
|---|---|---|
| **MoE** | Sparse activation, Router, Expert collapse, Load balancing loss | N experts, K active per token → big model, small inference cost |
| **Constitutional AI** | RLAIF, Critique-Revision, Constitution | Replace human preference labels with AI self-critique against explicit principles |
| **Whisper** | Mel spectrogram, Encoder-Decoder, Multi-task | Supervised trained on 680k hours; robust multilingual ASR |
| **AudioPaLM** | SoundStream, Unified tokens | Audio + text in one PaLM-backbone model |
| **Open-source LLMs** | Llama 3, Mistral, Phi-3, Qwen | Pick by: size constraint, license, domain, language |
| **Reasoning models** | RLVR, GRPO, Thinking tokens | Verifiable reward + group normalization → no critic, stable CoT training |
| **Diffusion** | U-Net/DiT, CFG, Latent space | T-step denoising; better diversity + quality than GANs |
| **Document AI** | ColPali, Nougat, Form Recognizer | Match tool to document complexity and scale requirements |
