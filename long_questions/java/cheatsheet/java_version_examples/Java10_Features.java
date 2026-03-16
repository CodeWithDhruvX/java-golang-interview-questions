import java.util.*;

public class Java10_Features {
    public static void main(String[] args) {
        // Java 10: Local Variable Type Inference (var)
        
        var message = "Hello, Java 10 Type Inference!";
        var list = new ArrayList<String>(); // Inferred as ArrayList<String>
        list.add(message);

        for (var item : list) {
            System.out.println("Inferred type item: " + item);
        }

        System.out.println("Note: 'var' can only be used for LOCAL variables with initialization.");
    }
}
