# Low-Level Design (LLD) - Ride Sharing System

## Problem Statement
Design a Ride Sharing System similar to Ola/Uber that connects riders with drivers for transportation services, optimized for the Indian market.

## Requirements
*   **Users:** Two main types - Rider and Driver
*   **Ride Operations:** Ride booking, driver matching, route calculation, fare calculation
*   **Real-time Tracking:** Live location tracking for riders and drivers
*   **Payment System:** Multiple payment options (Cash, UPI, Cards, Wallet)
*   **Ride Types:** Multiple ride categories (Auto, Go, Sedan, SUV, Premium)
*   **Driver Operations:** Accept/reject rides, availability management, earnings tracking
*   **Rating System:** Mutual rating between riders and drivers
*   **Surge Pricing:** Dynamic pricing based on demand and supply
*   **Emergency Features:** SOS button, trip sharing, safety features

## Core Entities / Classes

1.  **User (Abstract):** `id`, `name`, `email`, `phone`, `location`
    *   **Rider:** `paymentMethods`, `rideHistory`, `defaultPaymentMethod`, `homeAddress`, `workAddress`
    *   **Driver:** `vehicleDetails`, `licenseNumber`, `isAvailable`, `currentLocation`, `rating`, `earnings`

2.  **Vehicle:** `vehicleNumber`, `type`, `model`, `color`, `insuranceExpiry`, `pucExpiry`
3.  **RideType:** `name` (Auto, Go, Sedan), `baseFare`, `perKmRate`, `perMinuteRate`, `capacity`
4.  **Ride:** `rideId`, `rider`, `driver`, `vehicle`, `rideType`, `pickupLocation`, `dropLocation`, `status`, `fare`, `distance`, `duration`
5.  **Location:** `latitude`, `longitude`, `address`, `timestamp`
6.  **Payment:** `ride`, `amount`, `paymentMethod`, `status`, `transactionId`
7.  **Rating:** `ride`, `riderRating`, `driverRating`, `comments`, `timestamp`
8.  **SurgeCalculator:** Calculates dynamic pricing based on demand/supply
9.  **MatchingEngine:** Algorithm to match riders with nearby drivers
10. **RouteOptimizer:** Calculates optimal routes and ETA

## Key Design Patterns Applicable
*   **Strategy Pattern:** Different fare calculation strategies for ride types
*   **Observer Pattern:** Notify riders/drivers of ride status changes
*   **Factory Pattern:** Create different types of rides and vehicles
*   **Singleton Pattern:** `MatchingEngine` and `SurgeCalculator`
*   **State Pattern:** Ride status management (SEARCHING → MATCHED → ACCEPTED → ARRIVING → IN_PROGRESS → COMPLETED)
*   **Command Pattern:** Encapsulate ride operations for undo/redo
*   **Decorator Pattern:** Add surge pricing, discounts, or special features to base fare

## Code Snippet (Java/Go focus)

### Java Implementation
```java
// Ride State Management
public enum RideStatus {
    SEARCHING, MATCHED, ACCEPTED, ARRIVING, IN_PROGRESS, COMPLETED, CANCELLED
}

public class Ride {
    private String rideId;
    private Rider rider;
    private Driver driver;
    private Location pickupLocation;
    private Location dropLocation;
    private RideStatus status;
    private FareCalculator fareCalculator;
    
    public void updateStatus(RideStatus newStatus) {
        this.status = newStatus;
        notifyParticipants();
    }
    
    public double calculateFare() {
        return fareCalculator.calculateFare(this);
    }
}

// Strategy Pattern for Fare Calculation
public interface FareCalculator {
    double calculateFare(Ride ride);
}

public class StandardFareCalculator implements FareCalculator {
    private double baseFare;
    private double perKmRate;
    private double perMinuteRate;
    
    @Override
    public double calculateFare(Ride ride) {
        double distance = ride.getDistance();
        double duration = ride.getDuration();
        return baseFare + (distance * perKmRate) + (duration * perMinuteRate);
    }
}

public class SurgeFareCalculator implements FareCalculator {
    private FareCalculator baseCalculator;
    private SurgeCalculator surgeCalculator;
    
    @Override
    public double calculateFare(Ride ride) {
        double baseFare = baseCalculator.calculateFare(ride);
        double surgeMultiplier = surgeCalculator.getSurgeMultiplier(ride.getPickupLocation());
        return baseFare * surgeMultiplier;
    }
}
```

### Go Implementation
```go
// Ride Status Management
type RideStatus int

const (
    Searching RideStatus = iota
    Matched
    Accepted
    Arriving
    InProgress
    Completed
    Cancelled
)

type Ride struct {
    RideID         string
    Rider          *Rider
    Driver         *Driver
    PickupLocation *Location
    DropLocation   *Location
    Status         RideStatus
    FareCalculator FareCalculator
    Distance       float64
    Duration       time.Duration
}

func (r *Ride) UpdateStatus(status RideStatus) {
    r.Status = status
    r.notifyParticipants()
}

func (r *Ride) CalculateFare() float64 {
    return r.FareCalculator.CalculateFare(r)
}

// Strategy Pattern for Fare Calculation
type FareCalculator interface {
    CalculateFare(ride *Ride) float64
}

type StandardFareCalculator struct {
    BaseFare     float64
    PerKmRate    float64
    PerMinuteRate float64
}

func (sfc *StandardFareCalculator) CalculateFare(ride *Ride) float64 {
    return sfc.BaseFare + (ride.Distance * sfc.PerKmRate) + (float64(ride.Duration.Minutes()) * sfc.PerMinuteRate)
}

type SurgeFareCalculator struct {
    BaseCalculator   FareCalculator
    SurgeCalculator  *SurgeCalculator
}

func (sfc *SurgeFareCalculator) CalculateFare(ride *Ride) float64 {
    baseFare := sfc.BaseCalculator.CalculateFare(ride)
    multiplier := sfc.SurgeCalculator.GetSurgeMultiplier(ride.PickupLocation)
    return baseFare * multiplier
}
```

## Critical Design Considerations
*   **Real-time Location Updates:** WebSocket for live tracking
*   **Driver Matching Algorithm:** Optimize for pickup time, driver rating, and availability
*   **Concurrent Ride Processing:** Handle thousands of simultaneous rides
*   **Geospatial Indexing:** Efficient location-based queries
*   **Payment Processing:** Handle multiple payment methods and refunds
*   **Surge Pricing Algorithm:** Balance demand and supply effectively
*   **Safety Features:** Emergency contacts, trip sharing, driver verification

## Indian Market Specific Features
*   **Auto Rickshaws:** Three-wheeler vehicle support
*   **Regional Language Support:** Multi-language interface
*   **Cash Payments:** High cash transaction volume
*   **Local Address System:** Support for Indian address formats
*   **Traffic-aware Routing:** Consider Indian traffic conditions
*   **Regional Festivals:** Surge pricing during festivals and events

## Interview Success Tips
*   Focus on the driver matching algorithm and optimization
*   Discuss how to handle surge pricing fairly
*   Address real-time location tracking challenges
*   Explain database design for ride history and analytics
*   Discuss edge cases: ride cancellation, payment failures, driver unavailability
*   Talk about scalability for peak hours (office commute, nightlife)
*   Explain how to ensure driver and rider safety
