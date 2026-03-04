# Low-Level Design (LLD) - Traffic Light Control System

## Problem Statement
Design a software controller for a Traffic Light system at a 4-way intersection. The lights must transition smoothly and safely to control traffic flow.

## Requirements
*   **Intersection:** A standard 4-way intersection (North, South, East, West).
*   **States:** Each light can be RED, YELLOW, or GREEN.
*   **Timing:** Default timers for Green (e.g., 30s), Yellow (e.g., 5s), Red (Wait).
*   **Safety Constraints:** Opposing traffic streams must never have GREEN simultaneously. When North-South is Green/Yellow, East-West must be Red.
*   **Emergency Mode:** A manual override to set all lights to RED or a specific direction to GREEN.

## Core Entities / Classes

1.  **TrafficController:** Singleton class that manages the overall intersection timing and `TrafficLight` instances.
2.  **TrafficLight:** Represents the physical light pole. `id`, `direction` (N, S, E, W), `CurrentColor` (State).
3.  **LightState (Enum / State Pattern):**
    *   `RedState`
    *   `YellowState`
    *   `GreenState`
4.  **Direction (Enum):** NORTH_SOUTH, EAST_WEST.

## Key Design Patterns Applicable
*   **State Pattern:** The TrafficLight transitions logically from Green -> Yellow -> Red -> Green. Each state handles its own timer and transition logic.
*   **Observer Pattern:** The `TrafficController` broadcasts global events (like an Ambulance emergency override) to all `TrafficLight` registered observers to force a state change.

## Code Snippet (State Pattern for Traffic Light)

```java
public enum TrafficLightColor {
    RED, YELLOW, GREEN;
}

public class TrafficLight {
    private TrafficLightColor color;
    private int timerSeconds;

    // Changes to the next color in the sequence
    public void changeState() {
        if (color == TrafficLightColor.GREEN) {
            color = TrafficLightColor.YELLOW;
            timerSeconds = 5;
        } else if (color == TrafficLightColor.YELLOW) {
            color = TrafficLightColor.RED;
            timerSeconds = 35; // Waits for the other side to finish
        } else if (color == TrafficLightColor.RED) {
            color = TrafficLightColor.GREEN;
            timerSeconds = 30; // Green for 30 seconds
        }
    }
}
```

## Follow-up Questions for Candidate
1.  How will you implement the actual timers? (Discuss `ScheduledExecutorService` vs a daemon Thread).
2.  How would you design the system to handle a pedestrian crossing button?
3.  If a sensor detects no cars on East-West, how can the controller skip their Green phase to optimize traffic flow?
