import java.util.*;

public class RemoveNthNodeFromEnd {
    
    // Definition for singly-linked list.
    public static class ListNode {
        int val;
        ListNode next;
        
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    // 19. Remove Nth Node From End of List
    // Time: O(N), Space: O(1)
    public static ListNode removeNthFromEnd(ListNode head, int n) {
        if (head == null || n <= 0) {
            return head;
        }
        
        // Use two pointers with n distance apart
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode fast = dummy;
        ListNode slow = dummy;
        
        // Move fast pointer n steps ahead
        for (int i = 0; i <= n; i++) {
            if (fast == null) {
                return head; // n is larger than list length
            }
            fast = fast.next;
        }
        
        // Move both pointers until fast reaches end
        while (fast != null) {
            slow = slow.next;
            fast = fast.next;
        }
        
        // Remove the nth node from end
        slow.next = slow.next.next;
        
        return dummy.next;
    }

    // Helper method to create linked list from array
    public static ListNode createList(int[] arr) {
        if (arr == null || arr.length == 0) {
            return null;
        }
        
        ListNode dummy = new ListNode(0);
        ListNode current = dummy;
        
        for (int val : arr) {
            current.next = new ListNode(val);
            current = current.next;
        }
        
        return dummy.next;
    }

    // Helper method to convert linked list to array
    public static List<Integer> listToArray(ListNode head) {
        List<Integer> result = new ArrayList<>();
        ListNode current = head;
        
        while (current != null) {
            result.add(current.val);
            current = current.next;
        }
        
        return result;
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {new int[]{1, 2, 3, 4, 5}, 2},
            {new int[]{1}, 1},
            {new int[]{1, 2}, 1},
            {new int[]{1, 2, 3, 4, 5}, 1},
            {new int[]{1, 2, 3, 4, 5}, 5},
            {new int[]{1, 2, 3, 4, 5}, 3},
            {new int[]{}, 1},
            {new int[]{1, 2, 3}, 2},
            {new int[]{1, 2, 3, 4}, 4},
            {new int[]{1, 2, 3, 4, 5, 6, 7}, 6}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] arr = (int[]) testCases[i][0];
            int n = (int) testCases[i][1];
            
            ListNode head = createList(arr);
            List<Integer> original = listToArray(head);
            
            ListNode result = removeNthFromEnd(head, n);
            List<Integer> afterRemoval = listToArray(result);
            
            System.out.printf("Test Case %d: %s, n=%d -> %s\n", 
                i + 1, original, n, afterRemoval);
        }
    }
}
