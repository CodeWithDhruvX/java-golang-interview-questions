# Practical System Design Problem Statements

This document contains a list of practical System Design interview questions that require combining multiple technologies found in this project's ecosystem (Golang, Java, NodeJS, Angular, React, Kubernetes, Docker, etc.).

## 1. Global E-Commerce Microservices Architecture

**Problem Statement:**
Design a scalable, global e-commerce platform that handles high traffic during flash sales. The system must support user browsing, searching, cart management, order placement, and payments.

**Technology Stack Requirements:**
*   **Search & Catalog Service:** Implement using **Golang** for high-concurrency low-latency search requests.
*   **Order Management Service:** Implement using **Java (Spring Boot)** to handle complex business logic and transactions.
*   **User/Auth Service:** Implement using **Node.js** with JWT authentication.
*   **Frontend:** Build a responsive SPA using **React** or **Angular**.
*   **API Gateway:** Utilize **GraphQL** to aggregate data from multiple microservices for the frontend.
*   **Infrastructure:** Containerize all services using **Docker** and orchestrate with **Kubernetes** (HPA enabled).
*   **Messaging:** Use **Kafka** or **RabbitMQ** for asynchronous order processing (e.g., sending email confirmations, updating inventory).

**Key Challenges to Address:**
*   Handling distributed transactions across Java and Golang services.
*   Ensuring inventory consistency during flash sales.
*   Designing the Kubernetes cluster topology for multi-region availability.

---

## 2. Real-Time Collaborative Code Editor

**Problem Statement:**
Design a web-based code editor like VS Code in the browser or Google Docs for code, where multiple users can edit files simultaneously in real-time.

**Technology Stack Requirements:**
*   **Real-time Collaboration Service:** Implement the Operational Transformation (OT) or CRDT logic using **Golang** (WebSockets).
*   **File System Service:** A **Node.js** service to handle file storage and retrieval (interfacing with S3/Blob storage).
*   **User Session Management:** **Java** microservice for managing user permissions and project metadata.
*   **Frontend:** **React** application with a rich text/code editor component.
*   **Deployment:** Use **Docker Compose** for local development and **Kubernetes** StatefulSets for the collaboration service if needed.

**Key Challenges to Address:**
*   Resolving conflict concurrent edits.
*   Minimizing latency for a smooth typing experience.
*   Persisting the state of the document reliably.

---

## 3. High-Throughput Ride-Share Backend (Uber/Lyft Clone)

**Problem Statement:**
Design the backend for a ride-sharing application that matches riders with nearby drivers in real-time.

**Technology Stack Requirements:**
*   **Location Service:** High-performance **Golang** service using Geo-hashing (Redis Geo) to track driver locations.
*   **Matching Engine:** **Java** service to handle the complex algorithm of matching a rider to the best driver.
*   **Notification Service:** **Node.js** service to push notifications to mobile clients (iOS/Android).
*   **Data Ingestion:** Use **Kafka** to ingest massive streams of location updates from drivers.
*   **Orchestration:** Deploy using **Kubernetes** on a cloud provider (Azure/AWS).

**Key Challenges to Address:**
*   Handling high write throughput for location updates.
*   Efficient spatial indexing and querying.
*   Ensuring reliability of the matching process.

---

## 4. Video Content Delivery & Streaming Platform (Netflix/YouTube Clone)

**Problem Statement:**
Design a video streaming platform that allows users to upload videos, which are then processed and made available for streaming in various resolutions.

**Technology Stack Requirements:**
*   **Upload Service:** **Node.js** service to handle large file uploads and resume capabilities.
*   **Video Processing/Transcoding:** Resource-intensive workers written in **Golang** or **C++** (wrapped in Go) to transcode videos into HLS/DASH formats.
*   **Content Management:** **Java** service for video metadata, ratings, and comments.
*   **Frontend:** **Angular** application for the user dashboard and video player.
*   **Storage:** Object storage for raw and processed videos.
*   **Infrastructure:** Use **Docker** to containerize the transcoding workers and scale them dynamically on **Kubernetes** based on queue depth.

**Key Challenges to Address:**
*   Scalable video processing pipeline.
*   Efficient content delivery using CDNs.
*   Managing state of async processing jobs.

---

## 5. Distributed Task Scheduler

**Problem Statement:**
Design a distributed task scheduler (like Cron but distributed) that can schedule and execute millions of jobs with varying intervals and priorities.

**Technology Stack Requirements:**
*   **Scheduler Core:** **Golang** for precise timing and low-overhead task dispatching.
*   **Job Management API:** **Node.js** REST API for users to submit and monitor jobs.
*   **Execution Agents:** Lightweight **Golang** or **Java** agents running in **Docker** containers that actually run the tasks.
*   **Persistence:** SQL database for job schedules and NoSQL for execution logs.
*   **Coordination:** Use a distributed lock manager (e.g., via Etcd or ZooKeeper often used with **Kubernetes**).

**Key Challenges to Address:**
*   Guaranteeing "at least once" or "exactly once" execution.
*   Handling node failures (scheduler or worker).
*   Syncing clocks or managing drift in a distributed system.

---

## 6. Social Media News Feed

**Problem Statement:**
Design a News Feed generation system for a social media platform where users follow other users and see a timeline of posts.

**Technology Stack Requirements:**
*   **Feed Aggregation Service:** **Java** service that performs fan-out on write (push model) or fan-out on read (pull model).
*   **Post Service:** **Node.js** service for creating and storing posts (text, images).
*   **Graph Service:** **Golang** service or specialized GraphDB to handle followers/following relationships.
*   **Caching:** Extensive use of Redis/Memcached.
*   **Frontend:** **React** mobile-first web view.
*   **API Layer:** **GraphQL** to allow the client to request exactly the feed structure they need.

**Key Challenges to Address:**
*   Latency in generating feeds for celebrities (millions of followers).
*   Pagination and infinite scroll performance.
*   Data consistency across services.
