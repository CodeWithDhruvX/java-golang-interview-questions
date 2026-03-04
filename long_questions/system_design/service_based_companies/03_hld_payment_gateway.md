# System Design (HLD) - Payment Gateway Integration

## Problem Statement
Design a Payment Gateway integration layer for an enterprise. The system must act as an abstraction over multiple external Payment Providers (like Stripe, PayPal, Razorpay) and provide a unified API to internal services (E-commerce, Billing).

## 1. Requirements Clarification
### functional Requirements
*   **Unified API:** Internal services request a charge, refund, or subscription independently of the underlying provider.
*   **Smart Routing:** The system should dynamically route the payment to the cheapest or most reliable provider based on card type or region.
*   **Webhooks:** Handle asynchronous callbacks from providers (e.g., "Payment Success" or "Payment Failed").

### Non-Functional Requirements
*   **100% Reliability & Idempotency:** A customer must *never* be charged twice for the same transaction.
*   **Security & Compliance:** Must adhere to PCI-DSS standards (No raw credit card storage).
*   **Auditability:** Every step of money movement must be logged.

## 2. High-Level Architecture

```text
[ Internal Svc (e.g., E-commerce) ]
           |
           v (POST /v1/charge)
[ Payment Service Gateway ] <------> [ Database (PostgreSQL) ]
           |
       (Smart Router)
      /      |       \
     /       |        \
[Adyen]  [Stripe]  [PayPal]
```

## 3. The Core Challenge: Idempotency
Network issues are the biggest enemy of payment systems.
*   **Scenario:** Application calls Stripe API -> Stripe charges the card -> Network cuts out so Application never gets the HTTP 200 response -> Application thinks it failed and retries -> Customer charged twice.
*   **Solution (Idempotency Key):**
    1. The Internal Service generates a unique UUID (Idempotency Key) for the order.
    2. Payment Gateway saves this in the DB: `OrderID: 123, IdemKey: ABC, Status: PENDING`.
    3. Payment Gateway calls Stripe, passing `Idempotency-Key: ABC` in the HTTP headers.
    4. If it retries, Stripe sees the same Key, recognizes it as a duplicate, and returns the previous result without charging the card again.

## 4. Database Schema (Ledger System)
A rigid relational database (ACID compliant) is absolutely required. **Never use a NoSQL database for core payment transactions.**

**Table: Transactions**
*   `transaction_id` (PK)
*   `order_id` (FK)
*   `amount` & `currency`
*   `provider` (Stripe, PayPal)
*   `provider_ref_id` (The ID returned by Stripe)
*   `status` (INITIATED, PENDING, COMPLETED, FAILED, REFUNDED)
*   `idempotency_key` (Unique Constraint)

## 5. Webhook Handling (Asynchronous Updates)
Often, a payment takes time (e.g., Bank Transfers, UPI). Stripe returns "Pending", and 5 minutes later sends a webhook.
1. Stripe POSTs to `/webhooks/stripe`.
2. The Webhook endpoint immediately saves the raw payload to an `incoming_events` table and responds `HTTP 200 OK` to Stripe so Stripe stops retrying.
3. A background worker picks up the event from the DB, verifies the signature (to prevent fraud), updates the `Transactions` table to `COMPLETED`, and publishes a `PaymentSuccessful` event to Kafka.
4. Internal E-commerce service listens to Kafka and ships the product.

## 6. Reconciliation
At midnight every day, a **Cron Job / Reconciliation script** runs.
*   It downloads the settlement report from Stripe.
*   It compares the Stripe report against the internal `Transactions` table row by row.
*   If the internal DB says `FAILED` but Stripe says `COMPLETED`, human intervention is flagged immediately to investigate the anomaly.

## 7. Follow-up Questions for Candidate
1.  How do you handle PCI-DSS compliance? (Using Provider-hosted iframe injected fields or Tokenization; the Internal Service never touches the raw 16-digit card number, it only handles an opaque Token).
2.  How would you design the system to handle a scenario where Stripe goes down globally? (Implement Circuit Breakers. If Stripe failures trip the circuit, the Smart Router automatically fails-over and routes all new traffic to Adyen or PayPal).
