# MLOps and Deployment (Product-Based Companies)

At top companies, an ML model is useless if it sits in a Jupyter Notebook. MLOps is the discipline of reliably training, deploying, monitoring, and scaling models in production, treating ML components exactly like critical software infrastructure.

## 1. What is the difference between Data Parallelism and Tensor (Model) Parallelism when training/serving massive LLMs across clusters?
**Answer:**
When a model (like a 70B parameter LLM) is too large to fit the weights, gradients, and optimizer states into the VRAM of a single GPU (e.g., an 80GB A100), we must distribute it.

**Data Parallelism (DP):**
*   **Concept:** The entire model replica fits onto *one* GPU. We distribute the *Data*.
*   **Mechanism:** Every GPU in the cluster gets an exact copy of the model. We split the training batch (e.g., batch size 64 across 8 GPUs = 8 samples per GPU). Each GPU computes the forward and backward passes independently to find its gradients. At the end of the step, an all-reduce operation averages the gradients across all GPUs, updates the weights uniformly, and proceeds to the next step.
*   **Limitation:** It fails entirely if the model itself cannot fit on a single GPU's VRAM.

**Tensor Model Parallelism (TP):**
*   **Concept:** The model is too big for one GPU, so we distribute the *Model Weights (Tensors)*.
*   **Mechanism:** We slice the actual weight matrices (e.g., those inside the self-attention mechanism or MLPs) and place pieces of them on different GPUs. During the forward pass, GPU 1 calculates part of the math, GPU 2 calculates the other part. They must communicate rapidly over NVLink to perform matrix additions/concatenations *in the middle of a layer's execution*.
*   **Limitation:** Requires extremely high intra-node communication speed (NVLink). It is usually limited to the GPUs within a single physical server chassis (typically up to 8 GPUs) because going across the network (Infiniband) is too slow for mid-layer synchronization.

*(Note: Pipeline Parallelism is the third type, where different whole layers are placed on different GPUs. Training a 100B+ model requires 3D Parallelism: DP + TP + PP simultaneously).*

## 2. Walk me through the architecture of a robust ML Feature Store. Why do we need it instead of just querying the raw database?
**Answer:**
**The Problem:**
Data scientists write Python scripts to calculate features (e.g., "rolling 30-day average transaction count"). They train the model offline using a Data Warehouse. When deploying the model online, software engineers must rewrite that exact same feature-calculation logic in Java/Go to run on live production databases. This leads to **Training-Serving Skew** (the offline feature calculation drifts from the live logic, ruining predictions) and massive engineering duplication.

**The Feature Store Architecture (e.g., Feast, Hopsworks):**
A Feature Store acts as a centralized registry and data layer specifically for ML features.

1.  **Single Definition (The Registry):** Data Scientists define the feature calculation logic *once* (usually in PySpark or Pandas). This definition is checked into version control.
2.  **Dual Storage Systems:**
    *   **Offline Store (e.g., Snowflake, BigQuery):** Stores massive historical feature values based on the central definition. Used by Data Scientists to generate training datasets with point-in-time correctness (preventing data leakage).
    *   **Online/Real-Time Store (e.g., Redis, DynamoDB):** A low-latency key-value store. Streaming jobs or microservices continuously write the absolutely freshest feature values here based on the exact same central definition.
3.  **The API Layer:** During production inference, the serving application calls the Feature Store API: `get_features(user_id=123)`. The API pulls from the Online Store in milliseconds and passes the correctly formatted vector to the model backend.

**Why it's needed:** Eliminates training-serving skew, promotes re-usability (Team A can use Team B's curated embeddings), and dramatically reduces "time-to-production."

## 3. How do you monitor machine learning models in production? Conceptually, what is Data Drift vs. Concept Drift?
**Answer:**
Model accuracy usually degrades the moment it is deployed. Monitoring requires tracking more than just latency/throughput; you must track the statistical behavior of the model.

**1. Data Drift (Covariate Shift):**
*   **Definition:** The statistical distribution of the input features $P(X)$ changes significantly from what the model saw during training. The relationship between input and output might still be the same, but the model is operating in an unfamiliar region of the feature space.
*   **Example:** You train a credit risk model on a population with an average age of 40. A new marketing campaign brings in thousands of 18-year-olds. The $P(Age\_Feature)$ has drifted.
*   **Detection:** Calculate statistical distance metrics (KL Divergence, Population Stability Index (PSI), Wasserstein Distance) between the training dataset features and the currently incoming, real-time feature streams. If the distance crosses a threshold, trigger an alert.

**2. Concept Drift:**
*   **Definition:** The fundamental relationship between the inputs and the target variable $P(Y | X)$ changes. What was true yesterday is false today. The model's logic is fundamentally outdated.
*   **Example:** Predicting housing prices. The relationship between square footage and price changes dramatically before vs. after a sudden economic recession.
*   **Detection:** This is harder. Unless you have real-time ground-truth labels (like knowing if a transaction was fraudulent 5 seconds after it happened), you often detect this via degrading downstream business metrics (e.g., recommendations are generating fewer clicks over time).

**Action upon Detection:** Trigger automated re-training pipelines using fresh data. If drift is severe, fall back to simple heuristic rules until a new model is validated.

## 4. Why are classical REST APIs built with Flask/FastAPI often insufficient for serving large Deep Learning models? Explain technologies like Triton, TorchServe, or vLLM.
**Answer:**
**The FastAPI Limitations:**
Wrapping PyTorch inference in a simple FastAPI endpoint is fine for prototypes, but terrible for high-throughput production systems handling thousands of requests per second (RPS).
1.  **Inefficient Batching:** GPUs operate best when crunching large matrices. A standard REST API processes requests sequentially (batch size 1). You pay the massive overhead of GPU kernel launches for incredibly tiny workloads, wasting $>90\%$ of the GPU's potential.
2.  **Concurrency Management (GIL):** Python's Global Interpreter Lock prevents true multi-threading, causing bottlenecks when managing many concurrent IO requests.
3.  **Memory Management:** Loading models natively in Python across multiple workers duplicates memory instead of sharing it effectively.

**The Inference Server Solution (NVIDIA Triton, TorchServe):**
Robust inference servers are usually written in C++ and act as a specialized layer sitting directly on the GPU.

*   **Dynamic Batching (Crucial):** If 10 separate REST requests hit Triton at the exact same millisecond, Triton intercepts them, pauses for a configurable microsecond window (e.g., 5ms), concatenates all 10 inputs into a single batch, sends that batch to the GPU to calculate at once, slices the results back up, and returns them to the 10 separate users. This scales throughput exponentially with minimal latency hit.
*   **Model Versioning and Hot-Swapping:** Allows zero-downtime updates. You can load version 2 of a model, route 1% of traffic to it (A/B testing), and take down version 1 without restarting the container.

**For LLMs specifically (vLLM, TGI):**
Serving autoregressive LLMs requires extreme optimization beyond Triton. Tools like vLLM provide custom C++ CUDA kernels (PagedAttention) to eliminate memory fragmentation during generation and perform "Continuous Batching"—adding and evicting new requests dynamically at the token level, rather than waiting for an entire prompt to finish generating.
