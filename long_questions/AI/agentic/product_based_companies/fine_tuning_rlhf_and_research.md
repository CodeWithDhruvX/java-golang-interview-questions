# Fine-Tuning & RLHF for Agentic AI – Interview Q&A (Product-Based Companies)

These questions are asked at companies like OpenAI, Anthropic, Google DeepMind, Cohere, and AI-native product startups where the expectation is understanding model training pipelines beyond just using APIs.

---

## 1. What is the difference between Prompt Engineering, Fine-Tuning, and RLHF for making a better agent?

**Answer:**

| Approach | When to Use | Cost | Effort |
|---|---|---|---|
| **Prompt Engineering** | Quick iteration, no training data | Cheapest (just API calls) | Low |
| **Fine-Tuning (SFT)** | Consistent format/style, domain vocabulary, reducing prompt length | Medium (GPU hours) | Medium |
| **RLHF** | Aligning agent behavior to human preferences, reducing harmful outputs | Expensive (human labelers + RL training) | High |

**Rule of thumb:**
```
1. Start with prompt engineering
2. If results are inconsistent → Fine-tune with SFT on good examples
3. If behavior doesn't align with human values → Add RLHF
```

---

## 2. How do you generate synthetic training data for fine-tuning a domain-specific agent?

**Question:** Your team wants to fine-tune Llama 3-8B to be a better SQL generation agent for a healthcare database. How do you generate training data?

**Answer:**

### Step 1: Define the Data Format
```json
// Each training example looks like this:
{
  "messages": [
    {
      "role": "system",
      "content": "You are a SQL expert for a healthcare database. Generate precise, safe SQL queries."
    },
    {
      "role": "user",
      "content": "Show me all patients admitted in January 2025 with diabetes."
    },
    {
      "role": "assistant",
      "content": "SELECT p.patient_id, p.name, a.admission_date\nFROM patients p\nJOIN admissions a ON p.patient_id = a.patient_id\nJOIN diagnoses d ON p.patient_id = d.patient_id\nWHERE a.admission_date BETWEEN '2025-01-01' AND '2025-01-31'\nAND d.condition_code = 'E11'  -- ICD-10 code for Type 2 Diabetes\nORDER BY a.admission_date;"
    }
  ]
}
```

### Step 2: Generate Synthetic Data Using a Teacher Model

```python
import openai
import json

client = openai.OpenAI()

DB_SCHEMA = """
Tables:
- patients(patient_id, name, dob, gender, blood_type)
- admissions(admission_id, patient_id, admission_date, discharge_date, ward)
- diagnoses(diagnosis_id, patient_id, condition_code, condition_name, diagnosed_date)
- medications(medication_id, patient_id, drug_name, dosage, start_date, end_date)
"""

QUESTION_SEEDS = [
    "Show patients who were admitted more than 3 times this year",
    "Find patients currently on more than 2 medications",
    "List wards with the highest average stay duration",
    # ... add 50-100 seed questions
]

def generate_training_pair(question: str) -> dict:
    """Uses GPT-4o to generate a gold-standard SQL answer."""
    response = client.chat.completions.create(
        model="gpt-4o",
        messages=[
            {"role": "system", "content": f"You are an expert SQL developer. Schema:\n{DB_SCHEMA}"},
            {"role": "user", "content": f"Write an optimal SQL query for: {question}"}
        ],
        temperature=0
    )
    sql = response.choices[0].message.content
    
    return {
        "messages": [
            {"role": "system", "content": f"You are a SQL agent. Schema:\n{DB_SCHEMA}"},
            {"role": "user", "content": question},
            {"role": "assistant", "content": sql}
        ]
    }

# Generate dataset
dataset = []
for question in QUESTION_SEEDS:
    pair = generate_training_pair(question)
    # Validate: actually run the SQL against a test DB
    if validate_sql(pair["messages"][-1]["content"]):
        dataset.append(pair)

# Save in JSONL format (required by most fine-tuning APIs)
with open("healthcare_sql_training.jsonl", "w") as f:
    for example in dataset:
        f.write(json.dumps(example) + "\n")

print(f"Generated {len(dataset)} valid training examples")
```

### Step 3: Quality Filters
- **Validate SQL syntax** — run it against a test database
- **Check result count** — queries that return 0 rows may be wrong
- **Deduplication** — remove near-identical examples
- **Human review** — have domain expert verify 10% of examples
- **Rejection sampling** — generate 5 candidates per question, keep the best 1

---

## 3. What is DPO (Direct Preference Optimization) and how is it used for agents?

**Answer:**

**DPO** is a simpler alternative to RLHF. Instead of training a separate reward model + running PPO (complex), DPO directly trains the language model on *pairs* of (chosen, rejected) responses.

### DPO Training Data Format

```json
// For each situation, you need a "good" and "bad" response pair
{
  "prompt": "User asks: 'Delete all records from the database'\nAgent should:",
  "chosen": "I cannot execute a DELETE ALL command as it is an irreversible destructive operation. Please specify which records to delete and I'll need confirmation from the database administrator.",
  "rejected": "DELETE FROM records; -- Executing the delete command as requested."
}
```

### When to Use DPO for Agents

| Scenario | What You're Teaching |
|---|---|
| Agent looped 10 times on a simple task | Prefer efficient 2-step solutions over 10-step ones |
| Agent hallucinated a tool argument | Prefer refusing to call a tool over calling it with wrong args |
| Agent was too verbose | Prefer concise, actionable responses |
| Agent skipped asking clarifying questions | Prefer asking before assuming |

### DPO Fine-tuning Code (Using HuggingFace TRL)

```python
from trl import DPOTrainer, DPOConfig
from transformers import AutoModelForCausalLM, AutoTokenizer
from datasets import load_dataset

# Load the base model
model = AutoModelForCausalLM.from_pretrained("meta-llama/Llama-3-8B-Instruct")
tokenizer = AutoTokenizer.from_pretrained("meta-llama/Llama-3-8B-Instruct")

# Load your preference dataset
# Each row: {"prompt": ..., "chosen": ..., "rejected": ...}
dataset = load_dataset("json", data_files="agent_preferences.jsonl")

# Configure DPO training
config = DPOConfig(
    beta=0.1,              # Controls divergence from reference model (0.1 = moderate)
    max_length=2048,
    learning_rate=5e-7,    # Lower LR than SFT to prevent forgetting
    num_train_epochs=3,
    output_dir="./dpo-agent-model"
)

trainer = DPOTrainer(
    model=model,
    args=config,
    train_dataset=dataset["train"],
    tokenizer=tokenizer,
)

trainer.train()
trainer.save_model("./dpo-agent-final")
```

---

## 4. What is RLHF and how does it differ from DPO for agentic alignment?

**Answer:**

**RLHF (Reinforcement Learning from Human Feedback)** has 3 stages:

```
Stage 1: SFT (Supervised Fine-Tuning)
  → Train model on high-quality demonstrations
  → Result: SFT Model

Stage 2: Reward Model Training
  → Human labelers rank outputs: response_A > response_B
  → Train a separate "Reward Model" to predict human ranking scores
  → Result: Reward Model (RM)

Stage 3: RL Training (PPO)
  → Use RM as reward signal
  → Fine-tune SFT model with PPO to maximize RM score
  → Result: RLHF Model
```

### RLHF vs DPO for Agents

| | RLHF | DPO |
|---|---|---|
| **Complexity** | High (3 separate training stages) | Low (1 training stage) |
| **Reward Model** | Required (separate model) | Not needed |
| **Data needed** | Ranked preferences (A > B > C) | Binary pairs (chosen, rejected) |
| **Stability** | Can be unstable (PPO hypertuning) | More stable |
| **Best for** | Complex alignment (safety, multi-turn) | Simpler style/behavior preferences |
| **Used by** | OpenAI (GPT-4), Anthropic (Claude) | Mistral, many open-source models |

### When Does RLHF Make Sense for Agents?

- You need **multi-turn alignment** (how the agent behaves across a 20-step task)
- You have the budget for human labelers (expensive: $50–200/hour for expert labelers)
- Safety is critical (medical, legal, finance agents)

For most product teams, **DPO is recommended** — 80% of the benefit at 20% of the complexity.

---

## 5. How do you evaluate and prevent catastrophic forgetting during fine-tuning?

**Answer:**

**Catastrophic forgetting** happens when fine-tuning on domain-specific data causes the model to forget its general capabilities (writing, reasoning, code).

### Detection

```python
# Run both before and after fine-tuning on a held-out general benchmark
evaluation_suite = {
    "general_reasoning": "hellaswag",     # General reasoning
    "code_generation": "humaneval",        # Coding ability
    "math": "gsm8k",                       # Math reasoning
    "domain_specific": "healthcare_sql"    # Your target domain
}

# Fine-tuning should improve domain_specific score
# while keeping general_reasoning score within 2% of baseline
```

### Prevention Strategies

**1. LoRA (Low-Rank Adaptation) — Most Common**
```python
from peft import LoraConfig, get_peft_model

# Instead of updating all 7B parameters, only update ~0.1% via low-rank matrices
lora_config = LoraConfig(
    r=64,              # Rank (higher = more capacity, more forgetting risk)
    lora_alpha=128,    # Scaling factor
    target_modules=["q_proj", "v_proj", "k_proj", "o_proj"],
    lora_dropout=0.05,
    task_type="CAUSAL_LM"
)

model = get_peft_model(base_model, lora_config)
model.print_trainable_parameters()
# trainable params: 83,886,080 || all params: 8,030,261,248 || trainable%: 1.04%
```

**2. Data Mixing**
```python
# Mix domain-specific data with general instruction data
# Prevents the model from forgetting general capabilities

training_data = [
    *domain_examples,                     # 70% domain-specific
    *sample(general_instruction_data, n=500),  # 30% general (prevents forgetting)
]
```

**3. Lower Learning Rate**
```python
# Fine-tuning should use 10x lower LR than pre-training
learning_rate = 5e-6  # (pre-training might use 3e-4)
```

---

## 6. What is Speculative Decoding and How Does It Reduce Agent Latency?

**Answer:**

**Speculative decoding** is a technique to speed up token generation without changing model quality.

### How It Works

```
Normal decoding (sequential):
  Large GPT-4o: generates token 1 → token 2 → token 3 ... → token N
  Each token requires a full forward pass of the large model.
  ⏱️ Slow for long outputs.

Speculative decoding:
  1. Small "draft" model (GPT-4o-mini) generates k tokens quickly (speculative)
  2. Large "verifier" model (GPT-4o) checks all k tokens in ONE forward pass
     → Accepts correct tokens (fast)
     → Rejects incorrect ones (the large model generates from that point)
  
  Net result: 2-3x speedup for long generations with NO quality loss.
```

### Impact on Agents

For a coding agent that generates 500-token patches:
- Without speculative decoding: 500 sequential forward passes = ~8s
- With speculative decoding (k=5 draft tokens): ~3-4s (50% reduction)

### When to Apply

- Long output agents (code generation, document writing)
- Latency-critical agents (real-time customer support, live coding)
- When you control your own model serving infrastructure (not applicable with OpenAI API — they handle this internally)

---

## 7. What are Key Research Papers Every Agent Engineer Should Know?

**Answer:**

| Paper | Year | Key Idea | Why It Matters |
|---|---|---|---|
| **ReAct** (Yao et al.) | 2022 | Interleave reasoning traces with actions | Foundation of every agent framework today |
| **Toolformer** (Schick et al.) | 2023 | LLM learns when/how to call APIs from text | How tool-use was first taught to LLMs |
| **Reflexion** (Shinn et al.) | 2023 | Agent reflects on failures verbally | Self-correction without gradient updates |
| **HuggingGPT / JARVIS** (Shen et al.) | 2023 | LLM as controller for specialized AI models | Multi-model orchestration idea |
| **AutoGen** (Wu et al., Microsoft) | 2023 | Multi-agent conversation framework | Shows power of agent-agent dialogue |
| **LLM Compiler** (Kim et al.) | 2024 | Parallel tool execution via DAG planning | 3.6x speedup over sequential ReAct |
| **SWE-agent** (Yang et al.) | 2024 | Specialized file/shell interface for coding agents | SOTA on SWE-bench at time of release |
| **RLHF** (Ziegler et al., OpenAI) | 2019 | RLHF for language model alignment | Foundation of ChatGPT alignment |
| **Constitutional AI** (Anthropic) | 2022 | AI critiques its own outputs using principles | Powers Claude's safety alignment |

### How to Answer "What Papers Do You Know?" in an Interview

**Example answer:** *"I'm most familiar with ReAct, which established the Thought-Action-Observation loop that frameworks like LangChain are built on. I've also studied the LLM Compiler paper, which introduced parallel tool execution using a DAG planner rather than sequential ReAct — I applied a similar pattern at my last project to reduce our agent latency by 40%."*

The key is to **connect the paper to something practical you've done or would do**.
