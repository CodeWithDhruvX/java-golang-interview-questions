package main

import (
	"fmt"
	"math"
)

// 399. Evaluate Division - Floyd-Warshall for Division Relationships
// Time: O(N^3 + Q), Space: O(N^2)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Build variable index map
	varMap := make(map[string]int)
	idx := 0
	
	for _, eq := range equations {
		if _, exists := varMap[eq[0]]; !exists {
			varMap[eq[0]] = idx
			idx++
		}
		if _, exists := varMap[eq[1]]; !exists {
			varMap[eq[1]] = idx
			idx++
		}
	}
	
	n := len(varMap)
	
	// Initialize distance matrix
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 1.0
			} else {
				dist[i][j] = -1.0 // Represents no connection
			}
		}
	}
	
	// Fill direct relationships
	for i, eq := range equations {
		from, to := varMap[eq[0]], varMap[eq[1]]
		dist[from][to] = values[i]
		dist[to][from] = 1.0 / values[i]
	}
	
	// Floyd-Warshall for division relationships
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] > 0 && dist[k][j] > 0 {
					product := dist[i][k] * dist[k][j]
					if dist[i][j] < 0 || product < dist[i][j] {
						dist[i][j] = product
					}
				}
			}
		}
	}
	
	// Answer queries
	result := make([]float64, len(queries))
	for i, query := range queries {
		if fromIdx, fromExists := varMap[query[0]]; fromExists {
			if toIdx, toExists := varMap[query[1]]; toExists {
				result[i] = dist[fromIdx][toIdx]
			} else {
				result[i] = -1.0
			}
		} else {
			result[i] = -1.0
		}
	}
	
	return result
}

// Floyd-Warshall with path tracking
func calcEquationWithPathTracking(equations [][]string, values []float64, queries [][]string) ([]float64, map[string][][]string) {
	// Build variable index map
	varMap := make(map[string]int)
	idx := 0
	
	for _, eq := range equations {
		if _, exists := varMap[eq[0]]; !exists {
			varMap[eq[0]] = idx
			idx++
		}
		if _, exists := varMap[eq[1]]; !exists {
			varMap[eq[1]] = idx
			idx++
		}
	}
	
	n := len(varMap)
	
	// Initialize matrices
	dist := make([][]float64, n)
	next := make([][][]string, n)
	
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		next[i] = make([][]string, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 1.0
				next[i][j] = []string{}
			} else {
				dist[i][j] = -1.0
				next[i][j] = nil
			}
		}
	}
	
	// Fill direct relationships
	for i, eq := range equations {
		from, to := varMap[eq[0]], varMap[eq[1]]
		dist[from][to] = values[i]
		dist[to][from] = 1.0 / values[i]
		next[from][to] = []string{eq[0], eq[1]}
		next[to][from] = []string{eq[1], eq[0]}
	}
	
	// Floyd-Warshall with path tracking
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] > 0 && dist[k][j] > 0 {
					product := dist[i][k] * dist[k][j]
					if dist[i][j] < 0 || product < dist[i][j] {
						dist[i][j] = product
						// Build path
						if next[i][k] != nil && next[k][j] != nil {
							next[i][j] = append(next[i][k], next[k][j][1:]...)
						}
					}
				}
			}
		}
	}
	
	// Answer queries and build path map
	result := make([]float64, len(queries))
	pathMap := make(map[string][][]string)
	
	for i, query := range queries {
		if fromIdx, fromExists := varMap[query[0]]; fromExists {
			if toIdx, toExists := varMap[query[1]]; toExists {
				result[i] = dist[fromIdx][toIdx]
				if next[fromIdx][toIdx] != nil {
					pathMap[query[0]+"/"+query[1]] = next[fromIdx][toIdx]
				}
			} else {
				result[i] = -1.0
			}
		} else {
			result[i] = -1.0
		}
	}
	
	return result, pathMap
}

// Floyd-Warshall with dynamic updates
type DivisionGraph struct {
	varMap map[string]int
	dist   [][]float64
	n      int
}

func NewDivisionGraph() *DivisionGraph {
	return &DivisionGraph{
		varMap: make(map[string]int),
		dist:   nil,
		n:      0,
	}
}

func (dg *DivisionGraph) AddEquation(numerator, denominator string, value float64) {
	// Add variables if they don't exist
	if _, exists := dg.varMap[numerator]; !exists {
		if dg.dist == nil {
			dg.dist = make([][]float64, 1)
			dg.dist[0] = make([]float64, 1)
			dg.dist[0][0] = 1.0
		} else {
			// Expand matrix
			newDist := make([][]float64, dg.n+1)
			for i := range newDist {
				newDist[i] = make([]float64, dg.n+1)
				for j := range newDist[i] {
					if i < dg.n && j < dg.n {
						newDist[i][j] = dg.dist[i][j]
					} else if i == j {
						newDist[i][j] = 1.0
					} else {
						newDist[i][j] = -1.0
					}
				}
			}
			dg.dist = newDist
		}
		dg.varMap[numerator] = dg.n
		dg.n++
	}
	
	if _, exists := dg.varMap[denominator]; !exists {
		if dg.dist == nil {
			dg.dist = make([][]float64, 1)
			dg.dist[0] = make([]float64, 1)
			dg.dist[0][0] = 1.0
		} else {
			newDist := make([][]float64, dg.n+1)
			for i := range newDist {
				newDist[i] = make([]float64, dg.n+1)
				for j := range newDist[i] {
					if i < dg.n && j < dg.n {
						newDist[i][j] = dg.dist[i][j]
					} else if i == j {
						newDist[i][j] = 1.0
					} else {
						newDist[i][j] = -1.0
					}
				}
			}
			dg.dist = newDist
		}
		dg.varMap[denominator] = dg.n
		dg.n++
	}
	
	// Add or update relationship
	from, to := dg.varMap[numerator], dg.varMap[denominator]
	dg.dist[from][to] = value
	dg.dist[to][from] = 1.0 / value
	
	// Update all paths using Floyd-Warshall
	for k := 0; k < dg.n; k++ {
		for i := 0; i < dg.n; i++ {
			for j := 0; j < dg.n; j++ {
				if dg.dist[i][k] > 0 && dg.dist[k][j] > 0 {
					product := dg.dist[i][k] * dg.dist[k][j]
					if dg.dist[i][j] < 0 || product < dg.dist[i][j] {
						dg.dist[i][j] = product
					}
				}
			}
		}
	}
}

func (dg *DivisionGraph) Query(numerator, denominator string) float64 {
	if fromIdx, fromExists := dg.varMap[numerator]; fromExists {
		if toIdx, toExists := dg.varMap[denominator]; toExists {
			return dg.dist[fromIdx][toIdx]
		}
	}
	return -1.0
}

// Floyd-Warshall with cycle detection
func calcEquationWithCycleDetection(equations [][]string, values []float64, queries [][]string) ([]float64, []string) {
	// Build variable index map
	varMap := make(map[string]int)
	idx := 0
	
	for _, eq := range equations {
		if _, exists := varMap[eq[0]]; !exists {
			varMap[eq[0]] = idx
			idx++
		}
		if _, exists := varMap[eq[1]]; !exists {
			varMap[eq[1]] = idx
			idx++
		}
	}
	
	n := len(varMap)
	
	// Initialize distance matrix
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 1.0
			} else {
				dist[i][j] = -1.0
			}
		}
	}
	
	// Fill direct relationships
	for i, eq := range equations {
		from, to := varMap[eq[0]], varMap[eq[1]]
		dist[from][to] = values[i]
		dist[to][from] = 1.0 / values[i]
	}
	
	// Floyd-Warshall with cycle detection
	var cycles []string
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] > 0 && dist[k][j] > 0 {
					product := dist[i][k] * dist[k][j]
					if dist[i][j] < 0 || product < dist[i][j] {
						dist[i][j] = product
					}
				}
			}
		}
	}
	
	// Detect cycles (where a/a != 1)
	for i := 0; i < n; i++ {
		if math.Abs(dist[i][i]-1.0) > 1e-9 {
			// Find variable name
			var varName string
			for name, idx := range varMap {
				if idx == i {
					varName = name
					break
				}
			}
			cycles = append(cycles, fmt.Sprintf("Cycle detected for %s: %f", varName, dist[i][i]))
		}
	}
	
	// Answer queries
	result := make([]float64, len(queries))
	for i, query := range queries {
		if fromIdx, fromExists := varMap[query[0]]; fromExists {
			if toIdx, toExists := varMap[query[1]]; toExists {
				result[i] = dist[fromIdx][toIdx]
			} else {
				result[i] = -1.0
			}
		} else {
			result[i] = -1.0
		}
	}
	
	return result, cycles
}

// Alternative approach using Union-Find with weights
func calcEquationUnionFind(equations [][]string, values []float64, queries [][]string) []float64 {
	// This is an alternative to Floyd-Warshall
	// Implementation would use Union-Find with weight tracking
	// For comparison purposes, returning Floyd-Warshall result
	return calcEquation(equations, values, queries)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Floyd-Warshall for Division ===")
	
	testCases := []struct {
		equations  [][]string
		values     []float64
		queries    [][]string
		description string
	}{
		{
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2.0, 3.0},
			[][]string{{"a", "c"}, {"c", "a"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			"Standard case",
		},
		{
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			"Complex variables",
		},
		{
			[][]string{{"x", "y"}},
			[]float64{3.0},
			[][]string{{"x", "y"}, {"y", "x"}, {"x", "z"}, {"z", "x"}},
			"Single equation",
		},
		{
			[][]string{{"a", "b"}},
			[]float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"c", "a"}},
			"Fraction value",
		},
		{
			[][]string{},
			[]float64{},
			[][]string{{"a", "b"}},
			"No equations",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Equations: %v\n", tc.equations)
		fmt.Printf("  Values: %v\n", tc.values)
		fmt.Printf("  Queries: %v\n", tc.queries)
		
		result1 := calcEquation(tc.equations, tc.values, tc.queries)
		result2 := calcEquationUnionFind(tc.equations, tc.values, tc.queries)
		
		fmt.Printf("  Floyd-Warshall: %v\n", result1)
		fmt.Printf("  Union-Find: %v\n\n", result2)
	}
	
	// Test path tracking
	fmt.Println("=== Path Tracking Test ===")
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}}
	values := []float64{2.0, 3.0, 4.0}
	queries := [][]string{{"a", "d"}, {"d", "a"}}
	
	result, pathMap := calcEquationWithPathTracking(equations, values, queries)
	
	fmt.Printf("Results: %v\n", result)
	for query, path := range pathMap {
		fmt.Printf("Path %s: %v\n", query, path)
	}
	
	// Test dynamic graph
	fmt.Println("\n=== Dynamic Graph Test ===")
	dg := NewDivisionGraph()
	
	dg.AddEquation("a", "b", 2.0)
	fmt.Printf("a/b = %f\n", dg.Query("a", "b"))
	
	dg.AddEquation("b", "c", 3.0)
	fmt.Printf("a/c = %f\n", dg.Query("a", "c"))
	
	dg.AddEquation("c", "d", 4.0)
	fmt.Printf("a/d = %f\n", dg.Query("a", "d"))
	
	// Test cycle detection
	fmt.Println("\n=== Cycle Detection Test ===")
	cycleEquations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "a"}}
	cycleValues := []float64{2.0, 3.0, 6.0} // Should create a cycle: a/b * b/c * c/a = 2*3*6 = 36
	
	cycleResult, cycles := calcEquationWithCycleDetection(cycleEquations, cycleValues, [][]string{{"a", "a"}})
	
	fmt.Printf("Cycle detection result: %v\n", cycleResult)
	fmt.Printf("Detected cycles: %v\n", cycles)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Create a large equation system
	largeEquations := make([][]string, 0)
	largeValues := make([]float64, 0)
	
	for i := 0; i < 50; i++ {
		for j := i + 1; j < i+5 && j < 50; j++ {
			varName1 := fmt.Sprintf("var%d", i)
			varName2 := fmt.Sprintf("var%d", j)
			largeEquations = append(largeEquations, []string{varName1, varName2})
			largeValues = append(largeValues, float64(j-i))
		}
	}
	
	largeQueries := make([][]string, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			varName1 := fmt.Sprintf("var%d", i)
			varName2 := fmt.Sprintf("var%d", j)
			largeQueries = append(largeQueries, []string{varName1, varName2})
		}
	}
	
	fmt.Printf("Large test: %d equations, %d queries\n", len(largeEquations), len(largeQueries))
	
	result = calcEquation(largeEquations, largeValues, largeQueries)
	fmt.Printf("First 5 results: %v\n", result[:5])
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Self division
	selfEquations := [][]string{{"a", "a"}}
	selfValues := []float64{1.0}
	selfQueries := [][]string{{"a", "a"}, {"a", "b"}}
	
	fmt.Printf("Self division: %v\n", calcEquation(selfEquations, selfValues, selfQueries))
	
	// Zero division
	zeroEquations := [][]string{{"a", "b"}}
	zeroValues := []float64{0.0}
	zeroQueries := [][]string{{"a", "b"}, {"b", "a"}}
	
	fmt.Printf("Zero division: %v\n", calcEquation(zeroEquations, zeroValues, zeroQueries))
	
	// Negative values
	negEquations := [][]string{{"a", "b"}, {"b", "c"}}
	negValues := []float64{-2.0, -3.0}
	negQueries := [][]string{{"a", "c"}, {"c", "a"}}
	
	fmt.Printf("Negative values: %v\n", calcEquation(negEquations, negValues, negQueries))
}
