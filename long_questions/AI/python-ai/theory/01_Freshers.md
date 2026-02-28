# Freshers Level Python AI Interview Questions

## From 01 Basic Python for AI & NumPy/Pandas

# 🟢 **1–20: Basics & Data Manipulation**

### 1. What are Python's key features that make it suitable for AI?
"Python is my absolute go-to for AI primarily because of its **simplicity** and **massive ecosystem**.

The syntax is incredibly clean, allowing me to focus on creating models rather than wrestling with boilerplate code. Most importantly, it boasts libraries like NumPy, Pandas, Scikit-learn, and PyTorch, which abstract away complex math. 

I also love its interactivity; using Jupyter Notebooks lets me test ideas, visualize data, and document findings all in one place."

#### Indepth
While Python is inherently slow due to the Global Interpreter Lock (GIL) and being interpreted, AI libraries bypass this. NumPy and TensorFlow are built on C/C++ or CUDA. Python acts strictly as a "glue" or orchestrator language, combining researcher-friendly syntax with bare-metal execution speed.

---

### 2. How do lists differ from NumPy arrays?
"To me, the biggest difference comes down to **performance and memory**. 

A standard Python `list` can hold different data types, but this flexibility makes it slow because every element is a separate object in memory. A NumPy array (`ndarray`), on the other hand, requires all elements to be the same type.

I use NumPy arrays for almost all numeric data because they store elements contiguously in memory, making mathematical operations significantly faster."

#### Indepth
NumPy arrays support **vectorization**. If I want to add 5 to every element in a list, I need a `for` loop. In NumPy, I just write `arr + 5`. This pushes the loop down to the C level, executing in parallel without Python loop overhead, yielding massive performance gains.

---

### 3. What are lambda functions and how are they useful in AI?
"A **lambda function** is a small, anonymous, one-line function defined without the `def` keyword.

I find them incredibly useful in AI pipelines for quick data transformations where writing a full function feels like overkill. For example, I might use a lambda to quickly scale variables or extract specific substrings while prepping a Pandas dataset.

They keep the code compact and readable, especially when passed as arguments to other functions like `apply()` or `map()`."

#### Indepth
Lambdas are restricted to a single expression and cannot contain multiple statements or assignments. In Pandas, using `df['col'].apply(lambda x: math.log(x))` is an idiomatic pattern, although vectorization (like `np.log(df['col'])`) is statistically preferred over apply/lambda for performance reasons.

---

### 4. What is the use of `map()`, `filter()`, and `reduce()` in AI workflows?
"These are functional programming tools I use for clean data manipulation.

`map()` applies a function to every item in an iterable. `filter()` keeps only items satisfying a condition. `reduce()` (from `functools`) continuously applies a function to pairs of elements until a single value remains.

In AI workflows, I might use `filter()` to remove malformed records before ingestion, and `map()` to quickly format a list of strings before tokenizing them."

#### Indepth
In modern Python, I tend to replace `map()` and `filter()` with **list comprehensions** because they are often more readable and slightly faster. However, when working with enormous distributed datasets (like in Spark/PySpark implementations), the map-reduce paradigm remains the foundational approach for parallel computation.

---

### 5. How do Python generators work?
"A generator is a special type of function that returns a **lazy iterator**.

Instead of calculating a massive array of values and returning them all at once via `return`, a generator uses the `yield` keyword. It yields one value, pauses its state, and resumes right where it left off the next time I ask for a value.

I use generators constantly in AI when dealing with datasets that are larger than my system's RAM, like streaming thousands of images for training deep learning models."

#### Indepth
Under the hood, generators avoid memory exhaustion. Loading a 10GB dataset into a list crashes the system. Creating a generator `def read_chunks(file):` that yields 10MB batches keeps RAM usage flat. Keras' `ImageDataGenerator` and PyTorch's `DataLoader` rely entirely on this yielding concept.

---

### 6. Explain the difference between `@staticmethod` and `@classmethod`.
"A `@classmethod` takes the class itself (`cls`) as its first argument automatically. I use it mostly as an alternative constructor, for example, `Model.from_pretrained('bert')`.

A `@staticmethod` doesn't take the instance (`self`) or the class (`cls`). It acts like a plain function that just happens to live inside the class namespace for organizational purposes.

I use static methods for pure utility functions, like `DataPreprocessor.clean_text(text)`, which don't need access to any class state."

#### Indepth
Class methods are crucial for inheritance. If a subclass calls a `@classmethod` inherited from a parent, `cls` will correctly point to the subclass, allowing the factory method to return an instance of the specific subclass instead of hardcoding the parent class.

---

### 7. What is the use of `with` statement in Python?
"The `with` statement is my primary tool for resource management, typically file operations or database connections.

It sets up an execution block and guarantees that cleanup happens automatically when the block is exited, even if an error occurs inside. `with open('data.csv') as f:` ensures the file is closed automatically.

This saves me from writing explicit `try...finally` blocks and minimizes memory leaks."

#### Indepth
The `with` statement works through the **Context Manager** protocol, invoking `__enter__()` and `__exit__()` dunder methods dynamically. In ML, TensorFlow relies on this pattern extensively, such as `with tf.GradientTape() as tape:`, where the context manager explicitly tracks operations for automatic differentiation.

---

### 8. What are Python’s data serialization formats used in ML?
"I routinely use **Pickle** and **JSON**.

`json` is human-readable and perfect for exporting configuration dictionaries or model hyperparameters.

`pickle` is Python's native binary format. I use it to save trained Scikit-learn models or complex Python objects to disk. However, it's strictly Python-to-Python."

#### Indepth
Pickle is notoriously **insecure**, as unpickling can execute arbitrary code. I never unpickle files from untrusted sources. For large-scale ML data, I completely replace Pickle with formats like **Parquet** (columnar storage) or **HDF5 / Safetensors** for saving neural network weights safely and efficiently across languages.

---

### 9. How does list comprehension differ from generator expressions?
"A list comprehension `[x*2 for x in data]` computes all the values instantly and stores them in a new list in memory. 

A generator expression `(x*2 for x in data)` looks almost identical but uses parentheses. Instead of doing the work upfront, it creates a generator object that evaluates values on-the-fly.

If my `data` has a million items, the list comprehension will spike my RAM, while the generator expression uses practically zero memory."

#### Indepth
Generator expressions use "lazy evaluation." Time to execution for the first element is instant, but iterating entirely through the generator takes slightly longer than iterating over an already instantiated list. I default to list comprehensions for small data and generator expressions strictly for large pipelines.

---

### 10. How do you handle missing data in a dataset using Python?
"Missing data is a daily ML problem. I primarily use Pandas to handle it.

First, I identify missing values using `df.isnull().sum()`. The action I take depends on the data:
1. If only a few rows are missing, I just drop them via `df.dropna()`.
2. If a column represents continuous data, I impute it with the mean or median using `df.fillna()`.
3. If categorical, I might fill with the mode or create a separate 'Missing' category."

#### Indepth
Simple imputation can reduce data variance artificially. For critical features, I utilize more advanced techniques like `IterativeImputer` from `sklearn.impute`, which models the missing feature based on all other features. For time-series, I rely on `df.interpolate()` or forward/backward filling (`ffill`).

---

### 11. How does broadcasting work in NumPy?
"Broadcasting is NumPy's way of doing math between arrays of different shapes.

Instead of writing loops to pad or repeat the smaller array, NumPy automatically 'broadcasts' the smaller dimension across the larger one so they have compatible shapes.

If I try to add a scalar `10` to a matrix `A`, NumPy virtually stretches `10` into a matrix of the same size as `A` and performs element-wise addition. It makes my code significantly cleaner."

#### Indepth
Broadcasting requires dimensions to be compatible: starting from trailing dimensions, they must either be equal, or one of them must be 1. The operation happens purely in C without making literal copies of data, which means it’s exponentially faster and more memory-efficient than doing the equivalent logic manually in Python.

---

### 12. What is the difference between `loc[]` and `iloc[]` in Pandas?
"This is a distinction I use constantly.

`loc[]` is **label-based**. I use it when I want to select rows or columns by their actual index names or column names (e.g., `df.loc[:, 'Age']`).

`iloc[]` is **integer-position-based**. I use it when I want to slice by numerical position, exactly like a Python list (e.g., `df.iloc[0:5, 2]`), regardless of what the index labels are."

#### Indepth
A common trap is slicing. With `iloc[0:5]`, the result includes indices 0 up to 4 (exclusive of 5), mimicking Python lists. However, with `loc['A':'E']`, the slicing is **inclusive**, returning rows 'A' through 'E', which can cause off-by-one logical bugs if you mistake the behavior.

---

### 13. How to normalize a NumPy array?
"Normalization usually means scaling values to a range between 0 and 1. 

In NumPy, I do this manually using the formula: `(arr - np.min(arr)) / (np.max(arr) - np.min(arr))`.

For neural networks, standardizing the array (mean of 0, standard deviation of 1) is often better. I calculate that as `(arr - np.mean(arr)) / np.std(arr)`. This puts variables on the same playing field for gradient descent."

#### Indepth
While doing this manually is easy, in standard ML workflows I prefer using `MinMaxScaler` or `StandardScaler` from Scikit-learn. The scaler "fits" (learns the min/max) on the training set and "transforms" the test set, preventing **data leakage**. Manually normalizing entire arrays at once before splitting the data is a critical ML mistake.

---

### 14. What is the difference between `axis=0` and `axis=1` in Pandas?
"`axis=0` refers to the **rows** (the index). If I compute `df.mean(axis=0)`, it collapses the rows and calculates the mean downwards, giving me the average of each column.

`axis=1` refers to the **columns**. If I use `df.drop('Price', axis=1)`, it looks across the columns and drops that specific one.

I always visualize it as: axis=0 moves vertically (↓), axis=1 moves horizontally (→)."

#### Indepth
The nomenclature confuses almost every beginner because it feels inverted. A good mnemonic is that `axis=0` means the operation is applied *along the 0th dimension* (changing the number of rows). Using `axis='index'` or `axis='columns'` is supported in Pandas and I prefer it in production code for maximum readability.

---

### 15. How do you merge two DataFrames with `concat()` vs `merge()`?
"I use `concat()` when I just want to 'stack' or 'glue' DataFrames together. If I have sales data from Jan and Feb, `pd.concat([df_jan, df_feb], axis=0)` just places Feb underneath Jan.

I use `merge()` when I need SQL-like relational joins. If I have a `users` DataFrame and an `orders` DataFrame, I use `pd.merge(users, orders, on='user_id', how='left')` to match records intelligently based on a key."

#### Indepth
`concat` performs rapid indexing alignments but expects structural similarity. `merge` uses heavy hash-based or sort-merge algorithms for foreign key resolution. By default, `merge` runs an 'inner' join. For complex time-series joins (matching close but not exact timestamps), I rely on `pd.merge_asof()`.

---

### 16. How do you perform group-wise operations in Pandas?
"I use the `groupby()` method. It follows the **Split-Apply-Combine** principle.

First, I split: `grouped = df.groupby('Department')`. Then, I apply my operation and automatically combine the results, like `grouped['Salary'].mean()`.

It is incredibly powerful for aggregating data quickly without writing iteration loops."

#### Indepth
The true power of `groupby` is exposed with the `.agg()` function, allowing me to compute multiple statistics at once on different columns cleanly: `df.groupby('Dept').agg({'Salary': ['mean', 'max'], 'Age': 'median'})`. This is an essential step during Exploratory Data Analysis (EDA).

---

### 17. How do you handle time series data in Pandas?
"Pandas was originally built for financial time series. First, I always convert my date columns into actual datetime objects using `pd.to_datetime()`.

Then, I usually set that datetime column as the DataFrame's index.

This unlocks powerful time-based slicing (e.g., `df['2023-01':'2023-06']`), shifting (`df.shift(1)` for lag features), and my favorite: `resample()`, which lets me aggregate daily data into weekly or monthly sums easily."

#### Indepth
Dealing with timezones is notoriously difficult, but setting `df.tz_localize()` handles it. An important step in ML is addressing gaps in time series. I use `df.asfreq('D')` to force a daily frequency, which injects NaN rows for missing days, making it obvious where I need to apply interpolation logic.

---

### 18. How do you drop duplicates in Pandas?
"I use the straightforward `df.drop_duplicates()` method.

By default, it looks at all columns and drops completely identical rows, keeping the first occurrence. 

If I only want uniqueness based on specific keys (like preventing the same user from appearing twice), I specify `subset=['user_id']`. Usually, I tack on `keep='last'` if newer entries appear at the bottom of the DataFrame."

#### Indepth
Often, purely dropping duplicates is not correct. E.g., if a user makes two distinct payments on the same timestamp, dropping one deletes financial data. I always pair `duplicated()` analysis with contextual business logic prior to bluntly executing a drop to ensure data integrity.

---

### 19. How do you replace missing values with median in Pandas?
"It’s a two-step logic combined into one chain.

I use `df['Column'].fillna(df['Column'].median(), inplace=True)`.

The `.median()` computes the 50th percentile, ignoring missing values by default. `fillna()` then finds all the `NaN` records and injects this computed median. I use median instead of mean when the data has heavy outliers (like income data) that would skew the average."

#### Indepth
If I need to impute medians *per category*, `transform` is essential. Instead of a single flat median, I use `df['A'] = df['A'].fillna(df.groupby('Category')['A'].transform('median'))`. This ensures that a missing "Engineering" salary is filled with the Engineering median, not the company-wide flat median.

---

### 20. How to compute correlation between features using Pandas?
"I use the simple `df.corr()` command.

It immediately generates a square matrix showing the correlation coefficients between all numeric columns. If a value is near 1, they are perfectly positively correlated; near -1, perfectly negatively correlated.

I heavily rely on this during feature selection to drop highly collinear features (redundant info) and identify which features correlate most strongly with the target variable."

#### Indepth
By default, `corr()` utilizes the **Pearson** correlation coefficient, which only measures linear relationships. If I am dealing with monotonic non-linear relationships or ordinal categories, I explicitly pass `method='spearman'` to calculate rank-based correlations. Typically, I pair the result with `sns.heatmap()` to spot issues visually.

---
