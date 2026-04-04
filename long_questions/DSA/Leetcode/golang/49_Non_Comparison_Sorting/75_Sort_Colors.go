package main

import (
	"fmt"
	"math"
)

// 75. Sort Colors - Non-Comparison Sorting
// Time: O(N), Space: O(1) for counting sort

// Counting Sort implementation
func sortColorsCounting(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Count occurrences of each color
	counts := make([]int, 3)
	for _, num := range nums {
		counts[num]++
	}
	
	// Reconstruct array
	index := 0
	for color, count := range counts {
		for i := 0; i < count; i++ {
			nums[index] = color
			index++
		}
	}
	
	return nums
}

// Dutch National Flag algorithm (optimized counting sort)
func sortColorsDutchNationalFlag(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Three pointers: low, mid, high
	low := 0
	mid := 0
	high := len(nums) - 1
	
	for mid <= high {
		switch nums[mid] {
		case 0: // Red
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1: // White
			mid++
		case 2: // Blue
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
	
	return nums
}

// Radix Sort implementation
func sortColorsRadix(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Find maximum number to determine number of digits
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	
	// Perform counting sort for each digit
	for exp := 1; maxNum/exp > 0; exp *= 10 {
		nums = countingSortByDigit(nums, exp)
	}
	
	return nums
}

func countingSortByDigit(nums []int, exp int) []int {
	n := len(nums)
	output := make([]int, n)
	count := make([]int, 10)
	
	// Count occurrences of each digit
	for i := 0; i < n; i++ {
		digit := (nums[i] / exp) % 10
		count[digit]++
	}
	
	// Calculate cumulative count
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	
	// Build output array
	for i := n - 1; i >= 0; i-- {
		digit := (nums[i] / exp) % 10
		output[count[digit]-1] = nums[i]
		count[digit]--
	}
	
	// Copy back to original array
	copy(nums, output)
	return nums
}

// Bucket Sort implementation
func sortColorsBucket(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Create buckets for each color
	buckets := make([][]int, 3)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	
	// Distribute numbers into buckets
	for _, num := range nums {
		buckets[num] = append(buckets[num], num)
	}
	
	// Concatenate buckets
	index := 0
	for color := 0; color < 3; color++ {
		for _, num := range buckets[color] {
			nums[index] = num
			index++
		}
	}
	
	return nums
}

// Pigeonhole Sort implementation
func sortColorsPigeonhole(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Find min and max values
	minVal := nums[0]
	maxVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	
	// Create pigeonholes
	rangeSize := maxVal - minVal + 1
	pigeonholes := make([][]int, rangeSize)
	
	// Distribute numbers into pigeonholes
	for _, num := range nums {
		index := num - minVal
		pigeonholes[index] = append(pigeonholes[index], num)
	}
	
	// Collect numbers from pigeonholes
	index := 0
	for i := 0; i < rangeSize; i++ {
		for _, num := range pigeonholes[i] {
			nums[index] = num
			index++
		}
	}
	
	return nums
}

// Flash Sort implementation (simplified)
func sortColorsFlash(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Find min and max
	minVal := nums[0]
	maxVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	
	// Calculate bucket count
	m := int(float64(len(nums)) * 0.43)
	if m < 1 {
		m = 1
	}
	
	// Create buckets
	buckets := make([][]int, m)
	
	// Distribute into buckets
	for _, num := range nums {
		k := int((float64(num-minVal) / float64(maxVal-minVal+1)) * float64(m-1))
		if k < 0 {
			k = 0
		}
		if k >= m {
			k = m - 1
		}
		buckets[k] = append(buckets[k], num)
	}
	
	// Sort each bucket and concatenate
	index := 0
	for i := 0; i < m; i++ {
		if len(buckets[i]) > 0 {
			// Sort bucket using insertion sort
			for j := 1; j < len(buckets[i]); j++ {
				key := buckets[i][j]
				k := j - 1
				for k >= 0 && buckets[i][k] > key {
					buckets[i][k+1] = buckets[i][k]
					k--
				}
				buckets[i][k+1] = key
			}
			
			// Copy to original array
			for _, num := range buckets[i] {
				nums[index] = num
				index++
			}
		}
	}
	
	return nums
}

// Bead Sort implementation (simplified for positive integers)
func sortColorsBead(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Find maximum value
	maxVal := 0
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	
	// Create bead grid
	beads := make([][]int, len(nums))
	for i := range beads {
		beads[i] = make([]int, maxVal)
		for j := 0; j < nums[i]; j++ {
			beads[i][j] = 1
		}
	}
	
	// Let beads fall
	for j := 0; j < maxVal; j++ {
		count := 0
		for i := 0; i < len(nums); i++ {
			if beads[i][j] == 1 {
				count++
			}
		}
		
		for i := 0; i < len(nums); i++ {
			if i < count {
				beads[i][j] = 1
			} else {
				beads[i][j] = 0
			}
		}
	}
	
	// Count beads in each row
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < maxVal; j++ {
			if beads[i][j] == 1 {
				count++
			}
		}
		nums[i] = count
	}
	
	return nums
}

// Gnome Sort implementation
func sortColorsGnome(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	i := 0
	for i < len(nums) {
		if i == 0 || nums[i] >= nums[i-1] {
			i++
		} else {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i--
		}
	}
	
	return nums
}

// Cocktail Sort implementation
func sortColorsCocktail(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	swapped := true
	start := 0
	end := len(nums) - 1
	
	for swapped {
		swapped = false
		
		// Forward pass
		for i := start; i < end; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
		
		if !swapped {
			break
		}
		
		swapped = false
		end--
		
		// Backward pass
		for i := end - 1; i >= start; i-- {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
		
		start++
	}
	
	return nums
}

// Odd-Even Sort implementation
func sortColorsOddEven(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	sorted := false
	for !sorted {
		sorted = true
		
		// Even indices
		for i := 0; i < len(nums)-1; i += 2 {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sorted = false
			}
		}
		
		// Odd indices
		for i := 1; i < len(nums)-1; i += 2 {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sorted = false
			}
		}
	}
	
	return nums
}

// Comb Sort implementation
func sortColorsComb(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	gap := len(nums)
	shrink := 1.3
	sorted := false
	
	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		}
		
		for i := 0; i < len(nums)-gap; i++ {
			if nums[i] > nums[i+gap] {
				nums[i], nums[i+gap] = nums[i+gap], nums[i]
				sorted = false
			}
		}
	}
	
	return nums
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Non-Comparison Sorting for Limited Range Data
- **Counting Sort**: Count occurrences of each value, then reconstruct
- **Dutch National Flag**: Three-way partitioning with single pass
- **Radix Sort**: Sort digit by digit using counting sort
- **Bucket Sort**: Distribute into buckets, sort each bucket

## 2. PROBLEM CHARACTERISTICS
- **Limited Range**: Small number of distinct values (0, 1, 2)
- **Color Sorting**: Sort red, white, blue balls in order
- **In-Place Requirement**: Some variants require O(1) extra space
- **Linear Time**: Can achieve O(N) time with non-comparison sorts

## 3. SIMILAR PROBLEMS
- Sort Colors (LeetCode 75) - Same problem
- Sort an Array with Three Types - Three-way partitioning
- Counting Sort Applications - Frequency-based sorting
- Radix Sort Applications - Multi-digit sorting

## 4. KEY OBSERVATIONS
- **Limited Values**: Only 3 possible values (0, 1, 2)
- **Counting Natural**: Perfect for counting sort
- **Three-Way Partition**: Dutch National Flag algorithm optimal
- **Linear Time Possible**: Non-comparison sorts achieve O(N)

## 5. VARIATIONS & EXTENSIONS
- **Counting Sort**: O(N + k) time, O(k) space
- **Dutch National Flag**: O(N) time, O(1) space
- **Radix Sort**: O(d × (N + k)) time, O(N + k) space
- **Bucket Sort**: O(N + k) time, O(N + k) space

## 6. INTERVIEW INSIGHTS
- Always clarify: "Value range? Space constraints? Stability required?"
- Edge cases: empty array, single element, all same values
- Time complexity: O(N) for counting/DNF, O(N log N) for comparison sorts
- Space complexity: O(1) for DNF, O(k) for counting sort
- Key insight: limited range makes non-comparison sorting optimal

## 7. COMMON MISTAKES
- Wrong counting sort range (should be 0-2 for colors)
- Missing Dutch National Flag edge cases
- Incorrect bucket boundaries
- Wrong radix sort digit processing
- Not handling empty/single element arrays

## 8. OPTIMIZATION STRATEGIES
- **Counting Sort**: O(N) time, O(1) space for fixed range
- **Dutch National Flag**: O(N) time, O(1) space - optimal
- **Radix Sort**: O(N) time for bounded integers
- **Bucket Sort**: O(N) time for uniform distribution

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like sorting colored balls in boxes:**
- You have balls of three colors: red (0), white (1), blue (2)
- You want to arrange them in color order: reds first, then whites, then blues
- You can count how many of each color, then place them accordingly
- Or you can use three boxes and move balls between them
- Like organizing laundry by color: whites, colors, darks

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array containing only 0s, 1s, and 2s
2. **Goal**: Sort array so all 0s come first, then 1s, then 2s
3. **Constraints**: Should be done in-place if possible
4. **Output**: Sorted array in the same order

#### Phase 2: Key Insight Recognition
- **"Limited values"** → Only 3 possible values (0, 1, 2)
- **"Counting natural"** → Can count occurrences of each value
- **"Three-way partition"** → Dutch National Flag algorithm perfect
- **"Linear time possible"** → Non-comparison sorts achieve O(N)

#### Phase 3: Strategy Development
```
Human thought process:
"I need to sort array with only 3 values.
Comparison sort: O(N log N) time, unnecessary overhead.

Counting Sort Approach:
1. Count occurrences of 0, 1, 2
2. Fill array with 0s, then 1s, then 2s
Time: O(N), Space: O(1) for fixed range

Dutch National Flag Approach:
1. Use three pointers: low, mid, high
2. low tracks 0s boundary, high tracks 2s boundary
3. mid processes current element
Time: O(N), Space: O(1) - even better!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return as-is
- **Single element**: Already sorted
- **All same values**: No changes needed
- **Already sorted**: Algorithm handles efficiently

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [2, 0, 2, 1, 1, 0]

Human thinking:
"Dutch National Flag Process:
Step 1: Initialize
low = 0, mid = 0, high = 5
Array: [2, 0, 2, 1, 1, 0]

Step 2: Process nums[mid] = 2 (blue)
Swap nums[mid] with nums[high]
Array: [0, 0, 2, 1, 1, 2]
high = 4

Step 3: Process nums[mid] = 0 (red)
Swap nums[mid] with nums[low]
Array: [0, 0, 2, 1, 1, 2]
low = 1, mid = 1

Step 4: Process nums[mid] = 0 (red)
Swap nums[mid] with nums[low] (same position)
Array: [0, 0, 2, 1, 1, 2]
low = 2, mid = 2

Step 5: Process nums[mid] = 2 (blue)
Swap nums[mid] with nums[high]
Array: [0, 0, 1, 1, 2, 2]
high = 3

Step 6: Process nums[mid] = 1 (white)
mid = 3

Step 7: Process nums[mid] = 1 (white)
mid = 4

Step 8: mid > high, done
Result: [0, 0, 1, 1, 2, 2] ✓"
```

#### Phase 6: Intuition Validation
- **Why counting sort**: Simple and effective for limited range
- **Why Dutch National Flag**: Single pass, O(1) space, optimal
- **Why three pointers**: Separate tracking for each color boundary
- **Why O(N) time**: Each element processed at most once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use quicksort?"** → O(N log N) vs O(N), unnecessary for limited range
2. **"Should I use built-in sort?"** → Comparison sort, not optimal
3. **"What about stability?"** → Not required for this problem
4. **"Can I use extra space?"** → Dutch National Flag is in-place
5. **"Why three pointers?"** → Efficient three-way partitioning

### Real-World Analogy
**Like sorting laundry by color:**
- You have a pile of mixed colored clothes (whites, colors, darks)
- You want to separate them into three piles by color
- You can count each type first, then create separate piles
- Or you can use three baskets and move clothes between them
- Like doing laundry: whites first, then colors, then darks

### Human-Readable Pseudocode
```
function dutchNationalFlag(nums):
    low = 0
    mid = 0
    high = len(nums) - 1
    
    while mid <= high:
        if nums[mid] == 0:  # Red
            swap(nums[low], nums[mid])
            low += 1
            mid += 1
        elif nums[mid] == 1:  # White
            mid += 1
        else:  # Blue (nums[mid] == 2)
            swap(nums[mid], nums[high])
            high -= 1
    
    return nums
```

### Execution Visualization

### Example: nums = [2, 0, 2, 1, 1, 0]
```
Initial: [2, 0, 2, 1, 1, 0]
         low=0, mid=0, high=5

Step 1: nums[mid]=2 (blue)
Swap mid and high: [0, 0, 2, 1, 1, 2]
high=4

Step 2: nums[mid]=0 (red)
Swap low and mid: [0, 0, 2, 1, 1, 2]
low=1, mid=1

Step 3: nums[mid]=0 (red)
Swap low and mid: [0, 0, 2, 1, 1, 2]
low=2, mid=2

Step 4: nums[mid]=2 (blue)
Swap mid and high: [0, 0, 1, 1, 2, 2]
high=3

Step 5: nums[mid]=1 (white)
mid=3

Step 6: nums[mid]=1 (white)
mid=4

Step 7: mid > high, done
Result: [0, 0, 1, 1, 2, 2] ✓
```

### Key Visualization Points:
- **Three Regions**: [0..low-1] = 0s, [low..mid-1] = 1s, [high+1..end] = 2s
- **Processing**: mid pointer processes unsorted region [mid..high]
- **Swaps**: 0s move left, 2s move right, 1s stay in middle
- **Termination**: When mid crosses high, all elements sorted

### Dutch National Flag Algorithm Visualization:
```
[0, 0, 0, | 1, 1, ?, ?, | 2, 2, 2]
 ^low-1    ^mid-1   ^mid   ^high   ^high+1
  0s      1s      unsorted   2s
```

### Time Complexity Breakdown:
- **Counting Sort**: O(N) time, O(1) space for fixed range
- **Dutch National Flag**: O(N) time, O(1) space - optimal
- **Radix Sort**: O(N) time for bounded integers
- **Bucket Sort**: O(N) time for uniform distribution

### Alternative Approaches:

#### 1. Counting Sort (O(N) time, O(1) space)
```go
func sortColorsCounting(nums []int) []int {
    counts := [3]int{0, 0, 0}
    
    // Count occurrences
    for _, num := range nums {
        counts[num]++
    }
    
    // Reconstruct array
    index := 0
    for color := 0; color < 3; color++ {
        for i := 0; i < counts[color]; i++ {
            nums[index] = color
            index++
        }
    }
    
    return nums
}
```
- **Pros**: Simple, stable, easy to understand
- **Cons**: Two passes, requires counting array

#### 2. Two-Pass Counting (O(N) time, O(1) space)
```go
func sortColorsTwoPass(nums []int) []int {
    // First pass: count
    zeros, ones, twos := 0, 0, 0
    for _, num := range nums {
        switch num {
        case 0: zeros++
        case 1: ones++
        case 2: twos++
        }
    }
    
    // Second pass: fill
    for i := 0; i < zeros; i++ {
        nums[i] = 0
    }
    for i := zeros; i < zeros+ones; i++ {
        nums[i] = 1
    }
    for i := zeros+ones; i < len(nums); i++ {
        nums[i] = 2
    }
    
    return nums
}
```
- **Pros**: No extra array, clear logic
- **Cons**: Two passes, more code than Dutch National Flag

#### 3. Bucket Sort (O(N) time, O(N) space)
```go
func sortColorsBucket(nums []int) []int {
    buckets := make([][]int, 3)
    
    // Distribute
    for _, num := range nums {
        buckets[num] = append(buckets[num], num)
    }
    
    // Collect
    index := 0
    for color := 0; color < 3; color++ {
        for _, num := range buckets[color] {
            nums[index] = num
            index++
        }
    }
    
    return nums
}
```
- **Pros**: Extensible to more values
- **Cons**: More memory, overhead for small range

### Extensions for Interviews:
- **General Dutch National Flag**: Sort k different values
- **Stable Sorting**: Maintain relative order within groups
- **Counting Sort Variants**: Handle larger ranges, negative numbers
- **Radix Sort Applications**: Multi-digit numbers, strings
- **Real-world Applications**: Color sorting, grade classification, data partitioning
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Non-Comparison Sorting ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{2, 0, 2, 1, 1, 0}, "Standard case"},
		{[]int{2, 0, 1}, "Small case"},
		{[]int{0, 0, 0, 0}, "All zeros"},
		{[]int{1, 1, 1, 1}, "All ones"},
		{[]int{2, 2, 2, 2}, "All twos"},
		{[]int{0, 1, 2, 0, 1, 2}, "Already sorted"},
		{[]int{2, 1, 0, 2, 1, 0}, "Reverse sorted"},
		{[]int{1, 0, 1, 0, 1, 0}, "Alternating"},
		{[]int{0}, "Single element"},
		{[]int{}, "Empty array"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Original: %v\n", tc.nums)
		
		// Test counting sort
		countingResult := make([]int, len(tc.nums))
		copy(countingResult, tc.nums)
		sortColorsCounting(countingResult)
		fmt.Printf("  Counting sort: %v\n", countingResult)
		
		// Test Dutch National Flag
		dutchResult := make([]int, len(tc.nums))
		copy(dutchResult, tc.nums)
		sortColorsDutchNationalFlag(dutchResult)
		fmt.Printf("  Dutch National Flag: %v\n", dutchResult)
		
		// Test radix sort
		radixResult := make([]int, len(tc.nums))
		copy(radixResult, tc.nums)
		sortColorsRadix(radixResult)
		fmt.Printf("  Radix sort: %v\n", radixResult)
		
		// Test bucket sort
		bucketResult := make([]int, len(tc.nums))
		copy(bucketResult, tc.nums)
		sortColorsBucket(bucketResult)
		fmt.Printf("  Bucket sort: %v\n", bucketResult)
		
		// Test pigeonhole sort
		pigeonholeResult := make([]int, len(tc.nums))
		copy(pigeonholeResult, tc.nums)
		sortColorsPigeonhole(pigeonholeResult)
		fmt.Printf("  Pigeonhole sort: %v\n", pigeonholeResult)
		
		// Test flash sort
		flashResult := make([]int, len(tc.nums))
		copy(flashResult, tc.nums)
		sortColorsFlash(flashResult)
		fmt.Printf("  Flash sort: %v\n", flashResult)
		
		// Test bead sort
		beadResult := make([]int, len(tc.nums))
		copy(beadResult, tc.nums)
		sortColorsBead(beadResult)
		fmt.Printf("  Bead sort: %v\n", beadResult)
		
		// Test gnome sort
		gnomeResult := make([]int, len(tc.nums))
		copy(gnomeResult, tc.nums)
		sortColorsGnome(gnomeResult)
		fmt.Printf("  Gnome sort: %v\n", gnomeResult)
		
		// Test cocktail sort
		cocktailResult := make([]int, len(tc.nums))
		copy(cocktailResult, tc.nums)
		sortColorsCocktail(cocktailResult)
		fmt.Printf("  Cocktail sort: %v\n", cocktailResult)
		
		// Test odd-even sort
		oddEvenResult := make([]int, len(tc.nums))
		copy(oddEvenResult, tc.nums)
		sortColorsOddEven(oddEvenResult)
		fmt.Printf("  Odd-even sort: %v\n", oddEvenResult)
		
		// Test comb sort
		combResult := make([]int, len(tc.nums))
		copy(combResult, tc.nums)
		sortColorsComb(combResult)
		fmt.Printf("  Comb sort: %v\n", combResult)
		
		fmt.Println()
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	// Large array
	largeNums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		largeNums[i] = i % 3
	}
	
	fmt.Printf("Large array with %d elements\n", len(largeNums))
	
	start := time.Now()
	countingLarge := make([]int, len(largeNums))
	copy(countingLarge, largeNums)
	sortColorsCounting(countingLarge)
	countingDuration := time.Since(start)
	
	start = time.Now()
	dutchLarge := make([]int, len(largeNums))
	copy(dutchLarge, largeNums)
	sortColorsDutchNationalFlag(dutchLarge)
	dutchDuration := time.Since(start)
	
	start = time.Now()
	radixLarge := make([]int, len(largeNums))
	copy(radixLarge, largeNums)
	sortColorsRadix(radixLarge)
	radixDuration := time.Since(start)
	
	fmt.Printf("Counting sort: %v\n", countingDuration)
	fmt.Printf("Dutch National Flag: %v\n", dutchDuration)
	fmt.Printf("Radix sort: %v\n", radixDuration)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single element
	single := []int{1}
	fmt.Printf("Single element: %v\n", sortColorsCounting(single))
	
	// Already sorted
	sorted := []int{0, 0, 1, 1, 2, 2}
	fmt.Printf("Already sorted: %v\n", sortColorsCounting(sorted))
	
	// Reverse sorted
	reverse := []int{2, 2, 1, 1, 0, 0}
	fmt.Printf("Reverse sorted: %v\n", sortColorsCounting(reverse))
	
	// All same
	allSame := []int{1, 1, 1, 1, 1}
	fmt.Printf("All same: %v\n", sortColorsCounting(allSame))
	
	// Large values
	largeValues := []int{100, 200, 300, 100, 200, 300}
	fmt.Printf("Large values: %v\n", sortColorsCounting(largeValues))
	
	// Test space complexity
	fmt.Println("\n=== Space Complexity Test ===")
	
	spaceTest := []int{2, 0, 2, 1, 1, 0}
	fmt.Printf("Original: %v\n", spaceTest)
	
	// Counting sort uses O(k) space where k is range
	countingSpace := make([]int, len(spaceTest))
	copy(countingSpace, spaceTest)
	sortColorsCounting(countingSpace)
	fmt.Printf("Counting sort (O(k) space): %v\n", countingSpace)
	
	// Dutch National Flag uses O(1) space
	dutchSpace := make([]int, len(spaceTest))
	copy(dutchSpace, spaceTest)
	sortColorsDutchNationalFlag(dutchSpace)
	fmt.Printf("Dutch National Flag (O(1) space): %v\n", dutchSpace)
	
	// Test time complexity
	fmt.Println("\n=== Time Complexity Test ===")
	
	// Test with different sizes
	sizes := []int{100, 1000, 10000}
	
	for _, size := range sizes {
		testArray := make([]int, size)
		for i := 0; i < size; i++ {
			testArray[i] = i % 3
		}
		
		start = time.Now()
		sortColorsCounting(testArray)
		duration = time.Since(start)
		
		fmt.Printf("Size %d: %v\n", size, duration)
	}
	
	// Test stability
	fmt.Println("\n=== Stability Test ===")
	
	// Create array with additional data to test stability
	type ColorItem struct {
		color int
		index int
	}
	
	stableTest := []ColorItem{
		{2, 0}, {0, 1}, {2, 2}, {1, 3}, {1, 4}, {0, 5},
	}
	
	fmt.Printf("Original stable test: %v\n", stableTest)
	
	// Extract colors and sort
	colors := make([]int, len(stableTest))
	for i, item := range stableTest {
		colors[i] = item.color
	}
	
	sortColorsCounting(colors)
	fmt.Printf("Sorted colors: %v\n", colors)
	
	// Test with negative numbers (should handle gracefully)
	fmt.Println("\n=== Negative Numbers Test ===")
	
	negativeTest := []int{-1, 0, 1, -1, 0, 1}
	fmt.Printf("Negative numbers: %v\n", negativeTest)
	
	// Note: counting sort for colors assumes non-negative integers
	// For negative numbers, we'd need to adjust the algorithm
}
