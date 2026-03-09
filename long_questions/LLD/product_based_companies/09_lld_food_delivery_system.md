# Low-Level Design (LLD) - Food Delivery System

## Problem Statement
Design a Food Delivery System similar to Swiggy/Zomato that connects customers with restaurants and delivery partners for food ordering and delivery.

## Requirements
*   **Users:** Three main types - Customer, Restaurant, Delivery Partner
*   **Restaurant Operations:** Menu management, order acceptance, order preparation status
*   **Customer Operations:** Browse restaurants, search food, place orders, track delivery, payment
*   **Delivery Operations:** Order assignment, route optimization, delivery status updates
*   **Order Management:** Order lifecycle from placement to delivery completion
*   **Payment System:** Multiple payment modes (COD, UPI, Cards, Wallet)
*   **Real-time Tracking:** Live order status and delivery partner location
*   **Ratings & Reviews:** Customer feedback for restaurants and delivery partners

## Core Entities / Classes

1.  **User (Abstract):** `id`, `name`, `email`, `phone`, `address`, `status`
    *   **Customer:** `orderHistory`, `favorites`, `paymentMethods`
    *   **Restaurant:** `name`, `cuisine`, `menu`, `operatingHours`, `address`, `rating`
    *   **DeliveryPartner:** `vehicleDetails`, `currentLocation`, `isAvailable`, `earnings`

2.  **MenuItem:** `id`, `name`, `description`, `price`, `category`, `isAvailable`, `preparationTime`
3.  **Menu:** `restaurantId`, `menuItems`, `categories`
4.  **Order:** `orderId`, `customer`, `restaurant`, `items`, `totalAmount`, `status`, `paymentStatus`, `deliveryAddress`
5.  **OrderItem:** `menuItem`, `quantity`, `price`, `customizations`
6.  **Delivery:** `order`, `deliveryPartner`, `pickupTime`, `deliveryTime`, `route`, `status`
7.  **Payment (Interface):** `processPayment(amount)`, `refund(amount)`
    *   **UPIPayment**, **CardPayment**, **CashOnDelivery**, **WalletPayment**
8.  **NotificationService:** Handles order updates, delivery tracking, promotional messages
9.  **LocationTracker:** Real-time location tracking for delivery partners

## Key Design Patterns Applicable
*   **Strategy Pattern:** For different payment methods and delivery algorithms
*   **Observer Pattern:** For order status notifications (Customer, Restaurant, Delivery Partner)
*   **Factory Pattern:** For creating different types of users and payments
*   **Singleton Pattern:** For `NotificationService` and `LocationTracker`
*   **State Pattern:** For order status management (PLACED → CONFIRMED → PREPARING → READY → PICKED_UP → DELIVERED)
*   **Decorator Pattern:** For adding extra toppings or customizations to menu items

## Code Snippet (Java/Go focus)

### Java Implementation
```java
// Order State Pattern
public enum OrderStatus {
    PLACED, CONFIRMED, PREPARING, READY, PICKED_UP, DELIVERED, CANCELLED
}

public class Order {
    private String orderId;
    private Customer customer;
    private Restaurant restaurant;
    private List<OrderItem> items;
    private OrderStatus status;
    private Payment payment;
    
    public void updateStatus(OrderStatus newStatus) {
        this.status = newStatus;
        notifyObservers();
    }
}

// Strategy Pattern for Payment
public interface PaymentStrategy {
    boolean processPayment(double amount);
    boolean refund(double amount);
}

public class UPIPayment implements PaymentStrategy {
    private String upiId;
    
    @Override
    public boolean processPayment(double amount) {
        // UPI payment logic
        return true;
    }
}
```

### Go Implementation
```go
// Order State Management
type OrderStatus int

const (
    Placed OrderStatus = iota
    Confirmed
    Preparing
    Ready
    PickedUp
    Delivered
    Cancelled
)

type Order struct {
    OrderID     string
    Customer    *Customer
    Restaurant  *Restaurant
    Items       []OrderItem
    Status      OrderStatus
    Payment     PaymentStrategy
}

// Strategy Pattern for Payment
type PaymentStrategy interface {
    ProcessPayment(amount float64) bool
    Refund(amount float64) bool
}

type UPIPayment struct {
    UPIID string
}

func (upi *UPIPayment) ProcessPayment(amount float64) bool {
    // UPI payment logic
    return true
}
```

## Critical Design Considerations
*   **Concurrent Order Processing:** Multiple orders from same restaurant
*   **Real-time Updates:** WebSocket for live tracking
*   **Database Design:** Normalized vs denormalized for performance
*   **API Rate Limiting:** Prevent abuse during peak hours
*   **Caching Strategy:** Restaurant menus, popular items
*   **Scalability:** Handle 10K+ concurrent orders during peak hours

## Interview Success Tips
*   Focus on order lifecycle and state transitions
*   Discuss how to handle delivery partner assignment algorithm
*   Address edge cases: order cancellation, payment failures, delivery delays
*   Explain database schema for order tracking
*   Discuss how to ensure data consistency across distributed services
