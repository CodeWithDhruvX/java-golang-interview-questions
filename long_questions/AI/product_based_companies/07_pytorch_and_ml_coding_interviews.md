# PyTorch & ML Coding Interviews (Product-Based Companies)

For core Machine Learning Engineer roles at top product companies, passing LeetCode DSA is not enough. You will face ML-specific live coding rounds where you must implement standard layers, loss functions, or training loops from scratch using PyTorch or pure NumPy.

## 1. Implement a complete PyTorch training loop for a simple classification model.
**Answer:**
A robust PyTorch training loop must explicitly manage devices (CPU/GPU), gradients, forward passes, and backpropagation. Interviewers look for proper zeroing of gradients and evaluation toggling (`model.train()` vs `model.eval()`).

```python
import torch
import torch.nn as nn
import torch.optim as optim

def train_model(model, train_loader, val_loader, epochs=5, lr=0.001):
    # 1. Setup Device
    device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
    model.to(device)
    
    # 2. Define Loss and Optimizer
    criterion = nn.CrossEntropyLoss()
    optimizer = optim.Adam(model.parameters(), lr=lr)
    
    for epoch in range(epochs):
        # --------------------
        # TRAINING PHASE
        # --------------------
        model.train() # Set to training mode (enables Dropout/BatchNorm)
        running_loss = 0.0
        
        for inputs, labels in train_loader:
            inputs, labels = inputs.to(device), labels.to(device)
            
            # Step A: Zero the parameter gradients
            # (Crucial: PyTorch accumulates gradients by default)
            optimizer.zero_grad() 
            
            # Step B: Forward pass
            outputs = model(inputs)
            loss = criterion(outputs, labels)
            
            # Step C: Backward pass (Compute gradients)
            loss.backward()
            
            # Step D: Optimize (Update weights)
            optimizer.step()
            
            running_loss += loss.item() * inputs.size(0)
            
        epoch_loss = running_loss / len(train_loader.dataset)
        
        # --------------------
        # VALIDATION PHASE
        # --------------------
        model.eval() # Set to evaluation mode (disables Dropout/BatchNorm)
        val_loss = 0.0
        correct = 0
        
        # Disable gradient calculation for inference to save memory/compute
        with torch.no_grad():
            for val_inputs, val_labels in val_loader:
                val_inputs, val_labels = val_inputs.to(device), val_labels.to(device)
                
                val_outputs = model(val_inputs)
                loss = criterion(val_outputs, val_labels)
                val_loss += loss.item() * val_inputs.size(0)
                
                # Calculate accuracy
                _, preds = torch.max(val_outputs, 1)
                correct += torch.sum(preds == val_labels.data)
                
        epoch_val_loss = val_loss / len(val_loader.dataset)
        epoch_val_acc = correct.double() / len(val_loader.dataset)
        
        print(f'Epoch {epoch+1}/{epochs} | Train Loss: {epoch_loss:.4f} | Val Loss: {epoch_val_loss:.4f} | Val Acc: {epoch_val_acc:.4f}')

    return model
```

## 2. Implement the Scaled Dot-Product Attention mechanism from scratch using PyTorch.
**Answer:**
This is a very common interview question to test your understanding of Transformer architecture. The interviewer expects matrix multiplications and correct dimensional broadcasting.

**Math recap:** $\text{Attention}(Q, K, V) = \text{softmax}(\frac{QK^T}{\sqrt{d_k}})V$

```python
import torch
import torch.nn as nn
import torch.nn.functional as F
import math

class ScaledDotProductAttention(nn.Module):
    def __init__(self):
        super(ScaledDotProductAttention, self).__init__()

    def forward(self, query, key, value, mask=None):
        # Expected shape of Q, K, V: 
        # (batch_size, num_heads, sequence_length, depth/d_k)
        
        d_k = query.size(-1) # The depth of the key vector
        
        # 1. QKT: Transpose the last two dimensions of Key for dot product
        # Shape becomes: (batch_size, num_heads, seq_len, seq_len)
        scores = torch.matmul(query, key.transpose(-2, -1))
        
        # 2. Scale by sqrt(d_k)
        scores = scores / math.sqrt(d_k)
        
        # 3. Apply optional causality mask (e.g., for Decoder)
        if mask is not None:
            # Replace 0s in mask with -infinity so softmax makes them 0
            scores = scores.masked_fill(mask == 0, -1e9)
            
        # 4. Softmax over the last dimension (key sequence length)
        # to get probability attention weights
        p_attn = F.softmax(scores, dim=-1)
        
        # 5. Multiply by Value
        # Result shape: (batch_size, num_heads, seq_len, depth)
        output = torch.matmul(p_attn, value)
        
        return output, p_attn

# Dummy data test
batch_size, num_heads, seq_len, d_k = 2, 8, 10, 64
q = torch.rand(batch_size, num_heads, seq_len, d_k)
k = torch.rand(batch_size, num_heads, seq_len, d_k)
v = torch.rand(batch_size, num_heads, seq_len, d_k)

attention_layer = ScaledDotProductAttention()
output, weights = attention_layer(q, k, v)
print(f"Output shape: {output.shape}") # Should be: [2, 8, 10, 64]
```

## 3. Write a pure NumPy implementation of the Sigmoid function and its derivative.
**Answer:**
Interviewers ask for pure NumPy to ensure you aren't just relying on `torch.nn.functional` and actually understand the underlying math acting on arrays.

```python
import numpy as np

def sigmoid(x):
    """
    Compute the sigmoid of x.
    x can be a scalar or a numpy array.
    """
    # Note: np.exp(-x) can overflow if x is a large negative number.
    # A robust implementation handles this, but for interviews, this is usually acceptable.
    return 1.0 / (1.0 + np.exp(-x))

def sigmoid_derivative(x):
    """
    Compute the derivative of the sigmoid function with respect to its input x.
    Derivative of sigmoid(x) is sigmoid(x) * (1 - sigmoid(x))
    """
    s = sigmoid(x)
    return s * (1.0 - s)

# Example usage over a vector
x = np.array([-2.0, 0.0, 2.0])
print(f"Sigmoid: {sigmoid(x)}")
print(f"Derivative: {sigmoid_derivative(x)}")
```

## 4. How would you implement early stopping in PyTorch?
**Answer:**
Early stopping prevents overfitting by halting training when the validation loss stops improving, rather than running for a fixed number of epochs.

```python
class EarlyStopping:
    def __init__(self, patience=5, min_delta=0.0):
        """
        Args:
            patience (int): How many epochs to wait after last time validation loss improved.
            min_delta (float): Minimum change in the monitored quantity to qualify as an improvement.
        """
        self.patience = patience
        self.min_delta = min_delta
        self.counter = 0
        self.best_loss = None
        self.early_stop = False

    def __call__(self, val_loss):
        # First epoch
        if self.best_loss is None:
            self.best_loss = val_loss
        # If validation loss doesn't improve by at least min_delta
        elif val_loss > self.best_loss - self.min_delta:
            self.counter += 1
            print(f'EarlyStopping counter: {self.counter} out of {self.patience}')
            if self.counter >= self.patience:
                self.early_stop = True
        # If validation loss improves
        else:
            self.best_loss = val_loss
            self.counter = 0 # Reset counter

# --- Inside your training loop ---
# early_stopping = EarlyStopping(patience=3)
# for epoch in range(epochs):
#     ... train ...
#     val_loss = ... calculate validation loss ...
#
#     early_stopping(val_loss)
#     if early_stopping.early_stop:
#         print("Early stopping triggered. Halting training.")
#         break
```

## 5. Implement IoU (Intersection over Union) using PyTorch tensors for an Object Detection output.
**Answer:**
IoU evaluates how well the predicted bounding box matches the ground truth bounding box.
Assumptions: Boxes are given as `[x1, y1, x2, y2]` where `(x1, y1)` is the top-left coordinate and `(x2, y2)` is bottom-right.

```python
import torch

def intersection_over_union(boxes_preds, boxes_labels):
    """
    Calculates Intersection over Union
    boxes_preds: tensor of shape (N, 4)
    boxes_labels: tensor of shape (N, 4)
    """
    # 1. Find the coordinates of the intersecting rectangle
    # We want the maximum of the top-left coordinates and minimum of bottom-right
    box1_x1 = boxes_preds[..., 0:1]
    box1_y1 = boxes_preds[..., 1:2]
    box1_x2 = boxes_preds[..., 2:3]
    box1_y2 = boxes_preds[..., 3:4]
    
    box2_x1 = boxes_labels[..., 0:1]
    box2_y1 = boxes_labels[..., 1:2]
    box2_x2 = boxes_labels[..., 2:3]
    box2_y2 = boxes_labels[..., 3:4]
    
    x1 = torch.max(box1_x1, box2_x1)
    y1 = torch.max(box1_y1, box2_y1)
    x2 = torch.min(box1_x2, box2_x2)
    y2 = torch.min(box1_y2, box2_y2)
    
    # 2. Calculate the area of intersection
    # Use clamp(0) to ensure that if boxes don't intersect, the width/height is 0, not negative
    intersection_width = (x2 - x1).clamp(0)
    intersection_height = (y2 - y1).clamp(0)
    intersection_area = intersection_width * intersection_height
    
    # 3. Calculate areas of both individual boxes
    box1_area = abs((box1_x2 - box1_x1) * (box1_y2 - box1_y1))
    box2_area = abs((box2_x2 - box2_x1) * (box2_y2 - box2_y1))
    
    # 4. Calculate Union
    # Union Area = Area(A) + Area(B) - Area(Intersection)
    # Add a tiny epsilon (1e-6) to the denominator to prevent division by zero
    union_area = box1_area + box2_area - intersection_area + 1e-6
    
    return intersection_area / union_area

# Example usage
# [x1, y1, x2, y2]
pred_box = torch.tensor([[2.0, 2.0, 6.0, 6.0]])
true_box = torch.tensor([[4.0, 4.0, 8.0, 8.0]])

iou = intersection_over_union(pred_box, true_box)
print(f"IoU: {iou.item():.4f}") # Should be 0.1428 (Area 4 / Area 28)
```
