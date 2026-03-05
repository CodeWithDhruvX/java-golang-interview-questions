# DevOps Engineer & SRE Roadmap (6 Weeks)

Target Audience: Infrastructure, CI/CD, and Reliability Engineers.
Primary Focus: Docker, Kubernetes, Azure, Git, Scripting (Python/Golang), and Architecture Observability.

## Overview
This 6-week structured roadmap shifts focus from application-level code to platform-level engineering, infrastructure as code (IaC), and site reliability. It heavily leverages `docker`, `kubernetes`, `azure`, `git`, and `architecture`.

---

## Week 1: Operating Systems & Containerization Fundamentals
*Goal: Understanding process isolation and reliable packaging.*

* **Day 1-2: Advanced Linux & OS Basics**  
  * **Resource**: `CS_Fundamentals`.
  * **Action**: cgroups, namespaces, iptables, file descriptors, CPU & Memory management.
* **Day 3-4: Docker Deep Dive**  
  * **Resource**: `docker`.
  * **Action**: Multi-stage builds, Docker networking (bridge, host, overlay), volumes vs bind mounts.
* **Day 5-6: Advanced Container Debugging**  
  * **Action**: Distroless images, inspecting intermediate layers, tackling container bloat.
* **Day 7: Implementation**  
  * **Action**: Write an optimized Dockerfile and a `docker-compose.yml` for a 3-tier application (React + Go + Postgres).

## Week 2: Source Control & CI/CD Pipelines
*Goal: Automating the path to production.*

* **Day 1: Git Mastery**  
  * **Resource**: `git`.
  * **Action**: Git internals, rebasing vs merging strategies, reflogs, Git hooks, resolving complex conflicts.
* **Day 2-3: Deployment Strategies**  
  * **Resource**: `git`, `architecture`.
  * **Action**: Canary deployments, Blue/Green deployments, A/B testing, Rollbacks.
* **Day 4-6: CI/CD Pipeline Construction**  
  * **Action**: Setting up GitLab CI/GitHub Actions. Implementing linters, unit tests, integration tests, and container registries in the pipeline.
* **Day 7: Security in Pipelines (DevSecOps)**  
  * **Action**: Static Application Security Testing (SAST), Dynamic Application Security Testing (DAST), container vulnerability scanning (e.g., Trivy).

## Week 3: Kubernetes Administration (Part 1)
*Goal: The standard orchestrator.*

* **Day 1: Kubernetes Architecture**  
  * **Resource**: `kubernetes`.
  * **Action**: Control plane components (API Server, etcd, Scheduler, Controller Manager) vs Worker Node components (Kubelet, Kube-proxy).
* **Day 2-3: Core Workloads**  
  * **Resource**: `kubernetes`.
  * **Action**: Pods, ReplicaSets, Deployments, DaemonSets, StatefulSets, Jobs, CronJobs.
* **Day 4-5: Networking in K8s**  
  * **Action**: ClusterIP, NodePort, LoadBalancer, Ingress Controllers, Network Policies.
* **Day 6-7: Practice**  
  * **Action**: Deploy the 3-tier app (from Week 1) into a local `kind` or `minikube` cluster.

## Week 4: Kubernetes Admin (Part 2) & Infrastructure as Code
*Goal: Production-ready K8s and Automating Infrastructure.*

* **Day 1-2: K8s Storage & Configuration**  
  * **Resource**: `kubernetes`.
  * **Action**: ConfigMaps, Secrets, PersistentVolumes (PV), PersistentVolumeClaims (PVC), StorageClasses.
* **Day 3-4: Scaling & Resource Management**  
  * **Action**: Requests vs Limits, Horizontal Pod Autoscaler (HPA), Vertical Pod Autoscaler (VPA), Cluster Autoscaler.
* **Day 5-6: Infrastructure as Code (IaC)**  
  * **Resource**: `azure` (principles apply to Terraform/ARM).
  * **Action**: Terraform basics: state management, modules, providers. Using IaC to provision cloud resources.
* **Day 7: Implementation**  
  * **Action**: Write Terraform scripts to spin up a managed K8s cluster (AKS on Azure).

## Week 5: Cloud Engineering (Azure Focus) & SRE
*Goal: Cloud-native infrastructure and keeping the system up.*

* **Day 1-2: Azure Fundamentals**  
  * **Resource**: `azure`.
  * **Action**: Azure AD (Entra ID), Virtual Networks, App Services, Azure Blob Storage, Azure SQL.
* **Day 3-4: SRE Principles**  
  * **Resource**: `architecture`.
  * **Action**: Service Level Indicators (SLI), Objectives (SLO), and Agreements (SLA). Error Budgets, Incident Response, Post-Mortems.
* **Day 5-6: Observability (Monitoring & Logging)**  
  * **Action**: Prometheus metrics scraping, Grafana dashboards, ELK/Fluentd log aggregation, Jaeger distributed tracing.
* **Day 7: Scenario Practice**  
  * **Action**: Document a blameless post-mortem for a mocked Database outage.

## Week 6: Scripting & Interview Prep
*Goal: Polish scripting and handle scenario questions.*

* **Day 1-2: Scripting (Python / Golang)**  
  * **Resource**: `Python`, `golang`.
  * **Action**: Automating boring tasks. E.g., script a backup of a Postgres DB to AWS S3 using Boto3 (Python) or AWS SDK (Go).
* **Day 3-5: Scenario & Architecture Questions**  
  * **Resource**: `Scenariobase_questions_bank`, `system_design`.
  * **Action**: "How would you design the deployment pipeline for a high-frequency trading firm?" "How do you zero-downtime migrate a database?"
* **Day 6-7: Mock Interviews**  
  * **Action**: 45-minute strict timed mock interviews focusing on K8s troubleshooting and CI/CD design.
