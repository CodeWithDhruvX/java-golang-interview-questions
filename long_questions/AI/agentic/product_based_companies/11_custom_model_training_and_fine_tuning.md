# Custom Model Training and Fine-Tuning (Product-Based)

While prompt engineering and RAG solve many problems, top product companies often fine-tune open-weights models (like Llama 3, Mistral) for specific agentic tasks to reduce costs, improve speed, or handle highly specialized domains.

## 1. When would you choose to Fine-Tune a model instead of using RAG or Prompt Engineering?

**Answer:**
We follow a hierarchy of complexity. We only move to fine-tuning when prompt engineering and RAG fail to meet requirements:
*   **Format and Tone Specialization:** When the model needs to output complex, specific structures consistently (e.g., proprietary JSON schemas, specific programming languages, or a highly specific brand voice) that it struggles to grasp even with few-shot prompting.
*   **Task Specialization (Tool Calling):** To train a smaller, cheaper open-source model (like an 8B parameter model) to be extremely good at identifying when to use specific internal APIs and extracting the correct parameters. We can fine-tune it to match GPT-4's performance *on our specific tools*.
*   **Domain Expertise (Vocabulary):** When dealing with highly specialized jargon (e.g., advanced medical research or proprietary legal terminology) where the base model lacks the foundational vocabulary, making RAG inefficient.
*   **Latency and Cost Optimization:** Once a prompt is fully optimized for a task, it might be very long. Fine-tuning bakes those instructions into the model weights, allowing us to use much shorter prompts (saving token costs) and smaller models (improving latency and reducing inference costs).
*   **Data Privacy:** When regulatory or privacy concerns dictate that data cannot be sent to commercial APIs, forcing us to run custom models within our VPC.

*Note: RAG and Fine-tuning are not mutually exclusive. We often fine-tune a model to be better at using RAG (e.g., better at understanding context or better at knowing when to search).*

## 2. Explain the concept of PEFT (Parameter-Efficient Fine-Tuning) and LoRA. Why are they important?

**Answer:**
Full fine-tuning (updating all billions of weights in an LLM) is computationally prohibitive, requiring massive GPU clusters.
*   **PEFT (Parameter-Efficient Fine-Tuning):** Techniques that fine-tune only a small subset of parameters or add a small number of new parameters, while keeping the pre-trained base model weights frozen.
*   **LoRA (Low-Rank Adaptation):** This is the most popular PEFT method. Instead of updating the massive weight matrices directly, LoRA injects small, trainable "rank decomposition" matrices into the model's layers (usually the attention blocks).
*   **Why it's important:**
    *   **Cost/Resource Reduction:** LoRA allows us to fine-tune large models (like a 70B parameter model) on a single, consumer-grade GPU or a small cloud instance, making it highly accessible.
    *   **Storage Efficiency:** The resulting "adapter" (the LoRA weights) is very small (often just a few megabytes), compared to the tens of gigabytes for the base model.
    *   **Swappability:** We can keep one large base model loaded in VRAM and dynamically swap out different LoRA adapters for different tasks or customers on the fly without reloading the massive base model.

## 3. How do you prepare a dataset for fine-tuning a model for function calling/tool use?

**Answer:**
Data quality is the most critical factor in fine-tuning. For tool calling, the dataset must teach the model *when* to use a tool, *how* to format the request, and *how* to handle the tool's response.
1.  **Format Selection:** We choose a standardized format for representing tool calls (e.g., OpenAI's function calling JSON format, or a specialized XML schema).
2.  **Data Generation:** We need thousands of high-quality examples. We often use a "Teacher Model" (like GPT-4) to generate this synthetic data. We provide the Teacher with the tool definitions and various user scenarios, and ask it to generate the proper interaction.
3.  **Constructing the Conversation format:** The dataset must be formatted in a chat structure (e.g., ChatML), ensuring roles are clear:
    *   `System`: Defines the available tools and instructions.
    *   `User`: The request.
    *   `Assistant`: The generated tool call (JSON).
    *   `Tool`: The simulated output of the tool execution.
    *   `Assistant`: The final answer based on the tool output.
4.  **Negative Examples:** Crucially, we include examples where the user asks something irrelevant, and the model must learn *not* to call any tools and just respond normally.
5.  **Validation:** We rigorously clean the data, ensuring the JSON tool calls are perfectly valid and the logic is sound. Junk data leads to a broken model.

## 4. How do you evaluate the performance of a fine-tuned model?

**Answer:**
Evaluation cannot rely on standard NLP metrics (like BLEU or ROUGE). We need task-specific evaluation:
*   **Holdout Validation Set:** Evaluating loss and accuracy during training on a validation dataset that the model hasn't seen. This checks for overfitting.
*   **LLM-as-a-Judge:** We use a high-capability model (GPT-4) to evaluate the fine-tuned model's outputs based on specific criteria (e.g., rubric scoring for tone, accuracy, formatting).
*   **Automated Tool Evaluation:** For tool calling, we run automated tests:
    1.  Provide the model a prompt.
    2.  Check if it outputs a syntactically valid tool call.
    3.  Check if it selected the *correct* tool.
    4.  Check if the arguments provided to the tool are accurate.
*   **A/B Testing in Shadow Mode:** Before replacing the production model, run the fine-tuned model in parallel ("shadow mode"), logging its responses alongside the current production model, and compare the results for anomalies.
*   **Human Eval:** For subjective things like tone and brand voice, human review on a sample of outputs is unmatched.

## 5. What is DPO (Direct Preference Optimization) or RLHF (Reinforcement Learning from Human Feedback), and when would you use them?

**Answer:**
Supervised Fine-Tuning (SFT) teaches a model *how* to do something (the format, the tone). RLHF and DPO teach the model what is *preferred* or *helpful* behavior.
*   **RLHF:** Involves training a "reward model" based on human rankings of different LLM outputs, then using reinforcement learning (like PPO) to optimize the LLM to maximize that reward. It's complex and unstable.
*   **DPO (Direct Preference Optimization):** A more modern, stable approach. It bypasses the separate reward model phase. You provide a dataset of preference pairs: (Prompt, Chosen Response, Rejected Response). DPO mathematically optimizes the model directly to increase the probability of the chosen response and decrease the probability of the rejected one.
*   **When to use:** We use SFT first to get the formatting right. We use DPO when we need to align the model's behavior with human preferences—for example, teaching an agent to be more concise, less apologetic, or less prone to generating "I am an AI..." disclaimers when it shouldn't. It's about stylistic and behavioral alignment.
