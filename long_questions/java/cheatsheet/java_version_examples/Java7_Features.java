import java.io.*;
import java.util.*;

public class Java7_Features {
    public static void main(String[] args) {
        // Java 7: Try-with-resources, Multi-catch, Diamond operator, String in switch
        
        // 1. Diamond operator (Type inference for constructors)
        List<String> list = new ArrayList<>(); 
        list.add("Java 7 Diamond");

        // 2. String in switch
        String day = "FRIDAY";
        switch (day) {
            case "FRIDAY":
                System.out.println("It's FRIDAY (Strings work in switch!)");
                break;
        }

        // 3. Try-with-resources (Auto-closable)
        try (BufferedReader br = new BufferedReader(new StringReader("Safe Auto Closing!"))) {
            System.out.println("File read: " + br.readLine());
        } catch (IOException e) { // Could be Multi-catch: IOException | SQLException
            e.printStackTrace();
        }
        
        System.out.println("Java 7 also introduced: NIO.2, Fork/Join Framework, and Numeric Literals with Underscores (e.g., 1_000_000).");
    }
}
