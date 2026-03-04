# Low-Level Design (LLD) - Parking Lot System

## Problem Statement
Design a fully functional Parking Lot system. A parking lot or car park is a dedicated cleared area that is intended for parking vehicles. Give a low-level design for a parking lot.

## Requirements
*   **Multiple Floors:** The parking lot should have multiple levels/floors.
*   **Multiple Entry/Exit Points:** There should be multiple entry and exit gates.
*   **Ticket Generation:** Users can collect a ticket at the entry point.
*   **Spot Types:** The parking lot can accommodate different types of vehicles (Motorcycle, Car, Truck). Each vehicle type requires a specific parking spot size (Compact, Large, Handicapped).
*   **Payment System:** Users can pay at the exit gate or using an automated payment machine. Payment calculation is based on an hourly rate.
*   **Availability:** Display boards showing available parking spots per floor for each vehicle type.

## Core Entities / Classes

1.  **Vehicle (Abstract):** Car, Truck, Motorcycle. Properties: `licensePlate`, `vehicleType`.
2.  **ParkingSpot:** `id`, `isFree`, `Vehicle`, `SpotType` (COMPACT, LARGE, MOTORBIKE, HANDICAPPED).
3.  **ParkingLevel/Floor:** Contains a list of `ParkingSpot`. Methods to find free spots and update displays.
4.  **ParkingLot:** Singleton class that manages all levels, entries, and exits.
5.  **ParkingTicket:** Contains `ticketNumber`, `issuedAt`, `paidAt`, `Vehicle`, `payStatus` (ACTIVE, PAID, LOST).
6.  **EntryPanel / ExitPanel:** Handles ticket issuance and collection/payment processing.
7.  **Payment Strategy (Interface):** Defines payment modes (CreditCard, Cash, UPI).

## Key Design Patterns Applicable
*   **Singleton Pattern:** `ParkingLot` instance to ensure only one system manages the overall state.
*   **Factory Pattern:** To create different types of `Vehicle` objects or `Payment` modes.
*   **Strategy Pattern:** For calculating fees based on different strategies (e.g., hourly, daily, weekend rate).
*   **Observer Pattern:** When a spot becomes free or occupied, the `ParkingFloor` notifies the `DisplayBoard` to update the count.

## Code Snippet (Java/Go focus)

```java
public abstract class Vehicle {
    private String licenseNumber;
    private VehicleType type;
    // getters and setters
}

public class ParkingSpot {
    private int id;
    private boolean isFree;
    private Vehicle vehicle;
    private ParkingSpotType type;

    public boolean assignVehicle(Vehicle vehicle) {
        if(isFree && this.type == vehicle.getType()) { // Simplify match logic
            this.vehicle = vehicle;
            isFree = false;
            return true;
        }
        return false;
    }
    public boolean removeVehicle() {
        this.vehicle = null;
        isFree = true;
        return true;
    }
}
```

## Follow-up Questions for Candidate
1.  How do you handle concurrent requests at multiple entry gates?
2.  What happens if the system crashes? How do you recover the state?
3.  How would you design the database schema for this?
