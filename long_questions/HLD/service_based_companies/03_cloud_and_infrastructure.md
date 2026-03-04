# High-Level Design (HLD): Cloud and Infrastructure Basics

Service-based companies deploy and manage enterprise applications on major cloud platforms (AWS, Azure, GCP). Understanding basic cloud architecture is essential.

## 1. IaaS vs. PaaS vs. SaaS: Give examples.
**Answer:**
*   **IaaS (Infrastructure as a Service):** You rent raw virtual hardware (servers, networking, storage). You manage the OS, runtime, data, and application.
    *   *Examples:* AWS EC2, Azure Virtual Machines, Google Compute Engine.
*   **PaaS (Platform as a Service):** The cloud provider manages the OS, runtime (Java/Node.js), and servers. You just deploy your code and data.
    *   *Examples:* AWS Elastic Beanstalk, Heroku, Azure App Services.
*   **SaaS (Software as a Service):** Fully functional, hosted software accessed over the internet. You just consume it.
    *   *Examples:* Salesforce, Google Workspace, Microsoft Office 365, Atlassian Jira.

## 2. How do you achieve High Availability (HA) in a Cloud environment?
**Answer:**
High Availability ensures a system stays operational even if individual components fail.
*   **Redundancy across Availability Zones (AZs):** An AZ is a distinct, isolated physical data center within a cloud region. Deploying multiple application servers across AZs ensures that if one data center goes down (e.g., due to power loss), the app stays online via the other AZ.
*   **Load Balancing:** Use Cloud Load Balancers (like AWS ALB or Azure Application Gateway) to distribute traffic to healthy nodes and automatically stop routing traffic to failed ones.
*   **Database Replication:** Run a Master DB in AZ-1 and a synchronous Standby Replica in AZ-2 (e.g., AWS RDS Multi-AZ). If the Master crashes, the cloud provider automatically fails over to the Standby.

## 3. Explain Auto-Scaling and how it works.
**Answer:**
Auto-scaling is the process of automatically adding (scaling out) or removing (scaling in) compute resources based on real-time demand.
*   **Metrics:** It monitors metrics like CPU utilization, Memory usage, Network traffic, or custom queue lengths.
*   **Triggers:** If CPU > 70% for 5 minutes, an alarm triggers a scaling policy.
*   **Action:** The policy instructs the Auto Scaling Group (ASG) to spin up 2 new Virtual Machines (based on an established image/template) and attach them to the Load Balancer.
*   *Why it's important:* It handles traffic spikes gracefully without manual intervention, and saves the client money during low-traffic periods by terminating unneeded servers.

## 4. What is a VPC (Virtual Private Cloud) and why is it important for Enterprise Security?
**Answer:**
A VPC is a logically isolated virtual network within a public cloud dedicated solely to your resources. It's essentially your own private data center in the cloud.
*   **Subnets:** It allows dividing the network into Public and Private subnets.
*   **Enterprise Architecture Pattern:**
    *   *Public Subnet:* Contains resources that must be internet-facing. Specifically, the Load Balancer, API Gateway, and NAT Gateways.
    *   *Private Subnet:* Contains the Application/Web Servers and the Databases. These have NO direct inbound internet access.
*   **Security:** By placing Databases deep within private subnets, protected by strict Security Groups and Network ACLs (Access Control Lists), the enterprise protects its sensitive data from direct external attacks.

## 5. Explain Docker and Kubernetes at a high level.
**Answer:**
*   **Docker (Containerization):** Packages an application and all its dependencies (libraries, runtime, OS binaries) into a single, standardized unit called a *Container*. This solves the "it works on my machine" problem, ensuring the app runs exactly the same everywhere.
*   **Kubernetes (K8s) (Orchestration):** When an enterprise has hundreds of containers across dozens of servers, managing them is impossible manually. K8s automates the deployment, scaling, healing, and management of these containers.
    *   If a container crashes, K8s restarts it. 
    *   It load-balances traffic among duplicate containers.
    *   It handles zero-downtime rolling deployments.
