package main

import "fmt"

// 733. Flood Fill
// Time: O(M*N), Space: O(M*N) for recursion stack
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}
	
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	
	dfs(image, sr, sc, originalColor, color)
	return image
}

func dfs(image [][]int, row, col, originalColor, newColor int) {
	m, n := len(image), len(image[0])
	
	// Boundary check and color check
	if row < 0 || row >= m || col < 0 || col >= n || image[row][col] != originalColor {
		return
	}
	
	// Fill the current cell
	image[row][col] = newColor
	
	// Recursively fill all 4 directions
	dfs(image, row+1, col, originalColor, newColor)
	dfs(image, row-1, col, originalColor, newColor)
	dfs(image, row, col+1, originalColor, newColor)
	dfs(image, row, col-1, originalColor, newColor)
}

func main() {
	// Test cases
	testCases := []struct {
		image    [][]int
		sr, sc   int
		color    int
	}{
		{
			[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			1, 1, 2,
		},
		{
			[][]int{{0, 0, 0}, {0, 0, 0}},
			0, 0, 0,
		},
		{
			[][]int{{0}},
			0, 0, 1,
		},
		{
			[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			0, 0, 2,
		},
		{
			[][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			1, 1, 2,
		},
		{
			[][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			2, 2, 3,
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			0, 0, 10,
		},
	}
	
	for i, tc := range testCases {
		// Make a copy to preserve original for display
		original := make([][]int, len(tc.image))
		for j := range tc.image {
			original[j] = make([]int, len(tc.image[j]))
			copy(original[j], tc.image[j])
		}
		
		result := floodFill(tc.image, tc.sr, tc.sc, tc.color)
		fmt.Printf("Test Case %d: image=%v, start=(%d,%d), color=%d\n", 
			i+1, original, tc.sr, tc.sc, tc.color)
		fmt.Printf("  Result: %v\n\n", result)
	}
}
