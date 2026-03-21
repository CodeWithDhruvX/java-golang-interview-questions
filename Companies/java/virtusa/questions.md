Virtusa Java Developer - AI Technical Interview Insights

I recently attended the Virtusa Java Developer (1 - 3 YOE) AI Technical Round, and it was one of the most practical interviews I’ve taken.

The questions were topic-wise, focusing on both theory + scenario-based reasoning across:

1️⃣ArrayList vs LinkedList

2️⃣Overloading vs Overriding

3️⃣Optional.of(), Optional.ofNullable(), Optional.empty()

4️⃣Spring Framework vs Spring Boot

5️⃣@Component vs @Service vs @Repository

7️⃣REST API basics, HTTP methods

7️⃣@RequestParam vs @PathVariable

8️⃣Java Streams coding
Coding Task: Print employee names with salary > 55,000 using Streams
(I completed both technical and coding parts strongly.)

I didn’t perform enough in the scenario-based questions, so I couldn’t clear this round. But overall, it was a great learning experience, and I’m glad I performed well technically.

On to the next opportunity! 🚀
If you’re preparing for Java interviews, happy to share more insights.



Recently, I attended an interview for the Java Full Stack Developer role at Virtusa. The interview lasted around 1 hour and focused mainly on Java, Spring Boot, and database concepts. I wanted to share some of the questions that were asked, which may help others preparing for similar roles.
🔹 Some of the questions asked during the interview:
• Write a program to convert a List into a Map in Java.
• Write code to find the maximum sum of a subarray for a given array.
• Write a small code example demonstrating a Lambda expression.
• What is the difference between Lambda expressions and Functional Interfaces?
• What are Exceptions in Java?
• Can we handle exceptions inside a Lambda expression?
• What are Optional classes in Java and why are they used?
• What are the SOLID principles in object-oriented design?
• How do you handle global exceptions in Spring Boot?
• How do you secure a Spring Boot application?
• What is Kafka and where is it used?
• What is Normalization in SQL?
• What is the difference between DELETE, TRUNCATE, and DROP in SQL?
• What is the purpose of annotations in Spring Boot / Java?
🔹 Discussion also included:
• Java 8 features
• REST API concepts
• Spring Boot fundamentals
• SQL and database design

1. Core Java – Immutability & Concurrency
What is immutability?
Explain generics in Java.
What new features were introduced in Java 17?
Explain String immutability with an example.
Why is String immutable in Java?
What is ConcurrentModificationException?
What is an Iterator?
If there are 1000 threads, will you create 1000 object copies?
Is creating multiple copies cost-effective?
Discussion on immutability vs thread safety.
Why are immutable objects thread-safe?
2. Collections & Concurrency
Difference between HashMap and ConcurrentHashMap.
How does ConcurrentHashMap work internally?
Why are immutable objects preferred in concurrent environments?
3. Algorithms / Coding
What is a subsequence?
What is the Longest Common Subsequence (LCS)?
Coding Problem:
Input:
 String1 = ABC
 String2 = ACD
Output:
 AC
 Length: 2
4. Spring Core
Difference between @Component and @Service annotation.
5. Spring MVC / REST
Difference between @Controller and @RestController.
Difference between Request Parameter and Path Parameter.
6. Microservices Architecture
Why do we use microservices architecture?
If multiple microservices are communicating and one service goes down, how will you handle it?


🔹 Spring Boot / REST API
Your REST API handles high volume traffic — how do you add caching and circuit breaking?
You’re asked to expose two versions of an API while maintaining backward compatibility. How do you do it?
Your POST request is returning 200 OK but not saving data — how do you debug this in a layered architecture?

🔹 JDBC / Transaction Management
You’ve implemented a batch insert using JDBC, but it fails on large datasets. How do you handle rollback and recovery?
A transaction is marked as @Transactional, but data still rolls back partially. What could be wrong?
You’re using multiple DAOs in a service layer — how do you maintain atomicity across method calls?

🔹 Multithreading / Core Java
You implemented parallel processing using CompletableFuture — but threads are blocking. What’s your fix?
You’ve built a shared in-memory cache — how do you prevent data inconsistency across threads?
How do you safely shutdown a thread pool in production without losing tasks?

🔹 Servlets, Filters & JSP
You used a servlet for file downloads, but some files are getting corrupted — how do you fix the stream handling?
You’re applying filters for request logging — but async requests aren’t captured. What’s missing?
A legacy JSP page is exposing server config data — how do you secure it without rewriting the whole page?

🔹 Security & Session Management
You’re using Spring Security and JWT — how do you implement token refresh?
Users are reporting unexpected logout issues — how would you debug session management under a load balancer?
How do you implement IP whitelisting for sensitive APIs in Spring Boot?




Today, I had the opportunity to attend a technical interview with #Virtusa.

It was a truly insightful experience where the discussion went deep into real-time production scenarios rather than just theoretical questions.
What I really appreciated was the focus on advanced, scenario-based problem solving across DevOps, Cloud, and SRE practices.

Here are some of the advanced topics and scenarios discussed:

🔹 AWS – Real Production Scenarios
• How would you design a highly available 3-tier architecture across multiple AZs?
 • If an ALB is healthy but users still get 502 errors, how would you troubleshoot?
 • How would you implement cross-account access securely?
 • Design DR strategy with RTO/RPO considerations.
 • How would you reduce AWS cost in a Kubernetes-based production workload?
 • Difference between NLB vs ALB in microservices architecture.

🔹 Jenkins – Enterprise CI/CD
• How do you design a scalable Jenkins architecture for 100+ microservices?
 • How would you implement dynamic agents using Docker or Kubernetes?
 • If a pipeline succeeds but deployment fails in production, how do you handle rollback?
 • How do you secure secrets in Jenkins?
 • How would you implement parallel stages efficiently?

🔹 Docker – Production Challenges
• How do you reduce Docker image size in production?
 • Multi-stage builds – real use case?
 • How do you handle container security vulnerabilities?
 • What causes container OOMKilled and how to fix it?
 • How would you debug a container that works locally but fails in Kubernetes?

🔹 Kubernetes – Advanced Troubleshooting
• Pod stuck in CrashLoopBackOff – debugging steps?
 • How to design HPA for CPU + custom metrics?
 • Difference between Deployment vs StatefulSet in real production?
 • How to implement blue-green or canary deployment?
 • If cluster nodes are under memory pressure, what actions will you take?
 • How do you manage secrets securely in Kubernetes?

🔹 Terraform – Infrastructure as Code
• How do you manage remote state securely?
 • How to handle drift detection?
 • How do you structure Terraform modules for enterprise projects?
 • What is the difference between count and for_each?
 • How do you implement multi-environment (dev/stage/prod) architecture?

🔹 SRE & Reliability Engineering
• How do you define SLI, SLO, and SLA in production?
 • How do you calculate error budget?
 • Incident handling strategy in Sev-1 outage?
 • How do you reduce MTTR?
 • How would you design monitoring and alerting for microservices?
 • Difference between reactive vs proactive reliability engineering?
Overall, it was a strong reminder that: