# Low-Level Design (LLD) - E-commerce Shopping Cart

## Problem Statement
Design an E-commerce Shopping Cart system. The system should allow users to browse products, add items to a cart, apply discount codes or coupons, and checkout.

## Requirements
*   **Catalogue:** Users can search and view products.
*   **Cart Management:** Add, remove, update quantities of items in the cart.
*   **Pricing:** Calculate total price, tax, and shipping costs.
*   **Discounts:** Define and apply different types of discounts (Flat amount, Percentage, BOGO - Buy One Get One).
*   **Order processing:** Convert the cart into an Order upon successful checkout.

## Core Entities / Classes

1.  **Product:** `id`, `name`, `description`, `price`, `availableQuantity`.
2.  **Item / CartItem:** Wraps a `Product` with a `quantity` selected by the user.
3.  **ShoppingCart:** Contains a List of `CartItem`s. Has methods `addItem()`, `removeItem()`, `updateQuantity()`, `checkout()`.
4.  **Order:** Created from a `ShoppingCart`. Holds `OrderDetails`, `totalAmount`, `shippingAddress`, `status`.
5.  **Discount (Strategy Interface):** Defines calculation for reducing price.
    *   `PercentageDiscount`
    *   `FixedAmountDiscount`
    *   `ThresholdDiscount` (e.g., $10 off if cart > $100)
6.  **PaymentGateway:** Interface to handle the actual monetary transaction.

## Key Design Patterns Applicable
*   **Strategy Pattern:** Absolutely necessary for the Pricing and Discount logic. The rules for calculating price and applying coupons change frequently in e-commerce.
*   **Observer Pattern:** When an Order is placed (status changes to CONFIRMED), notify the Inventory Service to deduct stock, and notify the Email Service to send a receipt.
*   **Decorator Pattern (Optional):** Can be used to layer multiple discounts (e.g., 10% off + Free Shipping).

## Code Snippet (Strategy Pattern for Discounts)

```java
public interface DiscountStrategy {
    double calculateDiscount(ShoppingCart cart);
}

public class PercentageDiscount implements DiscountStrategy {
    private double percentage;
    public PercentageDiscount(double percentage) { this.percentage = percentage; }

    @Override
    public double calculateDiscount(ShoppingCart cart) {
        return cart.getSubTotal() * (percentage / 100);
    }
}

public class ShoppingCart {
    private List<CartItem> items;
    private DiscountStrategy discountStrategy;

    public void setDiscountStrategy(DiscountStrategy strategy) {
        this.discountStrategy = strategy;
    }

    public double getFinalTotal() {
        double subtotal = getSubTotal();
        double discount = discountStrategy != null ? discountStrategy.calculateDiscount(this) : 0;
        return subtotal - discount;
    }
}
```

## Follow-up Questions for Candidate
1.  How do you handle "Cart Abandonment"? Do you store cart data in a Cache (Redis) or the Database?
2.  If a user adds an item but doesn't checkout for 3 hours, how do you handle the situation where the product price changed in the meantime?
3.  How would you implement the observer pattern to notify 5 different microservices when an order is successful?
