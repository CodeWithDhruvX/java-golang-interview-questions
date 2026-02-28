# Intermediate Level Python AI Interview Questions

## From 01 Advanced Python for AI & Data Wrangling

# 🟡 **101–120: Advanced Python & Pandas**

### 101. How do decorators help in AI pipelines?
"Decorators are my go-to for adding reusable logic across an AI pipeline without cluttering the core modeling code.

I use them extensively for three things: **logging**, **timing**, and **caching**. 
For example, training functions can take hours. I wrap my `train_model()` function with a custom `@time_it` decorator to log exactly how long execution took. I also use `@functools.lru_cache` on data-loading functions so if I request the same clean dataset twice, it loads instantly from memory instead of hitting the disk again."

#### Indepth
In large pipelines (like Airflow or Prefect), decorators like `@task` physically transform a standard Python function into a node on a Directed Acyclic Graph (DAG). This allows the orchestrator to manage retries, parallel execution, and dependency mapping automatically just by adding a single line above the function definition.

---

### 102. Explain how context managers can be used in data processing.
"Context managers (using the `with` statement) guarantee that a resource is cleanly allocated and then safely destroyed, even if my code crashes in the middle of processing.

Beyond `with open('file.csv')`, I use them to manage **database connections** or **GPU memory limits**.
For instance, `with psycopg2.connect() as conn:` ensures the database lock is released. In deep learning, `with torch.no_grad():` is a context manager telling PyTorch temporarily to stop tracking gradients during inference, which slashes RAM usage by half."

#### Indepth
I frequently write custom context managers using the `@contextlib.contextmanager` decorator. For example, a `with timer("Data Prep"):` manager that logs the start time when entering the block, `yields` execution to the inner code, and calculates the elapsed time automatically when the block exits.

---

### 103. What is duck typing and how does it apply in ML codebases?
"Duck typing means 'If it walks like a duck and quacks like a duck, it is a duck.' Python doesn't care about the *actual type* of an object; it only cares if it has the *methods* we need.

In an ML codebase, this is why Scikit-learn is so beautiful. I can pass a Pandas DataFrame, a NumPy array, or even a standard Python list of lists into `model.fit()`. Scikit-learn doesn't strictly check the type; it just expects the object to behave like a 2D matrix that can be iterated and sliced."

#### Indepth
This heavily drives the design of custom ML estimators. If I build a completely custom neural network but implement the `.fit()`, `.predict()`, and `.score()` methods, I can instantly plug my custom model into Scikit-learn's `GridSearchCV` or `Pipeline` tools. They will treat my deep learning model exactly as if it were a standard logistic regression because it perfectly mimics the expected API.

---

### 104. How can you use `__slots__` to optimize memory in AI apps?
"When I instantiate standard Python classes (like creating 10 million custom `Record` objects for a massive NLP corpus), Python dynamically creates a `__dict__` for every single object to store its attributes. 

This `__dict__` consumes a shocking amount of memory overhead.

By declaring `__slots__ = ['text', 'label']` inside my class definition, I explicitly tell Python, 'This class will *only* ever have these two attributes.' Python ditches the dynamic dictionary and uses a static C-struct instead. This simple line of code can reduce RAM usage by 40-50% for millions of objects."

#### Indepth
While brilliant for memory, `__slots__` sacrifices the flexibility of dynamically attaching new attributes to an object at runtime (`obj.new_feature = 5` will throw an AttributeError). In modern data engineering, if I'm worrying about `__slots__`, I usually realize I shouldn't be using custom Python classes for bulk data anyway, and I switch to Pandas DataFrames or **PyArrow tables**, which are memory-mapped and infinitely more efficient.

---

### 105. How do you create custom exceptions for ML applications?
"Standard exceptions like `ValueError` are often too vague for debugging complex ML pipelines. I create custom exceptions by inheriting from Python's base `Exception` class.

For example, I might define `class ModelNotConvergedError(Exception): pass`. 
If my training loop hits 1000 epochs without the loss dropping, I `raise ModelNotConvergedError("Loss plateaued.")`. 

This allows my high-level orchestration pipeline (like an Airflow DAG) to `except ModelNotConvergedError:` specifically, and automatically trigger a fallback strategy (like restarting training with a different learning rate) rather than crashing the entire system."

#### Indepth
I also create things like `DataDriftDetectedError`. This is vital in production. If the incoming live distribution diverges 20% from the training distribution, throwing a standard `ValueError` tells the DevOps team nothing. Throwing `DataDriftDetectedError` immediately points the investigation toward the data engineering pipeline rather than the infrastructure.

---

### 106. What is the difference between `deepcopy` and `copy` in Python?
"This distinction has caused me numerous debugging headaches.

A `copy` (shallow copy) creates a new object, but it inserts *references* to the child objects found in the original. If I shallow copy a list of lists, modifying a number inside the inner list modifies it in *both* the original and the copy.

A `deepcopy` recursively generates a completely new copy of the object *and all its children*. I use `copy.deepcopy(model_config)` when I want to take a base hyperparameter dictionary, radically alter nested values for an experiment, but ensure the original dictionary remains pristine."

#### Indepth
In PyTorch, duplicating a model architecture requires `copy.deepcopy(model)`. A shallow copy would just create two variables pointing to the exact same neural network weights in GPU memory. Updating weights on the copy would invisibly update the weights on the original.

---

### 107. How does Python's GIL affect multi-threaded ML applications?
"The Global Interpreter Lock (GIL) is Python's infamous bottleneck. It is a mutex that prevents multiple native threads from executing Python bytecodes simultaneously.

Because of the GIL, using the `threading` module for CPU-heavy ML tasks (like running 4 Random Forests in parallel) completely fails; they just take turns running on a single CPU core, offering zero speedup.

However, the GIL is largely irrelevant for I/O bound tasks (like downloading 1,000 images), where threading works beautifully. Crucially, C-extensions like NumPy and PyTorch actively release the GIL when performing heavy math, allowing them to utilize all CPU cores fully despite being called from Python."

#### Indepth
To achieve true parallel CPU execution for pure Python code (like complex custom data augmentation loops), I completely abandon `threading` and use the `multiprocessing` module. This bypasses the GIL by spinning up entirely separate Python processes, each with its own memory space and its own GIL, distributed across multiple physical CPU cores.

---

### 108. What is multiprocessing and when should you use it in AI?
"Multiprocessing spins up completely independent Python processes, utilizing all physical cores on my CPU to achieve true parallelism and bypass the GIL constraint.

I use `from multiprocessing import Pool`. 
If I have 10,000 messy text documents that require heavy Regex cleaning and spaCy lemmatization (pure CPU work), running it in a standard `for` loop might take an hour. By using `Pool(processes=8).map(clean_text, docs)`, I chop the workload into 8 chunks, process them simultaneously on my 8-core CPU, and literally finish the job in 8 minutes."

#### Indepth
Multiprocessing carries heavy overhead because data must be 'pickled' (serialized), sent to the new process, processed, deeply pickled again, and returned to the main process. If the task is incredibly fast but the data payload is massive (like transmitting a 5GB DataFrame to a worker to add two columns), the serialization overhead will make multiprocessing *slower* than a single-threaded loop. It is only useful for computationally dense tasks.

---

### 109. How do you manage memory when working with large datasets in Python?
"When data exceeds my RAM (e.g., a 50GB CSV on a 16GB laptop), I never try to load it into Pandas at once (`pd.read_csv` will just crash).

My primary defense is **Chunking**. `for chunk in pd.read_csv('massive.csv', chunksize=100000):`. I process the data 100K rows at a time, update a running aggregate or write the filtered results back to disk, keeping RAM usage flat.

My second defense is **Dtypes**. Changing default `float64` to `float32` instantly halves memory usage. Converting string columns with repeating values (like 'State') into Pandas `category` types drastically shrinks the memory footprint by storing integers instead of raw text strings."

#### Indepth
For truly massive datasets in a production environment, I abandon Pandas entirely and switch to **Dask** or **Polars**. Dask provides a Pandas-like API but secretly orchestrates operations out-of-core (spilling to disk when RAM is full) and parallelizes computations across multiple cores automatically.

---

### 110. What are type hints and how do they help in large AI codebases?
"Type hints (`def predict(text: str) -> float:`) do not affect Python's runtime; it remains dynamically typed. However, they are a superpower for developer experience.

In a large ML codebase with 10 engineers, passing undocumented tensors around is a nightmare. Does `process(data)` expect a Pandas DataFrame, a Numpy array, or a PyTorch tensor? Type hints (`data: np.ndarray`) solve this instantly.

They allow my IDE (VSCode/PyCharm) to provide spectacular auto-complete and instantly highlight bugs before I ever run the code. Integrating `mypy` into our CI/CD pipeline ensures no type-mismatch bugs ever reach production."

#### Indepth
In advanced ML architectures, specifying generic shapes is becoming standard. While `np.ndarray` is helpful, using specialized typing libraries (like `jaxtyping`) allows me to specify exact tensor dimensions: `def forward(x: Float[Tensor, "batch channels height width"]) -> Float[Tensor, "batch classes"]:`. This makes debugging massive matrix multiplication shape mismatch errors 100x easier.

---

### 111. How do you pivot and unpivot data in Pandas?
"Pivoting turns vertical, 'long' data into horizontal, 'wide' data. Unpivoting does the reverse.

If I have transactional data (Date, Store, Sales), and I want a report showing Stores as rows and Dates as columns, I use `df.pivot_table(index='Store', columns='Date', values='Sales', aggfunc='sum')`.

To unpivot (e.g., turning those Date columns back into a single 'Date' column for feeding into an ML model), I use `df.melt(id_vars=['Store'], var_name='Date', value_name='Sales')`."

#### Indepth
`pivot()` and `pivot_table()` are different. `pivot()` fails instantly if there are duplicate entries for the index/column intersection because it doesn't know which value to keep. `pivot_table()` is vastly superior because it incorporates an aggregation function (like `mean` or `sum`) to logically combine duplicates instantly.

---

### 112. How do you handle multi-index DataFrames?
"When I run a `groupby` on multiple columns, Pandas creates a Hierarchical Index (MultiIndex). For example, row labels might be ('New York', '2023').

While powerful, MultiIndexes are notoriously annoying to slice and feed into Scikit-learn. 
My standard practice is to immediately flatten them. By calling `df.reset_index()`, Pandas converts the nested indices right back into standard, flat columns.

If I *must* keep them for complex aggregations, I use the `.xs()` (cross-section) method, which allows me to slice down cleanly through specific levels, like extracting all data for '2023' regardless of the city."

#### Indepth
A major pain point arrives when columns become multi-indexed after doing a `.agg(['mean', 'sum'])`. The columns become tuples like `('Salary', 'mean')`. To flatten this for exporting or modeling, I iterate and join the tuples: `df.columns = ['_'.join(col) for col in df.columns]`, resulting in clean `Salary_mean` style columns.

---

### 113. What are rolling windows and how are they useful in ML?
"Rolling windows are the core of time-series feature engineering.

If I want to predict stock prices tomorrow, knowing the price today is good, but knowing the trend is better. 
I use `df['Price'].rolling(window=7).mean()` to create a 7-day Moving Average feature. 

I use `.std()` to calculate rolling volatility. These rolling aggregations capture the recent 'momentum' of the data, providing drastically stronger predictive signals to an ML model than isolated raw data points."

#### Indepth
A vital trap in rolling windows is **Data Leakage**. By default, `rolling(7).mean()` at Row 7 calculates the average of Rows 1 through 7. If I use this feature to *predict* Row 7, I have leaked the answer into the training data. I absolutely must append a `.shift(1)` to ensure the rolling average available at Row 7 only contains data from Rows 1 through 6.

---

### 114. How do you merge datasets with mismatched keys?
"Real-world databases are messy. If Table A has `user_id='AB123'` but Table B has `client_id='ab-123'`, a standard `pd.merge()` completely fails.

First, I normalize the keys into a consistent format using text manipulation: I lowercase everything and strip dashes: `df['key'] = df['key'].str.lower().str.replace('-', '')`. Then I perform the merge.

If the keys are simply misspelled (e.g., 'Jon Doe' vs 'John Doe'), exact merging is impossible. I pull in fuzzy matching libraries like `FuzzyWuzzy` or `RapidFuzz` to calculate the Levenshtein distance between strings, assigning matches if the similarity score is extremely high (e.g., >95%)."

#### Indepth
For merging time-series events where timestamps don't align perfectly (e.g., linking a 'click' event at 10:01:05 with a 'purchase' event at 10:01:45), exact joins fail. I rely heavily on `pd.merge_asof()`. It performs a 'nearest' merge, allowing me to specify parameters like "Match the purchase to the most recent click, provided it occurred within the last 5 minutes (`tolerance=pd.Timedelta('5m')`)."

---

### 115. How do you convert categorical columns into binary efficiently?
"Converting categories to numbers is mandatory because ML models only ingest math. This is One-Hot Encoding.

In Pandas during Exploratory Data Analysis, I use `pd.get_dummies(df, columns=['Color'])`. It instantly converts the 'Color' column into 'Color_Red', 'Color_Blue', filled with 1s and 0s.

However, in an ML production pipeline, I strictly use Scikit-learn's `OneHotEncoder(handle_unknown='ignore')`. If my web app encounters a brand new 'Color_Purple' during live deployment that wasn't in the training data, `get_dummies` crashes the pipeline. `OneHotEncoder` safely ignores it, outputting all zeros for that specific variable so the model proceeds without crashing."

#### Indepth
If a categorical column has 500 distinct zip codes, One-Hot Encoding creates 500 sparse, mostly-zero columns. This destroys tree-model performance and spikes memory (the Curse of Dimensionality). For high-cardinality categories, I switch to **Target Encoding** (replacing the category with the average target value of that category) or use **Embedding Layers** in deep learning.

---

### 116. How do you resample time-series data in Pandas?
"Resampling allows me to change the frequency of my time-series data, acting exactly like a timeline `groupby`. The DataFrame index *must* be a datetime object.

If I possess data recorded every minute, but I want to predict daily aggregates, I downsample using `df.resample('D').sum()`. The 'D' stands for Daily, and it aggregates all 1,440 minute-rows into a single daily sum.

If I have monthly data but need it interpolated to daily (upsampling), I use `df.resample('D').ffill()`, which propagates the monthly value forward filling the brand new blank days."

#### Indepth
When resampled data generates blank periods (NaNs), filling strategies depend entirely on business logic. Forward fill (`ffill`) assumes the last known value holds true (like a bank balance). Interpolation (`interpolate(method='linear')`) draws a straight line between the known points (good for temperature averages). Never interpolate directly into the future, as that creates massive data leakage in predictive models.

---

### 117. How to detect outliers in data using statistical methods?
"Outliers destroy the mathematical assumptions of algorithms like Linear Regression.

The most common method is the **Z-Score**. Assuming the data is normally distributed, I calculate how many standard deviations a point is away from the mean: `(Value - Mean) / StdDev`. Any absolute Z-score strictly greater than 3 is flagged as an extreme outlier.

For data that is heavily skewed (not a bell curve), the Z-score is useless. I use the **IQR (Interquartile Range) Method**. I find the 25th (Q1) and 75th (Q3) percentiles. The IQR is Q3 - Q1. Any value falling below `Q1 - 1.5 * IQR` or above `Q3 + 1.5 * IQR` is an outlier. This is exactly what a Seaborn boxplot uses to draw 'whiskers'."

#### Indepth
Statistical methods look at one variable isolation. But what if $500,000 for a house isn't an outlier, and 1 bedroom isn't an outlier, but a $500,000 1-bedroom house *is* an outlier? This requires multivariate detection algorithms. I use Scikit-Learn's **Isolation Forest**, which builds random trees designed explicitly to isolate anomalous data points rapidly based on their multiple dimensions simultaneously.

---

### 118. How to handle skewed data before training a model?
"Highly right-skewed data (like Income: mostly $50k, but a few billionaires at $10B) ruins regression models because the billionaires pull the entire prediction line towards them.

The most standard fix is applying a **Log Transformation**: `df['Income'] = np.log1p(df['Income'])` (using `log1p` to handle zeros safely). This mathematically squashes the massive outlier tails while preserving the rank order, transforming the skewed distribution into something much closer to a normal, symmetrical bell curve."

#### Indepth
If the data contains negative values, log transforms instantly fail (yielding NaNs). The modern ML standard is the **Yeo-Johnson Power Transformation** (available in `sklearn.preprocessing.PowerTransformer`). It mathematically determines the optimal $\lambda$ parameter to normalize the curve, supporting both strictly positive and negative numerical ranges seamlessly.

---

### 119. How to vectorize a function for use with NumPy arrays?
"Vectorization is the process of eliminating slow Python `for` loops by pushing the computation down to highly optimized C code within NumPy.

If I write a custom Python function `def custom_math(x): ...`, and I want to apply it to a massive array, `[custom_math(i) for i in arr]` is terribly slow.
I wrap it using `vectorized_math = np.vectorize(custom_math)`. Then I can just call `vectorized_math(arr)`.

This allows my complex custom logic to execute significantly faster across the entire array matrix."

#### Indepth
`np.vectorize` is syntactic sugar; it essentially runs a glorified C-loop under the hood, making it slightly faster, but it isn't true parallel vector math. To achieve staggering performance gains, I avoid `np.vectorize` entirely and rewrite my algorithm logically utilizing core NumPy primitive operators (e.g., using boolean masking like `arr[arr > 0] = 1`) which leverage native SIMD (Single Instruction Multiple Data) registers on the CPU.

---

### 120. What is memory-mapping in NumPy and when to use it?
"Memory-mapping (`np.memmap`) is my lifeline when visual or tabular data completely exceeds available RAM (e.g., a 100GB matrix on a 16GB machine).

Instead of loading the array into memory, `np.memmap` creates a lightweight virtual pointer to the massive binary file on the hard drive. I can slice it like a totally normal array: `chunk = mmap_array[100:200]`. 

NumPy only loads *that tiny specific chunk* into RAM from the disk instantly. When I update the slice, it writes directly back to the disk. It allows Deep Learning models to stream massive matrices flawlessly."

#### Indepth
While brilliant for avoiding OOM crashes, `memmap` relies on the OS virtual memory page cache, which is terrifyingly slow if the data is on a spinning HDD. If the access patterns are entirely random, it causes "thrashing". `memmap` operates best when slicing and reading sequential continuous blocks of data stored on high-speed NVMe SSDs.

---
EOF
