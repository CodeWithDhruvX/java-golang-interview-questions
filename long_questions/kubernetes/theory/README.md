# Kubernetes Theory — Deep-Dive Reference

This directory contains spoken-style theory files covering all Kubernetes concepts from fundamentals to expert-level advanced topics. Each file covers both a concise spoken answer and an in-depth technical explanation.

## 📚 Table of Contents

- [01_Basics_And_Architecture.md](./01_Basics_And_Architecture.md) — What is K8s, Pods, Nodes, Clusters, Namespaces, kubelet, kubectl, API Server, etcd, Scheduler, Controller Manager, kube-proxy, Service Discovery, apply vs create, Minikube vs Kind vs K3s
- [02_Pods_ReplicaSets_Deployments.md](./02_Pods_ReplicaSets_Deployments.md) — Pod lifecycle, ReplicaSets, Deployments, Rolling Updates, Rollbacks, DaemonSets, StatefulSets, Jobs, CronJobs, Init Containers, Sidecar pattern
- [03_Services_And_Networking.md](./03_Services_And_Networking.md) — Service types (ClusterIP/NodePort/LoadBalancer/ExternalName), Ingress, NetworkPolicy, CoreDNS, CNI plugins, Calico, Flannel, pod-to-pod communication
- [04_Storage.md](./04_Storage.md) — Volumes, emptyDir, hostPath, PersistentVolume, PersistentVolumeClaim, StorageClass, dynamic provisioning, StatefulSet storage, CSI
- [05_ConfigMaps_Secrets_Security.md](./05_ConfigMaps_Secrets_Security.md) — ConfigMaps, Secrets, RBAC (Role/ClusterRole/ServiceAccount), SecurityContext, Pod Security Standards, securityContext, encryption at rest
- [06_Autoscaling_And_Scheduling.md](./06_Autoscaling_And_Scheduling.md) — HPA, VPA, KEDA, Cluster Autoscaler, Karpenter, Taints & Tolerations, Node Affinity, Pod Affinity/AntiAffinity, Topology Spread Constraints, QoS classes, PodDisruptionBudget
- [07_Monitoring_Logging_Observability.md](./07_Monitoring_Logging_Observability.md) — Prometheus, Grafana, kube-state-metrics, metrics-server, kubectl top, Loki, Fluent Bit, Jaeger, OpenTelemetry, distributed tracing, alerting
- [08_Helm_And_GitOps.md](./08_Helm_And_GitOps.md) — Helm charts, values/templates/hooks, Helm lifecycle, ArgoCD, Flux, GitOps principles, ApplicationSet, sync policies
- [09_Advanced_And_Production.md](./09_Advanced_And_Production.md) — CRDs, Operators, Admission Webhooks, Multi-tenancy, Cluster HA, etcd backup/restore, node upgrade strategy, production hardening
- [10_Troubleshooting_And_Debugging.md](./10_Troubleshooting_And_Debugging.md) — CrashLoopBackOff, OOMKilled, Pending pods, ImagePullBackOff, node not ready, pod eviction, kubectl debug, ephemeral containers, events, resource pressure
- [11_eBPF_Operators_Sandboxing.md](./11_eBPF_Operators_Sandboxing.md) — eBPF internals (verifier, JIT, hook types, BPF maps), Cilium (kube-proxy replacement, L7 NetworkPolicy, XDP DDoS mitigation), Operator pattern deep dive (Informers, Finalizers, Status subresource, Webhook), RuntimeClass, gVisor Sentry architecture, Kata Containers, sandboxing threat model & trade-offs
