import java.util.*;

public class DailyTemperatures {
    
    // 739. Daily Temperatures
    // Time: O(N), Space: O(N)
    public static int[] dailyTemperatures(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        Deque<Integer> stack = new ArrayDeque<>(); // Store indices
        
        for (int i = 0; i < n; i++) {
            // While current temperature is greater than temperature at stack top
            while (!stack.isEmpty() && temperatures[i] > temperatures[stack.peek()]) {
                int prevIndex = stack.pop();
                result[prevIndex] = i - prevIndex;
            }
            stack.push(i);
        }
        
        // Remaining indices in stack have no warmer day (result is already 0)
        return result;
    }

    // Alternative approach using reverse iteration
    public static int[] dailyTemperaturesReverse(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        Deque<Integer> stack = new ArrayDeque<>(); // Store indices
        
        // Process from right to left
        for (int i = n - 1; i >= 0; i--) {
            // Remove indices that are not warmer than current temperature
            while (!stack.isEmpty() && temperatures[i] >= temperatures[stack.peek()]) {
                stack.pop();
            }
            
            if (!stack.isEmpty()) {
                result[i] = stack.peek() - i;
            }
            
            stack.push(i);
        }
        
        return result;
    }

    // Brute force approach for comparison (O(N^2))
    public static int[] dailyTemperaturesBruteForce(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (temperatures[j] > temperatures[i]) {
                    result[i] = j - i;
                    break;
                }
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {73, 74, 75, 71, 69, 72, 76, 73},
            {30, 40, 50, 60},
            {30, 60, 90},
            {90, 80, 70, 60},
            {55, 60, 65, 70, 75},
            {65, 70, 65, 60, 65},
            {50},
            {},
            {73, 73, 73, 73},
            {30, 40, 30, 50, 30, 60, 30}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] result1 = dailyTemperatures(testCases[i]);
            int[] result2 = dailyTemperaturesReverse(testCases[i]);
            int[] result3 = dailyTemperaturesBruteForce(testCases[i]);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Forward: %s\n", Arrays.toString(result1));
            System.out.printf("  Reverse: %s\n", Arrays.toString(result2));
            System.out.printf("  Brute:   %s\n\n", Arrays.toString(result3));
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Monotonic Stack for Next Greater Element
- **Monotonic Stack**: Stack maintains decreasing temperature indices
- **Next Greater**: Find first warmer day for each temperature
- **Distance Calculation**: Current index - previous greater index
- **Stack Operations**: Push new indices, pop smaller temperatures

## 2. PROBLEM CHARACTERISTICS
- **Temperature Array**: Daily temperature readings
- **Next Warmer Day**: Find first day with higher temperature
- **Distance Calculation**: Number of days to wait
- **Zero Default**: No warmer day means distance = 0

## 3. SIMILAR PROBLEMS
- Next Greater Element I & II
- Largest Rectangle in Histogram
- Trapping Rain Water
- Stock Buy/Sell problems

## 4. KEY OBSERVATIONS
- Need to find next greater element to the right
- Stack helps maintain candidates for "next greater"
- Monotonic decreasing stack: temperatures[stack[i]] > temperatures[stack[i+1]]
- When current temperature > stack top, we found answer for stack top

## 5. VARIATIONS & EXTENSIONS
- Find next smaller element (reverse comparison)
- Handle circular array (wrap around)
- Multiple queries on same temperature array
- Find k-th next greater element

## 6. INTERVIEW INSIGHTS
- Clarify: "Are temperatures guaranteed to be integers?"
- Edge cases: empty array, single element, all same temperature
- Why stack works better than nested loops
- Time complexity: O(N) vs O(N²) brute force

## 7. COMMON MISTAKES
- Using nested loops (O(N²) solution)
- Not handling stack empty case properly
- Wrong comparison direction (should be > not <)
- Forgetting to initialize result array with zeros

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal
- For multiple queries, consider preprocessing
- For streaming data, maintain running stack
- For memory constraints, use array instead of stack

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like waiting for a warmer day:**
- You're checking weather forecast each day
- For each day, you want to know when it will get warmer
- You keep a list of days that haven't found warmer weather yet
- When a hot day comes, it resolves the wait for all cooler previous days

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of daily temperatures
2. **Goal**: For each day, find days until warmer temperature
3. **Output**: Array of wait times (0 if no warmer day exists)

#### Phase 2: Key Insight Recognition
- **"Need next greater element!"** → Classic problem
- **"How to find efficiently?"** → Stack to track candidates
- **"What makes a good candidate?"** → Decreasing temperatures
- **"When to resolve?"** → Current temperature > stack top

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use a stack to store indices of temperatures
that haven't found a warmer day yet.
The stack will maintain decreasing temperatures.

For each day, I'll:
1. Pop all days that are cooler than current temperature
2. For each popped day, current day is their warmer day
3. Push current day onto the stack"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single element**: Return [0]
- **All same temperatures**: All zeros (no warmer days)
- **Decreasing sequence**: Each day waits for next warmer day

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [73, 74, 75, 71, 69, 72, 76, 73]

Human thinking:
"Initialize stack = [], result = [0,0,0,0,0,0,0,0,0]

Day 0 (73°C):
→ Stack is empty, push 0
→ Stack: [0], result: [0,0,0,0,0,0,0,0,0]

Day 1 (74°C):
→ 74 > 73 (stack top), so day 0 waits 1 day
→ Pop 0, set result[0] = 1-0 = 1
→ Stack is empty, push 1
→ Stack: [1], result: [1,0,0,0,0,0,0,0,0]

Day 2 (75°C):
→ 75 > 74 (stack top), so day 1 waits 1 day
→ Pop 1, set result[1] = 2-1 = 1
→ Stack is empty, push 2
→ Stack: [2], result: [1,1,0,0,0,0,0,0,0]

Day 3 (71°C):
→ 71 ≤ 75 (stack top), no pop
→ Push 3
→ Stack: [2,3], result: [1,1,0,0,0,0,0,0,0]

Continue this process..."
```

#### Phase 6: Intuition Validation
- **Why it works**: Stack maintains candidates in decreasing order
- **Why it's efficient**: Each index pushed and popped at most once
- **Why it's correct**: When we pop, we've found the first warmer day

### Common Human Pitfalls & How to Avoid Them
1. **"Why not nested loops?"** → Too slow O(N²)
2. **"What stack order?"** → Must be decreasing temperatures
3. **"When to pop?"** → Current temperature > stack top temperature
4. **"What about equal temperatures?"** → Don't pop (need strictly greater)

### Real-World Analogy
**Like waiting for a bus with fewer passengers:**
- You're tracking bus capacity each day (temperatures)
- You want to know when you'll find a less crowded bus
- You keep a list of days still waiting for better buses
- When a good bus arrives, everyone waiting gets on board
- The distance shows how long they waited

### Human-Readable Pseudocode
```
function dailyTemperatures(temps):
    result = array of zeros same length as temps
    stack = empty stack  // stores indices
    
    for i from 0 to temps.length-1:
        while stack not empty and temps[i] > temps[stack.top()]:
            prevIndex = stack.pop()
            result[prevIndex] = i - prevIndex
        
        stack.push(i)
    
    return result
```

### Execution Visualization

### Example: [73, 74, 75, 71, 69, 72, 76, 73]
```
Initial: temps = [73, 74, 75, 71, 69, 72, 76, 73]
stack = [], result = [0,0,0,0,0,0,0,0,0]

Day 0: temp=73
→ stack empty, push 0
→ stack=[0], result=[0,0,0,0,0,0,0,0,0]

Day 1: temp=74
→ 74 > 73, pop 0, result[0]=1-0=1
→ stack empty, push 1
→ stack=[1], result=[1,0,0,0,0,0,0,0,0]

Day 2: temp=75
→ 75 > 74, pop 1, result[1]=2-1=1
→ stack empty, push 2
→ stack=[2], result=[1,1,0,0,0,0,0,0,0]

Day 3: temp=71
→ 71 ≤ 75, no pop
→ push 3
→ stack=[2,3], result=[1,1,0,0,0,0,0,0,0]

Day 4: temp=69
→ 69 ≤ 71, no pop
→ push 4
→ stack=[2,3,4], result=[1,1,0,0,0,0,0,0,0]

Day 5: temp=72
→ 72 > 69, pop 4, result[4]=5-4=1
→ 72 > 71, pop 3, result[3]=5-3=2
→ push 5
→ stack=[2,5], result=[1,1,0,2,1,0,0,0,0]

Day 6: temp=76
→ 76 > 72, pop 5, result[5]=6-5=1
→ 76 > 75, pop 2, result[2]=6-2=4
→ push 6
→ stack=[6], result=[1,1,0,2,1,0,4,0,0]

Day 7: temp=73
→ 73 ≤ 76, no pop
→ push 7
→ stack=[6,7], result=[1,1,0,2,1,0,4,0,0]

Final result: [1,1,4,2,1,0,4,0]
```

### Key Visualization Points:
- **Stack** maintains decreasing temperature indices
- **Pop operation** resolves wait time for cooler days
- **Push operation** adds current day as candidate
- **Distance calculation** uses current index - previous index

### Memory Layout Visualization:
```
Day:     0  1  2  3  4  5  6  7
Temp:     73 74 75 71 69 72 76 73
Stack:     [0]→[1]→[2]→[2,3]→[2,3,4]→[2,3,4,5]→[6]→[6,7]
Result:    [1,1,4,2,1,0,4,0]
           ↑  ↑  ↑     ↑  ↑     ↑
           Days 0,1 wait 1 day each
           Days 2,3 wait for warmer days 4,6 respectively
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each day processed once
- **Stack Operations**: Each index pushed and popped at most once
- **Total**: O(N) time, O(N) space for stack and result
- **Optimal**: Cannot do better than O(N) for this problem
*/
