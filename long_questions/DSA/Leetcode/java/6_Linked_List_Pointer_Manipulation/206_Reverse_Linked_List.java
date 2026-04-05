import java.util.*;

public class ReverseLinkedList {
    
    // Definition for singly-linked list.
    public static class ListNode {
        int val;
        ListNode next;
        
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { 
            this.val = val; 
            this.next = next; 
        }
    }

    // 206. Reverse Linked List
    // Time: O(N), Space: O(1)
    public static ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode current = head;
        
        while (current != null) {
            ListNode next = current.next;
            current.next = prev;
            prev = current;
            current = next;
        }
        
        return prev;
    }

    // Helper function to create a linked list from array
    public static ListNode createLinkedList(int[] nums) {
        if (nums.length == 0) {
            return null;
        }
        
        ListNode head = new ListNode(nums[0]);
        ListNode current = head;
        
        for (int i = 1; i < nums.length; i++) {
            current.next = new ListNode(nums[i]);
            current = current.next;
        }
        
        return head;
    }

    // Helper function to convert linked list to array
    public static int[] linkedListToArray(ListNode head) {
        List<Integer> result = new ArrayList<>();
        ListNode current = head;
        
        while (current != null) {
            result.add(current.val);
            current = current.next;
        }
        
        return result.stream().mapToInt(i -> i).toArray();
    }

    // Alternative recursive approach
    public static ListNode reverseListRecursive(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        
        ListNode newHead = reverseListRecursive(head.next);
        head.next.next = head;
        head.next = null;
        
        return newHead;
    }

    // Alternative approach using array (not space efficient but educational)
    public static ListNode reverseListUsingArray(ListNode head) {
        if (head == null) {
            return null;
        }
        
        // Store values in array
        List<Integer> values = new ArrayList<>();
        ListNode current = head;
        while (current != null) {
            values.add(current.val);
            current = current.next;
        }
        
        // Create new reversed list
        ListNode newHead = new ListNode(values.get(values.size() - 1));
        current = newHead;
        
        for (int i = values.size() - 2; i >= 0; i--) {
            current.next = new ListNode(values.get(i));
            current = current.next;
        }
        
        return newHead;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 4, 5},
            {1, 2},
            {},
            {1},
            {1, 2, 3, 4},
            {5, 4, 3, 2, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            ListNode head = createLinkedList(testCases[i]);
            ListNode reversedHead = reverseList(head);
            int[] result = linkedListToArray(reversedHead);
            
            System.out.printf("Test Case %d: %s -> Reversed: %s\n", 
                i + 1, Arrays.toString(testCases[i]), Arrays.toString(result));
        }
        
        // Test alternative approaches
        System.out.println("\n=== Testing Alternative Approaches ===");
        ListNode testHead = createLinkedList(new int[]{1, 2, 3, 4, 5});
        
        System.out.println("Original: " + Arrays.toString(linkedListToArray(testHead)));
        
        ListNode recursiveResult = reverseListRecursive(testHead);
        System.out.println("Recursive: " + Arrays.toString(linkedListToArray(recursiveResult)));
        
        ListNode arrayResult = reverseListUsingArray(testHead);
        System.out.println("Array-based: " + Arrays.toString(linkedListToArray(arrayResult)));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: In-place Pointer Manipulation
- **Three Pointers**: Previous, Current, Next nodes
- **Link Reversal**: Point current.next to previous
- **Pointer Movement**: Advance all three pointers
- **New Head**: Previous node becomes new head

## 2. PROBLEM CHARACTERISTICS
- **Singly Linked List**: Unidirectional node connections
- **In-place Modification**: Cannot use extra array/list
- **Pointer Manipulation**: Careful management of node references
- **Edge Cases**: Empty list, single node

## 3. SIMILAR PROBLEMS
- Reverse Linked List II (partial reversal)
- Swap Nodes in Pairs
- Rotate List
- Palindrome Linked List

## 4. KEY OBSERVATIONS
- Need three pointers to avoid losing references
- Previous pointer tracks new list head
- Current pointer processes original list
- Next pointer preserves remaining list
- Order of operations is critical

## 5. VARIATIONS & EXTENSIONS
- Recursive approach (uses call stack)
- Reverse in groups of K nodes
- Reverse first K nodes only
- Handle circular linked lists

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I modify original list?"
- Edge cases: empty list, single node, two nodes
- Memory management: avoid memory leaks
- Alternative approaches: recursive vs iterative

## 7. COMMON MISTAKES
- Losing reference to remaining list
- Not handling empty list case
- Creating cycles in the reversed list
- Wrong order of pointer updates

## 8. OPTIMIZATION STRATEGIES
- Iterative approach: O(1) space, O(N) time
- Recursive approach: O(N) space for call stack
- For very large lists, consider tail recursion
- For memory constraints, iterative is better

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like reversing a train track:**
- You have train cars connected in one direction (linked list)
- You need to reverse the direction of all connections
- You carefully disconnect each car and reconnect it backward
- You keep track of where the new train starts


### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Head of singly linked list
2. **Goal**: Reverse all node connections
3. **Output**: New head of reversed list

#### Phase 2: Key Insight Recognition
- **"Need three pointers!"** → Previous, Current, Next
- **"Why three?"** → To avoid losing list reference
- **"How to reverse?"** → Point current.next to previous
- **"What becomes new head?"** → Last node becomes first

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use three pointers:
- prev: tracks the reversed list so far
- curr: processes the original list
- next: saves the rest of the original list

For each node, I'll:
1. Save next node
2. Point current.next to prev
3. Move prev to current
4. Move current to saved next"
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return null
- **Single node**: Node points to null, return same node
- **Two nodes**: Simple swap
- **Null pointer checks**: Always validate before accessing

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: 1 → 2 → 3 → 4 → 5

Human thinking:
"Initialize: prev = null, curr = 1 (head)

Step 1: curr = 1
           next = 2 (save curr.next)
           curr.next = prev (1.next = null)
           prev = curr (prev = 1)
           curr = next (curr = 2)
           List: null ← 1 → 2 → 3 → 4 → 5

Step 2: curr = 2
           next = 3 (save curr.next)
           curr.next = prev (2.next = 1)
           prev = curr (prev = 2)
           curr = next (curr = 3)
           List: null ← 1 ← 2 → 3 → 4 → 5

Continue until curr = null...
Final head = prev = 5"
```

#### Phase 6: Intuition Validation
- **Why it works**: We systematically reverse each connection
- **Why it's efficient**: Single pass, constant space
- **Why it's correct**: Every node gets reconnected exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just swap values?"** → Need to restructure connections
2. **"Can I use two pointers?"** → Need three to avoid losing list
3. **"What about order?"** → Save next BEFORE modifying current.next
4. **"Recursive vs iterative?"** → Iterative uses less space

### Real-World Analogy
**Like reversing a chain of paper clips:**
- You have paper clips linked together (linked list)
- You need to reverse the entire chain direction
- You carefully unclip each clip and reattach it backward
- You keep track of where the new chain starts
- Each clip must be reattached exactly once

### Human-Readable Pseudocode
```
function reverseLinkedList(head):
    prev = null
    curr = head
    
    while curr is not null:
        next = curr.next      // Save rest of list
        curr.next = prev     // Reverse connection
        prev = curr          // Move prev forward
        curr = next          // Move to next node
    
    return prev  // New head
```

### Execution Visualization

### Example: 1 → 2 → 3 → 4 → 5
```
Initial: head = 1→2→3→4→5
prev = null, curr = 1

Step 1: curr = 1
→ next = 2
→ 1.next = null
→ prev = 1, curr = 2
→ List: null←1→2→3→4→5

Step 2: curr = 2
→ next = 3
→ 2.next = 1
→ prev = 2, curr = 3
→ List: null←1←2→3→4→5

Step 3: curr = 3
→ next = 4
→ 3.next = 2
→ prev = 3, curr = 4
→ List: null←1←2←3→4→5

Step 4: curr = 4
→ next = 5
→ 4.next = 3
→ prev = 4, curr = 5
→ List: null←1←2←3←4→5

Step 5: curr = 5
→ next = null
→ 5.next = 4
→ prev = 5, curr = null
→ List: null←1←2←3←4←5

Final: return prev = 5
Reversed: 5→4→3→2→1 ✓
```

### Key Visualization Points:
- **Three pointers** maintain list integrity
- **Link reversal** happens node by node
- **Previous pointer** becomes new head
- **Current pointer** traverses original list

### Memory Layout Visualization:
```
Original: 1 → 2 → 3 → 4 → 5
         ↓  ↓  ↓  ↓  ↓
Reversed: 5 ← 4 ← 3 ← 2 ← 1

Step-by-step:
Step 1: null←1→2→3→4→5
Step 2: null←1←2→3→4→5
Step 3: null←1←2←3→4→5
Step 4: null←1←2←3←4→5
Step 5: null←1←2←3←4→5
Final:   5→4→3→2→1
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each node visited once
- **Constant Space**: O(1) - only three pointers
- **In-place**: No additional data structures
- **Optimal**: Cannot do better than O(N) for this problem
*/
