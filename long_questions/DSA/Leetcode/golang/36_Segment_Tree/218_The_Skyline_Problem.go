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
