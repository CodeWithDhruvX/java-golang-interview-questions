# System Design (HLD) - Video Streaming (YouTube/Netflix)

## Problem Statement
Design a global video streaming platform like YouTube or Netflix. Users can upload videos, search for videos, and stream videos seamlessly across various devices and network conditions.

## 1. Requirements Clarification
### Functional Requirements
*   **Upload Video:** Content creators can upload videos.
*   **Watch Video:** Users can stream videos without buffering.
*   **Search/Thumbnails:** Users can search for videos and see metadata (thumbnails, title, views).

### Non-Functional Requirements
*   **High Availability:** The streaming service must be highly reliable.
*   **Low Latency Streaming:** Video must start playing quickly.
*   **Scale:** Support millions of concurrent viewers and petabytes of data storage.

## 2. High-Level Design (Architecture)

The architecture is split logically into the **Video Upload Flow** and the **Video Streaming Flow**.

### A. Video Upload Flow
1. **Client** uploads a large raw video file directly to **Cloud Storage** (e.g., AWS S3) using multipart upload.
2. The Client tells the **API Servers**, "I uploaded a video."
3. The API Server drops a message into a **Message Queue** (e.g., Kafka or RabbitMQ) containing the video metadata and S3 location.
4. **Video Transcoding Workers** (Processing nodes) pick up the message.
   *   They download the raw video.
   *   They transcode it into multiple formats (MP4, HLS, DASH) and bitrates/resolutions (360p, 720p, 1080p, 4K).
   *   They extract thumbnails.
5. Transcoded videos are saved to **S3 (Processed Storage)**.
6. Processed videos are propagated to the **CDN** (Content Delivery Network).
7. Metadata DB is updated to "Ready to Stream".

### B. Video Streaming Flow
1. **Client** requests to watch a video.
2. The user's device figures out the nearest Edge Server via the **CDN**.
3. The video is streamed directly from the **CDN edge node** closest to the user.
4. If the CDN doesn't have the video (Cache Miss), it fetches it from the origin S3 bucket.

## 3. Streaming Protocols & Transcoding
Streaming isn't just downloading an MP4 file. Modern platforms use **Adaptive Bitrate Streaming**.
*   **Dash / Apple HLS:** The video is chopped into small chunks (e.g., 2-10 second segments).
*   **Adaptability:** The client's video player constantly monitors internet speed. If the network drops, the player seamlessly requests the next 5-second chunk in a lower resolution (e.g., drops from 1080p to 480p) to prevent buffering.

## 4. Components & Databases
*   **Video Storage:** BLOB storage (Amazon S3 / Google Cloud Storage).
*   **CDN:** Akamai, Cloudflare, or custom Edge nodes.
*   **Metadata DB:** Relational Database (MySQL/PostgreSQL) or NoSQL (MongoDB) for video titles, descriptions, uploaders. Must be heavily cached with Redis/Memcached.
*   **Message Brokers:** Kafka is highly recommended to decouple the heavy transcoding jobs from the user-facing APIs.

## 5. Security and Copyrights (DRM)
*   **DRM (Digital Rights Management):** Used to encrypt the video chunks so they cannot be played outside the official app (Widevine for Android/Chrome, FairPlay for Apple).
*   **Pre-signed URLs:** When a user clicks play, the server generates a time-limited pre-signed URL to the CDNs so that video links cannot be shared publicly.

## 6. Follow-up Questions for Candidate
1.  How do you handle a viral video that gets 10 million views in an hour? (CDN caching strategies, Cache Stampede prevention).
2.  How would you design the "View Counter"? (Cannot update the DB directly 10 million times. Use a distributed counter in Redis, or aggregate events in Kafka first).
