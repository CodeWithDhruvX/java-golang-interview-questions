# Kubernetes Interview Questions - PRODUCT BASED COMPANIES

This directory contains comprehensive networking, scaling, and systems design interview questions tailored for Product Based Companies.

## 📚 Table of Contents

- [01_kubernetes_product_based_qa.md](./01_kubernetes_product_based_qa.md) — etcd internals, Pod lifecycle, Scheduler, CNI, CRDs & Operators, StatefulSets, HPA/VPA
- [02_kubernetes_product_based_architecture_networking.md](./02_kubernetes_product_based_architecture_networking.md) — Admission Webhooks, Custom Operators, API Aggregation, eBPF/Cilium, CoreDNS tuning, Dual-stack
- [03_kubernetes_product_based_reliability_tuning.md](./03_kubernetes_product_based_reliability_tuning.md) — GPU scheduling, Descheduler, QoS/OOMKiller, KEDA, CSI, StatefulSet rolling updates, Headroom
- [04_kubernetes_product_based_security_governance.md](./04_kubernetes_product_based_security_governance.md) — Hard multi-tenancy, OPA Gatekeeper, Network Policy L7, Istio mTLS, Bound ServiceAccount Tokens, Global cluster design
- [05_kubernetes_product_based_observability_pdb_canary.md](./05_kubernetes_product_based_observability_pdb_canary.md) — PDB, Karpenter vs CA, Topology Spread, cert-manager, OPA Rego policies, External Secrets rotation, Prometheus Operator, Hubble/eBPF, Pixie, Canary deployments
- [06_kubernetes_product_based_scheduler_istio_security.md](./06_kubernetes_product_based_scheduler_istio_security.md) — Custom Scheduler Framework, Gang Scheduling, Preemption, Istio VirtualService/DestinationRule, Circuit Breaker, Fault Injection, Retry/Timeout, Audit Logging, Falco, FinOps, Multi-cluster GitOps, API deprecation
- [07_kubernetes_product_based_identity_tracing_final.md](./07_kubernetes_product_based_identity_tracing_final.md) — AWS IRSA, GKE Workload Identity, in-cluster API access, Cluster API (CAPI), VCluster multi-tenancy, Jaeger distributed tracing, OpenTelemetry Operator, Velero backup/restore, Kubernetes Gateway API (HTTPRoute/GRPCRoute/TCPRoute/Policies)
- [08_kubernetes_ebpf_operators_sandboxing.md](./08_kubernetes_ebpf_operators_sandboxing.md) — eBPF internals (BPF maps, XDP, TC hooks), Cilium kube-proxy replacement & L7 NetworkPolicy, writing CRDs + Operators in Go (controller-runtime, Informers, Finalizers, Admission Webhooks), RuntimeClass, gVisor vs Kata Containers sandboxing, Submariner cross-cluster networking, Cluster API (CAPI) cluster lifecycle
