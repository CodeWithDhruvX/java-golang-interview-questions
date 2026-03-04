# Applied Machine Learning (Service-Based Companies)

Service-based interviews focus heavily on your ability to wrangle data and push it through standard ML pipelines using Python libraries like `pandas` and `scikit-learn`.

## 1. What are the common steps you take to preprocess a raw dataset before passing it to a Machine Learning model?
**Answer:**
Raw real-world client data is almost never ready for a model. The typical preprocessing steps using tools like Pandas and Scikit-Learn include:

1.  **Handling Missing Data:**
    *   *Drop Rows/Columns:* If a column is 90% null, drop it.
    *   *Imputation:* Fill missing numerical values with the Mean, Median, or a constant (like `-1`). Use Median if there are heavy outliers. For categorical data, impute with the most frequent value (Mode) or a specific "Unknown" category.
2.  **Encoding Categorical Variables:** Models only understand numbers, not text like "Red" or "New York".
    *   *Label Encoding/Ordinal Encoding:* Convert categories to integers (e.g., Small=1, Medium=2, Large=3). Best for ordinal data with a natural hierarchy.
    *   *One-Hot Encoding:* Create dummy columns of 1s and 0s for each category (e.g., `City_NY: 1`, `City_SF: 0`). Best for nominal data where no hierarchy exists.
3.  **Feature Scaling (Normalization/Standardization):** Many algorithms (like KNN, SVMs, or Neural Networks) calculate distances. If "Salary" is $100,000$ and "Age" is $30$, Salary will dominate the calculation.
    *   *Min-Max Scaler (Normalization):* Squashes values strictly between 0 and 1.
    *   *Standard Scaler (Standardization):* Centers data around a mean of 0 with a standard deviation of 1.
4.  **Handling Outliers:** Cap extreme values (using the IQR method or Z-scores) so they don't skew linear models.

## 2. In Scikit-Learn, how do you prevent Data Leakage during preprocessing, especially when using K-Fold Cross Validation?
**Answer:**
**The Pitfall (Data Leakage):**
A very common mistake is scaling the *entire dataset* (e.g., `StandardScaler().fit_transform(X)`) before splitting it into train and test sets. When you scale using the whole dataset, information about the 'Test' set (like the global mean and variance) "leaks" into the training process. The model will look artificially good on the test set, but fail in production on truly unseen data.

**The Solution:**
You must fit your preprocessing objects ONLY on the training data.

1.  `X_train, X_test = train_test_split(...)`
2.  `scaler = StandardScaler()`
3.  `X_train_scaled = scaler.fit_transform(X_train)` -> (**Fit** learns the mean/variance of the *train set only*, **transform** applies it).
4.  `X_test_scaled = scaler.transform(X_test)` -> (Applies the *train set's* math to the test set).

**With Cross-Validation:**
When using `cross_val_score`, applying the above logic manually is tricky because the folds rotate. The correct Scikit-Learn approach is to use a **Pipeline** (`from sklearn.pipeline import Pipeline`).
```python
pipe = Pipeline([('scaler', StandardScaler()), ('svc', SVC())])
cross_val_score(pipe, X, y, cv=5)
```
The Pipeline guarantees that for *every single fold*, the Scaler is strictly fit only on that specific fold's training portion, entirely preventing leakage.

## 3. A client provides a dataset where they want to predict house prices, but there are 500 features. How do you reduce the feature space?
**Answer:**
Feeding 500 features into a model often leads to the "Curse of Dimensionality" (overfitting, slow training, and noise). I would use Dimensionality Reduction or Feature Selection techniques:

*   **1. Statistical Feature Selection (Filter Methods):**
    *   Drop features with very low variance (e.g., if a column is 99% the same value, it holds no predictive power).
    *   Use correlation matrices. If Feature A and Feature B are 95% correlated (e.g., "House Area in Sq Ft" and "House Area in Sq Meters"), drop one of them to prevent multicollinearity.
*   **2. Model-Based Selection (Embedded/Wrapper Methods):**
    *   Train a Random Forest or XGBoost model on all 500 features. These models inherently calculate **Feature Importance**. I would plot the importances, find the "elbow," and drop the 400 least important features.
    *   Use L1 Regularization (Lasso Regression). Lasso actively pushes the coefficients of useless features to exactly zero. You simply keep the non-zero features.
*   **3. Dimensionality Reduction (PCA):**
    *   Principal Component Analysis (PCA) mathematically projects the 500 original features down into a smaller set of completely uncorrelated "Principal Components" (e.g., 20 components) that still capture 95% of the variance in the data. *Drawback:* We lose interpretability. You can no longer tell the client exactly which original feature drove the price.

## 4. How would you explain regular Logistic Regression to a non-technical project manager?
**Answer:**
"Imagine you want to predict whether a customer will click on an advertisement (Yes or No).

If I graph customer data (like their age or income), you can't use a straight line (a regular Linear Regression) to predict a 'Yes' or 'No' because a straight line goes on forever into infinity. It and might give a prediction of '1.5' or '-0.3', which doesn't make sense when we just want a probability between 0% and 100%.

**Logistic Regression** solves this by taking that straight line and 'bending' it into an S-shape curve (called a Sigmoid curve). This curve acts as a ceiling and a floor. No matter how extreme the customer data gets, the curve squashes the prediction so it always falls exactly between 0 and 1.

So, if the model outputs 0.85, we can confidently tell the business there is an 85% probability the user will click, and we can set a threshold (e.g., above 50% = Yes) to make a final decision."

## 5. You trained a RandomForest model and it is overfitting terribly. What hyperparameters would you adjust to fix this?
**Answer:**
A Random Forest naturally handles variance well by combining many trees, but individual deep trees can still overfit. To constrain the model in Scikit-Learn (`RandomForestClassifier` or `RandomForestRegressor`), I would adjust:

*   `max_depth`: Decrease this. The most common cause of overfitting is trees growing too deep and memorizing every data point. Limiting depth acts as severe regularization.
*   `min_samples_split`: Increase this. This dictates the minimum number of samples required to split an internal node. Changing it from 2 to 10 forces the tree to make broader, more generalized splits rather than isolating tiny edge cases.
*   `min_samples_leaf`: Increase this. Ensures that the final end nodes ("leaves") of the tree have a minimum number of samples (e.g., 5). It prevents the tree from creating a leaf for just one specific, noisy data point.
*   `max_features`: Decrease this. This dictates how many random features a tree looks at when deciding how to split. Lowering it (e.g., from `sqrt` to `log2` or a smaller integer) forces the individual trees to be more diverse and uncorrelated, improving the overall ensemble's generalization.
