package main

import "fmt"

// 54. Spiral Matrix
// Time: O(M*N), Space: O(1) (excluding output)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	
	top, bottom := 0, m-1
	left, right := 0, n-1
	
	for top <= bottom && left <= right {
		// Traverse from left to right (top row)
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++
		
		// Traverse from top to bottom (right column)
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--
		
		// Traverse from right to left (bottom row) if still valid
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}
		
		// Traverse from bottom to top (left column) if still valid
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}
	
	return result
}

// Alternative approach using direction vectors
func spiralOrderDirections(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	
	// Directions: right, down, left, up
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex := 0
	
	// Boundaries
	top, bottom := 0, m-1
	left, right := 0, n-1
	
	row, col := 0, 0
	
	for len(result) < m*n {
		result = append(result, matrix[row][col])
		
		// Calculate next position
		nextRow := row + directions[dirIndex][0]
		nextCol := col + directions[dirIndex][1]
		
		// Check if we need to change direction
		if nextRow < top || nextRow > bottom || nextCol < left || nextCol > right {
			// Update boundaries and change direction
			switch dirIndex {
			case 0: // Moving right
				top++
			case 1: // Moving down
				right--
			case 2: // Moving left
				bottom--
			case 3: // Moving up
				left++
			}
			
			dirIndex = (dirIndex + 1) % 4
			nextRow = row + directions[dirIndex][0]
			nextCol = col + directions[dirIndex][1]
		}
		
		row, col = nextRow, nextCol
	}
	
	return result
}

// Recursive approach
func spiralOrderRecursive(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	var result []int
	spiralHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, &result)
	return result
}

func spiralHelper(matrix [][]int, top, bottom, left, right int, result *[]int) {
	if top > bottom || left > right {
		return
	}
	
	// Traverse from left to right (top row)
	for col := left; col <= right; col++ {
		*result = append(*result, matrix[top][col])
	}
	
	// Traverse from top to bottom (right column)
	for row := top + 1; row <= bottom; row++ {
		*result = append(*result, matrix[row][right])
	}
	
	// Traverse from right to left (bottom row) if still valid
	if top < bottom {
		for col := right - 1; col >= left; col-- {
			*result = append(*result, matrix[bottom][col])
		}
	}
	
	// Traverse from bottom to top (left column) if still valid
	if left < right {
		for row := bottom - 1; row > top; row-- {
			*result = append(*result, matrix[row][left])
		}
	}
	
	// Recursively process inner matrix
	spiralHelper(matrix, top+1, bottom-1, left+1, right-1, result)
}

func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		{{1, 2, 3}, {4, 5, 6}},
		{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		{{1}},
		{{1, 2, 3, 4, 5}},
		{{1}, {2}, {3}, {4}, {5}},
		{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5}, {6, 7, 8}},
		{{}},
	}
	
	for i, matrix := range testCases {
		fmt.Printf("Test Case %d: %v\n", i+1, matrix)
		
		result1 := spiralOrder(matrix)
		result2 := spiralOrderDirections(matrix)
		result3 := spiralOrderRecursive(matrix)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Directions: %v\n", result2)
		fmt.Printf("  Recursive: %v\n\n", result3)
	}
}
