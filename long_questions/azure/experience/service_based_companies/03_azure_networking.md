# 📘 03 — Azure Networking
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Virtual Networks (VNet) and Subnets
- Network Security Groups (NSG) and Application Security Groups (ASG)
- VNet Peering vs VPN Gateway
- Azure Load Balancer vs Application Gateway vs Traffic Manager
- ExpressRoute

---

## ❓ Most Asked Questions

### Q1. What is an Azure Virtual Network (VNet) and a Subnet?
- **VNet:** The fundamental building block for your private network in Azure. It enables Azure resources (VMs) to securely communicate with each other, the internet, and on-premises networks. VNets are scoped to a single Region.
- **Subnet:** A VNet can be segmented into one or more Subnets (e.g., Frontend Subnet, Backend Subnet, Database Subnet). Subnets help organize resources and apply granular network security rules.

---

### Q2. What is a Network Security Group (NSG)?
An NSG is used to filter network traffic to and from Azure resources in a VNet. It acts as a basic firewall at the Layer 4 (Transport layer - TCP/UDP) level.

**Key Features:**
- Contains inbound and outbound security rules (Allow/Deny based on Source IP, Destination IP, Port, and Protocol).
- Evaluated by **Priority** (lower number = higher priority, e.g., Rule 100 runs before Rule 200).
- Has Default Rules (e.g., Allow VNet Inbound, Allow Internet Outbound).
- Can be applied to a **Subnet** or a **Network Interface (NIC)** of a VM.

---

### Q3. Explain VNet Peering vs VPN Gateway to connect two VNets.

| Feature | VNet Peering | VPN Gateway (VNet-to-VNet) |
|---------|--------------|-----------------------------|
| **Routing** | Traffic traverses the Microsoft backbone (private). | Traffic is encrypted and goes over an IPsec tunnel. |
| **Complexity** | Very simple to set up. | Requires creating Gateways (takes 30-45 mins). |
| **Bandwidth** | High bandwidth, low latency. | Limited by the VPN Gateway SKU bandwidth. |
| **Cost** | Charged per GB of data transferred. | Charged per hour for the Gateway + data transfer. |
| **Cross-Region** | Global VNet Peering supported. | Supported. |

> **Use Case:** Always prefer VNet Peering for Azure-to-Azure communication. Use VPN Gateways for connecting Azure to On-Premises.

---

### Q4. Compare Azure Load Balancer, Application Gateway, Traffic Manager, and Front Door.

| Service | OSI Layer | Scope | Protocol Support | Primary Use Case |
|---------|-----------|--------|------------------|------------------|
| **Azure Load Balancer** | Layer 4 (Transport) | Regional | TCP, UDP | Internal/External balancing for VMs. No SSL termination. |
| **Application Gateway** | Layer 7 (Application)| Regional | HTTP, HTTPS, WebSockets | Web traffic balancing, SSL termination, WAF (Web App Firewall), URL-based routing. |
| **Traffic Manager** | DNS Level | Global | Any | DNS-based geographical routing and failover across regions. |
| **Front Door** | Layer 7 (Application)| Global | HTTP, HTTPS | Global CDN, WAF, SSL termination, and routing for web apps across regions. |

---

### Q5. What is an Application Security Group (ASG)?
ASGs enable you to configure network security as a natural extension of an application's structure, allowing you to group VMs and define network security policies based on those groups.

**Why use it?**
Instead of writing an NSG rule based on explicit IP addresses:
`Allow Port 80 from Any to IPs (10.0.0.4, 10.0.0.5)`
You apply an ASG to the VMs and write a cleaner rule:
`Allow Port 80 from Any to ASG "WebServers"`
It simplifies NSG management, especially when IPs change natively via auto-scaling.

---

### Q6. What is Azure ExpressRoute?
ExpressRoute lets you extend your on-premises networks into the Microsoft cloud over a **private connection** facilitated by a connectivity provider.

**ExpressRoute vs Site-to-Site VPN:**
- **VPN:** Travels over the public internet (encrypted), latency varies, max bandwidth ~1.25 Gbps to 10 Gbps depending on SKU.
- **ExpressRoute:** Does NOT go over the public internet, more reliable, faster (up to 100 Gbps), lower latency, but is more expensive and requires a telecom provider.

---

### Q7. What are Service Endpoints vs Private Endpoints?
Used to secure Azure PaaS services (like Azure SQL or Storage Accounts) directly to your VNet.

- **Service Endpoints (VNet Service Endpoints):** Provides a secure route over the Azure backbone network. The PaaS resource still uses a **Public IP**, but its firewall rules reject all traffic except traffic originating from your specific VNet subnet.
- **Private Endpoints (Azure Private Link):** Assigns a **Private IP** from your VNet directly to the PaaS service. The service essentially becomes part of your VNet. Highly preferred for strict enterprise security.
