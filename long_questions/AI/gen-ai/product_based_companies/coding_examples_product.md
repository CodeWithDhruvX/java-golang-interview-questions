# GenAI Coding Examples — Product-Based Companies

Product-based companies (Google, OpenAI, Anthropic, Databricks, Cohere, Hugging Face) often include **coding rounds** that test your ability to implement core GenAI components from scratch in Python/PyTorch. This file contains the most commonly asked coding challenges with full solutions.

---

## 1. Scaled Dot-Product Attention (From Scratch)

> **Asked at:** Google DeepMind, Hugging Face, Cohere, AI research roles

```python
import torch
import torch.nn.functional as F
import math

def scaled_dot_product_attention(
    Q: torch.Tensor,  # (batch, heads, seq_len, d_k)
    K: torch.Tensor,  # (batch, heads, seq_len, d_k)
    V: torch.Tensor,  # (batch, heads, seq_len, d_k)
    mask: torch.Tensor = None  # (batch, 1, 1, seq_len) — causal mask
) -> torch.Tensor:
    """
    Computes Scaled Dot-Product Attention.
    
    Attention(Q, K, V) = softmax(QK^T / sqrt(d_k)) * V
    """
    d_k = Q.size(-1)  # key dimension
    
    # Step 1: Compute raw attention scores
    # (batch, heads, seq_len, seq_len)
    scores = torch.matmul(Q, K.transpose(-2, -1)) / math.sqrt(d_k)
    
    # Step 2: Apply causal mask (upper triangle = -inf → softmax → 0)
    if mask is not None:
        scores = scores.masked_fill(mask == 0, float('-inf'))
    
    # Step 3: Softmax to get attention weights
    attn_weights = F.softmax(scores, dim=-1)  # (batch, heads, seq_len, seq_len)
    
    # Step 4: Weighted sum of values
    output = torch.matmul(attn_weights, V)  # (batch, heads, seq_len, d_k)
    
    return output, attn_weights


def causal_mask(seq_len: int) -> torch.Tensor:
    """Creates a lower-triangular mask for autoregressive (decoder) attention."""
    mask = torch.tril(torch.ones(seq_len, seq_len))
    return mask.unsqueeze(0).unsqueeze(0)  # (1, 1, seq_len, seq_len)


# --- Example usage ---
batch, heads, seq_len, d_k = 2, 4, 10, 64
Q = torch.randn(batch, heads, seq_len, d_k)
K = torch.randn(batch, heads, seq_len, d_k)
V = torch.randn(batch, heads, seq_len, d_k)
mask = causal_mask(seq_len)

output, weights = scaled_dot_product_attention(Q, K, V, mask)
print(f"Output shape: {output.shape}")   # (2, 4, 10, 64)
print(f"Weight shape: {weights.shape}")  # (2, 4, 10, 10)
```

**Key interview points:**
- Division by √d_k prevents gradients from vanishing in softmax for large d_k
- Causal mask ensures each position only attends to earlier positions (no future peeking)
- Time complexity: O(n²·d_k) — the quadratic bottleneck in standard attention

---

## 2. Multi-Head Attention Module (From Scratch)

> **Asked at:** Anthropic, Meta FAIR, Microsoft Research

```python
import torch
import torch.nn as nn
import math

class MultiHeadAttention(nn.Module):
    def __init__(self, d_model: int, num_heads: int, dropout: float = 0.1):
        super().__init__()
        assert d_model % num_heads == 0, "d_model must be divisible by num_heads"
        
        self.d_model = d_model
        self.num_heads = num_heads
        self.d_k = d_model // num_heads  # dimension per head
        
        # Weight matrices (single linear layer, then split)
        self.W_q = nn.Linear(d_model, d_model, bias=False)
        self.W_k = nn.Linear(d_model, d_model, bias=False)
        self.W_v = nn.Linear(d_model, d_model, bias=False)
        self.W_o = nn.Linear(d_model, d_model, bias=False)
        
        self.dropout = nn.Dropout(dropout)
        
    def split_heads(self, x: torch.Tensor) -> torch.Tensor:
        """Reshape (batch, seq, d_model) → (batch, heads, seq, d_k)"""
        batch, seq, d_model = x.shape
        x = x.view(batch, seq, self.num_heads, self.d_k)
        return x.transpose(1, 2)  # (batch, heads, seq, d_k)
    
    def forward(self, Q, K, V, mask=None):
        batch = Q.size(0)
        
        # 1. Linear projections
        Q = self.split_heads(self.W_q(Q))  # (batch, heads, seq, d_k)
        K = self.split_heads(self.W_k(K))
        V = self.split_heads(self.W_v(V))
        
        # 2. Scaled dot-product attention per head
        scores = torch.matmul(Q, K.transpose(-2, -1)) / math.sqrt(self.d_k)
        if mask is not None:
            scores = scores.masked_fill(mask == 0, float('-inf'))
        attn_weights = torch.softmax(scores, dim=-1)
        attn_weights = self.dropout(attn_weights)
        context = torch.matmul(attn_weights, V)  # (batch, heads, seq, d_k)
        
        # 3. Concatenate heads and final projection
        context = context.transpose(1, 2).contiguous()
        context = context.view(batch, -1, self.d_model)  # (batch, seq, d_model)
        
        return self.W_o(context)


# --- Test ---
d_model, num_heads, seq_len = 512, 8, 20
mha = MultiHeadAttention(d_model=d_model, num_heads=num_heads)
x = torch.randn(2, seq_len, d_model)
# Self-attention: Q = K = V = x
out = mha(x, x, x)
print(f"MHA output shape: {out.shape}")  # (2, 20, 512)
```

---

## 3. BPE (Byte Pair Encoding) Tokenizer

> **Asked at:** Hugging Face, OpenAI, Cohere — "Implement a basic BPE tokenizer"

```python
from collections import Counter, defaultdict
from typing import Dict, List, Tuple

def get_vocab(corpus: List[str]) -> Dict[str, int]:
    """
    Build initial character-level vocabulary.
    Each word is split into characters separated by spaces.
    End-of-word marker </w> is added to distinguish word boundaries.
    """
    vocab = Counter()
    for sentence in corpus:
        for word in sentence.split():
            # "low" → "l o w </w>"
            word_chars = ' '.join(list(word)) + ' </w>'
            vocab[word_chars] += 1
    return dict(vocab)


def get_stats(vocab: Dict[str, int]) -> Dict[Tuple[str, str], int]:
    """Count frequency of every adjacent pair of symbols in the vocabulary."""
    pairs = defaultdict(int)
    for word, freq in vocab.items():
        symbols = word.split()
        for i in range(len(symbols) - 1):
            pairs[(symbols[i], symbols[i + 1])] += freq
    return pairs


def merge_vocab(pair: Tuple[str, str], vocab: Dict[str, int]) -> Dict[str, int]:
    """Merge the most frequent pair across all vocabulary entries."""
    new_vocab = {}
    bigram = ' '.join(pair)
    replacement = ''.join(pair)
    for word in vocab:
        new_word = word.replace(bigram, replacement)
        new_vocab[new_word] = vocab[word]
    return new_vocab


def train_bpe(corpus: List[str], num_merges: int) -> List[Tuple[str, str]]:
    """
    Train BPE for num_merges steps.
    Returns the list of merge rules (in order they were learned).
    """
    vocab = get_vocab(corpus)
    merges = []
    
    print("Initial vocab:", vocab)
    
    for i in range(num_merges):
        pairs = get_stats(vocab)
        if not pairs:
            break
        
        # Pick the most frequent pair
        best_pair = max(pairs, key=pairs.get)
        vocab = merge_vocab(best_pair, vocab)
        merges.append(best_pair)
        
        print(f"Merge {i+1}: {best_pair} → {''.join(best_pair)}  (freq={pairs[best_pair]})")
    
    return merges


# --- Example ---
corpus = [
    "low low low low",
    "lower lower",
    "newest newest newest newest newest newest",
    "widest widest widest"
]

merges = train_bpe(corpus, num_merges=10)
print("\nLearned BPE merges:", merges)
```

**Key interview points:**
- BPE starts with character-level vocabulary and iteratively merges the most frequent adjacent pair
- The merge order learned on training data is then applied to encode new text
- GPT-2 uses BPE; GPT-4 uses `tiktoken` which is a Rust-optimized BPE variant

---

## 4. LoRA (Low-Rank Adaptation) Module

> **Asked at:** Cohere, Databricks, Hugging Face — "Implement a LoRA adapter layer"

```python
import torch
import torch.nn as nn

class LoRALinear(nn.Module):
    """
    LoRA adapter wrapping a frozen Linear layer.
    
    Instead of updating W (d_out, d_in), we learn:
        W_adapted = W_frozen + α/r * B @ A
    where A ∈ R(r, d_in) and B ∈ R(d_out, r); r << min(d_in, d_out)
    """
    def __init__(
        self,
        d_in: int,
        d_out: int,
        rank: int = 4,
        alpha: float = 1.0,
        bias: bool = True
    ):
        super().__init__()
        self.rank = rank
        self.alpha = alpha
        self.scaling = alpha / rank  # typical LoRA scaling factor
        
        # Frozen pre-trained weight (not trained)
        self.linear = nn.Linear(d_in, d_out, bias=bias)
        self.linear.weight.requires_grad = False
        if bias:
            self.linear.bias.requires_grad = False
        
        # Trainable low-rank decomposition: A (r × d_in), B (d_out × r)
        self.lora_A = nn.Parameter(torch.randn(rank, d_in) * 0.01)
        self.lora_B = nn.Parameter(torch.zeros(d_out, rank))
        # Note: B initialized to 0 so initial LoRA delta = 0 (no disturbance)
    
    def forward(self, x: torch.Tensor) -> torch.Tensor:
        # Original frozen output
        base_output = self.linear(x)
        # LoRA delta: x @ A^T @ B^T * scaling
        lora_delta = (x @ self.lora_A.T) @ self.lora_B.T * self.scaling
        return base_output + lora_delta
    
    def merge_weights(self):
        """
        Merge LoRA weights into the base linear layer.
        After merging, the adapter has zero overhead at inference.
        """
        with torch.no_grad():
            self.linear.weight += self.scaling * (self.lora_B @ self.lora_A)
        # Set LoRA params to zero so they don't double-count
        self.lora_A.data.zero_()
        self.lora_B.data.zero_()
    
    def trainable_parameters(self) -> int:
        return self.lora_A.numel() + self.lora_B.numel()


# --- Example ---
d_in, d_out, rank = 768, 768, 8
lora_layer = LoRALinear(d_in=d_in, d_out=d_out, rank=rank, alpha=16.0)

# Only LoRA params are trainable
trainable = sum(p.numel() for p in lora_layer.parameters() if p.requires_grad)
total = sum(p.numel() for p in lora_layer.parameters())
print(f"Trainable params: {trainable:,} / Total params: {total:,}")
print(f"LoRA efficiency: {100 * trainable / total:.2f}% of total parameters")
# Expected: only ~12,288 trainable vs 590,592 total (~2% efficiency)

x = torch.randn(4, 32, d_in)
out = lora_layer(x)
print(f"Output shape: {out.shape}")  # (4, 32, 768)
```

**Key interview points:**
- `lora_B` initialized to 0 ensures the initial LoRA delta is 0 — no disturbance to the pretrained weights
- `merge_weights()` eliminates adapter overhead at inference (critical for production)
- QLoRA = LoRA on a 4-bit NF4 quantized base model via `bitsandbytes`

---

## 5. KV Cache Implementation

> **Asked at:** vLLM team interviews, NVIDIA, Inference-focused engineering roles

```python
import torch
from typing import Optional, Tuple

class KVCache:
    """
    Simple Key-Value Cache for autoregressive LLM decoding.
    
    During autoregressive generation, we avoid recomputing K and V
    for all previous tokens at each new step.
    """
    def __init__(self, max_batch_size: int, max_seq_len: int, num_heads: int, d_k: int):
        self.max_seq_len = max_seq_len
        # Pre-allocate buffers
        self.k_cache = torch.zeros(max_batch_size, num_heads, max_seq_len, d_k)
        self.v_cache = torch.zeros(max_batch_size, num_heads, max_seq_len, d_k)
        self.current_len = 0
    
    def update(self, K_new: torch.Tensor, V_new: torch.Tensor) -> Tuple[torch.Tensor, torch.Tensor]:
        """
        Append new K/V for the current token step.
        Returns full K and V sequence (all past + current).
        """
        step_len = K_new.shape[2]  # number of new tokens (usually 1 in autoregressive mode)
        
        start = self.current_len
        end = self.current_len + step_len
        assert end <= self.max_seq_len, "Sequence length exceeds cache capacity"
        
        self.k_cache[:K_new.shape[0], :, start:end, :] = K_new
        self.v_cache[:V_new.shape[0], :, start:end, :] = V_new
        self.current_len = end
        
        # Return full accumulated K, V up to current position
        K_full = self.k_cache[:K_new.shape[0], :, :end, :]
        V_full = self.v_cache[:V_new.shape[0], :, :end, :]
        return K_full, V_full
    
    def reset(self):
        self.k_cache.zero_()
        self.v_cache.zero_()
        self.current_len = 0


# --- Simulate autoregressive decoding with KV cache ---
batch, heads, d_k = 2, 4, 64
cache = KVCache(max_batch_size=batch, max_seq_len=100, num_heads=heads, d_k=d_k)

print("Simulating 5-step autoregressive generation with KV cache:")
for step in range(5):
    # Each step processes only 1 new token
    K_new = torch.randn(batch, heads, 1, d_k)
    V_new = torch.randn(batch, heads, 1, d_k)
    
    K_full, V_full = cache.update(K_new, V_new)
    print(f"  Step {step+1}: K_full shape = {K_full.shape}, V_full shape = {V_full.shape}")
    # Without cache: would recompute ALL K, V from scratch every step → O(n²) wasted work
    # With cache: only compute K, V for new token, retrieve past from cache → O(1) per step
```

---

## 6. Simple Transformer Block

> **Asked at:** DeepMind, Meta FAIR — "Implement a full Transformer decoder block"

```python
import torch
import torch.nn as nn
import math

class FeedForward(nn.Module):
    def __init__(self, d_model: int, d_ff: int, dropout: float = 0.1):
        super().__init__()
        # Standard FFN: d_model → d_ff → d_model (with ReLU or GELU)
        self.net = nn.Sequential(
            nn.Linear(d_model, d_ff),
            nn.GELU(),          # GELU used by GPT-2/3/4; original paper used ReLU
            nn.Dropout(dropout),
            nn.Linear(d_ff, d_model),
            nn.Dropout(dropout),
        )
    
    def forward(self, x):
        return self.net(x)


class TransformerDecoderBlock(nn.Module):
    """
    GPT-style decoder block (no cross-attention — decoder-only architecture).
    Uses Pre-LayerNorm (modern variant, used in GPT-2 and later).
    """
    def __init__(self, d_model: int, num_heads: int, d_ff: int, dropout: float = 0.1):
        super().__init__()
        self.ln1 = nn.LayerNorm(d_model)
        self.attn = nn.MultiheadAttention(d_model, num_heads, dropout=dropout, batch_first=True)
        self.ln2 = nn.LayerNorm(d_model)
        self.ff = FeedForward(d_model, d_ff, dropout)
    
    def forward(self, x: torch.Tensor, causal_mask: torch.Tensor = None) -> torch.Tensor:
        # Pre-norm + self-attention + residual
        x_norm = self.ln1(x)
        attn_out, _ = self.attn(x_norm, x_norm, x_norm, attn_mask=causal_mask)
        x = x + attn_out  # residual connection
        
        # Pre-norm + feed-forward + residual
        x = x + self.ff(self.ln2(x))
        return x


def make_causal_mask(seq_len: int) -> torch.Tensor:
    """Upper-triangular mask for PyTorch MHA (additive: 0 = attend, -inf = mask)."""
    mask = torch.triu(torch.ones(seq_len, seq_len), diagonal=1)
    return mask.masked_fill(mask == 1, float('-inf'))


# --- Test a small GPT-like stack of decoder blocks ---
d_model, num_heads, d_ff, seq_len = 256, 4, 1024, 15
block = TransformerDecoderBlock(d_model=d_model, num_heads=num_heads, d_ff=d_ff)

x = torch.randn(2, seq_len, d_model)
causal_mask = make_causal_mask(seq_len)
out = block(x, causal_mask)
print(f"Decoder block output shape: {out.shape}")  # (2, 15, 256)
```

---

## 7. RAGAS-Style Faithfulness Check (Mini Implementation)

> **Asked at:** Evaluation-focused roles at Cohere, Scale AI, Weights & Biases

```python
from openai import OpenAI

client = OpenAI()

def check_faithfulness(context: str, answer: str) -> dict:
    """
    Checks if every claim in 'answer' is supported by 'context'.
    Returns a score between 0 and 1.
    Uses LLM-as-a-judge (RAGAS Faithfulness metric pattern).
    """
    prompt = f"""
You are a strict fact-verifier. Given a CONTEXT and an ANSWER, identify each claim made in the answer.
For each claim, determine if it is supported by the context or not.

CONTEXT:
{context}

ANSWER:
{answer}

Respond in this exact JSON format:
{{
  "claims": [
    {{"claim": "<claim text>", "supported": true/false, "reason": "<brief reason>"}}
  ],
  "faithfulness_score": <float between 0 and 1, = supported_claims / total_claims>
}}
"""
    response = client.chat.completions.create(
        model="gpt-4o-mini",
        messages=[{"role": "user", "content": prompt}],
        response_format={"type": "json_object"},
        temperature=0
    )
    import json
    return json.loads(response.choices[0].message.content)


# --- Test ---
context = """
The Eiffel Tower is located in Paris, France. It was built between 1887 and 1889 
as the entrance arch for the 1889 World's Fair. It stands 330 meters tall.
"""

answer_good = "The Eiffel Tower is in Paris and was completed in 1889."
answer_bad = "The Eiffel Tower is in London and was built in 1950. It stands 500 meters tall."

print("=== Faithful Answer ===")
result = check_faithfulness(context, answer_good)
print(f"Faithfulness score: {result.get('faithfulness_score')}")

print("\n=== Unfaithful Answer ===")
result = check_faithfulness(context, answer_bad)
print(f"Faithfulness score: {result.get('faithfulness_score')}")
for claim in result.get("claims", []):
    print(f"  - [{claim['supported']}] {claim['claim']}")
```

---

## Interview Tips for Coding Rounds

| Topic | Common Ask | Key Concept to Demonstrate |
|---|---|---|
| **Attention** | Implement from scratch | Scaled dot-product, causal mask, multi-head split |
| **BPE** | Implement training loop | Iterative pair merging, `</w>` boundary token |
| **LoRA** | Implement adapter | Low-rank decomposition, weight merging at inference |
| **KV Cache** | Explain + implement | Pre-allocation, O(1) per step vs O(n) recompute |
| **Evaluation** | LLM-as-judge pattern | Faithfulness = claims supported / total claims |
| **Quantization** | Explain INT4/NF4 | Calibration, GPTQ vs AWQ vs NF4 (for QLoRA) |
