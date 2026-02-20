# 🔵 Services & Networking

---

### 1. What is a Service in Kubernetes?

"A Service is a **stable network endpoint** that provides a consistent way to access a set of pods.

Pods are ephemeral — they come and go, changing IPs constantly. A Service gives you a fixed DNS name and virtual IP (ClusterIP) that always routes to healthy pods matching its label selector.

I define Services alongside every Deployment. Without one, there's no reliable way for other services or external clients to reach your pods."

#### In Depth
Services use **endpoint slices** (and legacy endpoints) to track ready pod IPs. The kube-proxy on each node programs iptables/IPVS rules based on these. When pods fail their readiness probe, they're removed from endpoints automatically — this is the key integration between health checks and traffic routing.

---

### 2. What are the types of Kubernetes Services?

"There are four types:

1. **ClusterIP** (default): Accessible only within the cluster. Perfect for service-to-service communication.
2. **NodePort**: Exposes the service on a port on each node's IP. Limited use — mostly for dev/testing.
3. **LoadBalancer**: Provisions a cloud load balancer (ALB, NLB, etc.) for external access. The standard for production external exposure.
4. **ExternalName**: Maps a service to an external DNS name (CNAME). Useful for abstracting external dependencies.

In production, I almost exclusively use ClusterIP for internal services and LoadBalancer or Ingress for anything external-facing."

#### In Depth
NodePort range is 30000-32767 by default. LoadBalancer type implicitly creates a NodePort and ClusterIP too. ExternalName works at the DNS level — no proxying happens. There's also a less-known pseudo-type: **Headless** service (ClusterIP: None), which bypasses the virtual IP and returns pod IPs directly from DNS.

---

### 3. What is a ClusterIP service?

"ClusterIP is the **default service type**. It creates an internal virtual IP that's only routable within the cluster.

kube-proxy programs iptables rules on every node so that any traffic to the ClusterIP:Port gets DNAT'd to one of the backing pod IPs.

This is what I use for all internal microservice communication — no external exposure, just stable DNS-based addressing within the cluster."

#### In Depth
The ClusterIP is allocated from a configurable CIDR (`--service-cluster-ip-range`). It's a **virtual IP that doesn't exist on any interface** — it only exists in iptables/IPVS rules. It can't be pinged directly. If you're debugging connectivity, `kubectl exec` into a pod and use `curl` or `nslookup` — not `ping`.

---

### 4. What is a NodePort service?

"NodePort exposes your service on **every node's IP at a static port** (30000-32767).

External traffic can reach `<any-node-ip>:<nodePort>`, which K8s forwards to the service.

I use NodePort rarely — mostly for exposing services in Minikube during development. In production, it's too coarse: you're tied to specific node IPs, there's no SSL termination, and it bypasses load balancing across nodes."

#### In Depth
NodePort has a critical limitation: `externalTrafficPolicy: Cluster` (default) does an extra hop — traffic may arrive at Node A but be forwarded to a pod on Node B, adding latency. Setting `externalTrafficPolicy: Local` avoids the extra hop but can cause uneven load distribution if pods aren't evenly spread.

---

### 5. What is a LoadBalancer service?

"LoadBalancer extends NodePort and **provisions a cloud provider's load balancer** automatically.

When you create a Service of type LoadBalancer in GKE, EKS, or AKS, the cloud controller manager calls the cloud API to provision an ALB or NLB. The service gets an external IP.

This is the simplest way to expose a service externally, but it's expensive — each service gets its own LB. For multiple services, I use **Ingress** to share a single LB."

#### In Depth
The LoadBalancer type can be customized with annotations specific to the cloud provider: e.g., `service.beta.kubernetes.io/aws-load-balancer-type: nlb` for AWS NLB instead of classic ELB. MetalLB is used in bare-metal environments to achieve the same behavior without a cloud provider.

---

### 6. What is an Ingress?

"Ingress is an **L7 (HTTP/HTTPS) routing layer** that sits in front of multiple services and routes based on hostname or path.

It lets you expose many services through a **single external IP/load balancer**, reducing cost.

Example: `api.example.com/users` → users-service, `api.example.com/orders` → orders-service. All through one Ingress and one load balancer."

#### In Depth
An Ingress resource is just a configuration object — you also need an **Ingress Controller** (NGINX, Traefik, AWS ALB Ingress Controller, etc.) to implement it. The controller watches Ingress resources and configures the underlying proxy. This separation means you can swap controllers without changing your Ingress specs.

---

### 7. Difference between Ingress and LoadBalancer?

"LoadBalancer creates one cloud LB per service — expensive and not path-routable.

Ingress uses **one load balancer shared across many services** with L7 routing rules (by host, path, header, etc.).

In production: use LoadBalancer only for TCP/UDP services that can't go through HTTP routing (e.g., a gRPC service that needs direct LB), and use Ingress for all HTTP/HTTPS APIs."

#### In Depth
Another key difference: Ingress supports **TLS termination** natively — specify a TLS secret, and the Ingress controller handles SSL. With raw LoadBalancer services, you'd need to configure SSL at the pod level. Ingress also handles redirects, rewrites, and rate limiting through controller-specific annotations.

---

### 8. How does DNS work in Kubernetes?

"CoreDNS runs as a Deployment in the `kube-system` namespace and handles all cluster DNS.

Every Service and headless Pod gets a DNS record. The format is `<service>.<namespace>.svc.cluster.local`. Pods have `resolv.conf` pre-configured to search `<namespace>.svc.cluster.local` so you can use just the service name within the same namespace.

I debug DNS issues with `kubectl exec -it <debug-pod> -- nslookup my-service` or by running a `dnsutils` pod."

#### In Depth
CoreDNS is highly configurable via its `Corefile` ConfigMap. Common production customizations include: adding external DNS forwarding rules (for custom internal domains), increasing cache TTL for performance, and adding rewrite rules. `ndots:5` in pod resolv.conf means DNS queries are tried with search domains appended before a direct query — this can cause latency for external DNS lookups.

---

### 9. What is a NetworkPolicy?

"A NetworkPolicy is a **firewall rule for pods** — it controls which pods can communicate with which other pods or external endpoints.

By default, all pods can talk to all other pods. NetworkPolicies allow you to implement **least-privilege networking**.

In production, I always define a default-deny policy and then explicitly allow necessary traffic. For example: only the payments-api pods can reach the database pods."

#### In Depth
NetworkPolicies require a CNI plugin that supports them — Calico, Cilium, Weave Net, etc. Flannel alone does NOT enforce NetworkPolicies. The policy applies to ingress, egress, or both. Entries can match by podSelector, namespaceSelector, or ipBlock. Testing policies is tricky — I use `kubectl exec` to test connectivity and Cilium's policy visualization tooling in Hubble.

---

### 10. How do you expose a service to the outside world?

"Three main options:

1. **Ingress** (preferred for HTTP): One shared LB, L7 routing, TLS termination.
2. **LoadBalancer Service**: Simple but one LB per service — use for non-HTTP or when you need a dedicated external IP.
3. **NodePort**: Fine for dev/testing but not production.

For internal tools, **Ingress + basic auth or OAuth2-proxy** is my go-to. For production APIs, Ingress with cert-manager for automatic TLS via Let's Encrypt."

#### In Depth
There's also the option of a service mesh gateway (Istio Gateway or Envoy-based) for very fine-grained traffic control. This is an L7 proxy at the edge with features like rate limiting, JWT validation, and retries built in. In AWS, you can also use the AWS Load Balancer Controller which creates ALBs directly from Ingress resources.

---

### 11. What is CNI (Container Network Interface)?

"CNI is the **plugin interface that Kubernetes uses to configure pod networking**.

When a pod is created, K8s calls the CNI plugin to: assign an IP address, configure the network interface inside the pod, and set up routing on the host.

Popular CNI plugins: **Calico** (BGP routing, rich NetworkPolicy), **Cilium** (eBPF-based, best performance), **Flannel** (simple VXLAN overlay), **Weave Net** (encrypted VXLAN)."

#### In Depth
CNI plugins implement the `/opt/cni/bin/` interface — they're scripts or binaries called by the kubelet. Kubernetes itself doesn't know or care about the underlying networking implementation. This is why you can swap CNI plugins (with some downtime). Each plugin has different IPAM, routing, and policy capabilities.

---

### 12. What are common CNI plugins and their differences?

"The main plugins I've worked with:

- **Flannel**: Simple VXLAN overlay. No NetworkPolicy support. Good for simple clusters.
- **Calico**: BGP routing (no overlay in L3 networks), full NetworkPolicy, excellent for on-prem.
- **Cilium**: eBPF-based, replaces kube-proxy, best for high-performance environments. Rich observability with Hubble.
- **Weave Net**: Encrypted overlay, simpler configuration. Falls behind Calico/Cilium in features.

In production, I default to **Calico** for compatibility or **Cilium** when we need performance and observability."

#### In Depth
Cilium's eBPF advantage: it hooks into the kernel at a lower level than iptables, reducing packet processing latency significantly. Benchmark tests show 30-40% lower latency vs iptables-based networking. Cilium can also replace kube-proxy entirely. The tradeoff: requires a recent kernel (4.9+, preferably 5.10+) and more operational expertise.

---

### 13. What is Calico and how does it work?

"Calico implements K8s networking using **BGP** (Border Gateway Protocol) instead of overlay networks.

In a datacenter with L3 routing, Calico has each node advertise pod CIDRs via BGP. Pods communicate with native IP routing — no encapsulation overhead.

It also has a rich NetworkPolicy engine that supports egress and ingress policies, global policies, and host policies. In AWS VPCs where BGP isn't available, it falls back to VXLAN."

#### In Depth
Calico uses **Felix** (the policy daemon), **BIRD** (the BGP daemon), and **Confd** (for configuration). The calico-node DaemonSet runs all three. For large clusters, Calico recommends using a Route Reflector topology rather than full mesh BGP to reduce the number of BGP sessions. Calico Enterprise adds threat detection and compliance reporting.

---

### 14. What is kube-proxy and what does it do?

"kube-proxy runs on every node as a DaemonSet and implements the **Service networking layer** by programming iptables or IPVS rules.

When a Service is created, kube-proxy adds rules so that traffic to the ClusterIP gets routed to actual pod IPs. It watches the API server for Service and Endpoint changes and keeps rules in sync.

In high-scale deployments, IPVS mode is significantly faster — iptables scans rules linearly, IPVS uses hash tables. With 10,000+ services, the difference is dramatic."

#### In Depth
kube-proxy has three modes: `userspace` (deprecated, traffic goes through user space), `iptables` (default, kernel space NAT), and `ipvs` (kernel space, hash-based LB). Cilium can **replace kube-proxy entirely** using eBPF socket hooks, eliminating the need for NAT and iptables rules altogether.

---

### 15. How does service discovery work with headless services?

"A headless service (`spec.clusterIP: None`) returns **individual pod IPs** from DNS instead of a virtual IP.

For StatefulSets, each pod gets a stable DNS name: `<pod-name>.<service-name>.<namespace>.svc.cluster.local`. This is how distributed systems like Cassandra, Kafka, and Zookeeper know how to address specific peers.

I use headless services whenever I need clients to connect to specific pod instances rather than any healthy pod."

#### In Depth
With a headless service, the DNS query for the service name returns multiple A records — one per ready pod. Client-side load balancing can then be implemented in the app itself (e.g., gRPC's built-in load balancing). This is different from regular services where the client always gets the ClusterIP and kube-proxy handles the LB.

---

### 16. What is an ExternalName service?

"ExternalName creates a Service that resolves to an **external DNS name** (CNAME record).

Example: `spec.externalName: my-database.prod.us-east-1.rds.amazonaws.com` — pods in the cluster can reach `my-db-service` and they'll resolve to the RDS endpoint.

I use this as an abstraction layer for external dependencies. If I need to migrate from RDS to Aurora, I update the ExternalName service — no application config changes needed."

#### In Depth
ExternalName doesn't provision a ClusterIP or load balancer — it's purely DNS. The kube-proxy doesn't set up any iptables rules for it. This means you can't apply NetworkPolicies to ExternalName services in the same way. Also, some old browsers and apps don't follow CNAME chains correctly, so test thoroughly.

---

### 17. What is the role of DNSPolicy?

"DNSPolicy determines how DNS is configured inside a pod. There are four options:

- `ClusterFirst` (default): Use cluster DNS (CoreDNS) first, fall back to node's DNS.
- `ClusterFirstWithHostNet`: Like ClusterFirst but for pods using `hostNetwork: true`.
- `Default`: Use the node's DNS settings (not cluster DNS).
- `None`: Customize DNS completely via `dnsConfig`.

I set `None` with custom `dnsConfig` when I need pods to use specific resolvers — for example, in hybrid cloud setups where certain domains should resolve via a private DNS server."

#### In Depth
The `ndots` setting in `resolv.conf` (default 5 in K8s) causes every hostname with fewer than 5 dots to first get tried with search domain suffixes appended. This means `google.com` triggers 5+ DNS queries before resolving! For latency-sensitive applications, set `dnsConfig.options: [{name: ndots, value: "1"}]`.

---

### 18. What is network isolation with NetworkPolicy?

"NetworkPolicy implements **microsegmentation** — fine-grained firewall rules at the pod level.

A common production pattern:
1. Apply a **default-deny-all** policy in every namespace.
2. Explicitly allow necessary communication paths.

Example: only `api-pod` (label `role: api`) can reach `db-pod` (label `role: db`) on port 5432. Everything else is blocked.

NetworkPolicies are stateful — if you allow ingress, the response traffic is automatically allowed."

#### In Depth
NetworkPolicies are additive — multiple policies applying to the same pod are ORed together. There's no concept of a priority or 'deny' rule within NetworkPolicy itself. For more expressive policies (e.g., deny rules, FQDN-based rules), Calico's GlobalNetworkPolicy or Cilium's CiliumNetworkPolicy CRDs are needed.

---
