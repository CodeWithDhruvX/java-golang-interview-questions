# Low-Level Design (LLD) - Payment Gateway System

## Problem Statement
Design a Payment Gateway System that processes online payments through multiple channels including UPI, credit/debit cards, net banking, and digital wallets. The system should handle Indian payment ecosystem requirements.

## Requirements
*   **Multiple Payment Methods:** UPI, Credit Cards, Debit Cards, Net Banking, Wallets (PayTM, PhonePe, Amazon Pay)
*   **Merchant Integration:** APIs for merchants to integrate payment processing
*   **Transaction Processing:** Handle payment authorization, capture, refund, and settlement
*   **Security:** PCI DSS compliance, fraud detection, tokenization
*   **Webhooks:** Real-time payment status notifications to merchants
*   **Dashboard:** Analytics, reporting, and transaction monitoring
*   **Recurring Payments:** Subscription and EMI support
*   **International Payments:** Multi-currency support

## Core Entities / Classes

1.  **PaymentGateway (Singleton):** Central orchestrator for all payment operations
2.  **Merchant:** `merchantId`, `businessDetails`, `apiKey`, `webhookUrl`, `settlementBankAccount`
3.  **Customer:** `customerId`, `email`, `phone`, `paymentMethods`, `defaultPaymentMethod`
4.  **PaymentMethod (Abstract):** `type`, `isSaved`, `token`
    *   **UPIMethod:** `upiId`, `virtualPaymentAddress`
    *   **CardMethod:** `cardNumber`, `expiry`, `cvv`, `cardType`, `bank`
    *   **NetBankingMethod:** `bankCode`, `accountNumber`
    *   **WalletMethod:** `walletProvider`, `walletId`
5.  **Transaction:** `transactionId`, `merchant`, `customer`, `amount`, `currency`, `status`, `paymentMethod`, `createdAt`
6.  **PaymentProcessor (Interface):** `processPayment(transaction)`, `refund(transaction, amount)`
    *   **UPIProcessor**, **CardProcessor**, **NetBankingProcessor**, **WalletProcessor**
7.  **FraudDetectionEngine:** Analyzes transactions for suspicious patterns
8.  **SettlementEngine:** Daily settlement processing to merchant accounts
9.  **NotificationService:** Webhook and SMS/email notifications
10. **RetryManager:** Handles failed payment retries

## Key Design Patterns Applicable
*   **Strategy Pattern:** Different payment processing strategies for each payment type
*   **Factory Pattern:** Create appropriate payment processors based on payment method
*   **Observer Pattern:** Notify merchants of payment status changes
*   **Chain of Responsibility:** Fraud detection → Risk assessment → Payment processing
*   **Command Pattern:** Encapsulate payment operations for undo/redo
*   **Singleton Pattern:** `PaymentGateway` and `SettlementEngine`
*   **Adapter Pattern:** Integrate with different bank APIs and payment providers

## Code Snippet (Java/Go focus)

### Java Implementation
```java
// Strategy Pattern for Payment Processing
public interface PaymentProcessor {
    PaymentResult processPayment(Transaction transaction);
    PaymentResult refund(Transaction transaction, double amount);
}

public class UPIProcessor implements PaymentProcessor {
    private UPIClient upiClient;
    
    @Override
    public PaymentResult processPayment(Transaction transaction) {
        UPIPaymentRequest request = new UPIPaymentRequest(
            transaction.getCustomer().getUpiId(),
            transaction.getAmount(),
            transaction.getMerchant().getUpiMerchantId()
        );
        return upiClient.processPayment(request);
    }
}

// Transaction Status State Pattern
public enum TransactionStatus {
    INITIATED, PROCESSING, SUCCESS, FAILED, REFUNDED, PARTIALLY_REFUNDED
}

public class Transaction {
    private String transactionId;
    private Merchant merchant;
    private Customer customer;
    private double amount;
    private TransactionStatus status;
    private PaymentMethod paymentMethod;
    
    public void updateStatus(TransactionStatus newStatus) {
        this.status = newStatus;
        notifyMerchant();
    }
}
```

### Go Implementation
```go
// Payment Processing Strategy
type PaymentProcessor interface {
    ProcessPayment(transaction *Transaction) *PaymentResult
    Refund(transaction *Transaction, amount float64) *PaymentResult
}

type UPIProcessor struct {
    upiClient UPIApiClient
}

func (upi *UPIProcessor) ProcessPayment(transaction *Transaction) *PaymentResult {
    request := &UPIPaymentRequest{
        UPIID:       transaction.Customer.UpiId,
        Amount:      transaction.Amount,
        MerchantID:  transaction.Merchant.UpiMerchantId,
    }
    return upi.upiClient.ProcessPayment(request)
}

// Transaction Management
type TransactionStatus int

const (
    Initiated TransactionStatus = iota
    Processing
    Success
    Failed
    Refunded
    PartiallyRefunded
)

type Transaction struct {
    TransactionID  string
    Merchant       *Merchant
    Customer       *Customer
    Amount         float64
    Status         TransactionStatus
    PaymentMethod  PaymentMethod
    CreatedAt      time.Time
}
```

## Critical Design Considerations
*   **Idempotency:** Ensure duplicate requests don't cause multiple charges
*   **Transaction Atomicity:** ACID compliance for financial operations
*   **Rate Limiting:** Prevent abuse and manage system load
*   **Data Encryption:** Sensitive data protection at rest and in transit
*   **Audit Trail:** Complete transaction history for compliance
*   **High Availability:** 99.99% uptime for payment processing
*   **Latency:** Sub-second response times for payment authorization

## UPI-Specific Considerations
*   **Virtual Payment Address (VPA):** Handle UPI ID validation
*   **2FA Authentication:** Integrate with bank authentication flows
*   **Transaction Limits:** Handle per-transaction and daily limits
*   **Bank Integration:** Connect with multiple UPI providers
*   **Settlement Cycle:** T+1 or T+2 settlement to merchant accounts

## PIT (Payment Interface Token) Security Questions

### Q: What is the primary token time used for stateless security in PIT?

**Answer:**
The primary token time in PIT (Payment Interface Token) refers to the **JWT expiration time (`exp` claim)** which is typically set to **15-30 minutes** for stateless security in payment systems.

**Detailed Explanation:**

**1. Primary Token Time Components:**
- **Access Token TTL:** 15-30 minutes (primary token time)
- **Refresh Token TTL:** 7-30 days 
- **Token Issuance Time (`iat`):** When token was created
- **Not Before Time (`nbf`):** Token becomes valid after this time

**2. Why 15-30 Minutes for Primary Token?**
- **Security:** Limits damage if token is compromised
- **Performance:** Balances security with user experience
- **Scalability:** Reduces token revocation overhead
- **Compliance:** Meets PCI DSS requirements for session timeout

**3. Stateless Security Implementation:**
```java
// PIT Token Generation with Primary Time
public class PITTokenGenerator {
    private static final long PRIMARY_TOKEN_TTL = 15 * 60 * 1000; // 15 minutes
    
    public String generatePITToken(User user, PaymentContext context) {
        long now = System.currentTimeMillis();
        long expiryTime = now + PRIMARY_TOKEN_TTL;
        
        return Jwts.builder()
            .setSubject(user.getUserId())
            .claim("paymentContext", context)
            .claim("merchantId", context.getMerchantId())
            .claim("transactionLimit", context.getTransactionLimit())
            .setIssuedAt(new Date(now))
            .setExpiration(new Date(expiryTime)) // Primary token time
            .signWith(SignatureAlgorithm.HS256, getSigningKey())
            .compact();
    }
}
```

**4. Token Renewal Strategy:**
```go
// Go Implementation for PIT Token Renewal
type PITTokenManager struct {
    primaryTokenTTL time.Duration // 15 minutes
    refreshTTL      time.Duration // 7 days
}

func (ptm *PITTokenManager) RenewToken(refreshToken string) (*TokenPair, error) {
    // Validate refresh token
    claims, err := ptm.validateRefreshToken(refreshToken)
    if err != nil {
        return nil, err
    }
    
    // Generate new primary token with fresh expiration
    newAccessToken := ptm.generatePrimaryToken(claims.UserID)
    
    return &TokenPair{
        AccessToken:  newAccessToken,  // New 15-min token
        RefreshToken: refreshToken,    // Same refresh token
        ExpiresIn:    int(ptm.primaryTokenTTL.Seconds()),
    }, nil
}
```

**5. Security Benefits of Short Primary Token Time:**
- **Reduced Attack Window:** Limited time for token exploitation
- **Automatic Session Timeout:** Enforces session limits
- **Token Rotation:** Frequent renewal reduces token theft impact
- **Compliance:** Meets regulatory requirements for session management

**6. Trade-offs and Considerations:**
- **User Experience:** Balance between security and convenience
- **System Load:** More frequent token renewal requests
- **Network Reliability:** Handle token expiry during network issues
- **Offline Support:** Graceful degradation for poor connectivity

**7. PIT-Specific Security Features:**
- **Transaction Context Binding:** Token tied to specific payment parameters
- **Merchant Validation:** Token restricted to specific merchant
- **Amount Limits:** Token validity based on transaction amounts
- **Geo-Fencing:** Location-based token validation

**Interview Follow-up Points:**
- How do you handle token expiry during payment processing?
- What's your strategy for token refresh in mobile apps?
- How do you implement token revocation for fraud detection?
- What happens when primary token expires mid-transaction?

## Interview Success Tips
*   Discuss how to handle payment failures and retries
*   Explain idempotency in payment systems
*   Address security concerns and PCI compliance
*   Design for high availability and disaster recovery
*   Discuss how to handle concurrent payment processing
*   Explain webhook reliability and retry mechanisms
*   Talk about fraud detection algorithms and risk scoring
*   **NEW:** Explain PIT token security and primary token time strategy
