package main

import (
	"fmt"
	"sort"
)

// Pattern: Event-Based Sweeping
// Difficulty: Medium
// Key Concept: Decomposing intervals into discrete "events" (points in time) and processing them in order.

/*
INTUITION:
This is very similar to "Line Sweep" but often applied when you have specific capacity constraints or more complex event types.

Imagine a bus driving along a straight road (0 to 1000).
- trip 1: pick up 2 people at km 150, drop them at km 400.
- trip 2: pick up 5 people at km 300, drop them at km 700.

Problem: "Does the bus ever exceed its capacity of 4 seats?"

Method:
- At km 150: +2 passengers. (Current: 2)
- At km 300: +5 passengers. (Current: 7) -> OVERFLOW! (7 > 4)
- At km 400: -2 passengers. (Current: 5)
...

We just sort these "Events" and sweep. If `current_load` ever > `capacity`, return False.

PROBLEM:
LeetCode 1094. Car Pooling
There is a car with `capacity` empty seats. The vehicle travels east to west.
Given a list of `trips`, `trip[i] = [numPassengers, from, to]`, return true if it is possible to pick up and drop off all passengers for all the given trips, or false otherwise.

ALGORITHM:
1. Create an event list.
   - For each trip `[num, start, end]`:
     - Add `(start, +num)`
     - Add `(end, -num)`
2. Sort events by location.
   - If locations are same, process Drop-offs (-num) before Pick-ups (+num). (Usually depends on problem statement, here we assume passengers get off before new ones get on at the same stop).
3. Iterate through events, maintaining `currentLoad`.
4. If `currentLoad > capacity`, return `false`.
5. Return `true` if loop finishes.

Alternative (Bucket Sort approach): Since max location is small (1000), we could use an array `diff[1001]` and prefix sum.
However, the "Event Sorting" approach works even if locations are huge (e.g., 0 to 10^9), so we demonstrate the Sorting approach here as the generic pattern.
*/

type TripEvent struct {
	location int
	change   int // +passengers or -passengers
}

func carPooling(trips [][]int, capacity int) bool {
	events := make([]TripEvent, 0, len(trips)*2)

	for _, t := range trips {
		numPassengers := t[0]
		start := t[1]
		end := t[2]

		events = append(events, TripEvent{location: start, change: numPassengers})
		events = append(events, TripEvent{location: end, change: -numPassengers})
	}

	// Sort events
	sort.Slice(events, func(i, j int) bool {
		if events[i].location == events[j].location {
			// Optimization: Drop off passengers first to free up seats?
			// The problem states: "drop off ... at location 'to'".
			// Usually standard logic: process END before START if at same time.
			// So negative change (drop off) should come first.
			return events[i].change < events[j].change
		}
		return events[i].location < events[j].location
	})

	currentLoad := 0
	for _, e := range events {
		currentLoad += e.change
		if currentLoad > capacity {
			return false
		}
	}

	return true
}

func main() {
	// Trips: [[2,1,5], [3,3,7]], Capacity: 4
	// Loc 1: +2. Load 2.
	// Loc 3: +3. Load 5. (>4) -> False.
	trips1 := [][]int{{2, 1, 5}, {3, 3, 7}}
	cap1 := 4
	fmt.Printf("Trips: %v, Cap: %d -> Possible? %v\n", trips1, cap1, carPooling(trips1, cap1))

	// Trips: [[2,1,5], [3,3,7]], Capacity: 5
	// Loc 1: +2. Load 2.
	// Loc 3: +3. Load 5. (<=5) -> OK.
	// Loc 5: -2. Load 3.
	// Loc 7: -3. Load 0.
	cap2 := 5
	fmt.Printf("Trips: %v, Cap: %d -> Possible? %v\n", trips1, cap2, carPooling(trips1, cap2))
}
