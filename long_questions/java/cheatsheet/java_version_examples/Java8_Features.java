import java.util.*;
import java.util.stream.*;
import java.time.*;

public class Java8_Features {
    public static void main(String[] args) {
        // Java 8: Lambdas, Streams, Optional, New Date Time API, Default methods
        
        List<String> names = Arrays.asList("Java", "8", "Lambdas", "Streams");

        // 1. Lambda Expressions & Stream API
        System.out.println("Sorted & Filtered via Streams:");
        names.stream()
            .filter(n -> n.length() > 2)
            .sorted()
            .forEach(System.out::println);

        // 2. Optional Class (Null safety)
        Optional<String> checkNull = Optional.ofNullable(null);
        System.out.println("Optional result: " + checkNull.orElse("Default Value"));

        // 3. New Date & Time API (java.time)
        LocalDate today = LocalDate.now();
        System.out.println("Current Date (Java 8 API): " + today);

        // 4. Method References (System.out::println used above)
        System.out.println("Java 8 features changed Java from strictly Imperative to Functional hybrid!");
    }
}
