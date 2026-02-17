import java.util.*;
import java.util.function.UnaryOperator;

/**
 * Demonstrates List Interface methods and implementations.
 * 
 * Implementations: ArrayList, LinkedList, Vector, Stack
 * 
 * Methods covered:
 * - add(index, element), get(index), set(index, element), remove(index)
 * - indexOf(o), lastIndexOf(o)
 * - listIterator(), subList(from, to)
 * - sort(Comparator), replaceAll(UnaryOperator)
 */
public class ListDemo {

    public static void main(String[] args) {
        System.out.println("=== 1. ArrayList Implementation ===");
        List<String> arrayList = new ArrayList<>();
        arrayList.add("Java");
        arrayList.add("Python");
        arrayList.add("C++");
        arrayList.add("Java"); // Duplicate
        
        // Basic List Operations
        System.out.println("Original List: " + arrayList);
        
        // add(int index, E element)
        arrayList.add(1, "JavaScript");
        System.out.println("After add at index 1: " + arrayList);
        
        // get(int index)
        System.out.println("Element at index 2: " + arrayList.get(2));
        
        // set(int index, E element)
        arrayList.set(0, "Golang");
        System.out.println("After set at index 0: " + arrayList);
        
        // indexOf / lastIndexOf
        System.out.println("IndexOf 'Java': " + arrayList.indexOf("Java"));
        System.out.println("LastIndexOf 'Java': " + arrayList.lastIndexOf("Java"));
        
        // remove(int index)
        arrayList.remove(3); // Removes C++
        System.out.println("After remove at index 3: " + arrayList);

        // subList
        List<String> sub = arrayList.subList(0, 2);
        System.out.println("SubList (0-2): " + sub);
        
        // sort
        arrayList.sort(Comparator.naturalOrder());
        System.out.println("Sorted: " + arrayList);
        
        // replaceAll
        arrayList.replaceAll(s -> s.toUpperCase());
        System.out.println("After replaceAll (UpperCase): " + arrayList);


        System.out.println("\n=== 2. LinkedList Implementation ===");
        // Good for frequent insertions/deletions
        LinkedList<String> linkedList = new LinkedList<>();
        linkedList.add("Node1");
        linkedList.add("Node2");
        linkedList.addFirst("Head"); // Deque method available in LinkedList
        linkedList.addLast("Tail");
        System.out.println("LinkedList: " + linkedList);


        System.out.println("\n=== 3. Vector Implementation ===");
        // Synchronized, legacy
        Vector<Integer> vector = new Vector<>();
        vector.add(10);
        vector.addElement(20); // Legacy method
        System.out.println("Vector: " + vector);


        System.out.println("\n=== 4. Stack Implementation ===");
        // Extends Vector, LIFO
        Stack<String> stack = new Stack<>();
        stack.push("Bottom");
        stack.push("Middle");
        stack.push("Top");
        
        System.out.println("Stack: " + stack);
        System.out.println("Pop: " + stack.pop());
        System.out.println("Peek: " + stack.peek());
        System.out.println("Stack after pop: " + stack);
    }
}
