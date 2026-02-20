# Freshers Level GenAI Interview Questions

## 🟢 **1–20: Fundamentals & Basic Concepts**

### 1. What is Generative AI and how does it differ from traditional AI?
"Generative AI focuses on creating **new content** rather than just analyzing existing data. While traditional AI (like a classifier) might look at a picture and tell me 'this is a cat,' Generative AI can take a prompt like 'a cat in space' and actually draw it.

Traditional AI is largely analytical or predictive—it finds patterns to make decisions. Generative AI learns the underlying distribution of the training data and samples from it to produce novel outputs, such as text, images, or code. 

I think of traditional AI as a highly skilled critic, whereas Generative AI is like an artist."

#### Indepth
Generative AI models estimate the joint probability $P(X, Y)$ to learn the distribution of the data, whereas discriminative (traditional AI) models learn the conditional probability $P(Y|X)$ to map inputs to labels. This fundamental difference allows generative models to create synthetic data points that mimic the original dataset's distribution.

---

### 2. What are the main types of generative models?
"The four main architectures I encounter are **Generative Adversarial Networks (GANs)**, **Variational Autoencoders (VAEs)**, **Autoregressive Models (like Transformers)**, and **Diffusion Models**. 

LLMs like GPT use Transformers to predict the next token. Midjourney or Stable Diffusion rely heavily on Diffusion models. GANs are famous for generating highly realistic faces or deepfakes, and VAEs are great when we need a continuous, structured latent space for things like interpolating between images."

#### Indepth
Each architecture has a trade-off. GANs generate sharp images but suffer from training instability (mode collapse). VAEs are stable but often produce blurry images. Diffusion models provide state-of-the-art image quality but are computationally expensive during inference because of the iterative denoising process.

---

### 3. What are some common use cases of Generative AI?
"The use cases are incredibly broad. 
In text, it's used for **drafting emails, summarizing long documents, and generating code** (like GitHub Copilot). 
In visual arts, it’s used for **concept art, stock image generation, and varying ad creatives**. 
In audio, it can do **text-to-speech** with human-like intonation or generate background music. 
It’s also revolutionizing science, for example by **predicting protein structures** or accelerating drug discovery."

#### Indepth
Beyond creative tasks, GenAI is heavily utilized in synthetic data generation. If an enterprise lacks diverse training data for a traditional ML task (like fraud detection), GenAI can create synthetic, anonymized profiles that maintain statistical properties without violating PII constraints.

---

### 4. What role does randomness play in generative models?
"Randomness is the spark that allows GenAI to be 'creative.' If models were entirely deterministic, asking the same question would always yield the exact same response or image. 

In LLMs, randomness is controlled by parameters like **Temperature**. A low temperature makes the model pick only the most probable next words (good for factual answers). A high temperature allows it to pick less probable words, resulting in more varied, creative, or surprising text. Similarly, image generators often start from a seed of random noise."

#### Indepth
In Diffusion models, the reverse process starts with a matrix of pure Gaussian noise. The model uses the prompt condition to iteratively denoise it. The initial random noise dictates the final composition of the image, which is why saving the 'seed' allows you to reproduce the exact same generation.

---

### 5. What is a Transformer model, and why is it suited for GenAI?
"The Transformer is a neural network architecture introduced by Google in 2017. Its superpower is the **Self-Attention Mechanism**. 

Before Transformers, we used RNNs or LSTMs, which read text sequentially. Transformers process the entire sequence at once (in parallel). It looks at every word in a sentence and calculates how relevant it is to every other word. This allows the model to understand context and long-range dependencies perfectly, which is why it excels at generating coherent paragraphs of text."

#### Indepth
Transformers eliminated the bottleneck of sequential computation, allowing massive parallelization on modern GPUs. This permitted the training of massive models (LLMs) on trillions of tokens. The standard architecture has an Encoder (good for understanding/classification like BERT) and a Decoder (good for generation like GPT).

---

### 6. What is the difference between GPT and BERT?
"Both are Transformers, but they excel at different things because of their architecture.

**GPT (Generative Pre-trained Transformer)** uses only the **Decoder** part of the transformer. It reads text strictly from left to right and is meant to predict the next word. It’s perfect for generating text.

**BERT (Bidirectional Encoder Representations from Transformers)** uses only the **Encoder**. It looks at text bidirectionally (left and right simultaneously). It’s not meant for generation; I would use it for understanding tasks like sentiment analysis, search ranking, or question answering."

#### Indepth
BERT is trained using Masked Language Modeling (MLM)—it randomly hides 15% of the words in a sentence and tries to guess them based on the surrounding context. GPT uses Causal Language Modeling—it predicts the $Nth$ word using only words 1 to $N-1$. 

---

### 7. How do LLMs generate coherent text?
"LLMs generate text exactly one 'token' (a word or piece of a word) at a time. It’s fundamentally an incredibly advanced autocomplete. 

When I give it a prompt, it calculates the mathematical probability of what the next token should be based on everything it has seen before. It picks a token, appends it to the prompt, feeds the whole thing back into itself, and guesses the next one. It repeats this loop until it generates a special 'stop token'."

#### Indepth
This process is called **Autoregressive Generation**. Because the model has to run a full forward pass through billions of parameters for every single token it generates, inference is heavily memory-bandwidth bound (memory wall). This is why techniques like Key-Value (KV) caching are essential to speed up text generation.

---

### 8. What are prompt engineering techniques?
"Prompt engineering is the art of giving an LLM the right instructions to get the best output. 

Some basic techniques I use include **Zero-shot prompting** (just asking a question), **Few-shot prompting** (showing it 2 or 3 examples of what I want first), and **Role-playing** ('Act as a senior software engineer'). 

For complex tasks, I use **Chain of Thought (CoT)**, where I ask the model to 'think step by step.' This forces the model to generate intermediate reasoning tokens, significantly improving its logic and math capabilities."

#### Indepth
Prompt engineering works because LLMs are patterned sequential predictors. By forcing the model to output steps (Chain of Thought), you give it more computational 'scratchpad' space. It can look at its own intermediate logical steps to predict the correct final answer, rather than trying to jump to the final answer in one direct probability hop.

---

### 9. Explain the concept of context window and token limits.
"The context window is the model's 'short-term memory' for a single conversation. It is measured in tokens.

If a model has a 4,000 token context window, it can 'see' about 3,000 words at once. If our conversation gets longer than that, the model literally 'forgets' the beginning of the chat because those tokens fall out of the window. When I build apps, I often have to write code to summarize older messages to keep the important context within this hard limit."

#### Indepth
The context limit exists because the Self-Attention mechanism of standard transformers scales quadratically $O(N^2)$ with the sequence length. A 100k context window requires exponentially more GPU VRAM than a 4k window. Newer models use techniques like Ring Attention or Sparse Attention to achieve 1M+ token windows.

---

### 10. How does text-to-image generation work?
"Text-to-image models usually involve two main parts: a **Text Encoder** and an **Image Generator**.

When I type 'a cute dog,' the Text Encoder (like CLIP) translates those words into a mathematical vector (numbers) that represents the 'meaning' of that phrase. Then, the Image Generator (usually a Diffusion model) starts with a canvas of random static. It looks at the text vector and iteratively removes the static step-by-step, shaping the pixels until an image matching the text emerges."

#### Indepth
CLIP (Contrastive Language-Image Pre-training) is crucial here. It was trained to map images and their text captions into the same latent space. Because they live in the same space, the diffusion model can use the text embedding as a direct mathematical guide to condition the denoising steps.

---

### 11. What is super-resolution in Generative AI?
"Super-resolution is the process of taking a low-resolution, blurry, or pixelated image and using AI to upscale it into a sharp, high-definition image. 

Unlike older software that just stretched the pixels and blurred the edges (interpolation), GenAI actually 'hallucinates' or predicts the missing details based on what it learned during training. If it sees a blurry eye, it adds realistic eyelashes and reflections because it knows what an eye *should* look like."

#### Indepth
Super-resolution is heavily used in Deep Learning Super Sampling (DLSS) in modern video games, where the game renders at 1080p and a generative model upscales it to 4K in real-time, saving immense GPU compute power.

---

### 12. Name top industries using GenAI today.
"GenAI is touching almost everything, but the biggest adopters are:
1. **Software Engineering:** Tools like Copilot are accelerating coding.
2. **Customer Service:** Advanced chatbots handle complex support workflows, not just FAQs.
3. **Marketing & Content:** Automatically drafting blogs, ad copy, and generating ad visuals.
4. **Healthcare:** Analyzing patient notes, generating medical reports, and synthesizing new drug compounds.
5. **Legal:** Summarizing massive legal briefs and drafting standard contracts."

#### Indepth
The adoption barrier is mainly data privacy and hallucination risks. Industries like Finance and Healthcare are adopting GenAI slower for client-facing applications out of regulatory fear, but are heavily utilizing it internally for developer productivity and data summarization.

---

### 13. How do chatbots use GenAI?
"Modern chatbots use GenAI to actually parse intent and formulate dynamic responses. 

Older chatbots were simple decision trees: 'If user says X, print Y.' They failed if the phrasing was slightly off. GenAI chatbots (like ChatGPT) understand the semantic meaning of the user's query. They can handle spelling errors, complex multi-part questions, and can maintain the context of the conversation over time, generating a completely unique, human-like response."

#### Indepth
Enterprise chatbots rarely rely purely on the GenAI model's internal memory (which hallucinates). Instead, they use RAG (Retrieval-Augmented Generation). The chatbot intercepts the query, searches a private database for facts, and then instructs the LLM: "Answer the user using strictly these retrieved facts."

---

### 14. Explain AI-assisted code generation.
"AI code generation acts like an incredibly smart pair-programmer. 

I can write a comment like `// function to parse a JSON file and return active users` and the AI will auto-complete the entire function. It saves me from looking up boilerplate syntax or library documentation. It’s powered by LLMs trained specifically on massive datasets of public code (like GitHub), so it understands the syntax and logic of almost every programming language."

#### Indepth
Code models (like Codex or StarCoder) are often trained with a Fill-In-The-Middle (FIM) objective. Instead of just predicting the end of a file, they are trained to see the code *above* your cursor and *below* your cursor, and accurately predict the missing chunk in the middle.

---

### 15. What are virtual avatars or AI influencers?
"Virtual avatars are completely AI-generated personas. They have computer-generated faces and bodies, and their scripts or captions are written by LLMs. 

Brands use them for marketing because they don't age, they don't get into scandals, and they can be customized to appeal perfectly to a target demographic. Technologically, they combine image generation (for their photos), voice synthesis (for videos), and text generation (for interacting with fans)."

#### Indepth
Creating consistent virtual avatars is technically challenging because diffusion models struggle with character consistency across different angles and scenes. Creators use advanced control structures like ControlNet and LoRA fine-tuning to force the model to render the exact same face structure and clothing repeatedly.

---

### 16. How does OpenAI's API work for generative tasks?
"Using the OpenAI API involves sending an HTTP request containing my prompt and an API key. 

I send a JSON payload with an array of 'messages' representing the conversation history, and a 'model' parameter (like `gpt-4o`). The OpenAI servers run the massive model in the cloud and return a JSON response containing the generated text. 

I don't have to worry about GPU hosting, scaling, or model updates; it’s strictly a 'pay-per-token' service."

#### Indepth
For production apps, it's vital to handle the API dynamically using Server-Sent Events (SSE). Instead of waiting 10 seconds for the whole response, I stream the tokens back one-by-one to create the 'typing' effect in the UI, drastically improving perceived performance for the end-user.

---

### 17. What are the ethical concerns with generative AI?
"The ethical concerns are massive. 
First is **Copyright**: models are trained on billions of images and texts without compensating the original authors. 
Second is **Deepfakes**: generating fake audio or video to impersonate politicians or individuals for fraud or blackmail. 
Third is **Bias**: models learn from the internet, which means they easily replicate racist, sexist, or cultural biases if not carefully filtered."

#### Indepth
There is a massive legal gray area around 'Fair Use'. Currently, courts are debating whether training a neural network on copyrighted data constitutes transformative fair use or copyright infringement. Additionally, the proliferation of AI-generated content pollutes the internet, leading to "model collapse" when future AIs train on the synthetic output of current AIs.

---

### 18. How can GenAI impact jobs and the economy?
"GenAI is a huge productivity multiplier, acting as a co-pilot for knowledge workers. 

It will automate highly repetitive cognitive tasks—like writing basic boilerplate code, translating texts, summarizing data, or generating stock images. While it might displace entry-level roles tied strictly to content generation, history shows new technologies shift employment rather than destroy it. Prompt engineering, AI safety auditing, and AI system architecture are entirely new job categories."

#### Indepth
Economic reports suggest GenAI will boost global GDP by trillions. The shift moves human labor from 'Creation' to 'Curation and Verification'. A software logic error by an AI still needs a human engineer to spot and fix it; an AI legal draft still requires a lawyer's signature and liability assumption.

---

### 19. What are “hallucinations” in LLMs?
"A hallucination is when an AI confidently presents information that is factually incorrect or completely entirely made up.

If I ask an LLM about a niche, non-existent historical event, it might invent dates, names, and a convincing narrative. This happens because LLMs don't have a 'database of facts.' They just predict which words probabilistically look good together based on patterns. They don’t inherently know true from false."

#### Indepth
Hallucinations are the biggest blocker for enterprise adoption. You can't have a customer service bot promising free laptops. We combat this using RAG (grounding the model in factual documents), lowering the extraction temperature to 0, and designing strict system prompts that command: "If the information is not in the text, reply 'I don't know'."

---

### 20. Discuss bias in generative models.
"Bias occurs when a model's outputs reflect unfair prejudices. For example, if I ask an image generator for a 'CEO,' it might only generate old white men, and if I ask for a 'nurse,' it might only generate young women.

This happens because the model is trained on historical data from the internet, which objectively contains these societal biases. The model isn't maliciously biased; it’s mathematically lazy, gravitating toward the largest statistical clusters in its dataset."

#### Indepth
Mitigating bias requires intervention at multiple steps: curating a mathematically balanced training dataset (hard), utilizing RLHF to penalize biased outputs during human alignment (expensive), and secretly injecting diverse modifiers into user prompts behind the scenes (e.g., quietly changing "draw a doctor" to "draw a doctor of hispanic descent").
