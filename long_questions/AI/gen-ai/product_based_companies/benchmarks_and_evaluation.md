# GenAI Benchmarks & Evaluation — Product-Based Companies

Understanding how to **benchmark, evaluate, and compare LLMs** is essential for roles at Google, OpenAI, Anthropic, Cohere, and AI research teams. This file covers the most important benchmarks, evaluation methodologies, and how to build custom evaluation harnesses.

---

## 1. Core LLM Benchmarks — Know These Cold

### 📐 Reasoning & Knowledge

| Benchmark | Full Name | What It Tests | SOTA (2025) |
|---|---|---|---|
| **MMLU** | Massive Multitask Language Understanding | 57 academic subjects (STEM, law, ethics, medicine) — MCQ | GPT-4o / Gemini 1.5 Pro ~88% |
| **HellaSwag** | — | Commonsense NLI — complete a sentence logically | Near-saturated ~95%+ |
| **ARC-Challenge** | AI2 Reasoning Challenge | Grade school science MCQs; harder subset | ~90%+ for frontier models |
| **WinoGrande** | — | Commonsense pronoun resolution | ~85–90% |
| **TruthfulQA** | — | Does the model give *true* answers to tricky questions? (tests for false beliefs) | ~60–70% — models still fail |
| **BIG-Bench Hard (BBH)** | Beyond the Imitation Game - Hard | 23 challenging tasks that GPT-3 couldn't solve | GPT-4 ~85% |

### 💻 Code

| Benchmark | What It Tests | SOTA |
|---|---|---|
| **HumanEval** | 164 Python coding problems — execution-based (pass@1) | GPT-4o / Claude 3.5 ~90% |
| **MBPP** | Mostly Basic Python Problems — 374 problems | Similar to HumanEval |
| **SWE-Bench** | Real GitHub issue resolution — agent-based evaluation | Top agents ~45–50% (2025) |
| **LiveCodeBench** | Contamination-free coding benchmark (new problems monthly) | More honest signal than HumanEval |

### 🧮 Math

| Benchmark | What It Tests | SOTA |
|---|---|---|
| **GSM8K** | Grade school math word problems | Near-perfect ~95%+ with CoT |
| **MATH** | Competition math (AMC, AIME level) | GPT-4o ~72%, o1 ~94% |
| **AIME** | American Invitational Mathematics Exam | o1/DeepSeek-R1 ~60–80% |

### 🗣️ Instruction Following & Alignment

| Benchmark | What It Tests |
|---|---|
| **MT-Bench** | Multi-turn conversation quality (LLM-as-judge with GPT-4) |
| **AlpacaEval 2.0** | Head-to-head win rate vs GPT-4 Turbo (length-controlled) |
| **LMSYS Chatbot Arena** | Human preference ranking (Elo-style blind A/B voting) |
| **IFEval** | Instruction following evaluation — precise constraint satisfaction |

### 🔍 Long Context & RAG

| Benchmark | What It Tests |
|---|---|
| **RULER** | Long-context recall: needle-in-a-haystack, variable-length retrieval |
| **SCROLLS** | Summarization + QA on very long documents |
| **RAGAS** | RAG-specific: Faithfulness, Answer Relevancy, Context Recall/Precision |

---

## 2. Evaluation Frameworks (Production-Grade)

**Q1: What is the HELM evaluation framework? How is it different from MMLU?**

- **HELM (Holistic Evaluation of Language Models)** — Stanford CRFM.
- Tests models across **42 scenarios** and **7 metric types**: accuracy, calibration, robustness, fairness, bias, toxicity, efficiency.
- Unlike MMLU (just accuracy on one task type), HELM evaluates the *full profile* of a model.
- **Key insight:** A model that scores high on accuracy might score poorly on calibration (overconfident) or fairness (biased by demographic).

**Q2: What is the difference between pass@1, pass@10, and pass@100 for code benchmarks?**

```python
# pass@k metric (Chen et al. 2021):
# Generate k code samples per problem.
# pass@k = probability that at least 1 of the k samples passes all unit tests.

# Unbiased estimator (avoids sampling variance):
from math import comb

def pass_at_k(n: int, c: int, k: int) -> float:
    """
    n = total samples generated per problem
    c = number of correct samples (passed all tests)
    k = k in pass@k
    Returns: 1 - P(all k samples fail)
    """
    if n - c < k:
        return 1.0
    return 1.0 - comb(n - c, k) / comb(n, k)

# Example: Generate 200 samples, 40 pass → what is pass@1, pass@10, pass@100?
n, c = 200, 40
for k in [1, 10, 100]:
    score = pass_at_k(n, c, k)
    print(f"pass@{k}: {score:.3f}")
# pass@1: 0.200 (20% chance any single sample is correct)
# pass@10: 0.894
# pass@100: 1.000
```

**Why it matters:** pass@1 is what users see in practice. pass@100 tests the maximum capability. The gap shows reliability.

---

## 3. LLM-as-Judge (MT-Bench Pattern)

**Q3: How does LLM-as-Judge evaluation work? What are its biases?**

```python
from openai import OpenAI

client = OpenAI()

def llm_judge_score(question: str, answer: str, reference: str = None) -> dict:
    """
    G-Eval style LLM-as-judge for evaluating answer quality.
    Returns a score on a 1-10 scale with justification.
    """
    ref_section = f"\nReference Answer:\n{reference}" if reference else ""
    
    prompt = f"""You are an expert judge evaluating an AI assistant's response quality.
Rate the following response on a scale of 1-10 based on:
- Accuracy (is it factually correct?)
- Completeness (does it fully address the question?)
- Clarity (is it well-explained?)
- Conciseness (no unnecessary padding?){ref_section}

Question: {question}

Response to evaluate:
{answer}

Respond in JSON:
{{"score": <1-10>, "accuracy": <1-10>, "completeness": <1-10>, "reasoning": "<brief justification>"}}
"""
    response = client.chat.completions.create(
        model="gpt-4o",
        messages=[{"role": "user", "content": prompt}],
        response_format={"type": "json_object"},
        temperature=0
    )
    import json
    return json.loads(response.choices[0].message.content)


# Common biases in LLM-as-judge:
biases = {
    "Verbosity bias": "Longer answers rated higher regardless of quality",
    "Self-enhancement bias": "GPT-4 favors GPT-4 style outputs",
    "Position bias": "First answer in A/B comparison rated higher",
    "Sycophancy": "Agreeing with whatever the question implies is 'better'",
}
print("Known biases in LLM-as-judge:", biases)
```

**Mitigations:**
- Use reference answers (reduces verbosity bias).
- Swap answer positions and average (reduces position bias).
- Use multiple judges and take the median.
- Fine-tune the judge on human preference labels (Prometheus, JudgeLM).

---

## 4. Calibration & Confidence

**Q4: What is calibration in LLMs? Why does it matter?**

- A **perfectly calibrated** model: when it says "I'm 80% confident," it's correct exactly 80% of the time.
- **Overconfidence:** Says 90% confident but only right 60% of the time → dangerous in medical/legal/financial AI.
- **Underconfidence:** Says 50% but right 90% of the time → unhelpful, too hedged.

**Expected Calibration Error (ECE):**
```python
import numpy as np

def compute_ece(confidences: list, correctness: list, n_bins: int = 10) -> float:
    """
    Compute Expected Calibration Error.
    confidences: model's stated probability for each prediction
    correctness: 0 or 1 for each prediction
    """
    confidences = np.array(confidences)
    correctness = np.array(correctness)
    
    bin_boundaries = np.linspace(0, 1, n_bins + 1)
    ece = 0.0
    
    for i in range(n_bins):
        in_bin = (confidences > bin_boundaries[i]) & (confidences <= bin_boundaries[i + 1])
        if in_bin.sum() == 0:
            continue
        bin_accuracy = correctness[in_bin].mean()
        bin_confidence = confidences[in_bin].mean()
        bin_weight = in_bin.sum() / len(confidences)
        ece += bin_weight * abs(bin_accuracy - bin_confidence)
    
    return ece

# Example
confs = [0.9, 0.8, 0.7, 0.6, 0.9, 0.8, 0.7, 0.6]
correct = [1, 1, 0, 0, 1, 0, 1, 0]
print(f"ECE: {compute_ece(confs, correct):.3f}")  # Lower is better; 0 = perfect calibration
```

---

## 5. Contamination & Benchmark Leakage

**Q5: What is benchmark contamination? How do you detect and mitigate it?**

- **Contamination:** The test set questions from a benchmark appeared in the model's training data → inflated scores.
- **How common:** Multiple studies (e.g., Open LLM Leaderboard analysis) have found significant contamination in MMLU, HumanEval, and GSM8K for models trained on large web crawls.

**Detection methods:**
1. **N-gram overlap:** Check if test examples appear in training data (Common Crawl / C4 / The Pile membership inference).
2. **Canary tokens:** Insert unique synthetic examples into benchmarks; check if models regurgitate them.
3. **Performance cliff analysis:** If a model scores 90% on seen benchmark but 50% on a held-out equivalent, likely contaminated.
4. **LiveBench / LiveCodeBench:** Use temporally-gated questions (released after training cutoff) to avoid contamination.
5. **Eleuther AI's `lm-evaluation-harness`** includes decontamination utilities.

**Mitigation:**
- During pretraining: deduplicate against all known benchmark test splits.
- Evaluate on dynamic/rolling benchmarks that update monthly.
- Report both seen and unseen benchmark scores.

---

## 6. Evaluation Harness — Building One

**Q6: How do you build a custom evaluation harness for a domain-specific LLM?**

```python
import json
from openai import OpenAI
from dataclasses import dataclass, field
from typing import Callable

client = OpenAI()

@dataclass
class EvalExample:
    question: str
    ground_truth: str
    metadata: dict = field(default_factory=dict)

@dataclass 
class EvalResult:
    question: str
    prediction: str
    ground_truth: str
    scores: dict  # metric_name → score
    passed: bool

def exact_match(prediction: str, ground_truth: str) -> float:
    return float(prediction.strip().lower() == ground_truth.strip().lower())

def contains_key_fact(prediction: str, ground_truth: str) -> float:
    return float(ground_truth.strip().lower() in prediction.strip().lower())

def run_evaluation(
    dataset: list[EvalExample],
    model: str = "gpt-4o-mini",
    system_prompt: str = "You are a helpful assistant.",
    metrics: dict[str, Callable] = None
) -> dict:
    """Run evaluation on a golden dataset and return aggregate scores."""
    if metrics is None:
        metrics = {
            "exact_match": exact_match,
            "contains_key_fact": contains_key_fact
        }
    
    results = []
    for example in dataset:
        # Generate prediction
        response = client.chat.completions.create(
            model=model,
            messages=[
                {"role": "system", "content": system_prompt},
                {"role": "user", "content": example.question}
            ],
            temperature=0
        )
        prediction = response.choices[0].message.content
        
        # Score on all metrics
        scores = {
            name: fn(prediction, example.ground_truth) 
            for name, fn in metrics.items()
        }
        passed = all(s >= 0.5 for s in scores.values())
        
        results.append(EvalResult(
            question=example.question,
            prediction=prediction,
            ground_truth=example.ground_truth,
            scores=scores,
            passed=passed
        ))
    
    # Aggregate
    n = len(results)
    aggregate = {
        name: sum(r.scores[name] for r in results) / n
        for name in metrics
    }
    aggregate["pass_rate"] = sum(r.passed for r in results) / n
    aggregate["n_examples"] = n
    
    return {"aggregate_scores": aggregate, "results": results}


# --- Golden dataset example ---
golden_dataset = [
    EvalExample(
        question="What is the capital of France?",
        ground_truth="Paris"
    ),
    EvalExample(
        question="What does RAG stand for?",
        ground_truth="Retrieval-Augmented Generation"
    ),
]

eval_output = run_evaluation(golden_dataset)
print(json.dumps(eval_output["aggregate_scores"], indent=2))
```

---

## 7. Benchmark Quick Reference Card

| Category | Benchmark | Key Metric | Why It Matters in Interviews |
|---|---|---|---|
| Knowledge | MMLU | Accuracy (MCQ) | Universal model capability baseline |
| Reasoning | BBH | Accuracy | Tests tasks GPT-3 couldn't do |
| Math | MATH, AIME | Accuracy | Differentiates o1/R1-class reasoning models |
| Code | HumanEval | pass@1 | Coding assistant quality signal |
| Code (no contamination) | LiveCodeBench | pass@1 | More honest signal |
| Real-world code | SWE-Bench | % resolved | Agent capability for software engineering |
| Alignment | Chatbot Arena | Elo rating | Human preference ground truth |
| Instruction following | IFEval | Strict accuracy | Tests precise constraint adherence |
| RAG | RAGAS | Faithfulness, AR | Production RAG system evaluation |
| Long context | RULER | Recall@depth | Context window quality |
| Safety | TruthfulQA | Truthfulness % | Model does not propagate false beliefs |
| Calibration | ECE | Lower is better | Confidence reliability for high-stakes |
