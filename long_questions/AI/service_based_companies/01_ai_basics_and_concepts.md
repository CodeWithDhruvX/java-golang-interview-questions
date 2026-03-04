# AI Basics and Concepts (Service-Based Companies)

In service-based interviews, you must clearly distinguish fundamental concepts, explain them simply (as you would to a non-technical client), and know which evaluation metric fits a given business problem.

## 1. Explain the differences between Artificial Intelligence (AI), Machine Learning (ML), and Deep Learning (DL).
**Answer:**
*   **Artificial Intelligence (AI):** The broadest concept. It is the simulation of human intelligence by machines. It includes anything that allows a computer to mimic human behavior, from simple rule-based if/else expert systems to complex neural networks. It's the overarching goal.
*   **Machine Learning (ML):** A subset of AI. Instead of explicitly programming the rules (as in traditional software), ML uses statistical methods to enable machines to *learn from data* and improve from experience without being explicitly programmed. Example: A spam filter learning what words indicate spam based on past emails.
*   **Deep Learning (DL):** A specialized subset of ML. It specifically utilizes multi-layered Artificial Neural Networks (inspired by the human brain architecture). DL is capable of learning massive amounts of unstructured data (images, text, audio) automatically, often removing the need for human engineers to perform manual "feature extraction."

## 2. What is the difference between Supervised, Unsupervised, and Reinforcement Learning? Provide a business example for each.
**Answer:**
*   **Supervised Learning:** The model is trained on labeled data. For every input ($X$), you provide the exact "correct answer" or label ($Y$). The model learns the mapping function.
    *   *Business Example:* Predicting House Prices (Regression) where the data includes features like square footage and the historical sale price. Or predicting Customer Churn (Classification) where the data indicates whether past customers left (Yes/No).
*   **Unsupervised Learning:** The model is trained on unlabeled data. The algorithm must discover patterns, groupings, or structures on its own without predefined correct answers.
    *   *Business Example:* Customer Segmentation. Grouping a company's millions of customers into distinct clusters based on their purchasing behavior so marketing can target them differently.
*   **Reinforcement Learning (RL):** An agent learns to make decisions by performing actions in an environment to maximize some notion of cumulative reward. It learns via trial and error.
    *   *Business Example:* Optimizing logistics and delivery routes in real-time, algorithmic trading in finance, or training robots to navigate a warehouse.

## 3. Explain Precision, Recall, and the F1-Score. When would you prioritize Recall over Precision?
**Answer:**
These are metrics used to evaluate classification models, especially on imbalanced data. Let's use the example of predicting if a patient has Cancer (Positive class).
*   **Precision:** Out of all the patients the model *predicted* as having cancer, what percentage *actually* had cancer? ($TP / (TP + FP)$). High precision means fewer false alarms.
*   **Recall (Sensitivity):** Out of all the patients who *actually* have cancer, what percentage did the model successfully find? ($TP / (TP + FN)$). High recall means fewer missed cases.
*   **F1-Score:** The harmonic mean of Precision and Recall. It provides a single metric that balances both, useful when you care about both false positives and false negatives equivalently.

**When to prioritize Recall (minimize False Negatives):**
You prioritize Recall in healthcare/disease prediction or fraud detection.
*   *Why:* Telling a healthy person they might have cancer (a False Positive, hurting Precision) leads to an unnecessary biopsy (bad, but manageable). But missing a cancer patient and telling them they are healthy (a False Negative, hurting Recall) can be fatal. It is better to "over-warn".

**When to prioritize Precision (minimize False Positives):**
You prioritize Precision when the cost of a false alarm is very high.
*   *Why:* In an automated email spam filter, classifying a legitimate work email from the CEO as spam (False Positive) is disastrous. It's better to accidentally let a few spam emails into the inbox (False Negative) than to hide an important legitimate email.

## 4. What is the Bias-Variance Tradeoff, and how does it relate to Overfitting and Underfitting?
**Answer:**
This is the fundamental problem in machine learning. It describes how well a model generalizes to *new, unseen* data. Let's imagine trying to shoot arrows at a target's bullseye.

*   **Bias:** The error from erroneous assumptions. A high bias model pays very little attention to the training data and oversimplifies the problem.
    *   *Symptom:* **Underfitting**. The model performs poorly on the training data AND the test data. It's too rigid (e.g., trying to fit a complex curve with a straight line). High Bias = Misses the bullseye consistently.
*   **Variance:** The error from sensitivity to small fluctuations in the training set. A high variance model pays *too much* attention to the training data, memorizing the noise and outliers instead of the signal.
    *   *Symptom:* **Overfitting**. The model performs perfectly on the training data (near 0% error) but performs terribly on new test data because it memorized the specific examples rather than learning the general rule. High Variance = Arrows are scattered widely based on the wind (noise).

**The Tradeoff:**
You cannot minimize both simultaneously. Making a model more complex (adding features, deepening a tree) reduces Bias but increases Variance (risk of overfitting). Making a model simpler reduces Variance but increases Bias (risk of underfitting). The goal is to find the "sweet spot" in the middle.

## 5. What are Hyperparameters vs. Model Parameters? Support your answer with examples.
**Answer:**
*   **Model Parameters:** These are the configuration variables internal to the model whose values can be estimated directly from the data during the training process. They are *learned* automatically by the algorithm.
    *   *Examples:* The weights ($w_i$) and bias ($b$) in a Linear/Logistic Regression model. The split points acting as rules within nodes of a Decision Tree. The weights in an Artificial Neural Network.
    *   *Key trait:* You do not set these manually. The training algorithm adjusts them.

*   **Hyperparameters:** These are the configuration variables set *before* the training process begins. They dictate *how* the algorithm should learn. They cannot be learned directly from the standard training data and must be tuned manually or via search algorithms (like Grid Search).
    *   *Examples:* The learning rate ($\alpha$) in Gradient Descent. The `max_depth` of a Decision Tree (to prevent overfitting). The number of estimators (trees) in a Random Forest. The $K$ value in K-Nearest Neighbors.
    *   *Key trait:* You configure these as an engineer to try and find the best performing model setup.
