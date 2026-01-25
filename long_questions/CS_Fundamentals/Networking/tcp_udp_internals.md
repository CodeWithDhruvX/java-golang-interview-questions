# TCP vs UDP Internals

## 1. TCP (Transmission Control Protocol)
**Connection-oriented** protocol that ensures reliable, ordered, and error-checked delivery of a stream of bytes.

### Key Features
*   **Reliable**: Uses Acknowledgments (ACKs) and Retransmissions.
*   **Ordered**: specific sequence numbers ensure packets are reassembled in order.
*   **Flow Control**: Uses **Sliding Window** to prevent overwhelming the receiver.
*   **Congestion Control**: Uses algorithms (Slow Start, Congestion Avoidance) to prevent overwhelming the network.

### 3-Way Handshake (Connection Establishment)
1.  **SYN**: Client sends SYN (Synchronize) packet. "I want to connect."
2.  **SYN-ACK**: Server responds with SYN-ACK. "I received your request, here is my SYN."
3.  **ACK**: Client sends ACK. "Connection established."

### 4-Way Handshake (Connection Termination)
1.  **FIN**: Client sends FIN. "I'm done sending."
2.  **ACK**: Server acknowledges.
3.  **FIN**: Server processes remaining data, then sends its own FIN. "I'm also done."
4.  **ACK**: Client acknowledges.

## 2. UDP (User Datagram Protocol)
**Connectionless** protocol that sends datagrams without establishing a connection.

### Key Features
*   **Unreliable**: No guarantees of delivery. Examples: Dropped packets are just lost.
*   **Unordered**: Packets may arrive out of sequence.
*   **Lightweight**: Low overhead (smalled header size: 8 bytes vs TCP's 20+ bytes).
*   **Fast**: No handshake latency.

## 3. Comparison Table

| Feature | TCP | UDP |
| :--- | :--- | :--- |
| **Connection** | Connection-oriented (Handshake) | Connection-less |
| **Reliability** | Reliable (Retries, ACKs) | Best-effort (No retries) |
| **Ordering** | Guaranteed | Not guaranteed |
| **Speed** | Slower (Overhead) | Faster |
| **Use Cases** | Web (HTTP), Email (SMTP), File Transfer (FTP) | Video Streaming, VoIP, Online Gaming, DNS |

## 4. Interview Questions
1.  **Why does video streaming use UDP?**
    *   *Ans*: In live video, speed matters more than perfection. If a packet is lost, it's better to skip a frame than to pause the video and wait for retransmission (buffering).
2.  **What happens if a TCP packet is lost?**
    *   *Ans*: The sender waits for an ACK. If timeout occurs (or duplicate ACKs received), it retransmits the packet. This slows down throughput (Head-of-Line Blocking).
3.  **Does HTTP/3 use TCP or UDP?**
    *   *Ans*: HTTP/3 uses **QUIC**, which is built on top of **UDP**. It implements reliability and congestion control in the application layer ("userspace TCP") to avoid TCP's head-of-line blocking issues.
