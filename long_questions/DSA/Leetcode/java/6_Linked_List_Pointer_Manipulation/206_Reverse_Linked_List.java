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
}
