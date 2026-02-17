import java.util.*;

/**
 * Demonstrates ListIterator Interface methods.
 * Only available for Lists.
 * 
 * Methods covered:
 * - hasPrevious(), previous()
 * - nextIndex(), previousIndex()
 * - set(), add()
 */
public class ListIteratorDemo {

    public static void main(String[] args) {
        List<String> library = new ArrayList<>();
        library.add("Book 1");
        library.add("Book 2");
        library.add("Book 3");

        System.out.println("Original List: " + library);

        ListIterator<String> listIt = library.listIterator();

        System.out.println("\n--- Forward Traversal ---");
        while (listIt.hasNext()) {
            System.out.println("Index " + listIt.nextIndex() + ": " + listIt.next());
        }

        System.out.println("\n--- Backward Traversal ---");
        // Cursor is now at the end
        while (listIt.hasPrevious()) {
            System.out.println("Index " + listIt.previousIndex() + ": " + listIt.previous());
        }

        System.out.println("\n--- Modification during Traversal ---");
        // Reset iterator to start
        listIt = library.listIterator();

        while (listIt.hasNext()) {
            String book = listIt.next();

            if (book.equals("Book 2")) {
                listIt.set("Updated Book 2"); // Replace
                listIt.add("Inserted Book 2.5"); // Insert after
            }
        }

        System.out.println("Modified List: " + library);
    }
}
