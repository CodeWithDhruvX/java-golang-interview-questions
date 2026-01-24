package main

import (
	"fmt"
	"sort"
)

// Pattern: Line Sweep
// Difficulty: Medium/Hard
// Key Concept: Process events (start/end points) in sorted order to track active intervals or overlaps.

/*
INTUITION:
Imagine you are the manager of a conference center. You have a list of meeting times: [1pm-3pm], [2pm-4pm], [3:30pm-5pm].
You want to know: "What is the maximum number of rooms I need at the same time?"

Instead of checking every minute of the day (which is slow), you just look at the specific times when something CHANGES.
- 1:00pm: Meeting Starts (+1 room needed)
- 2:00pm: Meeting Starts (+1 room needed) -> Total 2
- 3:00pm: Meeting Ends (-1 room needed) -> Total 1
- 3:30pm: Meeting Starts (+1) -> Total 2
- ...

By sweeping a vertical line across the timeline (sorted events), you can track the state (number of active concurrent meetings) efficiently.

PROBLEM:
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...], find the minimum number of conference rooms required. (LeetCode: Meeting Rooms II / Car Pooling concept)

ALGORITHM:
1. Separate Start times and End times.
2. Sort Start times. Sort End times.
3. Use two pointers: `s_ptr` (for starts) and `e_ptr` (for ends).
4. Iterate through Start times:
   - If `start[s_ptr] < end[e_ptr]`:
     - A meeting has started before the earliest one ended.
     - Increment `count`. Move `s_ptr`.
   - Else (`start[s_ptr] >= end[e_ptr]`):
     - A meeting ended by the time this one started. One room frees up.
     - Move both `s_ptr` and `e_ptr` (effectively reusing the room).
   - Update `max_count`.

NOTE: This specific "Two Pointer" approach is an optimized version of Line Sweep for this specific problem.
The more *general* Line Sweep approach (explicit Event struct) is:
1. Create events: (time, +1) for start, (time, -1) for end.
2. Sort events by time.
3. Iterate and sum values.

We will implement the GENERAL Event-Based approach here as it's more extensible to other Line Sweep problems (like Skyline).
*/

type Event struct {
	time int
	kind int // +1 for start, -1 for end
}

func minMeetingRooms(intervals [][]int) int {
	events := make([]Event, 0, len(intervals)*2)

	for _, interval := range intervals {
		// Start of a meeting: need a room (+1)
		events = append(events, Event{time: interval[0], kind: 1})
		// End of a meeting: free a room (-1)
		// Note: If a meeting ends at 3 and another starts at 3, usually reuse is allowed.
		// To process "end" before "start" at same time, sort End events before Start events if needed.
		// For Meeting Rooms, usually [1, 3] and [3, 4] DOES NOT overlap.
		// So we should process END (-1) before START (+1) if times are equal.
		events = append(events, Event{time: interval[1], kind: -1})
	}

	// Sort Events
	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			// If times are equal, process END (-1) before START (1) to reuse room
			return events[i].kind < events[j].kind
		}
		return events[i].time < events[j].time
	})

	maxRooms := 0
	currentRooms := 0

	for _, e := range events {
		currentRooms += e.kind
		if currentRooms > maxRooms {
			maxRooms = currentRooms
		}
	}

	return maxRooms
}

func main() {
	// Intervals: [0, 30], [5, 10], [15, 20]
	// 0--Start (1)
	// 5--Start (2)
	// 10-End   (1)
	// 15-Start (2)
	// 20-End   (1)
	// 30-End   (0)
	// Max overlaps: 2
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Printf("Intervals: %v\n", intervals)
	fmt.Printf("Min Rooms Needed: %d\n", minMeetingRooms(intervals))

	// Intervals: [7,10], [2,4]
	// No overlap. Max 1.
	intervals2 := [][]int{{7, 10}, {2, 4}}
	fmt.Printf("Intervals: %v\n", intervals2)
	fmt.Printf("Min Rooms Needed: %d\n", minMeetingRooms(intervals2))
}
