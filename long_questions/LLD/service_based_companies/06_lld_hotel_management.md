# Low-Level Design (LLD) - Hotel Management System

## Problem Statement
Design a Hotel Management System to allow users to search for rooms, book them, and allow hotel staff to manage check-ins, check-outs, and housekeeping.

## Requirements
*   **Hotels & Branches:** Manage multiple hotel locations.
*   **Search:** Users can search for rooms by location, date ranges, and room style (Suite, Standard, Deluxe).
*   **Booking:** Users can reserve a room, preventing double bookings. 
*   **Operations:** Receptionists can Check-in and Check-out guests, process payments. Housekeepers can update room status (e.g., "Cleaning", "Ready").
*   **Amenities & Services:** Guests can order room service which is added to their final bill.

## Core Entities / Classes

1.  **HotelSystem:** Main controller.
2.  **HotelLocation (Branch):** Contains the physical location details and a list of `Room`s.
3.  **Room:** `roomNumber`, `style` (Enum), `bookingPrice`, `status` (AVAILABLE, OCCUPIED, RESERVED, MAINTENANCE).
    *   Can have a list of `RoomKey`s linked to it.
4.  **Account:** Base class.
    *   `Guest`: Books rooms.
    *   `Receptionist`: Can issue RoomKeys.
    *   `Housekeeper`: Can change `room` status to AVAILABLE after cleaning.
5.  **RoomBooking / Reservation:** `startDate`, `endDate`, `status` (PENDING, CONFIRMED, CANCELLED), `Guest`.
6.  **Invoice (Bill):** Base charge + `List<RoomServiceCharge>` + Taxes.
7.  **Search (Interface):** Defines search methods via Catalogue.

## Key Design Patterns Applicable
*   **Decorator Pattern:** Perfect for `Invoice` calculations (Base charge + Spa + Food + Parking).
*   **Observer Pattern:** When a guest Checks-out, the system notifies Housekeeping that the room needs cleaning. When Housekeeping finishes, the Room state is observed by the booking engine to show it as available for the next guest.
*   **State Pattern:** Room goes through distinct states: `Available` -> `Reserved` -> `Occupied` -> `Cleaning` -> `Available`.

## Code Snippet (Room State Update)

```java
public class Housekeeper extends Account {
    public boolean updateRoomStatus(Room room, RoomStatus status) {
        if(room.getStatus() == RoomStatus.OCCUPIED && status == RoomStatus.AVAILABLE) {
            // Business Rule: Cannot suddenly make an occupied room available, must go through Checkout
            return false; 
        }
        room.setStatus(status);
        if(status == RoomStatus.AVAILABLE) {
             // System notification that it can be assigned to new check-ins
             SystemNotifier.getInstance().notifyFrontDesk(room);
        }
        return true;
    }
}
```

## Follow-up Questions for Candidate
1.  How do you manage Room Inventory over continuous days? If you search for June 1 to June 5, how do you know Room 101 isn't booked for June 3?
2.  What happens if the guest accidentally loses their physical RFID key card? How do you deactivate it in your system?
3.  How do you implement dynamic pricing (where weekends are 30% more expensive)?
