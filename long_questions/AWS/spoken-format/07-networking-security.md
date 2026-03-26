# AWS Networking and Security - Spoken Format

## 1. How do you secure an EC2 instance?

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you secure an EC2 instance?
**Your Response:** I secure EC2 instances using a defense-in-depth approach. First, I use Security Groups as the primary network firewall, allowing only necessary ports and protocols. I place instances in private subnets when possible and use Bastion hosts for management access. I disable password authentication and use key pairs for SSH access, implementing IAM roles instead of storing credentials on the instance. I keep the operating system and applications updated using Systems Manager Patch Manager. For additional security, I enable Amazon Inspector for vulnerability assessment and install the CloudWatch agent for monitoring. I also encrypt EBS volumes and enable instance metadata service version 2 to protect against SSRF attacks. Finally, I regularly audit security groups and review CloudTrail logs for suspicious activity.

---

## 2. What are key pairs in EC2?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are key pairs in EC2?
**Your Response:** EC2 key pairs are SSH credentials used to securely access Linux instances or decrypt Windows administrator passwords. Each key pair consists of a public key that AWS stores and a private key that you download and keep secure. When you launch an instance, AWS places the public key on the instance, and you use the corresponding private key to authenticate. It's crucial to never share your private key and to store it securely. For production environments, I recommend using IAM roles instead of key pairs for applications, and using AWS Systems Manager Session Manager for secure access without opening SSH ports in security groups. Key pairs are essential for initial access but should be part of a broader security strategy.

---

## 3. What is Bastion Host?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Bastion Host?
**Your Response:** A Bastion Host is a dedicated server designed to be the single entry point for accessing instances in private subnets from the internet. Think of it as a secure gateway or gatekeeper - it sits in a public subnet with strict security controls, and administrators connect to it first, then use it to access other private resources. I configure Bastion Hosts with minimal services, strong security groups allowing only specific IP addresses, and comprehensive logging. Modern best practice is to use AWS Systems Manager Session Manager instead of traditional Bastion Hosts, as it provides secure access without opening inbound ports. Bastion Hosts are essential for managing private infrastructure while maintaining security boundaries.

---

## 4. How does VPC peering work?

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does VPC peering work?
**Your Response:** VPC peering allows two VPCs to communicate with each other as if they were part of the same network. Think of it as creating a private network connection between VPCs that allows resources to communicate using private IP addresses. The peering connection is a one-to-one relationship - it's not transitive, meaning if VPC A peers with VPC B, and VPC B peers with VPC C, VPC A cannot communicate with VPC C unless there's a direct peering connection. I use VPC peering for sharing resources between accounts or environments, like allowing development VPCs to access a shared database VPC. For complex networks with many VPCs, I prefer using Transit Gateway which provides a hub-and-spoke model that's more scalable than multiple peering connections.

---

## 5. What is AWS Direct Connect?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS Direct Connect?
**Your Response:** AWS Direct Connect provides a dedicated private network connection between your datacenter and AWS. Think of it as building a private highway to AWS instead of using the public internet. This provides more consistent network performance, lower latency, and higher bandwidth than internet-based connections. Direct Connect offers two connection types: dedicated connections providing 1 Gbps or 10 Gbps capacity, and hosted connections sharing capacity from partners. I recommend Direct Connect for workloads requiring consistent performance like large data transfers, real-time applications, or when you have high bandwidth requirements. It's also essential for hybrid cloud architectures where you need reliable connectivity between on-premises and AWS resources.

---

## 6. What are transit gateways?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are transit gateways?
**Your Response:** AWS Transit Gateway is a cloud router that connects VPCs and on-premises networks through a central hub. Think of it as an airport hub - instead of creating individual connections between every VPC, you connect all VPCs to the Transit Gateway, and it handles routing between them. This simplifies network management as you add more VPCs - you just connect new VPCs to the Transit Gateway instead of creating multiple VPC peering connections. Transit Gateway supports routing between VPCs, VPN connections, and Direct Connect gateways. It also provides features like route tables for controlling traffic flow and multicast support for streaming applications. For organizations with many VPCs or hybrid architectures, Transit Gateway is much more scalable than traditional VPC peering.

---

## 7. What are the differences between Security Groups and NACLs?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the differences between Security Groups and NACLs?
**Your Response:** Security Groups and NACLs both control network traffic, but work differently. Security Groups operate at the instance level and are stateful - if you allow inbound traffic, the return traffic is automatically allowed. NACLs operate at the subnet level and are stateless - you must explicitly allow both inbound and outbound traffic. Security Groups evaluate allow rules only, while NACLs evaluate both allow and deny rules. I use Security Groups as the primary security mechanism because they're more flexible and easier to manage. NACLs provide an additional layer of subnet-level protection and are useful for blocking specific IP ranges or implementing network-wide rules. The key difference is that Security Groups are for instance protection, while NACLs are for subnet protection.

---

## 8. How do you restrict access to S3 buckets?

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you restrict access to S3 buckets?
**Your Response:** I restrict S3 bucket access using multiple layers of security controls. First, I enable "Block Public Access" at the account and bucket levels to prevent accidental public exposure. Then I use S3 bucket policies to specify exactly who can access the bucket and what actions they can perform. For fine-grained control, I use IAM policies to grant access to specific users or roles. For additional security, I implement VPC endpoints to ensure S3 traffic stays within the AWS network. I also use S3 access points to create individual access endpoints with their own policies. For sensitive data, I enable MFA Delete and use object-level ACLs or policies. The principle is defense-in-depth - combining multiple controls to ensure only authorized access.

---

## 9. What is the difference between symmetric and asymmetric encryption in AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between symmetric and asymmetric encryption in AWS?
**Your Response:** Symmetric encryption uses the same key for both encryption and decryption, while asymmetric encryption uses a pair of keys - a public key for encryption and a private key for decryption. In AWS, symmetric encryption is commonly used for encrypting data at rest, like EBS volumes or S3 objects, because it's faster and more efficient for large amounts of data. Asymmetric encryption is typically used for key exchange and digital signatures. AWS KMS supports both types - symmetric CMKs for most use cases and asymmetric CMKs for specific scenarios like encryption and decryption or signing and verification. I use symmetric encryption for bulk data and asymmetric encryption for secure key exchange or digital signatures in my applications.

---

## 10. What is AWS WAF?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS WAF?
**Your Response:** AWS WAF, or Web Application Firewall, helps protect web applications from common web exploits that could affect availability or consume excessive resources. Think of it as a security guard for your web applications that sits in front of services like CloudFront or Application Load Balancer. WAF allows you to create rules to block traffic based on conditions like IP address, HTTP headers, or request body content. It provides pre-configured rule sets for common threats like SQL injection and cross-site scripting. I use WAF to protect against OWASP Top 10 vulnerabilities, implement rate limiting to prevent abuse, and create custom rules for application-specific threats. WAF is essential for any public-facing web application to protect against automated attacks and malicious traffic.
