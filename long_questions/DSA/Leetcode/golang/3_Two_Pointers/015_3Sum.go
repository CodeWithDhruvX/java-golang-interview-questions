package main

import (
	"fmt"
	"sort"
)

// 15. 3Sum
// Time: O(N^2), Space: O(1) (ignoring output space)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	
	for i := 0; i < n-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		// Two pointers approach for the remaining two elements
		left, right := i+1, n-1
		target := -nums[i]
		
		for left < right {
			sum := nums[left] + nums[right]
			
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				
				// Skip duplicates for the second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for the third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Sorting + Two Pointers + Deduplication
- **Sorting**: Enables two-pointer approach and duplicate handling
- **Fixed Element**: Iterate through array, fixing one element
- **Two Pointers**: Find remaining two elements that sum to target
- **Deduplication**: Skip duplicate elements to avoid duplicate triplets

## 2. PROBLEM CHARACTERISTICS
- **Triplet Finding**: Need three numbers that sum to zero
- **No Duplicates**: Each triplet must be unique
- **Order Independence**: [a,b,c] is same as [c,b,a]
- **Multiple Solutions**: Can have multiple valid triplets

## 3. SIMILAR PROBLEMS
- Two Sum (LeetCode 1)
- Two Sum II (LeetCode 167)
- 3Sum Closest (LeetCode 16)
- 4Sum (LeetCode 18)

## 4. KEY OBSERVATIONS
- **Sorting helps**: Enables two-pointer technique and duplicate detection
- **Target transformation**: For fixed element a, find b + c = -a
- **Duplicate elimination**: Skip same values at all three positions
- **Early termination**: Can skip when fixed element is too large

## 5. VARIATIONS & EXTENSIONS
- **3Sum Closest**: Find triplet closest to target
- **4Sum**: Extend to four numbers
- **K-Sum**: Generalize to K numbers
- **Different Target**: Sum to specific value instead of zero

## 6. INTERVIEW INSIGHTS
- Always clarify: "Should output be sorted? Can input have duplicates?"
- Edge cases: empty array, less than 3 elements, all zeros
- Space complexity: O(1) extra space (ignoring output)
- Time complexity: O(N²) after sorting

## 7. COMMON MISTAKES
- Not handling duplicates properly
- Using O(N³) brute force approach
- Not sorting first (two-pointer requires sorted input)
- Forgetting to skip duplicates for all three positions
- Not handling edge cases (empty array, etc.)

## 8. OPTIMIZATION STRATEGIES
- **Early termination**: Skip when fixed element > 0 (can't sum to zero)
- **Duplicate skipping**: Avoid processing same values multiple times
- **Pruning**: Skip impossible ranges based on sorted order
- **Hash optimization**: Use hash set for duplicate detection

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding three friends whose ages sum to zero:**
- You have a list of people's ages (can be negative)
- You want to find groups of three where ages sum to zero
- First, sort everyone by age (helps with organization)
- For each person, you need to find two others whose ages cancel out theirs
- Use two pointers to efficiently find the right pair
- Make sure not to count the same group multiple times

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers (can be positive, negative, zero)
2. **Goal**: Find all unique triplets that sum to zero
3. **Output**: List of unique triplets (order within triplet doesn't matter)
4. **Constraint**: No duplicate triplets in output

#### Phase 2: Key Insight Recognition
- **"Sorting advantage"** → Enables two-pointer technique and easy duplicate detection
- **"Fix and find"** → Fix one element, find pair for remaining sum
- **"Two-pointer efficiency"** → Find pairs in O(N) instead of O(N²)
- **"Deduplication necessity"** → Must skip duplicates at all levels

#### Phase 3: Strategy Development
```
Human thought process:
"If I sort the array first, I can use two pointers to find pairs efficiently.
For each element nums[i], I need to find two other numbers that sum to -nums[i].
I'll use left and right pointers to search for this pair.
When I find a valid triplet, I'll add it to results.
Then I'll skip any duplicates to avoid duplicate triplets.
This way, I can find all unique triplets in O(N²) time."
```

#### Phase 4: Edge Case Handling
- **Less than 3 elements**: Return empty array
- **All zeros**: Return [[0,0,0]] if at least 3 zeros
- **No solution**: Return empty array
- **Large arrays**: Early termination when possible

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [-1, 0, 1, 2, -1, -4]
After sorting: [-4, -1, -1, 0, 1, 2]

Human thinking:
"I'll go through each element and find pairs:

i=0, nums[0] = -4:
   Need pair summing to 4
   left=1(-1), right=5(2): sum = 1 < 4, move left
   left=2(-1), right=5(2): sum = 1 < 4, move left  
   left=3(0), right=5(2): sum = 2 < 4, move left
   left=4(1), right=5(2): sum = 3 < 4, move left
   left=5, left >= right → no pairs found

i=1, nums[1] = -1:
   Need pair summing to 1
   left=2(-1), right=5(2): sum = 1 → FOUND [-1,-1,2]
   Skip duplicates: left moves to 3, right moves to 4
   left=3(0), right=4(1): sum = 1 → FOUND [-1,0,1]
   Skip duplicates: left moves to 4, right moves to 3
   left >= right → done with i=1

i=2, nums[2] = -1:
   Same as previous, skip (duplicate)

i=3, nums[3] = 0:
   Need pair summing to 0
   left=4(1), right=5(2): sum = 3 > 0, move right
   left=4, right=4 → done

i=4, nums[4] = 1:
   Need pair summing to -1, but left=5 > right=4 → done

Final result: [[-1,-1,2], [-1,0,1]]"
```

#### Phase 6: Intuition Validation
- **Why sorting works**: Enables two-pointer technique and duplicate detection
- **Why two pointers**: Efficiently finds pairs in O(N) time
- **Why deduplication works**: Skipping duplicates prevents repeated triplets
- **Why O(N²)**: Outer loop O(N), inner two-pointer O(N)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all triplets?"** → That's O(N³), too slow
2. **"Do I really need to sort?"** → Yes, two-pointer requires sorted input
3. **"What about duplicate handling?"** → Must skip duplicates at all three positions
4. **"Can I use hash instead?"** → Possible but more complex for deduplication

### Real-World Analogy
**Like forming balanced teams of three:**
- You have players with different skill levels (positive/negative)
- You want teams where total skill sums to zero (balanced)
- Sort players by skill level
- For each player, find two others that balance them out
- Use two pointers to efficiently search from both ends
- Make sure not to form the same team multiple times

### Human-Readable Pseudocode
```
function threeSum(numbers):
    sort(numbers)
    result = []
    
    for i from 0 to length(numbers)-3:
        if i > 0 and numbers[i] == numbers[i-1]:
            continue  // skip duplicate first element
            
        left = i + 1
        right = length(numbers) - 1
        target = -numbers[i]
        
        while left < right:
            sum = numbers[left] + numbers[right]
            
            if sum == target:
                add [numbers[i], numbers[left], numbers[right]] to result
                
                // skip duplicates
                while left < right and numbers[left] == numbers[left+1]:
                    left++
                while left < right and numbers[right] == numbers[right-1]:
                    right--
                    
                left++
                right--
            else if sum < target:
                left++
            else:
                right--
    
    return result
```

### Execution Visualization

### Example: [-1, 0, 1, 2, -1, -4]
```
Input: [-1, 0, 1, 2, -1, -4]
Sorted: [-4, -1, -1, 0, 1, 2]

=== i=0, nums[0]=-4, target=4 ===
left=1(-1), right=5(2): sum=1 < 4 → left=2
left=2(-1), right=5(2): sum=1 < 4 → left=3
left=3(0), right=5(2): sum=2 < 4 → left=4
left=4(1), right=5(2): sum=3 < 4 → left=5
left >= right → no result

=== i=1, nums[1]=-1, target=1 ===
left=2(-1), right=5(2): sum=1 → FOUND [-1,-1,2]
Skip duplicates: left=3, right=4
left=3(0), right=4(1): sum=1 → FOUND [-1,0,1]
Skip duplicates: left=4, right=3
left >= right → done

=== i=2, nums[2]=-1 ===
Duplicate of nums[1], skip

=== Continue... ===
Final result: [[-1,-1,2], [-1,0,1]]
```

### Key Visualization Points:
- **Sorting**: [-4, -1, -1, 0, 1, 2] enables two-pointer
- **Fixed element**: nums[i] is fixed, find pair for -nums[i]
- **Two pointers**: left and right move inward based on sum comparison
- **Duplicate skipping**: Prevents duplicate triplets

### Memory Layout Visualization:
```
Sorted:  [-4][-1][-1][0][1][2]
Index:    0   1   2  3 4  5
          ^        ^     ^
         i=0     left=1 right=5
          sum=-1+2=1 < target=4
```

### Time Complexity Breakdown:
- **Sorting**: O(N log N)
- **Outer Loop**: O(N) - iterate through each element
- **Inner Two-Pointer**: O(N) - each element visited at most once
- **Total**: O(N²) - dominated by nested loops
- **Space**: O(1) extra space (ignoring output storage)

### Alternative Approaches:

#### 1. Brute Force (O(N³))
```go
func threeSum(nums []int) [][]int {
    result := [][]int{}
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                if nums[i] + nums[j] + nums[k] == 0 {
                    triplet := []int{nums[i], nums[j], nums[k]}
                    sort.Ints(triplet)
                    // Add to result with duplicate checking...
                }
            }
        }
    }
    return result
}
```
- **Pros**: Simple, finds all combinations
- **Cons**: O(N³) time complexity

#### 2. Hash Set Approach (O(N²))
```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        seen := make(map[int]int)
        target := -nums[i]
        
        for j := i + 1; j < len(nums); j++ {
            complement := target - nums[j]
            if idx, exists := seen[complement]; exists {
                result = append(result, []int{nums[i], complement, nums[j]})
                // Skip duplicates...
            }
            seen[nums[j]] = j
        }
    }
    return result
}
```
- **Pros**: Still O(N²), different approach
- **Cons**: More complex duplicate handling

### Extensions for Interviews:
- **3Sum Closest**: Find triplet closest to target
- **4Sum**: Extend approach to four numbers
- **K-Sum**: Generalize using recursion
- **Count Triplets**: Count instead of returning actual triplets
*/
func main() {
	// Test cases
	testCases := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
		{-2, 0, 1, 1, 2},
		{-1, -2, -3, -4, -5},
		{1, 2, -2, -1},
		{3, -2, 1, 0, -1, 2, -3},
		{},
		{0},
	}
	
	for i, nums := range testCases {
		result := threeSum(nums)
		fmt.Printf("Test Case %d: %v -> Triplets: %v\n", i+1, nums, result)
	}
}
