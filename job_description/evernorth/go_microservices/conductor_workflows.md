# Netflix Conductor Workflow Orchestration with Go

## Overview

Netflix Conductor is a workflow orchestration engine that enables building complex, resilient applications. This guide covers Conductor integration patterns specifically for Go microservices.

## Core Concepts

### 1. Conductor Architecture

- **Workflow Definition**: JSON-based workflow definitions
- **Task Definitions**: Individual task configurations
- **Workers**: Go services that execute specific tasks
- **Server**: Conductor server that manages workflow execution

### 2. Go Client Integration

The Conductor Go client provides:
- Workflow execution and management
- Task polling and execution
- Event handling and callbacks
- Monitoring and metrics integration

## Implementation Patterns

### 1. Worker Implementation

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"
    
    "github.com/conductor-sdk/conductor-go/sdk/model"
    "github.com/conductor-sdk/conductor-go/sdk/workers"
)

type PaymentWorker struct {
    taskType string
    conductor *workers.TaskRunner
}

func NewPaymentWorker(taskType string) *PaymentWorker {
    return &PaymentWorker{
        taskType: taskType,
    }
}

func (pw *PaymentWorker) ExecuteTask(ctx context.Context, task *model.Task) (*model.TaskResult, error) {
    log.Printf("Executing payment task: %s", task.TaskId)
    
    // Extract task input
    var paymentData struct {
        OrderID    string  `json:"orderId"`
        Amount     float64 `json:"amount"`
        Currency   string  `json:"currency"`
        PaymentMethod string `json:"paymentMethod"`
    }
    
    if err := json.Unmarshal(task.InputData, &paymentData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal payment data: %w", err)
    }
    
    // Process payment
    result, err := pw.processPayment(ctx, paymentData)
    if err != nil {
        log.Printf("Payment processing failed: %v", err)
        return &model.TaskResult{
            Status:     model.Failed,
            ReasonForIncompletion: err.Error(),
            OutputData: nil,
        }, nil
    }
    
    // Return success result
    outputData, _ := json.Marshal(result)
    return &model.TaskResult{
        Status:     model.Completed,
        OutputData: outputData,
    }, nil
}

func (pw *PaymentWorker) processPayment(ctx context.Context, paymentData struct {
    OrderID    string  `json:"orderId"`
    Amount     float64 `json:"amount"`
    Currency   string  `json:"currency"`
    PaymentMethod string `json:"paymentMethod"`
}) (map[string]interface{}, error) {
    // Simulate payment processing
    log.Printf("Processing payment for order %s: %.2f %s via %s", 
        paymentData.OrderID, paymentData.Amount, paymentData.Currency, paymentData.PaymentMethod)
    
    // Add business logic here
    // - Validate payment method
    // - Check fraud detection
    // - Process payment gateway
    // - Update order status
    
    time.Sleep(2 * time.Second) // Simulate processing time
    
    return map[string]interface{}{
        "orderId":        paymentData.OrderID,
        "transactionId":  fmt.Sprintf("txn_%d", time.Now().Unix()),
        "status":         "completed",
        "processedAt":    time.Now().Format(time.RFC3339),
        "amount":         paymentData.Amount,
        "currency":       paymentData.Currency,
    }, nil
}

func (pw *PaymentWorker) PollAndExecute(ctx context.Context) error {
    return pw.conductor.StartWorkers(ctx, pw.taskType, pw.ExecuteTask)
}
```

### 2. Workflow Management

```go
type WorkflowManager struct {
    client *conductor.Client
}

func NewWorkflowManager(apiURL string) (*WorkflowManager, error) {
    client, err := conductor.NewClient(&conductor.Config{
        RootURI: apiURL,
        // Add authentication if needed
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create conductor client: %w", err)
    }
    
    return &WorkflowManager{client: client}, nil
}

func (wm *WorkflowManager) StartWorkflow(workflowName string, input map[string]interface{}) (string, error) {
    workflowRequest := &model.StartWorkflowRequest{
        Name:  workflowName,
        Input: input,
        // Add correlation ID for tracking
        CorrelationId: generateCorrelationID(),
    }
    
    workflowID, err := wm.client.WorkflowClient.StartWorkflow(workflowRequest)
    if err != nil {
        return "", fmt.Errorf("failed to start workflow: %w", err)
    }
    
    return workflowID, nil
}

func (wm *WorkflowManager) GetWorkflowStatus(workflowID string) (*model.Workflow, error) {
    workflow, err := wm.client.WorkflowClient.GetWorkflow(workflowID, false)
    if err != nil {
        return nil, fmt.Errorf("failed to get workflow status: %w", err)
    }
    
    return workflow, nil
}

func (wm *WorkflowManager) WaitForCompletion(ctx context.Context, workflowID string, timeout time.Duration) (*model.Workflow, error) {
    deadline := time.After(timeout)
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return nil, ctx.Err()
        case <-deadline:
            return nil, fmt.Errorf("workflow timeout after %v", timeout)
        case <-ticker.C:
            workflow, err := wm.GetWorkflowStatus(workflowID)
            if err != nil {
                return nil, err
            }
            
            if workflow.Status == model.CompletedWorkflow {
                return workflow, nil
            } else if workflow.Status == model.FailedWorkflow {
                return workflow, fmt.Errorf("workflow failed")
            } else if workflow.Status == model.TerminatedWorkflow {
                return workflow, fmt.Errorf("workflow terminated")
            }
        }
    }
}
```

### 3. Self-Healing Patterns

```go
type SelfHealingWorker struct {
    taskType    string
    maxRetries  int
    retryDelay  time.Duration
    circuitBreaker *CircuitBreaker
}

type CircuitBreaker struct {
    maxFailures int
    resetTimeout time.Duration
    failures    int
    lastFailure time.Time
    state       CircuitBreakerState
    mutex       sync.RWMutex
}

type CircuitBreakerState int

const (
    Closed CircuitBreakerState = iota
    Open
    HalfOpen
)

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mutex.Lock()
    defer cb.mutex.Unlock()
    
    // Check if circuit breaker should reset
    if cb.state == Open && time.Since(cb.lastFailure) > cb.resetTimeout {
        cb.state = HalfOpen
        cb.failures = 0
    }
    
    // Reject calls if circuit is open
    if cb.state == Open {
        return fmt.Errorf("circuit breaker is open")
    }
    
    // Execute the function
    err := fn()
    
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
        
        if cb.failures >= cb.maxFailures {
            cb.state = Open
        }
        
        return err
    }
    
    // Reset on success
    if cb.state == HalfOpen {
        cb.state = Closed
    }
    cb.failures = 0
    
    return nil
}

func (shw *SelfHealingWorker) ExecuteTaskWithRetry(ctx context.Context, task *model.Task) (*model.TaskResult, error) {
    var lastErr error
    
    for attempt := 0; attempt <= shw.maxRetries; attempt++ {
        if attempt > 0 {
            select {
            case <-ctx.Done():
                return nil, ctx.Err()
            case <-time.After(shw.retryDelay):
                // Wait before retry
            }
        }
        
        // Check circuit breaker
        err := shw.circuitBreaker.Call(func() error {
            result, err := shw.ExecuteTask(ctx, task)
            if err != nil {
                return err
            }
            lastErr = nil
            return nil
        })
        
        if err == nil {
            return result, nil
        }
        
        lastErr = err
        log.Printf("Task execution failed (attempt %d/%d): %v", attempt+1, shw.maxRetries, err)
    }
    
    return &model.TaskResult{
        Status:     model.Failed,
        ReasonForIncompletion: fmt.Sprintf("Max retries exceeded: %v", lastErr),
        OutputData: nil,
    }, lastErr
}
```

### 4. Task Recovery and Compensation

```go
type CompensationWorker struct {
    taskType string
    compensations map[string]CompensationFunc
}

type CompensationFunc func(ctx context.Context, task *model.Task) error

func (cw *CompensationWorker) RegisterCompensation(taskName string, fn CompensationFunc) {
    if cw.compensations == nil {
        cw.compensations = make(map[string]CompensationFunc)
    }
    cw.compensations[taskName] = fn
}

func (cw *CompensationWorker) ExecuteCompensation(ctx context.Context, task *model.Task) error {
    compensationFn, exists := cw.compensations[task.TaskType]
    if !exists {
        return fmt.Errorf("no compensation function found for task type: %s", task.TaskType)
    }
    
    return compensationFn(ctx, task)
}

// Example compensation functions
func compensatePayment(ctx context.Context, task *model.Task) error {
    var paymentData struct {
        TransactionID string `json:"transactionId"`
        Amount        float64 `json:"amount"`
    }
    
    if err := json.Unmarshal(task.InputData, &paymentData); err != nil {
        return err
    }
    
    log.Printf("Compensating payment transaction: %s", paymentData.TransactionID)
    
    // Implement refund logic
    // - Call payment gateway refund API
    // - Update order status
    // - Send notification
    
    return nil
}

func compensateInventory(ctx context.Context, task *model.Task) error {
    var inventoryData struct {
        ProductID string `json:"productId"`
        Quantity   int    `json:"quantity"`
    }
    
    if err := json.Unmarshal(task.InputData, &inventoryData); err != nil {
        return err
    }
    
    log.Printf("Restoring inventory for product: %s, quantity: %d", inventoryData.ProductID, inventoryData.Quantity)
    
    // Implement inventory restoration
    // - Update inventory count
    // - Release reserved items
    // - Log compensation action
    
    return nil
}
```

### 5. Event-Driven Workflow Integration

```go
type EventHandler struct {
    conductorClient *conductor.Client
    eventQueue      chan Event
}

type Event struct {
    Type      string                 `json:"type"`
    Data      map[string]interface{} `json:"data"`
    Timestamp time.Time              `json:"timestamp"`
    Source    string                 `json:"source"`
}

func (eh *EventHandler) StartEventProcessor(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        case event := <-eh.eventQueue:
            go eh.processEvent(ctx, event)
        }
    }
}

func (eh *EventHandler) processEvent(ctx context.Context, event Event) {
    switch event.Type {
    case "order.created":
        eh.startOrderWorkflow(ctx, event.Data)
    case "payment.completed":
        eh.updateOrderStatus(ctx, event.Data)
    case "inventory.updated":
        eh.notifyShipping(ctx, event.Data)
    default:
        log.Printf("Unknown event type: %s", event.Type)
    }
}

func (eh *EventHandler) startOrderWorkflow(ctx context.Context, eventData map[string]interface{}) {
    workflowName := "order_processing_workflow"
    
    workflowID, err := eh.conductorClient.WorkflowClient.StartWorkflow(&model.StartWorkflowRequest{
        Name:  workflowName,
        Input: eventData,
        CorrelationId: generateCorrelationID(),
    })
    
    if err != nil {
        log.Printf("Failed to start order workflow: %v", err)
        return
    }
    
    log.Printf("Started order workflow: %s", workflowID)
}
```

## Workflow Definitions

### 1. Order Processing Workflow

```json
{
  "name": "order_processing_workflow",
  "description": "Process customer orders with payment and inventory",
  "version": 1,
  "tasks": [
    {
      "name": "validate_order",
      "taskReferenceName": "validate_order_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "order": "${workflow.input.order}"
      }
    },
    {
      "name": "reserve_inventory",
      "taskReferenceName": "reserve_inventory_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "orderId": "${validate_order_ref.output.orderId}",
        "items": "${validate_order_ref.output.items}"
      }
    },
    {
      "name": "process_payment",
      "taskReferenceName": "process_payment_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "orderId": "${reserve_inventory_ref.output.orderId}",
        "amount": "${reserve_inventory_ref.output.total}",
        "paymentMethod": "${workflow.input.paymentMethod}"
      }
    },
    {
      "name": "update_order_status",
      "taskReferenceName": "update_order_status_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "orderId": "${process_payment_ref.output.orderId}",
        "status": "paid"
      }
    },
    {
      "name": "notify_shipping",
      "taskReferenceName": "notify_shipping_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "orderId": "${update_order_status_ref.output.orderId}"
      }
    }
  ],
  "failureWorkflow": "order_compensation_workflow",
  "restartable": true,
  "workflowTimeoutPolicy": "TIME_OUT",
  "timeoutSeconds": 3600
}
```

### 2. Compensation Workflow

```json
{
  "name": "order_compensation_workflow",
  "description": "Compensate failed order processing",
  "version": 1,
  "tasks": [
    {
      "name": "refund_payment",
      "taskReferenceName": "refund_payment_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "transactionId": "${workflow.input.transactionId}",
        "amount": "${workflow.input.amount}"
      }
    },
    {
      "name": "restore_inventory",
      "taskReferenceName": "restore_inventory_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "items": "${workflow.input.items}"
      }
    },
    {
      "name": "update_order_status",
      "taskReferenceName": "update_order_status_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "orderId": "${workflow.input.orderId}",
        "status": "failed"
      }
    }
  ]
}
```

## Monitoring and Metrics

### 1. Conductor Metrics Integration

```go
type ConductorMetrics struct {
    workflowsStarted     prometheus.Counter
    workflowsCompleted   prometheus.Counter
    workflowsFailed      prometheus.Counter
    tasksExecuted        prometheus.Counter
    taskDuration        *prometheus.HistogramVec
    circuitBreakerTrips prometheus.Counter
}

func NewConductorMetrics() *ConductorMetrics {
    return &ConductorMetrics{
        workflowsStarted: prometheus.NewCounter(prometheus.CounterOpts{
            Name: "conductor_workflows_started_total",
            Help: "Total number of workflows started",
        }),
        workflowsCompleted: prometheus.NewCounter(prometheus.CounterOpts{
            Name: "conductor_workflows_completed_total",
            Help: "Total number of workflows completed",
        }),
        workflowsFailed: prometheus.NewCounter(prometheus.CounterOpts{
            Name: "conductor_workflows_failed_total",
            Help: "Total number of workflows failed",
        }),
        tasksExecuted: prometheus.NewCounter(prometheus.CounterOpts{
            Name: "conductor_tasks_executed_total",
            Help: "Total number of tasks executed",
        }),
        taskDuration: prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Name: "conductor_task_duration_seconds",
                Help: "Task execution duration in seconds",
                Buckets: prometheus.DefBuckets,
            },
            []string{"task_type", "status"},
        ),
        circuitBreakerTrips: prometheus.NewCounter(prometheus.CounterOpts{
            Name: "conductor_circuit_breaker_trips_total",
            Help: "Total number of circuit breaker trips",
        }),
    }
}

func (cm *ConductorMetrics) RecordWorkflowStart(workflowName string) {
    cm.workflowsStarted.Inc()
}

func (cm *ConductorMetrics) RecordWorkflowCompletion(workflowName string, status string, duration time.Duration) {
    if status == "COMPLETED" {
        cm.workflowsCompleted.Inc()
    } else {
        cm.workflowsFailed.Inc()
    }
}

func (cm *ConductorMetrics) RecordTaskExecution(taskType string, status string, duration time.Duration) {
    cm.tasksExecuted.Inc()
    cm.taskDuration.WithLabelValues(taskType, status).Observe(duration.Seconds())
}
```

## Interview Questions

### Technical Questions
1. How do you implement a Conductor worker in Go?
2. Explain the difference between synchronous and asynchronous task execution.
3. How do you handle workflow failures and compensation?
4. What are self-healing patterns in workflow orchestration?
5. How do you implement circuit breakers in Conductor workers?

### Design Questions
1. Design an order processing workflow using Conductor and Go.
2. How would you implement distributed transactions with Conductor?
3. Explain how to handle long-running tasks in Conductor.
4. Design a monitoring strategy for Conductor workflows.
5. How do you ensure workflow idempotency and safety?

### Practical Coding
1. Write a Go worker that processes payment tasks with retry logic.
2. Implement a compensation workflow for failed operations.
3. Create a metrics collection system for Conductor workflows.
4. Design an event-driven workflow integration pattern.
5. Implement a circuit breaker pattern for task execution.

## Best Practices

1. **Idempotent Tasks**: Design tasks to be safely retriable
2. **Timeout Management**: Set appropriate timeouts for tasks and workflows
3. **Error Handling**: Implement comprehensive error handling and logging
4. **Resource Management**: Properly clean up resources and connections
5. **Monitoring**: Track workflow performance and failure rates
6. **Testing**: Unit test workers and integration test workflows
7. **Security**: Secure task data and workflow communications
