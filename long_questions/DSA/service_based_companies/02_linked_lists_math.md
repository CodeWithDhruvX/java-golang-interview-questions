# Linked Lists and Math (Service-Based Companies)

Linked Lists and fundamental Math/Logic questions are very common in technical interviews for service companies like TCS Ninja/Digital, Infosys HackWithInfy, Capgemini, and Wipro.

## Question 1: Reverse a Linked List
**Problem Statement:** Given the `head` of a singly linked list, reverse the list, and return the reversed list.

### Answer:
To reverse a linked list, we need to iterate through the list and change the `next` pointer of the current node to point to the `prev` node. We also need to keep track of the `next` node to continue the iteration.

**Code Implementation (Java):**
```java
class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

public class ReverseLinkedList {
    public ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode curr = head;
        
        while (curr != null) {
            ListNode nextTemp = curr.next; // Store next node
            curr.next = prev;              // Reverse the pointer
            prev = curr;                   // Move prev forward
            curr = nextTemp;               // Move curr forward
        }
        return prev;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 2: Detect Cycle in a Linked List
**Problem Statement:** Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

### Answer:
This is commonly solved using Floyd's Tortoise and Hare algorithm. We use two pointers, a slow pointer that moves one step at a time and a fast pointer that moves two steps at a time. If there is a cycle, the two pointers will eventually meet.

**Code Implementation (Java):**
```java
public class LinkedListCycle {
    public boolean hasCycle(ListNode head) {
        if (head == null || head.next == null) return false;
        
        ListNode slow = head;
        ListNode fast = head.next;
        
        while (slow != fast) {
            if (fast == null || fast.next == null) {
                return false;
            }
            slow = slow.next;
            fast = fast.next.next;
        }
        return true;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 3: Find Middle of the Linked List
**Problem Statement:** Given the `head` of a singly linked list, return the middle node of the linked list. If there are two middle nodes, return the second middle node.

### Answer:
We can use the fast and slow pointer technique. The slow pointer moves one step at a time, and the fast pointer moves two steps. When the fast pointer reaches the end, the slow pointer will be at the middle.

**Code Implementation (Java):**
```java
public class MiddleOfLinkedList {
    public ListNode middleNode(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;
        
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        return slow;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 4: FizzBuzz
**Problem Statement:** Given an integer `n`, return a string array where:
- `answer[i] == "FizzBuzz"` if `i` is divisible by 3 and 5.
- `answer[i] == "Fizz"` if `i` is divisible by 3.
- `answer[i] == "Buzz"` if `i` is divisible by 5.
- `answer[i] == i` (as a string) if none of the above conditions are true.

### Answer:
This is a very classic and popular screening question. Loop from 1 to `n`, checking the divisibility conditions using the modulo operator `%`. Always check the combined condition (divisible by 15) first.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.List;

public class FizzBuzz {
    public List<String> fizzBuzz(int n) {
        List<String> answer = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            if (i % 3 == 0 && i % 5 == 0) {
                answer.add("FizzBuzz");
            } else if (i % 3 == 0) {
                answer.add("Fizz");
            } else if (i % 5 == 0) {
                answer.add("Buzz");
            } else {
                answer.add(String.valueOf(i));
            }
        }
        return answer;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(N) (for output array)

---

## Question 5: Fibonacci Series (Nth Fibonacci)
**Problem Statement:** The Fibonacci numbers form a sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. Given `N`, calculate `F(N)`.

### Answer:
While recursion is easy to write, it's not optimal (O(2^N)). We should use an iterative approach (Dynamic Programming with space optimization) to get O(N) time and O(1) space.

**Code Implementation (Java):**
```java
public class Fibonacci {
    public int fib(int n) {
        if (n <= 1) return n;
        
        int prev2 = 0;
        int prev1 = 1;
        int current = 0;
        
        for (int i = 2; i <= n; i++) {
            current = prev1 + prev2;
            prev2 = prev1;
            prev1 = current;
        }
        
        return current;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)
