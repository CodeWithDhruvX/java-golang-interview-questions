import java.util.*;

/**
 * Demonstrates Deque (Double Ended Queue) Interface.
 * 
 * Implementations: ArrayDeque, LinkedList
 * 
 * Methods covered:
 * - addFirst, addLast, offerFirst, offerLast
 * - removeFirst, removeLast, pollFirst, pollLast
 * - peekFirst, peekLast
 * - push, pop (Stack behavior)
 */
public class DequeDemo {

    public static void main(String[] args) {
        System.out.println("=== ArrayDeque Implementation ===");
        Deque<String> deque = new ArrayDeque<>();

        // Add elements
        deque.addFirst("First");
        deque.addLast("Last");
        deque.offerFirst("New First");
        deque.offerLast("New Last");

        System.out.println("Deque: " + deque);

        // Peek
        System.out.println("Peek First: " + deque.peekFirst());
        System.out.println("Peek Last: " + deque.peekLast());

        // Remove
        System.out.println("Poll First: " + deque.pollFirst());
        System.out.println("Poll Last: " + deque.pollLast());
        System.out.println("Deque after polls: " + deque);

        // Stack Operation Simulation
        System.out.println("\n=== Stack Behavior (LIFO) using Deque ===");
        Deque<Integer> stack = new ArrayDeque<>();

        stack.push(10); // same as addFirst
        stack.push(20);
        stack.push(30);

        System.out.println("Stack: " + stack);

        System.out.println("Pop: " + stack.pop()); // same as removeFirst
        System.out.println("Pop: " + stack.pop());
        System.out.println("Stack after pops: " + stack);
    }
}
