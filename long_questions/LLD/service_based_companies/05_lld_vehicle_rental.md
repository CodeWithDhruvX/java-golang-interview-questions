# Low-Level Design (LLD) - Vehicle Rental System

## Problem Statement
Design a Vehicle Rental System (like Zoomcar, Hertz). Users can search for cars, motorcycles, or vans, select a rental duration, and book the vehicle.

## Requirements
*   **Vehicles:** Standard properties like `licensePlate`, `make`, `model`, `type` (Car, Truck, Motorcycle).
*   **Store / Branch:** Vehicles are parked at specific physical locations.
*   **Search:** Users can search for available vehicles at a specific store for a specific date range.
*   **Booking:** Create a reservation with a start date, end date, and user details.
*   **Inventory Status:** A vehicle can be Available, Booked, In-Maintenance, or Rented.
*   **Payment & Late Fees:** Calculate base rate + optionally add insurance + possible late fee upon return.

## Core Entities / Classes

1.  **RentalSystem (Facade):** Main entry point for searching and booking.
2.  **Branch / StoreLocation:** Contains `Address` and `List<Vehicle>`.
3.  **Vehicle (Abstract):** Tracks `barcode`, `dailyRate`, `VehicleStatus`.
4.  **Reservation:** `id`, `creationDate`, `returnDate`, `pickupDate`, `dropoffLocation`, `ReservationStatus`.
5.  **RentalEquipment (Optional):** Child seats, GPS units that can be added to a reservation.
6.  **Bill / Invoice:** Handles the monetary calculation, including `BaseCharge`, addons, taxes.
7.  **Account:** User or Admin.

## Key Design Patterns Applicable
*   **Decorator Pattern:** The `Bill` can be decorated with additional items. Base Rental Cost + GPS Addon + Insurance Addon + Child Seat Addon. This avoids creating rigid subclasses.
*   **Factory Pattern:** For creating different types of vehicles.
*   **Observer Pattern:** To notify users via email when their rental period is about to expire.

## Code Snippet (Decorator Pattern for Pricing)

```java
// Component Interface
public interface RentalCost {
    double getCost();
}

// Concrete Component
public class BaseVehicleCost implements RentalCost {
    private double dailyRate;
    private int days;

    public BaseVehicleCost(double dailyRate, int days) {
        this.dailyRate = dailyRate;
        this.days = days;
    }

    @Override
    public double getCost() { return this.dailyRate * this.days; }
}

// Base Decorator
public abstract class CostDecorator implements RentalCost {
    protected RentalCost rentalCost;
    public CostDecorator(RentalCost req) { this.rentalCost = req; }
    public abstract double getCost();
}

// Concrete Decorator
public class InsuranceDecorator extends CostDecorator {
    public InsuranceDecorator(RentalCost req) { super(req); }

    @Override
    public double getCost() {
        return rentalCost.getCost() + 50.0; // Flat $50 insurance
    }
}
```

## Follow-up Questions for Candidate
1.  How do you ensure a car isn't double-booked for intersecting date ranges? (Explain database interval checking or maintaining a calendar ledger per vehicle).
2.  If a vehicle goes into unscheduled maintenance, how do you handle existing future reservations for that vehicle?
3.  How can the Search service be optimized for high read throughput using an in-memory cache?
