# Intermediate Level GenAI Interview Questions

## 🟡 **21–40: Core Architectures & Workflows**

### 21. Explain the difference between discriminative and generative models.
"Discriminative models draw boundaries. If I give it a picture, it discriminates between 'cat' and 'dog' by finding the dividing line between those classes in the data. Mathematically, it models $P(Y|X)$—the probability of label Y given data X.

Generative models, however, learn *how the data is distributed in space*. They don't just classify; they learn exactly what makes a 'cat' a cat. They model the joint probability $P(X, Y)$, allowing me to sample from it to generate an entirely new, unseen 'cat' that fits the learned mathematical distribution."

#### Indepth
Generative models are inherently much harder to train than discriminative models because they have to capture the entire feature space distribution, not just the decision boundary. If a generative model misses the variance of a dataset, all its output will look identically average (mode collapse), whereas a discriminative model might still classify perfectly fine with only the boundary.

---

### 22. How do GANs differ from VAEs?
"Both generate new data, but their architectures are entirely different. 

A **Variational Autoencoder (VAE)** consists of an Encoder that compresses an image into a structured latent space (a 'map' of numbers), and a Decoder that tries to perfectly rebuild the image from that map. It forces the map to be smooth, so I can pick a point halfway between a 'smile' map and a 'frown' map, and get a realistic intermediate face.

A **Generative Adversarial Network (GAN)** uses two networks fighting each other. A Generator tries to forge fake images from random noise, and a Discriminator acts as an art expert trying to detect the fakes. They train until the forgeries are indistinguishable from reality."

#### Indepth
VAEs optimize a lower bound on the data likelihood (ELBO) and explicitly define the probability distribution (usually Gaussian). GANs use an implicit distribution—they don't explicitly calculate probabilities, they just use a min-max game (Nash Equilibrium). GANs produce much sharper, crisper images, while VAEs produce slightly blurrier images but offer better latent space interpolation.

---

### 23. Describe how a VAE works.
"A VAE is just an autoencoder with a probabilistic spin. 

Standard autoencoders compress a 1080p image down to, say, a 512-number vector, and then decompress it back. The problem is, that 512-vector space is messy. 

A VAE fixes this by forcing the encoder to output two things: the **Mean** and the **Standard Deviation** of a normal distribution. To decode, the model randomly *samples* from that distribution. This mathematical trick ensures that points close together in the latent space represent images that look similar, allowing me to generate smooth transitions between faces or objects."

#### Indepth
The 'magic' of the VAE is the **Reparameterization Trick**. Since you can't backpropagate through a random sampling node, the VAE treats the sampling as an external input ($\epsilon$) multiplied by the standard deviation. This allows the gradient to flow cleanly back into the encoder during the training process, balancing the reconstruction loss and the Kullback-Leibler (KL) divergence.

---

### 24. What is a latent space in generative models?
"Latent space is the hidden 'mathematical map' where the AI organizes concepts. 

If I train a model on human faces, the latent space is a multi-dimensional graph. One direction on this graph might represent 'adding glasses.' Another direction might represent 'getting older.' 

It's called 'latent' because these features (glasses, age, gender) aren't explicitly labeled; the model discovers them as continuous mathematical vectors. By navigating this space—taking the coordinates of a young face and adding the 'older' vector—I generate an old face."

#### Indepth
High-quality latent spaces demonstrate **disentanglement**. If a latent space is entangled, moving the "add glasses" vector might accidentally make the person blonde. Disentangling factors of variation is a major research goal, allowing users to control specific features (like pose or lighting) without altering the subject's identity.

---

### 25. How is sampling done in Generative AI?
"Sampling is how I extract a discrete output from the model's complex probability distribution. 

In a Language Model, the model predicts the probabilities for the next 50,000 words in its vocabulary. I 'sample' from this array. I could just pick the highest probability (greedy search), but that makes the text boring. Instead, I use techniques like **Top-K** (limit to the top 50 choices and randomly pick based on probability) or **Top-p/Nucleus** (limit to a dynamic subset of words whose probabilities add up to 90%)."

#### Indepth
Sampling temperature manipulates the logits *before* the softmax function turns them into probabilities. A temperature of 1.0 does nothing. A temperature $< 1$ makes the peaks higher and the valleys lower (more deterministic). A temperature $> 1$ flattens the distribution, making lower-ranked words more likely to be sampled, increasing 'creativity' but risking gibberish.

---

### 26. Explain diffusion models in simple terms.
"I like to explain diffusion models using a cup of coffee. 

Imagine you drop a cube of sugar into coffee. Over time, it diffuses until the coffee is uniformly sweet (pure noise). 

A diffusion model is trained to reverse this exact process. It takes a canvas of pure TV static (noise) and step-by-step, 'un-diffuses' it, predicting how to remove exactly one layer of noise to reveal the hidden image underneath. If I condition it with a text prompt like 'draw a dog,' it guides the noise removal to eventually reveal a dog."

#### Indepth
Mathematically, this uses a Markov chain of Gaussian noise injection (forward process) and a U-Net architecture predicting the added noise (reverse process). The breakthrough was DDPM (Denoising Diffusion Probabilistic Models), proving that parameterizing the model to predict the *noise* rather than the original image results in much stabler training and better sample quality than GANs.

---

### 27. What role do attention mechanisms play in multimodal generation?
"Attention is what allows models to bridge entirely different types of data, like text and images. 

If I use an AI to caption an image, the Attention mechanism calculates which specific pixels in the image correspond to which specific words in the text. 

When generating the word 'frisbee', the Cross-Attention mechanism 'looks' specifically at the red circular pixels in the image feature map, ignoring the background grass. It dynamically weighs the importance of the visual input based on the text it is currently generating."

#### Indepth
Cross-attention acts as the bridge. In an image generator like Stable Diffusion, the visual U-Net treats the text embeddings (from CLIP) as the Key and Value matrices in the attention calculation, while treating the noisy image latent pixels as the Query. This forces the noisy image to align itself mathematically with the text prompt.

---

### 28. How does attention work in transformers?
"Attention mathematically answers the question: 'While reading this specific word, which other words in the sentence should I care about?' 

For the sentence 'The bank of the river,' the word 'bank' could mean money or a river edge. The attention mechanism compares the 'bank' vector against the 'river' vector. It calculates a high compatibility score between them, signaling to the model that in this specific context, 'bank' refers to nature, not finance."

#### Indepth
It works using three vectors per token: Query (Q), Key (K), and Value (V). The similarity score is $Q . K^T / \sqrt{d_k}$. This score (after softmax) acts as a weight. The new representation of the word 'bank' becomes a weighted sum of the Values (V) of all surrounding words, effectively baking the context directly into the word's embedding vector.

---

### 29. Explain the training objective of GPT.
"GPT relies on an incredibly simple objective: **Next Token Prediction** (Causal Language Modeling). 

I take 1 terabyte of text data. I feed the model 'The cat sat on the'. The model predicts 'dog'. I calculate the loss because the correct word was 'mat'. I use backpropagation to adjust the weights, and move to the next word. 

Because we have essentially infinite text data on the internet, this self-supervised process allows GPT to learn grammar, facts, reasoning, and coding syntax entirely on its own."

#### Indepth
The specific loss function is Cross-Entropy Loss over the vocabulary distribution. While simple, predicting the next token perfectly requires immense world knowledge. To predict the next word in 'The capital of France is ___', the model *must* internalize geography. This emergent understanding of reality makes GPT vastly more than just an autocomplete algorithm.

---

### 30. What is transfer learning and how is it used in GenAI?
"Transfer learning is taking knowledge learned from one broad task and applying it to a specific task. 

Training an LLM from scratch on a medical database would require thousands of GPUs and months of time. Instead, I download an open-source model like Llama-3 that has already spent months learning English grammar, general facts, and reasoning properties. 

I then 'transfer' that foundational knowledge by fine-tuning it for a few hours on my specific medical dataset. It learns the medical jargon rapidly because it already understands the underlying structure of language."

#### Indepth
In image generation, Stable Diffusion serves as the broad 'foundation model'. You don't train a new diffusion model to generate anime. You use transfer learning (via LoRA or Dreambooth) to adjust the cross-attention weights of the pre-trained model so it maps its immense existing knowledge of anatomy, shading, and lighting to the specific anime visual style.

---

### 31. How are embeddings used in LLMs?
"An embedding is a translation of a word into a massive list of numbers (a vector) that captures its semantic meaning. 

In a 1024-dimensional embedding space, the vector for 'King' and 'Queen' are mathematically very close together (high cosine similarity), while 'King' and 'Car' are far apart. 

When a user types a prompt into an LLM, the very first thing the architecture does is look up the embedding vectors for those words. It is the fundamental input layer that turns human language into math."

#### Indepth
Interestingly, embeddings support vector arithmetic: $Embedding(King) - Embedding(Man) + Embedding(Woman) \approx Embedding(Queen)$. This geometric relationship proves the network has inferred semantic properties without explicit instruction. The initial embedding layer is $V \times d$, where $V$ is vocab size (e.g. 50,000) and $d$ is embedding dimension (e.g. 4096).

---

### 32. Explain few-shot, zero-shot, and one-shot learning in LLMs.
"These refer to how many examples I give the model in my prompt to teach it a new task *without* changing its weights.

**Zero-shot:** 'Translate Apple to French.' (No examples provided, relying entirely on pre-training).
**One-shot:** 'Dog -> Perro. Apple -> ?' (One example provided to show the desired pattern).
**Few-shot:** Showing it 5 examples of analyzing sentiment before asking it to analyze a 6th review.

I use few-shot heavily in production to force the LLM to output specific JSON structures or edge-case labels."

#### Indepth
This capability (In-Context Learning) is an emergent property of large transformer models. The model temporarily uses its self-attention heads to map the implicit relationship established in your few-shot prompt to the final query, mimicking the behavior of gradient updates (training) simply through forward-pass matrix multiplication.

---

### 33. What are diffusion models like Stable Diffusion used for?
"Stable Diffusion is fundamentally an **Image Generation** model. 

I use it primarily for text-to-image ('generate a castle on a cloud'), but it's incredibly versatile. 
It supports **Image-to-Image** (turning a rough sketch into a photorealistic rendering), **Inpainting** (erasing a person from a photo and predicting the background), and **ControlNet** (forcing the generated image to perfectly match the pose of a human in a reference photo)."

#### Indepth
The 'Stable' in Stable Diffusion signifies it operates in a *Latent* space, not pixel space. Unlike earlier models that diffused huge 1024x1024 pixel grids (which required absurd VRAM), SD uses a VAE to compress the image into a tiny 64x64 latent grid, runs the diffusion denoising there, and decodes it back to pixels. This made it the first model capable of running on consumer gaming GPUs.

---

### 34. What is CLIP and how is it related to image generation?
"CLIP (Contrastive Language-Image Pre-training) is the translator between pictures and words. 

OpenAI trained it by showing it 400 million images and their alt-text captions from the internet. It projects both text and images into the exact same mathematical embedding space. 

It is the heart of models like DALL-E or Stable Diffusion. When I prompt 'a red apple', CLIP turns that text into a vector. Because the diffusion model expects a CLIP vector as guidance, it generates an image that perfectly matches that vector's position in the embedding space."

#### Indepth
CLIP uses a contrastive loss function. Given a batch of $N$ image-text pairs, it maximizes the cosine similarity between the $N$ correct pairs on the diagonal of an $N \times N$ matrix while minimizing the similarity of the $N^2 - N$ incorrect pairs. This forces the model to learn incredibly robust, generalized concepts linking vision and language.

---

### 35. What is a spectrogram and how is it used in generative audio?
"If audio is a time-series wave, a Spectrogram is a 2D image of that wave representing Time (X-axis), Frequency/Pitch (Y-axis), and Loudness (Color Intensity). 

When we train generative audio models (like text-to-speech or music generators), predicting a waveform millisecond by millisecond is computationally terrible. Instead, we use GenAI algorithms originally meant for images (like CNNs or Diffusion) to generate a customized Spectrogram. 
Then, a vocoder algorithm translates that 2D image back into an audible 1D wave."

#### Indepth
Usually, a **Mel-Spectrogram** is used, which compresses frequencies using a logarithmic scale matching human hearing perception. Modern audio models (like AudioLM) bypass the continuous waveform entirely by tokenizing audio into discrete "acoustic tokens" and training a Transformer LLM on them, exactly like predicting the next text word.

---

### 36. How does Retrieval-Augmented Generation (RAG) work?
"RAG fixes the Achilles heel of LLMs: hallucinations and lack of proprietary data. 

In my enterprise architecture, when a user asks 'What is our Q3 refund policy?', I don't send that straight to the LLM. 

1. **Retrieval**: I take the user's query and search my private Vector Database for relevant internal documents.
2. **Augmentation**: I take the top 3 retrieved paragraphs and paste them invisibly into the prompt.
3. **Generation**: I send the LLM the prompt: 'Using *only* the following paragraphs, answer the user's question.'

This guarantees the LLM's answer is factually grounded in my private, up-to-date data."

#### Indepth
Naïve RAG often fails in production because user queries are messy. Advanced RAG uses **Query Rewriting** (using a small LLM to rephrase 'it's broken' into a searchable 'Refund Policy Hardware Failure 2024'), **Re-ranking** (using a Cross-Encoder to guarantee the 3 retrieved documents are logically perfect for the specific query), and **Chunking Strategies** (semantic vs character-based chunking of the PDFs).

---

### 37. What is the role of vector databases in GenAI applications?
"Vector Databases (like Pinecone, Milvus, or pgvector) are built specifically to do extremely fast 'similarity searches' on embeddings. 

Traditional SQL databases search by exact keyword matches (`WHERE word = 'Dog'`). Vector databases search by math. They store the numerical vectors generated from all my company PDFs. When a user queries 'Puppy', it calculates the cosine distance between 'Puppy' and every paragraph I have stored, instantly returning the paragraphs regarding 'Dogs', even if the exact keyword never appears."

#### Indepth
Vector databases achieve this speed using Approximate Nearest Neighbor (ANN) algorithms instead of brute-force $O(N)$ comparisons. The most common algorithm is HNSW (Hierarchical Navigable Small World). It builds a multi-layered graph linking nearby vectors. Searching it is $O(\log N)$, allowing me to find the most relevant paragraph out of billions in milliseconds.

---

### 38. What is LangChain and how is it used with GenAI?
"LangChain is a Python/JavaScript orchestration framework for building complex LLM apps. 

If I just want a chatbot, I deal with the OpenAI API directly. But if I want an 'AI Agent' that can search Google, query my Postgres database, execute Python scripts, and summarize the output—doing that manually turns into spaghetti code. 

LangChain provides standard interfaces to link prompts, LLMs, Memory, Vector Stores, and External Tools into logical 'Chains'. It handles the messy backend plumbing of AI workflows."

#### Indepth
LangChain abstracts concepts like `DocumentLoaders` (turning PDFs/Websites into raw text), `TextSplitters` (chunking tokens for embedding), and `Memory` (keeping a rolling summary of chat history). For Agentic workflows, LangGraph is actively replacing basic LangChain by allowing developers to model complex cyclical node-based state machines for autonomous LLM processes.

---

### 39. How do you integrate a GenAI model into a web app?
"I treat the LLM identically to any external REST API or third-party service. 

1. **Backend Proxy:** The browser never calls the LLM directly because that exposes my API keys.
2. **Streaming:** Text generation takes seconds. I use Server-Sent Events (SSE) or WebSockets on the backend to stream tokens up to the frontend React component to simulate the 'typing' animation.
3. **Guardrails:** On the backend, before querying the LLM, I validate the user prompt for toxicity or injection attacks. After generation, I parse the JSON (using tools like Zod) to ensure the UI won't crash on malformed output."

#### Indepth
If I am hosting my own open-source model (like Llama 3) instead of using OpenAI, I use an inference engine like **vLLM** or **Ollama** on GPU servers. These wrap the raw Python ML code in high-performance HTTP servers utilizing techniques like PagedAttention to serve multiple web requests concurrently without Out-Of-Memory (OOM) crashes.

---

### 40. How do you reduce hallucinations in GenAI systems?
"I use a multi-layered strategy for enterprise apps:

1. **Lower Temperature:** Setting `temperature=0.0` or `0.1` restricts the model from 'creative guessing'.
2. **Strict System Prompting:** Giving explicit rules: 'If the retrieved context does not contain the answer, reply exactly: "I do not know." Do not infer.'
3. **RAG (Retrieval-Augmented Generation):** Forcing the model to look at hard facts rather than its internal memory.
4. **Citation Extraction:** Forcing the LLM to output a precise quote or document ID alongside its answer."

#### Indepth
For critical workflows (like legal review), I employ **Self-Reflection** (or LLM-as-a-Judge). After LLM-A generates the summary, an invisible LLM-B is prompted to review the summary against the source text to detect if LLM-A hallucinated any clauses. Only if LLM-B passes it does it render to the user UI.
