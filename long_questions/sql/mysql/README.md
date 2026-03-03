# 📚 MySQL Interview Questions

> A complete, structured set of MySQL interview questions organized by company type.
> Mirrors the folder structure of Go and Java interview questions in this repo.

---

## 🗂️ Folder Structure

```
mysql/
├── product_based_companies/              ← Hard topics — Amazon, Google, Flipkart, Paytm
│   ├── 01_mysql_indexes_performance.md
│   ├── 02_mysql_transactions_acid.md
│   ├── 03_mysql_advanced_queries.md
│   ├── 04_mysql_replication_architecture.md
│   ├── README.md
│   └── theory/
│       ├── 01_mysql_indexes_performance_theory.md
│       ├── 02_mysql_transactions_acid_theory.md
│       ├── 03_mysql_advanced_queries_theory.md
│       └── 04_mysql_replication_architecture_theory.md
│
├── service_based_companies/              ← Core topics — TCS, Infosys, Wipro, Cognizant
│   ├── 01_mysql_basics.md
│   ├── 02_mysql_joins_subqueries.md
│   ├── 03_mysql_stored_procedures_triggers.md
│   ├── 04_mysql_security_administration.md
│   ├── README.md
│   └── theory/
│       ├── 01_mysql_basics_theory.md
│       ├── 02_mysql_joins_subqueries_theory.md
│       ├── 03_mysql_stored_procedures_triggers_theory.md
│       └── 04_mysql_security_administration_theory.md
│
├── theory/                               ← Original 250 Q theory files (legacy)
│   ├── 01_Basic_MySQL_(Q1-25).md
│   ├── 02_Intermediate_MySQL_(Q26-60).md
│   ├── 03_Advanced_MySQL_(Q61-100).md
│   ├── 04_Additional_Beginner_(Q101-120).md
│   ├── 05_Additional_Intermediate_(Q121-160).md
│   ├── 06_Senior_Level_(Q161-210).md
│   └── 07_Architect_Expert_(Q211-250).md
│
├── ecommerce_sql_interview_questions.md
├── sql_joins_100_questions.md
└── sql_zero_to_hero_interview_questions.md
```

---

## 📖 How to Use

| File | Use For |
|------|---------|
| **`service_based_companies/*.md`** | SQL coding practice — basics to intermediate |
| **`product_based_companies/*.md`** | Advanced SQL — performance, ACID, architecture |
| **`*/theory/*.md`** | Verbal prep — read aloud as spoken interview answers |

---

## 🗺️ Quick Reference

| Topic | Company Type | File |
|-------|-------------|------|
| DDL, DML, Keys, Normalization | 🔵 Service | `service_based_companies/01_mysql_basics.md` |
| JOINs, CTEs, Subqueries | 🔵 Service | `service_based_companies/02_mysql_joins_subqueries.md` |
| Stored Procedures & Triggers | 🔵 Service | `service_based_companies/03_mysql_stored_procedures_triggers.md` |
| Security, Backup, Monitoring | 🔵 Service | `service_based_companies/04_mysql_security_administration.md` |
| Indexes, EXPLAIN, Performance | 🔴 Product | `product_based_companies/01_mysql_indexes_performance.md` |
| ACID, Isolation, Deadlocks | 🔴 Product | `product_based_companies/02_mysql_transactions_acid.md` |
| Window Functions, JSON, Pivot | 🔴 Product | `product_based_companies/03_mysql_advanced_queries.md` |
| Replication, Sharding, HA | 🔴 Product | `product_based_companies/04_mysql_replication_architecture.md` |

