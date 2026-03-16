import java.util.*;

public class Java5_Features {
    // Enum introduced in Java 5
    enum Color { RED, GREEN, BLUE }

    public static void main(String[] args) {
        // Java 5: Generics, Enums, Autoboxing, Enhanced For-loop, Varargs, Annotations
        
        // Generics & Type Safety
        List<String> list = new ArrayList<String>();
        list.add("Java 5 Generics");

        // Enhanced For-loop (foreach)
        for (String s : list) {
            System.out.println(s);
        }

        // Autoboxing (int to Integer)
        Integer autoBoxed = 42; 
        System.out.println("Autoboxed value: " + autoBoxed);

        // Varargs usage
        printItems("Item1", "Item2", "Item3");
    }

    // Varargs method
    public static void printItems(String... items) {
        System.out.print("Varargs output: ");
        for (String item : items) {
            System.out.print(item + " ");
        }
        System.out.println();
    }
}
