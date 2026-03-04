# AI System Design (Product-Based Companies)

AI System Design is arguably the most critical interview round for senior ML Engineering roles. You are tested on your ability to map business problems to ML solutions, handle massive scale, and manage infrastructure trade-offs.

## 1. Design a large-scale Video Recommendation System (e.g., YouTube/TikTok). Walk through the architecture.
**Answer:**
A standard large-scale recommendation system follows a **Two-Stage Funnel Architecture: Retrieval (Candidate Generation) and Ranking.**

**1. Retrieval (Candidate Generation):**
*   **Goal:** Quickly narrow down millions/billions of videos to a few hundred (e.g., Top 500) relevant candidates per user. This step must be extremely fast (low latency) and prioritize *recall* over precision.
*   **Architecture:**
    *   **Two-Tower Model:** One neural network embedding the User (features like watch history, demographics, device type), and another neural network embedding the Item (video ID, metadata, tags). Use contrastive loss on historical clicks to train the embeddings.
    *   **Vector Database/ANN:** Pre-compute the Item embeddings offline and store them in an Approximate Nearest Neighbor (ANN) index (like FAISS or HNSW). At request time, fetch the User embedding and perform a lightning-fast vector similarity search ($K$ nearest neighbors) to retrieve the 500 closest videos.
    *   **Fallback Heuristics:** In addition to the Two-Tower, pull candidates based on "Trending Globally," "Trending in Location," and "Collaborative Filtering (users like you also liked)."

**2. Ranking (Scoring):**
*   **Goal:** Take the 500 candidates and precisely rank them to maximize the target objective (e.g., user engagement time, click-through rate, completion rate). This operates on a smaller set, allowing for computationally intensive models.
*   **Architecture:**
    *   **Complex Model:** Assemble a Deep Learning Ranker combining multiple specific features. Common architectures include **Wide & Deep** (memorizes rules but generalizes well) or **DLRM** (Deep Learning Recommendation Model).
    *   **Multi-Task Learning (MMoE):** Crucially, you don't just predict "clicks." A user clicking then leaving in 2 seconds is bad. Predict multiple objectives simultaneously: $P_{click}$, $P_{watch\_time}$, $P_{like}$, $P_{comment}$. Combine them using a weighted score: $Score = W_{click}P_{click} + W_{watch}P_{watch\_time}$.
*   **Outputs:** Produce an ordered list of 100 videos.

**3. Re-Ranking (Filtering/Business Logic):**
*   **Goal:** Final polished list ready for the UI.
*   **Logic:** Apply filters (remove videos the user already watched), enforce diversity (don't show 10 Fortnite videos in a row), ensure fairness, and boost sponsored content based on active ad campaigns.

**4. Data Pipeline & Serving:**
*   **Offline Training:** Nightly training of the heavy Ranking model using Airflow/Spark on a Data Warehouse. Generate new Item Embeddings weekly.
*   **Online Serving:** When a user opens the app, an API Gateway routes the request to the Serving Layer. Use Redis/Feature Stores (Feast) to look up real-time user features (e.g., "what did they click 5 minutes ago") to inject into the ranking model.

## 2. In an e-commerce platform, how do you solve the "Cold Start Problem" for new items and new users?
**Answer:**
*   **New Item Cold Start (No historical clicks):**
    *   **Content-Based Filtering:** The model cannot rely on Collaborative Filtering. It must rely entirely on the item's metadata (Item features). If it's a new "Blue Nike Shoe", push it to users whose profile embedding strongly matches shoes, Nike, or the color blue.
    *   **Exploration/Exploitation Strategy (Multi-Armed Bandits):** We must allocate a small portion of traffic specifically to test new items. Randomly (but intelligently) surface the new item to a generic audience slice. Gather the initial impression/click data, then feed it back into the model to quickly establish its baseline engagement rate. Epsilon-Greedy or Thompson Sampling are common MAB algorithms.
*   **New User Cold Start (No user history):**
    *   **Onboarding Flow/Heuristics:** Ask the user explicitly what they like during signup (e.g., selecting favorite genres on Netflix). Serve a hardcoded "Popular Items" list based on their immediate context (location, time of day, device language).
    *   **Contextual Bandits:** Use whatever immediate implicit features you have (session ID, referring site, browser fingerprint) to group the user into broad cohorts and guess recommendations until behavioral data accumulates.

## 3. Describe the architecture and trade-offs of building a Retrieval-Augmented Generation (RAG) system over millions of company documents.
**Answer:**
**The Architecture:**
A large-scale RAG system involves complex ingestion pipelines and optimized querying mechanisms.

**1. Data Ingestion Pipeline (Offline):**
*   **Extraction:** Crawling Confluence, Jira, PDFs, Slack. Handling complex formats (tables in PDFs are notoriously hard; often requires specialized models like LayoutLM).
*   **Chunking Strategy:** Chunking the text. Semantic chunking (breaking at paragraphs/headers) is vastly superior to fixed-length character chunking because it preserves context.
*   **Embedding & Storage:** Pass the chunks through a high-quality embedding model (e.g., OpenAI `text-embedding-3-large` or an open-source BGE model). Store the vectors alongside metadata (Author, Date, Source URL, Tags) in a specialized Vector Database (Pinecone, Milvus, Qdrant).

**2. The Retrieval & Generation Pipeline (Online):**
*   **Query Transformation:** When the user types a query ("How do I deploy?"), don't blindly embed it. Treat the query. Use an LLM to expand the query, fix synonyms, or generate hypothetical answers (HyDE) to embed instead.
*   **Hybrid Search:** Pure vector similarity search fails when users ask for exact keywords or IDs. You must implement Hybrid Search: combines Dense Vector Search (semantic similarity) with Sparse Keyword Search (BM25 algorithms on ElasticSearch), and fuse the scores using tools like Reciprocal Rank Fusion (RRF).
*   **Re-Ranking:** The Database returns top 50 chunks. Pass these through a specialized "Cross-Encoder" Re-Ranker model (e.g., Cohere Rerank) to precisely sort the top 5 chunks.
*   **Generation & Citation:** Feed the system prompt + top 5 chunks + user query into the final LLM to generate the answer. The system must cite which chunk provided which fact.

**Trade-offs at Scale:**
*   **Chunk Size:** Small chunks give highly precise embeddings but lose broad context. Large chunks capture context but the embedding might dilute the specific fact you are searching for.
*   **Embedding Delay:** If a new document is added to Confluence, it takes time to chunk, embed, and index it. Does the business need real-time RAG (seconds) or is nightly batch indexing acceptable?
*   **Vector DB Cost:** Storing billions of 1536-dimensional vectors in RAM (necessary for low-latency HNSW search) is extremely expensive. You must use Product Quantization (PQ) to compress the vectors and trade perfect accuracy for massive cost savings.

## 4. How do you design an ML system to detect fraudulent credit card transactions in real-time?
**Answer:**
**Key Constraints:**
*   Extreme Class Imbalance (99.9% legitimate, 0.1% fraud).
*   Strict Latency (Must approve/decline a swipe in $<50$ milliseconds).
*   Concept Drift (Fraudsters constantly change tactics; models degrade quickly).

**Architecture:**
1.  **Event Ingestion & Stream Processing:** Swipe happens. Event hits an API Gateway and is pushed onto a high-throughput message queue (Apache Kafka). A stream processor (Apache Flink or Spark Streaming) picks it up instantly.
2.  **Real-Time Feature Store:** The ML model needs context. The Flink job queries a low-latency database (Redis/Aerospike, which acts as a Feature Store) to fetch moving aggregates: "How many times did this user try to buy jewelry in the last 10 minutes from a different IP?" These aggregates must be updated continuously by the stream processor.
3.  **The ML Model:**
    *   **The Classifier:** Given the 50ms constraint, complex deep learning models might be too slow unless heavily optimized via TensorRT. The industry standard is highly optimized Gradient Boosted Trees (XGBoost/LightGBM) deployed directly in C++ or via specialized inference servers.
    *   **Graph Features:** Fraud often occurs in rings. Using Graph Neural Networks (GNNs) on the backend to detect suspicious clusters of transaction nodes and turning those graph embeddings into features for the XGBoost model is highly effective.
4.  **Ensembling & Rules:** The ML model score is combined with a hardcoded Rules Engine. Even if the ML score is low, if a rule says "Reject all transactions over \$10k from IP ranges in country X," the transaction is blocked.
5.  **Offline Re-Training Pipeline:** Continuous monitoring is vital. Transactions that are later reported as chargebacks are labeled "fraud." Nightly pipelines must retrain the model and automatically deploy shadow models to check for concept drift.
