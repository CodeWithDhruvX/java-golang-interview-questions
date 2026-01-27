package java_solutions;

import java.util.*;

public class AdditionalDS {

    // --- LINKED LIST (BASIC) ---
    static class ListNode {
        int val;
        ListNode next;

        ListNode(int val) {
            this.val = val;
        }
    }

    // 51. Create LL (Helper)
    public static ListNode createLL(int val) {
        return new ListNode(val);
    }

    // 52. Traverse
    public static void traverseLL(ListNode head) {
        ListNode curr = head;
        while (curr != null) {
            System.out.print(curr.val + " -> ");
            curr = curr.next;
        }
        System.out.println("NULL");
    }

    // 53. Insert Begin
    public static ListNode insertBegin(ListNode head, int val) {
        ListNode newNode = new ListNode(val);
        newNode.next = head;
        return newNode;
    }

    // 54. Insert End
    public static ListNode insertEnd(ListNode head, int val) {
        ListNode newNode = new ListNode(val);
        if (head == null)
            return newNode;
        ListNode curr = head;
        while (curr.next != null)
            curr = curr.next;
        curr.next = newNode;
        return head;
    }

    // 55. Delete Node
    public static ListNode deleteNode(ListNode head, int val) {
        if (head == null)
            return null;
        if (head.val == val)
            return head.next;
        ListNode curr = head;
        while (curr.next != null) {
            if (curr.next.val == val) {
                curr.next = curr.next.next;
                return head;
            }
            curr = curr.next;
        }
        return head;
    }

    // 56. Reverse LL
    public static ListNode reverseLL(ListNode head) {
        ListNode prev = null, curr = head;
        while (curr != null) {
            ListNode nextTemp = curr.next;
            curr.next = prev;
            prev = curr;
            curr = nextTemp;
        }
        return prev;
    }

    // 57. Middle Element
    public static ListNode middleNode(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        return slow;
    }

    // 58. Detect Loop
    public static boolean hasCycle(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast)
                return true;
        }
        return false;
    }

    // 59. Count Nodes
    public static int countNodes(ListNode head) {
        int count = 0;
        ListNode curr = head;
        while (curr != null) {
            count++;
            curr = curr.next;
        }
        return count;
    }

    // 60. Merge 2 Sorted Lists
    public static ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0);
        ListNode curr = dummy;
        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                curr.next = l1;
                l1 = l1.next;
            } else {
                curr.next = l2;
                l2 = l2.next;
            }
            curr = curr.next;
        }
        curr.next = (l1 != null) ? l1 : l2;
        return dummy.next;
    }

    // --- STACK & QUEUE (BASIC) ---

    // 61. Stack using Array
    static class MyStack {
        int[] items = new int[100];
        int top = -1;

        void push(int x) {
            items[++top] = x;
        }

        int pop() {
            return (top == -1) ? -1 : items[top--];
        }
    }

    // 62. Queue using Array
    static class MyQueue {
        int[] items = new int[100];
        int front = 0, rear = 0;

        void enqueue(int x) {
            items[rear++] = x;
        }

        int dequeue() {
            return (front == rear) ? -1 : items[front++];
        }
    }

    // 63. Reverse String Stack
    public static String reverseStringStack(String str) {
        Stack<Character> stack = new Stack<>();
        for (char c : str.toCharArray())
            stack.push(c);
        StringBuilder sb = new StringBuilder();
        while (!stack.isEmpty())
            sb.append(stack.pop());
        return sb.toString();
    }

    // 67. Next Greater Element
    public static int[] nextGreaterElement(int[] nums) {
        int[] res = new int[nums.length];
        Arrays.fill(res, -1);
        Stack<Integer> stack = new Stack<>(); // stores indices
        for (int i = 0; i < nums.length; i++) {
            while (!stack.isEmpty() && nums[stack.peek()] < nums[i]) {
                res[stack.pop()] = nums[i];
            }
            stack.push(i);
        }
        return res;
    }

    // 68. Evaluate Postfix
    public static int evalPostfix(String exp) {
        Stack<Integer> stack = new Stack<>();
        for (char c : exp.toCharArray()) {
            if (Character.isDigit(c))
                stack.push(c - '0');
            else {
                int val1 = stack.pop();
                int val2 = stack.pop();
                switch (c) {
                    case '+':
                        stack.push(val2 + val1);
                        break;
                    case '-':
                        stack.push(val2 - val1);
                        break;
                    case '*':
                        stack.push(val2 * val1);
                        break;
                    case '/':
                        stack.push(val2 / val1);
                        break;
                }
            }
        }
        return stack.pop();
    }

    // --- HASHING ---

    // 71. Freq Elements
    public static void freqElements(int[] arr) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int v : arr)
            map.put(v, map.getOrDefault(v, 0) + 1);
        System.out.println(map);
    }

    // 72. First Repeating
    public static int firstRepeating(int[] arr) {
        Set<Integer> seen = new HashSet<>();
        for (int v : arr) {
            if (seen.contains(v))
                return v;
            seen.add(v);
        }
        return -1;
    }

    // 74. Two Sum
    public static int[] twoSum(int[] arr, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < arr.length; i++) {
            int needed = target - arr[i];
            if (map.containsKey(needed)) {
                return new int[] { map.get(needed), i };
            }
            map.put(arr[i], i);
        }
        return null;
    }

    // 75. Group Anagrams
    public static List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String s : strs) {
            char[] ca = s.toCharArray();
            Arrays.sort(ca);
            String key = new String(ca);
            if (!map.containsKey(key))
                map.put(key, new ArrayList<>());
            map.get(key).add(s);
        }
        return new ArrayList<>(map.values());
    }

    // 80. Longest Substring No Repeats
    public static int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> map = new HashMap<>();
        int maxLen = 0, start = 0;
        for (int i = 0; i < s.length(); i++) {
            if (map.containsKey(s.charAt(i))) {
                start = Math.max(start, map.get(s.charAt(i)) + 1);
            }
            map.put(s.charAt(i), i);
            maxLen = Math.max(maxLen, i - start + 1);
        }
        return maxLen;
    }

    public static void main(String[] args) {
        System.out.println("--- LINKED LIST ---");
        ListNode head = createLL(1);
        head = insertEnd(head, 2);
        head = insertEnd(head, 3);
        traverseLL(head);
        System.out.println("Middle: " + middleNode(head).val);
        head = reverseLL(head);
        traverseLL(head);

        System.out.println("\n--- STACK/QUEUE ---");
        MyStack s = new MyStack();
        s.push(10);
        s.push(20);
        System.out.println("Pop: " + s.pop());
        System.out.println("Reverse String Stack: " + reverseStringStack("hello"));
        System.out.println("Next Greater: " + Arrays.toString(nextGreaterElement(new int[] { 4, 5, 2, 25 })));
        System.out.println("Postfix: 231*+9- -> " + evalPostfix("231*+9-"));

        System.out.println("\n--- HASHING ---");
        System.out.print("Freq: ");
        freqElements(new int[] { 1, 2, 2, 3 });
        System.out.println("First Repeating: " + firstRepeating(new int[] { 1, 2, 3, 2 }));
        System.out.println("Two Sum: " + Arrays.toString(twoSum(new int[] { 2, 7, 11, 15 }, 9)));
        System.out
                .println("Group Anagrams: " + groupAnagrams(new String[] { "eat", "tea", "tan", "ate", "nat", "bat" }));
        System.out.println("Longest Substr: abcabcbb -> " + lengthOfLongestSubstring("abcabcbb"));
    }
}
