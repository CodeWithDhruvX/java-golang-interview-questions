import java.util.*;

public class LongestConsecutiveSequence {
    
    // 128. Longest Consecutive Sequence
    // Time: O(N), Space: O(N)
    public static int longestConsecutive(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        Set<Integer> numSet = new HashSet<>();
        for (int num : nums) {
            numSet.add(num);
        }
        
        int maxLength = 0;
        
        for (int num : numSet) {
            // Only start counting from the beginning of a sequence
            if (!numSet.contains(num - 1)) {
                int currentNum = num;
                int currentLength = 1;
                
                // Count the length of consecutive sequence
                while (numSet.contains(currentNum + 1)) {
                    currentNum++;
                    currentLength++;
                }
                
                maxLength = Math.max(maxLength, currentLength);
            }
        }
        
        return maxLength;
    }

    // Union-Find approach
    public static int longestConsecutiveUnionFind(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        // Initialize Union-Find
        Map<Integer, Integer> parent = new HashMap<>();
        Map<Integer, Integer> rank = new HashMap<>();
        
        for (int num : nums) {
            parent.put(num, num);
            rank.put(num, 0);
        }
        
        // Find function with path compression
        class UnionFind {
            int find(int x) {
                if (parent.get(x) != x) {
                    parent.put(x, find(parent.get(x)));
                }
                return parent.get(x);
            }
            
            void union(int x, int y) {
                int rootX = find(x);
                int rootY = find(y);
                
                if (rootX == rootY) {
                    return;
                }
                
                if (rank.get(rootX) < rank.get(rootY)) {
                    parent.put(rootX, rootY);
                } else if (rank.get(rootX) > rank.get(rootY)) {
                    parent.put(rootY, rootX);
                } else {
                    parent.put(rootY, rootX);
                    rank.put(rootX, rank.get(rootX) + 1);
                }
            }
        }
        
        UnionFind uf = new UnionFind();
        
        // Union consecutive numbers
        for (int num : nums) {
            if (parent.containsKey(num - 1)) {
                uf.union(num, num - 1);
            }
            if (parent.containsKey(num + 1)) {
                uf.union(num, num + 1);
            }
        }
        
        // Count the size of each component
        Map<Integer, Integer> componentSize = new HashMap<>();
        for (int num : nums) {
            int root = uf.find(num);
            componentSize.put(root, componentSize.getOrDefault(root, 0) + 1);
        }
        
        int maxLength = 0;
        for (int size : componentSize.values()) {
            maxLength = Math.max(maxLength, size);
        }
        
        return maxLength;
    }

    // Sorting approach (O(N log N))
    public static int longestConsecutiveSorting(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        // Remove duplicates and sort
        Set<Integer> unique = new HashSet<>();
        for (int num : nums) {
            unique.add(num);
        }
        
        int[] sorted = new int[unique.size()];
        int index = 0;
        for (int num : unique) {
            sorted[index++] = num;
        }
        
        // Simple bubble sort for demonstration (in practice, use Arrays.sort)
        for (int i = 0; i < sorted.length - 1; i++) {
            for (int j = 0; j < sorted.length - i - 1; j++) {
                if (sorted[j] > sorted[j + 1]) {
                    int temp = sorted[j];
                    sorted[j] = sorted[j + 1];
                    sorted[j + 1] = temp;
                }
            }
        }
        
        int maxLength = 1;
        int currentLength = 1;
        
        for (int i = 1; i < sorted.length; i++) {
            if (sorted[i] == sorted[i - 1] + 1) {
                currentLength++;
            } else {
                maxLength = Math.max(maxLength, currentLength);
                currentLength = 1;
            }
        }
        
        maxLength = Math.max(maxLength, currentLength);
        
        return maxLength;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {100, 4, 200, 1, 3, 2},
            {0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
            {},
            {1},
            {1, 2, 0, 1},
            {9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
            {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
            {1, 3, 5, 7, 9},
            {-1, -2, -3, 0, 1, 2, 3}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result1 = longestConsecutive(testCases[i]);
            int result2 = longestConsecutiveUnionFind(testCases[i]);
            int result3 = longestConsecutiveSorting(testCases[i]);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  HashSet: %d, Union-Find: %d, Sorting: %d\n\n", 
                result1, result2, result3);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Union-Find (Disjoint Set Union)
- **Union-Find**: Efficient connectivity data structure
- **Path Compression**: Optimize find operation
- **Union by Rank**: Keep tree shallow
- **Component Analysis**: Track connected component sizes
- **Consecutive Numbers**: Find longest sequence of consecutive integers

## 2. PROBLEM CHARACTERISTICS
- **Connectivity**: Determine if numbers are connected
- **Dynamic Union**: Numbers can be added/connected during processing
- **Component Tracking**: Need to know size of each component
- **Consecutive Constraint**: Numbers must differ by exactly 1

## 3. SIMILAR PROBLEMS
- Number of Provinces
- Graph Valid Tree
- Redundant Connection
- Accounts Merge

## 4. KEY OBSERVATIONS
- Union-Find maintains forest of trees
- Path compression flattens tree structure
- Union by rank prevents tall trees
- Component size tracking enables longest sequence finding
- Time complexity nearly constant for practical purposes

## 5. VARIATIONS & EXTENSIONS
- Union by size instead of rank
- Dynamic connectivity with deletions
- Offline processing of multiple queries
- Different union criteria

## 6. INTERVIEW INSIGHTS
- Clarify: "Are numbers guaranteed to be unique?"
- Edge cases: empty array, single element, all consecutive
- Time complexity: O(α(N)) vs O(N²) naive approach
- Space complexity: O(N) for Union-Find

## 7. COMMON MISTAKES
- Not implementing path compression
- Using union by size instead of rank
- Forgetting to update component sizes
- Off-by-one errors in consecutive checking

## 8. OPTIMIZATION STRATEGIES
- Path compression is crucial for performance
- Union by rank keeps trees balanced
- Lazy union (only when needed)
- Component size tracking for longest sequence

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like managing social networks:**
- Each person is a node in a social network
- Union-Find tracks who's connected to whom
- When two people become friends, their networks merge
- Path compression finds the "ultimate friend" quickly
- Component size tells you how big each friend group is

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers
2. **Goal**: Find length of longest consecutive sequence
3. **Output**: Maximum length of consecutive numbers

#### Phase 2: Key Insight Recognition
- **"What makes numbers consecutive?"** → Numbers that differ by exactly 1
- **"How to track connectivity?"** → Union-Find data structure
- **"What's the key insight?"** → Consecutive numbers must be in same component
- **"How to find longest?"** → Track maximum component size

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Union-Find:
1. For each number, create its own component initially
2. For consecutive numbers, union their components
3. Track size of each component
4. The maximum component size is the answer"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0
- **Single element**: Return 1
- **All consecutive**: Return array length
- **No consecutive pairs**: Return 1

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Numbers: [100, 4, 200, 1, 3, 2]

Human thinking:
"Initialize Union-Find with each number in its own set:
- 100: {100}, 4: {4}, 200: {200}, 1: {1}, 3: {3}, 2: {2}

Check consecutive pairs:
- 100 and 99? 99 not in array
- 100 and 101? 101 not in array
- 4 and 3? Both exist, union(4,3)
- 200 and 199? 199 not in array
- 200 and 201? 201 not in array
- 1 and 0? 0 not in array
- 1 and 2? Both exist, union(1,2)
- 3 and 2? Both exist, union(3,2)
- 3 and 4? Both exist, union(3,4)

Components after unions:
- {100}, {4,3}, {200}, {1,2}, {3,2}

Component sizes:
- {100}: size 1
- {4,3}: size 2
- {200}: size 1
- {1,2}: size 2
- {3,2}: size 2

Maximum size: 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Union-Find correctly tracks connectivity
- **Why it's efficient**: Path compression makes operations nearly O(1)
- **Why it's correct**: Consecutive numbers must be connected

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort and scan?"** → O(N log N) but Union-Find is more elegant
2. **"What about path compression?"** → Essential for performance
3. **"How to track component sizes?"** → Update during union operations
4. **"What about duplicate numbers?"** → Need to handle properly

### Real-World Analogy
**Like managing company departments:**
- Each employee starts in their own department (individual component)
- When employees collaborate, their departments merge
- Union-Find tracks which department each employee belongs to
- Path compression helps find the "ultimate department" quickly
- Component size shows how many employees are in each department
- Largest department size is like the biggest team collaboration

### Human-Readable Pseudocode
```
function longestConsecutive(nums):
    if nums is empty:
        return 0
    
    uf = UnionFind()
    
    // Initialize each number as its own component
    for num in nums:
        uf.makeSet(num)
    
    // Union consecutive numbers
    for num in nums:
        if uf.contains(num - 1):
            uf.union(num, num - 1)
        if uf.contains(num + 1):
            uf.union(num, num + 1)
    
    // Find maximum component size
    maxSize = 0
    for num in nums:
        root = uf.find(num)
        size = uf.getComponentSize(root)
        maxSize = max(maxSize, size)
    
    return maxSize
```

### Execution Visualization

### Example: nums = [100, 4, 200, 1, 3, 2]
```
Initial: Each number in its own component
Components: {100}, {4}, {200}, {1}, {3}, {2}

Union consecutive pairs:
- 4 and 3: union(4,3) → {4,3}
- 1 and 2: union(1,2) → {1,2}
- 3 and 2: union(3,2) → {1,2,3,2}

Final components:
- {100}: size 1
- {4,3,200}: size 3
- {1,2,3,2}: size 4

Maximum size: 4 ✓

Longest consecutive sequence: [1,2,3,2] or [3,2,4,3] (length 4)"
```

### Key Visualization Points:
- **Union-Find** maintains dynamic connectivity
- **Path compression** optimizes find operations
- **Component tracking** enables size queries
- **Consecutive checking** unites appropriate numbers
- **Maximum size** gives longest sequence length

### Memory Layout Visualization:
```
Union-Find Forest Evolution:
Initial:  {100} {4} {200} {1} {3} {2}
         |     |     |      |     |     |

After unions:
 {100} {4,3,200} {1,2,3,2}
         |     \     |      /     |
         |     |     |      |

Component Roots:
100 → rootA (size 1)
4,3,200 → rootB (size 3)
1,2,3,2 → rootC (size 4)

Path Compression:
find(3) → rootC (direct)
find(1) → rootC (direct)
```

### Time Complexity Breakdown:
- **Make Set**: O(N) operations
- **Union Operations**: O(α(N)) where α is inverse Ackermann function
- **Find Operations**: O(α(N)) with path compression
- **Component Tracking**: O(N) additional space
- **Total**: O(N) time for practical purposes, O(N) space
- **Optimal**: Nearly constant time per operation
*/
