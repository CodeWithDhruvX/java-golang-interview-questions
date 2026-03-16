import java.util.*;
import java.util.stream.*;

public class Java9_Features {
    public static void main(String[] args) {
        // Java 9: Factory methods for Collections, Stream enhancements (takeWhile/dropWhile)
        
        // 1. Immutable Collection Factory Methods
        List<String> immutableList = List.of("Java", "9", "Modules");
        System.out.println("Immutable List: " + immutableList);

        // 2. Stream takeWhile
        List<Integer> numbers = List.of(1, 2, 3, 4, 5, 2, 1);
        List<Integer> head = numbers.stream()
            .takeWhile(n -> n < 4)
            .collect(Collectors.toList());
        System.out.println("Stream takeWhile (< 4): " + head);

        System.out.println("Major Feature: Project Jigsaw (Module System), JShell (REPL), Private methods in interfaces.");
    }
}
