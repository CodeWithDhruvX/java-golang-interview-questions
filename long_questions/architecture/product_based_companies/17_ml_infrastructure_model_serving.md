# 🤖 ML Infrastructure & Model Serving

> **Focus:** Product-Based Companies (Google, Meta, Amazon, Netflix, Spotify)
> **Level:** 🔴 Senior – 🟠 Staff

---

## 📋 Table of Contents

1. [ML System Architecture](#1-ml-system-architecture)
2. [Model Training Infrastructure](#2-model-training-infrastructure)
3. [Feature Engineering & Stores](#3-feature-engineering--stores)
4. [Model Serving Patterns](#4-model-serving-patterns)
5. [Real-time Inference](#5-real-time-inference)
6. [Batch Inference & Offline Processing](#6-batch-inference--offline-processing)
7. [ML Monitoring & Observability](#7-ml-monitoring--observability)
8. [MLOps & CI/CD for ML](#8-mlops--cicd-for-ml)
9. [Scaling ML Systems](#9-scaling-ml-systems)
10. [Common Interview Questions](#10-common-interview-questions)

---

## 1. ML System Architecture

### Q1: What are the key components of a production ML system?

**Answer:**

**Core Components:**
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Data Source   │───▶│ Feature Store    │───▶│ Model Training  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                                                        │
                                                        ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Monitoring    │◀───│ Model Serving    │◀───│ Model Registry  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

**Detailed Architecture:**
- **Data Pipeline:** ETL/ELT processes for data collection
- **Feature Store:** Centralized feature engineering and storage
- **Training Infrastructure:** Distributed training compute
- **Model Registry:** Version control for models
- **Serving Layer:** Real-time and batch inference
- **Monitoring:** Drift detection, performance tracking

### Q2: Explain online vs offline ML systems

**Answer:**

**Online ML Systems:**
```java
@RestController
public class OnlineMLController {
    
    @Autowired
    private ModelService modelService;
    
    @Autowired
    private FeatureStore featureStore;
    
    @PostMapping("/predict")
    public Prediction predict(@RequestBody Request request) {
        // Real-time feature generation
        Features features = featureStore.getFeatures(request.getUserId());
        
        // Real-time inference
        Prediction prediction = modelService.predict(features);
        
        // Log for monitoring
        monitoringService.logPrediction(request, prediction);
        
        return prediction;
    }
}
```

**Characteristics:**
- **Low latency:** Milliseconds response time
- **Real-time features:** Up-to-date user behavior
- **High availability:** 99.9%+ uptime required
- **Use Cases:** Recommendation, fraud detection, search ranking

**Offline ML Systems:**
```python
# Batch processing pipeline
def batch_inference():
    # Load model
    model = load_model_from_registry(model_version)
    
    # Process batch data
    batch_data = load_batch_data(date=yesterday)
    
    # Generate predictions
    predictions = model.predict(batch_data)
    
    # Store results
    save_predictions(predictions, destination="s3://predictions/")
```

**Characteristics:**
- **High throughput:** Process millions of records
- **Complex features:** Heavy feature engineering possible
- **Cost-effective:** Cheaper compute resources
- **Use Cases:** User segmentation, content analysis, reporting

---

## 2. Model Training Infrastructure

### Q3: How do you design distributed training infrastructure?

**Answer:**

**Data Parallel Training:**
```python
# PyTorch Distributed Data Parallel
import torch.distributed as dist
import torch.multiprocessing as mp
from torch.nn.parallel import DistributedDataParallel as DDP

def train_worker(rank, world_size):
    # Initialize distributed process group
    dist.init_process_group("nccl", rank=rank, world_size=world_size)
    
    # Create model and move to GPU
    model = YourModel().to(rank)
    ddp_model = DDP(model, device_ids=[rank])
    
    # Distributed data loader
    train_dataset = YourDataset()
    train_sampler = torch.utils.data.distributed.DistributedSampler(
        train_dataset, num_replicas=world_size, rank=rank
    )
    train_loader = torch.utils.data.DataLoader(
        train_dataset, batch_size=32, sampler=train_sampler
    )
    
    # Training loop
    for epoch in range(epochs):
        for batch in train_loader:
            output = ddp_model(batch)
            loss = compute_loss(output, batch)
            loss.backward()
            optimizer.step()
            optimizer.zero_grad()
```

**Kubernetes Training Job:**
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: distributed-training
spec:
  parallelism: 4  # Number of workers
  completions: 4
  template:
    spec:
      containers:
      - name: training-worker
        image: ml-training:latest
        resources:
          requests:
            nvidia.com/gpu: 1
          limits:
            nvidia.com/gpu: 1
        env:
        - name: WORLD_SIZE
          value: "4"
        - name: RANK
          valueFrom:
            fieldRef:
              fieldPath: metadata.annotations['batch.kubernetes.io/job-completion-index']
      restartPolicy: OnFailure
```

**Training Infrastructure Components:**
- **Compute Clusters:** GPU instances (NVIDIA A100, V100)
- **Storage:** High-performance storage for datasets
- **Orchestration:** Kubernetes, Slurm, or cloud-specific
- **Experiment Tracking:** MLflow, Weights & Biases

### Q4: What is hyperparameter optimization at scale?

**Answer:**

**Distributed Hyperparameter Search:**
```python
import optuna
from optuna.integration import DaskStorage

def objective(trial):
    # Define hyperparameter search space
    learning_rate = trial.suggest_loguniform('lr', 1e-5, 1e-1)
    batch_size = trial.suggest_categorical('batch_size', [32, 64, 128])
    num_layers = trial.suggest_int('num_layers', 2, 8)
    
    # Train model with these hyperparameters
    model = create_model(num_layers)
    accuracy = train_model(model, learning_rate, batch_size)
    
    return accuracy

# Distributed optimization
storage = DaskStorage()
study = optuna.create_study(direction='maximize', storage=storage)
study.optimize(objective, n_trials=1000, n_jobs=50)
```

**Bayesian Optimization:**
```python
from skopt import gp_minimize
from skopt.space import Real, Integer, Categorical

space = [
    Real(1e-5, 1e-1, name='learning_rate', prior='log-uniform'),
    Integer(2, 8, name='num_layers'),
    Categorical(['adam', 'sgd', 'rmsprop'], name='optimizer'),
    Integer(32, 256, name='batch_size')
]

@use_named_args(space)
def objective(learning_rate, num_layers, optimizer, batch_size):
    model = create_model(num_layers, optimizer)
    return -train_and_evaluate(model, learning_rate, batch_size)

result = gp_minimize(objective, space, n_calls=100, n_random_starts=10)
```

---

## 3. Feature Engineering & Stores

### Q5: What is a feature store and why is it needed?

**Answer:**

**Feature Store Architecture:**
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Raw Data      │───▶│ Feature Store    │───▶│  Training       │
│ (Events, Logs)  │    │  - Online Store  │    │  - Batch Jobs   │
└─────────────────┘    │  - Offline Store │    │  - Real-time    │
                       │  - Metadata      │    └─────────────────┘
                       └──────────────────┘              │
                                │                        ▼
                                ▼                ┌─────────────────┐
┌─────────────────┐    ┌──────────────────┐    │   Inference    │
│   Monitoring    │◀───│ Feature Serving   │◀───│   - Online     │
│   - Drift       │    │   - Low latency   │    │   - Batch      │
│   - Quality     │    │   - Consistency   │    │   - Streaming  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

**Key Benefits:**
- **Feature Reusability:** Same features for training and serving
- **Consistency:** Prevents training-serving skew
- **Low Latency:** Optimized for real-time serving
- **Versioning:** Track feature evolution over time

**Implementation Example:**
```java
@Service
public class FeatureStoreService {
    
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;  // Online store
    
    @Autowired
    private JdbcTemplate jdbcTemplate;  // Offline store
    
    public Features getOnlineFeatures(String userId) {
        // Real-time feature lookup
        String key = "features:" + userId;
        return (Features) redisTemplate.opsForValue().get(key);
    }
    
    public List<Features> getOfflineFeatures(List<String> userIds) {
        // Batch feature lookup for training
        String sql = "SELECT * FROM features WHERE user_id IN (?)";
        return jdbcTemplate.query(sql, new Object[]{userIds}, new FeatureRowMapper());
    }
    
    public void updateFeatures(String userId, Features features) {
        // Update both online and offline stores
        redisTemplate.opsForValue().set("features:" + userId, features);
        
        String sql = "INSERT INTO features (user_id, feature_data) VALUES (?, ?) " +
                    "ON CONFLICT (user_id) DO UPDATE SET feature_data = ?";
        jdbcTemplate.update(sql, userId, features.toJson(), features.toJson());
    }
}
```

### Q6: How do you handle feature engineering for real-time vs batch?

**Answer:**

**Real-time Feature Engineering:**
```java
@Component
public class RealTimeFeatureProcessor {
    
    @KafkaListener(topics = "user-events")
    public void processUserEvent(UserEvent event) {
        // Calculate real-time features
        Features features = calculateRealTimeFeatures(event);
        
        // Update online feature store
        featureStore.updateFeatures(event.getUserId(), features);
        
        // Trigger real-time inference if needed
        if (shouldTriggerInference(event)) {
            Prediction prediction = modelService.predict(features);
            eventPublisher.publishPrediction(event.getUserId(), prediction);
        }
    }
    
    private Features calculateRealTimeFeatures(UserEvent event) {
        Features features = new Features();
        
        // Time-based features
        features.setLastActivityTime(event.getTimestamp());
        features.setDaysSinceLastActivity(calculateDaysSince(event.getTimestamp()));
        
        // Behavioral features
        features.setRecentClickCount(getRecentClickCount(event.getUserId(), Duration.ofHours(1)));
        features.setSessionDuration(calculateSessionDuration(event.getUserId()));
        
        // Contextual features
        features.setDeviceType(event.getDeviceType());
        features.setLocation(event.getLocation());
        
        return features;
    }
}
```

**Batch Feature Engineering:**
```python
# Spark job for batch feature computation
def compute_batch_features(date):
    # Load raw data
    events_df = spark.read.parquet(f"s3://events/date={date}")
    
    # Compute user engagement features
    user_features = events_df.groupBy("user_id").agg(
        count("*").alias("total_events"),
        avg("session_duration").alias("avg_session_duration"),
        max("timestamp").alias("last_activity"),
        collect_list("event_type").alias("event_types")
    )
    
    # Compute temporal features
    temporal_features = compute_temporal_features(user_features, date)
    
    # Merge with historical features
    final_features = user_features.join(temporal_features, "user_id")
    
    # Save to offline feature store
    final_features.write.parquet(f"s3://features/date={date}")
    
    # Update online store for active users
    active_users = final_features.filter(col("last_activity") > date_sub(current_date(), 7))
    update_online_store(active_users)
```

---

## 4. Model Serving Patterns

### Q7: What are different model serving patterns?

**Answer:**

**REST API Serving:**
```java
@RestController
@RequestMapping("/api/v1/predict")
public class PredictionController {
    
    @Autowired
    private ModelService modelService;
    
    @PostMapping
    public ResponseEntity<PredictionResponse> predict(@RequestBody PredictionRequest request) {
        try {
            // Validate input
            validateRequest(request);
            
            // Convert to model input format
            ModelInput input = convertToModelInput(request);
            
            // Get prediction
            Prediction prediction = modelService.predict(input);
            
            // Convert response
            PredictionResponse response = convertToResponse(prediction);
            
            // Log for monitoring
            monitoringService.logPrediction(request, response);
            
            return ResponseEntity.ok(response);
            
        } catch (Exception e) {
            return ResponseEntity.badRequest().body(
                new PredictionResponse("error", e.getMessage())
            );
        }
    }
}
```

**gRPC Serving:**
```protobuf
// prediction.proto
service PredictionService {
    rpc Predict(PredictionRequest) returns (PredictionResponse);
    rpc BatchPredict(stream PredictionRequest) returns (stream PredictionResponse);
}

message PredictionRequest {
    string user_id = 1;
    repeated Feature features = 2;
    map<string, string> metadata = 3;
}

message PredictionResponse {
    float prediction = 1;
    float confidence = 2;
    map<string, float> probabilities = 3;
}
```

```java
@Service
public class PredictionServiceImpl extends PredictionServiceGrpc.PredictionServiceImplBase {
    
    @Override
    public void predict(PredictionRequest request, StreamObserver<PredictionResponse> responseObserver) {
        try {
            // Process prediction
            PredictionResponse response = modelService.predict(request);
            
            // Send response
            responseObserver.onNext(response);
            responseObserver.onCompleted();
            
        } catch (Exception e) {
            responseObserver.onError(e);
        }
    }
}
```

**Serverless Serving:**
```python
# AWS Lambda function
import json
import tensorflow as tf
import boto3

model = None

def load_model():
    global model
    if model is None:
        s3 = boto3.client('s3')
        s3.download_file('models-bucket', 'model.h5', '/tmp/model.h5')
        model = tf.keras.models.load_model('/tmp/model.h5')

def lambda_handler(event, context):
    load_model()
    
    # Parse input
    input_data = json.loads(event['body'])
    
    # Make prediction
    prediction = model.predict(input_data['features'])
    
    # Return response
    return {
        'statusCode': 200,
        'body': json.dumps({
            'prediction': prediction.tolist(),
            'confidence': float(max(prediction[0]))
        })
    }
```

---

## 5. Real-time Inference

### Q8: How do you optimize for low-latency inference?

**Answer:**

**Model Optimization Techniques:**
```python
# Model quantization for faster inference
import tensorflow as tf

def optimize_model_for_inference(model_path):
    # Load model
    converter = tf.lite.TFLiteConverter.from_saved_model(model_path)
    
    # Enable optimizations
    converter.optimizations = [tf.lite.Optimize.DEFAULT]
    
    # Quantization
    converter.target_spec.supported_types = [tf.float16]
    
    # Convert model
    tflite_model = converter.convert()
    
    # Save optimized model
    with open('optimized_model.tflite', 'wb') as f:
        f.write(tflite_model)
    
    return 'optimized_model.tflite'

# Inference with optimized model
import numpy as np

def predict_with_tflite(model_path, input_data):
    interpreter = tf.lite.Interpreter(model_path=model_path)
    
    # Allocate tensors
    interpreter.allocate_tensors()
    
    # Get input/output details
    input_details = interpreter.get_input_details()
    output_details = interpreter.get_output_details()
    
    # Set input tensor
    interpreter.set_tensor(input_details[0]['index'], input_data)
    
    # Run inference
    interpreter.invoke()
    
    # Get output
    output_data = interpreter.get_tensor(output_details[0]['index'])
    
    return output_data
```

**Caching Strategy:**
```java
@Service
public class InferenceCacheService {
    
    @Autowired
    private RedisTemplate<String, Object> redisTemplate;
    
    @Cacheable(value = "predictions", key = "#featureHash")
    public Prediction getCachedPrediction(String featureHash) {
        // Cache miss - compute prediction
        return null;
    }
    
    public Prediction predictWithCache(Features features) {
        String featureHash = computeFeatureHash(features);
        
        // Try cache first
        Prediction cached = getCachedPrediction(featureHash);
        if (cached != null) {
            return cached;
        }
        
        // Compute prediction
        Prediction prediction = modelService.predict(features);
        
        // Cache result with TTL
        redisTemplate.opsForValue().set(
            "pred:" + featureHash, 
            prediction, 
            Duration.ofMinutes(30)
        );
        
        return prediction;
    }
}
```

**GPU Inference Service:**
```python
# FastAPI with GPU inference
from fastapi import FastAPI, HTTPException
import torch
import uvicorn

app = FastAPI()

class ModelService:
    def __init__(self):
        self.device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        self.model = self.load_model()
        self.model.to(self.device)
        self.model.eval()
    
    def load_model(self):
        return torch.load("model.pth")
    
    def predict(self, features):
        with torch.no_grad():
            input_tensor = torch.tensor(features).to(self.device)
            output = self.model(input_tensor)
            return output.cpu().numpy()

model_service = ModelService()

@app.post("/predict")
async def predict(request: PredictionRequest):
    try:
        features = request.features
        prediction = model_service.predict(features)
        return {"prediction": prediction.tolist()}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
```

---

## 6. Batch Inference & Offline Processing

### Q9: How do you design batch inference pipelines?

**Answer:**

**Spark Batch Inference:**
```python
from pyspark.sql import SparkSession
from pyspark.sql.functions import col, udf
from pyspark.sql.types import FloatType

spark = SparkSession.builder.appName("BatchInference").getOrCreate()

# Load model (broadcast to all workers)
model = load_model()
broadcast_model = spark.sparkContext.broadcast(model)

def predict_batch(features):
    model = broadcast_model.value
    return float(model.predict(features))

predict_udf = udf(predict_batch, FloatType())

# Batch inference pipeline
def batch_inference(input_path, output_path, date):
    # Load input data
    df = spark.read.parquet(input_path)
    
    # Filter for processing date
    df = df.filter(col("date") == date)
    
    # Apply model inference
    predictions_df = df.withColumn("prediction", predict_udf(col("features")))
    
    # Save results
    predictions_df.write.parquet(output_path, mode="overwrite")
    
    # Generate summary statistics
    summary = predictions_df.agg(
        count("*").alias("total_predictions"),
        avg("prediction").alias("avg_prediction"),
        stddev("prediction").alias("std_prediction")
    ).collect()[0]
    
    return summary
```

**Airflow DAG for Batch Processing:**
```python
from airflow import DAG
from airflow.operators.python_operator import PythonOperator
from airflow.providers.spark.operators.spark_submit import SparkSubmitOperator
from datetime import datetime, timedelta

default_args = {
    'owner': 'ml-team',
    'depends_on_past': False,
    'start_date': datetime(2024, 1, 1),
    'email_on_failure': True,
    'retries': 2,
    'retry_delay': timedelta(minutes=5)
}

dag = DAG(
    'batch_inference_pipeline',
    default_args=default_args,
    schedule_interval='@daily',
    catchup=False
)

# Task 1: Prepare input data
prepare_data = PythonOperator(
    task_id='prepare_input_data',
    python_callable=prepare_batch_input,
    dag=dag
)

# Task 2: Run Spark inference
spark_inference = SparkSubmitOperator(
    task_id='spark_batch_inference',
    application='/opt/spark/applications/batch_inference.py',
    conn_id='spark_default',
    application_args=['--date', '{{ ds }}'],
    dag=dag
)

# Task 3: Post-process results
post_process = PythonOperator(
    task_id='post_process_results',
    python_callable=post_process_predictions,
    dag=dag
)

# Task 4: Update online store
update_online = PythonOperator(
    task_id='update_online_store',
    python_callable=update_feature_store,
    dag=dag
)

# Define pipeline
prepare_data >> spark_inference >> post_process >> update_online
```

---

## 7. ML Monitoring & Observability

### Q10: How do you monitor ML models in production?

**Answer:**

**Drift Detection:**
```python
import numpy as np
from scipy import stats
from sklearn.metrics import jensen_shannon_distance

class DriftDetector:
    
    def __init__(self, reference_data):
        self.reference_stats = self.compute_statistics(reference_data)
    
    def compute_statistics(self, data):
        return {
            'mean': np.mean(data, axis=0),
            'std': np.std(data, axis=0),
            'distribution': np.histogram(data, bins=50, density=True)[0]
        }
    
    def detect_drift(self, current_data, threshold=0.1):
        current_stats = self.compute_statistics(current_data)
        
        # Statistical tests
        drift_scores = {}
        
        # Mean shift detection
        mean_shift = np.abs(self.reference_stats['mean'] - current_stats['mean'])
        drift_scores['mean_shift'] = np.mean(mean_shift)
        
        # Distribution shift (Jensen-Shannon divergence)
        js_divergence = jensen_shannon_distance(
            self.reference_stats['distribution'],
            current_stats['distribution']
        )
        drift_scores['distribution_shift'] = js_divergence
        
        # Overall drift score
        overall_drift = np.mean(list(drift_scores.values()))
        
        return {
            'drift_detected': overall_drift > threshold,
            'drift_score': overall_drift,
            'detailed_scores': drift_scores
        }
```

**Performance Monitoring:**
```java
@Component
public class ModelMonitor {
    
    @Autowired
    private MeterRegistry meterRegistry;
    
    @EventListener
    public void handlePrediction(PredictionEvent event) {
        // Record prediction metrics
        Timer.Sample sample = Timer.start(meterRegistry);
        
        try {
            // Record prediction count
            Counter.builder("model.predictions.count")
                .tag("model", event.getModelName())
                .tag("version", event.getModelVersion())
                .register(meterRegistry)
                .increment();
            
            // Record prediction latency
            sample.stop(Timer.builder("model.predictions.latency")
                .tag("model", event.getModelName())
                .register(meterRegistry));
            
            // Record prediction distribution
            Gauge.builder("model.predictions.value", event, e -> e.getPrediction())
                .tag("model", event.getModelName())
                .register(meterRegistry);
                
        } catch (Exception e) {
            // Record errors
            Counter.builder("model.predictions.errors")
                .tag("model", event.getModelName())
                .tag("error_type", e.getClass().getSimpleName())
                .register(meterRegistry)
                .increment();
        }
    }
    
    @Scheduled(fixedRate = 300000) // Every 5 minutes
    public void checkModelHealth() {
        // Check prediction accuracy if ground truth available
        double recentAccuracy = calculateRecentAccuracy();
        
        Gauge.builder("model.accuracy", () -> recentAccuracy)
            .tag("model", "recommendation")
            .register(meterRegistry);
        
        // Alert if accuracy drops below threshold
        if (recentAccuracy < 0.8) {
            alertService.sendAlert("Model accuracy dropped: " + recentAccuracy);
        }
    }
}
```

**Explainability Monitoring:**
```python
import shap
import numpy as np

class ExplainabilityMonitor:
    
    def __init__(self, model, reference_data):
        self.model = model
        self.explainer = shap.Explainer(model, reference_data)
        self.reference_importance = self.compute_reference_importance()
    
    def compute_reference_importance(self):
        # Compute SHAP values for reference dataset
        shap_values = self.explainer(self.reference_data)
        return np.abs(shap_values.values).mean(axis=0)
    
    def monitor_explanation(self, features, prediction):
        # Compute SHAP values for current prediction
        shap_values = self.explainer(features)
        current_importance = np.abs(shap_values.values[0])
        
        # Compare with reference importance
        importance_drift = np.linalg.norm(current_importance - self.reference_importance)
        
        # Check for explanation drift
        if importance_drift > 0.5:  # Threshold
            return {
                'explanation_drift': True,
                'drift_score': importance_drift,
                'current_importance': current_importance.tolist(),
                'reference_importance': self.reference_importance.tolist()
            }
        
        return {'explanation_drift': False}
```

---

## 8. MLOps & CI/CD for ML

### Q11: How do you implement CI/CD for ML models?

**Answer:**

**ML Pipeline with GitHub Actions:**
```yaml
# .github/workflows/ml-pipeline.yml
name: ML Pipeline

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    
    - name: Set up Python
      uses: actions/setup-python@v2
      with:
        python-version: 3.8
    
    - name: Install dependencies
      run: |
        pip install -r requirements.txt
        pip install pytest pytest-cov
    
    - name: Run tests
      run: |
        pytest tests/ --cov=ml_pipeline --cov-report=xml
    
    - name: Upload coverage
      uses: codecov/codecov-action@v1

  train:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v2
    
    - name: Train model
      run: |
        python train.py --output-dir model_output/
    
    - name: Evaluate model
      run: |
        python evaluate.py --model-path model_output/model.pkl
    
    - name: Register model
      run: |
        python register_model.py --model-path model_output/model.pkl

  deploy:
    needs: train
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - name: Deploy to staging
      run: |
        kubectl set image deployment/ml-serving ml-serving=ml-model:${{ github.sha }}
        kubectl rollout status deployment/ml-serving
    
    - name: Run integration tests
      run: |
        python integration_tests.py --endpoint http://staging-api
    
    - name: Deploy to production
      if: success()
      run: |
        kubectl set image deployment/ml-serving-prod ml-serving=ml-model:${{ github.sha }}
        kubectl rollout status deployment/ml-serving-prod
```

**Model Registry with MLflow:**
```python
import mlflow
import mlflow.sklearn
from sklearn.ensemble import RandomForestClassifier

def train_and_register_model():
    # Load data
    X_train, X_test, y_train, y_test = load_data()
    
    # Train model
    model = RandomForestClassifier(n_estimators=100, random_state=42)
    model.fit(X_train, y_train)
    
    # Evaluate
    accuracy = model.score(X_test, y_test)
    
    # Log to MLflow
    with mlflow.start_run() as run:
        # Log parameters
        mlflow.log_param("n_estimators", 100)
        mlflow.log_param("random_state", 42)
        
        # Log metrics
        mlflow.log_metric("accuracy", accuracy)
        
        # Log model
        mlflow.sklearn.log_model(model, "model")
        
        # Register model
        model_uri = f"runs:/{run.info.run_id}/model"
        mlflow.register_model(model_uri, "production-model")
    
    return run.info.run_id

# Model serving with MLflow
import mlflow.pyfunc

def load_production_model():
    # Load latest production model
    model_uri = "models:/production-model/Production"
    return mlflow.pyfunc.load_model(model_uri)

def predict_with_production_model(features):
    model = load_production_model()
    return model.predict(features)
```

---

## 9. Scaling ML Systems

### Q12: How do you scale ML inference systems?

**Answer:**

**Horizontal Scaling with Kubernetes:**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: ml-serving-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: ml-serving
  minReplicas: 2
  maxReplicas: 20
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  - type: Pods
    pods:
      metric:
        name: requests_per_second
      target:
        type: AverageValue
        averageValue: "100"
```

**Model Sharding for Large Models:**
```python
import torch
import torch.nn as nn

class ShardedModel(nn.Module):
    def __init__(self, model, num_shards):
        super().__init__()
        self.num_shards = num_shards
        self.shards = self.create_shards(model)
    
    def create_shards(self, model):
        # Split model into shards
        shards = []
        layers = list(model.children())
        shard_size = len(layers) // self.num_shards
        
        for i in range(self.num_shards):
            start_idx = i * shard_size
            end_idx = (i + 1) * shard_size if i < self.num_shards - 1 else len(layers)
            shard = nn.Sequential(*layers[start_idx:end_idx])
            shards.append(shard.to(f'cuda:{i}'))
        
        return shards
    
    def forward(self, x):
        # Pipeline execution through shards
        for i, shard in enumerate(self.shards):
            x = x.to(f'cuda:{i}')
            x = shard(x)
        
        return x

# Distributed inference
def distributed_inference(model_shards, input_data):
    futures = []
    
    # Parallel processing on different GPUs
    for i, shard in enumerate(model_shards):
        with torch.cuda.device(i):
            future = torch.distributed.rpc.rpc_async(
                f"worker_{i}",
                shard.forward,
                args=(input_data,)
            )
            futures.append(future)
    
    # Collect results
    results = [future.wait() for future in futures]
    return torch.cat(results, dim=0)
```

**Load Balancing Strategies:**
```java
@RestController
public class LoadBalancedInferenceController {
    
    @Autowired
    private List<ModelService> modelServices;  // Multiple model instances
    
    private final AtomicInteger counter = new AtomicInteger(0);
    
    @PostMapping("/predict")
    public Prediction predict(@RequestBody Request request) {
        // Round-robin load balancing
        int index = counter.getAndIncrement() % modelServices.size();
        ModelService service = modelServices.get(index);
        
        try {
            return service.predict(request);
        } catch (Exception e) {
            // Fallback to next available service
            return handleFallback(request, index);
        }
    }
    
    private Prediction handleFallback(Request request, int failedIndex) {
        for (int i = 0; i < modelServices.size(); i++) {
            if (i != failedIndex) {
                try {
                    return modelServices.get(i).predict(request);
                } catch (Exception e) {
                    // Continue trying other services
                }
            }
        }
        throw new ServiceUnavailableException("All model services failed");
    }
}
```

---

## 10. Common Interview Questions

### Q13: Design a recommendation system for Netflix

**Answer:**
"I'll design a multi-stage recommendation system:

**Architecture:**
```
User Request → Candidate Generation → Ranking → Re-ranking → Response
     │                │               │           │
     ▼                ▼               ▼           ▼
Feature Store    Content-based   Deep Learning  Business Rules
Real-time Data   Collaborative   Neural Nets    Diversity
User Profile     Filtering       Transformer    Freshness
```

**Components:**

1. **Candidate Generation (Millions → Thousands):**
   - Content-based filtering: Similar items based on metadata
   - Collaborative filtering: User-item interactions
   - Real-time features: Recent viewing history

2. **Ranking (Thousands → Hundreds):**
   - Deep neural network with attention mechanisms
   - Features: User profile, item features, context, interactions
   - Output: Click-through rate prediction

3. **Re-ranking (Hundreds → Final):**
   - Business rules: Diversity, freshness, fairness
   - Multi-armed bandit for exploration
   - Personalization filters

**Implementation:**
```python
class RecommendationSystem:
    
    def __init__(self):
        self.candidate_generator = CandidateGenerator()
        self.ranking_model = RankingModel()
        self.reranker = Reranker()
    
    def recommend(self, user_id, context):
        # Generate candidates
        candidates = self.candidate_generator.generate(
            user_id, 
            num_candidates=1000,
            context=context
        )
        
        # Rank candidates
        ranked_candidates = self.ranking_model.rank(
            user_id, 
            candidates,
            context=context
        )
        
        # Re-rank with business rules
        final_recommendations = self.reranker.rerank(
            ranked_candidates,
            diversity_weight=0.3,
            freshness_weight=0.2
        )
        
        return final_recommendations[:50]  # Return top 50
```

**Scaling Considerations:**
- **Pre-computation:** Offline candidate generation for popular content
- **Real-time inference:** Low-latency ranking models
- **A/B testing:** Multiple recommendation strategies
- **Feedback loop:** Continuous learning from user interactions"

### Q14: How would you handle model versioning and A/B testing?

**Answer:**
"I'll implement a comprehensive model management system:

**Model Versioning Strategy:**
```python
class ModelRegistry:
    
    def __init__(self):
        self.models = {}  # version -> model mapping
        self.metadata = {}  # version -> metadata
    
    def register_model(self, model, version, metadata):
        self.models[version] = model
        self.metadata[version] = {
            'created_at': datetime.now(),
            'performance': metadata['performance'],
            'features': metadata['features'],
            'training_data': metadata['training_data']
        }
    
    def get_model(self, version):
        return self.models.get(version)
    
    def promote_model(self, from_version, to_stage):
        # Promote model through stages: dev → staging → production
        if to_stage == 'production':
            # Run validation checks
            if self.validate_model(from_version):
                return self.deploy_to_production(from_version)
        return False
```

**A/B Testing Framework:**
```java
@Service
public class ABTestingService {
    
    @Autowired
    private ModelRegistry modelRegistry;
    
    @Autowired
    private TrafficSplitter trafficSplitter;
    
    public Prediction predictWithABTest(Request request) {
        // Determine which model to use
        String modelVersion = trafficSplitter.getModelVersion(
            request.getUserId(), 
            request.getExperimentId()
        );
        
        // Get prediction from selected model
        Model model = modelRegistry.getModel(modelVersion);
        Prediction prediction = model.predict(request);
        
        // Log A/B test data
        abTestLogger.logExperiment(
            request.getUserId(),
            modelVersion,
            prediction,
            request.getExperimentId()
        );
        
        return prediction;
    }
    
    @Scheduled(fixedRate = 3600000)  # Hourly
    public void analyzeExperimentResults() {
        ExperimentResults results = abTestAnalyzer.analyzeResults();
        
        // Check statistical significance
        if (results.isStatisticallySignificant()) {
            // Promote winning model
            String winner = results.getWinningModel();
            modelRegistry.promoteModel(winner, "production");
            
            // Stop experiment
            trafficSplitter.stopExperiment(results.getExperimentId());
        }
    }
}
```

**Traffic Splitting:**
```python
class TrafficSplitter:
    
    def __init__(self):
        self.experiments = {}  # experiment_id → config
    
    def get_model_version(self, user_id, experiment_id):
        if experiment_id not in self.experiments:
            return "control"  # Default model
        
        config = self.experiments[experiment_id]
        
        # Consistent hashing for user assignment
        user_hash = hash(user_id) % 100
        
        if user_hash < config['control_percentage']:
            return "control"
        else:
            return config['treatment_model']
    
    def update_experiment(self, experiment_id, config):
        self.experiments[experiment_id] = config
```

**Monitoring and Rollback:**
```java
@Component
public class ExperimentMonitor {
    
    @EventListener
    public void handlePrediction(PredictionEvent event) {
        // Track performance metrics per model version
        String modelVersion = event.getModelVersion();
        
        // Update metrics
        metricsCollector.recordPrediction(modelVersion, event);
        
        // Check for degradation
        if (detectPerformanceDegradation(modelVersion)) {
            alertService.sendAlert(
                "Performance degradation detected for model: " + modelVersion
            );
            
            // Auto-rollback if severe
            if (isSevereDegradation(modelVersion)) {
                rollbackToPreviousVersion(modelVersion);
            }
        }
    }
    
    private void rollbackToPreviousVersion(String currentVersion) {
        String previousVersion = getPreviousStableVersion();
        trafficSplitter.updateExperiment(
            currentVersion,
            {"control_percentage": 100, "treatment_model": previousVersion}
        );
    }
}
```"

### Q15: How do you handle training-serving skew?

**Answer:**
"Training-serving skew occurs when the data distribution during training differs from serving. Here's how to address it:

**Detection:**
```python
class SkewDetector:
    
    def __init__(self, training_stats):
        self.training_stats = training_stats
    
    def detect_skew(self, serving_data):
        # Compare feature distributions
        skew_scores = {}
        
        for feature_name in self.training_stats:
            training_dist = self.training_stats[feature_name]
            serving_dist = self.compute_distribution(serving_data[feature_name])
            
            # Statistical tests
            ks_statistic, p_value = ks_2samp(training_dist, serving_dist)
            
            skew_scores[feature_name] = {
                'ks_statistic': ks_statistic,
                'p_value': p_value,
                'skew_detected': p_value < 0.05
            }
        
        return skew_scores
```

**Prevention Strategies:**

1. **Feature Store Integration:**
```java
@Service
public class ConsistentFeatureService {
    
    public Features getTrainingFeatures(String userId, LocalDateTime timestamp) {
        // Use same feature logic for training and serving
        return featureCalculator.calculateFeatures(userId, timestamp);
    }
    
    public Features getServingFeatures(String userId) {
        // Use same feature logic with current timestamp
        return featureCalculator.calculateFeatures(userId, LocalDateTime.now());
    }
}
```

2. **Online Training:**
```python
class OnlineLearning:
    
    def __init__(self, model):
        self.model = model
        self.feature_store = FeatureStore()
    
    def update_with_serving_data(self, user_id, features, true_label):
        # Use serving features for online updates
        self.model.partial_fit([features], [true_label])
        
        # Update feature store
        self.feature_store.update_features(user_id, features, true_label)
```

3. **Data Pipeline Consistency:**
```yaml
# Unified data processing
feature_pipeline:
  training:
    source: historical_data
    transformations:
      - normalize_features
      - encode_categorical
      - handle_missing_values
    output: training_features
  
  serving:
    source: real_time_events
    transformations:
      - normalize_features  # Same as training
      - encode_categorical  # Same as training
      - handle_missing_values  # Same as training
    output: serving_features
```

**Monitoring and Alerting:**
```java
@Component
public class SkewMonitor {
    
    @Scheduled(fixedRate = 300000)  # Every 5 minutes
    public void checkForSkew() {
        // Collect recent serving statistics
        ServingStats servingStats = collectServingStats();
        
        // Compare with training statistics
        SkewReport report = skewDetector.compare(trainingStats, servingStats);
        
        if (report.hasSignificantSkew()) {
            // Alert team
            alertService.sendSkewAlert(report);
            
            // Trigger retraining if needed
            if (report.shouldRetrain()) {
                trainingService.triggerRetraining();
            }
        }
    }
}
```"

---

## 🎯 Quick Reference

### Key ML Infrastructure Components
- **Feature Store:** Centralized feature management
- **Model Registry:** Version control for models
- **Serving Layer:** Real-time and batch inference
- **Monitoring:** Drift detection, performance tracking
- **CI/CD:** Automated training and deployment

### Common Technologies
- **Training:** TensorFlow, PyTorch, Spark ML
- **Serving:** TensorFlow Serving, TorchServe, Seldon
- **Feature Store:** Feast, Tecton
- **Monitoring:** MLflow, Weights & Biases
- **Orchestration:** Kubeflow, Airflow, Prefect

### Interview Focus Areas
- **System design:** End-to-end ML pipelines
- **Scalability:** Distributed training and serving
- **Reliability:** Monitoring, drift detection, A/B testing
- **Performance:** Optimization techniques, caching
- **MLOps:** CI/CD, versioning, automation
