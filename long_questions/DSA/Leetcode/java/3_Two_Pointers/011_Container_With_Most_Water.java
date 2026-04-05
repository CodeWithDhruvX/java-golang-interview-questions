import java.util.Arrays;

public class ContainerWithMostWater {
    
    // 11. Container With Most Water
    // Time: O(N), Space: O(1)
    public static int maxArea(int[] height) {
        int left = 0, right = height.length - 1;
        int maxWater = 0;
        
        while (left < right) {
            // Calculate current area
            int width = right - left;
            int height1 = height[left];
            int height2 = height[right];
            int currentHeight = Math.min(height1, height2);
            
            int currentArea = width * currentHeight;
            maxWater = Math.max(maxWater, currentArea);
            
            // Move the pointer with smaller height
            if (height1 < height2) {
                left++;
            } else {
                right--;
            }
        }
        
        return maxWater;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 8, 6, 2, 5, 4, 8, 3, 7},
            {1, 1},
            {4, 3, 2, 1, 4},
            {1, 2, 1},
            {2, 3, 4, 5, 18, 17, 6},
            {1, 2, 3, 4, 5},
            {5, 4, 3, 2, 1},
            {1, 3, 2, 5, 25, 24, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = maxArea(testCases[i]);
            System.out.printf("Test Case %d: %s -> Max Area: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers with Greedy Selection
- **Two Pointers**: Start at both ends of array
- **Area Calculation**: width × min(height[left], height[right])
- **Greedy Move**: Move pointer with smaller height
- **Maximum Tracking**: Keep track of best area found

## 2. PROBLEM CHARACTERISTICS
- **Height Array**: Vertical lines at different positions
- **Container Formation**: Two lines form container with x-axis
- **Area Calculation**: Width × height (limited by shorter line)
- **Goal**: Find maximum possible area

## 3. SIMILAR PROBLEMS
- Trapping Rain Water
- Largest Rectangle in Histogram
- Max Consecutive Ones
- Two Sum variations

## 4. KEY OBSERVATIONS
- Area limited by shorter of two lines
- Moving taller line inward cannot increase area
- Moving shorter line might find taller line
- Width decreases as pointers move inward

## 5. VARIATIONS & EXTENSIONS
- Find indices of maximum area
- Handle circular container
- Multiple queries on same array
- 3D version (container with volume)

## 6. INTERVIEW INSIGHTS
- Clarify: "Can lines have zero height?"
- Edge cases: empty array, single element, all same height
- Why greedy approach works: mathematical proof
- Time complexity: O(N) vs O(N²) brute force

## 7. COMMON MISTAKES
- Using max height instead of min height
- Moving wrong pointer (taller instead of shorter)
- Not handling edge cases properly
- Forgetting to update maximum area

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal
- For multiple queries, consider preprocessing
- For streaming data, use sliding window
- For memory constraints, use in-place processing

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like choosing two walls for a pool:**
- You have walls of different heights at different positions
- You want to build the largest possible pool between two walls
- The water height is limited by the shorter wall
- To get more water, you need taller walls or wider distance

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of heights representing vertical lines
2. **Goal**: Find two lines that form maximum area container
3. **Output**: Maximum area value

#### Phase 2: Key Insight Recognition
- **"Area depends on shorter line!"** → min(height[left], height[right])
- **"Width decreases as we move inward!"** → right - left
- **"Which line to move?"** → The shorter one!
- **"Why move shorter?"** → Might find taller line

#### Phase 3: Strategy Development
```
Human thought process:
"I'll start with the widest possible container (first and last lines).
I'll calculate the area and remember the maximum.
Then I'll move the shorter wall inward, hoping to find a taller one.
I'll keep doing this until the walls meet."
```

#### Phase 4: Edge Case Handling
- **Empty array**: No container possible, return 0
- **Single element**: Cannot form container, return 0
- **Two elements**: Only one possible container
- **All zero heights**: Area will be 0

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [1,8,6,2,5,4,8,3,7]

Human thinking:
"Start with left=0 (height=1) and right=7 (height=7)
Area = 7 × min(1,7) = 7 × 1 = 7
Max area so far = 7

Left is shorter (1 < 7), so move left to 1
Now left=1 (height=8), right=7 (height=7)
Area = 6 × min(8,7) = 6 × 7 = 42
Max area so far = 42

Right is shorter (7 < 8), so move right to 6
Continue this process..."
```

#### Phase 6: Intuition Validation
- **Why it works**: We never miss optimal pair
- **Why it's efficient**: Single pass with two pointers
- **Why it's correct**: Greedy choice is mathematically optimal

### Common Human Pitfalls & How to Avoid Them
1. **"Why not move both pointers?"** → Only one move per iteration
2. **"Which pointer to move?"** → Always the shorter one
3. **"Can I use brute force?"** → Too slow O(N²)
4. **"What about equal heights?"** → Move either pointer

### Real-World Analogy
**Like finding the best spot for a bridge:**
- You have cliffs of different heights along a river
- You want to build the longest possible bridge between two cliffs
- Bridge height is limited by the shorter cliff
- You start with the widest span and adjust based on height

### Human-Readable Pseudocode
```
function maxContainerArea(heights):
    left = 0
    right = heights.length - 1
    maxArea = 0
    
    while left < right:
        width = right - left
        currentHeight = min(heights[left], heights[right])
        currentArea = width * currentHeight
        maxArea = max(maxArea, currentArea)
        
        if heights[left] < heights[right]:
            left++
        else:
            right--
    
    return maxArea
```

### Execution Visualization

### Example: [1,8,6,2,5,4,8,3,7]
```
Initial: heights = [1,8,6,2,5,4,8,3,7]
         left = 0 (1), right = 7 (7)
         maxArea = 0

Step 1: left=0, right=7
→ width = 7, height = min(1,7) = 1
→ area = 7 × 1 = 7
→ maxArea = 7
→ left is shorter, left = 1

Step 2: left=1, right=7
→ width = 6, height = min(8,7) = 7
→ area = 6 × 7 = 42
→ maxArea = 42
→ right is shorter, right = 6

Step 3: left=1, right=6
→ width = 5, height = min(8,3) = 3
→ area = 5 × 3 = 15
→ maxArea = 42 (unchanged)
→ right is shorter, right = 5

Continue until left >= right...
Final maxArea = 49
```

### Key Visualization Points:
- **Two pointers** start at extremes and move inward
- **Area calculation** uses minimum height
- **Greedy choice**: move shorter pointer
- **Maximum tracking** updates on each iteration

### Memory Layout Visualization:
```
Array: [1][8][6][2][5][4][8][3][7]
        ^                   ^
        |                   |
      left=0               right=7
      height=1            height=7
      width=7             area=7

After move:
Array: [1][8][6][2][5][4][8][3][7]
           ^                   ^
           |                   |
         left=1             right=7
         height=8            height=7
         width=6             area=42 ✓
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited at most once
- **Constant Space**: O(1) - only a few variables
- **Optimal**: Cannot do better than O(N) for this problem
- **Greedy Proof**: Mathematically optimal strategy
*/
