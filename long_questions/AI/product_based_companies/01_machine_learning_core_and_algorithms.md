# Core Machine Learning & Algorithms (Product-Based Companies)

This section focuses on the deep foundational knowledge required by top-tier product companies. You are expected to know the math behind the models, understand trade-offs perfectly, and know how algorithms scale.

## 1. Explain the mathematical intuition behind Gradient Boosting algorithms (like XGBoost or LightGBM) and how they differ from Random Forests.
**Answer:**
Both Random Forests and Gradient Boosting are ensemble methods, but they operate on fundamentally different principles.

*   **Random Forests (Bagging):**
    *   **Intuition:** Builds numerous independent decision trees (usually deep ones) in parallel. Each tree is trained on a random subset of the data (bootstrapping) and a random subset of the features.
    *   **Math:** The final prediction is the average (regression) or majority vote (classification) of all trees. This reduces the *variance* of the model without significantly increasing bias.
    *   **Loss Function:** Trees are built independently; there's no iterative loss optimization across the ensemble.

*   **Gradient Boosting (Boosting):**
    *   **Intuition:** Builds shallow trees *sequentially*. Each new tree tries to predict the *errors (residuals)* of the combined ensemble of all previous trees.
    *   **Math:** It's an optimization algorithm on a cost function. At iteration $m$, the model $F_m(x) = F_{m-1}(x) + \alpha h_m(x)$, where $h_m(x)$ is the new tree trained to predict the negative gradient of the loss function with respect to the previous model's output (pseudo-residuals). $\alpha$ is the learning rate.
    *   **Variance/Bias:** It reduces *bias* (by iteratively fitting the data better). To prevent overfitting (variance), we use a slow learning rate and shallow trees.

*   **XGBoost Specifics:** XGBoost adds *regularization* ($\Omega(f) = \gamma T + \frac{1}{2}\lambda||w||^2$) directly into the objective function to penalize complex trees (number of leaves $T$ and leaf weights $w$). It also uses a second-order Taylor approximation (using both Gradient and Hessian) of the loss function to find the optimal split points faster and more accurately than standard Gradient Boosting.

## 2. Derive the update rule for Logistic Regression using Gradient Descent. Explain the role of the sigmoid function.
**Answer:**
**Role of Sigmoid:**
The sigmoid function $\sigma(z) = \frac{1}{1 + e^{-z}}$ squashes any real number into the range $(0, 1)$. In logistic regression, we model the probability of the positive class as $P(y=1|x) = \hat{y} = \sigma(w^T x + b)$.

**Derivation:**
1.  **Hypothesis:** $\hat{y} = \sigma(z)$ where $z = w^T x + b$
2.  **Loss Function (Binary Cross-Entropy / Log Loss):** We want to maximize the likelihood, which is equivalent to minimizing the negative log-likelihood over $m$ examples.
    $$L(w,b) = -\frac{1}{m} \sum_{i=1}^{m} \left[ y^{(i)} \log(\hat{y}^{(i)}) + (1 - y^{(i)}) \log(1 - \hat{y}^{(i)}) \right]$$
3.  **Derivative of Sigmoid:** A crucial property is $\frac{\partial \sigma(z)}{\partial z} = \sigma(z)(1 - \sigma(z)) = \hat{y}(1 - \hat{y})$.
4.  **Chain Rule (Gradient with respect to weight $w_j$):**
    $$ \frac{\partial L}{\partial w_j} = \frac{\partial L}{\partial \hat{y}} \cdot \frac{\partial \hat{y}}{\partial z} \cdot \frac{\partial z}{\partial w_j} $$
5.  **Calculating parts:**
    *   $\frac{\partial L}{\partial \hat{y}} = -\left(\frac{y}{\hat{y}} - \frac{1-y}{1-\hat{y}}\right) = \frac{\hat{y}-y}{\hat{y}(1-\hat{y})}$
    *   $\frac{\partial \hat{y}}{\partial z} = \hat{y}(1-\hat{y})$ (from step 3)
    *   $\frac{\partial z}{\partial w_j} = x_j$
6.  **Combining:**
    $$ \frac{\partial L}{\partial w_j} = \left(\frac{\hat{y}-y}{\hat{y}(1-\hat{y})}\right) \cdot (\hat{y}(1-\hat{y})) \cdot x_j = (\hat{y} - y)x_j $$
7.  **Update Rule:** Over all $m$ examples, the gradient is the average.
    $$ w_j := w_j - \alpha \frac{1}{m} \sum_{i=1}^{m} (\hat{y}^{(i)} - y^{(i)})x_j^{(i)} $$
    *(Notice how elegantly the derivative simplifies, which is why Cross-Entropy is the natural pairing for Sigmoid).*

## 3. How do you handle highly imbalanced datasets in classification tasks at scale? (Beyond just "SMOTE")
**Answer:**
While SMOTE (Synthetic Minority Over-sampling Technique) is a common answer, in large-scale production systems, naive oversampling often fails or creates significant computational overhead.

*   **1. Algorithm Level Changes (Preferred at Scale):**
    *   **Cost-Sensitive Learning (Class Weights):** This is usually the first and most efficient step. We assign a higher penalty (weight) to misclassifying the minority class in the loss function. In libraries like XGBoost or Scikit-Learn, this is the `scale_pos_weight` or `class_weight` parameter.
        *   $Loss = - [ W_{pos} \cdot y \log(\hat{y}) + W_{neg} \cdot (1-y) \log(1-\hat{y}) ]$
    *   **Focal Loss:** Originally designed for Object Detection (RetinaNet), Focal loss dynamically scales the cross-entropy loss based on prediction confidence. It heavily penalizes "easy" examples (where the model is already confident) and focuses the model's attention on hard, misclassified examples (which usually belong to the minority class).
        *   $FL(p_t) = -\alpha_t (1 - p_t)^\gamma \log(p_t)$ (where $\gamma$ focuses on hard examples).

*   **2. Data Level (Careful Application):**
    *   **Strategic Downsampling:** Instead of just random undersampling, use techniques to remove redundant majority class examples (e.g., Tomek Links) or keep examples near the decision boundary. If we have 1 billion negative examples and 1 million positive, random downsampling of the negatives to 10 million is often better than trying to SMOTE the positives to 1 billion.
    *   **Two-Phase Training:** Train the initial model on a balanced dataset (downsampled negatives) to learn the basic features. Then, freeze the initial layers and fine-tune only the final decision layers on the true, imbalanced distribution so the model learns the correct prior probabilities.

*   **3. Evaluation Metric Shift:**
    *   **Never use Accuracy.**
    *   Rely on **Precision-Recall AUC (PR-AUC)** rather than ROC-AUC. ROC-AUC can be overly optimistic when the negative class is massive, whereas PR curves focus entirely on the positive (minority) class.

## 4. What is the difference between Generative and Discriminative models? Give examples of each and explain when to use which.
**Answer:**
*   **Discriminative Models:**
    *   **Goal:** Learn the boundary between classes. They model the conditional probability $P(Y | X)$. Given data $X$, what is the probability it belongs to class $Y$?
    *   **Function:** They directly map inputs to classes without caring how the data was generated.
    *   **Examples:** Logistic Regression, Support Vector Machines (SVM), Random Forests, standard Deep Neural Networks for classification.
    *   **When to use:** When your sole goal is prediction/classification and you have plenty of labeled data. They generally have lower asymptotic error because they solve a simpler problem (just finding the boundary, not modeling the whole distribution).

*   **Generative Models:**
    *   **Goal:** Learn the underlying distribution of the data. They model the joint probability $P(X, Y)$, or just $P(X)$ if unsupervised. They can understand how the data $X$ was generated.
    *   **Function:** Because they learn $P(X, Y)$, you can use Bayes' rule to calculate $P(Y|X) = \frac{P(X,Y)}{P(X)}$ for classification, OR you can sample from $P(X)$ to generate *new* data.
    *   **Examples:** Naive Bayes, Hidden Markov Models (HMMs), Gaussian Mixture Models (GMMs), GANs, VAEs, Autoregressive LLMs.
    *   **When to use:** When you need to generate new samples, deal with significant missing data, or when you have a lot of unlabeled data and little labeled data (semi-supervised learning). They handle outliers better because they know the "shape" of the normal data.

## 5. Explain L1 vs L2 Regularization mathematically and intuitively. Why does L1 lead to sparsity?
**Answer:**
Regularization aims to prevent overfitting by penalizing complex models. It adds a penalty term to the loss function: $Loss = Unregularized\_Loss + \lambda \cdot Penalty$.

*   **L2 Regularization (Ridge):**
    *   **Penalty:** $\lambda \sum w_i^2$ (Sum of squared weights).
    *   **Intuition:** It heavily penalizes *large* weights because the penalty grows quadratically. It encourages the model to use all features a little bit, rather than relying heavily on one feature. The weights get pushed towards zero but rarely exactly *to* zero.

*   **L1 Regularization (Lasso):**
    *   **Penalty:** $\lambda \sum |w_i|$ (Sum of absolute weights).
    *   **Intuition:** It penalizes weights linearly. Crucially, it acts as a feature selector by driving the weights of less important features exactly to zero.

*   **Mathematical/Geometric reason for Sparsity in L1:**
    *   Think of optimization as finding the point where the contours of the Unregularized Loss function intersect with the boundary of the Regularization penalty.
    *   In 2D (two weights $w_1, w_2$), the L2 penalty region $\lambda(w_1^2 + w_2^2) \le C$ is a **circle**. The loss contours usually hit the circle somewhere along the curve (where both $w_1$ and $w_2$ have non-zero values).
    *   The L1 penalty region $\lambda(|w_1| + |w_2|) \le C$ is a **diamond** (rhombus) with sharp corners exactly on the axes (where one weight is 0). Because it has sharp corners on the axes, the elliptical contours of the loss function are highly likely to hit the constraint region exactly at one of these corners, forcing the other weight to exactly 0.
