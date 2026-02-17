import java.util.*;

/**
 * Demonstrates Queue Interface methods and implementations.
 * 
 * Implementations: PriorityQueue, ArrayDeque (as Queue), LinkedList (as Queue)
 * 
 * Methods covered:
 * - add(e) vs offer(e)
 * - remove() vs poll()
 * - element() vs peek()
 */
public class QueueDemo {

    public static void main(String[] args) {
        System.out.println("=== 1. PriorityQueue Implementation ===");
        // Orders elements according to their natural ordering or a Comparator
        Queue<Integer> pq = new PriorityQueue<>();

        // offer(e) - secure, returns false if queue full (though PQ is unbounded)
        pq.offer(10);
        pq.offer(5);
        pq.offer(20);
        pq.add(1); // add(e) throws exception if full

        System.out.println("PriorityQueue (Sorted/Heap): " + pq);

        // peek() - Retrieves head without removing, returns null if empty
        System.out.println("Peek Head: " + pq.peek());
        // element() - Retrieves head, throws Exception if empty
        System.out.println("Element Head: " + pq.element());

        // poll() - Retrieves and removes head, returns null if empty
        System.out.println("Poll Head: " + pq.poll());
        System.out.println("Queue after poll: " + pq);

        // remove() - Retrieves and removes head, throws Exception if empty
        System.out.println("Remove Head: " + pq.remove());
        System.out.println("Queue after remove: " + pq);

        // Clearing to test empty behavior
        pq.clear();
        System.out.println("Poll empty queue: " + pq.poll()); // null
        try {
            pq.remove(); // Exception
        } catch (NoSuchElementException e) {
            System.out.println("Remove on empty queue threw Exception");
        }

        System.out.println("\n=== 2. LinkedList as Queue ===");
        // FIFO (First In First Out)
        Queue<String> linkedQueue = new LinkedList<>();
        linkedQueue.offer("First");
        linkedQueue.offer("Second");
        linkedQueue.offer("Third");

        System.out.println("LinkedList Queue: " + linkedQueue);
        System.out.println("Poll: " + linkedQueue.poll()); // First
        System.out.println("Queue after poll: " + linkedQueue);

        System.out.println("\n=== 3. ArrayDeque as Queue ===");
        // Faster than LinkedList for queue operations
        Queue<String> arrayDequeQueue = new ArrayDeque<>();
        arrayDequeQueue.offer("A");
        arrayDequeQueue.offer("B");

        System.out.println("ArrayDeque Queue: " + arrayDequeQueue);
        System.out.println("Poll: " + arrayDequeQueue.poll());
    }
}
