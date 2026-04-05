import java.util.Arrays;
import java.util.PriorityQueue;

public class KthLargestElementInAnArray {
    
    // 215. Kth Largest Element in an Array
    // Time: O(N log K), Space: O(K)
    public static int findKthLargest(int[] nums, int k) {
        // Use a min-heap of size k
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        
        for (int num : nums) {
            minHeap.offer(num);
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }
        
        return minHeap.peek();
    }

    // Alternative solution using QuickSelect (O(N) average case)
    public static int findKthLargestQuickSelect(int[] nums, int k) {
        return quickSelect(nums, 0, nums.length - 1, nums.length - k);
    }

    private static int quickSelect(int[] nums, int left, int right, int kthSmallest) {
        if (left == right) {
            return nums[left];
        }
        
        int pivotIndex = partition(nums, left, right);
        
        if (kthSmallest == pivotIndex) {
            return nums[pivotIndex];
        } else if (kthSmallest < pivotIndex) {
            return quickSelect(nums, left, pivotIndex - 1, kthSmallest);
        } else {
            return quickSelect(nums, pivotIndex + 1, right, kthSmallest);
        }
    }

    private static int partition(int[] nums, int left, int right) {
        int pivot = nums[right];
        int i = left;
        
        for (int j = left; j < right; j++) {
            if (nums[j] <= pivot) {
                swap(nums, i, j);
                i++;
            }
        }
        
        swap(nums, i, right);
        return i;
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {3, 2, 1, 5, 6, 4},
            {3, 2, 3, 1, 2, 4, 5, 5, 6},
            {1},
            {2, 1},
            {1, 2, 3, 4, 5},
            {5, 4, 3, 2, 1},
            {3, 2, 3, 1, 2, 4, 5, 5, 6},
            {7, 10, 4, 3, 20, 15},
            {-1, -2, -3, -4, -5},
            {100, 200, 300, 400, 500}
        };
        
        int[] kValues = {2, 4, 1, 1, 5, 1, 1, 3, 2, 4};
        
        for (int i = 0; i < testArrays.length; i++) {
            // Make copies for both methods
            int[] nums1 = testArrays[i].clone();
            int[] nums2 = testArrays[i].clone();
            
            int result1 = findKthLargest(nums1, kValues[i]);
            int result2 = findKthLargestQuickSelect(nums2, kValues[i]);
            
            System.out.printf("Test Case %d: nums=%s, k=%d -> Heap: %d, QuickSelect: %d\n", 
                i + 1, Arrays.toString(testArrays[i]), kValues[i], result1, result2);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Heap/Priority Queue
- **Kth Largest Element**: Find k-th largest using min-heap
- **Heap Maintenance**: Keep only k largest elements
- **Priority Queue**: Always maintain smallest of k largest
- **Selection Problem**: Order statistics without full sorting

## 2. PROBLEM CHARACTERISTICS
- **Find Kth Largest**: Need k-th largest element in array
- **Partial Sorting**: Don't need full sort, just k elements
- **Heap Properties**: Min-heap keeps smallest at root
- **Space Tradeoff**: O(K) space vs O(N log N) full sort

## 3. SIMILAR PROBLEMS
- Kth Smallest Element
- Find Median from Data Stream
- Top K Frequent Elements
- Sliding Window Maximum

## 4. KEY OBSERVATIONS
- Min-heap of size k keeps k largest elements
- Root of min-heap is k-th largest element
- Each insertion/poll is O(log k) operation
- Total time is O(N log k) vs O(N log N) for full sort

## 5. VARIATIONS & EXTENSIONS
- Kth smallest element (using max-heap)
- Dynamic k value
- Multiple queries on same array
- Streaming data processing

## 6. INTERVIEW INSIGHTS
- Clarify: "Is k always valid?"
- Edge cases: empty array, k=1, k=N
- Time complexity: O(N log K) vs O(N log N) sorting
- Space complexity: O(K) vs O(1) for QuickSelect

## 7. COMMON MISTAKES
- Using max-heap for k-th largest
- Not handling k > array length
- Integer overflow in comparisons
- Forgetting to check heap size

## 8. OPTIMIZATION STRATEGIES
- QuickSelect for O(N) average case
- Early termination for small k
- In-place heap operations
- Hybrid approaches for different k ranges

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like maintaining a top-k leaderboard:**
- You have contestants (array elements) competing
- You want to know who's ranked k-th place
- Maintain a leaderboard showing only top k performers
- When someone new joins, check if they make top k
- If leaderboard is full, remove the lowest performer
- The lowest in top k is exactly the k-th rank

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers and integer k
2. **Goal**: Find k-th largest element
3. **Output**: The element at rank k from largest

#### Phase 2: Key Insight Recognition
- **"What does k-th largest mean?"** → k-th element when sorted descending
- **"How to maintain top k?"** → Use min-heap of size k
- **"Why min-heap?"** → Root is smallest among top k
- **"When to update heap?"** → For each new element

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use min-heap:
1. Create empty min-heap
2. For each number in array:
   - Add to heap
   - If heap size > k, remove smallest
3. After processing all, heap root is k-th largest
4. This keeps only top k elements"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return error or handle gracefully
- **k = 1**: Return maximum element
- **k = N**: Return minimum element
- **k > N**: Return error

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [3, 2, 1, 5, 6, 4], k = 2

Human thinking:
"Let's maintain top 2 elements:

Process 3:
- Heap: [3], size = 1 ≤ 2
- Keep 3

Process 2:
- Add 2: Heap: [2,3], size = 2 ≤ 2
- Keep [2,3]

Process 1:
- Add 1: Heap: [1,3,2], size = 3 > 2
- Remove smallest (1): Heap: [2,3]
- Keep [2,3]

Process 5:
- Add 5: Heap: [2,3,5], size = 3 > 2
- Remove smallest (2): Heap: [3,5]
- Keep [3,5]

Process 6:
- Add 6: Heap: [3,5,6], size = 3 > 2
- Remove smallest (3): Heap: [5,6]
- Keep [5,6]

Process 4:
- Add 4: Heap: [4,6,5], size = 3 > 2
- Remove smallest (4): Heap: [5,6]
- Keep [5,6]

Final heap: [5,6]
Root (smallest): 5
This is the 2nd largest element ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Heap always contains k largest elements
- **Why it's efficient**: Each operation is O(log k)
- **Why it's correct**: Root is smallest among k largest

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N log K) when K << N
2. **"What about max-heap?"** → Would give k-th smallest, not largest
3. **"How to handle duplicates?"** → Heap handles naturally
4. **"What about QuickSelect?"** → O(N) average but O(N²) worst case

### Real-World Analogy
**Like managing a VIP waiting list:**
- You have customers arriving (array elements)
- You can only serve k VIP customers at once
- When a new VIP arrives, check if they qualify
- If list is full, remove the lowest priority VIP
- The lowest priority VIP in the list is at rank k
- This maintains exactly the k highest priority customers

### Human-Readable Pseudocode
```
function findKthLargest(nums, k):
    if k > nums.length:
        return error
    
    minHeap = empty priority queue
    
    for num in nums:
        minHeap.offer(num)
        if minHeap.size() > k:
            minHeap.poll()  // remove smallest
    
    return minHeap.peek()  // k-th largest
```

### Execution Visualization

### Example: nums = [3, 2, 1, 5, 6, 4], k = 2
```
Heap Evolution:
Process 3: [3]              size=1
Process 2: [2,3]            size=2
Process 1: [2,3] → [3,2]   size=2 (removed 1)
Process 5: [3,2] → [5,3]   size=2 (removed 2)
Process 6: [5,3] → [6,5]   size=2 (removed 3)
Process 4: [6,5] → [5,6]   size=2 (removed 4)

Final heap: [5,6]
Root: 5 (2nd largest) ✓

The top 2 elements are [6,5] ✓
```

### Key Visualization Points:
- **Min-heap** maintains k largest elements
- **Size control** ensures only k elements kept
- **Root element** is k-th largest
- **Insertion/removal** are O(log k) operations

### Memory Layout Visualization:
```
Heap Structure Evolution:
[3]                →    [3]
[2,3]              →   [2,3]
[1,2,3] → [2,3]  (removed 1)
[2,3,5] → [3,5]  (removed 2)
[3,5,6] → [5,6]  (removed 3)
[4,5,6] → [5,6]  (removed 4)

Final: [5,6] where 5 is root (2nd largest) ✓
```

### Time Complexity Breakdown:
- **Each Insertion**: O(log k) for heap maintenance
- **Each Removal**: O(log k) when size exceeds k
- **Total Elements**: N insertions, at most N-k removals
- **Total**: O(N log k) time, O(k) space
- **Optimal**: Better than O(N log N) when k << N
- **vs QuickSelect**: O(N) average but O(N²) worst case
*/
