# Deep Learning, NLP & Computer Vision (Product-Based Companies)

This section covers Deep Neural Networks (DNNs), Convolutional Neural Networks (CNNs), Recurrent Neural Networks (RNNs), and their application in specialized Natural Language Processing (NLP) and Computer Vision (CV) tasks.

## 1. Explain the Vanishing and Exploding Gradient problems in deep neural networks. How do we mitigate them?
**Answer:**
**The Problem:**
During backpropagation, errors are passed backward through the network layers using the chain rule of calculus. If we have a deep network with many layers, we multiply many gradients together.
*   **Vanishing Gradients:** If the gradients of the activation functions are small (e.g., less than 1, like in the saturated regions of Sigmoid or Tanh), multiplying many small numbers together causes the overall gradient to shrink exponentially. Early layers receive almost zero gradient and stop learning.
*   **Exploding Gradients:** Conversely, if the weights are initialized too large and the activation derivatives are $>1$, the gradients grow exponentially, leading to massive weight updates, making the model unstable (often resulting in `NaN` loss).

**Mitigation Strategies:**
1.  **Activation Functions:**
    *   Switch from Sigmoid/Tanh to **ReLU** (Rectified Linear Unit, $f(x) = \max(0, x)$). The gradient of ReLU is either 0 or exactly 1, preventing the gradient from shrinking during multiplication in the positive domain.
    *   Variants like Leaky ReLU or ELU address the "dying ReLU" problem (where units get stuck returning 0).
2.  **Weight Initialization:**
    *   **Xavier/Glorot Initialization:** Designed for Tanh/Sigmoid. It initializes weights such that the variance of the outputs of a layer is equal to the variance of its inputs.
    *   **He Initialization:** Designed for ReLU. Similar to Xavier but accounts for the fact that ReLU zeroes out half the distribution.
3.  **Batch Normalization:**
    *   Normalizes the activations of each layer to have a mean of 0 and variance of 1 across the mini-batch. This keeps activations in the "sweet spot" of non-linearities, reducing the likelihood of vanishing/exploding gradients.
4.  **Architectural Changes:**
    *   **Residual Connections (ResNets):** Allow gradients to bypass layers via shortcut connections ($F(x) + x$). This creates a direct, uninterrupted gradient path from the output back to the earlier layers.
5.  **Gradient Clipping:**
    *   Specifically for exploding gradients (common in RNNs/LSTMs). If the gradient norm exceeds a threshold, simply cap it (clip it) to that threshold.

## 2. How do LSTMs solve the vanishing gradient problem compared to vanilla RNNs? Explain cell state vs. hidden state.
**Answer:**
**Vanilla RNN Issue:**
Standard RNNs suffer heavily from vanishing gradients when processing long sequences because the hidden state is repeatedly multiplied by the same weight matrix $W_h$ at every time step.

**The LSTM Solution:**
Long Short-Term Memory (LSTM) networks solve this by introducing an internal memory mechanism called the **Cell State** and a set of **Gates** to control information flow.

*   **Cell State ($C_t$):** The "conveyor belt". It runs continuously down the entire sequence with only minor linear interactions. This allows information to flow practically unchanged over many time steps, preserving the gradient.
*   **Hidden State ($h_t$):** The short-term memory or the "output" of the LSTM cell at the current time step for the next layer (or the final prediction).

**The Gates (Learnable sigmoid layers that output between 0 and 1):**
1.  **Forget Gate ($f_t$):** Decides what information to throw away from the *previous* cell state ($C_{t-1}$). `0 = forget completely, 1 = keep completely`.
2.  **Input Gate ($i_t$):** Decides what *new* information (from the candidate cell state $\tilde{C}_t$) we are going to store in the current cell state.
3.  **Output Gate ($o_t$):** Decides what parts of the updated cell state ($C_t$) we are going to output as the new hidden state ($h_t$).

**Mathematical intuition for gradient flow:** The update rule for the cell state is $C_t = f_t * C_{t-1} + i_t * \tilde{C}_t$. The derivative of $C_t$ with respect to $C_{t-1}$ contains the forget gate $f_t$. If $f_t$ is close to 1, the gradient passes through almost unchanged, effectively bypassing the severe vanishing gradient problem of repeated matrix multiplication seen in vanilla RNNs.

## 3. Explain the difference between 1D, 2D, and 3D Convolutions. Provide a use case for each.
**Answer:**
*   **1D Convolutions (Conv1D):**
    *   **Mechanism:** The filter sweeps along a single dimension.
    *   **Data Type:** Sequence data, time-series, or text.
    *   **Use Case:** Analyzing financial time-series data to predict stock prices, analyzing sensor data (e.g., electrocardiograms), or text classification (where words are embedded as vectors, and the filter slides over the word sequence to capture n-grams).
*   **2D Convolutions (Conv2D):**
    *   **Mechanism:** The filter sweeps across two dimensions (height and width).
    *   **Data Type:** Images or spatial data.
    *   **Use Case:** Image classification (ResNet, VGG), object detection (YOLO), or analyzing spectrograms (audio represented as a 2D image).
*   **3D Convolutions (Conv3D):**
    *   **Mechanism:** The filter sweeps across three dimensions (height, width, and depth/time).
    *   **Data Type:** Volumetric data or video data.
    *   **Use Case:** Medical imaging (e.g., MRI or CT scans which have depth), or Video Action Recognition (where the third dimension is time/frames, capturing both spatial features in frames and temporal motion across frames).

## 4. In Computer Vision, what is the role of the Receptive Field? How does it increase in a CNN?
**Answer:**
**Definition:**
The receptive field of a specific neuron in a deep CNN layer is the region in the *original input image* that affects that neuron's activation.

**Why it matters:**
Early layers in a CNN have small receptive fields and learn local features like edges and corners. Deeper layers capture higher-level concepts (like a "wheel" or a "face") because their receptive fields are much larger, allowing them to "see" a larger portion of the input image.

**How it increases:**
1.  **Successive Convolutions:** Even if you use a small $3\times3$ filter, applying it layer after layer implicitly increases the receptive field. A neuron in layer 2 looking at a $3\times3$ patch of layer 1 is actually looking at a typical $5\times5$ patch of the original image.
2.  **Pooling Layers (Max/Average Pooling):** Pooling immediately scales down the spatial dimensions (e.g., halving height and width). A $3\times3$ convolution performed *after* a $2\times2$ pooling layer effectively covers twice as much area of the original input compared to performing it before pooling.
3.  **Dilated (Atrous) Convolutions:** This involves inserting spaces between the values in a kernel. A $3\times3$ dilated convolution with a dilation rate of 2 acts like a $5\times5$ filter (but with fewer parameters). This is a very efficient way to rapidly increase the receptive field without losing resolution (often used in segmentation tasks like DeepLab).

## 5. Before Transformers dominated NLP, how did architecture like Seq2Seq with Attention work for machine translation?
**Answer:**
**The Pre-Attention Problem:**
Basic Seq2Seq models consist of an Encoder RNN and a Decoder RNN. The Encoder processes the entire input sentence and compresses it into a single, fixed-size vector (the final hidden state, often called the "context vector"). The Decoder must reconstruct the translation from this single vector. This creates a severe bottleneck for long sentences.

**The Attention Mechanism (Bahdanau/Luong):**
Instead of forcing the Encoder to compress everything into one vector, "Attention" allows the Decoder to look back at *all* the hidden states of the Encoder at every step of generation.

1.  **Encoding:** The Encoder calculates a hidden state ($h_i$) for every word in the input sentence.
2.  **Decoding Step (generating word $t$):**
    *   The Decoder has its current hidden state ($s_t$).
    *   **Alignment/Scoring:** The model compares $s_t$ with *every* Encoder hidden state ($h_i$) to calculate a "score" of how relevant input word $i$ is to generating the output word $t$.
    *   **Attention Weights:** These scores are passed through a Softmax function to create a probability distribution (Attention Weights, $\alpha_{ti}$). They sum to 1.
    *   **Context Vector:** We calculate an attention-weighted sum of the Encoder hidden states: $C_t = \sum \alpha_{ti} h_i$. This custom context vector $C_t$ focuses heavily on the relevant input words.
    *   **Generation:** The Decoder uses this targeted context vector $C_t$ AND its own state to predict the next translated word.

**Why it was revolutionary:** It solved the long-range dependency problem in RNNs because the model could directly "attend" to a distant word without the signal having to pass through many recurrent steps.
