# System Design (HLD) - Distributed Web Crawler

## Problem Statement
Design a distributed Web Crawler (like Googlebot) that fetches web pages, extracts data/links, and stores them for a search engine.

## 1. Requirements Clarification
### functional Requirements
*   **Crawling:** Given a set of "seed" URLs, begin downloading pages.
*   **Extraction:** Extract all `<a href>` links from the downloaded HTML to find new pages to crawl.
*   **Storage:** Store the raw HTML content.

### Non-Functional Requirements
*   **Scale:** The web has over 50 Billion pages. The crawler must be distributed across hundreds of machines.
*   **Politeness:** Do not aggressively hammer a single domain (Obey `robots.txt` and rate limits).
*   **Deduplication:** Never crawl the same exact URL twice.

## 2. High-Level Architecture (The Crawl Loop)

```text
       [ Initial Seed URLs ]
               |
               v
      +-> [ URL Frontier (Queue) ] ----+
      |                                |
      |                        [ Crawler Workers ] -----> [ DNS Resolver ]
      |                                |                     (Fetch IP)
      |                                v
      |                       [ HTML Downloader ] -------> (Internet)
      |                                |
      |                                v
      |                     [ Parser / Link Extractor ] -> [ HTML Storage DB ]
      |                                |
      +--------- (New Links) ----------+
          (Passed through Deduplicator)
```

## 3. Core Component Design

### A. URL Frontier (The Brain)
*   The Frontier is essentially a massive priority queue of URLs waiting to be downloaded.
*   **Politeness Policy:** The Frontier queues URLs based on domains. It ensures that a Worker isn't sent 100 links for `wikipedia.org` all at once, which would DDOS Wikipedia. It spaces them out.

### B. The Deduplicator (Bloom Filters)
If we download `google.com/about`, it has a link back to `google.com`. If we blindly follow it, we create an infinite loop. We must track every URL we've ever seen.
*   **Database lookups:** Checking a database of 50 Billion URLs for every single newly discovered link is too slow ($O(log n)$ or network bound).
*   **Memory Set:** A `HashSet` of 50 Billion strings would take Terabytes of RAM.
*   **Solution: Bloom Filter.** A probabilistic data structure that takes up fixed, small memory (e.g., a few hundred MBs). It can tell us very quickly:
    *   "This URL is *definitely not* in the set" (Safe to crawl).
    *   "This URL is *probably* in the set" (Do a slower DB lookup to confirm, or just skip it).

### C. Downloader & DNS Resolution
*   DNS lookups take 10-50ms. A crawler doing millions of requests per second cannot rely on the default OS DNS.
*   We must build or use a custom caching DNS resolver physically close to the Crawler Workers to bring DNS resolution down to < 1ms.

## 4. Storage Design
*   **Metadata DB (URLs, Hash, Last Crawled Date):** A fast NoSQL database like Cassandra or MongoDB.
*   **Raw HTML Storage:** Saving billions of small 50kb HTML files to a standard file system will destroy the inodes and disk read/write heads.
*   **Blob Storage:** Store HTML files in Amazon S3, or bundle thousands of HTML pages together into large binary files (like Hadoop SequenceFiles / BigTable) to optimize disk I/O.

## 5. Follow-up Questions for Candidate
1.  How do you detect if a page is a duplicate, even if the URL is different? (e.g., `site.com/home` and `site.com/`). (Generate a Simhash or MD5 checksum of the stripped HTML content and compare it).
2.  How do you handle crawler traps? (e.g., `site.com/a/b/c/d/e...` generating infinite subdirectories dynamically). (Set a maximum depth limit in the URL Frontier).
