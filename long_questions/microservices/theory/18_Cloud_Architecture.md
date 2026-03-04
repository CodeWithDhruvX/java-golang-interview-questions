# 🟣 **281–290: Cloud Architecture & Serverless**

### 281. What is the difference between IaaS, PaaS, and SaaS?
"**IaaS (Infrastructure as a Service):** You rent bare-metal servers, VMs, storage, and networking from AWS (EC2/S3). You are 100% responsible for installing the OS, updating patches, configuring the firewall, and managing the runtime (Docker/Java). High control, high operational overhead.

**PaaS (Platform as a Service):** The cloud provider abstracts away the OS and infrastructure. You just provide the application code (e.g., a `.jar` file). The platform (like AWS Elastic Beanstalk, Heroku, or Google App Engine) automatically handles provisioning, load balancing, and OS patching. Medium control, low overhead.

**SaaS (Software as a Service):** A completely finished, hosted application. You just log in and use it (e.g., Salesforce, Google Workspace). Zero control, zero operational overhead."

#### Indepth
Serverless computing (FaaS - Function as a Service) sits between PaaS and SaaS. It completely abstracts server management, but uniquely, it scales to absolutely zero when not in use, meaning you pay exactly per millisecond of execution time, unlike standard PaaS which charges a flat hourly rate for the underlying VM capacity.

---

### 282. When should you use Serverless architectures?
"Serverless (like AWS Lambda or Azure Functions) is perfect for highly irregular, extremely spiky workloads. If you have an API that gets 0 requests most of the day, but 10,000 requests for exactly 5 minutes at midnight, provisioning EC2 or K8s nodes 24/7 is a massive waste of money. Lambda scales from 0 to 10,000 instantly to handle the burst, and scales back to 0, charging you only for those 5 minutes.

It is also excellent for event-driven 'glue' logic: e.g., triggering a Lambda function automatically whenever a user uploads a photo to an S3 bucket, which generates a thumbnail and saves it to DynamoDB, without managing any polling infrastructure."

#### Indepth
Serverless has significant drawbacks. **Cold Starts:** If a Lambda function hasn't been called in a while, AWS kills the underlying container. The next request forces a new container to spin up, JVM to Boot, and frameworks to load, adding 2-5 seconds of latency to that specific user's request. **Cost at Scale:** While cheap at low volume, running Lambda functions constantly for massive, consistent traffic (e.g., millions of requests per hour) is profoundly more expensive than running dedicated containers on Kubernetes.

---

### 283. How do you manage infrastructure deployments in the cloud?
"In modern cloud architecture, clicking around the AWS Console to create databases and servers is strictly forbidden. It is unrepeatable, untrackable, and profoundly error-prone.

I use **Infrastructure as Code (IaC)**, specifically **Terraform** or AWS CloudFormation. 

I write declarative configuration files (`main.tf`) describing precisely what infrastructure I want (e.g., '1 VPC, 3 Subnets, 1 RDS Postgres DB, 5 EC2 instances'). I check this file into Git. 
When I run `terraform apply`, Terraform reads the state of the real world, compares it to my code, and makes exactly the API calls necessary to make the real world match the code."

#### Indepth
IaC enforces immutable infrastructure. If a server is misconfigured, nobody SSHes into the server to fix it. We modify the Terraform code, tear down the broken server entirely, and instantly provision a brand new, identical, correct server from scratch.

---

### 284. Explain the difference between Multi-AZ and Multi-Region architecture.
"**Multi-AZ (Availability Zone):** A Region (e.g., us-east-1) contains multiple physically isolated data centers called AZs, located miles apart. Deploying an application across three AZs protects against a localized power outage or a severe flood in one building. The network latency between AZs is extremely low (<2ms), allowing synchronous database replication.

**Multi-Region:** Deploying across entirely different geographic areas (e.g., us-east-1 in Virginia and eu-west-1 in Ireland). This protects against massive catastrophic failures (a meteor hitting Virginia, or a region-wide DNS routing error). However, because they are thousands of miles apart, network latency is high (100ms+), meaning databases must use asynchronous replication, leading to severe eventual consistency challenges."

#### Indepth
Global, latency-sensitive applications (like Netflix or Uber) heavily leverage Multi-Region active-active architectures. Users connect to their geographically closest region for speed. Behind the scenes, Cassandra or DynamoDB Global Tables replicate their actions asynchronously across the ocean to maintain global state.

---

### 285. What is the role of a Content Delivery Network (CDN) at scale?
"If my application is hosted in New York, a user downloading a 5MB image in Tokyo experiences severe network latency because the data has to physically travel halfway across the globe.

A CDN (like Cloudflare, AWS CloudFront, or Akamai) solves this by placing hundreds of "Edge Nodes" in data centers in almost every major city globally. 

When the Tokyo user requests the image, they actually hit the Tokyo Edge Node. If the node has the image cached, it returns it instantly (<10ms). If it doesn't, the Tokyo node fetches it from New York once, caches it locally, and serves all subsequent Tokyo users instantly."

#### Indepth
CDNs are not just for static image caching anymore. Edge Computing (like Cloudflare Workers) allows you to deploy actual Javascript/WASM microservices directly to these edge nodes globally. You can execute authentication, A/B testing logic, and dynamic API routing directly at the edge, drastically reducing load on your core origin servers.
