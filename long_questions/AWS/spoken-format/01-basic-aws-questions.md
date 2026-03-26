# AWS Basic Interview Questions - Spoken Format

## 1. What is AWS and why is it so popular?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS and why is it so popular?
**Your Response:** AWS, or Amazon Web Services, is Amazon's cloud computing platform that provides on-demand computing resources and services over the internet. It's popular because it offers pay-as-you-go pricing, massive scalability, reliability with 99.99% uptime SLAs, and a comprehensive suite of services from computing to storage to machine learning. Companies love AWS because it eliminates the need for upfront hardware investments and allows them to focus on building applications rather than managing infrastructure.

---

## 2. What are the main services provided by AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the main services provided by AWS?
**Your Response:** AWS provides a comprehensive portfolio of services organized into categories. The core services include: Compute services like EC2 for virtual servers and Lambda for serverless computing; Storage services like S3 for object storage and EBS for block storage; Database services including RDS for relational databases and DynamoDB for NoSQL; Networking services like VPC, CloudFront CDN, and Route 53 DNS; and specialized services for AI/ML, IoT, analytics, and more. This extensive service catalog allows businesses to build virtually any type of application in the cloud.

---

## 3. What is EC2 in AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is EC2 in AWS?
**Your Response:** EC2, or Elastic Compute Cloud, is AWS's core compute service that provides scalable virtual servers in the cloud. Think of it as renting a computer in AWS's datacenter that you can configure with any operating system, software, and resources you need. The key benefits are elasticity - you can scale up or down based on demand, pay-as-you-go pricing, and the ability to launch instances in minutes. EC2 is the foundation for most cloud applications on AWS.

---

## 4. What is an AMI?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is an AMI?
**Your Response:** An AMI, or Amazon Machine Image, is essentially a template for creating EC2 instances. It contains the operating system, application server, and applications required to launch an instance. Think of it as a snapshot or backup of a configured server that you can use to launch identical instances quickly. AMIs are crucial for maintaining consistency across deployments and for disaster recovery scenarios.

---

## 5. What is the difference between stopping and terminating an EC2 instance?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between stopping and terminating an EC2 instance?
**Your Response:** When you stop an EC2 instance, it's like shutting down a computer - the instance is no longer running but you still pay for the EBS volume storage, and you can start it again later. When you terminate an instance, it's completely deleted along with its attached EBS volumes (unless configured otherwise). Stopping preserves your data and instance configuration, while terminating permanently removes the instance. It's important to choose the right action based on whether you need the instance again.

---

## 6. What is the use of Elastic Load Balancer (ELB)?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the use of Elastic Load Balancer (ELB)?
**Your Response:** Elastic Load Balancer distributes incoming traffic across multiple EC2 instances to ensure high availability and reliability. It's like having a traffic manager that automatically routes requests to healthy instances, preventing any single instance from being overwhelmed. ELB also performs health checks to detect unhealthy instances and automatically stops sending traffic to them. This is essential for building scalable, fault-tolerant applications that can handle varying loads.

---

## 7. What is Auto Scaling?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Auto Scaling?
**Your Response:** Auto Scaling automatically adjusts the number of EC2 instances based on demand. It's like having a smart assistant that adds more servers when traffic increases and removes them when traffic decreases. This ensures you have enough capacity to handle peak loads while saving costs during low traffic periods. Auto Scaling works with ELB to maintain application performance and availability automatically, making it a cornerstone of cloud-native architecture.

---

## 8. What is Amazon S3?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Amazon S3?
**Your Response:** Amazon S3, or Simple Storage Service, is AWS's object storage service that's designed for virtually unlimited scalability. It's like having a massive hard drive in the cloud where you can store any amount of data - from small files to petabytes. S3 is incredibly durable with 99.999999999% durability, meaning your data is highly protected. It's commonly used for website hosting, backup and restore, data archiving, and big data analytics. The best part is you only pay for what you use, making it extremely cost-effective.

---

## 9. What are S3 storage classes?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are S3 storage classes?
**Your Response:** S3 offers different storage classes optimized for different use cases and cost points. Standard is for frequently accessed data with high performance. Standard-IA and One Zone-IA are for infrequently accessed data at lower cost. Glacier is for long-term archival where data retrieval takes minutes to hours. Glacier Deep Archive is the cheapest option for data you rarely access, with retrieval taking 12 hours. The key is choosing the right class based on how often you access the data to optimize costs.

---

## 10. What is Glacier in AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Glacier in AWS?
**Your Response:** AWS Glacier is a secure, durable, and low-cost storage service for data archiving and long-term backup. It's designed for data that doesn't need to be accessed frequently, like compliance records, medical data, or media archives. Glacier offers different retrieval options from minutes to hours, with costs decreasing as retrieval time increases. It's perfect for meeting regulatory requirements for data retention while keeping storage costs minimal.

---

## 11. What is IAM?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is IAM?
**Your Response:** IAM, or Identity and Access Management, is AWS's security service that controls who can access what resources in your AWS account. Think of it as the security guard of your AWS infrastructure. With IAM, you can create users, groups, and roles, and define fine-grained permissions using policies. The principle of least privilege is key here - give users only the permissions they absolutely need to do their job. IAM is fundamental to securing your AWS environment and is often the first service configured when setting up any AWS account.

---

## 12. What are IAM Roles and Policies?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are IAM Roles and Policies?
**Your Response:** IAM policies are JSON documents that define permissions - they specify what actions can be performed on which resources. IAM roles are similar to users but are meant to be assumed by AWS services or applications. For example, an EC2 instance can assume a role to access S3, or a Lambda function can assume a role to call other AWS services. The beauty of roles is that you don't need to manage long-term credentials - AWS handles temporary security tokens automatically.

---

## 13. What is the root account in AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the root account in AWS?
**Your Response:** The root account is the original AWS account that has complete access to all AWS services and resources in the account. It's created when you first sign up for AWS and has unrestricted permissions. Because of its power, AWS recommends using the root account only for tasks that require root access, like creating IAM users or changing account settings. For daily operations, you should create IAM users with appropriate permissions and enable MFA on the root account for security.

---

## 14. What is an AWS Region and Availability Zone?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is an AWS Region and Availability Zone?
**Your Response:** AWS Regions are geographical areas where AWS has multiple data centers. For example, us-east-1 is Northern Virginia, and eu-west-1 is Ireland. Within each region, there are Availability Zones, which are distinct data centers separated by distance to protect against failures. Think of regions as countries and AZs as cities within those countries. When you deploy applications across multiple AZs, you ensure high availability - if one AZ goes down, your application continues running in another AZ.

---

## 15. How is data replicated in AWS across regions?

### How to Explain in Interview (Spoken style format)
**Interviewer:** How is data replicated in AWS across regions?
**Your Response:** AWS provides several ways to replicate data across regions. For S3, you can use Cross-Region Replication which automatically copies objects to a destination bucket in another region. For databases, RDS supports cross-region read replicas and Aurora Global Database for multi-region deployments. For disaster recovery, you can use AWS Backup or create custom replication solutions. The key is choosing the right replication strategy based on your RTO and RPO requirements.

---

## 16. What is EBS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is EBS?
**Your Response:** EBS, or Elastic Block Store, is AWS's block storage service for EC2 instances. Think of it as a virtual hard drive that you can attach to your EC2 instances. EBS volumes are persistent - they continue to exist even after you stop or terminate the instance (unless you configure deletion on termination). EBS offers different volume types optimized for different workloads, from general-purpose SSD to provisioned IOPS for high-performance databases. It's the primary storage solution for most applications running on EC2.

---

## 17. Difference between EBS and Instance Store?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between EBS and Instance Store?
**Your Response:** The key difference is persistence. EBS volumes are network-attached storage that persists independently of the EC2 instance lifecycle, while instance store is temporary storage physically attached to the host computer. If you stop or terminate an EC2 instance, EBS data remains but instance store data is lost. EBS is better for long-term storage and databases, while instance store is ideal for temporary data, buffers, or caching where high I/O performance is critical and data persistence isn't required.

---

## 18. What is AWS Lambda?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS Lambda?
**Your Response:** AWS Lambda is AWS's serverless compute service that lets you run code without provisioning or managing servers. You simply upload your code, and Lambda handles everything needed to run and scale it. You pay only for the compute time you consume - there's no charge when your code isn't running. Lambda automatically scales from a few requests per day to thousands per second, making it perfect for event-driven architectures, microservices, and backend processing.

---

## 19. What triggers can invoke a Lambda function?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What triggers can invoke a Lambda function?
**Your Response:** Lambda functions can be triggered by over 200 AWS services and custom applications. Common triggers include API Gateway for HTTP requests, S3 for object creation or deletion events, DynamoDB streams for database changes, CloudWatch Events for scheduled tasks, and SNS for message notifications. You can also invoke Lambda directly from other Lambda functions, EC2 instances, or external applications using the AWS SDK. This flexibility makes Lambda the glue that connects different services in event-driven architectures.

---

## 20. What is a VPC?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a VPC?
**Your Response:** A VPC, or Virtual Private Cloud, is your own isolated network environment in AWS. Think of it as having your own private data center in the cloud where you have complete control over networking components like IP address ranges, subnets, route tables, and network gateways. VPC allows you to create a secure network environment for your AWS resources, define how they communicate with each other and the internet, and implement network security best practices. Every AWS account starts with a default VPC, but best practice is to create custom VPCs for production workloads.
