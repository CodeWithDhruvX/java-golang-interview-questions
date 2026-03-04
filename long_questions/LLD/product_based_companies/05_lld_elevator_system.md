# Low-Level Design (LLD) - Elevator System

## Problem Statement
Design an Elevator System for a multi-story building. The system should handle requests efficiently to minimize wait time and power usage. This is a classic concurrency problem.

## Requirements
*   **Multiple Elevators:** The building has multiple elevators serving multiple floors.
*   **State Tracking:** Elevators can be moving up, moving down, or idle.
*   **Buttons:** Inside buttons (destination floors) and outside buttons (up/down requests per floor).
*   **Dispatching Algorithm:** An efficient algorithm to dispatch the correct elevator to a floor.
*   **Safety/Limits:** Capacity constraints (weight/number of people).

## Core Entities / Classes

1.  **ElevatorSystem / ElevatorController:** Singleton or centralized dispatcher. Receives requests from outside buttons.
2.  **ElevatorCar:** Properties: `id`, `currentFloor`, `direction` (UP, DOWN, IDLE), `state` (MOVING, STOPPED).
3.  **Request:** Source floor, Destination floor, Direction.
4.  **Button:** 
    *   `HallButton` (Outside: Up/Down)
    *   `ElevatorButton` (Inside: Floor number)
5.  **Display:** Inside the elevator and outside on each floor (shows current floor and direction).
6.  **Door:** `isOpen`, `open()`, `close()`.

## Key Design Patterns Applicable
*   **State Pattern:** An elevator changes state between moving, stopped, idle, door open, door closed.
*   **Strategy Pattern:** Dispatch algorithms. (e.g., Shortest Seek Time First (SSTF), SCAN algorithm, etc.)
*   **Observer Pattern:** Display boards observing the Elevator Car's state to update the floor numbers.

## Dispatch Algorithm Concept (SCAN / LOOK Algorithm)

A common interview approach is the LOOK algorithm, where the elevator continues traveling in its current direction while there are remaining requests in that direction. If there are no more requests in that direction, it changes direction.

```java
public class ElevatorCar implements Runnable {
    private int id;
    private int currentFloor;
    private Direction direction;
    private PriorityQueue<Integer> upRequests;
    private PriorityQueue<Integer> downRequests; // Custom comparator for max-heap

    public void processRequests() {
        while (true) {
            if (direction == Direction.UP) {
                if (!upRequests.isEmpty()) {
                    int nextFloor = upRequests.poll();
                    moveToFloor(nextFloor);
                } else {
                    direction = Direction.DOWN;
                }
            } else if (direction == Direction.DOWN) {
                // process downRequests
            } else {
                // IDLE
            }
        }
    }
}
```

## Follow-up Questions for Candidate
1.  How will you handle VIP requests or Fire Emergencies? (Priority Dispatching)
2.  Can you explain how the SCAN array/queues handle new requests that come in while the elevator is moving?
3.  Discuss thread safety—how do you guarantee multiple users pressing buttons concurrently don't corrupt the request queue?
