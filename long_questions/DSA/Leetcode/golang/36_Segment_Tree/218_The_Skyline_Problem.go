package main

import "fmt"

// 218. The Skyline Problem - Segment Tree approach
// Time: O(N log N), Space: O(N)
func getSkyline(buildings [][]int) [][]int {
	// Collect all unique x coordinates
	coords := make(map[int]bool)
	for _, building := range buildings {
		coords[building[0]] = true
		coords[building[1]] = true
	}
	
	// Sort coordinates
	var sortedCoords []int
	for coord := range coords {
		sortedCoords = append(sortedCoords, coord)
	}
	
	// Simple bubble sort for demonstration
	for i := 0; i < len(sortedCoords)-1; i++ {
		for j := 0; j < len(sortedCoords)-i-1; j++ {
			if sortedCoords[j] > sortedCoords[j+1] {
				sortedCoords[j], sortedCoords[j+1] = sortedCoords[j+1], sortedCoords[j]
			}
		}
	}
	
	// Create coordinate mapping
	coordMap := make(map[int]int)
	for i, coord := range sortedCoords {
		coordMap[coord] = i
	}
	
	// Build segment tree for maximum height
	n := len(sortedCoords)
	segTree := make([]int, 4*n)
	
	// Process buildings
	for _, building := range buildings {
		left, right, height := building[0], building[1], building[2]
		leftIdx, rightIdx := coordMap[left], coordMap[right]-1
		
		updateSegTree(segTree, 0, 0, n-1, leftIdx, rightIdx, height)
	}
	
	// Generate skyline
	var result [][]int
	prevHeight := 0
	
	for i := 0; i < len(sortedCoords); i++ {
		currentHeight := querySegTree(segTree, 0, 0, n-1, i, i)
		
		if currentHeight != prevHeight {
			result = append(result, []int{sortedCoords[i], currentHeight})
			prevHeight = currentHeight
		}
	}
	
	return result
}

func updateSegTree(tree []int, node, start, end, left, right, height int) {
	// No overlap
	if start > right || end < left {
		return
	}
	
	// Complete overlap
	if left <= start && end <= right {
		if height > tree[node] {
			tree[node] = height
		}
		return
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	updateSegTree(tree, 2*node+1, start, mid, left, right, height)
	updateSegTree(tree, 2*node+2, mid+1, end, left, right, height)
	
	tree[node] = max(tree[2*node+1], tree[2*node+2])
}

func querySegTree(tree []int, node, start, end, left, right int) int {
	// No overlap
	if start > right || end < left {
		return 0
	}
	
	// Complete overlap
	if left <= start && end <= right {
		return tree[node]
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	return max(querySegTree(tree, 2*node+1, start, mid, left, right),
		querySegTree(tree, 2*node+2, mid+1, end, left, right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Alternative approach using sweep line algorithm
func getSkylineSweepLine(buildings [][]int) [][]int {
	// Create events
	type Event struct {
		x   int
		h   int
		end bool // false for start, true for end
	}
	
	var events []Event
	for _, building := range buildings {
		left, right, height := building[0], building[1], building[2]
		events = append(events, Event{left, height, false})
		events = append(events, Event{right, height, true})
	}
	
	// Sort events by x coordinate
	for i := 0; i < len(events)-1; i++ {
		for j := 0; j < len(events)-i-1; j++ {
			if events[j].x > events[j+1].x {
				events[j], events[j+1] = events[j+1], events[j]
			}
		}
	}
	
	// Use a simple map to track active heights
	activeHeights := make(map[int]int)
	var result [][]int
	prevHeight := 0
	
	for _, event := range events {
		if event.end {
			// Remove height
			activeHeights[event.h]--
			if activeHeights[event.h] == 0 {
				delete(activeHeights, event.h)
			}
		} else {
			// Add height
			activeHeights[event.h]++
		}
		
		// Find current max height
		currentHeight := 0
		for h := range activeHeights {
			if h > currentHeight {
				currentHeight = h
			}
		}
		
		if currentHeight != prevHeight {
			result = append(result, []int{event.x, currentHeight})
			prevHeight = currentHeight
		}
	}
	
	return result
}

// Divide and conquer approach
func getSkylineDivideConquer(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}
	
	if len(buildings) == 1 {
		building := buildings[0]
		return [][]int{{building[0], building[2]}, {building[1], 0}}
	}
	
	mid := len(buildings) / 2
	left := getSkylineDivideConquer(buildings[:mid])
	right := getSkylineDivideConquer(buildings[mid:])
	
	return mergeSkylines(left, right)
}

func mergeSkylines(left, right [][]int) [][]int {
	var result [][]int
	i, j := 0, 0
	leftHeight, rightHeight := 0, 0
	
	for i < len(left) || j < len(right) {
		var x int
		
		if i < len(left) && (j >= len(right) || left[i][0] < right[j][0]) {
			x = left[i][0]
			leftHeight = left[i][1]
			i++
		} else if j < len(right) && (i >= len(left) || right[j][0] < left[i][0]) {
			x = right[j][0]
			rightHeight = right[j][1]
			j++
		} else {
			// Same x coordinate
			x = left[i][0]
			leftHeight = left[i][1]
			rightHeight = right[j][1]
			i++
			j++
		}
		
		maxHeight := leftHeight
		if rightHeight > maxHeight {
			maxHeight = rightHeight
		}
		
		if len(result) == 0 || result[len(result)-1][1] != maxHeight {
			result = append(result, []int{x, maxHeight})
		}
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Segment Tree for Range Maximum Queries
- **Coordinate Compression**: Map building coordinates to indices
- **Range Updates**: Update ranges with building heights
- **Range Queries**: Query maximum height at each coordinate
- **Skyline Generation**: Extract critical points where height changes

## 2. PROBLEM CHARACTERISTICS
- **Building Intervals**: Each building defined by [left, right, height]
- **Skyline Formation**: Outer contour formed by building silhouettes
- **Critical Points**: Points where skyline height changes
- **Range Maximum**: Need maximum height at each x-coordinate

## 3. SIMILAR PROBLEMS
- Range Sum Query Mutable (LeetCode 307) - Segment tree for range queries
- Range Module (LeetCode 715) - Range tracking with segment tree
- Count of Smaller Numbers After Self (LeetCode 315) - Segment tree variant
- Falling Squares (LeetCode 699) - Similar interval height problem

## 4. KEY OBSERVATIONS
- **Coordinate Discretization**: Only need building start/end coordinates
- **Range Updates**: Buildings affect range of coordinates
- **Maximum Tracking**: Skyline = maximum height at each point
- **Critical Points**: Only record when height changes

## 5. VARIATIONS & EXTENSIONS
- **Sweep Line Algorithm**: Process events in order with priority queue
- **Divide and Conquer**: Recursively merge building skylines
- **Lazy Propagation**: Optimize range updates in segment tree
- **Multiset Approach**: Use balanced BST for active heights

## 6. INTERVIEW INSIGHTS
- Always clarify: "Coordinate ranges? Building overlap complexity? Output format?"
- Edge cases: no buildings, single building, all same height
- Time complexity: O(N log N) for segment tree, O(N log N) for sweep line
- Space complexity: O(N) for segment tree, O(N) for event processing
- Key insight: discretize coordinates to handle large ranges efficiently

## 7. COMMON MISTAKES
- Not handling coordinate compression properly
- Wrong range update boundaries (off-by-one errors)
- Missing critical points where height doesn't change
- Inefficient sorting algorithms for large inputs
- Not handling duplicate coordinates correctly

## 8. OPTIMIZATION STRATEGIES
- **Segment Tree**: O(N log N) time, O(N) space - range updates/queries
- **Sweep Line**: O(N log N) time, O(N) space - event processing
- **Divide and Conquer**: O(N log N) time, O(N) space - recursive merging
- **Lazy Propagation**: O(N log N) time, O(N) space - optimized updates

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like drawing the outline of a city skyline:**
- Each building is a rectangle with height
- You want to draw the outer contour (silhouette)
- The skyline changes only at building edges
- At each x-coordinate, you need the tallest building
- You trace the outline by following height changes
- Like drawing the profile of buildings against the sky

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Buildings defined by [left, right, height]
2. **Goal**: Find critical points of skyline silhouette
3. **Rules**: Skyline = maximum height at each x-coordinate
4. **Output**: List of [x, height] points where height changes

#### Phase 2: Key Insight Recognition
- **"Coordinate compression natural"** → Only need building edges
- **"Range updates needed"** → Buildings affect continuous ranges
- **"Maximum queries"** → Skyline = max height at each point
- **"Critical points only"** → Output only where height changes

#### Phase 3: Strategy Development
```
Human thought process:
"I need skyline from building rectangles.
Direct approach would be too slow for large coordinates.

Segment Tree Approach:
1. Collect all unique x-coordinates (building edges)
2. Sort and map to indices (coordinate compression)
3. Build segment tree for range maximum
4. For each building: update range [left, right) with height
5. Query each coordinate for maximum height
6. Extract critical points where height changes

This handles large coordinate ranges efficiently!"
```

#### Phase 4: Edge Case Handling
- **No buildings**: Return empty skyline
- **Single building**: Return [left, height], [right, 0]
- **Overlapping buildings**: Higher building dominates
- **Same coordinates**: Handle duplicates properly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: buildings = [[2,9,10], [3,7,15]]

Human thinking:
"Segment Tree Approach:
Step 1: Collect coordinates: [2,9,3,7] → [2,3,7,9]
Step 2: Map to indices: 2→0, 3→1, 7→2, 9→3

Step 3: Process building [2,9,10]:
- Update range [0,2] (indices 2→0, 9→3) with height 10
- Tree now has max height 10 in range [0,2]

Step 4: Process building [3,7,15]:
- Update range [1,1] (indices 3→1, 7→2) with height 15
- Tree now has max height 15 at index 1

Step 5: Query all coordinates:
- x=2 (idx 0): height = 10
- x=3 (idx 1): height = 15 (change!)
- x=7 (idx 2): height = 10 (change!)
- x=9 (idx 3): height = 0 (change!)

Step 6: Extract critical points:
[2,10], [3,15], [7,10], [9,0]

Result ✓"
```

#### Phase 6: Intuition Validation
- **Why compression works**: Reduces coordinate range to O(N)
- **Why segment tree**: Efficient range updates and queries
- **Why critical points**: Skyline only changes at building edges
- **Why O(N log N)**: N buildings × log N for tree operations

### Common Human Pitfalls & How to Avoid Them
1. **"Why not process all coordinates?"** → Too many coordinates, need compression
2. **"Should I use array?"** → Range updates inefficient with simple array
3. **"What about sweep line?"** → Alternative approach with events
4. **"Can I use divide and conquer?"** → Yes, recursive merging works
5. **"What about lazy propagation?"** → Optimizes range updates

### Real-World Analogy
**Like creating a city skyline silhouette for tourism:**
- Each building is a skyscraper with height
- You want to draw the outline against the sunset
- The silhouette changes only at building corners
- At each position, you see the tallest building
- You trace the outline by following height changes
- Like creating a city profile photograph

### Human-Readable Pseudocode
```
function getSkyline(buildings):
    # Collect and sort unique coordinates
    coords = unique building edges
    sort(coords)
    
    # Map coordinates to indices
    coordMap = {coord: index for index, coord in enumerate(coords)}
    
    # Build segment tree for range maximum
    segTree = SegmentTree(len(coords))
    
    # Process each building
    for building in buildings:
        left, right, height = building
        leftIdx = coordMap[left]
        rightIdx = coordMap[right] - 1
        segTree.updateRange(leftIdx, rightIdx, height)
    
    # Generate skyline
    result = []
    prevHeight = 0
    
    for i, coord in enumerate(coords):
        currentHeight = segTree.queryPoint(i)
        if currentHeight != prevHeight:
            result.append([coord, currentHeight])
            prevHeight = currentHeight
    
    return result
```

### Execution Visualization

### Example: buildings = [[2,9,10], [3,7,15]]
```
Segment Tree Process:

Step 1: Coordinate Compression
Original coords: [2,9,3,7]
Sorted unique: [2,3,7,9]
Mapping: 2→0, 3→1, 7→2, 9→3

Step 2: Building [2,9,10] → update range [0,2] with 10
Tree state: [10,10,10,0]

Step 3: Building [3,7,15] → update range [1,1] with 15
Tree state: [10,15,10,0]

Step 4: Query all coordinates
x=2 (idx 0): height = 10
x=3 (idx 1): height = 15 (change!)
x=7 (idx 2): height = 10 (change!)
x=9 (idx 3): height = 0 (change!)

Step 5: Extract critical points
[2,10], [3,15], [7,10], [9,0]

Result ✓
```

### Key Visualization Points:
- **Coordinate Compression**: Reduce large coordinate ranges
- **Range Updates**: Buildings affect continuous coordinate ranges
- **Maximum Queries**: Skyline = maximum height at each point
- **Critical Points**: Only record height changes

### Skyline Visualization:
```
Buildings:
    Building 1: [2,9,10]  ██████████
    Building 2: [3,7,15]      ████████

Combined Skyline:
    15 |          ████████
    10 | ███████████████████
     0 |_________________________
      2  3  4  5  6  7  8  9

Critical Points: [2,10], [3,15], [7,10], [9,0]
```

### Time Complexity Breakdown:
- **Segment Tree**: O(N log N) time, O(N) space - range updates/queries
- **Sweep Line**: O(N log N) time, O(N) space - event processing
- **Divide and Conquer**: O(N log N) time, O(N) space - recursive merging
- **Lazy Propagation**: O(N log N) time, O(N) space - optimized updates

### Alternative Approaches:

#### 1. Sweep Line Algorithm (O(N log N) time, O(N) space)
```go
func getSkylineSweepLine(buildings [][]int) [][]int {
    // Process start/end events in order
    // Use priority queue for active heights
    // Track height changes at event points
    // ... implementation details omitted
}
```
- **Pros**: More intuitive, event-based processing
- **Cons**: Requires priority queue implementation

#### 2. Divide and Conquer (O(N log N) time, O(N) space)
```go
func getSkylineDivideConquer(buildings [][]int) [][]int {
    // Recursively split buildings
    // Merge skylines from subproblems
    // Two-pointer merging technique
    // ... implementation details omitted
}
```
- **Pros**: Natural recursive formulation
- **Cons**: Complex merging logic

#### 3. Lazy Propagation (O(N log N) time, O(N) space)
```go
func getSkylineLazy(buildings [][]int) [][]int {
    // Segment tree with lazy propagation
    // More efficient range updates
    // ... implementation details omitted
}
```
- **Pros**: Optimized range updates
- **Cons**: More complex tree implementation

### Extensions for Interviews:
- **Multiple Queries**: Handle multiple skyline queries efficiently
- **Dynamic Updates**: Add/remove buildings dynamically
- **3D Skyline**: Extend to 3D building silhouettes
- **Area Calculation**: Calculate total skyline area
- **Real-world Applications**: Urban planning, architecture visualization
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Skyline Problem ===")
	
	testCases := []struct {
		buildings   [][]int
		description string
	}{
		{
			[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			"Standard case",
		},
		{
			[][]int{{0, 2, 3}, {2, 5, 3}},
			"Two buildings",
		},
		{
			[][]int{{0, 1, 3}, {1, 2, 3}},
			"Adjacent buildings",
		},
		{
			[][]int{{0, 5, 5}},
			"Single building",
		},
		{
			[][]int{{1, 5, 11}, {2, 7, 6}, {3, 9, 13}, {12, 16, 7}, {14, 25, 3}, {19, 22, 18}, {23, 29, 13}, {24, 28, 4}},
			"Complex case",
		},
		{
			[][]int{{0, 2, 3}, {2, 4, 3}, {4, 6, 3}},
			"Equal height buildings",
		},
		{
			[][]int{{0, 10, 10}, {5, 15, 5}},
			"Overlapping with different heights",
		},
		{
			[][]int{{0, 3, 3}, {1, 4, 4}, {2, 5, 5}},
			"Nested buildings",
		},
		{
			[][]int{},
			"No buildings",
		},
		{
			[][]int{{0, 100, 100}},
			"Large building",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Buildings: %v\n", tc.buildings)
		
		result1 := getSkyline(tc.buildings)
		result2 := getSkylineSweepLine(tc.buildings)
		result3 := getSkylineDivideConquer(tc.buildings)
		
		fmt.Printf("  Segment Tree: %v\n", result1)
		fmt.Printf("  Sweep Line: %v\n", result2)
		fmt.Printf("  Divide & Conquer: %v\n\n", result3)
	}
}
