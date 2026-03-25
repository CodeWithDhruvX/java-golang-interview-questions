# 11. Sorting Customization and Methods in Go

## Table of Contents
1. [Built-in Sorting Functions](#built-in-sorting-functions)
2. [Custom Sort Implementation](#custom-sort-implementation)
3. [Sort Interface Implementation](#sort-interface-implementation)
4. [Advanced Sorting Techniques](#advanced-sorting-techniques)
5. [Performance Considerations](#performance-considerations)
6. [Interview Questions](#interview-questions)

---

## Built-in Sorting Functions

### 1. Basic Sort Functions

```go
package main

import (
	"fmt"
	"sort"
)

// Integer sorting
func basicIntSorting() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	
	// Sort in ascending order
	sort.Ints(numbers)
	fmt.Println("Sorted integers:", numbers) // [11 12 22 25 34 64 90]
	
	// Check if sorted
	isSorted := sort.IntsAreSorted(numbers)
	fmt.Println("Is sorted:", isSorted) // true
}

// Float sorting
func basicFloatSorting() {
	floats := []float64{3.14, 2.71, 1.41, 1.73, 0.57}
	
	sort.Float64s(floats)
	fmt.Println("Sorted floats:", floats) // [0.57 1.41 1.73 2.71 3.14]
	
	isSorted := sort.Float64sAreSorted(floats)
	fmt.Println("Is sorted:", isSorted) // true
}

// String sorting
func basicStringSorting() {
	strings := []string{"banana", "apple", "cherry", "date"}
	
	sort.Strings(strings)
	fmt.Println("Sorted strings:", strings) // [apple banana cherry date]
	
	isSorted := sort.StringsAreSorted(strings)
	fmt.Println("Is sorted:", isSorted) // true
}
```

### 2. Generic Sort Functions (Go 1.18+)

```go
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Generic sorting using constraints
func genericSorting[T constraints.Ordered](slice []T) {
	slices.Sort(slice)
	fmt.Printf("Sorted %T: %v\n", slice, slice)
}

// Usage examples
func useGenericSorting() {
	ints := []int{5, 2, 8, 1, 9}
	floats := []float64{3.14, 1.59, 2.65, 1.41}
	strings := []string{"zebra", "apple", "monkey"}
	
	genericSorting(ints)    // Sorted []int: [1 2 5 8 9]
	genericSorting(floats)  // Sorted []float64: [1.41 1.59 2.65 3.14]
	genericSorting(strings)  // Sorted []string: [apple monkey zebra]
}
```

---

## Custom Sort Implementation

### 1. Implementing Sort Interface

```go
package main

import (
	"fmt"
	"sort"
)

// Person struct for custom sorting
type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on Age field
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for []Person based on Name field
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func customSortExample() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
		{"Diana", 35},
	}
	
	fmt.Println("Original:", people)
	
	// Sort by age
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:", people)
	
	// Sort by name
	sort.Sort(ByName(people))
	fmt.Println("Sorted by name:", people)
}
```

### 2. Multi-Criteria Sorting

```go
package main

import (
	"fmt"
	"sort"
)

// Student with multiple fields
type Student struct {
	Name   string
	Grade  float64
	Age    int
	Active bool
}

// Multi-criteria sort: Grade (desc), then Age (asc), then Name (asc)
type ByMultipleCriteria []Student

func (s ByMultipleCriteria) Len() int { return len(s) }
func (s ByMultipleCriteria) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByMultipleCriteria) Less(i, j int) bool {
	// First criterion: Grade (descending)
	if s[i].Grade != s[j].Grade {
		return s[i].Grade > s[j].Grade
	}
	
	// Second criterion: Age (ascending)
	if s[i].Age != s[j].Age {
		return s[i].Age < s[j].Age
	}
	
	// Third criterion: Name (ascending)
	return s[i].Name < s[j].Name
}

func multiCriteriaSortExample() {
	students := []Student{
		{"Alice", 85.5, 20, true},
		{"Bob", 92.0, 22, true},
		{"Charlie", 85.5, 19, false},
		{"Diana", 92.0, 21, true},
		{"Eve", 78.0, 23, true},
	}
	
	fmt.Println("Original:", students)
	sort.Sort(ByMultipleCriteria(students))
	fmt.Println("Multi-criteria sorted:", students)
}
```

### 3. Stable vs Unstable Sorting

```go
package main

import (
	"fmt"
	"sort"
)

// Item with equal values but different original order
type Item struct {
	Value    int
	Original int // Original position
}

func demonstrateStableSort() {
	items := []Item{
		{2, 1}, {1, 2}, {2, 3}, {1, 4}, {3, 5}, {2, 6},
	}
	
	fmt.Println("Original:", items)
	
	// Stable sort (sort.Stable preserves order of equal elements)
	sort.Stable(sort.Slice(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	}))
	fmt.Println("Stable sorted:", items)
	
	// Reset and try unstable sort
	items = []Item{{2, 1}, {1, 2}, {2, 3}, {1, 4}, {3, 5}, {2, 6}}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	})
	fmt.Println("Unstable sorted:", items)
}
```

---

## Sort Interface Implementation

### 1. Complete Sort Interface Examples

```go
package main

import (
	"fmt"
	"sort"
)

// Product struct for e-commerce sorting
type Product struct {
	ID       int
	Name     string
	Price    float64
	Rating   float64
	Category string
}

// ByPrice implements sort.Interface
type ByPrice []Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }

// ByRating implements sort.Interface
type ByRating []Product

func (p ByRating) Len() int           { return len(p) }
func (p ByRating) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByRating) Less(i, j int) bool { return p[i].Rating > p[j].Rating } // Higher rating first

// ByCategoryThenPrice for nested sorting
type ByCategoryThenPrice []Product

func (p ByCategoryThenPrice) Len() int { return len(p) }
func (p ByCategoryThenPrice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByCategoryThenPrice) Less(i, j int) bool {
	if p[i].Category != p[j].Category {
		return p[i].Category < p[j].Category
	}
	return p[i].Price < p[j].Price
}

func productSortExample() {
	products := []Product{
		{1, "Laptop", 999.99, 4.5, "Electronics"},
		{2, "Book", 19.99, 4.2, "Books"},
		{3, "Phone", 699.99, 4.7, "Electronics"},
		{4, "Pen", 2.99, 3.8, "Stationery"},
		{5, "Tablet", 399.99, 4.3, "Electronics"},
		{6, "Notebook", 9.99, 4.1, "Stationery"},
	}
	
	fmt.Println("Original products:")
	for _, p := range products {
		fmt.Printf("%s: $%.2f (%.1f★) - %s\n", p.Name, p.Price, p.Rating, p.Category)
	}
	
	// Sort by price
	sort.Sort(ByPrice(products))
	fmt.Println("\nSorted by price:")
	for _, p := range products {
		fmt.Printf("%s: $%.2f\n", p.Name, p.Price)
	}
	
	// Sort by rating (desc)
	sort.Sort(ByRating(products))
	fmt.Println("\nSorted by rating (desc):")
	for _, p := range products {
		fmt.Printf("%s: %.1f★\n", p.Name, p.Rating)
	}
	
	// Sort by category then price
	sort.Sort(ByCategoryThenPrice(products))
	fmt.Println("\nSorted by category then price:")
	for _, p := range products {
		fmt.Printf("%s: %s - $%.2f\n", p.Name, p.Category, p.Price)
	}
}
```

### 2. Using sort.Slice for Custom Logic

```go
package main

import (
	"fmt"
	"sort"
)

// Employee struct
type Employee struct {
	Name     string
	Salary   int
	Department string
	Experience int
}

func sliceSortExamples() {
	employees := []Employee{
		{"John", 50000, "Engineering", 3},
		{"Jane", 60000, "Marketing", 5},
		{"Bob", 55000, "Engineering", 2},
		{"Alice", 65000, "Engineering", 4},
		{"Charlie", 45000, "Marketing", 1},
	}
	
	fmt.Println("Original:", employees)
	
	// Sort by salary (ascending)
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Salary < employees[j].Salary
	})
	fmt.Println("Sorted by salary:", employees)
	
	// Sort by experience (descending)
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Experience > employees[j].Experience
	})
	fmt.Println("Sorted by experience (desc):", employees)
	
	// Sort by department, then by salary
	sort.Slice(employees, func(i, j int) bool {
		if employees[i].Department != employees[j].Department {
			return employees[i].Department < employees[j].Department
		}
		return employees[i].Salary < employees[j].Salary
	})
	fmt.Println("Sorted by department then salary:", employees)
	
	// Custom logic: sort by salary-to-experience ratio
	sort.Slice(employees, func(i, j int) bool {
		ratioI := float64(employees[i].Salary) / float64(employees[i].Experience)
		ratioJ := float64(employees[j].Salary) / float64(employees[j].Experience)
		return ratioI > ratioJ
	})
	fmt.Println("Sorted by salary/experience ratio:", employees)
}
```

---

## Advanced Sorting Techniques

### 1. Partial Sorting

```go
package main

import (
	"fmt"
	"sort"
)

// Find top k elements using partial sorting
func topKElements(nums []int, k int) []int {
	if k >= len(nums) {
		sort.Ints(nums)
		return nums
	}
	
	// Use nth_element-like approach with sort.Slice
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Descending order
	})
	
	return nums[:k]
}

// Find k smallest elements
func kSmallestElements(nums []int, k int) []int {
	if k >= len(nums) {
		sort.Ints(nums)
		return nums
	}
	
	sort.Ints(nums)
	return nums[:k]
}

func partialSortExample() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}
	
	fmt.Println("Original:", numbers)
	
	// Get top 3 elements
	top3 := topKElements(append([]int{}, numbers...), 3)
	fmt.Println("Top 3 elements:", top3)
	
	// Get 4 smallest elements
	smallest4 := kSmallestElements(append([]int{}, numbers...), 4)
	fmt.Println("4 smallest elements:", smallest4)
}
```

### 2. Custom Sorting Algorithms

```go
package main

import (
	"fmt"
)

// QuickSort implementation
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	pivot := arr[len(arr)/2]
	var left, right, equal []int
	
	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}
	
	return append(append(quickSort(left), equal...), quickSort(right)...)
}

// MergeSort implementation
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

// Counting Sort for integers in a range
func countingSort(arr []int, maxVal int) []int {
	count := make([]int, maxVal+1)
	
	// Count occurrences
	for _, v := range arr {
		count[v]++
	}
	
	// Build result
	result := make([]int, 0, len(arr))
	for i, c := range count {
		for j := 0; j < c; j++ {
			result = append(result, i)
		}
	}
	
	return result
}

func customAlgorithmsExample() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}
	
	fmt.Println("Original:", numbers)
	
	fmt.Println("QuickSort:", quickSort(append([]int{}, numbers...)))
	fmt.Println("MergeSort:", mergeSort(append([]int{}, numbers...)))
	fmt.Println("CountingSort:", countingSort(append([]int{}, numbers...), 90))
}
```

### 3. Sorting with Custom Comparators

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Case-insensitive string sorting
type CaseInsensitive []string

func (s CaseInsensitive) Len() int           { return len(s) }
func (s CaseInsensitive) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(s[i]) < strings.ToLower(s[j])
}

// Roman numeral comparison
func romanToValue(roman string) int {
	values := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0
	
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && values[rune(roman[i])] < values[rune(roman[i+1])] {
			total += values[rune(roman[i+1])] - values[rune(roman[i])]
			i++
		} else {
			total += values[rune(roman[i])]
		}
	}
	
	return total
}

type ByRomanNumeral []string

func (r ByRomanNumeral) Len() int           { return len(r) }
func (r ByRomanNumeral) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRomanNumeral) Less(i, j int) bool {
	return romanToValue(r[i]) < romanToValue(r[j])
}

func customComparatorExample() {
	// Case-insensitive sorting
	words := []string{"Banana", "apple", "Cherry", "date", "Elderberry"}
	fmt.Println("Original strings:", words)
	sort.Sort(CaseInsensitive(words))
	fmt.Println("Case-insensitive sorted:", words)
	
	// Roman numeral sorting
	romans := []string{"X", "II", "IV", "IX", "I", "V", "III"}
	fmt.Println("\nRoman numerals:", romans)
	sort.Sort(ByRomanNumeral(romans))
	fmt.Println("Sorted by value:", romans)
}
```

---

## Performance Considerations

### 1. Time and Space Complexity Analysis

```go
package main

import (
	"fmt"
	"sort"
	"time"
)

// Performance comparison
func compareSortingPerformance() {
	sizes := []int{1000, 10000, 100000}
	
	for _, size := range sizes {
		// Generate random data
		data := make([]int, size)
		for i := range data {
			data[i] = i * 12345 % size // Pseudo-random
		}
		
		// Test built-in sort
		start := time.Now()
		sort.Ints(append([]int{}, data...))
		builtinTime := time.Since(start)
		
		// Test custom quicksort
		start = time.Now()
		quickSort(append([]int{}, data...))
		quickSortTime := time.Since(start)
		
		fmt.Printf("Size: %d | Built-in: %v | QuickSort: %v\n", 
			size, builtinTime, quickSortTime)
	}
}

// Memory-efficient sorting for large datasets
func memoryEfficientSort(data []int) {
	// Use in-place sorting to minimize memory usage
	sort.Ints(data)
}

// Chunked sorting for very large datasets
func chunkedSort(data []int, chunkSize int) []int {
	if len(data) <= chunkSize {
		sort.Ints(data)
		return data
	}
	
	// Sort chunks
	var chunks [][]int
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		
		chunk := make([]int, end-i)
		copy(chunk, data[i:end])
		sort.Ints(chunk)
		chunks = append(chunks, chunk)
	}
	
	// Merge chunks
	result := make([]int, 0, len(data))
	for _, chunk := range chunks {
		result = merge(result, chunk)
	}
	
	return result
}
```

### 2. Optimization Techniques

```go
package main

import (
	"fmt"
	"sort"
)

// Optimized sorting for nearly sorted data
func insertionSortOptimized(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Hybrid sort: use insertion sort for small arrays
func hybridSort(arr []int) {
	const threshold = 16
	
	if len(arr) <= threshold {
		insertionSortOptimized(arr)
		return
	}
	
	// Use built-in sort for larger arrays
	sort.Ints(arr)
}

// Adaptive sorting based on data characteristics
func adaptiveSort(arr []int) {
	// Check if already sorted
	if sort.IntsAreSorted(arr) {
		return
	}
	
	// Check if nearly sorted (few inversions)
	inversions := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			inversions++
		}
	}
	
	// If less than 10% inversions, use insertion sort
	if float64(inversions)/float64(len(arr)) < 0.1 {
		insertionSortOptimized(arr)
	} else {
		sort.Ints(arr)
	}
}

func optimizationExample() {
	data := []int{1, 2, 3, 5, 4, 6, 7, 8, 9, 10} // Nearly sorted
	
	fmt.Println("Original:", data)
	
	// Test adaptive sort
	testData := make([]int, len(data))
	copy(testData, data)
	adaptiveSort(testData)
	fmt.Println("Adaptive sort result:", testData)
}
```

---

## Interview Questions

### 1. Basic Questions

**Q1: How do you sort a slice of custom structs in Go?**
```go
// Answer: Implement sort.Interface or use sort.Slice
type Person struct {
	Name string
	Age  int
}

// Method 1: Implement sort.Interface
type ByAge []Person
func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Method 2: Use sort.Slice (simpler)
people := []Person{{"Alice", 25}, {"Bob", 30}}
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

**Q2: What's the difference between sort.Sort() and sort.Stable()?**
```go
// sort.Sort() may be unstable (equal elements may change order)
// sort.Stable() preserves the relative order of equal elements

// Example demonstrating the difference
type Item struct {
	Value int
	Index int
}

items := []Item{{2, 0}, {1, 1}, {2, 2}, {1, 3}}

// Unstable sort - may change order of equal elements
sort.Slice(items, func(i, j int) bool {
    return items[i].Value < items[j].Value
})

// Stable sort - preserves order of equal elements
sort.Stable(sort.Slice(items, func(i, j int) bool {
    return items[i].Value < items[j].Value
}))
```

### 2. Advanced Questions

**Q3: Implement a custom sorting algorithm for a specific use case**
```go
// Sort strings by their length, then alphabetically
type ByLengthThenAlpha []string

func (s ByLengthThenAlpha) Len() int { return len(s) }
func (s ByLengthThenAlpha) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByLengthThenAlpha) Less(i, j int) bool {
    if len(s[i]) != len(s[j]) {
        return len(s[i]) < len(s[j])
    }
    return s[i] < s[j]
}
```

**Q4: How would you sort a large dataset that doesn't fit in memory?**
```go
// External sorting approach:
// 1. Divide data into chunks that fit in memory
// 2. Sort each chunk individually
// 3. Write sorted chunks to disk
// 4. Merge chunks using a priority queue

func externalSort(data []int, chunkSize int) []int {
    // Implementation would involve file I/O and merging
    // This is a simplified in-memory version
    return chunkedSort(data, chunkSize)
}
```

### 3. Problem-Solving Questions

**Q5: Sort an array of numbers where each number is at most k positions away from its correct position**
```go
func sortNearlySorted(arr []int, k int) []int {
    if k == 0 {
        return arr
    }
    
    // Use min-heap for optimal O(n log k) solution
    // Simplified version using insertion sort for demonstration
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        
        // Only need to check last k elements
        start := max(0, i-k)
        for j >= start && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
    
    return arr
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**Q6: Implement a sorting function that can sort by multiple criteria dynamically**
```go
type SortCriterion struct {
	Field string
	Asc   bool
}

func dynamicSort(data []map[string]interface{}, criteria []SortCriterion) {
    sort.Slice(data, func(i, j int) bool {
        for _, criterion := range criteria {
            valI := data[i][criterion.Field]
            valJ := data[j][criterion.Field]
            
            // Compare based on type
            switch v := valI.(type) {
            case int:
                if v != valJ.(int) {
                    if criterion.Asc {
                        return v < valJ.(int)
                    }
                    return v > valJ.(int)
                }
            case string:
                if v != valJ.(string) {
                    if criterion.Asc {
                        return v < valJ.(string)
                    }
                    return v > valJ.(string)
                }
            }
        }
        return false // All criteria equal
    })
}
```

---

## Summary

### Key Takeaways:

1. **Built-in Functions**: Use `sort.Ints()`, `sort.Strings()`, `sort.Float64s()` for basic types
2. **Custom Sorting**: Implement `sort.Interface` or use `sort.Slice()` for custom logic
3. **Stability**: Use `sort.Stable()` when preserving order of equal elements matters
4. **Performance**: Built-in sort is highly optimized (hybrid of quicksort, heapsort, and insertion sort)
5. **Flexibility**: Go's sorting is very flexible with multiple approaches for different needs

### Best Practices:

- Prefer `sort.Slice()` for simple custom sorting
- Implement `sort.Interface` for reusable sorting logic
- Use `sort.Stable()` when order preservation is important
- Consider memory usage for large datasets
- Profile performance when dealing with critical sorting operations

### Time Complexity Reference:
- **Built-in sort**: O(n log n) average, O(n log n) worst case
- **QuickSort**: O(n log n) average, O(n²) worst case
- **MergeSort**: O(n log n) always, O(n) space
- **Insertion Sort**: O(n²) worst case, O(n) best case (nearly sorted)
- **Counting Sort**: O(n + k) where k is range of values
