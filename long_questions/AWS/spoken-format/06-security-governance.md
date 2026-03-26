# AWS Cloud Security & Governance - Spoken Format

## 1. What is AWS Inspector?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS Inspector?
**Your Response:** AWS Inspector is an automated security assessment service that helps improve the security and compliance of applications deployed on AWS. Think of it as a security scanner that automatically analyzes your EC2 instances for vulnerabilities or deviations from best practices. Inspector checks for things like exposed ports, insecure software versions, and potential security vulnerabilities. It provides detailed findings with severity ratings and remediation recommendations. I use Inspector to continuously monitor my infrastructure for security issues and integrate it with my CI/CD pipeline to catch security problems before deployment. It's particularly valuable for compliance requirements and for maintaining a strong security posture without manual security audits.

---

## 2. What is AWS GuardDuty?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS GuardDuty?
**Your Response:** AWS GuardDuty is a threat detection service that continuously monitors for malicious activity and unauthorized behavior in your AWS accounts and workloads. Think of it as a smart security guard that uses machine learning and anomaly detection to identify potential threats. GuardDuty analyzes CloudTrail logs, VPC Flow Logs, and DNS logs to detect things like unusual API calls, compromised instances, or data exfiltration attempts. It provides detailed findings with context and remediation steps. I use GuardDuty as part of my defense-in-depth strategy - it catches things that traditional security tools might miss, like credential abuse or lateral movement. The best part is it's fully managed and requires minimal setup.

---

## 3. What is Macie in AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Macie in AWS?
**Your Response:** Amazon Macie is a fully managed data security and data privacy service that uses machine learning and pattern matching to discover and protect sensitive data in AWS. Think of it as a data classification and protection tool that automatically finds sensitive information like PII, financial data, or health records in your S3 buckets. Macie continuously monitors data access patterns and alerts you to unusual access that could indicate a data breach. It provides detailed reports on what sensitive data you have, where it's stored, and who has access to it. I use Macie to help with compliance requirements like GDPR or HIPAA and to ensure we're not accidentally exposing sensitive data in public S3 buckets.

---

## 4. What is AWS Shield?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is AWS Shield?
**Your Response:** AWS Shield is a managed Distributed Denial of Service protection service that safeguards applications running on AWS. There are two tiers: Shield Standard is automatically enabled for all AWS customers at no additional cost and protects against common network and transport layer DDoS attacks. Shield Advanced provides additional protection for sophisticated attacks, with 24/7 access to the AWS DDoS Response Team and cost protection against DDoS-related usage spikes. I use Shield Standard for basic protection and recommend Shield Advanced for critical applications that require higher availability and advanced protection against sophisticated attacks. Shield integrates with CloudFront, ELB, Route 53, and other AWS services to provide comprehensive DDoS protection.

---

## 5. What is SCP in AWS Organizations?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is SCP in AWS Organizations?
**Your Response:** Service Control Policies (SCP) are a type of policy in AWS Organizations that you use to manage permissions in your AWS organization. Think of SCps as guardrails that set the maximum permissions for accounts in your organization. They work at the organization level and apply to all IAM users and roles in the affected accounts. SCps don't grant permissions - they only restrict what's allowed. For example, I might create an SCP that prevents anyone from creating S3 buckets with public access, or restricts which regions can be used. SCps are essential for implementing governance controls across multiple accounts and ensuring compliance with organizational policies. They're a key part of the defense-in-depth security model.

---

## 6. What is the difference between IAM User and IAM Role?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between IAM User and IAM Role?
**Your Response:** The key difference is that IAM Users represent actual people or applications with long-term credentials, while IAM Roles are meant to be assumed temporarily by AWS services or applications. IAM Users have permanent access keys and are typically for human users who need direct access to the AWS console or CLI. IAM Roles don't have credentials - instead, they provide temporary security credentials when assumed by trusted entities. For example, an EC2 instance assumes a role to access S3, or a Lambda function assumes a role to call other AWS services. Roles are more secure because they use temporary credentials and eliminate the need to manage long-term access keys. Best practice is to use roles whenever possible and only create users for people who need console access.

---

## 7. What is the principle of least privilege?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the principle of least privilege?
**Your Response:** The principle of least privilege is a fundamental security concept that means granting only the minimum permissions necessary for users and services to perform their intended tasks. In AWS, this means creating IAM policies that allow only specific actions on specific resources that are absolutely required. For example, instead of giving a Lambda function full S3 access, I'd grant it access only to the specific bucket and operations it needs. This principle limits the potential damage if credentials are compromised or if there's a configuration error. I implement least privilege by regularly auditing permissions, using IAM Access Analyzer to identify unused or excessive permissions, and starting with deny-all policies and adding only necessary permissions. This approach is crucial for maintaining a strong security posture in AWS.

---

## 8. How do you audit changes in AWS resources?

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you audit changes in AWS resources?
**Your Response:** I audit AWS resource changes using multiple services. CloudTrail is the primary tool - it logs every API call made in my account, showing who did what, when, and from where. I enable CloudTrail in all regions and send logs to a central S3 bucket for long-term storage. AWS Config tracks configuration changes and maintains a history of resource configurations, allowing me to see how resources have changed over time. I use Config rules to continuously evaluate resources against compliance policies and alert on violations. For security events, I use GuardDuty to detect malicious activity. I also set up SNS notifications for critical changes and use CloudWatch Logs Insights to analyze log patterns. For compliance reporting, I use AWS Audit Manager to automate evidence collection.

---

## 9. What is multi-factor authentication (MFA) in AWS?

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is multi-factor authentication (MFA) in AWS?
**Your Response:** MFA in AWS adds an extra layer of security by requiring users to provide two or more verification factors to access AWS resources. Typically, this combines something they know (password) with something they have (like a mobile app or hardware token). In AWS, I enforce MFA for the root account and all IAM users with console access. I use virtual MFA devices like Google Authenticator or AWS's own authenticator app for convenience, or hardware tokens for higher security. MFA is crucial because even if a password is compromised, attackers still can't access the account without the second factor. I also use MFA-protected API access for highly privileged operations and implement MFA requirements in IAM policies for sensitive actions.

---

## 10. How do you prevent accidental deletion of AWS resources?

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent accidental deletion of AWS resources?
**Your Response:** I prevent accidental deletion using multiple layers of protection. For S3 buckets, I enable versioning and MFA Delete, which requires MFA authentication to delete versions. For critical resources like databases, I enable deletion protection and create automated backups using AWS Backup. I use IAM policies to explicitly deny delete actions for critical resources and implement SCPs at the organization level. I also use resource tagging and lifecycle policies to manage deletion in a controlled way. For additional safety, I implement approval workflows for deletion operations using services like AWS Service Catalog or custom approval processes. Finally, I regularly review CloudTrail logs for delete actions and set up alerts for any deletion of critical resources.
