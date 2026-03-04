# Low-Level Design (LLD) - BookMyShow (Ticket Booking)

## Problem Statement
Design an online movie ticket booking system like BookMyShow or Fandango. It should allow users to browse movies, theaters, showtimes, and book seats.

## Requirements
*   **Search:** Users should be able to search movies by title, language, genre, city, and release date.
*   **Theater & Showtimes:** Users can view theaters running a movie and their showtimes.
*   **Seat Selection:** Users can select available seats. Support different seat types (Silver, Gold, Platinum).
*   **Concurrency:** To prevent double booking, the system should hold (reserve) a seat temporarily (e.g., 10 minutes) while the user pays.
*   **Payment & Notification:** The system must process payments and send an email/SMS confirmation.

## Core Entities / Classes

1.  **City / Location:** Properties like `name`, `zipcode`. Contains a list of `Cinema`.
2.  **Cinema (Theater):** Contains multiple `CinemaHall` (Screens).
3.  **CinemaHall (Screen):** Contains a list of `Seat`.
4.  **Movie & Show:** `Movie` has info like `title`, `duration`. `Show` maps a `Movie` to a `CinemaHall` at a specific time.
5.  **Seat:** `id`, `type`, `price`.
6.  **ShowSeat:** Maps a `Seat` to a `Show`. Has states: `AVAILABLE`, `RESERVED`, `BOOKED`.
7.  **Booking:** Maps a User to a list of `ShowSeat`s, tracks `totalAmount`, `status` (PENDING, CONFIRMED, FAILED).
8.  **Payment:** Payment details and status.

## Key Design Patterns Applicable
*   **Strategy Pattern:** For dynamic pricing logic (e.g., weekend surge pricing).
*   **Observer Pattern:** Booking service notifying Email Service and SMS Service upon successful ticket confirmation.
*   **Concurrency/Locking mechanisms:** (In Database) Optimistic or Pessimistic locking for `ShowSeat` booking.

## Code Snippet (Seat Locking Concept)

```java
public class BookingController {

    // Using synchronized block or a distributed lock mechanism (Redis)
    public synchronized boolean reserveSeats(List<ShowSeat> seats, User user) {
        for(ShowSeat seat : seats) {
            if(seat.getStatus() != SeatStatus.AVAILABLE) {
                return false; // Transaction fails if any seat is unavailable
            }
        }
        
        for(ShowSeat seat : seats) {
            seat.setStatus(SeatStatus.RESERVED);
            seat.setReservedBy(user);
            seat.setReservedAt(new Date());
        }
        return true;
    }
}
```

## Follow-up Questions for Candidate
1.  How do you implement the "Seat Hold" mechanism? Do you use a database job or a Cache like Redis?
2.  If the payment gateway takes 15 minutes to respond, but a seat hold expires in 10 minutes, what exactly happens?
3.  How do you prevent SQL injection and ensure ACID properties for booking transactions?
