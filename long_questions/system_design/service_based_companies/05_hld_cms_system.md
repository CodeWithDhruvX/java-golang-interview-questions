# System Design (HLD) - Content Management System (CMS)

## Problem Statement
Design a massive Content Management System (CMS) used by reporters (like the New York Times or an enterprise blog network) to write, review, and publish millions of articles, distributing them to readers globally.

## 1. Requirements Clarification
### Functional Requirements
*   **Authoring:** Writers can create rich-text drafts, upload images, and save revisions.
*   **Editorial Workflow:** Editors can review, approve, and schedule articles for publication.
*   **Publishing:** Articles go live globally and are searchable.
*   **Reading:** Millions of users browse the published articles.

### Non-Functional Requirements
*   **Read-Heavy:** The ratio of readers to writers is astronomical (e.g., 100,000:1). Reading must be fiercely optimized.
*   **Strong Consistency for Editors:** When a writer saves a draft, they must see their saved draft instantly.
*   **Eventual Consistency for Readers:** It's acceptable if a new article takes 60 seconds to propagate to the edge CDNs.

## 2. High-Level Architecture (The CQRS Approach)

Because reads and writes have vastly different traffic patterns and requirements, the most elegant architectural pattern here is **CQRS (Command Query Responsibility Segregation)**. We physically split the write-side (Authoring) from the read-side (Publishing).

### A. The Write Path (Content Authoring)
```text
[ Reporters ] ---> [ Authoring API Gateway ]
                          |
                  [ Drafts Microservice ] ---> [ RDBMS (PostgreSQL) ]
                          |
                (Event "Article_Approved")
                          |
                 [ Kafka Event Stream ]
```
*   **Database (RDBMS):** We use a relational SQL database for the authoring side because we need strict ACID transactions. We need to track lock status (so two editors don't overwrite each other) and maintain a heavy `Revisions` table so authors can undo changes.

### B. The Publish Process (The "Compiler")
```text
                 [ Kafka Event Stream ]
                          |
                [ Publishing Worker ] (Consumes "Article_Approved")
               /          |            \
             /            |              \
   [ ElasticSearch ]   [ MongoDB ]     [ Pre-render to S3 ]
    (For Search)      (For API reads)  (Static HTML gen)
```
*   When approved, the Publishing Worker takes the complex relational SQL tables (Article, Tags, Author, Categories) and *flattens* them into a single, massive JSON blob.
*   It saves this JSON into a NoSQL Read-Database (like MongoDB) because NoSQL has extremely fast `_id` fetching.
*   It indexes the text in Elasticsearch.

### C. The Read Path (Content Delivery)
```text
[ Global Readers ] ---> [ CDN (Cloudflare/Akamai) ] ---> [ Delivery API Gateway ]
                                                                 |
                                                         [ Redis Cache ]
                                                                 |
                                                         [ MongoDB (Reads) ]
```
*   **CDN First:** The absolute most critical component. The actual text/HTML and images are cached at CDN Edge Nodes worldwide. 95% of traffic never even touches the backend servers.
*   **Redis Second:** For dynamic API requests (e.g., "Get the latest 5 headlines"), the Delivery API checks a Redis cluster.
*   **MongoDB Third:** Only if Redis misses does the API query the NoSQL document store.

## 3. Media Handling
Images cannot be stored in SQL/NoSQL.
*   Reporter uploads a 5MB TIFF image.
*   Saved to a raw S3 bucket.
*   An Image Processing lambda resizes it into WebP, creates thumbnails, and saves it to a public S3 bucket fronted by a CDN.
*   The CMS saves only the CDN URL strings in the database.

## 4. Follow-up Questions for Candidate
1.  How do you implement the "Save Draft" feature safely without constant DB locking? (Autosaving to an in-memory cache temporarily, or pessimistic locking of the article record).
2.  How do you invalidate the CDN cache when breaking news gets updated? (CDNs offer API endpoints to "Purge" cache by URL or Tag. The Publishing Worker calls the Purge API when an update event flies through Kafka).
